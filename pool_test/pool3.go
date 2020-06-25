package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func connectToService() interface{} {
	time.Sleep(1*time.Second)
	return struct {}{}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server, _ := net.Listen("tcp", "localhost:8080")
		defer server.Close()

		wg.Done()

		for {
			conn, _ := server.Accept()
			connectToService()
			fmt.Println(conn, "")
			conn.Close()
		}
	}()
	return &wg
}

func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}
