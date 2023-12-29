package writer

import (
	"go/format"
	"os"
	"regexp"
	"strings"
)

// InsertPointerWriter 在文件指针下方插入代码
func InsertPointerWriter(path string, pointer string, handlers []string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if strings.Contains(line, pointer) {
			lines = append(lines[:i+1], append(handlers, lines[i+1:]...)...)
			break
		}
	}

	output := strings.Join(lines, "\n")
	if err = os.WriteFile(path, []byte(output), 0644); err != nil {
		return err
	}

	return nil
}

// GoFunctionFilter 过滤已存在的函数
func GoFunctionFilter(handlers map[string]string, src string) (map[string]string, error) {
	res := make(map[string]string)
	fCode, err := format.Source([]byte(src))
	if err != nil {
		return nil, err
	}
	for handler, code := range handlers {
		regFunction := regexp.MustCompile(`func .*` + handler + `\(.*\).*{`)
		if !regFunction.MatchString(string(fCode)) {
			res[handler] = code
		}
	}
	return res, nil
}

// GoStructFilter 过滤已存在的go结构体
func GoStructFilter(messages map[string]string, src string) (map[string]string, error) {
	res := make(map[string]string)
	fCode, err := format.Source([]byte(src))
	if err != nil {
		return nil, err
	}
	for msg, code := range messages {
		regFunction := regexp.MustCompile(msg + ` struct {`)
		if !regFunction.MatchString(string(fCode)) {
			res[msg] = code
		}
	}
	return res, nil
}
