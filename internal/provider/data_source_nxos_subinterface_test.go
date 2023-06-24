// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNxosSubinterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosSubinterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_subinterface.test", "interface_id", "eth1/10.124"),
					resource.TestCheckResourceAttr("data.nxos_subinterface.test", "admin_state", "down"),
					resource.TestCheckResourceAttr("data.nxos_subinterface.test", "bandwidth", "1000"),
					resource.TestCheckResourceAttr("data.nxos_subinterface.test", "delay", "10"),
					resource.TestCheckResourceAttr("data.nxos_subinterface.test", "description", "My Description"),
					resource.TestCheckResourceAttr("data.nxos_subinterface.test", "encap", "vlan-124"),
					resource.TestCheckResourceAttr("data.nxos_subinterface.test", "link_logging", "enable"),
					resource.TestCheckResourceAttr("data.nxos_subinterface.test", "medium", "broadcast"),
					resource.TestCheckResourceAttr("data.nxos_subinterface.test", "mtu", "1500"),
				),
			},
		},
	})
}

const testAccDataSourceNxosSubinterfaceConfig = `

resource "nxos_subinterface" "test" {
  interface_id = "eth1/10.124"
  admin_state = "down"
  bandwidth = 1000
  delay = 10
  description = "My Description"
  encap = "vlan-124"
  link_logging = "enable"
  medium = "broadcast"
  mtu = 1500
}

data "nxos_subinterface" "test" {
  interface_id = "eth1/10.124"
  depends_on = [nxos_subinterface.test]
}
`
