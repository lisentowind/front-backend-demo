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

- 用户列表展示
- 用户分页查询
- 用户删除操作
- 前后端跨域支持

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

```bash
# 同时启动前端和后端
pnpm dev

# 或分别启动
pnpm front-dev   # 启动前端
pnpm backend-dev # 启动后端
```

### 构建

```bash
# 执行构建脚本
./build.sh
```

## 作者

tingfeng

## License

MIT
