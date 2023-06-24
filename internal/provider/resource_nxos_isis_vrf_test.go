// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNxosISISVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosISISVRFPrerequisitesConfig + testAccNxosISISVRFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "instance_name", "ISIS1"),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "name", "default"),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "authentication_check_l1", "false"),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "authentication_check_l2", "false"),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "authentication_key_l1", ""),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "authentication_key_l2", ""),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "authentication_type_l1", "unknown"),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "authentication_type_l2", "unknown"),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "bandwidth_reference", "400000"),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "banwidth_reference_unit", "mbps"),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "is_type", "l2"),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "metric_type", "wide"),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "mtu", "2000"),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "net", "49.0001.0000.0000.3333.00"),
					resource.TestCheckResourceAttr("nxos_isis_vrf.test", "passive_default", "l12"),
				),
			},
			{
				ResourceName:  "nxos_isis_vrf.test",
				ImportState:   true,
				ImportStateId: "sys/isis/inst-[ISIS1]/dom-[default]",
			},
		},
	})
}

const testAccNxosISISVRFPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/isis"
  class_name = "fmIsis"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/isis"
  class_name = "isisEntity"
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/isis/inst-[ISIS1]"
  class_name = "isisInst"
  content = {
      name = "ISIS1"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

`

func testAccNxosISISVRFConfig_minimum() string {
	return `
	resource "nxos_isis_vrf" "test" {
		instance_name = "ISIS1"
		name = "default"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}

func testAccNxosISISVRFConfig_all() string {
	return `
	resource "nxos_isis_vrf" "test" {
		instance_name = "ISIS1"
		name = "default"
		admin_state = "enabled"
		authentication_check_l1 = false
		authentication_check_l2 = false
		authentication_key_l1 = ""
		authentication_key_l2 = ""
		authentication_type_l1 = "unknown"
		authentication_type_l2 = "unknown"
		bandwidth_reference = 400000
		banwidth_reference_unit = "mbps"
		is_type = "l2"
		metric_type = "wide"
		mtu = 2000
		net = "49.0001.0000.0000.3333.00"
		passive_default = "l12"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}
