// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNxosHMMInstance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosHMMInstancePrerequisitesConfig + testAccDataSourceNxosHMMInstanceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_hmm_instance.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_hmm_instance.test", "anycast_mac", "20:20:00:00:10:10"),
				),
			},
		},
	})
}

const testAccDataSourceNxosHMMInstancePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/hmm"
  class_name = "fmHmm"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/fm/evpn"
  class_name = "fmEvpn"
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/hmm"
  class_name = "hmmEntity"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

`

const testAccDataSourceNxosHMMInstanceConfig = `

resource "nxos_hmm_instance" "test" {
  admin_state = "enabled"
  anycast_mac = "20:20:00:00:10:10"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
}

data "nxos_hmm_instance" "test" {
  depends_on = [nxos_hmm_instance.test]
}
`
