package conf

import (
	_ "embed"
	"gopkg.in/yaml.v3"
)

// 配置文件
type (
	Service struct {
		Name        string   `yaml:"name"`
		Middlewares []string `yaml:"middleware"`
		Apis        []string `yaml:"apis"`
	}
	Idl struct {
		Messages string    `yaml:"messages"`
		Services []Service `yaml:"services"`
	}

	Template struct {
		Register   string `yaml:"register"`
		Controller string `yaml:"controller"`
		Api        string `yaml:"api"`
	}
	Log struct {
		Type    string   `yaml:"type"`
		Imports []string `yaml:"imports"`
	}
	LayoutConfig struct {
		Server struct {
			ContextName   string `yaml:"context-name"`
			ContextImport string `yaml:"context-import"`
			EngineName    string `yaml:"engine-name"`
			EngineImport  string `yaml:"engine-import"`
		} `yaml:"server"`
		Template Template `yaml:"template"`
		Log      Log      `yaml:"log"`
	}
)

func GetConfig(layoutYaml, idlYaml []byte) (*LayoutConfig, *Idl, error) {
	var apiConf Idl
	if err := yaml.Unmarshal(idlYaml, &apiConf); err != nil {
		return nil, nil, err
	}
	var layoutConf LayoutConfig
	if err := yaml.Unmarshal(layoutYaml, &layoutConf); err != nil {
		return nil, nil, err
	}
	return &layoutConf, &apiConf, nil
}
