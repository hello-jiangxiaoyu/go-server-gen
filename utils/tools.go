package utils

import (
	"errors"
	"os/exec"
	"strings"
)

// DeferErr handle defer function err
func DeferErr(errFunc func() error) {
	if err := errFunc(); err != nil {
		println("### Defer err: ", err.Error())
	}
}

// WithMessage err和msg有一个不为空则返回error
func WithMessage(err error, msg string) error {
	if err == nil && msg == "" {
		return nil
	} else if err == nil {
		return errors.New(msg)
	} else if msg == "" {
		return err
	}

	return errors.New(msg + ": " + err.Error())
}

// DeduplicateStrings 字符串数组去重并去除空字符串
func DeduplicateStrings(arr []string) []string {
	visited := make(map[string]bool)
	result := make([]string, 0)
	for _, str := range arr {
		if str == "" {
			continue
		}
		if !visited[str] {
			visited[str] = true
			result = append(result, str)
		}
	}

	return result
}

var projectName = ""

// GetProjectName 获取当前项目名
func GetProjectName() (string, error) {
	if projectName != "" {
		return projectName, nil
	}
	cmd := exec.Command("go", "list", "-m")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	projectName = strings.TrimSpace(string(output))
	return projectName, nil
}
