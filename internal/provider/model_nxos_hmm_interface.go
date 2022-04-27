// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type HMMInterface struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	Id      types.String `tfsdk:"interface_id"`
	AdminSt types.String `tfsdk:"admin_state"`
	Mode    types.String `tfsdk:"mode"`
}

func (data HMMInterface) getDn() string {
	return fmt.Sprintf("sys/hmm/fwdinst/if-[%s]", data.Id.Value)
}

func (data HMMInterface) getClassName() string {
	return "hmmFwdIf"
}

func (data HMMInterface) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("id", data.Id.Value).
		Set("adminSt", data.AdminSt.Value).
		Set("mode", data.Mode.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *HMMInterface) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Id.Value = res.Get("*.attributes.id").String()
	data.AdminSt.Value = res.Get("*.attributes.adminSt").String()
	data.Mode.Value = res.Get("*.attributes.mode").String()
}

func (data *HMMInterface) fromPlan(plan HMMInterface) {
	data.Device = plan.Device
}