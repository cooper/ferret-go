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
		return NewNumberFloat(v)
	case float32:
		return NewNumberFloat(float64(v))
	case int:
		return NewNumberInt(int64(v))
	case int8:
		return NewNumberInt(int64(v))
	case int16:
		return NewNumberInt(int64(v))
	case int32:
		return NewNumberInt(int64(v))
	case int64:
		return NewNumberInt(v)
	case uint:
		return NewNumberInt(int64(v))
	case uint8:
		return NewNumberInt(int64(v))
	case uint16:
		return NewNumberInt(int64(v))
	case uint32:
		return NewNumberInt(int64(v))
	case uint64:
		return NewNumberInt(int64(v))
	default:
		return nil
	}
}
