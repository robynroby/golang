package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Name struct{
	fname string
	lname string
};

func main() {
		slice := make([]Name, 0, 3)
		fmt.Println("Enter file name (same directory):")
		
		var fileName string
		fmt.Scan(&fileName)
		
		file, e := os.Open(fileName)
		if e != nil {
			fmt.Println("Error is = ",e)
		}
		
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
		
			s := strings.Split(scanner.Text(), " ")
			var namee Name
			namee.fname, namee.lname = s[0], s[1]
			slice = append(slice, namee)
		
		}
		
		file.Close()
		fmt.Println(" ")	
		
		for _, v := range slice {
			fmt.Println(v.fname, v.lname)
		}
		
}	
	

