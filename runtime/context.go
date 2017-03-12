package runtime

import "fmt"

var MainContext = NewContext("main", nil)

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
	return c
}

func (c *Context) Description(d *DescriptionOption) string {
	return fmt.Sprintf("[ Context '%s' ] %s", c.name, c.genericObject.Description(d))
}

func (c *Context) String() string {
	return c.Description(nil)
}
