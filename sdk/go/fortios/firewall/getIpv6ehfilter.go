// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on fortios firewall ipv6ehfilter
func LookupIpv6ehfilter(ctx *pulumi.Context, args *LookupIpv6ehfilterArgs, opts ...pulumi.InvokeOption) (*LookupIpv6ehfilterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpv6ehfilterResult
	err := ctx.Invoke("fortios:firewall/getIpv6ehfilter:getIpv6ehfilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpv6ehfilter.
type LookupIpv6ehfilterArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getIpv6ehfilter.
type LookupIpv6ehfilterResult struct {
	// Enable/disable blocking packets with the Authentication header (default = disable).
	Auth string `pulumi:"auth"`
	// Enable/disable blocking packets with Destination Options headers (default = disable).
	DestOpt string `pulumi:"destOpt"`
	// Enable/disable blocking packets with the Fragment header (default = disable).
	Fragment string `pulumi:"fragment"`
	// Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
	HdoptType int `pulumi:"hdoptType"`
	// Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable).
	HopOpt string `pulumi:"hopOpt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable blocking packets with the No Next header (default = disable)
	NoNext string `pulumi:"noNext"`
	// Enable/disable blocking packets with Routing headers (default = enable).
	Routing string `pulumi:"routing"`
	// Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
	RoutingType int     `pulumi:"routingType"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

func LookupIpv6ehfilterOutput(ctx *pulumi.Context, args LookupIpv6ehfilterOutputArgs, opts ...pulumi.InvokeOption) LookupIpv6ehfilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpv6ehfilterResult, error) {
			args := v.(LookupIpv6ehfilterArgs)
			r, err := LookupIpv6ehfilter(ctx, &args, opts...)
			var s LookupIpv6ehfilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpv6ehfilterResultOutput)
}

// A collection of arguments for invoking getIpv6ehfilter.
type LookupIpv6ehfilterOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupIpv6ehfilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpv6ehfilterArgs)(nil)).Elem()
}

// A collection of values returned by getIpv6ehfilter.
type LookupIpv6ehfilterResultOutput struct{ *pulumi.OutputState }

func (LookupIpv6ehfilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpv6ehfilterResult)(nil)).Elem()
}

func (o LookupIpv6ehfilterResultOutput) ToLookupIpv6ehfilterResultOutput() LookupIpv6ehfilterResultOutput {
	return o
}

func (o LookupIpv6ehfilterResultOutput) ToLookupIpv6ehfilterResultOutputWithContext(ctx context.Context) LookupIpv6ehfilterResultOutput {
	return o
}

// Enable/disable blocking packets with the Authentication header (default = disable).
func (o LookupIpv6ehfilterResultOutput) Auth() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv6ehfilterResult) string { return v.Auth }).(pulumi.StringOutput)
}

// Enable/disable blocking packets with Destination Options headers (default = disable).
func (o LookupIpv6ehfilterResultOutput) DestOpt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv6ehfilterResult) string { return v.DestOpt }).(pulumi.StringOutput)
}

// Enable/disable blocking packets with the Fragment header (default = disable).
func (o LookupIpv6ehfilterResultOutput) Fragment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv6ehfilterResult) string { return v.Fragment }).(pulumi.StringOutput)
}

// Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
func (o LookupIpv6ehfilterResultOutput) HdoptType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv6ehfilterResult) int { return v.HdoptType }).(pulumi.IntOutput)
}

// Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable).
func (o LookupIpv6ehfilterResultOutput) HopOpt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv6ehfilterResult) string { return v.HopOpt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIpv6ehfilterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv6ehfilterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable blocking packets with the No Next header (default = disable)
func (o LookupIpv6ehfilterResultOutput) NoNext() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv6ehfilterResult) string { return v.NoNext }).(pulumi.StringOutput)
}

// Enable/disable blocking packets with Routing headers (default = enable).
func (o LookupIpv6ehfilterResultOutput) Routing() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv6ehfilterResult) string { return v.Routing }).(pulumi.StringOutput)
}

// Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
func (o LookupIpv6ehfilterResultOutput) RoutingType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv6ehfilterResult) int { return v.RoutingType }).(pulumi.IntOutput)
}

func (o LookupIpv6ehfilterResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpv6ehfilterResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpv6ehfilterResultOutput{})
}
