CREATE TABLE `carts` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`uid`  int(10) unsigned NOT NULL DEFAULT '0'  COMMENT 'uid',
    `gid` int(10) unsigned NOT NULL DEFAULT '0'  COMMENT '商品id',
    `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '产品名称',
	`num` int(10) unsigned NOT NULL DEFAULT '0'  COMMENT '数量',
	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
