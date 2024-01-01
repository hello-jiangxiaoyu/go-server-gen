package utils

import (
	"bytes"
	"go/format"
	"strings"
	"text/template"
)

// ParseTemplate 模板解析
func ParseTemplate(tpl string, data any) (string, error) {
	t := template.New("gen").Funcs(defaultFuncMap)
	t = template.Must(t.Parse(tpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// ParseSource 解析代码
func ParseSource(key, body string, data any) (string, string, error) {
	body, err := ParseTemplate(body, data)
	if err != nil {
		return "", "", err
	}
	if strings.Contains(key, ".go") { // 格式化go代码
		codeBody, err := format.Source([]byte(body))
		if err != nil {
			return "", "", err
		}
		body = string(codeBody)
	}

	if strings.Contains(key, "{{") { // 需要解析模板
		key, err = ParseTemplate(key, data)
		if err != nil {
			return "", "", err
		}
	}

	return key, body, nil
}
