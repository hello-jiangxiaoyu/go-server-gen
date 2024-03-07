package internal

const (
	recoverMiddleware = `package middleware
      import (
        {{ if eq .Pkg.ContextType "*app.RequestContext" -}}
        "context"
        {{- end}}
        "errors"
        "fmt"
        "net"
        "os"
        "runtime"
        "strings"
        "{{.Pkg.ContextImport}}"
        "{{.ProjectName}}/{{.Pkg.Resp}}"
      )

      func stackInfo(msg string, skip int) string {
        for i := skip; i < 30; i++ {
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
      func Recovery({{- if eq .Pkg.ContextType "echo.Context" -}}next {{.Pkg.HandleFuncType}}{{- end }}) {{.Pkg.HandleFuncType}} {
        return func(
        {{- if eq .Pkg.ContextType "*app.RequestContext" -}}
        ctx context.Context, c {{.Pkg.ContextType}}
        {{- else -}}
        c {{.Pkg.ContextType}}
        {{- end -}}
        ) {{.Pkg.ReturnType}} {
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
                response.ErrorPanic(c)
              }
            }
          }()
          {{ if eq .Pkg.ContextType "*app.RequestContext" -}}
          c.Next(ctx)
          {{- else if eq .Pkg.ContextType "echo.Context" -}}
          return next(c)
          {{- else -}}
          {{.Pkg.Return}} c.Next()
          {{- end }}
        }
      }`

	requestMiddleware = `package middleware
	import (
        {{ if eq .Pkg.ContextType "*app.RequestContext" -}}
        "context"
        {{- end}}
        "{{.Pkg.ContextImport}}"
		"math/rand"
		"strconv"
	)
	
	func GenerateRequestID({{- if eq .Pkg.ContextType "echo.Context" -}}next {{.Pkg.HandleFuncType}}{{- end }}) {{.Pkg.HandleFuncType}} {
		return func(
			{{- if eq .Pkg.ContextType "*app.RequestContext" -}}
			ctx context.Context, c {{.Pkg.ContextType}}
			{{- else -}}
			c {{.Pkg.ContextType}}
			{{- end -}}
			) {{.Pkg.ReturnType}} {
			{{ if eq .Pkg.ContextType "*gin.Context" -}}
			requestID := c.GetHeader("X-Request-ID")
			{{- else if eq .Pkg.ContextType "*fiber.Ctx" -}}
			requestID := ""
			requestIDs := c.GetReqHeaders()["X-Request-Id"]
			if len(requestIDs) == 0 {
				requestID = requestIDs[0]
			}
			{{- else if eq .Pkg.ContextType "echo.Context" -}}
			requestID := c.Request().Header.Get("X-Request-Id")
			{{- else if eq .Pkg.ContextType "*app.RequestContext" -}}
			requestID := c.Request.Header.Get("X-Request-Id")
			{{- else -}}
			panic(nil)
			{{- end }}
			if requestID == "" {
				requestID = strconv.FormatInt(rand.Int63(), 10)
			}
			c.Set("requestID", requestID)
			{{- if .Pkg.Return }}
			return nil
			{{- end }}
		}
	}`
)

var MiddlewareMap = map[string]string{
	"recover": recoverMiddleware,
	"request": requestMiddleware,
}
