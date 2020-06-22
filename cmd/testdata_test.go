package cmd

import (
	"go/ast"

	"github.com/google/uuid"
)

// TestInterface1 is a dummy interface to test the program output.
// This interface tests //-style method comments.
/* Test comment */
type TestInterface1 interface {
	// Method1 is the first method of TestInterface1.
	Method1(arg1 string, arg2 string) (result string, err error)
	// Method2 is the second method of TestInterface1.
	Method2(arg1 int, arg2 int) (result int, err error)
	/* Method3 is the third method of TestInterface1.
	Continue comment for method */
	Method3(arg1 bool, arg2 bool) (result bool, err error)
}

// EmptyInterface is a dummy interface to test program
type EmptyInterface interface {
}

// UnnammedParameters is a dummy interface to test program
type UnnammedParametersInterface interface {
	UnnammedParameter(string) error
	UnnammedParameters(string, int, int, bool) (bool, string, int, error)
	WithoutReturn(string)
	EmptyMethod()
}

// ImportedParamTypeInterface is a dummy interface to test program
type ImportedParamTypeInterface interface {
	// Pointer type param
	PointerTypeParam(typ1 *ast.TypeSpec) *ast.InterfaceType
	// Multiple params with same type
	MultipleParamsWithSameType(typ1, typ2 *ast.TypeSpec, uuid1, uuid2 uuid.UUID) (ret1, ret2 *ast.InterfaceType)
}

// VariadicParamTypeInterface is a dummy interface to test program
type VariadicParamTypeInterface interface {
	// Variadic param type
	VariadicFunction(prefix string, values ...int)
	// Variadic param type
	VariadicPointerFunction(values ...*int)
}

// FuncTypeParamsInterface is a dummy interface to test program
type FuncTypeParamsInterface interface {
	// Function param type
	FuncTypeParams(f func(int, int) int, a, b int) func(string) error
}
