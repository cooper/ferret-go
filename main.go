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
	o3.Set("hey", "hi fam")
	o3.Set("hayy", "hows it going")
	o1.Set("hello", o3)
	runtime.MainContext.Set("ayy", o2)

	runtime.MainContext.Set("weakObject", runtime.NewObject())
	runtime.MainContext.Set("undef", runtime.Undefined)
	runtime.MainContext.Set("tru", true)
	runtime.MainContext.Set("fals", false)
	runtime.MainContext.Weaken("weakObject")

	runtime.BindFunction(runtime.FunctionBinding{
		Name: "say",
		Code: func(c runtime.Call) {
			fmt.Println(c.Args["message"])
		},
		Need: "$message",
	})

	runtime.MainContext.Get("say").Call(runtime.Call{
		Args: map[string]runtime.Object{"message": runtime.Fstring("Hello World!")},
		Urgs: []runtime.Object{runtime.Fstring("This should not override the named argument")},
	})

	fmt.Println(runtime.MainContext)
}
