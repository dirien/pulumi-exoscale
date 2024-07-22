// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

// Manage Exoscale [Instance Pools](https://community.exoscale.com/documentation/compute/instance-pools/).
//
// Corresponding data sources: exoscale_instance_pool, exoscale_instance_pool_list.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myTemplate, err := exoscale.GetTemplate(ctx, &exoscale.GetTemplateArgs{
//				Zone: "ch-gva-2",
//				Name: pulumi.StringRef("Linux Ubuntu 22.04 LTS 64-bit"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = exoscale.NewInstancePool(ctx, "myInstancePool", &exoscale.InstancePoolArgs{
//				Zone:         pulumi.String("ch-gva-2"),
//				TemplateId:   pulumi.String(myTemplate.Id),
//				InstanceType: pulumi.String("standard.medium"),
//				DiskSize:     pulumi.Int(10),
//				Size:         pulumi.Int(3),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Please refer to the examples
// directory for complete configuration examples.
//
// ## Import
//
// An existing instance pool may be imported by `<ID>@<zone>`:
//
// ```sh
// $ pulumi import exoscale:index/instancePool:InstancePool \
// ```
//
//	exoscale_instance_pool.my_instance_pool \
//
//	f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
type InstancePool struct {
	pulumi.CustomResourceState

	// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
	//
	// Deprecated: Use antiAffinityGroupIds instead.
	AffinityGroupIds pulumi.StringArrayOutput `pulumi:"affinityGroupIds"`
	// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
	AntiAffinityGroupIds pulumi.StringArrayOutput `pulumi:"antiAffinityGroupIds"`
	// A deploy target ID.
	DeployTargetId pulumi.StringPtrOutput `pulumi:"deployTargetId"`
	// A free-form text describing the pool.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The managed instances disk size (GiB).
	DiskSize pulumi.IntOutput `pulumi:"diskSize"`
	// A list of exoscale*elastic*ip (IDs).
	ElasticIpIds pulumi.StringArrayOutput `pulumi:"elasticIpIds"`
	// The string used to prefix managed instances name (default: `pool`).
	InstancePrefix pulumi.StringPtrOutput `pulumi:"instancePrefix"`
	// The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The list of managed instances. Structure is documented below.
	Instances InstancePoolInstanceArrayOutput `pulumi:"instances"`
	// Enable IPv6 on managed instances (boolean; default: `false`).
	Ipv6 pulumi.BoolPtrOutput `pulumi:"ipv6"`
	// The exoscale*ssh*key (name) to authorize in the managed instances.
	KeyPair pulumi.StringPtrOutput `pulumi:"keyPair"`
	// A map of key/value labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The instance pool name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of exoscale*private*network (IDs).
	NetworkIds pulumi.StringArrayOutput `pulumi:"networkIds"`
	// A list of exoscale*security*group (IDs).
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The managed instances type. Please use the `instanceType` argument instead.
	//
	// Deprecated: This attribute has been replaced by "instanceType".
	ServiceOffering pulumi.StringOutput `pulumi:"serviceOffering"`
	// The number of managed instances.
	Size  pulumi.IntOutput    `pulumi:"size"`
	State pulumi.StringOutput `pulumi:"state"`
	// The getTemplate (ID) to use when creating the managed instances.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
	// [cloud-init](http://cloudinit.readthedocs.io/) configuration to apply to the managed instances.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// The list of managed instances (IDs). Please use the `instances.*.id` attribute instead.
	//
	// Deprecated: Use the instances exported attribute instead.
	VirtualMachines pulumi.StringArrayOutput `pulumi:"virtualMachines"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstancePool registers a new resource with the given unique name, arguments, and options.
func NewInstancePool(ctx *pulumi.Context,
	name string, args *InstancePoolArgs, opts ...pulumi.ResourceOption) (*InstancePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstancePool
	err := ctx.RegisterResource("exoscale:index/instancePool:InstancePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstancePool gets an existing InstancePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstancePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstancePoolState, opts ...pulumi.ResourceOption) (*InstancePool, error) {
	var resource InstancePool
	err := ctx.ReadResource("exoscale:index/instancePool:InstancePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstancePool resources.
type instancePoolState struct {
	// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
	//
	// Deprecated: Use antiAffinityGroupIds instead.
	AffinityGroupIds []string `pulumi:"affinityGroupIds"`
	// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
	AntiAffinityGroupIds []string `pulumi:"antiAffinityGroupIds"`
	// A deploy target ID.
	DeployTargetId *string `pulumi:"deployTargetId"`
	// A free-form text describing the pool.
	Description *string `pulumi:"description"`
	// The managed instances disk size (GiB).
	DiskSize *int `pulumi:"diskSize"`
	// A list of exoscale*elastic*ip (IDs).
	ElasticIpIds []string `pulumi:"elasticIpIds"`
	// The string used to prefix managed instances name (default: `pool`).
	InstancePrefix *string `pulumi:"instancePrefix"`
	// The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
	InstanceType *string `pulumi:"instanceType"`
	// The list of managed instances. Structure is documented below.
	Instances []InstancePoolInstance `pulumi:"instances"`
	// Enable IPv6 on managed instances (boolean; default: `false`).
	Ipv6 *bool `pulumi:"ipv6"`
	// The exoscale*ssh*key (name) to authorize in the managed instances.
	KeyPair *string `pulumi:"keyPair"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// The instance pool name.
	Name *string `pulumi:"name"`
	// A list of exoscale*private*network (IDs).
	NetworkIds []string `pulumi:"networkIds"`
	// A list of exoscale*security*group (IDs).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The managed instances type. Please use the `instanceType` argument instead.
	//
	// Deprecated: This attribute has been replaced by "instanceType".
	ServiceOffering *string `pulumi:"serviceOffering"`
	// The number of managed instances.
	Size  *int    `pulumi:"size"`
	State *string `pulumi:"state"`
	// The getTemplate (ID) to use when creating the managed instances.
	TemplateId *string `pulumi:"templateId"`
	// [cloud-init](http://cloudinit.readthedocs.io/) configuration to apply to the managed instances.
	UserData *string `pulumi:"userData"`
	// The list of managed instances (IDs). Please use the `instances.*.id` attribute instead.
	//
	// Deprecated: Use the instances exported attribute instead.
	VirtualMachines []string `pulumi:"virtualMachines"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `pulumi:"zone"`
}

type InstancePoolState struct {
	// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
	//
	// Deprecated: Use antiAffinityGroupIds instead.
	AffinityGroupIds pulumi.StringArrayInput
	// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
	AntiAffinityGroupIds pulumi.StringArrayInput
	// A deploy target ID.
	DeployTargetId pulumi.StringPtrInput
	// A free-form text describing the pool.
	Description pulumi.StringPtrInput
	// The managed instances disk size (GiB).
	DiskSize pulumi.IntPtrInput
	// A list of exoscale*elastic*ip (IDs).
	ElasticIpIds pulumi.StringArrayInput
	// The string used to prefix managed instances name (default: `pool`).
	InstancePrefix pulumi.StringPtrInput
	// The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
	InstanceType pulumi.StringPtrInput
	// The list of managed instances. Structure is documented below.
	Instances InstancePoolInstanceArrayInput
	// Enable IPv6 on managed instances (boolean; default: `false`).
	Ipv6 pulumi.BoolPtrInput
	// The exoscale*ssh*key (name) to authorize in the managed instances.
	KeyPair pulumi.StringPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// The instance pool name.
	Name pulumi.StringPtrInput
	// A list of exoscale*private*network (IDs).
	NetworkIds pulumi.StringArrayInput
	// A list of exoscale*security*group (IDs).
	SecurityGroupIds pulumi.StringArrayInput
	// The managed instances type. Please use the `instanceType` argument instead.
	//
	// Deprecated: This attribute has been replaced by "instanceType".
	ServiceOffering pulumi.StringPtrInput
	// The number of managed instances.
	Size  pulumi.IntPtrInput
	State pulumi.StringPtrInput
	// The getTemplate (ID) to use when creating the managed instances.
	TemplateId pulumi.StringPtrInput
	// [cloud-init](http://cloudinit.readthedocs.io/) configuration to apply to the managed instances.
	UserData pulumi.StringPtrInput
	// The list of managed instances (IDs). Please use the `instances.*.id` attribute instead.
	//
	// Deprecated: Use the instances exported attribute instead.
	VirtualMachines pulumi.StringArrayInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringPtrInput
}

func (InstancePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*instancePoolState)(nil)).Elem()
}

type instancePoolArgs struct {
	// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
	//
	// Deprecated: Use antiAffinityGroupIds instead.
	AffinityGroupIds []string `pulumi:"affinityGroupIds"`
	// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
	AntiAffinityGroupIds []string `pulumi:"antiAffinityGroupIds"`
	// A deploy target ID.
	DeployTargetId *string `pulumi:"deployTargetId"`
	// A free-form text describing the pool.
	Description *string `pulumi:"description"`
	// The managed instances disk size (GiB).
	DiskSize *int `pulumi:"diskSize"`
	// A list of exoscale*elastic*ip (IDs).
	ElasticIpIds []string `pulumi:"elasticIpIds"`
	// The string used to prefix managed instances name (default: `pool`).
	InstancePrefix *string `pulumi:"instancePrefix"`
	// The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
	InstanceType *string `pulumi:"instanceType"`
	// The list of managed instances. Structure is documented below.
	Instances []InstancePoolInstance `pulumi:"instances"`
	// Enable IPv6 on managed instances (boolean; default: `false`).
	Ipv6 *bool `pulumi:"ipv6"`
	// The exoscale*ssh*key (name) to authorize in the managed instances.
	KeyPair *string `pulumi:"keyPair"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// The instance pool name.
	Name *string `pulumi:"name"`
	// A list of exoscale*private*network (IDs).
	NetworkIds []string `pulumi:"networkIds"`
	// A list of exoscale*security*group (IDs).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The managed instances type. Please use the `instanceType` argument instead.
	//
	// Deprecated: This attribute has been replaced by "instanceType".
	ServiceOffering *string `pulumi:"serviceOffering"`
	// The number of managed instances.
	Size  int     `pulumi:"size"`
	State *string `pulumi:"state"`
	// The getTemplate (ID) to use when creating the managed instances.
	TemplateId string `pulumi:"templateId"`
	// [cloud-init](http://cloudinit.readthedocs.io/) configuration to apply to the managed instances.
	UserData *string `pulumi:"userData"`
	// The list of managed instances (IDs). Please use the `instances.*.id` attribute instead.
	//
	// Deprecated: Use the instances exported attribute instead.
	VirtualMachines []string `pulumi:"virtualMachines"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a InstancePool resource.
type InstancePoolArgs struct {
	// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
	//
	// Deprecated: Use antiAffinityGroupIds instead.
	AffinityGroupIds pulumi.StringArrayInput
	// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
	AntiAffinityGroupIds pulumi.StringArrayInput
	// A deploy target ID.
	DeployTargetId pulumi.StringPtrInput
	// A free-form text describing the pool.
	Description pulumi.StringPtrInput
	// The managed instances disk size (GiB).
	DiskSize pulumi.IntPtrInput
	// A list of exoscale*elastic*ip (IDs).
	ElasticIpIds pulumi.StringArrayInput
	// The string used to prefix managed instances name (default: `pool`).
	InstancePrefix pulumi.StringPtrInput
	// The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
	InstanceType pulumi.StringPtrInput
	// The list of managed instances. Structure is documented below.
	Instances InstancePoolInstanceArrayInput
	// Enable IPv6 on managed instances (boolean; default: `false`).
	Ipv6 pulumi.BoolPtrInput
	// The exoscale*ssh*key (name) to authorize in the managed instances.
	KeyPair pulumi.StringPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// The instance pool name.
	Name pulumi.StringPtrInput
	// A list of exoscale*private*network (IDs).
	NetworkIds pulumi.StringArrayInput
	// A list of exoscale*security*group (IDs).
	SecurityGroupIds pulumi.StringArrayInput
	// The managed instances type. Please use the `instanceType` argument instead.
	//
	// Deprecated: This attribute has been replaced by "instanceType".
	ServiceOffering pulumi.StringPtrInput
	// The number of managed instances.
	Size  pulumi.IntInput
	State pulumi.StringPtrInput
	// The getTemplate (ID) to use when creating the managed instances.
	TemplateId pulumi.StringInput
	// [cloud-init](http://cloudinit.readthedocs.io/) configuration to apply to the managed instances.
	UserData pulumi.StringPtrInput
	// The list of managed instances (IDs). Please use the `instances.*.id` attribute instead.
	//
	// Deprecated: Use the instances exported attribute instead.
	VirtualMachines pulumi.StringArrayInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput
}

func (InstancePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instancePoolArgs)(nil)).Elem()
}

type InstancePoolInput interface {
	pulumi.Input

	ToInstancePoolOutput() InstancePoolOutput
	ToInstancePoolOutputWithContext(ctx context.Context) InstancePoolOutput
}

func (*InstancePool) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancePool)(nil)).Elem()
}

func (i *InstancePool) ToInstancePoolOutput() InstancePoolOutput {
	return i.ToInstancePoolOutputWithContext(context.Background())
}

func (i *InstancePool) ToInstancePoolOutputWithContext(ctx context.Context) InstancePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePoolOutput)
}

// InstancePoolArrayInput is an input type that accepts InstancePoolArray and InstancePoolArrayOutput values.
// You can construct a concrete instance of `InstancePoolArrayInput` via:
//
//	InstancePoolArray{ InstancePoolArgs{...} }
type InstancePoolArrayInput interface {
	pulumi.Input

	ToInstancePoolArrayOutput() InstancePoolArrayOutput
	ToInstancePoolArrayOutputWithContext(context.Context) InstancePoolArrayOutput
}

type InstancePoolArray []InstancePoolInput

func (InstancePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstancePool)(nil)).Elem()
}

func (i InstancePoolArray) ToInstancePoolArrayOutput() InstancePoolArrayOutput {
	return i.ToInstancePoolArrayOutputWithContext(context.Background())
}

func (i InstancePoolArray) ToInstancePoolArrayOutputWithContext(ctx context.Context) InstancePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePoolArrayOutput)
}

// InstancePoolMapInput is an input type that accepts InstancePoolMap and InstancePoolMapOutput values.
// You can construct a concrete instance of `InstancePoolMapInput` via:
//
//	InstancePoolMap{ "key": InstancePoolArgs{...} }
type InstancePoolMapInput interface {
	pulumi.Input

	ToInstancePoolMapOutput() InstancePoolMapOutput
	ToInstancePoolMapOutputWithContext(context.Context) InstancePoolMapOutput
}

type InstancePoolMap map[string]InstancePoolInput

func (InstancePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstancePool)(nil)).Elem()
}

func (i InstancePoolMap) ToInstancePoolMapOutput() InstancePoolMapOutput {
	return i.ToInstancePoolMapOutputWithContext(context.Background())
}

func (i InstancePoolMap) ToInstancePoolMapOutputWithContext(ctx context.Context) InstancePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePoolMapOutput)
}

type InstancePoolOutput struct{ *pulumi.OutputState }

func (InstancePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancePool)(nil)).Elem()
}

func (o InstancePoolOutput) ToInstancePoolOutput() InstancePoolOutput {
	return o
}

func (o InstancePoolOutput) ToInstancePoolOutputWithContext(ctx context.Context) InstancePoolOutput {
	return o
}

// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
//
// Deprecated: Use antiAffinityGroupIds instead.
func (o InstancePoolOutput) AffinityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringArrayOutput { return v.AffinityGroupIds }).(pulumi.StringArrayOutput)
}

// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
func (o InstancePoolOutput) AntiAffinityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringArrayOutput { return v.AntiAffinityGroupIds }).(pulumi.StringArrayOutput)
}

// A deploy target ID.
func (o InstancePoolOutput) DeployTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringPtrOutput { return v.DeployTargetId }).(pulumi.StringPtrOutput)
}

// A free-form text describing the pool.
func (o InstancePoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The managed instances disk size (GiB).
func (o InstancePoolOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.IntOutput { return v.DiskSize }).(pulumi.IntOutput)
}

// A list of exoscale*elastic*ip (IDs).
func (o InstancePoolOutput) ElasticIpIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringArrayOutput { return v.ElasticIpIds }).(pulumi.StringArrayOutput)
}

// The string used to prefix managed instances name (default: `pool`).
func (o InstancePoolOutput) InstancePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringPtrOutput { return v.InstancePrefix }).(pulumi.StringPtrOutput)
}

// The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
func (o InstancePoolOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// The list of managed instances. Structure is documented below.
func (o InstancePoolOutput) Instances() InstancePoolInstanceArrayOutput {
	return o.ApplyT(func(v *InstancePool) InstancePoolInstanceArrayOutput { return v.Instances }).(InstancePoolInstanceArrayOutput)
}

// Enable IPv6 on managed instances (boolean; default: `false`).
func (o InstancePoolOutput) Ipv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.BoolPtrOutput { return v.Ipv6 }).(pulumi.BoolPtrOutput)
}

// The exoscale*ssh*key (name) to authorize in the managed instances.
func (o InstancePoolOutput) KeyPair() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringPtrOutput { return v.KeyPair }).(pulumi.StringPtrOutput)
}

// A map of key/value labels.
func (o InstancePoolOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The instance pool name.
func (o InstancePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of exoscale*private*network (IDs).
func (o InstancePoolOutput) NetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringArrayOutput { return v.NetworkIds }).(pulumi.StringArrayOutput)
}

// A list of exoscale*security*group (IDs).
func (o InstancePoolOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The managed instances type. Please use the `instanceType` argument instead.
//
// Deprecated: This attribute has been replaced by "instanceType".
func (o InstancePoolOutput) ServiceOffering() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.ServiceOffering }).(pulumi.StringOutput)
}

// The number of managed instances.
func (o InstancePoolOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

func (o InstancePoolOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The getTemplate (ID) to use when creating the managed instances.
func (o InstancePoolOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

// [cloud-init](http://cloudinit.readthedocs.io/) configuration to apply to the managed instances.
func (o InstancePoolOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringPtrOutput { return v.UserData }).(pulumi.StringPtrOutput)
}

// The list of managed instances (IDs). Please use the `instances.*.id` attribute instead.
//
// Deprecated: Use the instances exported attribute instead.
func (o InstancePoolOutput) VirtualMachines() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringArrayOutput { return v.VirtualMachines }).(pulumi.StringArrayOutput)
}

// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o InstancePoolOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type InstancePoolArrayOutput struct{ *pulumi.OutputState }

func (InstancePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstancePool)(nil)).Elem()
}

func (o InstancePoolArrayOutput) ToInstancePoolArrayOutput() InstancePoolArrayOutput {
	return o
}

func (o InstancePoolArrayOutput) ToInstancePoolArrayOutputWithContext(ctx context.Context) InstancePoolArrayOutput {
	return o
}

func (o InstancePoolArrayOutput) Index(i pulumi.IntInput) InstancePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstancePool {
		return vs[0].([]*InstancePool)[vs[1].(int)]
	}).(InstancePoolOutput)
}

type InstancePoolMapOutput struct{ *pulumi.OutputState }

func (InstancePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstancePool)(nil)).Elem()
}

func (o InstancePoolMapOutput) ToInstancePoolMapOutput() InstancePoolMapOutput {
	return o
}

func (o InstancePoolMapOutput) ToInstancePoolMapOutputWithContext(ctx context.Context) InstancePoolMapOutput {
	return o
}

func (o InstancePoolMapOutput) MapIndex(k pulumi.StringInput) InstancePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstancePool {
		return vs[0].(map[string]*InstancePool)[vs[1].(string)]
	}).(InstancePoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstancePoolInput)(nil)).Elem(), &InstancePool{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancePoolArrayInput)(nil)).Elem(), InstancePoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancePoolMapInput)(nil)).Elem(), InstancePoolMap{})
	pulumi.RegisterOutputType(InstancePoolOutput{})
	pulumi.RegisterOutputType(InstancePoolArrayOutput{})
	pulumi.RegisterOutputType(InstancePoolMapOutput{})
}
