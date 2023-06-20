// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemAdmin`.
func GetSystemAdminlist(ctx *pulumi.Context, args *GetSystemAdminlistArgs, opts ...pulumi.InvokeOption) (*GetSystemAdminlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemAdminlistResult
	err := ctx.Invoke("fortios:index/getSystemAdminlist:getSystemAdminlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemAdminlist.
type GetSystemAdminlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemAdminlist.
type GetSystemAdminlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemAdmin`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemAdminlistOutput(ctx *pulumi.Context, args GetSystemAdminlistOutputArgs, opts ...pulumi.InvokeOption) GetSystemAdminlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemAdminlistResult, error) {
			args := v.(GetSystemAdminlistArgs)
			r, err := GetSystemAdminlist(ctx, &args, opts...)
			var s GetSystemAdminlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemAdminlistResultOutput)
}

// A collection of arguments for invoking getSystemAdminlist.
type GetSystemAdminlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemAdminlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAdminlistArgs)(nil)).Elem()
}

// A collection of values returned by getSystemAdminlist.
type GetSystemAdminlistResultOutput struct{ *pulumi.OutputState }

func (GetSystemAdminlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAdminlistResult)(nil)).Elem()
}

func (o GetSystemAdminlistResultOutput) ToGetSystemAdminlistResultOutput() GetSystemAdminlistResultOutput {
	return o
}

func (o GetSystemAdminlistResultOutput) ToGetSystemAdminlistResultOutputWithContext(ctx context.Context) GetSystemAdminlistResultOutput {
	return o
}

func (o GetSystemAdminlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAdminlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemAdminlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemAdminlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemAdmin`.
func (o GetSystemAdminlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemAdminlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemAdminlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAdminlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemAdminlistResultOutput{})
}