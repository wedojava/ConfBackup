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
	CMDList  `json:"cmd_list"`
	HostList `json:"host_list"`
}

type CMDList struct {
	HistoryBackup  []string `json:"history_backup"`
	GetConfig      []string `json:"get_config"`
	HistoryRecover []string `json:"history_recover"`
}

type HostList []Host

type Host struct {
	IP               string `json:"ip"`
	Port             string `json:"port"`
	Type             string `json:"type"`
	IsAuthentication bool   `json:"is_authentication"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	SuPassword       string `json:"su_password"`
}

// Check errors and log out
func Check(e error) {
	if e != nil {
		// panic(e)
		log.Fatalln("Error:", e)
	}
}

//GetConfRaw will get the raw data of config file to string
//func GetConfRaw(rootpath, confName string) string {
//	filename := filepath.Join(rootpath, confName)
//	dat, err := ioutil.ReadFile(filename)
//	Check(err)
//	return string(dat)
//}

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
	return confObj
}
