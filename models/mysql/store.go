package mysql

import (
	"fmt"
	"sync"

	"github.com/block-listener/conf"
	"github.com/block-listener/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var gdb *gorm.DB
var store models.Store
var storeOnce sync.Once

// var dbPass = utils.GetDBPass()
var dbPass = "password"

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func SharedStore() models.Store {
	storeOnce.Do(func() {
		err := initDb()
		if err != nil {
			panic(err)
		}
		store = NewStore(gdb)
	})
	return store
}

func initDb() error {
	cfg := conf.GetConfig()

	url := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		cfg.DataSource.User, dbPass, cfg.DataSource.Addr, cfg.DataSource.Database)
	var err error
	gdb, err = gorm.Open(cfg.DataSource.DriverName, url)
	if err != nil {
		return err
	}

	gdb.SingularTable(true)
	gdb.DB().SetMaxIdleConns(10)
	gdb.DB().SetMaxOpenConns(50)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "g_" + defaultTableName
	}

	if cfg.DataSource.EnableAutoMigrate {
		var tables = []interface{}{

			&models.Txhash{},
		}
		for _, table := range tables {
			if err = gdb.AutoMigrate(table).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Store) BeginTx() (models.Store, error) {
	db := s.db.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	return NewStore(db), nil
}

func (s *Store) Rollback() error {
	return s.db.Rollback().Error
}

func (s *Store) CommitTx() error {
	return s.db.Commit().Error
}
