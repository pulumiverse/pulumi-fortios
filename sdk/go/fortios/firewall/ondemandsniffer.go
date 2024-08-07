// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure on-demand packet sniffer. Applies to FortiOS Version `>= 7.4.4`.
//
// ## Import
//
// Firewall OnDemandSniffer can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/ondemandsniffer:Ondemandsniffer labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/ondemandsniffer:Ondemandsniffer labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Ondemandsniffer struct {
	pulumi.CustomResourceState

	// Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
	AdvancedFilter pulumi.StringPtrOutput `pulumi:"advancedFilter"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
	Hosts OndemandsnifferHostArrayOutput `pulumi:"hosts"`
	// Interface name that on-demand packet sniffer will take place.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Maximum number of packets to capture per on-demand packet sniffer.
	MaxPacketCount pulumi.IntOutput `pulumi:"maxPacketCount"`
	// On-demand packet sniffer name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Include non-IP packets. Valid values: `enable`, `disable`.
	NonIpPacket pulumi.StringOutput `pulumi:"nonIpPacket"`
	// Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
	Ports OndemandsnifferPortArrayOutput `pulumi:"ports"`
	// Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
	Protocols OndemandsnifferProtocolArrayOutput `pulumi:"protocols"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewOndemandsniffer registers a new resource with the given unique name, arguments, and options.
