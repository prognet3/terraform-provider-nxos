---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_port_channel_interface_member Resource - terraform-provider-nxos"
subcategory: "Interface"
description: |-
  This resource can manage a port-channel interface member.
  API Documentation: pcRsMbrIfs https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Interfaces/pc:RsMbrIfs/
  Parent resources
  nxosportchannel_interface https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/port_channel_interface
  Referenced resources
  nxosphysicalinterface https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/physical_interface
---

# nxos_port_channel_interface_member (Resource)

This resource can manage a port-channel interface member.

- API Documentation: [pcRsMbrIfs](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Interfaces/pc:RsMbrIfs/)

### Parent resources

- [nxos_port_channel_interface](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/port_channel_interface)

### Referenced resources

- [nxos_physical_interface](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/physical_interface)

## Example Usage

```terraform
resource "nxos_port_channel_interface_member" "example" {
  interface_id = "po1"
  interface_dn = "sys/intf/phys-[eth1/11]"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_dn` (String) DN of interface. For example: `sys/intf/phys-[eth1/1]`.
- `interface_id` (String) Must match first field in the output of `show intf brief`. Example: `po1`.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_port_channel_interface_member.example "sys/intf/aggr-[po1]/rsmbrIfs-[sys/intf/phys-[eth1/11]]"
```
