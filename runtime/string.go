package runtime

import "fmt"

type String struct {
	value string
	*genericObject
}

func (s *String) Description() string {
	return fmt.Sprintf("%#v", s.value)
}