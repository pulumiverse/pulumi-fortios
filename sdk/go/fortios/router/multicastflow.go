// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure multicast-flow.
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
//			_, err := router.NewMulticastflow(ctx, "trname", &router.MulticastflowArgs{
//				Flows: router.MulticastflowFlowArray{
//					&router.MulticastflowFlowArgs{
//						GroupAddr:  pulumi.String("224.252.0.0"),
//						SourceAddr: pulumi.String("224.112.0.0"),
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
// Router MulticastFlow can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:router/multicastflow:Multicastflow labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:router/multicastflow:Multicastflow labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Multicastflow struct {
	pulumi.CustomResourceState

	// Comment.
	Comments pulumi.StringOutput `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Multicast-flow entries. The structure of `flows` block is documented below.
	Flows MulticastflowFlowArrayOutput `pulumi:"flows"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewMulticastflow registers a new resource with the given unique name, arguments, and options.
func NewMulticastflow(ctx *pulumi.Context,
	name string, args *MulticastflowArgs, opts ...pulumi.ResourceOption) (*Multicastflow, error) {
	if args == nil {
		args = &MulticastflowArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Multicastflow
	err := ctx.RegisterResource("fortios:router/multicastflow:Multicastflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMulticastflow gets an existing Multicastflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMulticastflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MulticastflowState, opts ...pulumi.ResourceOption) (*Multicastflow, error) {
	var resource Multicastflow
	err := ctx.ReadResource("fortios:router/multicastflow:Multicastflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Multicastflow resources.
type multicastflowState struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Multicast-flow entries. The structure of `flows` block is documented below.
	Flows []MulticastflowFlow `pulumi:"flows"`
	// Name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type MulticastflowState struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Multicast-flow entries. The structure of `flows` block is documented below.
	Flows MulticastflowFlowArrayInput
	// Name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (MulticastflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*multicastflowState)(nil)).Elem()
}

type multicastflowArgs struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Multicast-flow entries. The structure of `flows` block is documented below.
	Flows []MulticastflowFlow `pulumi:"flows"`
	// Name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Multicastflow resource.
type MulticastflowArgs struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Multicast-flow entries. The structure of `flows` block is documented below.
	Flows MulticastflowFlowArrayInput
	// Name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (MulticastflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*multicastflowArgs)(nil)).Elem()
}

type MulticastflowInput interface {
	pulumi.Input

	ToMulticastflowOutput() MulticastflowOutput
	ToMulticastflowOutputWithContext(ctx context.Context) MulticastflowOutput
}

func (*Multicastflow) ElementType() reflect.Type {
	return reflect.TypeOf((**Multicastflow)(nil)).Elem()
}

func (i *Multicastflow) ToMulticastflowOutput() MulticastflowOutput {
	return i.ToMulticastflowOutputWithContext(context.Background())
}

func (i *Multicastflow) ToMulticastflowOutputWithContext(ctx context.Context) MulticastflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MulticastflowOutput)
}

// MulticastflowArrayInput is an input type that accepts MulticastflowArray and MulticastflowArrayOutput values.
// You can construct a concrete instance of `MulticastflowArrayInput` via:
//
//	MulticastflowArray{ MulticastflowArgs{...} }
type MulticastflowArrayInput interface {
	pulumi.Input

	ToMulticastflowArrayOutput() MulticastflowArrayOutput
	ToMulticastflowArrayOutputWithContext(context.Context) MulticastflowArrayOutput
}

type MulticastflowArray []MulticastflowInput

func (MulticastflowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Multicastflow)(nil)).Elem()
}

func (i MulticastflowArray) ToMulticastflowArrayOutput() MulticastflowArrayOutput {
	return i.ToMulticastflowArrayOutputWithContext(context.Background())
}

func (i MulticastflowArray) ToMulticastflowArrayOutputWithContext(ctx context.Context) MulticastflowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MulticastflowArrayOutput)
}

// MulticastflowMapInput is an input type that accepts MulticastflowMap and MulticastflowMapOutput values.
// You can construct a concrete instance of `MulticastflowMapInput` via:
//
//	MulticastflowMap{ "key": MulticastflowArgs{...} }
type MulticastflowMapInput interface {
	pulumi.Input

	ToMulticastflowMapOutput() MulticastflowMapOutput
	ToMulticastflowMapOutputWithContext(context.Context) MulticastflowMapOutput
}

type MulticastflowMap map[string]MulticastflowInput

func (MulticastflowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Multicastflow)(nil)).Elem()
}

func (i MulticastflowMap) ToMulticastflowMapOutput() MulticastflowMapOutput {
	return i.ToMulticastflowMapOutputWithContext(context.Background())
}

func (i MulticastflowMap) ToMulticastflowMapOutputWithContext(ctx context.Context) MulticastflowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MulticastflowMapOutput)
}

type MulticastflowOutput struct{ *pulumi.OutputState }

func (MulticastflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Multicastflow)(nil)).Elem()
}

func (o MulticastflowOutput) ToMulticastflowOutput() MulticastflowOutput {
	return o
}

func (o MulticastflowOutput) ToMulticastflowOutputWithContext(ctx context.Context) MulticastflowOutput {
	return o
}

// Comment.
func (o MulticastflowOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastflow) pulumi.StringOutput { return v.Comments }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o MulticastflowOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multicastflow) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Multicast-flow entries. The structure of `flows` block is documented below.
func (o MulticastflowOutput) Flows() MulticastflowFlowArrayOutput {
	return o.ApplyT(func(v *Multicastflow) MulticastflowFlowArrayOutput { return v.Flows }).(MulticastflowFlowArrayOutput)
}

// Name.
func (o MulticastflowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicastflow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o MulticastflowOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multicastflow) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type MulticastflowArrayOutput struct{ *pulumi.OutputState }

func (MulticastflowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Multicastflow)(nil)).Elem()
}

func (o MulticastflowArrayOutput) ToMulticastflowArrayOutput() MulticastflowArrayOutput {
	return o
}

func (o MulticastflowArrayOutput) ToMulticastflowArrayOutputWithContext(ctx context.Context) MulticastflowArrayOutput {
	return o
}

func (o MulticastflowArrayOutput) Index(i pulumi.IntInput) MulticastflowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Multicastflow {
		return vs[0].([]*Multicastflow)[vs[1].(int)]
	}).(MulticastflowOutput)
}

type MulticastflowMapOutput struct{ *pulumi.OutputState }

func (MulticastflowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Multicastflow)(nil)).Elem()
}

func (o MulticastflowMapOutput) ToMulticastflowMapOutput() MulticastflowMapOutput {
	return o
}

func (o MulticastflowMapOutput) ToMulticastflowMapOutputWithContext(ctx context.Context) MulticastflowMapOutput {
	return o
}

func (o MulticastflowMapOutput) MapIndex(k pulumi.StringInput) MulticastflowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Multicastflow {
		return vs[0].(map[string]*Multicastflow)[vs[1].(string)]
	}).(MulticastflowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MulticastflowInput)(nil)).Elem(), &Multicastflow{})
	pulumi.RegisterInputType(reflect.TypeOf((*MulticastflowArrayInput)(nil)).Elem(), MulticastflowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MulticastflowMapInput)(nil)).Elem(), MulticastflowMap{})
	pulumi.RegisterOutputType(MulticastflowOutput{})
	pulumi.RegisterOutputType(MulticastflowArrayOutput{})
	pulumi.RegisterOutputType(MulticastflowMapOutput{})
}
