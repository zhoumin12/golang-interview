package main

import "fmt"

func main() {
	map1 := make(map[int]int, 10)
	map1[12] = 12
	fmt.Println(len(map1))
}
