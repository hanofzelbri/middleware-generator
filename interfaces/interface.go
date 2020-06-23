package interfaces

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"
)

// BuildInterface creates an Interface object for provided options
func BuildInterface(config *Config) (*Interface, error) {
	iface, ok := config.Object.Type().Underlying().(*types.Interface)
	if !ok {
		return nil, fmt.Errorf("Passed type name %q in package %q is not an interface", config.InterfaceName, config.Package.Path())
	}

	inter := &Interface{
		Name:                   config.InterfaceName,
		Comment:                commentText(config, config.Object.Pos()),
		Functions:              interfaceFunctions(config, iface),
		WrapperStructName:      config.WrapperStructName,
		WrapperPackageName:     config.WrapperPackageName,
		MiddleWareFunctionName: config.options.MiddlewareFunctionName,
	}

	fixupInterface(inter)

	return inter, nil
}

func commentText(config *Config, pos token.Pos) string {
	_, paths, _ := config.Program.PathEnclosingInterval(pos, pos)
	for _, n := range paths {
		switch n := n.(type) {
		case *ast.FuncDecl:
			return commentGroupToString(n.Doc)
		case *ast.GenDecl:
			return commentGroupToString(n.Doc)
		case *ast.Field:
			return commentGroupToString(n.Doc)
		}
	}

	return ""
}

func commentGroupToString(commentGroup *ast.CommentGroup) string {
	s := ""

	if commentGroup == nil {
		return s
	}

	for _, comment := range commentGroup.List {
		s += fmt.Sprintf("%v\n", comment.Text)
	}

	return s
}

func interfaceFunctions(config *Config, iface *types.Interface) []Func {
	funcs := []Func{}

	for i := 0; i < iface.NumMethods(); i++ {
		meth := iface.Method(i)

		sig, ok := meth.Type().(*types.Signature)
		if !ok {
			continue
		}

		f := Func{
			Name:       meth.Name(),
			Comment:    commentText(config, meth.Pos()),
			Params:     signatureVariables(sig.Params(), config.options.EmptyFunctionParamNamePrefix),
			Res:        signatureVariables(sig.Results(), config.options.EmptyFunctionReturnParamNamePrefix),
			IsVariadic: sig.Variadic(),
		}

		funcs = append(funcs, f)
	}

	return funcs
}

func signatureVariables(tuple *types.Tuple, emptyNamePrefix string) []Param {
	params := make([]Param, tuple.Len())

	for i := 0; i < tuple.Len(); i++ {
		param := tuple.At(i)

		name := param.Name()
		if name == "" {
			name = fmt.Sprintf("%v%v", emptyNamePrefix, i+1)
		}

		t := &Type{}
		configureParamType(t, param.Type())

		params[i] = Param{
			Name: name,
			Type: *t,
		}
	}

	return params
}

func configureParamType(t *Type, typ types.Type) {
	switch typ := typ.(type) {
	case *types.Basic:
		configureParamTypeName(t, typ.Name())
	case *types.Pointer:
		configureParamTypeName(t, typ.String())
		configureParamType(t, typ.Elem())
	case *types.Named:
		configureParamTypeName(t, typ.Obj().Name())
		if pkg := typ.Obj().Pkg(); pkg != nil {
			t.Imports = append(t.Imports, Import{Package: pkg.Name(), Path: pkg.Path()})
		}
	case *types.Slice:
		configureParamTypeName(t, typ.String())
		configureParamType(t, typ.Elem())
	case *types.Signature:
		configureParamTypeName(t, typ.String())
		variables := signatureVariables(typ.Params(), "")
		variables = append(variables, signatureVariables(typ.Results(), "")...)
		for _, v := range variables {
			t.Imports = append(t.Imports, v.Type.Imports...)
		}
	case *types.Array:
		configureParamTypeName(t, typ.String())
		configureParamType(t, typ.Elem())
	case *types.Chan:
		configureParamTypeName(t, typ.String())
		configureParamType(t, typ.Elem())
	case *types.Map:
		configureParamTypeName(t, typ.String())
		configureParamType(t, typ.Elem())
	}
}

func configureParamTypeName(t *Type, name string) {
	if t.Name == "" {
		t.Name = name
	}
}

func fixupInterface(inter *Interface) {
	imports := map[string]Import{}

	for fi, f := range inter.Functions {
		for pi, p := range f.Params {
			for _, i := range p.Type.Imports {
				imports[i.Path] = i
				inter.Functions[fi].Params[pi].Type.Name = strings.ReplaceAll(p.Type.Name, i.Path, i.Package)
			}

			if f.IsVariadic && pi == len(f.Params)-1 {
				inter.Functions[fi].Params[pi].Type.Name = strings.Replace(p.Type.Name, "[]", "...", 1)
			}
		}
		for ri, p := range f.Res {
			for _, i := range p.Type.Imports {
				imports[i.Path] = i
				inter.Functions[fi].Params[ri].Type.Name = strings.ReplaceAll(p.Type.Name, i.Path, i.Package)
			}
		}
	}

	for _, i := range imports {
		inter.Imports = append(inter.Imports, i)
	}
}
