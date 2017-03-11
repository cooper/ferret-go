package runtime

type Scope struct {
	*genericObject
}

func NewScope() *Scope {
	return &Scope{objectBase()}
}

func (s *Scope) Description(d *DescriptionOption) string {
	return "[ Scope ] " + s.genericObject.Description(d)
}
