package interfaces

import (
	"reflect"
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
		want    *Interface
		wantErr bool
	}{
		{
			name: "io.Reader",
			options: func() Options {
				o.Query = "io.Reader"
				return o
			},
			want:    ReaderInterface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.TestInterface1",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.TestInterface1"
				return o
			},
			want:    TestInterface1Interface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.EmptyInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.EmptyInterface"
				return o
			},
			want:    EmptyInterfaceInterface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.UnnammedParametersInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.UnnammedParametersInterface"
				return o
			},
			want:    UnnammedParametersInterfaceInterface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.ImportedParamTypeInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.ImportedParamTypeInterface"
				return o
			},
			want:    ImportedParamTypeInterfaceInterface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.VariadicParamTypeInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.VariadicParamTypeInterface"
				return o
			},
			want:    VariadicParamTypeInterfaceInterface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.FuncTypeParamsInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.FuncTypeParamsInterface"
				return o
			},
			want:    FuncTypeParamsInterfaceInterface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.CompositeParamsInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.CompositeParamsInterface"
				return o
			},
			want:    CompositeParamsInterfaceInterface,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildInterface(tt.options())
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}
