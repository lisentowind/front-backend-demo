# Front-Backend Demo

一个前后端分离的全栈示例项目，实现基础的用户管理功能。

## 技术栈

### 前端
- Vue 3 + TypeScript
- Vite 构建工具
- Ant Design Vue UI 组件库
- Pinia 状态管理
- Axios 网络请求
- UnoCSS 样式方案

### 后端
- Go 1.25
- Gin Web 框架
- GORM ORM
- MySQL 数据库
- JWT 认证 (golang-jwt/jwt/v5)
- Bcrypt 密码加密

## 项目结构

```
front-backend-demo/
├── front/          # 前端项目
│   ├── src/        # 源代码
│   ├── public/     # 静态资源
│   └── config/     # 配置文件
├── backend/        # 后端项目
│   ├── cmd/        # 命令入口
│   ├── internal/   # 内部模块
│   ├── config/     # 配置文件
│   └── utils/      # 工具函数
└── build.sh        # 构建脚本
```

## 功能特性

- ✅ JWT 用户认证与授权
- ✅ 用户注册/登录
- ✅ 用户列表展示
- ✅ 用户分页查询
- ✅ 用户删除操作
- ✅ 前后端跨域支持
- ✅ 密码 bcrypt 加密存储

## 快速开始

### 环境要求

- Node.js >= 18
- pnpm >= 10
- Go >= 1.25
- MySQL >= 5.7

### 安装依赖

```bash
# 安装根目录依赖
pnpm install

# 安装前端依赖
cd front && pnpm install

# 安装后端依赖
cd backend && go mod download
```

### 开发模式

#### 方式一：全栈启动（推荐）

**macOS / Linux:**
```bash
pnpm dev:mac
```

**Windows:**
```bash
pnpm dev:windows
```

这会同时启动前端和后端服务。

#### 方式二：仅启动后端

**macOS / Linux:**
```bash
# 直接启动后端
pnpm backend-dev:unix

# 或先进行环境检测和配置
pnpm backend-dev:unix-setup
```

**Windows:**
```bash
pnpm backend-dev:windows
```

#### 方式三：分别启动

```bash
# 启动前端（所有平台）
pnpm front-dev

# 启动后端（根据你的平台选择）
pnpm backend-dev:unix    # macOS / Linux
# 或
pnpm backend-dev:windows # Windows
```

### 平台特定配置

如果你的后端在 macOS/Linux 上无法启动，请运行以下命令进行平台检测和配置：

```bash
pnpm backend-dev:unix-setup
```

这个脚本会：
1. 检测你的操作系统平台
2. 自动选择合适的 Air 配置文件
3. 验证 Air 二进制文件是否存在
4. 设置正确的二进制文件路径

> 注意：Windows 用户直接使用 `pnpm backend-dev:windows` 即可，无需额外配置。

### 构建

```bash
# 执行构建脚本（仅适用于 Unix-like 系统）
./build.sh
```

## 常见问题

### 后端在 macOS/Linux 上无法启动

**问题**: 后端在 Windows 上可以正常启动，但在 macOS/Linux 上不行。

**原因**: Air 配置文件使用了 Windows 特有的路径分隔符和文件扩展名。

**解决方案**:
1. 运行 `pnpm backend-dev:unix-setup` 自动配置
2. 或确保使用正确的平台特定命令（`pnpm backend-dev:unix`）

### 数据库连接失败

确保 MySQL 服务正在运行，并且配置文件中的数据库连接信息正确：
- 数据库地址: `127.0.0.1:3306`
- 数据库名: `zxd`
- 用户名: `root`
- 密码: `zxd123`

### JWT 认证相关

**401 Unauthorized 错误**

如果访问接口时返回 401 错误，说明需要登录认证：

1. 先使用 `/api/v1/auth/login` 接口登录获取 Token
2. 在请求头中添加：`Authorization: Bearer {token}`
3. Token 有效期为 2 小时，过期后需要重新登录

**测试账号**

默认测试账号（密码：`123456`）：
- 用户名：`admin`
- 用户名：`test`

详细 JWT 认证文档请查看 [JWT_AUTH.md](./JWT_AUTH.md)

## 作者

tingfeng

## License

MIT
