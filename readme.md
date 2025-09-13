```
gin_blog/
├── config/         # 配置结构体目录
├── core/           # 初始化操作目录
├── docs/           # API 文档目录
├── flag/           # 命令行相关操作
├── global/         # 全局变量
├── handler/        # http
├── middleware/     # 中间件目录
├── models/         # 数据库表结构目录
├── routers/        # Gin 路由配置目录
├── service/        # 业务服务逻辑目录
├── testdata/       # 测试数据/文件目录
├── utils/          # 常用工具
├── main.go         # 程序入口文件
└── setting.yaml    # 项目配置文件
```
本项目使用gin开发，为学习项目。

// 启动前端
python -m http.server 8000

// 启动后端
go run main.go

// 访问前端
http://localhost:8000/
