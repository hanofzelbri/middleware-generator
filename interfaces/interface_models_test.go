package interfaces

var ReaderInterface = &Interface{
	Name:    "io.Reader",
	Comment: "// Reader is the interface that wraps the basic Read method.\n//\n// Read reads up to len(p) bytes into p. It returns the number of bytes\n// read (0 <= n <= len(p)) and any error encountered. Even if Read\n// returns n < len(p), it may use all of p as scratch space during the call.\n// If some data is available but not len(p) bytes, Read conventionally\n// returns what is available instead of waiting for more.\n//\n// When Read encounters an error or end-of-file condition after\n// successfully reading n > 0 bytes, it returns the number of\n// bytes read. It may return the (non-nil) error from the same call\n// or return the error (and n == 0) from a subsequent call.\n// An instance of this general case is that a Reader returning\n// a non-zero number of bytes at the end of the input stream may\n// return either err == EOF or err == nil. The next Read should\n// return 0, EOF.\n//\n// Callers should always process the n > 0 bytes returned before\n// considering the error err. Doing so correctly handles I/O errors\n// that happen after reading some bytes and also both of the\n// allowed EOF behaviors.\n//\n// Implementations of Read are discouraged from returning a\n// zero byte count with a nil error, except when len(p) == 0.\n// Callers should treat a return of 0 and nil as indicating that\n// nothing happened; in particular it does not indicate EOF.\n//\n// Implementations must not retain p.\n",
	Functions: []Func{
		{
			Name: "Read",
			Params: []Param{
				{
					Name: "p",
					Type: Type{
						Name:    "[]byte",
						Imports: nil,
					},
				},
			},
			Res: []Param{
				{
					Name: "n",
					Type: Type{
						Name:    "int",
						Imports: nil,
					},
				},
				{
					Name: "err",
					Type: Type{
						Name:    "error",
						Imports: nil,
					},
				},
			},
			Comment:    "",
			IsVariadic: false,
		},
	},
	Imports:                []Import{{Package: "io", Path: "io"}},
	WrapperPackageName:     "tests",
	WrapperStructName:      "Wrapper",
	MiddleWareFunctionName: "WithWrapper",
}
var TestInterface1Interface = &Interface{
	Name:    "TestInterface1",
	Comment: "// TestInterface1 is a dummy interface to test the program output.\n// This interface tests //-style method comments.\n/* Test comment */\n",
	Functions: []Func{
		{
			Name: "Method1",
			Params: []Param{
				{
					Name: "arg1",
					Type: Type{
						Name:    "string",
						Imports: nil,
					},
				},
				{
					Name: "arg2",
					Type: Type{
						Name:    "string",
						Imports: nil,
					},
				},
			},
			Res: []Param{
				{
					Name: "result",
					Type: Type{
						Name:    "string",
						Imports: nil,
					},
				},
				{
					Name: "err",
					Type: Type{
						Name:    "error",
						Imports: nil,
					},
				},
			},
			Comment:    "// Method1 is the first method of TestInterface1.\n",
			IsVariadic: false,
		},
		{
			Name: "Method2",
			Params: []Param{
				{
					Name: "arg1",
					Type: Type{
						Name:    "int",
						Imports: nil,
					},
				},
				{
					Name: "arg2",
					Type: Type{
						Name:    "int",
						Imports: nil,
					},
				},
			},
			Res: []Param{
				{
					Name: "result",
					Type: Type{
						Name:    "int",
						Imports: nil,
					},
				},
				{
					Name: "err",
					Type: Type{
						Name:    "error",
						Imports: nil,
					},
				},
			},
			Comment:    "// Method2 is the second method of TestInterface1.\n",
			IsVariadic: false,
		},
		{
			Name: "Method3",
			Params: []Param{
				{
					Name: "arg1",
					Type: Type{
						Name:    "bool",
						Imports: nil,
					},
				},
				{
					Name: "arg2",
					Type: Type{
						Name:    "bool",
						Imports: nil,
					},
				},
			},
			Res: []Param{
				{
					Name: "result",
					Type: Type{
						Name:    "bool",
						Imports: nil,
					},
				},
				{
					Name: "err",
					Type: Type{
						Name:    "error",
						Imports: nil,
					},
				},
			},
			Comment:    "/* Method3 is the third method of TestInterface1.\n\tContinue comment for method */\n",
			IsVariadic: false,
		},
	},
	Imports:                nil,
	WrapperPackageName:     "interfaces",
	WrapperStructName:      "testInterface1",
	MiddleWareFunctionName: "WithWrapper",
}
var EmptyInterfaceInterface = &Interface{
	Name:                   "interfaces.EmptyInterface",
	Comment:                "// EmptyInterface is a dummy interface to test program\n",
	Functions:              []Func{},
	Imports:                []Import{{Package: "interfaces", Path: "github.com/hanofzelbri/middleware-generator/interfaces"}},
	WrapperPackageName:     "tests",
	WrapperStructName:      "Wrapper",
	MiddleWareFunctionName: "WithWrapper",
}
var UnnammedParametersInterfaceInterface = &Interface{
	Name:    "UnnammedParametersInterface",
	Comment: "// UnnammedParameters is a dummy interface to test program\n",
	Functions: []Func{
		{
			Name:       "EmptyMethod",
			Params:     []Param{},
			Res:        []Param{},
			Comment:    "",
			IsVariadic: false,
		},
		{
			Name: "UnnammedParameter",
			Params: []Param{
				{
					Name: "paramName1",
					Type: Type{
						Name:    "string",
						Imports: nil,
					},
				},
			},
			Res: []Param{
				{
					Name: "returnName1",
					Type: Type{
						Name:    "error",
						Imports: nil,
					},
				},
			},
			Comment:    "",
			IsVariadic: false,
		},
		{
			Name: "UnnammedParameters",
			Params: []Param{
				{
					Name: "paramName1",
					Type: Type{
						Name:    "string",
						Imports: nil,
					},
				},
				{
					Name: "paramName2",
					Type: Type{
						Name:    "int",
						Imports: nil,
					},
				},
				{
					Name: "paramName3",
					Type: Type{
						Name:    "int",
						Imports: nil,
					},
				},
				{
					Name: "paramName4",
					Type: Type{
						Name:    "bool",
						Imports: nil,
					},
				},
			},
			Res: []Param{
				{
					Name: "returnName1",
					Type: Type{
						Name:    "bool",
						Imports: nil,
					},
				},
				{
					Name: "returnName2",
					Type: Type{
						Name:    "string",
						Imports: nil,
					},
				},
				{
					Name: "returnName3",
					Type: Type{
						Name:    "int",
						Imports: nil,
					},
				},
				{
					Name: "returnName4",
					Type: Type{
						Name:    "error",
						Imports: nil,
					},
				},
			},
			Comment:    "",
			IsVariadic: false,
		},
		{
			Name: "WithoutReturn",
			Params: []Param{
				{
					Name: "paramName1",
					Type: Type{
						Name:    "string",
						Imports: nil,
					},
				},
			},
			Res:        []Param{},
			Comment:    "",
			IsVariadic: false,
		},
	},
	Imports:                nil,
	WrapperPackageName:     "interfaces",
	WrapperStructName:      "unnammedParametersInterface",
	MiddleWareFunctionName: "WithWrapper",
}
var ImportedParamTypeInterfaceInterface = &Interface{
	Name:    "ImportedParamTypeInterface",
	Comment: "// ImportedParamTypeInterface is a dummy interface to test program\n/* Test comment */\n",
	Functions: []Func{
		{
			Name: "MultipleParamsWithSameType",
			Params: []Param{
				{
					Name: "typ1",
					Type: Type{
						Name: "*ast.TypeSpec",
						Imports: []Import{
							{Package: "ast", Path: "go/ast"},
						},
					},
				},
				{
					Name: "typ2",
					Type: Type{
						Name: "*ast.TypeSpec",
						Imports: []Import{
							{Package: "ast", Path: "go/ast"},
						},
					},
				},
				{
					Name: "uuid1",
					Type: Type{
						Name: "uuid.UUID",
						Imports: []Import{
							{Package: "uuid", Path: "github.com/google/uuid"},
						},
					},
				},
				{
					Name: "uuid2",
					Type: Type{
						Name: "uuid.UUID",
						Imports: []Import{
							{Package: "uuid", Path: "github.com/google/uuid"},
						},
					},
				},
			},
			Res: []Param{
				{
					Name: "ret1",
					Type: Type{
						Name: "ast.InterfaceType",
						Imports: []Import{
							{Package: "ast", Path: "go/ast"},
						},
					},
				},
				{
					Name: "ret2",
					Type: Type{
						Name: "ast.InterfaceType",
						Imports: []Import{
							{Package: "ast", Path: "go/ast"},
						},
					},
				},
			},
			Comment:    "// Multiple params with same type\n",
			IsVariadic: false,
		},
		{
			Name: "PointerTypeParam",
			Params: []Param{
				{
					Name: "typ1",
					Type: Type{
						Name: "*ast.TypeSpec",
						Imports: []Import{
							{Package: "ast", Path: "go/ast"},
						},
					},
				},
			},
			Res: []Param{
				{
					Name: "returnName1",
					Type: Type{
						Name: "*ast.InterfaceType",
						Imports: []Import{
							{Package: "ast", Path: "go/ast"},
						},
					},
				},
			},
			Comment:    "// Pointer type param\n",
			IsVariadic: false,
		},
		{
			Name:   "WithoutParameter",
			Params: []Param{},
			Res: []Param{
				{
					Name: "returnName1",
					Type: Type{
						Name: "*ast.InterfaceType",
						Imports: []Import{
							{Package: "ast", Path: "go/ast"},
						},
					},
				},
			},
			Comment:    "",
			IsVariadic: false,
		},
		{
			Name: "WithoutReturnParameter",
			Params: []Param{
				{
					Name: "paramName1",
					Type: Type{
						Name: "*ast.InterfaceType",
						Imports: []Import{
							{Package: "ast", Path: "go/ast"},
						},
					},
				},
			},
			Res:        []Param{},
			Comment:    "",
			IsVariadic: false,
		},
	},
	Imports: []Import{
		{Package: "uuid", Path: "github.com/google/uuid"},
		{Package: "ast", Path: "go/ast"},
	},
	WrapperPackageName:     "interfaces",
	WrapperStructName:      "importedParamTypeInterface",
	MiddleWareFunctionName: "WithWrapper",
}
var VariadicParamTypeInterfaceInterface = &Interface{
	Name:    "VariadicParamTypeInterface",
	Comment: "// VariadicParamTypeInterface is a dummy interface to test program\n",
	Functions: []Func{
		{
			Name: "VariadicFunction",
			Params: []Param{
				{
					Name: "prefix",
					Type: Type{
						Name:    "string",
						Imports: nil,
					},
				},
				{
					Name: "values",
					Type: Type{
						Name:    "...int",
						Imports: nil,
					},
				},
			},
			Res:        []Param{},
			Comment:    "// Variadic param type\n",
			IsVariadic: true,
		},
		{
			Name: "VariadicPointerFunction",
			Params: []Param{
				{
					Name: "values",
					Type: Type{
						Name:    "...*int",
						Imports: nil,
					},
				},
			},
			Res:        []Param{},
			Comment:    "// Variadic param type\n",
			IsVariadic: true,
		},
	},
	Imports:                nil,
	WrapperPackageName:     "interfaces",
	WrapperStructName:      "variadicParamTypeInterface",
	MiddleWareFunctionName: "WithWrapper",
}
var FuncTypeParamsInterfaceInterface = &Interface{
	Name:    "FuncTypeParamsInterface",
	Comment: "// FuncTypeParamsInterface is a dummy interface to test program\n",
	Functions: []Func{
		{
			Name: "FuncTypeParams",
			Params: []Param{
				{
					Name: "f",
					Type: Type{
						Name: "func(int, *ast.MapType) int",
						Imports: []Import{
							{Package: "ast", Path: "go/ast"},
						},
					},
				},
				{
					Name: "a",
					Type: Type{
						Name:    "int",
						Imports: nil,
					},
				},
				{
					Name: "b",
					Type: Type{
						Name:    "int",
						Imports: nil,
					},
				},
			},
			Res: []Param{
				{
					Name: "returnName1",
					Type: Type{
						Name: "func(uuid.UUID) error",
						Imports: []Import{
							{Package: "uuid", Path: "github.com/google/uuid"},
						},
					},
				},
			},
			Comment:    "// Function param type\n",
			IsVariadic: false,
		},
	},
	Imports: []Import{
		{Package: "uuid", Path: "github.com/google/uuid"},
		{Package: "ast", Path: "go/ast"},
	},
	WrapperPackageName:     "interfaces",
	WrapperStructName:      "funcTypeParamsInterface",
	MiddleWareFunctionName: "WithWrapper",
}

