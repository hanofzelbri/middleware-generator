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
    InterfaceName      string          `json:"interfaceName,omitempty"`
    PackageName        string          `json:"packageName,omitempty"`
    Program            *loader.Program `json:"program,omitempty"`
    Package            *types.Package  `json:"package,omitempty"`
    Object             types.Object    `json:"object,omitempty"`
    WrapperPackageName string          `json:"wrapperPackageName,omitempty"`
    WrapperStructName  string          `json:"wrapperStructName,omitempty"`
    Options            Options         `json:"options,omitempty"`
}

// Interface represents an interface signature
type Interface struct {
    Name                   string   `json:"name,omitempty"`
    Comment                string   `json:"comment,omitempty"`
    Functions              []Func   `json:"functions,omitempty"`
    Imports                []Import `json:"imports,omitempty"`
    WrapperPackageName     string   `json:"wrapperPackageName,omitempty"`
    WrapperStructName      string   `json:"wrapperStructName,omitempty"`
    MiddleWareFunctionName string   `json:"middlewareFunctionName,omitempty"`
}

// Func represents a function signature
type Func struct {
    Name       string  `json:"name,omitempty"`
    Params     []Param `json:"params,omitempty"`
    Res        []Param `json:"res,omitempty"`
    Comment    string  `json:"comment,omitempty"`
    IsVariadic bool    `json:"isVariadic,omitempty"`
}

// Param represents a parameter in a function or method signature
type Param struct {
    Name string `json:"name,omitempty"`
    Type Type   `json:"type,omitempty"`
}

// Type represents a simple representation of a single parameter type
type Type struct {
    Name    string   `json:"name,omitempty"`
    Imports []Import `json:"imports,omitempty"`
}

// Import defines imported package
type Import struct {
    Package string `json:"package,omitempty"`
    Path    string `json:"path,omitempty"`
}
