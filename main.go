package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func main() {
	var cmdBk = &cobra.Command{
		Use:   "bk [directory path]",
		Short: "bookmark your directory",
		Long:  "bk is bookmarker for your directory.",
		Run: func(cmd *cobra.Command, args []string) {
			// ブックマークの処理実行
			save()
		},
	}

	var cmdShow = &cobra.Command{
		Use:   "show",
		Short: "show bookmark directory",
		Long:  "show your bookmarked directory.",
		Run: func(cmd *cobra.Command, args []string) {
			// ブックマークの処理実行
			show()
		},
	}

	var rootCmd = &cobra.Command{Use: "bk"}
	rootCmd.AddCommand(cmdBk, cmdShow)
	rootCmd.Execute()
}

func HistoryFile() (string, error) {
	home, e := homedir.Dir()
	if e != nil {
		return "", e
	}
	historyFileName := home + "/.bk_history"
	return historyFileName, e
}

func save() error {
	historyFileName, e := HistoryFile()
	historyFile, e := os.OpenFile(historyFileName, os.O_RDWR|os.O_APPEND, 0600)
	defer historyFile.Close()
	if e != nil {
		historyFile, e = os.OpenFile(historyFileName, os.O_CREATE|os.O_WRONLY, 0600)
		if e != nil {
			return e
		}
	}
	currentDirName, e := os.Getwd()
	if e != nil {
		return e
	}
	var isDuplicate = false
	scanner := bufio.NewScanner(historyFile)
	for scanner.Scan() {
		if currentDirName == scanner.Text() {
			isDuplicate = true
			fmt.Println("This directory is already bookmarked.")
		}
	}
	if isDuplicate {
		return nil
	}
	fmt.Fprintln(historyFile, currentDirName)
	fmt.Println("Bookmark is successful.")
	return nil
}

func show() error {
	historyFileName, e := HistoryFile()
	historyFile, e := os.OpenFile(historyFileName, os.O_RDONLY, 0400)
	defer historyFile.Close()
	if e != nil {
		return e
	}
	bytes, e := ioutil.ReadAll(historyFile)
	if e != nil {
		return e
	}
	texts := string(bytes)
	texts = strings.TrimRight(texts, "\n")
	if len(texts) == 0 {
		return nil
	}
	fmt.Println(texts)
	return nil
}
