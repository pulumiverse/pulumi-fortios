// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiSwitch switch groups.
//
// ## Import
//
// SwitchController SwitchGroup can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/switchgroup:Switchgroup labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/switchgroup:Switchgroup labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Switchgroup struct {
	pulumi.CustomResourceState

	// Optional switch group description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// FortiLink interface to which switch group members belong.
	Fortilink pulumi.StringOutput `pulumi:"fortilink"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// FortiSwitch members belonging to this switch group. The structure of `members` block is documented below.
	Members SwitchgroupMemberArrayOutput `pulumi:"members"`
	// Switch group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewSwitchgroup registers a new resource with the given unique name, arguments, and options.
func NewSwitchgroup(ctx *pulumi.Context,
	name string, args *SwitchgroupArgs, opts ...pulumi.ResourceOption) (*Switchgroup, error) {
	if args == nil {
		args = &SwitchgroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Switchgroup
	err := ctx.RegisterResource("fortios:switchcontroller/switchgroup:Switchgroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchgroup gets an existing Switchgroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchgroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchgroupState, opts ...pulumi.ResourceOption) (*Switchgroup, error) {
	var resource Switchgroup
	err := ctx.ReadResource("fortios:switchcontroller/switchgroup:Switchgroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Switchgroup resources.
type switchgroupState struct {
	// Optional switch group description.
	Description *string `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiLink interface to which switch group members belong.
	Fortilink *string `pulumi:"fortilink"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// FortiSwitch members belonging to this switch group. The structure of `members` block is documented below.
	Members []SwitchgroupMember `pulumi:"members"`
	// Switch group name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchgroupState struct {
	// Optional switch group description.
	Description pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiLink interface to which switch group members belong.
	Fortilink pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// FortiSwitch members belonging to this switch group. The structure of `members` block is documented below.
	Members SwitchgroupMemberArrayInput
	// Switch group name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchgroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchgroupState)(nil)).Elem()
}

type switchgroupArgs struct {
	// Optional switch group description.
	Description *string `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiLink interface to which switch group members belong.
	Fortilink *string `pulumi:"fortilink"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// FortiSwitch members belonging to this switch group. The structure of `members` block is documented below.
	Members []SwitchgroupMember `pulumi:"members"`
	// Switch group name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Switchgroup resource.
type SwitchgroupArgs struct {
	// Optional switch group description.
	Description pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiLink interface to which switch group members belong.
	Fortilink pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// FortiSwitch members belonging to this switch group. The structure of `members` block is documented below.
	Members SwitchgroupMemberArrayInput
	// Switch group name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchgroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchgroupArgs)(nil)).Elem()
}

type SwitchgroupInput interface {
	pulumi.Input

	ToSwitchgroupOutput() SwitchgroupOutput
	ToSwitchgroupOutputWithContext(ctx context.Context) SwitchgroupOutput
}

func (*Switchgroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Switchgroup)(nil)).Elem()
}

func (i *Switchgroup) ToSwitchgroupOutput() SwitchgroupOutput {
	return i.ToSwitchgroupOutputWithContext(context.Background())
}

func (i *Switchgroup) ToSwitchgroupOutputWithContext(ctx context.Context) SwitchgroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchgroupOutput)
}

// SwitchgroupArrayInput is an input type that accepts SwitchgroupArray and SwitchgroupArrayOutput values.
// You can construct a concrete instance of `SwitchgroupArrayInput` via:
//
//	SwitchgroupArray{ SwitchgroupArgs{...} }
type SwitchgroupArrayInput interface {
	pulumi.Input

	ToSwitchgroupArrayOutput() SwitchgroupArrayOutput
	ToSwitchgroupArrayOutputWithContext(context.Context) SwitchgroupArrayOutput
}

type SwitchgroupArray []SwitchgroupInput

func (SwitchgroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Switchgroup)(nil)).Elem()
}

func (i SwitchgroupArray) ToSwitchgroupArrayOutput() SwitchgroupArrayOutput {
	return i.ToSwitchgroupArrayOutputWithContext(context.Background())
}

func (i SwitchgroupArray) ToSwitchgroupArrayOutputWithContext(ctx context.Context) SwitchgroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchgroupArrayOutput)
}

// SwitchgroupMapInput is an input type that accepts SwitchgroupMap and SwitchgroupMapOutput values.
// You can construct a concrete instance of `SwitchgroupMapInput` via:
//
//	SwitchgroupMap{ "key": SwitchgroupArgs{...} }
type SwitchgroupMapInput interface {
	pulumi.Input

	ToSwitchgroupMapOutput() SwitchgroupMapOutput
	ToSwitchgroupMapOutputWithContext(context.Context) SwitchgroupMapOutput
}

type SwitchgroupMap map[string]SwitchgroupInput

func (SwitchgroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Switchgroup)(nil)).Elem()
}

func (i SwitchgroupMap) ToSwitchgroupMapOutput() SwitchgroupMapOutput {
	return i.ToSwitchgroupMapOutputWithContext(context.Background())
}

func (i SwitchgroupMap) ToSwitchgroupMapOutputWithContext(ctx context.Context) SwitchgroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchgroupMapOutput)
}

type SwitchgroupOutput struct{ *pulumi.OutputState }

func (SwitchgroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Switchgroup)(nil)).Elem()
}

func (o SwitchgroupOutput) ToSwitchgroupOutput() SwitchgroupOutput {
	return o
}

func (o SwitchgroupOutput) ToSwitchgroupOutputWithContext(ctx context.Context) SwitchgroupOutput {
	return o
}

// Optional switch group description.
func (o SwitchgroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Switchgroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SwitchgroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Switchgroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// FortiLink interface to which switch group members belong.
func (o SwitchgroupOutput) Fortilink() pulumi.StringOutput {
	return o.ApplyT(func(v *Switchgroup) pulumi.StringOutput { return v.Fortilink }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o SwitchgroupOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Switchgroup) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// FortiSwitch members belonging to this switch group. The structure of `members` block is documented below.
func (o SwitchgroupOutput) Members() SwitchgroupMemberArrayOutput {
	return o.ApplyT(func(v *Switchgroup) SwitchgroupMemberArrayOutput { return v.Members }).(SwitchgroupMemberArrayOutput)
}

// Switch group name.
func (o SwitchgroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Switchgroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchgroupOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Switchgroup) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type SwitchgroupArrayOutput struct{ *pulumi.OutputState }

func (SwitchgroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Switchgroup)(nil)).Elem()
}

func (o SwitchgroupArrayOutput) ToSwitchgroupArrayOutput() SwitchgroupArrayOutput {
	return o
}

func (o SwitchgroupArrayOutput) ToSwitchgroupArrayOutputWithContext(ctx context.Context) SwitchgroupArrayOutput {
	return o
}

func (o SwitchgroupArrayOutput) Index(i pulumi.IntInput) SwitchgroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Switchgroup {
		return vs[0].([]*Switchgroup)[vs[1].(int)]
	}).(SwitchgroupOutput)
}

type SwitchgroupMapOutput struct{ *pulumi.OutputState }

func (SwitchgroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Switchgroup)(nil)).Elem()
}

func (o SwitchgroupMapOutput) ToSwitchgroupMapOutput() SwitchgroupMapOutput {
	return o
}

func (o SwitchgroupMapOutput) ToSwitchgroupMapOutputWithContext(ctx context.Context) SwitchgroupMapOutput {
	return o
}

func (o SwitchgroupMapOutput) MapIndex(k pulumi.StringInput) SwitchgroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Switchgroup {
		return vs[0].(map[string]*Switchgroup)[vs[1].(string)]
	}).(SwitchgroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchgroupInput)(nil)).Elem(), &Switchgroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchgroupArrayInput)(nil)).Elem(), SwitchgroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchgroupMapInput)(nil)).Elem(), SwitchgroupMap{})
	pulumi.RegisterOutputType(SwitchgroupOutput{})
	pulumi.RegisterOutputType(SwitchgroupArrayOutput{})
	pulumi.RegisterOutputType(SwitchgroupMapOutput{})
}
