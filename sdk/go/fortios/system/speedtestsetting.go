// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure speed test setting. Applies to FortiOS Version `7.2.6,7.4.1,7.4.2`.
//
// ## Import
//
// System SpeedTestSetting can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/speedtestsetting:Speedtestsetting labelname SystemSpeedTestSetting
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/speedtestsetting:Speedtestsetting labelname SystemSpeedTestSetting
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Speedtestsetting struct {
	pulumi.CustomResourceState

	// Speed test latency threshold in milliseconds (0 - 2000, default = 60) for the Auto mode. If the latency exceeds this threshold, the speed test will use the UDP protocol; otherwise, it will use the TCP protocol.
	LatencyThreshold pulumi.IntOutput `pulumi:"latencyThreshold"`
	// Number of parallel client streams (1 - 64, default = 4) for the TCP protocol to run during the speed test.
	MultipleTcpStream pulumi.IntOutput `pulumi:"multipleTcpStream"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSpeedtestsetting registers a new resource with the given unique name, arguments, and options.
func NewSpeedtestsetting(ctx *pulumi.Context,
	name string, args *SpeedtestsettingArgs, opts ...pulumi.ResourceOption) (*Speedtestsetting, error) {
	if args == nil {
		args = &SpeedtestsettingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Speedtestsetting
	err := ctx.RegisterResource("fortios:system/speedtestsetting:Speedtestsetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpeedtestsetting gets an existing Speedtestsetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpeedtestsetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpeedtestsettingState, opts ...pulumi.ResourceOption) (*Speedtestsetting, error) {
	var resource Speedtestsetting
	err := ctx.ReadResource("fortios:system/speedtestsetting:Speedtestsetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Speedtestsetting resources.
type speedtestsettingState struct {
	// Speed test latency threshold in milliseconds (0 - 2000, default = 60) for the Auto mode. If the latency exceeds this threshold, the speed test will use the UDP protocol; otherwise, it will use the TCP protocol.
	LatencyThreshold *int `pulumi:"latencyThreshold"`
	// Number of parallel client streams (1 - 64, default = 4) for the TCP protocol to run during the speed test.
	MultipleTcpStream *int `pulumi:"multipleTcpStream"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SpeedtestsettingState struct {
	// Speed test latency threshold in milliseconds (0 - 2000, default = 60) for the Auto mode. If the latency exceeds this threshold, the speed test will use the UDP protocol; otherwise, it will use the TCP protocol.
	LatencyThreshold pulumi.IntPtrInput
	// Number of parallel client streams (1 - 64, default = 4) for the TCP protocol to run during the speed test.
	MultipleTcpStream pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpeedtestsettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*speedtestsettingState)(nil)).Elem()
}

