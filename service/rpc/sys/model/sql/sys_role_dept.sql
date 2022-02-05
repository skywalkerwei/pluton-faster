CREATE TABLE `sys_role_dept`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
    `role_id`    bigint(20) NOT NULL DEFAULT '0' COMMENT '角色ID',
    `dept_id`    bigint(20) NOT NULL DEFAULT '0' COMMENT '机构ID',
    `created_by` varchar(50)  NOT NULL DEFAULT '' COMMENT '创建人',
    `updated_by` varchar(50)  NOT NULL DEFAULT '' COMMENT '更新人',
    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色机构';