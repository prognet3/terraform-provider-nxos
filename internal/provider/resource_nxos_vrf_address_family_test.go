// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNxosVRFAddressFamily(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosVRFAddressFamilyPrerequisitesConfig + testAccNxosVRFAddressFamilyConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_vrf_address_family.test", "vrf", "VRF1"),
					resource.TestCheckResourceAttr("nxos_vrf_address_family.test", "address_family", "ipv4-ucast"),
				),
			},
			{
				ResourceName:  "nxos_vrf_address_family.test",
				ImportState:   true,
				ImportStateId: "sys/inst-[VRF1]/dom-[VRF1]/af-[ipv4-ucast]",
			},
		},
	})
}

const testAccNxosVRFAddressFamilyPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/inst-[VRF1]"
  class_name = "l3Inst"
  content = {
      name = "VRF1"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/ipv4/inst/dom-[VRF1]"
  class_name = "ipv4Dom"
  content = {
      name = "VRF1"
  }
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/inst-[VRF1]/dom-[VRF1]"
  class_name = "rtctrlDom"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

`

func testAccNxosVRFAddressFamilyConfig_minimum() string {
	return `
	resource "nxos_vrf_address_family" "test" {
		vrf = "VRF1"
		address_family = "ipv4-ucast"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}

func testAccNxosVRFAddressFamilyConfig_all() string {
	return `
	resource "nxos_vrf_address_family" "test" {
		vrf = "VRF1"
		address_family = "ipv4-ucast"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}
