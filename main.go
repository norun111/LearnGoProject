package main

import (
	"github.com/tomoyaueno29/LearnGoProject/greeting"

	"os"

	"fmt"

)

func main(){
	fmt.Println(os.Args)
	fmt.Println(greeting.Do())
}
