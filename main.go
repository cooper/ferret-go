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
	runtime.MainContext.Weaken("weakObject")

	sayFunc := runtime.NewFunction("say", say)
	sayFunc.Signature.AddArgument(runtime.SignatureEntry{
		Name:  "message",
		Types: []string{"Str"},
	})
	runtime.MainContext.Set("say", sayFunc)
	sayFunc.Call(runtime.Call{
		Args: map[string]runtime.Object{"message": runtime.Fstring("Hello World!")},
		Urgs: []runtime.Object{runtime.Fstring("This should not override the named argument")},
	})

	fmt.Println(runtime.MainContext)
}

func say(c runtime.Call) {
	fmt.Println(c.Args["message"])
}
