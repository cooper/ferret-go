package runtime

import "fmt"

type String struct {
	value string
	*genericObject
}

func NewString(s string) *String {
	return &String{s, objectBase()}
}

func (s *String) Description(d *DescriptionOption) string {
	return fmt.Sprintf("%#v", s.value)
}
