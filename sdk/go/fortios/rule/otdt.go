// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rule

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Show OT detection signatures. Applies to FortiOS Version `>= 7.4.1`.
//
// ## Import
//
// Rule Otdt can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:rule/otdt:Otdt labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:rule/otdt:Otdt labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Otdt struct {
	pulumi.CustomResourceState

	// Application behavior.
	Behavior pulumi.StringOutput `pulumi:"behavior"`
	// Application category ID.
	Category pulumi.IntOutput `pulumi:"category"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Application ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas OtdtMetadataArrayOutput `pulumi:"metadatas"`
	// Application name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Application parameters. The structure of `parameters` block is documented below.
	Parameters OtdtParameterArrayOutput `pulumi:"parameters"`
	// Application popularity.
	Popularity pulumi.IntOutput `pulumi:"popularity"`
	// Application protocol.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Application risk.
	Risk pulumi.IntOutput `pulumi:"risk"`
	// Application technology.
	Technology pulumi.StringOutput `pulumi:"technology"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Application vendor.
	Vendor pulumi.StringOutput `pulumi:"vendor"`
	// Application weight.
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewOtdt registers a new resource with the given unique name, arguments, and options.
func NewOtdt(ctx *pulumi.Context,
	name string, args *OtdtArgs, opts ...pulumi.ResourceOption) (*Otdt, error) {
	if args == nil {
		args = &OtdtArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Otdt
	err := ctx.RegisterResource("fortios:rule/otdt:Otdt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOtdt gets an existing Otdt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOtdt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OtdtState, opts ...pulumi.ResourceOption) (*Otdt, error) {
	var resource Otdt
	err := ctx.ReadResource("fortios:rule/otdt:Otdt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Otdt resources.
type otdtState struct {
	// Application behavior.
	Behavior *string `pulumi:"behavior"`
	// Application category ID.
	Category *int `pulumi:"category"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Application ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas []OtdtMetadata `pulumi:"metadatas"`
	// Application name.
	Name *string `pulumi:"name"`
	// Application parameters. The structure of `parameters` block is documented below.
	Parameters []OtdtParameter `pulumi:"parameters"`
	// Application popularity.
	Popularity *int `pulumi:"popularity"`
	// Application protocol.
	Protocol *string `pulumi:"protocol"`
	// Application risk.
	Risk *int `pulumi:"risk"`
	// Application technology.
	Technology *string `pulumi:"technology"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Application vendor.
	Vendor *string `pulumi:"vendor"`
	// Application weight.
	Weight *int `pulumi:"weight"`
}

type OtdtState struct {
	// Application behavior.
	Behavior pulumi.StringPtrInput
	// Application category ID.
	Category pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Application ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas OtdtMetadataArrayInput
	// Application name.
	Name pulumi.StringPtrInput
	// Application parameters. The structure of `parameters` block is documented below.
	Parameters OtdtParameterArrayInput
	// Application popularity.
	Popularity pulumi.IntPtrInput
	// Application protocol.
	Protocol pulumi.StringPtrInput
	// Application risk.
	Risk pulumi.IntPtrInput
	// Application technology.
	Technology pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Application vendor.
	Vendor pulumi.StringPtrInput
	// Application weight.
	Weight pulumi.IntPtrInput
}

func (OtdtState) ElementType() reflect.Type {
	return reflect.TypeOf((*otdtState)(nil)).Elem()
}

type otdtArgs struct {
	// Application behavior.
	Behavior *string `pulumi:"behavior"`
	// Application category ID.
	Category *int `pulumi:"category"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Application ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas []OtdtMetadata `pulumi:"metadatas"`
	// Application name.
	Name *string `pulumi:"name"`
	// Application parameters. The structure of `parameters` block is documented below.
	Parameters []OtdtParameter `pulumi:"parameters"`
	// Application popularity.
	Popularity *int `pulumi:"popularity"`
	// Application protocol.
	Protocol *string `pulumi:"protocol"`
	// Application risk.
	Risk *int `pulumi:"risk"`
	// Application technology.
	Technology *string `pulumi:"technology"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Application vendor.
	Vendor *string `pulumi:"vendor"`
	// Application weight.
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a Otdt resource.
type OtdtArgs struct {
	// Application behavior.
	Behavior pulumi.StringPtrInput
	// Application category ID.
	Category pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Application ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas OtdtMetadataArrayInput
	// Application name.
	Name pulumi.StringPtrInput
	// Application parameters. The structure of `parameters` block is documented below.
	Parameters OtdtParameterArrayInput
	// Application popularity.
	Popularity pulumi.IntPtrInput
	// Application protocol.
	Protocol pulumi.StringPtrInput
	// Application risk.
	Risk pulumi.IntPtrInput
	// Application technology.
	Technology pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Application vendor.
	Vendor pulumi.StringPtrInput
	// Application weight.
	Weight pulumi.IntPtrInput
}

func (OtdtArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*otdtArgs)(nil)).Elem()
}

type OtdtInput interface {
	pulumi.Input

	ToOtdtOutput() OtdtOutput
	ToOtdtOutputWithContext(ctx context.Context) OtdtOutput
}

func (*Otdt) ElementType() reflect.Type {
	return reflect.TypeOf((**Otdt)(nil)).Elem()
}

func (i *Otdt) ToOtdtOutput() OtdtOutput {
	return i.ToOtdtOutputWithContext(context.Background())
}

func (i *Otdt) ToOtdtOutputWithContext(ctx context.Context) OtdtOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OtdtOutput)
}

// OtdtArrayInput is an input type that accepts OtdtArray and OtdtArrayOutput values.
// You can construct a concrete instance of `OtdtArrayInput` via:
//
//	OtdtArray{ OtdtArgs{...} }
type OtdtArrayInput interface {
	pulumi.Input

	ToOtdtArrayOutput() OtdtArrayOutput
	ToOtdtArrayOutputWithContext(context.Context) OtdtArrayOutput
}

type OtdtArray []OtdtInput

func (OtdtArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Otdt)(nil)).Elem()
}

