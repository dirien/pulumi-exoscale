// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the exoscale package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// Exoscale CloudStack API endpoint (by default: https://api.exoscale.com/v1)
	ComputeEndpoint pulumi.StringPtrOutput `pulumi:"computeEndpoint"`
	// CloudStack ini configuration filename (by default: cloudstack.ini)
	Config pulumi.StringPtrOutput `pulumi:"config"`
	// Exoscale DNS API endpoint (by default: https://api.exoscale.com/dns)
	DnsEndpoint pulumi.StringOutput    `pulumi:"dnsEndpoint"`
	Environment pulumi.StringPtrOutput `pulumi:"environment"`
	// Exoscale API key
	Key pulumi.StringPtrOutput `pulumi:"key"`
	// Deprecated: Use region instead
	Profile pulumi.StringPtrOutput `pulumi:"profile"`
	// CloudStack ini configuration section name (by default: cloudstack)
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// Exoscale API secret
	Secret pulumi.StringPtrOutput `pulumi:"secret"`
	// Deprecated: Use key instead
	Token pulumi.StringPtrOutput `pulumi:"token"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'DnsEndpoint'")
	}
	if args.Timeout == nil {
		return nil, errors.New("invalid value for required argument 'Timeout'")
	}
	if args.Secret != nil {
		args.Secret = pulumi.ToSecret(args.Secret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:exoscale", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// Exoscale CloudStack API endpoint (by default: https://api.exoscale.com/v1)
	ComputeEndpoint *string `pulumi:"computeEndpoint"`
	// CloudStack ini configuration filename (by default: cloudstack.ini)
	Config *string `pulumi:"config"`
	// Deprecated: Does nothing
	Delay *int `pulumi:"delay"`
	// Exoscale DNS API endpoint (by default: https://api.exoscale.com/dns)
	DnsEndpoint string  `pulumi:"dnsEndpoint"`
	Environment *string `pulumi:"environment"`
	// Defines if the user-data of compute instances should be gzipped (by default: true)
	GzipUserData *bool `pulumi:"gzipUserData"`
	// Exoscale API key
	Key *string `pulumi:"key"`
	// Deprecated: Use region instead
	Profile *string `pulumi:"profile"`
	// CloudStack ini configuration section name (by default: cloudstack)
	Region *string `pulumi:"region"`
	// Exoscale API secret
	Secret *string `pulumi:"secret"`
	// Timeout in seconds for waiting on compute resources to become available (by default: 300)
	Timeout int `pulumi:"timeout"`
	// Deprecated: Use key instead
	Token *string `pulumi:"token"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// Exoscale CloudStack API endpoint (by default: https://api.exoscale.com/v1)
	ComputeEndpoint pulumi.StringPtrInput
	// CloudStack ini configuration filename (by default: cloudstack.ini)
	Config pulumi.StringPtrInput
	// Deprecated: Does nothing
	Delay pulumi.IntPtrInput
	// Exoscale DNS API endpoint (by default: https://api.exoscale.com/dns)
	DnsEndpoint pulumi.StringInput
	Environment pulumi.StringPtrInput
	// Defines if the user-data of compute instances should be gzipped (by default: true)
	GzipUserData pulumi.BoolPtrInput
	// Exoscale API key
	Key pulumi.StringPtrInput
	// Deprecated: Use region instead
	Profile pulumi.StringPtrInput
	// CloudStack ini configuration section name (by default: cloudstack)
	Region pulumi.StringPtrInput
	// Exoscale API secret
	Secret pulumi.StringPtrInput
	// Timeout in seconds for waiting on compute resources to become available (by default: 300)
	Timeout pulumi.IntInput
	// Deprecated: Use key instead
	Token pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// Exoscale CloudStack API endpoint (by default: https://api.exoscale.com/v1)
func (o ProviderOutput) ComputeEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ComputeEndpoint }).(pulumi.StringPtrOutput)
}

// CloudStack ini configuration filename (by default: cloudstack.ini)
func (o ProviderOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Config }).(pulumi.StringPtrOutput)
}

// Exoscale DNS API endpoint (by default: https://api.exoscale.com/dns)
func (o ProviderOutput) DnsEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.DnsEndpoint }).(pulumi.StringOutput)
}

func (o ProviderOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Environment }).(pulumi.StringPtrOutput)
}

// Exoscale API key
func (o ProviderOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Key }).(pulumi.StringPtrOutput)
}

// Deprecated: Use region instead
func (o ProviderOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Profile }).(pulumi.StringPtrOutput)
}

// CloudStack ini configuration section name (by default: cloudstack)
func (o ProviderOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// Exoscale API secret
func (o ProviderOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Secret }).(pulumi.StringPtrOutput)
}

// Deprecated: Use key instead
func (o ProviderOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
