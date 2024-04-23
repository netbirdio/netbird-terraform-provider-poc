package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/netbirdio/netbird-terraform-provider/internal/sdk"
)

var _ provider.Provider = (*netbirdProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &netbirdProvider{}
	}
}

type netbirdProvider struct {
}

// NetbirdProviderModel describes the provider data model.
type NetbirdProviderModel struct {
	ServerURL  types.String `tfsdk:"server_url"`
	BearerAuth types.String `tfsdk:"bearer_auth"`
	TokenAuth  types.String `tfsdk:"token_auth"`
}

func (p *netbirdProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `NetBird REST API: API to manipulate groups, rules, policies and retrieve information about peers and users`,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://api.netbird.io)",
				Optional:            true,
				Required:            false,
			},
			"bearer_auth": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"token_auth": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *netbirdProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data NetbirdProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	serverURL := data.ServerURL.ValueString()
	if serverURL == "" {
		serverURL = "https://api.netbird.io"
	}

	bearerAuth := new(string)
	if !data.BearerAuth.IsUnknown() && !data.BearerAuth.IsNull() {
		*bearerAuth = data.BearerAuth.ValueString()
	} else {
		bearerAuth = nil
	}
	tokenAuth := new(string)
	if !data.TokenAuth.IsUnknown() && !data.TokenAuth.IsNull() {
		*tokenAuth = data.TokenAuth.ValueString()
	} else {
		tokenAuth = nil
	}

	addRequestAuth := func(ctx context.Context, req *http.Request) error {
		if bearerAuth != nil {
			req.Header.Set("Authorization", "Bearer "+*bearerAuth)
		} else if tokenAuth != nil {
			req.Header.Set("Authorization", "Token "+*tokenAuth)
		}
		return nil
	}
	client, err := sdk.NewClientWithResponses(serverURL, sdk.WithRequestEditorFn(addRequestAuth))
	if err != nil {
		resp.Diagnostics.AddError("failed to create client", err.Error())
		return
	}
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *netbirdProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "netbird"
}

func (p *netbirdProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *netbirdProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSetupKeyResource,
	}
}
