package commons

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/ziutek/telnet"
)

const timeout = 10 * time.Second

// execute command, t is obj of telnet.Conn, strSend is the command you send, strPrompt is the prompt string before you type the command
func execute(t *telnet.Conn, strSend string, strPrompt ...string) {
	// 1. get expect prompt string
	Check(t.SetReadDeadline(time.Now().Add(timeout)))
	Check(t.SkipUntil(strPrompt...))
	// 2. send command string
	Check(t.SetWriteDeadline(time.Now().Add(timeout)))
	buf := make([]byte, len(strSend)+1)
	copy(buf, strSend)
	buf[len(strSend)] = '\n'
	_, err := t.Write(buf)
	Check(err)
}

// Login is used to login Unix, Juniper or Cisco devices.
func (ti *TaskItem) Login() {
	typ, dst, user, passwd := "juniper", net.JoinHostPort(ti.IP, ti.Port), ti.Username, ti.Password
	t, err := telnet.Dial("tcp", dst)
	Check(err)
	t.SetUnixWriteMode(true)
	var data []byte
	switch typ {
	case "unix":
		execute(t, user, "login: ")
		execute(t, passwd, "ssword:")
		data, err = t.ReadBytes('$')
	case "juniper":
		execute(t, user, "login: ")
		execute(t, passwd, "ssword:")
		data, err = t.ReadBytes('>')
	case "cisco":
		execute(t, user, "name: ")
		execute(t, passwd, "ssword: ")
		data, err = t.ReadBytes('>')
	default:
		log.Fatalln("bad host type: " + typ)
	}
	Check(err)
	os.Stdout.Write(data)
	os.Stdout.WriteString("\n")
}

// HistoryBackup will backup what described in conf.json
//func (ti *TaskItem) HistoryBackup() {
//
//}
