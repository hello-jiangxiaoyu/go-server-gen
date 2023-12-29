package utils

import (
	"go/format"
	"os/exec"
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
