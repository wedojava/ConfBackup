package commons

import "testing"

func TestHostLogin(t *testing.T) {
	conf := LoadConf("../..", "conf.json")
	t.Run("The right conn info.", func(t *testing.T) {
		host := conf.HostList[0]
		_, err := host.HostLogin()
		if err != nil {
			t.Errorf("Login success but err not nil.")
		}
	})
	t.Run("The wrong conn info.", func(t *testing.T) {
		host := conf.HostList[1]
		_, err := host.HostLogin()
		if err == nil {
			t.Errorf("Login failure but err is nil.")
		}
	})
}

func TestConf_ConfGetConfig(t *testing.T) {
	conf := LoadConf("../..", "conf.json")
	t.Run("save with diff names.", func(t *testing.T) {
		host := conf.HostList[0]
		telnetObj, err := host.HostLogin()
		conf.ConfGetConfig(telnetObj, host)
		if err != nil {
			t.Errorf("Login success but err not nil.")
		}
	})
}

func TestConfHistoryBackup(t *testing.T) {
	conf := LoadConf("../..", "conf.json")
	t.Run("The right conn info.", func(t *testing.T) {
		host := conf.HostList[0]
		telnetObj, err := host.HostLogin()
		conf.ConfHistoryBackup(telnetObj)
		conf.ConfGetConfig(telnetObj, host)
		conf.ConfHistoryRecover(telnetObj, host)
		if err != nil {
			t.Errorf("Login success but err not nil.")
		}
	})
	t.Run("The wrong conn info.", func(t *testing.T) {
		host := conf.HostList[1]
		telnetObj, err := host.HostLogin()
		conf.ConfHistoryBackup(telnetObj)
		conf.ConfGetConfig(telnetObj, host)
		conf.ConfHistoryRecover(telnetObj, host)
		if err == nil {
			t.Errorf("Login failure but err is nil.")
		}
	})
}
