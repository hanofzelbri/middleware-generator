package interfaces

import (
	"errors"
	"fmt"
	"go/parser"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/tools/go/loader"
)

// SetupConfig provide factory function to generate an options object
// which then will be used for further handling
func SetupConfig(options Options) (*Config, error) {
	idx := strings.LastIndex(options.Query, ".")
	if idx == -1 || options.Query[:idx] == "" || options.Query[idx+1:] == "" {
		return nil, errors.New("--interface (-i) flag should be like path/to/package.type")
	}

	interfaceName := options.Query[idx+1:]
	packageName := options.Query[:idx]

	program, err := loadProgram(packageName)
	if err != nil {
		return nil, err
	}

	pkg := program.Package(packageName).Pkg
	obj := pkg.Scope().Lookup(interfaceName)
	if obj == nil {
		return nil, fmt.Errorf("Interface %q not found in package %q", interfaceName, packageName)
	}

	return &Config{
		InterfaceName:      interfaceName,
		PackageName:        packageName,
		Program:            program,
		Package:            pkg,
		Object:             obj,
		WrapperStructName:  wrapperStructName(options.Wrapper, interfaceName),
		WrapperPackageName: wrapperPackageName(options.Wrapper, packageName),
		options:            options,
	}, nil
}

func loadProgram(packageName string) (*loader.Program, error) {
	cfg := &loader.Config{
		AllowErrors:         true,
		ImportPkgs:          map[string]bool{packageName: true},
		TypeCheckFuncBodies: func(string) bool { return false },
		ParserMode:          parser.ParseComments,
	}

	cfg.ImportWithTests(packageName)

	return cfg.Load()
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
