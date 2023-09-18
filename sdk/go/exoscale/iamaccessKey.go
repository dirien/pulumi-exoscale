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

type IAMAccessKey struct {
	pulumi.CustomResourceState

	// The IAM access key (identifier).
	Key pulumi.StringOutput `pulumi:"key"`
	// ❗ The IAM access key name.
	Name pulumi.StringOutput `pulumi:"name"`
	// ❗ A list of API operations to restrict the key to.
	Operations pulumi.StringArrayOutput `pulumi:"operations"`
	// ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
	Resources pulumi.StringArrayOutput `pulumi:"resources"`
	// The key secret.
	Secret pulumi.StringOutput `pulumi:"secret"`
	// ❗ A list of tags to restrict the key to.
	Tags           pulumi.StringArrayOutput `pulumi:"tags"`
	TagsOperations pulumi.StringArrayOutput `pulumi:"tagsOperations"`
}

// NewIAMAccessKey registers a new resource with the given unique name, arguments, and options.
func NewIAMAccessKey(ctx *pulumi.Context,
	name string, args *IAMAccessKeyArgs, opts ...pulumi.ResourceOption) (*IAMAccessKey, error) {
	if args == nil {
		args = &IAMAccessKeyArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"key",
		"secret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IAMAccessKey
	err := ctx.RegisterResource("exoscale:index/iAMAccessKey:IAMAccessKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMAccessKey gets an existing IAMAccessKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMAccessKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMAccessKeyState, opts ...pulumi.ResourceOption) (*IAMAccessKey, error) {
	var resource IAMAccessKey
	err := ctx.ReadResource("exoscale:index/iAMAccessKey:IAMAccessKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMAccessKey resources.
type iamaccessKeyState struct {
	// The IAM access key (identifier).
	Key *string `pulumi:"key"`
	// ❗ The IAM access key name.
	Name *string `pulumi:"name"`
	// ❗ A list of API operations to restrict the key to.
	Operations []string `pulumi:"operations"`
	// ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
	Resources []string `pulumi:"resources"`
	// The key secret.
	Secret *string `pulumi:"secret"`
	// ❗ A list of tags to restrict the key to.
	Tags           []string `pulumi:"tags"`
	TagsOperations []string `pulumi:"tagsOperations"`
}

type IAMAccessKeyState struct {
	// The IAM access key (identifier).
	Key pulumi.StringPtrInput
	// ❗ The IAM access key name.
	Name pulumi.StringPtrInput
	// ❗ A list of API operations to restrict the key to.
	Operations pulumi.StringArrayInput
	// ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
	Resources pulumi.StringArrayInput
	// The key secret.
	Secret pulumi.StringPtrInput
	// ❗ A list of tags to restrict the key to.
	Tags           pulumi.StringArrayInput
	TagsOperations pulumi.StringArrayInput
}

func (IAMAccessKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamaccessKeyState)(nil)).Elem()
}

type iamaccessKeyArgs struct {
	// ❗ The IAM access key name.
	Name *string `pulumi:"name"`
	// ❗ A list of API operations to restrict the key to.
	Operations []string `pulumi:"operations"`
	// ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
	Resources []string `pulumi:"resources"`
	// ❗ A list of tags to restrict the key to.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a IAMAccessKey resource.
type IAMAccessKeyArgs struct {
	// ❗ The IAM access key name.
	Name pulumi.StringPtrInput
	// ❗ A list of API operations to restrict the key to.
	Operations pulumi.StringArrayInput
	// ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
	Resources pulumi.StringArrayInput
	// ❗ A list of tags to restrict the key to.
	Tags pulumi.StringArrayInput
}

func (IAMAccessKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamaccessKeyArgs)(nil)).Elem()
}

type IAMAccessKeyInput interface {
	pulumi.Input

	ToIAMAccessKeyOutput() IAMAccessKeyOutput
	ToIAMAccessKeyOutputWithContext(ctx context.Context) IAMAccessKeyOutput
}

func (*IAMAccessKey) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMAccessKey)(nil)).Elem()
}

func (i *IAMAccessKey) ToIAMAccessKeyOutput() IAMAccessKeyOutput {
	return i.ToIAMAccessKeyOutputWithContext(context.Background())
}

func (i *IAMAccessKey) ToIAMAccessKeyOutputWithContext(ctx context.Context) IAMAccessKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAccessKeyOutput)
}

func (i *IAMAccessKey) ToOutput(ctx context.Context) pulumix.Output[*IAMAccessKey] {
	return pulumix.Output[*IAMAccessKey]{
		OutputState: i.ToIAMAccessKeyOutputWithContext(ctx).OutputState,
	}
}

// IAMAccessKeyArrayInput is an input type that accepts IAMAccessKeyArray and IAMAccessKeyArrayOutput values.
// You can construct a concrete instance of `IAMAccessKeyArrayInput` via:
//
//	IAMAccessKeyArray{ IAMAccessKeyArgs{...} }
type IAMAccessKeyArrayInput interface {
	pulumi.Input

	ToIAMAccessKeyArrayOutput() IAMAccessKeyArrayOutput
	ToIAMAccessKeyArrayOutputWithContext(context.Context) IAMAccessKeyArrayOutput
}

type IAMAccessKeyArray []IAMAccessKeyInput

func (IAMAccessKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMAccessKey)(nil)).Elem()
}

func (i IAMAccessKeyArray) ToIAMAccessKeyArrayOutput() IAMAccessKeyArrayOutput {
	return i.ToIAMAccessKeyArrayOutputWithContext(context.Background())
}

