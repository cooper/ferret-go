package runtime

type Class struct {
	Name    string
	Version float32
	*genericObject
}

var classPrototype = NewPrototype("Class") // TODO

func NewClass(opts Class) *Class {
	c := &opts
    c.genericObject = objectBase()

    // make the class inherit from the global class prototype
    c.AddParent(classPrototype)

    // create this class's prototype
    c.Set("proto", NewPrototype(c.Name))

    // set name and version properties
    c.Set("name", c.Name)
    c.Set("version", c.Version)

	return c
}

func (c *Class) Proto() *Prototype {
    if p, ok := c.Get("proto").(*Prototype); ok {
        return p
    }
    return nil
}
