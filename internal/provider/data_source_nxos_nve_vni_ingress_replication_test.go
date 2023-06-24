// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNxosNVEVNIIngressReplication(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosNVEVNIIngressReplicationPrerequisitesConfig + testAccDataSourceNxosNVEVNIIngressReplicationConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_nve_vni_ingress_replication.test", "protocol", "bgp"),
				),
			},
		},
	})
}

const testAccDataSourceNxosNVEVNIIngressReplicationPrerequisitesConfig = `
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

resource "nxos_rest" "PreReq4" {
  dn = "sys/eps/epId-[1]/nws/vni-[103100]"
  class_name = "nvoNw"
  content = {
      vni = "103100"
  }
  depends_on = [nxos_rest.PreReq3, ]
}

`

const testAccDataSourceNxosNVEVNIIngressReplicationConfig = `

resource "nxos_nve_vni_ingress_replication" "test" {
  vni = 103100
  protocol = "bgp"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, ]
}

data "nxos_nve_vni_ingress_replication" "test" {
  vni = 103100
  depends_on = [nxos_nve_vni_ingress_replication.test]
}
`
