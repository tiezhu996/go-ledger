package worker

import (
	"context"
	"sync"

	"ledger/internal/model"
	"ledger/internal/service"
	"ledger/internal/store"
)

type Auditor interface {
	Audit(ctx context.Context, t *model.Transaction) error
}

type Reconcile struct {
	store   *store.Store
	svc     *service.Service
	auditor Auditor
	workers int
}

func New(s *store.Store, svc *service.Service, a Auditor, workers int) *Reconcile {
	if workers <= 0 {
		workers = 1
	}
	return &Reconcile{store: s, svc: svc, auditor: a, workers: workers}
}

func (r *Reconcile) Run(ctx context.Context) model.Summary {
	pages := r.svc.ListPages()

	var wg sync.WaitGroup
	ch := make(chan []*model.Transaction, len(pages))

	go func() {
		defer close(ch)
		for _, p := range pages {
			select {
			case <-ctx.Done():
				return
			case ch <- p:
			}
		}
	}()

	var mu sync.Mutex
	var sum model.Summary

	for i := 0; i < r.workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for page := range ch {
				var local model.Summary
				for _, t := range page {
					select {
					case <-ctx.Done():
						return
					default:
					}
					if err := r.auditor.Audit(ctx, t); err != nil {
						continue
					}
					local.Audited++
					if t.Type == model.TxnTransfer {
						local.Transferred += t.Amount
					}
				}
				mu.Lock()
				sum = model.MergeSummary(sum, local)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return sum
}
