package interfaces

var IoReaderJSON = `{
  "name": "Reader",
  "comment": "// Reader is the interface that wraps the basic Read method.\n//\n// Read reads up to len(p) bytes into p. It returns the number of bytes\n// read (0 \u003c= n \u003c= len(p)) and any error encountered. Even if Read\n// returns n \u003c len(p), it may use all of p as scratch space during the call.\n// If some data is available but not len(p) bytes, Read conventionally\n// returns what is available instead of waiting for more.\n//\n// When Read encounters an error or end-of-file condition after\n// successfully reading n \u003e 0 bytes, it returns the number of\n// bytes read. It may return the (non-nil) error from the same call\n// or return the error (and n == 0) from a subsequent call.\n// An instance of this general case is that a Reader returning\n// a non-zero number of bytes at the end of the input stream may\n// return either err == EOF or err == nil. The next Read should\n// return 0, EOF.\n//\n// Callers should always process the n \u003e 0 bytes returned before\n// considering the error err. Doing so correctly handles I/O errors\n// that happen after reading some bytes and also both of the\n// allowed EOF behaviors.\n//\n// Implementations of Read are discouraged from returning a\n// zero byte count with a nil error, except when len(p) == 0.\n// Callers should treat a return of 0 and nil as indicating that\n// nothing happened; in particular it does not indicate EOF.\n//\n// Implementations must not retain p.\n",
  "functions": [
    {
      "name": "Read",
      "params": [
        {
          "name": "p",
          "type": {
            "name": "[]byte"
          }
        }
      ],
      "res": [
        {
          "name": "n",
          "type": {
            "name": "int"
          }
        },
        {
          "name": "err",
          "type": {
            "name": "error"
          }
        }
      ]
    }
  ],
  "wrapperPackageName": "tests",
  "wrapperStructName": "Wrapper",
  "middlewareFunctionName": "WithWrapper"
}`

var TestInterface1JSON = `{
  "name": "TestInterface1",
  "comment": "// TestInterface1 is a dummy interface to test the program output.\n// This interface tests //-style method comments.\n/* Test comment */\n",
  "functions": [
    {
      "name": "Method1",
      "params": [
        {
          "name": "arg1",
          "type": {
            "name": "string"
          }
        },
        {
          "name": "arg2",
          "type": {
            "name": "string"
          }
        }
      ],
      "res": [
        {
          "name": "result",
          "type": {
            "name": "string"
          }
        },
        {
          "name": "err",
          "type": {
            "name": "error"
          }
        }
      ],
      "comment": "// Method1 is the first method of TestInterface1.\n"
    },
    {
      "name": "Method2",
      "params": [
        {
          "name": "arg1",
          "type": {
            "name": "int"
          }
        },
        {
          "name": "arg2",
          "type": {
            "name": "int"
          }
        }
      ],
      "res": [
        {
          "name": "result",
          "type": {
            "name": "int"
          }
        },
        {
          "name": "err",
          "type": {
            "name": "error"
          }
        }
      ],
      "comment": "// Method2 is the second method of TestInterface1.\n"
    },
    {
      "name": "Method3",
      "params": [
        {
          "name": "arg1",
          "type": {
            "name": "bool"
          }
        },
        {
          "name": "arg2",
          "type": {
            "name": "bool"
          }
        }
      ],
      "res": [
        {
          "name": "result",
          "type": {
            "name": "bool"
          }
        },
        {
          "name": "err",
          "type": {
            "name": "error"
          }
        }
      ],
      "comment": "/* Method3 is the third method of TestInterface1.\n\tContinue comment for method */\n"
    }
  ],
  "wrapperPackageName": "tests",
  "wrapperStructName": "Wrapper",
  "middlewareFunctionName": "WithWrapper"
}`

var EmptyInterfaceJSON = `{
  "name": "EmptyInterface",
  "comment": "// EmptyInterface is a dummy interface to test program\n",
  "wrapperPackageName": "tests",
  "wrapperStructName": "Wrapper",
  "middlewareFunctionName": "WithWrapper"
}`

