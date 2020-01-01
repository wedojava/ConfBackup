package commons

// TaskItem get from config file.
type TaskItem struct {
	IP               string
	Port             string
	IsAuthentication bool
	Username         string
	Password         string
	SuPassword       string
}

// TaskRun will run one item's task.
func (ti *TaskItem) TaskRun() {
	// 3.1. login
	ti.Login()
	// 3.2. history backup
	// 3.3. do it
	// 3.4. history recover
}
