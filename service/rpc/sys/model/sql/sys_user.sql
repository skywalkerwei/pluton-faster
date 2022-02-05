CREATE TABLE `sys_user`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
    `name`       varchar(50)  NOT NULL DEFAULT '' COMMENT '用户名',
    `nick_name`  varchar(150)  NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar`     varchar(150)  NOT NULL DEFAULT '' COMMENT '头像',
    `password`   varchar(100)  NOT NULL DEFAULT '' COMMENT '密码',
    `salt`       varchar(40)   NOT NULL DEFAULT '' COMMENT '加密盐',
    `email`      varchar(100)  NOT NULL DEFAULT '' COMMENT '邮箱',
    `mobile`     varchar(100)  NOT NULL DEFAULT '' COMMENT '手机号',
    `dept_id`    bigint(20) NOT NULL DEFAULT '0' COMMENT '机构ID',
    `job_id`     int(11) NOT NULL DEFAULT '0' COMMENT '岗位Id',
    `status`     tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态  0：禁用   1：正常',
    `created_by` varchar(50)   NOT NULL DEFAULT '' COMMENT '创建人',
    `updated_by` varchar(50)   NOT NULL DEFAULT '' COMMENT '更新人',
    `created_at` timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户管理';