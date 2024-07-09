// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package application

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure application signatures.
//
// ## Import
//
// Application Name can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:application/name:Name labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:application/name:Name labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Name struct {
	pulumi.CustomResourceState

	// Application behavior.
	Behavior pulumi.StringOutput `pulumi:"behavior"`
	// Application category ID.
	Category pulumi.IntOutput `pulumi:"category"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Application ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas NameMetadataArrayOutput `pulumi:"metadatas"`
	// Application name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Application parameter name.
	Parameter pulumi.StringOutput `pulumi:"parameter"`
	// Application parameters. The structure of `parameters` block is documented below.
	Parameters NameParameterArrayOutput `pulumi:"parameters"`
	// Application popularity.
	Popularity pulumi.IntOutput `pulumi:"popularity"`
	// Application protocol.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Application risk.
	Risk pulumi.IntOutput `pulumi:"risk"`
	// Application sub-category ID.
	SubCategory pulumi.IntOutput `pulumi:"subCategory"`
	// Application technology.
	Technology pulumi.StringOutput `pulumi:"technology"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
	// Application vendor.
	Vendor pulumi.StringOutput `pulumi:"vendor"`
	// Application weight.
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewName registers a new resource with the given unique name, arguments, and options.
func NewName(ctx *pulumi.Context,
	name string, args *NameArgs, opts ...pulumi.ResourceOption) (*Name, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Name
	err := ctx.RegisterResource("fortios:application/name:Name", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetName gets an existing Name resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NameState, opts ...pulumi.ResourceOption) (*Name, error) {
	var resource Name
	err := ctx.ReadResource("fortios:application/name:Name", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Name resources.
type nameState struct {
	// Application behavior.
	Behavior *string `pulumi:"behavior"`
	// Application category ID.
	Category *int `pulumi:"category"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Application ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas []NameMetadata `pulumi:"metadatas"`
	// Application name.
	Name *string `pulumi:"name"`
	// Application parameter name.
	Parameter *string `pulumi:"parameter"`
	// Application parameters. The structure of `parameters` block is documented below.
	Parameters []NameParameter `pulumi:"parameters"`
	// Application popularity.
	Popularity *int `pulumi:"popularity"`
	// Application protocol.
	Protocol *string `pulumi:"protocol"`
	// Application risk.
	Risk *int `pulumi:"risk"`
	// Application sub-category ID.
	SubCategory *int `pulumi:"subCategory"`
	// Application technology.
	Technology *string `pulumi:"technology"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Application vendor.
	Vendor *string `pulumi:"vendor"`
	// Application weight.
	Weight *int `pulumi:"weight"`
}

type NameState struct {
	// Application behavior.
	Behavior pulumi.StringPtrInput
	// Application category ID.
	Category pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Application ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas NameMetadataArrayInput
	// Application name.
	Name pulumi.StringPtrInput
	// Application parameter name.
	Parameter pulumi.StringPtrInput
	// Application parameters. The structure of `parameters` block is documented below.
	Parameters NameParameterArrayInput
	// Application popularity.
	Popularity pulumi.IntPtrInput
	// Application protocol.
	Protocol pulumi.StringPtrInput
	// Application risk.
	Risk pulumi.IntPtrInput
	// Application sub-category ID.
	SubCategory pulumi.IntPtrInput
	// Application technology.
	Technology pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Application vendor.
	Vendor pulumi.StringPtrInput
	// Application weight.
	Weight pulumi.IntPtrInput
}

func (NameState) ElementType() reflect.Type {
	return reflect.TypeOf((*nameState)(nil)).Elem()
}

type nameArgs struct {
	// Application behavior.
	Behavior *string `pulumi:"behavior"`
	// Application category ID.
	Category int `pulumi:"category"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Application ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas []NameMetadata `pulumi:"metadatas"`
	// Application name.
	Name *string `pulumi:"name"`
	// Application parameter name.
	Parameter *string `pulumi:"parameter"`
	// Application parameters. The structure of `parameters` block is documented below.
	Parameters []NameParameter `pulumi:"parameters"`
	// Application popularity.
	Popularity *int `pulumi:"popularity"`
	// Application protocol.
	Protocol *string `pulumi:"protocol"`
	// Application risk.
	Risk *int `pulumi:"risk"`
	// Application sub-category ID.
	SubCategory *int `pulumi:"subCategory"`
	// Application technology.
	Technology *string `pulumi:"technology"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Application vendor.
	Vendor *string `pulumi:"vendor"`
	// Application weight.
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a Name resource.
type NameArgs struct {
	// Application behavior.
	Behavior pulumi.StringPtrInput
	// Application category ID.
	Category pulumi.IntInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Application ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas NameMetadataArrayInput
	// Application name.
	Name pulumi.StringPtrInput
	// Application parameter name.
	Parameter pulumi.StringPtrInput
	// Application parameters. The structure of `parameters` block is documented below.
	Parameters NameParameterArrayInput
	// Application popularity.
	Popularity pulumi.IntPtrInput
	// Application protocol.
	Protocol pulumi.StringPtrInput
	// Application risk.
	Risk pulumi.IntPtrInput
	// Application sub-category ID.
	SubCategory pulumi.IntPtrInput
	// Application technology.
	Technology pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Application vendor.
	Vendor pulumi.StringPtrInput
	// Application weight.
	Weight pulumi.IntPtrInput
}

func (NameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nameArgs)(nil)).Elem()
}

type NameInput interface {
	pulumi.Input

	ToNameOutput() NameOutput
	ToNameOutputWithContext(ctx context.Context) NameOutput
}

func (*Name) ElementType() reflect.Type {
	return reflect.TypeOf((**Name)(nil)).Elem()
}

func (i *Name) ToNameOutput() NameOutput {
	return i.ToNameOutputWithContext(context.Background())
}

func (i *Name) ToNameOutputWithContext(ctx context.Context) NameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameOutput)
}

