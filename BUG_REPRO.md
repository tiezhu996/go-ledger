# BUG_REPRO

## Bug 是什么
BuildPages / OrderIDs / ListPages 返回共享底层数组的子切片，worker 又用 `page[:len(page)-1]` 切掉最后一条，导致交易漏掉、列表互相串改。

## 如何触发
`go test ./...`

## 错误信息
- TestBuildPagesFresh / TestRecordTxnOrderFresh 失败（切片共享底层数组）。
- TestReconcileSummary 失败（Audited 计数不对）。
