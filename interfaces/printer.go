package interfaces

import (
	"bytes"
	"go/format"
	"html/template"
	"os"
)

// PrintInterface prints the generated Interface to provided filePath
// If filePath is empty then StdOut is used instead
func PrintInterface(i *Interface, filePath string) error {
	buf := &bytes.Buffer{}
	t := template.Must(template.New("tmpl").Parse(tmpl))
	t.Execute(buf, i)

	pretty, err := format.Source(buf.Bytes())
	if err != nil {
		os.Stderr.Write(buf.Bytes())
		return err
	}

	f := os.Stdout
	if filePath != "" {
		f, err = os.OpenFile(filePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
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
