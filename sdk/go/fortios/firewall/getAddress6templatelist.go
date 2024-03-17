// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `firewall.Address6template`.
func GetAddress6templatelist(ctx *pulumi.Context, args *GetAddress6templatelistArgs, opts ...pulumi.InvokeOption) (*GetAddress6templatelistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAddress6templatelistResult
	err := ctx.Invoke("fortios:firewall/getAddress6templatelist:getAddress6templatelist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAddress6templatelist.
type GetAddress6templatelistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getAddress6templatelist.
type GetAddress6templatelistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall.Address6template`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetAddress6templatelistOutput(ctx *pulumi.Context, args GetAddress6templatelistOutputArgs, opts ...pulumi.InvokeOption) GetAddress6templatelistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAddress6templatelistResult, error) {
			args := v.(GetAddress6templatelistArgs)
			r, err := GetAddress6templatelist(ctx, &args, opts...)
			var s GetAddress6templatelistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAddress6templatelistResultOutput)
}

// A collection of arguments for invoking getAddress6templatelist.
type GetAddress6templatelistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetAddress6templatelistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddress6templatelistArgs)(nil)).Elem()
}

// A collection of values returned by getAddress6templatelist.
type GetAddress6templatelistResultOutput struct{ *pulumi.OutputState }

func (GetAddress6templatelistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddress6templatelistResult)(nil)).Elem()
}

func (o GetAddress6templatelistResultOutput) ToGetAddress6templatelistResultOutput() GetAddress6templatelistResultOutput {
	return o
}

func (o GetAddress6templatelistResultOutput) ToGetAddress6templatelistResultOutputWithContext(ctx context.Context) GetAddress6templatelistResultOutput {
	return o
}

func (o GetAddress6templatelistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddress6templatelistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAddress6templatelistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAddress6templatelistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall.Address6template`.
func (o GetAddress6templatelistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAddress6templatelistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetAddress6templatelistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddress6templatelistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAddress6templatelistResultOutput{})
}
