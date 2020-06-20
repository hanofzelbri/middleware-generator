package cmd

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/tools/go/loader"
)

var (
	queryVar      string
	wrapperVar    string
	outputVar     string
	unexportedVar bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dmgmori-logging-generator",
	Short: "Generates logging middleware for golang interface",
	Long: `This golang generator can be used to generate a logging
middleware with the zerolog logging library for an provided interface.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		q, err := validateInterfaceParam(queryVar)
		if err != nil {
			return err
		}

		_, err = buildInterface(q)

		return err
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
	rootCmd.PersistentFlags().StringVarP(&queryVar, "interface", "i", "", "Interface to generate logging middleware for.")
	rootCmd.MarkPersistentFlagRequired("interface")

	rootCmd.PersistentFlags().StringVarP(&outputVar, "output", "o", "", "Output file.")

	rootCmd.PersistentFlags().StringVarP(&wrapperVar, "wrapper", "w", "", "Wrapper name for implementation of middleware interface.")
	rootCmd.MarkPersistentFlagRequired("wrapper")

	rootCmd.PersistentFlags().BoolVarP(&unexportedVar, "all", "a", false, "Include also unexported methods.")
}

func validateInterfaceParam(query string) (*Query, error) {
	idx := strings.LastIndex(query, ".")
	if idx == -1 || query[:idx] == "" || query[idx+1:] == "" {
		return nil, errors.New("--interface (-i) flag should be like path/to/package.type")
	}

	return &Query{
		InterfaceName: query[idx+1:],
		PackageName:   query[:idx],
	}, nil
}

func buildInterface(query *Query) (*Interface, error) {
	prog, err := loadProgram(query)
	i, err := buildInterfaceFromProgram(prog, query)
	return i, err
}

func loadProgram(query *Query) (*loader.Program, error) {
	cfg := &loader.Config{
		AllowErrors:         true,
		ImportPkgs:          map[string]bool{query.PackageName: true},
		TypeCheckFuncBodies: func(string) bool { return false },
	}

	return cfg.Load()
}

func buildInterfaceFromProgram(prog *loader.Program, query *Query) (*Interface, error) {
	inter := &Interface{
		Name:        query.InterfaceName,
		PackageName: wrapperPackageName(wrapperVar),
		StructName:  wrapperStructName(wrapperVar),
	}

	pkg := prog.Imported[query.PackageName]
	pos, err := interfaceTypeDefinitionPos(pkg, query)
	if err != nil {
		return nil, err
	}

	typFileName := prog.Fset.File(pos).Name()
	decl, err := interfaceGenericDeclaration(typFileName, query)
	if err != nil {
		return nil, err
	}

	inter.Comment = commentGroupToString(decl.Doc)

	typ, err := interfaceTypeSpec(decl, query)
	if err != nil {
		return nil, err
	}

	idecl, err := interfaceType(typ, query)
	if err != nil {
		return nil, err
	}

	inter.Functions, err = interfaceFunctions(idecl, query)
	if err != nil {
		return nil, err
    }

	return inter, nil
}

func interfaceTypeDefinitionPos(pkg *loader.PackageInfo, query *Query) (token.Pos, error) {
	for _, obj := range pkg.Defs {
		if obj == nil ||
			obj.Name() != query.InterfaceName ||
			obj.Pkg().Path() != query.PackageName {
			continue
		}

		_, ok := obj.Type().(*types.Named)
		if ok {
			return obj.Pos(), nil
		}
	}

	return token.Pos(-1), fmt.Errorf("Interface %q in package %q not found", query.InterfaceName, query.PackageName)
}

func interfaceGenericDeclaration(fileName string, query *Query) (*ast.GenDecl, error) {
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

		if _, err := interfaceTypeSpec(decl, query); err != nil {
			continue
		}

		interfaceGenericDeclaration = decl
		break
	}

	if interfaceGenericDeclaration == nil {
		return nil, fmt.Errorf("Interface generic declaration for interface %q in package %q not found", query.InterfaceName, query.PackageName)
	}

	return interfaceGenericDeclaration, nil
}

func interfaceTypeSpec(decl *ast.GenDecl, query *Query) (*ast.TypeSpec, error) {
	for _, spec := range decl.Specs {
		spec := spec.(*ast.TypeSpec)
		if spec.Name.Name == query.InterfaceName {
			return spec, nil
		}
	}

	return nil, fmt.Errorf("No typespec found for interface %q in package %q", query.InterfaceName, query.PackageName)
}

func interfaceType(typ *ast.TypeSpec, query *Query) (*ast.InterfaceType, error) {
	idecl, ok := typ.Type.(*ast.InterfaceType)
	if !ok {
		return nil, fmt.Errorf("Type for interface %q in package %q is not of type *ast.InterfaceType", query.InterfaceName, query.PackageName)
	}

	return idecl, nil
}

func interfaceFunctions(idecl *ast.InterfaceType, query *Query) ([]Func, error) {
	if idecl.Methods == nil {
		return nil, fmt.Errorf("Interface %q in package %q is empty", query.InterfaceName, query.PackageName)
	}

	functions := []Func{}

	for _, m := range idecl.Methods.List {
		typ := m.Type.(*ast.FuncType)
		functions = append(functions, Func{
			Name:    m.Names[0].Name,
			Comment: commentGroupToString(m.Comment),
			Params:  interfaceFunctionFields(typ.Params),
			Res:     interfaceFunctionFields(typ.Results),
		})
	}

	return functions, nil
}

func interfaceFunctionFields(fields *ast.FieldList) []Param {
	params := []Param{}

	for _, field := range fields.List {
		typ := interfaceFunctionFieldType(field.Type)

		if len(field.Names) == 0 {
			params = []Param{{Type: typ}}
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

func interfaceFunctionFieldType(e ast.Expr) Type {
	typ := Type{}

	ast.Inspect(e, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.StarExpr:
			typ.IsPointer = true
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

func wrapperPackageName(wrapper string) string {
	packageName := ""

	if i := strings.IndexRune(wrapper, '.'); i != -1 {
		packageName = (wrapper)[:i]
	}

	return packageName
}

func wrapperStructName(wrapper string) string {
	structName := wrapper

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
