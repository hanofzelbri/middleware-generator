package interfaces

import (
    "bytes"
    "go/format"
    "text/template"
)

// InterfaceWrapperTemplate returns the filled template with Interface data
func InterfaceWrapperTemplate(i *Interface) ([]byte, error) {
    buf := &bytes.Buffer{}
    t := template.Must(template.New("tmpl").Parse(tmpl))
    t.Execute(buf, i)

    pretty, err := format.Source(buf.Bytes())
    if err != nil {
        return buf.Bytes(), err
    }

    return pretty, nil
}

var tmpl = `// Code generated by middleware-generator; DO NOT EDIT

package {{.WrapperPackageName}}

import (
    "time"

    {{range .Imports}}
    "{{.Path}}"
    {{- end}}
    "github.com/rs/zerolog/log"
)

{{if .Comment}}{{.Comment}}{{end -}}
type {{.WrapperStructName}} struct {
    wrapper {{.Name}}
}

// {{.MiddleWareFunctionName}} adds logging for interface {{.Name}}
func {{.MiddleWareFunctionName}}(wrapper {{.Name}}) {{.Name}} {
    return &{{.WrapperStructName}}{
        wrapper: wrapper,
    }
}

{{range .Functions}}
{{if .Comment}}{{.Comment}}{{end -}}
func (l *{{$.WrapperStructName}}) {{.Name}}({{range .Params}}{{.Name}} {{.Type.Name}}, {{end}}) ({{range .Res}}{{.Name}} {{.Type.Name}}, {{end}}) {
    defer func(begin time.Time) {
        log.Info().
            {{range .Params}}
                Interface("{{.Name}}", {{.Name}}).
            {{end}}
            Dur("took", time.Since(begin)).
            {{range .Res}}
                Interface("{{.Name}}", {{.Name}}).
            {{end}}
            Msg("Method {{.Name}} called")
    }(time.Now())

    {{if .Res}}return{{end}} l.wrapper.{{.Name}}({{range .Params}}{{.Name}}, {{end}})
}
{{end}}
`
