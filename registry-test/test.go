package main

import _ "registry-test/tmpl/impl"
import "registry-test/tmpl"
import "fmt"

func main() {
	fmt.Println(tmpl.Parser("jetty", "hello"))
	fmt.Println(tmpl.Parser("kiev", "world"))
	fmt.Println(tmpl.Parser("task", "task"))
}
