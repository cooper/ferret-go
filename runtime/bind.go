package runtime

type ClassBinding struct {
	Name        string            // class name
	Version     float32           // version
	Package     string            // ferret package name (top-level)
	Initializer FunctionBinding   // default init function
	Functions   []FunctionBinding // class functions
	Methods     []FunctionBinding // instance methods
	Aliases     []string          // alternative class names
	Creator     func() Object     // create a new instance
	Prototype   *Prototype
}

type FunctionBinding struct {
	Name string       // function name
	Code FunctionCode // code
	Need string       // required arguments
	Want string       // optional arguments
	Prop bool         // true for computed properties
	Lazy bool         // true for lazy-evaluated properties
}

func BindFunction(o Object, f FunctionBinding) *Event {

	// create event
	e := NewEventWithCode(f.Name, f.Code)

	// add wants and needs
	e.Default.signature.AddNeedString(f.Need)
	e.Default.signature.AddWantString(f.Want)

	// if it's a property, wrap it
	var v PropertyValue = e
	if f.Lazy {
		v = ComputedProperty{e, true}
	} else if f.Prop {
		v = ComputedProperty{e, false}
	}

	// store the event
	o.Set(f.Name, v)

	return e
}

func BindClass(o Object, c ClassBinding) *Class {
	class := NewClass(Class{
		Name:    c.Name,
		Version: c.Version,
		Creator: c.Creator,
	}, c.Prototype)

	// initializer
	if c.Initializer.Code != nil {
		c.Initializer.Name = "initializer__"
		BindFunction(class, c.Initializer)
	}

	// class functions
	for _, f := range c.Functions {
		BindFunction(class, f)
	}

	// instance methods
	proto := class.Proto()
	for _, f := range c.Methods {
		if proto == nil {
			panic("class " + c.Name + "has a nil prototype!")
		}
		BindFunction(proto, f)
	}

	// TODO: Package

	// store the class
	for _, name := range append(c.Aliases, c.Name) {
		o.Set(name, class)
	}

	return class
}

func bindCoreClass(c ClassBinding) *Class {
	return BindClass(MainContext, c)
}
