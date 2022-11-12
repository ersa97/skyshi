CREATE TABLE IF NOT EXISTS `activities` (
  `activity_group_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `email` varchar(100) NOT NULL UNIQUE,
  `title` TEXT NOT NULL,
  PRIMARY KEY (`activity_group_id`),
  KEY `idx_activity_deleted_at` (`deleted_at`)
);