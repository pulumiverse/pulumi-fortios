// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure wireless access gateway (WAG) profiles used for tunnels on AP. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// WirelessController WagProfile can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/wagprofile:Wagprofile labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/wagprofile:Wagprofile labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Wagprofile struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// IP address of the monitoring DHCP request packet sent through the tunnel.
	DhcpIpAddr pulumi.StringOutput `pulumi:"dhcpIpAddr"`
	// Tunnel profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
	PingInterval pulumi.IntOutput `pulumi:"pingInterval"`
	// Number of the tunnel monitoring echo packets (1 - 65535, default = 5).
	PingNumber pulumi.IntOutput `pulumi:"pingNumber"`
	// Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
	ReturnPacketTimeout pulumi.IntOutput `pulumi:"returnPacketTimeout"`
	// Tunnel type. Valid values: `l2tpv3`, `gre`.
	TunnelType pulumi.StringOutput `pulumi:"tunnelType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// IP Address of the wireless access gateway.
	WagIp pulumi.StringOutput `pulumi:"wagIp"`
	// UDP port of the wireless access gateway.
	WagPort pulumi.IntOutput `pulumi:"wagPort"`
}

// NewWagprofile registers a new resource with the given unique name, arguments, and options.
func NewWagprofile(ctx *pulumi.Context,
	name string, args *WagprofileArgs, opts ...pulumi.ResourceOption) (*Wagprofile, error) {
	if args == nil {
		args = &WagprofileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Wagprofile
	err := ctx.RegisterResource("fortios:wirelesscontroller/wagprofile:Wagprofile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWagprofile gets an existing Wagprofile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWagprofile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WagprofileState, opts ...pulumi.ResourceOption) (*Wagprofile, error) {
	var resource Wagprofile
	err := ctx.ReadResource("fortios:wirelesscontroller/wagprofile:Wagprofile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Wagprofile resources.
type wagprofileState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// IP address of the monitoring DHCP request packet sent through the tunnel.
	DhcpIpAddr *string `pulumi:"dhcpIpAddr"`
	// Tunnel profile name.
	Name *string `pulumi:"name"`
	// Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
	PingInterval *int `pulumi:"pingInterval"`
	// Number of the tunnel monitoring echo packets (1 - 65535, default = 5).
	PingNumber *int `pulumi:"pingNumber"`
	// Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
	ReturnPacketTimeout *int `pulumi:"returnPacketTimeout"`
	// Tunnel type. Valid values: `l2tpv3`, `gre`.
	TunnelType *string `pulumi:"tunnelType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// IP Address of the wireless access gateway.
	WagIp *string `pulumi:"wagIp"`
	// UDP port of the wireless access gateway.
	WagPort *int `pulumi:"wagPort"`
}

type WagprofileState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// IP address of the monitoring DHCP request packet sent through the tunnel.
	DhcpIpAddr pulumi.StringPtrInput
	// Tunnel profile name.
	Name pulumi.StringPtrInput
	// Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
	PingInterval pulumi.IntPtrInput
	// Number of the tunnel monitoring echo packets (1 - 65535, default = 5).
	PingNumber pulumi.IntPtrInput
	// Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
	ReturnPacketTimeout pulumi.IntPtrInput
	// Tunnel type. Valid values: `l2tpv3`, `gre`.
	TunnelType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// IP Address of the wireless access gateway.
	WagIp pulumi.StringPtrInput
	// UDP port of the wireless access gateway.
	WagPort pulumi.IntPtrInput
}

func (WagprofileState) ElementType() reflect.Type {
	return reflect.TypeOf((*wagprofileState)(nil)).Elem()
}

type wagprofileArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// IP address of the monitoring DHCP request packet sent through the tunnel.
	DhcpIpAddr *string `pulumi:"dhcpIpAddr"`
	// Tunnel profile name.
	Name *string `pulumi:"name"`
	// Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
	PingInterval *int `pulumi:"pingInterval"`
	// Number of the tunnel monitoring echo packets (1 - 65535, default = 5).
	PingNumber *int `pulumi:"pingNumber"`
	// Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
	ReturnPacketTimeout *int `pulumi:"returnPacketTimeout"`
	// Tunnel type. Valid values: `l2tpv3`, `gre`.
	TunnelType *string `pulumi:"tunnelType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// IP Address of the wireless access gateway.
	WagIp *string `pulumi:"wagIp"`
	// UDP port of the wireless access gateway.
	WagPort *int `pulumi:"wagPort"`
}

// The set of arguments for constructing a Wagprofile resource.
type WagprofileArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// IP address of the monitoring DHCP request packet sent through the tunnel.
	DhcpIpAddr pulumi.StringPtrInput
	// Tunnel profile name.
	Name pulumi.StringPtrInput
	// Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
	PingInterval pulumi.IntPtrInput
	// Number of the tunnel monitoring echo packets (1 - 65535, default = 5).
	PingNumber pulumi.IntPtrInput
	// Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
	ReturnPacketTimeout pulumi.IntPtrInput
	// Tunnel type. Valid values: `l2tpv3`, `gre`.
	TunnelType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// IP Address of the wireless access gateway.
	WagIp pulumi.StringPtrInput
	// UDP port of the wireless access gateway.
	WagPort pulumi.IntPtrInput
}

func (WagprofileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wagprofileArgs)(nil)).Elem()
}

type WagprofileInput interface {
	pulumi.Input

	ToWagprofileOutput() WagprofileOutput
	ToWagprofileOutputWithContext(ctx context.Context) WagprofileOutput
}

func (*Wagprofile) ElementType() reflect.Type {
	return reflect.TypeOf((**Wagprofile)(nil)).Elem()
}

func (i *Wagprofile) ToWagprofileOutput() WagprofileOutput {
	return i.ToWagprofileOutputWithContext(context.Background())
}

func (i *Wagprofile) ToWagprofileOutputWithContext(ctx context.Context) WagprofileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WagprofileOutput)
}

// WagprofileArrayInput is an input type that accepts WagprofileArray and WagprofileArrayOutput values.
// You can construct a concrete instance of `WagprofileArrayInput` via:
//
//	WagprofileArray{ WagprofileArgs{...} }
type WagprofileArrayInput interface {
	pulumi.Input

	ToWagprofileArrayOutput() WagprofileArrayOutput
	ToWagprofileArrayOutputWithContext(context.Context) WagprofileArrayOutput
}

type WagprofileArray []WagprofileInput

func (WagprofileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wagprofile)(nil)).Elem()
}

func (i WagprofileArray) ToWagprofileArrayOutput() WagprofileArrayOutput {
	return i.ToWagprofileArrayOutputWithContext(context.Background())
}

func (i WagprofileArray) ToWagprofileArrayOutputWithContext(ctx context.Context) WagprofileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WagprofileArrayOutput)
}

// WagprofileMapInput is an input type that accepts WagprofileMap and WagprofileMapOutput values.
// You can construct a concrete instance of `WagprofileMapInput` via:
//
//	WagprofileMap{ "key": WagprofileArgs{...} }
type WagprofileMapInput interface {
	pulumi.Input

	ToWagprofileMapOutput() WagprofileMapOutput
	ToWagprofileMapOutputWithContext(context.Context) WagprofileMapOutput
}

type WagprofileMap map[string]WagprofileInput

func (WagprofileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wagprofile)(nil)).Elem()
}

func (i WagprofileMap) ToWagprofileMapOutput() WagprofileMapOutput {
	return i.ToWagprofileMapOutputWithContext(context.Background())
}

func (i WagprofileMap) ToWagprofileMapOutputWithContext(ctx context.Context) WagprofileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WagprofileMapOutput)
}

type WagprofileOutput struct{ *pulumi.OutputState }

func (WagprofileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Wagprofile)(nil)).Elem()
}

func (o WagprofileOutput) ToWagprofileOutput() WagprofileOutput {
	return o
}

func (o WagprofileOutput) ToWagprofileOutputWithContext(ctx context.Context) WagprofileOutput {
	return o
}

// Comment.
func (o WagprofileOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wagprofile) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// IP address of the monitoring DHCP request packet sent through the tunnel.
func (o WagprofileOutput) DhcpIpAddr() pulumi.StringOutput {
	return o.ApplyT(func(v *Wagprofile) pulumi.StringOutput { return v.DhcpIpAddr }).(pulumi.StringOutput)
}

// Tunnel profile name.
func (o WagprofileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Wagprofile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
func (o WagprofileOutput) PingInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Wagprofile) pulumi.IntOutput { return v.PingInterval }).(pulumi.IntOutput)
}

// Number of the tunnel monitoring echo packets (1 - 65535, default = 5).
func (o WagprofileOutput) PingNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Wagprofile) pulumi.IntOutput { return v.PingNumber }).(pulumi.IntOutput)
}

// Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
func (o WagprofileOutput) ReturnPacketTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Wagprofile) pulumi.IntOutput { return v.ReturnPacketTimeout }).(pulumi.IntOutput)
}

// Tunnel type. Valid values: `l2tpv3`, `gre`.
func (o WagprofileOutput) TunnelType() pulumi.StringOutput {
	return o.ApplyT(func(v *Wagprofile) pulumi.StringOutput { return v.TunnelType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WagprofileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wagprofile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// IP Address of the wireless access gateway.
func (o WagprofileOutput) WagIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Wagprofile) pulumi.StringOutput { return v.WagIp }).(pulumi.StringOutput)
}

// UDP port of the wireless access gateway.
func (o WagprofileOutput) WagPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Wagprofile) pulumi.IntOutput { return v.WagPort }).(pulumi.IntOutput)
}

type WagprofileArrayOutput struct{ *pulumi.OutputState }

func (WagprofileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wagprofile)(nil)).Elem()
}

func (o WagprofileArrayOutput) ToWagprofileArrayOutput() WagprofileArrayOutput {
	return o
}

func (o WagprofileArrayOutput) ToWagprofileArrayOutputWithContext(ctx context.Context) WagprofileArrayOutput {
	return o
}

func (o WagprofileArrayOutput) Index(i pulumi.IntInput) WagprofileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Wagprofile {
		return vs[0].([]*Wagprofile)[vs[1].(int)]
	}).(WagprofileOutput)
}

type WagprofileMapOutput struct{ *pulumi.OutputState }

func (WagprofileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wagprofile)(nil)).Elem()
}

func (o WagprofileMapOutput) ToWagprofileMapOutput() WagprofileMapOutput {
	return o
}

func (o WagprofileMapOutput) ToWagprofileMapOutputWithContext(ctx context.Context) WagprofileMapOutput {
	return o
}

func (o WagprofileMapOutput) MapIndex(k pulumi.StringInput) WagprofileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Wagprofile {
		return vs[0].(map[string]*Wagprofile)[vs[1].(string)]
	}).(WagprofileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WagprofileInput)(nil)).Elem(), &Wagprofile{})
	pulumi.RegisterInputType(reflect.TypeOf((*WagprofileArrayInput)(nil)).Elem(), WagprofileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WagprofileMapInput)(nil)).Elem(), WagprofileMap{})
	pulumi.RegisterOutputType(WagprofileOutput{})
	pulumi.RegisterOutputType(WagprofileArrayOutput{})
	pulumi.RegisterOutputType(WagprofileMapOutput{})
}