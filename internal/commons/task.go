package commons

// TaskRun will run one item's task.
func (host *Host) TaskRun() {
	// 3.1. login
	host.Login()
	// 3.2. history backup
	//ti.HistoryBackup()
	// 3.3. do it
	// 3.4. history recover
}
