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
		{"mysql", "192.168.1.104:3311", "192.168.1.100:3306"},
		{"mariadb", "192.168.1.104:3307", "192.168.1.100:3307"},
		{"hive", "192.168.1.104:10000", "192.168.1.171:10000"},
		{"dm", "192.168.1.104:5237", "192.168.1.116:5237"},
		{"sybase", "192.168.1.104:2638", "192.168.1.107:2638"},
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
