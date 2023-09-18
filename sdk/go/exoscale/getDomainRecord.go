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

func LookupDomainRecord(ctx *pulumi.Context, args *LookupDomainRecordArgs, opts ...pulumi.InvokeOption) (*LookupDomainRecordResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainRecordResult
	err := ctx.Invoke("exoscale:index/getDomainRecord:getDomainRecord", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomainRecord.
type LookupDomainRecordArgs struct {
	// The Domain name to match.
	Domain string `pulumi:"domain"`
	// Filter to apply when looking up domain records.
	Filter GetDomainRecordFilter `pulumi:"filter"`
}

// A collection of values returned by getDomainRecord.
type LookupDomainRecordResult struct {
	// The Domain name to match.
	Domain string `pulumi:"domain"`
	// Filter to apply when looking up domain records.
	Filter GetDomainRecordFilter `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of matching records. Structure is documented below.
	Records []GetDomainRecordRecord `pulumi:"records"`
}

func LookupDomainRecordOutput(ctx *pulumi.Context, args LookupDomainRecordOutputArgs, opts ...pulumi.InvokeOption) LookupDomainRecordResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainRecordResult, error) {
			args := v.(LookupDomainRecordArgs)
			r, err := LookupDomainRecord(ctx, &args, opts...)
			var s LookupDomainRecordResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainRecordResultOutput)
}

// A collection of arguments for invoking getDomainRecord.
type LookupDomainRecordOutputArgs struct {
	// The Domain name to match.
	Domain pulumi.StringInput `pulumi:"domain"`
	// Filter to apply when looking up domain records.
	Filter GetDomainRecordFilterInput `pulumi:"filter"`
}

func (LookupDomainRecordOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainRecordArgs)(nil)).Elem()
}

// A collection of values returned by getDomainRecord.
type LookupDomainRecordResultOutput struct{ *pulumi.OutputState }

func (LookupDomainRecordResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainRecordResult)(nil)).Elem()
}

func (o LookupDomainRecordResultOutput) ToLookupDomainRecordResultOutput() LookupDomainRecordResultOutput {
	return o
}

func (o LookupDomainRecordResultOutput) ToLookupDomainRecordResultOutputWithContext(ctx context.Context) LookupDomainRecordResultOutput {
	return o
}

func (o LookupDomainRecordResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDomainRecordResult] {
	return pulumix.Output[LookupDomainRecordResult]{
		OutputState: o.OutputState,
	}
}

// The Domain name to match.
func (o LookupDomainRecordResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) string { return v.Domain }).(pulumi.StringOutput)
}

// Filter to apply when looking up domain records.
func (o LookupDomainRecordResultOutput) Filter() GetDomainRecordFilterOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) GetDomainRecordFilter { return v.Filter }).(GetDomainRecordFilterOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDomainRecordResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of matching records. Structure is documented below.
func (o LookupDomainRecordResultOutput) Records() GetDomainRecordRecordArrayOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) []GetDomainRecordRecord { return v.Records }).(GetDomainRecordRecordArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainRecordResultOutput{})
}
