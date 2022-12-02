package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Swap(mySlice []int, i int) {
	
	tmp := mySlice[i]
	mySlice[i] = mySlice[i+1]
	mySlice[i+1] = tmp
}


func BubbleSort(mySlice []int) {

	l := len(mySlice)
	for j:=l-1; j>0; j-- {
		for i:=0; i<j; i++ {
			if mySlice[i] > mySlice[i+1] {
			Swap (mySlice, i)
			}
		}
	}
}

func readSlice()([]int, error){
	mySlice := make([]int, 0)

	fmt.Println("Enter slice elements separated by space:") 
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n') 
	if err != nil {
		return mySlice, err
	}

	str = str[:len(str)-1]
	
	i:=0
	for _, s := range strings.Fields (str){
		num, err := strconv.Atoi(s)
		if err == nil {
			mySlice = append(mySlice, num)
			if i>=9 {
				break
			}
			i++
		}
	}
	return mySlice, nil
	
}

func main(){
	mySlice, err := readSlice()
	if err != nil {
		log.Fatal(err)
	}
	BubbleSort(mySlice)
	fmt.Println("Sorted slice:", mySlice)
}