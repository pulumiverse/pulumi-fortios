// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios router static
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/router"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			sample1, err := router.LookupStatic(ctx, &router.LookupStaticArgs{
//				SeqNum: 1,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("output1", sample1)
//			return nil
//		})
//	}
//
// ```
func LookupStatic(ctx *pulumi.Context, args *LookupStaticArgs, opts ...pulumi.InvokeOption) (*LookupStaticResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStaticResult
	err := ctx.Invoke("fortios:router/getStatic:getStatic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStatic.
type LookupStaticArgs struct {
	// Specify the seqNum of the desired router static.
	SeqNum int `pulumi:"seqNum"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getStatic.
type LookupStaticResult struct {
	// Enable/disable Bidirectional Forwarding Detection (BFD).
	Bfd string `pulumi:"bfd"`
	// Enable/disable black hole.
	Blackhole string `pulumi:"blackhole"`
	// Optional comments.
	Comment string `pulumi:"comment"`
	// Gateway out interface or tunnel.
	Device string `pulumi:"device"`
	// Administrative distance (1 - 255).
	Distance int `pulumi:"distance"`
	// Destination IP and mask for this route.
	Dst string `pulumi:"dst"`
	// Name of firewall address or address group.
	Dstaddr string `pulumi:"dstaddr"`
	// Enable use of dynamic gateway retrieved from a DHCP or PPP server.
	DynamicGateway string `pulumi:"dynamicGateway"`
	// Gateway IP for this route.
	Gateway string `pulumi:"gateway"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Application ID in the Internet service database.
	InternetService int `pulumi:"internetService"`
	// Application name in the Internet service custom database.
	InternetServiceCustom string `pulumi:"internetServiceCustom"`
	// Enable/disable withdrawing this route when link monitor or health check is down.
	LinkMonitorExempt string `pulumi:"linkMonitorExempt"`
	// Preferred source IP for this route.
	PreferredSource string `pulumi:"preferredSource"`
	// Administrative priority (0 - 4294967295).
	Priority int `pulumi:"priority"`
	// Enable/disable egress through SD-WAN.
	Sdwan string `pulumi:"sdwan"`
	// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
	SdwanZones []GetStaticSdwanZone `pulumi:"sdwanZones"`
	// Sequence number.
	SeqNum int `pulumi:"seqNum"`
	// Source prefix for this route.
	Src string `pulumi:"src"`
	// Enable/disable this static route.
	Status string `pulumi:"status"`
	// Route tag.
	Tag       int     `pulumi:"tag"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable egress through the virtual-wan-link.
	VirtualWanLink string `pulumi:"virtualWanLink"`
	// Virtual Routing Forwarding ID.
	Vrf int `pulumi:"vrf"`
	// Administrative weight (0 - 255).
	Weight int `pulumi:"weight"`
}

func LookupStaticOutput(ctx *pulumi.Context, args LookupStaticOutputArgs, opts ...pulumi.InvokeOption) LookupStaticResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticResult, error) {
			args := v.(LookupStaticArgs)
			r, err := LookupStatic(ctx, &args, opts...)
			var s LookupStaticResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticResultOutput)
}

// A collection of arguments for invoking getStatic.
type LookupStaticOutputArgs struct {
	// Specify the seqNum of the desired router static.
	SeqNum pulumi.IntInput `pulumi:"seqNum"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupStaticOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticArgs)(nil)).Elem()
}

// A collection of values returned by getStatic.
type LookupStaticResultOutput struct{ *pulumi.OutputState }

func (LookupStaticResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticResult)(nil)).Elem()
}

func (o LookupStaticResultOutput) ToLookupStaticResultOutput() LookupStaticResultOutput {
	return o
}

func (o LookupStaticResultOutput) ToLookupStaticResultOutputWithContext(ctx context.Context) LookupStaticResultOutput {
	return o
}

// Enable/disable Bidirectional Forwarding Detection (BFD).
func (o LookupStaticResultOutput) Bfd() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.Bfd }).(pulumi.StringOutput)
}

// Enable/disable black hole.
func (o LookupStaticResultOutput) Blackhole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.Blackhole }).(pulumi.StringOutput)
}

// Optional comments.
func (o LookupStaticResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.Comment }).(pulumi.StringOutput)
}

// Gateway out interface or tunnel.
func (o LookupStaticResultOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.Device }).(pulumi.StringOutput)
}

// Administrative distance (1 - 255).
func (o LookupStaticResultOutput) Distance() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStaticResult) int { return v.Distance }).(pulumi.IntOutput)
}

// Destination IP and mask for this route.
func (o LookupStaticResultOutput) Dst() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.Dst }).(pulumi.StringOutput)
}

// Name of firewall address or address group.
func (o LookupStaticResultOutput) Dstaddr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.Dstaddr }).(pulumi.StringOutput)
}

// Enable use of dynamic gateway retrieved from a DHCP or PPP server.
func (o LookupStaticResultOutput) DynamicGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.DynamicGateway }).(pulumi.StringOutput)
}

// Gateway IP for this route.
func (o LookupStaticResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.Gateway }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupStaticResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.Id }).(pulumi.StringOutput)
}

// Application ID in the Internet service database.
func (o LookupStaticResultOutput) InternetService() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStaticResult) int { return v.InternetService }).(pulumi.IntOutput)
}

// Application name in the Internet service custom database.
func (o LookupStaticResultOutput) InternetServiceCustom() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.InternetServiceCustom }).(pulumi.StringOutput)
}

// Enable/disable withdrawing this route when link monitor or health check is down.
func (o LookupStaticResultOutput) LinkMonitorExempt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.LinkMonitorExempt }).(pulumi.StringOutput)
}

// Preferred source IP for this route.
func (o LookupStaticResultOutput) PreferredSource() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.PreferredSource }).(pulumi.StringOutput)
}

// Administrative priority (0 - 4294967295).
func (o LookupStaticResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStaticResult) int { return v.Priority }).(pulumi.IntOutput)
}

// Enable/disable egress through SD-WAN.
func (o LookupStaticResultOutput) Sdwan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.Sdwan }).(pulumi.StringOutput)
}

// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
func (o LookupStaticResultOutput) SdwanZones() GetStaticSdwanZoneArrayOutput {
	return o.ApplyT(func(v LookupStaticResult) []GetStaticSdwanZone { return v.SdwanZones }).(GetStaticSdwanZoneArrayOutput)
}

// Sequence number.
func (o LookupStaticResultOutput) SeqNum() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStaticResult) int { return v.SeqNum }).(pulumi.IntOutput)
}

// Source prefix for this route.
func (o LookupStaticResultOutput) Src() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.Src }).(pulumi.StringOutput)
}

// Enable/disable this static route.
func (o LookupStaticResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.Status }).(pulumi.StringOutput)
}

// Route tag.
func (o LookupStaticResultOutput) Tag() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStaticResult) int { return v.Tag }).(pulumi.IntOutput)
}

func (o LookupStaticResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable egress through the virtual-wan-link.
func (o LookupStaticResultOutput) VirtualWanLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticResult) string { return v.VirtualWanLink }).(pulumi.StringOutput)
}

// Virtual Routing Forwarding ID.
func (o LookupStaticResultOutput) Vrf() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStaticResult) int { return v.Vrf }).(pulumi.IntOutput)
}

// Administrative weight (0 - 255).
func (o LookupStaticResultOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStaticResult) int { return v.Weight }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticResultOutput{})
}
