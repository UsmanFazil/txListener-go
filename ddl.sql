CREATE TABLE `g_txhash` (
  `txhash`      varchar(255) NOT NULL,
  `blocknum`    bigint(20) NOT NULL,
  `contractadd` varchar(255) NOT NULL,
  `chainid`     bigint(20) NOT NULL,
  PRIMARY KEY (`txhash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `g_blocksyncinfo` (
  `blocksyncnum`    bigint(20) NOT NULL,
  `syncstatus`      TINYINT,
  `backupsyncnum`   bigint(20) NOT NULL,
  `chainId`     bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `g_txmintinfo` (
  `txhash`    varchar(255) NOT NULL,
  `address`   varchar(255) NOT NULL,
  `amount`    varchar(255)  NOT NULL,
  `burnid`    varchar(255)  NOT NULL,
  `originchainid`   bigint(15)    NOT NULL,
  `tochainid`   bigint(15)    NOT NULL,    
  `status`    varchar(15) NOT NULL,
  PRIMARY KEY (`txhash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `g_txburninfo` (
  `txhash`    varchar(255) NOT NULL,
  `address`   varchar(255) NOT NULL,
  `amount`    varchar(255)  NOT NULL,
  `signature` varchar(255)  DEFAULT NULL,
  `originchainid`   bigint(15)    NOT NULL,
  `tochainid`   bigint(15)    NOT NULL,   
  `status`    varchar(15) NOT NULL,
  `burnid`    varchar(255) NOT NULL,
  PRIMARY KEY (`txhash`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;