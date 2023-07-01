// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use AntiAffinityGroup instead.
type Affinity struct {
	pulumi.CustomResourceState

	// ❗ A free-form text describing the group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// ❗ The anti-affinity group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// ❗ The type of the group (`host anti-affinity` is the only supported value).
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The compute instances (IDs) members of the group.
	VirtualMachineIds pulumi.StringArrayOutput `pulumi:"virtualMachineIds"`
}

// NewAffinity registers a new resource with the given unique name, arguments, and options.
func NewAffinity(ctx *pulumi.Context,
	name string, args *AffinityArgs, opts ...pulumi.ResourceOption) (*Affinity, error) {
	if args == nil {
		args = &AffinityArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Affinity
	err := ctx.RegisterResource("exoscale:index/affinity:Affinity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAffinity gets an existing Affinity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAffinity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AffinityState, opts ...pulumi.ResourceOption) (*Affinity, error) {
	var resource Affinity
	err := ctx.ReadResource("exoscale:index/affinity:Affinity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Affinity resources.
type affinityState struct {
	// ❗ A free-form text describing the group.
	Description *string `pulumi:"description"`
	// ❗ The anti-affinity group name.
	Name *string `pulumi:"name"`
	// ❗ The type of the group (`host anti-affinity` is the only supported value).
	Type *string `pulumi:"type"`
	// The compute instances (IDs) members of the group.
	VirtualMachineIds []string `pulumi:"virtualMachineIds"`
}

type AffinityState struct {
	// ❗ A free-form text describing the group.
	Description pulumi.StringPtrInput
	// ❗ The anti-affinity group name.
	Name pulumi.StringPtrInput
	// ❗ The type of the group (`host anti-affinity` is the only supported value).
	Type pulumi.StringPtrInput
	// The compute instances (IDs) members of the group.
	VirtualMachineIds pulumi.StringArrayInput
}

func (AffinityState) ElementType() reflect.Type {
	return reflect.TypeOf((*affinityState)(nil)).Elem()
}

type affinityArgs struct {
	// ❗ A free-form text describing the group.
	Description *string `pulumi:"description"`
	// ❗ The anti-affinity group name.
	Name *string `pulumi:"name"`
	// ❗ The type of the group (`host anti-affinity` is the only supported value).
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Affinity resource.
type AffinityArgs struct {
	// ❗ A free-form text describing the group.
	Description pulumi.StringPtrInput
	// ❗ The anti-affinity group name.
	Name pulumi.StringPtrInput
	// ❗ The type of the group (`host anti-affinity` is the only supported value).
	Type pulumi.StringPtrInput
}

func (AffinityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*affinityArgs)(nil)).Elem()
}

type AffinityInput interface {
	pulumi.Input

	ToAffinityOutput() AffinityOutput
	ToAffinityOutputWithContext(ctx context.Context) AffinityOutput
}

func (*Affinity) ElementType() reflect.Type {
	return reflect.TypeOf((**Affinity)(nil)).Elem()
}

func (i *Affinity) ToAffinityOutput() AffinityOutput {
	return i.ToAffinityOutputWithContext(context.Background())
}

func (i *Affinity) ToAffinityOutputWithContext(ctx context.Context) AffinityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AffinityOutput)
}

// AffinityArrayInput is an input type that accepts AffinityArray and AffinityArrayOutput values.
// You can construct a concrete instance of `AffinityArrayInput` via:
//
//	AffinityArray{ AffinityArgs{...} }
type AffinityArrayInput interface {
	pulumi.Input

	ToAffinityArrayOutput() AffinityArrayOutput
	ToAffinityArrayOutputWithContext(context.Context) AffinityArrayOutput
}

type AffinityArray []AffinityInput

func (AffinityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Affinity)(nil)).Elem()
}

func (i AffinityArray) ToAffinityArrayOutput() AffinityArrayOutput {
	return i.ToAffinityArrayOutputWithContext(context.Background())
}

func (i AffinityArray) ToAffinityArrayOutputWithContext(ctx context.Context) AffinityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AffinityArrayOutput)
}

// AffinityMapInput is an input type that accepts AffinityMap and AffinityMapOutput values.
// You can construct a concrete instance of `AffinityMapInput` via:
//
//	AffinityMap{ "key": AffinityArgs{...} }
type AffinityMapInput interface {
	pulumi.Input

	ToAffinityMapOutput() AffinityMapOutput
	ToAffinityMapOutputWithContext(context.Context) AffinityMapOutput
}

type AffinityMap map[string]AffinityInput

func (AffinityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Affinity)(nil)).Elem()
}

func (i AffinityMap) ToAffinityMapOutput() AffinityMapOutput {
	return i.ToAffinityMapOutputWithContext(context.Background())
}

func (i AffinityMap) ToAffinityMapOutputWithContext(ctx context.Context) AffinityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AffinityMapOutput)
}

type AffinityOutput struct{ *pulumi.OutputState }

func (AffinityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Affinity)(nil)).Elem()
}

func (o AffinityOutput) ToAffinityOutput() AffinityOutput {
	return o
}

func (o AffinityOutput) ToAffinityOutputWithContext(ctx context.Context) AffinityOutput {
	return o
}

// ❗ A free-form text describing the group.
func (o AffinityOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Affinity) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// ❗ The anti-affinity group name.
func (o AffinityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Affinity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ❗ The type of the group (`host anti-affinity` is the only supported value).
func (o AffinityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Affinity) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// The compute instances (IDs) members of the group.
func (o AffinityOutput) VirtualMachineIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Affinity) pulumi.StringArrayOutput { return v.VirtualMachineIds }).(pulumi.StringArrayOutput)
}

type AffinityArrayOutput struct{ *pulumi.OutputState }

func (AffinityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Affinity)(nil)).Elem()
}

func (o AffinityArrayOutput) ToAffinityArrayOutput() AffinityArrayOutput {
	return o
}

func (o AffinityArrayOutput) ToAffinityArrayOutputWithContext(ctx context.Context) AffinityArrayOutput {
	return o
}

func (o AffinityArrayOutput) Index(i pulumi.IntInput) AffinityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Affinity {
		return vs[0].([]*Affinity)[vs[1].(int)]
	}).(AffinityOutput)
}

type AffinityMapOutput struct{ *pulumi.OutputState }

func (AffinityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Affinity)(nil)).Elem()
}

func (o AffinityMapOutput) ToAffinityMapOutput() AffinityMapOutput {
	return o
}

func (o AffinityMapOutput) ToAffinityMapOutputWithContext(ctx context.Context) AffinityMapOutput {
	return o
}

func (o AffinityMapOutput) MapIndex(k pulumi.StringInput) AffinityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Affinity {
		return vs[0].(map[string]*Affinity)[vs[1].(string)]
	}).(AffinityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AffinityInput)(nil)).Elem(), &Affinity{})
	pulumi.RegisterInputType(reflect.TypeOf((*AffinityArrayInput)(nil)).Elem(), AffinityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AffinityMapInput)(nil)).Elem(), AffinityMap{})
	pulumi.RegisterOutputType(AffinityOutput{})
	pulumi.RegisterOutputType(AffinityArrayOutput{})
	pulumi.RegisterOutputType(AffinityMapOutput{})
}
