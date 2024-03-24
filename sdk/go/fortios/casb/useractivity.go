// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package casb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure CASB user activity. Applies to FortiOS Version `>= 7.4.1`.
//
// ## Import
//
// Casb UserActivity can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:casb/useractivity:Useractivity labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:casb/useractivity:Useractivity labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Useractivity struct {
	pulumi.CustomResourceState

	// CASB SaaS application name.
	Application pulumi.StringOutput `pulumi:"application"`
	// CASB user activity signature name.
	CasbName pulumi.StringOutput `pulumi:"casbName"`
	// CASB user activity category. Valid values: `activity-control`, `tenant-control`, `domain-control`, `safe-search-control`, `other`.
	Category pulumi.StringOutput `pulumi:"category"`
	// CASB control options. The structure of `controlOptions` block is documented below.
	ControlOptions UseractivityControlOptionArrayOutput `pulumi:"controlOptions"`
	// CASB user activity description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// CASB user activity match strategy. Valid values: `and`, `or`.
	MatchStrategy pulumi.StringOutput `pulumi:"matchStrategy"`
	// CASB user activity match rules. The structure of `match` block is documented below.
	Matches UseractivityMatchArrayOutput `pulumi:"matches"`
	// CASB user activity name.
	Name pulumi.StringOutput `pulumi:"name"`
	// CASB user activity status. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// CASB user activity type. Valid values: `built-in`, `customized`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUseractivity registers a new resource with the given unique name, arguments, and options.
