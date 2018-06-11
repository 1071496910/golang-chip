package main

import (
	"fmt"
)

func main() {
	lru := NewLRU(3)
	lru.Add("1", "1")
	lru.Add("2", "2")
	lru.Add("3", "3")
	lru.Peek("1")
	lru.Add("4", "4")
	v1 := lru.Contains("1")
	fmt.Println(v1)
	v2 := lru.Contains("2")
	fmt.Println(v2)
	lru.Remove("2")
	v3 := lru.Contains("2")
	fmt.Println(v3)

}
