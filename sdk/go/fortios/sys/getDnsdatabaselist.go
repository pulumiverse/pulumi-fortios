// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `sys.Dnsdatabase`.
func GetDnsdatabaselist(ctx *pulumi.Context, args *GetDnsdatabaselistArgs, opts ...pulumi.InvokeOption) (*GetDnsdatabaselistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDnsdatabaselistResult
	err := ctx.Invoke("fortios:sys/getDnsdatabaselist:getDnsdatabaselist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDnsdatabaselist.
type GetDnsdatabaselistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getDnsdatabaselist.
type GetDnsdatabaselistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `sys.Dnsdatabase`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetDnsdatabaselistOutput(ctx *pulumi.Context, args GetDnsdatabaselistOutputArgs, opts ...pulumi.InvokeOption) GetDnsdatabaselistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDnsdatabaselistResult, error) {
			args := v.(GetDnsdatabaselistArgs)
			r, err := GetDnsdatabaselist(ctx, &args, opts...)
			var s GetDnsdatabaselistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDnsdatabaselistResultOutput)
}

// A collection of arguments for invoking getDnsdatabaselist.
type GetDnsdatabaselistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetDnsdatabaselistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsdatabaselistArgs)(nil)).Elem()
}

// A collection of values returned by getDnsdatabaselist.
type GetDnsdatabaselistResultOutput struct{ *pulumi.OutputState }

func (GetDnsdatabaselistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsdatabaselistResult)(nil)).Elem()
}

func (o GetDnsdatabaselistResultOutput) ToGetDnsdatabaselistResultOutput() GetDnsdatabaselistResultOutput {
	return o
}

func (o GetDnsdatabaselistResultOutput) ToGetDnsdatabaselistResultOutputWithContext(ctx context.Context) GetDnsdatabaselistResultOutput {
	return o
}

func (o GetDnsdatabaselistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDnsdatabaselistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDnsdatabaselistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsdatabaselistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `sys.Dnsdatabase`.
func (o GetDnsdatabaselistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDnsdatabaselistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetDnsdatabaselistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDnsdatabaselistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDnsdatabaselistResultOutput{})
}
