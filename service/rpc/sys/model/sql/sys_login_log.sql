CREATE TABLE `sys_login_log`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
    `user_name`  varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
    `status`     varchar(50) NOT NULL DEFAULT 'online' COMMENT '登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）',
    `ip`         varchar(64) NOT NULL DEFAULT '' COMMENT 'IP地址',
    `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '创建人',
    `updated_by` varchar(50) NOT NULL DEFAULT '' COMMENT '更新人',
    `created_at` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统登录日志';