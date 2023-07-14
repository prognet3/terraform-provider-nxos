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

func TestAccNxosOSPFInstance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosOSPFInstancePrerequisitesConfig + testAccNxosOSPFInstanceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_ospf_instance.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("nxos_ospf_instance.test", "name", "OSPF1"),
				),
			},
			{
				ResourceName:  "nxos_ospf_instance.test",
				ImportState:   true,
				ImportStateId: "sys/ospf/inst-[OSPF1]",
			},
		},
	})
}

const testAccNxosOSPFInstancePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/ospf"
  class_name = "fmOspf"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/ospf"
  class_name = "ospfEntity"
  depends_on = [nxos_rest.PreReq0, ]
}

`

func testAccNxosOSPFInstanceConfig_minimum() string {
	return `
	resource "nxos_ospf_instance" "test" {
		name = "OSPF1"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}

func testAccNxosOSPFInstanceConfig_all() string {
	return `
	resource "nxos_ospf_instance" "test" {
		admin_state = "enabled"
		name = "OSPF1"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}
