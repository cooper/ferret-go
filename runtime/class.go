package runtime

type Class struct {
	Name    string
	Version float32
	Creator func() Object
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

func (c *Class) Initializer() *Event {
	if e, ok := c.Get("initializer__").(*Event); ok {
		return e
	}
	return nil
}

func (c *Class) Signature() *Signature {
	if c.Initializer() == nil {
		return &Signature{}
	}
	return c.Initializer().Signature()
}

func (class *Class) Call(c Call) Object {

	// if we have a creator, use it
	if class.Creator != nil {
		return class.Creator()
	}

	// otherwise, make a new object, then initialize it
	obj := NewObject()
	// TODO: initialize
	return obj
}

func (c *Class) Description(d *DescriptionOption) string {
	s := "Class"
	if c.Name != "" {
		s += " '" + c.Name + "'"
	}
	if sig := c.Signature().DetailedString(); sig != "" {
		s += " { " + sig + " }"
	}
	return s
}

func (c *Class) String() string {
	return c.Description(nil)
}
