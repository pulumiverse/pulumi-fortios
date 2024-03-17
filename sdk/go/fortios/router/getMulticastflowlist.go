// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `router.Multicastflow`.
func GetMulticastflowlist(ctx *pulumi.Context, args *GetMulticastflowlistArgs, opts ...pulumi.InvokeOption) (*GetMulticastflowlistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMulticastflowlistResult
	err := ctx.Invoke("fortios:router/getMulticastflowlist:getMulticastflowlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMulticastflowlist.
type GetMulticastflowlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getMulticastflowlist.
type GetMulticastflowlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `router.Multicastflow`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetMulticastflowlistOutput(ctx *pulumi.Context, args GetMulticastflowlistOutputArgs, opts ...pulumi.InvokeOption) GetMulticastflowlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMulticastflowlistResult, error) {
			args := v.(GetMulticastflowlistArgs)
			r, err := GetMulticastflowlist(ctx, &args, opts...)
			var s GetMulticastflowlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMulticastflowlistResultOutput)
}

// A collection of arguments for invoking getMulticastflowlist.
type GetMulticastflowlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetMulticastflowlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMulticastflowlistArgs)(nil)).Elem()
}

// A collection of values returned by getMulticastflowlist.
type GetMulticastflowlistResultOutput struct{ *pulumi.OutputState }

func (GetMulticastflowlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMulticastflowlistResult)(nil)).Elem()
}

func (o GetMulticastflowlistResultOutput) ToGetMulticastflowlistResultOutput() GetMulticastflowlistResultOutput {
	return o
}

func (o GetMulticastflowlistResultOutput) ToGetMulticastflowlistResultOutputWithContext(ctx context.Context) GetMulticastflowlistResultOutput {
	return o
}

func (o GetMulticastflowlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMulticastflowlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMulticastflowlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMulticastflowlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `router.Multicastflow`.
func (o GetMulticastflowlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMulticastflowlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetMulticastflowlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMulticastflowlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMulticastflowlistResultOutput{})
}