// NameArrayInput is an input type that accepts NameArray and NameArrayOutput values.
// You can construct a concrete instance of `NameArrayInput` via:
//
//	NameArray{ NameArgs{...} }
type NameArrayInput interface {
	pulumi.Input

	ToNameArrayOutput() NameArrayOutput
	ToNameArrayOutputWithContext(context.Context) NameArrayOutput
}

type NameArray []NameInput

func (NameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Name)(nil)).Elem()
}

func (i NameArray) ToNameArrayOutput() NameArrayOutput {
	return i.ToNameArrayOutputWithContext(context.Background())
}

func (i NameArray) ToNameArrayOutputWithContext(ctx context.Context) NameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameArrayOutput)
}

// NameMapInput is an input type that accepts NameMap and NameMapOutput values.
// You can construct a concrete instance of `NameMapInput` via:
//
//	NameMap{ "key": NameArgs{...} }
type NameMapInput interface {
	pulumi.Input

	ToNameMapOutput() NameMapOutput
	ToNameMapOutputWithContext(context.Context) NameMapOutput
}

type NameMap map[string]NameInput

func (NameMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Name)(nil)).Elem()
}

func (i NameMap) ToNameMapOutput() NameMapOutput {
	return i.ToNameMapOutputWithContext(context.Background())
}

func (i NameMap) ToNameMapOutputWithContext(ctx context.Context) NameMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameMapOutput)
}

type NameOutput struct{ *pulumi.OutputState }

func (NameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Name)(nil)).Elem()
}

func (o NameOutput) ToNameOutput() NameOutput {
	return o
}

func (o NameOutput) ToNameOutputWithContext(ctx context.Context) NameOutput {
	return o
}

// Application behavior.
func (o NameOutput) Behavior() pulumi.StringOutput {
	return o.ApplyT(func(v *Name) pulumi.StringOutput { return v.Behavior }).(pulumi.StringOutput)
}

// Application category ID.
func (o NameOutput) Category() pulumi.IntOutput {
	return o.ApplyT(func(v *Name) pulumi.IntOutput { return v.Category }).(pulumi.IntOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o NameOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Name) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Application ID.
func (o NameOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Name) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o NameOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Name) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Meta data. The structure of `metadata` block is documented below.
func (o NameOutput) Metadatas() NameMetadataArrayOutput {
	return o.ApplyT(func(v *Name) NameMetadataArrayOutput { return v.Metadatas }).(NameMetadataArrayOutput)
}

// Application name.
func (o NameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Name) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Application parameter name.
func (o NameOutput) Parameter() pulumi.StringOutput {
	return o.ApplyT(func(v *Name) pulumi.StringOutput { return v.Parameter }).(pulumi.StringOutput)
}

// Application parameters. The structure of `parameters` block is documented below.
func (o NameOutput) Parameters() NameParameterArrayOutput {
	return o.ApplyT(func(v *Name) NameParameterArrayOutput { return v.Parameters }).(NameParameterArrayOutput)
}

// Application popularity.
func (o NameOutput) Popularity() pulumi.IntOutput {
	return o.ApplyT(func(v *Name) pulumi.IntOutput { return v.Popularity }).(pulumi.IntOutput)
}

// Application protocol.
func (o NameOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Name) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Application risk.
func (o NameOutput) Risk() pulumi.IntOutput {
	return o.ApplyT(func(v *Name) pulumi.IntOutput { return v.Risk }).(pulumi.IntOutput)
}

// Application sub-category ID.
func (o NameOutput) SubCategory() pulumi.IntOutput {
	return o.ApplyT(func(v *Name) pulumi.IntOutput { return v.SubCategory }).(pulumi.IntOutput)
}

// Application technology.
func (o NameOutput) Technology() pulumi.StringOutput {
	return o.ApplyT(func(v *Name) pulumi.StringOutput { return v.Technology }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o NameOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Name) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

// Application vendor.
func (o NameOutput) Vendor() pulumi.StringOutput {
	return o.ApplyT(func(v *Name) pulumi.StringOutput { return v.Vendor }).(pulumi.StringOutput)
}

// Application weight.
func (o NameOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v *Name) pulumi.IntOutput { return v.Weight }).(pulumi.IntOutput)
}

type NameArrayOutput struct{ *pulumi.OutputState }

func (NameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Name)(nil)).Elem()
}

func (o NameArrayOutput) ToNameArrayOutput() NameArrayOutput {
	return o
}

func (o NameArrayOutput) ToNameArrayOutputWithContext(ctx context.Context) NameArrayOutput {
	return o
}

func (o NameArrayOutput) Index(i pulumi.IntInput) NameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Name {
		return vs[0].([]*Name)[vs[1].(int)]
	}).(NameOutput)
}

type NameMapOutput struct{ *pulumi.OutputState }

func (NameMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Name)(nil)).Elem()
}

func (o NameMapOutput) ToNameMapOutput() NameMapOutput {
	return o
}

func (o NameMapOutput) ToNameMapOutputWithContext(ctx context.Context) NameMapOutput {
	return o
}

func (o NameMapOutput) MapIndex(k pulumi.StringInput) NameOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Name {
		return vs[0].(map[string]*Name)[vs[1].(string)]
	}).(NameOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NameInput)(nil)).Elem(), &Name{})
	pulumi.RegisterInputType(reflect.TypeOf((*NameArrayInput)(nil)).Elem(), NameArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NameMapInput)(nil)).Elem(), NameMap{})
	pulumi.RegisterOutputType(NameOutput{})
	pulumi.RegisterOutputType(NameArrayOutput{})
	pulumi.RegisterOutputType(NameMapOutput{})
}
