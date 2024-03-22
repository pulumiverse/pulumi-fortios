// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure system-wide switch controller settings.
//
// ## Import
//
// SwitchController System can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/system:System labelname SwitchControllerSystem
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/system:System labelname SwitchControllerSystem
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type System struct {
	pulumi.CustomResourceState

	// Echo interval for the caputp echo requests from swtp.
	CaputpEchoInterval pulumi.IntOutput `pulumi:"caputpEchoInterval"`
	// Maximum retransmission count for the caputp tunnel packets.
	CaputpMaxRetransmit pulumi.IntOutput `pulumi:"caputpMaxRetransmit"`
	// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
	DataSyncInterval pulumi.IntOutput `pulumi:"dataSyncInterval"`
	// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
	DynamicPeriodicInterval pulumi.IntOutput `pulumi:"dynamicPeriodicInterval"`
	// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
	IotHoldoff pulumi.IntOutput `pulumi:"iotHoldoff"`
	// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
	IotMacIdle pulumi.IntOutput `pulumi:"iotMacIdle"`
	// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
	IotScanInterval pulumi.IntOutput `pulumi:"iotScanInterval"`
	// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
	IotWeightThreshold pulumi.IntOutput `pulumi:"iotWeightThreshold"`
	// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
	NacPeriodicInterval pulumi.IntOutput `pulumi:"nacPeriodicInterval"`
	// Maximum number of parallel processes (1 - 300, default = 1).
	ParallelProcess pulumi.IntOutput `pulumi:"parallelProcess"`
	// Enable/disable parallel process override. Valid values: `disable`, `enable`.
	ParallelProcessOverride pulumi.StringOutput `pulumi:"parallelProcessOverride"`
	// Compatible/strict tunnel mode.
	TunnelMode pulumi.StringOutput `pulumi:"tunnelMode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystem registers a new resource with the given unique name, arguments, and options.
func NewSystem(ctx *pulumi.Context,
	name string, args *SystemArgs, opts ...pulumi.ResourceOption) (*System, error) {
	if args == nil {
		args = &SystemArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource System
	err := ctx.RegisterResource("fortios:switchcontroller/system:System", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystem gets an existing System resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemState, opts ...pulumi.ResourceOption) (*System, error) {
	var resource System
	err := ctx.ReadResource("fortios:switchcontroller/system:System", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering System resources.
type systemState struct {
	// Echo interval for the caputp echo requests from swtp.
	CaputpEchoInterval *int `pulumi:"caputpEchoInterval"`
	// Maximum retransmission count for the caputp tunnel packets.
	CaputpMaxRetransmit *int `pulumi:"caputpMaxRetransmit"`
	// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
	DataSyncInterval *int `pulumi:"dataSyncInterval"`
	// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
	DynamicPeriodicInterval *int `pulumi:"dynamicPeriodicInterval"`
	// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
	IotHoldoff *int `pulumi:"iotHoldoff"`
	// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
	IotMacIdle *int `pulumi:"iotMacIdle"`
	// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
	IotScanInterval *int `pulumi:"iotScanInterval"`
	// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
	IotWeightThreshold *int `pulumi:"iotWeightThreshold"`
	// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
	NacPeriodicInterval *int `pulumi:"nacPeriodicInterval"`
	// Maximum number of parallel processes (1 - 300, default = 1).
	ParallelProcess *int `pulumi:"parallelProcess"`
	// Enable/disable parallel process override. Valid values: `disable`, `enable`.
	ParallelProcessOverride *string `pulumi:"parallelProcessOverride"`
	// Compatible/strict tunnel mode.
	TunnelMode *string `pulumi:"tunnelMode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemState struct {
	// Echo interval for the caputp echo requests from swtp.
	CaputpEchoInterval pulumi.IntPtrInput
	// Maximum retransmission count for the caputp tunnel packets.
	CaputpMaxRetransmit pulumi.IntPtrInput
	// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
	DataSyncInterval pulumi.IntPtrInput
	// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
	DynamicPeriodicInterval pulumi.IntPtrInput
	// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
	IotHoldoff pulumi.IntPtrInput
	// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
	IotMacIdle pulumi.IntPtrInput
	// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
	IotScanInterval pulumi.IntPtrInput
	// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
	IotWeightThreshold pulumi.IntPtrInput
	// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
	NacPeriodicInterval pulumi.IntPtrInput
	// Maximum number of parallel processes (1 - 300, default = 1).
	ParallelProcess pulumi.IntPtrInput
	// Enable/disable parallel process override. Valid values: `disable`, `enable`.
	ParallelProcessOverride pulumi.StringPtrInput
	// Compatible/strict tunnel mode.
	TunnelMode pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemState)(nil)).Elem()
}

