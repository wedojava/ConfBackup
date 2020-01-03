package commons

import "testing"

func TestHostLogin(t *testing.T) {
	host := Host{
		IP:               "192.168.111.82",
		Port:             "23",
		IsAuthentication: true,
		Username:         "foobar",
		Password:         "123qwe!@#QWE",
		SuPassword:       "123qwe!@#QWE",
	}
	host.Login()

}