var UnnammedParametersInterfaceJSON = `{
  "name": "UnnammedParametersInterface",
  "comment": "// UnnammedParameters is a dummy interface to test program\n",
  "functions": [
    {
      "name": "EmptyMethod"
    },
    {
      "name": "UnnammedParameter",
      "params": [
        {
          "name": "paramName1",
          "type": {
            "name": "string"
          }
        }
      ],
      "res": [
        {
          "name": "returnName1",
          "type": {
            "name": "error"
          }
        }
      ]
    },
    {
      "name": "UnnammedParameters",
      "params": [
        {
          "name": "paramName1",
          "type": {
            "name": "string"
          }
        },
        {
          "name": "paramName2",
          "type": {
            "name": "int"
          }
        },
        {
          "name": "paramName3",
          "type": {
            "name": "int"
          }
        },
        {
          "name": "paramName4",
          "type": {
            "name": "bool"
          }
        }
      ],
      "res": [
        {
          "name": "returnName1",
          "type": {
            "name": "bool"
          }
        },
        {
          "name": "returnName2",
          "type": {
            "name": "string"
          }
        },
        {
          "name": "returnName3",
          "type": {
            "name": "int"
          }
        },
        {
          "name": "returnName4",
          "type": {
            "name": "error"
          }
        }
      ]
    },
    {
      "name": "WithoutReturn",
      "params": [
        {
          "name": "paramName1",
          "type": {
            "name": "string"
          }
        }
      ]
    }
  ],
  "wrapperPackageName": "tests",
  "wrapperStructName": "Wrapper",
  "middlewareFunctionName": "WithWrapper"
}`

var ImportedParamTypeInterfaceJSON = `{
  "name": "ImportedParamTypeInterface",
  "comment": "// ImportedParamTypeInterface is a dummy interface to test program\n/* Test comment */\n",
  "functions": [
    {
      "name": "MultipleParamsWithSameType",
      "params": [
        {
          "name": "typ1",
          "type": {
            "name": "*ast.TypeSpec",
            "imports": [
              {
                "package": "ast",
                "path": "go/ast"
              }
            ]
          }
        },
        {
          "name": "typ2",
          "type": {
            "name": "*ast.TypeSpec",
            "imports": [
              {
                "package": "ast",
                "path": "go/ast"
              }
            ]
          }
        },
        {
          "name": "uuid1",
          "type": {
            "name": "UUID",
            "imports": [
              {
                "package": "uuid",
                "path": "github.com/google/uuid"
              }
            ]
          }
        },
        {
          "name": "uuid2",
          "type": {
            "name": "UUID",
            "imports": [
              {
                "package": "uuid",
                "path": "github.com/google/uuid"
              }
            ]
          }
        }
      ],
      "res": [
        {
          "name": "ret1",
          "type": {
            "name": "*ast.InterfaceType",
            "imports": [
              {
                "package": "ast",
                "path": "go/ast"
              }
            ]
          }
        },
        {
          "name": "ret2",
          "type": {
            "name": "*ast.InterfaceType",
            "imports": [
              {
                "package": "ast",
                "path": "go/ast"
              }
            ]
          }
        }
      ],
      "comment": "// Multiple params with same type\n"
    },
    {
      "name": "PointerTypeParam",
      "params": [
        {
          "name": "typ1",
          "type": {
            "name": "*ast.TypeSpec",
            "imports": [
              {
                "package": "ast",
                "path": "go/ast"
              }
            ]
          }
        }
      ],
      "res": [
        {
          "name": "returnName1",
          "type": {
            "name": "*ast.InterfaceType",
            "imports": [
              {
                "package": "ast",
                "path": "go/ast"
              }
            ]
          }
        }
      ],
      "comment": "// Pointer type param\n"
    },
    {
      "name": "WithoutParameter",
      "res": [
        {
          "name": "returnName1",
          "type": {
            "name": "*ast.InterfaceType",
            "imports": [
              {
                "package": "ast",
                "path": "go/ast"
              }
            ]
          }
        }
      ]
    },
    {
      "name": "WithoutReturnParameter",
      "params": [
        {
          "name": "paramName1",
          "type": {
            "name": "*ast.InterfaceType",
            "imports": [
              {
                "package": "ast",
                "path": "go/ast"
              }
            ]
          }
        }
      ]
    }
  ],
  "imports": [
    {
      "package": "uuid",
      "path": "github.com/google/uuid"
    },
    {
      "package": "ast",
      "path": "go/ast"
    }
  ],
  "wrapperPackageName": "tests",
  "wrapperStructName": "Wrapper",
  "middlewareFunctionName": "WithWrapper"
}`

var VariadicParamTypeInterfaceJSON = `{
  "name": "VariadicParamTypeInterface",
  "comment": "// VariadicParamTypeInterface is a dummy interface to test program\n",
  "functions": [
    {
      "name": "VariadicFunction",
      "params": [
        {
          "name": "prefix",
          "type": {
            "name": "string"
          }
        },
        {
          "name": "values",
          "type": {
            "name": "...int"
          }
        }
      ],
      "comment": "// Variadic param type\n",
      "isVariadic": true
    },
    {
      "name": "VariadicPointerFunction",
      "params": [
        {
          "name": "values",
          "type": {
            "name": "...*int"
          }
        }
      ],
      "comment": "// Variadic param type\n",
      "isVariadic": true
    }
  ],
  "wrapperPackageName": "tests",
  "wrapperStructName": "Wrapper",
  "middlewareFunctionName": "WithWrapper"
}`

