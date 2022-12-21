// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &IPv4AccessListPolicyIngressInterfaceInstaceDataSource{}
	_ datasource.DataSourceWithConfigure = &IPv4AccessListPolicyIngressInterfaceInstaceDataSource{}
)

func NewIPv4AccessListPolicyIngressInterfaceInstaceDataSource() datasource.DataSource {
	return &IPv4AccessListPolicyIngressInterfaceInstaceDataSource{}
}

type IPv4AccessListPolicyIngressInterfaceInstaceDataSource struct {
	data *NxosProviderData
}

func (d *IPv4AccessListPolicyIngressInterfaceInstaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv4_access_list_policy_ingress_interface_instace"
}

func (d *IPv4AccessListPolicyIngressInterfaceInstaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read an IPv4 Access List Policy Ingress Interface Instance.", "aclInst", "Security%20and%20Policing/acl:Inst/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: "Must match first field in the output of `show intf brief`. Example: `eth1/1`.",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Access list name.",
				Computed:            true,
			},
		},
	}
}

func (d *IPv4AccessListPolicyIngressInterfaceInstaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *IPv4AccessListPolicyIngressInterfaceInstaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state IPv4AccessListPolicyIngressInterfaceInstace

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	res, err := d.data.client.GetDn(config.getDn(), nxos.OverrideUrl(d.data.devices[config.Device.ValueString()]))

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(config)
	state.Dn = types.StringValue(config.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
