// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure BFD.
//
// ## Import
//
// # Router Bfd can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/routerBfd:RouterBfd labelname RouterBfd
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/routerBfd:RouterBfd labelname RouterBfd
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type RouterBfd struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// BFD multi-hop template table. The structure of `multihopTemplate` block is documented below.
	MultihopTemplates RouterBfdMultihopTemplateArrayOutput `pulumi:"multihopTemplates"`
	// neighbor The structure of `neighbor` block is documented below.
	Neighbors RouterBfdNeighborArrayOutput `pulumi:"neighbors"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterBfd registers a new resource with the given unique name, arguments, and options.
func NewRouterBfd(ctx *pulumi.Context,
	name string, args *RouterBfdArgs, opts ...pulumi.ResourceOption) (*RouterBfd, error) {
	if args == nil {
		args = &RouterBfdArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource RouterBfd
	err := ctx.RegisterResource("fortios:index/routerBfd:RouterBfd", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterBfd gets an existing RouterBfd resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterBfd(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterBfdState, opts ...pulumi.ResourceOption) (*RouterBfd, error) {
	var resource RouterBfd
	err := ctx.ReadResource("fortios:index/routerBfd:RouterBfd", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterBfd resources.
type routerBfdState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// BFD multi-hop template table. The structure of `multihopTemplate` block is documented below.
	MultihopTemplates []RouterBfdMultihopTemplate `pulumi:"multihopTemplates"`
	// neighbor The structure of `neighbor` block is documented below.
	Neighbors []RouterBfdNeighbor `pulumi:"neighbors"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterBfdState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// BFD multi-hop template table. The structure of `multihopTemplate` block is documented below.
	MultihopTemplates RouterBfdMultihopTemplateArrayInput
	// neighbor The structure of `neighbor` block is documented below.
	Neighbors RouterBfdNeighborArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterBfdState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerBfdState)(nil)).Elem()
}

type routerBfdArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// BFD multi-hop template table. The structure of `multihopTemplate` block is documented below.
	MultihopTemplates []RouterBfdMultihopTemplate `pulumi:"multihopTemplates"`
	// neighbor The structure of `neighbor` block is documented below.
	Neighbors []RouterBfdNeighbor `pulumi:"neighbors"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterBfd resource.
type RouterBfdArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// BFD multi-hop template table. The structure of `multihopTemplate` block is documented below.
	MultihopTemplates RouterBfdMultihopTemplateArrayInput
	// neighbor The structure of `neighbor` block is documented below.
	Neighbors RouterBfdNeighborArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterBfdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerBfdArgs)(nil)).Elem()
}

type RouterBfdInput interface {
	pulumi.Input

	ToRouterBfdOutput() RouterBfdOutput
	ToRouterBfdOutputWithContext(ctx context.Context) RouterBfdOutput
}

func (*RouterBfd) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterBfd)(nil)).Elem()
}

func (i *RouterBfd) ToRouterBfdOutput() RouterBfdOutput {
	return i.ToRouterBfdOutputWithContext(context.Background())
}

func (i *RouterBfd) ToRouterBfdOutputWithContext(ctx context.Context) RouterBfdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBfdOutput)
}

// RouterBfdArrayInput is an input type that accepts RouterBfdArray and RouterBfdArrayOutput values.
// You can construct a concrete instance of `RouterBfdArrayInput` via:
//
//	RouterBfdArray{ RouterBfdArgs{...} }
type RouterBfdArrayInput interface {
	pulumi.Input

	ToRouterBfdArrayOutput() RouterBfdArrayOutput
	ToRouterBfdArrayOutputWithContext(context.Context) RouterBfdArrayOutput
}

type RouterBfdArray []RouterBfdInput

func (RouterBfdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterBfd)(nil)).Elem()
}

func (i RouterBfdArray) ToRouterBfdArrayOutput() RouterBfdArrayOutput {
	return i.ToRouterBfdArrayOutputWithContext(context.Background())
}

func (i RouterBfdArray) ToRouterBfdArrayOutputWithContext(ctx context.Context) RouterBfdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBfdArrayOutput)
}

// RouterBfdMapInput is an input type that accepts RouterBfdMap and RouterBfdMapOutput values.
// You can construct a concrete instance of `RouterBfdMapInput` via:
//
//	RouterBfdMap{ "key": RouterBfdArgs{...} }
type RouterBfdMapInput interface {
	pulumi.Input

	ToRouterBfdMapOutput() RouterBfdMapOutput
	ToRouterBfdMapOutputWithContext(context.Context) RouterBfdMapOutput
}

type RouterBfdMap map[string]RouterBfdInput

func (RouterBfdMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterBfd)(nil)).Elem()
}

func (i RouterBfdMap) ToRouterBfdMapOutput() RouterBfdMapOutput {
	return i.ToRouterBfdMapOutputWithContext(context.Background())
}

func (i RouterBfdMap) ToRouterBfdMapOutputWithContext(ctx context.Context) RouterBfdMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBfdMapOutput)
}

type RouterBfdOutput struct{ *pulumi.OutputState }

func (RouterBfdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterBfd)(nil)).Elem()
}

func (o RouterBfdOutput) ToRouterBfdOutput() RouterBfdOutput {
	return o
}

func (o RouterBfdOutput) ToRouterBfdOutputWithContext(ctx context.Context) RouterBfdOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o RouterBfdOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterBfd) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// BFD multi-hop template table. The structure of `multihopTemplate` block is documented below.
func (o RouterBfdOutput) MultihopTemplates() RouterBfdMultihopTemplateArrayOutput {
	return o.ApplyT(func(v *RouterBfd) RouterBfdMultihopTemplateArrayOutput { return v.MultihopTemplates }).(RouterBfdMultihopTemplateArrayOutput)
}

// neighbor The structure of `neighbor` block is documented below.
func (o RouterBfdOutput) Neighbors() RouterBfdNeighborArrayOutput {
	return o.ApplyT(func(v *RouterBfd) RouterBfdNeighborArrayOutput { return v.Neighbors }).(RouterBfdNeighborArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RouterBfdOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterBfd) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterBfdArrayOutput struct{ *pulumi.OutputState }

func (RouterBfdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterBfd)(nil)).Elem()
}

func (o RouterBfdArrayOutput) ToRouterBfdArrayOutput() RouterBfdArrayOutput {
	return o
}

func (o RouterBfdArrayOutput) ToRouterBfdArrayOutputWithContext(ctx context.Context) RouterBfdArrayOutput {
	return o
}

func (o RouterBfdArrayOutput) Index(i pulumi.IntInput) RouterBfdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterBfd {
		return vs[0].([]*RouterBfd)[vs[1].(int)]
	}).(RouterBfdOutput)
}

type RouterBfdMapOutput struct{ *pulumi.OutputState }

func (RouterBfdMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterBfd)(nil)).Elem()
}

func (o RouterBfdMapOutput) ToRouterBfdMapOutput() RouterBfdMapOutput {
	return o
}

func (o RouterBfdMapOutput) ToRouterBfdMapOutputWithContext(ctx context.Context) RouterBfdMapOutput {
	return o
}

func (o RouterBfdMapOutput) MapIndex(k pulumi.StringInput) RouterBfdOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterBfd {
		return vs[0].(map[string]*RouterBfd)[vs[1].(string)]
	}).(RouterBfdOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterBfdInput)(nil)).Elem(), &RouterBfd{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterBfdArrayInput)(nil)).Elem(), RouterBfdArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterBfdMapInput)(nil)).Elem(), RouterBfdMap{})
	pulumi.RegisterOutputType(RouterBfdOutput{})
	pulumi.RegisterOutputType(RouterBfdArrayOutput{})
	pulumi.RegisterOutputType(RouterBfdMapOutput{})
}
