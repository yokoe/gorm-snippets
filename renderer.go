package main

import (
	"bytes"
	"text/template"

	"golang.org/x/xerrors"
)

func renderTpl(body string, values map[string]interface{}) (string, error) {
	tpl := template.Must(template.New("").Parse(body))
	buf := bytes.Buffer{}
	err := tpl.Execute(&buf, values)
	if err != nil {
		return "", xerrors.Errorf("template rendering error: %w", err)
	}
	return buf.String(), nil
}
