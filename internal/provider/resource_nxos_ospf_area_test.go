// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNxosOSPFArea(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosOSPFAreaPrerequisitesConfig + testAccNxosOSPFAreaConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_ospf_area.test", "instance_name", "OSPF1"),
					resource.TestCheckResourceAttr("nxos_ospf_area.test", "vrf_name", "VRF1"),
					resource.TestCheckResourceAttr("nxos_ospf_area.test", "area_id", "0.0.0.10"),
					resource.TestCheckResourceAttr("nxos_ospf_area.test", "authentication_type", "none"),
					resource.TestCheckResourceAttr("nxos_ospf_area.test", "cost", "10"),
					resource.TestCheckResourceAttr("nxos_ospf_area.test", "type", "stub"),
				),
			},
			{
				ResourceName:  "nxos_ospf_area.test",
				ImportState:   true,
				ImportStateId: "sys/ospf/inst-[OSPF1]/dom-[VRF1]/area-[0.0.0.10]",
			},
		},
	})
}

const testAccNxosOSPFAreaPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/ospf"
  class_name = "fmOspf"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/ospf"
  class_name = "ospfEntity"
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/ospf/inst-[OSPF1]"
  class_name = "ospfInst"
  content = {
      name = "OSPF1"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/ospf/inst-[OSPF1]/dom-[VRF1]"
  class_name = "ospfDom"
  content = {
      name = "VRF1"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

`

func testAccNxosOSPFAreaConfig_minimum() string {
	return `
	resource "nxos_ospf_area" "test" {
		instance_name = "OSPF1"
		vrf_name = "VRF1"
		area_id = "0.0.0.10"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}

func testAccNxosOSPFAreaConfig_all() string {
	return `
	resource "nxos_ospf_area" "test" {
		instance_name = "OSPF1"
		vrf_name = "VRF1"
		area_id = "0.0.0.10"
		authentication_type = "none"
		cost = 10
		type = "stub"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}
