// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallAddress6template`.
func GetFirewallAddress6templatelist(ctx *pulumi.Context, args *GetFirewallAddress6templatelistArgs, opts ...pulumi.InvokeOption) (*GetFirewallAddress6templatelistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallAddress6templatelistResult
	err := ctx.Invoke("fortios:index/getFirewallAddress6templatelist:getFirewallAddress6templatelist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallAddress6templatelist.
type GetFirewallAddress6templatelistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallAddress6templatelist.
type GetFirewallAddress6templatelistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallAddress6template`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallAddress6templatelistOutput(ctx *pulumi.Context, args GetFirewallAddress6templatelistOutputArgs, opts ...pulumi.InvokeOption) GetFirewallAddress6templatelistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallAddress6templatelistResult, error) {
			args := v.(GetFirewallAddress6templatelistArgs)
			r, err := GetFirewallAddress6templatelist(ctx, &args, opts...)
			var s GetFirewallAddress6templatelistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallAddress6templatelistResultOutput)
}

// A collection of arguments for invoking getFirewallAddress6templatelist.
type GetFirewallAddress6templatelistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallAddress6templatelistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallAddress6templatelistArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallAddress6templatelist.
type GetFirewallAddress6templatelistResultOutput struct{ *pulumi.OutputState }

func (GetFirewallAddress6templatelistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallAddress6templatelistResult)(nil)).Elem()
}

func (o GetFirewallAddress6templatelistResultOutput) ToGetFirewallAddress6templatelistResultOutput() GetFirewallAddress6templatelistResultOutput {
	return o
}

func (o GetFirewallAddress6templatelistResultOutput) ToGetFirewallAddress6templatelistResultOutputWithContext(ctx context.Context) GetFirewallAddress6templatelistResultOutput {
	return o
}

func (o GetFirewallAddress6templatelistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallAddress6templatelistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallAddress6templatelistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallAddress6templatelistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallAddress6template`.
func (o GetFirewallAddress6templatelistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallAddress6templatelistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallAddress6templatelistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallAddress6templatelistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallAddress6templatelistResultOutput{})
}
