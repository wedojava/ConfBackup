package main

import (
	"github.com/wedojava/ConfBackup/internal/commons"
	"os"
)

func passArgVerify() bool {
	if len(os.Args) > 1 {
		if os.Args[1] == "yeezy350" {
			return true
		}
	}
	return false

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// 1. args identify
	if !(passArgVerify()) {
		os.Exit(3)
	}
	// 2. get config
	config := commons.LoadConf("./", "conf.json")
	// 3. task execute
	config.ConfTaskListRun()
}
