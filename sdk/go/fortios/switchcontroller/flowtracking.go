// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiSwitch flow tracking and export via ipfix/netflow. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// SwitchController FlowTracking can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/flowtracking:Flowtracking labelname SwitchControllerFlowTracking
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/flowtracking:Flowtracking labelname SwitchControllerFlowTracking
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Flowtracking struct {
	pulumi.CustomResourceState

	// Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
	Aggregates FlowtrackingAggregateArrayOutput `pulumi:"aggregates"`
	// Configure collector ip address.
	CollectorIp pulumi.StringOutput `pulumi:"collectorIp"`
	// Configure collector port number(0-65535, default=0).
	CollectorPort pulumi.IntOutput `pulumi:"collectorPort"`
	// Configure collectors for the flow. The structure of `collectors` block is documented below.
	Collectors FlowtrackingCollectorArrayOutput `pulumi:"collectors"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
	Format pulumi.StringOutput `pulumi:"format"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
	Level pulumi.StringOutput `pulumi:"level"`
	// Configure flow max export packet size (512-9216, default=512 bytes).
	MaxExportPktSize pulumi.IntOutput `pulumi:"maxExportPktSize"`
	// Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
	SampleMode pulumi.StringOutput `pulumi:"sampleMode"`
	// Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
	SampleRate pulumi.IntOutput `pulumi:"sampleRate"`
	// Configure template export period (1-60, default=5 minutes).
	TemplateExportPeriod pulumi.IntOutput `pulumi:"templateExportPeriod"`
	// Configure flow session general timeout (60-604800, default=3600 seconds).
	TimeoutGeneral pulumi.IntOutput `pulumi:"timeoutGeneral"`
	// Configure flow session ICMP timeout (60-604800, default=300 seconds).
	TimeoutIcmp pulumi.IntOutput `pulumi:"timeoutIcmp"`
	// Configure flow session max timeout (60-604800, default=604800 seconds).
	TimeoutMax pulumi.IntOutput `pulumi:"timeoutMax"`
	// Configure flow session TCP timeout (60-604800, default=3600 seconds).
	TimeoutTcp pulumi.IntOutput `pulumi:"timeoutTcp"`
	// Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
	TimeoutTcpFin pulumi.IntOutput `pulumi:"timeoutTcpFin"`
	// Configure flow session TCP RST timeout (60-604800, default=120 seconds).
	TimeoutTcpRst pulumi.IntOutput `pulumi:"timeoutTcpRst"`
	// Configure flow session UDP timeout (60-604800, default=300 seconds).
	TimeoutUdp pulumi.IntOutput `pulumi:"timeoutUdp"`
	// Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
	Transport pulumi.StringOutput `pulumi:"transport"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewFlowtracking registers a new resource with the given unique name, arguments, and options.
func NewFlowtracking(ctx *pulumi.Context,
	name string, args *FlowtrackingArgs, opts ...pulumi.ResourceOption) (*Flowtracking, error) {
	if args == nil {
		args = &FlowtrackingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Flowtracking
	err := ctx.RegisterResource("fortios:switchcontroller/flowtracking:Flowtracking", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlowtracking gets an existing Flowtracking resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowtracking(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowtrackingState, opts ...pulumi.ResourceOption) (*Flowtracking, error) {
	var resource Flowtracking
	err := ctx.ReadResource("fortios:switchcontroller/flowtracking:Flowtracking", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Flowtracking resources.
type flowtrackingState struct {
	// Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
	Aggregates []FlowtrackingAggregate `pulumi:"aggregates"`
	// Configure collector ip address.
	CollectorIp *string `pulumi:"collectorIp"`
	// Configure collector port number(0-65535, default=0).
	CollectorPort *int `pulumi:"collectorPort"`
	// Configure collectors for the flow. The structure of `collectors` block is documented below.
	Collectors []FlowtrackingCollector `pulumi:"collectors"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
	Format *string `pulumi:"format"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
	Level *string `pulumi:"level"`
	// Configure flow max export packet size (512-9216, default=512 bytes).
	MaxExportPktSize *int `pulumi:"maxExportPktSize"`
	// Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
	SampleMode *string `pulumi:"sampleMode"`
	// Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
	SampleRate *int `pulumi:"sampleRate"`
	// Configure template export period (1-60, default=5 minutes).
	TemplateExportPeriod *int `pulumi:"templateExportPeriod"`
	// Configure flow session general timeout (60-604800, default=3600 seconds).
	TimeoutGeneral *int `pulumi:"timeoutGeneral"`
	// Configure flow session ICMP timeout (60-604800, default=300 seconds).
	TimeoutIcmp *int `pulumi:"timeoutIcmp"`
	// Configure flow session max timeout (60-604800, default=604800 seconds).
	TimeoutMax *int `pulumi:"timeoutMax"`
	// Configure flow session TCP timeout (60-604800, default=3600 seconds).
	TimeoutTcp *int `pulumi:"timeoutTcp"`
	// Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
	TimeoutTcpFin *int `pulumi:"timeoutTcpFin"`
	// Configure flow session TCP RST timeout (60-604800, default=120 seconds).
	TimeoutTcpRst *int `pulumi:"timeoutTcpRst"`
	// Configure flow session UDP timeout (60-604800, default=300 seconds).
	TimeoutUdp *int `pulumi:"timeoutUdp"`
	// Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
	Transport *string `pulumi:"transport"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FlowtrackingState struct {
	// Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
	Aggregates FlowtrackingAggregateArrayInput
	// Configure collector ip address.
	CollectorIp pulumi.StringPtrInput
	// Configure collector port number(0-65535, default=0).
	CollectorPort pulumi.IntPtrInput
	// Configure collectors for the flow. The structure of `collectors` block is documented below.
	Collectors FlowtrackingCollectorArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
	Format pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
	Level pulumi.StringPtrInput
	// Configure flow max export packet size (512-9216, default=512 bytes).
	MaxExportPktSize pulumi.IntPtrInput
	// Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
	SampleMode pulumi.StringPtrInput
	// Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
	SampleRate pulumi.IntPtrInput
	// Configure template export period (1-60, default=5 minutes).
	TemplateExportPeriod pulumi.IntPtrInput
	// Configure flow session general timeout (60-604800, default=3600 seconds).
	TimeoutGeneral pulumi.IntPtrInput
	// Configure flow session ICMP timeout (60-604800, default=300 seconds).
	TimeoutIcmp pulumi.IntPtrInput
	// Configure flow session max timeout (60-604800, default=604800 seconds).
	TimeoutMax pulumi.IntPtrInput
	// Configure flow session TCP timeout (60-604800, default=3600 seconds).
	TimeoutTcp pulumi.IntPtrInput
	// Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
	TimeoutTcpFin pulumi.IntPtrInput
	// Configure flow session TCP RST timeout (60-604800, default=120 seconds).
	TimeoutTcpRst pulumi.IntPtrInput
	// Configure flow session UDP timeout (60-604800, default=300 seconds).
	TimeoutUdp pulumi.IntPtrInput
	// Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
	Transport pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FlowtrackingState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowtrackingState)(nil)).Elem()
}

