// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Override global FortiCloud logging settings for this VDOM.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewLogfortiguardOverridesetting(ctx, "trname", &fortios.LogfortiguardOverridesettingArgs{
//				Override:       pulumi.String("disable"),
//				Status:         pulumi.String("disable"),
//				UploadInterval: pulumi.String("daily"),
//				UploadOption:   pulumi.String("5-minute"),
//				UploadTime:     pulumi.String("00:00"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # LogFortiguard OverrideSetting can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/logfortiguardOverridesetting:LogfortiguardOverridesetting labelname LogFortiguardOverrideSetting
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/logfortiguardOverridesetting:LogfortiguardOverridesetting labelname LogFortiguardOverrideSetting
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type LogfortiguardOverridesetting struct {
	pulumi.CustomResourceState

	// Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig pulumi.StringOutput `pulumi:"accessConfig"`
	// FortiCloud maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntOutput `pulumi:"maxLogRate"`
	// Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
	Override pulumi.StringOutput `pulumi:"override"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringOutput `pulumi:"priority"`
	// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Day of week to roll logs.
	UploadDay pulumi.StringOutput `pulumi:"uploadDay"`
	// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval pulumi.StringOutput `pulumi:"uploadInterval"`
	// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption pulumi.StringOutput `pulumi:"uploadOption"`
	// Time of day to roll logs (hh:mm).
	UploadTime pulumi.StringOutput `pulumi:"uploadTime"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogfortiguardOverridesetting registers a new resource with the given unique name, arguments, and options.
func NewLogfortiguardOverridesetting(ctx *pulumi.Context,
	name string, args *LogfortiguardOverridesettingArgs, opts ...pulumi.ResourceOption) (*LogfortiguardOverridesetting, error) {
	if args == nil {
		args = &LogfortiguardOverridesettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogfortiguardOverridesetting
	err := ctx.RegisterResource("fortios:index/logfortiguardOverridesetting:LogfortiguardOverridesetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogfortiguardOverridesetting gets an existing LogfortiguardOverridesetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogfortiguardOverridesetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogfortiguardOverridesettingState, opts ...pulumi.ResourceOption) (*LogfortiguardOverridesetting, error) {
	var resource LogfortiguardOverridesetting
	err := ctx.ReadResource("fortios:index/logfortiguardOverridesetting:LogfortiguardOverridesetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogfortiguardOverridesetting resources.
type logfortiguardOverridesettingState struct {
	// Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig *string `pulumi:"accessConfig"`
	// FortiCloud maximum log rate in MBps (0 = unlimited).
	MaxLogRate *int `pulumi:"maxLogRate"`
	// Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
	Override *string `pulumi:"override"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Day of week to roll logs.
	UploadDay *string `pulumi:"uploadDay"`
	// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval *string `pulumi:"uploadInterval"`
	// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption *string `pulumi:"uploadOption"`
	// Time of day to roll logs (hh:mm).
	UploadTime *string `pulumi:"uploadTime"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogfortiguardOverridesettingState struct {
	// Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig pulumi.StringPtrInput
	// FortiCloud maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntPtrInput
	// Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
	Override pulumi.StringPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Day of week to roll logs.
	UploadDay pulumi.StringPtrInput
	// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval pulumi.StringPtrInput
	// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption pulumi.StringPtrInput
	// Time of day to roll logs (hh:mm).
	UploadTime pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogfortiguardOverridesettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logfortiguardOverridesettingState)(nil)).Elem()
}

type logfortiguardOverridesettingArgs struct {
	// Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig *string `pulumi:"accessConfig"`
	// FortiCloud maximum log rate in MBps (0 = unlimited).
	MaxLogRate *int `pulumi:"maxLogRate"`
	// Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
	Override *string `pulumi:"override"`
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority *string `pulumi:"priority"`
	// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Day of week to roll logs.
	UploadDay *string `pulumi:"uploadDay"`
	// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval *string `pulumi:"uploadInterval"`
	// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption *string `pulumi:"uploadOption"`
	// Time of day to roll logs (hh:mm).
	UploadTime *string `pulumi:"uploadTime"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogfortiguardOverridesetting resource.
type LogfortiguardOverridesettingArgs struct {
	// Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
	AccessConfig pulumi.StringPtrInput
	// FortiCloud maximum log rate in MBps (0 = unlimited).
	MaxLogRate pulumi.IntPtrInput
	// Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
	Override pulumi.StringPtrInput
	// Set log transmission priority. Valid values: `default`, `low`.
	Priority pulumi.StringPtrInput
	// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Day of week to roll logs.
	UploadDay pulumi.StringPtrInput
	// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
	UploadInterval pulumi.StringPtrInput
	// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
	UploadOption pulumi.StringPtrInput
	// Time of day to roll logs (hh:mm).
	UploadTime pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogfortiguardOverridesettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logfortiguardOverridesettingArgs)(nil)).Elem()
}

type LogfortiguardOverridesettingInput interface {
	pulumi.Input

	ToLogfortiguardOverridesettingOutput() LogfortiguardOverridesettingOutput
	ToLogfortiguardOverridesettingOutputWithContext(ctx context.Context) LogfortiguardOverridesettingOutput
}

func (*LogfortiguardOverridesetting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogfortiguardOverridesetting)(nil)).Elem()
}

func (i *LogfortiguardOverridesetting) ToLogfortiguardOverridesettingOutput() LogfortiguardOverridesettingOutput {
	return i.ToLogfortiguardOverridesettingOutputWithContext(context.Background())
}

func (i *LogfortiguardOverridesetting) ToLogfortiguardOverridesettingOutputWithContext(ctx context.Context) LogfortiguardOverridesettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogfortiguardOverridesettingOutput)
}

// LogfortiguardOverridesettingArrayInput is an input type that accepts LogfortiguardOverridesettingArray and LogfortiguardOverridesettingArrayOutput values.
// You can construct a concrete instance of `LogfortiguardOverridesettingArrayInput` via:
//
//	LogfortiguardOverridesettingArray{ LogfortiguardOverridesettingArgs{...} }
type LogfortiguardOverridesettingArrayInput interface {
	pulumi.Input

	ToLogfortiguardOverridesettingArrayOutput() LogfortiguardOverridesettingArrayOutput
	ToLogfortiguardOverridesettingArrayOutputWithContext(context.Context) LogfortiguardOverridesettingArrayOutput
}

type LogfortiguardOverridesettingArray []LogfortiguardOverridesettingInput

func (LogfortiguardOverridesettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogfortiguardOverridesetting)(nil)).Elem()
}

func (i LogfortiguardOverridesettingArray) ToLogfortiguardOverridesettingArrayOutput() LogfortiguardOverridesettingArrayOutput {
	return i.ToLogfortiguardOverridesettingArrayOutputWithContext(context.Background())
}

func (i LogfortiguardOverridesettingArray) ToLogfortiguardOverridesettingArrayOutputWithContext(ctx context.Context) LogfortiguardOverridesettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogfortiguardOverridesettingArrayOutput)
}

// LogfortiguardOverridesettingMapInput is an input type that accepts LogfortiguardOverridesettingMap and LogfortiguardOverridesettingMapOutput values.
// You can construct a concrete instance of `LogfortiguardOverridesettingMapInput` via:
//
//	LogfortiguardOverridesettingMap{ "key": LogfortiguardOverridesettingArgs{...} }
type LogfortiguardOverridesettingMapInput interface {
	pulumi.Input

	ToLogfortiguardOverridesettingMapOutput() LogfortiguardOverridesettingMapOutput
	ToLogfortiguardOverridesettingMapOutputWithContext(context.Context) LogfortiguardOverridesettingMapOutput
}

type LogfortiguardOverridesettingMap map[string]LogfortiguardOverridesettingInput

func (LogfortiguardOverridesettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogfortiguardOverridesetting)(nil)).Elem()
}

func (i LogfortiguardOverridesettingMap) ToLogfortiguardOverridesettingMapOutput() LogfortiguardOverridesettingMapOutput {
	return i.ToLogfortiguardOverridesettingMapOutputWithContext(context.Background())
}

func (i LogfortiguardOverridesettingMap) ToLogfortiguardOverridesettingMapOutputWithContext(ctx context.Context) LogfortiguardOverridesettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogfortiguardOverridesettingMapOutput)
}

type LogfortiguardOverridesettingOutput struct{ *pulumi.OutputState }

func (LogfortiguardOverridesettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogfortiguardOverridesetting)(nil)).Elem()
}

func (o LogfortiguardOverridesettingOutput) ToLogfortiguardOverridesettingOutput() LogfortiguardOverridesettingOutput {
	return o
}

func (o LogfortiguardOverridesettingOutput) ToLogfortiguardOverridesettingOutputWithContext(ctx context.Context) LogfortiguardOverridesettingOutput {
	return o
}

// Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
func (o LogfortiguardOverridesettingOutput) AccessConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *LogfortiguardOverridesetting) pulumi.StringOutput { return v.AccessConfig }).(pulumi.StringOutput)
}

// FortiCloud maximum log rate in MBps (0 = unlimited).
func (o LogfortiguardOverridesettingOutput) MaxLogRate() pulumi.IntOutput {
	return o.ApplyT(func(v *LogfortiguardOverridesetting) pulumi.IntOutput { return v.MaxLogRate }).(pulumi.IntOutput)
}

// Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
func (o LogfortiguardOverridesettingOutput) Override() pulumi.StringOutput {
	return o.ApplyT(func(v *LogfortiguardOverridesetting) pulumi.StringOutput { return v.Override }).(pulumi.StringOutput)
}

// Set log transmission priority. Valid values: `default`, `low`.
func (o LogfortiguardOverridesettingOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v *LogfortiguardOverridesetting) pulumi.StringOutput { return v.Priority }).(pulumi.StringOutput)
}

// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
func (o LogfortiguardOverridesettingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LogfortiguardOverridesetting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Day of week to roll logs.
func (o LogfortiguardOverridesettingOutput) UploadDay() pulumi.StringOutput {
	return o.ApplyT(func(v *LogfortiguardOverridesetting) pulumi.StringOutput { return v.UploadDay }).(pulumi.StringOutput)
}

// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
func (o LogfortiguardOverridesettingOutput) UploadInterval() pulumi.StringOutput {
	return o.ApplyT(func(v *LogfortiguardOverridesetting) pulumi.StringOutput { return v.UploadInterval }).(pulumi.StringOutput)
}

// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
func (o LogfortiguardOverridesettingOutput) UploadOption() pulumi.StringOutput {
	return o.ApplyT(func(v *LogfortiguardOverridesetting) pulumi.StringOutput { return v.UploadOption }).(pulumi.StringOutput)
}

// Time of day to roll logs (hh:mm).
func (o LogfortiguardOverridesettingOutput) UploadTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LogfortiguardOverridesetting) pulumi.StringOutput { return v.UploadTime }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o LogfortiguardOverridesettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogfortiguardOverridesetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type LogfortiguardOverridesettingArrayOutput struct{ *pulumi.OutputState }

func (LogfortiguardOverridesettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogfortiguardOverridesetting)(nil)).Elem()
}

func (o LogfortiguardOverridesettingArrayOutput) ToLogfortiguardOverridesettingArrayOutput() LogfortiguardOverridesettingArrayOutput {
	return o
}

func (o LogfortiguardOverridesettingArrayOutput) ToLogfortiguardOverridesettingArrayOutputWithContext(ctx context.Context) LogfortiguardOverridesettingArrayOutput {
	return o
}

func (o LogfortiguardOverridesettingArrayOutput) Index(i pulumi.IntInput) LogfortiguardOverridesettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogfortiguardOverridesetting {
		return vs[0].([]*LogfortiguardOverridesetting)[vs[1].(int)]
	}).(LogfortiguardOverridesettingOutput)
}

type LogfortiguardOverridesettingMapOutput struct{ *pulumi.OutputState }

func (LogfortiguardOverridesettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogfortiguardOverridesetting)(nil)).Elem()
}

func (o LogfortiguardOverridesettingMapOutput) ToLogfortiguardOverridesettingMapOutput() LogfortiguardOverridesettingMapOutput {
	return o
}

func (o LogfortiguardOverridesettingMapOutput) ToLogfortiguardOverridesettingMapOutputWithContext(ctx context.Context) LogfortiguardOverridesettingMapOutput {
	return o
}

func (o LogfortiguardOverridesettingMapOutput) MapIndex(k pulumi.StringInput) LogfortiguardOverridesettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogfortiguardOverridesetting {
		return vs[0].(map[string]*LogfortiguardOverridesetting)[vs[1].(string)]
	}).(LogfortiguardOverridesettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogfortiguardOverridesettingInput)(nil)).Elem(), &LogfortiguardOverridesetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogfortiguardOverridesettingArrayInput)(nil)).Elem(), LogfortiguardOverridesettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogfortiguardOverridesettingMapInput)(nil)).Elem(), LogfortiguardOverridesettingMap{})
	pulumi.RegisterOutputType(LogfortiguardOverridesettingOutput{})
	pulumi.RegisterOutputType(LogfortiguardOverridesettingArrayOutput{})
	pulumi.RegisterOutputType(LogfortiguardOverridesettingMapOutput{})
}
