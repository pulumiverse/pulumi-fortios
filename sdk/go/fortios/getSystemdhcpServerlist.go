// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemdhcpServer`.
func GetSystemdhcpServerlist(ctx *pulumi.Context, args *GetSystemdhcpServerlistArgs, opts ...pulumi.InvokeOption) (*GetSystemdhcpServerlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemdhcpServerlistResult
	err := ctx.Invoke("fortios:index/getSystemdhcpServerlist:getSystemdhcpServerlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemdhcpServerlist.
type GetSystemdhcpServerlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemdhcpServerlist.
type GetSystemdhcpServerlistResult struct {
	Filter *string `pulumi:"filter"`
	// A list of the `SystemdhcpServer`.
	Fosidlists []int `pulumi:"fosidlists"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetSystemdhcpServerlistOutput(ctx *pulumi.Context, args GetSystemdhcpServerlistOutputArgs, opts ...pulumi.InvokeOption) GetSystemdhcpServerlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemdhcpServerlistResult, error) {
			args := v.(GetSystemdhcpServerlistArgs)
			r, err := GetSystemdhcpServerlist(ctx, &args, opts...)
			var s GetSystemdhcpServerlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemdhcpServerlistResultOutput)
}

// A collection of arguments for invoking getSystemdhcpServerlist.
type GetSystemdhcpServerlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemdhcpServerlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemdhcpServerlistArgs)(nil)).Elem()
}

// A collection of values returned by getSystemdhcpServerlist.
type GetSystemdhcpServerlistResultOutput struct{ *pulumi.OutputState }

func (GetSystemdhcpServerlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemdhcpServerlistResult)(nil)).Elem()
}

func (o GetSystemdhcpServerlistResultOutput) ToGetSystemdhcpServerlistResultOutput() GetSystemdhcpServerlistResultOutput {
	return o
}

func (o GetSystemdhcpServerlistResultOutput) ToGetSystemdhcpServerlistResultOutputWithContext(ctx context.Context) GetSystemdhcpServerlistResultOutput {
	return o
}

func (o GetSystemdhcpServerlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemdhcpServerlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// A list of the `SystemdhcpServer`.
func (o GetSystemdhcpServerlistResultOutput) Fosidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetSystemdhcpServerlistResult) []int { return v.Fosidlists }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemdhcpServerlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemdhcpServerlistResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemdhcpServerlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemdhcpServerlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemdhcpServerlistResultOutput{})
}