var FuncTypeParamsInterfaceJSON = `{
  "name": "FuncTypeParamsInterface",
  "comment": "// FuncTypeParamsInterface is a dummy interface to test program\n",
  "functions": [
    {
      "name": "FuncTypeParams",
      "params": [
        {
          "name": "f",
          "type": {
            "name": "func(int, *ast.MapType) int",
            "imports": [
              {
                "package": "ast",
                "path": "go/ast"
              }
            ]
          }
        },
        {
          "name": "a",
          "type": {
            "name": "int"
          }
        },
        {
          "name": "b",
          "type": {
            "name": "int"
          }
        }
      ],
      "res": [
        {
          "name": "returnName1",
          "type": {
            "name": "func(uuid.UUID) error",
            "imports": [
              {
                "package": "uuid",
                "path": "github.com/google/uuid"
              }
            ]
          }
        }
      ],
      "comment": "// Function param type\n"
    }
  ],
  "imports": [
    {
      "package": "uuid",
      "path": "github.com/google/uuid"
    },
    {
      "package": "ast",
      "path": "go/ast"
    }
  ],
  "wrapperPackageName": "tests",
  "wrapperStructName": "Wrapper",
  "middlewareFunctionName": "WithWrapper"
}`

var CompositeParamsInterfaceJSON = `{
  "name": "CompositeParamsInterface",
  "comment": "// CompositeParamsInterface is a dummy interface to test program\n",
  "functions": [
    {
      "name": "Array",
      "params": [
        {
          "name": "a",
          "type": {
            "name": "[3]uuid.UUID",
            "imports": [
              {
                "package": "uuid",
                "path": "github.com/google/uuid"
              }
            ]
          }
        }
      ],
      "res": [
        {
          "name": "r",
          "type": {
            "name": "[10]bool"
          }
        }
      ],
      "comment": "// Array param types\n"
    },
    {
      "name": "Channel",
      "params": [
        {
          "name": "paramName1",
          "type": {
            "name": "chan string"
          }
        },
        {
          "name": "paramName2",
          "type": {
            "name": "\u003c-chan bool"
          }
        },
        {
          "name": "paramName3",
          "type": {
            "name": "chan\u003c- int"
          }
        }
      ],
      "res": [
        {
          "name": "returnName1",
          "type": {
            "name": "chan int"
          }
        }
      ],
      "comment": "// Channel param types\n"
    },
    {
      "name": "Composite",
      "params": [
        {
          "name": "m",
          "type": {
            "name": "map[string]chan int"
          }
        },
        {
          "name": "d",
          "type": {
            "name": "[2]chan func(string) map[bool]*ast.MapType",
            "imports": [
              {
                "package": "ast",
                "path": "go/ast"
              }
            ]
          }
        }
      ],
      "res": [
        {
          "name": "returnName1",
          "type": {
            "name": "[]chan func(string) error"
          }
        }
      ],
      "comment": "// Composite param types\n"
    },
    {
      "name": "Map",
      "params": [
        {
          "name": "paramName1",
          "type": {
            "name": "map[string]uuid.UUID",
            "imports": [
              {
                "package": "uuid",
                "path": "github.com/google/uuid"
              }
            ]
          }
        }
      ],
      "res": [
        {
          "name": "returnName1",
          "type": {
            "name": "map[bool]int"
          }
        }
      ],
      "comment": "// Map param types\n"
    },
    {
      "name": "Slice",
      "params": [
        {
          "name": "paramName1",
          "type": {
            "name": "[]uuid.UUID",
            "imports": [
              {
                "package": "uuid",
                "path": "github.com/google/uuid"
              }
            ]
          }
        },
        {
          "name": "paramName2",
          "type": {
            "name": "[]int"
          }
        }
      ],
      "res": [
        {
          "name": "returnName1",
          "type": {
            "name": "[]bool"
          }
        }
      ],
      "comment": "// Slice param types\n"
    }
  ],
  "imports": [
    {
      "package": "uuid",
      "path": "github.com/google/uuid"
    },
    {
      "package": "ast",
      "path": "go/ast"
    }
  ],
  "wrapperPackageName": "tests",
  "wrapperStructName": "Wrapper",
  "middlewareFunctionName": "WithWrapper"
}`
