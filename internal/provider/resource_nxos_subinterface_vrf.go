// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &SubinterfaceVRFResource{}
var _ resource.ResourceWithImportState = &SubinterfaceVRFResource{}

func NewSubinterfaceVRFResource() resource.Resource {
	return &SubinterfaceVRFResource{}
}

type SubinterfaceVRFResource struct {
	data *NxosProviderData
}

func (r *SubinterfaceVRFResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subinterface_vrf"
}

func (r *SubinterfaceVRFResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage a subinterface VRF association.", "nwRtVrfMbr", "Routing%20and%20Forwarding/nw:RtVrfMbr/").AddParents("subinterface").AddReferences("vrf").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"vrf_dn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DN of VRF. For example: `sys/inst-VRF1`.").String,
				Required:            true,
			},
		},
	}
}

func (r *SubinterfaceVRFResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*NxosProviderData)
}

func (r *SubinterfaceVRFResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, state SubinterfaceVRF

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

func (r *SubinterfaceVRFResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SubinterfaceVRF

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

func (r *SubinterfaceVRFResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SubinterfaceVRF

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

func (r *SubinterfaceVRFResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SubinterfaceVRF

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

func (r *SubinterfaceVRFResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
