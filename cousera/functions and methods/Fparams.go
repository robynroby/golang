package main

import(
	"fmt"
)

func foo(x int, y int){
	fmt.Println(x * y)
}

func foo2 (x int) (int, int){
	return x, x+1
}

func main(){
	foo(2,3)
	a, b:= foo2(3)
	fmt.Println(a,b)
}