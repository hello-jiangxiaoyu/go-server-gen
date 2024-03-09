package internal

// 失败响应
const (
	ErrorRequest = `package response
      import (
        "{{.Pkg.ContextImport}}"
        "{{.Pkg.StatusCodeImport}}"
      )

      // ErrorRequest 请求参数错误
      func ErrorRequest(c {{.Pkg.ContextType}}, err error) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusBadRequest, CodeRequestPara, err, "invalidate para")
      }

      // ErrorRequestWithMsg 请求参数错误
      func ErrorRequestWithMsg(c {{.Pkg.ContextType}}, err error, msg string) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusBadRequest, CodeRequestPara, err, msg)
      }

      // ErrorForbidden 无权访问
      func ErrorForbidden(c {{.Pkg.ContextType}}, msg string) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusForbidden, CodeForbidden, nil, msg)
      }

      // ErrorInvalidateToken token 无效
      func ErrorInvalidateToken(c {{.Pkg.ContextType}}, err error) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusForbidden, CodeInvalidToken, err, "invalidate token")
      }

      // ErrorNoLogin 用户未登录
      func ErrorNoLogin(c {{.Pkg.ContextType}}) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusUnauthorized, CodeNotLogin, nil, "user not login")
      }
      
      // ErrorNoRoute 
      func ErrorNoRoute(c {{.Pkg.ContextType}}) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusNotFound, CodeNoSuchRoute, nil, "no such router")
      }`

	ErrorSql = `package response
      import (
        "strings"
        "gorm.io/gorm"
        "{{.Pkg.ContextImport}}"
        "{{.Pkg.StatusCodeImport}}"
      )

      // ErrorUpdate SQL更新失败
      func ErrorUpdate(c {{.Pkg.ContextType}}, err error, respMsg string) {{.Pkg.ReturnType}} {
        if err != nil && strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
          {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusConflict, CodeSqlModifyDuplicate, err, respMsg)
        } else {
          {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusInternalServerError, CodeSqlModify, err, respMsg)
        }
      }

      // ErrorCreate SQL创建失败
      func ErrorCreate(c {{.Pkg.ContextType}}, err error, respMsg string) {{.Pkg.ReturnType}} {
        if err != nil && strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
          {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusConflict, CodeSqlCreateDuplicate, err, respMsg)
        } else {
          {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusInternalServerError, CodeSqlCreate, err, respMsg)
        }
      }

      // ErrorSelect 数据库查询错误
      func ErrorSelect(c {{.Pkg.ContextType}}, err error, respMsg string) {{.Pkg.ReturnType}} {
        if err != nil && strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) { // gorm First操作record not found
          {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusNotFound, CodeSqlSelectNotFound, err, respMsg)
        } else {
          {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusInternalServerError, CodeSqlSelect, err, respMsg)
        }
      }

      // ErrorDelete SQL删除失败
      func ErrorDelete(c {{.Pkg.ContextType}}, err error, respMsg string) {{.Pkg.ReturnType}} {
        if err != nil && strings.Contains(err.Error(), gorm.ErrForeignKeyViolated.Error()) { // 外键依赖导致无法删除
          {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusConflict, CodeSqlDeleteForKey, err, respMsg)
        } else {
          {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusInternalServerError, CodeSqlDelete, err, respMsg)
        }
      }`

	ErrorUnknown = `package response
      import (
        "{{.Pkg.ContextImport}}"
        "{{.Pkg.StatusCodeImport}}"
      )

      // ErrorUnknown 未知错误
      func ErrorUnknown(c {{.Pkg.ContextType}}, err error, respMsg string) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusInternalServerError, CodeUnknown, err, respMsg)
      }

      // ErrorNotFound 资源未找到
      func ErrorNotFound(c {{.Pkg.ContextType}}, err error, respMsg string) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusInternalServerError, CodeNotFound, err, respMsg)
      }

      func ErrorSaveSession(c {{.Pkg.ContextType}}, err error) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusInternalServerError, CodeSaveSession, err, "failed to save session")
      }

      func ErrorPanic(c {{.Pkg.ContextType}}, err error) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusInternalServerError, CodePanic, err, "server panic")
      }

      // ErrorSendRequest 发送 http 请求失败
      func ErrorSendRequest(c {{.Pkg.ContextType}}, err error, respMsg string) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} errorResponse(c, {{.Pkg.StatusCodePackage}}.StatusInternalServerError, CodeSendRequest, err, respMsg)
      }`
)

// 成功响应
const (
	Success = `package response

      import (
        "{{.Pkg.ContextImport}}"
        "{{.Pkg.StatusCodeImport}}"
      )

      func success(c {{.Pkg.ContextType}}, data any, total int) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} response(c, {{.Pkg.StatusCodePackage}}.StatusOK, CodeSuccess, nil, MsgSuccess, data, total)
      }

      func Success(c {{.Pkg.ContextType}}) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} response(c, {{.Pkg.StatusCodePackage}}.StatusOK, CodeSuccess, nil, MsgSuccess, struct{}{}, 0)
      }
      func SuccessWithData(c {{.Pkg.ContextType}}, data any) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} response(c, {{.Pkg.StatusCodePackage}}.StatusOK, CodeSuccess, nil, MsgSuccess, data, 0)
      }
      func SuccessWithArrayData(c {{.Pkg.ContextType}}, data any, total int) {{.Pkg.ReturnType}} {
        {{.Pkg.Return}} response(c, {{.Pkg.StatusCodePackage}}.StatusOK, CodeSuccess, nil, MsgSuccess, data, total)
      }`

	ServiceCode = `package response
      const (
        MsgSuccess = ""
      )

      const (
        CodeSuccess = 200
        CodeAccept  = 202
      )

      // 系统相关错误码
      const (
        CodeNoSuchRoute = 1000 // 路由不存在
        CodeRequestPara = 1001 // 请求参数错误
        CodeForbidden   = 1002 // 无权访问
      )

      // 业务相关错误码
      const (
        CodeNoSuchHost   = 2000 // 非法host
        CodeNotLogin     = 2001 // 用户未登录
        CodeInvalidToken = 2002 // 非法token
      )

      // SQL相关错误码
      const (
        CodeSqlSelect          = iota + 3000 // 查询失败
        CodeSqlSelectNotFound                // 不存在该数据
        CodeSqlModify                        // 修改失败
        CodeSqlModifyDuplicate               // 数据冲突导致修改失败
        CodeSqlCreate                        // 创建失败
        CodeSqlCreateDuplicate               // 数据重复导致创建失败
        CodeSqlDelete                        // 删除失败
        CodeSqlDeleteForKey                  // 外键依赖导致删除失败
      )

      // error相关错误码
      const (
        CodeServerPanic = iota + 5000 // 发生panic
        CodeUnknown                   // 未知错误
		CodePanic                     // panic
        CodeNotFound                  // 未找到
        CodeSaveSession               // session保存错误
        CodeSendRequest               // 发送http请求错误
      )`
)
