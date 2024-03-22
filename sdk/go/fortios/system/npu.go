// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure NPU attributes. Applies to FortiOS Version `7.0.4`.
//
// ## Import
//
// System Npu can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/npu:Npu labelname SystemNpu
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/npu:Npu labelname SystemNpu
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Npu struct {
	pulumi.CustomResourceState

	// Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
	CapwapOffload pulumi.StringOutput `pulumi:"capwapOffload"`
	// Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
	DedicatedManagementAffinity pulumi.StringOutput `pulumi:"dedicatedManagementAffinity"`
	// Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
	DedicatedManagementCpu pulumi.StringOutput `pulumi:"dedicatedManagementCpu"`
	// Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
	Fastpath pulumi.StringOutput `pulumi:"fastpath"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
	IpsecDecSubengineMask pulumi.StringOutput `pulumi:"ipsecDecSubengineMask"`
	// IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
	IpsecEncSubengineMask pulumi.StringOutput `pulumi:"ipsecEncSubengineMask"`
	// Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
	IpsecInboundCache pulumi.StringOutput `pulumi:"ipsecInboundCache"`
	// Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
	IpsecMtuOverride pulumi.StringOutput `pulumi:"ipsecMtuOverride"`
	// Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
	IpsecOverVlink pulumi.StringOutput `pulumi:"ipsecOverVlink"`
	// Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
	McastSessionAccounting pulumi.StringOutput `pulumi:"mcastSessionAccounting"`
	// Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
	Np6CpsOptimizationMode pulumi.StringOutput `pulumi:"np6CpsOptimizationMode"`
	// Configure NPU priority protocol. The structure of `priorityProtocol` block is documented below.
	PriorityProtocol NpuPriorityProtocolOutput `pulumi:"priorityProtocol"`
	// Enable/disable rdp offload. Valid values: `enable`, `disable`.
	RdpOffload pulumi.StringOutput `pulumi:"rdpOffload"`
	// Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
	SessionDeniedOffload pulumi.StringOutput `pulumi:"sessionDeniedOffload"`
	// Enable/disable sse backpressure. Valid values: `enable`, `disable`.
	SseBackpressure pulumi.StringOutput `pulumi:"sseBackpressure"`
	// Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
	StripClearTextPadding pulumi.StringOutput `pulumi:"stripClearTextPadding"`
	// Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
	StripEspPadding pulumi.StringOutput `pulumi:"stripEspPadding"`
	// Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
	SwNpBandwidth pulumi.StringOutput `pulumi:"swNpBandwidth"`
	// Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
	UespOffload pulumi.StringOutput `pulumi:"uespOffload"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewNpu registers a new resource with the given unique name, arguments, and options.
func NewNpu(ctx *pulumi.Context,
	name string, args *NpuArgs, opts ...pulumi.ResourceOption) (*Npu, error) {
	if args == nil {
		args = &NpuArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Npu
	err := ctx.RegisterResource("fortios:system/npu:Npu", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNpu gets an existing Npu resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNpu(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NpuState, opts ...pulumi.ResourceOption) (*Npu, error) {
	var resource Npu
	err := ctx.ReadResource("fortios:system/npu:Npu", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Npu resources.
type npuState struct {
	// Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
	CapwapOffload *string `pulumi:"capwapOffload"`
	// Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
	DedicatedManagementAffinity *string `pulumi:"dedicatedManagementAffinity"`
	// Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
	DedicatedManagementCpu *string `pulumi:"dedicatedManagementCpu"`
	// Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
	Fastpath *string `pulumi:"fastpath"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
	IpsecDecSubengineMask *string `pulumi:"ipsecDecSubengineMask"`
	// IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
	IpsecEncSubengineMask *string `pulumi:"ipsecEncSubengineMask"`
	// Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
	IpsecInboundCache *string `pulumi:"ipsecInboundCache"`
	// Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
	IpsecMtuOverride *string `pulumi:"ipsecMtuOverride"`
	// Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
	IpsecOverVlink *string `pulumi:"ipsecOverVlink"`
	// Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
	McastSessionAccounting *string `pulumi:"mcastSessionAccounting"`
	// Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
	Np6CpsOptimizationMode *string `pulumi:"np6CpsOptimizationMode"`
	// Configure NPU priority protocol. The structure of `priorityProtocol` block is documented below.
	PriorityProtocol *NpuPriorityProtocol `pulumi:"priorityProtocol"`
	// Enable/disable rdp offload. Valid values: `enable`, `disable`.
	RdpOffload *string `pulumi:"rdpOffload"`
	// Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
	SessionDeniedOffload *string `pulumi:"sessionDeniedOffload"`
	// Enable/disable sse backpressure. Valid values: `enable`, `disable`.
	SseBackpressure *string `pulumi:"sseBackpressure"`
	// Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
	StripClearTextPadding *string `pulumi:"stripClearTextPadding"`
	// Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
	StripEspPadding *string `pulumi:"stripEspPadding"`
	// Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
	SwNpBandwidth *string `pulumi:"swNpBandwidth"`
	// Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
	UespOffload *string `pulumi:"uespOffload"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type NpuState struct {
	// Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
	CapwapOffload pulumi.StringPtrInput
	// Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
	DedicatedManagementAffinity pulumi.StringPtrInput
	// Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
	DedicatedManagementCpu pulumi.StringPtrInput
	// Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
	Fastpath pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
	IpsecDecSubengineMask pulumi.StringPtrInput
	// IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
	IpsecEncSubengineMask pulumi.StringPtrInput
	// Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
	IpsecInboundCache pulumi.StringPtrInput
	// Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
	IpsecMtuOverride pulumi.StringPtrInput
	// Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
	IpsecOverVlink pulumi.StringPtrInput
	// Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
	McastSessionAccounting pulumi.StringPtrInput
	// Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
	Np6CpsOptimizationMode pulumi.StringPtrInput
	// Configure NPU priority protocol. The structure of `priorityProtocol` block is documented below.
	PriorityProtocol NpuPriorityProtocolPtrInput
	// Enable/disable rdp offload. Valid values: `enable`, `disable`.
	RdpOffload pulumi.StringPtrInput
	// Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
	SessionDeniedOffload pulumi.StringPtrInput
	// Enable/disable sse backpressure. Valid values: `enable`, `disable`.
	SseBackpressure pulumi.StringPtrInput
	// Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
	StripClearTextPadding pulumi.StringPtrInput
	// Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
	StripEspPadding pulumi.StringPtrInput
	// Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
	SwNpBandwidth pulumi.StringPtrInput
	// Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
	UespOffload pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NpuState) ElementType() reflect.Type {
	return reflect.TypeOf((*npuState)(nil)).Elem()
}

type npuArgs struct {
	// Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
	CapwapOffload *string `pulumi:"capwapOffload"`
	// Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
	DedicatedManagementAffinity *string `pulumi:"dedicatedManagementAffinity"`
	// Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
	DedicatedManagementCpu *string `pulumi:"dedicatedManagementCpu"`
	// Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
	Fastpath *string `pulumi:"fastpath"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
	IpsecDecSubengineMask *string `pulumi:"ipsecDecSubengineMask"`
	// IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
	IpsecEncSubengineMask *string `pulumi:"ipsecEncSubengineMask"`
	// Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
	IpsecInboundCache *string `pulumi:"ipsecInboundCache"`
	// Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
	IpsecMtuOverride *string `pulumi:"ipsecMtuOverride"`
	// Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
	IpsecOverVlink *string `pulumi:"ipsecOverVlink"`
	// Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
	McastSessionAccounting *string `pulumi:"mcastSessionAccounting"`
	// Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
	Np6CpsOptimizationMode *string `pulumi:"np6CpsOptimizationMode"`
	// Configure NPU priority protocol. The structure of `priorityProtocol` block is documented below.
	PriorityProtocol *NpuPriorityProtocol `pulumi:"priorityProtocol"`
	// Enable/disable rdp offload. Valid values: `enable`, `disable`.
	RdpOffload *string `pulumi:"rdpOffload"`
	// Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
	SessionDeniedOffload *string `pulumi:"sessionDeniedOffload"`
	// Enable/disable sse backpressure. Valid values: `enable`, `disable`.
	SseBackpressure *string `pulumi:"sseBackpressure"`
	// Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
	StripClearTextPadding *string `pulumi:"stripClearTextPadding"`
	// Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
	StripEspPadding *string `pulumi:"stripEspPadding"`
	// Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
	SwNpBandwidth *string `pulumi:"swNpBandwidth"`
	// Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
	UespOffload *string `pulumi:"uespOffload"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Npu resource.
type NpuArgs struct {
	// Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
	CapwapOffload pulumi.StringPtrInput
	// Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
	DedicatedManagementAffinity pulumi.StringPtrInput
	// Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
	DedicatedManagementCpu pulumi.StringPtrInput
	// Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
	Fastpath pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
	IpsecDecSubengineMask pulumi.StringPtrInput
	// IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
	IpsecEncSubengineMask pulumi.StringPtrInput
	// Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
	IpsecInboundCache pulumi.StringPtrInput
	// Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
	IpsecMtuOverride pulumi.StringPtrInput
	// Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
	IpsecOverVlink pulumi.StringPtrInput
	// Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
	McastSessionAccounting pulumi.StringPtrInput
	// Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
	Np6CpsOptimizationMode pulumi.StringPtrInput
	// Configure NPU priority protocol. The structure of `priorityProtocol` block is documented below.
	PriorityProtocol NpuPriorityProtocolPtrInput
	// Enable/disable rdp offload. Valid values: `enable`, `disable`.
	RdpOffload pulumi.StringPtrInput
	// Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
	SessionDeniedOffload pulumi.StringPtrInput
	// Enable/disable sse backpressure. Valid values: `enable`, `disable`.
	SseBackpressure pulumi.StringPtrInput
	// Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
	StripClearTextPadding pulumi.StringPtrInput
	// Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
	StripEspPadding pulumi.StringPtrInput
	// Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
	SwNpBandwidth pulumi.StringPtrInput
	// Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
	UespOffload pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NpuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*npuArgs)(nil)).Elem()
}

type NpuInput interface {
	pulumi.Input

	ToNpuOutput() NpuOutput
	ToNpuOutputWithContext(ctx context.Context) NpuOutput
}

func (*Npu) ElementType() reflect.Type {
	return reflect.TypeOf((**Npu)(nil)).Elem()
}

func (i *Npu) ToNpuOutput() NpuOutput {
	return i.ToNpuOutputWithContext(context.Background())
}

func (i *Npu) ToNpuOutputWithContext(ctx context.Context) NpuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NpuOutput)
}

// NpuArrayInput is an input type that accepts NpuArray and NpuArrayOutput values.
// You can construct a concrete instance of `NpuArrayInput` via:
//
//	NpuArray{ NpuArgs{...} }
type NpuArrayInput interface {
	pulumi.Input

	ToNpuArrayOutput() NpuArrayOutput
	ToNpuArrayOutputWithContext(context.Context) NpuArrayOutput
}

type NpuArray []NpuInput

func (NpuArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Npu)(nil)).Elem()
}

func (i NpuArray) ToNpuArrayOutput() NpuArrayOutput {
	return i.ToNpuArrayOutputWithContext(context.Background())
}

func (i NpuArray) ToNpuArrayOutputWithContext(ctx context.Context) NpuArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NpuArrayOutput)
}

