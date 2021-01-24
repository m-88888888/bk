package main

import (
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestHistoryFile(t *testing.T) {
	result, err := HistoryFile()
	if err != nil {
		t.Errorf("something is wrong. %v", err)
	}
	home, e := homedir.Dir()
	if e != nil {
		t.Errorf("error")
	}
	historyFileName := home + "/.bk_history"
	if result != historyFileName {
		t.Errorf("historyFileName() = %v, want %v", result, historyFileName)
	}
}
