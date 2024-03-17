// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `system.Pppoeinterface`.
func GetPppoeinterfacelist(ctx *pulumi.Context, args *GetPppoeinterfacelistArgs, opts ...pulumi.InvokeOption) (*GetPppoeinterfacelistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPppoeinterfacelistResult
	err := ctx.Invoke("fortios:system/getPppoeinterfacelist:getPppoeinterfacelist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPppoeinterfacelist.
type GetPppoeinterfacelistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPppoeinterfacelist.
type GetPppoeinterfacelistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `system.Pppoeinterface`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetPppoeinterfacelistOutput(ctx *pulumi.Context, args GetPppoeinterfacelistOutputArgs, opts ...pulumi.InvokeOption) GetPppoeinterfacelistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPppoeinterfacelistResult, error) {
			args := v.(GetPppoeinterfacelistArgs)
			r, err := GetPppoeinterfacelist(ctx, &args, opts...)
			var s GetPppoeinterfacelistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPppoeinterfacelistResultOutput)
}

// A collection of arguments for invoking getPppoeinterfacelist.
type GetPppoeinterfacelistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetPppoeinterfacelistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPppoeinterfacelistArgs)(nil)).Elem()
}

// A collection of values returned by getPppoeinterfacelist.
type GetPppoeinterfacelistResultOutput struct{ *pulumi.OutputState }

func (GetPppoeinterfacelistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPppoeinterfacelistResult)(nil)).Elem()
}

func (o GetPppoeinterfacelistResultOutput) ToGetPppoeinterfacelistResultOutput() GetPppoeinterfacelistResultOutput {
	return o
}

func (o GetPppoeinterfacelistResultOutput) ToGetPppoeinterfacelistResultOutputWithContext(ctx context.Context) GetPppoeinterfacelistResultOutput {
	return o
}

func (o GetPppoeinterfacelistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPppoeinterfacelistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPppoeinterfacelistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPppoeinterfacelistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `system.Pppoeinterface`.
func (o GetPppoeinterfacelistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPppoeinterfacelistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetPppoeinterfacelistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPppoeinterfacelistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPppoeinterfacelistResultOutput{})
}
