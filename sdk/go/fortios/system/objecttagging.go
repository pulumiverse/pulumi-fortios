// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure object tagging.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewObjecttagging(ctx, "trname", &system.ObjecttaggingArgs{
//				Address:   pulumi.String("disable"),
//				Category:  pulumi.String("s1"),
//				Color:     pulumi.Int(0),
//				Device:    pulumi.String("mandatory"),
//				Interface: pulumi.String("disable"),
//				Multiple:  pulumi.String("enable"),
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
// System ObjectTagging can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/objecttagging:Objecttagging labelname {{category}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/objecttagging:Objecttagging labelname {{category}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Objecttagging struct {
	pulumi.CustomResourceState

	// Address. Valid values: `disable`, `mandatory`, `optional`.
	Address pulumi.StringOutput `pulumi:"address"`
	// Tag Category.
	Category pulumi.StringOutput `pulumi:"category"`
	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Device. Valid values: `disable`, `mandatory`, `optional`.
	Device pulumi.StringOutput `pulumi:"device"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Interface. Valid values: `disable`, `mandatory`, `optional`.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Allow multiple tag selection. Valid values: `enable`, `disable`.
	Multiple pulumi.StringOutput `pulumi:"multiple"`
	// Tags. The structure of `tags` block is documented below.
	Tags ObjecttaggingTagArrayOutput `pulumi:"tags"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewObjecttagging registers a new resource with the given unique name, arguments, and options.
func NewObjecttagging(ctx *pulumi.Context,
	name string, args *ObjecttaggingArgs, opts ...pulumi.ResourceOption) (*Objecttagging, error) {
	if args == nil {
		args = &ObjecttaggingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Objecttagging
	err := ctx.RegisterResource("fortios:system/objecttagging:Objecttagging", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjecttagging gets an existing Objecttagging resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjecttagging(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjecttaggingState, opts ...pulumi.ResourceOption) (*Objecttagging, error) {
	var resource Objecttagging
	err := ctx.ReadResource("fortios:system/objecttagging:Objecttagging", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Objecttagging resources.
type objecttaggingState struct {
	// Address. Valid values: `disable`, `mandatory`, `optional`.
	Address *string `pulumi:"address"`
	// Tag Category.
	Category *string `pulumi:"category"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Device. Valid values: `disable`, `mandatory`, `optional`.
	Device *string `pulumi:"device"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Interface. Valid values: `disable`, `mandatory`, `optional`.
	Interface *string `pulumi:"interface"`
	// Allow multiple tag selection. Valid values: `enable`, `disable`.
	Multiple *string `pulumi:"multiple"`
	// Tags. The structure of `tags` block is documented below.
	Tags []ObjecttaggingTag `pulumi:"tags"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ObjecttaggingState struct {
	// Address. Valid values: `disable`, `mandatory`, `optional`.
	Address pulumi.StringPtrInput
	// Tag Category.
	Category pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Device. Valid values: `disable`, `mandatory`, `optional`.
	Device pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Interface. Valid values: `disable`, `mandatory`, `optional`.
	Interface pulumi.StringPtrInput
	// Allow multiple tag selection. Valid values: `enable`, `disable`.
	Multiple pulumi.StringPtrInput
	// Tags. The structure of `tags` block is documented below.
	Tags ObjecttaggingTagArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ObjecttaggingState) ElementType() reflect.Type {
	return reflect.TypeOf((*objecttaggingState)(nil)).Elem()
}

type objecttaggingArgs struct {
	// Address. Valid values: `disable`, `mandatory`, `optional`.
	Address *string `pulumi:"address"`
	// Tag Category.
	Category *string `pulumi:"category"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Device. Valid values: `disable`, `mandatory`, `optional`.
	Device *string `pulumi:"device"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Interface. Valid values: `disable`, `mandatory`, `optional`.
	Interface *string `pulumi:"interface"`
	// Allow multiple tag selection. Valid values: `enable`, `disable`.
	Multiple *string `pulumi:"multiple"`
	// Tags. The structure of `tags` block is documented below.
	Tags []ObjecttaggingTag `pulumi:"tags"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Objecttagging resource.
type ObjecttaggingArgs struct {
	// Address. Valid values: `disable`, `mandatory`, `optional`.
	Address pulumi.StringPtrInput
	// Tag Category.
	Category pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Device. Valid values: `disable`, `mandatory`, `optional`.
	Device pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Interface. Valid values: `disable`, `mandatory`, `optional`.
	Interface pulumi.StringPtrInput
	// Allow multiple tag selection. Valid values: `enable`, `disable`.
	Multiple pulumi.StringPtrInput
	// Tags. The structure of `tags` block is documented below.
	Tags ObjecttaggingTagArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ObjecttaggingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objecttaggingArgs)(nil)).Elem()
}

type ObjecttaggingInput interface {
	pulumi.Input

	ToObjecttaggingOutput() ObjecttaggingOutput
	ToObjecttaggingOutputWithContext(ctx context.Context) ObjecttaggingOutput
}

func (*Objecttagging) ElementType() reflect.Type {
	return reflect.TypeOf((**Objecttagging)(nil)).Elem()
}

func (i *Objecttagging) ToObjecttaggingOutput() ObjecttaggingOutput {
	return i.ToObjecttaggingOutputWithContext(context.Background())
}

func (i *Objecttagging) ToObjecttaggingOutputWithContext(ctx context.Context) ObjecttaggingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjecttaggingOutput)
}

// ObjecttaggingArrayInput is an input type that accepts ObjecttaggingArray and ObjecttaggingArrayOutput values.
// You can construct a concrete instance of `ObjecttaggingArrayInput` via:
//
//	ObjecttaggingArray{ ObjecttaggingArgs{...} }
type ObjecttaggingArrayInput interface {
	pulumi.Input

	ToObjecttaggingArrayOutput() ObjecttaggingArrayOutput
	ToObjecttaggingArrayOutputWithContext(context.Context) ObjecttaggingArrayOutput
}

type ObjecttaggingArray []ObjecttaggingInput

func (ObjecttaggingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Objecttagging)(nil)).Elem()
}

