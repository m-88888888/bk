package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/m-88888888/bk/util"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show bookmark directory",
	Long:  "show your bookmarked directory.",
	Run: func(cmd *cobra.Command, args []string) {
		show()
	},
}

func show() error {
	historyFileName, e := util.HistoryFile()
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

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
