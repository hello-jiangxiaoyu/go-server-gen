package internal

const (
	ginMain = `package main
import (
	"github.com/gin-gonic/gin"
	"{{.ProjectName}}/biz"
	"os"
)

func main() {
	server := gin.New()
	biz.Register(server)
	if err := server.Run(":1234"); err != nil {
		println("server run err: ", err.Error())
		os.Exit(1)
	}
}`

	fiberMain = `package main
import (
	"github.com/gofiber/fiber/v2"
	"{{.ProjectName}}/biz"
	"os"
)

func main() {
	server := fiber.New()
	biz.Register(server)
	if err := server.Listen(":1234"); err != nil {
		println("server run err: ", err.Error())
		os.Exit(1)
	}
}`

	echoMain = `package main
import (
	"github.com/labstack/echo/v4"
	"{{.ProjectName}}/biz"
	"os"
)

func main() {
	server := echo.New()
	biz.Register(server)
	if err := server.Start(":1234"); err != nil {
		println("server run err: ", err.Error())
		os.Exit(1)
	}
}`

	hertzMain = `package main
import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"{{.ProjectName}}/biz"
)

func main() {
	svc := server.New()
	biz.Register(svc)
	svc.Spin()
}`
)

var MainCodeMap = map[string]string{
	"gin":   ginMain,
	"echo":  echoMain,
	"fiber": fiberMain,
	"hertz": hertzMain,
}

const RegisterCode = `// Code generated by go-server-gen. DO NOT EDIT.
      
package biz
import (
  "{{.Pkg.EngineImport}}"
)
func Register(e {{.Pkg.EngineType}}) {
  //INSERT_POINT: DO NOT DELETE THIS LINE!

}
`

const Readme = `# {{.ProjectName}}

` + "生成文档\n\n```bash\nswag init -o pkg/docs\n```\n"

const DockerFile = `##
## build
##
FROM golang:1.21-alpine as builder
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN go mod download

RUN  go build -ldflags="-s -w" -o server .


##
## deploy
##
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server ./

EXPOSE 1234
ENTRYPOINT ./server
`

const Makefile = `
BINARY_NAME = server

build:
	go build -ldflags="-s -w" -o $(BINARY_NAME) .

clean:
	rm -f $(BINARY_NAME)

.PHONY: build clean
`

const GitIgnore = `.DS_Store
.idea
.vscode
.DS_Store

*.log
*.exe
*.so
*.dll
*.lib
`
