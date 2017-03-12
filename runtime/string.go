package runtime

import "fmt"

type String struct {
	Value string
	*genericObject
}

func NewString(s string) *String {
	return &String{s, objectBase()}
}

func (s *String) Description(d *DescriptionOption) string {
	return fmt.Sprintf("%#v", s.Value)
}

func (s *String) String() string {
	return s.Description(nil)
}
