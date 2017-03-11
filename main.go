package main

import "./runtime"
import "fmt"

func main() {
	o1 := runtime.NewObject()
	o1.Set("hi", runtime.Fstring("hey"))
	o2 := runtime.NewObject()
	o2.Set("hi2", runtime.Fstring("hey there"))
	o2.AddParent(o1)
	o3 := runtime.NewObject()
	o3.Set("hey", runtime.Fstring("hi fam"))
	o3.Set("hayy", runtime.Fstring("hows it going"))
	o1.Set("hello", o3)
	runtime.MainContext.Set("ayy", o2)
	fmt.Println(runtime.MainContext.Description(nil))
}
