package main

import (
        "fmt"
        "sort"
        "strconv"
)

func main() {
        slices := make([]int, 3)
        var number string

        for i := 0; i < len(slices); i++ {
                fmt.Println(" Enter a number")
                fmt.Scanln(&number)
                if number == "X" {
                        break
                }

                newSlice, err := strconv.Atoi(number)
                if err != nil {
                        fmt.Println("Wrong input")
                        continue
                }

                
                slices = append(slices, newSlice)
                

                sort.Ints(slices)

                for _,v := range slices{
                        if v != 0 {
                                fmt.Printf("%v",v)
                        }
                }
                fmt.Println("")

        }

}
