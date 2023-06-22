// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system fortimanager
func LookupSystemFortimanager(ctx *pulumi.Context, args *LookupSystemFortimanagerArgs, opts ...pulumi.InvokeOption) (*LookupSystemFortimanagerResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemFortimanagerResult
	err := ctx.Invoke("fortios:index/getSystemFortimanager:getSystemFortimanager", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemFortimanager.
type LookupSystemFortimanagerArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemFortimanager.
type LookupSystemFortimanagerResult struct {
	// Enable/disable FortiManager central management.
	CentralManagement string `pulumi:"centralManagement"`
	// Enable/disable central management auto backup.
	CentralMgmtAutoBackup string `pulumi:"centralMgmtAutoBackup"`
	// Enable/disable central management schedule config restore.
	CentralMgmtScheduleConfigRestore string `pulumi:"centralMgmtScheduleConfigRestore"`
	// Enable/disable central management schedule script restore.
	CentralMgmtScheduleScriptRestore string `pulumi:"centralMgmtScheduleScriptRestore"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IP address.
	Ip string `pulumi:"ip"`
	// Enable/disable FortiManager IPsec tunnel.
	Ipsec string `pulumi:"ipsec"`
	// Virtual domain name.
	Vdom      string  `pulumi:"vdom"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemFortimanagerOutput(ctx *pulumi.Context, args LookupSystemFortimanagerOutputArgs, opts ...pulumi.InvokeOption) LookupSystemFortimanagerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemFortimanagerResult, error) {
			args := v.(LookupSystemFortimanagerArgs)
			r, err := LookupSystemFortimanager(ctx, &args, opts...)
			var s LookupSystemFortimanagerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemFortimanagerResultOutput)
}

// A collection of arguments for invoking getSystemFortimanager.
type LookupSystemFortimanagerOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemFortimanagerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemFortimanagerArgs)(nil)).Elem()
}

// A collection of values returned by getSystemFortimanager.
type LookupSystemFortimanagerResultOutput struct{ *pulumi.OutputState }

func (LookupSystemFortimanagerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemFortimanagerResult)(nil)).Elem()
}

func (o LookupSystemFortimanagerResultOutput) ToLookupSystemFortimanagerResultOutput() LookupSystemFortimanagerResultOutput {
	return o
}

func (o LookupSystemFortimanagerResultOutput) ToLookupSystemFortimanagerResultOutputWithContext(ctx context.Context) LookupSystemFortimanagerResultOutput {
	return o
}

// Enable/disable FortiManager central management.
func (o LookupSystemFortimanagerResultOutput) CentralManagement() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortimanagerResult) string { return v.CentralManagement }).(pulumi.StringOutput)
}

// Enable/disable central management auto backup.
func (o LookupSystemFortimanagerResultOutput) CentralMgmtAutoBackup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortimanagerResult) string { return v.CentralMgmtAutoBackup }).(pulumi.StringOutput)
}

// Enable/disable central management schedule config restore.
func (o LookupSystemFortimanagerResultOutput) CentralMgmtScheduleConfigRestore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortimanagerResult) string { return v.CentralMgmtScheduleConfigRestore }).(pulumi.StringOutput)
}

// Enable/disable central management schedule script restore.
func (o LookupSystemFortimanagerResultOutput) CentralMgmtScheduleScriptRestore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortimanagerResult) string { return v.CentralMgmtScheduleScriptRestore }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemFortimanagerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortimanagerResult) string { return v.Id }).(pulumi.StringOutput)
}

// IP address.
func (o LookupSystemFortimanagerResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortimanagerResult) string { return v.Ip }).(pulumi.StringOutput)
}

// Enable/disable FortiManager IPsec tunnel.
func (o LookupSystemFortimanagerResultOutput) Ipsec() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortimanagerResult) string { return v.Ipsec }).(pulumi.StringOutput)
}

// Virtual domain name.
func (o LookupSystemFortimanagerResultOutput) Vdom() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortimanagerResult) string { return v.Vdom }).(pulumi.StringOutput)
}

func (o LookupSystemFortimanagerResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemFortimanagerResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemFortimanagerResultOutput{})
}
