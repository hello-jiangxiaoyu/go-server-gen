package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"
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
func WrapError(err1 error, err2 error) error {
	if err1 == nil && err2 == nil {
		return nil
	} else if err1 == nil {
		return err2
	} else if err2 == nil {
		return err1
	}

	return errors.New(err2.Error() + ": " + err1.Error())
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

func Log(a ...any) {
	_, file, line, _ := runtime.Caller(1)
	res := []any{fmt.Sprintf("%v %s:%d", time.Now().Format("2006-01-02 15:04:05"), file, line)}
	res = append(res, a...)
	fmt.Println(res...)
}

func Logf(format string, a ...any) {
	_, file, line, _ := runtime.Caller(1)
	prefix := fmt.Sprintf("%v %s:%d", time.Now().Format("2006-01-02 15:04:05"), file, line)
	fmt.Println(prefix, fmt.Sprintf(format, a...))
}