type speedtestsettingArgs struct {
	// Speed test latency threshold in milliseconds (0 - 2000, default = 60) for the Auto mode. If the latency exceeds this threshold, the speed test will use the UDP protocol; otherwise, it will use the TCP protocol.
	LatencyThreshold *int `pulumi:"latencyThreshold"`
	// Number of parallel client streams (1 - 64, default = 4) for the TCP protocol to run during the speed test.
	MultipleTcpStream *int `pulumi:"multipleTcpStream"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Speedtestsetting resource.
type SpeedtestsettingArgs struct {
	// Speed test latency threshold in milliseconds (0 - 2000, default = 60) for the Auto mode. If the latency exceeds this threshold, the speed test will use the UDP protocol; otherwise, it will use the TCP protocol.
	LatencyThreshold pulumi.IntPtrInput
	// Number of parallel client streams (1 - 64, default = 4) for the TCP protocol to run during the speed test.
	MultipleTcpStream pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpeedtestsettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*speedtestsettingArgs)(nil)).Elem()
}

type SpeedtestsettingInput interface {
	pulumi.Input

	ToSpeedtestsettingOutput() SpeedtestsettingOutput
	ToSpeedtestsettingOutputWithContext(ctx context.Context) SpeedtestsettingOutput
}

func (*Speedtestsetting) ElementType() reflect.Type {
	return reflect.TypeOf((**Speedtestsetting)(nil)).Elem()
}

func (i *Speedtestsetting) ToSpeedtestsettingOutput() SpeedtestsettingOutput {
	return i.ToSpeedtestsettingOutputWithContext(context.Background())
}

func (i *Speedtestsetting) ToSpeedtestsettingOutputWithContext(ctx context.Context) SpeedtestsettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpeedtestsettingOutput)
}

// SpeedtestsettingArrayInput is an input type that accepts SpeedtestsettingArray and SpeedtestsettingArrayOutput values.
// You can construct a concrete instance of `SpeedtestsettingArrayInput` via:
//
//	SpeedtestsettingArray{ SpeedtestsettingArgs{...} }
type SpeedtestsettingArrayInput interface {
	pulumi.Input

	ToSpeedtestsettingArrayOutput() SpeedtestsettingArrayOutput
	ToSpeedtestsettingArrayOutputWithContext(context.Context) SpeedtestsettingArrayOutput
}

type SpeedtestsettingArray []SpeedtestsettingInput

func (SpeedtestsettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Speedtestsetting)(nil)).Elem()
}

func (i SpeedtestsettingArray) ToSpeedtestsettingArrayOutput() SpeedtestsettingArrayOutput {
	return i.ToSpeedtestsettingArrayOutputWithContext(context.Background())
}

func (i SpeedtestsettingArray) ToSpeedtestsettingArrayOutputWithContext(ctx context.Context) SpeedtestsettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpeedtestsettingArrayOutput)
}

// SpeedtestsettingMapInput is an input type that accepts SpeedtestsettingMap and SpeedtestsettingMapOutput values.
// You can construct a concrete instance of `SpeedtestsettingMapInput` via:
//
//	SpeedtestsettingMap{ "key": SpeedtestsettingArgs{...} }
type SpeedtestsettingMapInput interface {
	pulumi.Input

	ToSpeedtestsettingMapOutput() SpeedtestsettingMapOutput
	ToSpeedtestsettingMapOutputWithContext(context.Context) SpeedtestsettingMapOutput
}

type SpeedtestsettingMap map[string]SpeedtestsettingInput

func (SpeedtestsettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Speedtestsetting)(nil)).Elem()
}

func (i SpeedtestsettingMap) ToSpeedtestsettingMapOutput() SpeedtestsettingMapOutput {
	return i.ToSpeedtestsettingMapOutputWithContext(context.Background())
}

func (i SpeedtestsettingMap) ToSpeedtestsettingMapOutputWithContext(ctx context.Context) SpeedtestsettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpeedtestsettingMapOutput)
}

type SpeedtestsettingOutput struct{ *pulumi.OutputState }

func (SpeedtestsettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Speedtestsetting)(nil)).Elem()
}

func (o SpeedtestsettingOutput) ToSpeedtestsettingOutput() SpeedtestsettingOutput {
	return o
}

func (o SpeedtestsettingOutput) ToSpeedtestsettingOutputWithContext(ctx context.Context) SpeedtestsettingOutput {
	return o
}

// Speed test latency threshold in milliseconds (0 - 2000, default = 60) for the Auto mode. If the latency exceeds this threshold, the speed test will use the UDP protocol; otherwise, it will use the TCP protocol.
func (o SpeedtestsettingOutput) LatencyThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Speedtestsetting) pulumi.IntOutput { return v.LatencyThreshold }).(pulumi.IntOutput)
}

// Number of parallel client streams (1 - 64, default = 4) for the TCP protocol to run during the speed test.
func (o SpeedtestsettingOutput) MultipleTcpStream() pulumi.IntOutput {
	return o.ApplyT(func(v *Speedtestsetting) pulumi.IntOutput { return v.MultipleTcpStream }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SpeedtestsettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Speedtestsetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SpeedtestsettingArrayOutput struct{ *pulumi.OutputState }

func (SpeedtestsettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Speedtestsetting)(nil)).Elem()
}

func (o SpeedtestsettingArrayOutput) ToSpeedtestsettingArrayOutput() SpeedtestsettingArrayOutput {
	return o
}

func (o SpeedtestsettingArrayOutput) ToSpeedtestsettingArrayOutputWithContext(ctx context.Context) SpeedtestsettingArrayOutput {
	return o
}

func (o SpeedtestsettingArrayOutput) Index(i pulumi.IntInput) SpeedtestsettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Speedtestsetting {
		return vs[0].([]*Speedtestsetting)[vs[1].(int)]
	}).(SpeedtestsettingOutput)
}

type SpeedtestsettingMapOutput struct{ *pulumi.OutputState }

func (SpeedtestsettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Speedtestsetting)(nil)).Elem()
}

func (o SpeedtestsettingMapOutput) ToSpeedtestsettingMapOutput() SpeedtestsettingMapOutput {
	return o
}

func (o SpeedtestsettingMapOutput) ToSpeedtestsettingMapOutputWithContext(ctx context.Context) SpeedtestsettingMapOutput {
	return o
}

func (o SpeedtestsettingMapOutput) MapIndex(k pulumi.StringInput) SpeedtestsettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Speedtestsetting {
		return vs[0].(map[string]*Speedtestsetting)[vs[1].(string)]
	}).(SpeedtestsettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpeedtestsettingInput)(nil)).Elem(), &Speedtestsetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpeedtestsettingArrayInput)(nil)).Elem(), SpeedtestsettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpeedtestsettingMapInput)(nil)).Elem(), SpeedtestsettingMap{})
	pulumi.RegisterOutputType(SpeedtestsettingOutput{})
	pulumi.RegisterOutputType(SpeedtestsettingArrayOutput{})
	pulumi.RegisterOutputType(SpeedtestsettingMapOutput{})
}
