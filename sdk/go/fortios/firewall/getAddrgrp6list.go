// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `firewall.Addrgrp6`.
func GetAddrgrp6list(ctx *pulumi.Context, args *GetAddrgrp6listArgs, opts ...pulumi.InvokeOption) (*GetAddrgrp6listResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAddrgrp6listResult
	err := ctx.Invoke("fortios:firewall/getAddrgrp6list:getAddrgrp6list", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAddrgrp6list.
type GetAddrgrp6listArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getAddrgrp6list.
type GetAddrgrp6listResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall.Addrgrp6`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetAddrgrp6listOutput(ctx *pulumi.Context, args GetAddrgrp6listOutputArgs, opts ...pulumi.InvokeOption) GetAddrgrp6listResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAddrgrp6listResult, error) {
			args := v.(GetAddrgrp6listArgs)
			r, err := GetAddrgrp6list(ctx, &args, opts...)
			var s GetAddrgrp6listResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAddrgrp6listResultOutput)
}

// A collection of arguments for invoking getAddrgrp6list.
type GetAddrgrp6listOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetAddrgrp6listOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddrgrp6listArgs)(nil)).Elem()
}

// A collection of values returned by getAddrgrp6list.
type GetAddrgrp6listResultOutput struct{ *pulumi.OutputState }

func (GetAddrgrp6listResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddrgrp6listResult)(nil)).Elem()
}

func (o GetAddrgrp6listResultOutput) ToGetAddrgrp6listResultOutput() GetAddrgrp6listResultOutput {
	return o
}

func (o GetAddrgrp6listResultOutput) ToGetAddrgrp6listResultOutputWithContext(ctx context.Context) GetAddrgrp6listResultOutput {
	return o
}

func (o GetAddrgrp6listResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddrgrp6listResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAddrgrp6listResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAddrgrp6listResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall.Addrgrp6`.
func (o GetAddrgrp6listResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAddrgrp6listResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetAddrgrp6listResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddrgrp6listResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAddrgrp6listResultOutput{})
}
