package runtime

import "github.com/cooper/ferret-go/utils"

type Event struct {
	Name      string
	Default   *Function
	Functions []*Function
	*genericObject
}

func NewEvent(name string, defaultFunc *Function) *Event {
	e := &Event{
		name,
		defaultFunc,
		[]*Function{defaultFunc},
		objectBase(),
	}
	e.genericObject.object = e
	return e
}

func NewEventWithCode(name string, code FunctionCode) *Event {
	defaultFunc := NewFunction("default", code)
	return NewEvent(name, defaultFunc)
}

// FIXME: this is fine for now
func (e *Event) Call(c Call) Object {
	c.Self = e.GetLastParent()
	return e.Default.Call(c)
}

func (e *Event) Signature() *utils.Signature {
	if e.Default == nil {
		return new(utils.Signature)
	}
	return e.Default.signature
}

func (e *Event) Description(d *DescriptionOption) string {
	s := "Event"
	if e.Name != "" {
		s += " '" + e.Name + "'"
	}
	if sig := e.Signature().DetailedString(); sig != "" {
		s += " { " + sig + " }"
	}
	return s
}
