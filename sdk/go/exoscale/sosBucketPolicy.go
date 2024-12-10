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

// Manage Exoscale [SOS Bucket Policies](https://community.exoscale.com/documentation/storage/bucketpolicy/).
type SosBucketPolicy struct {
	pulumi.CustomResourceState

	// ❗ The name of the bucket to which the policy is to be applied.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The content of the policy
	Policy   pulumi.StringOutput              `pulumi:"policy"`
	Timeouts SosBucketPolicyTimeoutsPtrOutput `pulumi:"timeouts"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewSosBucketPolicy registers a new resource with the given unique name, arguments, and options.
func NewSosBucketPolicy(ctx *pulumi.Context,
	name string, args *SosBucketPolicyArgs, opts ...pulumi.ResourceOption) (*SosBucketPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SosBucketPolicy
	err := ctx.RegisterResource("exoscale:index/sosBucketPolicy:SosBucketPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSosBucketPolicy gets an existing SosBucketPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSosBucketPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SosBucketPolicyState, opts ...pulumi.ResourceOption) (*SosBucketPolicy, error) {
	var resource SosBucketPolicy
	err := ctx.ReadResource("exoscale:index/sosBucketPolicy:SosBucketPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SosBucketPolicy resources.
type sosBucketPolicyState struct {
	// ❗ The name of the bucket to which the policy is to be applied.
	Bucket *string `pulumi:"bucket"`
	// The content of the policy
	Policy   *string                  `pulumi:"policy"`
	Timeouts *SosBucketPolicyTimeouts `pulumi:"timeouts"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `pulumi:"zone"`
}

type SosBucketPolicyState struct {
	// ❗ The name of the bucket to which the policy is to be applied.
	Bucket pulumi.StringPtrInput
	// The content of the policy
	Policy   pulumi.StringPtrInput
	Timeouts SosBucketPolicyTimeoutsPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringPtrInput
}

func (SosBucketPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sosBucketPolicyState)(nil)).Elem()
}

