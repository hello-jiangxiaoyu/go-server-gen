package fiber

const (
	MsgSuccess = ""
)

const (
	CodeSuccess = 200
	CodeAccept  = 202

	CodeNoSuchRoute = 1000 // 系统相关错误码
	CodeRequestPara = 1001
	CodeForbidden   = 1002
	CodeNoRoute     = 1003

	CodeNoSuchHost   = 2000 // 业务相关错误码
	CodeNotLogin     = 2001
	CodeInvalidToken = 2002

	CodeSqlSelect          = 3000 // 查询失败
	CodeSqlSelectNotFound  = 3001 // 不存在该数据
	CodeSqlModify          = 3002 // 修改失败
	CodeSqlModifyDuplicate = 3003 // 数据冲突导致修改失败
	CodeSqlCreate          = 3004 // 创建失败
	CodeSqlCreateDuplicate = 3005 // 数据重复导致创建失败
	CodeSqlDelete          = 3006 // 删除失败
	CodeSqlDeleteForKey    = 3007 // 外键依赖导致删除失败

	CodeServerPanic = 5000 // error相关错误码
	CodeUnknown     = 5001
	CodeNotFound    = 5002
	CodeSaveSession = 5003
	CodeSendRequest = 5004
)
