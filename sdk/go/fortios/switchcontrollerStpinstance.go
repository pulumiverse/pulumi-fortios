// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch multiple spanning tree protocol (MSTP) instances. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # SwitchController StpInstance can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerStpinstance:SwitchcontrollerStpinstance labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerStpinstance:SwitchcontrollerStpinstance labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SwitchcontrollerStpinstance struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Instance ID.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
	VlanRanges SwitchcontrollerStpinstanceVlanRangeArrayOutput `pulumi:"vlanRanges"`
}

// NewSwitchcontrollerStpinstance registers a new resource with the given unique name, arguments, and options.
func NewSwitchcontrollerStpinstance(ctx *pulumi.Context,
	name string, args *SwitchcontrollerStpinstanceArgs, opts ...pulumi.ResourceOption) (*SwitchcontrollerStpinstance, error) {
	if args == nil {
		args = &SwitchcontrollerStpinstanceArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchcontrollerStpinstance
	err := ctx.RegisterResource("fortios:index/switchcontrollerStpinstance:SwitchcontrollerStpinstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchcontrollerStpinstance gets an existing SwitchcontrollerStpinstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchcontrollerStpinstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchcontrollerStpinstanceState, opts ...pulumi.ResourceOption) (*SwitchcontrollerStpinstance, error) {
	var resource SwitchcontrollerStpinstance
	err := ctx.ReadResource("fortios:index/switchcontrollerStpinstance:SwitchcontrollerStpinstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchcontrollerStpinstance resources.
type switchcontrollerStpinstanceState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Instance ID.
	Fosid *string `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
	VlanRanges []SwitchcontrollerStpinstanceVlanRange `pulumi:"vlanRanges"`
}

type SwitchcontrollerStpinstanceState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Instance ID.
	Fosid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
	VlanRanges SwitchcontrollerStpinstanceVlanRangeArrayInput
}

func (SwitchcontrollerStpinstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerStpinstanceState)(nil)).Elem()
}

type switchcontrollerStpinstanceArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Instance ID.
	Fosid *string `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
	VlanRanges []SwitchcontrollerStpinstanceVlanRange `pulumi:"vlanRanges"`
}

// The set of arguments for constructing a SwitchcontrollerStpinstance resource.
type SwitchcontrollerStpinstanceArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Instance ID.
	Fosid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
	VlanRanges SwitchcontrollerStpinstanceVlanRangeArrayInput
}

func (SwitchcontrollerStpinstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerStpinstanceArgs)(nil)).Elem()
}

type SwitchcontrollerStpinstanceInput interface {
	pulumi.Input

	ToSwitchcontrollerStpinstanceOutput() SwitchcontrollerStpinstanceOutput
	ToSwitchcontrollerStpinstanceOutputWithContext(ctx context.Context) SwitchcontrollerStpinstanceOutput
}

func (*SwitchcontrollerStpinstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerStpinstance)(nil)).Elem()
}

func (i *SwitchcontrollerStpinstance) ToSwitchcontrollerStpinstanceOutput() SwitchcontrollerStpinstanceOutput {
	return i.ToSwitchcontrollerStpinstanceOutputWithContext(context.Background())
}

func (i *SwitchcontrollerStpinstance) ToSwitchcontrollerStpinstanceOutputWithContext(ctx context.Context) SwitchcontrollerStpinstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerStpinstanceOutput)
}

// SwitchcontrollerStpinstanceArrayInput is an input type that accepts SwitchcontrollerStpinstanceArray and SwitchcontrollerStpinstanceArrayOutput values.
// You can construct a concrete instance of `SwitchcontrollerStpinstanceArrayInput` via:
//
//	SwitchcontrollerStpinstanceArray{ SwitchcontrollerStpinstanceArgs{...} }
type SwitchcontrollerStpinstanceArrayInput interface {
	pulumi.Input

	ToSwitchcontrollerStpinstanceArrayOutput() SwitchcontrollerStpinstanceArrayOutput
	ToSwitchcontrollerStpinstanceArrayOutputWithContext(context.Context) SwitchcontrollerStpinstanceArrayOutput
}

type SwitchcontrollerStpinstanceArray []SwitchcontrollerStpinstanceInput

func (SwitchcontrollerStpinstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerStpinstance)(nil)).Elem()
}

func (i SwitchcontrollerStpinstanceArray) ToSwitchcontrollerStpinstanceArrayOutput() SwitchcontrollerStpinstanceArrayOutput {
	return i.ToSwitchcontrollerStpinstanceArrayOutputWithContext(context.Background())
}

func (i SwitchcontrollerStpinstanceArray) ToSwitchcontrollerStpinstanceArrayOutputWithContext(ctx context.Context) SwitchcontrollerStpinstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerStpinstanceArrayOutput)
}

