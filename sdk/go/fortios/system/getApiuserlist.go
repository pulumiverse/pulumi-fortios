// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `system.Apiuser`.
func GetApiuserlist(ctx *pulumi.Context, args *GetApiuserlistArgs, opts ...pulumi.InvokeOption) (*GetApiuserlistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApiuserlistResult
	err := ctx.Invoke("fortios:system/getApiuserlist:getApiuserlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApiuserlist.
type GetApiuserlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getApiuserlist.
type GetApiuserlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `system.Apiuser`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetApiuserlistOutput(ctx *pulumi.Context, args GetApiuserlistOutputArgs, opts ...pulumi.InvokeOption) GetApiuserlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApiuserlistResult, error) {
			args := v.(GetApiuserlistArgs)
			r, err := GetApiuserlist(ctx, &args, opts...)
			var s GetApiuserlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetApiuserlistResultOutput)
}

// A collection of arguments for invoking getApiuserlist.
type GetApiuserlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetApiuserlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApiuserlistArgs)(nil)).Elem()
}

// A collection of values returned by getApiuserlist.
type GetApiuserlistResultOutput struct{ *pulumi.OutputState }

func (GetApiuserlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApiuserlistResult)(nil)).Elem()
}

func (o GetApiuserlistResultOutput) ToGetApiuserlistResultOutput() GetApiuserlistResultOutput {
	return o
}

func (o GetApiuserlistResultOutput) ToGetApiuserlistResultOutputWithContext(ctx context.Context) GetApiuserlistResultOutput {
	return o
}

func (o GetApiuserlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApiuserlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApiuserlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApiuserlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `system.Apiuser`.
func (o GetApiuserlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetApiuserlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetApiuserlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApiuserlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApiuserlistResultOutput{})
}
