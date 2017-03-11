package runtime

type Scope struct {
	*genericObject
}

func (s *Scope) Description(d *DescriptionOption) string {
	return "[ Scope ] " + s.genericObject.Description(d)
}
