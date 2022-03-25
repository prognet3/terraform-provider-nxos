// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosOSPFVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosOSPFConfig_all(),
			},
			{
				Config: testAccNxosOSPFConfig_all() + testAccNxosOSPFInstanceConfig_all(),
			},
			{
				Config: testAccNxosOSPFConfig_all() + testAccNxosOSPFInstanceConfig_all() + testAccNxosOSPFVRFConfig_all(),
			},
			{
				Config: testAccDataSourceNxosOSPFVRFConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_ospf_vrf.test", "name", "default"),
					resource.TestCheckResourceAttr("data.nxos_ospf_vrf.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_ospf_vrf.test", "bandwidth_reference", "400000"),
					resource.TestCheckResourceAttr("data.nxos_ospf_vrf.test", "banwidth_reference_unit", "mbps"),
					resource.TestCheckResourceAttr("data.nxos_ospf_vrf.test", "distance", "110"),
					resource.TestCheckResourceAttr("data.nxos_ospf_vrf.test", "router_id", "34.56.78.90"),
				),
			},
		},
	})
}

const testAccDataSourceNxosOSPFVRFConfig = `
data "nxos_ospf_vrf" "test" {
  instance_name = "OSPF1"
  name = "default"
}
`