package runtime

type Function struct {
	Name string
	Code FunctionCode
	*genericObject
}

type FunctionCode func(c Call)

type Call struct {
	Scope *Scope
	Ret   *Return
}

func NewFunction(name string, code FunctionCode) *Function {
	return &Function{name, code, objectBase()}
}

func (f *Function) Call(c Call) Object {
	return c.Ret.Return()
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
