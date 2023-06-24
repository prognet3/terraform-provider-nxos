// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNxosNVEVNIContainer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosNVEVNIContainerPrerequisitesConfig + testAccNxosNVEVNIContainerConfig_all(),
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				ResourceName:  "nxos_nve_vni_container.test",
				ImportState:   true,
				ImportStateId: "sys/eps/epId-[1]/nws",
			},
		},
	})
}

const testAccNxosNVEVNIContainerPrerequisitesConfig = `
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
  }
  depends_on = [nxos_rest.PreReq1, ]
}

`

func testAccNxosNVEVNIContainerConfig_minimum() string {
	return `
	resource "nxos_nve_vni_container" "test" {
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}

func testAccNxosNVEVNIContainerConfig_all() string {
	return `
	resource "nxos_nve_vni_container" "test" {
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}
