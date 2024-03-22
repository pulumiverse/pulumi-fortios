// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dlp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure exact-data-match template used by DLP scan. Applies to FortiOS Version `>= 7.4.2`.
//
// ## Import
//
// Dlp ExactDataMatch can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:dlp/exactdatamatch:Exactdatamatch labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:dlp/exactdatamatch:Exactdatamatch labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Exactdatamatch struct {
	pulumi.CustomResourceState

	// DLP exact-data-match column types. The structure of `columns` block is documented below.
	Columns ExactdatamatchColumnArrayOutput `pulumi:"columns"`
	// External resource for exact data match.
	Data pulumi.StringOutput `pulumi:"data"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Name of table containing the exact-data-match template.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of optional columns need to match.
	Optional pulumi.IntOutput `pulumi:"optional"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewExactdatamatch registers a new resource with the given unique name, arguments, and options.
func NewExactdatamatch(ctx *pulumi.Context,
	name string, args *ExactdatamatchArgs, opts ...pulumi.ResourceOption) (*Exactdatamatch, error) {
	if args == nil {
		args = &ExactdatamatchArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Exactdatamatch
	err := ctx.RegisterResource("fortios:dlp/exactdatamatch:Exactdatamatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExactdatamatch gets an existing Exactdatamatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExactdatamatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExactdatamatchState, opts ...pulumi.ResourceOption) (*Exactdatamatch, error) {
	var resource Exactdatamatch
	err := ctx.ReadResource("fortios:dlp/exactdatamatch:Exactdatamatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Exactdatamatch resources.
type exactdatamatchState struct {
	// DLP exact-data-match column types. The structure of `columns` block is documented below.
	Columns []ExactdatamatchColumn `pulumi:"columns"`
	// External resource for exact data match.
	Data *string `pulumi:"data"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Name of table containing the exact-data-match template.
	Name *string `pulumi:"name"`
	// Number of optional columns need to match.
	Optional *int `pulumi:"optional"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ExactdatamatchState struct {
	// DLP exact-data-match column types. The structure of `columns` block is documented below.
	Columns ExactdatamatchColumnArrayInput
	// External resource for exact data match.
	Data pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Name of table containing the exact-data-match template.
	Name pulumi.StringPtrInput
	// Number of optional columns need to match.
	Optional pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExactdatamatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*exactdatamatchState)(nil)).Elem()
}

type exactdatamatchArgs struct {
	// DLP exact-data-match column types. The structure of `columns` block is documented below.
	Columns []ExactdatamatchColumn `pulumi:"columns"`
	// External resource for exact data match.
	Data *string `pulumi:"data"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Name of table containing the exact-data-match template.
	Name *string `pulumi:"name"`
	// Number of optional columns need to match.
	Optional *int `pulumi:"optional"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Exactdatamatch resource.
type ExactdatamatchArgs struct {
	// DLP exact-data-match column types. The structure of `columns` block is documented below.
	Columns ExactdatamatchColumnArrayInput
	// External resource for exact data match.
	Data pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Name of table containing the exact-data-match template.
	Name pulumi.StringPtrInput
	// Number of optional columns need to match.
	Optional pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExactdatamatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*exactdatamatchArgs)(nil)).Elem()
}

type ExactdatamatchInput interface {
	pulumi.Input

	ToExactdatamatchOutput() ExactdatamatchOutput
	ToExactdatamatchOutputWithContext(ctx context.Context) ExactdatamatchOutput
}

func (*Exactdatamatch) ElementType() reflect.Type {
	return reflect.TypeOf((**Exactdatamatch)(nil)).Elem()
}

func (i *Exactdatamatch) ToExactdatamatchOutput() ExactdatamatchOutput {
	return i.ToExactdatamatchOutputWithContext(context.Background())
}

func (i *Exactdatamatch) ToExactdatamatchOutputWithContext(ctx context.Context) ExactdatamatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExactdatamatchOutput)
}

// ExactdatamatchArrayInput is an input type that accepts ExactdatamatchArray and ExactdatamatchArrayOutput values.
// You can construct a concrete instance of `ExactdatamatchArrayInput` via:
//
//	ExactdatamatchArray{ ExactdatamatchArgs{...} }
type ExactdatamatchArrayInput interface {
	pulumi.Input

	ToExactdatamatchArrayOutput() ExactdatamatchArrayOutput
	ToExactdatamatchArrayOutputWithContext(context.Context) ExactdatamatchArrayOutput
}

type ExactdatamatchArray []ExactdatamatchInput

func (ExactdatamatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Exactdatamatch)(nil)).Elem()
}

