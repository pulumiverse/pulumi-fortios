// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios router bfd6
func LookupRouterBfd6(ctx *pulumi.Context, args *LookupRouterBfd6Args, opts ...pulumi.InvokeOption) (*LookupRouterBfd6Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterBfd6Result
	err := ctx.Invoke("fortios:index/getRouterBfd6:getRouterBfd6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouterBfd6.
type LookupRouterBfd6Args struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRouterBfd6.
type LookupRouterBfd6Result struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// BFD IPv6 multi-hop template table. The structure of `multihopTemplate` block is documented below.
	MultihopTemplates []GetRouterBfd6MultihopTemplate `pulumi:"multihopTemplates"`
	// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
	Neighbors []GetRouterBfd6Neighbor `pulumi:"neighbors"`
	Vdomparam *string                 `pulumi:"vdomparam"`
}

func LookupRouterBfd6Output(ctx *pulumi.Context, args LookupRouterBfd6OutputArgs, opts ...pulumi.InvokeOption) LookupRouterBfd6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterBfd6Result, error) {
			args := v.(LookupRouterBfd6Args)
			r, err := LookupRouterBfd6(ctx, &args, opts...)
			var s LookupRouterBfd6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterBfd6ResultOutput)
}

// A collection of arguments for invoking getRouterBfd6.
type LookupRouterBfd6OutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterBfd6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterBfd6Args)(nil)).Elem()
}

// A collection of values returned by getRouterBfd6.
type LookupRouterBfd6ResultOutput struct{ *pulumi.OutputState }

func (LookupRouterBfd6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterBfd6Result)(nil)).Elem()
}

func (o LookupRouterBfd6ResultOutput) ToLookupRouterBfd6ResultOutput() LookupRouterBfd6ResultOutput {
	return o
}

func (o LookupRouterBfd6ResultOutput) ToLookupRouterBfd6ResultOutputWithContext(ctx context.Context) LookupRouterBfd6ResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterBfd6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBfd6Result) string { return v.Id }).(pulumi.StringOutput)
}

// BFD IPv6 multi-hop template table. The structure of `multihopTemplate` block is documented below.
func (o LookupRouterBfd6ResultOutput) MultihopTemplates() GetRouterBfd6MultihopTemplateArrayOutput {
	return o.ApplyT(func(v LookupRouterBfd6Result) []GetRouterBfd6MultihopTemplate { return v.MultihopTemplates }).(GetRouterBfd6MultihopTemplateArrayOutput)
}

// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
func (o LookupRouterBfd6ResultOutput) Neighbors() GetRouterBfd6NeighborArrayOutput {
	return o.ApplyT(func(v LookupRouterBfd6Result) []GetRouterBfd6Neighbor { return v.Neighbors }).(GetRouterBfd6NeighborArrayOutput)
}

func (o LookupRouterBfd6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterBfd6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterBfd6ResultOutput{})
}