type sosBucketPolicyArgs struct {
	// ❗ The name of the bucket to which the policy is to be applied.
	Bucket string `pulumi:"bucket"`
	// The content of the policy
	Policy   string                   `pulumi:"policy"`
	Timeouts *SosBucketPolicyTimeouts `pulumi:"timeouts"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a SosBucketPolicy resource.
type SosBucketPolicyArgs struct {
	// ❗ The name of the bucket to which the policy is to be applied.
	Bucket pulumi.StringInput
	// The content of the policy
	Policy   pulumi.StringInput
	Timeouts SosBucketPolicyTimeoutsPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput
}

func (SosBucketPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sosBucketPolicyArgs)(nil)).Elem()
}

type SosBucketPolicyInput interface {
	pulumi.Input

	ToSosBucketPolicyOutput() SosBucketPolicyOutput
	ToSosBucketPolicyOutputWithContext(ctx context.Context) SosBucketPolicyOutput
}

func (*SosBucketPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SosBucketPolicy)(nil)).Elem()
}

func (i *SosBucketPolicy) ToSosBucketPolicyOutput() SosBucketPolicyOutput {
	return i.ToSosBucketPolicyOutputWithContext(context.Background())
}

func (i *SosBucketPolicy) ToSosBucketPolicyOutputWithContext(ctx context.Context) SosBucketPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SosBucketPolicyOutput)
}

// SosBucketPolicyArrayInput is an input type that accepts SosBucketPolicyArray and SosBucketPolicyArrayOutput values.
// You can construct a concrete instance of `SosBucketPolicyArrayInput` via:
//
//	SosBucketPolicyArray{ SosBucketPolicyArgs{...} }
type SosBucketPolicyArrayInput interface {
	pulumi.Input

	ToSosBucketPolicyArrayOutput() SosBucketPolicyArrayOutput
	ToSosBucketPolicyArrayOutputWithContext(context.Context) SosBucketPolicyArrayOutput
}

type SosBucketPolicyArray []SosBucketPolicyInput

func (SosBucketPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SosBucketPolicy)(nil)).Elem()
}

func (i SosBucketPolicyArray) ToSosBucketPolicyArrayOutput() SosBucketPolicyArrayOutput {
	return i.ToSosBucketPolicyArrayOutputWithContext(context.Background())
}

func (i SosBucketPolicyArray) ToSosBucketPolicyArrayOutputWithContext(ctx context.Context) SosBucketPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SosBucketPolicyArrayOutput)
}

// SosBucketPolicyMapInput is an input type that accepts SosBucketPolicyMap and SosBucketPolicyMapOutput values.
// You can construct a concrete instance of `SosBucketPolicyMapInput` via:
//
//	SosBucketPolicyMap{ "key": SosBucketPolicyArgs{...} }
type SosBucketPolicyMapInput interface {
	pulumi.Input

	ToSosBucketPolicyMapOutput() SosBucketPolicyMapOutput
	ToSosBucketPolicyMapOutputWithContext(context.Context) SosBucketPolicyMapOutput
}

type SosBucketPolicyMap map[string]SosBucketPolicyInput

func (SosBucketPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SosBucketPolicy)(nil)).Elem()
}

func (i SosBucketPolicyMap) ToSosBucketPolicyMapOutput() SosBucketPolicyMapOutput {
	return i.ToSosBucketPolicyMapOutputWithContext(context.Background())
}

func (i SosBucketPolicyMap) ToSosBucketPolicyMapOutputWithContext(ctx context.Context) SosBucketPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SosBucketPolicyMapOutput)
}

type SosBucketPolicyOutput struct{ *pulumi.OutputState }

func (SosBucketPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SosBucketPolicy)(nil)).Elem()
}

func (o SosBucketPolicyOutput) ToSosBucketPolicyOutput() SosBucketPolicyOutput {
	return o
}

func (o SosBucketPolicyOutput) ToSosBucketPolicyOutputWithContext(ctx context.Context) SosBucketPolicyOutput {
	return o
}

// ❗ The name of the bucket to which the policy is to be applied.
func (o SosBucketPolicyOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *SosBucketPolicy) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The content of the policy
func (o SosBucketPolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *SosBucketPolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

func (o SosBucketPolicyOutput) Timeouts() SosBucketPolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *SosBucketPolicy) SosBucketPolicyTimeoutsPtrOutput { return v.Timeouts }).(SosBucketPolicyTimeoutsPtrOutput)
}

// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o SosBucketPolicyOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *SosBucketPolicy) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type SosBucketPolicyArrayOutput struct{ *pulumi.OutputState }

func (SosBucketPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SosBucketPolicy)(nil)).Elem()
}

func (o SosBucketPolicyArrayOutput) ToSosBucketPolicyArrayOutput() SosBucketPolicyArrayOutput {
	return o
}

func (o SosBucketPolicyArrayOutput) ToSosBucketPolicyArrayOutputWithContext(ctx context.Context) SosBucketPolicyArrayOutput {
	return o
}

func (o SosBucketPolicyArrayOutput) Index(i pulumi.IntInput) SosBucketPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SosBucketPolicy {
		return vs[0].([]*SosBucketPolicy)[vs[1].(int)]
	}).(SosBucketPolicyOutput)
}

type SosBucketPolicyMapOutput struct{ *pulumi.OutputState }

func (SosBucketPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SosBucketPolicy)(nil)).Elem()
}

func (o SosBucketPolicyMapOutput) ToSosBucketPolicyMapOutput() SosBucketPolicyMapOutput {
	return o
}

func (o SosBucketPolicyMapOutput) ToSosBucketPolicyMapOutputWithContext(ctx context.Context) SosBucketPolicyMapOutput {
	return o
}

func (o SosBucketPolicyMapOutput) MapIndex(k pulumi.StringInput) SosBucketPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SosBucketPolicy {
		return vs[0].(map[string]*SosBucketPolicy)[vs[1].(string)]
	}).(SosBucketPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SosBucketPolicyInput)(nil)).Elem(), &SosBucketPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SosBucketPolicyArrayInput)(nil)).Elem(), SosBucketPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SosBucketPolicyMapInput)(nil)).Elem(), SosBucketPolicyMap{})
	pulumi.RegisterOutputType(SosBucketPolicyOutput{})
	pulumi.RegisterOutputType(SosBucketPolicyArrayOutput{})
	pulumi.RegisterOutputType(SosBucketPolicyMapOutput{})
}
