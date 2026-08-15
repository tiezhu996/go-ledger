# ledger

一个用 Go 写的内存记账/转账服务，演示分层、并发转账、对账 worker 与分页查询。

## 功能
- 开户、存款、转账、余额查询
- 交易记录分页查询
- 并发对账 worker 汇总审计结果

## 目录结构
```
cmd/ledger/          程序入口
internal/config/     环境配置
internal/model/      模型与纯工具函数
internal/store/      内存存储（账户 + 交易 + 锁）
internal/service/    业务逻辑
internal/worker/     对账 worker 池
```

## 运行与测试
```bash
go build ./...
go test ./...
go run ./cmd/ledger
```

## 环境变量
| 变量 | 说明 | 默认值 |
|------|------|--------|
| `LEDGER_WORKERS` | worker 数量 | `2` |
| `LEDGER_PAGE_SIZE` | 分页大小 | `2` |
