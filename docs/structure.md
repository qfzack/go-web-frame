# Project Structure

Golang项目没有规范统一的分层架构，但是在实际开发中，社区中形成了一种广泛采用的架构分层规范：

```bash
project-root/
├── cmd/                          # 应用程序入口点
│   ├── user-service/
│   │   └── main.go
│   ├── order-service/
│   │   └── main.go
│   └── payment-service/
│       └── main.go
├── internal/                     # 私有代码，不对外暴露
│   ├── user/
│   │   ├── handler/              # HTTP handlers
│   │   ├── service/              # 业务逻辑层
│   │   ├── repository/           # 数据访问层
│   │   └── model/                # 数据模型
│   ├── order/
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   └── model/
│   └── payment/
│       ├── handler/
│       ├── service/
│       ├── repository/
│       └── model/
├── pkg/                          # 可复用的公共代码
│   ├── database/                 # 数据库连接和配置
│   │   ├── mysql/
│   │   ├── postgres/
│   │   └── migration/
│   ├── cache/                    # 缓存相关
│   │   ├── redis/
│   │   └── memory/
│   ├── mq/                       # 消息队列
│   │   ├── rabbitmq/
│   │   ├── kafka/
│   │   └── nsq/
│   ├── grpc/                     # gRPC相关工具
│   │   ├── client/
│   │   ├── server/
│   │   └── middleware/
│   ├── http/                     # HTTP相关工具
│   │   ├── middleware/
│   │   ├── response/
│   │   └── validator/
│   ├── logger/                   # 日志工具
│   ├── config/                   # 配置管理
│   └── utils/                    # 通用工具函数
├── api/                          # API定义
│   ├── proto/                    # Protocol Buffers定义
│   │   ├── user/
│   │   │   ├── user.proto
│   │   │   └── user.pb.go        # *_grpc.pb.go一般同目录
│   │   ├── order/
│   │   │   ├── order.proto
│   │   │   └── order.pb.go
│   │   └── payment/
│   │       ├── payment.proto
│   │       └── payment.pb.go
│   └── openapi/                  # OpenAPI/Swagger规范
│       ├── user.yaml
│       ├── order.yaml
│       └── payment.yaml
├── configs/                      # 配置文件
│   ├── local.yaml
│   ├── dev.yaml
│   ├── staging.yaml
│   └── prod.yaml
├── scripts/                      # 脚本文件
│   ├── build.sh
│   ├── deploy.sh
│   └── migrate.sh
├── deployments/                  # 部署相关
│   ├── docker/
│   │   ├── user-service/
│   │   │   └── Dockerfile
│   │   ├── order-service/
│   │   │   └── Dockerfile
│   │   └── payment-service/
│   │       └── Dockerfile
│   ├── k8s/                      # Kubernetes配置
│   └── docker-compose.yaml   
├── docs/                         # 文档
├── test/                         # 测试相关
│   ├── integration/   
│   └── e2e/
├── tools/                        # 工具和代码生成
│   └── protoc/
├── go.mod
├── go.sum
├── Makefile
└── README.md
```
