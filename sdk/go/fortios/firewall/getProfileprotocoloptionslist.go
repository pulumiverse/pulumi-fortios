// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `firewall.Profileprotocoloptions`.
func GetProfileprotocoloptionslist(ctx *pulumi.Context, args *GetProfileprotocoloptionslistArgs, opts ...pulumi.InvokeOption) (*GetProfileprotocoloptionslistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetProfileprotocoloptionslistResult
	err := ctx.Invoke("fortios:firewall/getProfileprotocoloptionslist:getProfileprotocoloptionslist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProfileprotocoloptionslist.
type GetProfileprotocoloptionslistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getProfileprotocoloptionslist.
type GetProfileprotocoloptionslistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall.Profileprotocoloptions`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetProfileprotocoloptionslistOutput(ctx *pulumi.Context, args GetProfileprotocoloptionslistOutputArgs, opts ...pulumi.InvokeOption) GetProfileprotocoloptionslistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProfileprotocoloptionslistResult, error) {
			args := v.(GetProfileprotocoloptionslistArgs)
			r, err := GetProfileprotocoloptionslist(ctx, &args, opts...)
			var s GetProfileprotocoloptionslistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProfileprotocoloptionslistResultOutput)
}

// A collection of arguments for invoking getProfileprotocoloptionslist.
type GetProfileprotocoloptionslistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetProfileprotocoloptionslistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProfileprotocoloptionslistArgs)(nil)).Elem()
}

// A collection of values returned by getProfileprotocoloptionslist.
type GetProfileprotocoloptionslistResultOutput struct{ *pulumi.OutputState }

func (GetProfileprotocoloptionslistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProfileprotocoloptionslistResult)(nil)).Elem()
}

func (o GetProfileprotocoloptionslistResultOutput) ToGetProfileprotocoloptionslistResultOutput() GetProfileprotocoloptionslistResultOutput {
	return o
}

func (o GetProfileprotocoloptionslistResultOutput) ToGetProfileprotocoloptionslistResultOutputWithContext(ctx context.Context) GetProfileprotocoloptionslistResultOutput {
	return o
}

func (o GetProfileprotocoloptionslistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProfileprotocoloptionslistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProfileprotocoloptionslistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProfileprotocoloptionslistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall.Profileprotocoloptions`.
func (o GetProfileprotocoloptionslistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProfileprotocoloptionslistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetProfileprotocoloptionslistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProfileprotocoloptionslistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProfileprotocoloptionslistResultOutput{})
}
