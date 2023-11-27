// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

func LookupElasticIP(ctx *pulumi.Context, args *LookupElasticIPArgs, opts ...pulumi.InvokeOption) (*LookupElasticIPResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupElasticIPResult
	err := ctx.Invoke("exoscale:index/getElasticIP:getElasticIP", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getElasticIP.
type LookupElasticIPArgs struct {
	// The Elastic IP (EIP) ID to match (conflicts with `ipAddress` and `labels`).
	Id *string `pulumi:"id"`
	// The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
	IpAddress *string `pulumi:"ipAddress"`
	// The EIP labels to match (conflicts with `ipAddress` and `id`).
	Labels map[string]string `pulumi:"labels"`
	// The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getElasticIP.
type LookupElasticIPResult struct {
	// The Elastic IP (EIP) address family (`inet4` or `inet6`).
	AddressFamily string `pulumi:"addressFamily"`
	// The Elastic IP (EIP) CIDR.
	Cidr string `pulumi:"cidr"`
	// The Elastic IP (EIP) description.
	Description string `pulumi:"description"`
	// The *managed* EIP healthcheck configuration.
	Healthchecks []GetElasticIPHealthcheck `pulumi:"healthchecks"`
	// The Elastic IP (EIP) ID to match (conflicts with `ipAddress` and `labels`).
	Id *string `pulumi:"id"`
	// The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
	IpAddress *string `pulumi:"ipAddress"`
	// The EIP labels to match (conflicts with `ipAddress` and `id`).
	Labels map[string]string `pulumi:"labels"`
	// Domain name for reverse DNS record.
	ReverseDns string `pulumi:"reverseDns"`
	// The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

func LookupElasticIPOutput(ctx *pulumi.Context, args LookupElasticIPOutputArgs, opts ...pulumi.InvokeOption) LookupElasticIPResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupElasticIPResult, error) {
			args := v.(LookupElasticIPArgs)
			r, err := LookupElasticIP(ctx, &args, opts...)
			var s LookupElasticIPResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupElasticIPResultOutput)
}

// A collection of arguments for invoking getElasticIP.
type LookupElasticIPOutputArgs struct {
	// The Elastic IP (EIP) ID to match (conflicts with `ipAddress` and `labels`).
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	// The EIP labels to match (conflicts with `ipAddress` and `id`).
	Labels pulumi.StringMapInput `pulumi:"labels"`
	// The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupElasticIPOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticIPArgs)(nil)).Elem()
}

// A collection of values returned by getElasticIP.
type LookupElasticIPResultOutput struct{ *pulumi.OutputState }

func (LookupElasticIPResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticIPResult)(nil)).Elem()
}

func (o LookupElasticIPResultOutput) ToLookupElasticIPResultOutput() LookupElasticIPResultOutput {
	return o
}

func (o LookupElasticIPResultOutput) ToLookupElasticIPResultOutputWithContext(ctx context.Context) LookupElasticIPResultOutput {
	return o
}

// The Elastic IP (EIP) address family (`inet4` or `inet6`).
func (o LookupElasticIPResultOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticIPResult) string { return v.AddressFamily }).(pulumi.StringOutput)
}

// The Elastic IP (EIP) CIDR.
func (o LookupElasticIPResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticIPResult) string { return v.Cidr }).(pulumi.StringOutput)
}

// The Elastic IP (EIP) description.
func (o LookupElasticIPResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticIPResult) string { return v.Description }).(pulumi.StringOutput)
}

// The *managed* EIP healthcheck configuration.
func (o LookupElasticIPResultOutput) Healthchecks() GetElasticIPHealthcheckArrayOutput {
	return o.ApplyT(func(v LookupElasticIPResult) []GetElasticIPHealthcheck { return v.Healthchecks }).(GetElasticIPHealthcheckArrayOutput)
}

// The Elastic IP (EIP) ID to match (conflicts with `ipAddress` and `labels`).
func (o LookupElasticIPResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupElasticIPResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
func (o LookupElasticIPResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupElasticIPResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// The EIP labels to match (conflicts with `ipAddress` and `id`).
func (o LookupElasticIPResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupElasticIPResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Domain name for reverse DNS record.
func (o LookupElasticIPResultOutput) ReverseDns() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticIPResult) string { return v.ReverseDns }).(pulumi.StringOutput)
}

// The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
func (o LookupElasticIPResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticIPResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupElasticIPResultOutput{})
}
