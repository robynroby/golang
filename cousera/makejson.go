package main

import (
	"encoding/json"
	"fmt"
)

// type Person struct{
// 	name string
// 	address string
// }

func main() {
	var name string
	var address string

	fmt.Printf("Enter a name: ")
	fmt.Scan(&name)
	fmt.Printf("Enter an address: ")
	fmt.Scan(&address)

	m := make(map[string]string)
	m["name"] = name
	m["address"] = address

	u, _ := json.Marshal(m)
	 fmt.Println(string(u))
}
