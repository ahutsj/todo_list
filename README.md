# todo_list
todo_list based on Golang,mainly uses grpc,gin,mysql,etcd...

分为api-gateway和user模块
gin实现api-gateway的router，etcd做服务发现，grpc实现api-gateway和user层的通信调用，user层做handler和dao逻辑，mysql做持久化使用gorm操作，其他的（jwt/cors等）放api-gateway

## 目录

```
api-gateway/
├── cmd                   // 启动入口
├── config                // 配置文件
├── discovery             // etcd服务注册、keep-alive、获取服务信息等等
├── internal              // 业务逻辑（不对外暴露）
│   ├── handler           // 视图层
│   └── service           // 服务层
│       └──pb             // 放置生成的pb文件
├── logs                  // 放置打印日志模块
├── middleware            // 中间件
├── pkg                   // 各种包
│   ├── e                 // 统一错误状态码
│   ├── res               // 统一response接口返回
│   └── util              // 各种工具、JWT、Logger等等..
├── routes                // http路由模块
└── wrappers              // 各个服务之间的熔断降级
```

```
user/
├── cmd                   // 启动入口
├── config                // 配置文件
├── discovery             // etcd服务注册、keep-alive、获取服务信息等等
├── internal              // 业务逻辑（不对外暴露）
│   ├── handler           // 视图层
│   ├── cache             // 缓存模块
│   ├── repository        // 持久层
│   └── service           // 服务层
│       └──pb             // 放置生成的pb文件
├── logs                  // 放置打印日志模块
└── pkg                   // 各种包
    ├── e                 // 统一错误状态码
    ├── res               // 统一response接口返回
    └── util              // 各种工具、JWT、Logger等等..
```

## 启动顺序

```bash
cd X:\Program Files (x86)\mysql-8.0.31-winx64\bin>
net start mysql

cd X:\Program Files (x86)\etcd-v3.4.22-windows-amd64
etcd.exe

// 进入工作目录
cd /user/cmd
go run main.go

cd /api-gateway/cmd
go run main.go
```