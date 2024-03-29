// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `firewall.Policy46`.
func GetPolicy46list(ctx *pulumi.Context, args *GetPolicy46listArgs, opts ...pulumi.InvokeOption) (*GetPolicy46listResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPolicy46listResult
	err := ctx.Invoke("fortios:firewall/getPolicy46list:getPolicy46list", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicy46list.
type GetPolicy46listArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPolicy46list.
type GetPolicy46listResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall.Policy46`.
	Policyidlists []int   `pulumi:"policyidlists"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

func GetPolicy46listOutput(ctx *pulumi.Context, args GetPolicy46listOutputArgs, opts ...pulumi.InvokeOption) GetPolicy46listResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPolicy46listResult, error) {
			args := v.(GetPolicy46listArgs)
			r, err := GetPolicy46list(ctx, &args, opts...)
			var s GetPolicy46listResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPolicy46listResultOutput)
}

// A collection of arguments for invoking getPolicy46list.
type GetPolicy46listOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetPolicy46listOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicy46listArgs)(nil)).Elem()
}

// A collection of values returned by getPolicy46list.
type GetPolicy46listResultOutput struct{ *pulumi.OutputState }

func (GetPolicy46listResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicy46listResult)(nil)).Elem()
}

func (o GetPolicy46listResultOutput) ToGetPolicy46listResultOutput() GetPolicy46listResultOutput {
	return o
}

func (o GetPolicy46listResultOutput) ToGetPolicy46listResultOutputWithContext(ctx context.Context) GetPolicy46listResultOutput {
	return o
}

func (o GetPolicy46listResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicy46listResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPolicy46listResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicy46listResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall.Policy46`.
func (o GetPolicy46listResultOutput) Policyidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetPolicy46listResult) []int { return v.Policyidlists }).(pulumi.IntArrayOutput)
}

func (o GetPolicy46listResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicy46listResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPolicy46listResultOutput{})
}
