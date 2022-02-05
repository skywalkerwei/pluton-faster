CREATE TABLE `sys_job`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
    `job_name`   varchar(50) NOT NULL DEFAULT '' COMMENT '职位名称',
    `order_num`  int(11) NOT NULL DEFAULT '0' COMMENT '排序',
    `remarks`    varchar(68) NOT NULL DEFAULT '' COMMENT '备注',
    `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '创建人',
    `updated_by` varchar(50) NOT NULL DEFAULT '' COMMENT '更新人',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` int(11) DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='职位管理';