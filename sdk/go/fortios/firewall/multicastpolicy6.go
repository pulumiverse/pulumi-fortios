// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPv6 multicast NAT policies.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/firewall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firewall.NewMulticastpolicy6(ctx, "trname", &firewall.Multicastpolicy6Args{
//				Action: pulumi.String("accept"),
//				Dstaddrs: firewall.Multicastpolicy6DstaddrArray{
//					&firewall.Multicastpolicy6DstaddrArgs{
//						Name: pulumi.String("all"),
//					},
//				},
//				Dstintf:    pulumi.String("port4"),
//				EndPort:    pulumi.Int(65535),
//				Fosid:      pulumi.Int(1),
//				Logtraffic: pulumi.String("disable"),
//				Protocol:   pulumi.Int(0),
//				Srcaddrs: firewall.Multicastpolicy6SrcaddrArray{
//					&firewall.Multicastpolicy6SrcaddrArgs{
//						Name: pulumi.String("all"),
//					},
//				},
//				Srcintf:   pulumi.String("port3"),
//				StartPort: pulumi.Int(1),
//				Status:    pulumi.String("enable"),
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
// Firewall MulticastPolicy6 can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/multicastpolicy6:Multicastpolicy6 labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/multicastpolicy6:Multicastpolicy6 labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Multicastpolicy6 struct {
	pulumi.CustomResourceState

	// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
	AutoAsicOffload pulumi.StringOutput `pulumi:"autoAsicOffload"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// IPv6 destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs Multicastpolicy6DstaddrArrayOutput `pulumi:"dstaddrs"`
	// IPv6 destination interface name.
	Dstintf pulumi.StringOutput `pulumi:"dstintf"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
	EndPort pulumi.IntOutput `pulumi:"endPort"`
	// Policy ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Name of an existing IPS sensor.
	IpsSensor pulumi.StringOutput `pulumi:"ipsSensor"`
	// Enable/disable logging traffic accepted by this policy.
	Logtraffic pulumi.StringOutput `pulumi:"logtraffic"`
	// Policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
	Protocol pulumi.IntOutput `pulumi:"protocol"`
	// IPv6 source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs Multicastpolicy6SrcaddrArrayOutput `pulumi:"srcaddrs"`
	// IPv6 source interface name.
	Srcintf pulumi.StringOutput `pulumi:"srcintf"`
	// Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
	StartPort pulumi.IntOutput `pulumi:"startPort"`
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Enable to add an IPS security profile to the policy. Valid values: `enable`, `disable`.
	UtmStatus pulumi.StringOutput `pulumi:"utmStatus"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewMulticastpolicy6 registers a new resource with the given unique name, arguments, and options.
func NewMulticastpolicy6(ctx *pulumi.Context,
	name string, args *Multicastpolicy6Args, opts ...pulumi.ResourceOption) (*Multicastpolicy6, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dstaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Dstaddrs'")
	}
	if args.Dstintf == nil {
		return nil, errors.New("invalid value for required argument 'Dstintf'")
	}
	if args.Srcaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Srcaddrs'")
	}
	if args.Srcintf == nil {
		return nil, errors.New("invalid value for required argument 'Srcintf'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Multicastpolicy6
	err := ctx.RegisterResource("fortios:firewall/multicastpolicy6:Multicastpolicy6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMulticastpolicy6 gets an existing Multicastpolicy6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMulticastpolicy6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Multicastpolicy6State, opts ...pulumi.ResourceOption) (*Multicastpolicy6, error) {
	var resource Multicastpolicy6
	err := ctx.ReadResource("fortios:firewall/multicastpolicy6:Multicastpolicy6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Multicastpolicy6 resources.
type multicastpolicy6State struct {
	// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
	Action *string `pulumi:"action"`
	// Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
	AutoAsicOffload *string `pulumi:"autoAsicOffload"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// IPv6 destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs []Multicastpolicy6Dstaddr `pulumi:"dstaddrs"`
	// IPv6 destination interface name.
	Dstintf *string `pulumi:"dstintf"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
	EndPort *int `pulumi:"endPort"`
	// Policy ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Name of an existing IPS sensor.
	IpsSensor *string `pulumi:"ipsSensor"`
	// Enable/disable logging traffic accepted by this policy.
	Logtraffic *string `pulumi:"logtraffic"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
	Protocol *int `pulumi:"protocol"`
	// IPv6 source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs []Multicastpolicy6Srcaddr `pulumi:"srcaddrs"`
	// IPv6 source interface name.
	Srcintf *string `pulumi:"srcintf"`
	// Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
	StartPort *int `pulumi:"startPort"`
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Enable to add an IPS security profile to the policy. Valid values: `enable`, `disable`.
	UtmStatus *string `pulumi:"utmStatus"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Multicastpolicy6State struct {
	// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
	Action pulumi.StringPtrInput
	// Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
	AutoAsicOffload pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// IPv6 destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs Multicastpolicy6DstaddrArrayInput
	// IPv6 destination interface name.
	Dstintf pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
	EndPort pulumi.IntPtrInput
	// Policy ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Name of an existing IPS sensor.
	IpsSensor pulumi.StringPtrInput
	// Enable/disable logging traffic accepted by this policy.
	Logtraffic pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
	Protocol pulumi.IntPtrInput
	// IPv6 source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs Multicastpolicy6SrcaddrArrayInput
	// IPv6 source interface name.
	Srcintf pulumi.StringPtrInput
	// Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
	StartPort pulumi.IntPtrInput
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Enable to add an IPS security profile to the policy. Valid values: `enable`, `disable`.
	UtmStatus pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Multicastpolicy6State) ElementType() reflect.Type {
	return reflect.TypeOf((*multicastpolicy6State)(nil)).Elem()
}

type multicastpolicy6Args struct {
	// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
	Action *string `pulumi:"action"`
	// Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
	AutoAsicOffload *string `pulumi:"autoAsicOffload"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// IPv6 destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs []Multicastpolicy6Dstaddr `pulumi:"dstaddrs"`
	// IPv6 destination interface name.
	Dstintf string `pulumi:"dstintf"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
	EndPort *int `pulumi:"endPort"`
	// Policy ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Name of an existing IPS sensor.
	IpsSensor *string `pulumi:"ipsSensor"`
	// Enable/disable logging traffic accepted by this policy.
	Logtraffic *string `pulumi:"logtraffic"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
	Protocol *int `pulumi:"protocol"`
	// IPv6 source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs []Multicastpolicy6Srcaddr `pulumi:"srcaddrs"`
	// IPv6 source interface name.
	Srcintf string `pulumi:"srcintf"`
	// Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
	StartPort *int `pulumi:"startPort"`
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Enable to add an IPS security profile to the policy. Valid values: `enable`, `disable`.
	UtmStatus *string `pulumi:"utmStatus"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Multicastpolicy6 resource.
type Multicastpolicy6Args struct {
	// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
	Action pulumi.StringPtrInput
	// Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
	AutoAsicOffload pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// IPv6 destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs Multicastpolicy6DstaddrArrayInput
	// IPv6 destination interface name.
	Dstintf pulumi.StringInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
	EndPort pulumi.IntPtrInput
	// Policy ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Name of an existing IPS sensor.
	IpsSensor pulumi.StringPtrInput
	// Enable/disable logging traffic accepted by this policy.
	Logtraffic pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
	Protocol pulumi.IntPtrInput
	// IPv6 source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs Multicastpolicy6SrcaddrArrayInput
	// IPv6 source interface name.
	Srcintf pulumi.StringInput
	// Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
	StartPort pulumi.IntPtrInput
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Enable to add an IPS security profile to the policy. Valid values: `enable`, `disable`.
	UtmStatus pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Multicastpolicy6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*multicastpolicy6Args)(nil)).Elem()
}

type Multicastpolicy6Input interface {
	pulumi.Input

	ToMulticastpolicy6Output() Multicastpolicy6Output
	ToMulticastpolicy6OutputWithContext(ctx context.Context) Multicastpolicy6Output
}

func (*Multicastpolicy6) ElementType() reflect.Type {
	return reflect.TypeOf((**Multicastpolicy6)(nil)).Elem()
}

func (i *Multicastpolicy6) ToMulticastpolicy6Output() Multicastpolicy6Output {
	return i.ToMulticastpolicy6OutputWithContext(context.Background())
}

func (i *Multicastpolicy6) ToMulticastpolicy6OutputWithContext(ctx context.Context) Multicastpolicy6Output {
	return pulumi.ToOutputWithContext(ctx, i).(Multicastpolicy6Output)
}

// Multicastpolicy6ArrayInput is an input type that accepts Multicastpolicy6Array and Multicastpolicy6ArrayOutput values.
// You can construct a concrete instance of `Multicastpolicy6ArrayInput` via:
//
//	Multicastpolicy6Array{ Multicastpolicy6Args{...} }
type Multicastpolicy6ArrayInput interface {
	pulumi.Input

	ToMulticastpolicy6ArrayOutput() Multicastpolicy6ArrayOutput
	ToMulticastpolicy6ArrayOutputWithContext(context.Context) Multicastpolicy6ArrayOutput
}

type Multicastpolicy6Array []Multicastpolicy6Input

func (Multicastpolicy6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Multicastpolicy6)(nil)).Elem()
}

func (i Multicastpolicy6Array) ToMulticastpolicy6ArrayOutput() Multicastpolicy6ArrayOutput {
	return i.ToMulticastpolicy6ArrayOutputWithContext(context.Background())
}

func (i Multicastpolicy6Array) ToMulticastpolicy6ArrayOutputWithContext(ctx context.Context) Multicastpolicy6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Multicastpolicy6ArrayOutput)
}

// Multicastpolicy6MapInput is an input type that accepts Multicastpolicy6Map and Multicastpolicy6MapOutput values.
// You can construct a concrete instance of `Multicastpolicy6MapInput` via:
//
//	Multicastpolicy6Map{ "key": Multicastpolicy6Args{...} }
type Multicastpolicy6MapInput interface {
	pulumi.Input

	ToMulticastpolicy6MapOutput() Multicastpolicy6MapOutput
	ToMulticastpolicy6MapOutputWithContext(context.Context) Multicastpolicy6MapOutput
}

type Multicastpolicy6Map map[string]Multicastpolicy6Input

func (Multicastpolicy6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Multicastpolicy6)(nil)).Elem()
}

func (i Multicastpolicy6Map) ToMulticastpolicy6MapOutput() Multicastpolicy6MapOutput {
	return i.ToMulticastpolicy6MapOutputWithContext(context.Background())
}

func (i Multicastpolicy6Map) ToMulticastpolicy6MapOutputWithContext(ctx context.Context) Multicastpolicy6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Multicastpolicy6MapOutput)
}

type Multicastpolicy6Output struct{ *pulumi.OutputState }

func (Multicastpolicy6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Multicastpolicy6)(nil)).Elem()
}

func (o Multicastpolicy6Output) ToMulticastpolicy6Output() Multicastpolicy6Output {
	return o
}

func (o Multicastpolicy6Output) ToMulticastpolicy6OutputWithContext(ctx context.Context) Multicastpolicy6Output {
	return o
}

// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
func (o Multicastpolicy6Output) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
func (o Multicastpolicy6Output) AutoAsicOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringOutput { return v.AutoAsicOffload }).(pulumi.StringOutput)
}

// Comment.
func (o Multicastpolicy6Output) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// IPv6 destination address name. The structure of `dstaddr` block is documented below.
func (o Multicastpolicy6Output) Dstaddrs() Multicastpolicy6DstaddrArrayOutput {
	return o.ApplyT(func(v *Multicastpolicy6) Multicastpolicy6DstaddrArrayOutput { return v.Dstaddrs }).(Multicastpolicy6DstaddrArrayOutput)
}

// IPv6 destination interface name.
func (o Multicastpolicy6Output) Dstintf() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringOutput { return v.Dstintf }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Multicastpolicy6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
func (o Multicastpolicy6Output) EndPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.IntOutput { return v.EndPort }).(pulumi.IntOutput)
}

// Policy ID.
func (o Multicastpolicy6Output) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o Multicastpolicy6Output) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Name of an existing IPS sensor.
func (o Multicastpolicy6Output) IpsSensor() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringOutput { return v.IpsSensor }).(pulumi.StringOutput)
}

// Enable/disable logging traffic accepted by this policy.
func (o Multicastpolicy6Output) Logtraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringOutput { return v.Logtraffic }).(pulumi.StringOutput)
}

// Policy name.
func (o Multicastpolicy6Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
func (o Multicastpolicy6Output) Protocol() pulumi.IntOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.IntOutput { return v.Protocol }).(pulumi.IntOutput)
}

// IPv6 source address name. The structure of `srcaddr` block is documented below.
func (o Multicastpolicy6Output) Srcaddrs() Multicastpolicy6SrcaddrArrayOutput {
	return o.ApplyT(func(v *Multicastpolicy6) Multicastpolicy6SrcaddrArrayOutput { return v.Srcaddrs }).(Multicastpolicy6SrcaddrArrayOutput)
}

// IPv6 source interface name.
func (o Multicastpolicy6Output) Srcintf() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringOutput { return v.Srcintf }).(pulumi.StringOutput)
}

// Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
func (o Multicastpolicy6Output) StartPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.IntOutput { return v.StartPort }).(pulumi.IntOutput)
}

// Enable/disable this policy. Valid values: `enable`, `disable`.
func (o Multicastpolicy6Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Enable to add an IPS security profile to the policy. Valid values: `enable`, `disable`.
func (o Multicastpolicy6Output) UtmStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringOutput { return v.UtmStatus }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o Multicastpolicy6Output) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Multicastpolicy6Output) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastpolicy6) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type Multicastpolicy6ArrayOutput struct{ *pulumi.OutputState }

func (Multicastpolicy6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Multicastpolicy6)(nil)).Elem()
}

func (o Multicastpolicy6ArrayOutput) ToMulticastpolicy6ArrayOutput() Multicastpolicy6ArrayOutput {
	return o
}

func (o Multicastpolicy6ArrayOutput) ToMulticastpolicy6ArrayOutputWithContext(ctx context.Context) Multicastpolicy6ArrayOutput {
	return o
}

func (o Multicastpolicy6ArrayOutput) Index(i pulumi.IntInput) Multicastpolicy6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Multicastpolicy6 {
		return vs[0].([]*Multicastpolicy6)[vs[1].(int)]
	}).(Multicastpolicy6Output)
}

type Multicastpolicy6MapOutput struct{ *pulumi.OutputState }

func (Multicastpolicy6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Multicastpolicy6)(nil)).Elem()
}

func (o Multicastpolicy6MapOutput) ToMulticastpolicy6MapOutput() Multicastpolicy6MapOutput {
	return o
}

func (o Multicastpolicy6MapOutput) ToMulticastpolicy6MapOutputWithContext(ctx context.Context) Multicastpolicy6MapOutput {
	return o
}

func (o Multicastpolicy6MapOutput) MapIndex(k pulumi.StringInput) Multicastpolicy6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Multicastpolicy6 {
		return vs[0].(map[string]*Multicastpolicy6)[vs[1].(string)]
	}).(Multicastpolicy6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Multicastpolicy6Input)(nil)).Elem(), &Multicastpolicy6{})
	pulumi.RegisterInputType(reflect.TypeOf((*Multicastpolicy6ArrayInput)(nil)).Elem(), Multicastpolicy6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Multicastpolicy6MapInput)(nil)).Elem(), Multicastpolicy6Map{})
	pulumi.RegisterOutputType(Multicastpolicy6Output{})
	pulumi.RegisterOutputType(Multicastpolicy6ArrayOutput{})
	pulumi.RegisterOutputType(Multicastpolicy6MapOutput{})
}