func NewUseractivity(ctx *pulumi.Context,
	name string, args *UseractivityArgs, opts ...pulumi.ResourceOption) (*Useractivity, error) {
	if args == nil {
		args = &UseractivityArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Useractivity
	err := ctx.RegisterResource("fortios:casb/useractivity:Useractivity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUseractivity gets an existing Useractivity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUseractivity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UseractivityState, opts ...pulumi.ResourceOption) (*Useractivity, error) {
	var resource Useractivity
	err := ctx.ReadResource("fortios:casb/useractivity:Useractivity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Useractivity resources.
type useractivityState struct {
	// CASB SaaS application name.
	Application *string `pulumi:"application"`
	// CASB user activity signature name.
	CasbName *string `pulumi:"casbName"`
	// CASB user activity category. Valid values: `activity-control`, `tenant-control`, `domain-control`, `safe-search-control`, `other`.
	Category *string `pulumi:"category"`
	// CASB control options. The structure of `controlOptions` block is documented below.
	ControlOptions []UseractivityControlOption `pulumi:"controlOptions"`
	// CASB user activity description.
	Description *string `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// CASB user activity match strategy. Valid values: `and`, `or`.
	MatchStrategy *string `pulumi:"matchStrategy"`
	// CASB user activity match rules. The structure of `match` block is documented below.
	Matches []UseractivityMatch `pulumi:"matches"`
	// CASB user activity name.
	Name *string `pulumi:"name"`
	// CASB user activity status. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// CASB user activity type. Valid values: `built-in`, `customized`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UseractivityState struct {
	// CASB SaaS application name.
	Application pulumi.StringPtrInput
	// CASB user activity signature name.
	CasbName pulumi.StringPtrInput
	// CASB user activity category. Valid values: `activity-control`, `tenant-control`, `domain-control`, `safe-search-control`, `other`.
	Category pulumi.StringPtrInput
	// CASB control options. The structure of `controlOptions` block is documented below.
	ControlOptions UseractivityControlOptionArrayInput
	// CASB user activity description.
	Description pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// CASB user activity match strategy. Valid values: `and`, `or`.
	MatchStrategy pulumi.StringPtrInput
	// CASB user activity match rules. The structure of `match` block is documented below.
	Matches UseractivityMatchArrayInput
	// CASB user activity name.
	Name pulumi.StringPtrInput
	// CASB user activity status. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// CASB user activity type. Valid values: `built-in`, `customized`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UseractivityState) ElementType() reflect.Type {
	return reflect.TypeOf((*useractivityState)(nil)).Elem()
}

type useractivityArgs struct {
	// CASB SaaS application name.
	Application *string `pulumi:"application"`
	// CASB user activity signature name.
	CasbName *string `pulumi:"casbName"`
	// CASB user activity category. Valid values: `activity-control`, `tenant-control`, `domain-control`, `safe-search-control`, `other`.
	Category *string `pulumi:"category"`
	// CASB control options. The structure of `controlOptions` block is documented below.
	ControlOptions []UseractivityControlOption `pulumi:"controlOptions"`
	// CASB user activity description.
	Description *string `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// CASB user activity match strategy. Valid values: `and`, `or`.
	MatchStrategy *string `pulumi:"matchStrategy"`
	// CASB user activity match rules. The structure of `match` block is documented below.
	Matches []UseractivityMatch `pulumi:"matches"`
	// CASB user activity name.
	Name *string `pulumi:"name"`
	// CASB user activity status. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// CASB user activity type. Valid values: `built-in`, `customized`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Useractivity resource.
type UseractivityArgs struct {
	// CASB SaaS application name.
	Application pulumi.StringPtrInput
	// CASB user activity signature name.
	CasbName pulumi.StringPtrInput
	// CASB user activity category. Valid values: `activity-control`, `tenant-control`, `domain-control`, `safe-search-control`, `other`.
	Category pulumi.StringPtrInput
	// CASB control options. The structure of `controlOptions` block is documented below.
	ControlOptions UseractivityControlOptionArrayInput
	// CASB user activity description.
	Description pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// CASB user activity match strategy. Valid values: `and`, `or`.
	MatchStrategy pulumi.StringPtrInput
	// CASB user activity match rules. The structure of `match` block is documented below.
	Matches UseractivityMatchArrayInput
	// CASB user activity name.
	Name pulumi.StringPtrInput
	// CASB user activity status. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// CASB user activity type. Valid values: `built-in`, `customized`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UseractivityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*useractivityArgs)(nil)).Elem()
}

type UseractivityInput interface {
	pulumi.Input

	ToUseractivityOutput() UseractivityOutput
	ToUseractivityOutputWithContext(ctx context.Context) UseractivityOutput
}

func (*Useractivity) ElementType() reflect.Type {
	return reflect.TypeOf((**Useractivity)(nil)).Elem()
}

func (i *Useractivity) ToUseractivityOutput() UseractivityOutput {
	return i.ToUseractivityOutputWithContext(context.Background())
}

func (i *Useractivity) ToUseractivityOutputWithContext(ctx context.Context) UseractivityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UseractivityOutput)
}

// UseractivityArrayInput is an input type that accepts UseractivityArray and UseractivityArrayOutput values.
// You can construct a concrete instance of `UseractivityArrayInput` via:
//
//	UseractivityArray{ UseractivityArgs{...} }
type UseractivityArrayInput interface {
	pulumi.Input

	ToUseractivityArrayOutput() UseractivityArrayOutput
	ToUseractivityArrayOutputWithContext(context.Context) UseractivityArrayOutput
}

type UseractivityArray []UseractivityInput

func (UseractivityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Useractivity)(nil)).Elem()
}

func (i UseractivityArray) ToUseractivityArrayOutput() UseractivityArrayOutput {
	return i.ToUseractivityArrayOutputWithContext(context.Background())
}

func (i UseractivityArray) ToUseractivityArrayOutputWithContext(ctx context.Context) UseractivityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UseractivityArrayOutput)
}

// UseractivityMapInput is an input type that accepts UseractivityMap and UseractivityMapOutput values.
// You can construct a concrete instance of `UseractivityMapInput` via:
//
//	UseractivityMap{ "key": UseractivityArgs{...} }
type UseractivityMapInput interface {
	pulumi.Input

	ToUseractivityMapOutput() UseractivityMapOutput
	ToUseractivityMapOutputWithContext(context.Context) UseractivityMapOutput
}

type UseractivityMap map[string]UseractivityInput

func (UseractivityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Useractivity)(nil)).Elem()
}

func (i UseractivityMap) ToUseractivityMapOutput() UseractivityMapOutput {
	return i.ToUseractivityMapOutputWithContext(context.Background())
}

