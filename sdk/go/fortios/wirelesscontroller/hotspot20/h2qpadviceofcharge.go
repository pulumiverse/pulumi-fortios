// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hotspot20

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure advice of charge. Applies to FortiOS Version `>= 7.0.2`.
//
// ## Import
//
// WirelessControllerHotspot20 H2QpAdviceOfCharge can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/hotspot20/h2qpadviceofcharge:H2qpadviceofcharge labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/hotspot20/h2qpadviceofcharge:H2qpadviceofcharge labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type H2qpadviceofcharge struct {
	pulumi.CustomResourceState

	// AOC list. The structure of `aocList` block is documented below.
	AocLists H2qpadviceofchargeAocListArrayOutput `pulumi:"aocLists"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Plan name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewH2qpadviceofcharge registers a new resource with the given unique name, arguments, and options.
func NewH2qpadviceofcharge(ctx *pulumi.Context,
	name string, args *H2qpadviceofchargeArgs, opts ...pulumi.ResourceOption) (*H2qpadviceofcharge, error) {
	if args == nil {
		args = &H2qpadviceofchargeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource H2qpadviceofcharge
	err := ctx.RegisterResource("fortios:wirelesscontroller/hotspot20/h2qpadviceofcharge:H2qpadviceofcharge", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetH2qpadviceofcharge gets an existing H2qpadviceofcharge resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetH2qpadviceofcharge(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *H2qpadviceofchargeState, opts ...pulumi.ResourceOption) (*H2qpadviceofcharge, error) {
	var resource H2qpadviceofcharge
	err := ctx.ReadResource("fortios:wirelesscontroller/hotspot20/h2qpadviceofcharge:H2qpadviceofcharge", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering H2qpadviceofcharge resources.
type h2qpadviceofchargeState struct {
	// AOC list. The structure of `aocList` block is documented below.
	AocLists []H2qpadviceofchargeAocList `pulumi:"aocLists"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Plan name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type H2qpadviceofchargeState struct {
	// AOC list. The structure of `aocList` block is documented below.
	AocLists H2qpadviceofchargeAocListArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Plan name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (H2qpadviceofchargeState) ElementType() reflect.Type {
	return reflect.TypeOf((*h2qpadviceofchargeState)(nil)).Elem()
}

type h2qpadviceofchargeArgs struct {
	// AOC list. The structure of `aocList` block is documented below.
	AocLists []H2qpadviceofchargeAocList `pulumi:"aocLists"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Plan name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a H2qpadviceofcharge resource.
type H2qpadviceofchargeArgs struct {
	// AOC list. The structure of `aocList` block is documented below.
	AocLists H2qpadviceofchargeAocListArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Plan name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (H2qpadviceofchargeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*h2qpadviceofchargeArgs)(nil)).Elem()
}

type H2qpadviceofchargeInput interface {
	pulumi.Input

	ToH2qpadviceofchargeOutput() H2qpadviceofchargeOutput
	ToH2qpadviceofchargeOutputWithContext(ctx context.Context) H2qpadviceofchargeOutput
}

func (*H2qpadviceofcharge) ElementType() reflect.Type {
	return reflect.TypeOf((**H2qpadviceofcharge)(nil)).Elem()
}

func (i *H2qpadviceofcharge) ToH2qpadviceofchargeOutput() H2qpadviceofchargeOutput {
	return i.ToH2qpadviceofchargeOutputWithContext(context.Background())
}

func (i *H2qpadviceofcharge) ToH2qpadviceofchargeOutputWithContext(ctx context.Context) H2qpadviceofchargeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H2qpadviceofchargeOutput)
}

// H2qpadviceofchargeArrayInput is an input type that accepts H2qpadviceofchargeArray and H2qpadviceofchargeArrayOutput values.
// You can construct a concrete instance of `H2qpadviceofchargeArrayInput` via:
//
//	H2qpadviceofchargeArray{ H2qpadviceofchargeArgs{...} }
type H2qpadviceofchargeArrayInput interface {
	pulumi.Input

	ToH2qpadviceofchargeArrayOutput() H2qpadviceofchargeArrayOutput
	ToH2qpadviceofchargeArrayOutputWithContext(context.Context) H2qpadviceofchargeArrayOutput
}

type H2qpadviceofchargeArray []H2qpadviceofchargeInput

func (H2qpadviceofchargeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*H2qpadviceofcharge)(nil)).Elem()
}

func (i H2qpadviceofchargeArray) ToH2qpadviceofchargeArrayOutput() H2qpadviceofchargeArrayOutput {
	return i.ToH2qpadviceofchargeArrayOutputWithContext(context.Background())
}

func (i H2qpadviceofchargeArray) ToH2qpadviceofchargeArrayOutputWithContext(ctx context.Context) H2qpadviceofchargeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H2qpadviceofchargeArrayOutput)
}

// H2qpadviceofchargeMapInput is an input type that accepts H2qpadviceofchargeMap and H2qpadviceofchargeMapOutput values.
// You can construct a concrete instance of `H2qpadviceofchargeMapInput` via:
//
//	H2qpadviceofchargeMap{ "key": H2qpadviceofchargeArgs{...} }
type H2qpadviceofchargeMapInput interface {
	pulumi.Input

	ToH2qpadviceofchargeMapOutput() H2qpadviceofchargeMapOutput
	ToH2qpadviceofchargeMapOutputWithContext(context.Context) H2qpadviceofchargeMapOutput
}

type H2qpadviceofchargeMap map[string]H2qpadviceofchargeInput

func (H2qpadviceofchargeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*H2qpadviceofcharge)(nil)).Elem()
}

func (i H2qpadviceofchargeMap) ToH2qpadviceofchargeMapOutput() H2qpadviceofchargeMapOutput {
	return i.ToH2qpadviceofchargeMapOutputWithContext(context.Background())
}

func (i H2qpadviceofchargeMap) ToH2qpadviceofchargeMapOutputWithContext(ctx context.Context) H2qpadviceofchargeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H2qpadviceofchargeMapOutput)
}

type H2qpadviceofchargeOutput struct{ *pulumi.OutputState }

func (H2qpadviceofchargeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**H2qpadviceofcharge)(nil)).Elem()
}

func (o H2qpadviceofchargeOutput) ToH2qpadviceofchargeOutput() H2qpadviceofchargeOutput {
	return o
}

func (o H2qpadviceofchargeOutput) ToH2qpadviceofchargeOutputWithContext(ctx context.Context) H2qpadviceofchargeOutput {
	return o
}

// AOC list. The structure of `aocList` block is documented below.
func (o H2qpadviceofchargeOutput) AocLists() H2qpadviceofchargeAocListArrayOutput {
	return o.ApplyT(func(v *H2qpadviceofcharge) H2qpadviceofchargeAocListArrayOutput { return v.AocLists }).(H2qpadviceofchargeAocListArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o H2qpadviceofchargeOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *H2qpadviceofcharge) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Plan name.
func (o H2qpadviceofchargeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qpadviceofcharge) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o H2qpadviceofchargeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *H2qpadviceofcharge) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type H2qpadviceofchargeArrayOutput struct{ *pulumi.OutputState }

func (H2qpadviceofchargeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*H2qpadviceofcharge)(nil)).Elem()
}

func (o H2qpadviceofchargeArrayOutput) ToH2qpadviceofchargeArrayOutput() H2qpadviceofchargeArrayOutput {
	return o
}

func (o H2qpadviceofchargeArrayOutput) ToH2qpadviceofchargeArrayOutputWithContext(ctx context.Context) H2qpadviceofchargeArrayOutput {
	return o
}

func (o H2qpadviceofchargeArrayOutput) Index(i pulumi.IntInput) H2qpadviceofchargeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *H2qpadviceofcharge {
		return vs[0].([]*H2qpadviceofcharge)[vs[1].(int)]
	}).(H2qpadviceofchargeOutput)
}

type H2qpadviceofchargeMapOutput struct{ *pulumi.OutputState }

func (H2qpadviceofchargeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*H2qpadviceofcharge)(nil)).Elem()
}

func (o H2qpadviceofchargeMapOutput) ToH2qpadviceofchargeMapOutput() H2qpadviceofchargeMapOutput {
	return o
}

func (o H2qpadviceofchargeMapOutput) ToH2qpadviceofchargeMapOutputWithContext(ctx context.Context) H2qpadviceofchargeMapOutput {
	return o
}

func (o H2qpadviceofchargeMapOutput) MapIndex(k pulumi.StringInput) H2qpadviceofchargeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *H2qpadviceofcharge {
		return vs[0].(map[string]*H2qpadviceofcharge)[vs[1].(string)]
	}).(H2qpadviceofchargeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*H2qpadviceofchargeInput)(nil)).Elem(), &H2qpadviceofcharge{})
	pulumi.RegisterInputType(reflect.TypeOf((*H2qpadviceofchargeArrayInput)(nil)).Elem(), H2qpadviceofchargeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*H2qpadviceofchargeMapInput)(nil)).Elem(), H2qpadviceofchargeMap{})
	pulumi.RegisterOutputType(H2qpadviceofchargeOutput{})
	pulumi.RegisterOutputType(H2qpadviceofchargeArrayOutput{})
	pulumi.RegisterOutputType(H2qpadviceofchargeMapOutput{})
}
