// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hotspot20

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IP address type availability.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/wirelesscontroller"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := wirelesscontroller.NewAnqpipaddresstype(ctx, "trname", &wirelesscontroller.AnqpipaddresstypeArgs{
//				Ipv4AddressType: pulumi.String("public"),
//				Ipv6AddressType: pulumi.String("not-available"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// WirelessControllerHotspot20 AnqpIpAddressType can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/hotspot20/anqpipaddresstype:Anqpipaddresstype labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/hotspot20/anqpipaddresstype:Anqpipaddresstype labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Anqpipaddresstype struct {
	pulumi.CustomResourceState

	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType pulumi.StringOutput `pulumi:"ipv4AddressType"`
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType pulumi.StringOutput `pulumi:"ipv6AddressType"`
	// IP type name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewAnqpipaddresstype registers a new resource with the given unique name, arguments, and options.
func NewAnqpipaddresstype(ctx *pulumi.Context,
	name string, args *AnqpipaddresstypeArgs, opts ...pulumi.ResourceOption) (*Anqpipaddresstype, error) {
	if args == nil {
		args = &AnqpipaddresstypeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Anqpipaddresstype
	err := ctx.RegisterResource("fortios:wirelesscontroller/hotspot20/anqpipaddresstype:Anqpipaddresstype", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnqpipaddresstype gets an existing Anqpipaddresstype resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnqpipaddresstype(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnqpipaddresstypeState, opts ...pulumi.ResourceOption) (*Anqpipaddresstype, error) {
	var resource Anqpipaddresstype
	err := ctx.ReadResource("fortios:wirelesscontroller/hotspot20/anqpipaddresstype:Anqpipaddresstype", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Anqpipaddresstype resources.
type anqpipaddresstypeState struct {
	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType *string `pulumi:"ipv4AddressType"`
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType *string `pulumi:"ipv6AddressType"`
	// IP type name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type AnqpipaddresstypeState struct {
	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType pulumi.StringPtrInput
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType pulumi.StringPtrInput
	// IP type name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AnqpipaddresstypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*anqpipaddresstypeState)(nil)).Elem()
}

type anqpipaddresstypeArgs struct {
	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType *string `pulumi:"ipv4AddressType"`
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType *string `pulumi:"ipv6AddressType"`
	// IP type name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Anqpipaddresstype resource.
type AnqpipaddresstypeArgs struct {
	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType pulumi.StringPtrInput
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType pulumi.StringPtrInput
	// IP type name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AnqpipaddresstypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*anqpipaddresstypeArgs)(nil)).Elem()
}

type AnqpipaddresstypeInput interface {
	pulumi.Input

	ToAnqpipaddresstypeOutput() AnqpipaddresstypeOutput
	ToAnqpipaddresstypeOutputWithContext(ctx context.Context) AnqpipaddresstypeOutput
}

func (*Anqpipaddresstype) ElementType() reflect.Type {
	return reflect.TypeOf((**Anqpipaddresstype)(nil)).Elem()
}

func (i *Anqpipaddresstype) ToAnqpipaddresstypeOutput() AnqpipaddresstypeOutput {
	return i.ToAnqpipaddresstypeOutputWithContext(context.Background())
}

func (i *Anqpipaddresstype) ToAnqpipaddresstypeOutputWithContext(ctx context.Context) AnqpipaddresstypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnqpipaddresstypeOutput)
}

// AnqpipaddresstypeArrayInput is an input type that accepts AnqpipaddresstypeArray and AnqpipaddresstypeArrayOutput values.
// You can construct a concrete instance of `AnqpipaddresstypeArrayInput` via:
//
//	AnqpipaddresstypeArray{ AnqpipaddresstypeArgs{...} }
type AnqpipaddresstypeArrayInput interface {
	pulumi.Input

	ToAnqpipaddresstypeArrayOutput() AnqpipaddresstypeArrayOutput
	ToAnqpipaddresstypeArrayOutputWithContext(context.Context) AnqpipaddresstypeArrayOutput
}

type AnqpipaddresstypeArray []AnqpipaddresstypeInput

func (AnqpipaddresstypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Anqpipaddresstype)(nil)).Elem()
}

func (i AnqpipaddresstypeArray) ToAnqpipaddresstypeArrayOutput() AnqpipaddresstypeArrayOutput {
	return i.ToAnqpipaddresstypeArrayOutputWithContext(context.Background())
}

func (i AnqpipaddresstypeArray) ToAnqpipaddresstypeArrayOutputWithContext(ctx context.Context) AnqpipaddresstypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnqpipaddresstypeArrayOutput)
}

