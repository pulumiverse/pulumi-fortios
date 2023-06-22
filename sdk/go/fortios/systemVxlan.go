// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure VXLAN devices.
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
//			_, err := fortios.NewSystemVxlan(ctx, "trname", &fortios.SystemVxlanArgs{
//				Dstport:   pulumi.Int(4789),
//				Interface: pulumi.String("port3"),
//				IpVersion: pulumi.String("ipv4-unicast"),
//				RemoteIps: fortios.SystemVxlanRemoteIpArray{
//					&fortios.SystemVxlanRemoteIpArgs{
//						Ip: pulumi.String("1.1.1.1"),
//					},
//				},
//				Vni: pulumi.Int(3),
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
// # System Vxlan can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemVxlan:SystemVxlan labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemVxlan:SystemVxlan labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemVxlan struct {
	pulumi.CustomResourceState

	// VXLAN destination port (1 - 65535, default = 4789).
	Dstport pulumi.IntOutput `pulumi:"dstport"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Outgoing interface for VXLAN encapsulated traffic.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast. Valid values: `ipv4-unicast`, `ipv6-unicast`, `ipv4-multicast`, `ipv6-multicast`.
	IpVersion pulumi.StringOutput `pulumi:"ipVersion"`
	// VXLAN multicast TTL (1-255, default = 0).
	MulticastTtl pulumi.IntOutput `pulumi:"multicastTtl"`
	// VXLAN device or interface name. Must be a unique interface name.
	Name pulumi.StringOutput `pulumi:"name"`
	// IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp6` block is documented below.
	RemoteIp6s SystemVxlanRemoteIp6ArrayOutput `pulumi:"remoteIp6s"`
	// IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp` block is documented below.
	RemoteIps SystemVxlanRemoteIpArrayOutput `pulumi:"remoteIps"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// VXLAN network ID.
	Vni pulumi.IntOutput `pulumi:"vni"`
}

// NewSystemVxlan registers a new resource with the given unique name, arguments, and options.
func NewSystemVxlan(ctx *pulumi.Context,
	name string, args *SystemVxlanArgs, opts ...pulumi.ResourceOption) (*SystemVxlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.IpVersion == nil {
		return nil, errors.New("invalid value for required argument 'IpVersion'")
	}
	if args.Vni == nil {
		return nil, errors.New("invalid value for required argument 'Vni'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemVxlan
	err := ctx.RegisterResource("fortios:index/systemVxlan:SystemVxlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemVxlan gets an existing SystemVxlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemVxlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemVxlanState, opts ...pulumi.ResourceOption) (*SystemVxlan, error) {
	var resource SystemVxlan
	err := ctx.ReadResource("fortios:index/systemVxlan:SystemVxlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemVxlan resources.
type systemVxlanState struct {
	// VXLAN destination port (1 - 65535, default = 4789).
	Dstport *int `pulumi:"dstport"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Outgoing interface for VXLAN encapsulated traffic.
	Interface *string `pulumi:"interface"`
	// IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast. Valid values: `ipv4-unicast`, `ipv6-unicast`, `ipv4-multicast`, `ipv6-multicast`.
	IpVersion *string `pulumi:"ipVersion"`
	// VXLAN multicast TTL (1-255, default = 0).
	MulticastTtl *int `pulumi:"multicastTtl"`
	// VXLAN device or interface name. Must be a unique interface name.
	Name *string `pulumi:"name"`
	// IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp6` block is documented below.
	RemoteIp6s []SystemVxlanRemoteIp6 `pulumi:"remoteIp6s"`
	// IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp` block is documented below.
	RemoteIps []SystemVxlanRemoteIp `pulumi:"remoteIps"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VXLAN network ID.
	Vni *int `pulumi:"vni"`
}

type SystemVxlanState struct {
	// VXLAN destination port (1 - 65535, default = 4789).
	Dstport pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Outgoing interface for VXLAN encapsulated traffic.
	Interface pulumi.StringPtrInput
	// IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast. Valid values: `ipv4-unicast`, `ipv6-unicast`, `ipv4-multicast`, `ipv6-multicast`.
	IpVersion pulumi.StringPtrInput
	// VXLAN multicast TTL (1-255, default = 0).
	MulticastTtl pulumi.IntPtrInput
	// VXLAN device or interface name. Must be a unique interface name.
	Name pulumi.StringPtrInput
	// IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp6` block is documented below.
	RemoteIp6s SystemVxlanRemoteIp6ArrayInput
	// IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp` block is documented below.
	RemoteIps SystemVxlanRemoteIpArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VXLAN network ID.
	Vni pulumi.IntPtrInput
}