func NewOndemandsniffer(ctx *pulumi.Context,
	name string, args *OndemandsnifferArgs, opts ...pulumi.ResourceOption) (*Ondemandsniffer, error) {
	if args == nil {
		args = &OndemandsnifferArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ondemandsniffer
	err := ctx.RegisterResource("fortios:firewall/ondemandsniffer:Ondemandsniffer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOndemandsniffer gets an existing Ondemandsniffer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOndemandsniffer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OndemandsnifferState, opts ...pulumi.ResourceOption) (*Ondemandsniffer, error) {
	var resource Ondemandsniffer
	err := ctx.ReadResource("fortios:firewall/ondemandsniffer:Ondemandsniffer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ondemandsniffer resources.
type ondemandsnifferState struct {
	// Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
	AdvancedFilter *string `pulumi:"advancedFilter"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
	Hosts []OndemandsnifferHost `pulumi:"hosts"`
	// Interface name that on-demand packet sniffer will take place.
	Interface *string `pulumi:"interface"`
	// Maximum number of packets to capture per on-demand packet sniffer.
	MaxPacketCount *int `pulumi:"maxPacketCount"`
	// On-demand packet sniffer name.
	Name *string `pulumi:"name"`
	// Include non-IP packets. Valid values: `enable`, `disable`.
	NonIpPacket *string `pulumi:"nonIpPacket"`
	// Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
	Ports []OndemandsnifferPort `pulumi:"ports"`
	// Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
	Protocols []OndemandsnifferProtocol `pulumi:"protocols"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type OndemandsnifferState struct {
	// Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
	AdvancedFilter pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
	Hosts OndemandsnifferHostArrayInput
	// Interface name that on-demand packet sniffer will take place.
	Interface pulumi.StringPtrInput
	// Maximum number of packets to capture per on-demand packet sniffer.
	MaxPacketCount pulumi.IntPtrInput
	// On-demand packet sniffer name.
	Name pulumi.StringPtrInput
	// Include non-IP packets. Valid values: `enable`, `disable`.
	NonIpPacket pulumi.StringPtrInput
	// Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
	Ports OndemandsnifferPortArrayInput
	// Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
	Protocols OndemandsnifferProtocolArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (OndemandsnifferState) ElementType() reflect.Type {
	return reflect.TypeOf((*ondemandsnifferState)(nil)).Elem()
}

type ondemandsnifferArgs struct {
	// Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
	AdvancedFilter *string `pulumi:"advancedFilter"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
	Hosts []OndemandsnifferHost `pulumi:"hosts"`
	// Interface name that on-demand packet sniffer will take place.
	Interface *string `pulumi:"interface"`
	// Maximum number of packets to capture per on-demand packet sniffer.
	MaxPacketCount *int `pulumi:"maxPacketCount"`
	// On-demand packet sniffer name.
	Name *string `pulumi:"name"`
	// Include non-IP packets. Valid values: `enable`, `disable`.
	NonIpPacket *string `pulumi:"nonIpPacket"`
	// Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
	Ports []OndemandsnifferPort `pulumi:"ports"`
	// Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
	Protocols []OndemandsnifferProtocol `pulumi:"protocols"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Ondemandsniffer resource.
type OndemandsnifferArgs struct {
	// Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
	AdvancedFilter pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
	Hosts OndemandsnifferHostArrayInput
	// Interface name that on-demand packet sniffer will take place.
	Interface pulumi.StringPtrInput
	// Maximum number of packets to capture per on-demand packet sniffer.
	MaxPacketCount pulumi.IntPtrInput
	// On-demand packet sniffer name.
	Name pulumi.StringPtrInput
	// Include non-IP packets. Valid values: `enable`, `disable`.
	NonIpPacket pulumi.StringPtrInput
	// Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
	Ports OndemandsnifferPortArrayInput
	// Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
	Protocols OndemandsnifferProtocolArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (OndemandsnifferArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ondemandsnifferArgs)(nil)).Elem()
}

type OndemandsnifferInput interface {
	pulumi.Input

	ToOndemandsnifferOutput() OndemandsnifferOutput
	ToOndemandsnifferOutputWithContext(ctx context.Context) OndemandsnifferOutput
}

func (*Ondemandsniffer) ElementType() reflect.Type {
	return reflect.TypeOf((**Ondemandsniffer)(nil)).Elem()
}

func (i *Ondemandsniffer) ToOndemandsnifferOutput() OndemandsnifferOutput {
	return i.ToOndemandsnifferOutputWithContext(context.Background())
}

func (i *Ondemandsniffer) ToOndemandsnifferOutputWithContext(ctx context.Context) OndemandsnifferOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OndemandsnifferOutput)
}

// OndemandsnifferArrayInput is an input type that accepts OndemandsnifferArray and OndemandsnifferArrayOutput values.
// You can construct a concrete instance of `OndemandsnifferArrayInput` via:
//
//	OndemandsnifferArray{ OndemandsnifferArgs{...} }
type OndemandsnifferArrayInput interface {
	pulumi.Input

	ToOndemandsnifferArrayOutput() OndemandsnifferArrayOutput
	ToOndemandsnifferArrayOutputWithContext(context.Context) OndemandsnifferArrayOutput
}

type OndemandsnifferArray []OndemandsnifferInput

func (OndemandsnifferArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ondemandsniffer)(nil)).Elem()
}

func (i OndemandsnifferArray) ToOndemandsnifferArrayOutput() OndemandsnifferArrayOutput {
	return i.ToOndemandsnifferArrayOutputWithContext(context.Background())
}

func (i OndemandsnifferArray) ToOndemandsnifferArrayOutputWithContext(ctx context.Context) OndemandsnifferArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OndemandsnifferArrayOutput)
}

// OndemandsnifferMapInput is an input type that accepts OndemandsnifferMap and OndemandsnifferMapOutput values.
// You can construct a concrete instance of `OndemandsnifferMapInput` via:
//
//	OndemandsnifferMap{ "key": OndemandsnifferArgs{...} }
type OndemandsnifferMapInput interface {
	pulumi.Input

	ToOndemandsnifferMapOutput() OndemandsnifferMapOutput
	ToOndemandsnifferMapOutputWithContext(context.Context) OndemandsnifferMapOutput
}

type OndemandsnifferMap map[string]OndemandsnifferInput

func (OndemandsnifferMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ondemandsniffer)(nil)).Elem()
}

func (i OndemandsnifferMap) ToOndemandsnifferMapOutput() OndemandsnifferMapOutput {
	return i.ToOndemandsnifferMapOutputWithContext(context.Background())
}

func (i OndemandsnifferMap) ToOndemandsnifferMapOutputWithContext(ctx context.Context) OndemandsnifferMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OndemandsnifferMapOutput)
}

type OndemandsnifferOutput struct{ *pulumi.OutputState }

func (OndemandsnifferOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ondemandsniffer)(nil)).Elem()
}

