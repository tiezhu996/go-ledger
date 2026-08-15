package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Setenv("LEDGER_WORKERS", "")
	t.Setenv("LEDGER_PAGE_SIZE", "")
	c := Load()
	if c.Workers != 2 || c.PageSize != 2 {
		t.Fatalf("%+v", c)
	}
}

func TestLoadEnv(t *testing.T) {
	t.Setenv("LEDGER_WORKERS", "5")
	t.Setenv("LEDGER_PAGE_SIZE", "3")
	c := Load()
	if c.Workers != 5 || c.PageSize != 3 {
		t.Fatalf("%+v", c)
	}
}
