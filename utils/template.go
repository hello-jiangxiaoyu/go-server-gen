package utils

import (
	"bytes"
	"errors"
	"go/format"
	"text/template"
)

// PhaseTemplate 模板解析
func PhaseTemplate(tpl string, data any) (string, error) {
	t := template.New("gen").Funcs(defaultFuncMap)
	t = template.Must(t.Parse(tpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	if buf.Len() == 0 {
		return "", errors.New("code is empty")
	}
	return buf.String(), nil
}

// PhaseAndFormat 解析并格式化go代码
func PhaseAndFormat(tpl string, data any) (string, error) {
	buf, err := PhaseTemplate(tpl, data)
	if err != nil {
		return "", err
	}
	code, err := format.Source([]byte(buf))
	if err != nil {
		return "", err
	}

	return string(code), nil
}
