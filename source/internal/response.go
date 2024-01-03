package internal

// 失败响应
const (
	ErrorRequest = `package response
      import (
        "{{.Pkg.Context.Import}}"
        "{{.Pkg.StatusCode.Import}}"
      )

      // ErrorRequest 请求参数错误
      func ErrorRequest(c {{.Pkg.Context.Value}}, err error) {{.Pkg.ReturnType.Value}} {
        {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusBadRequest, CodeRequestPara, err, "invalidate para")
      }

      // ErrorRequestWithMsg 请求参数错误
      func ErrorRequestWithMsg(c {{.Pkg.Context.Value}}, err error, msg string) {{.Pkg.ReturnType.Value}} {
        {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusBadRequest, CodeRequestPara, err, msg)
      }

      // ErrorForbidden 无权访问
      func ErrorForbidden(c {{.Pkg.Context.Value}}, msg string) {{.Pkg.ReturnType.Value}} {
        {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusForbidden, CodeForbidden, nil, msg)
      }

      // ErrorInvalidateToken token 无效
      func ErrorInvalidateToken(c {{.Pkg.Context.Value}}, err error) {{.Pkg.ReturnType.Value}} {
        {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusForbidden, CodeInvalidToken, err, "invalidate token")
      }

      // ErrorNoLogin 用户未登录
      func ErrorNoLogin(c {{.Pkg.Context.Value}}) {{.Pkg.ReturnType.Value}} {
        {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusUnauthorized, CodeNotLogin, nil, "user not login")
      }`

	ErrorSql = `package response
      import (
        "strings"
        "gorm.io/gorm"
        "{{.Pkg.Context.Import}}"
        "{{.Pkg.StatusCode.Import}}"
      )

      // ErrorUpdate SQL更新失败
      func ErrorUpdate(c {{.Pkg.Context.Value}}, err error, respMsg string) {{.Pkg.ReturnType.Value}} {
        if err != nil && strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
          {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusConflict, CodeSqlModifyDuplicate, err, respMsg)
        } else {
          {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusInternalServerError, CodeSqlModify, err, respMsg)
        }
      }

      // ErrorCreate SQL创建失败
      func ErrorCreate(c {{.Pkg.Context.Value}}, err error, respMsg string) {{.Pkg.ReturnType.Value}} {
        if err != nil && strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
          {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusConflict, CodeSqlCreateDuplicate, err, respMsg)
        } else {
          {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusInternalServerError, CodeSqlCreate, err, respMsg)
        }
      }

      // ErrorSelect 数据库查询错误
      func ErrorSelect(c {{.Pkg.Context.Value}}, err error, respMsg string) {{.Pkg.ReturnType.Value}} {
        if err != nil && strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) { // gorm First操作record not found
          {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusNotFound, CodeSqlSelectNotFound, err, respMsg)
        } else {
          {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusInternalServerError, CodeSqlSelect, err, respMsg)
        }
      }

      // ErrorDelete SQL删除失败
      func ErrorDelete(c {{.Pkg.Context.Value}}, err error, respMsg string) {{.Pkg.ReturnType.Value}} {
        if err != nil && strings.Contains(err.Error(), gorm.ErrForeignKeyViolated.Error()) { // 外键依赖导致无法删除
          {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusConflict, CodeSqlDeleteForKey, err, respMsg)
        } else {
          {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusInternalServerError, CodeSqlDelete, err, respMsg)
        }
      }`

	ErrorUnknown = `package response
      import (
        "{{.Pkg.Context.Import}}"
        "{{.Pkg.StatusCode.Import}}"
      )

      // ErrorUnknown 未知错误
      func ErrorUnknown(c {{.Pkg.Context.Value}}, err error, respMsg string) {{.Pkg.ReturnType.Value}} {
        {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusInternalServerError, CodeUnknown, err, respMsg)
      }

      // ErrorNotFound 资源未找到
      func ErrorNotFound(c {{.Pkg.Context.Value}}, err error, respMsg string) {{.Pkg.ReturnType.Value}} {
        {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusInternalServerError, CodeNotFound, err, respMsg)
      }

      func ErrorSaveSession(c {{.Pkg.Context.Value}}, err error) {{.Pkg.ReturnType.Value}} {
        {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusInternalServerError, CodeSaveSession, err, "failed to save session")
      }

      // ErrorSendRequest 发送 fast http 请求失败
      func ErrorSendRequest(c {{.Pkg.Context.Value}}, err error, respMsg string) {{.Pkg.ReturnType.Value}} {
        {{.Pkg.Return.Value}} errorResponse(c, {{.Pkg.StatusCode.Value}}.StatusInternalServerError, CodeSendRequest, err, respMsg)
      }`
)

// 成功响应
const (
	Success = `package response

      import (
        "{{.Pkg.Context.Import}}"
        "{{.Pkg.StatusCode.Import}}"
      )

      func success(c {{.Pkg.Context.Value}}, data any, total int) {{.Pkg.ReturnType.Value}} {
        {{.Pkg.Return.Value}} response(c, {{.Pkg.StatusCode.Value}}.StatusOK, CodeSuccess, nil, MsgSuccess, data, total)
      }

      func Success(c {{.Pkg.Context.Value}}) {{.Pkg.ReturnType.Value}} {
        {{.Pkg.Return.Value}} response(c, {{.Pkg.StatusCode.Value}}.StatusOK, CodeSuccess, nil, MsgSuccess, struct{}{}, 0)
      }
      func SuccessWithData(c {{.Pkg.Context.Value}}, data any) {{.Pkg.ReturnType.Value}} {
        {{.Pkg.Return.Value}} response(c, {{.Pkg.StatusCode.Value}}.StatusOK, CodeSuccess, nil, MsgSuccess, data, 0)
      }
      func SuccessWithArrayData(c {{.Pkg.Context.Value}}, data any, total int) {{.Pkg.ReturnType.Value}} {
        {{.Pkg.Return.Value}} response(c, {{.Pkg.StatusCode.Value}}.StatusOK, CodeSuccess, nil, MsgSuccess, data, total)
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
        CodeNotFound                  // 未找到
        CodeSaveSession               // session保存错误
        CodeSendRequest               // 发送http请求错误
      )`
)
