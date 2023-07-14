// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNxosBGPPeerTemplateAddressFamily(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosBGPPeerTemplateAddressFamilyPrerequisitesConfig + testAccDataSourceNxosBGPPeerTemplateAddressFamilyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_bgp_peer_template_address_family.test", "address_family", "ipv4-ucast"),
					resource.TestCheckResourceAttr("data.nxos_bgp_peer_template_address_family.test", "control", "nh-self,rr-client"),
					resource.TestCheckResourceAttr("data.nxos_bgp_peer_template_address_family.test", "send_community_extended", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_bgp_peer_template_address_family.test", "send_community_standard", "enabled"),
				),
			},
		},
	})
}

const testAccDataSourceNxosBGPPeerTemplateAddressFamilyPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/bgp"
  class_name = "fmBgp"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/bgp"
  class_name = "bgpEntity"
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/bgp/inst"
  class_name = "bgpInst"
  content = {
      adminSt = "enabled"
      asn = "65001"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/bgp/inst/dom-[default]"
  class_name = "bgpDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

resource "nxos_rest" "PreReq4" {
  dn = "sys/bgp/inst/dom-[default]/peercont-[SPINE-PEERS]"
  class_name = "bgpPeerCont"
  content = {
      name = "SPINE-PEERS"
  }
  depends_on = [nxos_rest.PreReq3, ]
}

`

const testAccDataSourceNxosBGPPeerTemplateAddressFamilyConfig = `

resource "nxos_bgp_peer_template_address_family" "test" {
  asn = "65001"
  template_name = "SPINE-PEERS"
  address_family = "ipv4-ucast"
  control = "nh-self,rr-client"
  send_community_extended = "enabled"
  send_community_standard = "enabled"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, ]
}

data "nxos_bgp_peer_template_address_family" "test" {
  asn = "65001"
  template_name = "SPINE-PEERS"
  address_family = "ipv4-ucast"
  depends_on = [nxos_bgp_peer_template_address_family.test]
}
`
