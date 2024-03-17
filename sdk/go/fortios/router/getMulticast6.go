// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on fortios router multicast6
func LookupMulticast6(ctx *pulumi.Context, args *LookupMulticast6Args, opts ...pulumi.InvokeOption) (*LookupMulticast6Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMulticast6Result
	err := ctx.Invoke("fortios:router/getMulticast6:getMulticast6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMulticast6.
type LookupMulticast6Args struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getMulticast6.
type LookupMulticast6Result struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
	Interfaces []GetMulticast6Interface `pulumi:"interfaces"`
	// Enable/disable PMTU for IPv6 multicast.
	MulticastPmtu string `pulumi:"multicastPmtu"`
	// Enable/disable IPv6 multicast routing.
	MulticastRouting string `pulumi:"multicastRouting"`
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobals []GetMulticast6PimSmGlobal `pulumi:"pimSmGlobals"`
	Vdomparam    *string                    `pulumi:"vdomparam"`
}

func LookupMulticast6Output(ctx *pulumi.Context, args LookupMulticast6OutputArgs, opts ...pulumi.InvokeOption) LookupMulticast6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMulticast6Result, error) {
			args := v.(LookupMulticast6Args)
			r, err := LookupMulticast6(ctx, &args, opts...)
			var s LookupMulticast6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMulticast6ResultOutput)
}

// A collection of arguments for invoking getMulticast6.
type LookupMulticast6OutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupMulticast6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMulticast6Args)(nil)).Elem()
}

// A collection of values returned by getMulticast6.
type LookupMulticast6ResultOutput struct{ *pulumi.OutputState }

func (LookupMulticast6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMulticast6Result)(nil)).Elem()
}

func (o LookupMulticast6ResultOutput) ToLookupMulticast6ResultOutput() LookupMulticast6ResultOutput {
	return o
}

func (o LookupMulticast6ResultOutput) ToLookupMulticast6ResultOutputWithContext(ctx context.Context) LookupMulticast6ResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMulticast6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMulticast6Result) string { return v.Id }).(pulumi.StringOutput)
}

// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
func (o LookupMulticast6ResultOutput) Interfaces() GetMulticast6InterfaceArrayOutput {
	return o.ApplyT(func(v LookupMulticast6Result) []GetMulticast6Interface { return v.Interfaces }).(GetMulticast6InterfaceArrayOutput)
}

// Enable/disable PMTU for IPv6 multicast.
func (o LookupMulticast6ResultOutput) MulticastPmtu() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMulticast6Result) string { return v.MulticastPmtu }).(pulumi.StringOutput)
}

// Enable/disable IPv6 multicast routing.
func (o LookupMulticast6ResultOutput) MulticastRouting() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMulticast6Result) string { return v.MulticastRouting }).(pulumi.StringOutput)
}

// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
func (o LookupMulticast6ResultOutput) PimSmGlobals() GetMulticast6PimSmGlobalArrayOutput {
	return o.ApplyT(func(v LookupMulticast6Result) []GetMulticast6PimSmGlobal { return v.PimSmGlobals }).(GetMulticast6PimSmGlobalArrayOutput)
}

func (o LookupMulticast6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMulticast6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMulticast6ResultOutput{})
}
