package main

import f "./runtime"
import "fmt"

func main() {
	o1 := f.NewObject()
	o1.Set("hi", f.Fstring("hey"))
	o2 := f.NewObject()
	o2.Set("hi2", f.Fstring("hey there"))
	o2.AddParent(o1)
	o3 := f.NewObject()
	o3.Set("hey", "hi fam")
	o3.Set("hayy", "hows it going")
	o1.Set("hello", o3)
	f.MainContext.Set("ayy", o2)

	f.MainContext.Set("weakObject", f.NewObject())
	f.MainContext.Set("undef", f.Undefined)
	f.MainContext.Set("tru", true)
	f.MainContext.Set("fals", false)
	f.MainContext.Weaken("weakObject")

	f.MainContext.Get("say").Call(f.Call{
		Args: map[string]f.Object{"message": f.Fstring("Hello World!")},
		Urgs: []f.Object{f.Fstring("This should not override the named argument")},
	})

	f.MainContext.Get("say").Call(f.Call{
		Urgs: []f.Object{f.Fstring("unnamed arg")},
	})

	myFunc := f.NewEventWithCode("myFunc", func(c f.Call) {
		fmt.Println("called myFunc!")
	})
	f.MainContext.Set("myFunc", myFunc)
	f.MainContext.Get("myFunc").Call(f.Call{})

	f.MainContext.Set("createdString", f.MainContext.Get("String").Call(f.Call{}))
	f.MainContext.Set("myString", "testing")
	f.MainContext.Set("length", f.MainContext.Get("myString").Get("length"))
	f.MainContext.Set("ceil73", f.Fnum(7.3).Get("ceil"))
	fmt.Println(f.MainContext.String())
}
