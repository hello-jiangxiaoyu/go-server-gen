package conf

type PkgConfig struct {
	Context    Package
	Engine     Package
	Return     Package
	ReturnType Package
}

var PkgMap = map[string]PkgConfig{
	"gin": {
		Context: Package{
			Value:  "*gin.Context",
			Import: "github.com/gin-gonic/gin",
		}, Engine: Package{
			Value:  "*gin.engine",
			Import: "github.com/gin-gonic/gin",
		},
	},
	"hertz": {
		Context: Package{
			Value:  "*app.RequestContext",
			Import: "github.com/cloudwego/hertz/pkg/app",
		}, Engine: Package{
			Value:  "*server.Hertz",
			Import: "github.com/cloudwego/hertz/pkg/app/server",
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
		},
	},
}
