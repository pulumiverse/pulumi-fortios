// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios firewall ipv6ehfilter
func LookupFirewallIpv6ehfilter(ctx *pulumi.Context, args *LookupFirewallIpv6ehfilterArgs, opts ...pulumi.InvokeOption) (*LookupFirewallIpv6ehfilterResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallIpv6ehfilterResult
	err := ctx.Invoke("fortios:index/getFirewallIpv6ehfilter:getFirewallIpv6ehfilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallIpv6ehfilter.
type LookupFirewallIpv6ehfilterArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallIpv6ehfilter.
type LookupFirewallIpv6ehfilterResult struct {
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

func LookupFirewallIpv6ehfilterOutput(ctx *pulumi.Context, args LookupFirewallIpv6ehfilterOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallIpv6ehfilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallIpv6ehfilterResult, error) {
			args := v.(LookupFirewallIpv6ehfilterArgs)
			r, err := LookupFirewallIpv6ehfilter(ctx, &args, opts...)
			var s LookupFirewallIpv6ehfilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallIpv6ehfilterResultOutput)
}

// A collection of arguments for invoking getFirewallIpv6ehfilter.
type LookupFirewallIpv6ehfilterOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallIpv6ehfilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallIpv6ehfilterArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallIpv6ehfilter.
type LookupFirewallIpv6ehfilterResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallIpv6ehfilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallIpv6ehfilterResult)(nil)).Elem()
}

func (o LookupFirewallIpv6ehfilterResultOutput) ToLookupFirewallIpv6ehfilterResultOutput() LookupFirewallIpv6ehfilterResultOutput {
	return o
}

func (o LookupFirewallIpv6ehfilterResultOutput) ToLookupFirewallIpv6ehfilterResultOutputWithContext(ctx context.Context) LookupFirewallIpv6ehfilterResultOutput {
	return o
}

// Enable/disable blocking packets with the Authentication header (default = disable).
func (o LookupFirewallIpv6ehfilterResultOutput) Auth() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6ehfilterResult) string { return v.Auth }).(pulumi.StringOutput)
}

// Enable/disable blocking packets with Destination Options headers (default = disable).
func (o LookupFirewallIpv6ehfilterResultOutput) DestOpt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6ehfilterResult) string { return v.DestOpt }).(pulumi.StringOutput)
}

// Enable/disable blocking packets with the Fragment header (default = disable).
func (o LookupFirewallIpv6ehfilterResultOutput) Fragment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6ehfilterResult) string { return v.Fragment }).(pulumi.StringOutput)
}

// Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
func (o LookupFirewallIpv6ehfilterResultOutput) HdoptType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallIpv6ehfilterResult) int { return v.HdoptType }).(pulumi.IntOutput)
}

// Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable).
func (o LookupFirewallIpv6ehfilterResultOutput) HopOpt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6ehfilterResult) string { return v.HopOpt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallIpv6ehfilterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6ehfilterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable blocking packets with the No Next header (default = disable)
func (o LookupFirewallIpv6ehfilterResultOutput) NoNext() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6ehfilterResult) string { return v.NoNext }).(pulumi.StringOutput)
}

// Enable/disable blocking packets with Routing headers (default = enable).
func (o LookupFirewallIpv6ehfilterResultOutput) Routing() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6ehfilterResult) string { return v.Routing }).(pulumi.StringOutput)
}

// Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
func (o LookupFirewallIpv6ehfilterResultOutput) RoutingType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallIpv6ehfilterResult) int { return v.RoutingType }).(pulumi.IntOutput)
}

func (o LookupFirewallIpv6ehfilterResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallIpv6ehfilterResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallIpv6ehfilterResultOutput{})
}