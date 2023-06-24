// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNxosFeatureISIS(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosFeatureISISConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_feature_isis.test", "admin_state", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_feature_isis.test",
				ImportState:   true,
				ImportStateId: "sys/fm/isis",
			},
		},
	})
}

func testAccNxosFeatureISISConfig_minimum() string {
	return `
	resource "nxos_feature_isis" "test" {
		admin_state = "enabled"
	}
	`
}

func testAccNxosFeatureISISConfig_all() string {
	return `
	resource "nxos_feature_isis" "test" {
		admin_state = "enabled"
	}
	`
}
