# NetBird Terraform Provider

This provider is built on the [Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework) and uses code generation for easy maintenance and scalability. Currently, it implements a single resource for setup keys as a proof of concept.

## Prerequisites

Install the required tools:

```shell
go install github.com/hashicorp/terraform-plugin-codegen-framework/cmd/tfplugingen-framework@latest
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
go install github.com/hashicorp/terraform-plugin-codegen-openapi/cmd/tfplugingen-openapi@latest
go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest
```

## Key Components and Development Workflow

1. `openapi.yml`: NetBird OpenAPI specification
    - Update as needed

2. Generate NetBird Go SDK:
   ```shell
   oapi-codegen -generate=types,client,spec -o ./internal/sdk/sdk_gen.go -package sdk openapi.yml
   ```

3. `generator_config.yml`: OpenAPI to Terraform resource mapping
    - Update as needed

4. Generate provider code specification:
   ```shell
   tfplugingen-openapi generate --config ./generator_config.yml --output ./provider_code_spec.json openapi.yml
   ```

5. Generate resource definitions:
   ```shell
   tfplugingen-framework generate resources --input ./provider_code_spec.json --output internal/provider
   ```

6. Implement resource lifecycle in `internal/provider/setup_key_resource`
    - Scaffold new resources with:
      ```shell
      tfplugingen-framework scaffold resource --name <resource_name> --output-dir ./internal/provider
      ```

7. Generate documentation:
   ```shell
   tfplugindocs generate --provider-name netbird --provider-dir .
   ```

## Building and Testing

Build the provider:
```shell
go install
```

For local testing, override the terraform config (~/.terraformrc):
```hcl
provider_installation {
dev_overrides {
"github.com/netbirdio/netbird" = "<GO_BIN_PATH>"
}
direct {}
}
```

## Adding New Resources

1. Define the resource in `generator_config.yml`
2. Generate new provider code specifications
3. Generate resources
4. Scaffold the resource implementation
5. Add actual implementation

Refer to the Terraform documentation for more details on provider development.

## Files Overview

- `internal/provider/`: Contains resources and data sources
- `internal/sdk/`: NetBird Go SDK generated from OpenAPI
- `provider_code_spec.json`: Provider code specification
- `internal/provider/resource_setup_key`: Setup key resource definition
- `docs/`: Provider documentation
- `examples/`: Usage examples
