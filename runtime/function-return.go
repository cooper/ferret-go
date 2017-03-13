package runtime

type Return struct {
	override Object
	*genericObject
}

func NewReturn() *Return {
	return &Return{nil, objectBase()}
}

func (r *Return) Override(obj Object) {
	r.override = obj
}

func (r *Return) Return() Object {
	if r.override != nil {
		return r.override
	}
	return r
}

func (r *Return) Description(d *DescriptionOption) string {
	return "[ Return ] " + r.genericObject.Description(d)
}

func (r *Return) Object() Object {
	return r
}
