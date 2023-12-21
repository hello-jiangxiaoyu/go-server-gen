package conf

type (
	Param struct {
		Name        string
		From        string
		Type        string
		Required    string
		Description string
	}
	Api struct {
		Method     string
		Path       string
		DocPath    string
		Summary    string
		FuncName   string
		ReqName    string
		Tag        string
		Context    string
		ReqParam   []Param
		Middleware []string
	}
	Group struct {
		Apis          []Api
		Import        []string
		Middleware    []string
		ApiTemplate   string
		GroupTemplate string
	}
)

func ConfigToGen(layout *LayoutConfig, api *ApiConfig) {

}
