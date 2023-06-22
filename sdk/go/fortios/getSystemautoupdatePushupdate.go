// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios systemautoupdate pushupdate
func LookupSystemautoupdatePushupdate(ctx *pulumi.Context, args *LookupSystemautoupdatePushupdateArgs, opts ...pulumi.InvokeOption) (*LookupSystemautoupdatePushupdateResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemautoupdatePushupdateResult
	err := ctx.Invoke("fortios:index/getSystemautoupdatePushupdate:getSystemautoupdatePushupdate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemautoupdatePushupdate.
type LookupSystemautoupdatePushupdateArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemautoupdatePushupdate.
type LookupSystemautoupdatePushupdateResult struct {
	// Push update override server.
	Address string `pulumi:"address"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable push update override server.
	Override string `pulumi:"override"`
	// Push update override port. (Do not overlap with other service ports)
	Port int `pulumi:"port"`
	// Enable/disable push updates.
	Status    string  `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemautoupdatePushupdateOutput(ctx *pulumi.Context, args LookupSystemautoupdatePushupdateOutputArgs, opts ...pulumi.InvokeOption) LookupSystemautoupdatePushupdateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemautoupdatePushupdateResult, error) {
			args := v.(LookupSystemautoupdatePushupdateArgs)
			r, err := LookupSystemautoupdatePushupdate(ctx, &args, opts...)
			var s LookupSystemautoupdatePushupdateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemautoupdatePushupdateResultOutput)
}

// A collection of arguments for invoking getSystemautoupdatePushupdate.
type LookupSystemautoupdatePushupdateOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemautoupdatePushupdateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemautoupdatePushupdateArgs)(nil)).Elem()
}

// A collection of values returned by getSystemautoupdatePushupdate.
type LookupSystemautoupdatePushupdateResultOutput struct{ *pulumi.OutputState }

func (LookupSystemautoupdatePushupdateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemautoupdatePushupdateResult)(nil)).Elem()
}

func (o LookupSystemautoupdatePushupdateResultOutput) ToLookupSystemautoupdatePushupdateResultOutput() LookupSystemautoupdatePushupdateResultOutput {
	return o
}

func (o LookupSystemautoupdatePushupdateResultOutput) ToLookupSystemautoupdatePushupdateResultOutputWithContext(ctx context.Context) LookupSystemautoupdatePushupdateResultOutput {
	return o
}

// Push update override server.
func (o LookupSystemautoupdatePushupdateResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemautoupdatePushupdateResult) string { return v.Address }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemautoupdatePushupdateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemautoupdatePushupdateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable push update override server.
func (o LookupSystemautoupdatePushupdateResultOutput) Override() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemautoupdatePushupdateResult) string { return v.Override }).(pulumi.StringOutput)
}

// Push update override port. (Do not overlap with other service ports)
func (o LookupSystemautoupdatePushupdateResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemautoupdatePushupdateResult) int { return v.Port }).(pulumi.IntOutput)
}

// Enable/disable push updates.
func (o LookupSystemautoupdatePushupdateResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemautoupdatePushupdateResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSystemautoupdatePushupdateResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemautoupdatePushupdateResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemautoupdatePushupdateResultOutput{})
}
