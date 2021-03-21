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
		save()
	},
}

func save() {
	currentDirName, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	msg, err := SaveFilePath(currentDirName)
	if err != nil {
		panic(err)
	}
	fmt.Println(msg)
}

func SaveFilePath(filePath string) (string, error) {
	historyFileName, err := util.HistoryFile()
	historyFile, err := os.OpenFile(historyFileName, os.O_RDWR|os.O_APPEND, 0600)
	defer historyFile.Close()
	if err != nil {
		historyFile, err = os.OpenFile(historyFileName, os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			return "", err
		}
	}

	scanner := bufio.NewScanner(historyFile)
	for scanner.Scan() {
		if filePath == scanner.Text() {
			return "This directory is already bookmarked.", nil
		}
	}
	fmt.Fprintln(historyFile, filePath)
	return "Bookmark is successful.", nil
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
