package conf

import (
	_ "embed"
	"go-server-gen/utils"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

var (
	//go:embed layout.yaml
	GinYaml []byte

	//go:embed ts-fetch.yaml
	TsFetchYaml []byte

	IdlName    = "idl"
	layoutYaml []byte
)

func GetConfig(serverType, logType, layoutPath, idlPath string) (LayoutConfig, Idl, error) {
	var apiConf Idl
	idlYaml, err := os.ReadFile(idlPath)
	if err != nil {
		return LayoutConfig{}, Idl{}, err
	}
	if err = yaml.Unmarshal(idlYaml, &apiConf); err != nil {
		return LayoutConfig{}, Idl{}, err
	}

	IdlName = getFileName(idlPath)
	layoutConf, err := GetLayoutConfig(serverType, logType, layoutPath)
	if err != nil {
		return LayoutConfig{}, Idl{}, err
	}

	return layoutConf, apiConf, nil
}

func GetLayoutConfig(serverType, logType, layoutPath string) (layoutConf LayoutConfig, err error) {
	if layoutPath == "" {
		layoutYaml = GinYaml
	} else if layoutPath == "__ts" {
		layoutYaml = TsFetchYaml
	} else {
		layoutYaml, err = os.ReadFile(layoutPath)
		if err != nil {
			return LayoutConfig{}, err
		}
	}

	if err = yaml.Unmarshal(layoutYaml, &layoutConf); err != nil {
		return LayoutConfig{}, err
	}

	if layoutConf.Pkg == nil {
		layoutConf.Pkg = make(map[string]Package)
	}

	projectName, err := utils.GetProjectName()
	if err != nil {
		return LayoutConfig{}, err
	}
	layoutConf.ProjectName = projectName
	layoutConf.IdlName = IdlName
	layoutConf.ServerType = serverType
	layoutConf.LogType = logType
	svc, ok := PkgMap[serverType]
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

func getFileName(path string) string {
	filenameWithExtension := filepath.Base(path)
	fileExtension := filepath.Ext(filenameWithExtension)
	filename := strings.TrimSuffix(filenameWithExtension, fileExtension)
	return filename
}
