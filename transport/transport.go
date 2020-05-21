package transport

import (
	"fmt"
	"gomask/tools"
	"log"
	"net"
)

var BuffSize = 1024 * 1024

func Forward(locAddr, remAddr string) {
	locListener, err := net.Listen("tcp", locAddr)
	if err != nil {
		log.Fatalf("unable to listen on: %s, error: %s\n", locAddr, err.Error())
	}
	defer func() {
		err = locListener.Close()
		if err != nil {
			fmt.Printf("%s closed faild\n", locListener)
		}
	}()

	for {
		locConn, err := locListener.Accept()
		if err != nil {
			fmt.Printf("unable to to accept a request, error: %s\n", err.Error())
		} else {
			// 校验是否能够连接
			fmt.Println("new connect:", locConn.RemoteAddr().String())
		}

		remConn, err := net.Dial("tcp", remAddr)
		if err != nil {
			fmt.Printf("unable to to connect %s, error: %s\n", remAddr, err.Error())
			continue
		}

		go interact(locConn, remConn)
		go interact(remConn, locConn)
	}
}

func interact(r, w net.Conn) {

	defer func() {
		rErr := r.Close()
		if rErr != nil {
			fmt.Println(rErr)
		}

		wErr := w.Close()
		if wErr != nil {
			fmt.Println(wErr)
		}
	}()

	var buffer = make([]byte, BuffSize)
	for {
		n, err := r.Read(buffer)
		if err != nil {
			fmt.Printf("unable to read from input, error: %s\n", err.Error())
			break
		}

		buff := buffer[:n]

		if tools.IsSameIPV4Host(r.RemoteAddr().String(), r.LocalAddr().String()) {

			strBuff := string(buff)

			//var stream analysis.Stream
			//buff = stream.Distribute(buff)
			//fmt.Println("res:", buff)

			fmt.Println("--------------------vvv------------------")
			fmt.Println("new:")
			fmt.Printf("%v\n", strBuff)
			fmt.Println()
			fmt.Printf("%v\n", []byte(strBuff))
			fmt.Println("------------------^^^--------------------")

		} else {
			// Receiver
			fmt.Println("receiver passed")
		}

		n, err = w.Write(buff)
		if err != nil {
			fmt.Printf("unable to write to output, error: %s\n", err.Error())
			break
		}
	}
}
