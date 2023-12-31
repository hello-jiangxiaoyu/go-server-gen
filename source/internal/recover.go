package middleware

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

func stackInfo(msg string, err any, skip int) string {
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		// 不打印第三方库栈信息
		if !strings.Contains(file, "github.com/") && !strings.Contains(file, "gorm.io/") && !strings.Contains(file, "net/http") {
			msg += fmt.Sprintf("\n\t%s:%d %s", file, line, runtime.FuncForPC(pc).Name())
		}
	}
	return msg + "\n\n"
}

// MyRecoveryHandler panic处理
func MyRecoveryHandler(ctx context.Context, c *app.RequestContext) {
	defer func() {
		if err := recover(); err != nil {
			// Check for a broken connection, as it is not really a
			var brokenPipe bool
			if ne, ok := err.(*net.OpError); ok {
				var se *os.SyscallError
				if errors.As(ne, &se) {
					seStr := strings.ToLower(se.Error())
					if strings.Contains(seStr, "broken pipe") || strings.Contains(seStr, "connection reset by peer") {
						brokenPipe = true
					}
				}
			}
			req := fmt.Sprintf("panic recovered: %s; method:%s path:%s", err, c.Method(), c.Path())
			println(stackInfo(req, err, 3))
			if !brokenPipe {
				// response.ErrorPanic(c)
			}
		}
	}()
	c.Next(ctx)
}
