package server

var apiMap = map[string]string{
	"get":         `GET("{{.Prefix}}/{{lowercaseFirst .ServiceName}}", Get{{uppercaseFirst .ServiceName}}List)  // {{uppercaseFirst .ServiceName}}Req`,
	"list":        `GET("{{.Prefix}}/{{lowercaseFirst .ServiceName}}/:{{lowercaseFirst .ServiceName}}ID", Get{{uppercaseFirst .ServiceName}})  // {{uppercaseFirst .ServiceName}}Req`,
	"create":      `POST("{{.Prefix}}/{{lowercaseFirst .ServiceName}}", Create{{uppercaseFirst .ServiceName}})  // {{uppercaseFirst .ServiceName}}Req`,
	"update":      `PUT("{{.Prefix}}/{{lowercaseFirst .ServiceName}}/:{{lowercaseFirst .ServiceName}}ID", Update{{uppercaseFirst .ServiceName}})  // {{uppercaseFirst .ServiceName}}Req`,
	"delete":      `DELETE("{{.Prefix}}/{{lowercaseFirst .ServiceName}}/:{{lowercaseFirst .ServiceName}}ID", Delete{{uppercaseFirst .ServiceName}})  // {{uppercaseFirst .ServiceName}}Req`,
	"batchDelete": `DELETE("{{.Prefix}}/{{lowercaseFirst .ServiceName}}", BatchDelete{{uppercaseFirst .ServiceName}})  // {{uppercaseFirst .ServiceName}}Req`,
}
