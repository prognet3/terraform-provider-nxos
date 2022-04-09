// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type BGPAddressFamily struct {
	Device           types.String `tfsdk:"device"`
	Dn               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"vrf"`
	Type             types.String `tfsdk:"address_family"`
	CritNhTimeout    types.Int64  `tfsdk:"critical_nexthop_timeout"`
	NonCritNhTimeout types.Int64  `tfsdk:"non_critical_nexthop_timeout"`
}

func (data BGPAddressFamily) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/af-[%s]", data.Name.Value, data.Type.Value)
}

func (data BGPAddressFamily) getClassName() string {
	return "bgpGr"
}

func (data BGPAddressFamily) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("type", data.Type.Value).
		Set("critNhTimeout", strconv.FormatInt(data.CritNhTimeout.Value, 10)).
		Set("nonCritNhTimeout", strconv.FormatInt(data.NonCritNhTimeout.Value, 10))
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *BGPAddressFamily) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Type.Value = res.Get("*.attributes.type").String()
	data.CritNhTimeout.Value = res.Get("*.attributes.critNhTimeout").Int()
	data.NonCritNhTimeout.Value = res.Get("*.attributes.nonCritNhTimeout").Int()
}

func (data *BGPAddressFamily) fromPlan(plan BGPAddressFamily) {
	data.Device = plan.Device
	data.Name.Value = plan.Name.Value
}