// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// OSPF neighbor configuration are used when OSPF runs on non-broadcast media
//
// > The provider supports the definition of Neighbor in Router Ospf `RouterOspf`, and also allows the definition of separate Neighbor resources `RouterospfNeighbor`, but do not use a `RouterOspf` with in-line Neighbor in conjunction with any `RouterospfNeighbor` resources, otherwise conflicts and overwrite will occur.
//
// ## Import
//
// # Routerospf Neighbor can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/routerospfNeighbor:RouterospfNeighbor labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/routerospfNeighbor:RouterospfNeighbor labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type RouterospfNeighbor struct {
	pulumi.CustomResourceState

	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost pulumi.IntOutput `pulumi:"cost"`
	// Neighbor entry ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Interface IP address of the neighbor.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Poll interval time in seconds.
	PollInterval pulumi.IntOutput `pulumi:"pollInterval"`
	// Priority.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterospfNeighbor registers a new resource with the given unique name, arguments, and options.
func NewRouterospfNeighbor(ctx *pulumi.Context,
	name string, args *RouterospfNeighborArgs, opts ...pulumi.ResourceOption) (*RouterospfNeighbor, error) {
	if args == nil {
		args = &RouterospfNeighborArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource RouterospfNeighbor
	err := ctx.RegisterResource("fortios:index/routerospfNeighbor:RouterospfNeighbor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterospfNeighbor gets an existing RouterospfNeighbor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterospfNeighbor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterospfNeighborState, opts ...pulumi.ResourceOption) (*RouterospfNeighbor, error) {
	var resource RouterospfNeighbor
	err := ctx.ReadResource("fortios:index/routerospfNeighbor:RouterospfNeighbor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterospfNeighbor resources.
type routerospfNeighborState struct {
	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost *int `pulumi:"cost"`
	// Neighbor entry ID.
	Fosid *int `pulumi:"fosid"`
	// Interface IP address of the neighbor.
	Ip *string `pulumi:"ip"`
	// Poll interval time in seconds.
	PollInterval *int `pulumi:"pollInterval"`
	// Priority.
	Priority *int `pulumi:"priority"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterospfNeighborState struct {
	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost pulumi.IntPtrInput
	// Neighbor entry ID.
	Fosid pulumi.IntPtrInput
	// Interface IP address of the neighbor.
	Ip pulumi.StringPtrInput
	// Poll interval time in seconds.
	PollInterval pulumi.IntPtrInput
	// Priority.
	Priority pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterospfNeighborState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerospfNeighborState)(nil)).Elem()
}

type routerospfNeighborArgs struct {
	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost *int `pulumi:"cost"`
	// Neighbor entry ID.
	Fosid *int `pulumi:"fosid"`
	// Interface IP address of the neighbor.
	Ip *string `pulumi:"ip"`
	// Poll interval time in seconds.
	PollInterval *int `pulumi:"pollInterval"`
	// Priority.
	Priority *int `pulumi:"priority"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterospfNeighbor resource.
type RouterospfNeighborArgs struct {
	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost pulumi.IntPtrInput
	// Neighbor entry ID.
	Fosid pulumi.IntPtrInput
	// Interface IP address of the neighbor.
	Ip pulumi.StringPtrInput
	// Poll interval time in seconds.
	PollInterval pulumi.IntPtrInput
	// Priority.
	Priority pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterospfNeighborArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerospfNeighborArgs)(nil)).Elem()
}

type RouterospfNeighborInput interface {
	pulumi.Input

	ToRouterospfNeighborOutput() RouterospfNeighborOutput
	ToRouterospfNeighborOutputWithContext(ctx context.Context) RouterospfNeighborOutput
}

func (*RouterospfNeighbor) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterospfNeighbor)(nil)).Elem()
}

func (i *RouterospfNeighbor) ToRouterospfNeighborOutput() RouterospfNeighborOutput {
	return i.ToRouterospfNeighborOutputWithContext(context.Background())
}

func (i *RouterospfNeighbor) ToRouterospfNeighborOutputWithContext(ctx context.Context) RouterospfNeighborOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterospfNeighborOutput)
}

// RouterospfNeighborArrayInput is an input type that accepts RouterospfNeighborArray and RouterospfNeighborArrayOutput values.
// You can construct a concrete instance of `RouterospfNeighborArrayInput` via:
//
//	RouterospfNeighborArray{ RouterospfNeighborArgs{...} }
type RouterospfNeighborArrayInput interface {
	pulumi.Input

	ToRouterospfNeighborArrayOutput() RouterospfNeighborArrayOutput
	ToRouterospfNeighborArrayOutputWithContext(context.Context) RouterospfNeighborArrayOutput
}

