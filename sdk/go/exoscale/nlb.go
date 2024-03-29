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

// Manage Exoscale [Network Load Balancers (NLB)](https://community.exoscale.com/documentation/compute/network-load-balancer/).
//
// Corresponding data source: exoscale_nlb.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := exoscale.NewNlb(ctx, "myNlb", &exoscale.NlbArgs{
//				Zone: pulumi.String("ch-gva-2"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// Next step is to attach exoscale_nlb_service(s) to the NLB.
//
// Please refer to the examples
// directory for complete configuration examples.
//
// ## Import
//
// An existing network load balancer (NLB) may be imported by `<ID>@<zone>`:
//
// console
//
// ```sh
// $ pulumi import exoscale:index/nlb:Nlb \
// ```
//
//	exoscale_nlb.my_nlb \
//
//	f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
type Nlb struct {
	pulumi.CustomResourceState

	// The NLB creation date.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A free-form text describing the NLB.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The NLB IPv4 address.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// A map of key/value labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The network load balancer (NLB) name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of the exoscale*nlb*service (names).
	Services pulumi.StringArrayOutput `pulumi:"services"`
	// The current NLB state.
	State pulumi.StringOutput `pulumi:"state"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewNlb registers a new resource with the given unique name, arguments, and options.
func NewNlb(ctx *pulumi.Context,
	name string, args *NlbArgs, opts ...pulumi.ResourceOption) (*Nlb, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Nlb
	err := ctx.RegisterResource("exoscale:index/nlb:Nlb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNlb gets an existing Nlb resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNlb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NlbState, opts ...pulumi.ResourceOption) (*Nlb, error) {
	var resource Nlb
	err := ctx.ReadResource("exoscale:index/nlb:Nlb", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Nlb resources.
type nlbState struct {
	// The NLB creation date.
	CreatedAt *string `pulumi:"createdAt"`
	// A free-form text describing the NLB.
	Description *string `pulumi:"description"`
	// The NLB IPv4 address.
	IpAddress *string `pulumi:"ipAddress"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// The network load balancer (NLB) name.
	Name *string `pulumi:"name"`
	// The list of the exoscale*nlb*service (names).
	Services []string `pulumi:"services"`
	// The current NLB state.
	State *string `pulumi:"state"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `pulumi:"zone"`
}

type NlbState struct {
	// The NLB creation date.
	CreatedAt pulumi.StringPtrInput
	// A free-form text describing the NLB.
	Description pulumi.StringPtrInput
	// The NLB IPv4 address.
	IpAddress pulumi.StringPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// The network load balancer (NLB) name.
	Name pulumi.StringPtrInput
	// The list of the exoscale*nlb*service (names).
	Services pulumi.StringArrayInput
	// The current NLB state.
	State pulumi.StringPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringPtrInput
}

func (NlbState) ElementType() reflect.Type {
	return reflect.TypeOf((*nlbState)(nil)).Elem()
}

type nlbArgs struct {
	// A free-form text describing the NLB.
	Description *string `pulumi:"description"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// The network load balancer (NLB) name.
	Name *string `pulumi:"name"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a Nlb resource.
type NlbArgs struct {
	// A free-form text describing the NLB.
	Description pulumi.StringPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// The network load balancer (NLB) name.
	Name pulumi.StringPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput
}

func (NlbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nlbArgs)(nil)).Elem()
}

type NlbInput interface {
	pulumi.Input

	ToNlbOutput() NlbOutput
	ToNlbOutputWithContext(ctx context.Context) NlbOutput
}

func (*Nlb) ElementType() reflect.Type {
	return reflect.TypeOf((**Nlb)(nil)).Elem()
}

func (i *Nlb) ToNlbOutput() NlbOutput {
	return i.ToNlbOutputWithContext(context.Background())
}

func (i *Nlb) ToNlbOutputWithContext(ctx context.Context) NlbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NlbOutput)
}

// NlbArrayInput is an input type that accepts NlbArray and NlbArrayOutput values.
// You can construct a concrete instance of `NlbArrayInput` via:
//
//	NlbArray{ NlbArgs{...} }
type NlbArrayInput interface {
	pulumi.Input

	ToNlbArrayOutput() NlbArrayOutput
	ToNlbArrayOutputWithContext(context.Context) NlbArrayOutput
}

type NlbArray []NlbInput

func (NlbArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nlb)(nil)).Elem()
}

func (i NlbArray) ToNlbArrayOutput() NlbArrayOutput {
	return i.ToNlbArrayOutputWithContext(context.Background())
}

func (i NlbArray) ToNlbArrayOutputWithContext(ctx context.Context) NlbArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NlbArrayOutput)
}

// NlbMapInput is an input type that accepts NlbMap and NlbMapOutput values.
// You can construct a concrete instance of `NlbMapInput` via:
//
//	NlbMap{ "key": NlbArgs{...} }
type NlbMapInput interface {
	pulumi.Input

	ToNlbMapOutput() NlbMapOutput
	ToNlbMapOutputWithContext(context.Context) NlbMapOutput
}

type NlbMap map[string]NlbInput

func (NlbMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nlb)(nil)).Elem()
}

func (i NlbMap) ToNlbMapOutput() NlbMapOutput {
	return i.ToNlbMapOutputWithContext(context.Background())
}

func (i NlbMap) ToNlbMapOutputWithContext(ctx context.Context) NlbMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NlbMapOutput)
}

type NlbOutput struct{ *pulumi.OutputState }

func (NlbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nlb)(nil)).Elem()
}

func (o NlbOutput) ToNlbOutput() NlbOutput {
	return o
}

func (o NlbOutput) ToNlbOutputWithContext(ctx context.Context) NlbOutput {
	return o
}

// The NLB creation date.
func (o NlbOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Nlb) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// A free-form text describing the NLB.
func (o NlbOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nlb) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The NLB IPv4 address.
func (o NlbOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Nlb) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// A map of key/value labels.
func (o NlbOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Nlb) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The network load balancer (NLB) name.
func (o NlbOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Nlb) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of the exoscale*nlb*service (names).
func (o NlbOutput) Services() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Nlb) pulumi.StringArrayOutput { return v.Services }).(pulumi.StringArrayOutput)
}

// The current NLB state.
func (o NlbOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Nlb) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o NlbOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Nlb) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type NlbArrayOutput struct{ *pulumi.OutputState }

func (NlbArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nlb)(nil)).Elem()
}

func (o NlbArrayOutput) ToNlbArrayOutput() NlbArrayOutput {
	return o
}

func (o NlbArrayOutput) ToNlbArrayOutputWithContext(ctx context.Context) NlbArrayOutput {
	return o
}

func (o NlbArrayOutput) Index(i pulumi.IntInput) NlbOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Nlb {
		return vs[0].([]*Nlb)[vs[1].(int)]
	}).(NlbOutput)
}

type NlbMapOutput struct{ *pulumi.OutputState }

func (NlbMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nlb)(nil)).Elem()
}

func (o NlbMapOutput) ToNlbMapOutput() NlbMapOutput {
	return o
}

func (o NlbMapOutput) ToNlbMapOutputWithContext(ctx context.Context) NlbMapOutput {
	return o
}

func (o NlbMapOutput) MapIndex(k pulumi.StringInput) NlbOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Nlb {
		return vs[0].(map[string]*Nlb)[vs[1].(string)]
	}).(NlbOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NlbInput)(nil)).Elem(), &Nlb{})
	pulumi.RegisterInputType(reflect.TypeOf((*NlbArrayInput)(nil)).Elem(), NlbArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NlbMapInput)(nil)).Elem(), NlbMap{})
	pulumi.RegisterOutputType(NlbOutput{})
	pulumi.RegisterOutputType(NlbArrayOutput{})
	pulumi.RegisterOutputType(NlbMapOutput{})
}
