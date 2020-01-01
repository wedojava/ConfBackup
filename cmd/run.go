package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/wedojava/Cbackup/internal/commons"
)

func passArgVerify() bool {
	if os.Args[1] == "yeezy350" {
		return true
	} 
	return false
	
}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
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
	path := getCurrentPath()
	fmt.Println(path)
	config := commons.LoadConf("../..", "conf.json")
	fmt.Println(config)
	// 3. task execute

}
