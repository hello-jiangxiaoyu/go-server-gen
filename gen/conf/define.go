package conf

// IDL配置文件
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
)

// Layout 配置
type (
	Package struct {
		Name   string `yaml:"name"`
		Import string `yaml:"import"`
	}
	Template struct {
		Name       string `yaml:"name"`
		Path       string `yaml:"path"`
		Write      string `yaml:"write"`
		Handler    string `yaml:"handler"`
		HandlerTpl string `yaml:"handler-tpl"`
		Body       string `yaml:"body"`
	}
	LayoutConfig struct {
		Pkg         map[string]Package `yaml:"package"`
		Service     []Template         `yaml:"service"`
		Messages    []Template         `yaml:"messages"`
		ProjectName string             `yaml:"-"`
	}
)
