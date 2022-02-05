CREATE TABLE `sys_menu`
(
    `id`             bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
    `name`           varchar(50)  NOT NULL DEFAULT '' COMMENT '菜单名称',
    `parent_id`      bigint(20) NOT NULL DEFAULT '0' COMMENT '父菜单ID，一级菜单为0',
    `url`            varchar(200) NOT NULL DEFAULT '',
    `perms`          varchar(500) NOT NULL DEFAULT '' COMMENT '授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)',
    `tp`             int(11) NOT NULL DEFAULT '0' COMMENT '类型   0：目录   1：菜单   2：按钮',
    `icon`           varchar(50)  NOT NULL DEFAULT '' COMMENT '菜单图标',
    `order_num`      int(11) NOT NULL DEFAULT '0' COMMENT '排序',
    `vue_path`       varchar(64)  NOT NULL DEFAULT '' COMMENT 'vue系统的path',
    `vue_component`  varchar(64)  NOT NULL DEFAULT '' COMMENT 'vue的页面',
    `vue_icon`       varchar(64)  NOT NULL DEFAULT '' COMMENT 'vue的图标',
    `vue_redirect`   varchar(64)  NOT NULL DEFAULT '' COMMENT 'vue的路由重定向',
    `background_url` varchar(128) COLLATE utf8mb4_unicode_ci                       NOT NULL DEFAULT '' COMMENT '后台地址',
    `created_by`     varchar(50)  NOT NULL DEFAULT '' COMMENT '创建人',
    `update_by`      varchar(50)  NOT NULL DEFAULT '' COMMENT '更新人',
    `created_at`     timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`     timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单管理';