func (SystemVxlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVxlanState)(nil)).Elem()
}

type systemVxlanArgs struct {
	// VXLAN destination port (1 - 65535, default = 4789).
	Dstport *int `pulumi:"dstport"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Outgoing interface for VXLAN encapsulated traffic.
	Interface string `pulumi:"interface"`
	// IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast. Valid values: `ipv4-unicast`, `ipv6-unicast`, `ipv4-multicast`, `ipv6-multicast`.
	IpVersion string `pulumi:"ipVersion"`
	// VXLAN multicast TTL (1-255, default = 0).
	MulticastTtl *int `pulumi:"multicastTtl"`
	// VXLAN device or interface name. Must be a unique interface name.
	Name *string `pulumi:"name"`
	// IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp6` block is documented below.
	RemoteIp6s []SystemVxlanRemoteIp6 `pulumi:"remoteIp6s"`
	// IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp` block is documented below.
	RemoteIps []SystemVxlanRemoteIp `pulumi:"remoteIps"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VXLAN network ID.
	Vni int `pulumi:"vni"`
}

// The set of arguments for constructing a SystemVxlan resource.
type SystemVxlanArgs struct {
	// VXLAN destination port (1 - 65535, default = 4789).
	Dstport pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Outgoing interface for VXLAN encapsulated traffic.
	Interface pulumi.StringInput
	// IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast. Valid values: `ipv4-unicast`, `ipv6-unicast`, `ipv4-multicast`, `ipv6-multicast`.
	IpVersion pulumi.StringInput
	// VXLAN multicast TTL (1-255, default = 0).
	MulticastTtl pulumi.IntPtrInput
	// VXLAN device or interface name. Must be a unique interface name.
	Name pulumi.StringPtrInput
	// IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp6` block is documented below.
	RemoteIp6s SystemVxlanRemoteIp6ArrayInput
	// IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp` block is documented below.
	RemoteIps SystemVxlanRemoteIpArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VXLAN network ID.
	Vni pulumi.IntInput
}

func (SystemVxlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVxlanArgs)(nil)).Elem()
}

type SystemVxlanInput interface {
	pulumi.Input

	ToSystemVxlanOutput() SystemVxlanOutput
	ToSystemVxlanOutputWithContext(ctx context.Context) SystemVxlanOutput
}

func (*SystemVxlan) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVxlan)(nil)).Elem()
}

func (i *SystemVxlan) ToSystemVxlanOutput() SystemVxlanOutput {
	return i.ToSystemVxlanOutputWithContext(context.Background())
}

func (i *SystemVxlan) ToSystemVxlanOutputWithContext(ctx context.Context) SystemVxlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVxlanOutput)
}

// SystemVxlanArrayInput is an input type that accepts SystemVxlanArray and SystemVxlanArrayOutput values.
// You can construct a concrete instance of `SystemVxlanArrayInput` via:
//
//	SystemVxlanArray{ SystemVxlanArgs{...} }
type SystemVxlanArrayInput interface {
	pulumi.Input

	ToSystemVxlanArrayOutput() SystemVxlanArrayOutput
	ToSystemVxlanArrayOutputWithContext(context.Context) SystemVxlanArrayOutput
}

type SystemVxlanArray []SystemVxlanInput

func (SystemVxlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVxlan)(nil)).Elem()
}

func (i SystemVxlanArray) ToSystemVxlanArrayOutput() SystemVxlanArrayOutput {
	return i.ToSystemVxlanArrayOutputWithContext(context.Background())
}

func (i SystemVxlanArray) ToSystemVxlanArrayOutputWithContext(ctx context.Context) SystemVxlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVxlanArrayOutput)
}

// SystemVxlanMapInput is an input type that accepts SystemVxlanMap and SystemVxlanMapOutput values.
// You can construct a concrete instance of `SystemVxlanMapInput` via:
//
//	SystemVxlanMap{ "key": SystemVxlanArgs{...} }
type SystemVxlanMapInput interface {
	pulumi.Input

	ToSystemVxlanMapOutput() SystemVxlanMapOutput
	ToSystemVxlanMapOutputWithContext(context.Context) SystemVxlanMapOutput
}

