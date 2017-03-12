package runtime

var coreClasses = []ClassBinding{
	stringClass,
}

func addCoreClasses(ctx *Context) *Context {
	for _, c := range coreClasses {
		BindClass(ctx, c)
	}
	return ctx
}
