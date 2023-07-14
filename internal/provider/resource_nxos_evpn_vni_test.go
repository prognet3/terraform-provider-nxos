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

func TestAccNxosEVPNVNI(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosEVPNVNIPrerequisitesConfig + testAccNxosEVPNVNIConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_evpn_vni.test", "encap", "vxlan-123456"),
					resource.TestCheckResourceAttr("nxos_evpn_vni.test", "route_distinguisher", "rd:unknown:0:0"),
				),
			},
			{
				ResourceName:  "nxos_evpn_vni.test",
				ImportState:   true,
				ImportStateId: "sys/evpn/bdevi-[vxlan-123456]",
			},
		},
	})
}

const testAccNxosEVPNVNIPrerequisitesConfig = `
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
  dn = "sys/evpn"
  class_name = "rtctrlL2Evpn"
  content = {
      adminSt = "enabled"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

`

func testAccNxosEVPNVNIConfig_minimum() string {
	return `
	resource "nxos_evpn_vni" "test" {
		encap = "vxlan-123456"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}

func testAccNxosEVPNVNIConfig_all() string {
	return `
	resource "nxos_evpn_vni" "test" {
		encap = "vxlan-123456"
		route_distinguisher = "rd:unknown:0:0"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}
