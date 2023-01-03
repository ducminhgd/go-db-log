DROP TABLE IF EXISTS `event_log`;
CREATE TABLE `event_log` (
    `id` VARCHAR(36) NOT NULL,
    `event_type` VARCHAR(100) NOT NULL,
    `object_type` VARCHAR(100) NOT NULL,
    `object_id` VARCHAR(100) NOT NULL,
    `actor_type` VARCHAR(100) NOT NULL,
    `actor_id` VARCHAR(100) NOT NULL,
    `data` TEXT NULL,
    `result` TINYINT UNSIGNED NOT NULL DEFAULT 1,
    `timestamps` TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    PRIMARY KEY (`id`)
);