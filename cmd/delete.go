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
	r, w := io.Pipe()
	show.Stdout = w
	peco.Stdin = r
	var out bytes.Buffer
	peco.Stdout = &out

	show.Start()
	peco.Start()
	show.Wait()
	w.Close()
	peco.Wait()

	return DeleteFilePath(out.String())
}

func DeleteFilePath(filePath string) error {
	deletePath := strings.TrimRight(filePath, "\n")

	historyFileName, err := util.HistoryFile()
	historyFile, err := os.OpenFile(historyFileName, os.O_RDONLY, 0600)
	if err != nil {
		return err
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
		historyFileW, err := os.OpenFile(historyFileName, os.O_WRONLY|os.O_TRUNC, 0600)
		if err != nil {
			return err
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
