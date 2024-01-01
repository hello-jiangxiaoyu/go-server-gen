package source

import (
	"embed"
	"go-server-gen/conf"
)

var (
	//go:embed internal
	code embed.FS
	//go:embed server.yaml
	serverTpl []byte
)

type ResponsePackage struct {
	BindCode     conf.Package
	ResponseCode conf.Package
	Context      conf.Package
	Return       conf.Package
	ReturnType   conf.Package
	HandleFunc   conf.Package
	Code         conf.Package
}

var responseMap = map[string]ResponsePackage{
	"gin": {
		BindCode:     conf.Package{Value: GetEmbedContent("internal/gin/bind.go")},
		ResponseCode: conf.Package{Value: GetEmbedContent("internal/gin/response.go")},
		HandleFunc:   conf.Package{Value: "gin.HandlerFunc"},
		Context: conf.Package{
			Value:  "*gin.Context",
			Import: `"github.com/gin-gonic/gin"`,
		},
		Code: conf.Package{
			Value:  "http",
			Import: `"net/http"`,
		},
	},
	"hertz": {
		BindCode:     conf.Package{Value: GetEmbedContent("internal/hertz/bind.go")},
		ResponseCode: conf.Package{Value: GetEmbedContent("internal/hertz/response.go")},
		HandleFunc:   conf.Package{Value: "app.HandlerFunc"},
		Context: conf.Package{
			Value:  "*app.RequestContext",
			Import: `"github.com/cloudwego/hertz/pkg/app"`,
		},
		Code: conf.Package{
			Value:  "consts",
			Import: `"github.com/cloudwego/hertz/pkg/protocol/consts"`,
		},
	},
	"fiber": {
		BindCode:     conf.Package{Value: GetEmbedContent("internal/fiber/bind.go")},
		ResponseCode: conf.Package{Value: GetEmbedContent("internal/fiber/response.go")},
		HandleFunc:   conf.Package{Value: "fiber.Handler"},
		Context: conf.Package{
			Value:  "*fiber.Ctx",
			Import: `"github.com/gofiber/fiber/v2"`,
		},
		Code:       conf.Package{Value: "fiber"},
		Return:     conf.Package{Value: "return"},
		ReturnType: conf.Package{Value: "error"},
	},
	"echo": {
		BindCode:     conf.Package{Value: GetEmbedContent("internal/echo/bind.go")},
		ResponseCode: conf.Package{Value: GetEmbedContent("internal/echo/response.go")},
		HandleFunc:   conf.Package{Value: "echo.HandlerFunc"},
		Context: conf.Package{
			Value:  "echo.Context",
			Import: `"github.com/labstack/echo/v4"`,
		},
		Code: conf.Package{
			Value:  "http",
			Import: `"net/http"`,
		},
		Return:     conf.Package{Value: "return"},
		ReturnType: conf.Package{Value: "error"},
	},
}