type SystemVxlanMap map[string]SystemVxlanInput

func (SystemVxlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVxlan)(nil)).Elem()
}

func (i SystemVxlanMap) ToSystemVxlanMapOutput() SystemVxlanMapOutput {
	return i.ToSystemVxlanMapOutputWithContext(context.Background())
}

func (i SystemVxlanMap) ToSystemVxlanMapOutputWithContext(ctx context.Context) SystemVxlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVxlanMapOutput)
}

type SystemVxlanOutput struct{ *pulumi.OutputState }

func (SystemVxlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVxlan)(nil)).Elem()
}

func (o SystemVxlanOutput) ToSystemVxlanOutput() SystemVxlanOutput {
	return o
}

func (o SystemVxlanOutput) ToSystemVxlanOutputWithContext(ctx context.Context) SystemVxlanOutput {
	return o
}

// VXLAN destination port (1 - 65535, default = 4789).
func (o SystemVxlanOutput) Dstport() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemVxlan) pulumi.IntOutput { return v.Dstport }).(pulumi.IntOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SystemVxlanOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemVxlan) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Outgoing interface for VXLAN encapsulated traffic.
func (o SystemVxlanOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVxlan) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast. Valid values: `ipv4-unicast`, `ipv6-unicast`, `ipv4-multicast`, `ipv6-multicast`.
func (o SystemVxlanOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVxlan) pulumi.StringOutput { return v.IpVersion }).(pulumi.StringOutput)
}

// VXLAN multicast TTL (1-255, default = 0).
func (o SystemVxlanOutput) MulticastTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemVxlan) pulumi.IntOutput { return v.MulticastTtl }).(pulumi.IntOutput)
}

// VXLAN device or interface name. Must be a unique interface name.
func (o SystemVxlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVxlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp6` block is documented below.
func (o SystemVxlanOutput) RemoteIp6s() SystemVxlanRemoteIp6ArrayOutput {
	return o.ApplyT(func(v *SystemVxlan) SystemVxlanRemoteIp6ArrayOutput { return v.RemoteIp6s }).(SystemVxlanRemoteIp6ArrayOutput)
}

// IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp` block is documented below.
func (o SystemVxlanOutput) RemoteIps() SystemVxlanRemoteIpArrayOutput {
	return o.ApplyT(func(v *SystemVxlan) SystemVxlanRemoteIpArrayOutput { return v.RemoteIps }).(SystemVxlanRemoteIpArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemVxlanOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemVxlan) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// VXLAN network ID.
func (o SystemVxlanOutput) Vni() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemVxlan) pulumi.IntOutput { return v.Vni }).(pulumi.IntOutput)
}

type SystemVxlanArrayOutput struct{ *pulumi.OutputState }

func (SystemVxlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVxlan)(nil)).Elem()
}

func (o SystemVxlanArrayOutput) ToSystemVxlanArrayOutput() SystemVxlanArrayOutput {
	return o
}

func (o SystemVxlanArrayOutput) ToSystemVxlanArrayOutputWithContext(ctx context.Context) SystemVxlanArrayOutput {
	return o
}

func (o SystemVxlanArrayOutput) Index(i pulumi.IntInput) SystemVxlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemVxlan {
		return vs[0].([]*SystemVxlan)[vs[1].(int)]
	}).(SystemVxlanOutput)
}

type SystemVxlanMapOutput struct{ *pulumi.OutputState }

func (SystemVxlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVxlan)(nil)).Elem()
}

func (o SystemVxlanMapOutput) ToSystemVxlanMapOutput() SystemVxlanMapOutput {
	return o
}

func (o SystemVxlanMapOutput) ToSystemVxlanMapOutputWithContext(ctx context.Context) SystemVxlanMapOutput {
	return o
}

func (o SystemVxlanMapOutput) MapIndex(k pulumi.StringInput) SystemVxlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemVxlan {
		return vs[0].(map[string]*SystemVxlan)[vs[1].(string)]
	}).(SystemVxlanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVxlanInput)(nil)).Elem(), &SystemVxlan{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVxlanArrayInput)(nil)).Elem(), SystemVxlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVxlanMapInput)(nil)).Elem(), SystemVxlanMap{})
	pulumi.RegisterOutputType(SystemVxlanOutput{})
	pulumi.RegisterOutputType(SystemVxlanArrayOutput{})
	pulumi.RegisterOutputType(SystemVxlanMapOutput{})
}
