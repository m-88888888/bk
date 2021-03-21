package cmd

import (
	"bufio"
	"os"
	"testing"

	"github.com/m-88888888/bk/util"
)

func TestSave(t *testing.T) {
	err := Save()
	if err != nil {
		t.Fatalf("%v", err)
	}

	fileName, err := util.HistoryFile()
	if err != nil {
		t.Fatalf("%v", err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("%v", err)
	}
	currentDirName, err := os.Getwd()
	if err != nil {
		t.Fatalf("%v", err)
	}
	scanner := bufio.NewScanner(file)
	result := false
	for scanner.Scan() {
		if scanner.Text() == currentDirName {
			result = true
		}
	}
	if !result {
		t.Errorf("currentDirName don't contain in historyFile")
	}
}