type RouterospfNeighborArray []RouterospfNeighborInput

func (RouterospfNeighborArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterospfNeighbor)(nil)).Elem()
}

func (i RouterospfNeighborArray) ToRouterospfNeighborArrayOutput() RouterospfNeighborArrayOutput {
	return i.ToRouterospfNeighborArrayOutputWithContext(context.Background())
}

func (i RouterospfNeighborArray) ToRouterospfNeighborArrayOutputWithContext(ctx context.Context) RouterospfNeighborArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterospfNeighborArrayOutput)
}

// RouterospfNeighborMapInput is an input type that accepts RouterospfNeighborMap and RouterospfNeighborMapOutput values.
// You can construct a concrete instance of `RouterospfNeighborMapInput` via:
//
//	RouterospfNeighborMap{ "key": RouterospfNeighborArgs{...} }
type RouterospfNeighborMapInput interface {
	pulumi.Input

	ToRouterospfNeighborMapOutput() RouterospfNeighborMapOutput
	ToRouterospfNeighborMapOutputWithContext(context.Context) RouterospfNeighborMapOutput
}

type RouterospfNeighborMap map[string]RouterospfNeighborInput

func (RouterospfNeighborMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterospfNeighbor)(nil)).Elem()
}

func (i RouterospfNeighborMap) ToRouterospfNeighborMapOutput() RouterospfNeighborMapOutput {
	return i.ToRouterospfNeighborMapOutputWithContext(context.Background())
}

func (i RouterospfNeighborMap) ToRouterospfNeighborMapOutputWithContext(ctx context.Context) RouterospfNeighborMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterospfNeighborMapOutput)
}

type RouterospfNeighborOutput struct{ *pulumi.OutputState }

func (RouterospfNeighborOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterospfNeighbor)(nil)).Elem()
}

func (o RouterospfNeighborOutput) ToRouterospfNeighborOutput() RouterospfNeighborOutput {
	return o
}

func (o RouterospfNeighborOutput) ToRouterospfNeighborOutputWithContext(ctx context.Context) RouterospfNeighborOutput {
	return o
}

// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
func (o RouterospfNeighborOutput) Cost() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfNeighbor) pulumi.IntOutput { return v.Cost }).(pulumi.IntOutput)
}

// Neighbor entry ID.
func (o RouterospfNeighborOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfNeighbor) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Interface IP address of the neighbor.
func (o RouterospfNeighborOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterospfNeighbor) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Poll interval time in seconds.
func (o RouterospfNeighborOutput) PollInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfNeighbor) pulumi.IntOutput { return v.PollInterval }).(pulumi.IntOutput)
}

// Priority.
func (o RouterospfNeighborOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfNeighbor) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RouterospfNeighborOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterospfNeighbor) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterospfNeighborArrayOutput struct{ *pulumi.OutputState }

func (RouterospfNeighborArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterospfNeighbor)(nil)).Elem()
}

func (o RouterospfNeighborArrayOutput) ToRouterospfNeighborArrayOutput() RouterospfNeighborArrayOutput {
	return o
}

func (o RouterospfNeighborArrayOutput) ToRouterospfNeighborArrayOutputWithContext(ctx context.Context) RouterospfNeighborArrayOutput {
	return o
}

func (o RouterospfNeighborArrayOutput) Index(i pulumi.IntInput) RouterospfNeighborOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterospfNeighbor {
		return vs[0].([]*RouterospfNeighbor)[vs[1].(int)]
	}).(RouterospfNeighborOutput)
}

type RouterospfNeighborMapOutput struct{ *pulumi.OutputState }

func (RouterospfNeighborMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterospfNeighbor)(nil)).Elem()
}

func (o RouterospfNeighborMapOutput) ToRouterospfNeighborMapOutput() RouterospfNeighborMapOutput {
	return o
}

func (o RouterospfNeighborMapOutput) ToRouterospfNeighborMapOutputWithContext(ctx context.Context) RouterospfNeighborMapOutput {
	return o
}

func (o RouterospfNeighborMapOutput) MapIndex(k pulumi.StringInput) RouterospfNeighborOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterospfNeighbor {
		return vs[0].(map[string]*RouterospfNeighbor)[vs[1].(string)]
	}).(RouterospfNeighborOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterospfNeighborInput)(nil)).Elem(), &RouterospfNeighbor{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterospfNeighborArrayInput)(nil)).Elem(), RouterospfNeighborArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterospfNeighborMapInput)(nil)).Elem(), RouterospfNeighborMap{})
	pulumi.RegisterOutputType(RouterospfNeighborOutput{})
	pulumi.RegisterOutputType(RouterospfNeighborArrayOutput{})
	pulumi.RegisterOutputType(RouterospfNeighborMapOutput{})
}
