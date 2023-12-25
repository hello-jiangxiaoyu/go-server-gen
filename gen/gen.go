package gen

import (
	_ "embed"
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
