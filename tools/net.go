package tools

import (
	"net"
	"strings"
	"time"
)

func IsSameIPV4Host(s1, s2 string) bool {
	return strings.Split(s1, ":")[0] == strings.Split(s2, ":")[0]
}

func IsSameIPV4Port(s1, s2 string) bool {
	return strings.Split(s1, ":")[1] == strings.Split(s2, ":")[1]
}

func ScanPort() {

}

func isOccupiedPort(port uint16) error {
	var err error

	tcpAddress, err := net.ResolveTCPAddr("tcp4", ":"+string(port))
	if err != nil {
		return err
	}

	for i := 0; i < 3; i++ {
		listener, err := net.ListenTCP("tcp", tcpAddress)
		if err != nil {
			time.Sleep(time.Duration(100) * time.Microsecond)
			if i == 3 {
				return err
			}
			continue
		} else {
			_ = listener.Close()
			break
		}
	}
	return nil
}