func (i IAMAccessKeyArray) ToIAMAccessKeyArrayOutputWithContext(ctx context.Context) IAMAccessKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAccessKeyArrayOutput)
}

func (i IAMAccessKeyArray) ToOutput(ctx context.Context) pulumix.Output[[]*IAMAccessKey] {
	return pulumix.Output[[]*IAMAccessKey]{
		OutputState: i.ToIAMAccessKeyArrayOutputWithContext(ctx).OutputState,
	}
}

// IAMAccessKeyMapInput is an input type that accepts IAMAccessKeyMap and IAMAccessKeyMapOutput values.
// You can construct a concrete instance of `IAMAccessKeyMapInput` via:
//
//	IAMAccessKeyMap{ "key": IAMAccessKeyArgs{...} }
type IAMAccessKeyMapInput interface {
	pulumi.Input

	ToIAMAccessKeyMapOutput() IAMAccessKeyMapOutput
	ToIAMAccessKeyMapOutputWithContext(context.Context) IAMAccessKeyMapOutput
}

type IAMAccessKeyMap map[string]IAMAccessKeyInput

func (IAMAccessKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMAccessKey)(nil)).Elem()
}

func (i IAMAccessKeyMap) ToIAMAccessKeyMapOutput() IAMAccessKeyMapOutput {
	return i.ToIAMAccessKeyMapOutputWithContext(context.Background())
}

func (i IAMAccessKeyMap) ToIAMAccessKeyMapOutputWithContext(ctx context.Context) IAMAccessKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAccessKeyMapOutput)
}

func (i IAMAccessKeyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*IAMAccessKey] {
	return pulumix.Output[map[string]*IAMAccessKey]{
		OutputState: i.ToIAMAccessKeyMapOutputWithContext(ctx).OutputState,
	}
}

type IAMAccessKeyOutput struct{ *pulumi.OutputState }

func (IAMAccessKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMAccessKey)(nil)).Elem()
}

func (o IAMAccessKeyOutput) ToIAMAccessKeyOutput() IAMAccessKeyOutput {
	return o
}

func (o IAMAccessKeyOutput) ToIAMAccessKeyOutputWithContext(ctx context.Context) IAMAccessKeyOutput {
	return o
}

func (o IAMAccessKeyOutput) ToOutput(ctx context.Context) pulumix.Output[*IAMAccessKey] {
	return pulumix.Output[*IAMAccessKey]{
		OutputState: o.OutputState,
	}
}

// The IAM access key (identifier).
func (o IAMAccessKeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMAccessKey) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// ❗ The IAM access key name.
func (o IAMAccessKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMAccessKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ❗ A list of API operations to restrict the key to.
func (o IAMAccessKeyOutput) Operations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IAMAccessKey) pulumi.StringArrayOutput { return v.Operations }).(pulumi.StringArrayOutput)
}

// ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
func (o IAMAccessKeyOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IAMAccessKey) pulumi.StringArrayOutput { return v.Resources }).(pulumi.StringArrayOutput)
}

// The key secret.
func (o IAMAccessKeyOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMAccessKey) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

// ❗ A list of tags to restrict the key to.
func (o IAMAccessKeyOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IAMAccessKey) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o IAMAccessKeyOutput) TagsOperations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IAMAccessKey) pulumi.StringArrayOutput { return v.TagsOperations }).(pulumi.StringArrayOutput)
}

type IAMAccessKeyArrayOutput struct{ *pulumi.OutputState }

func (IAMAccessKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMAccessKey)(nil)).Elem()
}

func (o IAMAccessKeyArrayOutput) ToIAMAccessKeyArrayOutput() IAMAccessKeyArrayOutput {
	return o
}

func (o IAMAccessKeyArrayOutput) ToIAMAccessKeyArrayOutputWithContext(ctx context.Context) IAMAccessKeyArrayOutput {
	return o
}

func (o IAMAccessKeyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*IAMAccessKey] {
	return pulumix.Output[[]*IAMAccessKey]{
		OutputState: o.OutputState,
	}
}

func (o IAMAccessKeyArrayOutput) Index(i pulumi.IntInput) IAMAccessKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IAMAccessKey {
		return vs[0].([]*IAMAccessKey)[vs[1].(int)]
	}).(IAMAccessKeyOutput)
}

type IAMAccessKeyMapOutput struct{ *pulumi.OutputState }

func (IAMAccessKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMAccessKey)(nil)).Elem()
}

func (o IAMAccessKeyMapOutput) ToIAMAccessKeyMapOutput() IAMAccessKeyMapOutput {
	return o
}

func (o IAMAccessKeyMapOutput) ToIAMAccessKeyMapOutputWithContext(ctx context.Context) IAMAccessKeyMapOutput {
	return o
}

func (o IAMAccessKeyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*IAMAccessKey] {
	return pulumix.Output[map[string]*IAMAccessKey]{
		OutputState: o.OutputState,
	}
}

func (o IAMAccessKeyMapOutput) MapIndex(k pulumi.StringInput) IAMAccessKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IAMAccessKey {
		return vs[0].(map[string]*IAMAccessKey)[vs[1].(string)]
	}).(IAMAccessKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IAMAccessKeyInput)(nil)).Elem(), &IAMAccessKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMAccessKeyArrayInput)(nil)).Elem(), IAMAccessKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMAccessKeyMapInput)(nil)).Elem(), IAMAccessKeyMap{})
	pulumi.RegisterOutputType(IAMAccessKeyOutput{})
	pulumi.RegisterOutputType(IAMAccessKeyArrayOutput{})
	pulumi.RegisterOutputType(IAMAccessKeyMapOutput{})
}
