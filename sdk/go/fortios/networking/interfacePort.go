// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a resource to configure interface settings of FortiOS.
//
// !> **Warning:** The resource will be deprecated and replaced by new resource `system.Interface`, we recommend that you use the new resource.
//
// ## Example Usage
//
// ### Loopback Interface
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/networking"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networking.NewInterfacePort(ctx, "loopback1", &networking.InterfacePortArgs{
//				Alias:       pulumi.String("cc1"),
//				Allowaccess: pulumi.String("ping http"),
//				Description: pulumi.String("description"),
//				Ip:          pulumi.String("23.123.33.10 255.255.255.0"),
//				Mode:        pulumi.String("static"),
//				Role:        pulumi.String("lan"),
//				Status:      pulumi.String("up"),
//				Type:        pulumi.String("loopback"),
//				Vdom:        pulumi.String("root"),
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
// ### VLAN Interface
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/networking"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networking.NewInterfacePort(ctx, "vlan1", &networking.InterfacePortArgs{
//				Allowaccess: pulumi.String("ping"),
//				Defaultgw:   pulumi.String("enable"),
//				Distance:    pulumi.String("33"),
//				Interface:   pulumi.String("port2"),
//				Ip:          pulumi.String("3.123.33.10 255.255.255.0"),
//				Mode:        pulumi.String("static"),
//				Role:        pulumi.String("lan"),
//				Type:        pulumi.String("vlan"),
//				Vdom:        pulumi.String("root"),
//				Vlanid:      pulumi.String("3"),
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
// ### Physical Interface
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/networking"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networking.NewInterfacePort(ctx, "test1", &networking.InterfacePortArgs{
//				Alias:                pulumi.String("dkeeew"),
//				Allowaccess:          pulumi.String("ping https"),
//				Defaultgw:            pulumi.String("enable"),
//				Description:          pulumi.String("description"),
//				DeviceIdentification: pulumi.String("enable"),
//				Distance:             pulumi.String("33"),
//				DnsServerOverride:    pulumi.String("enable"),
//				Ip:                   pulumi.String("93.133.133.110 255.255.255.0"),
//				Mode:                 pulumi.String("static"),
//				Mtu:                  pulumi.String("2933"),
//				MtuOverride:          pulumi.String("enable"),
//				Role:                 pulumi.String("lan"),
//				Speed:                pulumi.String("auto"),
//				Status:               pulumi.String("up"),
//				TcpMss:               pulumi.String("3232"),
//				Type:                 pulumi.String("physical"),
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
type InterfacePort struct {
	pulumi.CustomResourceState

	// Alias will be displayed with the interface name to make it easier to distinguish.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Permitted types of management access to this interface.
	Allowaccess pulumi.StringOutput `pulumi:"allowaccess"`
	// Enable to get the gateway IP from the DHCP or PPPoE server.
	Defaultgw pulumi.StringOutput `pulumi:"defaultgw"`
	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
	DeviceIdentification pulumi.StringOutput `pulumi:"deviceIdentification"`
	// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
	Distance pulumi.StringOutput `pulumi:"distance"`
	// Enable/disable use DNS acquired by DHCP or PPPoE.
	DnsServerOverride pulumi.StringOutput `pulumi:"dnsServerOverride"`
	// Interface name.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Addressing mode.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// MTU value for this interface.
	Mtu pulumi.StringOutput `pulumi:"mtu"`
	// Enable to set a custom MTU for this interface.
	MtuOverride pulumi.StringOutput `pulumi:"mtuOverride"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Interface role.
	Role pulumi.StringOutput `pulumi:"role"`
	// Interface speed. The default setting and the options available depend on the interface hardware.
	Speed pulumi.StringOutput `pulumi:"speed"`
	// Bring the interface up or shut the interface down.
	Status pulumi.StringOutput `pulumi:"status"`
	// TCP maximum segment size. 0 means do not change segment size.
	TcpMss pulumi.StringOutput `pulumi:"tcpMss"`
	// Interface type (support physical, vlan, loopback).
	Type pulumi.StringOutput `pulumi:"type"`
	// Interface is in this virtual domain (VDOM).
	Vdom pulumi.StringOutput `pulumi:"vdom"`
	// VLAN ID.
	Vlanid pulumi.StringOutput `pulumi:"vlanid"`
}

// NewInterfacePort registers a new resource with the given unique name, arguments, and options.
func NewInterfacePort(ctx *pulumi.Context,
	name string, args *InterfacePortArgs, opts ...pulumi.ResourceOption) (*InterfacePort, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InterfacePort
	err := ctx.RegisterResource("fortios:networking/interfacePort:InterfacePort", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInterfacePort gets an existing InterfacePort resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInterfacePort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InterfacePortState, opts ...pulumi.ResourceOption) (*InterfacePort, error) {
	var resource InterfacePort
	err := ctx.ReadResource("fortios:networking/interfacePort:InterfacePort", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InterfacePort resources.
type interfacePortState struct {
	// Alias will be displayed with the interface name to make it easier to distinguish.
	Alias *string `pulumi:"alias"`
	// Permitted types of management access to this interface.
	Allowaccess *string `pulumi:"allowaccess"`
	// Enable to get the gateway IP from the DHCP or PPPoE server.
	Defaultgw *string `pulumi:"defaultgw"`
	// Description.
	Description *string `pulumi:"description"`
	// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
	DeviceIdentification *string `pulumi:"deviceIdentification"`
	// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
	Distance *string `pulumi:"distance"`
	// Enable/disable use DNS acquired by DHCP or PPPoE.
	DnsServerOverride *string `pulumi:"dnsServerOverride"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
	Ip *string `pulumi:"ip"`
	// Addressing mode.
	Mode *string `pulumi:"mode"`
	// MTU value for this interface.
	Mtu *string `pulumi:"mtu"`
	// Enable to set a custom MTU for this interface.
	MtuOverride *string `pulumi:"mtuOverride"`
	// Name.
	Name *string `pulumi:"name"`
	// Interface role.
	Role *string `pulumi:"role"`
	// Interface speed. The default setting and the options available depend on the interface hardware.
	Speed *string `pulumi:"speed"`
	// Bring the interface up or shut the interface down.
	Status *string `pulumi:"status"`
	// TCP maximum segment size. 0 means do not change segment size.
	TcpMss *string `pulumi:"tcpMss"`
	// Interface type (support physical, vlan, loopback).
	Type *string `pulumi:"type"`
	// Interface is in this virtual domain (VDOM).
	Vdom *string `pulumi:"vdom"`
	// VLAN ID.
	Vlanid *string `pulumi:"vlanid"`
}

