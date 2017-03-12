package runtime

type Function struct {
	Name      string
	Code      FunctionCode
	signature *Signature
	*genericObject
}

type FunctionCode func(c Call)

type Call struct {
	Self  Object    // *self
	This  Object    // *this
	Topic Object    // $_
	Urgs  []Object  // unnamed arguments
	Args  Arguments // named arguments
	Scope *Scope    // function body scope
	Ret   *Return   // return object
	Func  *Function // the function itself
	// call scope? maybe we don't need it anymore
}

func NewFunction(name string, code FunctionCode) *Function {
	return &Function{name, code, new(Signature), objectBase()}
}

func (f *Function) Call(c Call) Object {

	// the function itself
	c.Func = f

	// map unnamed arguments
	if c.Urgs != nil && len(c.Urgs) != 0 {
		f.handleArguments(&c)
	}

	// return object
	if c.Ret == nil {
		c.Ret = NewReturn()
	}

	// call
	f.Code(c)

	return c.Ret.Return()
}

// map arguments to their respective names
func (f *Function) handleArguments(c *Call) {

	// we don't have any signature to work with
	if len(f.signature.Arguments) == 0 {
		return
	}

	// initialize the argument map if there isn't one
	if c.Args == nil {
		c.Args = make(map[string]Object)
	}

	// for each signature entry, map the next argument to it
	for i, e := range f.signature.Arguments {

		// there are no unnamed arguments left
		if len(c.Urgs)-1 < i {
			break
		}

		arg := c.Urgs[i]

		// there already is a named argument here
		if _, ok := c.Args[e.Name]; ok {
			continue
		}

		// this is a hungry argument
		if e.Hungry {
			// TODO: Flist with the remaining Urgs
			break
		}

		// otherwise this is OK; map it
		c.Args[e.Name] = arg
	}
}

// verify that the provided arguments satisfy the signature
func (f *Function) verifyArguments(c *Call) {

}

func (f *Function) Signature() *Signature {
	return f.signature
}

func (f *Function) Description(d *DescriptionOption) string {
	s := "Function"
	if f.Name != "" {
		s += " '" + f.Name + "'"
	}
	if sig := f.signature.DetailedString(); sig != "" {
		s += " { " + sig + " }"
	}
	return s
}

func (f *Function) String() string {
	return f.Description(nil)
}
