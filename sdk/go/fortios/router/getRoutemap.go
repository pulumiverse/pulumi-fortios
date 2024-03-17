// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios router routemap
func LookupRoutemap(ctx *pulumi.Context, args *LookupRoutemapArgs, opts ...pulumi.InvokeOption) (*LookupRoutemapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRoutemapResult
	err := ctx.Invoke("fortios:router/getRoutemap:getRoutemap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRoutemap.
type LookupRoutemapArgs struct {
	// Specify the name of the desired router routemap.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRoutemap.
type LookupRoutemapResult struct {
	// Optional comments.
	Comments string `pulumi:"comments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name.
	Name string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules     []GetRoutemapRule `pulumi:"rules"`
	Vdomparam *string           `pulumi:"vdomparam"`
}

func LookupRoutemapOutput(ctx *pulumi.Context, args LookupRoutemapOutputArgs, opts ...pulumi.InvokeOption) LookupRoutemapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoutemapResult, error) {
			args := v.(LookupRoutemapArgs)
			r, err := LookupRoutemap(ctx, &args, opts...)
			var s LookupRoutemapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoutemapResultOutput)
}

// A collection of arguments for invoking getRoutemap.
type LookupRoutemapOutputArgs struct {
	// Specify the name of the desired router routemap.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRoutemapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutemapArgs)(nil)).Elem()
}

// A collection of values returned by getRoutemap.
type LookupRoutemapResultOutput struct{ *pulumi.OutputState }

func (LookupRoutemapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutemapResult)(nil)).Elem()
}

func (o LookupRoutemapResultOutput) ToLookupRoutemapResultOutput() LookupRoutemapResultOutput {
	return o
}

func (o LookupRoutemapResultOutput) ToLookupRoutemapResultOutputWithContext(ctx context.Context) LookupRoutemapResultOutput {
	return o
}

// Optional comments.
func (o LookupRoutemapResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutemapResult) string { return v.Comments }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRoutemapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutemapResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name.
func (o LookupRoutemapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutemapResult) string { return v.Name }).(pulumi.StringOutput)
}

// Rule. The structure of `rule` block is documented below.
func (o LookupRoutemapResultOutput) Rules() GetRoutemapRuleArrayOutput {
	return o.ApplyT(func(v LookupRoutemapResult) []GetRoutemapRule { return v.Rules }).(GetRoutemapRuleArrayOutput)
}

func (o LookupRoutemapResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoutemapResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoutemapResultOutput{})
}
