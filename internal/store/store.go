package store

import (
	"errors"
	"sync"

	"ledger/internal/model"
)

var (
	ErrAccountNotFound   = errors.New("account not found")
	ErrAccountExists     = errors.New("account already exists")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrTxnExists         = errors.New("transaction already exists")
)

type Store struct {
	mu       sync.RWMutex
	accounts map[string]*model.Account
	txns     map[string]*model.Transaction
	order    []string
}

func New() *Store {
	return &Store{
		accounts: make(map[string]*model.Account),
		txns:     make(map[string]*model.Transaction),
		order:    []string{},
	}
}

func (s *Store) CreateAccount(a *model.Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.accounts[a.ID]; ok {
		return ErrAccountExists
	}
	s.accounts[a.ID] = a
	return nil
}

func (s *Store) GetAccount(id string) (*model.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	a, ok := s.accounts[id]
	if !ok {
		return nil, ErrAccountNotFound
	}
	return a, nil
}

func (s *Store) Deposit(id string, amount int64) error {
	a, ok := s.accounts[id]
	if !ok {
		return ErrAccountNotFound
	}
	a.Balance += amount
	return nil
}

func (s *Store) Transfer(from, to string, amount int64) error {
	f, ok := s.accounts[from]
	if !ok {
		return ErrAccountNotFound
	}
	t, ok := s.accounts[to]
	if !ok {
		return ErrAccountNotFound
	}
	if f.Balance < amount {
		return ErrInsufficientFunds
	}
	f.Balance -= amount
	t.Balance += amount
	return nil
}

func (s *Store) RecordTxn(t *model.Transaction) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.txns[t.ID]; ok {
		return ErrTxnExists
	}
	s.txns[t.ID] = t
	s.order = append(s.order, t.ID)
	return nil
}

func (s *Store) ListTransactions() []*model.Transaction {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]*model.Transaction, 0, len(s.order))
	for _, id := range s.order {
		out = append(out, s.txns[id])
	}
	return out
}

func (s *Store) OrderIDs() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]string, len(s.order))
	copy(out, s.order)
	return out
}

func (s *Store) Balance(id string) (int64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	a, ok := s.accounts[id]
	if !ok {
		return 0, ErrAccountNotFound
	}
	return a.Balance, nil
}
