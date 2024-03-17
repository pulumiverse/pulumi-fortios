// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `system.Automationtrigger`.
func GetAutomationtriggerlist(ctx *pulumi.Context, args *GetAutomationtriggerlistArgs, opts ...pulumi.InvokeOption) (*GetAutomationtriggerlistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAutomationtriggerlistResult
	err := ctx.Invoke("fortios:system/getAutomationtriggerlist:getAutomationtriggerlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAutomationtriggerlist.
type GetAutomationtriggerlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getAutomationtriggerlist.
type GetAutomationtriggerlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `system.Automationtrigger`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetAutomationtriggerlistOutput(ctx *pulumi.Context, args GetAutomationtriggerlistOutputArgs, opts ...pulumi.InvokeOption) GetAutomationtriggerlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAutomationtriggerlistResult, error) {
			args := v.(GetAutomationtriggerlistArgs)
			r, err := GetAutomationtriggerlist(ctx, &args, opts...)
			var s GetAutomationtriggerlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAutomationtriggerlistResultOutput)
}

// A collection of arguments for invoking getAutomationtriggerlist.
type GetAutomationtriggerlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetAutomationtriggerlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAutomationtriggerlistArgs)(nil)).Elem()
}

// A collection of values returned by getAutomationtriggerlist.
type GetAutomationtriggerlistResultOutput struct{ *pulumi.OutputState }

func (GetAutomationtriggerlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAutomationtriggerlistResult)(nil)).Elem()
}

func (o GetAutomationtriggerlistResultOutput) ToGetAutomationtriggerlistResultOutput() GetAutomationtriggerlistResultOutput {
	return o
}

func (o GetAutomationtriggerlistResultOutput) ToGetAutomationtriggerlistResultOutputWithContext(ctx context.Context) GetAutomationtriggerlistResultOutput {
	return o
}

func (o GetAutomationtriggerlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAutomationtriggerlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAutomationtriggerlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAutomationtriggerlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `system.Automationtrigger`.
func (o GetAutomationtriggerlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAutomationtriggerlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetAutomationtriggerlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAutomationtriggerlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAutomationtriggerlistResultOutput{})
}
