CREATE TABLE `g_txhash` (
  `txhash`      varchar(255) NOT NULL,
  `blocknum`    bigint(20) NOT NULL,
  `contractadd` varchar(255) NOT NULL,
  PRIMARY KEY (`txhash`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `g_lastconfirmed` (
  `blocknum`    bigint(20) NOT NULL,
  PRIMARY KEY (`blocknum`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;
