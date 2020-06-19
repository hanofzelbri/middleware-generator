package cmd

// Query represents a named type request.
type Query struct {
	InterfaceName string
	PackageName   string
}

// Interface represents an interface signature
type Interface struct {
	Name                   string
	Comment                string
	PackageName            string
	Functions              []Func
	StructName             string
	MiddleWareFunctionName string
}

// Func represents a function signature
type Func struct {
	Name       string
	Params     []Param
	Res        []Param
	Comment    string
	IsVariadic bool
}

// Param represents a parameter in a function or method signature
type Param struct {
	Name string
	Type Type
}

// Type represents a simple representation of a single parameter type
type Type struct {
	Name        string
	Package     string
	ImportPath  string
	IsPointer   bool
	IsComposite bool
	IsFunc      bool
}