type InterfacePortState struct {
	// Alias will be displayed with the interface name to make it easier to distinguish.
	Alias pulumi.StringPtrInput
	// Permitted types of management access to this interface.
	Allowaccess pulumi.StringPtrInput
	// Enable to get the gateway IP from the DHCP or PPPoE server.
	Defaultgw pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
	DeviceIdentification pulumi.StringPtrInput
	// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
	Distance pulumi.StringPtrInput
	// Enable/disable use DNS acquired by DHCP or PPPoE.
	DnsServerOverride pulumi.StringPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
	Ip pulumi.StringPtrInput
	// Addressing mode.
	Mode pulumi.StringPtrInput
	// MTU value for this interface.
	Mtu pulumi.StringPtrInput
	// Enable to set a custom MTU for this interface.
	MtuOverride pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Interface role.
	Role pulumi.StringPtrInput
	// Interface speed. The default setting and the options available depend on the interface hardware.
	Speed pulumi.StringPtrInput
	// Bring the interface up or shut the interface down.
	Status pulumi.StringPtrInput
	// TCP maximum segment size. 0 means do not change segment size.
	TcpMss pulumi.StringPtrInput
	// Interface type (support physical, vlan, loopback).
	Type pulumi.StringPtrInput
	// Interface is in this virtual domain (VDOM).
	Vdom pulumi.StringPtrInput
	// VLAN ID.
	Vlanid pulumi.StringPtrInput
}

func (InterfacePortState) ElementType() reflect.Type {
	return reflect.TypeOf((*interfacePortState)(nil)).Elem()
}

