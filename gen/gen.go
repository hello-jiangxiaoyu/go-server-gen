package gen

import (
	_ "embed"
	"go-server-gen/gen/conf"
	"go-server-gen/gen/data"
	"go-server-gen/gen/parse"
	"go-server-gen/utils"
	"os"
)

var (
	//go:embed idl.yaml
	IdlYaml []byte
	//go:embed layout.yaml
	LayoutYaml []byte
)

func InitConfig(layoutPath, idlPath string) error {
	var err error
	if len(layoutPath) != 0 {
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

func Execute() error {
	layout, idl, err := conf.GetConfig(LayoutYaml, IdlYaml)
	if err != nil {
		return utils.WithMessage(err, "unmarshal yaml config err")
	}

	groups, _, err := data.ConfigToData(layout, idl)
	if err != nil {
		return utils.WithMessage(err, "config to data err")
	}

	res, err := parse.GenServiceCode(layout, groups)
	if err != nil {
		return utils.WithMessage(err, "gen code err")
	}
	for _, v := range res {
		println(v.File, "\n"+v.Code)
	}
	return nil
}
