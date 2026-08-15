package model

import "sort"

type Account struct {
	ID      string
	Owner   string
	Balance int64
}

type TxnType string

const (
	TxnDeposit  TxnType = "deposit"
	TxnTransfer TxnType = "transfer"
)

type Transaction struct {
	ID     string
	From   string
	To     string
	Amount int64
	Type   TxnType
}

type Summary struct {
	Audited    int
	Failed     int
	Transferred int64
}

func ValidAmount(amount int64) bool {
	return amount > 0
}

func SortTransactions(txns []*Transaction) []*Transaction {
	sort.SliceStable(txns, func(i, j int) bool { return txns[i].ID < txns[j].ID })
	return txns
}

func BuildPages(txns []*Transaction, size int) [][]*Transaction {
	if size <= 0 {
		size = 1
	}
	out := make([][]*Transaction, 0, (len(txns)+size-1)/size)
	for i := 0; i < len(txns); i += size {
		end := i + size
		if end > len(txns) {
			end = len(txns)
		}
		p := make([]*Transaction, end-i)
		copy(p, txns[i:end])
		out = append(out, p)
	}
	return out
}

func MergeSummary(dst Summary, src Summary) Summary {
	dst.Audited += src.Audited
	dst.Failed += src.Failed
	dst.Transferred += src.Transferred
	return dst
}
