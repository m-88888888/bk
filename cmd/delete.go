package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/m-88888888/bk/util"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete bookmark directory.",
	Long:  "delete your bookmark directory from history.",
	Run: func(cmd *cobra.Command, args []string) {
		delete()
	},
}

func delete() error {
	show := exec.Command("bk", "show")
	peco := exec.Command("peco")
	// io.Writerとio.Readerをつなげる
	// bk showの出力内容をpecoにつなげるため？
	r, w := io.Pipe()
	// showの標準出力にio.Pipeを渡すことで、pecoの標準入力に値が渡される？
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

	historyFileName, e := util.HistoryFile()
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

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
