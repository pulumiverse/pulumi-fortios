// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IP address type availability.
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
//			_, err := fortios.NewWirelesscontrollerhotspot20Anqpipaddresstype(ctx, "trname", &fortios.Wirelesscontrollerhotspot20AnqpipaddresstypeArgs{
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
//
// ## Import
//
// # WirelessControllerHotspot20 AnqpIpAddressType can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerhotspot20Anqpipaddresstype:Wirelesscontrollerhotspot20Anqpipaddresstype labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerhotspot20Anqpipaddresstype:Wirelesscontrollerhotspot20Anqpipaddresstype labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Wirelesscontrollerhotspot20Anqpipaddresstype struct {
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

// NewWirelesscontrollerhotspot20Anqpipaddresstype registers a new resource with the given unique name, arguments, and options.
func NewWirelesscontrollerhotspot20Anqpipaddresstype(ctx *pulumi.Context,
	name string, args *Wirelesscontrollerhotspot20AnqpipaddresstypeArgs, opts ...pulumi.ResourceOption) (*Wirelesscontrollerhotspot20Anqpipaddresstype, error) {
	if args == nil {
		args = &Wirelesscontrollerhotspot20AnqpipaddresstypeArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Wirelesscontrollerhotspot20Anqpipaddresstype
	err := ctx.RegisterResource("fortios:index/wirelesscontrollerhotspot20Anqpipaddresstype:Wirelesscontrollerhotspot20Anqpipaddresstype", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelesscontrollerhotspot20Anqpipaddresstype gets an existing Wirelesscontrollerhotspot20Anqpipaddresstype resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelesscontrollerhotspot20Anqpipaddresstype(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Wirelesscontrollerhotspot20AnqpipaddresstypeState, opts ...pulumi.ResourceOption) (*Wirelesscontrollerhotspot20Anqpipaddresstype, error) {
	var resource Wirelesscontrollerhotspot20Anqpipaddresstype
	err := ctx.ReadResource("fortios:index/wirelesscontrollerhotspot20Anqpipaddresstype:Wirelesscontrollerhotspot20Anqpipaddresstype", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Wirelesscontrollerhotspot20Anqpipaddresstype resources.
type wirelesscontrollerhotspot20AnqpipaddresstypeState struct {
	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType *string `pulumi:"ipv4AddressType"`
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType *string `pulumi:"ipv6AddressType"`
	// IP type name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Wirelesscontrollerhotspot20AnqpipaddresstypeState struct {
	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType pulumi.StringPtrInput
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType pulumi.StringPtrInput
	// IP type name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Wirelesscontrollerhotspot20AnqpipaddresstypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerhotspot20AnqpipaddresstypeState)(nil)).Elem()
}

type wirelesscontrollerhotspot20AnqpipaddresstypeArgs struct {
	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType *string `pulumi:"ipv4AddressType"`
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType *string `pulumi:"ipv6AddressType"`
	// IP type name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Wirelesscontrollerhotspot20Anqpipaddresstype resource.
type Wirelesscontrollerhotspot20AnqpipaddresstypeArgs struct {
	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType pulumi.StringPtrInput
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType pulumi.StringPtrInput
	// IP type name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Wirelesscontrollerhotspot20AnqpipaddresstypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerhotspot20AnqpipaddresstypeArgs)(nil)).Elem()
}

type Wirelesscontrollerhotspot20AnqpipaddresstypeInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20AnqpipaddresstypeOutput() Wirelesscontrollerhotspot20AnqpipaddresstypeOutput
	ToWirelesscontrollerhotspot20AnqpipaddresstypeOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpipaddresstypeOutput
}

func (*Wirelesscontrollerhotspot20Anqpipaddresstype) ElementType() reflect.Type {
	return reflect.TypeOf((**Wirelesscontrollerhotspot20Anqpipaddresstype)(nil)).Elem()
}

func (i *Wirelesscontrollerhotspot20Anqpipaddresstype) ToWirelesscontrollerhotspot20AnqpipaddresstypeOutput() Wirelesscontrollerhotspot20AnqpipaddresstypeOutput {
	return i.ToWirelesscontrollerhotspot20AnqpipaddresstypeOutputWithContext(context.Background())
}

func (i *Wirelesscontrollerhotspot20Anqpipaddresstype) ToWirelesscontrollerhotspot20AnqpipaddresstypeOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpipaddresstypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20AnqpipaddresstypeOutput)
}

// Wirelesscontrollerhotspot20AnqpipaddresstypeArrayInput is an input type that accepts Wirelesscontrollerhotspot20AnqpipaddresstypeArray and Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput values.
// You can construct a concrete instance of `Wirelesscontrollerhotspot20AnqpipaddresstypeArrayInput` via:
//
//	Wirelesscontrollerhotspot20AnqpipaddresstypeArray{ Wirelesscontrollerhotspot20AnqpipaddresstypeArgs{...} }
type Wirelesscontrollerhotspot20AnqpipaddresstypeArrayInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput() Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput
	ToWirelesscontrollerhotspot20AnqpipaddresstypeArrayOutputWithContext(context.Context) Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput
}

type Wirelesscontrollerhotspot20AnqpipaddresstypeArray []Wirelesscontrollerhotspot20AnqpipaddresstypeInput

func (Wirelesscontrollerhotspot20AnqpipaddresstypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wirelesscontrollerhotspot20Anqpipaddresstype)(nil)).Elem()
}

