package conf

type PkgConfig struct {
	Context    Package
	Engine     Package
	Return     Package
	ReturnType Package
	HandleFunc Package
	StatusCode Package
}

var PkgMap = map[string]PkgConfig{
	"gin": {
		Context: Package{
			Value:  "*gin.Context",
			Import: "github.com/gin-gonic/gin",
		}, Engine: Package{
			Value:  "*gin.Engine",
			Import: "github.com/gin-gonic/gin",
		}, HandleFunc: Package{
			Value:  "gin.HandlerFunc",
			Import: "github.com/gin-gonic/gin",
		}, StatusCode: Package{
			Value:  "http",
			Import: "net/http",
		},
	},
	"fiber": {
		Return:     Package{Value: "return"},
		ReturnType: Package{Value: "error"},
		Context: Package{
			Value:  "*fiber.Ctx",
			Import: "github.com/gofiber/fiber/v2",
		}, Engine: Package{
			Value:  "*fiber.App",
			Import: "github.com/gofiber/fiber/v2",
		}, HandleFunc: Package{
			Value:  "fiber.Handler",
			Import: "github.com/gofiber/fiber/v2",
		}, StatusCode: Package{
			Value:  "fiber",
			Import: "github.com/gofiber/fiber/v2",
		},
	},
	"echo": {
		Return:     Package{Value: "return"},
		ReturnType: Package{Value: "error"},
		Context: Package{
			Value:  "echo.Context",
			Import: "github.com/labstack/echo/v4",
		}, Engine: Package{
			Value:  "*echo.Echo",
			Import: "github.com/labstack/echo/v4",
		}, HandleFunc: Package{
			Value:  "echo.HandlerFunc",
			Import: "github.com/labstack/echo/v4",
		}, StatusCode: Package{
			Value:  "http",
			Import: "net/http",
		},
	},
	"hertz": {
		Context: Package{
			Value:  "*app.RequestContext",
			Import: "github.com/cloudwego/hertz/pkg/app",
		}, Engine: Package{
			Value:  "*server.Hertz",
			Import: "github.com/cloudwego/hertz/pkg/app/server",
		}, HandleFunc: Package{
			Value:  "app.HandlerFunc",
			Import: "github.com/cloudwego/hertz/pkg/app",
		}, StatusCode: Package{
			Value:  "consts",
			Import: "github.com/cloudwego/hertz/pkg/protocol/consts",
		},
	},
}
