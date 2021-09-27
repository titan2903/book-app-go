CREATE TABLE IF NOT EXISTS `books` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` int UNSIGNED,
    `name` varchar(150),
    `genre` varchar(150),
    `release` int,
    `is_read` boolean,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);