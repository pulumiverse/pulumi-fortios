// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dhcp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `system/dhcp.Server`.
func GetServerlist(ctx *pulumi.Context, args *GetServerlistArgs, opts ...pulumi.InvokeOption) (*GetServerlistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServerlistResult
	err := ctx.Invoke("fortios:system/dhcp/getServerlist:getServerlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerlist.
type GetServerlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getServerlist.
type GetServerlistResult struct {
	Filter *string `pulumi:"filter"`
	// A list of the `system/dhcp.Server`.
	Fosidlists []int `pulumi:"fosidlists"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetServerlistOutput(ctx *pulumi.Context, args GetServerlistOutputArgs, opts ...pulumi.InvokeOption) GetServerlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServerlistResult, error) {
			args := v.(GetServerlistArgs)
			r, err := GetServerlist(ctx, &args, opts...)
			var s GetServerlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServerlistResultOutput)
}

// A collection of arguments for invoking getServerlist.
type GetServerlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetServerlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerlistArgs)(nil)).Elem()
}

// A collection of values returned by getServerlist.
type GetServerlistResultOutput struct{ *pulumi.OutputState }

func (GetServerlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerlistResult)(nil)).Elem()
}

func (o GetServerlistResultOutput) ToGetServerlistResultOutput() GetServerlistResultOutput {
	return o
}

func (o GetServerlistResultOutput) ToGetServerlistResultOutputWithContext(ctx context.Context) GetServerlistResultOutput {
	return o
}

func (o GetServerlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// A list of the `system/dhcp.Server`.
func (o GetServerlistResultOutput) Fosidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetServerlistResult) []int { return v.Fosidlists }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServerlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerlistResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServerlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServerlistResultOutput{})
}
