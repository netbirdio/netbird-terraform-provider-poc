package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*netbirdProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &netbirdProvider{}
	}
}

type netbirdProvider struct {
}

func (p *netbirdProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {

}

func (p *netbirdProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
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
