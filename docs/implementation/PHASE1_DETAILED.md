# Phase 1 实施 - 快速启动指南

## Phase 1 核心同步闭环实施（2-3 周）

### Week 1: 云端基础设施

#### 1.1 SQLite 数据库设计

**数据表设计:**

```sql
-- 用户表
CREATE TABLE users (
    id TEXT PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 设备表
CREATE TABLE devices (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    device_code TEXT UNIQUE NOT NULL,
    device_name TEXT,
    paired BOOLEAN DEFAULT FALSE,
    code_expires_at DATETIME,
    last_sync DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 同步日志表
CREATE TABLE sync_log (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    device_id TEXT NOT NULL,
    file_path TEXT NOT NULL,
    action TEXT NOT NULL,
    content_hash TEXT,
    version_id TEXT UNIQUE,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    synced BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (device_id) REFERENCES devices(id)
);

-- Markdown 元数据表
CREATE TABLE markdown_metadata (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    file_path TEXT UNIQUE NOT NULL,
    s3_key TEXT,
    content_hash TEXT,
    file_size INTEGER,
    last_modified DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 冲突记录表
CREATE TABLE conflicts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    file_path TEXT NOT NULL,
    device_id TEXT NOT NULL,
    version_a TEXT,
    version_b TEXT,
    resolved BOOLEAN DEFAULT FALSE,
    resolution TEXT,
    resolved_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (device_id) REFERENCES devices(id)
);
```

#### 1.2 设备码认证系统

**API 端点:**
```
POST /api/auth/device
请求: { "email": "user@example.com", "device_name": "iPhone 14" }
响应: { "device_code": "123456", "expires_at": "2026-05-27T11:30:00Z" }

POST /api/auth/device/verify
请求: { "device_code": "123456" }
响应: { "device_id": "dev_xxx", "user_id": "user_xxx", "token": "jwt_token" }
```

### Week 2: 同步核心引擎

#### 2.1 推模式 (Obsidian → Cloud)

```
POST /api/sync/push

请求体:
{
  "device_id": "dev_xxx",
  "device_token": "jwt_token",
  "changes": [
    {
      "file_path": ".obsidian-zhushou/notes/note1.md",
      "action": "create",
      "content": "# Note Content",
      "content_hash": "sha256_hash",
      "timestamp": "2026-05-20T11:30:00Z"
    }
  ]
}

响应:
{
  "success": true,
  "sync_version": "v_123",
  "next_sync_after": "2026-05-20T11:45:00Z"
}
```

#### 2.2 拉模式 (Cloud → Mobile)

```
GET /api/sync/pull?since=<last_sync_timestamp>

响应:
{
  "files": [
    {
      "file_path": ".obsidian-zhushou/notes/note1.md",
      "content": "# Note Content",
      "content_hash": "sha256_hash",
      "timestamp": "2026-05-20T11:30:00Z",
      "has_conflict": false
    }
  ],
  "sync_version": "v_123",
  "next_sync_after": "2026-05-20T11:45:00Z"
}
```

### Week 3-4: Obsidian 插件 & 手机网页端

## 验收标准 (Phase 1)

- [x] 云端 API 完整可用
- [x] 便签端到端同步
- [x] 冲突处理
- [x] 离线支持
