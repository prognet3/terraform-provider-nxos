// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type BGPPeerTemplateMaxPrefix struct {
	Device      types.String `tfsdk:"device"`
	Dn          types.String `tfsdk:"id"`
	Vrf_name    types.String `tfsdk:"vrf"`
	Name        types.String `tfsdk:"template_name"`
	Type        types.String `tfsdk:"address_family"`
	Action      types.String `tfsdk:"action"`
	MaxPfx      types.Int64  `tfsdk:"maximum_prefix"`
	RestartTime types.Int64  `tfsdk:"restart_time"`
	Thresh      types.Int64  `tfsdk:"threshold"`
}

func (data BGPPeerTemplateMaxPrefix) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/peercont-[%s]/af-[%s]/maxpfxp", data.Vrf_name.Value, data.Name.Value, data.Type.Value)
}

func (data BGPPeerTemplateMaxPrefix) getClassName() string {
	return "bgpMaxPfxP"
}

func (data BGPPeerTemplateMaxPrefix) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("action", data.Action.Value).
		Set("maxPfx", strconv.FormatInt(data.MaxPfx.Value, 10)).
		Set("restartTime", strconv.FormatInt(data.RestartTime.Value, 10)).
		Set("thresh", strconv.FormatInt(data.Thresh.Value, 10))
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *BGPPeerTemplateMaxPrefix) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Action.Value = res.Get("*.attributes.action").String()
	data.MaxPfx.Value = res.Get("*.attributes.maxPfx").Int()
	data.RestartTime.Value = res.Get("*.attributes.restartTime").Int()
	data.Thresh.Value = res.Get("*.attributes.thresh").Int()
}

func (data *BGPPeerTemplateMaxPrefix) fromPlan(plan BGPPeerTemplateMaxPrefix) {
	data.Device = plan.Device
	data.Vrf_name.Value = plan.Vrf_name.Value
	data.Name.Value = plan.Name.Value
	data.Type.Value = plan.Type.Value
}