// AnqpipaddresstypeMapInput is an input type that accepts AnqpipaddresstypeMap and AnqpipaddresstypeMapOutput values.
// You can construct a concrete instance of `AnqpipaddresstypeMapInput` via:
//
//	AnqpipaddresstypeMap{ "key": AnqpipaddresstypeArgs{...} }
type AnqpipaddresstypeMapInput interface {
	pulumi.Input

	ToAnqpipaddresstypeMapOutput() AnqpipaddresstypeMapOutput
	ToAnqpipaddresstypeMapOutputWithContext(context.Context) AnqpipaddresstypeMapOutput
}

type AnqpipaddresstypeMap map[string]AnqpipaddresstypeInput

func (AnqpipaddresstypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Anqpipaddresstype)(nil)).Elem()
}

func (i AnqpipaddresstypeMap) ToAnqpipaddresstypeMapOutput() AnqpipaddresstypeMapOutput {
	return i.ToAnqpipaddresstypeMapOutputWithContext(context.Background())
}

func (i AnqpipaddresstypeMap) ToAnqpipaddresstypeMapOutputWithContext(ctx context.Context) AnqpipaddresstypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnqpipaddresstypeMapOutput)
}

type AnqpipaddresstypeOutput struct{ *pulumi.OutputState }

func (AnqpipaddresstypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Anqpipaddresstype)(nil)).Elem()
}

func (o AnqpipaddresstypeOutput) ToAnqpipaddresstypeOutput() AnqpipaddresstypeOutput {
	return o
}

func (o AnqpipaddresstypeOutput) ToAnqpipaddresstypeOutputWithContext(ctx context.Context) AnqpipaddresstypeOutput {
	return o
}

// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
func (o AnqpipaddresstypeOutput) Ipv4AddressType() pulumi.StringOutput {
	return o.ApplyT(func(v *Anqpipaddresstype) pulumi.StringOutput { return v.Ipv4AddressType }).(pulumi.StringOutput)
}

// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
func (o AnqpipaddresstypeOutput) Ipv6AddressType() pulumi.StringOutput {
	return o.ApplyT(func(v *Anqpipaddresstype) pulumi.StringOutput { return v.Ipv6AddressType }).(pulumi.StringOutput)
}

// IP type name.
func (o AnqpipaddresstypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Anqpipaddresstype) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AnqpipaddresstypeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Anqpipaddresstype) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AnqpipaddresstypeArrayOutput struct{ *pulumi.OutputState }

func (AnqpipaddresstypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Anqpipaddresstype)(nil)).Elem()
}

func (o AnqpipaddresstypeArrayOutput) ToAnqpipaddresstypeArrayOutput() AnqpipaddresstypeArrayOutput {
	return o
}

func (o AnqpipaddresstypeArrayOutput) ToAnqpipaddresstypeArrayOutputWithContext(ctx context.Context) AnqpipaddresstypeArrayOutput {
	return o
}

func (o AnqpipaddresstypeArrayOutput) Index(i pulumi.IntInput) AnqpipaddresstypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Anqpipaddresstype {
		return vs[0].([]*Anqpipaddresstype)[vs[1].(int)]
	}).(AnqpipaddresstypeOutput)
}

type AnqpipaddresstypeMapOutput struct{ *pulumi.OutputState }

func (AnqpipaddresstypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Anqpipaddresstype)(nil)).Elem()
}

func (o AnqpipaddresstypeMapOutput) ToAnqpipaddresstypeMapOutput() AnqpipaddresstypeMapOutput {
	return o
}

func (o AnqpipaddresstypeMapOutput) ToAnqpipaddresstypeMapOutputWithContext(ctx context.Context) AnqpipaddresstypeMapOutput {
	return o
}

func (o AnqpipaddresstypeMapOutput) MapIndex(k pulumi.StringInput) AnqpipaddresstypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Anqpipaddresstype {
		return vs[0].(map[string]*Anqpipaddresstype)[vs[1].(string)]
	}).(AnqpipaddresstypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnqpipaddresstypeInput)(nil)).Elem(), &Anqpipaddresstype{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnqpipaddresstypeArrayInput)(nil)).Elem(), AnqpipaddresstypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnqpipaddresstypeMapInput)(nil)).Elem(), AnqpipaddresstypeMap{})
	pulumi.RegisterOutputType(AnqpipaddresstypeOutput{})
	pulumi.RegisterOutputType(AnqpipaddresstypeArrayOutput{})
	pulumi.RegisterOutputType(AnqpipaddresstypeMapOutput{})
}
