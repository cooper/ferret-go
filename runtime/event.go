package runtime

type Event struct {
	Name      string
	Default   *Function
	Functions []*Function
	*genericObject
}

func (e *Event) Signature() *Signature {
	return e.Default.Signature
}