// SwitchcontrollerStpinstanceMapInput is an input type that accepts SwitchcontrollerStpinstanceMap and SwitchcontrollerStpinstanceMapOutput values.
// You can construct a concrete instance of `SwitchcontrollerStpinstanceMapInput` via:
//
//	SwitchcontrollerStpinstanceMap{ "key": SwitchcontrollerStpinstanceArgs{...} }
type SwitchcontrollerStpinstanceMapInput interface {
	pulumi.Input

	ToSwitchcontrollerStpinstanceMapOutput() SwitchcontrollerStpinstanceMapOutput
	ToSwitchcontrollerStpinstanceMapOutputWithContext(context.Context) SwitchcontrollerStpinstanceMapOutput
}

type SwitchcontrollerStpinstanceMap map[string]SwitchcontrollerStpinstanceInput

func (SwitchcontrollerStpinstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerStpinstance)(nil)).Elem()
}

func (i SwitchcontrollerStpinstanceMap) ToSwitchcontrollerStpinstanceMapOutput() SwitchcontrollerStpinstanceMapOutput {
	return i.ToSwitchcontrollerStpinstanceMapOutputWithContext(context.Background())
}

func (i SwitchcontrollerStpinstanceMap) ToSwitchcontrollerStpinstanceMapOutputWithContext(ctx context.Context) SwitchcontrollerStpinstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerStpinstanceMapOutput)
}

type SwitchcontrollerStpinstanceOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerStpinstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerStpinstance)(nil)).Elem()
}

func (o SwitchcontrollerStpinstanceOutput) ToSwitchcontrollerStpinstanceOutput() SwitchcontrollerStpinstanceOutput {
	return o
}

func (o SwitchcontrollerStpinstanceOutput) ToSwitchcontrollerStpinstanceOutputWithContext(ctx context.Context) SwitchcontrollerStpinstanceOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SwitchcontrollerStpinstanceOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerStpinstance) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Instance ID.
func (o SwitchcontrollerStpinstanceOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerStpinstance) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchcontrollerStpinstanceOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerStpinstance) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
func (o SwitchcontrollerStpinstanceOutput) VlanRanges() SwitchcontrollerStpinstanceVlanRangeArrayOutput {
	return o.ApplyT(func(v *SwitchcontrollerStpinstance) SwitchcontrollerStpinstanceVlanRangeArrayOutput {
		return v.VlanRanges
	}).(SwitchcontrollerStpinstanceVlanRangeArrayOutput)
}

type SwitchcontrollerStpinstanceArrayOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerStpinstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerStpinstance)(nil)).Elem()
}

func (o SwitchcontrollerStpinstanceArrayOutput) ToSwitchcontrollerStpinstanceArrayOutput() SwitchcontrollerStpinstanceArrayOutput {
	return o
}

func (o SwitchcontrollerStpinstanceArrayOutput) ToSwitchcontrollerStpinstanceArrayOutputWithContext(ctx context.Context) SwitchcontrollerStpinstanceArrayOutput {
	return o
}

func (o SwitchcontrollerStpinstanceArrayOutput) Index(i pulumi.IntInput) SwitchcontrollerStpinstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchcontrollerStpinstance {
		return vs[0].([]*SwitchcontrollerStpinstance)[vs[1].(int)]
	}).(SwitchcontrollerStpinstanceOutput)
}

type SwitchcontrollerStpinstanceMapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerStpinstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerStpinstance)(nil)).Elem()
}

func (o SwitchcontrollerStpinstanceMapOutput) ToSwitchcontrollerStpinstanceMapOutput() SwitchcontrollerStpinstanceMapOutput {
	return o
}

func (o SwitchcontrollerStpinstanceMapOutput) ToSwitchcontrollerStpinstanceMapOutputWithContext(ctx context.Context) SwitchcontrollerStpinstanceMapOutput {
	return o
}

func (o SwitchcontrollerStpinstanceMapOutput) MapIndex(k pulumi.StringInput) SwitchcontrollerStpinstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchcontrollerStpinstance {
		return vs[0].(map[string]*SwitchcontrollerStpinstance)[vs[1].(string)]
	}).(SwitchcontrollerStpinstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerStpinstanceInput)(nil)).Elem(), &SwitchcontrollerStpinstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerStpinstanceArrayInput)(nil)).Elem(), SwitchcontrollerStpinstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerStpinstanceMapInput)(nil)).Elem(), SwitchcontrollerStpinstanceMap{})
	pulumi.RegisterOutputType(SwitchcontrollerStpinstanceOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerStpinstanceArrayOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerStpinstanceMapOutput{})
}
