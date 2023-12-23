package utils

import (
	"go/format"
	"os/exec"
	"regexp"
	"strings"
)

// GetProjectName 获取当前项目名
func GetProjectName() (string, error) {
	cmd := exec.Command("go", "list", "-m")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

// FormatCode 格式化go代码
func FormatCode(code []byte) ([]byte, error) {
	f, err := format.Source(code)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// GoFunctionFilter 过滤已存在的函数
func GoFunctionFilter(handlers map[string]string, src string) (map[string]string, error) {
	res := make(map[string]string)
	fCode, err := FormatCode([]byte(src))
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
	fCode, err := FormatCode([]byte(src))
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
