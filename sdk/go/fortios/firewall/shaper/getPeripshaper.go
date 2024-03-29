// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package shaper

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios firewallshaper peripshaper
func LookupPeripshaper(ctx *pulumi.Context, args *LookupPeripshaperArgs, opts ...pulumi.InvokeOption) (*LookupPeripshaperResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPeripshaperResult
	err := ctx.Invoke("fortios:firewall/shaper/getPeripshaper:getPeripshaper", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPeripshaper.
type LookupPeripshaperArgs struct {
	// Specify the name of the desired firewallshaper peripshaper.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPeripshaper.
type LookupPeripshaperResult struct {
	// Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps).
	BandwidthUnit string `pulumi:"bandwidthUnit"`
	// Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper.
	DiffservForward string `pulumi:"diffservForward"`
	// Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper.
	DiffservReverse string `pulumi:"diffservReverse"`
	// Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
	DiffservcodeForward string `pulumi:"diffservcodeForward"`
	// Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
	DiffservcodeRev string `pulumi:"diffservcodeRev"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
	MaxBandwidth int `pulumi:"maxBandwidth"`
	// Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentSession int `pulumi:"maxConcurrentSession"`
	// Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentTcpSession int `pulumi:"maxConcurrentTcpSession"`
	// Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentUdpSession int `pulumi:"maxConcurrentUdpSession"`
	// Traffic shaper name.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupPeripshaperOutput(ctx *pulumi.Context, args LookupPeripshaperOutputArgs, opts ...pulumi.InvokeOption) LookupPeripshaperResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPeripshaperResult, error) {
			args := v.(LookupPeripshaperArgs)
			r, err := LookupPeripshaper(ctx, &args, opts...)
			var s LookupPeripshaperResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPeripshaperResultOutput)
}

// A collection of arguments for invoking getPeripshaper.
type LookupPeripshaperOutputArgs struct {
	// Specify the name of the desired firewallshaper peripshaper.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupPeripshaperOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPeripshaperArgs)(nil)).Elem()
}

// A collection of values returned by getPeripshaper.
type LookupPeripshaperResultOutput struct{ *pulumi.OutputState }

func (LookupPeripshaperResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPeripshaperResult)(nil)).Elem()
}

func (o LookupPeripshaperResultOutput) ToLookupPeripshaperResultOutput() LookupPeripshaperResultOutput {
	return o
}

func (o LookupPeripshaperResultOutput) ToLookupPeripshaperResultOutputWithContext(ctx context.Context) LookupPeripshaperResultOutput {
	return o
}

// Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps).
func (o LookupPeripshaperResultOutput) BandwidthUnit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeripshaperResult) string { return v.BandwidthUnit }).(pulumi.StringOutput)
}

// Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper.
func (o LookupPeripshaperResultOutput) DiffservForward() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeripshaperResult) string { return v.DiffservForward }).(pulumi.StringOutput)
}

// Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper.
func (o LookupPeripshaperResultOutput) DiffservReverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeripshaperResult) string { return v.DiffservReverse }).(pulumi.StringOutput)
}

// Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
func (o LookupPeripshaperResultOutput) DiffservcodeForward() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeripshaperResult) string { return v.DiffservcodeForward }).(pulumi.StringOutput)
}

// Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
func (o LookupPeripshaperResultOutput) DiffservcodeRev() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeripshaperResult) string { return v.DiffservcodeRev }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPeripshaperResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeripshaperResult) string { return v.Id }).(pulumi.StringOutput)
}

// Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
func (o LookupPeripshaperResultOutput) MaxBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPeripshaperResult) int { return v.MaxBandwidth }).(pulumi.IntOutput)
}

// Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
func (o LookupPeripshaperResultOutput) MaxConcurrentSession() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPeripshaperResult) int { return v.MaxConcurrentSession }).(pulumi.IntOutput)
}

// Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
func (o LookupPeripshaperResultOutput) MaxConcurrentTcpSession() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPeripshaperResult) int { return v.MaxConcurrentTcpSession }).(pulumi.IntOutput)
}

// Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
func (o LookupPeripshaperResultOutput) MaxConcurrentUdpSession() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPeripshaperResult) int { return v.MaxConcurrentUdpSession }).(pulumi.IntOutput)
}

// Traffic shaper name.
func (o LookupPeripshaperResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeripshaperResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPeripshaperResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPeripshaperResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPeripshaperResultOutput{})
}
