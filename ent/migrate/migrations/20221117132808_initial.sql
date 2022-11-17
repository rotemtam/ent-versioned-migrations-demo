-- create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `email` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `email` (`email`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "blogs" table
CREATE TABLE `blogs` (`id` bigint NOT NULL AUTO_INCREMENT, `title` varchar(255) NOT NULL, `body` longtext NOT NULL, `created_at` timestamp NOT NULL, `user_blog_posts` bigint NULL, PRIMARY KEY (`id`), CONSTRAINT `blogs_users_blog_posts` FOREIGN KEY (`user_blog_posts`) REFERENCES `users` (`id`) ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
