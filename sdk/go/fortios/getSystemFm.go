// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system fm
func LookupSystemFm(ctx *pulumi.Context, args *LookupSystemFmArgs, opts ...pulumi.InvokeOption) (*LookupSystemFmResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemFmResult
	err := ctx.Invoke("fortios:index/getSystemFm:getSystemFm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemFm.
type LookupSystemFmArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemFm.
type LookupSystemFmResult struct {
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

func LookupSystemFmOutput(ctx *pulumi.Context, args LookupSystemFmOutputArgs, opts ...pulumi.InvokeOption) LookupSystemFmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemFmResult, error) {
			args := v.(LookupSystemFmArgs)
			r, err := LookupSystemFm(ctx, &args, opts...)
			var s LookupSystemFmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemFmResultOutput)
}

// A collection of arguments for invoking getSystemFm.
type LookupSystemFmOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemFmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemFmArgs)(nil)).Elem()
}

// A collection of values returned by getSystemFm.
type LookupSystemFmResultOutput struct{ *pulumi.OutputState }

func (LookupSystemFmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemFmResult)(nil)).Elem()
}

func (o LookupSystemFmResultOutput) ToLookupSystemFmResultOutput() LookupSystemFmResultOutput {
	return o
}

func (o LookupSystemFmResultOutput) ToLookupSystemFmResultOutputWithContext(ctx context.Context) LookupSystemFmResultOutput {
	return o
}

// Enable/disable automatic backup.
func (o LookupSystemFmResultOutput) AutoBackup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFmResult) string { return v.AutoBackup }).(pulumi.StringOutput)
}

// ID.
func (o LookupSystemFmResultOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFmResult) string { return v.Fosid }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemFmResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFmResult) string { return v.Id }).(pulumi.StringOutput)
}

// IP address.
func (o LookupSystemFmResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFmResult) string { return v.Ip }).(pulumi.StringOutput)
}

// Enable/disable IPsec.
func (o LookupSystemFmResultOutput) Ipsec() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFmResult) string { return v.Ipsec }).(pulumi.StringOutput)
}

// Enable/disable scheduled configuration restore.
func (o LookupSystemFmResultOutput) ScheduledConfigRestore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFmResult) string { return v.ScheduledConfigRestore }).(pulumi.StringOutput)
}

// Enable/disable FM.
func (o LookupSystemFmResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFmResult) string { return v.Status }).(pulumi.StringOutput)
}

// VDOM.
func (o LookupSystemFmResultOutput) Vdom() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFmResult) string { return v.Vdom }).(pulumi.StringOutput)
}

func (o LookupSystemFmResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemFmResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemFmResultOutput{})
}
