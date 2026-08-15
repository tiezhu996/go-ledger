package model

import "testing"

func TestValidAmount(t *testing.T) {
	if !ValidAmount(1) || ValidAmount(0) || ValidAmount(-1) {
		t.Fatal("ValidAmount wrong")
	}
}

func TestSortTransactions(t *testing.T) {
	in := []*Transaction{{ID: "c"}, {ID: "a"}, {ID: "b"}}
	got := SortTransactions(in)
	want := []string{"a", "b", "c"}
	for i, id := range want {
		if got[i].ID != id {
			t.Fatalf("order=%v want %v", got, want)
		}
	}
}

func TestBuildPagesFresh(t *testing.T) {
	in := []*Transaction{{ID: "1"}, {ID: "2"}, {ID: "3"}}
	p := BuildPages(in, 2)
	if len(p) != 2 {
		t.Fatalf("len=%d", len(p))
	}
	p[0][0] = &Transaction{ID: "x"}
	if in[0].ID != "1" {
		t.Fatal("mutating page corrupted input")
	}
}

func TestMergeSummary(t *testing.T) {
	got := MergeSummary(Summary{Audited: 1}, Summary{Audited: 2, Failed: 3})
	if got.Audited != 3 || got.Failed != 3 {
		t.Fatalf("%+v", got)
	}
}
