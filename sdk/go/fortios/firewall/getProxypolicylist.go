// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `firewall.Proxypolicy`.
func GetProxypolicylist(ctx *pulumi.Context, args *GetProxypolicylistArgs, opts ...pulumi.InvokeOption) (*GetProxypolicylistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetProxypolicylistResult
	err := ctx.Invoke("fortios:firewall/getProxypolicylist:getProxypolicylist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProxypolicylist.
type GetProxypolicylistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getProxypolicylist.
type GetProxypolicylistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall.Proxypolicy`.
	Policyidlists []int   `pulumi:"policyidlists"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

func GetProxypolicylistOutput(ctx *pulumi.Context, args GetProxypolicylistOutputArgs, opts ...pulumi.InvokeOption) GetProxypolicylistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProxypolicylistResult, error) {
			args := v.(GetProxypolicylistArgs)
			r, err := GetProxypolicylist(ctx, &args, opts...)
			var s GetProxypolicylistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProxypolicylistResultOutput)
}

// A collection of arguments for invoking getProxypolicylist.
type GetProxypolicylistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetProxypolicylistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicylistArgs)(nil)).Elem()
}

// A collection of values returned by getProxypolicylist.
type GetProxypolicylistResultOutput struct{ *pulumi.OutputState }

func (GetProxypolicylistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicylistResult)(nil)).Elem()
}

func (o GetProxypolicylistResultOutput) ToGetProxypolicylistResultOutput() GetProxypolicylistResultOutput {
	return o
}

func (o GetProxypolicylistResultOutput) ToGetProxypolicylistResultOutputWithContext(ctx context.Context) GetProxypolicylistResultOutput {
	return o
}

func (o GetProxypolicylistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProxypolicylistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProxypolicylistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxypolicylistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall.Proxypolicy`.
func (o GetProxypolicylistResultOutput) Policyidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetProxypolicylistResult) []int { return v.Policyidlists }).(pulumi.IntArrayOutput)
}

func (o GetProxypolicylistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProxypolicylistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProxypolicylistResultOutput{})
}
