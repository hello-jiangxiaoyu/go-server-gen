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
		HandlerKey string `yaml:"handler-key"`
		Body       string `yaml:"body"`
	}
	LayoutConfig struct {
		Pkg             map[string]Package `yaml:"package"`  // 全局变量
		ApiTemplate     []Template         `yaml:"api"`      // api模板配置（todo: 实现api级别模板解析）
		ServiceTemplate []Template         `yaml:"service"`  // service模板配置
		GlobalTemplate  []Template         `yaml:"global"`   // global模板配置
		MessageTemplate []Template         `yaml:"messages"` // 请求参数模板配置
		ProjectName     string             `yaml:"-"`
		IdlName         string             `yaml:"-"`
	}
)
