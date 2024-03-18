package conf

// IDL配置文件
type (
	Service struct {
		Name        string   `yaml:"name"`
		Middlewares []string `yaml:"middlewares"` // 中间件名称
		Apis        []string `yaml:"apis"`
	}
	Idl struct {
		Messages string    `yaml:"messages"`
		Ts       string    `yaml:"ts"`
		Services []Service `yaml:"services"`
	}
)

// Layout 配置
type (
	Template struct {
		Name       string `yaml:"name"`        // 模板名
		Path       string `yaml:"path"`        // 文件写入路径
		Write      string `yaml:"write"`       // 写入方式
		Handler    string `yaml:"handler"`     // 函数或结构体
		HandlerKey string `yaml:"handler-key"` // 写文件关键字
		Body       string `yaml:"body"`        // 文件内容
	}
	LayoutConfig struct {
		Pkg             map[string]any `yaml:"pkg"`      // 全局变量
		ApiTemplate     []Template     `yaml:"api"`      // api模板配置（todo: 实现api级别模板解析）
		ServiceTemplate []Template     `yaml:"service"`  // service模板配置
		GlobalTemplate  []Template     `yaml:"global"`   // global模板配置
		MessageTemplate []Template     `yaml:"messages"` // 请求参数模板配置

		ProjectName string `yaml:"-"` // 项目名
		IdlName     string `yaml:"-"` // idl名
		LogType     string `yaml:"-"` // 日志类型zap或zero
		ServerType  string `yaml:"-"` // server类型, gin, echo, fiber, hertz
	}
)
