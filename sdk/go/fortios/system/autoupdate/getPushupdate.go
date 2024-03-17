// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoupdate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on fortios systemautoupdate pushupdate
func LookupPushupdate(ctx *pulumi.Context, args *LookupPushupdateArgs, opts ...pulumi.InvokeOption) (*LookupPushupdateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPushupdateResult
	err := ctx.Invoke("fortios:system/autoupdate/getPushupdate:getPushupdate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPushupdate.
type LookupPushupdateArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPushupdate.
type LookupPushupdateResult struct {
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

func LookupPushupdateOutput(ctx *pulumi.Context, args LookupPushupdateOutputArgs, opts ...pulumi.InvokeOption) LookupPushupdateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPushupdateResult, error) {
			args := v.(LookupPushupdateArgs)
			r, err := LookupPushupdate(ctx, &args, opts...)
			var s LookupPushupdateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPushupdateResultOutput)
}

// A collection of arguments for invoking getPushupdate.
type LookupPushupdateOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupPushupdateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPushupdateArgs)(nil)).Elem()
}

// A collection of values returned by getPushupdate.
type LookupPushupdateResultOutput struct{ *pulumi.OutputState }

func (LookupPushupdateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPushupdateResult)(nil)).Elem()
}

func (o LookupPushupdateResultOutput) ToLookupPushupdateResultOutput() LookupPushupdateResultOutput {
	return o
}

func (o LookupPushupdateResultOutput) ToLookupPushupdateResultOutputWithContext(ctx context.Context) LookupPushupdateResultOutput {
	return o
}

// Push update override server.
func (o LookupPushupdateResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPushupdateResult) string { return v.Address }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPushupdateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPushupdateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable push update override server.
func (o LookupPushupdateResultOutput) Override() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPushupdateResult) string { return v.Override }).(pulumi.StringOutput)
}

// Push update override port. (Do not overlap with other service ports)
func (o LookupPushupdateResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPushupdateResult) int { return v.Port }).(pulumi.IntOutput)
}

// Enable/disable push updates.
func (o LookupPushupdateResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPushupdateResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupPushupdateResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPushupdateResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPushupdateResultOutput{})
}
