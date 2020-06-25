package interfaces

import (
	"go/parser"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/tools/go/loader"
)

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
