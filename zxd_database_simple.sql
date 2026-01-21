-- =============================================
-- zxd Database - Simplified Initialization Script
-- =============================================

-- Create database
CREATE DATABASE IF NOT EXISTS `zxd`
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

USE `zxd`;

-- Create user table
CREATE TABLE `user` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `age` INT,
    `email` VARCHAR(100),
    `password` VARCHAR(255) DEFAULT '',
    `role` VARCHAR(20) DEFAULT 'admin',
    `project` VARCHAR(100) NOT NULL,
    `create_time` VARCHAR(50),
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_name` (`name`),
    KEY `idx_role` (`role`),
    KEY `idx_project` (`project`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert initial users (passwords will be set via API)
INSERT INTO `user` (`name`, `age`, `email`, `password`, `role`, `project`, `create_time`) VALUES
('admin', 25, 'admin@zxd.com', '', 'admin', 'zxd-project', '2026-01-21 10:00:00'),
('editor1', 30, 'editor1@zxd.com', '', 'editor', 'zxd-project', '2026-01-21 10:00:00'),
('guest1', 28, 'guest1@zxd.com', '', 'guest', 'zxd-project', '2026-01-21 10:00:00');

-- Verify setup
SHOW TABLES;
DESCRIBE `user`;
SELECT * FROM `user`;