type flowtrackingArgs struct {
	// Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
	Aggregates []FlowtrackingAggregate `pulumi:"aggregates"`
	// Configure collector ip address.
	CollectorIp *string `pulumi:"collectorIp"`
	// Configure collector port number(0-65535, default=0).
	CollectorPort *int `pulumi:"collectorPort"`
	// Configure collectors for the flow. The structure of `collectors` block is documented below.
	Collectors []FlowtrackingCollector `pulumi:"collectors"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
	Format *string `pulumi:"format"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
	Level *string `pulumi:"level"`
	// Configure flow max export packet size (512-9216, default=512 bytes).
	MaxExportPktSize *int `pulumi:"maxExportPktSize"`
	// Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
	SampleMode *string `pulumi:"sampleMode"`
	// Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
	SampleRate *int `pulumi:"sampleRate"`
	// Configure template export period (1-60, default=5 minutes).
	TemplateExportPeriod *int `pulumi:"templateExportPeriod"`
	// Configure flow session general timeout (60-604800, default=3600 seconds).
	TimeoutGeneral *int `pulumi:"timeoutGeneral"`
	// Configure flow session ICMP timeout (60-604800, default=300 seconds).
	TimeoutIcmp *int `pulumi:"timeoutIcmp"`
	// Configure flow session max timeout (60-604800, default=604800 seconds).
	TimeoutMax *int `pulumi:"timeoutMax"`
	// Configure flow session TCP timeout (60-604800, default=3600 seconds).
	TimeoutTcp *int `pulumi:"timeoutTcp"`
	// Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
	TimeoutTcpFin *int `pulumi:"timeoutTcpFin"`
	// Configure flow session TCP RST timeout (60-604800, default=120 seconds).
	TimeoutTcpRst *int `pulumi:"timeoutTcpRst"`
	// Configure flow session UDP timeout (60-604800, default=300 seconds).
	TimeoutUdp *int `pulumi:"timeoutUdp"`
	// Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
	Transport *string `pulumi:"transport"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Flowtracking resource.
type FlowtrackingArgs struct {
	// Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
	Aggregates FlowtrackingAggregateArrayInput
	// Configure collector ip address.
	CollectorIp pulumi.StringPtrInput
	// Configure collector port number(0-65535, default=0).
	CollectorPort pulumi.IntPtrInput
	// Configure collectors for the flow. The structure of `collectors` block is documented below.
	Collectors FlowtrackingCollectorArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
	Format pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
	Level pulumi.StringPtrInput
	// Configure flow max export packet size (512-9216, default=512 bytes).
	MaxExportPktSize pulumi.IntPtrInput
	// Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
	SampleMode pulumi.StringPtrInput
	// Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
	SampleRate pulumi.IntPtrInput
	// Configure template export period (1-60, default=5 minutes).
	TemplateExportPeriod pulumi.IntPtrInput
	// Configure flow session general timeout (60-604800, default=3600 seconds).
	TimeoutGeneral pulumi.IntPtrInput
	// Configure flow session ICMP timeout (60-604800, default=300 seconds).
	TimeoutIcmp pulumi.IntPtrInput
	// Configure flow session max timeout (60-604800, default=604800 seconds).
	TimeoutMax pulumi.IntPtrInput
	// Configure flow session TCP timeout (60-604800, default=3600 seconds).
	TimeoutTcp pulumi.IntPtrInput
	// Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
	TimeoutTcpFin pulumi.IntPtrInput
	// Configure flow session TCP RST timeout (60-604800, default=120 seconds).
	TimeoutTcpRst pulumi.IntPtrInput
	// Configure flow session UDP timeout (60-604800, default=300 seconds).
	TimeoutUdp pulumi.IntPtrInput
	// Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
	Transport pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FlowtrackingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowtrackingArgs)(nil)).Elem()
}

type FlowtrackingInput interface {
	pulumi.Input

	ToFlowtrackingOutput() FlowtrackingOutput
	ToFlowtrackingOutputWithContext(ctx context.Context) FlowtrackingOutput
}

func (*Flowtracking) ElementType() reflect.Type {
	return reflect.TypeOf((**Flowtracking)(nil)).Elem()
}

func (i *Flowtracking) ToFlowtrackingOutput() FlowtrackingOutput {
	return i.ToFlowtrackingOutputWithContext(context.Background())
}

func (i *Flowtracking) ToFlowtrackingOutputWithContext(ctx context.Context) FlowtrackingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowtrackingOutput)
}

// FlowtrackingArrayInput is an input type that accepts FlowtrackingArray and FlowtrackingArrayOutput values.
// You can construct a concrete instance of `FlowtrackingArrayInput` via:
//
//	FlowtrackingArray{ FlowtrackingArgs{...} }
type FlowtrackingArrayInput interface {
	pulumi.Input

	ToFlowtrackingArrayOutput() FlowtrackingArrayOutput
	ToFlowtrackingArrayOutputWithContext(context.Context) FlowtrackingArrayOutput
}

type FlowtrackingArray []FlowtrackingInput

func (FlowtrackingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Flowtracking)(nil)).Elem()
}

func (i FlowtrackingArray) ToFlowtrackingArrayOutput() FlowtrackingArrayOutput {
	return i.ToFlowtrackingArrayOutputWithContext(context.Background())
}

func (i FlowtrackingArray) ToFlowtrackingArrayOutputWithContext(ctx context.Context) FlowtrackingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowtrackingArrayOutput)
}

// FlowtrackingMapInput is an input type that accepts FlowtrackingMap and FlowtrackingMapOutput values.
// You can construct a concrete instance of `FlowtrackingMapInput` via:
//
//	FlowtrackingMap{ "key": FlowtrackingArgs{...} }
type FlowtrackingMapInput interface {
	pulumi.Input

	ToFlowtrackingMapOutput() FlowtrackingMapOutput
	ToFlowtrackingMapOutputWithContext(context.Context) FlowtrackingMapOutput
}

type FlowtrackingMap map[string]FlowtrackingInput

func (FlowtrackingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Flowtracking)(nil)).Elem()
}

func (i FlowtrackingMap) ToFlowtrackingMapOutput() FlowtrackingMapOutput {
	return i.ToFlowtrackingMapOutputWithContext(context.Background())
}

func (i FlowtrackingMap) ToFlowtrackingMapOutputWithContext(ctx context.Context) FlowtrackingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowtrackingMapOutput)
}

type FlowtrackingOutput struct{ *pulumi.OutputState }

func (FlowtrackingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Flowtracking)(nil)).Elem()
}

func (o FlowtrackingOutput) ToFlowtrackingOutput() FlowtrackingOutput {
	return o
}

func (o FlowtrackingOutput) ToFlowtrackingOutputWithContext(ctx context.Context) FlowtrackingOutput {
	return o
}

// Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
func (o FlowtrackingOutput) Aggregates() FlowtrackingAggregateArrayOutput {
	return o.ApplyT(func(v *Flowtracking) FlowtrackingAggregateArrayOutput { return v.Aggregates }).(FlowtrackingAggregateArrayOutput)
}

// Configure collector ip address.
func (o FlowtrackingOutput) CollectorIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.StringOutput { return v.CollectorIp }).(pulumi.StringOutput)
}

// Configure collector port number(0-65535, default=0).
func (o FlowtrackingOutput) CollectorPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.IntOutput { return v.CollectorPort }).(pulumi.IntOutput)
}

// Configure collectors for the flow. The structure of `collectors` block is documented below.
func (o FlowtrackingOutput) Collectors() FlowtrackingCollectorArrayOutput {
	return o.ApplyT(func(v *Flowtracking) FlowtrackingCollectorArrayOutput { return v.Collectors }).(FlowtrackingCollectorArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o FlowtrackingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
func (o FlowtrackingOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o FlowtrackingOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
func (o FlowtrackingOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.StringOutput { return v.Level }).(pulumi.StringOutput)
}

// Configure flow max export packet size (512-9216, default=512 bytes).
func (o FlowtrackingOutput) MaxExportPktSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.IntOutput { return v.MaxExportPktSize }).(pulumi.IntOutput)
}

// Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
func (o FlowtrackingOutput) SampleMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.StringOutput { return v.SampleMode }).(pulumi.StringOutput)
}

// Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
func (o FlowtrackingOutput) SampleRate() pulumi.IntOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.IntOutput { return v.SampleRate }).(pulumi.IntOutput)
}

// Configure template export period (1-60, default=5 minutes).
func (o FlowtrackingOutput) TemplateExportPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.IntOutput { return v.TemplateExportPeriod }).(pulumi.IntOutput)
}

// Configure flow session general timeout (60-604800, default=3600 seconds).
func (o FlowtrackingOutput) TimeoutGeneral() pulumi.IntOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.IntOutput { return v.TimeoutGeneral }).(pulumi.IntOutput)
}

// Configure flow session ICMP timeout (60-604800, default=300 seconds).
func (o FlowtrackingOutput) TimeoutIcmp() pulumi.IntOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.IntOutput { return v.TimeoutIcmp }).(pulumi.IntOutput)
}

// Configure flow session max timeout (60-604800, default=604800 seconds).
func (o FlowtrackingOutput) TimeoutMax() pulumi.IntOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.IntOutput { return v.TimeoutMax }).(pulumi.IntOutput)
}

// Configure flow session TCP timeout (60-604800, default=3600 seconds).
func (o FlowtrackingOutput) TimeoutTcp() pulumi.IntOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.IntOutput { return v.TimeoutTcp }).(pulumi.IntOutput)
}

// Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
func (o FlowtrackingOutput) TimeoutTcpFin() pulumi.IntOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.IntOutput { return v.TimeoutTcpFin }).(pulumi.IntOutput)
}

// Configure flow session TCP RST timeout (60-604800, default=120 seconds).
func (o FlowtrackingOutput) TimeoutTcpRst() pulumi.IntOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.IntOutput { return v.TimeoutTcpRst }).(pulumi.IntOutput)
}

// Configure flow session UDP timeout (60-604800, default=300 seconds).
func (o FlowtrackingOutput) TimeoutUdp() pulumi.IntOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.IntOutput { return v.TimeoutUdp }).(pulumi.IntOutput)
}

// Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
func (o FlowtrackingOutput) Transport() pulumi.StringOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.StringOutput { return v.Transport }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FlowtrackingOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Flowtracking) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type FlowtrackingArrayOutput struct{ *pulumi.OutputState }

func (FlowtrackingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Flowtracking)(nil)).Elem()
}

func (o FlowtrackingArrayOutput) ToFlowtrackingArrayOutput() FlowtrackingArrayOutput {
	return o
}

func (o FlowtrackingArrayOutput) ToFlowtrackingArrayOutputWithContext(ctx context.Context) FlowtrackingArrayOutput {
	return o
}

func (o FlowtrackingArrayOutput) Index(i pulumi.IntInput) FlowtrackingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Flowtracking {
		return vs[0].([]*Flowtracking)[vs[1].(int)]
	}).(FlowtrackingOutput)
}

type FlowtrackingMapOutput struct{ *pulumi.OutputState }

func (FlowtrackingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Flowtracking)(nil)).Elem()
}

func (o FlowtrackingMapOutput) ToFlowtrackingMapOutput() FlowtrackingMapOutput {
	return o
}

func (o FlowtrackingMapOutput) ToFlowtrackingMapOutputWithContext(ctx context.Context) FlowtrackingMapOutput {
	return o
}

func (o FlowtrackingMapOutput) MapIndex(k pulumi.StringInput) FlowtrackingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Flowtracking {
		return vs[0].(map[string]*Flowtracking)[vs[1].(string)]
	}).(FlowtrackingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowtrackingInput)(nil)).Elem(), &Flowtracking{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowtrackingArrayInput)(nil)).Elem(), FlowtrackingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowtrackingMapInput)(nil)).Elem(), FlowtrackingMap{})
	pulumi.RegisterOutputType(FlowtrackingOutput{})
	pulumi.RegisterOutputType(FlowtrackingArrayOutput{})
	pulumi.RegisterOutputType(FlowtrackingMapOutput{})
}
