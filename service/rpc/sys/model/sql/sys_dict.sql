CREATE TABLE `sys_dict`
(
    `id`          bigint(20)                                                    NOT NULL AUTO_INCREMENT COMMENT '编号',
    `value`       varchar(100) NOT NULL DEFAULT '' COMMENT '数据值',
    `label`       varchar(100) NOT NULL DEFAULT '' COMMENT '标签名',
    `tp`          varchar(100) NOT NULL DEFAULT '' COMMENT '类型',
    `description` varchar(100) NOT NULL DEFAULT '' COMMENT '描述',
    `sort`        int(10)                                                       NOT NULL DEFAULT '0' COMMENT '排序（升序）',
    `remarks`     varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
    `created_by`  varchar(50)  NOT NULL DEFAULT '' COMMENT '创建人',
    `updated_by`  varchar(50)  NOT NULL DEFAULT '' COMMENT '更新人',
    `created_at`  timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`  timestamp   NULL     DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='字典表';