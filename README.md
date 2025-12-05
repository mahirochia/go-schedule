# go-schedule

[![CI](https://github.com/your-username/go-schedule/actions/workflows/ci.yml/badge.svg)](https://github.com/your-username/go-schedule/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

Golang 实现的日程管理系统，后端使用 Gin + GORM，前端使用 Vue3 + Element Plus。

## ✨ 功能特性

- 📅 日程管理（增删改查）
- 📰 新闻采集与展示
- 🔄 定时任务支持
- 🐳 Docker 容器化部署

## 🚀 快速开始

### 环境要求

- Go 1.24+
- Node.js 22+
- MySQL 8.0+
- Docker (可选)

### 本地开发

```bash
# 克隆项目
git clone https://github.com/your-username/go-schedule.git
cd go-schedule

# 安装后端依赖
go mod download

# 启动后端服务
go run main.go

# 前端开发
cd frontend
npm install
npm run dev
```

### Docker 部署

```bash
# 使用 docker-compose 一键部署
docker-compose up -d
```

## 📖 API 文档

项目使用 [Swagger](https://swagger.io/) 自动生成 API 文档。

### 在线文档

文档自动部署到 GitHub Pages：**https://your-username.github.io/go-schedule/**

### 本地生成文档

```bash
# 安装 swag 工具
go install github.com/swaggo/swag/cmd/swag@latest

# 生成 Swagger 文档
swag init -g main.go -o docs/swagger

# 文档将生成在 docs/swagger 目录下
```

### API 接口概览

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | `/schedule/query` | 查询指定日期的日程 |
| POST | `/schedule/queryMonth` | 查询整月的日程 |
| POST | `/schedule/store` | 创建新日程 |
| POST | `/schedule/update` | 更新日程 |
| GET | `/news/start` | 启动新闻采集 |
| POST | `/news/query` | 查询新闻列表 |

## 🏗️ 项目结构

```
go-schedule/
├── config/          # 配置文件
├── controller/      # 控制器层
├── dao/             # 数据访问层
├── model/           # 数据模型
├── plugin/          # 插件（数据库、定时任务、爬虫等）
├── router/          # 路由配置
├── frontend/        # Vue3 前端项目
├── mysql/           # MySQL 初始化脚本
├── nginx/           # Nginx 配置
├── docs/            # 文档目录
├── Dockerfile       # Docker 构建文件
├── docker-compose.yml
└── main.go          # 入口文件
```

## 🔧 GitHub Actions

项目配置了完整的 CI/CD 工作流：

- **CI** (`ci.yml`) - 代码构建、测试、Docker 镜像构建
- **Release** (`release.yml`) - 版本发布、镜像推送到 GHCR
- **Documentation** (`docs.yml`) - 自动生成并部署 API 文档

## 📝 开发指南

### 添加新的 API 接口

1. 在 `controller/` 中添加控制器函数
2. 添加 Swagger 注释：

```go
// CreateItem 创建新项目
// @Summary      创建项目
// @Description  创建一个新项目
// @Tags         项目管理
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreateReq  true  "创建参数"
// @Success      200      {object}  system.Response
// @Router       /item/create [post]
func CreateItem(c *gin.Context) {
    // ...
}
```

3. 在 `router/router.go` 中注册路由
4. 提交代码后，文档会自动更新

## 📄 License

[MIT License](LICENSE)
