// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `RouterStatic`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			sample1, err := fortios.GetRouterStaticlist(ctx, &fortios.GetRouterStaticlistArgs{
//				Filter: pulumi.StringRef("seq_num>1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("output1", sample1.SeqNumlists)
//			return nil
//		})
//	}
//
// ```
func GetRouterStaticlist(ctx *pulumi.Context, args *GetRouterStaticlistArgs, opts ...pulumi.InvokeOption) (*GetRouterStaticlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRouterStaticlistResult
	err := ctx.Invoke("fortios:index/getRouterStaticlist:getRouterStaticlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouterStaticlist.
type GetRouterStaticlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRouterStaticlist.
type GetRouterStaticlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `RouterStatic`.
	SeqNumlists []int   `pulumi:"seqNumlists"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

func GetRouterStaticlistOutput(ctx *pulumi.Context, args GetRouterStaticlistOutputArgs, opts ...pulumi.InvokeOption) GetRouterStaticlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouterStaticlistResult, error) {
			args := v.(GetRouterStaticlistArgs)
			r, err := GetRouterStaticlist(ctx, &args, opts...)
			var s GetRouterStaticlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRouterStaticlistResultOutput)
}

// A collection of arguments for invoking getRouterStaticlist.
type GetRouterStaticlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetRouterStaticlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterStaticlistArgs)(nil)).Elem()
}

// A collection of values returned by getRouterStaticlist.
type GetRouterStaticlistResultOutput struct{ *pulumi.OutputState }

func (GetRouterStaticlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterStaticlistResult)(nil)).Elem()
}

func (o GetRouterStaticlistResultOutput) ToGetRouterStaticlistResultOutput() GetRouterStaticlistResultOutput {
	return o
}

func (o GetRouterStaticlistResultOutput) ToGetRouterStaticlistResultOutputWithContext(ctx context.Context) GetRouterStaticlistResultOutput {
	return o
}

func (o GetRouterStaticlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterStaticlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouterStaticlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouterStaticlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `RouterStatic`.
func (o GetRouterStaticlistResultOutput) SeqNumlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetRouterStaticlistResult) []int { return v.SeqNumlists }).(pulumi.IntArrayOutput)
}

func (o GetRouterStaticlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterStaticlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouterStaticlistResultOutput{})
}
