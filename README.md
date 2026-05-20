# OBzhushou

Obsidian three-part collaborative system: Obsidian plugin + Cloud service + Mobile web

## Project Overview

OBzhushou (「O」meant Obsidian, 「B」meant Backup, and 「zhushou」meaning Assistant) is a comprehensive solution for extending Obsidian with cloud sync, mobile access, and AI-powered collaboration.

## Architecture

```
┌──────────────────┐      HTTPS API       ┌──────────────────┐
│  Obsidian 插件    │ ◄──────────────────► │  云端 Go 服务     │
│  (TypeScript)     │                      │  Go + Gin        │
│  本地 Vault 读写   │                      │  SQLite + S3     │
│                  │                      │  (同进程渲染前端)  │
└──────────────────┘                      └────────┬─────────┘
                                                  │
                                         Go 模板 + Alpine.js
                                                  │
                                         ┌────────▼─────────┐
                                         │  手机网页端       │
                                         │  移动端浏览器      │
                                         └──────────────────┘
```

## Quick Start

See `docs/implementation/PHASE1_DETAILED.md` for detailed development guide.

## License

MIT
