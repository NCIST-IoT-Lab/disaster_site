# 灾难现场信息管理系统

这是一个使用Gin和GORM开发的RESTful API应用，用于管理灾难现场信息。

## 功能特点

- 基于Gin框架的RESTful API
- 使用GORM进行数据库操作
- MySQL数据库支持
- 标准化API响应
- 中间件支持（日志、CORS等）
- Docker支持

## 技术栈

- Go 1.13+
- Gin Web框架
- GORM ORM库
- MySQL 8.0
- Docker & Docker Compose

## 目录结构

```
.
├── cmd
│   └── api
│       └── main.go        # 应用入口
├── configs                # 配置文件
├── docs                   # 文档
├── internal
│   ├── database           # 数据库相关
│   ├── handlers           # 请求处理器
│   ├── middleware         # 中间件
│   ├── models             # 数据模型
│   └── utils              # 工具函数
├── .env.example           # 环境变量示例
├── Dockerfile             # Docker构建文件
├── docker-compose.yml     # Docker Compose配置
├── go.mod                 # Go模块文件
└── go.sum                 # Go依赖校验文件
```

## 开始使用

### 前提条件

- Go 1.13+
- MySQL 8.0
- Docker（可选）

### 本地运行

1. 克隆仓库

```bash
git clone https://your-repository-url.git
cd disaster_site_information_management_system
```

2. 设置环境变量（可选）

```bash
cp configs/.env.example .env
# 编辑.env文件设置您的配置
```

3. 创建数据库

```bash
mysql -u root -p -e "CREATE DATABASE disaster_site_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
mysql -u root -p -e "CREATE USER 'app_user'@'localhost' IDENTIFIED BY 'password'; GRANT ALL PRIVILEGES ON disaster_site_db.* TO 'app_user'@'localhost'; FLUSH PRIVILEGES;"
```

4. 运行应用

```bash
go run cmd/api/main.go
```

### 使用Docker

1. 使用Docker Compose启动应用和数据库

```bash
docker-compose up -d
```

2. 停止服务

```bash
docker-compose down
```

## API 端点

- `GET /health` - 健康检查
- `GET /api/sites` - 获取所有灾难现场
- `GET /api/sites/:id` - 获取特定灾难现场
- `POST /api/sites` - 创建新的灾难现场
- `PUT /api/sites/:id` - 更新灾难现场
- `DELETE /api/sites/:id` - 删除灾难现场

## 贡献

欢迎提交问题和功能请求。对于大型更改，请先开一个issue讨论您想要更改的内容。

## 许可证

[MIT](LICENSE)
