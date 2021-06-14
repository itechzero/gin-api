-- user表
CREATE TABLE `sys_users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(191) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户UUID',
  `username` varchar(191) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录名',
  `password` varchar(191) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录密码',
  `nick_name` varchar(191) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `header_img` varchar(191) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `authority_id` varchar(90) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '888' COMMENT '用户角色ID',
  `created_at` datetime DEFAULT NOW(),
  `updated_at` datetime DEFAULT NOW(),
  `deleted_at` datetime DEFAULT NOW(),
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci