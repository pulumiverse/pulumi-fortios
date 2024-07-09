// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPv4 static routing tables.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/router"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := router.NewStatic(ctx, "trname", &router.StaticArgs{
//				Bfd:               pulumi.String("disable"),
//				Blackhole:         pulumi.String("disable"),
//				Device:            pulumi.String("port4"),
//				Distance:          pulumi.Int(10),
//				Dst:               pulumi.String("1.0.0.0 255.240.0.0"),
//				DynamicGateway:    pulumi.String("disable"),
//				Gateway:           pulumi.String("0.0.0.0"),
//				InternetService:   pulumi.Int(0),
//				LinkMonitorExempt: pulumi.String("disable"),
//				Priority:          pulumi.Int(22),
//				SeqNum:            pulumi.Int(1),
//				Src:               pulumi.String("0.0.0.0 0.0.0.0"),
//				Status:            pulumi.String("enable"),
//				VirtualWanLink:    pulumi.String("disable"),
//				Vrf:               pulumi.Int(0),
//				Weight:            pulumi.Int(2),
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
// Router Static can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:router/static:Static labelname {{seq_num}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:router/static:Static labelname {{seq_num}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Static struct {
	pulumi.CustomResourceState

	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
	Bfd pulumi.StringOutput `pulumi:"bfd"`
	// Enable/disable black hole. Valid values: `enable`, `disable`.
	Blackhole pulumi.StringOutput `pulumi:"blackhole"`
	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Gateway out interface or tunnel.
	Device pulumi.StringOutput `pulumi:"device"`
	// Administrative distance (1 - 255).
	Distance pulumi.IntOutput `pulumi:"distance"`
	// Destination IP and mask for this route.
	Dst pulumi.StringOutput `pulumi:"dst"`
	// Name of firewall address or address group.
	Dstaddr pulumi.StringOutput `pulumi:"dstaddr"`
	// Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `enable`, `disable`.
	DynamicGateway pulumi.StringOutput `pulumi:"dynamicGateway"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Gateway IP for this route.
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Application ID in the Internet service database.
	InternetService pulumi.IntOutput `pulumi:"internetService"`
	// Application name in the Internet service custom database.
	InternetServiceCustom pulumi.StringOutput `pulumi:"internetServiceCustom"`
	// Enable/disable withdrawing this route when link monitor or health check is down. Valid values: `enable`, `disable`.
	LinkMonitorExempt pulumi.StringOutput `pulumi:"linkMonitorExempt"`
	// Preferred source IP for this route.
	PreferredSource pulumi.StringOutput `pulumi:"preferredSource"`
	// Administrative priority. On FortiOS versions 6.2.0-6.4.1: 0 - 4294967295. On FortiOS versions 6.4.2-7.0.3: 0 - 65535. On FortiOS versions >= 7.0.4: 1 - 65535.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Enable/disable egress through SD-WAN. Valid values: `enable`, `disable`.
	Sdwan pulumi.StringOutput `pulumi:"sdwan"`
	// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
	SdwanZones StaticSdwanZoneArrayOutput `pulumi:"sdwanZones"`
	// Sequence number.
	SeqNum pulumi.IntOutput `pulumi:"seqNum"`
	// Source prefix for this route.
	Src pulumi.StringOutput `pulumi:"src"`
	// Enable/disable this static route. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Route tag.
	Tag pulumi.IntOutput `pulumi:"tag"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
	// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
	VirtualWanLink pulumi.StringOutput `pulumi:"virtualWanLink"`
	// Virtual Routing Forwarding ID.
	Vrf pulumi.IntOutput `pulumi:"vrf"`
	// Administrative weight (0 - 255).
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewStatic registers a new resource with the given unique name, arguments, and options.
func NewStatic(ctx *pulumi.Context,
	name string, args *StaticArgs, opts ...pulumi.ResourceOption) (*Static, error) {
	if args == nil {
		args = &StaticArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Static
	err := ctx.RegisterResource("fortios:router/static:Static", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStatic gets an existing Static resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStatic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticState, opts ...pulumi.ResourceOption) (*Static, error) {
	var resource Static
	err := ctx.ReadResource("fortios:router/static:Static", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Static resources.
type staticState struct {
	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
	Bfd *string `pulumi:"bfd"`
	// Enable/disable black hole. Valid values: `enable`, `disable`.
	Blackhole *string `pulumi:"blackhole"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Gateway out interface or tunnel.
	Device *string `pulumi:"device"`
	// Administrative distance (1 - 255).
	Distance *int `pulumi:"distance"`
	// Destination IP and mask for this route.
	Dst *string `pulumi:"dst"`
	// Name of firewall address or address group.
	Dstaddr *string `pulumi:"dstaddr"`
	// Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `enable`, `disable`.
	DynamicGateway *string `pulumi:"dynamicGateway"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Gateway IP for this route.
	Gateway *string `pulumi:"gateway"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Application ID in the Internet service database.
	InternetService *int `pulumi:"internetService"`
	// Application name in the Internet service custom database.
	InternetServiceCustom *string `pulumi:"internetServiceCustom"`
	// Enable/disable withdrawing this route when link monitor or health check is down. Valid values: `enable`, `disable`.
	LinkMonitorExempt *string `pulumi:"linkMonitorExempt"`
	// Preferred source IP for this route.
	PreferredSource *string `pulumi:"preferredSource"`
	// Administrative priority. On FortiOS versions 6.2.0-6.4.1: 0 - 4294967295. On FortiOS versions 6.4.2-7.0.3: 0 - 65535. On FortiOS versions >= 7.0.4: 1 - 65535.
	Priority *int `pulumi:"priority"`
	// Enable/disable egress through SD-WAN. Valid values: `enable`, `disable`.
	Sdwan *string `pulumi:"sdwan"`
	// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
	SdwanZones []StaticSdwanZone `pulumi:"sdwanZones"`
	// Sequence number.
	SeqNum *int `pulumi:"seqNum"`
	// Source prefix for this route.
	Src *string `pulumi:"src"`
	// Enable/disable this static route. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Route tag.
	Tag *int `pulumi:"tag"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
	VirtualWanLink *string `pulumi:"virtualWanLink"`
	// Virtual Routing Forwarding ID.
	Vrf *int `pulumi:"vrf"`
	// Administrative weight (0 - 255).
	Weight *int `pulumi:"weight"`
}

type StaticState struct {
	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
	Bfd pulumi.StringPtrInput
	// Enable/disable black hole. Valid values: `enable`, `disable`.
	Blackhole pulumi.StringPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Gateway out interface or tunnel.
	Device pulumi.StringPtrInput
	// Administrative distance (1 - 255).
	Distance pulumi.IntPtrInput
	// Destination IP and mask for this route.
	Dst pulumi.StringPtrInput
	// Name of firewall address or address group.
	Dstaddr pulumi.StringPtrInput
	// Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `enable`, `disable`.
	DynamicGateway pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Gateway IP for this route.
	Gateway pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Application ID in the Internet service database.
	InternetService pulumi.IntPtrInput
	// Application name in the Internet service custom database.
	InternetServiceCustom pulumi.StringPtrInput
	// Enable/disable withdrawing this route when link monitor or health check is down. Valid values: `enable`, `disable`.
	LinkMonitorExempt pulumi.StringPtrInput
	// Preferred source IP for this route.
	PreferredSource pulumi.StringPtrInput
	// Administrative priority. On FortiOS versions 6.2.0-6.4.1: 0 - 4294967295. On FortiOS versions 6.4.2-7.0.3: 0 - 65535. On FortiOS versions >= 7.0.4: 1 - 65535.
	Priority pulumi.IntPtrInput
	// Enable/disable egress through SD-WAN. Valid values: `enable`, `disable`.
	Sdwan pulumi.StringPtrInput
	// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
	SdwanZones StaticSdwanZoneArrayInput
	// Sequence number.
	SeqNum pulumi.IntPtrInput
	// Source prefix for this route.
	Src pulumi.StringPtrInput
	// Enable/disable this static route. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Route tag.
	Tag pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
	VirtualWanLink pulumi.StringPtrInput
	// Virtual Routing Forwarding ID.
	Vrf pulumi.IntPtrInput
	// Administrative weight (0 - 255).
	Weight pulumi.IntPtrInput
}

func (StaticState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticState)(nil)).Elem()
}

type staticArgs struct {
	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
	Bfd *string `pulumi:"bfd"`
	// Enable/disable black hole. Valid values: `enable`, `disable`.
	Blackhole *string `pulumi:"blackhole"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Gateway out interface or tunnel.
	Device *string `pulumi:"device"`
	// Administrative distance (1 - 255).
	Distance *int `pulumi:"distance"`
	// Destination IP and mask for this route.
	Dst *string `pulumi:"dst"`
	// Name of firewall address or address group.
	Dstaddr *string `pulumi:"dstaddr"`
	// Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `enable`, `disable`.
	DynamicGateway *string `pulumi:"dynamicGateway"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Gateway IP for this route.
	Gateway *string `pulumi:"gateway"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Application ID in the Internet service database.
	InternetService *int `pulumi:"internetService"`
	// Application name in the Internet service custom database.
	InternetServiceCustom *string `pulumi:"internetServiceCustom"`
	// Enable/disable withdrawing this route when link monitor or health check is down. Valid values: `enable`, `disable`.
	LinkMonitorExempt *string `pulumi:"linkMonitorExempt"`
	// Preferred source IP for this route.
	PreferredSource *string `pulumi:"preferredSource"`
	// Administrative priority. On FortiOS versions 6.2.0-6.4.1: 0 - 4294967295. On FortiOS versions 6.4.2-7.0.3: 0 - 65535. On FortiOS versions >= 7.0.4: 1 - 65535.
	Priority *int `pulumi:"priority"`
	// Enable/disable egress through SD-WAN. Valid values: `enable`, `disable`.
	Sdwan *string `pulumi:"sdwan"`
	// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
	SdwanZones []StaticSdwanZone `pulumi:"sdwanZones"`
	// Sequence number.
	SeqNum *int `pulumi:"seqNum"`
	// Source prefix for this route.
	Src *string `pulumi:"src"`
	// Enable/disable this static route. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Route tag.
	Tag *int `pulumi:"tag"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
	VirtualWanLink *string `pulumi:"virtualWanLink"`
	// Virtual Routing Forwarding ID.
	Vrf *int `pulumi:"vrf"`
	// Administrative weight (0 - 255).
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a Static resource.
type StaticArgs struct {
	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
	Bfd pulumi.StringPtrInput
	// Enable/disable black hole. Valid values: `enable`, `disable`.
	Blackhole pulumi.StringPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Gateway out interface or tunnel.
	Device pulumi.StringPtrInput
	// Administrative distance (1 - 255).
	Distance pulumi.IntPtrInput
	// Destination IP and mask for this route.
	Dst pulumi.StringPtrInput
	// Name of firewall address or address group.
	Dstaddr pulumi.StringPtrInput
	// Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `enable`, `disable`.
	DynamicGateway pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Gateway IP for this route.
	Gateway pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Application ID in the Internet service database.
	InternetService pulumi.IntPtrInput
	// Application name in the Internet service custom database.
	InternetServiceCustom pulumi.StringPtrInput
	// Enable/disable withdrawing this route when link monitor or health check is down. Valid values: `enable`, `disable`.
	LinkMonitorExempt pulumi.StringPtrInput
	// Preferred source IP for this route.
	PreferredSource pulumi.StringPtrInput
	// Administrative priority. On FortiOS versions 6.2.0-6.4.1: 0 - 4294967295. On FortiOS versions 6.4.2-7.0.3: 0 - 65535. On FortiOS versions >= 7.0.4: 1 - 65535.
	Priority pulumi.IntPtrInput
	// Enable/disable egress through SD-WAN. Valid values: `enable`, `disable`.
	Sdwan pulumi.StringPtrInput
	// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
	SdwanZones StaticSdwanZoneArrayInput
	// Sequence number.
	SeqNum pulumi.IntPtrInput
	// Source prefix for this route.
	Src pulumi.StringPtrInput
	// Enable/disable this static route. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Route tag.
	Tag pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
	VirtualWanLink pulumi.StringPtrInput
	// Virtual Routing Forwarding ID.
	Vrf pulumi.IntPtrInput
	// Administrative weight (0 - 255).
	Weight pulumi.IntPtrInput
}

func (StaticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticArgs)(nil)).Elem()
}

type StaticInput interface {
	pulumi.Input

	ToStaticOutput() StaticOutput
	ToStaticOutputWithContext(ctx context.Context) StaticOutput
}

func (*Static) ElementType() reflect.Type {
	return reflect.TypeOf((**Static)(nil)).Elem()
}

func (i *Static) ToStaticOutput() StaticOutput {
	return i.ToStaticOutputWithContext(context.Background())
}

func (i *Static) ToStaticOutputWithContext(ctx context.Context) StaticOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticOutput)
}

