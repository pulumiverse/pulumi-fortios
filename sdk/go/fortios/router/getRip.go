// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on fortios router rip
func LookupRip(ctx *pulumi.Context, args *LookupRipArgs, opts ...pulumi.InvokeOption) (*LookupRipResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRipResult
	err := ctx.Invoke("fortios:router/getRip:getRip", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRip.
type LookupRipArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRip.
type LookupRipResult struct {
	// Enable/disable generation of default route.
	DefaultInformationOriginate string `pulumi:"defaultInformationOriginate"`
	// Default metric.
	DefaultMetric int `pulumi:"defaultMetric"`
	// Distance (1 - 255).
	Distances []GetRipDistance `pulumi:"distances"`
	// Distribute list. The structure of `distributeList` block is documented below.
	DistributeLists []GetRipDistributeList `pulumi:"distributeLists"`
	// Garbage timer in seconds.
	GarbageTimer int `pulumi:"garbageTimer"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Interface name.
	Interfaces []GetRipInterface `pulumi:"interfaces"`
	// Maximum metric allowed to output(0 means 'not set').
	MaxOutMetric int `pulumi:"maxOutMetric"`
	// neighbor The structure of `neighbor` block is documented below.
	Neighbors []GetRipNeighbor `pulumi:"neighbors"`
	// network The structure of `network` block is documented below.
	Networks []GetRipNetwork `pulumi:"networks"`
	// Offset list. The structure of `offsetList` block is documented below.
	OffsetLists []GetRipOffsetList `pulumi:"offsetLists"`
	// Passive interface configuration. The structure of `passiveInterface` block is documented below.
	PassiveInterfaces []GetRipPassiveInterface `pulumi:"passiveInterfaces"`
	// Receiving buffer size.
	RecvBufferSize int `pulumi:"recvBufferSize"`
	// Redistribute configuration. The structure of `redistribute` block is documented below.
	Redistributes []GetRipRedistribute `pulumi:"redistributes"`
	// Timeout timer in seconds.
	TimeoutTimer int `pulumi:"timeoutTimer"`
	// Update timer in seconds.
	UpdateTimer int     `pulumi:"updateTimer"`
	Vdomparam   *string `pulumi:"vdomparam"`
	// RIP version.
	Version string `pulumi:"version"`
}

func LookupRipOutput(ctx *pulumi.Context, args LookupRipOutputArgs, opts ...pulumi.InvokeOption) LookupRipResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRipResult, error) {
			args := v.(LookupRipArgs)
			r, err := LookupRip(ctx, &args, opts...)
			var s LookupRipResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRipResultOutput)
}

// A collection of arguments for invoking getRip.
type LookupRipOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRipOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRipArgs)(nil)).Elem()
}

// A collection of values returned by getRip.
type LookupRipResultOutput struct{ *pulumi.OutputState }

func (LookupRipResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRipResult)(nil)).Elem()
}

func (o LookupRipResultOutput) ToLookupRipResultOutput() LookupRipResultOutput {
	return o
}

func (o LookupRipResultOutput) ToLookupRipResultOutputWithContext(ctx context.Context) LookupRipResultOutput {
	return o
}

// Enable/disable generation of default route.
func (o LookupRipResultOutput) DefaultInformationOriginate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRipResult) string { return v.DefaultInformationOriginate }).(pulumi.StringOutput)
}

// Default metric.
func (o LookupRipResultOutput) DefaultMetric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRipResult) int { return v.DefaultMetric }).(pulumi.IntOutput)
}

// Distance (1 - 255).
func (o LookupRipResultOutput) Distances() GetRipDistanceArrayOutput {
	return o.ApplyT(func(v LookupRipResult) []GetRipDistance { return v.Distances }).(GetRipDistanceArrayOutput)
}

// Distribute list. The structure of `distributeList` block is documented below.
func (o LookupRipResultOutput) DistributeLists() GetRipDistributeListArrayOutput {
	return o.ApplyT(func(v LookupRipResult) []GetRipDistributeList { return v.DistributeLists }).(GetRipDistributeListArrayOutput)
}

// Garbage timer in seconds.
func (o LookupRipResultOutput) GarbageTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRipResult) int { return v.GarbageTimer }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRipResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRipResult) string { return v.Id }).(pulumi.StringOutput)
}

// Interface name.
func (o LookupRipResultOutput) Interfaces() GetRipInterfaceArrayOutput {
	return o.ApplyT(func(v LookupRipResult) []GetRipInterface { return v.Interfaces }).(GetRipInterfaceArrayOutput)
}

// Maximum metric allowed to output(0 means 'not set').
func (o LookupRipResultOutput) MaxOutMetric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRipResult) int { return v.MaxOutMetric }).(pulumi.IntOutput)
}

// neighbor The structure of `neighbor` block is documented below.
func (o LookupRipResultOutput) Neighbors() GetRipNeighborArrayOutput {
	return o.ApplyT(func(v LookupRipResult) []GetRipNeighbor { return v.Neighbors }).(GetRipNeighborArrayOutput)
}

// network The structure of `network` block is documented below.
func (o LookupRipResultOutput) Networks() GetRipNetworkArrayOutput {
	return o.ApplyT(func(v LookupRipResult) []GetRipNetwork { return v.Networks }).(GetRipNetworkArrayOutput)
}

// Offset list. The structure of `offsetList` block is documented below.
func (o LookupRipResultOutput) OffsetLists() GetRipOffsetListArrayOutput {
	return o.ApplyT(func(v LookupRipResult) []GetRipOffsetList { return v.OffsetLists }).(GetRipOffsetListArrayOutput)
}

// Passive interface configuration. The structure of `passiveInterface` block is documented below.
func (o LookupRipResultOutput) PassiveInterfaces() GetRipPassiveInterfaceArrayOutput {
	return o.ApplyT(func(v LookupRipResult) []GetRipPassiveInterface { return v.PassiveInterfaces }).(GetRipPassiveInterfaceArrayOutput)
}

// Receiving buffer size.
func (o LookupRipResultOutput) RecvBufferSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRipResult) int { return v.RecvBufferSize }).(pulumi.IntOutput)
}

// Redistribute configuration. The structure of `redistribute` block is documented below.
func (o LookupRipResultOutput) Redistributes() GetRipRedistributeArrayOutput {
	return o.ApplyT(func(v LookupRipResult) []GetRipRedistribute { return v.Redistributes }).(GetRipRedistributeArrayOutput)
}

// Timeout timer in seconds.
func (o LookupRipResultOutput) TimeoutTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRipResult) int { return v.TimeoutTimer }).(pulumi.IntOutput)
}

// Update timer in seconds.
func (o LookupRipResultOutput) UpdateTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRipResult) int { return v.UpdateTimer }).(pulumi.IntOutput)
}

func (o LookupRipResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRipResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// RIP version.
func (o LookupRipResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRipResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRipResultOutput{})
}