func (o OndemandsnifferOutput) ToOndemandsnifferOutput() OndemandsnifferOutput {
	return o
}

func (o OndemandsnifferOutput) ToOndemandsnifferOutputWithContext(ctx context.Context) OndemandsnifferOutput {
	return o
}

// Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
func (o OndemandsnifferOutput) AdvancedFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ondemandsniffer) pulumi.StringPtrOutput { return v.AdvancedFilter }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o OndemandsnifferOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ondemandsniffer) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o OndemandsnifferOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ondemandsniffer) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
func (o OndemandsnifferOutput) Hosts() OndemandsnifferHostArrayOutput {
	return o.ApplyT(func(v *Ondemandsniffer) OndemandsnifferHostArrayOutput { return v.Hosts }).(OndemandsnifferHostArrayOutput)
}

// Interface name that on-demand packet sniffer will take place.
func (o OndemandsnifferOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Ondemandsniffer) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Maximum number of packets to capture per on-demand packet sniffer.
func (o OndemandsnifferOutput) MaxPacketCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Ondemandsniffer) pulumi.IntOutput { return v.MaxPacketCount }).(pulumi.IntOutput)
}

// On-demand packet sniffer name.
func (o OndemandsnifferOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ondemandsniffer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Include non-IP packets. Valid values: `enable`, `disable`.
func (o OndemandsnifferOutput) NonIpPacket() pulumi.StringOutput {
	return o.ApplyT(func(v *Ondemandsniffer) pulumi.StringOutput { return v.NonIpPacket }).(pulumi.StringOutput)
}

// Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
func (o OndemandsnifferOutput) Ports() OndemandsnifferPortArrayOutput {
	return o.ApplyT(func(v *Ondemandsniffer) OndemandsnifferPortArrayOutput { return v.Ports }).(OndemandsnifferPortArrayOutput)
}

// Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
func (o OndemandsnifferOutput) Protocols() OndemandsnifferProtocolArrayOutput {
	return o.ApplyT(func(v *Ondemandsniffer) OndemandsnifferProtocolArrayOutput { return v.Protocols }).(OndemandsnifferProtocolArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o OndemandsnifferOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Ondemandsniffer) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type OndemandsnifferArrayOutput struct{ *pulumi.OutputState }

func (OndemandsnifferArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ondemandsniffer)(nil)).Elem()
}

func (o OndemandsnifferArrayOutput) ToOndemandsnifferArrayOutput() OndemandsnifferArrayOutput {
	return o
}

func (o OndemandsnifferArrayOutput) ToOndemandsnifferArrayOutputWithContext(ctx context.Context) OndemandsnifferArrayOutput {
	return o
}

func (o OndemandsnifferArrayOutput) Index(i pulumi.IntInput) OndemandsnifferOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ondemandsniffer {
		return vs[0].([]*Ondemandsniffer)[vs[1].(int)]
	}).(OndemandsnifferOutput)
}

type OndemandsnifferMapOutput struct{ *pulumi.OutputState }

func (OndemandsnifferMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ondemandsniffer)(nil)).Elem()
}

func (o OndemandsnifferMapOutput) ToOndemandsnifferMapOutput() OndemandsnifferMapOutput {
	return o
}

func (o OndemandsnifferMapOutput) ToOndemandsnifferMapOutputWithContext(ctx context.Context) OndemandsnifferMapOutput {
	return o
}

func (o OndemandsnifferMapOutput) MapIndex(k pulumi.StringInput) OndemandsnifferOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ondemandsniffer {
		return vs[0].(map[string]*Ondemandsniffer)[vs[1].(string)]
	}).(OndemandsnifferOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OndemandsnifferInput)(nil)).Elem(), &Ondemandsniffer{})
	pulumi.RegisterInputType(reflect.TypeOf((*OndemandsnifferArrayInput)(nil)).Elem(), OndemandsnifferArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OndemandsnifferMapInput)(nil)).Elem(), OndemandsnifferMap{})
	pulumi.RegisterOutputType(OndemandsnifferOutput{})
	pulumi.RegisterOutputType(OndemandsnifferArrayOutput{})
	pulumi.RegisterOutputType(OndemandsnifferMapOutput{})
}
