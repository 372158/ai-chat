本项目为一个AI聊天服务后端，技术栈为：GO + GIN + GORM + Redis + SSE.

| 组件    | 技术          | 说明            |
| ------- | ------------- | --------------- |
| Web框架 | Gin           | 路由 + 中间件   |
| ORM     | GORM          | MySQL 读写      |
| 数据库  | MySQL (Docker) | 会话缓存 + 限流 |
| LLM     | 智谱 GLM      | OpenAI 兼容协议 |
| 流式    | SSE           | 打字机          |