var CompositeParamsInterfaceInterface = &Interface{
	Name:    "CompositeParamsInterface",
	Comment: "// CompositeParamsInterface is a dummy interface to test program\n",
	Functions: []Func{
		{
			Name: "Array",
			Params: []Param{
				{
					Name: "a",
					Type: Type{
						Name: "[3]uuid.UUID",
						Imports: []Import{
							{Package: "uuid", Path: "github.com/google/uuid"},
						},
					},
				},
			},
			Res: []Param{
				{
					Name: "r",
					Type: Type{
						Name:    "[10]bool",
						Imports: nil,
					},
				},
			},
			Comment:    "// Array param types\n",
			IsVariadic: false,
		},
		{
			Name: "Channel",
			Params: []Param{
				{
					Name: "paramName1",
					Type: Type{
						Name:    "chan string",
						Imports: nil,
					},
				},
				{
					Name: "paramName2",
					Type: Type{
						Name:    "<-chan bool",
						Imports: nil,
					},
				},
				{
					Name: "paramName3",
					Type: Type{
						Name:    "chan<- int",
						Imports: nil,
					},
				},
			},
			Res: []Param{
				{
					Name: "returnName1",
					Type: Type{
						Name:    "chan int",
						Imports: nil,
					},
				},
			},
			Comment:    "// Channel param types\n",
			IsVariadic: false,
		},
		{
			Name: "Composite",
			Params: []Param{
				{
					Name: "m",
					Type: Type{
						Name:    "map[string]chan int",
						Imports: nil,
					},
				},
				{
					Name: "d",
					Type: Type{
						Name: "[2]chan func(string) map[bool]*ast.MapType",
						Imports: []Import{
							{Package: "ast", Path: "go/ast"},
						},
					},
				},
			},
			Res: []Param{
				{
					Name: "returnName1",
					Type: Type{
						Name:    "[]chan func(string) error",
						Imports: nil,
					},
				},
			},
			Comment:    "// Composite param types\n",
			IsVariadic: false,
		},
		{
			Name: "Map",
			Params: []Param{
				{
					Name: "paramName1",
					Type: Type{
						Name: "map[string]uuid.UUID",
						Imports: []Import{
							{Package: "uuid", Path: "github.com/google/uuid"},
						},
					},
				},
			},
			Res: []Param{
				{
					Name: "returnName1",
					Type: Type{
						Name:    "map[bool]int",
						Imports: nil,
					},
				},
			},
			Comment:    "// Map param types\n",
			IsVariadic: false,
		},
		{
			Name: "Slice",
			Params: []Param{
				{
					Name: "paramName1",
					Type: Type{
						Name: "[]uuid.UUID",
						Imports: []Import{
							{Package: "uuid", Path: "github.com/google/uuid"},
						},
					},
				},
				{
					Name: "paramName2",
					Type: Type{
						Name:    "[]int",
						Imports: nil,
					},
				},
			},
			Res: []Param{
				{
					Name: "returnName1",
					Type: Type{
						Name:    "[]bool",
						Imports: nil,
					},
				},
			},
			Comment:    "// Slice param types\n",
			IsVariadic: false,
		},
	},
	Imports: []Import{
		{Package: "uuid", Path: "github.com/google/uuid"},
		{Package: "ast", Path: "go/ast"},
	},
	WrapperPackageName:     "interfaces",
	WrapperStructName:      "compositeParamsInterface",
	MiddleWareFunctionName: "WithWrapper",
}
