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

func (sig Signature) AddArgument(e SignatureEntry) {
	sig.Arguments = append(sig.Arguments, e)
}

func (sig Signature) AddReturn(e SignatureEntry) {
	sig.Arguments = append(sig.Arguments, e)
}

func (e SignatureEntry) String() string {
	s := ""
	if e.Optional {
		s += "?"
	}
	s += "$" + e.Name
	if len(e.Types) != 0 {
		s += ": " + strings.Join(e.Types, "|")
	}
	if e.Hungry {
		s += "..."
	}
	return s
}

func (sig Signature) String() string {
	s := ""
	if len(sig.Arguments) != 0 {
		entries := make([]string, len(sig.Arguments))
		for i, e := range sig.Arguments {
			entries[i] = e.String()
		}
		s += strings.Join(entries, " ")
		if len(sig.Returns) != 0 {
			s += " "
		}
	}
	if len(sig.Returns) != 0 {
		s += "-> "
		entries := make([]string, len(sig.Returns))
		for i, e := range sig.Returns {
			entries[i] = e.String()
		}
		s += strings.Join(entries, " ")
	}
	return s
}
