// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNxosVRFRouteTarget(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosVRFRouteTargetPrerequisitesConfig + testAccDataSourceNxosVRFRouteTargetConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_vrf_route_target.test", "route_target", "route-target:as2-nn2:2:2"),
				),
			},
		},
	})
}

const testAccDataSourceNxosVRFRouteTargetPrerequisitesConfig = `
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

resource "nxos_rest" "PreReq3" {
  dn = "sys/inst-[VRF1]/dom-[VRF1]/af-[ipv4-ucast]"
  class_name = "rtctrlDomAf"
  content = {
      type = "ipv4-ucast"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

resource "nxos_rest" "PreReq4" {
  dn = "sys/inst-[VRF1]/dom-[VRF1]/af-[ipv4-ucast]/ctrl-[ipv4-ucast]"
  class_name = "rtctrlAfCtrl"
  content = {
      type = "ipv4-ucast"
  }
  depends_on = [nxos_rest.PreReq3, ]
}

resource "nxos_rest" "PreReq5" {
  dn = "sys/inst-[VRF1]/dom-[VRF1]/af-[ipv4-ucast]/ctrl-[ipv4-ucast]/rttp-[import]"
  class_name = "rtctrlRttP"
  content = {
      type = "import"
  }
  depends_on = [nxos_rest.PreReq4, ]
}

`

const testAccDataSourceNxosVRFRouteTargetConfig = `

resource "nxos_vrf_route_target" "test" {
  vrf = "VRF1"
  address_family = "ipv4-ucast"
  route_target_address_family = "ipv4-ucast"
  direction = "import"
  route_target = "route-target:as2-nn2:2:2"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, ]
}

data "nxos_vrf_route_target" "test" {
  vrf = "VRF1"
  address_family = "ipv4-ucast"
  route_target_address_family = "ipv4-ucast"
  direction = "import"
  route_target = "route-target:as2-nn2:2:2"
  depends_on = [nxos_vrf_route_target.test]
}
`
