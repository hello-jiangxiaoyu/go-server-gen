package utils

import (
	"bytes"
	"go/format"
	"text/template"
)

func FormatCode(code []byte) ([]byte, error) {
	f, err := format.Source(code)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func PhaseAndFormat(tpl string, data any) (string, error) {
	t := template.New("test")
	t = template.Must(t.Parse(tpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	code, err := FormatCode(buf.Bytes())
	if err != nil {
		return "", err
	}

	return string(code), nil
}
