package runtime

import "fmt"

var coreClasses = make([]ClassBinding, 0)

var coreFunctions = []FunctionBinding{
	{Name: "say", Code: _say, Need: "$message"},
}

func addCoreClasses(ctx *Context) *Context {
	for _, c := range coreClasses {
		BindClass(ctx, c)
	}
	return ctx
}

func addCoreFunctions(c *Context) *Context {
	for _, f := range coreFunctions {
		BindFunction(c, f)
	}
	return c
}

func _say(c Call) {
	fmt.Println(c.Args["message"])
}
