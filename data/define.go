package data

import (
	"go-server-gen/conf"
)

type (
	Param struct {
		Name        string
		From        string
		Type        string
		Required    string
		Description string
	}
	Message struct {
		Name   string  `json:"name"`
		Param  []Param `json:"param"`
		Source string  `json:"source"`
	}

	Api struct {
		ServiceName   string                  // API所属service名
		Method        string                  // HTTP方法
		Path          string                  // 接口路径
		Summary       string                  // swagger 文档summary
		FuncName      string                  // 接口业务处理函数名
		ReqName       string                  // 请求参数名
		Handlers      []string                // 接口处理者，包含中间件
		ReqParam      []Param                 // 请求参数详情
		HasMiddleware bool                    // 是否包含中间件
		ProjectName   string                  // 当前项目名称，go mod name
		IdlName       string                  // idl name
		Pkg           map[string]conf.Package // layout定义的全局变量
	}
	Service struct {
		ServiceName   string                  // 当前service名称
		ProjectName   string                  // 当前项目名称，go mod name
		IdlName       string                  // idl name
		Pkg           map[string]conf.Package // 全局变量
		HasMiddleware bool                    // api或group是否包含中间件
		Middleware    []string                // group中间件
		Apis          []Api                   // 接口列表
		Handlers      map[string]string       // apis 解析后的结果
	}
)
