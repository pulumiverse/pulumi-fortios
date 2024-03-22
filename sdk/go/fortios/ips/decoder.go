// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ips

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPS decoder.
//
// ## Import
//
// Ips Decoder can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:ips/decoder:Decoder labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:ips/decoder:Decoder labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Decoder struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Decoder name.
	Name pulumi.StringOutput `pulumi:"name"`
	// IPS group parameters. The structure of `parameter` block is documented below.
	Parameters DecoderParameterArrayOutput `pulumi:"parameters"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDecoder registers a new resource with the given unique name, arguments, and options.
func NewDecoder(ctx *pulumi.Context,
	name string, args *DecoderArgs, opts ...pulumi.ResourceOption) (*Decoder, error) {
	if args == nil {
		args = &DecoderArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Decoder
	err := ctx.RegisterResource("fortios:ips/decoder:Decoder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDecoder gets an existing Decoder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDecoder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DecoderState, opts ...pulumi.ResourceOption) (*Decoder, error) {
	var resource Decoder
	err := ctx.ReadResource("fortios:ips/decoder:Decoder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Decoder resources.
type decoderState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Decoder name.
	Name *string `pulumi:"name"`
	// IPS group parameters. The structure of `parameter` block is documented below.
	Parameters []DecoderParameter `pulumi:"parameters"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DecoderState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Decoder name.
	Name pulumi.StringPtrInput
	// IPS group parameters. The structure of `parameter` block is documented below.
	Parameters DecoderParameterArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DecoderState) ElementType() reflect.Type {
	return reflect.TypeOf((*decoderState)(nil)).Elem()
}

type decoderArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Decoder name.
	Name *string `pulumi:"name"`
	// IPS group parameters. The structure of `parameter` block is documented below.
	Parameters []DecoderParameter `pulumi:"parameters"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Decoder resource.
type DecoderArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Decoder name.
	Name pulumi.StringPtrInput
	// IPS group parameters. The structure of `parameter` block is documented below.
	Parameters DecoderParameterArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DecoderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*decoderArgs)(nil)).Elem()
}

type DecoderInput interface {
	pulumi.Input

	ToDecoderOutput() DecoderOutput
	ToDecoderOutputWithContext(ctx context.Context) DecoderOutput
}

func (*Decoder) ElementType() reflect.Type {
	return reflect.TypeOf((**Decoder)(nil)).Elem()
}

func (i *Decoder) ToDecoderOutput() DecoderOutput {
	return i.ToDecoderOutputWithContext(context.Background())
}

func (i *Decoder) ToDecoderOutputWithContext(ctx context.Context) DecoderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DecoderOutput)
}

// DecoderArrayInput is an input type that accepts DecoderArray and DecoderArrayOutput values.
// You can construct a concrete instance of `DecoderArrayInput` via:
//
//	DecoderArray{ DecoderArgs{...} }
type DecoderArrayInput interface {
	pulumi.Input

	ToDecoderArrayOutput() DecoderArrayOutput
	ToDecoderArrayOutputWithContext(context.Context) DecoderArrayOutput
}

type DecoderArray []DecoderInput

func (DecoderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Decoder)(nil)).Elem()
}

func (i DecoderArray) ToDecoderArrayOutput() DecoderArrayOutput {
	return i.ToDecoderArrayOutputWithContext(context.Background())
}

func (i DecoderArray) ToDecoderArrayOutputWithContext(ctx context.Context) DecoderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DecoderArrayOutput)
}

// DecoderMapInput is an input type that accepts DecoderMap and DecoderMapOutput values.
// You can construct a concrete instance of `DecoderMapInput` via:
//
//	DecoderMap{ "key": DecoderArgs{...} }
type DecoderMapInput interface {
	pulumi.Input

	ToDecoderMapOutput() DecoderMapOutput
	ToDecoderMapOutputWithContext(context.Context) DecoderMapOutput
}

type DecoderMap map[string]DecoderInput

func (DecoderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Decoder)(nil)).Elem()
}

func (i DecoderMap) ToDecoderMapOutput() DecoderMapOutput {
	return i.ToDecoderMapOutputWithContext(context.Background())
}

func (i DecoderMap) ToDecoderMapOutputWithContext(ctx context.Context) DecoderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DecoderMapOutput)
}

type DecoderOutput struct{ *pulumi.OutputState }

func (DecoderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Decoder)(nil)).Elem()
}

func (o DecoderOutput) ToDecoderOutput() DecoderOutput {
	return o
}

func (o DecoderOutput) ToDecoderOutputWithContext(ctx context.Context) DecoderOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o DecoderOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Decoder) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o DecoderOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Decoder) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Decoder name.
func (o DecoderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Decoder) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// IPS group parameters. The structure of `parameter` block is documented below.
func (o DecoderOutput) Parameters() DecoderParameterArrayOutput {
	return o.ApplyT(func(v *Decoder) DecoderParameterArrayOutput { return v.Parameters }).(DecoderParameterArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DecoderOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Decoder) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DecoderArrayOutput struct{ *pulumi.OutputState }

func (DecoderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Decoder)(nil)).Elem()
}

func (o DecoderArrayOutput) ToDecoderArrayOutput() DecoderArrayOutput {
	return o
}

func (o DecoderArrayOutput) ToDecoderArrayOutputWithContext(ctx context.Context) DecoderArrayOutput {
	return o
}

func (o DecoderArrayOutput) Index(i pulumi.IntInput) DecoderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Decoder {
		return vs[0].([]*Decoder)[vs[1].(int)]
	}).(DecoderOutput)
}

type DecoderMapOutput struct{ *pulumi.OutputState }

func (DecoderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Decoder)(nil)).Elem()
}

func (o DecoderMapOutput) ToDecoderMapOutput() DecoderMapOutput {
	return o
}

func (o DecoderMapOutput) ToDecoderMapOutputWithContext(ctx context.Context) DecoderMapOutput {
	return o
}

func (o DecoderMapOutput) MapIndex(k pulumi.StringInput) DecoderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Decoder {
		return vs[0].(map[string]*Decoder)[vs[1].(string)]
	}).(DecoderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DecoderInput)(nil)).Elem(), &Decoder{})
	pulumi.RegisterInputType(reflect.TypeOf((*DecoderArrayInput)(nil)).Elem(), DecoderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DecoderMapInput)(nil)).Elem(), DecoderMap{})
	pulumi.RegisterOutputType(DecoderOutput{})
	pulumi.RegisterOutputType(DecoderArrayOutput{})
	pulumi.RegisterOutputType(DecoderMapOutput{})
}
