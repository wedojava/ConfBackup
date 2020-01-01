package commons

import (
	"reflect"
	"testing"
)

// func TestGetConfRaw(t *testing.T) {
// 	expected := "cmd_list:"
// 	actual := GetConfRaw("../..", "conf.json")

// 	if actual != expected {
// 		t.Errorf("got %q expected %q", actual, expected)
// 	}
// }

func TestLoadConf(t *testing.T) {
	var expected Conf
	actual := LoadConf("../..", "conf.json")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}
