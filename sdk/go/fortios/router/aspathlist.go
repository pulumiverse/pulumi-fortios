// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure Autonomous System (AS) path lists.
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
//			_, err := router.NewAspathlist(ctx, "trname", &router.AspathlistArgs{
//				Rules: router.AspathlistRuleArray{
//					&router.AspathlistRuleArgs{
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Router AspathList can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:router/aspathlist:Aspathlist labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:router/aspathlist:Aspathlist labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Aspathlist struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// AS path list name.
	Name pulumi.StringOutput `pulumi:"name"`
	// AS path list rule. The structure of `rule` block is documented below.
	Rules AspathlistRuleArrayOutput `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewAspathlist registers a new resource with the given unique name, arguments, and options.
func NewAspathlist(ctx *pulumi.Context,
	name string, args *AspathlistArgs, opts ...pulumi.ResourceOption) (*Aspathlist, error) {
	if args == nil {
		args = &AspathlistArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Aspathlist
	err := ctx.RegisterResource("fortios:router/aspathlist:Aspathlist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAspathlist gets an existing Aspathlist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAspathlist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AspathlistState, opts ...pulumi.ResourceOption) (*Aspathlist, error) {
	var resource Aspathlist
	err := ctx.ReadResource("fortios:router/aspathlist:Aspathlist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Aspathlist resources.
type aspathlistState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// AS path list name.
	Name *string `pulumi:"name"`
	// AS path list rule. The structure of `rule` block is documented below.
	Rules []AspathlistRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type AspathlistState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// AS path list name.
	Name pulumi.StringPtrInput
	// AS path list rule. The structure of `rule` block is documented below.
	Rules AspathlistRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AspathlistState) ElementType() reflect.Type {
	return reflect.TypeOf((*aspathlistState)(nil)).Elem()
}

type aspathlistArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// AS path list name.
	Name *string `pulumi:"name"`
	// AS path list rule. The structure of `rule` block is documented below.
	Rules []AspathlistRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Aspathlist resource.
type AspathlistArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// AS path list name.
	Name pulumi.StringPtrInput
	// AS path list rule. The structure of `rule` block is documented below.
	Rules AspathlistRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AspathlistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aspathlistArgs)(nil)).Elem()
}

type AspathlistInput interface {
	pulumi.Input

	ToAspathlistOutput() AspathlistOutput
	ToAspathlistOutputWithContext(ctx context.Context) AspathlistOutput
}

func (*Aspathlist) ElementType() reflect.Type {
	return reflect.TypeOf((**Aspathlist)(nil)).Elem()
}

func (i *Aspathlist) ToAspathlistOutput() AspathlistOutput {
	return i.ToAspathlistOutputWithContext(context.Background())
}

func (i *Aspathlist) ToAspathlistOutputWithContext(ctx context.Context) AspathlistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AspathlistOutput)
}

// AspathlistArrayInput is an input type that accepts AspathlistArray and AspathlistArrayOutput values.
// You can construct a concrete instance of `AspathlistArrayInput` via:
//
//	AspathlistArray{ AspathlistArgs{...} }
type AspathlistArrayInput interface {
	pulumi.Input

	ToAspathlistArrayOutput() AspathlistArrayOutput
	ToAspathlistArrayOutputWithContext(context.Context) AspathlistArrayOutput
}

type AspathlistArray []AspathlistInput

func (AspathlistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Aspathlist)(nil)).Elem()
}

func (i AspathlistArray) ToAspathlistArrayOutput() AspathlistArrayOutput {
	return i.ToAspathlistArrayOutputWithContext(context.Background())
}

func (i AspathlistArray) ToAspathlistArrayOutputWithContext(ctx context.Context) AspathlistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AspathlistArrayOutput)
}

// AspathlistMapInput is an input type that accepts AspathlistMap and AspathlistMapOutput values.
// You can construct a concrete instance of `AspathlistMapInput` via:
//
//	AspathlistMap{ "key": AspathlistArgs{...} }
type AspathlistMapInput interface {
	pulumi.Input

	ToAspathlistMapOutput() AspathlistMapOutput
	ToAspathlistMapOutputWithContext(context.Context) AspathlistMapOutput
}

type AspathlistMap map[string]AspathlistInput

func (AspathlistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Aspathlist)(nil)).Elem()
}

func (i AspathlistMap) ToAspathlistMapOutput() AspathlistMapOutput {
	return i.ToAspathlistMapOutputWithContext(context.Background())
}

func (i AspathlistMap) ToAspathlistMapOutputWithContext(ctx context.Context) AspathlistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AspathlistMapOutput)
}

type AspathlistOutput struct{ *pulumi.OutputState }

func (AspathlistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Aspathlist)(nil)).Elem()
}

func (o AspathlistOutput) ToAspathlistOutput() AspathlistOutput {
	return o
}

func (o AspathlistOutput) ToAspathlistOutputWithContext(ctx context.Context) AspathlistOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o AspathlistOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Aspathlist) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// AS path list name.
func (o AspathlistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Aspathlist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// AS path list rule. The structure of `rule` block is documented below.
func (o AspathlistOutput) Rules() AspathlistRuleArrayOutput {
	return o.ApplyT(func(v *Aspathlist) AspathlistRuleArrayOutput { return v.Rules }).(AspathlistRuleArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AspathlistOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Aspathlist) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AspathlistArrayOutput struct{ *pulumi.OutputState }

func (AspathlistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Aspathlist)(nil)).Elem()
}

func (o AspathlistArrayOutput) ToAspathlistArrayOutput() AspathlistArrayOutput {
	return o
}

func (o AspathlistArrayOutput) ToAspathlistArrayOutputWithContext(ctx context.Context) AspathlistArrayOutput {
	return o
}

func (o AspathlistArrayOutput) Index(i pulumi.IntInput) AspathlistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Aspathlist {
		return vs[0].([]*Aspathlist)[vs[1].(int)]
	}).(AspathlistOutput)
}

type AspathlistMapOutput struct{ *pulumi.OutputState }

func (AspathlistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Aspathlist)(nil)).Elem()
}

func (o AspathlistMapOutput) ToAspathlistMapOutput() AspathlistMapOutput {
	return o
}

func (o AspathlistMapOutput) ToAspathlistMapOutputWithContext(ctx context.Context) AspathlistMapOutput {
	return o
}

func (o AspathlistMapOutput) MapIndex(k pulumi.StringInput) AspathlistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Aspathlist {
		return vs[0].(map[string]*Aspathlist)[vs[1].(string)]
	}).(AspathlistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AspathlistInput)(nil)).Elem(), &Aspathlist{})
	pulumi.RegisterInputType(reflect.TypeOf((*AspathlistArrayInput)(nil)).Elem(), AspathlistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AspathlistMapInput)(nil)).Elem(), AspathlistMap{})
	pulumi.RegisterOutputType(AspathlistOutput{})
	pulumi.RegisterOutputType(AspathlistArrayOutput{})
	pulumi.RegisterOutputType(AspathlistMapOutput{})
}
