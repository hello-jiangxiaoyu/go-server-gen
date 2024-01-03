package internal

const (
	recoverMiddleware = `package middleware
      import (
        {{ if eq .Pkg.Context.Value "*app.RequestContext" -}}
        "context"
        {{- end}}
        "errors"
        "fmt"
        "net"
        "os"
        "runtime"
        "strings"
        "{{.Pkg.Context.Import}}"
        "{{.ProjectName}}/{{.Pkg.Resp.Import}}"
      )

      func stackInfo(msg string, skip int) string {
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

      // Recovery panic处理
      func Recovery({{- if eq .Pkg.Context.Value "echo.Context" -}}next {{.Pkg.HandleFunc.Value}}{{- end }}) {{.Pkg.HandleFunc.Value}} {
        return func(
        {{- if eq .Pkg.Context.Value "*app.RequestContext" -}}
        ctx context.Context, c {{.Pkg.Context.Value}}
        {{- else -}}
        c {{.Pkg.Context.Value}}
        {{- end -}}
        ) {{.Pkg.ReturnType.Value}} {
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
              println(stackInfo(req, 3))
              if !brokenPipe {
                {{.Pkg.Resp.Value}}.ErrorPanic(c)
              }
            }
          }()
          {{ if eq .Pkg.Context.Value "*app.RequestContext" -}}
          c.Next(ctx)
          {{- else if eq .Pkg.Context.Value "echo.Context" -}}
          return next(c)
          {{- else -}}
          {{.Pkg.Return.Value}} c.Next()
          {{- end }}
        }
      }`
)

var MiddlewareMap = map[string]string{
	"recover": recoverMiddleware,
}
