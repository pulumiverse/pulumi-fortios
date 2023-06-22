// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure router settings.
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
//			_, err := fortios.NewRouterSetting(ctx, "trname", &fortios.RouterSettingArgs{
//				Hostname: pulumi.String("s1"),
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
// # Router Setting can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/routerSetting:RouterSetting labelname RouterSetting
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/routerSetting:RouterSetting labelname RouterSetting
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type RouterSetting struct {
	pulumi.CustomResourceState

	// bgp_debug_flags
	BgpDebugFlags pulumi.StringOutput `pulumi:"bgpDebugFlags"`
	// Hostname for this virtual domain router.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// igmp_debug_flags
	IgmpDebugFlags pulumi.StringOutput `pulumi:"igmpDebugFlags"`
	// imi_debug_flags
	ImiDebugFlags pulumi.StringOutput `pulumi:"imiDebugFlags"`
	// isis_debug_flags
	IsisDebugFlags pulumi.StringOutput `pulumi:"isisDebugFlags"`
	// ospf6_debug_events_flags
	Ospf6DebugEventsFlags pulumi.StringOutput `pulumi:"ospf6DebugEventsFlags"`
	// ospf6_debug_ifsm_flags
	Ospf6DebugIfsmFlags pulumi.StringOutput `pulumi:"ospf6DebugIfsmFlags"`
	// ospf6_debug_lsa_flags
	Ospf6DebugLsaFlags pulumi.StringOutput `pulumi:"ospf6DebugLsaFlags"`
	// ospf6_debug_nfsm_flags
	Ospf6DebugNfsmFlags pulumi.StringOutput `pulumi:"ospf6DebugNfsmFlags"`
	// ospf6_debug_nsm_flags
	Ospf6DebugNsmFlags pulumi.StringOutput `pulumi:"ospf6DebugNsmFlags"`
	// ospf6_debug_packet_flags
	Ospf6DebugPacketFlags pulumi.StringOutput `pulumi:"ospf6DebugPacketFlags"`
	// ospf6_debug_route_flags
	Ospf6DebugRouteFlags pulumi.StringOutput `pulumi:"ospf6DebugRouteFlags"`
	// ospf_debug_events_flags
	OspfDebugEventsFlags pulumi.StringOutput `pulumi:"ospfDebugEventsFlags"`
	// ospf_debug_ifsm_flags
	OspfDebugIfsmFlags pulumi.StringOutput `pulumi:"ospfDebugIfsmFlags"`
	// ospf_debug_lsa_flags
	OspfDebugLsaFlags pulumi.StringOutput `pulumi:"ospfDebugLsaFlags"`
	// ospf_debug_nfsm_flags
	OspfDebugNfsmFlags pulumi.StringOutput `pulumi:"ospfDebugNfsmFlags"`
	// ospf_debug_nsm_flags
	OspfDebugNsmFlags pulumi.StringOutput `pulumi:"ospfDebugNsmFlags"`
	// ospf_debug_packet_flags
	OspfDebugPacketFlags pulumi.StringOutput `pulumi:"ospfDebugPacketFlags"`
	// ospf_debug_route_flags
	OspfDebugRouteFlags pulumi.StringOutput `pulumi:"ospfDebugRouteFlags"`
	// pimdm_debug_flags
	PimdmDebugFlags pulumi.StringOutput `pulumi:"pimdmDebugFlags"`
	// pimsm_debug_joinprune_flags
	PimsmDebugJoinpruneFlags pulumi.StringOutput `pulumi:"pimsmDebugJoinpruneFlags"`
	// pimsm_debug_simple_flags
	PimsmDebugSimpleFlags pulumi.StringOutput `pulumi:"pimsmDebugSimpleFlags"`
	// pimsm_debug_timer_flags
	PimsmDebugTimerFlags pulumi.StringOutput `pulumi:"pimsmDebugTimerFlags"`
	// rip_debug_flags
	RipDebugFlags pulumi.StringOutput `pulumi:"ripDebugFlags"`
	// ripng_debug_flags
	RipngDebugFlags pulumi.StringOutput `pulumi:"ripngDebugFlags"`
	// Prefix-list as filter for showing routes.
	ShowFilter pulumi.StringOutput `pulumi:"showFilter"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterSetting registers a new resource with the given unique name, arguments, and options.
func NewRouterSetting(ctx *pulumi.Context,
	name string, args *RouterSettingArgs, opts ...pulumi.ResourceOption) (*RouterSetting, error) {
	if args == nil {
		args = &RouterSettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource RouterSetting
	err := ctx.RegisterResource("fortios:index/routerSetting:RouterSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterSetting gets an existing RouterSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterSettingState, opts ...pulumi.ResourceOption) (*RouterSetting, error) {
	var resource RouterSetting
	err := ctx.ReadResource("fortios:index/routerSetting:RouterSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterSetting resources.
type routerSettingState struct {
	// bgp_debug_flags
	BgpDebugFlags *string `pulumi:"bgpDebugFlags"`
	// Hostname for this virtual domain router.
	Hostname *string `pulumi:"hostname"`
	// igmp_debug_flags
	IgmpDebugFlags *string `pulumi:"igmpDebugFlags"`
	// imi_debug_flags
	ImiDebugFlags *string `pulumi:"imiDebugFlags"`
	// isis_debug_flags
	IsisDebugFlags *string `pulumi:"isisDebugFlags"`
	// ospf6_debug_events_flags
	Ospf6DebugEventsFlags *string `pulumi:"ospf6DebugEventsFlags"`
	// ospf6_debug_ifsm_flags
	Ospf6DebugIfsmFlags *string `pulumi:"ospf6DebugIfsmFlags"`
	// ospf6_debug_lsa_flags
	Ospf6DebugLsaFlags *string `pulumi:"ospf6DebugLsaFlags"`
	// ospf6_debug_nfsm_flags
	Ospf6DebugNfsmFlags *string `pulumi:"ospf6DebugNfsmFlags"`
	// ospf6_debug_nsm_flags
	Ospf6DebugNsmFlags *string `pulumi:"ospf6DebugNsmFlags"`
	// ospf6_debug_packet_flags
	Ospf6DebugPacketFlags *string `pulumi:"ospf6DebugPacketFlags"`
	// ospf6_debug_route_flags
	Ospf6DebugRouteFlags *string `pulumi:"ospf6DebugRouteFlags"`
	// ospf_debug_events_flags
	OspfDebugEventsFlags *string `pulumi:"ospfDebugEventsFlags"`
	// ospf_debug_ifsm_flags
	OspfDebugIfsmFlags *string `pulumi:"ospfDebugIfsmFlags"`
	// ospf_debug_lsa_flags
	OspfDebugLsaFlags *string `pulumi:"ospfDebugLsaFlags"`
	// ospf_debug_nfsm_flags
	OspfDebugNfsmFlags *string `pulumi:"ospfDebugNfsmFlags"`
	// ospf_debug_nsm_flags
	OspfDebugNsmFlags *string `pulumi:"ospfDebugNsmFlags"`
	// ospf_debug_packet_flags
	OspfDebugPacketFlags *string `pulumi:"ospfDebugPacketFlags"`
	// ospf_debug_route_flags
	OspfDebugRouteFlags *string `pulumi:"ospfDebugRouteFlags"`
	// pimdm_debug_flags
	PimdmDebugFlags *string `pulumi:"pimdmDebugFlags"`
	// pimsm_debug_joinprune_flags
	PimsmDebugJoinpruneFlags *string `pulumi:"pimsmDebugJoinpruneFlags"`
	// pimsm_debug_simple_flags
	PimsmDebugSimpleFlags *string `pulumi:"pimsmDebugSimpleFlags"`
	// pimsm_debug_timer_flags
	PimsmDebugTimerFlags *string `pulumi:"pimsmDebugTimerFlags"`
	// rip_debug_flags
	RipDebugFlags *string `pulumi:"ripDebugFlags"`
	// ripng_debug_flags
	RipngDebugFlags *string `pulumi:"ripngDebugFlags"`
	// Prefix-list as filter for showing routes.
	ShowFilter *string `pulumi:"showFilter"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterSettingState struct {
	// bgp_debug_flags
	BgpDebugFlags pulumi.StringPtrInput
	// Hostname for this virtual domain router.
	Hostname pulumi.StringPtrInput
	// igmp_debug_flags
	IgmpDebugFlags pulumi.StringPtrInput
	// imi_debug_flags
	ImiDebugFlags pulumi.StringPtrInput
	// isis_debug_flags
	IsisDebugFlags pulumi.StringPtrInput
	// ospf6_debug_events_flags
	Ospf6DebugEventsFlags pulumi.StringPtrInput
	// ospf6_debug_ifsm_flags
	Ospf6DebugIfsmFlags pulumi.StringPtrInput
	// ospf6_debug_lsa_flags
	Ospf6DebugLsaFlags pulumi.StringPtrInput
	// ospf6_debug_nfsm_flags
	Ospf6DebugNfsmFlags pulumi.StringPtrInput
	// ospf6_debug_nsm_flags
	Ospf6DebugNsmFlags pulumi.StringPtrInput
	// ospf6_debug_packet_flags
	Ospf6DebugPacketFlags pulumi.StringPtrInput
	// ospf6_debug_route_flags
	Ospf6DebugRouteFlags pulumi.StringPtrInput
	// ospf_debug_events_flags
	OspfDebugEventsFlags pulumi.StringPtrInput
	// ospf_debug_ifsm_flags
	OspfDebugIfsmFlags pulumi.StringPtrInput
	// ospf_debug_lsa_flags
	OspfDebugLsaFlags pulumi.StringPtrInput
	// ospf_debug_nfsm_flags
	OspfDebugNfsmFlags pulumi.StringPtrInput
	// ospf_debug_nsm_flags
	OspfDebugNsmFlags pulumi.StringPtrInput
	// ospf_debug_packet_flags
	OspfDebugPacketFlags pulumi.StringPtrInput
	// ospf_debug_route_flags
	OspfDebugRouteFlags pulumi.StringPtrInput
	// pimdm_debug_flags
	PimdmDebugFlags pulumi.StringPtrInput
	// pimsm_debug_joinprune_flags
	PimsmDebugJoinpruneFlags pulumi.StringPtrInput
	// pimsm_debug_simple_flags
	PimsmDebugSimpleFlags pulumi.StringPtrInput
	// pimsm_debug_timer_flags
	PimsmDebugTimerFlags pulumi.StringPtrInput
	// rip_debug_flags
	RipDebugFlags pulumi.StringPtrInput
	// ripng_debug_flags
	RipngDebugFlags pulumi.StringPtrInput
	// Prefix-list as filter for showing routes.
	ShowFilter pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerSettingState)(nil)).Elem()
}

type routerSettingArgs struct {
	// bgp_debug_flags
	BgpDebugFlags *string `pulumi:"bgpDebugFlags"`
	// Hostname for this virtual domain router.
	Hostname *string `pulumi:"hostname"`
	// igmp_debug_flags
	IgmpDebugFlags *string `pulumi:"igmpDebugFlags"`
	// imi_debug_flags
	ImiDebugFlags *string `pulumi:"imiDebugFlags"`
	// isis_debug_flags
	IsisDebugFlags *string `pulumi:"isisDebugFlags"`
	// ospf6_debug_events_flags
	Ospf6DebugEventsFlags *string `pulumi:"ospf6DebugEventsFlags"`
	// ospf6_debug_ifsm_flags
	Ospf6DebugIfsmFlags *string `pulumi:"ospf6DebugIfsmFlags"`
	// ospf6_debug_lsa_flags
	Ospf6DebugLsaFlags *string `pulumi:"ospf6DebugLsaFlags"`
	// ospf6_debug_nfsm_flags
	Ospf6DebugNfsmFlags *string `pulumi:"ospf6DebugNfsmFlags"`
	// ospf6_debug_nsm_flags
	Ospf6DebugNsmFlags *string `pulumi:"ospf6DebugNsmFlags"`
	// ospf6_debug_packet_flags
	Ospf6DebugPacketFlags *string `pulumi:"ospf6DebugPacketFlags"`
	// ospf6_debug_route_flags
	Ospf6DebugRouteFlags *string `pulumi:"ospf6DebugRouteFlags"`
	// ospf_debug_events_flags
	OspfDebugEventsFlags *string `pulumi:"ospfDebugEventsFlags"`
	// ospf_debug_ifsm_flags
	OspfDebugIfsmFlags *string `pulumi:"ospfDebugIfsmFlags"`
	// ospf_debug_lsa_flags
	OspfDebugLsaFlags *string `pulumi:"ospfDebugLsaFlags"`
	// ospf_debug_nfsm_flags
	OspfDebugNfsmFlags *string `pulumi:"ospfDebugNfsmFlags"`
	// ospf_debug_nsm_flags
	OspfDebugNsmFlags *string `pulumi:"ospfDebugNsmFlags"`
	// ospf_debug_packet_flags
	OspfDebugPacketFlags *string `pulumi:"ospfDebugPacketFlags"`
	// ospf_debug_route_flags
	OspfDebugRouteFlags *string `pulumi:"ospfDebugRouteFlags"`
	// pimdm_debug_flags
	PimdmDebugFlags *string `pulumi:"pimdmDebugFlags"`
	// pimsm_debug_joinprune_flags
	PimsmDebugJoinpruneFlags *string `pulumi:"pimsmDebugJoinpruneFlags"`
	// pimsm_debug_simple_flags
	PimsmDebugSimpleFlags *string `pulumi:"pimsmDebugSimpleFlags"`
	// pimsm_debug_timer_flags
	PimsmDebugTimerFlags *string `pulumi:"pimsmDebugTimerFlags"`
	// rip_debug_flags
	RipDebugFlags *string `pulumi:"ripDebugFlags"`
	// ripng_debug_flags
	RipngDebugFlags *string `pulumi:"ripngDebugFlags"`
	// Prefix-list as filter for showing routes.
	ShowFilter *string `pulumi:"showFilter"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterSetting resource.
type RouterSettingArgs struct {
	// bgp_debug_flags
	BgpDebugFlags pulumi.StringPtrInput
	// Hostname for this virtual domain router.
	Hostname pulumi.StringPtrInput
	// igmp_debug_flags
	IgmpDebugFlags pulumi.StringPtrInput
	// imi_debug_flags
	ImiDebugFlags pulumi.StringPtrInput
	// isis_debug_flags
	IsisDebugFlags pulumi.StringPtrInput
	// ospf6_debug_events_flags
	Ospf6DebugEventsFlags pulumi.StringPtrInput
	// ospf6_debug_ifsm_flags
	Ospf6DebugIfsmFlags pulumi.StringPtrInput
	// ospf6_debug_lsa_flags
	Ospf6DebugLsaFlags pulumi.StringPtrInput
	// ospf6_debug_nfsm_flags
	Ospf6DebugNfsmFlags pulumi.StringPtrInput
	// ospf6_debug_nsm_flags
	Ospf6DebugNsmFlags pulumi.StringPtrInput
	// ospf6_debug_packet_flags
	Ospf6DebugPacketFlags pulumi.StringPtrInput
	// ospf6_debug_route_flags
	Ospf6DebugRouteFlags pulumi.StringPtrInput
	// ospf_debug_events_flags
	OspfDebugEventsFlags pulumi.StringPtrInput
	// ospf_debug_ifsm_flags
	OspfDebugIfsmFlags pulumi.StringPtrInput
	// ospf_debug_lsa_flags
	OspfDebugLsaFlags pulumi.StringPtrInput
	// ospf_debug_nfsm_flags
	OspfDebugNfsmFlags pulumi.StringPtrInput
	// ospf_debug_nsm_flags
	OspfDebugNsmFlags pulumi.StringPtrInput
	// ospf_debug_packet_flags
	OspfDebugPacketFlags pulumi.StringPtrInput
	// ospf_debug_route_flags
	OspfDebugRouteFlags pulumi.StringPtrInput
	// pimdm_debug_flags
	PimdmDebugFlags pulumi.StringPtrInput
	// pimsm_debug_joinprune_flags
	PimsmDebugJoinpruneFlags pulumi.StringPtrInput
	// pimsm_debug_simple_flags
	PimsmDebugSimpleFlags pulumi.StringPtrInput
	// pimsm_debug_timer_flags
	PimsmDebugTimerFlags pulumi.StringPtrInput
	// rip_debug_flags
	RipDebugFlags pulumi.StringPtrInput
	// ripng_debug_flags
	RipngDebugFlags pulumi.StringPtrInput
	// Prefix-list as filter for showing routes.
	ShowFilter pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerSettingArgs)(nil)).Elem()
}

type RouterSettingInput interface {
	pulumi.Input

	ToRouterSettingOutput() RouterSettingOutput
	ToRouterSettingOutputWithContext(ctx context.Context) RouterSettingOutput
}

func (*RouterSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterSetting)(nil)).Elem()
}

func (i *RouterSetting) ToRouterSettingOutput() RouterSettingOutput {
	return i.ToRouterSettingOutputWithContext(context.Background())
}

func (i *RouterSetting) ToRouterSettingOutputWithContext(ctx context.Context) RouterSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterSettingOutput)
}

// RouterSettingArrayInput is an input type that accepts RouterSettingArray and RouterSettingArrayOutput values.
// You can construct a concrete instance of `RouterSettingArrayInput` via:
//
//	RouterSettingArray{ RouterSettingArgs{...} }
type RouterSettingArrayInput interface {
	pulumi.Input

	ToRouterSettingArrayOutput() RouterSettingArrayOutput
	ToRouterSettingArrayOutputWithContext(context.Context) RouterSettingArrayOutput
}

type RouterSettingArray []RouterSettingInput

func (RouterSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterSetting)(nil)).Elem()
}

func (i RouterSettingArray) ToRouterSettingArrayOutput() RouterSettingArrayOutput {
	return i.ToRouterSettingArrayOutputWithContext(context.Background())
}

func (i RouterSettingArray) ToRouterSettingArrayOutputWithContext(ctx context.Context) RouterSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterSettingArrayOutput)
}

// RouterSettingMapInput is an input type that accepts RouterSettingMap and RouterSettingMapOutput values.
// You can construct a concrete instance of `RouterSettingMapInput` via:
//
//	RouterSettingMap{ "key": RouterSettingArgs{...} }
type RouterSettingMapInput interface {
	pulumi.Input

	ToRouterSettingMapOutput() RouterSettingMapOutput
	ToRouterSettingMapOutputWithContext(context.Context) RouterSettingMapOutput
}

type RouterSettingMap map[string]RouterSettingInput

func (RouterSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterSetting)(nil)).Elem()
}

func (i RouterSettingMap) ToRouterSettingMapOutput() RouterSettingMapOutput {
	return i.ToRouterSettingMapOutputWithContext(context.Background())
}

func (i RouterSettingMap) ToRouterSettingMapOutputWithContext(ctx context.Context) RouterSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterSettingMapOutput)
}

type RouterSettingOutput struct{ *pulumi.OutputState }

func (RouterSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterSetting)(nil)).Elem()
}

func (o RouterSettingOutput) ToRouterSettingOutput() RouterSettingOutput {
	return o
}

func (o RouterSettingOutput) ToRouterSettingOutputWithContext(ctx context.Context) RouterSettingOutput {
	return o
}

// bgp_debug_flags
func (o RouterSettingOutput) BgpDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.BgpDebugFlags }).(pulumi.StringOutput)
}

// Hostname for this virtual domain router.
func (o RouterSettingOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// igmp_debug_flags
func (o RouterSettingOutput) IgmpDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.IgmpDebugFlags }).(pulumi.StringOutput)
}

// imi_debug_flags
func (o RouterSettingOutput) ImiDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.ImiDebugFlags }).(pulumi.StringOutput)
}

// isis_debug_flags
func (o RouterSettingOutput) IsisDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.IsisDebugFlags }).(pulumi.StringOutput)
}

// ospf6_debug_events_flags
func (o RouterSettingOutput) Ospf6DebugEventsFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugEventsFlags }).(pulumi.StringOutput)
}

// ospf6_debug_ifsm_flags
func (o RouterSettingOutput) Ospf6DebugIfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugIfsmFlags }).(pulumi.StringOutput)
}

// ospf6_debug_lsa_flags
func (o RouterSettingOutput) Ospf6DebugLsaFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugLsaFlags }).(pulumi.StringOutput)
}

// ospf6_debug_nfsm_flags
func (o RouterSettingOutput) Ospf6DebugNfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugNfsmFlags }).(pulumi.StringOutput)
}

// ospf6_debug_nsm_flags
func (o RouterSettingOutput) Ospf6DebugNsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugNsmFlags }).(pulumi.StringOutput)
}

// ospf6_debug_packet_flags
func (o RouterSettingOutput) Ospf6DebugPacketFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugPacketFlags }).(pulumi.StringOutput)
}

// ospf6_debug_route_flags
func (o RouterSettingOutput) Ospf6DebugRouteFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugRouteFlags }).(pulumi.StringOutput)
}

// ospf_debug_events_flags
func (o RouterSettingOutput) OspfDebugEventsFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugEventsFlags }).(pulumi.StringOutput)
}

// ospf_debug_ifsm_flags
func (o RouterSettingOutput) OspfDebugIfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugIfsmFlags }).(pulumi.StringOutput)
}

// ospf_debug_lsa_flags
func (o RouterSettingOutput) OspfDebugLsaFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugLsaFlags }).(pulumi.StringOutput)
}

// ospf_debug_nfsm_flags
func (o RouterSettingOutput) OspfDebugNfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugNfsmFlags }).(pulumi.StringOutput)
}

// ospf_debug_nsm_flags
func (o RouterSettingOutput) OspfDebugNsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugNsmFlags }).(pulumi.StringOutput)
}

// ospf_debug_packet_flags
func (o RouterSettingOutput) OspfDebugPacketFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugPacketFlags }).(pulumi.StringOutput)
}

// ospf_debug_route_flags
func (o RouterSettingOutput) OspfDebugRouteFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugRouteFlags }).(pulumi.StringOutput)
}

// pimdm_debug_flags
func (o RouterSettingOutput) PimdmDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.PimdmDebugFlags }).(pulumi.StringOutput)
}

// pimsm_debug_joinprune_flags
func (o RouterSettingOutput) PimsmDebugJoinpruneFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.PimsmDebugJoinpruneFlags }).(pulumi.StringOutput)
}

// pimsm_debug_simple_flags
func (o RouterSettingOutput) PimsmDebugSimpleFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.PimsmDebugSimpleFlags }).(pulumi.StringOutput)
}

// pimsm_debug_timer_flags
func (o RouterSettingOutput) PimsmDebugTimerFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.PimsmDebugTimerFlags }).(pulumi.StringOutput)
}

// rip_debug_flags
func (o RouterSettingOutput) RipDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.RipDebugFlags }).(pulumi.StringOutput)
}

// ripng_debug_flags
func (o RouterSettingOutput) RipngDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.RipngDebugFlags }).(pulumi.StringOutput)
}

// Prefix-list as filter for showing routes.
func (o RouterSettingOutput) ShowFilter() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.ShowFilter }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RouterSettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterSettingArrayOutput struct{ *pulumi.OutputState }

func (RouterSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterSetting)(nil)).Elem()
}

func (o RouterSettingArrayOutput) ToRouterSettingArrayOutput() RouterSettingArrayOutput {
	return o
}

func (o RouterSettingArrayOutput) ToRouterSettingArrayOutputWithContext(ctx context.Context) RouterSettingArrayOutput {
	return o
}

func (o RouterSettingArrayOutput) Index(i pulumi.IntInput) RouterSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterSetting {
		return vs[0].([]*RouterSetting)[vs[1].(int)]
	}).(RouterSettingOutput)
}

type RouterSettingMapOutput struct{ *pulumi.OutputState }

func (RouterSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterSetting)(nil)).Elem()
}

func (o RouterSettingMapOutput) ToRouterSettingMapOutput() RouterSettingMapOutput {
	return o
}

func (o RouterSettingMapOutput) ToRouterSettingMapOutputWithContext(ctx context.Context) RouterSettingMapOutput {
	return o
}

func (o RouterSettingMapOutput) MapIndex(k pulumi.StringInput) RouterSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterSetting {
		return vs[0].(map[string]*RouterSetting)[vs[1].(string)]
	}).(RouterSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterSettingInput)(nil)).Elem(), &RouterSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterSettingArrayInput)(nil)).Elem(), RouterSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterSettingMapInput)(nil)).Elem(), RouterSettingMap{})
	pulumi.RegisterOutputType(RouterSettingOutput{})
	pulumi.RegisterOutputType(RouterSettingArrayOutput{})
	pulumi.RegisterOutputType(RouterSettingMapOutput{})
}
