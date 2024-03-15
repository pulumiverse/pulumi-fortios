// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bgp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// BGP IPv6 network table.
//
// > The provider supports the definition of Network6 in Router Bgp `router.Bgp`, and also allows the definition of separate Network6 resources `router/bgp.Network6`, but do not use a `router.Bgp` with in-line Network6 in conjunction with any `router/bgp.Network6` resources, otherwise conflicts and overwrite will occur.
//
// ## Import
//
// Routerbgp Network6 can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:router/bgp/network6:Network6 labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:router/bgp/network6:Network6 labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Network6 struct {
	pulumi.CustomResourceState

	// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
	Backdoor pulumi.StringOutput `pulumi:"backdoor"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
	NetworkImportCheck pulumi.StringOutput `pulumi:"networkImportCheck"`
	// Network IPv6 prefix.
	Prefix6 pulumi.StringOutput `pulumi:"prefix6"`
	// Route map to modify generated route.
	RouteMap pulumi.StringOutput `pulumi:"routeMap"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewNetwork6 registers a new resource with the given unique name, arguments, and options.
func NewNetwork6(ctx *pulumi.Context,
	name string, args *Network6Args, opts ...pulumi.ResourceOption) (*Network6, error) {
	if args == nil {
		args = &Network6Args{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Network6
	err := ctx.RegisterResource("fortios:router/bgp/network6:Network6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetwork6 gets an existing Network6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetwork6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Network6State, opts ...pulumi.ResourceOption) (*Network6, error) {
	var resource Network6
	err := ctx.ReadResource("fortios:router/bgp/network6:Network6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Network6 resources.
type network6State struct {
	// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
	Backdoor *string `pulumi:"backdoor"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
	NetworkImportCheck *string `pulumi:"networkImportCheck"`
	// Network IPv6 prefix.
	Prefix6 *string `pulumi:"prefix6"`
	// Route map to modify generated route.
	RouteMap *string `pulumi:"routeMap"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Network6State struct {
	// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
	Backdoor pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
	NetworkImportCheck pulumi.StringPtrInput
	// Network IPv6 prefix.
	Prefix6 pulumi.StringPtrInput
	// Route map to modify generated route.
	RouteMap pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Network6State) ElementType() reflect.Type {
	return reflect.TypeOf((*network6State)(nil)).Elem()
}

type network6Args struct {
	// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
	Backdoor *string `pulumi:"backdoor"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
	NetworkImportCheck *string `pulumi:"networkImportCheck"`
	// Network IPv6 prefix.
	Prefix6 *string `pulumi:"prefix6"`
	// Route map to modify generated route.
	RouteMap *string `pulumi:"routeMap"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Network6 resource.
type Network6Args struct {
	// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
	Backdoor pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
	NetworkImportCheck pulumi.StringPtrInput
	// Network IPv6 prefix.
	Prefix6 pulumi.StringPtrInput
	// Route map to modify generated route.
	RouteMap pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Network6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*network6Args)(nil)).Elem()
}

type Network6Input interface {
	pulumi.Input

	ToNetwork6Output() Network6Output
	ToNetwork6OutputWithContext(ctx context.Context) Network6Output
}

func (*Network6) ElementType() reflect.Type {
	return reflect.TypeOf((**Network6)(nil)).Elem()
}

func (i *Network6) ToNetwork6Output() Network6Output {
	return i.ToNetwork6OutputWithContext(context.Background())
}

func (i *Network6) ToNetwork6OutputWithContext(ctx context.Context) Network6Output {
	return pulumi.ToOutputWithContext(ctx, i).(Network6Output)
}

// Network6ArrayInput is an input type that accepts Network6Array and Network6ArrayOutput values.
// You can construct a concrete instance of `Network6ArrayInput` via:
//
//	Network6Array{ Network6Args{...} }
type Network6ArrayInput interface {
	pulumi.Input

	ToNetwork6ArrayOutput() Network6ArrayOutput
	ToNetwork6ArrayOutputWithContext(context.Context) Network6ArrayOutput
}

type Network6Array []Network6Input

func (Network6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Network6)(nil)).Elem()
}

func (i Network6Array) ToNetwork6ArrayOutput() Network6ArrayOutput {
	return i.ToNetwork6ArrayOutputWithContext(context.Background())
}

func (i Network6Array) ToNetwork6ArrayOutputWithContext(ctx context.Context) Network6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Network6ArrayOutput)
}

// Network6MapInput is an input type that accepts Network6Map and Network6MapOutput values.
// You can construct a concrete instance of `Network6MapInput` via:
//
//	Network6Map{ "key": Network6Args{...} }
type Network6MapInput interface {
	pulumi.Input

	ToNetwork6MapOutput() Network6MapOutput
	ToNetwork6MapOutputWithContext(context.Context) Network6MapOutput
}

type Network6Map map[string]Network6Input

func (Network6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Network6)(nil)).Elem()
}

