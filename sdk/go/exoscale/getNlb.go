// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

// Fetch Exoscale [Network Load Balancers (NLB)](https://community.exoscale.com/documentation/compute/network-load-balancer/) data.
//
// Corresponding resource: exoscale_nlb.
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
//			myNlb, err := exoscale.LookupNlb(ctx, &exoscale.LookupNlbArgs{
//				Zone: "ch-gva-2",
//				Name: pulumi.StringRef("my-nlb"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("myNlbId", myNlb.Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// Please refer to the examples
// directory for complete configuration examples.
func LookupNlb(ctx *pulumi.Context, args *LookupNlbArgs, opts ...pulumi.InvokeOption) (*LookupNlbResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNlbResult
	err := ctx.Invoke("exoscale:index/getNlb:getNlb", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNlb.
type LookupNlbArgs struct {
	// The Network Load Balancers (NLB) ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// The NLB name to match (conflicts with `id`).
	Name *string `pulumi:"name"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getNlb.
type LookupNlbResult struct {
	// The NLB creation date.
	CreatedAt string `pulumi:"createdAt"`
	// The Network Load Balancers (NLB) description.
	Description string `pulumi:"description"`
	// The Network Load Balancers (NLB) ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// The NLB public IPv4 address.
	IpAddress string `pulumi:"ipAddress"`
	// The NLB name to match (conflicts with `id`).
	Name *string `pulumi:"name"`
	// The current NLB state.
	State string `pulumi:"state"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

func LookupNlbOutput(ctx *pulumi.Context, args LookupNlbOutputArgs, opts ...pulumi.InvokeOption) LookupNlbResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNlbResult, error) {
			args := v.(LookupNlbArgs)
			r, err := LookupNlb(ctx, &args, opts...)
			var s LookupNlbResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNlbResultOutput)
}

// A collection of arguments for invoking getNlb.
type LookupNlbOutputArgs struct {
	// The Network Load Balancers (NLB) ID to match (conflicts with `name`).
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The NLB name to match (conflicts with `id`).
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupNlbOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNlbArgs)(nil)).Elem()
}

// A collection of values returned by getNlb.
type LookupNlbResultOutput struct{ *pulumi.OutputState }

func (LookupNlbResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNlbResult)(nil)).Elem()
}

func (o LookupNlbResultOutput) ToLookupNlbResultOutput() LookupNlbResultOutput {
	return o
}

func (o LookupNlbResultOutput) ToLookupNlbResultOutputWithContext(ctx context.Context) LookupNlbResultOutput {
	return o
}

// The NLB creation date.
func (o LookupNlbResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNlbResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The Network Load Balancers (NLB) description.
func (o LookupNlbResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNlbResult) string { return v.Description }).(pulumi.StringOutput)
}

// The Network Load Balancers (NLB) ID to match (conflicts with `name`).
func (o LookupNlbResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNlbResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The NLB public IPv4 address.
func (o LookupNlbResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNlbResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

// The NLB name to match (conflicts with `id`).
func (o LookupNlbResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNlbResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The current NLB state.
func (o LookupNlbResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNlbResult) string { return v.State }).(pulumi.StringOutput)
}

// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o LookupNlbResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNlbResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNlbResultOutput{})
}
