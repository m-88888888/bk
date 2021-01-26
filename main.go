package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
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
			save()
		},
	}

	var cmdShow = &cobra.Command{
		Use:   "show",
		Short: "show bookmark directory",
		Long:  "show your bookmarked directory.",
		Run: func(cmd *cobra.Command, args []string) {
			show()
		},
	}

	var cmdDelete = &cobra.Command{
		Use:   "delete",
		Short: "delete bookmark directory.",
		Long:  "delete your bookmark directory from history.",
		Run: func(cmd *cobra.Command, args []string) {
			delete()
		},
	}

	var cmdJp = &cobra.Command{
		Use: "jp",
		Run: func(cmd *cobra.Command, args []string) {
			jump()
		},
	}

	var rootCmd = &cobra.Command{Use: "bk"}
	rootCmd.AddCommand(cmdSave, cmdShow, cmdDelete, cmdJp)
	rootCmd.Execute()
}

// HistoryFile ...
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

func delete() error {
	show := exec.Command("bk", "show")
	peco := exec.Command("peco")
	// io.Writerとio.Readerをつなげる
	r, w := io.Pipe()
	show.Stdout = w
	// bk showの内容をpecoに渡す
	peco.Stdin = r
	var out bytes.Buffer
	peco.Stdout = &out

	show.Start()
	peco.Start()
	show.Wait()
	w.Close()
	peco.Wait()
	// pecoで選択した値をoutで受け取る
	deletePath := strings.TrimRight(out.String(), "\n")

	historyFileName, e := HistoryFile()
	historyFile, e := os.OpenFile(historyFileName, os.O_RDONLY, 0600)
	if e != nil {
		return e
	}
	var texts string
	scanner := bufio.NewScanner(historyFile)
	for scanner.Scan() {
		t := scanner.Text()
		if deletePath != t {
			texts = texts + t + "\n"
		}
	}
	historyFile.Close()
	texts = strings.TrimRight(texts, "\n")
	if len(texts) == 0 {
		exec.Command("cp", "/dev/null", historyFileName).Start()
	} else {
		historyFileW, e := os.OpenFile(historyFileName, os.O_WRONLY|os.O_TRUNC, 0600)
		if e != nil {
			return e
		}
		fmt.Fprintln(historyFileW, texts)
		historyFileW.Close()
	}
	return nil
}

func jump() error {
	show := exec.Command("bk", "show")
	peco := exec.Command("peco")
	r, w := io.Pipe()
	show.Stdout = w
	peco.Stdin = r
	show.Start()
	peco.Start()
	return nil
}
