// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios router policy
func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyResult
	err := ctx.Invoke("fortios:router/getPolicy:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicy.
type LookupPolicyArgs struct {
	// Specify the seqNum of the desired router policy.
	SeqNum int `pulumi:"seqNum"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPolicy.
type LookupPolicyResult struct {
	// Action of the policy route.
	Action string `pulumi:"action"`
	// Optional comments.
	Comments string `pulumi:"comments"`
	// Enable/disable negating destination address match.
	DstNegate string `pulumi:"dstNegate"`
	// Destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs []GetPolicyDstaddr `pulumi:"dstaddrs"`
	// Destination IP and mask (x.x.x.x/x). The structure of `dst` block is documented below.
	Dsts []GetPolicyDst `pulumi:"dsts"`
	// End destination port number (0 - 65535).
	EndPort int `pulumi:"endPort"`
	// End source port number (0 - 65535).
	EndSourcePort int `pulumi:"endSourcePort"`
	// IP address of the gateway.
	Gateway string `pulumi:"gateway"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable negation of input device match.
	InputDeviceNegate string `pulumi:"inputDeviceNegate"`
	// Incoming interface name. The structure of `inputDevice` block is documented below.
	InputDevices []GetPolicyInputDevice `pulumi:"inputDevices"`
	// Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
	InternetServiceCustoms []GetPolicyInternetServiceCustom `pulumi:"internetServiceCustoms"`
	// Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
	InternetServiceIds []GetPolicyInternetServiceId `pulumi:"internetServiceIds"`
	// Outgoing interface name.
	OutputDevice string `pulumi:"outputDevice"`
	// Protocol number (0 - 255).
	Protocol int `pulumi:"protocol"`
	// Sequence number.
	SeqNum int `pulumi:"seqNum"`
	// Enable/disable negating source address match.
	SrcNegate string `pulumi:"srcNegate"`
	// Source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs []GetPolicySrcaddr `pulumi:"srcaddrs"`
	// Source IP and mask (x.x.x.x/x). The structure of `src` block is documented below.
	Srcs []GetPolicySrc `pulumi:"srcs"`
	// Start destination port number (0 - 65535).
	StartPort int `pulumi:"startPort"`
	// Start source port number (0 - 65535).
	StartSourcePort int `pulumi:"startSourcePort"`
	// Enable/disable this policy route.
	Status string `pulumi:"status"`
	// Type of service bit pattern.
	Tos string `pulumi:"tos"`
	// Type of service evaluated bits.
	TosMask   string  `pulumi:"tosMask"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupPolicyOutput(ctx *pulumi.Context, args LookupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyResult, error) {
			args := v.(LookupPolicyArgs)
			r, err := LookupPolicy(ctx, &args, opts...)
			var s LookupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyResultOutput)
}

// A collection of arguments for invoking getPolicy.
type LookupPolicyOutputArgs struct {
	// Specify the seqNum of the desired router policy.
	SeqNum pulumi.IntInput `pulumi:"seqNum"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getPolicy.
type LookupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResult)(nil)).Elem()
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutput() LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutputWithContext(ctx context.Context) LookupPolicyResultOutput {
	return o
}

// Action of the policy route.
func (o LookupPolicyResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Action }).(pulumi.StringOutput)
}

// Optional comments.
func (o LookupPolicyResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Comments }).(pulumi.StringOutput)
}

// Enable/disable negating destination address match.
func (o LookupPolicyResultOutput) DstNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.DstNegate }).(pulumi.StringOutput)
}

// Destination address name. The structure of `dstaddr` block is documented below.
func (o LookupPolicyResultOutput) Dstaddrs() GetPolicyDstaddrArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []GetPolicyDstaddr { return v.Dstaddrs }).(GetPolicyDstaddrArrayOutput)
}

// Destination IP and mask (x.x.x.x/x). The structure of `dst` block is documented below.
func (o LookupPolicyResultOutput) Dsts() GetPolicyDstArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []GetPolicyDst { return v.Dsts }).(GetPolicyDstArrayOutput)
}

// End destination port number (0 - 65535).
func (o LookupPolicyResultOutput) EndPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPolicyResult) int { return v.EndPort }).(pulumi.IntOutput)
}

// End source port number (0 - 65535).
func (o LookupPolicyResultOutput) EndSourcePort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPolicyResult) int { return v.EndSourcePort }).(pulumi.IntOutput)
}

// IP address of the gateway.
func (o LookupPolicyResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Gateway }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable negation of input device match.
func (o LookupPolicyResultOutput) InputDeviceNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.InputDeviceNegate }).(pulumi.StringOutput)
}

// Incoming interface name. The structure of `inputDevice` block is documented below.
func (o LookupPolicyResultOutput) InputDevices() GetPolicyInputDeviceArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []GetPolicyInputDevice { return v.InputDevices }).(GetPolicyInputDeviceArrayOutput)
}

// Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
func (o LookupPolicyResultOutput) InternetServiceCustoms() GetPolicyInternetServiceCustomArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []GetPolicyInternetServiceCustom { return v.InternetServiceCustoms }).(GetPolicyInternetServiceCustomArrayOutput)
}

// Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
func (o LookupPolicyResultOutput) InternetServiceIds() GetPolicyInternetServiceIdArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []GetPolicyInternetServiceId { return v.InternetServiceIds }).(GetPolicyInternetServiceIdArrayOutput)
}

// Outgoing interface name.
func (o LookupPolicyResultOutput) OutputDevice() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.OutputDevice }).(pulumi.StringOutput)
}

// Protocol number (0 - 255).
func (o LookupPolicyResultOutput) Protocol() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPolicyResult) int { return v.Protocol }).(pulumi.IntOutput)
}

// Sequence number.
func (o LookupPolicyResultOutput) SeqNum() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPolicyResult) int { return v.SeqNum }).(pulumi.IntOutput)
}

// Enable/disable negating source address match.
func (o LookupPolicyResultOutput) SrcNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.SrcNegate }).(pulumi.StringOutput)
}

// Source address name. The structure of `srcaddr` block is documented below.
func (o LookupPolicyResultOutput) Srcaddrs() GetPolicySrcaddrArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []GetPolicySrcaddr { return v.Srcaddrs }).(GetPolicySrcaddrArrayOutput)
}

// Source IP and mask (x.x.x.x/x). The structure of `src` block is documented below.
func (o LookupPolicyResultOutput) Srcs() GetPolicySrcArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []GetPolicySrc { return v.Srcs }).(GetPolicySrcArrayOutput)
}

// Start destination port number (0 - 65535).
func (o LookupPolicyResultOutput) StartPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPolicyResult) int { return v.StartPort }).(pulumi.IntOutput)
}

// Start source port number (0 - 65535).
func (o LookupPolicyResultOutput) StartSourcePort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPolicyResult) int { return v.StartSourcePort }).(pulumi.IntOutput)
}

// Enable/disable this policy route.
func (o LookupPolicyResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Status }).(pulumi.StringOutput)
}

// Type of service bit pattern.
func (o LookupPolicyResultOutput) Tos() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Tos }).(pulumi.StringOutput)
}

// Type of service evaluated bits.
func (o LookupPolicyResultOutput) TosMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.TosMask }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResultOutput{})
}
