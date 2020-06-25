package main

import (
	//"github.com/tomoyaueno29/LearnGoProject/greeting"
	"fmt"
	"github.com/tomoyaueno29/LearnGoProject/hello"
	//"os"
)

func main(){
	//fmt.Println(os.Args)
	//fmt.Println(greeting.Do())
	s := "tomoya"
	fmt.Println(hello.GetHello(s))
}
