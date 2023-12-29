package conf

import (
	_ "embed"
	"go-server-gen/utils"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	//go:embed gin.yaml
	GinYaml []byte
	//go:embed fiber.yaml
	FiberYaml []byte
	//go:embed echo.yaml
	EchoYaml []byte
	//go:embed hertz.yaml
	HertzYaml []byte
	//go:embed ts-fetch.yaml
	TsFetchYaml []byte

	LayoutYaml []byte

	//go:embed test-idl.yaml
	IdlYaml []byte
	IdlName = "idl"
)

func init() {
	LayoutYaml = GinYaml
}

func InitConfig(layoutPath, idlPath string) error {
	var err error
	if len(layoutPath) != 0 {
		IdlName = strings.Split(idlPath, ".")[0]
		LayoutYaml, err = os.ReadFile(layoutPath)
		if err != nil {
			return err
		}
	}

	if len(idlPath) != 0 {
		IdlYaml, err = os.ReadFile(idlPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetConfig() (*LayoutConfig, *Idl, error) {
	var apiConf Idl
	if err := yaml.Unmarshal(IdlYaml, &apiConf); err != nil {
		return nil, nil, err
	}
	var layoutConf LayoutConfig
	if err := yaml.Unmarshal(LayoutYaml, &layoutConf); err != nil {
		return nil, nil, err
	}

	projectName, err := utils.GetProjectName()
	if err != nil {
		return nil, nil, err
	}
	layoutConf.ProjectName = projectName
	layoutConf.IdlName = IdlName
	return &layoutConf, &apiConf, nil
}
