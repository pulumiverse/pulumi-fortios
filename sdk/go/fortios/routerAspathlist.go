// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Autonomous System (AS) path lists.
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
//			_, err := fortios.NewRouterAspathlist(ctx, "trname", &fortios.RouterAspathlistArgs{
//				Rules: fortios.RouterAspathlistRuleArray{
//					&fortios.RouterAspathlistRuleArgs{
//						Action: pulumi.String("deny"),
//						Regexp: pulumi.String("/d+/n"),
//					},
//				},
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
// # Router AspathList can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/routerAspathlist:RouterAspathlist labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/routerAspathlist:RouterAspathlist labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type RouterAspathlist struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// AS path list name.
	Name pulumi.StringOutput `pulumi:"name"`
	// AS path list rule. The structure of `rule` block is documented below.
	Rules RouterAspathlistRuleArrayOutput `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterAspathlist registers a new resource with the given unique name, arguments, and options.
func NewRouterAspathlist(ctx *pulumi.Context,
	name string, args *RouterAspathlistArgs, opts ...pulumi.ResourceOption) (*RouterAspathlist, error) {
	if args == nil {
		args = &RouterAspathlistArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource RouterAspathlist
	err := ctx.RegisterResource("fortios:index/routerAspathlist:RouterAspathlist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterAspathlist gets an existing RouterAspathlist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterAspathlist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterAspathlistState, opts ...pulumi.ResourceOption) (*RouterAspathlist, error) {
	var resource RouterAspathlist
	err := ctx.ReadResource("fortios:index/routerAspathlist:RouterAspathlist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterAspathlist resources.
type routerAspathlistState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// AS path list name.
	Name *string `pulumi:"name"`
	// AS path list rule. The structure of `rule` block is documented below.
	Rules []RouterAspathlistRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterAspathlistState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// AS path list name.
	Name pulumi.StringPtrInput
	// AS path list rule. The structure of `rule` block is documented below.
	Rules RouterAspathlistRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterAspathlistState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerAspathlistState)(nil)).Elem()
}

type routerAspathlistArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// AS path list name.
	Name *string `pulumi:"name"`
	// AS path list rule. The structure of `rule` block is documented below.
	Rules []RouterAspathlistRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterAspathlist resource.
type RouterAspathlistArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// AS path list name.
	Name pulumi.StringPtrInput
	// AS path list rule. The structure of `rule` block is documented below.
	Rules RouterAspathlistRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterAspathlistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerAspathlistArgs)(nil)).Elem()
}

type RouterAspathlistInput interface {
	pulumi.Input

	ToRouterAspathlistOutput() RouterAspathlistOutput
	ToRouterAspathlistOutputWithContext(ctx context.Context) RouterAspathlistOutput
}

func (*RouterAspathlist) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterAspathlist)(nil)).Elem()
}

func (i *RouterAspathlist) ToRouterAspathlistOutput() RouterAspathlistOutput {
	return i.ToRouterAspathlistOutputWithContext(context.Background())
}

func (i *RouterAspathlist) ToRouterAspathlistOutputWithContext(ctx context.Context) RouterAspathlistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAspathlistOutput)
}

// RouterAspathlistArrayInput is an input type that accepts RouterAspathlistArray and RouterAspathlistArrayOutput values.
// You can construct a concrete instance of `RouterAspathlistArrayInput` via:
//
//	RouterAspathlistArray{ RouterAspathlistArgs{...} }
type RouterAspathlistArrayInput interface {
	pulumi.Input

	ToRouterAspathlistArrayOutput() RouterAspathlistArrayOutput
	ToRouterAspathlistArrayOutputWithContext(context.Context) RouterAspathlistArrayOutput
}

type RouterAspathlistArray []RouterAspathlistInput

func (RouterAspathlistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterAspathlist)(nil)).Elem()
}

func (i RouterAspathlistArray) ToRouterAspathlistArrayOutput() RouterAspathlistArrayOutput {
	return i.ToRouterAspathlistArrayOutputWithContext(context.Background())
}

func (i RouterAspathlistArray) ToRouterAspathlistArrayOutputWithContext(ctx context.Context) RouterAspathlistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAspathlistArrayOutput)
}

// RouterAspathlistMapInput is an input type that accepts RouterAspathlistMap and RouterAspathlistMapOutput values.
// You can construct a concrete instance of `RouterAspathlistMapInput` via:
//
//	RouterAspathlistMap{ "key": RouterAspathlistArgs{...} }
type RouterAspathlistMapInput interface {
	pulumi.Input

	ToRouterAspathlistMapOutput() RouterAspathlistMapOutput
	ToRouterAspathlistMapOutputWithContext(context.Context) RouterAspathlistMapOutput
}

type RouterAspathlistMap map[string]RouterAspathlistInput

func (RouterAspathlistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterAspathlist)(nil)).Elem()
}

func (i RouterAspathlistMap) ToRouterAspathlistMapOutput() RouterAspathlistMapOutput {
	return i.ToRouterAspathlistMapOutputWithContext(context.Background())
}

func (i RouterAspathlistMap) ToRouterAspathlistMapOutputWithContext(ctx context.Context) RouterAspathlistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAspathlistMapOutput)
}

type RouterAspathlistOutput struct{ *pulumi.OutputState }

func (RouterAspathlistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterAspathlist)(nil)).Elem()
}

func (o RouterAspathlistOutput) ToRouterAspathlistOutput() RouterAspathlistOutput {
	return o
}

func (o RouterAspathlistOutput) ToRouterAspathlistOutputWithContext(ctx context.Context) RouterAspathlistOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o RouterAspathlistOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterAspathlist) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// AS path list name.
func (o RouterAspathlistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterAspathlist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// AS path list rule. The structure of `rule` block is documented below.
func (o RouterAspathlistOutput) Rules() RouterAspathlistRuleArrayOutput {
	return o.ApplyT(func(v *RouterAspathlist) RouterAspathlistRuleArrayOutput { return v.Rules }).(RouterAspathlistRuleArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RouterAspathlistOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterAspathlist) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterAspathlistArrayOutput struct{ *pulumi.OutputState }

func (RouterAspathlistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterAspathlist)(nil)).Elem()
}

func (o RouterAspathlistArrayOutput) ToRouterAspathlistArrayOutput() RouterAspathlistArrayOutput {
	return o
}

func (o RouterAspathlistArrayOutput) ToRouterAspathlistArrayOutputWithContext(ctx context.Context) RouterAspathlistArrayOutput {
	return o
}

func (o RouterAspathlistArrayOutput) Index(i pulumi.IntInput) RouterAspathlistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterAspathlist {
		return vs[0].([]*RouterAspathlist)[vs[1].(int)]
	}).(RouterAspathlistOutput)
}

type RouterAspathlistMapOutput struct{ *pulumi.OutputState }

func (RouterAspathlistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterAspathlist)(nil)).Elem()
}

func (o RouterAspathlistMapOutput) ToRouterAspathlistMapOutput() RouterAspathlistMapOutput {
	return o
}

func (o RouterAspathlistMapOutput) ToRouterAspathlistMapOutputWithContext(ctx context.Context) RouterAspathlistMapOutput {
	return o
}

func (o RouterAspathlistMapOutput) MapIndex(k pulumi.StringInput) RouterAspathlistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterAspathlist {
		return vs[0].(map[string]*RouterAspathlist)[vs[1].(string)]
	}).(RouterAspathlistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterAspathlistInput)(nil)).Elem(), &RouterAspathlist{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterAspathlistArrayInput)(nil)).Elem(), RouterAspathlistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterAspathlistMapInput)(nil)).Elem(), RouterAspathlistMap{})
	pulumi.RegisterOutputType(RouterAspathlistOutput{})
	pulumi.RegisterOutputType(RouterAspathlistArrayOutput{})
	pulumi.RegisterOutputType(RouterAspathlistMapOutput{})
}