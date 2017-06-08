package instruct

import r "github.com/cooper/ferret-go/runtime"

type State struct {
	context *r.Context
	scope   *r.Scope
}

func NewState() *State {
	return &State{r.MainContext, r.MainContext.Scope}
}
