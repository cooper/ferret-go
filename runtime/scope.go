package runtime

type Scope struct {
	*genericObject
}

func NewScope(parent Object) *Scope {
	s := &Scope{objectBase()}
	if parent != nil {
		s.AddParent(parent)
	}
	return s
}

// return a new scope which inherits from the same scopes but only has the
// given own properties
func (s *Scope) With(propNames ...string) *Scope {
	newScope := NewScope(nil)

	// add all the same parents
	newScope.isa = s.isa

	// only add the own properties that are used
	for _, propName := range propNames {
		if s.HasOwn(propName) {
			_, val := s.Property(propName)
			newScope.Set(propName, val)
		}
	}

	return newScope
}

func (s *Scope) Description(d *DescriptionOption) string {
	return "[ Scope ] " + s.genericObject.Description(d)
}
