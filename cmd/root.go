package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/spf13/cobra"
	"golang.org/x/tools/go/loader"
)

var (
	queryVar                      string
	wrapperVar                    string
	outputVar                     string
	middlewareFunctionNameVar     string
	emptyFunctionParamNamePrefix  string
	emptyFunctionReturnNamePrefix string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dmgmori-logging-generator",
	Short: "Generates logging middleware for golang interface",
	Long: `This golang generator can be used to generate a logging
middleware with the zerolog logging library for an provided interface.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		o, err := setupOptions(queryVar)
		if err != nil {
			return err
		}

		i, err := buildInterface(o)
		if err != nil {
			return err
		}

		return printInterface(i)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&queryVar, "interface", "i", "", "Interface definition to generate logging middleware for.")
	rootCmd.MarkPersistentFlagRequired("interface")

	rootCmd.PersistentFlags().StringVarP(&wrapperVar, "wrapper", "w", "", "Wrapper definition for implementation of middleware interface.")
	rootCmd.PersistentFlags().StringVarP(&middlewareFunctionNameVar, "functionname", "f", "WithMiddleware", "Function name for middleware")
	rootCmd.PersistentFlags().StringVarP(&outputVar, "output", "o", "", "Output file. If empty StdOut is used")
	rootCmd.PersistentFlags().StringVarP(&emptyFunctionParamNamePrefix, "emptyFunctionParamNamePrefix", "p", "param", "If there is no function parameter name provided this prefix will be used")
	rootCmd.PersistentFlags().StringVarP(&emptyFunctionReturnNamePrefix, "emptyFunctionReturnNamePrefix", "r", "ret", "If there is no function parameter return name provided this prefix will be used")
}

func setupOptions(query string) (*Options, error) {
	idx := strings.LastIndex(query, ".")
	if idx == -1 || query[:idx] == "" || query[idx+1:] == "" {
		return nil, errors.New("--interface (-i) flag should be like path/to/package.type")
	}

	interfaceName := query[idx+1:]
	packageName := query[:idx]

	prog, err := loadProgram(packageName)
	if err != nil {
		return nil, err
	}

	return &Options{
		InterfaceName:                      interfaceName,
		PackageName:                        packageName,
		Program:                            prog,
		PackageInfo:                        prog.Imported[packageName],
		FSet:                               prog.Fset,
		EmptyFunctionParamNamePrefix:       emptyFunctionParamNamePrefix,
		EmptyFunctionReturnParamNamePrefix: emptyFunctionReturnNamePrefix,
	}, nil
}

func loadProgram(packageName string) (*loader.Program, error) {
	cfg := &loader.Config{
		AllowErrors:         true,
		ImportPkgs:          map[string]bool{packageName: true},
		TypeCheckFuncBodies: func(string) bool { return false },
	}

	return cfg.Load()
}

func buildInterface(options *Options) (*Interface, error) {
	inter := &Interface{
		Name:                   options.InterfaceName,
		PackageName:            wrapperPackageName(wrapperVar, options.PackageName),
		StructName:             wrapperStructName(wrapperVar, options.InterfaceName),
		MiddleWareFunctionName: middlewareFunctionNameVar,
	}

	pos, err := interfaceTypeDefinitionPos(options)
	if err != nil {
		return nil, err
	}

	typFileName := options.FSet.File(pos).Name()
	decl, err := interfaceGenericDeclaration(typFileName, options)
	if err != nil {
		return nil, err
	}

	inter.Comment = commentGroupToString(decl.Doc)

	typ, err := interfaceTypeSpec(decl, options)
	if err != nil {
		return nil, err
	}

	idecl, err := interfaceType(typ, options)
	if err != nil {
		return nil, err
	}

	inter.Functions, err = interfaceFunctions(idecl, options)
	if err != nil {
		return nil, err
	}

	imports := options.PackageInfo.Pkg.Imports()
	inter.Imports, err = interfaceImports(imports, inter.Functions)
	if err != nil {
		return nil, err
	}

	return inter, nil
}

func interfaceTypeDefinitionPos(options *Options) (token.Pos, error) {
	for _, obj := range options.PackageInfo.Defs {
		if obj == nil ||
			obj.Name() != options.InterfaceName ||
			obj.Pkg().Path() != options.PackageName {
			continue
		}

		_, ok := obj.Type().(*types.Named)
		if ok {
			return obj.Pos(), nil
		}
	}

	return token.Pos(-1), fmt.Errorf("Interface %q in package %q not found", options.InterfaceName, options.PackageName)
}

func interfaceGenericDeclaration(fileName string, options *Options) (*ast.GenDecl, error) {
	var interfaceGenericDeclaration *ast.GenDecl

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, decl := range file.Decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			continue
		}

		if _, err := interfaceTypeSpec(decl, options); err != nil {
			continue
		}

		interfaceGenericDeclaration = decl
		break
	}

	if interfaceGenericDeclaration == nil {
		return nil, fmt.Errorf("Interface generic declaration for interface %q in package %q not found", options.InterfaceName, options.PackageName)
	}

	return interfaceGenericDeclaration, nil
}

func interfaceTypeSpec(decl *ast.GenDecl, options *Options) (*ast.TypeSpec, error) {
	for _, spec := range decl.Specs {
		spec := spec.(*ast.TypeSpec)
		if spec.Name.Name == options.InterfaceName {
			return spec, nil
		}
	}

	return nil, fmt.Errorf("No typespec found for interface %q in package %q", options.InterfaceName, options.PackageName)
}

func interfaceType(typ *ast.TypeSpec, options *Options) (*ast.InterfaceType, error) {
	idecl, ok := typ.Type.(*ast.InterfaceType)
	if !ok {
		return nil, fmt.Errorf("Type for interface %q in package %q is not of type *ast.InterfaceType", options.InterfaceName, options.PackageName)
	}

	return idecl, nil
}

func interfaceFunctions(idecl *ast.InterfaceType, options *Options) ([]Func, error) {
	if idecl.Methods == nil {
		return nil, fmt.Errorf("Interface %q in package %q is empty", options.InterfaceName, options.PackageName)
	}

	functions := []Func{}

	for _, m := range idecl.Methods.List {
		typ := m.Type.(*ast.FuncType)
		functions = append(functions, Func{
			Name:    m.Names[0].Name,
			Comment: commentGroupToString(m.Doc),
			Params:  interfaceFunctionFields(typ.Params, options, options.EmptyFunctionParamNamePrefix),
			Res:     interfaceFunctionFields(typ.Results, options, options.EmptyFunctionReturnParamNamePrefix),
		})
	}

	return functions, nil
}

func interfaceFunctionFields(fields *ast.FieldList, options *Options, emptyFieldNamePrefix string) []Param {
	params := []Param{}

	if fields == nil {
		return params
	}

	for _, field := range fields.List {
		typ := interfaceFunctionFieldType(field.Type, options)

		if len(field.Names) == 0 {
			params = append(params, Param{
				Name: fmt.Sprintf("%v%v", emptyFieldNamePrefix, len(params)+1),
				Type: typ,
			})
		}

		for _, name := range field.Names {
			params = append(params, Param{
				Name: name.Name,
				Type: typ,
			})
		}
	}

	return params
}

func interfaceFunctionFieldType(e ast.Expr, options *Options) Type {
	typ := Type{}

	ast.Inspect(e, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.StarExpr:
			typ.IsPointer = true
		case *ast.Ellipsis:
			typ.IsVariadic = true
		case *ast.FuncType:
			typ.IsFunction = true
			var typeNameBuf bytes.Buffer
			printer.Fprint(&typeNameBuf, options.FSet, n)
			typ.Name = typeNameBuf.String()
			return false
		case *ast.Ident:
			if typ.Name != "" {
				typ.Package = typ.Name
				typ.Name = n.Name
			} else {
				typ.Name = n.Name
			}
		}
		return true
	})

	return typ
}

func wrapperPackageName(wrapper string, packageName string) string {
	if i := strings.IndexRune(wrapper, '.'); i != -1 {
		packageName = (wrapper)[:i]
	}

	return filepath.Base(packageName)
}

func wrapperStructName(wrapper string, interfaceName string) string {
	structName := wrapper

	if structName == "" {
		structName = string(unicode.ToLower(rune(interfaceName[0]))) + interfaceName[1:]
	}

	if i := strings.IndexRune(wrapper, '.'); i != -1 {
		structName = (wrapper)[i+1:]
	}

	return structName
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

func interfaceImports(imports []*types.Package, interfaceFunctions []Func) ([]Import, error) {
	interfaceImportsMap := map[string]Import{}

	importPackages := make(map[string]Import, len(imports))
	for _, i := range imports {
		importPackages[i.Name()] = Import{Name: i.Name(), Path: i.Path()}
	}

	for _, f := range interfaceFunctions {
		for _, p := range f.Params {
			pkg := p.Type.Package

			if pkg == "" {
				continue
			}

			i, ok := importPackages[pkg]
			if !ok {
				return []Import{}, fmt.Errorf("Import type definition for package %q is not available", p.Type.Package)
			}
			interfaceImportsMap[p.Type.Package] = i
		}
	}

	ret := []Import{}
	for _, val := range interfaceImportsMap {
		ret = append(ret, val)
	}

	return ret, nil
}

func printInterface(i *Interface) error {
	buf := &bytes.Buffer{}
	t := template.Must(template.New("tmpl").Parse(tmpl))
	t.Execute(buf, i)

	pretty, err := format.Source(buf.Bytes())
	if err != nil {
		os.Stderr.Write(buf.Bytes())
		return err
	}

	f := os.Stdout
	if outputVar != "" {
		f, err = os.OpenFile(outputVar, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}
	}

	_, err = f.Write(pretty)
	if err != nil {
		return err
	}

	return f.Close()
}
