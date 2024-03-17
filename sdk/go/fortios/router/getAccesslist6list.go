// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `router.Accesslist6`.
func GetAccesslist6list(ctx *pulumi.Context, args *GetAccesslist6listArgs, opts ...pulumi.InvokeOption) (*GetAccesslist6listResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccesslist6listResult
	err := ctx.Invoke("fortios:router/getAccesslist6list:getAccesslist6list", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccesslist6list.
type GetAccesslist6listArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getAccesslist6list.
type GetAccesslist6listResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `router.Accesslist6`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetAccesslist6listOutput(ctx *pulumi.Context, args GetAccesslist6listOutputArgs, opts ...pulumi.InvokeOption) GetAccesslist6listResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccesslist6listResult, error) {
			args := v.(GetAccesslist6listArgs)
			r, err := GetAccesslist6list(ctx, &args, opts...)
			var s GetAccesslist6listResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccesslist6listResultOutput)
}

// A collection of arguments for invoking getAccesslist6list.
type GetAccesslist6listOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetAccesslist6listOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccesslist6listArgs)(nil)).Elem()
}

// A collection of values returned by getAccesslist6list.
type GetAccesslist6listResultOutput struct{ *pulumi.OutputState }

func (GetAccesslist6listResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccesslist6listResult)(nil)).Elem()
}

func (o GetAccesslist6listResultOutput) ToGetAccesslist6listResultOutput() GetAccesslist6listResultOutput {
	return o
}

func (o GetAccesslist6listResultOutput) ToGetAccesslist6listResultOutputWithContext(ctx context.Context) GetAccesslist6listResultOutput {
	return o
}

func (o GetAccesslist6listResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccesslist6listResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccesslist6listResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccesslist6listResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `router.Accesslist6`.
func (o GetAccesslist6listResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccesslist6listResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetAccesslist6listResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccesslist6listResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccesslist6listResultOutput{})
}
