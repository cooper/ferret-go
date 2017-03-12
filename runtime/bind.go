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
}

type FunctionBinding struct {
	Name string       // function name
	Code FunctionCode // code
	Need string       // required arguments
	Want string       // optional arguments
}

func BindFunction(o Object, f FunctionBinding) *Event {
	e := NewEventWithCode(f.Name, f.Code)
	e.Default.Signature.AddNeedString(f.Need)
	e.Default.Signature.AddWantString(f.Want)
	o.Set(f.Name, e)
	return e
}

func BindClass(o Object, c ClassBinding) *Class {
	class := NewClass(Class{Name: c.Name, Version: c.Version})

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
		BindFunction(proto, f)
	}

	// TODO: Package

	// store the class
	for _, name := range append(c.Aliases, c.Name) {
		o.Set(name, class)
	}

	return class
}
