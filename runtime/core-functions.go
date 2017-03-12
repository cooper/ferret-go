package runtime

import "fmt"

var coreFunctions = []FunctionBinding{
	{"say", _say, "$message", ""},
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
