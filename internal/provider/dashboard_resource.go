package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/qustavo/terraform-provider-voltage/internal/voltage"
)

var dashboardSchemaV1 = schema.Schema{
	Description: "Creates and manage a dashboard in Voltage",
	Version:     1,
	Attributes: map[string]schema.Attribute{
		"dashboard_id": schema.StringAttribute{
			Computed: true,
		},
		"created": schema.StringAttribute{
			Computed: true,
		},
		"node_name": schema.StringAttribute{
			Computed: true,
		},
		"dashboard_name": schema.StringAttribute{
			Computed: true,
		},
		"status": schema.StringAttribute{
			Computed: true,
		},
		"endpoint": schema.StringAttribute{
			Computed: true,
		},
		"version": schema.StringAttribute{
			Computed: true,
		},
		"update_available": schema.BoolAttribute{
			Computed: true,
		},
		"node_id": schema.StringAttribute{
			Description: "Node which connects to the dashboard",
			Required:    true,
		},
		"type": schema.StringAttribute{
			Description: "Purchase type of the node. Can be either 'trial', 'paid', or 'ondemand'.",
			Required:    true,
			Validators: []validator.String{
				stringvalidator.OneOf("thunderhub", "lnbits"),
			},
		},
	},
}

type dashboardModel struct {
	DashboardID     types.String `tfsdk:"dashboard_id"`
	Created         types.String `tfsdk:"created"`
	DashboardName   types.String `tfsdk:"dashboard_name"`
	NodeName        types.String `tfsdk:"node_name"`
	NodeID          types.String `tfsdk:"node_id"`
	Status          types.String `tfsdk:"status"`
	Type            types.String `tfsdk:"type"`
	Endpoint        types.String `tfsdk:"endpoint"`
	Version         types.String `tfsdk:"version"`
	UpdateAvailable types.Bool   `tfsdk:"update_available"`
}

type DashboardResource struct {
	client *Client
}

func NewDashboardResource() resource.Resource {
	return &DashboardResource{}
}

func (r *DashboardResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dashboard"
}

func (r *DashboardResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = dashboardSchemaV1
}

func (r *DashboardResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*voltage.ClientWithResponses)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected '*voltage.Client', got: '%T'. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = NewClient(client)
}

// func errToDiags(err error) diag.Diagnostics {
// 	if err == nil {
// 		return nil
// 	}

// 	var (
// 		diags   diag.Diagnostics
// 		cErr    *ClientError
// 		summary string
// 	)

// 	if errors.As(err, &cErr) {
// 		summary = cErr.op
// 	} else if errors.Is(err, ErrInvalidAPIResponseBody) {
// 		summary = "The API server response was invalid"
// 	} else {
// 		summary = "There was an API error"
// 	}

// 	diags.AddError(summary, err.Error())

// 	return diags
// }

func (r *DashboardResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan dashboardModel

	// Get the state from the plan.
	diag := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.CreateDashboard(ctx, &plan); err != nil {
		resp.Diagnostics.Append(errToDiags(err)...)

		return
	}

	resp.Diagnostics.Append(
		resp.State.Set(ctx, &plan)...,
	)
}

func (r *DashboardResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state dashboardModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.ReadDashboard(ctx, state.NodeID.ValueString()); err != nil {
		resp.Diagnostics.Append(errToDiags(err)...)

		return
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
func (r *DashboardResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not implemented", "You cannot update a dashboard")
}

func (r *DashboardResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state dashboardModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteDashboard(ctx, state.DashboardID.ValueString()); err != nil {
		resp.Diagnostics.Append(errToDiags(err)...)

		return
	}
}
