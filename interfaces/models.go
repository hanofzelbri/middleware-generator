package interfaces

import (
	"go/types"

	"golang.org/x/tools/go/loader"
)

// Options represents commandline arguments
type Options struct {
	Query                              string
	Wrapper                            string
	Output                             string
	MiddlewareFunctionName             string
	EmptyFunctionParamNamePrefix       string
	EmptyFunctionReturnParamNamePrefix string
}

// Config represents a named type request.
type Config struct {
	InterfaceName      string
	PackageName        string
	Program            *loader.Program
	Package            *types.Package
	Object             types.Object
	WrapperPackageName string
	WrapperStructName  string
	options            Options
}

// Interface represents an interface signature
type Interface struct {
	Name                   string
	Comment                string
	Functions              []Func
	Imports                []Import
	WrapperPackageName     string
	WrapperStructName      string
	MiddleWareFunctionName string
}

// Func represents a function signature
type Func struct {
	Name    string
	Params  []Param
	Res     []Param
	Comment string
	IsVariadic  bool
}

// Param represents a parameter in a function or method signature
type Param struct {
	Name string
	Type Type
}

// Type represents a simple representation of a single parameter type
type Type struct {
	Name        string
	Imports     []Import
}

// Import defines imported package
type Import struct {
	Package string
	Path string
}
