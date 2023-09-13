// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `sys.Replacemsggroup`.
func GetReplacemsggrouplist(ctx *pulumi.Context, args *GetReplacemsggrouplistArgs, opts ...pulumi.InvokeOption) (*GetReplacemsggrouplistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetReplacemsggrouplistResult
	err := ctx.Invoke("fortios:sys/getReplacemsggrouplist:getReplacemsggrouplist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReplacemsggrouplist.
type GetReplacemsggrouplistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getReplacemsggrouplist.
type GetReplacemsggrouplistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `sys.Replacemsggroup`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetReplacemsggrouplistOutput(ctx *pulumi.Context, args GetReplacemsggrouplistOutputArgs, opts ...pulumi.InvokeOption) GetReplacemsggrouplistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetReplacemsggrouplistResult, error) {
			args := v.(GetReplacemsggrouplistArgs)
			r, err := GetReplacemsggrouplist(ctx, &args, opts...)
			var s GetReplacemsggrouplistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetReplacemsggrouplistResultOutput)
}

// A collection of arguments for invoking getReplacemsggrouplist.
type GetReplacemsggrouplistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetReplacemsggrouplistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReplacemsggrouplistArgs)(nil)).Elem()
}

// A collection of values returned by getReplacemsggrouplist.
type GetReplacemsggrouplistResultOutput struct{ *pulumi.OutputState }

func (GetReplacemsggrouplistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReplacemsggrouplistResult)(nil)).Elem()
}

func (o GetReplacemsggrouplistResultOutput) ToGetReplacemsggrouplistResultOutput() GetReplacemsggrouplistResultOutput {
	return o
}

func (o GetReplacemsggrouplistResultOutput) ToGetReplacemsggrouplistResultOutputWithContext(ctx context.Context) GetReplacemsggrouplistResultOutput {
	return o
}

func (o GetReplacemsggrouplistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetReplacemsggrouplistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetReplacemsggrouplistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetReplacemsggrouplistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `sys.Replacemsggroup`.
func (o GetReplacemsggrouplistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetReplacemsggrouplistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetReplacemsggrouplistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetReplacemsggrouplistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetReplacemsggrouplistResultOutput{})
}
