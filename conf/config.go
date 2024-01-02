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

func GetConfig(server string) (LayoutConfig, Idl, error) {
	var apiConf Idl
	if err := yaml.Unmarshal(IdlYaml, &apiConf); err != nil {
		return LayoutConfig{}, Idl{}, err
	}

	layoutConf, err := GetLayoutConfig(server)
	if err != nil {
		return LayoutConfig{}, Idl{}, err
	}

	return layoutConf, apiConf, nil
}

func GetLayoutConfig(server string) (LayoutConfig, error) {
	var layoutConf LayoutConfig
	if err := yaml.Unmarshal(LayoutYaml, &layoutConf); err != nil {
		return LayoutConfig{}, err
	}

	projectName, err := utils.GetProjectName()
	if err != nil {
		return LayoutConfig{}, err
	}
	layoutConf.ProjectName = projectName
	layoutConf.IdlName = IdlName
	svc, ok := PkgMap[server]
	if ok {
		layoutConf.Pkg["Context"] = svc.Context
		layoutConf.Pkg["Engine"] = svc.Engine
		layoutConf.Pkg["Return"] = svc.Return
		layoutConf.Pkg["ReturnType"] = svc.ReturnType
		layoutConf.Pkg["HandleFunc"] = svc.HandleFunc
		layoutConf.Pkg["StatusCode"] = svc.StatusCode
	}
	return layoutConf, nil
}
