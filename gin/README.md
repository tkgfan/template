# Gin 项目模板

> 项目简要说明

## 配置说明

### 环境变量

- MODEL: 项目运行环境（dev、test、prod）
- PORT: 端口号（默认 8888）
- TIMEOUT: 请求超时时间（默认 60000 毫秒）
- LOG_LEVEL: 日志级别（info、warn、error、panic 默认 info）

## 目录说明

- common: 通用工具代码
  - cerr: 错误码定义
  - ctxdata: 上下文数据
  - middleware: 中间件
  - result: 结果封装
- conf: 项目配置目录
- models: 模型存放目录 
- router: 路由目录
- service: 业务逻辑目录
- ci.sh: 自动构建 Docker 镜像脚本
- Dockerfile
- main.go: 入口函数
- CHANGELOG.md: 变更日志
- version: 版本号
- README: 项目说明
  
## 相关文档

