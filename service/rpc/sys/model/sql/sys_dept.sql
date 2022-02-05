CREATE TABLE `sys_dept`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
    `name`       varchar(50) NOT NULL DEFAULT '' COMMENT '机构名称',
    `parent_id`  bigint(20) NOT NULL DEFAULT '0' COMMENT '上级机构ID，一级机构为0',
    `order_num`  int(11) NOT NULL DEFAULT '0' COMMENT '排序',
    `created_by` varchar(50)  NOT NULL DEFAULT '' COMMENT '创建人',
    `updated_by` varchar(50)  NOT NULL DEFAULT '' COMMENT '更新人',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='机构管理';