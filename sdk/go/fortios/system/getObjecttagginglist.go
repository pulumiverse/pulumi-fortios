// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `system.Objecttagging`.
func GetObjecttagginglist(ctx *pulumi.Context, args *GetObjecttagginglistArgs, opts ...pulumi.InvokeOption) (*GetObjecttagginglistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetObjecttagginglistResult
	err := ctx.Invoke("fortios:system/getObjecttagginglist:getObjecttagginglist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getObjecttagginglist.
type GetObjecttagginglistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getObjecttagginglist.
type GetObjecttagginglistResult struct {
	// A list of the `system.Objecttagging`.
	Categorylists []string `pulumi:"categorylists"`
	Filter        *string  `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetObjecttagginglistOutput(ctx *pulumi.Context, args GetObjecttagginglistOutputArgs, opts ...pulumi.InvokeOption) GetObjecttagginglistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetObjecttagginglistResult, error) {
			args := v.(GetObjecttagginglistArgs)
			r, err := GetObjecttagginglist(ctx, &args, opts...)
			var s GetObjecttagginglistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetObjecttagginglistResultOutput)
}

// A collection of arguments for invoking getObjecttagginglist.
type GetObjecttagginglistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetObjecttagginglistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetObjecttagginglistArgs)(nil)).Elem()
}

// A collection of values returned by getObjecttagginglist.
type GetObjecttagginglistResultOutput struct{ *pulumi.OutputState }

func (GetObjecttagginglistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetObjecttagginglistResult)(nil)).Elem()
}

func (o GetObjecttagginglistResultOutput) ToGetObjecttagginglistResultOutput() GetObjecttagginglistResultOutput {
	return o
}

func (o GetObjecttagginglistResultOutput) ToGetObjecttagginglistResultOutputWithContext(ctx context.Context) GetObjecttagginglistResultOutput {
	return o
}

// A list of the `system.Objecttagging`.
func (o GetObjecttagginglistResultOutput) Categorylists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetObjecttagginglistResult) []string { return v.Categorylists }).(pulumi.StringArrayOutput)
}

func (o GetObjecttagginglistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetObjecttagginglistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetObjecttagginglistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjecttagginglistResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetObjecttagginglistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetObjecttagginglistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetObjecttagginglistResultOutput{})
}
