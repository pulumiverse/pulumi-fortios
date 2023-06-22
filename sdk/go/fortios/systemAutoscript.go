// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure auto script.
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
//			_, err := fortios.NewSystemAutoscript(ctx, "auto2", &fortios.SystemAutoscriptArgs{
//				Interval:   pulumi.Int(1),
//				OutputSize: pulumi.Int(10),
//				Repeat:     pulumi.Int(1),
//				Script:     pulumi.String("config firewall address\n    edit \"111\"\n        set color 3\n        set subnet 1.1.1.1 255.255.255.255\n    next\nend\n\n"),
//				Start:      pulumi.String("auto"),
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
// # System AutoScript can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemAutoscript:SystemAutoscript labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemAutoscript:SystemAutoscript labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemAutoscript struct {
	pulumi.CustomResourceState

	// Repeat interval in seconds.
	Interval pulumi.IntOutput `pulumi:"interval"`
	// Auto script name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of megabytes to limit script output to (10 - 1024, default = 10).
	OutputSize pulumi.IntOutput `pulumi:"outputSize"`
	// Number of times to repeat this script (0 = infinite).
	Repeat pulumi.IntOutput `pulumi:"repeat"`
	// List of FortiOS CLI commands to repeat.
	Script pulumi.StringPtrOutput `pulumi:"script"`
	// Script starting mode. Valid values: `manual`, `auto`.
	Start pulumi.StringOutput `pulumi:"start"`
	// Maximum running time for this script in seconds (0 = no timeout).
	Timeout pulumi.IntOutput `pulumi:"timeout"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemAutoscript registers a new resource with the given unique name, arguments, and options.
func NewSystemAutoscript(ctx *pulumi.Context,
	name string, args *SystemAutoscriptArgs, opts ...pulumi.ResourceOption) (*SystemAutoscript, error) {
	if args == nil {
		args = &SystemAutoscriptArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemAutoscript
	err := ctx.RegisterResource("fortios:index/systemAutoscript:SystemAutoscript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAutoscript gets an existing SystemAutoscript resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAutoscript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAutoscriptState, opts ...pulumi.ResourceOption) (*SystemAutoscript, error) {
	var resource SystemAutoscript
	err := ctx.ReadResource("fortios:index/systemAutoscript:SystemAutoscript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAutoscript resources.
type systemAutoscriptState struct {
	// Repeat interval in seconds.
	Interval *int `pulumi:"interval"`
	// Auto script name.
	Name *string `pulumi:"name"`
	// Number of megabytes to limit script output to (10 - 1024, default = 10).
	OutputSize *int `pulumi:"outputSize"`
	// Number of times to repeat this script (0 = infinite).
	Repeat *int `pulumi:"repeat"`
	// List of FortiOS CLI commands to repeat.
	Script *string `pulumi:"script"`
	// Script starting mode. Valid values: `manual`, `auto`.
	Start *string `pulumi:"start"`
	// Maximum running time for this script in seconds (0 = no timeout).
	Timeout *int `pulumi:"timeout"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemAutoscriptState struct {
	// Repeat interval in seconds.
	Interval pulumi.IntPtrInput
	// Auto script name.
	Name pulumi.StringPtrInput
	// Number of megabytes to limit script output to (10 - 1024, default = 10).
	OutputSize pulumi.IntPtrInput
	// Number of times to repeat this script (0 = infinite).
	Repeat pulumi.IntPtrInput
	// List of FortiOS CLI commands to repeat.
	Script pulumi.StringPtrInput
	// Script starting mode. Valid values: `manual`, `auto`.
	Start pulumi.StringPtrInput
	// Maximum running time for this script in seconds (0 = no timeout).
	Timeout pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAutoscriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutoscriptState)(nil)).Elem()
}

