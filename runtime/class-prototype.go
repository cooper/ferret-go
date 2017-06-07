package runtime

import "fmt"

type Prototype struct {
	Name string
	*genericObject
}

func NewPrototype(name string) *Prototype {
	return &Prototype{name, objectBase()}
}

func (p *Prototype) Description(d *DescriptionOption) string {
	return fmt.Sprintf("[ Prototype '%s' ] %s", p.Name, p.genericObject.Description(d))
}
