package runtime

// import "reflect"
import "fmt"

type Number struct {
	IntValue   int64
	FloatValue float64
	*genericObject
}

func (n *Number) AsFloat() float64 {
	if n.FloatValue != 0 {
		return n.FloatValue
	}
	return float64(n.IntValue)
}

func (n *Number) Description(d *DescriptionOption) string {
	return fmt.Sprintf("%v", n.AsFloat())
}

func (n *Number) String() string {
	return n.Description(nil)
}
