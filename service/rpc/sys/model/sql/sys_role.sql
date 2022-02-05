CREATE TABLE `sys_role`
(
    `id`         bigint(20)  NOT NULL AUTO_INCREMENT COMMENT '编号',
    `name`       varchar(100) NOT NULL DEFAULT '' COMMENT '角色名称',
    `remark`     varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
    `status`     tinyint(3)   NOT NULL DEFAULT '1' COMMENT '状态  1:启用,0:禁用',
    `created_by` varchar(50)  NOT NULL DEFAULT '' COMMENT '创建人',
    `updated_by` varchar(50)  NOT NULL DEFAULT '' COMMENT '更新人',
    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp   NULL     DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='角色管理';