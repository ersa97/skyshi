CREATE TABLE IF NOT EXISTS `todo` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `activity_group_id` bigint(20) DEFAULT NULL,
  `title` TEXT DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `priority` TEXT DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_todo_deleted_at` (`deleted_at`)
);