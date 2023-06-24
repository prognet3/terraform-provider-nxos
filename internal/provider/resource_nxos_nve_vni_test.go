// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNxosNVEVNI(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosNVEVNIPrerequisitesConfig + testAccNxosNVEVNIConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_nve_vni.test", "vni", "103100"),
					resource.TestCheckResourceAttr("nxos_nve_vni.test", "associate_vrf", "false"),
					resource.TestCheckResourceAttr("nxos_nve_vni.test", "multicast_group", "239.1.1.1"),
					resource.TestCheckResourceAttr("nxos_nve_vni.test", "multisite_ingress_replication", "disable"),
					resource.TestCheckResourceAttr("nxos_nve_vni.test", "suppress_arp", "off"),
				),
			},
			{
				ResourceName:  "nxos_nve_vni.test",
				ImportState:   true,
				ImportStateId: "sys/eps/epId-[1]/nws/vni-[103100]",
			},
		},
	})
}

const testAccNxosNVEVNIPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/nvo"
  class_name = "fmNvo"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/fm/evpn"
  class_name = "fmEvpn"
  delete = false
  content = {
      adminSt = "enabled"
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/eps/epId-[1]"
  class_name = "nvoEp"
  content = {
      adminSt = "enabled"
      hostReach = "bgp"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/eps/epId-[1]/nws"
  class_name = "nvoNws"
  depends_on = [nxos_rest.PreReq2, ]
}

`

func testAccNxosNVEVNIConfig_minimum() string {
	return `
	resource "nxos_nve_vni" "test" {
		vni = 103100
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}

func testAccNxosNVEVNIConfig_all() string {
	return `
	resource "nxos_nve_vni" "test" {
		vni = 103100
		associate_vrf = false
		multicast_group = "239.1.1.1"
		multisite_ingress_replication = "disable"
		suppress_arp = "off"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}
