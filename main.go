package main

import (
	"fmt"
	"awesomeProject/logic"
)

func PrintAdd(n int){
	for i := 0; i < n;i++{
		if i % 2  == 1 {
			fmt.Printf("%d\n",i)
		}
	}
}


func main() {
	logic.AttachStart()

}