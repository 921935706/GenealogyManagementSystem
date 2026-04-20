# 族谱管理系统 - 数据模型与 API 设计文档

## 一、技术栈说明
- 后端：GoLang
- 数据库：SQLite
- API 风格：RESTful
- swaggo/swag + gin-swagger 用于生成 API 文档和 UI
---

## 项目结构（DDD 分层架构） 

GenealogyManagementSystem/
├── cmd/
│   └── server/
│       └── main.go                 # 应用入口
├── internal/
│   ├── domain/                     # 领域层（核心业务逻辑）
│   │   ├── person/
│   │   │   ├── entity.go          # 实体定义
│   │   │   ├── repository.go      # 仓储接口
│   │   │   └── service.go         # 领域服务
│   │   ├── relationship/
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   └── pedigree/
│   │       ├── entity.go
│   │       ├── repository.go
│   │       └── service.go
│   ├── application/               # 应用层（用例编排）
│   │   ├── person/
│   │   │   └── person_app.go      # 人员应用服务
│   │   └── relationship/
│   │       └── relationship_app.go
│   ├── infrastructure/            # 基础设施层
│   │   ├── persistence/
│   │   │   ├── sqlite/
│   │   │   │   ├── person_repo.go    # SQLite 仓储实现 
│   │   │   │   ├── relationship_repo.go
│   │   │   │   └── db.go             # 数据库连接
│   │   │   └── redis/                 # Redis 缓存
│   │   ├── web/
│   │   │   ├── handler/           # HTTP 处理器
│   │   │   │   ├── person_handler.go
│   │   │   │   └── relationship_handler.go
│   │   │   ├── middleware/        # 中间件
│   │   │   └── router.go          # 路由配置
│   │   └── excel/                 # Excel 导入导出
│   └── pkg/                       # 公共包
│       ├── errors/                # 自定义错误
│       ├── validator/             # 参数校验
│       └── response/              # 统一响应
├── go.mod
└── go.sum

--- 


