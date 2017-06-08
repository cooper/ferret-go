package runtime

type Class struct {
	Name    string
	Version float32
	Creator func() Object
	*genericObject
}

var classPrototype = NewPrototype("Class") // TODO

func NewClass(opts Class, proto *Prototype) *Class {
	c := &opts
	c.genericObject = objectBase()
	c.genericObject.object = c

	// make the class inherit from the global class prototype
	c.AddParent(classPrototype)

	// create this class's prototype
	if proto == nil {
		proto = NewPrototype(c.Name)
	}
	c.Set("proto", proto)

	// set name and version properties
	c.Set("name", c.Name)
	c.Set("version", c.Version)

	return c
}

func (c *Class) Proto() *Prototype {
	switch proto := c.Get("proto").(type) {
	case *Prototype:
		return proto
	case Object:
		proto = proto.Object()
		if p, ok := proto.Object().(*Prototype); ok {
			return p
		}
		return nil
	default:
		return nil
	}
}

func (c *Class) Initializer() *Event {
	if e, ok := c.Get("initializer__").(*Event); ok {
		return e
	}
	return nil
}

func (c *Class) Signature() *Signature {
	if init := c.Initializer(); init != nil {
		return init.Signature()
	}
	return new(Signature)
}

func (class *Class) Init(obj Object, c Call) *Return {

	// can't initialize as a builtin type
	if class.Creator != nil {
		panic("cannot initialize existing object as type " + class.Name)
	}

	// TODO: actually initialize

	return c.Ret
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
	// TODO: version
	if sig := c.Signature().DetailedString(); sig != "" {
		s += " { " + sig + " }"
	}
	return s
}
