// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `firewall.Policy`.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/firewall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			sample1, err := firewall.GetPolicylist(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("output1", sample1)
//			sample2, err := firewall.GetPolicylist(ctx, &firewall.GetPolicylistArgs{
//				Filter: pulumi.StringRef("schedule==always&action==accept,action==deny"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("sample2Output", sample2.Policyidlists)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetPolicylist(ctx *pulumi.Context, args *GetPolicylistArgs, opts ...pulumi.InvokeOption) (*GetPolicylistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPolicylistResult
	err := ctx.Invoke("fortios:firewall/getPolicylist:getPolicylist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicylist.
type GetPolicylistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPolicylist.
type GetPolicylistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall.Policy`.
	Policyidlists []int   `pulumi:"policyidlists"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

func GetPolicylistOutput(ctx *pulumi.Context, args GetPolicylistOutputArgs, opts ...pulumi.InvokeOption) GetPolicylistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPolicylistResult, error) {
			args := v.(GetPolicylistArgs)
			r, err := GetPolicylist(ctx, &args, opts...)
			var s GetPolicylistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPolicylistResultOutput)
}

// A collection of arguments for invoking getPolicylist.
type GetPolicylistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetPolicylistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicylistArgs)(nil)).Elem()
}

// A collection of values returned by getPolicylist.
type GetPolicylistResultOutput struct{ *pulumi.OutputState }

func (GetPolicylistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicylistResult)(nil)).Elem()
}

func (o GetPolicylistResultOutput) ToGetPolicylistResultOutput() GetPolicylistResultOutput {
	return o
}

func (o GetPolicylistResultOutput) ToGetPolicylistResultOutputWithContext(ctx context.Context) GetPolicylistResultOutput {
	return o
}

func (o GetPolicylistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicylistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPolicylistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicylistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall.Policy`.
func (o GetPolicylistResultOutput) Policyidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetPolicylistResult) []int { return v.Policyidlists }).(pulumi.IntArrayOutput)
}

func (o GetPolicylistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicylistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPolicylistResultOutput{})
}
