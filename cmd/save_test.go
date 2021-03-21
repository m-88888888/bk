package cmd

import (
	"bufio"
	"os"
	"testing"

	"github.com/m-88888888/bk/util"
)

func TestSaveFilePath(t *testing.T) {
	testPath := "/Users/username/dev"
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
	result := false
	for scanner.Scan() {
		if scanner.Text() == testPath {
			result = true
		}
	}
	if !result && len(msg) == 0 {
		t.Errorf("currentDirName don't contain in historyFile")
	}
}
