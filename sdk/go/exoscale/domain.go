// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

// ## Import
//
// An existing DNS domain may be imported by `ID`
//
// ```sh
//
//	$ pulumi import exoscale:index/domain:Domain \
//
// ```
//
//	exoscale_domain.my_domain \
//
//	89083a5c-b648-474a-0000-0000000f67bd
type Domain struct {
	pulumi.CustomResourceState

	// Whether the DNS domain has automatic renewal enabled (boolean).
	//
	// Deprecated: Not used, will be removed in the future
	AutoRenew pulumi.BoolOutput `pulumi:"autoRenew"`
	// The domain expiration date, if known.
	//
	// Deprecated: Not used, will be removed in the future
	ExpiresOn pulumi.StringOutput `pulumi:"expiresOn"`
	// ❗ The DNS domain name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The domain state.
	//
	// Deprecated: Not used, will be removed in the future
	State pulumi.StringOutput `pulumi:"state"`
	// A security token that can be used as an alternative way to manage DNS domains via the Exoscale API.
	//
	// Deprecated: Not used, will be removed in the future
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		args = &DomainArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Domain
	err := ctx.RegisterResource("exoscale:index/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("exoscale:index/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// Whether the DNS domain has automatic renewal enabled (boolean).
	//
	// Deprecated: Not used, will be removed in the future
	AutoRenew *bool `pulumi:"autoRenew"`
	// The domain expiration date, if known.
	//
	// Deprecated: Not used, will be removed in the future
	ExpiresOn *string `pulumi:"expiresOn"`
	// ❗ The DNS domain name.
	Name *string `pulumi:"name"`
	// The domain state.
	//
	// Deprecated: Not used, will be removed in the future
	State *string `pulumi:"state"`
	// A security token that can be used as an alternative way to manage DNS domains via the Exoscale API.
	//
	// Deprecated: Not used, will be removed in the future
	Token *string `pulumi:"token"`
}

type DomainState struct {
	// Whether the DNS domain has automatic renewal enabled (boolean).
	//
	// Deprecated: Not used, will be removed in the future
	AutoRenew pulumi.BoolPtrInput
	// The domain expiration date, if known.
	//
	// Deprecated: Not used, will be removed in the future
	ExpiresOn pulumi.StringPtrInput
	// ❗ The DNS domain name.
	Name pulumi.StringPtrInput
	// The domain state.
	//
	// Deprecated: Not used, will be removed in the future
	State pulumi.StringPtrInput
	// A security token that can be used as an alternative way to manage DNS domains via the Exoscale API.
	//
	// Deprecated: Not used, will be removed in the future
	Token pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// ❗ The DNS domain name.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// ❗ The DNS domain name.
	Name pulumi.StringPtrInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

func (i *Domain) ToOutput(ctx context.Context) pulumix.Output[*Domain] {
	return pulumix.Output[*Domain]{
		OutputState: i.ToDomainOutputWithContext(ctx).OutputState,
	}
}

// DomainArrayInput is an input type that accepts DomainArray and DomainArrayOutput values.
// You can construct a concrete instance of `DomainArrayInput` via:
//
//	DomainArray{ DomainArgs{...} }
type DomainArrayInput interface {
	pulumi.Input

	ToDomainArrayOutput() DomainArrayOutput
	ToDomainArrayOutputWithContext(context.Context) DomainArrayOutput
}

type DomainArray []DomainInput

func (DomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
}

func (i DomainArray) ToDomainArrayOutput() DomainArrayOutput {
	return i.ToDomainArrayOutputWithContext(context.Background())
}

func (i DomainArray) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainArrayOutput)
}

func (i DomainArray) ToOutput(ctx context.Context) pulumix.Output[[]*Domain] {
	return pulumix.Output[[]*Domain]{
		OutputState: i.ToDomainArrayOutputWithContext(ctx).OutputState,
	}
}

// DomainMapInput is an input type that accepts DomainMap and DomainMapOutput values.
// You can construct a concrete instance of `DomainMapInput` via:
//
//	DomainMap{ "key": DomainArgs{...} }
type DomainMapInput interface {
	pulumi.Input

	ToDomainMapOutput() DomainMapOutput
	ToDomainMapOutputWithContext(context.Context) DomainMapOutput
}

type DomainMap map[string]DomainInput

func (DomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (i DomainMap) ToDomainMapOutput() DomainMapOutput {
	return i.ToDomainMapOutputWithContext(context.Background())
}

func (i DomainMap) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMapOutput)
}

func (i DomainMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Domain] {
	return pulumix.Output[map[string]*Domain]{
		OutputState: i.ToDomainMapOutputWithContext(ctx).OutputState,
	}
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func (o DomainOutput) ToOutput(ctx context.Context) pulumix.Output[*Domain] {
	return pulumix.Output[*Domain]{
		OutputState: o.OutputState,
	}
}

// Whether the DNS domain has automatic renewal enabled (boolean).
//
// Deprecated: Not used, will be removed in the future
func (o DomainOutput) AutoRenew() pulumi.BoolOutput {
	return o.ApplyT(func(v *Domain) pulumi.BoolOutput { return v.AutoRenew }).(pulumi.BoolOutput)
}

// The domain expiration date, if known.
//
// Deprecated: Not used, will be removed in the future
func (o DomainOutput) ExpiresOn() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ExpiresOn }).(pulumi.StringOutput)
}

// ❗ The DNS domain name.
func (o DomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The domain state.
//
// Deprecated: Not used, will be removed in the future
func (o DomainOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// A security token that can be used as an alternative way to manage DNS domains via the Exoscale API.
//
// Deprecated: Not used, will be removed in the future
func (o DomainOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

type DomainArrayOutput struct{ *pulumi.OutputState }

func (DomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
}

func (o DomainArrayOutput) ToDomainArrayOutput() DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Domain] {
	return pulumix.Output[[]*Domain]{
		OutputState: o.OutputState,
	}
}

func (o DomainArrayOutput) Index(i pulumi.IntInput) DomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Domain {
		return vs[0].([]*Domain)[vs[1].(int)]
	}).(DomainOutput)
}

type DomainMapOutput struct{ *pulumi.OutputState }

func (DomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (o DomainMapOutput) ToDomainMapOutput() DomainMapOutput {
	return o
}

func (o DomainMapOutput) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return o
}

func (o DomainMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Domain] {
	return pulumix.Output[map[string]*Domain]{
		OutputState: o.OutputState,
	}
}

func (o DomainMapOutput) MapIndex(k pulumi.StringInput) DomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Domain {
		return vs[0].(map[string]*Domain)[vs[1].(string)]
	}).(DomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainInput)(nil)).Elem(), &Domain{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainArrayInput)(nil)).Elem(), DomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainMapInput)(nil)).Elem(), DomainMap{})
	pulumi.RegisterOutputType(DomainOutput{})
	pulumi.RegisterOutputType(DomainArrayOutput{})
	pulumi.RegisterOutputType(DomainMapOutput{})
}