func (i ObjecttaggingArray) ToObjecttaggingArrayOutput() ObjecttaggingArrayOutput {
	return i.ToObjecttaggingArrayOutputWithContext(context.Background())
}

func (i ObjecttaggingArray) ToObjecttaggingArrayOutputWithContext(ctx context.Context) ObjecttaggingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjecttaggingArrayOutput)
}

// ObjecttaggingMapInput is an input type that accepts ObjecttaggingMap and ObjecttaggingMapOutput values.
// You can construct a concrete instance of `ObjecttaggingMapInput` via:
//
//	ObjecttaggingMap{ "key": ObjecttaggingArgs{...} }
type ObjecttaggingMapInput interface {
	pulumi.Input

	ToObjecttaggingMapOutput() ObjecttaggingMapOutput
	ToObjecttaggingMapOutputWithContext(context.Context) ObjecttaggingMapOutput
}

type ObjecttaggingMap map[string]ObjecttaggingInput

func (ObjecttaggingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Objecttagging)(nil)).Elem()
}

func (i ObjecttaggingMap) ToObjecttaggingMapOutput() ObjecttaggingMapOutput {
	return i.ToObjecttaggingMapOutputWithContext(context.Background())
}

func (i ObjecttaggingMap) ToObjecttaggingMapOutputWithContext(ctx context.Context) ObjecttaggingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjecttaggingMapOutput)
}

type ObjecttaggingOutput struct{ *pulumi.OutputState }

func (ObjecttaggingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Objecttagging)(nil)).Elem()
}

func (o ObjecttaggingOutput) ToObjecttaggingOutput() ObjecttaggingOutput {
	return o
}

func (o ObjecttaggingOutput) ToObjecttaggingOutputWithContext(ctx context.Context) ObjecttaggingOutput {
	return o
}

// Address. Valid values: `disable`, `mandatory`, `optional`.
func (o ObjecttaggingOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Objecttagging) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// Tag Category.
func (o ObjecttaggingOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *Objecttagging) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

// Color of icon on the GUI.
func (o ObjecttaggingOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *Objecttagging) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Device. Valid values: `disable`, `mandatory`, `optional`.
func (o ObjecttaggingOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v *Objecttagging) pulumi.StringOutput { return v.Device }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ObjecttaggingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Objecttagging) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o ObjecttaggingOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Objecttagging) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Interface. Valid values: `disable`, `mandatory`, `optional`.
func (o ObjecttaggingOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Objecttagging) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Allow multiple tag selection. Valid values: `enable`, `disable`.
func (o ObjecttaggingOutput) Multiple() pulumi.StringOutput {
	return o.ApplyT(func(v *Objecttagging) pulumi.StringOutput { return v.Multiple }).(pulumi.StringOutput)
}

// Tags. The structure of `tags` block is documented below.
func (o ObjecttaggingOutput) Tags() ObjecttaggingTagArrayOutput {
	return o.ApplyT(func(v *Objecttagging) ObjecttaggingTagArrayOutput { return v.Tags }).(ObjecttaggingTagArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ObjecttaggingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Objecttagging) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ObjecttaggingArrayOutput struct{ *pulumi.OutputState }

func (ObjecttaggingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Objecttagging)(nil)).Elem()
}

func (o ObjecttaggingArrayOutput) ToObjecttaggingArrayOutput() ObjecttaggingArrayOutput {
	return o
}

func (o ObjecttaggingArrayOutput) ToObjecttaggingArrayOutputWithContext(ctx context.Context) ObjecttaggingArrayOutput {
	return o
}

func (o ObjecttaggingArrayOutput) Index(i pulumi.IntInput) ObjecttaggingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Objecttagging {
		return vs[0].([]*Objecttagging)[vs[1].(int)]
	}).(ObjecttaggingOutput)
}

type ObjecttaggingMapOutput struct{ *pulumi.OutputState }

func (ObjecttaggingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Objecttagging)(nil)).Elem()
}

func (o ObjecttaggingMapOutput) ToObjecttaggingMapOutput() ObjecttaggingMapOutput {
	return o
}

func (o ObjecttaggingMapOutput) ToObjecttaggingMapOutputWithContext(ctx context.Context) ObjecttaggingMapOutput {
	return o
}

func (o ObjecttaggingMapOutput) MapIndex(k pulumi.StringInput) ObjecttaggingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Objecttagging {
		return vs[0].(map[string]*Objecttagging)[vs[1].(string)]
	}).(ObjecttaggingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjecttaggingInput)(nil)).Elem(), &Objecttagging{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjecttaggingArrayInput)(nil)).Elem(), ObjecttaggingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjecttaggingMapInput)(nil)).Elem(), ObjecttaggingMap{})
	pulumi.RegisterOutputType(ObjecttaggingOutput{})
	pulumi.RegisterOutputType(ObjecttaggingArrayOutput{})
	pulumi.RegisterOutputType(ObjecttaggingMapOutput{})
}
