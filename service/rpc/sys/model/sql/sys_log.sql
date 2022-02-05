CREATE TABLE `sys_log`
(
    `id`         bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '编号',
    `user_name`  varchar(50)   NOT NULL DEFAULT '' COMMENT '用户名',
    `operation`  varchar(50)   NOT NULL DEFAULT '' COMMENT '用户操作',
    `method`     varchar(200)  NOT NULL DEFAULT '' COMMENT '请求方法',
    `params`     varchar(5000) NOT NULL DEFAULT '' COMMENT '请求参数',
    `time`       bigint(20)    NOT NULL DEFAULT '0' COMMENT '执行时长(毫秒)',
    `ip`         varchar(64)   NOT NULL DEFAULT '' COMMENT 'IP地址',
    `created_by` varchar(50)   NOT NULL DEFAULT '' COMMENT '创建人',
    `updated_by` varchar(50)   NOT NULL DEFAULT '' COMMENT '更新人',
    `created_at` timestamp    NULL     DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp    NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='系统操作日志';