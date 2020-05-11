package proxy

import (
	"fmt"
	"gomask/tools"
	"net"
	"os"
)

var BuffSize = 1024 * 1024

func ForwarderStart(local, remote string) {
	proxyAddr := fmt.Sprintf("%s", local)
	proxyListener, err := net.Listen("tcp", proxyAddr)
	if err != nil {
		fmt.Printf("Unable to listen on: %s, error: %s\n", proxyAddr, err.Error())
		os.Exit(1)
	}
	defer func() {
		err := proxyListener.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(1111)
	for {
		proxyConn, err := proxyListener.Accept()
		fmt.Println(2222)
		fmt.Println(proxyConn)
		if err != nil {
			fmt.Printf("Unable to accept a request, error: %s\n", err.Error())
			continue
		}
		fmt.Println(3333)
		buffer := make([]byte, BuffSize)
		fmt.Println("3334")
		fmt.Printf("%v\n", proxyConn)
		n, err := proxyConn.Read(buffer)

		fmt.Println(4444)
		fmt.Println(n)

		strBuff := string(buffer[:n])
		fmt.Printf("NewSSS:\n")
		fmt.Printf("%v\n", strBuff)
		fmt.Println()
		fmt.Printf("%v\n", []byte(strBuff))

		if err != nil {
			fmt.Printf("Unable to read from input, error: %s\n", err.Error())
			continue
		}

		targetAddr := fmt.Sprintf("%s", remote)
		targetConn, err := net.Dial("tcp", targetAddr)
		fmt.Println("bbb", targetConn)
		if err != nil {
			fmt.Printf("Unable to connect to: %s, error: %s\n", targetAddr, err.Error())
			_ = proxyConn.Close()
			continue
		}
		fmt.Println("AAA", buffer[:n])
		n, err = targetConn.Write(buffer[:n])
		if err != nil {
			fmt.Printf("Unable to write to output, error: %s\n", err.Error())
			_ = proxyConn.Close()
			_ = targetConn.Close()
			continue
		}

		go proxyRequest(proxyConn, targetConn)
		go proxyRequest(targetConn, proxyConn)
	}
}

func proxyRequest(r net.Conn, w net.Conn) {
	defer func() {
		re := r.Close()
		if re != nil {
			fmt.Println(re)
		}

		we := w.Close()
		if we != nil {
			fmt.Println(we)
		}
	}()

	var buffer = make([]byte, BuffSize)
	for {
		n, err := r.Read(buffer)

		if err != nil {
			fmt.Printf("Unable to read from input, error: %s\n", err.Error())
			break
		}
		buff := buffer[:n]

		// Sender
		if tools.IsSameIPV4Host(r.RemoteAddr().String(), r.LocalAddr().String()) {

			strBuff := string(buff)

			//var stream analysis.Stream
			//buff = stream.Distribute(buff)
			//fmt.Println("res:", buff)

			fmt.Println("--------------------vvv------------------")
			fmt.Println("New:")
			fmt.Printf("%v\n", strBuff)
			fmt.Println()
			fmt.Printf("%v\n", []byte(strBuff))
			fmt.Println("------------------^^^--------------------")
		} else {
			// Receiver
			fmt.Println("Receiver passed")
		}

		n, err = w.Write(buff)

		if err != nil {
			fmt.Printf("Unable to write to output, error: %s\n", err.Error())
			break
		}
	}
}
