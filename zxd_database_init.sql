-- =============================================
-- zxd Database Initialization Script
-- Generated based on Go backend code analysis
-- Date: 2026-01-21
-- =============================================

-- Create the database if it doesn't exist
CREATE DATABASE IF NOT EXISTS `zxd`
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

-- Use the zxd database
USE `zxd`;

-- =============================================
-- User Table Structure
-- Based on: backend/internal/model/user.go
-- =============================================

-- Drop the table if it exists (for clean setup)
DROP TABLE IF EXISTS `user`;

-- Create the user table
CREATE TABLE `user` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT '用户ID，主键，自增',
    `name` VARCHAR(100) NOT NULL COMMENT '用户名，必填字段',
    `age` INT COMMENT '用户年龄',
    `email` VARCHAR(100) COMMENT '用户邮箱',
    `password` VARCHAR(255) DEFAULT '' COMMENT '用户密码（加密存储），默认为空',
    `role` VARCHAR(20) DEFAULT 'admin' COMMENT '用户角色：admin, editor, guest，默认为admin',
    `project` VARCHAR(100) NOT NULL COMMENT '所属项目，必填字段',
    `create_time` VARCHAR(50) COMMENT '创建时间',

    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_name` (`name`),
    KEY `idx_role` (`role`),
    KEY `idx_project` (`project`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- =============================================
-- Initial Data Insertion
-- =============================================

-- Insert sample admin user (password will be initialized via API)
-- Note: The password field is empty initially, it will be set via the /auth/register endpoint
INSERT INTO `user` (`name`, `age`, `email`, `password`, `role`, `project`, `create_time`) VALUES
('admin', 25, 'admin@zxd.com', '', 'admin', 'zxd-project', '2026-01-21 10:00:00'),
('editor1', 30, 'editor1@zxd.com', '', 'editor', 'zxd-project', '2026-01-21 10:00:00'),
('guest1', 28, 'guest1@zxd.com', '', 'guest', 'zxd-project', '2026-01-21 10:00:00');

-- =============================================
-- Database User Creation (Optional)
-- For security, create a dedicated database user
-- =============================================

-- Create database user (uncomment if needed)
-- CREATE USER 'zxd_user'@'localhost' IDENTIFIED BY 'zxd_password_123';
-- GRANT ALL PRIVILEGES ON `zxd`.* TO 'zxd_user'@'localhost';
-- FLUSH PRIVILEGES;

-- =============================================
-- Table Structure Verification
-- =============================================

-- Show table structure
SHOW TABLES;

-- Show user table structure
DESCRIBE `user`;

-- Show sample data
SELECT * FROM `user`;

-- =============================================
-- Additional Indexes (Optional)
-- For better query performance
-- =============================================

-- Add these indexes if you frequently query by email
-- ALTER TABLE `user` ADD UNIQUE KEY `idx_email` (`email`);

-- Add composite index for role and project queries
-- ALTER TABLE `user` ADD INDEX `idx_role_project` (`role`, `project`);

-- =============================================
-- Notes for Application Usage
-- =============================================

/*
1. Database Connection:
   - Host: 127.0.0.1
   - Port: 3306
   - Database: zxd
   - Username: root (or custom user)
   - Password: zxd123 (as per mysql.go config)

2. Initial Setup:
   - The password field is initially empty for all users
   - Use the /auth/register endpoint to initialize passwords
   - Default admin user: admin (no password set initially)

3. User Roles:
   - admin: Full access
   - editor: Limited access (can edit content)
   - guest: Read-only access

4. Project Structure:
   - All users belong to 'zxd-project' initially
   - The project field is required and can be used for multi-tenancy

5. Security Notes:
   - Passwords are hashed using bcrypt before storage
   - JWT tokens are used for authentication
   - Consider changing default passwords in production
*/

-- =============================================
-- End of Script
-- =============================================