// StaticArrayInput is an input type that accepts StaticArray and StaticArrayOutput values.
// You can construct a concrete instance of `StaticArrayInput` via:
//
//	StaticArray{ StaticArgs{...} }
type StaticArrayInput interface {
	pulumi.Input

	ToStaticArrayOutput() StaticArrayOutput
	ToStaticArrayOutputWithContext(context.Context) StaticArrayOutput
}

type StaticArray []StaticInput

func (StaticArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Static)(nil)).Elem()
}

func (i StaticArray) ToStaticArrayOutput() StaticArrayOutput {
	return i.ToStaticArrayOutputWithContext(context.Background())
}

func (i StaticArray) ToStaticArrayOutputWithContext(ctx context.Context) StaticArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticArrayOutput)
}

// StaticMapInput is an input type that accepts StaticMap and StaticMapOutput values.
// You can construct a concrete instance of `StaticMapInput` via:
//
//	StaticMap{ "key": StaticArgs{...} }
type StaticMapInput interface {
	pulumi.Input

	ToStaticMapOutput() StaticMapOutput
	ToStaticMapOutputWithContext(context.Context) StaticMapOutput
}

type StaticMap map[string]StaticInput

func (StaticMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Static)(nil)).Elem()
}

func (i StaticMap) ToStaticMapOutput() StaticMapOutput {
	return i.ToStaticMapOutputWithContext(context.Background())
}

