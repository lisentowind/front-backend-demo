-- 修改 password 字段为可空
-- 使用方法: mysql -u root -p zxd < migrate-password-nullable.sql

USE zxd;

-- 修改 password 字段允许为 NULL 并设置默认值为空字符串
ALTER TABLE user MODIFY COLUMN password VARCHAR(255) DEFAULT '';

SELECT '数据库迁移完成！password 字段已修改为可空' AS message;
