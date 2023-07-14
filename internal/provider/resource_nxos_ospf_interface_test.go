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

func TestAccNxosOSPFInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosOSPFInterfacePrerequisitesConfig + testAccNxosOSPFInterfaceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "instance_name", "OSPF1"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "vrf_name", "VRF1"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "interface_id", "eth1/10"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "advertise_secondaries", "false"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "area", "0.0.0.10"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "bfd", "disabled"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "cost", "1000"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "dead_interval", "60"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "hello_interval", "15"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "network_type", "p2p"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "passive", "enabled"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "priority", "10"),
				),
			},
			{
				ResourceName:  "nxos_ospf_interface.test",
				ImportState:   true,
				ImportStateId: "sys/ospf/inst-[OSPF1]/dom-[VRF1]/if-[eth1/10]",
			},
		},
	})
}

const testAccNxosOSPFInterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/bfd"
  class_name = "fmBfd"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/fm/ospf"
  class_name = "fmOspf"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/ospf"
  class_name = "ospfEntity"
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/ospf/inst-[OSPF1]"
  class_name = "ospfInst"
  content = {
      name = "OSPF1"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

resource "nxos_rest" "PreReq4" {
  dn = "sys/ospf/inst-[OSPF1]/dom-[VRF1]"
  class_name = "ospfDom"
  content = {
      name = "VRF1"
  }
  depends_on = [nxos_rest.PreReq3, ]
}

resource "nxos_rest" "PreReq5" {
  dn = "sys/intf/phys-[eth1/10]"
  class_name = "l1PhysIf"
  content = {
      layer = "Layer3"
  }
  depends_on = [nxos_rest.PreReq4, ]
}

resource "nxos_rest" "PreReq6" {
  dn = "sys/intf/phys-[eth1/10]/rtvrfMbr"
  class_name = "nwRtVrfMbr"
  content = {
      tDn = "sys/inst-VRF1"
  }
  depends_on = [nxos_rest.PreReq5, ]
}

`

func testAccNxosOSPFInterfaceConfig_minimum() string {
	return `
	resource "nxos_ospf_interface" "test" {
		instance_name = "OSPF1"
		vrf_name = "VRF1"
		interface_id = "eth1/10"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, nxos_rest.PreReq6, ]
	}
	`
}

func testAccNxosOSPFInterfaceConfig_all() string {
	return `
	resource "nxos_ospf_interface" "test" {
		instance_name = "OSPF1"
		vrf_name = "VRF1"
		interface_id = "eth1/10"
		advertise_secondaries = false
		area = "0.0.0.10"
		bfd = "disabled"
		cost = 1000
		dead_interval = 60
		hello_interval = 15
		network_type = "p2p"
		passive = "enabled"
		priority = 10
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, nxos_rest.PreReq6, ]
	}
	`
}
