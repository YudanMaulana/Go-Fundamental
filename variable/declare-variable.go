package main

import (
	"fmt"
)

func main() {
	var coba1 string = "John" //type is string
	var coba2 = "Jane"        //type is inferred
	coba3 := 2                //type is inferred

	fmt.Println(coba1)
	fmt.Println(coba2)
	fmt.Println(coba3)
}
