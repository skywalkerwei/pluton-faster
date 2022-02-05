CREATE TABLE `sys_user_role`
(
    `id`         bigint(20)  NOT NULL AUTO_INCREMENT COMMENT '编号',
    `user_id`    bigint(20)  NOT NULL DEFAULT '0' COMMENT '用户ID',
    `role_id`    bigint(20)  NOT NULL DEFAULT '0' COMMENT '角色ID',
    `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '创建人',
    `updated_by` varchar(50) NOT NULL DEFAULT '' COMMENT '更新人',
    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp   NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='用户角色';