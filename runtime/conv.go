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

func Fbool(i interface{}) Boolean {
	switch v := i.(type) {
	case Boolean:
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

func Fnum(i interface{}) *Number {
	switch v := i.(type) {
	case *Number:
		return v
	case float64:
		return &Number{FloatValue: v, genericObject: objectBase()}
	case float32:
		return &Number{FloatValue: float64(v), genericObject: objectBase()}
	case int:
		return &Number{IntValue: int64(v), genericObject: objectBase()}
	case int8:
		return &Number{IntValue: int64(v), genericObject: objectBase()}
	case int16:
		return &Number{IntValue: int64(v), genericObject: objectBase()}
	case int32:
		return &Number{IntValue: int64(v), genericObject: objectBase()}
	case int64:
		return &Number{IntValue: int64(v), genericObject: objectBase()}
	case uint:
		return &Number{IntValue: int64(v), genericObject: objectBase()}
	case uint8:
		return &Number{IntValue: int64(v), genericObject: objectBase()}
	case uint16:
		return &Number{IntValue: int64(v), genericObject: objectBase()}
	case uint32:
		return &Number{IntValue: int64(v), genericObject: objectBase()}
	case uint64:
		return &Number{IntValue: int64(v), genericObject: objectBase()}
	default:
		return nil
	}
}
