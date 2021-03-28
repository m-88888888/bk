package cmd

import (
	"bufio"
	"os"
	"testing"

	"github.com/m-88888888/bk/util"
)

func TestSaveFilePath(t *testing.T) {
	testPath := "/usr/local/bin"
	msg, err := SaveFilePath(testPath)
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
	scanner := bufio.NewScanner(file)
	pathNotExists := true
	for scanner.Scan() {
		if scanner.Text() == testPath {
			pathNotExists = false
		}
	}
	if pathNotExists || len(msg) == 0 {
		t.Errorf("%v don't contain in historyFile", testPath)
	}
}
