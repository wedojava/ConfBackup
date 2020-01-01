package commons

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// Conf is the struction of conf.json
type Conf struct {
	CmdList struct {
		HistoryBackup  []string
		GetConfig      []string
		HistoryRecover []string
	}
	IPList []struct {
		IP         string
		Port       string
		Type       string
		Username   string
		Password   string
		SuPassword string
	}
}

// Check errors and log out
func Check(e error) {
	if e != nil {
		// panic(e)
		log.Fatalln("Error:", e)
	}
}

// GetConfRaw will get the raw data of config file to string
// func GetConfRaw(rootpath, confName string) string {
// 	filename := filepath.Join(rootpath, confName)
// 	dat, err := ioutil.ReadFile(filename)
// 	Check(err)
// 	return string(dat)
// }

// LoadConf will load json file and return conf struct
func LoadConf(rootpath, confName string) Conf {
	filename := filepath.Join(rootpath, confName)
	dat, err := ioutil.ReadFile(filename)
	Check(err)
	var confObj Conf
	errUnmarshal := json.Unmarshal(dat, &confObj)
	if errUnmarshal != nil {
		fmt.Println("error:", errUnmarshal)
	}
	fmt.Printf("%+v", confObj)
	return confObj
}
