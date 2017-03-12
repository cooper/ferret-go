package runtime

type ClassBinding struct {
	Name        string            // class name
	Version     float32           // version
	Package     string            // ferret package name (top-level)
	Initializer FunctionBinding   // default init function
	Functions   []FunctionBinding // class functions
	Methods     []FunctionBinding // instance methods
	Aliases     []string          // alternative class names
}

type FunctionBinding struct {
	Name string       // function name
	Code FunctionCode // code
	Need string       // required arguments
	Want string       // optional arguments
}
