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

func TestAccNxosPIMAnycastRP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosPIMAnycastRPPrerequisitesConfig + testAccNxosPIMAnycastRPConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_pim_anycast_rp.test", "vrf_name", "default"),
					resource.TestCheckResourceAttr("nxos_pim_anycast_rp.test", "local_interface", "eth1/10"),
					resource.TestCheckResourceAttr("nxos_pim_anycast_rp.test", "source_interface", "eth1/10"),
				),
			},
			{
				ResourceName:  "nxos_pim_anycast_rp.test",
				ImportState:   true,
				ImportStateId: "sys/pim/inst/dom-[default]/acastrpfunc",
			},
		},
	})
}

const testAccNxosPIMAnycastRPPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/pim"
  class_name = "fmPim"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/pim"
  class_name = "pimEntity"
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/pim/inst"
  class_name = "pimInst"
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/pim/inst/dom-[default]"
  class_name = "pimDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

`

func testAccNxosPIMAnycastRPConfig_minimum() string {
	return `
	resource "nxos_pim_anycast_rp" "test" {
		vrf_name = "default"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}

func testAccNxosPIMAnycastRPConfig_all() string {
	return `
	resource "nxos_pim_anycast_rp" "test" {
		vrf_name = "default"
		local_interface = "eth1/10"
		source_interface = "eth1/10"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}
