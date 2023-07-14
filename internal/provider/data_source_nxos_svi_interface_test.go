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

func TestAccDataSourceNxosSVIInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosSVIInterfacePrerequisitesConfig + testAccDataSourceNxosSVIInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "interface_id", "vlan293"),
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "admin_state", "down"),
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "bandwidth", "1000"),
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "delay", "10"),
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "description", "My Description"),
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "medium", "bcast"),
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "mtu", "9216"),
				),
			},
		},
	})
}

const testAccDataSourceNxosSVIInterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/ifvlan"
  class_name = "fmInterfaceVlan"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

`

const testAccDataSourceNxosSVIInterfaceConfig = `

resource "nxos_svi_interface" "test" {
  interface_id = "vlan293"
  admin_state = "down"
  bandwidth = 1000
  delay = 10
  description = "My Description"
  medium = "bcast"
  mtu = 9216
  depends_on = [nxos_rest.PreReq0, ]
}

data "nxos_svi_interface" "test" {
  interface_id = "vlan293"
  depends_on = [nxos_svi_interface.test]
}
`
