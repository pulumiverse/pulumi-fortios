// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on fortios router isis
func LookupIsis(ctx *pulumi.Context, args *LookupIsisArgs, opts ...pulumi.InvokeOption) (*LookupIsisResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIsisResult
	err := ctx.Invoke("fortios:router/getIsis:getIsis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIsis.
type LookupIsisArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getIsis.
type LookupIsisResult struct {
	// Enable/disable adjacency check.
	AdjacencyCheck string `pulumi:"adjacencyCheck"`
	// Enable/disable IPv6 adjacency check.
	AdjacencyCheck6 string `pulumi:"adjacencyCheck6"`
	// Enable/disable IS-IS advertisement of passive interfaces only.
	AdvPassiveOnly string `pulumi:"advPassiveOnly"`
	// Enable/disable IPv6 IS-IS advertisement of passive interfaces only.
	AdvPassiveOnly6 string `pulumi:"advPassiveOnly6"`
	// Authentication key-chain for level 1 PDUs.
	AuthKeychainL1 string `pulumi:"authKeychainL1"`
	// Authentication key-chain for level 2 PDUs.
	AuthKeychainL2 string `pulumi:"authKeychainL2"`
	// Level 1 authentication mode.
	AuthModeL1 string `pulumi:"authModeL1"`
	// Level 2 authentication mode.
	AuthModeL2 string `pulumi:"authModeL2"`
	// Authentication password for level 1 PDUs.
	AuthPasswordL1 string `pulumi:"authPasswordL1"`
	// Authentication password for level 2 PDUs.
	AuthPasswordL2 string `pulumi:"authPasswordL2"`
	// Enable/disable level 1 authentication send-only.
	AuthSendonlyL1 string `pulumi:"authSendonlyL1"`
	// Enable/disable level 2 authentication send-only.
	AuthSendonlyL2 string `pulumi:"authSendonlyL2"`
	// Enable/disable distribution of default route information.
	DefaultOriginate string `pulumi:"defaultOriginate"`
	// Enable/disable distribution of default IPv6 route information.
	DefaultOriginate6 string `pulumi:"defaultOriginate6"`
	// Enable/disable dynamic hostname.
	DynamicHostname string `pulumi:"dynamicHostname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable ignoring of LSP errors with bad checksums.
	IgnoreLspErrors string `pulumi:"ignoreLspErrors"`
	// IS type.
	IsType string `pulumi:"isType"`
	// IS-IS interface configuration. The structure of `isisInterface` block is documented below.
	IsisInterfaces []GetIsisIsisInterface `pulumi:"isisInterfaces"`
	// IS-IS net configuration. The structure of `isisNet` block is documented below.
	IsisNets []GetIsisIsisNet `pulumi:"isisNets"`
	// Minimum interval for level 1 LSP regenerating.
	LspGenIntervalL1 int `pulumi:"lspGenIntervalL1"`
	// Minimum interval for level 2 LSP regenerating.
	LspGenIntervalL2 int `pulumi:"lspGenIntervalL2"`
	// LSP refresh time in seconds.
	LspRefreshInterval int `pulumi:"lspRefreshInterval"`
	// Maximum LSP lifetime in seconds.
	MaxLspLifetime int `pulumi:"maxLspLifetime"`
	// Use old-style (ISO 10589) or new-style packet formats
	MetricStyle string `pulumi:"metricStyle"`
	// Enable/disable signal other routers not to use us in SPF.
	OverloadBit string `pulumi:"overloadBit"`
	// Overload-bit only temporarily after reboot.
	OverloadBitOnStartup int `pulumi:"overloadBitOnStartup"`
	// Suppress overload-bit for the specific prefixes.
	OverloadBitSuppress string `pulumi:"overloadBitSuppress"`
	// Enable/disable redistribution of level 1 IPv6 routes into level 2.
	Redistribute6L1 string `pulumi:"redistribute6L1"`
	// Access-list for IPv6 route redistribution from l1 to l2.
	Redistribute6L1List string `pulumi:"redistribute6L1List"`
	// Enable/disable redistribution of level 2 IPv6 routes into level 1.
	Redistribute6L2 string `pulumi:"redistribute6L2"`
	// Access-list for IPv6 route redistribution from l2 to l1.
	Redistribute6L2List string `pulumi:"redistribute6L2List"`
	// IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.
	Redistribute6s []GetIsisRedistribute6 `pulumi:"redistribute6s"`
	// Enable/disable redistribution of level 1 routes into level 2.
	RedistributeL1 string `pulumi:"redistributeL1"`
	// Access-list for route redistribution from l1 to l2.
	RedistributeL1List string `pulumi:"redistributeL1List"`
	// Enable/disable redistribution of level 2 routes into level 1.
	RedistributeL2 string `pulumi:"redistributeL2"`
	// Access-list for route redistribution from l2 to l1.
	RedistributeL2List string `pulumi:"redistributeL2List"`
	// IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
	Redistributes []GetIsisRedistribute `pulumi:"redistributes"`
	// Level 1 SPF calculation delay.
	SpfIntervalExpL1 string `pulumi:"spfIntervalExpL1"`
	// Level 2 SPF calculation delay.
	SpfIntervalExpL2 string `pulumi:"spfIntervalExpL2"`
	// IS-IS IPv6 summary address. The structure of `summaryAddress6` block is documented below.
	SummaryAddress6s []GetIsisSummaryAddress6 `pulumi:"summaryAddress6s"`
	// IS-IS summary addresses. The structure of `summaryAddress` block is documented below.
	SummaryAddresses []GetIsisSummaryAddress `pulumi:"summaryAddresses"`
	Vdomparam        *string                 `pulumi:"vdomparam"`
}

func LookupIsisOutput(ctx *pulumi.Context, args LookupIsisOutputArgs, opts ...pulumi.InvokeOption) LookupIsisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIsisResult, error) {
			args := v.(LookupIsisArgs)
			r, err := LookupIsis(ctx, &args, opts...)
			var s LookupIsisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIsisResultOutput)
}

// A collection of arguments for invoking getIsis.
type LookupIsisOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupIsisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIsisArgs)(nil)).Elem()
}

// A collection of values returned by getIsis.
type LookupIsisResultOutput struct{ *pulumi.OutputState }

func (LookupIsisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIsisResult)(nil)).Elem()
}

func (o LookupIsisResultOutput) ToLookupIsisResultOutput() LookupIsisResultOutput {
	return o
}

func (o LookupIsisResultOutput) ToLookupIsisResultOutputWithContext(ctx context.Context) LookupIsisResultOutput {
	return o
}

// Enable/disable adjacency check.
func (o LookupIsisResultOutput) AdjacencyCheck() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.AdjacencyCheck }).(pulumi.StringOutput)
}

// Enable/disable IPv6 adjacency check.
func (o LookupIsisResultOutput) AdjacencyCheck6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.AdjacencyCheck6 }).(pulumi.StringOutput)
}

// Enable/disable IS-IS advertisement of passive interfaces only.
func (o LookupIsisResultOutput) AdvPassiveOnly() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.AdvPassiveOnly }).(pulumi.StringOutput)
}

// Enable/disable IPv6 IS-IS advertisement of passive interfaces only.
func (o LookupIsisResultOutput) AdvPassiveOnly6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.AdvPassiveOnly6 }).(pulumi.StringOutput)
}

// Authentication key-chain for level 1 PDUs.
func (o LookupIsisResultOutput) AuthKeychainL1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.AuthKeychainL1 }).(pulumi.StringOutput)
}

// Authentication key-chain for level 2 PDUs.
func (o LookupIsisResultOutput) AuthKeychainL2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.AuthKeychainL2 }).(pulumi.StringOutput)
}

// Level 1 authentication mode.
func (o LookupIsisResultOutput) AuthModeL1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.AuthModeL1 }).(pulumi.StringOutput)
}

// Level 2 authentication mode.
func (o LookupIsisResultOutput) AuthModeL2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.AuthModeL2 }).(pulumi.StringOutput)
}

// Authentication password for level 1 PDUs.
func (o LookupIsisResultOutput) AuthPasswordL1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.AuthPasswordL1 }).(pulumi.StringOutput)
}

// Authentication password for level 2 PDUs.
func (o LookupIsisResultOutput) AuthPasswordL2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.AuthPasswordL2 }).(pulumi.StringOutput)
}

// Enable/disable level 1 authentication send-only.
func (o LookupIsisResultOutput) AuthSendonlyL1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.AuthSendonlyL1 }).(pulumi.StringOutput)
}

// Enable/disable level 2 authentication send-only.
func (o LookupIsisResultOutput) AuthSendonlyL2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.AuthSendonlyL2 }).(pulumi.StringOutput)
}

// Enable/disable distribution of default route information.
func (o LookupIsisResultOutput) DefaultOriginate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.DefaultOriginate }).(pulumi.StringOutput)
}

// Enable/disable distribution of default IPv6 route information.
func (o LookupIsisResultOutput) DefaultOriginate6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.DefaultOriginate6 }).(pulumi.StringOutput)
}

// Enable/disable dynamic hostname.
func (o LookupIsisResultOutput) DynamicHostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.DynamicHostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIsisResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable ignoring of LSP errors with bad checksums.
func (o LookupIsisResultOutput) IgnoreLspErrors() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.IgnoreLspErrors }).(pulumi.StringOutput)
}

// IS type.
func (o LookupIsisResultOutput) IsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.IsType }).(pulumi.StringOutput)
}

// IS-IS interface configuration. The structure of `isisInterface` block is documented below.
func (o LookupIsisResultOutput) IsisInterfaces() GetIsisIsisInterfaceArrayOutput {
	return o.ApplyT(func(v LookupIsisResult) []GetIsisIsisInterface { return v.IsisInterfaces }).(GetIsisIsisInterfaceArrayOutput)
}

// IS-IS net configuration. The structure of `isisNet` block is documented below.
func (o LookupIsisResultOutput) IsisNets() GetIsisIsisNetArrayOutput {
	return o.ApplyT(func(v LookupIsisResult) []GetIsisIsisNet { return v.IsisNets }).(GetIsisIsisNetArrayOutput)
}

// Minimum interval for level 1 LSP regenerating.
func (o LookupIsisResultOutput) LspGenIntervalL1() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIsisResult) int { return v.LspGenIntervalL1 }).(pulumi.IntOutput)
}

// Minimum interval for level 2 LSP regenerating.
func (o LookupIsisResultOutput) LspGenIntervalL2() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIsisResult) int { return v.LspGenIntervalL2 }).(pulumi.IntOutput)
}

// LSP refresh time in seconds.
func (o LookupIsisResultOutput) LspRefreshInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIsisResult) int { return v.LspRefreshInterval }).(pulumi.IntOutput)
}

// Maximum LSP lifetime in seconds.
func (o LookupIsisResultOutput) MaxLspLifetime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIsisResult) int { return v.MaxLspLifetime }).(pulumi.IntOutput)
}

// Use old-style (ISO 10589) or new-style packet formats
func (o LookupIsisResultOutput) MetricStyle() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.MetricStyle }).(pulumi.StringOutput)
}

// Enable/disable signal other routers not to use us in SPF.
func (o LookupIsisResultOutput) OverloadBit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.OverloadBit }).(pulumi.StringOutput)
}

// Overload-bit only temporarily after reboot.
func (o LookupIsisResultOutput) OverloadBitOnStartup() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIsisResult) int { return v.OverloadBitOnStartup }).(pulumi.IntOutput)
}

// Suppress overload-bit for the specific prefixes.
func (o LookupIsisResultOutput) OverloadBitSuppress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.OverloadBitSuppress }).(pulumi.StringOutput)
}

// Enable/disable redistribution of level 1 IPv6 routes into level 2.
func (o LookupIsisResultOutput) Redistribute6L1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.Redistribute6L1 }).(pulumi.StringOutput)
}

// Access-list for IPv6 route redistribution from l1 to l2.
func (o LookupIsisResultOutput) Redistribute6L1List() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.Redistribute6L1List }).(pulumi.StringOutput)
}

// Enable/disable redistribution of level 2 IPv6 routes into level 1.
func (o LookupIsisResultOutput) Redistribute6L2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.Redistribute6L2 }).(pulumi.StringOutput)
}

// Access-list for IPv6 route redistribution from l2 to l1.
func (o LookupIsisResultOutput) Redistribute6L2List() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.Redistribute6L2List }).(pulumi.StringOutput)
}

// IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.
func (o LookupIsisResultOutput) Redistribute6s() GetIsisRedistribute6ArrayOutput {
	return o.ApplyT(func(v LookupIsisResult) []GetIsisRedistribute6 { return v.Redistribute6s }).(GetIsisRedistribute6ArrayOutput)
}

// Enable/disable redistribution of level 1 routes into level 2.
func (o LookupIsisResultOutput) RedistributeL1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.RedistributeL1 }).(pulumi.StringOutput)
}

// Access-list for route redistribution from l1 to l2.
func (o LookupIsisResultOutput) RedistributeL1List() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.RedistributeL1List }).(pulumi.StringOutput)
}

// Enable/disable redistribution of level 2 routes into level 1.
func (o LookupIsisResultOutput) RedistributeL2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.RedistributeL2 }).(pulumi.StringOutput)
}

// Access-list for route redistribution from l2 to l1.
func (o LookupIsisResultOutput) RedistributeL2List() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.RedistributeL2List }).(pulumi.StringOutput)
}

// IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
func (o LookupIsisResultOutput) Redistributes() GetIsisRedistributeArrayOutput {
	return o.ApplyT(func(v LookupIsisResult) []GetIsisRedistribute { return v.Redistributes }).(GetIsisRedistributeArrayOutput)
}

// Level 1 SPF calculation delay.
func (o LookupIsisResultOutput) SpfIntervalExpL1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.SpfIntervalExpL1 }).(pulumi.StringOutput)
}

// Level 2 SPF calculation delay.
func (o LookupIsisResultOutput) SpfIntervalExpL2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsisResult) string { return v.SpfIntervalExpL2 }).(pulumi.StringOutput)
}

// IS-IS IPv6 summary address. The structure of `summaryAddress6` block is documented below.
func (o LookupIsisResultOutput) SummaryAddress6s() GetIsisSummaryAddress6ArrayOutput {
	return o.ApplyT(func(v LookupIsisResult) []GetIsisSummaryAddress6 { return v.SummaryAddress6s }).(GetIsisSummaryAddress6ArrayOutput)
}

// IS-IS summary addresses. The structure of `summaryAddress` block is documented below.
func (o LookupIsisResultOutput) SummaryAddresses() GetIsisSummaryAddressArrayOutput {
	return o.ApplyT(func(v LookupIsisResult) []GetIsisSummaryAddress { return v.SummaryAddresses }).(GetIsisSummaryAddressArrayOutput)
}

func (o LookupIsisResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIsisResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIsisResultOutput{})
}
