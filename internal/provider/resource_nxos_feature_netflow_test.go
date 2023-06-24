//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNxosFeatureNetflow(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosFeatureNetflowConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_feature_netflow.test", "admin_state", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_feature_netflow.test",
				ImportState:   true,
				ImportStateId: "sys/fm/netflow",
			},
		},
	})
}

func testAccNxosFeatureNetflowConfig_minimum() string {
	return `
	resource "nxos_feature_netflow" "test" {
		admin_state = "enabled"
	}
	`
}

func testAccNxosFeatureNetflowConfig_all() string {
	return `
	resource "nxos_feature_netflow" "test" {
		admin_state = "enabled"
	}
	`
}
