CREATE TABLE `g_txhash` (
  `txhash`      varchar(255) NOT NULL,
  `blocknum`    bigint(20) NOT NULL,
  `contractadd` varchar(255) NOT NULL,
  PRIMARY KEY (`txhash`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `g_blocksyncinfo` (
  `blocksyncnum`    bigint(20) NOT NULL,
  `syncstatus`    TINYINT,
  `backupsyncnum`    bigint(20) NOT NULL

) ENGINE=InnoDB DEFAULT CHARSET=utf8;