func (i Wirelesscontrollerhotspot20AnqpipaddresstypeArray) ToWirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput() Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput {
	return i.ToWirelesscontrollerhotspot20AnqpipaddresstypeArrayOutputWithContext(context.Background())
}

func (i Wirelesscontrollerhotspot20AnqpipaddresstypeArray) ToWirelesscontrollerhotspot20AnqpipaddresstypeArrayOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput)
}

// Wirelesscontrollerhotspot20AnqpipaddresstypeMapInput is an input type that accepts Wirelesscontrollerhotspot20AnqpipaddresstypeMap and Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput values.
// You can construct a concrete instance of `Wirelesscontrollerhotspot20AnqpipaddresstypeMapInput` via:
//
//	Wirelesscontrollerhotspot20AnqpipaddresstypeMap{ "key": Wirelesscontrollerhotspot20AnqpipaddresstypeArgs{...} }
type Wirelesscontrollerhotspot20AnqpipaddresstypeMapInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20AnqpipaddresstypeMapOutput() Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput
	ToWirelesscontrollerhotspot20AnqpipaddresstypeMapOutputWithContext(context.Context) Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput
}

type Wirelesscontrollerhotspot20AnqpipaddresstypeMap map[string]Wirelesscontrollerhotspot20AnqpipaddresstypeInput

func (Wirelesscontrollerhotspot20AnqpipaddresstypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wirelesscontrollerhotspot20Anqpipaddresstype)(nil)).Elem()
}

func (i Wirelesscontrollerhotspot20AnqpipaddresstypeMap) ToWirelesscontrollerhotspot20AnqpipaddresstypeMapOutput() Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput {
	return i.ToWirelesscontrollerhotspot20AnqpipaddresstypeMapOutputWithContext(context.Background())
}

func (i Wirelesscontrollerhotspot20AnqpipaddresstypeMap) ToWirelesscontrollerhotspot20AnqpipaddresstypeMapOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput)
}

type Wirelesscontrollerhotspot20AnqpipaddresstypeOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20AnqpipaddresstypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Wirelesscontrollerhotspot20Anqpipaddresstype)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20AnqpipaddresstypeOutput) ToWirelesscontrollerhotspot20AnqpipaddresstypeOutput() Wirelesscontrollerhotspot20AnqpipaddresstypeOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpipaddresstypeOutput) ToWirelesscontrollerhotspot20AnqpipaddresstypeOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpipaddresstypeOutput {
	return o
}

// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
func (o Wirelesscontrollerhotspot20AnqpipaddresstypeOutput) Ipv4AddressType() pulumi.StringOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20Anqpipaddresstype) pulumi.StringOutput { return v.Ipv4AddressType }).(pulumi.StringOutput)
}

// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
func (o Wirelesscontrollerhotspot20AnqpipaddresstypeOutput) Ipv6AddressType() pulumi.StringOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20Anqpipaddresstype) pulumi.StringOutput { return v.Ipv6AddressType }).(pulumi.StringOutput)
}

// IP type name.
func (o Wirelesscontrollerhotspot20AnqpipaddresstypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20Anqpipaddresstype) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Wirelesscontrollerhotspot20AnqpipaddresstypeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20Anqpipaddresstype) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wirelesscontrollerhotspot20Anqpipaddresstype)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput) ToWirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput() Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput) ToWirelesscontrollerhotspot20AnqpipaddresstypeArrayOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput) Index(i pulumi.IntInput) Wirelesscontrollerhotspot20AnqpipaddresstypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Wirelesscontrollerhotspot20Anqpipaddresstype {
		return vs[0].([]*Wirelesscontrollerhotspot20Anqpipaddresstype)[vs[1].(int)]
	}).(Wirelesscontrollerhotspot20AnqpipaddresstypeOutput)
}

type Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wirelesscontrollerhotspot20Anqpipaddresstype)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput) ToWirelesscontrollerhotspot20AnqpipaddresstypeMapOutput() Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput) ToWirelesscontrollerhotspot20AnqpipaddresstypeMapOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput) MapIndex(k pulumi.StringInput) Wirelesscontrollerhotspot20AnqpipaddresstypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Wirelesscontrollerhotspot20Anqpipaddresstype {
		return vs[0].(map[string]*Wirelesscontrollerhotspot20Anqpipaddresstype)[vs[1].(string)]
	}).(Wirelesscontrollerhotspot20AnqpipaddresstypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20AnqpipaddresstypeInput)(nil)).Elem(), &Wirelesscontrollerhotspot20Anqpipaddresstype{})
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20AnqpipaddresstypeArrayInput)(nil)).Elem(), Wirelesscontrollerhotspot20AnqpipaddresstypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20AnqpipaddresstypeMapInput)(nil)).Elem(), Wirelesscontrollerhotspot20AnqpipaddresstypeMap{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20AnqpipaddresstypeOutput{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20AnqpipaddresstypeArrayOutput{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20AnqpipaddresstypeMapOutput{})
}
