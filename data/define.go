package data

type (
	ViewColumn struct {
		Key         string `json:"key"`
		Column      string `json:"column"`
		Label       string `json:"label"`
		LabelWidth  int    `json:"labelWidth"`
		Type        string `json:"type"`
		ViewType    string `json:"viewType"` // 前端展示类型
		CanCreate   bool   `json:"canCreate"`
		CanEdit     bool   `json:"canEdit"`
		CanSearch   bool   `json:"canSearch"`
		Required    bool   `json:"required"`
		Placeholder string `json:"placeholder"`
		GoType      string `json:"goType"`
		TsType      string `json:"tsType"`
	}
	Message struct {
		Name    string       `json:"name"`
		Lang    string       `json:"lang"`
		Source  string       `json:"source"` // 具体message内容
		Columns []ViewColumn `json:"columns"`
	}

	Param struct {
		Name        string // 参数名
		From        string // 参数来源: path, query, body
		Type        string // 类型
		Required    string // 是否必填
		Description string // 参数描述
	}
	Api struct {
		ServiceName string             // API所属service名
		Method      string             // HTTP方法
		Path        string             // 接口路径
		Summary     string             // swagger 文档summary
		FuncName    string             // 接口业务处理函数名
		ReqName     string             // 请求参数名
		Handler     string             // 接口处理者，包含中间件
		ReqParam    []Param            // swagger请求参数详情
		ProjectName string             // 当前项目名称，go mod name
		IdlName     string             // idl name
		Msg         Message            // msg
		MsgMap      map[string]Message // 参数表
		BodyColumns []ViewColumn       // body 参数
		Pkg         map[string]string  // layout定义的全局变量
	}
	Service struct {
		ServiceName string             // 当前service名称
		ProjectName string             // 当前项目名称，go mod name
		IdlName     string             // idl name
		Pkg         map[string]string  // 全局变量
		Apis        []Api              // 接口列表
		MsgMap      map[string]Message // 参数表
		Handlers    map[string]string  // apis 解析后的结果
	}
)
