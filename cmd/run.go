package main

import (
	"github.com/wedojava/ConfBackup/internal/commons"
	"os"
)

func enterHurdle(keyword string) {
	if len(os.Args) > 1 {
		if os.Args[1] == keyword {
			println("Here we go.")
		}
	}
	os.Exit(3)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// 1. args identify
	enterHurdle("yeezy350")
	// 2. get config
	config := commons.LoadConf("./", "conf.json")
	// 3. task execute
	config.ConfTaskListRun()
}
