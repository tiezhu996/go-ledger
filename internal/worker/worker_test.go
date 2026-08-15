package worker

import (
	"context"
	"testing"

	"ledger/internal/model"
	"ledger/internal/service"
	"ledger/internal/store"
)

type okAuditor struct{}

func (okAuditor) Audit(ctx context.Context, t *model.Transaction) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return nil
}

func TestReconcileSummary(t *testing.T) {
	st := store.New()
	svc := service.New(st, 2)
	_, _ = svc.OpenAccount("a", "alice")
	_, _ = svc.OpenAccount("b", "bob")
	_ = svc.Deposit("a", 100)
	_ = svc.Transfer("a", "b", 30)
	_ = svc.Transfer("a", "b", 20)
	r := New(st, svc, okAuditor{}, 4)
	sum := r.Run(context.Background())
	if sum.Audited != 3 {
		t.Fatalf("Audited=%d want 3", sum.Audited)
	}
	if sum.Transferred != 50 {
		t.Fatalf("Transferred=%d want 50", sum.Transferred)
	}
}

func TestReconcileCancel(t *testing.T) {
	st := store.New()
	svc := service.New(st, 2)
	_, _ = svc.OpenAccount("a", "alice")
	for i := 0; i < 50; i++ {
		_ = svc.Deposit("a", 1)
	}
	r := New(st, svc, okAuditor{}, 1)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sum := r.Run(ctx)
	if sum.Audited != 0 {
		t.Fatalf("Audited=%d want 0 (cancelled before run)", sum.Audited)
	}
}
