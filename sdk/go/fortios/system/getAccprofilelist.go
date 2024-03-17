// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `system.Accprofile`.
func GetAccprofilelist(ctx *pulumi.Context, args *GetAccprofilelistArgs, opts ...pulumi.InvokeOption) (*GetAccprofilelistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccprofilelistResult
	err := ctx.Invoke("fortios:system/getAccprofilelist:getAccprofilelist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccprofilelist.
type GetAccprofilelistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getAccprofilelist.
type GetAccprofilelistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `system.Accprofile`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetAccprofilelistOutput(ctx *pulumi.Context, args GetAccprofilelistOutputArgs, opts ...pulumi.InvokeOption) GetAccprofilelistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccprofilelistResult, error) {
			args := v.(GetAccprofilelistArgs)
			r, err := GetAccprofilelist(ctx, &args, opts...)
			var s GetAccprofilelistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccprofilelistResultOutput)
}

// A collection of arguments for invoking getAccprofilelist.
type GetAccprofilelistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetAccprofilelistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccprofilelistArgs)(nil)).Elem()
}

// A collection of values returned by getAccprofilelist.
type GetAccprofilelistResultOutput struct{ *pulumi.OutputState }

func (GetAccprofilelistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccprofilelistResult)(nil)).Elem()
}

func (o GetAccprofilelistResultOutput) ToGetAccprofilelistResultOutput() GetAccprofilelistResultOutput {
	return o
}

func (o GetAccprofilelistResultOutput) ToGetAccprofilelistResultOutputWithContext(ctx context.Context) GetAccprofilelistResultOutput {
	return o
}

func (o GetAccprofilelistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccprofilelistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccprofilelistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccprofilelistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `system.Accprofile`.
func (o GetAccprofilelistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccprofilelistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetAccprofilelistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccprofilelistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccprofilelistResultOutput{})
}