func (i ExactdatamatchArray) ToExactdatamatchArrayOutput() ExactdatamatchArrayOutput {
	return i.ToExactdatamatchArrayOutputWithContext(context.Background())
}

func (i ExactdatamatchArray) ToExactdatamatchArrayOutputWithContext(ctx context.Context) ExactdatamatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExactdatamatchArrayOutput)
}

// ExactdatamatchMapInput is an input type that accepts ExactdatamatchMap and ExactdatamatchMapOutput values.
// You can construct a concrete instance of `ExactdatamatchMapInput` via:
//
//	ExactdatamatchMap{ "key": ExactdatamatchArgs{...} }
type ExactdatamatchMapInput interface {
	pulumi.Input

	ToExactdatamatchMapOutput() ExactdatamatchMapOutput
	ToExactdatamatchMapOutputWithContext(context.Context) ExactdatamatchMapOutput
}

type ExactdatamatchMap map[string]ExactdatamatchInput

func (ExactdatamatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Exactdatamatch)(nil)).Elem()
}

func (i ExactdatamatchMap) ToExactdatamatchMapOutput() ExactdatamatchMapOutput {
	return i.ToExactdatamatchMapOutputWithContext(context.Background())
}

func (i ExactdatamatchMap) ToExactdatamatchMapOutputWithContext(ctx context.Context) ExactdatamatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExactdatamatchMapOutput)
}

type ExactdatamatchOutput struct{ *pulumi.OutputState }

func (ExactdatamatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Exactdatamatch)(nil)).Elem()
}

func (o ExactdatamatchOutput) ToExactdatamatchOutput() ExactdatamatchOutput {
	return o
}

func (o ExactdatamatchOutput) ToExactdatamatchOutputWithContext(ctx context.Context) ExactdatamatchOutput {
	return o
}

// DLP exact-data-match column types. The structure of `columns` block is documented below.
func (o ExactdatamatchOutput) Columns() ExactdatamatchColumnArrayOutput {
	return o.ApplyT(func(v *Exactdatamatch) ExactdatamatchColumnArrayOutput { return v.Columns }).(ExactdatamatchColumnArrayOutput)
}

// External resource for exact data match.
func (o ExactdatamatchOutput) Data() pulumi.StringOutput {
	return o.ApplyT(func(v *Exactdatamatch) pulumi.StringOutput { return v.Data }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ExactdatamatchOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Exactdatamatch) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o ExactdatamatchOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Exactdatamatch) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Name of table containing the exact-data-match template.
func (o ExactdatamatchOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Exactdatamatch) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of optional columns need to match.
func (o ExactdatamatchOutput) Optional() pulumi.IntOutput {
	return o.ApplyT(func(v *Exactdatamatch) pulumi.IntOutput { return v.Optional }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ExactdatamatchOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Exactdatamatch) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ExactdatamatchArrayOutput struct{ *pulumi.OutputState }

func (ExactdatamatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Exactdatamatch)(nil)).Elem()
}

func (o ExactdatamatchArrayOutput) ToExactdatamatchArrayOutput() ExactdatamatchArrayOutput {
	return o
}

func (o ExactdatamatchArrayOutput) ToExactdatamatchArrayOutputWithContext(ctx context.Context) ExactdatamatchArrayOutput {
	return o
}

func (o ExactdatamatchArrayOutput) Index(i pulumi.IntInput) ExactdatamatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Exactdatamatch {
		return vs[0].([]*Exactdatamatch)[vs[1].(int)]
	}).(ExactdatamatchOutput)
}

type ExactdatamatchMapOutput struct{ *pulumi.OutputState }

func (ExactdatamatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Exactdatamatch)(nil)).Elem()
}

func (o ExactdatamatchMapOutput) ToExactdatamatchMapOutput() ExactdatamatchMapOutput {
	return o
}

func (o ExactdatamatchMapOutput) ToExactdatamatchMapOutputWithContext(ctx context.Context) ExactdatamatchMapOutput {
	return o
}

func (o ExactdatamatchMapOutput) MapIndex(k pulumi.StringInput) ExactdatamatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Exactdatamatch {
		return vs[0].(map[string]*Exactdatamatch)[vs[1].(string)]
	}).(ExactdatamatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExactdatamatchInput)(nil)).Elem(), &Exactdatamatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExactdatamatchArrayInput)(nil)).Elem(), ExactdatamatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExactdatamatchMapInput)(nil)).Elem(), ExactdatamatchMap{})
	pulumi.RegisterOutputType(ExactdatamatchOutput{})
	pulumi.RegisterOutputType(ExactdatamatchArrayOutput{})
	pulumi.RegisterOutputType(ExactdatamatchMapOutput{})
}
