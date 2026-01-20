-- 数据库迁移脚本：添加 role 字段
-- 使用方法: mysql -u root -p zxd < migrate-add-role.sql

USE zxd;

-- 添加 role 字段到 user 表
ALTER TABLE user ADD COLUMN IF NOT EXISTS role VARCHAR(20) DEFAULT 'admin' AFTER password;

-- 更新所有现有用户的 role 为 admin
UPDATE user SET role = 'admin' WHERE role IS NULL OR role = '';

-- 确保 role 字段不为空
ALTER TABLE user MODIFY COLUMN role VARCHAR(20) NOT NULL DEFAULT 'admin';

SELECT '数据库迁移完成！所有用户角色已设置为 admin' AS message;
SELECT id, name, role FROM user;
