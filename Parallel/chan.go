package main

import (
	"log"
	"time"
)

func main() {
	log.Print("Start.")

	finished := make(chan bool)

	//func()型のスライスを定義
	funcs := []func(){
		func() {
		//1秒かかるコマンド
			log.Print("sleep1 started.")
			time.Sleep(1 * time.Second)
			log.Print("sleep1 finished.")
			finished <- true
		},
		func() {
		//2秒かかるコマンド
			log.Print("sleep2 started.")
			time.Sleep(2*time.Second)
			log.Print("sleep2 finished.")
			finished <- true
		},
		func() {
		// 3秒かかるコマンド
			log.Print("sleep3 started.")
			time.Sleep(3 * time.Second)
			log.Print("sleep3 finished.")
			finished <- true
		},
	}

	for _, sleep := range funcs {
		go sleep()
	}

	for i := 0; i <= len(funcs) - 1; i++ {
		<-finished
	}

	log.Print("all finished.")
}
