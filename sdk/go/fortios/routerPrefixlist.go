// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv4 prefix lists.
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
//			_, err := fortios.NewRouterPrefixlist(ctx, "trname", nil)
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
// # Router PrefixList can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/routerPrefixlist:RouterPrefixlist labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/routerPrefixlist:RouterPrefixlist labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type RouterPrefixlist struct {
	pulumi.CustomResourceState

	// Comment.
	Comments pulumi.StringOutput `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// IPv4 prefix list rule. The structure of `rule` block is documented below.
	Rules RouterPrefixlistRuleArrayOutput `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterPrefixlist registers a new resource with the given unique name, arguments, and options.
func NewRouterPrefixlist(ctx *pulumi.Context,
	name string, args *RouterPrefixlistArgs, opts ...pulumi.ResourceOption) (*RouterPrefixlist, error) {
	if args == nil {
		args = &RouterPrefixlistArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource RouterPrefixlist
	err := ctx.RegisterResource("fortios:index/routerPrefixlist:RouterPrefixlist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterPrefixlist gets an existing RouterPrefixlist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterPrefixlist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterPrefixlistState, opts ...pulumi.ResourceOption) (*RouterPrefixlist, error) {
	var resource RouterPrefixlist
	err := ctx.ReadResource("fortios:index/routerPrefixlist:RouterPrefixlist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterPrefixlist resources.
type routerPrefixlistState struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// IPv4 prefix list rule. The structure of `rule` block is documented below.
	Rules []RouterPrefixlistRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterPrefixlistState struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// IPv4 prefix list rule. The structure of `rule` block is documented below.
	Rules RouterPrefixlistRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterPrefixlistState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerPrefixlistState)(nil)).Elem()
}

type routerPrefixlistArgs struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// IPv4 prefix list rule. The structure of `rule` block is documented below.
	Rules []RouterPrefixlistRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterPrefixlist resource.
type RouterPrefixlistArgs struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// IPv4 prefix list rule. The structure of `rule` block is documented below.
	Rules RouterPrefixlistRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterPrefixlistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerPrefixlistArgs)(nil)).Elem()
}

type RouterPrefixlistInput interface {
	pulumi.Input

	ToRouterPrefixlistOutput() RouterPrefixlistOutput
	ToRouterPrefixlistOutputWithContext(ctx context.Context) RouterPrefixlistOutput
}

func (*RouterPrefixlist) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterPrefixlist)(nil)).Elem()
}

func (i *RouterPrefixlist) ToRouterPrefixlistOutput() RouterPrefixlistOutput {
	return i.ToRouterPrefixlistOutputWithContext(context.Background())
}

func (i *RouterPrefixlist) ToRouterPrefixlistOutputWithContext(ctx context.Context) RouterPrefixlistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterPrefixlistOutput)
}

// RouterPrefixlistArrayInput is an input type that accepts RouterPrefixlistArray and RouterPrefixlistArrayOutput values.
// You can construct a concrete instance of `RouterPrefixlistArrayInput` via:
//
//	RouterPrefixlistArray{ RouterPrefixlistArgs{...} }
type RouterPrefixlistArrayInput interface {
	pulumi.Input

	ToRouterPrefixlistArrayOutput() RouterPrefixlistArrayOutput
	ToRouterPrefixlistArrayOutputWithContext(context.Context) RouterPrefixlistArrayOutput
}

type RouterPrefixlistArray []RouterPrefixlistInput

func (RouterPrefixlistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterPrefixlist)(nil)).Elem()
}

func (i RouterPrefixlistArray) ToRouterPrefixlistArrayOutput() RouterPrefixlistArrayOutput {
	return i.ToRouterPrefixlistArrayOutputWithContext(context.Background())
}

func (i RouterPrefixlistArray) ToRouterPrefixlistArrayOutputWithContext(ctx context.Context) RouterPrefixlistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterPrefixlistArrayOutput)
}

