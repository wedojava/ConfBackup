package commons

import "testing"

func TestTaskItem_Login(t *testing.T) {
	ti := TaskItem{
		IP:               "192.168.111.82",
		Port:             "23",
		IsAuthentication: true,
		Username:         "foobar",
		Password:         "123qwe!@#QWE",
		SuPassword:       "123qwe!@#QWE",
	}
	ti.Login()

}
