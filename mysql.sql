-- CREATED_AT = 17-01-2021 12:12:01
-- create table `users` with id using uuid
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` char(36) NOT NULL,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(60) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- CREATED_AT = 17-01-2021 12:12:02
-- create table `tasks` with id using uuid
DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks` (
  `id` char(36) NOT NULL,
  `user_id` varchar(36) DEFAULT NULL,
  `title` varchar(100) NOT NULL,
  `description` varchar(1000) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tasks_deleted_at` (`deleted_at`),
  KEY `fk_user_tasks` (`user_id`),
  CONSTRAINT `fk_user_tasks` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- CREATED_AT = 07-02-2021 10:38:00
-- create table `blogs` with id using uuid
DROP TABLE IF EXISTS `blogs`;
CREATE TABLE `blogs` (
  `id` char(36) NOT NULL,
  `title` varchar(100) NOT NULL,
  `body` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;