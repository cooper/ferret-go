package runtime

type Prototype struct {
	Name string
	*genericObject
}

func NewPrototype(name string) *Prototype {
	return &Prototype{name, objectBase()}
}
