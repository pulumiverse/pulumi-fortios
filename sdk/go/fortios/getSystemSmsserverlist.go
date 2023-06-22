// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemSmsserver`.
func GetSystemSmsserverlist(ctx *pulumi.Context, args *GetSystemSmsserverlistArgs, opts ...pulumi.InvokeOption) (*GetSystemSmsserverlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemSmsserverlistResult
	err := ctx.Invoke("fortios:index/getSystemSmsserverlist:getSystemSmsserverlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemSmsserverlist.
type GetSystemSmsserverlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemSmsserverlist.
type GetSystemSmsserverlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemSmsserver`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemSmsserverlistOutput(ctx *pulumi.Context, args GetSystemSmsserverlistOutputArgs, opts ...pulumi.InvokeOption) GetSystemSmsserverlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemSmsserverlistResult, error) {
			args := v.(GetSystemSmsserverlistArgs)
			r, err := GetSystemSmsserverlist(ctx, &args, opts...)
			var s GetSystemSmsserverlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemSmsserverlistResultOutput)
}

// A collection of arguments for invoking getSystemSmsserverlist.
type GetSystemSmsserverlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemSmsserverlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemSmsserverlistArgs)(nil)).Elem()
}

// A collection of values returned by getSystemSmsserverlist.
type GetSystemSmsserverlistResultOutput struct{ *pulumi.OutputState }

func (GetSystemSmsserverlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemSmsserverlistResult)(nil)).Elem()
}

func (o GetSystemSmsserverlistResultOutput) ToGetSystemSmsserverlistResultOutput() GetSystemSmsserverlistResultOutput {
	return o
}

func (o GetSystemSmsserverlistResultOutput) ToGetSystemSmsserverlistResultOutputWithContext(ctx context.Context) GetSystemSmsserverlistResultOutput {
	return o
}

func (o GetSystemSmsserverlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemSmsserverlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemSmsserverlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemSmsserverlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemSmsserver`.
func (o GetSystemSmsserverlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemSmsserverlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemSmsserverlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemSmsserverlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemSmsserverlistResultOutput{})
}
