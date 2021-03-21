package cmd

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/m-88888888/bk/util"
)

func TestDeleteFilePath(t *testing.T) {
	testPath := "/Users/username/directory"
	_, err := SaveFilePath(testPath)
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = DeleteFilePath(testPath)
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
	for scanner.Scan() {
		if scanner.Text() == testPath {
			fmt.Println(scanner.Text())
			t.Errorf("Could not delete the specified file path.\nfile path: %v", testPath)
		}
	}
}
