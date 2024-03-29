// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `firewall.Policy64`.
func GetPolicy64list(ctx *pulumi.Context, args *GetPolicy64listArgs, opts ...pulumi.InvokeOption) (*GetPolicy64listResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPolicy64listResult
	err := ctx.Invoke("fortios:firewall/getPolicy64list:getPolicy64list", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicy64list.
type GetPolicy64listArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPolicy64list.
type GetPolicy64listResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall.Policy64`.
	Policyidlists []int   `pulumi:"policyidlists"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

func GetPolicy64listOutput(ctx *pulumi.Context, args GetPolicy64listOutputArgs, opts ...pulumi.InvokeOption) GetPolicy64listResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPolicy64listResult, error) {
			args := v.(GetPolicy64listArgs)
			r, err := GetPolicy64list(ctx, &args, opts...)
			var s GetPolicy64listResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPolicy64listResultOutput)
}

// A collection of arguments for invoking getPolicy64list.
type GetPolicy64listOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetPolicy64listOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicy64listArgs)(nil)).Elem()
}

// A collection of values returned by getPolicy64list.
type GetPolicy64listResultOutput struct{ *pulumi.OutputState }

func (GetPolicy64listResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicy64listResult)(nil)).Elem()
}

func (o GetPolicy64listResultOutput) ToGetPolicy64listResultOutput() GetPolicy64listResultOutput {
	return o
}

func (o GetPolicy64listResultOutput) ToGetPolicy64listResultOutputWithContext(ctx context.Context) GetPolicy64listResultOutput {
	return o
}

func (o GetPolicy64listResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicy64listResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPolicy64listResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicy64listResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall.Policy64`.
func (o GetPolicy64listResultOutput) Policyidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetPolicy64listResult) []int { return v.Policyidlists }).(pulumi.IntArrayOutput)
}

func (o GetPolicy64listResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicy64listResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPolicy64listResultOutput{})
}
