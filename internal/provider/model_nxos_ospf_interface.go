// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
	"github.com/tidwall/gjson"
)

type OSPFInterface struct {
	Dn                   types.String `tfsdk:"id"`
	Inst                 types.String `tfsdk:"instance_name"`
	Name                 types.String `tfsdk:"vrf_name"`
	Id                   types.String `tfsdk:"interface_id"`
	AdvertiseSecondaries types.Bool   `tfsdk:"advertise_secondaries"`
	Area                 types.String `tfsdk:"area"`
	BfdCtrl              types.String `tfsdk:"bfd"`
	Cost                 types.Int64  `tfsdk:"cost"`
	DeadIntvl            types.Int64  `tfsdk:"dead_interval"`
	HelloIntvl           types.Int64  `tfsdk:"hello_interval"`
	NwT                  types.String `tfsdk:"network_type"`
	PassiveCtrl          types.String `tfsdk:"passive"`
	Prio                 types.Int64  `tfsdk:"priority"`
}

func (data OSPFInterface) getDn() string {
	return fmt.Sprintf("sys/ospf/inst-[%s]/dom-[%s]/if-[%s]", data.Inst.Value, data.Name.Value, data.Id.Value)
}

func (data OSPFInterface) getClassName() string {
	return "ospfIf"
}

func (data OSPFInterface) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("id", data.Id.Value).
		Set("advertiseSecondaries", strconv.FormatBool(data.AdvertiseSecondaries.Value)).
		Set("area", data.Area.Value).
		Set("bfdCtrl", data.BfdCtrl.Value).
		Set("cost", strconv.FormatInt(data.Cost.Value, 10)).
		Set("deadIntvl", strconv.FormatInt(data.DeadIntvl.Value, 10)).
		Set("helloIntvl", strconv.FormatInt(data.HelloIntvl.Value, 10)).
		Set("nwT", data.NwT.Value).
		Set("passiveCtrl", data.PassiveCtrl.Value).
		Set("prio", strconv.FormatInt(data.Prio.Value, 10))
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *OSPFInterface) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Id.Value = res.Get("*.attributes.id").String()
	data.AdvertiseSecondaries.Value = helpers.ParseNxosBoolean(res.Get("*.attributes.advertiseSecondaries").String())
	data.Area.Value = res.Get("*.attributes.area").String()
	data.BfdCtrl.Value = res.Get("*.attributes.bfdCtrl").String()
	data.Cost.Value = res.Get("*.attributes.cost").Int()
	data.DeadIntvl.Value = res.Get("*.attributes.deadIntvl").Int()
	data.HelloIntvl.Value = res.Get("*.attributes.helloIntvl").Int()
	data.NwT.Value = res.Get("*.attributes.nwT").String()
	data.PassiveCtrl.Value = res.Get("*.attributes.passiveCtrl").String()
	data.Prio.Value = res.Get("*.attributes.prio").Int()
}

func (data *OSPFInterface) fromPlan(plan OSPFInterface) {
	data.Inst.Value = plan.Inst.Value
	data.Name.Value = plan.Name.Value
}