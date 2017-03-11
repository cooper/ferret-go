package runtime

var True = &Boolean{true, objectBase()}
var False = &Boolean{false, objectBase()}

type Boolean struct {
	Value          bool
	*genericObject // only for now
}

func (b *Boolean) Description(d *DescriptionOption) string {
	if b.Value {
		return "true"
	}
	return "false"
}