func (i UseractivityMap) ToUseractivityMapOutputWithContext(ctx context.Context) UseractivityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UseractivityMapOutput)
}

type UseractivityOutput struct{ *pulumi.OutputState }

func (UseractivityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Useractivity)(nil)).Elem()
}

func (o UseractivityOutput) ToUseractivityOutput() UseractivityOutput {
	return o
}

func (o UseractivityOutput) ToUseractivityOutputWithContext(ctx context.Context) UseractivityOutput {
	return o
}

// CASB SaaS application name.
func (o UseractivityOutput) Application() pulumi.StringOutput {
	return o.ApplyT(func(v *Useractivity) pulumi.StringOutput { return v.Application }).(pulumi.StringOutput)
}

// CASB user activity signature name.
func (o UseractivityOutput) CasbName() pulumi.StringOutput {
	return o.ApplyT(func(v *Useractivity) pulumi.StringOutput { return v.CasbName }).(pulumi.StringOutput)
}

// CASB user activity category. Valid values: `activity-control`, `tenant-control`, `domain-control`, `safe-search-control`, `other`.
func (o UseractivityOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *Useractivity) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

// CASB control options. The structure of `controlOptions` block is documented below.
func (o UseractivityOutput) ControlOptions() UseractivityControlOptionArrayOutput {
	return o.ApplyT(func(v *Useractivity) UseractivityControlOptionArrayOutput { return v.ControlOptions }).(UseractivityControlOptionArrayOutput)
}

// CASB user activity description.
func (o UseractivityOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Useractivity) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o UseractivityOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Useractivity) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o UseractivityOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Useractivity) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// CASB user activity match strategy. Valid values: `and`, `or`.
func (o UseractivityOutput) MatchStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *Useractivity) pulumi.StringOutput { return v.MatchStrategy }).(pulumi.StringOutput)
}

// CASB user activity match rules. The structure of `match` block is documented below.
func (o UseractivityOutput) Matches() UseractivityMatchArrayOutput {
	return o.ApplyT(func(v *Useractivity) UseractivityMatchArrayOutput { return v.Matches }).(UseractivityMatchArrayOutput)
}

// CASB user activity name.
func (o UseractivityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Useractivity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// CASB user activity status. Valid values: `enable`, `disable`.
func (o UseractivityOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Useractivity) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// CASB user activity type. Valid values: `built-in`, `customized`.
func (o UseractivityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Useractivity) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o UseractivityOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Useractivity) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o UseractivityOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Useractivity) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type UseractivityArrayOutput struct{ *pulumi.OutputState }

func (UseractivityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Useractivity)(nil)).Elem()
}

func (o UseractivityArrayOutput) ToUseractivityArrayOutput() UseractivityArrayOutput {
	return o
}

func (o UseractivityArrayOutput) ToUseractivityArrayOutputWithContext(ctx context.Context) UseractivityArrayOutput {
	return o
}

func (o UseractivityArrayOutput) Index(i pulumi.IntInput) UseractivityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Useractivity {
		return vs[0].([]*Useractivity)[vs[1].(int)]
	}).(UseractivityOutput)
}

type UseractivityMapOutput struct{ *pulumi.OutputState }

func (UseractivityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Useractivity)(nil)).Elem()
}

func (o UseractivityMapOutput) ToUseractivityMapOutput() UseractivityMapOutput {
	return o
}

func (o UseractivityMapOutput) ToUseractivityMapOutputWithContext(ctx context.Context) UseractivityMapOutput {
	return o
}

func (o UseractivityMapOutput) MapIndex(k pulumi.StringInput) UseractivityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Useractivity {
		return vs[0].(map[string]*Useractivity)[vs[1].(string)]
	}).(UseractivityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UseractivityInput)(nil)).Elem(), &Useractivity{})
	pulumi.RegisterInputType(reflect.TypeOf((*UseractivityArrayInput)(nil)).Elem(), UseractivityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UseractivityMapInput)(nil)).Elem(), UseractivityMap{})
	pulumi.RegisterOutputType(UseractivityOutput{})
	pulumi.RegisterOutputType(UseractivityArrayOutput{})
	pulumi.RegisterOutputType(UseractivityMapOutput{})
}