package commons

import (
	"fmt"
	"testing"
)

//func TestGetConfRaw(t *testing.T) {
//	expected := "cmd_list:"
//	actual := GetConfRaw("../..", "conf.json")
//
//	if actual != expected {
//		t.Errorf("got %q expected %q", actual, expected)
//	}
//}

func TestLoadConf(t *testing.T) {
	actual := LoadConf("../..", "conf.json")
	fmt.Println(actual)
}
