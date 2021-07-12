package main

import (
	"fmt"
	"gomask/transport"
	"sync"
)

type server struct {
	name string
	loc  string
	rem  string
}

var wg sync.WaitGroup

func main() {
	servers := []server{
		{"oracle", "127.0.0.1:1521", "172.17.2.129:1521"},
		{"pgsql", "127.0.0.1:5432", "172.17.2.129:5432"},
		{"mariadb", "127.0.0.1:3306", "172.17.2.129:3306"},
		{"db2", "127.0.0.1:20000", "192.168.1.180:50000"},
	}

	for _, s := range servers {
		fmt.Printf("%s: %s ==> %s...\n", s.name, s.loc, s.rem)
		go func(s server) {
			wg.Add(1)
			transport.Forward(s.loc, s.rem)
		}(s)
	}
	wg.Wait()
}
