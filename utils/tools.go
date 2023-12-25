package utils

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

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

// DeferErr handle defer function err
func DeferErr(errFunc func() error) {
	if err := errFunc(); err != nil {
		fmt.Println("### Defer err: ", err)
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
