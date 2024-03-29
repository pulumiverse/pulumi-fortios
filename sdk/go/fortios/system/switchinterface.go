// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure software switch interfaces by grouping physical and WiFi interfaces.
//
// ## Import
//
// System SwitchInterface can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/switchinterface:Switchinterface labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/switchinterface:Switchinterface labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Switchinterface struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
	IntraSwitchPolicy pulumi.StringOutput `pulumi:"intraSwitchPolicy"`
	// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
	MacTtl pulumi.IntOutput `pulumi:"macTtl"`
	// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
	Members SwitchinterfaceMemberArrayOutput `pulumi:"members"`
	// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
	Span pulumi.StringOutput `pulumi:"span"`
	// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
	SpanDestPort pulumi.StringOutput `pulumi:"spanDestPort"`
	// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
	SpanDirection pulumi.StringOutput `pulumi:"spanDirection"`
	// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
	SpanSourcePorts SwitchinterfaceSpanSourcePortArrayOutput `pulumi:"spanSourcePorts"`
	// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
	Type pulumi.StringOutput `pulumi:"type"`
	// VDOM that the software switch belongs to.
	Vdom pulumi.StringOutput `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchinterface registers a new resource with the given unique name, arguments, and options.
func NewSwitchinterface(ctx *pulumi.Context,
	name string, args *SwitchinterfaceArgs, opts ...pulumi.ResourceOption) (*Switchinterface, error) {
	if args == nil {
		args = &SwitchinterfaceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Switchinterface
	err := ctx.RegisterResource("fortios:system/switchinterface:Switchinterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchinterface gets an existing Switchinterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchinterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchinterfaceState, opts ...pulumi.ResourceOption) (*Switchinterface, error) {
	var resource Switchinterface
	err := ctx.ReadResource("fortios:system/switchinterface:Switchinterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Switchinterface resources.
type switchinterfaceState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
	IntraSwitchPolicy *string `pulumi:"intraSwitchPolicy"`
	// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
	MacTtl *int `pulumi:"macTtl"`
	// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
	Members []SwitchinterfaceMember `pulumi:"members"`
	// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
	Name *string `pulumi:"name"`
	// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
	Span *string `pulumi:"span"`
	// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
	SpanDestPort *string `pulumi:"spanDestPort"`
	// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
	SpanDirection *string `pulumi:"spanDirection"`
	// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
	SpanSourcePorts []SwitchinterfaceSpanSourcePort `pulumi:"spanSourcePorts"`
	// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
	Type *string `pulumi:"type"`
	// VDOM that the software switch belongs to.
	Vdom *string `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchinterfaceState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
	IntraSwitchPolicy pulumi.StringPtrInput
	// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
	MacTtl pulumi.IntPtrInput
	// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
	Members SwitchinterfaceMemberArrayInput
	// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
	Name pulumi.StringPtrInput
	// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
	Span pulumi.StringPtrInput
	// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
	SpanDestPort pulumi.StringPtrInput
	// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
	SpanDirection pulumi.StringPtrInput
	// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
	SpanSourcePorts SwitchinterfaceSpanSourcePortArrayInput
	// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
	Type pulumi.StringPtrInput
	// VDOM that the software switch belongs to.
	Vdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchinterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchinterfaceState)(nil)).Elem()
}

type switchinterfaceArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
	IntraSwitchPolicy *string `pulumi:"intraSwitchPolicy"`
	// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
	MacTtl *int `pulumi:"macTtl"`
	// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
	Members []SwitchinterfaceMember `pulumi:"members"`
	// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
	Name *string `pulumi:"name"`
	// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
	Span *string `pulumi:"span"`
	// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
	SpanDestPort *string `pulumi:"spanDestPort"`
	// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
	SpanDirection *string `pulumi:"spanDirection"`
	// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
	SpanSourcePorts []SwitchinterfaceSpanSourcePort `pulumi:"spanSourcePorts"`
	// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
	Type *string `pulumi:"type"`
	// VDOM that the software switch belongs to.
	Vdom *string `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Switchinterface resource.
type SwitchinterfaceArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
	IntraSwitchPolicy pulumi.StringPtrInput
	// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
	MacTtl pulumi.IntPtrInput
	// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
	Members SwitchinterfaceMemberArrayInput
	// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
	Name pulumi.StringPtrInput
	// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
	Span pulumi.StringPtrInput
	// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
	SpanDestPort pulumi.StringPtrInput
	// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
	SpanDirection pulumi.StringPtrInput
	// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
	SpanSourcePorts SwitchinterfaceSpanSourcePortArrayInput
	// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
	Type pulumi.StringPtrInput
	// VDOM that the software switch belongs to.
	Vdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchinterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchinterfaceArgs)(nil)).Elem()
}

type SwitchinterfaceInput interface {
	pulumi.Input

	ToSwitchinterfaceOutput() SwitchinterfaceOutput
	ToSwitchinterfaceOutputWithContext(ctx context.Context) SwitchinterfaceOutput
}

func (*Switchinterface) ElementType() reflect.Type {
	return reflect.TypeOf((**Switchinterface)(nil)).Elem()
}

func (i *Switchinterface) ToSwitchinterfaceOutput() SwitchinterfaceOutput {
	return i.ToSwitchinterfaceOutputWithContext(context.Background())
}

func (i *Switchinterface) ToSwitchinterfaceOutputWithContext(ctx context.Context) SwitchinterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchinterfaceOutput)
}

// SwitchinterfaceArrayInput is an input type that accepts SwitchinterfaceArray and SwitchinterfaceArrayOutput values.
// You can construct a concrete instance of `SwitchinterfaceArrayInput` via:
//
//	SwitchinterfaceArray{ SwitchinterfaceArgs{...} }
type SwitchinterfaceArrayInput interface {
	pulumi.Input

	ToSwitchinterfaceArrayOutput() SwitchinterfaceArrayOutput
	ToSwitchinterfaceArrayOutputWithContext(context.Context) SwitchinterfaceArrayOutput
}

type SwitchinterfaceArray []SwitchinterfaceInput

func (SwitchinterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Switchinterface)(nil)).Elem()
}

func (i SwitchinterfaceArray) ToSwitchinterfaceArrayOutput() SwitchinterfaceArrayOutput {
	return i.ToSwitchinterfaceArrayOutputWithContext(context.Background())
}

func (i SwitchinterfaceArray) ToSwitchinterfaceArrayOutputWithContext(ctx context.Context) SwitchinterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchinterfaceArrayOutput)
}

// SwitchinterfaceMapInput is an input type that accepts SwitchinterfaceMap and SwitchinterfaceMapOutput values.
// You can construct a concrete instance of `SwitchinterfaceMapInput` via:
//
//	SwitchinterfaceMap{ "key": SwitchinterfaceArgs{...} }
type SwitchinterfaceMapInput interface {
	pulumi.Input

	ToSwitchinterfaceMapOutput() SwitchinterfaceMapOutput
	ToSwitchinterfaceMapOutputWithContext(context.Context) SwitchinterfaceMapOutput
}

type SwitchinterfaceMap map[string]SwitchinterfaceInput

func (SwitchinterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Switchinterface)(nil)).Elem()
}

func (i SwitchinterfaceMap) ToSwitchinterfaceMapOutput() SwitchinterfaceMapOutput {
	return i.ToSwitchinterfaceMapOutputWithContext(context.Background())
}

func (i SwitchinterfaceMap) ToSwitchinterfaceMapOutputWithContext(ctx context.Context) SwitchinterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchinterfaceMapOutput)
}

type SwitchinterfaceOutput struct{ *pulumi.OutputState }

func (SwitchinterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Switchinterface)(nil)).Elem()
}

func (o SwitchinterfaceOutput) ToSwitchinterfaceOutput() SwitchinterfaceOutput {
	return o
}

func (o SwitchinterfaceOutput) ToSwitchinterfaceOutputWithContext(ctx context.Context) SwitchinterfaceOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SwitchinterfaceOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Switchinterface) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o SwitchinterfaceOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Switchinterface) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
func (o SwitchinterfaceOutput) IntraSwitchPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Switchinterface) pulumi.StringOutput { return v.IntraSwitchPolicy }).(pulumi.StringOutput)
}

// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
func (o SwitchinterfaceOutput) MacTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *Switchinterface) pulumi.IntOutput { return v.MacTtl }).(pulumi.IntOutput)
}

// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
func (o SwitchinterfaceOutput) Members() SwitchinterfaceMemberArrayOutput {
	return o.ApplyT(func(v *Switchinterface) SwitchinterfaceMemberArrayOutput { return v.Members }).(SwitchinterfaceMemberArrayOutput)
}

// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
func (o SwitchinterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Switchinterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
func (o SwitchinterfaceOutput) Span() pulumi.StringOutput {
	return o.ApplyT(func(v *Switchinterface) pulumi.StringOutput { return v.Span }).(pulumi.StringOutput)
}

// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
func (o SwitchinterfaceOutput) SpanDestPort() pulumi.StringOutput {
	return o.ApplyT(func(v *Switchinterface) pulumi.StringOutput { return v.SpanDestPort }).(pulumi.StringOutput)
}

// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
func (o SwitchinterfaceOutput) SpanDirection() pulumi.StringOutput {
	return o.ApplyT(func(v *Switchinterface) pulumi.StringOutput { return v.SpanDirection }).(pulumi.StringOutput)
}

// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
func (o SwitchinterfaceOutput) SpanSourcePorts() SwitchinterfaceSpanSourcePortArrayOutput {
	return o.ApplyT(func(v *Switchinterface) SwitchinterfaceSpanSourcePortArrayOutput { return v.SpanSourcePorts }).(SwitchinterfaceSpanSourcePortArrayOutput)
}

// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
func (o SwitchinterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Switchinterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// VDOM that the software switch belongs to.
func (o SwitchinterfaceOutput) Vdom() pulumi.StringOutput {
	return o.ApplyT(func(v *Switchinterface) pulumi.StringOutput { return v.Vdom }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchinterfaceOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Switchinterface) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchinterfaceArrayOutput struct{ *pulumi.OutputState }

func (SwitchinterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Switchinterface)(nil)).Elem()
}

func (o SwitchinterfaceArrayOutput) ToSwitchinterfaceArrayOutput() SwitchinterfaceArrayOutput {
	return o
}

func (o SwitchinterfaceArrayOutput) ToSwitchinterfaceArrayOutputWithContext(ctx context.Context) SwitchinterfaceArrayOutput {
	return o
}

func (o SwitchinterfaceArrayOutput) Index(i pulumi.IntInput) SwitchinterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Switchinterface {
		return vs[0].([]*Switchinterface)[vs[1].(int)]
	}).(SwitchinterfaceOutput)
}

type SwitchinterfaceMapOutput struct{ *pulumi.OutputState }

func (SwitchinterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Switchinterface)(nil)).Elem()
}

func (o SwitchinterfaceMapOutput) ToSwitchinterfaceMapOutput() SwitchinterfaceMapOutput {
	return o
}

func (o SwitchinterfaceMapOutput) ToSwitchinterfaceMapOutputWithContext(ctx context.Context) SwitchinterfaceMapOutput {
	return o
}

func (o SwitchinterfaceMapOutput) MapIndex(k pulumi.StringInput) SwitchinterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Switchinterface {
		return vs[0].(map[string]*Switchinterface)[vs[1].(string)]
	}).(SwitchinterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchinterfaceInput)(nil)).Elem(), &Switchinterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchinterfaceArrayInput)(nil)).Elem(), SwitchinterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchinterfaceMapInput)(nil)).Elem(), SwitchinterfaceMap{})
	pulumi.RegisterOutputType(SwitchinterfaceOutput{})
	pulumi.RegisterOutputType(SwitchinterfaceArrayOutput{})
	pulumi.RegisterOutputType(SwitchinterfaceMapOutput{})
}
