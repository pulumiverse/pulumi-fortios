// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package report

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Report dataset configuration. Applies to FortiOS Version `<= 7.0.0`.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/report"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := report.NewDataset(ctx, "trname", &report.DatasetArgs{
//				Policy: pulumi.Int(0),
//				Query:  pulumi.String("select * from testdb"),
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
// Report Dataset can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:report/dataset:Dataset labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:report/dataset:Dataset labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Dataset struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Fields. The structure of `field` block is documented below.
	Fields DatasetFieldArrayOutput `pulumi:"fields"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parameters. The structure of `parameters` block is documented below.
	Parameters DatasetParameterArrayOutput `pulumi:"parameters"`
	// Used by monitor policy.
	Policy pulumi.IntOutput `pulumi:"policy"`
	// SQL query statement.
	Query pulumi.StringOutput `pulumi:"query"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDataset registers a new resource with the given unique name, arguments, and options.
func NewDataset(ctx *pulumi.Context,
	name string, args *DatasetArgs, opts ...pulumi.ResourceOption) (*Dataset, error) {
	if args == nil {
		args = &DatasetArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Dataset
	err := ctx.RegisterResource("fortios:report/dataset:Dataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataset gets an existing Dataset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetState, opts ...pulumi.ResourceOption) (*Dataset, error) {
	var resource Dataset
	err := ctx.ReadResource("fortios:report/dataset:Dataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dataset resources.
type datasetState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Fields. The structure of `field` block is documented below.
	Fields []DatasetField `pulumi:"fields"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Name.
	Name *string `pulumi:"name"`
	// Parameters. The structure of `parameters` block is documented below.
	Parameters []DatasetParameter `pulumi:"parameters"`
	// Used by monitor policy.
	Policy *int `pulumi:"policy"`
	// SQL query statement.
	Query *string `pulumi:"query"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DatasetState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Fields. The structure of `field` block is documented below.
	Fields DatasetFieldArrayInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Parameters. The structure of `parameters` block is documented below.
	Parameters DatasetParameterArrayInput
	// Used by monitor policy.
	Policy pulumi.IntPtrInput
	// SQL query statement.
	Query pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetState)(nil)).Elem()
}

type datasetArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Fields. The structure of `field` block is documented below.
	Fields []DatasetField `pulumi:"fields"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Name.
	Name *string `pulumi:"name"`
	// Parameters. The structure of `parameters` block is documented below.
	Parameters []DatasetParameter `pulumi:"parameters"`
	// Used by monitor policy.
	Policy *int `pulumi:"policy"`
	// SQL query statement.
	Query *string `pulumi:"query"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Dataset resource.
type DatasetArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Fields. The structure of `field` block is documented below.
	Fields DatasetFieldArrayInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Parameters. The structure of `parameters` block is documented below.
	Parameters DatasetParameterArrayInput
	// Used by monitor policy.
	Policy pulumi.IntPtrInput
	// SQL query statement.
	Query pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetArgs)(nil)).Elem()
}

type DatasetInput interface {
	pulumi.Input

	ToDatasetOutput() DatasetOutput
	ToDatasetOutputWithContext(ctx context.Context) DatasetOutput
}

func (*Dataset) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (i *Dataset) ToDatasetOutput() DatasetOutput {
	return i.ToDatasetOutputWithContext(context.Background())
}

func (i *Dataset) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetOutput)
}

// DatasetArrayInput is an input type that accepts DatasetArray and DatasetArrayOutput values.
// You can construct a concrete instance of `DatasetArrayInput` via:
//
//	DatasetArray{ DatasetArgs{...} }
type DatasetArrayInput interface {
	pulumi.Input

	ToDatasetArrayOutput() DatasetArrayOutput
	ToDatasetArrayOutputWithContext(context.Context) DatasetArrayOutput
}

type DatasetArray []DatasetInput

func (DatasetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dataset)(nil)).Elem()
}

func (i DatasetArray) ToDatasetArrayOutput() DatasetArrayOutput {
	return i.ToDatasetArrayOutputWithContext(context.Background())
}

func (i DatasetArray) ToDatasetArrayOutputWithContext(ctx context.Context) DatasetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetArrayOutput)
}

// DatasetMapInput is an input type that accepts DatasetMap and DatasetMapOutput values.
// You can construct a concrete instance of `DatasetMapInput` via:
//
//	DatasetMap{ "key": DatasetArgs{...} }
type DatasetMapInput interface {
	pulumi.Input

	ToDatasetMapOutput() DatasetMapOutput
	ToDatasetMapOutputWithContext(context.Context) DatasetMapOutput
}

type DatasetMap map[string]DatasetInput

func (DatasetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dataset)(nil)).Elem()
}

func (i DatasetMap) ToDatasetMapOutput() DatasetMapOutput {
	return i.ToDatasetMapOutputWithContext(context.Background())
}

func (i DatasetMap) ToDatasetMapOutputWithContext(ctx context.Context) DatasetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetMapOutput)
}

type DatasetOutput struct{ *pulumi.OutputState }

func (DatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (o DatasetOutput) ToDatasetOutput() DatasetOutput {
	return o
}

func (o DatasetOutput) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o DatasetOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Fields. The structure of `field` block is documented below.
func (o DatasetOutput) Fields() DatasetFieldArrayOutput {
	return o.ApplyT(func(v *Dataset) DatasetFieldArrayOutput { return v.Fields }).(DatasetFieldArrayOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o DatasetOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Name.
func (o DatasetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Parameters. The structure of `parameters` block is documented below.
func (o DatasetOutput) Parameters() DatasetParameterArrayOutput {
	return o.ApplyT(func(v *Dataset) DatasetParameterArrayOutput { return v.Parameters }).(DatasetParameterArrayOutput)
}

// Used by monitor policy.
func (o DatasetOutput) Policy() pulumi.IntOutput {
	return o.ApplyT(func(v *Dataset) pulumi.IntOutput { return v.Policy }).(pulumi.IntOutput)
}

// SQL query statement.
func (o DatasetOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DatasetOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DatasetArrayOutput struct{ *pulumi.OutputState }

func (DatasetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dataset)(nil)).Elem()
}

func (o DatasetArrayOutput) ToDatasetArrayOutput() DatasetArrayOutput {
	return o
}

func (o DatasetArrayOutput) ToDatasetArrayOutputWithContext(ctx context.Context) DatasetArrayOutput {
	return o
}

func (o DatasetArrayOutput) Index(i pulumi.IntInput) DatasetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dataset {
		return vs[0].([]*Dataset)[vs[1].(int)]
	}).(DatasetOutput)
}

type DatasetMapOutput struct{ *pulumi.OutputState }

func (DatasetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dataset)(nil)).Elem()
}

func (o DatasetMapOutput) ToDatasetMapOutput() DatasetMapOutput {
	return o
}

func (o DatasetMapOutput) ToDatasetMapOutputWithContext(ctx context.Context) DatasetMapOutput {
	return o
}

func (o DatasetMapOutput) MapIndex(k pulumi.StringInput) DatasetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dataset {
		return vs[0].(map[string]*Dataset)[vs[1].(string)]
	}).(DatasetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetInput)(nil)).Elem(), &Dataset{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetArrayInput)(nil)).Elem(), DatasetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetMapInput)(nil)).Elem(), DatasetMap{})
	pulumi.RegisterOutputType(DatasetOutput{})
	pulumi.RegisterOutputType(DatasetArrayOutput{})
	pulumi.RegisterOutputType(DatasetMapOutput{})
}
