CREATE TABLE IF NOT EXISTS `users` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(150) NOT NULL,
    `username` varchar(150) NOT NULL,
    `password` varchar(150) NOT NULL,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
);