CREATE TABLE IF NOT EXISTS `activity` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `email` TEXT NOT NULL UNIQUE,
  `title` TEXT NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_activity_deleted_at` (`deleted_at`)
);