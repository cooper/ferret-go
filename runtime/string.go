package runtime

import "fmt"

var stringClass = ClassBinding{
	Name:      "String",
	Aliases:   []string{"Str"},
	Creator:   func() Object { return NewString("") },
	Prototype: stringPrototype,
}

var stringPrototype = NewPrototype("String")

type String struct {
	Value string
	*genericObject
}

func NewString(s string) *String {
	str := &String{s, objectBase()}
	str.AddParent(stringPrototype)
	return str
}

func (s *String) Description(d *DescriptionOption) string {
	return fmt.Sprintf("%#v", s.Value)
}

func (s *String) String() string {
	return s.Value
}
