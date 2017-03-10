package main

import "./runtime"
import "fmt"

func main() {
	o1 := runtime.NewObject()
	o1.Set("hi", runtime.Fstring("hey"))
	o2 := runtime.NewObject()
	o2.Set("hi2", runtime.Fstring("hey there"))
	o2.AddParent(o1)
	fmt.Println(o2.Description())
}
