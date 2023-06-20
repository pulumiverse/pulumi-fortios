// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPS URL filter settings for IPv6.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewWebfilterIpsurlfiltersetting6(ctx, "trname", &fortios.WebfilterIpsurlfiltersetting6Args{
//				Distance: pulumi.Int(1),
//				Gateway6: pulumi.String("::"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # Webfilter IpsUrlfilterSetting6 can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/webfilterIpsurlfiltersetting6:WebfilterIpsurlfiltersetting6 labelname WebfilterIpsUrlfilterSetting6
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/webfilterIpsurlfiltersetting6:WebfilterIpsurlfiltersetting6 labelname WebfilterIpsUrlfilterSetting6
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WebfilterIpsurlfiltersetting6 struct {
	pulumi.CustomResourceState

	// Interface for this route.
	Device pulumi.StringOutput `pulumi:"device"`
	// Administrative distance (1 - 255) for this route.
	Distance pulumi.IntOutput `pulumi:"distance"`
	// Gateway IPv6 address for this route.
	Gateway6 pulumi.StringOutput `pulumi:"gateway6"`
	// Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
	GeoFilter pulumi.StringPtrOutput `pulumi:"geoFilter"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWebfilterIpsurlfiltersetting6 registers a new resource with the given unique name, arguments, and options.
func NewWebfilterIpsurlfiltersetting6(ctx *pulumi.Context,
	name string, args *WebfilterIpsurlfiltersetting6Args, opts ...pulumi.ResourceOption) (*WebfilterIpsurlfiltersetting6, error) {
	if args == nil {
		args = &WebfilterIpsurlfiltersetting6Args{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WebfilterIpsurlfiltersetting6
	err := ctx.RegisterResource("fortios:index/webfilterIpsurlfiltersetting6:WebfilterIpsurlfiltersetting6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebfilterIpsurlfiltersetting6 gets an existing WebfilterIpsurlfiltersetting6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebfilterIpsurlfiltersetting6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebfilterIpsurlfiltersetting6State, opts ...pulumi.ResourceOption) (*WebfilterIpsurlfiltersetting6, error) {
	var resource WebfilterIpsurlfiltersetting6
	err := ctx.ReadResource("fortios:index/webfilterIpsurlfiltersetting6:WebfilterIpsurlfiltersetting6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebfilterIpsurlfiltersetting6 resources.
type webfilterIpsurlfiltersetting6State struct {
	// Interface for this route.
	Device *string `pulumi:"device"`
	// Administrative distance (1 - 255) for this route.
	Distance *int `pulumi:"distance"`
	// Gateway IPv6 address for this route.
	Gateway6 *string `pulumi:"gateway6"`
	// Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
	GeoFilter *string `pulumi:"geoFilter"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WebfilterIpsurlfiltersetting6State struct {
	// Interface for this route.
	Device pulumi.StringPtrInput
	// Administrative distance (1 - 255) for this route.
	Distance pulumi.IntPtrInput
	// Gateway IPv6 address for this route.
	Gateway6 pulumi.StringPtrInput
	// Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
	GeoFilter pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebfilterIpsurlfiltersetting6State) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterIpsurlfiltersetting6State)(nil)).Elem()
}

