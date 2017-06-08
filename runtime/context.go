package runtime

import "fmt"

var MainContext = addCoreClasses(addCoreFunctions(NewContext("main", nil)))

type Context struct {
	name string
	*Scope
}

func NewContext(name string, parent *Context) *Context {
	var s *Scope
	if parent == nil {
		s = NewScope(nil)
	} else {
		s = NewScope(parent)
	}
	c := &Context{name, s}
	c.genericObject.object = c
	return c
}

func (c *Context) Description(d *DescriptionOption) string {
	return fmt.Sprintf("[ Context '%s' ] %s", c.name, c.genericObject.Description(d))
}
