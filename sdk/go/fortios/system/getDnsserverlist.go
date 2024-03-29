// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `system.Dnsserver`.
func GetDnsserverlist(ctx *pulumi.Context, args *GetDnsserverlistArgs, opts ...pulumi.InvokeOption) (*GetDnsserverlistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDnsserverlistResult
	err := ctx.Invoke("fortios:system/getDnsserverlist:getDnsserverlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDnsserverlist.
type GetDnsserverlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getDnsserverlist.
type GetDnsserverlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `system.Dnsserver`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetDnsserverlistOutput(ctx *pulumi.Context, args GetDnsserverlistOutputArgs, opts ...pulumi.InvokeOption) GetDnsserverlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDnsserverlistResult, error) {
			args := v.(GetDnsserverlistArgs)
			r, err := GetDnsserverlist(ctx, &args, opts...)
			var s GetDnsserverlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDnsserverlistResultOutput)
}

// A collection of arguments for invoking getDnsserverlist.
type GetDnsserverlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetDnsserverlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsserverlistArgs)(nil)).Elem()
}

// A collection of values returned by getDnsserverlist.
type GetDnsserverlistResultOutput struct{ *pulumi.OutputState }

func (GetDnsserverlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsserverlistResult)(nil)).Elem()
}

func (o GetDnsserverlistResultOutput) ToGetDnsserverlistResultOutput() GetDnsserverlistResultOutput {
	return o
}

func (o GetDnsserverlistResultOutput) ToGetDnsserverlistResultOutputWithContext(ctx context.Context) GetDnsserverlistResultOutput {
	return o
}

func (o GetDnsserverlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDnsserverlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDnsserverlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsserverlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `system.Dnsserver`.
func (o GetDnsserverlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDnsserverlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetDnsserverlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDnsserverlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDnsserverlistResultOutput{})
}
