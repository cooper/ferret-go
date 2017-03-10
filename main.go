package main

import "./runtime"
import "fmt"

func main() {
	o := runtime.NewObject()
	o.Set("hi", runtime.Fstring("hey"))
    o.Get("hi").Description()
	fmt.Println(o)
}
