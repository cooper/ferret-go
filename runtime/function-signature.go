package runtime

import "strings"

type SignatureEntry struct {
	Name     string   // argument name
	Types    []string // acceptable type names
	Optional bool     // true for wants
	Hungry   bool     // true if this is an ellipsis argument
}

type Signature struct {
	Arguments []SignatureEntry
	Returns   []SignatureEntry
}

func (sig *Signature) AddArgument(e SignatureEntry) {
	sig.Arguments = append(sig.Arguments, e)
}

func (sig *Signature) AddReturn(e SignatureEntry) {
	sig.Arguments = append(sig.Arguments, e)
}

func (e SignatureEntry) String() string {
	s := ""
	if e.Optional {
		s += "?"
	}
	s += "$" + e.Name
	if len(e.Types) != 0 {
		s += ":" + strings.Join(e.Types, "|")
	}
	if e.Hungry {
		s += "..."
	}
	return s
}

func (sig *Signature) String() string {
	if len(sig.Arguments) == 0 {
		return ""
	}
	s := ""
	entries := make([]string, len(sig.Arguments))
	for i, e := range sig.Arguments {
		entries[i] = e.String()
	}
	s += strings.Join(entries, " ")
	return s
}

func (sig *Signature) ReturnString() string {
	if len(sig.Returns) == 0 {
		return ""
	}
	s := "-> "
	entries := make([]string, len(sig.Returns))
	for i, e := range sig.Returns {
		entries[i] = e.String()
	}
	s += strings.Join(entries, " ")
	return s
}

func (sig *Signature) DetailedString() string {
	s := sig.String()
	r := sig.ReturnString()
	if len(s) != 0 && len(r) != 0 {
		return s + " " + r
	} else if len(r) != 0 {
		return r
	}
	return s
}
