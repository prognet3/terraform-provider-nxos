// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNxosPortChannelInterfaceMember(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosPortChannelInterfaceMemberPrerequisitesConfig + testAccNxosPortChannelInterfaceMemberConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_port_channel_interface_member.test", "interface_id", "po1"),
					resource.TestCheckResourceAttr("nxos_port_channel_interface_member.test", "interface_dn", "sys/intf/phys-[eth1/11]"),
				),
			},
			{
				ResourceName:  "nxos_port_channel_interface_member.test",
				ImportState:   true,
				ImportStateId: "sys/intf/aggr-[po1]/rsmbrIfs-[sys/intf/phys-[eth1/11]]",
			},
		},
	})
}

const testAccNxosPortChannelInterfaceMemberPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/intf/aggr-[po1]"
  class_name = "pcAggrIf"
  content = {
      id = "po1"
  }
}

`

func testAccNxosPortChannelInterfaceMemberConfig_minimum() string {
	return `
	resource "nxos_port_channel_interface_member" "test" {
		interface_id = "po1"
		interface_dn = "sys/intf/phys-[eth1/11]"
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}

func testAccNxosPortChannelInterfaceMemberConfig_all() string {
	return `
	resource "nxos_port_channel_interface_member" "test" {
		interface_id = "po1"
		interface_dn = "sys/intf/phys-[eth1/11]"
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}