type systemArgs struct {
	// Echo interval for the caputp echo requests from swtp.
	CaputpEchoInterval *int `pulumi:"caputpEchoInterval"`
	// Maximum retransmission count for the caputp tunnel packets.
	CaputpMaxRetransmit *int `pulumi:"caputpMaxRetransmit"`
	// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
	DataSyncInterval *int `pulumi:"dataSyncInterval"`
	// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
	DynamicPeriodicInterval *int `pulumi:"dynamicPeriodicInterval"`
	// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
	IotHoldoff *int `pulumi:"iotHoldoff"`
	// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
	IotMacIdle *int `pulumi:"iotMacIdle"`
	// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
	IotScanInterval *int `pulumi:"iotScanInterval"`
	// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
	IotWeightThreshold *int `pulumi:"iotWeightThreshold"`
	// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
	NacPeriodicInterval *int `pulumi:"nacPeriodicInterval"`
	// Maximum number of parallel processes (1 - 300, default = 1).
	ParallelProcess *int `pulumi:"parallelProcess"`
	// Enable/disable parallel process override. Valid values: `disable`, `enable`.
	ParallelProcessOverride *string `pulumi:"parallelProcessOverride"`
	// Compatible/strict tunnel mode.
	TunnelMode *string `pulumi:"tunnelMode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a System resource.
type SystemArgs struct {
	// Echo interval for the caputp echo requests from swtp.
	CaputpEchoInterval pulumi.IntPtrInput
	// Maximum retransmission count for the caputp tunnel packets.
	CaputpMaxRetransmit pulumi.IntPtrInput
	// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
	DataSyncInterval pulumi.IntPtrInput
	// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
	DynamicPeriodicInterval pulumi.IntPtrInput
	// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
	IotHoldoff pulumi.IntPtrInput
	// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
	IotMacIdle pulumi.IntPtrInput
	// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
	IotScanInterval pulumi.IntPtrInput
	// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
	IotWeightThreshold pulumi.IntPtrInput
	// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
	NacPeriodicInterval pulumi.IntPtrInput
	// Maximum number of parallel processes (1 - 300, default = 1).
	ParallelProcess pulumi.IntPtrInput
	// Enable/disable parallel process override. Valid values: `disable`, `enable`.
	ParallelProcessOverride pulumi.StringPtrInput
	// Compatible/strict tunnel mode.
	TunnelMode pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemArgs)(nil)).Elem()
}

type SystemInput interface {
	pulumi.Input

	ToSystemOutput() SystemOutput
	ToSystemOutputWithContext(ctx context.Context) SystemOutput
}

func (*System) ElementType() reflect.Type {
	return reflect.TypeOf((**System)(nil)).Elem()
}

func (i *System) ToSystemOutput() SystemOutput {
	return i.ToSystemOutputWithContext(context.Background())
}

func (i *System) ToSystemOutputWithContext(ctx context.Context) SystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemOutput)
}

// SystemArrayInput is an input type that accepts SystemArray and SystemArrayOutput values.
// You can construct a concrete instance of `SystemArrayInput` via:
//
//	SystemArray{ SystemArgs{...} }
type SystemArrayInput interface {
	pulumi.Input

	ToSystemArrayOutput() SystemArrayOutput
	ToSystemArrayOutputWithContext(context.Context) SystemArrayOutput
}

type SystemArray []SystemInput

func (SystemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*System)(nil)).Elem()
}

func (i SystemArray) ToSystemArrayOutput() SystemArrayOutput {
	return i.ToSystemArrayOutputWithContext(context.Background())
}

func (i SystemArray) ToSystemArrayOutputWithContext(ctx context.Context) SystemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemArrayOutput)
}

// SystemMapInput is an input type that accepts SystemMap and SystemMapOutput values.
// You can construct a concrete instance of `SystemMapInput` via:
//
//	SystemMap{ "key": SystemArgs{...} }
type SystemMapInput interface {
	pulumi.Input

	ToSystemMapOutput() SystemMapOutput
	ToSystemMapOutputWithContext(context.Context) SystemMapOutput
}

type SystemMap map[string]SystemInput

func (SystemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*System)(nil)).Elem()
}

func (i SystemMap) ToSystemMapOutput() SystemMapOutput {
	return i.ToSystemMapOutputWithContext(context.Background())
}

func (i SystemMap) ToSystemMapOutputWithContext(ctx context.Context) SystemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemMapOutput)
}

type SystemOutput struct{ *pulumi.OutputState }

func (SystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**System)(nil)).Elem()
}

func (o SystemOutput) ToSystemOutput() SystemOutput {
	return o
}

func (o SystemOutput) ToSystemOutputWithContext(ctx context.Context) SystemOutput {
	return o
}

// Echo interval for the caputp echo requests from swtp.
func (o SystemOutput) CaputpEchoInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *System) pulumi.IntOutput { return v.CaputpEchoInterval }).(pulumi.IntOutput)
}

// Maximum retransmission count for the caputp tunnel packets.
func (o SystemOutput) CaputpMaxRetransmit() pulumi.IntOutput {
	return o.ApplyT(func(v *System) pulumi.IntOutput { return v.CaputpMaxRetransmit }).(pulumi.IntOutput)
}

// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
func (o SystemOutput) DataSyncInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *System) pulumi.IntOutput { return v.DataSyncInterval }).(pulumi.IntOutput)
}

// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
func (o SystemOutput) DynamicPeriodicInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *System) pulumi.IntOutput { return v.DynamicPeriodicInterval }).(pulumi.IntOutput)
}

// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
func (o SystemOutput) IotHoldoff() pulumi.IntOutput {
	return o.ApplyT(func(v *System) pulumi.IntOutput { return v.IotHoldoff }).(pulumi.IntOutput)
}

// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
func (o SystemOutput) IotMacIdle() pulumi.IntOutput {
	return o.ApplyT(func(v *System) pulumi.IntOutput { return v.IotMacIdle }).(pulumi.IntOutput)
}

// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
func (o SystemOutput) IotScanInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *System) pulumi.IntOutput { return v.IotScanInterval }).(pulumi.IntOutput)
}

// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
func (o SystemOutput) IotWeightThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *System) pulumi.IntOutput { return v.IotWeightThreshold }).(pulumi.IntOutput)
}

// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
func (o SystemOutput) NacPeriodicInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *System) pulumi.IntOutput { return v.NacPeriodicInterval }).(pulumi.IntOutput)
}

// Maximum number of parallel processes (1 - 300, default = 1).
func (o SystemOutput) ParallelProcess() pulumi.IntOutput {
	return o.ApplyT(func(v *System) pulumi.IntOutput { return v.ParallelProcess }).(pulumi.IntOutput)
}

// Enable/disable parallel process override. Valid values: `disable`, `enable`.
func (o SystemOutput) ParallelProcessOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *System) pulumi.StringOutput { return v.ParallelProcessOverride }).(pulumi.StringOutput)
}

// Compatible/strict tunnel mode.
func (o SystemOutput) TunnelMode() pulumi.StringOutput {
	return o.ApplyT(func(v *System) pulumi.StringOutput { return v.TunnelMode }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *System) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemArrayOutput struct{ *pulumi.OutputState }

func (SystemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*System)(nil)).Elem()
}

func (o SystemArrayOutput) ToSystemArrayOutput() SystemArrayOutput {
	return o
}

func (o SystemArrayOutput) ToSystemArrayOutputWithContext(ctx context.Context) SystemArrayOutput {
	return o
}

func (o SystemArrayOutput) Index(i pulumi.IntInput) SystemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *System {
		return vs[0].([]*System)[vs[1].(int)]
	}).(SystemOutput)
}

type SystemMapOutput struct{ *pulumi.OutputState }

func (SystemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*System)(nil)).Elem()
}

func (o SystemMapOutput) ToSystemMapOutput() SystemMapOutput {
	return o
}

func (o SystemMapOutput) ToSystemMapOutputWithContext(ctx context.Context) SystemMapOutput {
	return o
}

func (o SystemMapOutput) MapIndex(k pulumi.StringInput) SystemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *System {
		return vs[0].(map[string]*System)[vs[1].(string)]
	}).(SystemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemInput)(nil)).Elem(), &System{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemArrayInput)(nil)).Elem(), SystemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemMapInput)(nil)).Elem(), SystemMap{})
	pulumi.RegisterOutputType(SystemOutput{})
	pulumi.RegisterOutputType(SystemArrayOutput{})
	pulumi.RegisterOutputType(SystemMapOutput{})
}
