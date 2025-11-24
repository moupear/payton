# Payton 项目说明

## 项目概述
Payton 是一个基于 Go 语言开发的 Web 服务项目，主要用于处理与 TON 区块链相关的操作。

## 项目结构
```
.
├── command/         # 命令行相关代码
├── controller/      # 控制器层
├── route/          # 路由配置
├── config/         # 配置相关
├── pkg/            # 公共包
├── main.go         # 程序入口
├── config.yml      # 配置文件
├── go.mod          # Go 模块文件
└── go.sum          # Go 依赖校验文件
```

## 主要功能
- HTTP 服务监听
- TON 区块链相关操作
- 配置管理

## 配置说明
配置文件 `config.yml` 包含以下主要配置项：
- 应用名称
- HTTP 服务监听地址
- TON 助记词
- API 密钥

## 运行说明
1. 确保已安装 Go 环境
2. 配置 `config.yml` 文件
3. 运行 `go run main.go` 启动服务

## 注意事项
- 默认 HTTP 服务监听地址为 `0.0.0.0:8000`
- 如需限制访问，可将监听地址改为 `127.0.0.1:8000`
- 请妥善保管 TON 助记词和 API 密钥 
