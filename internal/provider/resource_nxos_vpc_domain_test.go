// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNxosVPCDomain(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosVPCDomainPrerequisitesConfig + testAccNxosVPCDomainConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "domain_id", "100"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "auto_recovery", "enabled"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "auto_recovery_interval", "480"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "delay_restore_orphan_port", "10"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "delay_restore_svi", "20"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "delay_restore_vpc", "60"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "dscp", "0"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "fast_convergence", "enabled"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "graceful_consistency_check", "disabled"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "l3_peer_router", "enabled"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "l3_peer_router_syslog", "enabled"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "l3_peer_router_syslog_interval", "3000"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "peer_gateway", "enabled"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "peer_ip", "1.2.3.4"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "peer_switch", "enabled"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "role_priority", "100"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "sys_mac", "01:01:01:01:01:01"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "system_priority", "100"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "track", "10"),
					resource.TestCheckResourceAttr("nxos_vpc_domain.test", "virtual_ip", "1.1.1.1"),
				),
			},
			{
				ResourceName:  "nxos_vpc_domain.test",
				ImportState:   true,
				ImportStateId: "sys/vpc/inst/dom",
			},
		},
	})
}

const testAccNxosVPCDomainPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/vpc"
  class_name = "fmVpc"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/vpc/inst"
  class_name = "vpcInst"
  depends_on = [nxos_rest.PreReq0, ]
}

`

func testAccNxosVPCDomainConfig_minimum() string {
	return `
	resource "nxos_vpc_domain" "test" {
		domain_id = 100
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}

func testAccNxosVPCDomainConfig_all() string {
	return `
	resource "nxos_vpc_domain" "test" {
		admin_state = "enabled"
		domain_id = 100
		auto_recovery = "enabled"
		auto_recovery_interval = 480
		delay_restore_orphan_port = 10
		delay_restore_svi = 20
		delay_restore_vpc = 60
		dscp = 0
		fast_convergence = "enabled"
		graceful_consistency_check = "disabled"
		l3_peer_router = "enabled"
		l3_peer_router_syslog = "enabled"
		l3_peer_router_syslog_interval = 3000
		peer_gateway = "enabled"
		peer_ip = "1.2.3.4"
		peer_switch = "enabled"
		role_priority = 100
		sys_mac = "01:01:01:01:01:01"
		system_priority = 100
		track = 10
		virtual_ip = "1.1.1.1"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}
