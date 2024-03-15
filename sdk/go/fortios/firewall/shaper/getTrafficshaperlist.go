// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package shaper

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `firewall/shaper.Trafficshaper`.
func GetTrafficshaperlist(ctx *pulumi.Context, args *GetTrafficshaperlistArgs, opts ...pulumi.InvokeOption) (*GetTrafficshaperlistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTrafficshaperlistResult
	err := ctx.Invoke("fortios:firewall/shaper/getTrafficshaperlist:getTrafficshaperlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTrafficshaperlist.
type GetTrafficshaperlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getTrafficshaperlist.
type GetTrafficshaperlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall/shaper.Trafficshaper`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetTrafficshaperlistOutput(ctx *pulumi.Context, args GetTrafficshaperlistOutputArgs, opts ...pulumi.InvokeOption) GetTrafficshaperlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTrafficshaperlistResult, error) {
			args := v.(GetTrafficshaperlistArgs)
			r, err := GetTrafficshaperlist(ctx, &args, opts...)
			var s GetTrafficshaperlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTrafficshaperlistResultOutput)
}

// A collection of arguments for invoking getTrafficshaperlist.
type GetTrafficshaperlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetTrafficshaperlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTrafficshaperlistArgs)(nil)).Elem()
}

// A collection of values returned by getTrafficshaperlist.
type GetTrafficshaperlistResultOutput struct{ *pulumi.OutputState }

func (GetTrafficshaperlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTrafficshaperlistResult)(nil)).Elem()
}

func (o GetTrafficshaperlistResultOutput) ToGetTrafficshaperlistResultOutput() GetTrafficshaperlistResultOutput {
	return o
}

func (o GetTrafficshaperlistResultOutput) ToGetTrafficshaperlistResultOutputWithContext(ctx context.Context) GetTrafficshaperlistResultOutput {
	return o
}

func (o GetTrafficshaperlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficshaperlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTrafficshaperlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrafficshaperlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall/shaper.Trafficshaper`.
func (o GetTrafficshaperlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTrafficshaperlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetTrafficshaperlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficshaperlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTrafficshaperlistResultOutput{})
}