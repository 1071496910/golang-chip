package main

import (
	"fmt"
)

func main() {
	lru := NewLRU(2)
	lru.Add("1", "1")
	lru.Add("2", "2")
	lru.Add("3", "3")
	lru.Peek("1")
	{
		k, v, ok := lru.GetOldest()
		fmt.Println("2")
		fmt.Println(k, v, ok)
	}
	lru.Add("4", "4")
	{
		k, v, ok := lru.GetOldest()
		fmt.Println("3")
		fmt.Println(k, v, ok)
	}
	v1 := lru.Contains("1")
	fmt.Println(v1)
	v2 := lru.Contains("2")
	fmt.Println(v2)
	lru.Remove("2")
	v3 := lru.Contains("2")
	fmt.Println(v3)
	v4 := lru.Contains("4")
	fmt.Println(v4)
	lru.Remove("4")
	v5 := lru.Contains("4")
	fmt.Println(v5)
	lru.Add("5", "5")
	_, v6 := lru.Get("5")
	{
		k, v, ok := lru.GetOldest()
		fmt.Println("3")
		fmt.Println(k, v, ok)
	}
	fmt.Println(v6)

}
