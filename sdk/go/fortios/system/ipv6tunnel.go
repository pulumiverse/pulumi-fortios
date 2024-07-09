// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPv6/IPv4 in IPv6 tunnel.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewIpv6tunnel(ctx, "trname", &system.Ipv6tunnelArgs{
//				Destination: pulumi.String("2001:db8:85a3::8a2e:370:7324"),
//				Interface:   pulumi.String("port3"),
//				Source:      pulumi.String("2001:db8:85a3::8a2e:370:7334"),
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
// System Ipv6Tunnel can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/ipv6tunnel:Ipv6tunnel labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/ipv6tunnel:Ipv6tunnel labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Ipv6tunnel struct {
	pulumi.CustomResourceState

	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload pulumi.StringOutput `pulumi:"autoAsicOffload"`
	// Remote IPv6 address of the tunnel.
	Destination pulumi.StringOutput `pulumi:"destination"`
	// Interface name.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// IPv6 tunnel name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Local IPv6 address of the tunnel.
	Source pulumi.StringOutput `pulumi:"source"`
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan pulumi.StringOutput `pulumi:"useSdwan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewIpv6tunnel registers a new resource with the given unique name, arguments, and options.
func NewIpv6tunnel(ctx *pulumi.Context,
	name string, args *Ipv6tunnelArgs, opts ...pulumi.ResourceOption) (*Ipv6tunnel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ipv6tunnel
	err := ctx.RegisterResource("fortios:system/ipv6tunnel:Ipv6tunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpv6tunnel gets an existing Ipv6tunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpv6tunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ipv6tunnelState, opts ...pulumi.ResourceOption) (*Ipv6tunnel, error) {
	var resource Ipv6tunnel
	err := ctx.ReadResource("fortios:system/ipv6tunnel:Ipv6tunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipv6tunnel resources.
type ipv6tunnelState struct {
	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload *string `pulumi:"autoAsicOffload"`
	// Remote IPv6 address of the tunnel.
	Destination *string `pulumi:"destination"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// IPv6 tunnel name.
	Name *string `pulumi:"name"`
	// Local IPv6 address of the tunnel.
	Source *string `pulumi:"source"`
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan *string `pulumi:"useSdwan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Ipv6tunnelState struct {
	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload pulumi.StringPtrInput
	// Remote IPv6 address of the tunnel.
	Destination pulumi.StringPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// IPv6 tunnel name.
	Name pulumi.StringPtrInput
	// Local IPv6 address of the tunnel.
	Source pulumi.StringPtrInput
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Ipv6tunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6tunnelState)(nil)).Elem()
}

type ipv6tunnelArgs struct {
	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload *string `pulumi:"autoAsicOffload"`
	// Remote IPv6 address of the tunnel.
	Destination string `pulumi:"destination"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// IPv6 tunnel name.
	Name *string `pulumi:"name"`
	// Local IPv6 address of the tunnel.
	Source *string `pulumi:"source"`
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan *string `pulumi:"useSdwan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Ipv6tunnel resource.
type Ipv6tunnelArgs struct {
	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload pulumi.StringPtrInput
	// Remote IPv6 address of the tunnel.
	Destination pulumi.StringInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// IPv6 tunnel name.
	Name pulumi.StringPtrInput
	// Local IPv6 address of the tunnel.
	Source pulumi.StringPtrInput
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Ipv6tunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6tunnelArgs)(nil)).Elem()
}

type Ipv6tunnelInput interface {
	pulumi.Input

	ToIpv6tunnelOutput() Ipv6tunnelOutput
	ToIpv6tunnelOutputWithContext(ctx context.Context) Ipv6tunnelOutput
}

func (*Ipv6tunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6tunnel)(nil)).Elem()
}

func (i *Ipv6tunnel) ToIpv6tunnelOutput() Ipv6tunnelOutput {
	return i.ToIpv6tunnelOutputWithContext(context.Background())
}

func (i *Ipv6tunnel) ToIpv6tunnelOutputWithContext(ctx context.Context) Ipv6tunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6tunnelOutput)
}

// Ipv6tunnelArrayInput is an input type that accepts Ipv6tunnelArray and Ipv6tunnelArrayOutput values.
// You can construct a concrete instance of `Ipv6tunnelArrayInput` via:
//
//	Ipv6tunnelArray{ Ipv6tunnelArgs{...} }
type Ipv6tunnelArrayInput interface {
	pulumi.Input

	ToIpv6tunnelArrayOutput() Ipv6tunnelArrayOutput
	ToIpv6tunnelArrayOutputWithContext(context.Context) Ipv6tunnelArrayOutput
}

type Ipv6tunnelArray []Ipv6tunnelInput

func (Ipv6tunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv6tunnel)(nil)).Elem()
}

func (i Ipv6tunnelArray) ToIpv6tunnelArrayOutput() Ipv6tunnelArrayOutput {
	return i.ToIpv6tunnelArrayOutputWithContext(context.Background())
}

func (i Ipv6tunnelArray) ToIpv6tunnelArrayOutputWithContext(ctx context.Context) Ipv6tunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6tunnelArrayOutput)
}

// Ipv6tunnelMapInput is an input type that accepts Ipv6tunnelMap and Ipv6tunnelMapOutput values.
// You can construct a concrete instance of `Ipv6tunnelMapInput` via:
//
//	Ipv6tunnelMap{ "key": Ipv6tunnelArgs{...} }
type Ipv6tunnelMapInput interface {
	pulumi.Input

	ToIpv6tunnelMapOutput() Ipv6tunnelMapOutput
	ToIpv6tunnelMapOutputWithContext(context.Context) Ipv6tunnelMapOutput
}

type Ipv6tunnelMap map[string]Ipv6tunnelInput

func (Ipv6tunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv6tunnel)(nil)).Elem()
}

func (i Ipv6tunnelMap) ToIpv6tunnelMapOutput() Ipv6tunnelMapOutput {
	return i.ToIpv6tunnelMapOutputWithContext(context.Background())
}

func (i Ipv6tunnelMap) ToIpv6tunnelMapOutputWithContext(ctx context.Context) Ipv6tunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6tunnelMapOutput)
}

type Ipv6tunnelOutput struct{ *pulumi.OutputState }

func (Ipv6tunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6tunnel)(nil)).Elem()
}

func (o Ipv6tunnelOutput) ToIpv6tunnelOutput() Ipv6tunnelOutput {
	return o
}

func (o Ipv6tunnelOutput) ToIpv6tunnelOutputWithContext(ctx context.Context) Ipv6tunnelOutput {
	return o
}

// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
func (o Ipv6tunnelOutput) AutoAsicOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6tunnel) pulumi.StringOutput { return v.AutoAsicOffload }).(pulumi.StringOutput)
}

// Remote IPv6 address of the tunnel.
func (o Ipv6tunnelOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6tunnel) pulumi.StringOutput { return v.Destination }).(pulumi.StringOutput)
}

// Interface name.
func (o Ipv6tunnelOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6tunnel) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// IPv6 tunnel name.
func (o Ipv6tunnelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6tunnel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Local IPv6 address of the tunnel.
func (o Ipv6tunnelOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6tunnel) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
func (o Ipv6tunnelOutput) UseSdwan() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6tunnel) pulumi.StringOutput { return v.UseSdwan }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Ipv6tunnelOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6tunnel) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type Ipv6tunnelArrayOutput struct{ *pulumi.OutputState }

func (Ipv6tunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv6tunnel)(nil)).Elem()
}

func (o Ipv6tunnelArrayOutput) ToIpv6tunnelArrayOutput() Ipv6tunnelArrayOutput {
	return o
}

func (o Ipv6tunnelArrayOutput) ToIpv6tunnelArrayOutputWithContext(ctx context.Context) Ipv6tunnelArrayOutput {
	return o
}

func (o Ipv6tunnelArrayOutput) Index(i pulumi.IntInput) Ipv6tunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ipv6tunnel {
		return vs[0].([]*Ipv6tunnel)[vs[1].(int)]
	}).(Ipv6tunnelOutput)
}

type Ipv6tunnelMapOutput struct{ *pulumi.OutputState }

func (Ipv6tunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv6tunnel)(nil)).Elem()
}

func (o Ipv6tunnelMapOutput) ToIpv6tunnelMapOutput() Ipv6tunnelMapOutput {
	return o
}

func (o Ipv6tunnelMapOutput) ToIpv6tunnelMapOutputWithContext(ctx context.Context) Ipv6tunnelMapOutput {
	return o
}

func (o Ipv6tunnelMapOutput) MapIndex(k pulumi.StringInput) Ipv6tunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ipv6tunnel {
		return vs[0].(map[string]*Ipv6tunnel)[vs[1].(string)]
	}).(Ipv6tunnelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6tunnelInput)(nil)).Elem(), &Ipv6tunnel{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6tunnelArrayInput)(nil)).Elem(), Ipv6tunnelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6tunnelMapInput)(nil)).Elem(), Ipv6tunnelMap{})
	pulumi.RegisterOutputType(Ipv6tunnelOutput{})
	pulumi.RegisterOutputType(Ipv6tunnelArrayOutput{})
	pulumi.RegisterOutputType(Ipv6tunnelMapOutput{})
}
