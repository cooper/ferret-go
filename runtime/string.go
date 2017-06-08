package runtime

import "fmt"

var stringPrototype = NewPrototype("String")

var stringClass = bindCoreClass(ClassBinding{
	Name:      "String",
	Aliases:   []string{"Str"},
	Creator:   func() Object { return NewString("") },
	Prototype: stringPrototype,
	Methods: []FunctionBinding{
		{Name: "length", Code: _string_length, Prop: true},
	},
})

type String struct {
	Value string
	*genericObject
}

func NewString(s string) *String {
	str := &String{s, objectBase()}
	str.genericObject.object = str
	str.AddParent(stringPrototype)
	return str
}

func (s *String) Len() int {
	return len(s.Value)
}

func (s *String) Description(d *DescriptionOption) string {
	return fmt.Sprintf("%#v", s.Value)
}

func (s *String) String() string {
	return s.Description(nil)
}

func _string_length(c Call) {
	s := c.Self.(*String)
	c.Ret.Override(Fnum(s.Len()))
}
