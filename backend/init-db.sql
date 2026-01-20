-- 初始化数据库和测试用户
-- 使用方法: mysql -u root -p < init-db.sql

-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS zxd CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 切换到数据库
USE zxd;

-- 创建用户表（如果不存在）
CREATE TABLE IF NOT EXISTS user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    age INT,
    email VARCHAR(100),
    password VARCHAR(255) NOT NULL,
    create_time VARCHAR(50)
);

-- 插入测试用户（密码: 123456，使用 bcrypt 加密）
-- admin 用户
INSERT INTO user (name, age, email, password, create_time)
VALUES (
    'admin',
    25,
    'admin@example.com',
    '$2a$10$K8wjZ5sY2J3tL4q6X8z9QeO5r6n7m8p9q0r1s2t3u4v5w6x7y8z9A0B',
    DATE_FORMAT(NOW(), '%Y-%m-%d %H:%i:%s')
)
ON DUPLICATE KEY UPDATE
    age = VALUES(age),
    email = VALUES(email),
    password = VALUES(password),
    create_time = VALUES(create_time);

-- test 用户
INSERT INTO user (name, age, email, password, create_time)
VALUES (
    'test',
    22,
    'test@example.com',
    '$2a$10$K8wjZ5sY2J3tL4q6X8z9QeO5r6n7m8p9q0r1s2t3u4v5w6x7y8z9A0B',
    DATE_FORMAT(NOW(), '%Y-%m-%d %H:%i:%s')
)
ON DUPLICATE KEY UPDATE
    age = VALUES(age),
    email = VALUES(email),
    password = VALUES(password),
    create_time = VALUES(create_time);

-- 插入一些测试用户数据
INSERT INTO user (name, age, email, password, create_time)
VALUES
    ('张三', 28, 'zhangsan@example.com', '$2a$10$K8wjZ5sY2J3tL4q6X8z9QeO5r6n7m8p9q0r1s2t3u4v5w6x7y8z9A0B', DATE_FORMAT(NOW(), '%Y-%m-%d %H:%i:%s')),
    ('李四', 32, 'lisi@example.com', '$2a$10$K8wjZ5sY2J3tL4q6X8z9QeO5r6n7m8p9q0r1s2t3u4v5w6x7y8z9A0B', DATE_FORMAT(NOW(), '%Y-%m-%d %H:%i:%s')),
    ('王五', 25, 'wangwu@example.com', '$2a$10$K8wjZ5sY2J3tL4q6X8z9QeO5r6n7m8p9q0r1s2t3u4v5w6x7y8z9A0B', DATE_FORMAT(NOW(), '%Y-%m-%d %H:%i:%s'))
ON DUPLICATE KEY UPDATE
    age = VALUES(age),
    email = VALUES(email),
    create_time = VALUES(create_time);

SELECT '数据库初始化完成！' AS message;
SELECT '测试账号:' AS message;
SELECT '用户名: admin, 密码: 123456' AS message;
SELECT '用户名: test, 密码: 123456' AS message;
