// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNxosSubinterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosSubinterfaceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_subinterface.test", "interface_id", "eth1/10.124"),
					resource.TestCheckResourceAttr("nxos_subinterface.test", "admin_state", "down"),
					resource.TestCheckResourceAttr("nxos_subinterface.test", "bandwidth", "1000"),
					resource.TestCheckResourceAttr("nxos_subinterface.test", "delay", "10"),
					resource.TestCheckResourceAttr("nxos_subinterface.test", "description", "My Description"),
					resource.TestCheckResourceAttr("nxos_subinterface.test", "encap", "vlan-124"),
					resource.TestCheckResourceAttr("nxos_subinterface.test", "link_logging", "enable"),
					resource.TestCheckResourceAttr("nxos_subinterface.test", "medium", "broadcast"),
					resource.TestCheckResourceAttr("nxos_subinterface.test", "mtu", "1500"),
				),
			},
			{
				ResourceName:  "nxos_subinterface.test",
				ImportState:   true,
				ImportStateId: "sys/intf/encrtd-[eth1/10.124]",
			},
		},
	})
}

func testAccNxosSubinterfaceConfig_minimum() string {
	return `
	resource "nxos_subinterface" "test" {
		interface_id = "eth1/10.124"
	}
	`
}

func testAccNxosSubinterfaceConfig_all() string {
	return `
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
	`
}