// NpuMapInput is an input type that accepts NpuMap and NpuMapOutput values.
// You can construct a concrete instance of `NpuMapInput` via:
//
//	NpuMap{ "key": NpuArgs{...} }
type NpuMapInput interface {
	pulumi.Input

	ToNpuMapOutput() NpuMapOutput
	ToNpuMapOutputWithContext(context.Context) NpuMapOutput
}

type NpuMap map[string]NpuInput

func (NpuMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Npu)(nil)).Elem()
}

func (i NpuMap) ToNpuMapOutput() NpuMapOutput {
	return i.ToNpuMapOutputWithContext(context.Background())
}

func (i NpuMap) ToNpuMapOutputWithContext(ctx context.Context) NpuMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NpuMapOutput)
}

type NpuOutput struct{ *pulumi.OutputState }

func (NpuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Npu)(nil)).Elem()
}

func (o NpuOutput) ToNpuOutput() NpuOutput {
	return o
}

func (o NpuOutput) ToNpuOutputWithContext(ctx context.Context) NpuOutput {
	return o
}

// Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
func (o NpuOutput) CapwapOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.CapwapOffload }).(pulumi.StringOutput)
}

// Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
func (o NpuOutput) DedicatedManagementAffinity() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.DedicatedManagementAffinity }).(pulumi.StringOutput)
}

// Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
func (o NpuOutput) DedicatedManagementCpu() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.DedicatedManagementCpu }).(pulumi.StringOutput)
}

// Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
func (o NpuOutput) Fastpath() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.Fastpath }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o NpuOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
func (o NpuOutput) IpsecDecSubengineMask() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.IpsecDecSubengineMask }).(pulumi.StringOutput)
}

// IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
func (o NpuOutput) IpsecEncSubengineMask() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.IpsecEncSubengineMask }).(pulumi.StringOutput)
}

// Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
func (o NpuOutput) IpsecInboundCache() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.IpsecInboundCache }).(pulumi.StringOutput)
}

// Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
func (o NpuOutput) IpsecMtuOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.IpsecMtuOverride }).(pulumi.StringOutput)
}

// Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
func (o NpuOutput) IpsecOverVlink() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.IpsecOverVlink }).(pulumi.StringOutput)
}

// Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
func (o NpuOutput) McastSessionAccounting() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.McastSessionAccounting }).(pulumi.StringOutput)
}

// Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
func (o NpuOutput) Np6CpsOptimizationMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.Np6CpsOptimizationMode }).(pulumi.StringOutput)
}

// Configure NPU priority protocol. The structure of `priorityProtocol` block is documented below.
func (o NpuOutput) PriorityProtocol() NpuPriorityProtocolOutput {
	return o.ApplyT(func(v *Npu) NpuPriorityProtocolOutput { return v.PriorityProtocol }).(NpuPriorityProtocolOutput)
}

