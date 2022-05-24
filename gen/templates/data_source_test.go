//go:build ignore
{{if .ExcludeTest}}//go:build testAll{{end}}
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxos{{camelCase .Name}}(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: {{if .TestPrerequisites}}testAccDataSourceNxos{{camelCase .Name}}PrerequisitesConfig+{{end}}testAccDataSourceNxos{{camelCase .Name}}Config,
				Check: resource.ComposeTestCheckFunc(
					{{- $name := .Name }}
					{{- range  .Attributes}}
					{{- if and (ne .ReferenceOnly true) (ne .WriteOnly true)}}
					resource.TestCheckResourceAttr("data.nxos_{{snakeCase $name}}.test", "{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
				),
			},
		},
	})
}

{{- if .TestPrerequisites}}
const testAccDataSourceNxos{{camelCase .Name}}PrerequisitesConfig = `
{{- range $index, $item := .TestPrerequisites}}
resource "nxos_rest" "PreReq{{$index}}" {
  dn = "{{.Dn}}"
  class_name = "{{.ClassName}}"
  {{- if .NoDelete}}
  delete = false
  {{- end}}
  content = {
    {{- range  .Attributes}}
      {{.Name}} = {{if .Reference}}{{.Reference}}{{else}}"{{.Value}}"{{end}}
    {{- end}}
  }
  {{- if .Dependencies}}
  depends_on = [{{range .Dependencies}}nxos_rest.PreReq{{.}}, {{end}}]
  {{- end}}
}
{{ end}}
`
{{- end}}

const testAccDataSourceNxos{{camelCase .Name}}Config = `

resource "nxos_{{snakeCase $name}}" "test" {
{{- range  .Attributes}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- if .TestPrerequisites}}
  depends_on = [{{range $index, $item := .TestPrerequisites}}nxos_rest.PreReq{{$index}}, {{end}}]
{{- end}}
}

data "nxos_{{snakeCase .Name}}" "test" {
{{- range  .Attributes}}
{{- if or (eq .Id true) (eq .ReferenceOnly true)}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end}}
  depends_on = [nxos_{{snakeCase $name}}.test]
}
`
