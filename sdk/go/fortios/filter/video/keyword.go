// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package video

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure video filter keywords. Applies to FortiOS Version `>= 7.4.2`.
//
// ## Import
//
// Videofilter Keyword can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:filter/video/keyword:Keyword labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:filter/video/keyword:Keyword labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Keyword struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Keyword matching logic. Valid values: `or`, `and`.
	Match pulumi.StringOutput `pulumi:"match"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// List of keywords. The structure of `word` block is documented below.
	Words KeywordWordArrayOutput `pulumi:"words"`
}

// NewKeyword registers a new resource with the given unique name, arguments, and options.
func NewKeyword(ctx *pulumi.Context,
	name string, args *KeywordArgs, opts ...pulumi.ResourceOption) (*Keyword, error) {
	if args == nil {
		args = &KeywordArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Keyword
	err := ctx.RegisterResource("fortios:filter/video/keyword:Keyword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyword gets an existing Keyword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeywordState, opts ...pulumi.ResourceOption) (*Keyword, error) {
	var resource Keyword
	err := ctx.ReadResource("fortios:filter/video/keyword:Keyword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Keyword resources.
type keywordState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Keyword matching logic. Valid values: `or`, `and`.
	Match *string `pulumi:"match"`
	// Name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// List of keywords. The structure of `word` block is documented below.
	Words []KeywordWord `pulumi:"words"`
}

type KeywordState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Keyword matching logic. Valid values: `or`, `and`.
	Match pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// List of keywords. The structure of `word` block is documented below.
	Words KeywordWordArrayInput
}

func (KeywordState) ElementType() reflect.Type {
	return reflect.TypeOf((*keywordState)(nil)).Elem()
}

type keywordArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Keyword matching logic. Valid values: `or`, `and`.
	Match *string `pulumi:"match"`
	// Name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// List of keywords. The structure of `word` block is documented below.
	Words []KeywordWord `pulumi:"words"`
}

// The set of arguments for constructing a Keyword resource.
type KeywordArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Keyword matching logic. Valid values: `or`, `and`.
	Match pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// List of keywords. The structure of `word` block is documented below.
	Words KeywordWordArrayInput
}

func (KeywordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keywordArgs)(nil)).Elem()
}

type KeywordInput interface {
	pulumi.Input

	ToKeywordOutput() KeywordOutput
	ToKeywordOutputWithContext(ctx context.Context) KeywordOutput
}

func (*Keyword) ElementType() reflect.Type {
	return reflect.TypeOf((**Keyword)(nil)).Elem()
}

func (i *Keyword) ToKeywordOutput() KeywordOutput {
	return i.ToKeywordOutputWithContext(context.Background())
}

func (i *Keyword) ToKeywordOutputWithContext(ctx context.Context) KeywordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeywordOutput)
}

// KeywordArrayInput is an input type that accepts KeywordArray and KeywordArrayOutput values.
// You can construct a concrete instance of `KeywordArrayInput` via:
//
//	KeywordArray{ KeywordArgs{...} }
type KeywordArrayInput interface {
	pulumi.Input

	ToKeywordArrayOutput() KeywordArrayOutput
	ToKeywordArrayOutputWithContext(context.Context) KeywordArrayOutput
}

type KeywordArray []KeywordInput

func (KeywordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Keyword)(nil)).Elem()
}

func (i KeywordArray) ToKeywordArrayOutput() KeywordArrayOutput {
	return i.ToKeywordArrayOutputWithContext(context.Background())
}

func (i KeywordArray) ToKeywordArrayOutputWithContext(ctx context.Context) KeywordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeywordArrayOutput)
}

// KeywordMapInput is an input type that accepts KeywordMap and KeywordMapOutput values.
// You can construct a concrete instance of `KeywordMapInput` via:
//
//	KeywordMap{ "key": KeywordArgs{...} }
type KeywordMapInput interface {
	pulumi.Input

	ToKeywordMapOutput() KeywordMapOutput
	ToKeywordMapOutputWithContext(context.Context) KeywordMapOutput
}

type KeywordMap map[string]KeywordInput

func (KeywordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Keyword)(nil)).Elem()
}

func (i KeywordMap) ToKeywordMapOutput() KeywordMapOutput {
	return i.ToKeywordMapOutputWithContext(context.Background())
}

func (i KeywordMap) ToKeywordMapOutputWithContext(ctx context.Context) KeywordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeywordMapOutput)
}

type KeywordOutput struct{ *pulumi.OutputState }

func (KeywordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Keyword)(nil)).Elem()
}

func (o KeywordOutput) ToKeywordOutput() KeywordOutput {
	return o
}

func (o KeywordOutput) ToKeywordOutputWithContext(ctx context.Context) KeywordOutput {
	return o
}

// Comment.
func (o KeywordOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Keyword) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o KeywordOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Keyword) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// ID.
func (o KeywordOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Keyword) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o KeywordOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Keyword) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Keyword matching logic. Valid values: `or`, `and`.
func (o KeywordOutput) Match() pulumi.StringOutput {
	return o.ApplyT(func(v *Keyword) pulumi.StringOutput { return v.Match }).(pulumi.StringOutput)
}

// Name.
func (o KeywordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Keyword) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o KeywordOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Keyword) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// List of keywords. The structure of `word` block is documented below.
func (o KeywordOutput) Words() KeywordWordArrayOutput {
	return o.ApplyT(func(v *Keyword) KeywordWordArrayOutput { return v.Words }).(KeywordWordArrayOutput)
}

type KeywordArrayOutput struct{ *pulumi.OutputState }

func (KeywordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Keyword)(nil)).Elem()
}

func (o KeywordArrayOutput) ToKeywordArrayOutput() KeywordArrayOutput {
	return o
}

func (o KeywordArrayOutput) ToKeywordArrayOutputWithContext(ctx context.Context) KeywordArrayOutput {
	return o
}

func (o KeywordArrayOutput) Index(i pulumi.IntInput) KeywordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Keyword {
		return vs[0].([]*Keyword)[vs[1].(int)]
	}).(KeywordOutput)
}

type KeywordMapOutput struct{ *pulumi.OutputState }

func (KeywordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Keyword)(nil)).Elem()
}

func (o KeywordMapOutput) ToKeywordMapOutput() KeywordMapOutput {
	return o
}

func (o KeywordMapOutput) ToKeywordMapOutputWithContext(ctx context.Context) KeywordMapOutput {
	return o
}

func (o KeywordMapOutput) MapIndex(k pulumi.StringInput) KeywordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Keyword {
		return vs[0].(map[string]*Keyword)[vs[1].(string)]
	}).(KeywordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeywordInput)(nil)).Elem(), &Keyword{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeywordArrayInput)(nil)).Elem(), KeywordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeywordMapInput)(nil)).Elem(), KeywordMap{})
	pulumi.RegisterOutputType(KeywordOutput{})
	pulumi.RegisterOutputType(KeywordArrayOutput{})
	pulumi.RegisterOutputType(KeywordMapOutput{})
}