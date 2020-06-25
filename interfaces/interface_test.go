package interfaces

import (
	"encoding/json"
	"fmt"
	"testing"
)



func TestBuildInterface(t *testing.T) {
	o := Options{
		Query:                              "io.Reader",
		Wrapper:                            "tests.Wrapper",
		Output:                             "",
		MiddlewareFunctionName:             "WithWrapper",
		EmptyFunctionParamNamePrefix:       "paramName",
		EmptyFunctionReturnParamNamePrefix: "returnName",
	}

	tests := []struct {
		name    string
		options func() Options
		want    string
		wantErr bool
	}{
		{
			name: "io.Reader",
			options: func() Options {
				o.Query = "io.Reader"
				return o
			},
			want:    IoReaderJSON,
			wantErr: false,
		},
		{
            name: "github.com/hanofzelbri/middleware-generator/interfaces.TestInterface1",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.TestInterface1"
				return o
			},
			want:    TestInterface1JSON,
			wantErr: false,
		},
		{
            name: "github.com/hanofzelbri/middleware-generator/interfaces.EmptyInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.EmptyInterface"
				return o
			},
			want:    EmptyInterfaceJSON,
			wantErr: false,
		},
		{
            name: "github.com/hanofzelbri/middleware-generator/interfaces.UnnammedParametersInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.UnnammedParametersInterface"
				return o
			},
			want:    UnnammedParametersInterfaceJSON,
			wantErr: false,
		},
		{
            name: "github.com/hanofzelbri/middleware-generator/interfaces.ImportedParamTypeInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.ImportedParamTypeInterface"
				return o
			},
			want:    ImportedParamTypeInterfaceJSON,
			wantErr: false,
		},
		{
            name: "github.com/hanofzelbri/middleware-generator/interfaces.VariadicParamTypeInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.VariadicParamTypeInterface"
				return o
			},
			want:    VariadicParamTypeInterfaceJSON,
			wantErr: false,
		},
		{
            name: "github.com/hanofzelbri/middleware-generator/interfaces.FuncTypeParamsInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.FuncTypeParamsInterface"
				return o
			},
			want:    FuncTypeParamsInterfaceJSON,
			wantErr: false,
		},
		{
            name: "github.com/hanofzelbri/middleware-generator/interfaces.CompositeParamsInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.CompositeParamsInterface"
				return o
			},
			want:    CompositeParamsInterfaceJSON,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := BuildInterface(tt.options())
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			b, err := json.MarshalIndent(i, "", "  ")
			if err != nil {
				fmt.Println(err)
				t.Errorf("Marshalling interface error: %v", err)
				return
			}

            got := string(b)
			if got != tt.want {
				t.Errorf("BuildInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}
