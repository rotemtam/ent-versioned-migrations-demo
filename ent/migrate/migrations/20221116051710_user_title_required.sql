-- modify "users" table
UPDATE `users` SET `title` = "" WHERE `title` IS NULL;

ALTER TABLE `users` MODIFY COLUMN `title` varchar(255) NOT NULL;
