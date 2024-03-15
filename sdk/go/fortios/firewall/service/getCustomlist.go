// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package service

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `firewall/service.Custom`.
func GetCustomlist(ctx *pulumi.Context, args *GetCustomlistArgs, opts ...pulumi.InvokeOption) (*GetCustomlistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCustomlistResult
	err := ctx.Invoke("fortios:firewall/service/getCustomlist:getCustomlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomlist.
type GetCustomlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getCustomlist.
type GetCustomlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall/service.Custom`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetCustomlistOutput(ctx *pulumi.Context, args GetCustomlistOutputArgs, opts ...pulumi.InvokeOption) GetCustomlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCustomlistResult, error) {
			args := v.(GetCustomlistArgs)
			r, err := GetCustomlist(ctx, &args, opts...)
			var s GetCustomlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCustomlistResultOutput)
}

// A collection of arguments for invoking getCustomlist.
type GetCustomlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetCustomlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomlistArgs)(nil)).Elem()
}

// A collection of values returned by getCustomlist.
type GetCustomlistResultOutput struct{ *pulumi.OutputState }

func (GetCustomlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomlistResult)(nil)).Elem()
}

func (o GetCustomlistResultOutput) ToGetCustomlistResultOutput() GetCustomlistResultOutput {
	return o
}

func (o GetCustomlistResultOutput) ToGetCustomlistResultOutputWithContext(ctx context.Context) GetCustomlistResultOutput {
	return o
}

func (o GetCustomlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCustomlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall/service.Custom`.
func (o GetCustomlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCustomlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetCustomlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCustomlistResultOutput{})
}