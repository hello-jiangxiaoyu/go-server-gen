package conf

import (
	_ "embed"
	"gopkg.in/yaml.v3"
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
