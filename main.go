package main

import (
	"fmt"
	"os"
	"path/filepath"
	"task7/cmd"
	"task7/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	errorhandler(db.Init(dbPath))
	errorhandler(cmd.RootCmd.Execute())
}

func errorhandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
