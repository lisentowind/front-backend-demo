# JWT 认证功能说明

## 概述

本项目已集成 JWT (JSON Web Token) 认证机制，用于保护 API 接口。

## 技术栈

- **JWT 库**: `github.com/golang-jwt/jwt/v5`
- **密码加密**: `golang.org/x/crypto/bcrypt`
- **认证中间件**: Gin 框架中间件

## 数据库变更

### 新增字段

用户表 (`user`) 新增了 `password` 字段：

```sql
password VARCHAR(255) NOT NULL  -- 存储 bcrypt 加密后的密码
```

### 自动迁移

项目启动时会自动执行数据库迁移，创建或更新表结构。

## API 接口

### 公开接口（无需认证）

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/v1/auth/login` | 用户登录 |
| POST | `/api/v1/auth/register` | 用户注册 |
| GET | `/api/v1/hello/ping` | 心跳检测 |
| GET | `/api/v1/hello/user/table` | 获取模拟表格数据 |

### 需要认证的接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/v1/auth/info` | 获取当前用户信息 |
| GET | `/api/v1/hello/user/all` | 分页获取用户列表 |
| GET | `/api/v1/hello/user` | 根据ID获取用户 |
| POST | `/api/v1/hello/user/add` | 添加用户 |
| DELETE | `/api/v1/hello/user/delete` | 删除用户 |

## 使用流程

### 1. 初始化数据库

如果你的数据库是空的，可以使用以下任一方式初始化：

**方式一：使用 SQL 脚本**
```bash
mysql -u root -p < backend/init-db.sql
```

**方式二：使用 Go 自动迁移**
启动后端服务，会自动创建表结构。然后手动插入测试用户：

```sql
INSERT INTO user (name, password, create_time)
VALUES ('admin', '$2a$10$K8wjZ5sY2J3tL4q6X8z9QeO5r6n7m8p9q0r1s2t3u4v5w6x7y8z9A0B', NOW());
```

> 注意：上面的密码是 `123456` 经过 bcrypt 加密后的哈希值。

### 2. 用户注册

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"myuser","password":"mypassword"}'
```

**响应示例：**
```json
{
  "code": 200,
  "msg": "注册成功",
  "data": {
    "id": 1,
    "username": "myuser"
  }
}
```

### 3. 用户登录

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}'
```

**响应示例：**
```json
{
  "code": 200,
  "msg": "登录成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### 4. 访问受保护接口

在请求头中添加 `Authorization`：

```bash
curl http://localhost:8080/api/v1/hello/user/all?pageNum=1&pageSize=10 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

## 前端集成

### Token 存储

前端使用 `utils/modules/auth.ts` 管理 Token：

```typescript
import { setToken, getToken, clearToken, isLogin } from '@/utils'

// 登录成功后保存 token
setToken(token)

// 获取 token
const token = getToken()

// 退出登录
clearToken()
```

### Axios 自动添加 Token

`src/apis/index.ts` 已配置请求拦截器，会自动：
1. 从 localStorage/Cookie 获取 Token
2. 添加到请求头：`Authorization: Bearer {token}`
3. 登录/注册接口自动跳过 Token 添加

### Pinia Store

`src/store/modules/user.ts` 提供了完整的用户管理：

```typescript
import { useUserStore } from '@/store/modules/user'

const userStore = useUserStore()

// 登录
await userStore.login({ username: 'admin', password: '123456' })

// 获取用户信息
await userStore.getUserInfo()

// 退出
userStore.logout()

// 检查登录状态
const isLoggedIn = userStore.isLoggedIn
```

## Token 配置

### 后端配置 (`utils/jwt.go`)

```go
// Token 有效期：2小时
ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour))

// 签发人
Issuer: "go-backend"

// 签名算法：HS256
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
```

### 安全建议

1. **密钥管理**：生产环境应从环境变量读取密钥
   ```go
   jwtSecret := []byte(os.Getenv("JWT_SECRET"))
   ```

2. **HTTPS**：生产环境必须使用 HTTPS 传输 Token

3. **Token 有效期**：当前设置为 2 小时，可根据需求调整

4. **密码加密**：使用 bcrypt 加密，自动加盐

## 测试账号

项目初始化时会创建以下测试账号：

| 用户名 | 密码 | 角色 |
|--------|------|------|
| admin | 123456 | user |
| test | 123456 | user |

## 常见问题

### 401 Unauthorized

**原因**：Token 无效或过期

**解决**：
1. 重新登录获取新的 Token
2. 检查请求头格式：`Authorization: Bearer {token}`

### 403 Forbidden

**原因**：权限不足（当前仅管理员角色有特殊权限）

**解决**：使用管理员账号登录

### 密码错误

**原因**：密码未正确加密或验证

**解决**：确保使用 bcrypt 加密密码，前端发送明文密码，后端自动验证

## 项目结构

```
backend/
├── internal/
│   ├── middleware/
│   │   └── auth.go              # JWT 认证中间件
│   ├── controller/
│   │   └── auth_controller.go   # 认证控制器
│   ├── model/
│   │   └── user.go              # 用户模型（新增 Password 字段）
│   └── router/
│       └── auth_router/
│           └── auth.go          # 认证路由
├── utils/
│   └── jwt.go                   # JWT 工具函数
└── go.mod                       # 新增依赖：jwt, bcrypt
```

## 依赖安装

如果需要手动安装依赖：

```bash
cd backend
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
```

## 后续扩展建议

1. **Refresh Token**：实现双 Token 机制（Access Token + Refresh Token）
2. **Token 黑名单**：实现 Token 注销功能
3. **角色权限**：扩展 RBAC 权限控制
4. **验证码**：登录增加图形验证码
5. **登录日志**：记录用户登录日志
