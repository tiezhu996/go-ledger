package store

import (
	"testing"

	"ledger/internal/model"
)

func TestAccountLifecycle(t *testing.T) {
	s := New()
	if err := s.CreateAccount(&model.Account{ID: "a"}); err != nil {
		t.Fatal(err)
	}
	if err := s.CreateAccount(&model.Account{ID: "a"}); err != ErrAccountExists {
		t.Fatalf("dup err=%v", err)
	}
	if err := s.Deposit("a", 100); err != nil {
		t.Fatal(err)
	}
	if _, err := s.GetAccount("nope"); err != ErrAccountNotFound {
		t.Fatalf("err=%v", err)
	}
	b, _ := s.Balance("a")
	if b != 100 {
		t.Fatalf("balance=%d", b)
	}
}

func TestTransfer(t *testing.T) {
	s := New()
	_ = s.CreateAccount(&model.Account{ID: "a"})
	_ = s.CreateAccount(&model.Account{ID: "b"})
	_ = s.Deposit("a", 100)
	if err := s.Transfer("a", "b", 60); err != nil {
		t.Fatal(err)
	}
	if err := s.Transfer("a", "b", 60); err != ErrInsufficientFunds {
		t.Fatalf("err=%v", err)
	}
	ba, _ := s.Balance("a")
	bb, _ := s.Balance("b")
	if ba != 40 || bb != 60 {
		t.Fatalf("a=%d b=%d", ba, bb)
	}
}

func TestRecordTxnOrderFresh(t *testing.T) {
	s := New()
	_ = s.RecordTxn(&model.Transaction{ID: "t1"})
	_ = s.RecordTxn(&model.Transaction{ID: "t2"})
	if err := s.RecordTxn(&model.Transaction{ID: "t1"}); err != ErrTxnExists {
		t.Fatalf("err=%v", err)
	}
	ids := s.OrderIDs()
	ids[0] = "x"
	if s.OrderIDs()[0] != "t1" {
		t.Fatal("OrderIDs aliased internal order")
	}
	if len(s.ListTransactions()) != 2 {
		t.Fatal("ListTransactions len")
	}
}
