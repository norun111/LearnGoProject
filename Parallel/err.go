package main

import (
	"fmt"
	"net/http"
)

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan *http.Response {
		response := make(chan *http.Response)
		go func() {
			defer close(response)
			for _, url := range urls{
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case response <- resp:
				}
			}
		}()
		return response
	}

	done := make(chan interface{})
	//この関数の最後にdoneをcloseしてfor文内にdoneを送信してシグナルを送る
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for response := range checkStatus(done, urls...){
		fmt.Printf("Response: %v\n", response.Status)
	}
}
