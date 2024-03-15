// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `fortios_firewall_DoSpolicy6`.
func GetDoSpolicy6list(ctx *pulumi.Context, args *GetDoSpolicy6listArgs, opts ...pulumi.InvokeOption) (*GetDoSpolicy6listResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDoSpolicy6listResult
	err := ctx.Invoke("fortios:firewall/getDoSpolicy6list:getDoSpolicy6list", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDoSpolicy6list.
type GetDoSpolicy6listArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getDoSpolicy6list.
type GetDoSpolicy6listResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `fortios_firewall_DoSpolicy6`.
	Policyidlists []int   `pulumi:"policyidlists"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

func GetDoSpolicy6listOutput(ctx *pulumi.Context, args GetDoSpolicy6listOutputArgs, opts ...pulumi.InvokeOption) GetDoSpolicy6listResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDoSpolicy6listResult, error) {
			args := v.(GetDoSpolicy6listArgs)
			r, err := GetDoSpolicy6list(ctx, &args, opts...)
			var s GetDoSpolicy6listResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDoSpolicy6listResultOutput)
}

// A collection of arguments for invoking getDoSpolicy6list.
type GetDoSpolicy6listOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetDoSpolicy6listOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDoSpolicy6listArgs)(nil)).Elem()
}

// A collection of values returned by getDoSpolicy6list.
type GetDoSpolicy6listResultOutput struct{ *pulumi.OutputState }

func (GetDoSpolicy6listResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDoSpolicy6listResult)(nil)).Elem()
}

func (o GetDoSpolicy6listResultOutput) ToGetDoSpolicy6listResultOutput() GetDoSpolicy6listResultOutput {
	return o
}

func (o GetDoSpolicy6listResultOutput) ToGetDoSpolicy6listResultOutputWithContext(ctx context.Context) GetDoSpolicy6listResultOutput {
	return o
}

func (o GetDoSpolicy6listResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDoSpolicy6listResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDoSpolicy6listResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDoSpolicy6listResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `fortios_firewall_DoSpolicy6`.
func (o GetDoSpolicy6listResultOutput) Policyidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetDoSpolicy6listResult) []int { return v.Policyidlists }).(pulumi.IntArrayOutput)
}

func (o GetDoSpolicy6listResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDoSpolicy6listResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDoSpolicy6listResultOutput{})
}