// Enable/disable rdp offload. Valid values: `enable`, `disable`.
func (o NpuOutput) RdpOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.RdpOffload }).(pulumi.StringOutput)
}

// Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
func (o NpuOutput) SessionDeniedOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.SessionDeniedOffload }).(pulumi.StringOutput)
}

// Enable/disable sse backpressure. Valid values: `enable`, `disable`.
func (o NpuOutput) SseBackpressure() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.SseBackpressure }).(pulumi.StringOutput)
}

// Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
func (o NpuOutput) StripClearTextPadding() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.StripClearTextPadding }).(pulumi.StringOutput)
}

// Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
func (o NpuOutput) StripEspPadding() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.StripEspPadding }).(pulumi.StringOutput)
}

// Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
func (o NpuOutput) SwNpBandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.SwNpBandwidth }).(pulumi.StringOutput)
}

// Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
func (o NpuOutput) UespOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringOutput { return v.UespOffload }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o NpuOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Npu) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type NpuArrayOutput struct{ *pulumi.OutputState }

func (NpuArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Npu)(nil)).Elem()
}

func (o NpuArrayOutput) ToNpuArrayOutput() NpuArrayOutput {
	return o
}

func (o NpuArrayOutput) ToNpuArrayOutputWithContext(ctx context.Context) NpuArrayOutput {
	return o
}

func (o NpuArrayOutput) Index(i pulumi.IntInput) NpuOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Npu {
		return vs[0].([]*Npu)[vs[1].(int)]
	}).(NpuOutput)
}

type NpuMapOutput struct{ *pulumi.OutputState }

func (NpuMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Npu)(nil)).Elem()
}

func (o NpuMapOutput) ToNpuMapOutput() NpuMapOutput {
	return o
}

func (o NpuMapOutput) ToNpuMapOutputWithContext(ctx context.Context) NpuMapOutput {
	return o
}

func (o NpuMapOutput) MapIndex(k pulumi.StringInput) NpuOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Npu {
		return vs[0].(map[string]*Npu)[vs[1].(string)]
	}).(NpuOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NpuInput)(nil)).Elem(), &Npu{})
	pulumi.RegisterInputType(reflect.TypeOf((*NpuArrayInput)(nil)).Elem(), NpuArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NpuMapInput)(nil)).Elem(), NpuMap{})
	pulumi.RegisterOutputType(NpuOutput{})
	pulumi.RegisterOutputType(NpuArrayOutput{})
	pulumi.RegisterOutputType(NpuMapOutput{})
}