type webfilterIpsurlfiltersetting6Args struct {
	// Interface for this route.
	Device *string `pulumi:"device"`
	// Administrative distance (1 - 255) for this route.
	Distance *int `pulumi:"distance"`
	// Gateway IPv6 address for this route.
	Gateway6 *string `pulumi:"gateway6"`
	// Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
	GeoFilter *string `pulumi:"geoFilter"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebfilterIpsurlfiltersetting6 resource.
type WebfilterIpsurlfiltersetting6Args struct {
	// Interface for this route.
	Device pulumi.StringPtrInput
	// Administrative distance (1 - 255) for this route.
	Distance pulumi.IntPtrInput
	// Gateway IPv6 address for this route.
	Gateway6 pulumi.StringPtrInput
	// Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
	GeoFilter pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebfilterIpsurlfiltersetting6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterIpsurlfiltersetting6Args)(nil)).Elem()
}

type WebfilterIpsurlfiltersetting6Input interface {
	pulumi.Input

	ToWebfilterIpsurlfiltersetting6Output() WebfilterIpsurlfiltersetting6Output
	ToWebfilterIpsurlfiltersetting6OutputWithContext(ctx context.Context) WebfilterIpsurlfiltersetting6Output
}

func (*WebfilterIpsurlfiltersetting6) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterIpsurlfiltersetting6)(nil)).Elem()
}

func (i *WebfilterIpsurlfiltersetting6) ToWebfilterIpsurlfiltersetting6Output() WebfilterIpsurlfiltersetting6Output {
	return i.ToWebfilterIpsurlfiltersetting6OutputWithContext(context.Background())
}

func (i *WebfilterIpsurlfiltersetting6) ToWebfilterIpsurlfiltersetting6OutputWithContext(ctx context.Context) WebfilterIpsurlfiltersetting6Output {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterIpsurlfiltersetting6Output)
}

// WebfilterIpsurlfiltersetting6ArrayInput is an input type that accepts WebfilterIpsurlfiltersetting6Array and WebfilterIpsurlfiltersetting6ArrayOutput values.
// You can construct a concrete instance of `WebfilterIpsurlfiltersetting6ArrayInput` via:
//
//	WebfilterIpsurlfiltersetting6Array{ WebfilterIpsurlfiltersetting6Args{...} }
type WebfilterIpsurlfiltersetting6ArrayInput interface {
	pulumi.Input

	ToWebfilterIpsurlfiltersetting6ArrayOutput() WebfilterIpsurlfiltersetting6ArrayOutput
	ToWebfilterIpsurlfiltersetting6ArrayOutputWithContext(context.Context) WebfilterIpsurlfiltersetting6ArrayOutput
}

type WebfilterIpsurlfiltersetting6Array []WebfilterIpsurlfiltersetting6Input

func (WebfilterIpsurlfiltersetting6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebfilterIpsurlfiltersetting6)(nil)).Elem()
}

func (i WebfilterIpsurlfiltersetting6Array) ToWebfilterIpsurlfiltersetting6ArrayOutput() WebfilterIpsurlfiltersetting6ArrayOutput {
	return i.ToWebfilterIpsurlfiltersetting6ArrayOutputWithContext(context.Background())
}

func (i WebfilterIpsurlfiltersetting6Array) ToWebfilterIpsurlfiltersetting6ArrayOutputWithContext(ctx context.Context) WebfilterIpsurlfiltersetting6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterIpsurlfiltersetting6ArrayOutput)
}

// WebfilterIpsurlfiltersetting6MapInput is an input type that accepts WebfilterIpsurlfiltersetting6Map and WebfilterIpsurlfiltersetting6MapOutput values.
// You can construct a concrete instance of `WebfilterIpsurlfiltersetting6MapInput` via:
//
//	WebfilterIpsurlfiltersetting6Map{ "key": WebfilterIpsurlfiltersetting6Args{...} }
type WebfilterIpsurlfiltersetting6MapInput interface {
	pulumi.Input

	ToWebfilterIpsurlfiltersetting6MapOutput() WebfilterIpsurlfiltersetting6MapOutput
	ToWebfilterIpsurlfiltersetting6MapOutputWithContext(context.Context) WebfilterIpsurlfiltersetting6MapOutput
}

type WebfilterIpsurlfiltersetting6Map map[string]WebfilterIpsurlfiltersetting6Input

func (WebfilterIpsurlfiltersetting6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebfilterIpsurlfiltersetting6)(nil)).Elem()
}

func (i WebfilterIpsurlfiltersetting6Map) ToWebfilterIpsurlfiltersetting6MapOutput() WebfilterIpsurlfiltersetting6MapOutput {
	return i.ToWebfilterIpsurlfiltersetting6MapOutputWithContext(context.Background())
}

func (i WebfilterIpsurlfiltersetting6Map) ToWebfilterIpsurlfiltersetting6MapOutputWithContext(ctx context.Context) WebfilterIpsurlfiltersetting6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterIpsurlfiltersetting6MapOutput)
}

type WebfilterIpsurlfiltersetting6Output struct{ *pulumi.OutputState }

func (WebfilterIpsurlfiltersetting6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterIpsurlfiltersetting6)(nil)).Elem()
}

func (o WebfilterIpsurlfiltersetting6Output) ToWebfilterIpsurlfiltersetting6Output() WebfilterIpsurlfiltersetting6Output {
	return o
}

func (o WebfilterIpsurlfiltersetting6Output) ToWebfilterIpsurlfiltersetting6OutputWithContext(ctx context.Context) WebfilterIpsurlfiltersetting6Output {
	return o
}

// Interface for this route.
func (o WebfilterIpsurlfiltersetting6Output) Device() pulumi.StringOutput {
	return o.ApplyT(func(v *WebfilterIpsurlfiltersetting6) pulumi.StringOutput { return v.Device }).(pulumi.StringOutput)
}

// Administrative distance (1 - 255) for this route.
func (o WebfilterIpsurlfiltersetting6Output) Distance() pulumi.IntOutput {
	return o.ApplyT(func(v *WebfilterIpsurlfiltersetting6) pulumi.IntOutput { return v.Distance }).(pulumi.IntOutput)
}

// Gateway IPv6 address for this route.
func (o WebfilterIpsurlfiltersetting6Output) Gateway6() pulumi.StringOutput {
	return o.ApplyT(func(v *WebfilterIpsurlfiltersetting6) pulumi.StringOutput { return v.Gateway6 }).(pulumi.StringOutput)
}

// Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
func (o WebfilterIpsurlfiltersetting6Output) GeoFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebfilterIpsurlfiltersetting6) pulumi.StringPtrOutput { return v.GeoFilter }).(pulumi.StringPtrOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WebfilterIpsurlfiltersetting6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebfilterIpsurlfiltersetting6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WebfilterIpsurlfiltersetting6ArrayOutput struct{ *pulumi.OutputState }

func (WebfilterIpsurlfiltersetting6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebfilterIpsurlfiltersetting6)(nil)).Elem()
}

func (o WebfilterIpsurlfiltersetting6ArrayOutput) ToWebfilterIpsurlfiltersetting6ArrayOutput() WebfilterIpsurlfiltersetting6ArrayOutput {
	return o
}

func (o WebfilterIpsurlfiltersetting6ArrayOutput) ToWebfilterIpsurlfiltersetting6ArrayOutputWithContext(ctx context.Context) WebfilterIpsurlfiltersetting6ArrayOutput {
	return o
}

func (o WebfilterIpsurlfiltersetting6ArrayOutput) Index(i pulumi.IntInput) WebfilterIpsurlfiltersetting6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebfilterIpsurlfiltersetting6 {
		return vs[0].([]*WebfilterIpsurlfiltersetting6)[vs[1].(int)]
	}).(WebfilterIpsurlfiltersetting6Output)
}

type WebfilterIpsurlfiltersetting6MapOutput struct{ *pulumi.OutputState }

func (WebfilterIpsurlfiltersetting6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebfilterIpsurlfiltersetting6)(nil)).Elem()
}

func (o WebfilterIpsurlfiltersetting6MapOutput) ToWebfilterIpsurlfiltersetting6MapOutput() WebfilterIpsurlfiltersetting6MapOutput {
	return o
}

func (o WebfilterIpsurlfiltersetting6MapOutput) ToWebfilterIpsurlfiltersetting6MapOutputWithContext(ctx context.Context) WebfilterIpsurlfiltersetting6MapOutput {
	return o
}

func (o WebfilterIpsurlfiltersetting6MapOutput) MapIndex(k pulumi.StringInput) WebfilterIpsurlfiltersetting6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebfilterIpsurlfiltersetting6 {
		return vs[0].(map[string]*WebfilterIpsurlfiltersetting6)[vs[1].(string)]
	}).(WebfilterIpsurlfiltersetting6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterIpsurlfiltersetting6Input)(nil)).Elem(), &WebfilterIpsurlfiltersetting6{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterIpsurlfiltersetting6ArrayInput)(nil)).Elem(), WebfilterIpsurlfiltersetting6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterIpsurlfiltersetting6MapInput)(nil)).Elem(), WebfilterIpsurlfiltersetting6Map{})
	pulumi.RegisterOutputType(WebfilterIpsurlfiltersetting6Output{})
	pulumi.RegisterOutputType(WebfilterIpsurlfiltersetting6ArrayOutput{})
	pulumi.RegisterOutputType(WebfilterIpsurlfiltersetting6MapOutput{})
}