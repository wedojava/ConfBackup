package commons

import "testing"

func TestConfTaskListRun(t *testing.T) {
	conf := LoadConf("../..", "conf.json")
	conf.ConfTaskListRun()
}
