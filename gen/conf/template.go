package conf

const defaultApiTemplate = `
// {{.FuncName}}
// @Tags	{{.Tag}}
// @Summary	{{.Summary}}
{{- range $_, $Para := .ReqPara}}
// @Param	{{$Para.Name}}	{{$Para.From}}	{{$Para.Type}}	{{$Para.Required}}	"{{$Para.Description}}"
{{- end }}
// @Success	200
// @Router	{{.Path}} [{{.Method}}]
func {{.FuncName}}(c {{.Context}}) {
	// you controller code hear
}
`

const defaultGroupTemplate = `
`
