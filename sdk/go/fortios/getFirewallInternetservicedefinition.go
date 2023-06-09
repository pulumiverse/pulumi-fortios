// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall internetservicedefinition
func LookupFirewallInternetservicedefinition(ctx *pulumi.Context, args *LookupFirewallInternetservicedefinitionArgs, opts ...pulumi.InvokeOption) (*LookupFirewallInternetservicedefinitionResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallInternetservicedefinitionResult
	err := ctx.Invoke("fortios:index/getFirewallInternetservicedefinition:getFirewallInternetservicedefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallInternetservicedefinition.
type LookupFirewallInternetservicedefinitionArgs struct {
	// Specify the fosid of the desired firewall internetservicedefinition.
	Fosid int `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallInternetservicedefinition.
type LookupFirewallInternetservicedefinitionResult struct {
	// Protocol and port information in an Internet Service entry. The structure of `entry` block is documented below.
	Entries []GetFirewallInternetservicedefinitionEntry `pulumi:"entries"`
	// Internet Service application list ID.
	Fosid int `pulumi:"fosid"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupFirewallInternetservicedefinitionOutput(ctx *pulumi.Context, args LookupFirewallInternetservicedefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallInternetservicedefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallInternetservicedefinitionResult, error) {
			args := v.(LookupFirewallInternetservicedefinitionArgs)
			r, err := LookupFirewallInternetservicedefinition(ctx, &args, opts...)
			var s LookupFirewallInternetservicedefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallInternetservicedefinitionResultOutput)
}

// A collection of arguments for invoking getFirewallInternetservicedefinition.
type LookupFirewallInternetservicedefinitionOutputArgs struct {
	// Specify the fosid of the desired firewall internetservicedefinition.
	Fosid pulumi.IntInput `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallInternetservicedefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetservicedefinitionArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallInternetservicedefinition.
type LookupFirewallInternetservicedefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallInternetservicedefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetservicedefinitionResult)(nil)).Elem()
}

func (o LookupFirewallInternetservicedefinitionResultOutput) ToLookupFirewallInternetservicedefinitionResultOutput() LookupFirewallInternetservicedefinitionResultOutput {
	return o
}

func (o LookupFirewallInternetservicedefinitionResultOutput) ToLookupFirewallInternetservicedefinitionResultOutputWithContext(ctx context.Context) LookupFirewallInternetservicedefinitionResultOutput {
	return o
}

// Protocol and port information in an Internet Service entry. The structure of `entry` block is documented below.
func (o LookupFirewallInternetservicedefinitionResultOutput) Entries() GetFirewallInternetservicedefinitionEntryArrayOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicedefinitionResult) []GetFirewallInternetservicedefinitionEntry {
		return v.Entries
	}).(GetFirewallInternetservicedefinitionEntryArrayOutput)
}

// Internet Service application list ID.
func (o LookupFirewallInternetservicedefinitionResultOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicedefinitionResult) int { return v.Fosid }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallInternetservicedefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicedefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFirewallInternetservicedefinitionResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallInternetservicedefinitionResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallInternetservicedefinitionResultOutput{})
}
