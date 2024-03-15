// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure router multicast.
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
//			_, err := router.NewMulticast(ctx, "trname", &router.MulticastArgs{
//				MulticastRouting: pulumi.String("disable"),
//				PimSmGlobal: &router.MulticastPimSmGlobalArgs{
//					BsrAllowQuickRefresh:     pulumi.String("disable"),
//					BsrCandidate:             pulumi.String("disable"),
//					BsrHash:                  pulumi.Int(10),
//					BsrPriority:              pulumi.Int(0),
//					CiscoCrpPrefix:           pulumi.String("disable"),
//					CiscoIgnoreRpSetPriority: pulumi.String("disable"),
//					CiscoRegisterChecksum:    pulumi.String("disable"),
//					JoinPruneHoldtime:        pulumi.Int(210),
//					MessageInterval:          pulumi.Int(60),
//					NullRegisterRetries:      pulumi.Int(1),
//					RegisterRateLimit:        pulumi.Int(0),
//					RegisterRpReachability:   pulumi.String("enable"),
//					RegisterSource:           pulumi.String("disable"),
//					RegisterSourceIp:         pulumi.String("0.0.0.0"),
//					RegisterSupression:       pulumi.Int(60),
//					RpRegisterKeepalive:      pulumi.Int(185),
//					SptThreshold:             pulumi.String("enable"),
//					Ssm:                      pulumi.String("disable"),
//				},
//				RouteLimit:     pulumi.Int(2147483647),
//				RouteThreshold: pulumi.Int(2147483647),
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
// Router Multicast can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:router/multicast:Multicast labelname RouterMulticast
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:router/multicast:Multicast labelname RouterMulticast
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Multicast struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// PIM interfaces. The structure of `interface` block is documented below.
	Interfaces MulticastInterfaceArrayOutput `pulumi:"interfaces"`
	// Enable/disable IP multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting pulumi.StringOutput `pulumi:"multicastRouting"`
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal MulticastPimSmGlobalOutput `pulumi:"pimSmGlobal"`
	// Maximum number of multicast routes.
	RouteLimit pulumi.IntOutput `pulumi:"routeLimit"`
	// Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
	RouteThreshold pulumi.IntOutput `pulumi:"routeThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewMulticast registers a new resource with the given unique name, arguments, and options.
func NewMulticast(ctx *pulumi.Context,
	name string, args *MulticastArgs, opts ...pulumi.ResourceOption) (*Multicast, error) {
	if args == nil {
		args = &MulticastArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Multicast
	err := ctx.RegisterResource("fortios:router/multicast:Multicast", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMulticast gets an existing Multicast resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMulticast(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MulticastState, opts ...pulumi.ResourceOption) (*Multicast, error) {
	var resource Multicast
	err := ctx.ReadResource("fortios:router/multicast:Multicast", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Multicast resources.
type multicastState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// PIM interfaces. The structure of `interface` block is documented below.
	Interfaces []MulticastInterface `pulumi:"interfaces"`
	// Enable/disable IP multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting *string `pulumi:"multicastRouting"`
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal *MulticastPimSmGlobal `pulumi:"pimSmGlobal"`
	// Maximum number of multicast routes.
	RouteLimit *int `pulumi:"routeLimit"`
	// Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
	RouteThreshold *int `pulumi:"routeThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type MulticastState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// PIM interfaces. The structure of `interface` block is documented below.
	Interfaces MulticastInterfaceArrayInput
	// Enable/disable IP multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting pulumi.StringPtrInput
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal MulticastPimSmGlobalPtrInput
	// Maximum number of multicast routes.
	RouteLimit pulumi.IntPtrInput
	// Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
	RouteThreshold pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (MulticastState) ElementType() reflect.Type {
	return reflect.TypeOf((*multicastState)(nil)).Elem()
}

type multicastArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// PIM interfaces. The structure of `interface` block is documented below.
	Interfaces []MulticastInterface `pulumi:"interfaces"`
	// Enable/disable IP multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting *string `pulumi:"multicastRouting"`
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal *MulticastPimSmGlobal `pulumi:"pimSmGlobal"`
	// Maximum number of multicast routes.
	RouteLimit *int `pulumi:"routeLimit"`
	// Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
	RouteThreshold *int `pulumi:"routeThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Multicast resource.
type MulticastArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// PIM interfaces. The structure of `interface` block is documented below.
	Interfaces MulticastInterfaceArrayInput
	// Enable/disable IP multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting pulumi.StringPtrInput
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal MulticastPimSmGlobalPtrInput
	// Maximum number of multicast routes.
	RouteLimit pulumi.IntPtrInput
	// Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
	RouteThreshold pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (MulticastArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*multicastArgs)(nil)).Elem()
}

type MulticastInput interface {
	pulumi.Input

	ToMulticastOutput() MulticastOutput
	ToMulticastOutputWithContext(ctx context.Context) MulticastOutput
}

func (*Multicast) ElementType() reflect.Type {
	return reflect.TypeOf((**Multicast)(nil)).Elem()
}

func (i *Multicast) ToMulticastOutput() MulticastOutput {
	return i.ToMulticastOutputWithContext(context.Background())
}

func (i *Multicast) ToMulticastOutputWithContext(ctx context.Context) MulticastOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MulticastOutput)
}

// MulticastArrayInput is an input type that accepts MulticastArray and MulticastArrayOutput values.
// You can construct a concrete instance of `MulticastArrayInput` via:
//
//	MulticastArray{ MulticastArgs{...} }
type MulticastArrayInput interface {
	pulumi.Input

	ToMulticastArrayOutput() MulticastArrayOutput
	ToMulticastArrayOutputWithContext(context.Context) MulticastArrayOutput
}

type MulticastArray []MulticastInput

func (MulticastArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Multicast)(nil)).Elem()
}

func (i MulticastArray) ToMulticastArrayOutput() MulticastArrayOutput {
	return i.ToMulticastArrayOutputWithContext(context.Background())
}

func (i MulticastArray) ToMulticastArrayOutputWithContext(ctx context.Context) MulticastArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MulticastArrayOutput)
}

// MulticastMapInput is an input type that accepts MulticastMap and MulticastMapOutput values.
// You can construct a concrete instance of `MulticastMapInput` via:
//
//	MulticastMap{ "key": MulticastArgs{...} }
type MulticastMapInput interface {
	pulumi.Input

	ToMulticastMapOutput() MulticastMapOutput
	ToMulticastMapOutputWithContext(context.Context) MulticastMapOutput
}

type MulticastMap map[string]MulticastInput

func (MulticastMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Multicast)(nil)).Elem()
}

func (i MulticastMap) ToMulticastMapOutput() MulticastMapOutput {
	return i.ToMulticastMapOutputWithContext(context.Background())
}

func (i MulticastMap) ToMulticastMapOutputWithContext(ctx context.Context) MulticastMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MulticastMapOutput)
}

type MulticastOutput struct{ *pulumi.OutputState }

func (MulticastOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Multicast)(nil)).Elem()
}

func (o MulticastOutput) ToMulticastOutput() MulticastOutput {
	return o
}

func (o MulticastOutput) ToMulticastOutputWithContext(ctx context.Context) MulticastOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o MulticastOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multicast) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// PIM interfaces. The structure of `interface` block is documented below.
func (o MulticastOutput) Interfaces() MulticastInterfaceArrayOutput {
	return o.ApplyT(func(v *Multicast) MulticastInterfaceArrayOutput { return v.Interfaces }).(MulticastInterfaceArrayOutput)
}

// Enable/disable IP multicast routing. Valid values: `enable`, `disable`.
func (o MulticastOutput) MulticastRouting() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicast) pulumi.StringOutput { return v.MulticastRouting }).(pulumi.StringOutput)
}

// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
func (o MulticastOutput) PimSmGlobal() MulticastPimSmGlobalOutput {
	return o.ApplyT(func(v *Multicast) MulticastPimSmGlobalOutput { return v.PimSmGlobal }).(MulticastPimSmGlobalOutput)
}

// Maximum number of multicast routes.
func (o MulticastOutput) RouteLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Multicast) pulumi.IntOutput { return v.RouteLimit }).(pulumi.IntOutput)
}

// Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
func (o MulticastOutput) RouteThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Multicast) pulumi.IntOutput { return v.RouteThreshold }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o MulticastOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multicast) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type MulticastArrayOutput struct{ *pulumi.OutputState }

func (MulticastArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Multicast)(nil)).Elem()
}

func (o MulticastArrayOutput) ToMulticastArrayOutput() MulticastArrayOutput {
	return o
}

func (o MulticastArrayOutput) ToMulticastArrayOutputWithContext(ctx context.Context) MulticastArrayOutput {
	return o
}

func (o MulticastArrayOutput) Index(i pulumi.IntInput) MulticastOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Multicast {
		return vs[0].([]*Multicast)[vs[1].(int)]
	}).(MulticastOutput)
}

type MulticastMapOutput struct{ *pulumi.OutputState }

func (MulticastMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Multicast)(nil)).Elem()
}

func (o MulticastMapOutput) ToMulticastMapOutput() MulticastMapOutput {
	return o
}

func (o MulticastMapOutput) ToMulticastMapOutputWithContext(ctx context.Context) MulticastMapOutput {
	return o
}

func (o MulticastMapOutput) MapIndex(k pulumi.StringInput) MulticastOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Multicast {
		return vs[0].(map[string]*Multicast)[vs[1].(string)]
	}).(MulticastOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MulticastInput)(nil)).Elem(), &Multicast{})
	pulumi.RegisterInputType(reflect.TypeOf((*MulticastArrayInput)(nil)).Elem(), MulticastArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MulticastMapInput)(nil)).Elem(), MulticastMap{})
	pulumi.RegisterOutputType(MulticastOutput{})
	pulumi.RegisterOutputType(MulticastArrayOutput{})
	pulumi.RegisterOutputType(MulticastMapOutput{})
}