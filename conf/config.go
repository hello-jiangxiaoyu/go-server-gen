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

	//go:embed data.yaml
	data []byte

	IdlName    = ""
	LayoutYaml []byte
	Layout     LayoutConfig
	IdlYaml    []byte
	Idl        IdlConfig
	ConstData  ConstStruct
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

func GetConfig(serverType, logType, layoutPath, idlPath string) (LayoutConfig, IdlConfig, error) {
	if err := ReadConfig(layoutPath, idlPath); err != nil {
		return LayoutConfig{}, IdlConfig{}, err
	}

	if err := yaml.Unmarshal(IdlYaml, &Idl); err != nil {
		return LayoutConfig{}, IdlConfig{}, utils.WithMessage(err, "idl yaml unmarshal err")
	}
	if err := yaml.Unmarshal(LayoutYaml, &Layout); err != nil {
		return LayoutConfig{}, IdlConfig{}, utils.WithMessage(err, "layout yaml unmarshal err")
	}
	if err := yaml.Unmarshal(data, &ConstData); err != nil {
		return LayoutConfig{}, IdlConfig{}, utils.WithMessage(err, "data yaml unmarshal err")
	}

	if err := ModifyConfig(&Layout, serverType, logType); err != nil {
		return LayoutConfig{}, IdlConfig{}, err
	}
	return Layout, Idl, nil
}

func GetLayoutConfig(serverType, logType, layoutPath string) (LayoutConfig, error) {
	if err := ReadConfig(layoutPath, ""); err != nil {
		return LayoutConfig{}, err
	}

	if err := yaml.Unmarshal(LayoutYaml, &Layout); err != nil {
		return LayoutConfig{}, utils.WithMessage(err, "layout yaml unmarshal err")
	}
	if err := yaml.Unmarshal(data, &ConstData); err != nil {
		return LayoutConfig{}, utils.WithMessage(err, "data yaml unmarshal err")
	}

	Layout.Data = ConstData
	if err := ModifyConfig(&Layout, serverType, logType); err != nil {
		return LayoutConfig{}, err
	}
	return Layout, nil
}

func ModifyConfig(layoutConf *LayoutConfig, serverType, logType string) error {
	if layoutConf.Pkg == nil {
		layoutConf.Pkg = make(map[string]string)
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
