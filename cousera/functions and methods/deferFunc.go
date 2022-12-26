package main

import(
	"fmt"
)

func main(){
	defer fmt.Println("bye")
	fmt.Println("hello")
}