type systemAutoscriptArgs struct {
	// Repeat interval in seconds.
	Interval *int `pulumi:"interval"`
	// Auto script name.
	Name *string `pulumi:"name"`
	// Number of megabytes to limit script output to (10 - 1024, default = 10).
	OutputSize *int `pulumi:"outputSize"`
	// Number of times to repeat this script (0 = infinite).
	Repeat *int `pulumi:"repeat"`
	// List of FortiOS CLI commands to repeat.
	Script *string `pulumi:"script"`
	// Script starting mode. Valid values: `manual`, `auto`.
	Start *string `pulumi:"start"`
	// Maximum running time for this script in seconds (0 = no timeout).
	Timeout *int `pulumi:"timeout"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemAutoscript resource.
type SystemAutoscriptArgs struct {
	// Repeat interval in seconds.
	Interval pulumi.IntPtrInput
	// Auto script name.
	Name pulumi.StringPtrInput
	// Number of megabytes to limit script output to (10 - 1024, default = 10).
	OutputSize pulumi.IntPtrInput
	// Number of times to repeat this script (0 = infinite).
	Repeat pulumi.IntPtrInput
	// List of FortiOS CLI commands to repeat.
	Script pulumi.StringPtrInput
	// Script starting mode. Valid values: `manual`, `auto`.
	Start pulumi.StringPtrInput
	// Maximum running time for this script in seconds (0 = no timeout).
	Timeout pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAutoscriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutoscriptArgs)(nil)).Elem()
}

type SystemAutoscriptInput interface {
	pulumi.Input

	ToSystemAutoscriptOutput() SystemAutoscriptOutput
	ToSystemAutoscriptOutputWithContext(ctx context.Context) SystemAutoscriptOutput
}

func (*SystemAutoscript) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutoscript)(nil)).Elem()
}

func (i *SystemAutoscript) ToSystemAutoscriptOutput() SystemAutoscriptOutput {
	return i.ToSystemAutoscriptOutputWithContext(context.Background())
}

func (i *SystemAutoscript) ToSystemAutoscriptOutputWithContext(ctx context.Context) SystemAutoscriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutoscriptOutput)
}

// SystemAutoscriptArrayInput is an input type that accepts SystemAutoscriptArray and SystemAutoscriptArrayOutput values.
// You can construct a concrete instance of `SystemAutoscriptArrayInput` via:
//
//	SystemAutoscriptArray{ SystemAutoscriptArgs{...} }
type SystemAutoscriptArrayInput interface {
	pulumi.Input

	ToSystemAutoscriptArrayOutput() SystemAutoscriptArrayOutput
	ToSystemAutoscriptArrayOutputWithContext(context.Context) SystemAutoscriptArrayOutput
}

type SystemAutoscriptArray []SystemAutoscriptInput

func (SystemAutoscriptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAutoscript)(nil)).Elem()
}

func (i SystemAutoscriptArray) ToSystemAutoscriptArrayOutput() SystemAutoscriptArrayOutput {
	return i.ToSystemAutoscriptArrayOutputWithContext(context.Background())
}

func (i SystemAutoscriptArray) ToSystemAutoscriptArrayOutputWithContext(ctx context.Context) SystemAutoscriptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutoscriptArrayOutput)
}

// SystemAutoscriptMapInput is an input type that accepts SystemAutoscriptMap and SystemAutoscriptMapOutput values.
// You can construct a concrete instance of `SystemAutoscriptMapInput` via:
//
//	SystemAutoscriptMap{ "key": SystemAutoscriptArgs{...} }
type SystemAutoscriptMapInput interface {
	pulumi.Input

	ToSystemAutoscriptMapOutput() SystemAutoscriptMapOutput
	ToSystemAutoscriptMapOutputWithContext(context.Context) SystemAutoscriptMapOutput
}

type SystemAutoscriptMap map[string]SystemAutoscriptInput

func (SystemAutoscriptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAutoscript)(nil)).Elem()
}

func (i SystemAutoscriptMap) ToSystemAutoscriptMapOutput() SystemAutoscriptMapOutput {
	return i.ToSystemAutoscriptMapOutputWithContext(context.Background())
}

func (i SystemAutoscriptMap) ToSystemAutoscriptMapOutputWithContext(ctx context.Context) SystemAutoscriptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutoscriptMapOutput)
}

type SystemAutoscriptOutput struct{ *pulumi.OutputState }

func (SystemAutoscriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutoscript)(nil)).Elem()
}

func (o SystemAutoscriptOutput) ToSystemAutoscriptOutput() SystemAutoscriptOutput {
	return o
}

func (o SystemAutoscriptOutput) ToSystemAutoscriptOutputWithContext(ctx context.Context) SystemAutoscriptOutput {
	return o
}

// Repeat interval in seconds.
func (o SystemAutoscriptOutput) Interval() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAutoscript) pulumi.IntOutput { return v.Interval }).(pulumi.IntOutput)
}

// Auto script name.
func (o SystemAutoscriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutoscript) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of megabytes to limit script output to (10 - 1024, default = 10).
func (o SystemAutoscriptOutput) OutputSize() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAutoscript) pulumi.IntOutput { return v.OutputSize }).(pulumi.IntOutput)
}

// Number of times to repeat this script (0 = infinite).
func (o SystemAutoscriptOutput) Repeat() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAutoscript) pulumi.IntOutput { return v.Repeat }).(pulumi.IntOutput)
}

// List of FortiOS CLI commands to repeat.
func (o SystemAutoscriptOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutoscript) pulumi.StringPtrOutput { return v.Script }).(pulumi.StringPtrOutput)
}

// Script starting mode. Valid values: `manual`, `auto`.
func (o SystemAutoscriptOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutoscript) pulumi.StringOutput { return v.Start }).(pulumi.StringOutput)
}

// Maximum running time for this script in seconds (0 = no timeout).
func (o SystemAutoscriptOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAutoscript) pulumi.IntOutput { return v.Timeout }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemAutoscriptOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutoscript) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemAutoscriptArrayOutput struct{ *pulumi.OutputState }

func (SystemAutoscriptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAutoscript)(nil)).Elem()
}

func (o SystemAutoscriptArrayOutput) ToSystemAutoscriptArrayOutput() SystemAutoscriptArrayOutput {
	return o
}

func (o SystemAutoscriptArrayOutput) ToSystemAutoscriptArrayOutputWithContext(ctx context.Context) SystemAutoscriptArrayOutput {
	return o
}

func (o SystemAutoscriptArrayOutput) Index(i pulumi.IntInput) SystemAutoscriptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAutoscript {
		return vs[0].([]*SystemAutoscript)[vs[1].(int)]
	}).(SystemAutoscriptOutput)
}

type SystemAutoscriptMapOutput struct{ *pulumi.OutputState }

func (SystemAutoscriptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAutoscript)(nil)).Elem()
}

func (o SystemAutoscriptMapOutput) ToSystemAutoscriptMapOutput() SystemAutoscriptMapOutput {
	return o
}

func (o SystemAutoscriptMapOutput) ToSystemAutoscriptMapOutputWithContext(ctx context.Context) SystemAutoscriptMapOutput {
	return o
}

func (o SystemAutoscriptMapOutput) MapIndex(k pulumi.StringInput) SystemAutoscriptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAutoscript {
		return vs[0].(map[string]*SystemAutoscript)[vs[1].(string)]
	}).(SystemAutoscriptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutoscriptInput)(nil)).Elem(), &SystemAutoscript{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutoscriptArrayInput)(nil)).Elem(), SystemAutoscriptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutoscriptMapInput)(nil)).Elem(), SystemAutoscriptMap{})
	pulumi.RegisterOutputType(SystemAutoscriptOutput{})
	pulumi.RegisterOutputType(SystemAutoscriptArrayOutput{})
	pulumi.RegisterOutputType(SystemAutoscriptMapOutput{})
}
