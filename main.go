package main

import (
	"fmt"
	"gomask/proxy"
	"sync"
)

type Forwarder struct {
	name   string
	local  string
	remote string
}

var wg sync.WaitGroup

func main() {
	servers := []Forwarder{
		{"pgsql", "192.168.1.104:35432", "192.168.1.100:5432"},
		{"oracle", "192.168.1.104:31521", "192.168.1.116:1521"},
		{"db2", "192.168.1.104:30000", "192.168.1.180:50000"},
	}
	for _, v := range servers {
		fmt.Printf("%s: %s ==> %s...\n", v.name, v.local, v.remote)
		//fmt.Printf("#{v.name}: #{v.local} ==> #{v.remote}")
		go func(v Forwarder) {
			wg.Add(1)
			proxy.ForwarderStart(v.local, v.remote)
		}(v)
	}
	wg.Wait()
}
