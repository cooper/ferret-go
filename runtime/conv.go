package runtime

import "fmt"

func Fstring(i interface{}) *String {
	s := ""
	switch v := i.(type) {
	case *String:
		return v
	case string:
		s = v
	case nil:
		s = "(undefined)"
		// TODO: case Object, check for toString
	default:
		s = fmt.Sprintf("%v", v)
	}
	return NewString(s)
}

func Fbool(i interface{}) *Boolean {
	switch v := i.(type) {
	case *Boolean:
		return v
	case nil:
		return False
	case bool:
		if v {
			return True
		}
		return False
	default:
		return False
	}
}
