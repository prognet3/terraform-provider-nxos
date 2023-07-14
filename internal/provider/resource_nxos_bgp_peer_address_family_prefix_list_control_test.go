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

func TestAccNxosBGPPeerAddressFamilyPrefixListControl(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosBGPPeerAddressFamilyPrefixListControlPrerequisitesConfig + testAccNxosBGPPeerAddressFamilyPrefixListControlConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family_prefix_list_control.test", "asn", "65001"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family_prefix_list_control.test", "vrf", "default"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family_prefix_list_control.test", "address", "192.168.0.1"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family_prefix_list_control.test", "address_family", "ipv4-ucast"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family_prefix_list_control.test", "direction", "in"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family_prefix_list_control.test", "list", "PREFIX_LIST1"),
				),
			},
			{
				ResourceName:  "nxos_bgp_peer_address_family_prefix_list_control.test",
				ImportState:   true,
				ImportStateId: "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]/af-[ipv4-ucast]/pfxctrl-[in]",
			},
		},
	})
}

const testAccNxosBGPPeerAddressFamilyPrefixListControlPrerequisitesConfig = `
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
  dn = "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]"
  class_name = "bgpPeer"
  content = {
      addr = "192.168.0.1"
      asn = "65001"
  }
  depends_on = [nxos_rest.PreReq3, ]
}

resource "nxos_rest" "PreReq5" {
  dn = "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]/af-[ipv4-ucast]"
  class_name = "bgpPeerAf"
  content = {
      type = "ipv4-ucast"
  }
  depends_on = [nxos_rest.PreReq4, ]
}

`

func testAccNxosBGPPeerAddressFamilyPrefixListControlConfig_minimum() string {
	return `
	resource "nxos_bgp_peer_address_family_prefix_list_control" "test" {
		asn = "65001"
		vrf = "default"
		address = "192.168.0.1"
		address_family = "ipv4-ucast"
		direction = "in"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, ]
	}
	`
}

func testAccNxosBGPPeerAddressFamilyPrefixListControlConfig_all() string {
	return `
	resource "nxos_bgp_peer_address_family_prefix_list_control" "test" {
		asn = "65001"
		vrf = "default"
		address = "192.168.0.1"
		address_family = "ipv4-ucast"
		direction = "in"
		list = "PREFIX_LIST1"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, ]
	}
	`
}
