# BUG_REPRO

## Bug 是什么
ValidAmount 判断方向反了、Deposit/Transfer 去掉锁、service 去掉金额校验、对账汇总去掉锁并把 wg.Add 放进 goroutine，导致余额算错、负数金额可入账、对账结果错误并有数据竞争。

## 如何触发
`go test -race ./...`

## 错误信息
- TestValidAmount / TestOpenDepositTransfer 失败（负数金额被接受）。
- `go test -race` 报 data race；对账汇总计数不对。
