/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/m-88888888/bk/util"
	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save [directory path]",
	Short: "bookmark your directory",
	Long:  "save is bookmarker for your directory.",
	Run: func(cmd *cobra.Command, args []string) {
		Save()
	},
}

func Save() error {
	historyFileName, e := util.HistoryFile()
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

func init() {
	rootCmd.AddCommand(saveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
