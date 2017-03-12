package runtime

type Event struct {
	Name      string
	Default   *Function
	Functions []*Function
	*genericObject
}

func NewEvent(name string, defaultFunc *Function) *Event {
	return &Event{
		name,
		defaultFunc,
		[]*Function{defaultFunc},
		objectBase(),
	}
}

func NewEventWithCode(name string, code FunctionCode) *Event {
	defaultFunc := NewFunction("default", code)
	return NewEvent(name, defaultFunc)
}

func (e *Event) Signature() *Signature {
	return e.Default.Signature
}
