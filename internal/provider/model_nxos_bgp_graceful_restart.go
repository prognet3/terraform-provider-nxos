// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type BGPGracefulRestart struct {
	Device       types.String `tfsdk:"device"`
	Dn           types.String `tfsdk:"id"`
	Asn          types.String `tfsdk:"asn"`
	Name         types.String `tfsdk:"vrf"`
	RestartIntvl types.Int64  `tfsdk:"restart_interval"`
	StaleIntvl   types.Int64  `tfsdk:"stale_interval"`
}

func (data BGPGracefulRestart) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/gr", data.Name.ValueString())
}

func (data BGPGracefulRestart) getClassName() string {
	return "bgpGr"
}

func (data BGPGracefulRestart) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("restartIntvl", strconv.FormatInt(data.RestartIntvl.ValueInt64(), 10)).
		Set("staleIntvl", strconv.FormatInt(data.StaleIntvl.ValueInt64(), 10))
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *BGPGracefulRestart) fromBody(res gjson.Result) {
	data.RestartIntvl = types.Int64Value(res.Get("*.attributes.restartIntvl").Int())
	data.StaleIntvl = types.Int64Value(res.Get("*.attributes.staleIntvl").Int())
}

func (data *BGPGracefulRestart) fromPlan(plan BGPGracefulRestart) {
	data.Device = plan.Device
	data.Dn = plan.Dn
	data.Asn = plan.Asn
	data.Name = plan.Name
}
