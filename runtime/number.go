package runtime

// import "reflect"
import "fmt"

type Number struct {
	IntValue   int64
	FloatValue float64
	*genericObject
}

func NewNumber() *Number {
	n := &Number{0, 0, objectBase()}
	n.genericObject.object = n
	return n
}

func NewNumberFloat(f float64) *Number {
	n := NewNumber()
	n.FloatValue = f
	return n
}

func NewNumberInt(i int64) *Number {
	n := NewNumber()
	n.IntValue = i
	return n
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
