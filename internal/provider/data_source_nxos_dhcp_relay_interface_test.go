// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNxosDHCPRelayInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosDHCPRelayInterfacePrerequisitesConfig + testAccDataSourceNxosDHCPRelayInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_dhcp_relay_interface.test", "interface_id", "eth1/10"),
				),
			},
		},
	})
}

const testAccDataSourceNxosDHCPRelayInterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/dhcp"
  class_name = "fmDhcp"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/intf/phys-[eth1/10]"
  class_name = "l1PhysIf"
  content = {
      layer = "Layer3"
  }
}

`

const testAccDataSourceNxosDHCPRelayInterfaceConfig = `

resource "nxos_dhcp_relay_interface" "test" {
  interface_id = "eth1/10"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

data "nxos_dhcp_relay_interface" "test" {
  interface_id = "eth1/10"
  depends_on = [nxos_dhcp_relay_interface.test]
}
`
