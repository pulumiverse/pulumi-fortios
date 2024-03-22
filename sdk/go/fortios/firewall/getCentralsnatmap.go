// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios firewall centralsnatmap
func LookupCentralsnatmap(ctx *pulumi.Context, args *LookupCentralsnatmapArgs, opts ...pulumi.InvokeOption) (*LookupCentralsnatmapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCentralsnatmapResult
	err := ctx.Invoke("fortios:firewall/getCentralsnatmap:getCentralsnatmap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCentralsnatmap.
type LookupCentralsnatmapArgs struct {
	// Specify the policyid of the desired firewall centralsnatmap.
	Policyid int `pulumi:"policyid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getCentralsnatmap.
type LookupCentralsnatmapResult struct {
	// Comment.
	Comments string `pulumi:"comments"`
	// IPv6 Destination address. The structure of `dstAddr6` block is documented below.
	DstAddr6s []GetCentralsnatmapDstAddr6 `pulumi:"dstAddr6s"`
	// Destination address name from available addresses. The structure of `dstAddr` block is documented below.
	DstAddrs []GetCentralsnatmapDstAddr `pulumi:"dstAddrs"`
	// Destination port or port range (1 to 65535, 0 means any port).
	DstPort string `pulumi:"dstPort"`
	// Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
	Dstintfs []GetCentralsnatmapDstintf `pulumi:"dstintfs"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable source NAT.
	Nat string `pulumi:"nat"`
	// Enable/disable NAT46.
	Nat46 string `pulumi:"nat46"`
	// Enable/disable NAT64.
	Nat64 string `pulumi:"nat64"`
	// IPv6 pools to be used for source NAT. The structure of `natIppool6` block is documented below.
	NatIppool6s []GetCentralsnatmapNatIppool6 `pulumi:"natIppool6s"`
	// Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `natIppool` block is documented below.
	NatIppools []GetCentralsnatmapNatIppool `pulumi:"natIppools"`
	// Translated port or port range (0 to 65535).
	NatPort string `pulumi:"natPort"`
	// IPv6 Original address. The structure of `origAddr6` block is documented below.
	OrigAddr6s []GetCentralsnatmapOrigAddr6 `pulumi:"origAddr6s"`
	// Original address. The structure of `origAddr` block is documented below.
	OrigAddrs []GetCentralsnatmapOrigAddr `pulumi:"origAddrs"`
	// Original TCP port (0 to 65535).
	OrigPort string `pulumi:"origPort"`
	// Policy ID.
	Policyid int `pulumi:"policyid"`
	// Integer value for the protocol type (0 - 255).
	Protocol int `pulumi:"protocol"`
	// Source interface name from available interfaces. The structure of `srcintf` block is documented below.
	Srcintfs []GetCentralsnatmapSrcintf `pulumi:"srcintfs"`
	// Enable/disable the active status of this policy.
	Status string `pulumi:"status"`
	// IPv4/IPv6 source NAT.
	Type string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupCentralsnatmapOutput(ctx *pulumi.Context, args LookupCentralsnatmapOutputArgs, opts ...pulumi.InvokeOption) LookupCentralsnatmapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCentralsnatmapResult, error) {
			args := v.(LookupCentralsnatmapArgs)
			r, err := LookupCentralsnatmap(ctx, &args, opts...)
			var s LookupCentralsnatmapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCentralsnatmapResultOutput)
}

// A collection of arguments for invoking getCentralsnatmap.
type LookupCentralsnatmapOutputArgs struct {
	// Specify the policyid of the desired firewall centralsnatmap.
	Policyid pulumi.IntInput `pulumi:"policyid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupCentralsnatmapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCentralsnatmapArgs)(nil)).Elem()
}

// A collection of values returned by getCentralsnatmap.
type LookupCentralsnatmapResultOutput struct{ *pulumi.OutputState }

func (LookupCentralsnatmapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCentralsnatmapResult)(nil)).Elem()
}

func (o LookupCentralsnatmapResultOutput) ToLookupCentralsnatmapResultOutput() LookupCentralsnatmapResultOutput {
	return o
}

func (o LookupCentralsnatmapResultOutput) ToLookupCentralsnatmapResultOutputWithContext(ctx context.Context) LookupCentralsnatmapResultOutput {
	return o
}

// Comment.
func (o LookupCentralsnatmapResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) string { return v.Comments }).(pulumi.StringOutput)
}

// IPv6 Destination address. The structure of `dstAddr6` block is documented below.
func (o LookupCentralsnatmapResultOutput) DstAddr6s() GetCentralsnatmapDstAddr6ArrayOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) []GetCentralsnatmapDstAddr6 { return v.DstAddr6s }).(GetCentralsnatmapDstAddr6ArrayOutput)
}

// Destination address name from available addresses. The structure of `dstAddr` block is documented below.
func (o LookupCentralsnatmapResultOutput) DstAddrs() GetCentralsnatmapDstAddrArrayOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) []GetCentralsnatmapDstAddr { return v.DstAddrs }).(GetCentralsnatmapDstAddrArrayOutput)
}

// Destination port or port range (1 to 65535, 0 means any port).
func (o LookupCentralsnatmapResultOutput) DstPort() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) string { return v.DstPort }).(pulumi.StringOutput)
}

// Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
func (o LookupCentralsnatmapResultOutput) Dstintfs() GetCentralsnatmapDstintfArrayOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) []GetCentralsnatmapDstintf { return v.Dstintfs }).(GetCentralsnatmapDstintfArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCentralsnatmapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable source NAT.
func (o LookupCentralsnatmapResultOutput) Nat() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) string { return v.Nat }).(pulumi.StringOutput)
}

// Enable/disable NAT46.
func (o LookupCentralsnatmapResultOutput) Nat46() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) string { return v.Nat46 }).(pulumi.StringOutput)
}

// Enable/disable NAT64.
func (o LookupCentralsnatmapResultOutput) Nat64() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) string { return v.Nat64 }).(pulumi.StringOutput)
}

// IPv6 pools to be used for source NAT. The structure of `natIppool6` block is documented below.
func (o LookupCentralsnatmapResultOutput) NatIppool6s() GetCentralsnatmapNatIppool6ArrayOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) []GetCentralsnatmapNatIppool6 { return v.NatIppool6s }).(GetCentralsnatmapNatIppool6ArrayOutput)
}

// Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `natIppool` block is documented below.
func (o LookupCentralsnatmapResultOutput) NatIppools() GetCentralsnatmapNatIppoolArrayOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) []GetCentralsnatmapNatIppool { return v.NatIppools }).(GetCentralsnatmapNatIppoolArrayOutput)
}

// Translated port or port range (0 to 65535).
func (o LookupCentralsnatmapResultOutput) NatPort() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) string { return v.NatPort }).(pulumi.StringOutput)
}

// IPv6 Original address. The structure of `origAddr6` block is documented below.
func (o LookupCentralsnatmapResultOutput) OrigAddr6s() GetCentralsnatmapOrigAddr6ArrayOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) []GetCentralsnatmapOrigAddr6 { return v.OrigAddr6s }).(GetCentralsnatmapOrigAddr6ArrayOutput)
}

// Original address. The structure of `origAddr` block is documented below.
func (o LookupCentralsnatmapResultOutput) OrigAddrs() GetCentralsnatmapOrigAddrArrayOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) []GetCentralsnatmapOrigAddr { return v.OrigAddrs }).(GetCentralsnatmapOrigAddrArrayOutput)
}

// Original TCP port (0 to 65535).
func (o LookupCentralsnatmapResultOutput) OrigPort() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) string { return v.OrigPort }).(pulumi.StringOutput)
}

// Policy ID.
func (o LookupCentralsnatmapResultOutput) Policyid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) int { return v.Policyid }).(pulumi.IntOutput)
}

// Integer value for the protocol type (0 - 255).
func (o LookupCentralsnatmapResultOutput) Protocol() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) int { return v.Protocol }).(pulumi.IntOutput)
}

// Source interface name from available interfaces. The structure of `srcintf` block is documented below.
func (o LookupCentralsnatmapResultOutput) Srcintfs() GetCentralsnatmapSrcintfArrayOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) []GetCentralsnatmapSrcintf { return v.Srcintfs }).(GetCentralsnatmapSrcintfArrayOutput)
}

// Enable/disable the active status of this policy.
func (o LookupCentralsnatmapResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) string { return v.Status }).(pulumi.StringOutput)
}

// IPv4/IPv6 source NAT.
func (o LookupCentralsnatmapResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) string { return v.Type }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o LookupCentralsnatmapResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupCentralsnatmapResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCentralsnatmapResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCentralsnatmapResultOutput{})
}
