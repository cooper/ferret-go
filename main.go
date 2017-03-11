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

	runtime.MainContext.Set("weakObject", runtime.NewObject())
	runtime.MainContext.Set("undef", runtime.Undefined)
	runtime.MainContext.Set("tru", runtime.True)
	runtime.MainContext.Set("fals", runtime.False)

	fmt.Println(runtime.MainContext.Description(nil))
	runtime.MainContext.Weaken("weakObject")
	fmt.Println(runtime.MainContext.Description(nil))
}
