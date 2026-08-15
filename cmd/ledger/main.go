package main

import (
	"fmt"

	"ledger/internal/config"
	"ledger/internal/service"
	"ledger/internal/store"
)

func main() {
	cfg := config.Load()
	st := store.New()
	svc := service.New(st, cfg.PageSize)
	_ = svc
	fmt.Println("ledger ready")
}
