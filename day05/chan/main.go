package main

import (
	"fmt"
	"sync"
	"time"
)

type Server struct {
	value int
	mutex sync.Mutex
	tune  chan int
}

func NewServer() *Server {
	ch := make(chan int)
	server := &Server{
		value: 100,
		tune:  ch,
	}

	return server
}

func (s *Server) Run() {
	var i int = 0
	for {
		fmt.Println("main thread", i)
		time.Sleep(time.Second * 3)
		i++
	}
}

func singleHandler(s *Server) {
	var i int = 0
	for {
		fmt.Printf("singleHandler thread : %d\n", i)
		time.Sleep(time.Second * 2)
		i++
		if i == 5 {
			fmt.Println("The main thead will been stopped by chan")
			s.tune <- 1
		}
	}
}

func main() {
	server := NewServer()

	go singleHandler(server) // run in child thread

	go server.Run()  // run in child thread

	//the main thread will wait here to receive content from tune
	select {
	case <-server.tune:
		break
	}
	fmt.Println("main thread end")
}
