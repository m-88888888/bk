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
	var cmdSave = &cobra.Command{
		Use:   "save [directory path]",
		Short: "bookmark your directory",
		Long:  "save is bookmarker for your directory.",
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
	rootCmd.AddCommand(cmdSave, cmdShow)
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
	// 第3引数の'0600'はファイルモード。「ll」コマンドで出てくるrwxr-xr-xとかのこと。
	// ファイルを読み込みできたら書き込みし、できなければファイルを作って書き込む。
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
	scanner := bufio.NewScanner(historyFile)
	for scanner.Scan() {
		if currentDirName == scanner.Text() {
			fmt.Println("This directory is already bookmarked.")
			return nil
		}
	}
	// historyFIleに現在のディレクトリを書き込み
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

//func delete() error {
//}
