package runtime

import "fmt"

type Prototype struct {
	Name string
	*genericObject
}

func NewPrototype(name string) *Prototype {
	p := &Prototype{name, objectBase()}
	p.genericObject.object = p
	return p
}

func (p *Prototype) Description(d *DescriptionOption) string {
	return fmt.Sprintf("[ Prototype '%s' ] %s", p.Name, p.genericObject.Description(d))
}
