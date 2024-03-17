// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure auto script.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewAutoscript(ctx, "auto2", &system.AutoscriptArgs{
//				Interval:   pulumi.Int(1),
//				OutputSize: pulumi.Int(10),
//				Repeat:     pulumi.Int(1),
//				Script: pulumi.String(`config firewall address
//	    edit "111"
//	        set color 3
//	        set subnet 1.1.1.1 255.255.255.255
//	    next
//
// end
//
// `),
//
//				Start: pulumi.String("auto"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// System AutoScript can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/autoscript:Autoscript labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/autoscript:Autoscript labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Autoscript struct {
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

// NewAutoscript registers a new resource with the given unique name, arguments, and options.
func NewAutoscript(ctx *pulumi.Context,
	name string, args *AutoscriptArgs, opts ...pulumi.ResourceOption) (*Autoscript, error) {
	if args == nil {
		args = &AutoscriptArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Autoscript
	err := ctx.RegisterResource("fortios:system/autoscript:Autoscript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoscript gets an existing Autoscript resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscriptState, opts ...pulumi.ResourceOption) (*Autoscript, error) {
	var resource Autoscript
	err := ctx.ReadResource("fortios:system/autoscript:Autoscript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Autoscript resources.
type autoscriptState struct {
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

type AutoscriptState struct {
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

func (AutoscriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscriptState)(nil)).Elem()
}

type autoscriptArgs struct {
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

// The set of arguments for constructing a Autoscript resource.
type AutoscriptArgs struct {
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

func (AutoscriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscriptArgs)(nil)).Elem()
}

type AutoscriptInput interface {
	pulumi.Input

	ToAutoscriptOutput() AutoscriptOutput
	ToAutoscriptOutputWithContext(ctx context.Context) AutoscriptOutput
}

func (*Autoscript) ElementType() reflect.Type {
	return reflect.TypeOf((**Autoscript)(nil)).Elem()
}

func (i *Autoscript) ToAutoscriptOutput() AutoscriptOutput {
	return i.ToAutoscriptOutputWithContext(context.Background())
}

func (i *Autoscript) ToAutoscriptOutputWithContext(ctx context.Context) AutoscriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscriptOutput)
}

// AutoscriptArrayInput is an input type that accepts AutoscriptArray and AutoscriptArrayOutput values.
// You can construct a concrete instance of `AutoscriptArrayInput` via:
//
//	AutoscriptArray{ AutoscriptArgs{...} }
type AutoscriptArrayInput interface {
	pulumi.Input

	ToAutoscriptArrayOutput() AutoscriptArrayOutput
	ToAutoscriptArrayOutputWithContext(context.Context) AutoscriptArrayOutput
}

type AutoscriptArray []AutoscriptInput

func (AutoscriptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Autoscript)(nil)).Elem()
}

func (i AutoscriptArray) ToAutoscriptArrayOutput() AutoscriptArrayOutput {
	return i.ToAutoscriptArrayOutputWithContext(context.Background())
}

func (i AutoscriptArray) ToAutoscriptArrayOutputWithContext(ctx context.Context) AutoscriptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscriptArrayOutput)
}

// AutoscriptMapInput is an input type that accepts AutoscriptMap and AutoscriptMapOutput values.
// You can construct a concrete instance of `AutoscriptMapInput` via:
//
//	AutoscriptMap{ "key": AutoscriptArgs{...} }
type AutoscriptMapInput interface {
	pulumi.Input

	ToAutoscriptMapOutput() AutoscriptMapOutput
	ToAutoscriptMapOutputWithContext(context.Context) AutoscriptMapOutput
}

type AutoscriptMap map[string]AutoscriptInput

func (AutoscriptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Autoscript)(nil)).Elem()
}

func (i AutoscriptMap) ToAutoscriptMapOutput() AutoscriptMapOutput {
	return i.ToAutoscriptMapOutputWithContext(context.Background())
}

func (i AutoscriptMap) ToAutoscriptMapOutputWithContext(ctx context.Context) AutoscriptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscriptMapOutput)
}

type AutoscriptOutput struct{ *pulumi.OutputState }

func (AutoscriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Autoscript)(nil)).Elem()
}

func (o AutoscriptOutput) ToAutoscriptOutput() AutoscriptOutput {
	return o
}

func (o AutoscriptOutput) ToAutoscriptOutputWithContext(ctx context.Context) AutoscriptOutput {
	return o
}

// Repeat interval in seconds.
func (o AutoscriptOutput) Interval() pulumi.IntOutput {
	return o.ApplyT(func(v *Autoscript) pulumi.IntOutput { return v.Interval }).(pulumi.IntOutput)
}

// Auto script name.
func (o AutoscriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Autoscript) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of megabytes to limit script output to (10 - 1024, default = 10).
func (o AutoscriptOutput) OutputSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Autoscript) pulumi.IntOutput { return v.OutputSize }).(pulumi.IntOutput)
}

// Number of times to repeat this script (0 = infinite).
func (o AutoscriptOutput) Repeat() pulumi.IntOutput {
	return o.ApplyT(func(v *Autoscript) pulumi.IntOutput { return v.Repeat }).(pulumi.IntOutput)
}

// List of FortiOS CLI commands to repeat.
func (o AutoscriptOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Autoscript) pulumi.StringPtrOutput { return v.Script }).(pulumi.StringPtrOutput)
}

// Script starting mode. Valid values: `manual`, `auto`.
func (o AutoscriptOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v *Autoscript) pulumi.StringOutput { return v.Start }).(pulumi.StringOutput)
}

// Maximum running time for this script in seconds (0 = no timeout).
func (o AutoscriptOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Autoscript) pulumi.IntOutput { return v.Timeout }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AutoscriptOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Autoscript) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AutoscriptArrayOutput struct{ *pulumi.OutputState }

func (AutoscriptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Autoscript)(nil)).Elem()
}

func (o AutoscriptArrayOutput) ToAutoscriptArrayOutput() AutoscriptArrayOutput {
	return o
}

func (o AutoscriptArrayOutput) ToAutoscriptArrayOutputWithContext(ctx context.Context) AutoscriptArrayOutput {
	return o
}

func (o AutoscriptArrayOutput) Index(i pulumi.IntInput) AutoscriptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Autoscript {
		return vs[0].([]*Autoscript)[vs[1].(int)]
	}).(AutoscriptOutput)
}

type AutoscriptMapOutput struct{ *pulumi.OutputState }

func (AutoscriptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Autoscript)(nil)).Elem()
}

func (o AutoscriptMapOutput) ToAutoscriptMapOutput() AutoscriptMapOutput {
	return o
}

func (o AutoscriptMapOutput) ToAutoscriptMapOutputWithContext(ctx context.Context) AutoscriptMapOutput {
	return o
}

func (o AutoscriptMapOutput) MapIndex(k pulumi.StringInput) AutoscriptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Autoscript {
		return vs[0].(map[string]*Autoscript)[vs[1].(string)]
	}).(AutoscriptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscriptInput)(nil)).Elem(), &Autoscript{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscriptArrayInput)(nil)).Elem(), AutoscriptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscriptMapInput)(nil)).Elem(), AutoscriptMap{})
	pulumi.RegisterOutputType(AutoscriptOutput{})
	pulumi.RegisterOutputType(AutoscriptArrayOutput{})
	pulumi.RegisterOutputType(AutoscriptMapOutput{})
}
