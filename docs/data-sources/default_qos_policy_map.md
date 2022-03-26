---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_default_qos_policy_map Data Source - terraform-provider-nxos"
subcategory: ""
description: |-
  This data source can read the default QoS policy map configuration.
---

# nxos_default_qos_policy_map (Data Source)

This data source can read the default QoS policy map configuration.

## Example Usage

```terraform
data "nxos_default_qos_policy_map" "example" {
  name = "PM1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) Policy map name.

### Read-Only

- **id** (String) The distinguished name of the object.
- **match_type** (String) Match type.

