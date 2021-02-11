package util

import "testing"

func TestHistoryFile(t *testing.T) {
	f, err := HistoryFile()
	if err != nil {
		t.Fatal(err)
	}
	if f == "" {
		t.Errorf("HistoryFile() = %s, want .bk_history", f)
	}
}