func (i Network6Map) ToNetwork6MapOutput() Network6MapOutput {
	return i.ToNetwork6MapOutputWithContext(context.Background())
}

func (i Network6Map) ToNetwork6MapOutputWithContext(ctx context.Context) Network6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Network6MapOutput)
}

type Network6Output struct{ *pulumi.OutputState }

func (Network6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Network6)(nil)).Elem()
}

func (o Network6Output) ToNetwork6Output() Network6Output {
	return o
}

func (o Network6Output) ToNetwork6OutputWithContext(ctx context.Context) Network6Output {
	return o
}

// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
func (o Network6Output) Backdoor() pulumi.StringOutput {
	return o.ApplyT(func(v *Network6) pulumi.StringOutput { return v.Backdoor }).(pulumi.StringOutput)
}

// ID.
func (o Network6Output) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Network6) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
func (o Network6Output) NetworkImportCheck() pulumi.StringOutput {
	return o.ApplyT(func(v *Network6) pulumi.StringOutput { return v.NetworkImportCheck }).(pulumi.StringOutput)
}

// Network IPv6 prefix.
func (o Network6Output) Prefix6() pulumi.StringOutput {
	return o.ApplyT(func(v *Network6) pulumi.StringOutput { return v.Prefix6 }).(pulumi.StringOutput)
}

// Route map to modify generated route.
func (o Network6Output) RouteMap() pulumi.StringOutput {
	return o.ApplyT(func(v *Network6) pulumi.StringOutput { return v.RouteMap }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Network6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Network6ArrayOutput struct{ *pulumi.OutputState }

func (Network6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Network6)(nil)).Elem()
}

func (o Network6ArrayOutput) ToNetwork6ArrayOutput() Network6ArrayOutput {
	return o
}

func (o Network6ArrayOutput) ToNetwork6ArrayOutputWithContext(ctx context.Context) Network6ArrayOutput {
	return o
}

func (o Network6ArrayOutput) Index(i pulumi.IntInput) Network6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Network6 {
		return vs[0].([]*Network6)[vs[1].(int)]
	}).(Network6Output)
}

type Network6MapOutput struct{ *pulumi.OutputState }

func (Network6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Network6)(nil)).Elem()
}

func (o Network6MapOutput) ToNetwork6MapOutput() Network6MapOutput {
	return o
}

func (o Network6MapOutput) ToNetwork6MapOutputWithContext(ctx context.Context) Network6MapOutput {
	return o
}

func (o Network6MapOutput) MapIndex(k pulumi.StringInput) Network6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Network6 {
		return vs[0].(map[string]*Network6)[vs[1].(string)]
	}).(Network6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Network6Input)(nil)).Elem(), &Network6{})
	pulumi.RegisterInputType(reflect.TypeOf((*Network6ArrayInput)(nil)).Elem(), Network6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Network6MapInput)(nil)).Elem(), Network6Map{})
	pulumi.RegisterOutputType(Network6Output{})
	pulumi.RegisterOutputType(Network6ArrayOutput{})
	pulumi.RegisterOutputType(Network6MapOutput{})
}