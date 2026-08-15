package config

import (
	"os"
	"strconv"
)

type Config struct {
	Workers  int
	PageSize int
}

func Load() *Config {
	return &Config{
		Workers:  getInt("LEDGER_WORKERS", 2),
		PageSize: getInt("LEDGER_PAGE_SIZE", 2),
	}
}

func getInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil || n <= 0 {
		return def
	}
	return n
}
