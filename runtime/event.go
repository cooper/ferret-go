package runtime

type Event struct {
	Name      string
	Default   *Function
	Functions []*Function
	Signature *Signature
	*genericObject
}
