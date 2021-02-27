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

-- add one data to login // password=password
INSERT INTO users(id, name, email, password)
VALUES ('190e9c30-c119-40a2-a554-85b6fb8cffaf', 'admin', 'admin@admin.com',
        '$2y$10$M.VrmGnEdrJyYyYTGmkynexSg9MjHDhPckyyVJvxoZGRAOMKDL3fq ');