func (i StaticMap) ToStaticMapOutputWithContext(ctx context.Context) StaticMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticMapOutput)
}

type StaticOutput struct{ *pulumi.OutputState }

func (StaticOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Static)(nil)).Elem()
}

func (o StaticOutput) ToStaticOutput() StaticOutput {
	return o
}

func (o StaticOutput) ToStaticOutputWithContext(ctx context.Context) StaticOutput {
	return o
}

// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
func (o StaticOutput) Bfd() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.Bfd }).(pulumi.StringOutput)
}

// Enable/disable black hole. Valid values: `enable`, `disable`.
func (o StaticOutput) Blackhole() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.Blackhole }).(pulumi.StringOutput)
}

// Optional comments.
func (o StaticOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Static) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Gateway out interface or tunnel.
func (o StaticOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.Device }).(pulumi.StringOutput)
}

// Administrative distance (1 - 255).
func (o StaticOutput) Distance() pulumi.IntOutput {
	return o.ApplyT(func(v *Static) pulumi.IntOutput { return v.Distance }).(pulumi.IntOutput)
}

// Destination IP and mask for this route.
func (o StaticOutput) Dst() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.Dst }).(pulumi.StringOutput)
}

// Name of firewall address or address group.
func (o StaticOutput) Dstaddr() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.Dstaddr }).(pulumi.StringOutput)
}

// Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `enable`, `disable`.
func (o StaticOutput) DynamicGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.DynamicGateway }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o StaticOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Static) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Gateway IP for this route.
func (o StaticOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o StaticOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Static) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Application ID in the Internet service database.
func (o StaticOutput) InternetService() pulumi.IntOutput {
	return o.ApplyT(func(v *Static) pulumi.IntOutput { return v.InternetService }).(pulumi.IntOutput)
}

// Application name in the Internet service custom database.
func (o StaticOutput) InternetServiceCustom() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.InternetServiceCustom }).(pulumi.StringOutput)
}

// Enable/disable withdrawing this route when link monitor or health check is down. Valid values: `enable`, `disable`.
func (o StaticOutput) LinkMonitorExempt() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.LinkMonitorExempt }).(pulumi.StringOutput)
}

// Preferred source IP for this route.
func (o StaticOutput) PreferredSource() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.PreferredSource }).(pulumi.StringOutput)
}

// Administrative priority. On FortiOS versions 6.2.0-6.4.1: 0 - 4294967295. On FortiOS versions 6.4.2-7.0.3: 0 - 65535. On FortiOS versions >= 7.0.4: 1 - 65535.
func (o StaticOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *Static) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// Enable/disable egress through SD-WAN. Valid values: `enable`, `disable`.
func (o StaticOutput) Sdwan() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.Sdwan }).(pulumi.StringOutput)
}

// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
func (o StaticOutput) SdwanZones() StaticSdwanZoneArrayOutput {
	return o.ApplyT(func(v *Static) StaticSdwanZoneArrayOutput { return v.SdwanZones }).(StaticSdwanZoneArrayOutput)
}

// Sequence number.
func (o StaticOutput) SeqNum() pulumi.IntOutput {
	return o.ApplyT(func(v *Static) pulumi.IntOutput { return v.SeqNum }).(pulumi.IntOutput)
}

// Source prefix for this route.
func (o StaticOutput) Src() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.Src }).(pulumi.StringOutput)
}

// Enable/disable this static route. Valid values: `enable`, `disable`.
func (o StaticOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Route tag.
func (o StaticOutput) Tag() pulumi.IntOutput {
	return o.ApplyT(func(v *Static) pulumi.IntOutput { return v.Tag }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o StaticOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
func (o StaticOutput) VirtualWanLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Static) pulumi.StringOutput { return v.VirtualWanLink }).(pulumi.StringOutput)
}

// Virtual Routing Forwarding ID.
func (o StaticOutput) Vrf() pulumi.IntOutput {
	return o.ApplyT(func(v *Static) pulumi.IntOutput { return v.Vrf }).(pulumi.IntOutput)
}

// Administrative weight (0 - 255).
func (o StaticOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v *Static) pulumi.IntOutput { return v.Weight }).(pulumi.IntOutput)
}

type StaticArrayOutput struct{ *pulumi.OutputState }

func (StaticArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Static)(nil)).Elem()
}

func (o StaticArrayOutput) ToStaticArrayOutput() StaticArrayOutput {
	return o
}

func (o StaticArrayOutput) ToStaticArrayOutputWithContext(ctx context.Context) StaticArrayOutput {
	return o
}

func (o StaticArrayOutput) Index(i pulumi.IntInput) StaticOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Static {
		return vs[0].([]*Static)[vs[1].(int)]
	}).(StaticOutput)
}

type StaticMapOutput struct{ *pulumi.OutputState }

func (StaticMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Static)(nil)).Elem()
}

func (o StaticMapOutput) ToStaticMapOutput() StaticMapOutput {
	return o
}

func (o StaticMapOutput) ToStaticMapOutputWithContext(ctx context.Context) StaticMapOutput {
	return o
}

func (o StaticMapOutput) MapIndex(k pulumi.StringInput) StaticOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Static {
		return vs[0].(map[string]*Static)[vs[1].(string)]
	}).(StaticOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StaticInput)(nil)).Elem(), &Static{})
	pulumi.RegisterInputType(reflect.TypeOf((*StaticArrayInput)(nil)).Elem(), StaticArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StaticMapInput)(nil)).Elem(), StaticMap{})
	pulumi.RegisterOutputType(StaticOutput{})
	pulumi.RegisterOutputType(StaticArrayOutput{})
	pulumi.RegisterOutputType(StaticMapOutput{})
}
