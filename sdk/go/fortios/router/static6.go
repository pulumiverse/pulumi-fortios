// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPv6 static routing tables.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/router"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := router.NewStatic6(ctx, "trname", &router.Static6Args{
//				Bfd:            pulumi.String("disable"),
//				Blackhole:      pulumi.String("disable"),
//				Device:         pulumi.String("port3"),
//				Devindex:       pulumi.Int(5),
//				Distance:       pulumi.Int(10),
//				Dst:            pulumi.String("2001:db8::/32"),
//				Gateway:        pulumi.String("::"),
//				Priority:       pulumi.Int(32),
//				SeqNum:         pulumi.Int(1),
//				Status:         pulumi.String("enable"),
//				VirtualWanLink: pulumi.String("disable"),
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
// Router Static6 can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:router/static6:Static6 labelname {{seq_num}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:router/static6:Static6 labelname {{seq_num}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Static6 struct {
	pulumi.CustomResourceState

	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
	Bfd pulumi.StringOutput `pulumi:"bfd"`
	// Enable/disable black hole. Valid values: `enable`, `disable`.
	Blackhole pulumi.StringOutput `pulumi:"blackhole"`
	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Gateway out interface or tunnel.
	Device pulumi.StringOutput `pulumi:"device"`
	// Device index (0 - 4294967295).
	Devindex pulumi.IntOutput `pulumi:"devindex"`
	// Administrative distance (1 - 255).
	Distance pulumi.IntOutput `pulumi:"distance"`
	// Destination IPv6 prefix.
	Dst pulumi.StringOutput `pulumi:"dst"`
	// Name of firewall address or address group.
	Dstaddr pulumi.StringOutput `pulumi:"dstaddr"`
	// Enable use of dynamic gateway retrieved from Router Advertisement (RA). Valid values: `enable`, `disable`.
	DynamicGateway pulumi.StringOutput `pulumi:"dynamicGateway"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// IPv6 address of the gateway.
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Enable/disable withdrawal of this static route when link monitor or health check is down. Valid values: `enable`, `disable`.
	LinkMonitorExempt pulumi.StringOutput `pulumi:"linkMonitorExempt"`
	// Administrative priority (0 - 4294967295).
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Enable/disable egress through the SD-WAN. Valid values: `enable`, `disable`.
	Sdwan pulumi.StringOutput `pulumi:"sdwan"`
	// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
	SdwanZones Static6SdwanZoneArrayOutput `pulumi:"sdwanZones"`
	// Sequence number.
	SeqNum pulumi.IntOutput `pulumi:"seqNum"`
	// Enable/disable this static route. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
	VirtualWanLink pulumi.StringOutput `pulumi:"virtualWanLink"`
	// Virtual Routing Forwarding ID.
	Vrf pulumi.IntOutput `pulumi:"vrf"`
	// Administrative weight (0 - 255).
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewStatic6 registers a new resource with the given unique name, arguments, and options.
func NewStatic6(ctx *pulumi.Context,
	name string, args *Static6Args, opts ...pulumi.ResourceOption) (*Static6, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Device == nil {
		return nil, errors.New("invalid value for required argument 'Device'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Static6
	err := ctx.RegisterResource("fortios:router/static6:Static6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStatic6 gets an existing Static6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStatic6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Static6State, opts ...pulumi.ResourceOption) (*Static6, error) {
	var resource Static6
	err := ctx.ReadResource("fortios:router/static6:Static6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Static6 resources.
type static6State struct {
	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
	Bfd *string `pulumi:"bfd"`
	// Enable/disable black hole. Valid values: `enable`, `disable`.
	Blackhole *string `pulumi:"blackhole"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Gateway out interface or tunnel.
	Device *string `pulumi:"device"`
	// Device index (0 - 4294967295).
	Devindex *int `pulumi:"devindex"`
	// Administrative distance (1 - 255).
	Distance *int `pulumi:"distance"`
	// Destination IPv6 prefix.
	Dst *string `pulumi:"dst"`
	// Name of firewall address or address group.
	Dstaddr *string `pulumi:"dstaddr"`
	// Enable use of dynamic gateway retrieved from Router Advertisement (RA). Valid values: `enable`, `disable`.
	DynamicGateway *string `pulumi:"dynamicGateway"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// IPv6 address of the gateway.
	Gateway *string `pulumi:"gateway"`
	// Enable/disable withdrawal of this static route when link monitor or health check is down. Valid values: `enable`, `disable`.
	LinkMonitorExempt *string `pulumi:"linkMonitorExempt"`
	// Administrative priority (0 - 4294967295).
	Priority *int `pulumi:"priority"`
	// Enable/disable egress through the SD-WAN. Valid values: `enable`, `disable`.
	Sdwan *string `pulumi:"sdwan"`
	// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
	SdwanZones []Static6SdwanZone `pulumi:"sdwanZones"`
	// Sequence number.
	SeqNum *int `pulumi:"seqNum"`
	// Enable/disable this static route. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
	VirtualWanLink *string `pulumi:"virtualWanLink"`
	// Virtual Routing Forwarding ID.
	Vrf *int `pulumi:"vrf"`
	// Administrative weight (0 - 255).
	Weight *int `pulumi:"weight"`
}

type Static6State struct {
	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
	Bfd pulumi.StringPtrInput
	// Enable/disable black hole. Valid values: `enable`, `disable`.
	Blackhole pulumi.StringPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Gateway out interface or tunnel.
	Device pulumi.StringPtrInput
	// Device index (0 - 4294967295).
	Devindex pulumi.IntPtrInput
	// Administrative distance (1 - 255).
	Distance pulumi.IntPtrInput
	// Destination IPv6 prefix.
	Dst pulumi.StringPtrInput
	// Name of firewall address or address group.
	Dstaddr pulumi.StringPtrInput
	// Enable use of dynamic gateway retrieved from Router Advertisement (RA). Valid values: `enable`, `disable`.
	DynamicGateway pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// IPv6 address of the gateway.
	Gateway pulumi.StringPtrInput
	// Enable/disable withdrawal of this static route when link monitor or health check is down. Valid values: `enable`, `disable`.
	LinkMonitorExempt pulumi.StringPtrInput
	// Administrative priority (0 - 4294967295).
	Priority pulumi.IntPtrInput
	// Enable/disable egress through the SD-WAN. Valid values: `enable`, `disable`.
	Sdwan pulumi.StringPtrInput
	// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
	SdwanZones Static6SdwanZoneArrayInput
	// Sequence number.
	SeqNum pulumi.IntPtrInput
	// Enable/disable this static route. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
	VirtualWanLink pulumi.StringPtrInput
	// Virtual Routing Forwarding ID.
	Vrf pulumi.IntPtrInput
	// Administrative weight (0 - 255).
	Weight pulumi.IntPtrInput
}

func (Static6State) ElementType() reflect.Type {
	return reflect.TypeOf((*static6State)(nil)).Elem()
}

type static6Args struct {
	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
	Bfd *string `pulumi:"bfd"`
	// Enable/disable black hole. Valid values: `enable`, `disable`.
	Blackhole *string `pulumi:"blackhole"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Gateway out interface or tunnel.
	Device string `pulumi:"device"`
	// Device index (0 - 4294967295).
	Devindex *int `pulumi:"devindex"`
	// Administrative distance (1 - 255).
	Distance *int `pulumi:"distance"`
	// Destination IPv6 prefix.
	Dst *string `pulumi:"dst"`
	// Name of firewall address or address group.
	Dstaddr *string `pulumi:"dstaddr"`
	// Enable use of dynamic gateway retrieved from Router Advertisement (RA). Valid values: `enable`, `disable`.
	DynamicGateway *string `pulumi:"dynamicGateway"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// IPv6 address of the gateway.
	Gateway *string `pulumi:"gateway"`
	// Enable/disable withdrawal of this static route when link monitor or health check is down. Valid values: `enable`, `disable`.
	LinkMonitorExempt *string `pulumi:"linkMonitorExempt"`
	// Administrative priority (0 - 4294967295).
	Priority *int `pulumi:"priority"`
	// Enable/disable egress through the SD-WAN. Valid values: `enable`, `disable`.
	Sdwan *string `pulumi:"sdwan"`
	// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
	SdwanZones []Static6SdwanZone `pulumi:"sdwanZones"`
	// Sequence number.
	SeqNum *int `pulumi:"seqNum"`
	// Enable/disable this static route. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
	VirtualWanLink *string `pulumi:"virtualWanLink"`
	// Virtual Routing Forwarding ID.
	Vrf *int `pulumi:"vrf"`
	// Administrative weight (0 - 255).
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a Static6 resource.
type Static6Args struct {
	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
	Bfd pulumi.StringPtrInput
	// Enable/disable black hole. Valid values: `enable`, `disable`.
	Blackhole pulumi.StringPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Gateway out interface or tunnel.
	Device pulumi.StringInput
	// Device index (0 - 4294967295).
	Devindex pulumi.IntPtrInput
	// Administrative distance (1 - 255).
	Distance pulumi.IntPtrInput
	// Destination IPv6 prefix.
	Dst pulumi.StringPtrInput
	// Name of firewall address or address group.
	Dstaddr pulumi.StringPtrInput
	// Enable use of dynamic gateway retrieved from Router Advertisement (RA). Valid values: `enable`, `disable`.
	DynamicGateway pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// IPv6 address of the gateway.
	Gateway pulumi.StringPtrInput
	// Enable/disable withdrawal of this static route when link monitor or health check is down. Valid values: `enable`, `disable`.
	LinkMonitorExempt pulumi.StringPtrInput
	// Administrative priority (0 - 4294967295).
	Priority pulumi.IntPtrInput
	// Enable/disable egress through the SD-WAN. Valid values: `enable`, `disable`.
	Sdwan pulumi.StringPtrInput
	// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
	SdwanZones Static6SdwanZoneArrayInput
	// Sequence number.
	SeqNum pulumi.IntPtrInput
	// Enable/disable this static route. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
	VirtualWanLink pulumi.StringPtrInput
	// Virtual Routing Forwarding ID.
	Vrf pulumi.IntPtrInput
	// Administrative weight (0 - 255).
	Weight pulumi.IntPtrInput
}

func (Static6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*static6Args)(nil)).Elem()
}

type Static6Input interface {
	pulumi.Input

	ToStatic6Output() Static6Output
	ToStatic6OutputWithContext(ctx context.Context) Static6Output
}

func (*Static6) ElementType() reflect.Type {
	return reflect.TypeOf((**Static6)(nil)).Elem()
}

func (i *Static6) ToStatic6Output() Static6Output {
	return i.ToStatic6OutputWithContext(context.Background())
}

func (i *Static6) ToStatic6OutputWithContext(ctx context.Context) Static6Output {
	return pulumi.ToOutputWithContext(ctx, i).(Static6Output)
}

// Static6ArrayInput is an input type that accepts Static6Array and Static6ArrayOutput values.
// You can construct a concrete instance of `Static6ArrayInput` via:
//
//	Static6Array{ Static6Args{...} }
type Static6ArrayInput interface {
	pulumi.Input

	ToStatic6ArrayOutput() Static6ArrayOutput
	ToStatic6ArrayOutputWithContext(context.Context) Static6ArrayOutput
}

type Static6Array []Static6Input

func (Static6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Static6)(nil)).Elem()
}

func (i Static6Array) ToStatic6ArrayOutput() Static6ArrayOutput {
	return i.ToStatic6ArrayOutputWithContext(context.Background())
}

func (i Static6Array) ToStatic6ArrayOutputWithContext(ctx context.Context) Static6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Static6ArrayOutput)
}

// Static6MapInput is an input type that accepts Static6Map and Static6MapOutput values.
// You can construct a concrete instance of `Static6MapInput` via:
//
//	Static6Map{ "key": Static6Args{...} }
type Static6MapInput interface {
	pulumi.Input

	ToStatic6MapOutput() Static6MapOutput
	ToStatic6MapOutputWithContext(context.Context) Static6MapOutput
}

type Static6Map map[string]Static6Input

func (Static6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Static6)(nil)).Elem()
}

func (i Static6Map) ToStatic6MapOutput() Static6MapOutput {
	return i.ToStatic6MapOutputWithContext(context.Background())
}

func (i Static6Map) ToStatic6MapOutputWithContext(ctx context.Context) Static6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Static6MapOutput)
}

type Static6Output struct{ *pulumi.OutputState }

func (Static6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Static6)(nil)).Elem()
}

func (o Static6Output) ToStatic6Output() Static6Output {
	return o
}

func (o Static6Output) ToStatic6OutputWithContext(ctx context.Context) Static6Output {
	return o
}

// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
func (o Static6Output) Bfd() pulumi.StringOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringOutput { return v.Bfd }).(pulumi.StringOutput)
}

// Enable/disable black hole. Valid values: `enable`, `disable`.
func (o Static6Output) Blackhole() pulumi.StringOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringOutput { return v.Blackhole }).(pulumi.StringOutput)
}

// Optional comments.
func (o Static6Output) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Gateway out interface or tunnel.
func (o Static6Output) Device() pulumi.StringOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringOutput { return v.Device }).(pulumi.StringOutput)
}

// Device index (0 - 4294967295).
func (o Static6Output) Devindex() pulumi.IntOutput {
	return o.ApplyT(func(v *Static6) pulumi.IntOutput { return v.Devindex }).(pulumi.IntOutput)
}

// Administrative distance (1 - 255).
func (o Static6Output) Distance() pulumi.IntOutput {
	return o.ApplyT(func(v *Static6) pulumi.IntOutput { return v.Distance }).(pulumi.IntOutput)
}

// Destination IPv6 prefix.
func (o Static6Output) Dst() pulumi.StringOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringOutput { return v.Dst }).(pulumi.StringOutput)
}

// Name of firewall address or address group.
func (o Static6Output) Dstaddr() pulumi.StringOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringOutput { return v.Dstaddr }).(pulumi.StringOutput)
}

// Enable use of dynamic gateway retrieved from Router Advertisement (RA). Valid values: `enable`, `disable`.
func (o Static6Output) DynamicGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringOutput { return v.DynamicGateway }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Static6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// IPv6 address of the gateway.
func (o Static6Output) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// Enable/disable withdrawal of this static route when link monitor or health check is down. Valid values: `enable`, `disable`.
func (o Static6Output) LinkMonitorExempt() pulumi.StringOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringOutput { return v.LinkMonitorExempt }).(pulumi.StringOutput)
}

// Administrative priority (0 - 4294967295).
func (o Static6Output) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *Static6) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// Enable/disable egress through the SD-WAN. Valid values: `enable`, `disable`.
func (o Static6Output) Sdwan() pulumi.StringOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringOutput { return v.Sdwan }).(pulumi.StringOutput)
}

// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
func (o Static6Output) SdwanZones() Static6SdwanZoneArrayOutput {
	return o.ApplyT(func(v *Static6) Static6SdwanZoneArrayOutput { return v.SdwanZones }).(Static6SdwanZoneArrayOutput)
}

// Sequence number.
func (o Static6Output) SeqNum() pulumi.IntOutput {
	return o.ApplyT(func(v *Static6) pulumi.IntOutput { return v.SeqNum }).(pulumi.IntOutput)
}

// Enable/disable this static route. Valid values: `enable`, `disable`.
func (o Static6Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Static6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
func (o Static6Output) VirtualWanLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Static6) pulumi.StringOutput { return v.VirtualWanLink }).(pulumi.StringOutput)
}

// Virtual Routing Forwarding ID.
func (o Static6Output) Vrf() pulumi.IntOutput {
	return o.ApplyT(func(v *Static6) pulumi.IntOutput { return v.Vrf }).(pulumi.IntOutput)
}

// Administrative weight (0 - 255).
func (o Static6Output) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v *Static6) pulumi.IntOutput { return v.Weight }).(pulumi.IntOutput)
}

type Static6ArrayOutput struct{ *pulumi.OutputState }

func (Static6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Static6)(nil)).Elem()
}

func (o Static6ArrayOutput) ToStatic6ArrayOutput() Static6ArrayOutput {
	return o
}

func (o Static6ArrayOutput) ToStatic6ArrayOutputWithContext(ctx context.Context) Static6ArrayOutput {
	return o
}

func (o Static6ArrayOutput) Index(i pulumi.IntInput) Static6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Static6 {
		return vs[0].([]*Static6)[vs[1].(int)]
	}).(Static6Output)
}

type Static6MapOutput struct{ *pulumi.OutputState }

func (Static6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Static6)(nil)).Elem()
}

func (o Static6MapOutput) ToStatic6MapOutput() Static6MapOutput {
	return o
}

func (o Static6MapOutput) ToStatic6MapOutputWithContext(ctx context.Context) Static6MapOutput {
	return o
}

func (o Static6MapOutput) MapIndex(k pulumi.StringInput) Static6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Static6 {
		return vs[0].(map[string]*Static6)[vs[1].(string)]
	}).(Static6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Static6Input)(nil)).Elem(), &Static6{})
	pulumi.RegisterInputType(reflect.TypeOf((*Static6ArrayInput)(nil)).Elem(), Static6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Static6MapInput)(nil)).Elem(), Static6Map{})
	pulumi.RegisterOutputType(Static6Output{})
	pulumi.RegisterOutputType(Static6ArrayOutput{})
	pulumi.RegisterOutputType(Static6MapOutput{})
}
