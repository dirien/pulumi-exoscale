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

// Manage Exoscale [Private Networks](https://community.exoscale.com/documentation/compute/private-networks/).
//
// Corresponding data source: exoscale_private_network.
//
// ## Example Usage
//
// *Unmanaged* private network:
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
//			_, err := exoscale.NewPrivateNetwork(ctx, "myPrivateNetwork", &exoscale.PrivateNetworkArgs{
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
//
// *Managed* private network:
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
//			_, err := exoscale.NewPrivateNetwork(ctx, "myManagedPrivateNetwork", &exoscale.PrivateNetworkArgs{
//				EndIp:   pulumi.String("10.0.0.253"),
//				Netmask: pulumi.String("255.255.255.0"),
//				StartIp: pulumi.String("10.0.0.20"),
//				Zone:    pulumi.String("ch-gva-2"),
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
// An existing private network may be imported by `<ID>@<zone>`
//
// ```sh
//
//	$ pulumi import exoscale:index/privateNetwork:PrivateNetwork \
//
// ```
//
//	exoscale_private_network.my_private_network \
//
//	f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
type PrivateNetwork struct {
	pulumi.CustomResourceState

	// A free-form text describing the network.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
	EndIp pulumi.StringPtrOutput `pulumi:"endIp"`
	// A map of key/value labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The private network name.
	Name pulumi.StringOutput `pulumi:"name"`
	// (For managed Privnets) The network mask defining the IPv4 network allowed for static leases.
	Netmask pulumi.StringPtrOutput `pulumi:"netmask"`
	// (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
	StartIp pulumi.StringPtrOutput `pulumi:"startIp"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewPrivateNetwork registers a new resource with the given unique name, arguments, and options.
func NewPrivateNetwork(ctx *pulumi.Context,
	name string, args *PrivateNetworkArgs, opts ...pulumi.ResourceOption) (*PrivateNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrivateNetwork
	err := ctx.RegisterResource("exoscale:index/privateNetwork:PrivateNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateNetwork gets an existing PrivateNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateNetworkState, opts ...pulumi.ResourceOption) (*PrivateNetwork, error) {
	var resource PrivateNetwork
	err := ctx.ReadResource("exoscale:index/privateNetwork:PrivateNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateNetwork resources.
type privateNetworkState struct {
	// A free-form text describing the network.
	Description *string `pulumi:"description"`
	// (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
	EndIp *string `pulumi:"endIp"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// The private network name.
	Name *string `pulumi:"name"`
	// (For managed Privnets) The network mask defining the IPv4 network allowed for static leases.
	Netmask *string `pulumi:"netmask"`
	// (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
	StartIp *string `pulumi:"startIp"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `pulumi:"zone"`
}

type PrivateNetworkState struct {
	// A free-form text describing the network.
	Description pulumi.StringPtrInput
	// (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
	EndIp pulumi.StringPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// The private network name.
	Name pulumi.StringPtrInput
	// (For managed Privnets) The network mask defining the IPv4 network allowed for static leases.
	Netmask pulumi.StringPtrInput
	// (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
	StartIp pulumi.StringPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringPtrInput
}

func (PrivateNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateNetworkState)(nil)).Elem()
}

type privateNetworkArgs struct {
	// A free-form text describing the network.
	Description *string `pulumi:"description"`
	// (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
	EndIp *string `pulumi:"endIp"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// The private network name.
	Name *string `pulumi:"name"`
	// (For managed Privnets) The network mask defining the IPv4 network allowed for static leases.
	Netmask *string `pulumi:"netmask"`
	// (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
	StartIp *string `pulumi:"startIp"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a PrivateNetwork resource.
type PrivateNetworkArgs struct {
	// A free-form text describing the network.
	Description pulumi.StringPtrInput
	// (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
	EndIp pulumi.StringPtrInput
	// A map of key/value labels.
	Labels pulumi.StringMapInput
	// The private network name.
	Name pulumi.StringPtrInput
	// (For managed Privnets) The network mask defining the IPv4 network allowed for static leases.
	Netmask pulumi.StringPtrInput
	// (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
	StartIp pulumi.StringPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput
}

func (PrivateNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateNetworkArgs)(nil)).Elem()
}

type PrivateNetworkInput interface {
	pulumi.Input

	ToPrivateNetworkOutput() PrivateNetworkOutput
	ToPrivateNetworkOutputWithContext(ctx context.Context) PrivateNetworkOutput
}

func (*PrivateNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateNetwork)(nil)).Elem()
}

func (i *PrivateNetwork) ToPrivateNetworkOutput() PrivateNetworkOutput {
	return i.ToPrivateNetworkOutputWithContext(context.Background())
}

func (i *PrivateNetwork) ToPrivateNetworkOutputWithContext(ctx context.Context) PrivateNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateNetworkOutput)
}

// PrivateNetworkArrayInput is an input type that accepts PrivateNetworkArray and PrivateNetworkArrayOutput values.
// You can construct a concrete instance of `PrivateNetworkArrayInput` via:
//
//	PrivateNetworkArray{ PrivateNetworkArgs{...} }
type PrivateNetworkArrayInput interface {
	pulumi.Input

	ToPrivateNetworkArrayOutput() PrivateNetworkArrayOutput
	ToPrivateNetworkArrayOutputWithContext(context.Context) PrivateNetworkArrayOutput
}

type PrivateNetworkArray []PrivateNetworkInput

func (PrivateNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateNetwork)(nil)).Elem()
}

func (i PrivateNetworkArray) ToPrivateNetworkArrayOutput() PrivateNetworkArrayOutput {
	return i.ToPrivateNetworkArrayOutputWithContext(context.Background())
}

func (i PrivateNetworkArray) ToPrivateNetworkArrayOutputWithContext(ctx context.Context) PrivateNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateNetworkArrayOutput)
}

// PrivateNetworkMapInput is an input type that accepts PrivateNetworkMap and PrivateNetworkMapOutput values.
// You can construct a concrete instance of `PrivateNetworkMapInput` via:
//
//	PrivateNetworkMap{ "key": PrivateNetworkArgs{...} }
type PrivateNetworkMapInput interface {
	pulumi.Input

	ToPrivateNetworkMapOutput() PrivateNetworkMapOutput
	ToPrivateNetworkMapOutputWithContext(context.Context) PrivateNetworkMapOutput
}

type PrivateNetworkMap map[string]PrivateNetworkInput

func (PrivateNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateNetwork)(nil)).Elem()
}

func (i PrivateNetworkMap) ToPrivateNetworkMapOutput() PrivateNetworkMapOutput {
	return i.ToPrivateNetworkMapOutputWithContext(context.Background())
}

func (i PrivateNetworkMap) ToPrivateNetworkMapOutputWithContext(ctx context.Context) PrivateNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateNetworkMapOutput)
}

type PrivateNetworkOutput struct{ *pulumi.OutputState }

func (PrivateNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateNetwork)(nil)).Elem()
}

func (o PrivateNetworkOutput) ToPrivateNetworkOutput() PrivateNetworkOutput {
	return o
}

func (o PrivateNetworkOutput) ToPrivateNetworkOutputWithContext(ctx context.Context) PrivateNetworkOutput {
	return o
}

// A free-form text describing the network.
func (o PrivateNetworkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
func (o PrivateNetworkOutput) EndIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringPtrOutput { return v.EndIp }).(pulumi.StringPtrOutput)
}

// A map of key/value labels.
func (o PrivateNetworkOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The private network name.
func (o PrivateNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (For managed Privnets) The network mask defining the IPv4 network allowed for static leases.
func (o PrivateNetworkOutput) Netmask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringPtrOutput { return v.Netmask }).(pulumi.StringPtrOutput)
}

// (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
func (o PrivateNetworkOutput) StartIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringPtrOutput { return v.StartIp }).(pulumi.StringPtrOutput)
}

// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o PrivateNetworkOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type PrivateNetworkArrayOutput struct{ *pulumi.OutputState }

func (PrivateNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateNetwork)(nil)).Elem()
}

func (o PrivateNetworkArrayOutput) ToPrivateNetworkArrayOutput() PrivateNetworkArrayOutput {
	return o
}

func (o PrivateNetworkArrayOutput) ToPrivateNetworkArrayOutputWithContext(ctx context.Context) PrivateNetworkArrayOutput {
	return o
}

func (o PrivateNetworkArrayOutput) Index(i pulumi.IntInput) PrivateNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivateNetwork {
		return vs[0].([]*PrivateNetwork)[vs[1].(int)]
	}).(PrivateNetworkOutput)
}

type PrivateNetworkMapOutput struct{ *pulumi.OutputState }

func (PrivateNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateNetwork)(nil)).Elem()
}

func (o PrivateNetworkMapOutput) ToPrivateNetworkMapOutput() PrivateNetworkMapOutput {
	return o
}

func (o PrivateNetworkMapOutput) ToPrivateNetworkMapOutputWithContext(ctx context.Context) PrivateNetworkMapOutput {
	return o
}

func (o PrivateNetworkMapOutput) MapIndex(k pulumi.StringInput) PrivateNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivateNetwork {
		return vs[0].(map[string]*PrivateNetwork)[vs[1].(string)]
	}).(PrivateNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateNetworkInput)(nil)).Elem(), &PrivateNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateNetworkArrayInput)(nil)).Elem(), PrivateNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateNetworkMapInput)(nil)).Elem(), PrivateNetworkMap{})
	pulumi.RegisterOutputType(PrivateNetworkOutput{})
	pulumi.RegisterOutputType(PrivateNetworkArrayOutput{})
	pulumi.RegisterOutputType(PrivateNetworkMapOutput{})
}