// RouterPrefixlistMapInput is an input type that accepts RouterPrefixlistMap and RouterPrefixlistMapOutput values.
// You can construct a concrete instance of `RouterPrefixlistMapInput` via:
//
//	RouterPrefixlistMap{ "key": RouterPrefixlistArgs{...} }
type RouterPrefixlistMapInput interface {
	pulumi.Input

	ToRouterPrefixlistMapOutput() RouterPrefixlistMapOutput
	ToRouterPrefixlistMapOutputWithContext(context.Context) RouterPrefixlistMapOutput
}

type RouterPrefixlistMap map[string]RouterPrefixlistInput

func (RouterPrefixlistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterPrefixlist)(nil)).Elem()
}

func (i RouterPrefixlistMap) ToRouterPrefixlistMapOutput() RouterPrefixlistMapOutput {
	return i.ToRouterPrefixlistMapOutputWithContext(context.Background())
}

func (i RouterPrefixlistMap) ToRouterPrefixlistMapOutputWithContext(ctx context.Context) RouterPrefixlistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterPrefixlistMapOutput)
}

type RouterPrefixlistOutput struct{ *pulumi.OutputState }

func (RouterPrefixlistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterPrefixlist)(nil)).Elem()
}

func (o RouterPrefixlistOutput) ToRouterPrefixlistOutput() RouterPrefixlistOutput {
	return o
}

func (o RouterPrefixlistOutput) ToRouterPrefixlistOutputWithContext(ctx context.Context) RouterPrefixlistOutput {
	return o
}

// Comment.
func (o RouterPrefixlistOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterPrefixlist) pulumi.StringOutput { return v.Comments }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o RouterPrefixlistOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterPrefixlist) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Name.
func (o RouterPrefixlistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterPrefixlist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// IPv4 prefix list rule. The structure of `rule` block is documented below.
func (o RouterPrefixlistOutput) Rules() RouterPrefixlistRuleArrayOutput {
	return o.ApplyT(func(v *RouterPrefixlist) RouterPrefixlistRuleArrayOutput { return v.Rules }).(RouterPrefixlistRuleArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RouterPrefixlistOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterPrefixlist) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterPrefixlistArrayOutput struct{ *pulumi.OutputState }

func (RouterPrefixlistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterPrefixlist)(nil)).Elem()
}

func (o RouterPrefixlistArrayOutput) ToRouterPrefixlistArrayOutput() RouterPrefixlistArrayOutput {
	return o
}

func (o RouterPrefixlistArrayOutput) ToRouterPrefixlistArrayOutputWithContext(ctx context.Context) RouterPrefixlistArrayOutput {
	return o
}

func (o RouterPrefixlistArrayOutput) Index(i pulumi.IntInput) RouterPrefixlistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterPrefixlist {
		return vs[0].([]*RouterPrefixlist)[vs[1].(int)]
	}).(RouterPrefixlistOutput)
}

type RouterPrefixlistMapOutput struct{ *pulumi.OutputState }

func (RouterPrefixlistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterPrefixlist)(nil)).Elem()
}

func (o RouterPrefixlistMapOutput) ToRouterPrefixlistMapOutput() RouterPrefixlistMapOutput {
	return o
}

func (o RouterPrefixlistMapOutput) ToRouterPrefixlistMapOutputWithContext(ctx context.Context) RouterPrefixlistMapOutput {
	return o
}

func (o RouterPrefixlistMapOutput) MapIndex(k pulumi.StringInput) RouterPrefixlistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterPrefixlist {
		return vs[0].(map[string]*RouterPrefixlist)[vs[1].(string)]
	}).(RouterPrefixlistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterPrefixlistInput)(nil)).Elem(), &RouterPrefixlist{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterPrefixlistArrayInput)(nil)).Elem(), RouterPrefixlistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterPrefixlistMapInput)(nil)).Elem(), RouterPrefixlistMap{})
	pulumi.RegisterOutputType(RouterPrefixlistOutput{})
	pulumi.RegisterOutputType(RouterPrefixlistArrayOutput{})
	pulumi.RegisterOutputType(RouterPrefixlistMapOutput{})
}
