// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `firewall.Multicastaddress6`.
func GetMulticastaddress6list(ctx *pulumi.Context, args *GetMulticastaddress6listArgs, opts ...pulumi.InvokeOption) (*GetMulticastaddress6listResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMulticastaddress6listResult
	err := ctx.Invoke("fortios:firewall/getMulticastaddress6list:getMulticastaddress6list", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMulticastaddress6list.
type GetMulticastaddress6listArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getMulticastaddress6list.
type GetMulticastaddress6listResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall.Multicastaddress6`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetMulticastaddress6listOutput(ctx *pulumi.Context, args GetMulticastaddress6listOutputArgs, opts ...pulumi.InvokeOption) GetMulticastaddress6listResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMulticastaddress6listResult, error) {
			args := v.(GetMulticastaddress6listArgs)
			r, err := GetMulticastaddress6list(ctx, &args, opts...)
			var s GetMulticastaddress6listResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMulticastaddress6listResultOutput)
}

// A collection of arguments for invoking getMulticastaddress6list.
type GetMulticastaddress6listOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetMulticastaddress6listOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMulticastaddress6listArgs)(nil)).Elem()
}

// A collection of values returned by getMulticastaddress6list.
type GetMulticastaddress6listResultOutput struct{ *pulumi.OutputState }

func (GetMulticastaddress6listResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMulticastaddress6listResult)(nil)).Elem()
}

func (o GetMulticastaddress6listResultOutput) ToGetMulticastaddress6listResultOutput() GetMulticastaddress6listResultOutput {
	return o
}

func (o GetMulticastaddress6listResultOutput) ToGetMulticastaddress6listResultOutputWithContext(ctx context.Context) GetMulticastaddress6listResultOutput {
	return o
}

func (o GetMulticastaddress6listResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMulticastaddress6listResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMulticastaddress6listResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMulticastaddress6listResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall.Multicastaddress6`.
func (o GetMulticastaddress6listResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMulticastaddress6listResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetMulticastaddress6listResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMulticastaddress6listResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMulticastaddress6listResultOutput{})
}
