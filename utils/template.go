package utils

import (
	"bytes"
	"fmt"
	"go/format"
	"regexp"
	"strings"
	"text/template"
)

// ParseTemplate 模板解析
func ParseTemplate(tpl string, data any) (string, error) {
	var buf bytes.Buffer
	t, err := template.New("gen").Funcs(defaultFuncMap).Parse(tpl)
	if err != nil {
		return "", WithMessage(err, "invalid template config")
	}
	if err = t.Execute(&buf, data); err != nil {
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
		codeBody, err := format.Source([]byte(RemoveDuplicateImport(body)))
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

var (
	importReg = regexp.MustCompile(`import\s+\(([\s\S]*?)\)`)
	spaceReg  = regexp.MustCompile(`[\s\t]+`)
)

// RemoveDuplicateImport 去除源码中多于的import
func RemoveDuplicateImport(src string) string {
	matches := importReg.FindAllStringSubmatch(src, -1)
	if len(matches) == 0 {
		return src
	}

	dependencies := make(map[string]bool)
	replacement := ""
	for _, match := range matches {
		importContent := match[1]
		imports := strings.Split(importContent, "\n")
		for _, ipt := range imports {
			ipt = spaceReg.ReplaceAllString(ipt, "")
			if !dependencies[ipt] {
				dependencies[ipt] = true
				replacement += fmt.Sprintf("\t%s\n", ipt)
			}
		}
	}

	return importReg.ReplaceAllString(src, "import ("+replacement+")")
}
