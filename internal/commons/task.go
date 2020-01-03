package commons

// TaskListRun will run hosts that in conf.json.
func (conf *Conf) ConfTaskListRun() {
	hosts := conf.HostList
	for _, host := range hosts {
		// 3.1. login
		t, _ := host.HostLogin()
		// 3.2. Task RUN
		conf.ConfHistoryBackup(t)
		conf.ConfGetConfig(t)
		conf.ConfHistoryRecover(t, host)
	}
}
