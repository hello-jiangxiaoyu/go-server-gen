package conf

const DefaultApiTemplate = `
// {{.FuncName}}
// @Tags	{{.Tag}}
// @Summary	{{.Summary}}
{{- range $_, $Para := .ReqParam}}
// @Param	{{$Para.Name}}	{{$Para.From}}	{{$Para.Type}}	{{$Para.Required}}	"{{$Para.Description}}"
{{- end }}
// @Success	200
// @Router	{{.DocPath}} [{{.Method}}]
func {{.FuncName}}(c {{.Context}}) {
	// you controller code hear

}
`

const defaultGroupTemplate = `
`
