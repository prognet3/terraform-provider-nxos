resource "nxos_ipv4_prefix_list_rule_entry" "example" {
  rule_name  = "RULE1"
  order      = 10
  action     = "permit"
  criteria   = "exact"
  prefix     = "192.168.1.0/24"
  from_range = 0
  to_range   = 24
}