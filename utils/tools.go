package utils

import (
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
