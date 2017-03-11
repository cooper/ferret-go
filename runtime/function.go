package runtime

type Function struct {
	Name string
	Code FunctionCode
	*genericObject
}

type FunctionCode func(c Call)

type Call struct {
	Self  Object     // *self
	This  Object     // *this
	Topic Object     // $_
	Urgs  []Object   // unnamed arguments
	Args  Arguments // named arguments
	Scope *Scope     // function body scope
	Ret   *Return    // return object
	Func  *Function  // the function itself
	// call scope? maybe we don't need it anymore
}

func NewFunction(name string, code FunctionCode) *Function {
	return &Function{name, code, objectBase()}
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

func (f *Function) handleArguments(c *Call) {

}

func (f *Function) SignatureString() string {
	return ""
}

func (f *Function) DetailedSignatureString() string {
	return ""
}

func (f *Function) Description(d *DescriptionOption) string {
	s := "Function"
	if f.Name != "" {
		s += " '" + f.Name + "'"
	}
	if sig := f.DetailedSignatureString(); sig != "" {
		s += " { " + sig + " }"
	}
	return s
}
