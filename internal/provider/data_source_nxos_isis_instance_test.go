// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNxosISISInstance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosISISInstancePrerequisitesConfig + testAccDataSourceNxosISISInstanceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_isis_instance.test", "name", "ISIS1"),
					resource.TestCheckResourceAttr("data.nxos_isis_instance.test", "admin_state", "enabled"),
				),
			},
		},
	})
}

const testAccDataSourceNxosISISInstancePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/isis"
  class_name = "fmIsis"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/isis"
  class_name = "isisEntity"
  depends_on = [nxos_rest.PreReq0, ]
}

`

const testAccDataSourceNxosISISInstanceConfig = `

resource "nxos_isis_instance" "test" {
  name = "ISIS1"
  admin_state = "enabled"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

data "nxos_isis_instance" "test" {
  name = "ISIS1"
  depends_on = [nxos_isis_instance.test]
}
`
