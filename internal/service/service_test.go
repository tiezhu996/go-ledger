package service

import (
	"errors"
	"testing"

	"ledger/internal/store"
)

func TestOpenDepositTransfer(t *testing.T) {
	s := store.New()
	svc := New(s, 2)
	if _, err := svc.OpenAccount("a", "alice"); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.OpenAccount("b", "bob"); err != nil {
		t.Fatal(err)
	}
	if err := svc.Deposit("a", 100); err != nil {
		t.Fatal(err)
	}
	if err := svc.Transfer("a", "b", 40); err != nil {
		t.Fatal(err)
	}
	if err := svc.Deposit("a", -1); !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("invalid amount err=%v", err)
	}
	ba, _ := svc.Balance("a")
	bb, _ := svc.Balance("b")
	if ba != 60 || bb != 40 {
		t.Fatalf("a=%d b=%d", ba, bb)
	}
}

func TestWrapping(t *testing.T) {
	s := store.New()
	svc := New(s, 2)
	if err := svc.Deposit("nope", 10); !errors.Is(err, store.ErrAccountNotFound) {
		t.Fatalf("errors.Is=false err=%v", err)
	}
	if _, err := svc.Balance("nope"); !errors.Is(err, store.ErrAccountNotFound) {
		t.Fatalf("errors.Is=false err=%v", err)
	}
}

func TestListPagesOrder(t *testing.T) {
	s := store.New()
	svc := New(s, 2)
	_, _ = svc.OpenAccount("a", "alice")
	_, _ = svc.OpenAccount("b", "bob")
	_ = svc.Deposit("a", 100)
	_ = svc.Transfer("a", "b", 10)
	_ = svc.Transfer("a", "b", 20)
	pages := svc.ListPages()
	if len(pages) == 0 {
		t.Fatal("no pages")
	}
	if pages[0][0].ID != "d-a-100" {
		t.Fatalf("page0[0]=%s want d-a-100", pages[0][0].ID)
	}
}
