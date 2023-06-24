// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNxosPIMInstance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosPIMInstancePrerequisitesConfig + testAccNxosPIMInstanceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_pim_instance.test", "admin_state", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_pim_instance.test",
				ImportState:   true,
				ImportStateId: "sys/pim/inst",
			},
		},
	})
}

const testAccNxosPIMInstancePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/pim"
  class_name = "fmPim"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/pim"
  class_name = "pimEntity"
  depends_on = [nxos_rest.PreReq0, ]
}

`

func testAccNxosPIMInstanceConfig_minimum() string {
	return `
	resource "nxos_pim_instance" "test" {
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}

func testAccNxosPIMInstanceConfig_all() string {
	return `
	resource "nxos_pim_instance" "test" {
		admin_state = "enabled"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}
