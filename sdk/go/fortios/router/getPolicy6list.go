// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `router.Policy6`.
func GetPolicy6list(ctx *pulumi.Context, args *GetPolicy6listArgs, opts ...pulumi.InvokeOption) (*GetPolicy6listResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPolicy6listResult
	err := ctx.Invoke("fortios:router/getPolicy6list:getPolicy6list", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicy6list.
type GetPolicy6listArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPolicy6list.
type GetPolicy6listResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `router.Policy6`.
	SeqNumlists []int   `pulumi:"seqNumlists"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

func GetPolicy6listOutput(ctx *pulumi.Context, args GetPolicy6listOutputArgs, opts ...pulumi.InvokeOption) GetPolicy6listResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPolicy6listResult, error) {
			args := v.(GetPolicy6listArgs)
			r, err := GetPolicy6list(ctx, &args, opts...)
			var s GetPolicy6listResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPolicy6listResultOutput)
}

// A collection of arguments for invoking getPolicy6list.
type GetPolicy6listOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetPolicy6listOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicy6listArgs)(nil)).Elem()
}

// A collection of values returned by getPolicy6list.
type GetPolicy6listResultOutput struct{ *pulumi.OutputState }

func (GetPolicy6listResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicy6listResult)(nil)).Elem()
}

func (o GetPolicy6listResultOutput) ToGetPolicy6listResultOutput() GetPolicy6listResultOutput {
	return o
}

func (o GetPolicy6listResultOutput) ToGetPolicy6listResultOutputWithContext(ctx context.Context) GetPolicy6listResultOutput {
	return o
}

func (o GetPolicy6listResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicy6listResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPolicy6listResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicy6listResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `router.Policy6`.
func (o GetPolicy6listResultOutput) SeqNumlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetPolicy6listResult) []int { return v.SeqNumlists }).(pulumi.IntArrayOutput)
}

func (o GetPolicy6listResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicy6listResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPolicy6listResultOutput{})
}
