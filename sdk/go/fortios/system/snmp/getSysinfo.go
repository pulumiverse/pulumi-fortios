// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package snmp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on fortios systemsnmp sysinfo
func LookupSysinfo(ctx *pulumi.Context, args *LookupSysinfoArgs, opts ...pulumi.InvokeOption) (*LookupSysinfoResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSysinfoResult
	err := ctx.Invoke("fortios:system/snmp/getSysinfo:getSysinfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSysinfo.
type LookupSysinfoArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSysinfo.
type LookupSysinfoResult struct {
	// Contact information.
	ContactInfo string `pulumi:"contactInfo"`
	// System description.
	Description string `pulumi:"description"`
	// Local SNMP engineID string (maximum 24 characters).
	EngineId string `pulumi:"engineId"`
	// Local SNMP engineID type (text/hex/mac).
	EngineIdType string `pulumi:"engineIdType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// System location.
	Location string `pulumi:"location"`
	// Enable/disable SNMP.
	Status string `pulumi:"status"`
	// CPU usage when trap is sent.
	TrapHighCpuThreshold int `pulumi:"trapHighCpuThreshold"`
	// Log disk usage when trap is sent.
	TrapLogFullThreshold int `pulumi:"trapLogFullThreshold"`
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold int     `pulumi:"trapLowMemoryThreshold"`
	Vdomparam              *string `pulumi:"vdomparam"`
}

func LookupSysinfoOutput(ctx *pulumi.Context, args LookupSysinfoOutputArgs, opts ...pulumi.InvokeOption) LookupSysinfoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSysinfoResult, error) {
			args := v.(LookupSysinfoArgs)
			r, err := LookupSysinfo(ctx, &args, opts...)
			var s LookupSysinfoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSysinfoResultOutput)
}

// A collection of arguments for invoking getSysinfo.
type LookupSysinfoOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSysinfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSysinfoArgs)(nil)).Elem()
}

// A collection of values returned by getSysinfo.
type LookupSysinfoResultOutput struct{ *pulumi.OutputState }

func (LookupSysinfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSysinfoResult)(nil)).Elem()
}

func (o LookupSysinfoResultOutput) ToLookupSysinfoResultOutput() LookupSysinfoResultOutput {
	return o
}

func (o LookupSysinfoResultOutput) ToLookupSysinfoResultOutputWithContext(ctx context.Context) LookupSysinfoResultOutput {
	return o
}

// Contact information.
func (o LookupSysinfoResultOutput) ContactInfo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSysinfoResult) string { return v.ContactInfo }).(pulumi.StringOutput)
}

// System description.
func (o LookupSysinfoResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSysinfoResult) string { return v.Description }).(pulumi.StringOutput)
}

// Local SNMP engineID string (maximum 24 characters).
func (o LookupSysinfoResultOutput) EngineId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSysinfoResult) string { return v.EngineId }).(pulumi.StringOutput)
}

// Local SNMP engineID type (text/hex/mac).
func (o LookupSysinfoResultOutput) EngineIdType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSysinfoResult) string { return v.EngineIdType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSysinfoResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSysinfoResult) string { return v.Id }).(pulumi.StringOutput)
}

// System location.
func (o LookupSysinfoResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSysinfoResult) string { return v.Location }).(pulumi.StringOutput)
}

// Enable/disable SNMP.
func (o LookupSysinfoResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSysinfoResult) string { return v.Status }).(pulumi.StringOutput)
}

// CPU usage when trap is sent.
func (o LookupSysinfoResultOutput) TrapHighCpuThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSysinfoResult) int { return v.TrapHighCpuThreshold }).(pulumi.IntOutput)
}

// Log disk usage when trap is sent.
func (o LookupSysinfoResultOutput) TrapLogFullThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSysinfoResult) int { return v.TrapLogFullThreshold }).(pulumi.IntOutput)
}

// Memory usage when trap is sent.
func (o LookupSysinfoResultOutput) TrapLowMemoryThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSysinfoResult) int { return v.TrapLowMemoryThreshold }).(pulumi.IntOutput)
}

func (o LookupSysinfoResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSysinfoResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSysinfoResultOutput{})
}
