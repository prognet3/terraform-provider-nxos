// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &PhysicalInterfaceResource{}
var _ resource.ResourceWithImportState = &PhysicalInterfaceResource{}

func NewPhysicalInterfaceResource() resource.Resource {
	return &PhysicalInterfaceResource{}
}

type PhysicalInterfaceResource struct {
	data *NxosProviderData
}

func (r *PhysicalInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_physical_interface"
}

func (r *PhysicalInterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage a physical interface.", "l1PhysIf", "System/l1:PhysIf/").AddChildren("physical_interface_vrf").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Must match first field in the output of `show intf brief`. Example: `eth1/1`.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"fec_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("FEC mode.").AddStringEnumDescription("fc-fec", "rs-fec", "fec-off", "auto", "rs-ieee", "rs-cons16", "kp-fec").AddDefaultValueDescription("auto").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("fc-fec", "rs-fec", "fec-off", "auto", "rs-ieee", "rs-cons16", "kp-fec"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("auto"),
				},
			},
			"access_vlan": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Access VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.").AddDefaultValueDescription("vlan-1").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("vlan-1"),
				},
			},
			"admin_state": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port state.").AddStringEnumDescription("up", "down").AddDefaultValueDescription("up").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("up", "down"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("up"),
				},
			},
			"auto_negotiation": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port auto-negotiation.").AddStringEnumDescription("on", "off", "25G").AddDefaultValueDescription("on").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("on", "off", "25G"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("on"),
				},
			},
			"bandwidth": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The bandwidth parameter for a routed interface, port channel, or subinterface.").AddIntegerRangeDescription(0, 3200000000).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 3200000000),
				},
				PlanModifiers: []planmodifier.Int64{
					helpers.IntegerDefaultModifier(0),
				},
			},
			"delay": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The administrative port delay time.").AddIntegerRangeDescription(1, 16777215).AddDefaultValueDescription("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 16777215),
				},
				PlanModifiers: []planmodifier.Int64{
					helpers.IntegerDefaultModifier(1),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface description.").String,
				Optional:            true,
				Computed:            true,
			},
			"duplex": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Duplex.").AddStringEnumDescription("auto", "full", "half").AddDefaultValueDescription("auto").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("auto", "full", "half"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("auto"),
				},
			},
			"layer": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port layer.").AddStringEnumDescription("Layer2", "Layer3").AddDefaultValueDescription("Layer2").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Layer2", "Layer3"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("Layer2"),
				},
			},
			"link_logging": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative link logging.").AddStringEnumDescription("default", "enable", "disable").AddDefaultValueDescription("default").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "enable", "disable"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("default"),
				},
			},
			"link_debounce_down": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port link debounce interval.").AddIntegerRangeDescription(0, 20000).AddDefaultValueDescription("100").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 20000),
				},
				PlanModifiers: []planmodifier.Int64{
					helpers.IntegerDefaultModifier(100),
				},
			},
			"link_debounce_up": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Link Debounce Interval - LinkUp Event.").AddIntegerRangeDescription(0, 20000).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 20000),
				},
				PlanModifiers: []planmodifier.Int64{
					helpers.IntegerDefaultModifier(0),
				},
			},
			"medium": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The administrative port medium type.").AddStringEnumDescription("broadcast", "p2p").AddDefaultValueDescription("broadcast").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("broadcast", "p2p"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("broadcast"),
				},
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port mode.").AddStringEnumDescription("access", "trunk", "fex-fabric", "dot1q-tunnel", "promiscuous", "host", "trunk_secondary", "trunk_promiscuous", "vntag").AddDefaultValueDescription("access").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("access", "trunk", "fex-fabric", "dot1q-tunnel", "promiscuous", "host", "trunk_secondary", "trunk_promiscuous", "vntag"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("access"),
				},
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port MTU.").AddIntegerRangeDescription(576, 9216).AddDefaultValueDescription("1500").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(576, 9216),
				},
				PlanModifiers: []planmodifier.Int64{
					helpers.IntegerDefaultModifier(1500),
				},
			},
			"native_vlan": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Native VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.").AddDefaultValueDescription("vlan-1").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("vlan-1"),
				},
			},
			"speed": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port speed.").AddStringEnumDescription("unknown", "100M", "1G", "10G", "40G", "auto", "auto 100M", "auto 100M 1G", "100G", "25G", "10M", "50G", "200G", "400G", "2.5G", "5G", "auto 2.5G 5G 10G", "auto 100M 1G 2.5G 5G").AddDefaultValueDescription("auto").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "100M", "1G", "10G", "40G", "auto", "auto 100M", "auto 100M 1G", "100G", "25G", "10M", "50G", "200G", "400G", "2.5G", "5G", "auto 2.5G 5G 10G", "auto 100M 1G 2.5G 5G"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("auto"),
				},
			},
			"speed_group": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Speed group.").AddStringEnumDescription("unknown", "1000", "10000", "40000", "auto", "25000").AddDefaultValueDescription("auto").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "1000", "10000", "40000", "auto", "25000"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("auto"),
				},
			},
			"trunk_vlans": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of trunk VLANs.").AddDefaultValueDescription("1-4094").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("1-4094"),
				},
			},
			"uni_directional_ethernet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("UDE (Uni-Directional Ethernet).").AddStringEnumDescription("disable", "send-only", "receive-only").AddDefaultValueDescription("disable").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disable", "send-only", "receive-only"),
				},
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"user_configured_flags": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port User Config Flags.").String,
				Optional:            true,
				Computed:            true,
			},
		},
	}
}

func (r *PhysicalInterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*NxosProviderData)
}

func (r *PhysicalInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, state PhysicalInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	// Post object
	body := plan.toBody()
	_, err := r.data.client.Post(plan.getDn(), body.Str, nxos.OverrideUrl(r.data.devices[plan.Device.ValueString()]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.data.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.data.devices[plan.Device.ValueString()]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(plan)
	state.Dn = types.StringValue(plan.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *PhysicalInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PhysicalInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Dn.ValueString()))

	res, err := r.data.client.GetDn(state.Dn.ValueString(), nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.data.devices[state.Device.ValueString()]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *PhysicalInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PhysicalInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	body := plan.toBody()
	_, err := r.data.client.Post(plan.getDn(), body.Str, nxos.OverrideUrl(r.data.devices[plan.Device.ValueString()]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.data.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.data.devices[plan.Device.ValueString()]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *PhysicalInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PhysicalInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.ValueString()))

	res, err := r.data.client.DeleteDn(state.Dn.ValueString(), nxos.OverrideUrl(r.data.devices[state.Device.ValueString()]))
	if err != nil {
		errCode := res.Get("imdata.0.error.attributes.code").Str
		// Ignore errors of type "Cannot delete object"
		if errCode != "1" && errCode != "107" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Dn.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *PhysicalInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
