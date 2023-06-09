// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system proberesponse
func LookupSystemProberesponse(ctx *pulumi.Context, args *LookupSystemProberesponseArgs, opts ...pulumi.InvokeOption) (*LookupSystemProberesponseResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemProberesponseResult
	err := ctx.Invoke("fortios:index/getSystemProberesponse:getSystemProberesponse", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemProberesponse.
type LookupSystemProberesponseArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemProberesponse.
type LookupSystemProberesponseResult struct {
	// Value to respond to the monitoring server.
	HttpProbeValue string `pulumi:"httpProbeValue"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// SLA response mode.
	Mode string `pulumi:"mode"`
	// Twamp respondor password in authentication mode
	Password string `pulumi:"password"`
	// Port number to response.
	Port int `pulumi:"port"`
	// Twamp respondor security mode.
	SecurityMode string `pulumi:"securityMode"`
	// An inactivity timer for a twamp test session.
	Timeout int `pulumi:"timeout"`
	// Mode for TWAMP packet TTL modification.
	TtlMode   string  `pulumi:"ttlMode"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemProberesponseOutput(ctx *pulumi.Context, args LookupSystemProberesponseOutputArgs, opts ...pulumi.InvokeOption) LookupSystemProberesponseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemProberesponseResult, error) {
			args := v.(LookupSystemProberesponseArgs)
			r, err := LookupSystemProberesponse(ctx, &args, opts...)
			var s LookupSystemProberesponseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemProberesponseResultOutput)
}

// A collection of arguments for invoking getSystemProberesponse.
type LookupSystemProberesponseOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemProberesponseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemProberesponseArgs)(nil)).Elem()
}

// A collection of values returned by getSystemProberesponse.
type LookupSystemProberesponseResultOutput struct{ *pulumi.OutputState }

func (LookupSystemProberesponseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemProberesponseResult)(nil)).Elem()
}

func (o LookupSystemProberesponseResultOutput) ToLookupSystemProberesponseResultOutput() LookupSystemProberesponseResultOutput {
	return o
}

func (o LookupSystemProberesponseResultOutput) ToLookupSystemProberesponseResultOutputWithContext(ctx context.Context) LookupSystemProberesponseResultOutput {
	return o
}

// Value to respond to the monitoring server.
func (o LookupSystemProberesponseResultOutput) HttpProbeValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemProberesponseResult) string { return v.HttpProbeValue }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemProberesponseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemProberesponseResult) string { return v.Id }).(pulumi.StringOutput)
}

// SLA response mode.
func (o LookupSystemProberesponseResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemProberesponseResult) string { return v.Mode }).(pulumi.StringOutput)
}

// Twamp respondor password in authentication mode
func (o LookupSystemProberesponseResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemProberesponseResult) string { return v.Password }).(pulumi.StringOutput)
}

// Port number to response.
func (o LookupSystemProberesponseResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemProberesponseResult) int { return v.Port }).(pulumi.IntOutput)
}

// Twamp respondor security mode.
func (o LookupSystemProberesponseResultOutput) SecurityMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemProberesponseResult) string { return v.SecurityMode }).(pulumi.StringOutput)
}

// An inactivity timer for a twamp test session.
func (o LookupSystemProberesponseResultOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemProberesponseResult) int { return v.Timeout }).(pulumi.IntOutput)
}

// Mode for TWAMP packet TTL modification.
func (o LookupSystemProberesponseResultOutput) TtlMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemProberesponseResult) string { return v.TtlMode }).(pulumi.StringOutput)
}

func (o LookupSystemProberesponseResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemProberesponseResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemProberesponseResultOutput{})
}
