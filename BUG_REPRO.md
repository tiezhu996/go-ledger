# BUG_REPRO

## Bug 是什么
Store 构造时未初始化 txns map，service.Deposit/Transfer 内部走 store.RecordTxn，首次向 nil map 写入触发 panic。

## 如何触发
`go test ./internal/store -run TestRecordTxnOrderFresh`

## 错误信息
`panic: assignment to entry in nil map`
