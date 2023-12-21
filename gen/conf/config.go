package conf

import (
	_ "embed"
	"gopkg.in/yaml.v3"
)

var (
	//go:embed api.yaml
	ApiYaml []byte
	//go:embed layout.yaml
	LayoutYaml []byte
)

type (
	Service struct {
		Name        string   `yaml:"name"`
		Middlewares []string `yaml:"middleware"`
		Apis        []string `yaml:"apis"`
	}
	ApiConfig struct {
		Messages map[string]string `yaml:"messages"`
		Services []Service         `yaml:"services"`
	}
	Import struct {
		Context string `yaml:"context"`
		Engine  string `yaml:"engine"`
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
		Import   Import   `yaml:"import"`
		Template Template `yaml:"template"`
		Log      Log      `yaml:"log"`
	}
)

func init() {
	println(string(ApiYaml), string(LayoutYaml))
}

func initConfig() {
	var apiConf ApiConfig
	if err := yaml.Unmarshal(ApiYaml, &apiConf); err != nil {
		return
	}
	var layoutConf LayoutConfig
	if err := yaml.Unmarshal(LayoutYaml, &layoutConf); err != nil {
		return
	}
}
