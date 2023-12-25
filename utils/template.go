package utils

import (
	"bytes"
	"text/template"
)

// PhaseTemplate 模板解析
func PhaseTemplate(tpl string, data any, funcMap ...template.FuncMap) ([]byte, error) {
	if len(funcMap) == 0 || funcMap == nil {
		funcMap = []template.FuncMap{defaultFuncMap}
	}
	t := template.New("gen").Funcs(funcMap[0])
	t = template.Must(t.Parse(tpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// PhaseAndFormat 解析并格式化go代码
func PhaseAndFormat(tpl string, data any, funcMap ...template.FuncMap) (string, error) {
	buf, err := PhaseTemplate(tpl, data, funcMap...)
	if err != nil {
		return "", err
	}
	code, err := FormatCode(buf)
	if err != nil {
		return "", err
	}

	return string(code), nil
}
