package commons

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ziutek/telnet"
)

const timeout = 10 * time.Second

// telnetCommand will run telnet commands.
// t is obj of telnet.Conn, strSend is the command you send, strPrompt is the prompt string before you type the command
func telnetCommand(t *telnet.Conn, strSend string, strPrompt ...string) {
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

//func StdOut(t *telnet.Conn, bDelim byte) (*telnet.Conn, error) {
//	data, err := t.ReadBytes(bDelim)
//	Check(err)
//	os.Stdout.Write(data)
//	os.Stdout.WriteString("\n")
//	return t, err
//}

// Login is used to login Unix, Juniper or Cisco devices.
func (host *Host) HostLogin() (*telnet.Conn, error) {
	typ, dst, user, passwd := "juniper", net.JoinHostPort(host.IP, host.Port), host.Username, host.Password
	t, err := telnet.Dial("tcp", dst)
	if err != nil {
		fmt.Println("[-] "+host.IP+":"+host.Port + " cannot connect.")
		return nil, err
	}
	t.SetUnixWriteMode(true)
	//var data []byte
	switch typ {
	case "unix":
		telnetCommand(t, user, "login: ")
		telnetCommand(t, passwd, "ssword:")
		//t, err = StdOut(t, '$')
	case "juniper":
		telnetCommand(t, user, "login: ")
		telnetCommand(t, passwd, "ssword:")
		//t, err = StdOut(t, '>')
	case "cisco":
		telnetCommand(t, user, "name: ")
		telnetCommand(t, passwd, "ssword: ")
		//t, err = StdOut(t, '>')
	default:
		log.Fatalln("bad host type: " + typ)
	}
	fmt.Println("======== " + host.IP + " ==========")
	fmt.Println("[+] Connection established.")
	return t, err
}

func (conf *Conf) ConfHistoryBackup(t *telnet.Conn) {
	if t == nil{
		return
	}
	for _, cmd := range conf.HistoryBackup {
		telnetCommand(t, cmd[1:], cmd[:1])
	}
	fmt.Println("[+] History backup completed.")
}

func (conf *Conf) ConfGetConfig(t *telnet.Conn) {
	if t == nil{
		return
	}
	for _, cmd := range conf.GetConfig {
		telnetCommand(t, cmd[1:], cmd[:1])
	}
	fmt.Println("[+] Configuration save completed.")
}

func (conf *Conf) ConfHistoryRecover(t *telnet.Conn, host Host) {
	if t == nil{
		return
	}
	for _, cmd := range conf.HistoryRecover {
		telnetCommand(t, cmd[1:], cmd[:1])
		if cmd[1:] == "su" {
			telnetCommand(t, host.SuPassword, "ssword:")
			fmt.Println("[+] History Recover completed.")
		}
	}
}
