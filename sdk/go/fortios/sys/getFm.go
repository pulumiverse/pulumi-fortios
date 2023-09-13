// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system fm
func LookupFm(ctx *pulumi.Context, args *LookupFmArgs, opts ...pulumi.InvokeOption) (*LookupFmResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFmResult
	err := ctx.Invoke("fortios:sys/getFm:getFm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFm.
type LookupFmArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFm.
type LookupFmResult struct {
	// Enable/disable automatic backup.
	AutoBackup string `pulumi:"autoBackup"`
	// ID.
	Fosid string `pulumi:"fosid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IP address.
	Ip string `pulumi:"ip"`
	// Enable/disable IPsec.
	Ipsec string `pulumi:"ipsec"`
	// Enable/disable scheduled configuration restore.
	ScheduledConfigRestore string `pulumi:"scheduledConfigRestore"`
	// Enable/disable FM.
	Status string `pulumi:"status"`
	// VDOM.
	Vdom      string  `pulumi:"vdom"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupFmOutput(ctx *pulumi.Context, args LookupFmOutputArgs, opts ...pulumi.InvokeOption) LookupFmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFmResult, error) {
			args := v.(LookupFmArgs)
			r, err := LookupFm(ctx, &args, opts...)
			var s LookupFmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFmResultOutput)
}

// A collection of arguments for invoking getFm.
type LookupFmOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFmArgs)(nil)).Elem()
}

// A collection of values returned by getFm.
type LookupFmResultOutput struct{ *pulumi.OutputState }

func (LookupFmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFmResult)(nil)).Elem()
}

func (o LookupFmResultOutput) ToLookupFmResultOutput() LookupFmResultOutput {
	return o
}

func (o LookupFmResultOutput) ToLookupFmResultOutputWithContext(ctx context.Context) LookupFmResultOutput {
	return o
}

// Enable/disable automatic backup.
func (o LookupFmResultOutput) AutoBackup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFmResult) string { return v.AutoBackup }).(pulumi.StringOutput)
}

// ID.
func (o LookupFmResultOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFmResult) string { return v.Fosid }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFmResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFmResult) string { return v.Id }).(pulumi.StringOutput)
}

// IP address.
func (o LookupFmResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFmResult) string { return v.Ip }).(pulumi.StringOutput)
}

// Enable/disable IPsec.
func (o LookupFmResultOutput) Ipsec() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFmResult) string { return v.Ipsec }).(pulumi.StringOutput)
}

// Enable/disable scheduled configuration restore.
func (o LookupFmResultOutput) ScheduledConfigRestore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFmResult) string { return v.ScheduledConfigRestore }).(pulumi.StringOutput)
}

// Enable/disable FM.
func (o LookupFmResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFmResult) string { return v.Status }).(pulumi.StringOutput)
}

// VDOM.
func (o LookupFmResultOutput) Vdom() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFmResult) string { return v.Vdom }).(pulumi.StringOutput)
}

func (o LookupFmResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFmResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFmResultOutput{})
}
