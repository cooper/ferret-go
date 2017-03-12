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

func (sig *Signature) AddNeedString(s string) {
	entries := stringToEntries(s, false)
	sig.Arguments = append(sig.Arguments, entries...)
}

func (sig *Signature) AddWantString(s string) {
	entries := stringToEntries(s, true)
	sig.Arguments = append(sig.Arguments, entries...)
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

func (sig *Signature) ArgumentString() string {
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
	s := sig.ArgumentString()
	r := sig.ReturnString()
	if len(s) != 0 && len(r) != 0 {
		return s + " " + r
	} else if len(r) != 0 {
		return r
	}
	return s
}

func stringToEntries(s string, optional bool) []SignatureEntry {
	parts := strings.Split(s, " ")
	entries := make([]SignatureEntry, len(parts))
	j := 0
	for _, part := range parts {

		// whitespace
		if len(part) == 0 {
			continue
		}

		// must start with $
		if !strings.HasPrefix(part, "$") {
			panic("bad string argument " + part)
		}
		part = strings.TrimPrefix(part, "$")

		// if it ends in ..., it's a hungry argument
		hungry := strings.HasSuffix(part, "...")
		if hungry {
			part = strings.TrimSuffix(part, "...")
		}

		e := SignatureEntry{
			Optional: optional,
			Hungry:   hungry,
		}

		// might have types
		nameType := strings.SplitN(part, ":", 2)
		e.Name = nameType[0]
		if len(nameType) == 2 {
			e.Types = strings.Split(nameType[1], "|")
		}

		entries[j] = e
		j++
	}
	return entries[:j]
}
