package main

import r "github.com/cooper/ferret-go/runtime"
import "fmt"

func main() {
	o1 := r.NewObject()
	o1.Set("hi", r.Fstring("hey"))
	o2 := r.NewObject()
	o2.Set("hi2", r.Fstring("hey there"))
	o2.AddParent(o1)
	o3 := r.NewObject()
	o3.Set("hey", "hi fam")
	o3.Set("hayy", "hows it going")
	o1.Set("hello", o3)
	r.MainContext.Set("ayy", o2)

	r.MainContext.Set("weakObject", r.NewObject())
	r.MainContext.Set("undef", r.Undefined)
	r.MainContext.Set("tru", true)
	r.MainContext.Set("fals", false)
	r.MainContext.Weaken("weakObject")

	r.MainContext.Get("say").Call(r.Call{
		Args: map[string]r.Object{"message": r.Fstring("Hello World!")},
		Urgs: []r.Object{r.Fstring("This should not override the named argument")},
	})

	r.MainContext.Get("say").Call(r.Call{
		Urgs: []r.Object{r.Fstring("unnamed arg")},
	})

	myFunc := r.NewEventWithCode("myFunc", func(c r.Call) {
		fmt.Println("called myFunc!")
	})
	r.MainContext.Set("myFunc", myFunc)
	r.MainContext.Get("myFunc").Call(r.Call{})

	r.MainContext.Set("createdString", r.MainContext.Get("String").Call(r.Call{}))
	r.MainContext.Set("myString", "testing")
	r.MainContext.Set("length", r.MainContext.Get("myString").Get("length"))
	r.MainContext.Set("ceil73", r.Fnum(7.3).Get("ceil"))
	fmt.Println(r.MainContext.String())
}
