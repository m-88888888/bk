package util

import "github.com/mitchellh/go-homedir"

// HistoryFile is return history file's absolute path
func HistoryFile() (string, error) {
	home, e := homedir.Dir()
	if e != nil {
		return "", e
	}
	historyFileName := home + "/.bk_history"
	return historyFileName, e
}
