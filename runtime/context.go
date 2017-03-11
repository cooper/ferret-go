package runtime

import "fmt"

var mainContext = NewContext("main", nil)

type Context struct {
	name string
	*Scope
}

func NewContext(name string, parent *Context) *Context {
	c := &Context{name, NewScope()}
	if parent != nil {
		c.AddParent(parent)
	}
	return c
}

func (c *Context) Description(d *DescriptionOption) string {
	return fmt.Sprintf("[ Context '%s' ] %s", c.name, c.genericObject.Description(d))
}
