// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `system.Interface`.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.GetInterfacelist(ctx, &system.GetInterfacelistArgs{
//				Filter: pulumi.StringRef("name!=port1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("output1", data.Fortios_system_interfacelist.Sample2.Namelist)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetInterfacelist(ctx *pulumi.Context, args *GetInterfacelistArgs, opts ...pulumi.InvokeOption) (*GetInterfacelistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInterfacelistResult
	err := ctx.Invoke("fortios:system/getInterfacelist:getInterfacelist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInterfacelist.
type GetInterfacelistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getInterfacelist.
type GetInterfacelistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `system.Interface`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetInterfacelistOutput(ctx *pulumi.Context, args GetInterfacelistOutputArgs, opts ...pulumi.InvokeOption) GetInterfacelistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInterfacelistResult, error) {
			args := v.(GetInterfacelistArgs)
			r, err := GetInterfacelist(ctx, &args, opts...)
			var s GetInterfacelistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInterfacelistResultOutput)
}

// A collection of arguments for invoking getInterfacelist.
type GetInterfacelistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetInterfacelistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInterfacelistArgs)(nil)).Elem()
}

// A collection of values returned by getInterfacelist.
type GetInterfacelistResultOutput struct{ *pulumi.OutputState }

func (GetInterfacelistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInterfacelistResult)(nil)).Elem()
}

func (o GetInterfacelistResultOutput) ToGetInterfacelistResultOutput() GetInterfacelistResultOutput {
	return o
}

func (o GetInterfacelistResultOutput) ToGetInterfacelistResultOutputWithContext(ctx context.Context) GetInterfacelistResultOutput {
	return o
}

func (o GetInterfacelistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInterfacelistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInterfacelistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInterfacelistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `system.Interface`.
func (o GetInterfacelistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInterfacelistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetInterfacelistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInterfacelistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInterfacelistResultOutput{})
}
