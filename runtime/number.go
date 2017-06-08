package runtime

// import "reflect"
import "fmt"
import "math"

var numberPrototype = NewPrototype("Number")

var numberClass = bindCoreClass(ClassBinding{
	Name:      "Number",
	Aliases:   []string{"Num"},
	Creator:   func() Object { return NewNumber() },
	Prototype: numberPrototype,
	Methods:   []FunctionBinding{
		{Name: "ceil", Code: _number_ceil, Prop: true},
	},
})


type Number struct {
	IntValue   int64
	FloatValue float64
	*genericObject
}

func NewNumber() *Number {
	n := &Number{0, 0, objectBase()}
	n.genericObject.object = n
	n.AddParent(numberPrototype)
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

func _number_ceil(c Call) {
	n := c.Self.(*Number)
	ceil := math.Ceil(n.AsFloat())
	c.Ret.Override(Fnum(ceil))
}
