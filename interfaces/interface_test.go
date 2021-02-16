package interfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildInterface(t *testing.T) {
	o := Options{
		Query:                              "io.Reader",
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
				o.Wrapper = "tests.Wrapper"
				return o
			},
			want:    ReaderInterface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.TestInterface1",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.TestInterface1"
				o.Wrapper = ""
				return o
			},
			want:    TestInterface1Interface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.EmptyInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.EmptyInterface"
				o.Wrapper = "tests.Wrapper"
				return o
			},
			want:    EmptyInterfaceInterface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.UnnammedParametersInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.UnnammedParametersInterface"
				o.Wrapper = ""
				return o
			},
			want:    UnnammedParametersInterfaceInterface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.ImportedParamTypeInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.ImportedParamTypeInterface"
				o.Wrapper = ""
				return o
			},
			want:    ImportedParamTypeInterfaceInterface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.VariadicParamTypeInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.VariadicParamTypeInterface"
				o.Wrapper = ""
				return o
			},
			want:    VariadicParamTypeInterfaceInterface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.FuncTypeParamsInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.FuncTypeParamsInterface"
				o.Wrapper = ""
				return o
			},
			want:    FuncTypeParamsInterfaceInterface,
			wantErr: false,
		},
		{
			name: "github.com/hanofzelbri/middleware-generator/interfaces.CompositeParamsInterface",
			options: func() Options {
				o.Query = "github.com/hanofzelbri/middleware-generator/interfaces.CompositeParamsInterface"
				o.Wrapper = ""
				return o
			},
			want:    CompositeParamsInterfaceInterface,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildInterface(tt.options())
			assert.Equal(t, tt.want, got)
			assert.Equal(t, (err != nil), tt.wantErr)
		})
	}
}
