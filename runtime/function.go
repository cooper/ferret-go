package runtime

type Function struct {
	name string
	code FunctionCode
	*genericObject
}

type FunctionCode func(c Call)

func NewFunction(name string, code FunctionCode) *Function {
	return &Function{name, code, objectBase()}
}

func (f *Function) SignatureString() string {
	return ""
}

func (f *Function) DetailedSignatureString() string {
	return ""
}

func (f *Function) Description(d *DescriptionOption) string {
	s := "Function"
	if f.name != "" {
		s += " '" + f.name + "'"
	}
	if sig := f.DetailedSignatureString(); sig != "" {
		s += " { " + sig + " }"
	}
	return s
}
