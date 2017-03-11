package runtime

var Undefined = &UndefinedObject{objectBase()}

type UndefinedObject struct {
	*genericObject // only for now
}

func (u *UndefinedObject) Description(d *DescriptionOption) string {
	return "undefined"
}
