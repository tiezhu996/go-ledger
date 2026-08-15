package service

import (
	"errors"
	"fmt"

	"ledger/internal/model"
	"ledger/internal/store"
)

var ErrInvalidAmount = errors.New("invalid amount")

type Service struct {
	store    *store.Store
	pageSize int
}

func New(s *store.Store, pageSize int) *Service {
	if pageSize <= 0 {
		pageSize = 1
	}
	return &Service{store: s, pageSize: pageSize}
}

func (svc *Service) OpenAccount(id, owner string) (*model.Account, error) {
	a := &model.Account{ID: id, Owner: owner}
	if err := svc.store.CreateAccount(a); err != nil {
		return nil, fmt.Errorf("open account %s: %w", id, err)
	}
	return a, nil
}

func (svc *Service) Deposit(id string, amount int64) error {
	if err := svc.store.Deposit(id, amount); err != nil {
		return fmt.Errorf("deposit %s: %w", id, err)
	}
	t := &model.Transaction{ID: fmt.Sprintf("d-%s-%d", id, amount), To: id, Amount: amount, Type: model.TxnDeposit}
	return svc.record(t)
}

func (svc *Service) Transfer(from, to string, amount int64) error {
	if err := svc.store.Transfer(from, to, amount); err != nil {
		return fmt.Errorf("transfer %s->%s: %w", from, to, err)
	}
	t := &model.Transaction{ID: fmt.Sprintf("t-%s-%s-%d", from, to, amount), From: from, To: to, Amount: amount, Type: model.TxnTransfer}
	return svc.record(t)
}

func (svc *Service) record(t *model.Transaction) error {
	if err := svc.store.RecordTxn(t); err != nil {
		return fmt.Errorf("record txn %s: %w", t.ID, err)
	}
	return nil
}

func (svc *Service) Balance(id string) (int64, error) {
	b, err := svc.store.Balance(id)
	if err != nil {
		return 0, fmt.Errorf("balance %s: %w", id, err)
	}
	return b, nil
}

func (svc *Service) ListPages() [][]*model.Transaction {
	txns := svc.store.ListTransactions()
	model.SortTransactions(txns)
	return model.BuildPages(txns, svc.pageSize)
}