func (i OtdtArray) ToOtdtArrayOutput() OtdtArrayOutput {
	return i.ToOtdtArrayOutputWithContext(context.Background())
}

func (i OtdtArray) ToOtdtArrayOutputWithContext(ctx context.Context) OtdtArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OtdtArrayOutput)
}

// OtdtMapInput is an input type that accepts OtdtMap and OtdtMapOutput values.
// You can construct a concrete instance of `OtdtMapInput` via:
//
//	OtdtMap{ "key": OtdtArgs{...} }
type OtdtMapInput interface {
	pulumi.Input

	ToOtdtMapOutput() OtdtMapOutput
	ToOtdtMapOutputWithContext(context.Context) OtdtMapOutput
}

type OtdtMap map[string]OtdtInput

func (OtdtMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Otdt)(nil)).Elem()
}

func (i OtdtMap) ToOtdtMapOutput() OtdtMapOutput {
	return i.ToOtdtMapOutputWithContext(context.Background())
}

func (i OtdtMap) ToOtdtMapOutputWithContext(ctx context.Context) OtdtMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OtdtMapOutput)
}

type OtdtOutput struct{ *pulumi.OutputState }

func (OtdtOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Otdt)(nil)).Elem()
}

func (o OtdtOutput) ToOtdtOutput() OtdtOutput {
	return o
}

func (o OtdtOutput) ToOtdtOutputWithContext(ctx context.Context) OtdtOutput {
	return o
}

// Application behavior.
func (o OtdtOutput) Behavior() pulumi.StringOutput {
	return o.ApplyT(func(v *Otdt) pulumi.StringOutput { return v.Behavior }).(pulumi.StringOutput)
}

// Application category ID.
func (o OtdtOutput) Category() pulumi.IntOutput {
	return o.ApplyT(func(v *Otdt) pulumi.IntOutput { return v.Category }).(pulumi.IntOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o OtdtOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Otdt) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Application ID.
func (o OtdtOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Otdt) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o OtdtOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Otdt) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Meta data. The structure of `metadata` block is documented below.
func (o OtdtOutput) Metadatas() OtdtMetadataArrayOutput {
	return o.ApplyT(func(v *Otdt) OtdtMetadataArrayOutput { return v.Metadatas }).(OtdtMetadataArrayOutput)
}

// Application name.
func (o OtdtOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Otdt) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Application parameters. The structure of `parameters` block is documented below.
func (o OtdtOutput) Parameters() OtdtParameterArrayOutput {
	return o.ApplyT(func(v *Otdt) OtdtParameterArrayOutput { return v.Parameters }).(OtdtParameterArrayOutput)
}

// Application popularity.
func (o OtdtOutput) Popularity() pulumi.IntOutput {
	return o.ApplyT(func(v *Otdt) pulumi.IntOutput { return v.Popularity }).(pulumi.IntOutput)
}

// Application protocol.
func (o OtdtOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Otdt) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Application risk.
func (o OtdtOutput) Risk() pulumi.IntOutput {
	return o.ApplyT(func(v *Otdt) pulumi.IntOutput { return v.Risk }).(pulumi.IntOutput)
}

// Application technology.
func (o OtdtOutput) Technology() pulumi.StringOutput {
	return o.ApplyT(func(v *Otdt) pulumi.StringOutput { return v.Technology }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o OtdtOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Otdt) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Application vendor.
func (o OtdtOutput) Vendor() pulumi.StringOutput {
	return o.ApplyT(func(v *Otdt) pulumi.StringOutput { return v.Vendor }).(pulumi.StringOutput)
}

// Application weight.
func (o OtdtOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v *Otdt) pulumi.IntOutput { return v.Weight }).(pulumi.IntOutput)
}

type OtdtArrayOutput struct{ *pulumi.OutputState }

func (OtdtArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Otdt)(nil)).Elem()
}

func (o OtdtArrayOutput) ToOtdtArrayOutput() OtdtArrayOutput {
	return o
}

func (o OtdtArrayOutput) ToOtdtArrayOutputWithContext(ctx context.Context) OtdtArrayOutput {
	return o
}

func (o OtdtArrayOutput) Index(i pulumi.IntInput) OtdtOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Otdt {
		return vs[0].([]*Otdt)[vs[1].(int)]
	}).(OtdtOutput)
}

type OtdtMapOutput struct{ *pulumi.OutputState }

func (OtdtMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Otdt)(nil)).Elem()
}

func (o OtdtMapOutput) ToOtdtMapOutput() OtdtMapOutput {
	return o
}

func (o OtdtMapOutput) ToOtdtMapOutputWithContext(ctx context.Context) OtdtMapOutput {
	return o
}

func (o OtdtMapOutput) MapIndex(k pulumi.StringInput) OtdtOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Otdt {
		return vs[0].(map[string]*Otdt)[vs[1].(string)]
	}).(OtdtOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OtdtInput)(nil)).Elem(), &Otdt{})
	pulumi.RegisterInputType(reflect.TypeOf((*OtdtArrayInput)(nil)).Elem(), OtdtArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OtdtMapInput)(nil)).Elem(), OtdtMap{})
	pulumi.RegisterOutputType(OtdtOutput{})
	pulumi.RegisterOutputType(OtdtArrayOutput{})
	pulumi.RegisterOutputType(OtdtMapOutput{})
}
