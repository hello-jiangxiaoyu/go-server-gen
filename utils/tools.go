package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
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

var projectName = ""

// GetProjectName 获取当前项目名
func GetProjectName(dir ...string) (string, error) {
	if projectName != "" {
		return projectName, nil
	}
	cmd := exec.Command("go", "list", "-m")
	if len(dir) > 0 && dir[0] != "" {
		cmd.Dir = dir[0]
	}
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	projectName = strings.TrimSpace(string(output))
	return projectName, nil
}

func Log(a ...any) {
	_, file, line, _ := runtime.Caller(1)
	res := []any{fmt.Sprintf("%s:%d\t", file, line)}
	res = append(res, a...)
	fmt.Println(res...)
}

func Logf(format string, a ...any) {
	_, file, line, _ := runtime.Caller(1)
	prefix := fmt.Sprintf("%s:%d\t", file, line)
	fmt.Println(prefix, fmt.Sprintf(format, a...))
}

func FileExists(path string) bool {
	matches, err := filepath.Glob(path)
	if err != nil {
		return false
	}
	return len(matches) != 0
}
