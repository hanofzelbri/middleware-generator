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
	Imports                []Import
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
	Name      string
	Package   string
	IsPointer bool
}

func (t *Type) String() string {
	ret := ""

	if t.IsPointer {
		ret += "*"
	}

	if t.Package != "" {
		ret += t.Package + "."
	}

	ret += t.Name

	return ret
}

// Import defines imported package
type Import struct {
	Name string
	Path string
}
