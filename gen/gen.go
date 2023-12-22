package gen

import _ "embed"

var (
	//go:embed api.yaml
	IdlYaml []byte
	//go:embed layout.yaml
	LayoutYaml []byte
)
