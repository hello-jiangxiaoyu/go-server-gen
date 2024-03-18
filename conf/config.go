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

	IdlName    = ""
	LayoutYaml []byte
	IdlYaml    []byte
)

func ReadConfig(layoutPath, idlPath string) (err error) {
	if len(layoutPath) == 0 {
		LayoutYaml = GinYaml
	} else if layoutPath == "__ts" {
		LayoutYaml = TsFetchYaml
	} else {
		LayoutYaml, err = os.ReadFile(layoutPath)
		if err != nil {
			return utils.WithMessage(err, layoutPath)
		}
	}

	if len(idlPath) > 0 {
		IdlYaml, err = os.ReadFile(idlPath)
		if err != nil {
			return utils.WithMessage(err, idlPath)
		}
	}
	return nil
}

func GetConfig(serverType, logType, layoutPath, idlPath string) (LayoutConfig, Idl, error) {
	if err := ReadConfig(layoutPath, idlPath); err != nil {
		return LayoutConfig{}, Idl{}, err
	}
	var apiConf Idl
	if err := yaml.Unmarshal(IdlYaml, &apiConf); err != nil {
		return LayoutConfig{}, Idl{}, err
	}
	var layoutConf LayoutConfig
	if err := yaml.Unmarshal(LayoutYaml, &layoutConf); err != nil {
		return LayoutConfig{}, Idl{}, err
	}
	if err := ModifyConfig(&layoutConf, serverType, logType); err != nil {
		return LayoutConfig{}, Idl{}, err
	}
	return layoutConf, apiConf, nil
}

func GetLayoutConfig(serverType, logType, layoutPath string) (LayoutConfig, error) {
	if err := ReadConfig(layoutPath, ""); err != nil {
		return LayoutConfig{}, err
	}
	var layoutConf LayoutConfig
	if err := yaml.Unmarshal(LayoutYaml, &layoutConf); err != nil {
		return LayoutConfig{}, err
	}
	if err := ModifyConfig(&layoutConf, serverType, logType); err != nil {
		return LayoutConfig{}, err
	}
	return layoutConf, nil
}

func ModifyConfig(layoutConf *LayoutConfig, serverType, logType string) error {
	if layoutConf.Pkg == nil {
		layoutConf.Pkg = make(map[string]any)
	}

	projectName, err := utils.GetProjectName()
	if err != nil {
		return err
	}
	layoutConf.ProjectName = projectName
	layoutConf.IdlName = IdlName
	layoutConf.ServerType = serverType
	layoutConf.LogType = logType
	svc, ok := PkgMap[serverType]
	if ok {
		for key, value := range svc {
			layoutConf.Pkg[key] = value
		}
	}
	return nil
}

func getFileName(path string) string {
	filenameWithExtension := filepath.Base(path)
	fileExtension := filepath.Ext(filenameWithExtension)
	filename := strings.TrimSuffix(filenameWithExtension, fileExtension)
	return filename
}
