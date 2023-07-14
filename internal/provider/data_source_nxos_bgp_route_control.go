// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &BGPRouteControlDataSource{}
	_ datasource.DataSourceWithConfigure = &BGPRouteControlDataSource{}
)

func NewBGPRouteControlDataSource() datasource.DataSource {
	return &BGPRouteControlDataSource{}
}

type BGPRouteControlDataSource struct {
	clients map[string]*nxos.Client
}

func (d *BGPRouteControlDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bgp_route_control"
}

func (d *BGPRouteControlDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the BGP Route Control configuration.", "bgpRtCtrl", "Routing%20and%20Forwarding/bgp:RtCtrl/").String,

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
			"enforce_first_as": schema.StringAttribute{
				MarkdownDescription: "Enforce First AS For Ebgp. Can be configured only for VRF default.",
				Computed:            true,
			},
			"fib_accelerate": schema.StringAttribute{
				MarkdownDescription: "Accelerate the hardware updates for IP/IPv6 adjacencies for neighbor. Can be configured only for VRF default.",
				Computed:            true,
			},
			"log_neighbor_changes": schema.StringAttribute{
				MarkdownDescription: "Log Neighbor Changes.",
				Computed:            true,
			},
			"suppress_routes": schema.StringAttribute{
				MarkdownDescription: "Suppress Routes: Advertise only routes that are programmed in hardware to peers. Can be configured only for VRF default.",
				Computed:            true,
			},
		},
	}
}

func (d *BGPRouteControlDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (d *BGPRouteControlDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config BGPRouteControl

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	queries := []func(*nxos.Req){}
	res, err := d.clients[config.Device.ValueString()].GetDn(config.getDn(), queries...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(res, true)
	config.Dn = types.StringValue(config.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getDn()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
