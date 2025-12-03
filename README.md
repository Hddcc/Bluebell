# Bluebell - 简洁的社区论坛应用

一个使用 Go 语言开发的简洁社区论坛 Web 应用。

## 功能特性

*   **用户管理**：支持用户注册和登录功能。
*   **帖子发布**：用户可以创建和发布自己的帖子。
*   **帖子浏览**：按时间或分数排序，浏览帖子列表。
*   **社区功能**：查看社区分类及详情。
*   **投票系统**：用户可以对帖子进行投票。

## 技术栈

*   **后端**：Go
*   **Web 框架**：Gin
*   **数据库**：MySQL, Redis
*   **配置**：Viper
*   **日志**：Zap

## 快速开始

### 环境依赖

*   Go (建议 1.18 或更高版本)
*   一个正在运行的 MySQL 服务
*   一个正在运行的 Redis 服务

### 安装与运行

1.  **克隆仓库：**
    ```bash
    git clone https://github.com/your_username/bluebell.git
    cd bluebell
    ```

2.  **配置服务：**
    打开 `conf/config.yaml` 文件，修改其中的 `mysql` 和 `redis` 部分为您的本地连接信息。

3.  **安装依赖：**
    此命令将下载所有必需的 Go 模块。
    ```bash
    go mod tidy
    ```

4.  **运行程序：**
    ```bash
    go run main.go
    ```
    程序默认将在 `http://127.0.0.1:8080` 启动。

## 项目结构

```
.
├── conf/           # 配置文件 (config.yaml)
├── controllers/    # 处理请求和响应的控制器
├── dao/            # 数据访问对象 (MySQL, Redis 连接)
├── docs/           # Swagger API 文档
├── logger/         # 日志配置
├── logic/          # 业务逻辑处理
├── main.go         # 程序入口文件
├── middleware/     # 中间件 (JWT认证, 限流)
├── models/         # 数据模型结构体
├── pkg/            # 自定义工具包 (JWT, Snowflake)
├── routes/         # API 路由配置
├── settings/       # 配置加载
├── static/         # 编译后的前端静态资源 (JS, CSS)
└── templates/      # 主页 index.html 模板
```
