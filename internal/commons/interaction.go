package commons

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/ziutek/telnet"
)

const timeout = 10 * time.Second

func expect(t *telnet.Conn, d ...string) {
	Check(t.SetReadDeadline(time.Now().Add(timeout)))
	Check(t.SkipUntil(d...))
}

func send(t *telnet.Conn, s string) {
	Check(t.SetWriteDeadline(time.Now().Add(timeout)))
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
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
		expect(t, "login: ")
		send(t, user)
		expect(t, "ssword: ")
		send(t, passwd)
		expect(t, "$")
		send(t, "ls -l")
		data, err = t.ReadBytes('$')
	case "juniper":
		expect(t, "login: ")
		send(t, user)
		expect(t, "ssword:")
		send(t, passwd)
		expect(t, ">")
		send(t, "sh ver")
		data, err = t.ReadBytes('>')
	case "cisco":
		expect(t, "name: ")
		send(t, user)
		expect(t, "ssword: ")
		send(t, passwd)
		expect(t, ">")
		send(t, "sh ver")
		data, err = t.ReadBytes('>')
	default:
		log.Fatalln("bad host type: " + typ)
	}
	Check(err)
	os.Stdout.Write(data)
	os.Stdout.WriteString("\n")
}
