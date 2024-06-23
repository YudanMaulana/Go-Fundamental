package main

import (
	"fmt"
)

func main() {
	// dengan var [length]type data{value}
	var arr1 = [3]int{1, 2, 3}
	// dengan Short Variable :=
	arr2 := [5]int{4, 5, 6, 7, 8}
	// Kita bisa tidak mendeklarasikan bilangan dengan "..."
	arr3 := [...]int{4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
