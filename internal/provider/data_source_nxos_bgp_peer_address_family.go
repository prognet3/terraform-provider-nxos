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
	_ datasource.DataSource              = &BGPPeerAddressFamilyDataSource{}
	_ datasource.DataSourceWithConfigure = &BGPPeerAddressFamilyDataSource{}
)

func NewBGPPeerAddressFamilyDataSource() datasource.DataSource {
	return &BGPPeerAddressFamilyDataSource{}
}

type BGPPeerAddressFamilyDataSource struct {
	data *NxosProviderData
}

func (d *BGPPeerAddressFamilyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bgp_peer_address_family"
}

func (d *BGPPeerAddressFamilyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the BGP peer address family configuration.", "bgpPeerAf", "Routing%20and%20Forwarding/bgp:PeerAf/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"asn": schema.StringAttribute{
				MarkdownDescription: "Autonomous system number.",
				Required:            true,
			},
			"vrf": schema.StringAttribute{
				MarkdownDescription: "VRF name.",
				Required:            true,
			},
			"address": schema.StringAttribute{
				MarkdownDescription: "Peer address.",
				Required:            true,
			},
			"address_family": schema.StringAttribute{
				MarkdownDescription: "Address Family.",
				Required:            true,
			},
			"control": schema.StringAttribute{
				MarkdownDescription: "Peer address-family control. Choices: `rr-client`, `nh-self`, `dis-peer-as-check`, `allow-self-as`, `default-originate`, `advertisement-interval`, `suppress-inactive`, `nh-self-all`. Can be an empty string. Allowed formats:\n  - Single value. Example: `nh-self`\n  - Multiple values (comma-separated). Example: `dis-peer-as-check,nh-self,rr-client,suppress-inactive`. In this case values must be in alphabetical order.",
				Computed:            true,
			},
			"send_community_extended": schema.StringAttribute{
				MarkdownDescription: "Send-community extended.",
				Computed:            true,
			},
			"send_community_standard": schema.StringAttribute{
				MarkdownDescription: "Send-community standard.",
				Computed:            true,
			},
		},
	}
}

func (d *BGPPeerAddressFamilyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *BGPPeerAddressFamilyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state BGPPeerAddressFamily

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
