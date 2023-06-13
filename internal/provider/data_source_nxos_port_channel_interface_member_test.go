// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosPortChannelInterfaceMember(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosPortChannelInterfaceMemberPrerequisitesConfig + testAccDataSourceNxosPortChannelInterfaceMemberConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface_member.test", "interface_dn", "sys/intf/phys-[eth1/11]"),
				),
			},
		},
	})
}

const testAccDataSourceNxosPortChannelInterfaceMemberPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/intf/aggr-[po1]"
  class_name = "pcAggrIf"
  content = {
      id = "po1"
  }
}

`

const testAccDataSourceNxosPortChannelInterfaceMemberConfig = `

resource "nxos_port_channel_interface_member" "test" {
  interface_id = "po1"
  interface_dn = "sys/intf/phys-[eth1/11]"
  depends_on = [nxos_rest.PreReq0, ]
}

data "nxos_port_channel_interface_member" "test" {
  interface_id = "po1"
  interface_dn = "sys/intf/phys-[eth1/11]"
  depends_on = [nxos_port_channel_interface_member.test]
}
`