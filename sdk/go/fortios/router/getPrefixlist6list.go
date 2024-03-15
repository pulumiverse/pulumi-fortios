// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `router.Prefixlist6`.
func GetPrefixlist6list(ctx *pulumi.Context, args *GetPrefixlist6listArgs, opts ...pulumi.InvokeOption) (*GetPrefixlist6listResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPrefixlist6listResult
	err := ctx.Invoke("fortios:router/getPrefixlist6list:getPrefixlist6list", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrefixlist6list.
type GetPrefixlist6listArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPrefixlist6list.
type GetPrefixlist6listResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `router.Prefixlist6`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetPrefixlist6listOutput(ctx *pulumi.Context, args GetPrefixlist6listOutputArgs, opts ...pulumi.InvokeOption) GetPrefixlist6listResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPrefixlist6listResult, error) {
			args := v.(GetPrefixlist6listArgs)
			r, err := GetPrefixlist6list(ctx, &args, opts...)
			var s GetPrefixlist6listResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPrefixlist6listResultOutput)
}

// A collection of arguments for invoking getPrefixlist6list.
type GetPrefixlist6listOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetPrefixlist6listOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPrefixlist6listArgs)(nil)).Elem()
}

// A collection of values returned by getPrefixlist6list.
type GetPrefixlist6listResultOutput struct{ *pulumi.OutputState }

func (GetPrefixlist6listResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPrefixlist6listResult)(nil)).Elem()
}

func (o GetPrefixlist6listResultOutput) ToGetPrefixlist6listResultOutput() GetPrefixlist6listResultOutput {
	return o
}

func (o GetPrefixlist6listResultOutput) ToGetPrefixlist6listResultOutputWithContext(ctx context.Context) GetPrefixlist6listResultOutput {
	return o
}

func (o GetPrefixlist6listResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPrefixlist6listResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPrefixlist6listResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPrefixlist6listResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `router.Prefixlist6`.
func (o GetPrefixlist6listResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPrefixlist6listResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetPrefixlist6listResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPrefixlist6listResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPrefixlist6listResultOutput{})
}
