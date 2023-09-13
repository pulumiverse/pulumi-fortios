// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `sys.Tosbasedpriority`.
func GetTosbasedprioritylist(ctx *pulumi.Context, args *GetTosbasedprioritylistArgs, opts ...pulumi.InvokeOption) (*GetTosbasedprioritylistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetTosbasedprioritylistResult
	err := ctx.Invoke("fortios:sys/getTosbasedprioritylist:getTosbasedprioritylist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTosbasedprioritylist.
type GetTosbasedprioritylistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getTosbasedprioritylist.
type GetTosbasedprioritylistResult struct {
	Filter *string `pulumi:"filter"`
	// A list of the `sys.Tosbasedpriority`.
	Fosidlists []int `pulumi:"fosidlists"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetTosbasedprioritylistOutput(ctx *pulumi.Context, args GetTosbasedprioritylistOutputArgs, opts ...pulumi.InvokeOption) GetTosbasedprioritylistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTosbasedprioritylistResult, error) {
			args := v.(GetTosbasedprioritylistArgs)
			r, err := GetTosbasedprioritylist(ctx, &args, opts...)
			var s GetTosbasedprioritylistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTosbasedprioritylistResultOutput)
}

// A collection of arguments for invoking getTosbasedprioritylist.
type GetTosbasedprioritylistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetTosbasedprioritylistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTosbasedprioritylistArgs)(nil)).Elem()
}

// A collection of values returned by getTosbasedprioritylist.
type GetTosbasedprioritylistResultOutput struct{ *pulumi.OutputState }

func (GetTosbasedprioritylistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTosbasedprioritylistResult)(nil)).Elem()
}

func (o GetTosbasedprioritylistResultOutput) ToGetTosbasedprioritylistResultOutput() GetTosbasedprioritylistResultOutput {
	return o
}

func (o GetTosbasedprioritylistResultOutput) ToGetTosbasedprioritylistResultOutputWithContext(ctx context.Context) GetTosbasedprioritylistResultOutput {
	return o
}

func (o GetTosbasedprioritylistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTosbasedprioritylistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// A list of the `sys.Tosbasedpriority`.
func (o GetTosbasedprioritylistResultOutput) Fosidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetTosbasedprioritylistResult) []int { return v.Fosidlists }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTosbasedprioritylistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTosbasedprioritylistResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTosbasedprioritylistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTosbasedprioritylistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTosbasedprioritylistResultOutput{})
}
