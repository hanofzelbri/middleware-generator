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
	"unicode"

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
	typ, err := interfaceTypeDefinition(pkg, query)
	if err != nil {
		return nil, err
	}

	typFileName := prog.Fset.File(typ.Pos()).Name()
	inter.Comment, err = interfaceComment(typFileName, query)
	if err != nil {
		return nil, err
	}

	for _, function := range interfaceFunctions(typ) {
		sig, ok := function.Type().(*types.Signature)
		if !ok {
			continue
		}

		params := sig.Params()
		res := sig.Results()
		fn := Func{
			Name:       function.Name(),
			Params:     make([]Param, params.Len()),
			Res:        make([]Param, res.Len()),
			IsVariadic: sig.Variadic(),
		}

		_ = fn // TODO: Implement filling of fn object
	}

	return nil, nil
}

func interfaceTypeDefinition(pkg *loader.PackageInfo, query *Query) (typ *ast.TypeSpec, err error) {
	for ident, obj := range pkg.Defs {
		if obj == nil ||
			obj.Name() != query.InterfaceName ||
			obj.Pkg().Path() != query.PackageName {
			continue
		}

		t, ok := ident.Obj.Decl.(*ast.TypeSpec)
		if !ok {
			continue
		}

		_, ok = t.Type.(*ast.InterfaceType)
		if !ok {
			continue
		}

		typ = t
		break
	}

	if typ == nil {
		err = fmt.Errorf("Interface %q in package %q not found", query.InterfaceName, query.PackageName)
	}

	return
}

func interfaceComment(fileName string, query *Query) (string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}

	for _, decl := range file.Decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			continue
		}

		for _, spec := range decl.Specs {
			typ := spec.(*ast.TypeSpec)
			if typ.Name.Name == query.InterfaceName {
				return decl.Doc.Text(), nil
			}
		}
	}

	return "", fmt.Errorf("Comment for interface %q in package %q not found", query.InterfaceName, query.PackageName)
}

func interfaceFunctions(typ *ast.TypeSpec) map[string]*types.Func {
	functions := map[string]*types.Func{}

	// for i := 0; i < typ.NumMethods(); i++ {
	// 	m := typ.Method(i)

	// 	if !isFunctionExported(m) && !unexported {
	// 		continue
	// 	}

	// 	functions[m.Name()] = m
	// }

	return functions
}

func isFunctionExported(function *types.Func) bool {
	return unicode.IsLower(rune(function.Name()[0]))
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
