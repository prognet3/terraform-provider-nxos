// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapSetQOSGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosDefaultQOSClassMapConfig_all(),
			},
			{
				Config: testAccNxosDefaultQOSClassMapConfig_all() + testAccNxosDefaultQOSPolicyMapConfig_all(),
			},
			{
				Config: testAccNxosDefaultQOSClassMapConfig_all() + testAccNxosDefaultQOSPolicyMapConfig_all() + testAccNxosDefaultQOSPolicyMapMatchClassMapConfig_all(),
			},
			{
				Config: testAccNxosDefaultQOSClassMapConfig_all() + testAccNxosDefaultQOSPolicyMapConfig_all() + testAccNxosDefaultQOSPolicyMapMatchClassMapConfig_all() + testAccNxosDefaultQOSPolicyMapMatchClassMapSetQOSGroupConfig_all(),
			},
			{
				Config: testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapSetQOSGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_set_qos_group.test", "qos_group_id", "1"),
				),
			},
		},
	})
}

const testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapSetQOSGroupConfig = `
data "nxos_default_qos_policy_map_match_class_map_set_qos_group" "test" {
  policy_map_name = "PM1"
  class_map_name = "Voice"
}
`