type interfacePortArgs struct {
	// Alias will be displayed with the interface name to make it easier to distinguish.
	Alias *string `pulumi:"alias"`
	// Permitted types of management access to this interface.
	Allowaccess *string `pulumi:"allowaccess"`
	// Enable to get the gateway IP from the DHCP or PPPoE server.
	Defaultgw *string `pulumi:"defaultgw"`
	// Description.
	Description *string `pulumi:"description"`
	// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
	DeviceIdentification *string `pulumi:"deviceIdentification"`
	// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
	Distance *string `pulumi:"distance"`
	// Enable/disable use DNS acquired by DHCP or PPPoE.
	DnsServerOverride *string `pulumi:"dnsServerOverride"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
	Ip *string `pulumi:"ip"`
	// Addressing mode.
	Mode *string `pulumi:"mode"`
	// MTU value for this interface.
	Mtu *string `pulumi:"mtu"`
	// Enable to set a custom MTU for this interface.
	MtuOverride *string `pulumi:"mtuOverride"`
	// Name.
	Name *string `pulumi:"name"`
	// Interface role.
	Role *string `pulumi:"role"`
	// Interface speed. The default setting and the options available depend on the interface hardware.
	Speed *string `pulumi:"speed"`
	// Bring the interface up or shut the interface down.
	Status *string `pulumi:"status"`
	// TCP maximum segment size. 0 means do not change segment size.
	TcpMss *string `pulumi:"tcpMss"`
	// Interface type (support physical, vlan, loopback).
	Type string `pulumi:"type"`
	// Interface is in this virtual domain (VDOM).
	Vdom *string `pulumi:"vdom"`
	// VLAN ID.
	Vlanid *string `pulumi:"vlanid"`
}

// The set of arguments for constructing a InterfacePort resource.
type InterfacePortArgs struct {
	// Alias will be displayed with the interface name to make it easier to distinguish.
	Alias pulumi.StringPtrInput
	// Permitted types of management access to this interface.
	Allowaccess pulumi.StringPtrInput
	// Enable to get the gateway IP from the DHCP or PPPoE server.
	Defaultgw pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
	DeviceIdentification pulumi.StringPtrInput
	// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
	Distance pulumi.StringPtrInput
	// Enable/disable use DNS acquired by DHCP or PPPoE.
	DnsServerOverride pulumi.StringPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
	Ip pulumi.StringPtrInput
	// Addressing mode.
	Mode pulumi.StringPtrInput
	// MTU value for this interface.
	Mtu pulumi.StringPtrInput
	// Enable to set a custom MTU for this interface.
	MtuOverride pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Interface role.
	Role pulumi.StringPtrInput
	// Interface speed. The default setting and the options available depend on the interface hardware.
	Speed pulumi.StringPtrInput
	// Bring the interface up or shut the interface down.
	Status pulumi.StringPtrInput
	// TCP maximum segment size. 0 means do not change segment size.
	TcpMss pulumi.StringPtrInput
	// Interface type (support physical, vlan, loopback).
	Type pulumi.StringInput
	// Interface is in this virtual domain (VDOM).
	Vdom pulumi.StringPtrInput
	// VLAN ID.
	Vlanid pulumi.StringPtrInput
}

func (InterfacePortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*interfacePortArgs)(nil)).Elem()
}

type InterfacePortInput interface {
	pulumi.Input

	ToInterfacePortOutput() InterfacePortOutput
	ToInterfacePortOutputWithContext(ctx context.Context) InterfacePortOutput
}

func (*InterfacePort) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfacePort)(nil)).Elem()
}

func (i *InterfacePort) ToInterfacePortOutput() InterfacePortOutput {
	return i.ToInterfacePortOutputWithContext(context.Background())
}

func (i *InterfacePort) ToInterfacePortOutputWithContext(ctx context.Context) InterfacePortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfacePortOutput)
}

// InterfacePortArrayInput is an input type that accepts InterfacePortArray and InterfacePortArrayOutput values.
// You can construct a concrete instance of `InterfacePortArrayInput` via:
//
//	InterfacePortArray{ InterfacePortArgs{...} }
type InterfacePortArrayInput interface {
	pulumi.Input

	ToInterfacePortArrayOutput() InterfacePortArrayOutput
	ToInterfacePortArrayOutputWithContext(context.Context) InterfacePortArrayOutput
}

type InterfacePortArray []InterfacePortInput

func (InterfacePortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InterfacePort)(nil)).Elem()
}

func (i InterfacePortArray) ToInterfacePortArrayOutput() InterfacePortArrayOutput {
	return i.ToInterfacePortArrayOutputWithContext(context.Background())
}

func (i InterfacePortArray) ToInterfacePortArrayOutputWithContext(ctx context.Context) InterfacePortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfacePortArrayOutput)
}

// InterfacePortMapInput is an input type that accepts InterfacePortMap and InterfacePortMapOutput values.
// You can construct a concrete instance of `InterfacePortMapInput` via:
//
//	InterfacePortMap{ "key": InterfacePortArgs{...} }
type InterfacePortMapInput interface {
	pulumi.Input

	ToInterfacePortMapOutput() InterfacePortMapOutput
	ToInterfacePortMapOutputWithContext(context.Context) InterfacePortMapOutput
}

type InterfacePortMap map[string]InterfacePortInput

func (InterfacePortMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InterfacePort)(nil)).Elem()
}

func (i InterfacePortMap) ToInterfacePortMapOutput() InterfacePortMapOutput {
	return i.ToInterfacePortMapOutputWithContext(context.Background())
}

func (i InterfacePortMap) ToInterfacePortMapOutputWithContext(ctx context.Context) InterfacePortMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfacePortMapOutput)
}

type InterfacePortOutput struct{ *pulumi.OutputState }

func (InterfacePortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfacePort)(nil)).Elem()
}

func (o InterfacePortOutput) ToInterfacePortOutput() InterfacePortOutput {
	return o
}

func (o InterfacePortOutput) ToInterfacePortOutputWithContext(ctx context.Context) InterfacePortOutput {
	return o
}

// Alias will be displayed with the interface name to make it easier to distinguish.
func (o InterfacePortOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// Permitted types of management access to this interface.
func (o InterfacePortOutput) Allowaccess() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Allowaccess }).(pulumi.StringOutput)
}

// Enable to get the gateway IP from the DHCP or PPPoE server.
func (o InterfacePortOutput) Defaultgw() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Defaultgw }).(pulumi.StringOutput)
}

// Description.
func (o InterfacePortOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
func (o InterfacePortOutput) DeviceIdentification() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.DeviceIdentification }).(pulumi.StringOutput)
}

// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
func (o InterfacePortOutput) Distance() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Distance }).(pulumi.StringOutput)
}

// Enable/disable use DNS acquired by DHCP or PPPoE.
func (o InterfacePortOutput) DnsServerOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.DnsServerOverride }).(pulumi.StringOutput)
}

// Interface name.
func (o InterfacePortOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
func (o InterfacePortOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Addressing mode.
func (o InterfacePortOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// MTU value for this interface.
func (o InterfacePortOutput) Mtu() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Mtu }).(pulumi.StringOutput)
}

// Enable to set a custom MTU for this interface.
func (o InterfacePortOutput) MtuOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.MtuOverride }).(pulumi.StringOutput)
}

// Name.
func (o InterfacePortOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Interface role.
func (o InterfacePortOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Interface speed. The default setting and the options available depend on the interface hardware.
func (o InterfacePortOutput) Speed() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Speed }).(pulumi.StringOutput)
}

// Bring the interface up or shut the interface down.
func (o InterfacePortOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// TCP maximum segment size. 0 means do not change segment size.
func (o InterfacePortOutput) TcpMss() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.TcpMss }).(pulumi.StringOutput)
}

// Interface type (support physical, vlan, loopback).
func (o InterfacePortOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Interface is in this virtual domain (VDOM).
func (o InterfacePortOutput) Vdom() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Vdom }).(pulumi.StringOutput)
}

// VLAN ID.
func (o InterfacePortOutput) Vlanid() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfacePort) pulumi.StringOutput { return v.Vlanid }).(pulumi.StringOutput)
}

type InterfacePortArrayOutput struct{ *pulumi.OutputState }

func (InterfacePortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InterfacePort)(nil)).Elem()
}

func (o InterfacePortArrayOutput) ToInterfacePortArrayOutput() InterfacePortArrayOutput {
	return o
}

func (o InterfacePortArrayOutput) ToInterfacePortArrayOutputWithContext(ctx context.Context) InterfacePortArrayOutput {
	return o
}

func (o InterfacePortArrayOutput) Index(i pulumi.IntInput) InterfacePortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InterfacePort {
		return vs[0].([]*InterfacePort)[vs[1].(int)]
	}).(InterfacePortOutput)
}

type InterfacePortMapOutput struct{ *pulumi.OutputState }

func (InterfacePortMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InterfacePort)(nil)).Elem()
}

func (o InterfacePortMapOutput) ToInterfacePortMapOutput() InterfacePortMapOutput {
	return o
}

func (o InterfacePortMapOutput) ToInterfacePortMapOutputWithContext(ctx context.Context) InterfacePortMapOutput {
	return o
}

func (o InterfacePortMapOutput) MapIndex(k pulumi.StringInput) InterfacePortOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InterfacePort {
		return vs[0].(map[string]*InterfacePort)[vs[1].(string)]
	}).(InterfacePortOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InterfacePortInput)(nil)).Elem(), &InterfacePort{})
	pulumi.RegisterInputType(reflect.TypeOf((*InterfacePortArrayInput)(nil)).Elem(), InterfacePortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InterfacePortMapInput)(nil)).Elem(), InterfacePortMap{})
	pulumi.RegisterOutputType(InterfacePortOutput{})
	pulumi.RegisterOutputType(InterfacePortArrayOutput{})
	pulumi.RegisterOutputType(InterfacePortMapOutput{})
}
