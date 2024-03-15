// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiSwitch SNMP trap threshold values globally. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// SwitchController SnmpTrapThreshold can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold labelname SwitchControllerSnmpTrapThreshold
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold labelname SwitchControllerSnmpTrapThreshold
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Snmptrapthreshold struct {
	pulumi.CustomResourceState

	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntOutput `pulumi:"trapHighCpuThreshold"`
	// Log disk usage when trap is sent.
	TrapLogFullThreshold pulumi.IntOutput `pulumi:"trapLogFullThreshold"`
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold pulumi.IntOutput `pulumi:"trapLowMemoryThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSnmptrapthreshold registers a new resource with the given unique name, arguments, and options.
func NewSnmptrapthreshold(ctx *pulumi.Context,
	name string, args *SnmptrapthresholdArgs, opts ...pulumi.ResourceOption) (*Snmptrapthreshold, error) {
	if args == nil {
		args = &SnmptrapthresholdArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Snmptrapthreshold
	err := ctx.RegisterResource("fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnmptrapthreshold gets an existing Snmptrapthreshold resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnmptrapthreshold(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnmptrapthresholdState, opts ...pulumi.ResourceOption) (*Snmptrapthreshold, error) {
	var resource Snmptrapthreshold
	err := ctx.ReadResource("fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snmptrapthreshold resources.
type snmptrapthresholdState struct {
	// CPU usage when trap is sent.
	TrapHighCpuThreshold *int `pulumi:"trapHighCpuThreshold"`
	// Log disk usage when trap is sent.
	TrapLogFullThreshold *int `pulumi:"trapLogFullThreshold"`
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold *int `pulumi:"trapLowMemoryThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SnmptrapthresholdState struct {
	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntPtrInput
	// Log disk usage when trap is sent.
	TrapLogFullThreshold pulumi.IntPtrInput
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SnmptrapthresholdState) ElementType() reflect.Type {
	return reflect.TypeOf((*snmptrapthresholdState)(nil)).Elem()
}

type snmptrapthresholdArgs struct {
	// CPU usage when trap is sent.
	TrapHighCpuThreshold *int `pulumi:"trapHighCpuThreshold"`
	// Log disk usage when trap is sent.
	TrapLogFullThreshold *int `pulumi:"trapLogFullThreshold"`
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold *int `pulumi:"trapLowMemoryThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Snmptrapthreshold resource.
type SnmptrapthresholdArgs struct {
	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntPtrInput
	// Log disk usage when trap is sent.
	TrapLogFullThreshold pulumi.IntPtrInput
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SnmptrapthresholdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snmptrapthresholdArgs)(nil)).Elem()
}

type SnmptrapthresholdInput interface {
	pulumi.Input

	ToSnmptrapthresholdOutput() SnmptrapthresholdOutput
	ToSnmptrapthresholdOutputWithContext(ctx context.Context) SnmptrapthresholdOutput
}

func (*Snmptrapthreshold) ElementType() reflect.Type {
	return reflect.TypeOf((**Snmptrapthreshold)(nil)).Elem()
}

func (i *Snmptrapthreshold) ToSnmptrapthresholdOutput() SnmptrapthresholdOutput {
	return i.ToSnmptrapthresholdOutputWithContext(context.Background())
}

func (i *Snmptrapthreshold) ToSnmptrapthresholdOutputWithContext(ctx context.Context) SnmptrapthresholdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmptrapthresholdOutput)
}

// SnmptrapthresholdArrayInput is an input type that accepts SnmptrapthresholdArray and SnmptrapthresholdArrayOutput values.
// You can construct a concrete instance of `SnmptrapthresholdArrayInput` via:
//
//	SnmptrapthresholdArray{ SnmptrapthresholdArgs{...} }
type SnmptrapthresholdArrayInput interface {
	pulumi.Input

	ToSnmptrapthresholdArrayOutput() SnmptrapthresholdArrayOutput
	ToSnmptrapthresholdArrayOutputWithContext(context.Context) SnmptrapthresholdArrayOutput
}

type SnmptrapthresholdArray []SnmptrapthresholdInput

func (SnmptrapthresholdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snmptrapthreshold)(nil)).Elem()
}

func (i SnmptrapthresholdArray) ToSnmptrapthresholdArrayOutput() SnmptrapthresholdArrayOutput {
	return i.ToSnmptrapthresholdArrayOutputWithContext(context.Background())
}

func (i SnmptrapthresholdArray) ToSnmptrapthresholdArrayOutputWithContext(ctx context.Context) SnmptrapthresholdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmptrapthresholdArrayOutput)
}

// SnmptrapthresholdMapInput is an input type that accepts SnmptrapthresholdMap and SnmptrapthresholdMapOutput values.
// You can construct a concrete instance of `SnmptrapthresholdMapInput` via:
//
//	SnmptrapthresholdMap{ "key": SnmptrapthresholdArgs{...} }
type SnmptrapthresholdMapInput interface {
	pulumi.Input

	ToSnmptrapthresholdMapOutput() SnmptrapthresholdMapOutput
	ToSnmptrapthresholdMapOutputWithContext(context.Context) SnmptrapthresholdMapOutput
}

type SnmptrapthresholdMap map[string]SnmptrapthresholdInput

func (SnmptrapthresholdMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snmptrapthreshold)(nil)).Elem()
}

func (i SnmptrapthresholdMap) ToSnmptrapthresholdMapOutput() SnmptrapthresholdMapOutput {
	return i.ToSnmptrapthresholdMapOutputWithContext(context.Background())
}

func (i SnmptrapthresholdMap) ToSnmptrapthresholdMapOutputWithContext(ctx context.Context) SnmptrapthresholdMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmptrapthresholdMapOutput)
}

type SnmptrapthresholdOutput struct{ *pulumi.OutputState }

func (SnmptrapthresholdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snmptrapthreshold)(nil)).Elem()
}

func (o SnmptrapthresholdOutput) ToSnmptrapthresholdOutput() SnmptrapthresholdOutput {
	return o
}

func (o SnmptrapthresholdOutput) ToSnmptrapthresholdOutputWithContext(ctx context.Context) SnmptrapthresholdOutput {
	return o
}

// CPU usage when trap is sent.
func (o SnmptrapthresholdOutput) TrapHighCpuThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Snmptrapthreshold) pulumi.IntOutput { return v.TrapHighCpuThreshold }).(pulumi.IntOutput)
}

// Log disk usage when trap is sent.
func (o SnmptrapthresholdOutput) TrapLogFullThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Snmptrapthreshold) pulumi.IntOutput { return v.TrapLogFullThreshold }).(pulumi.IntOutput)
}

// Memory usage when trap is sent.
func (o SnmptrapthresholdOutput) TrapLowMemoryThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Snmptrapthreshold) pulumi.IntOutput { return v.TrapLowMemoryThreshold }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SnmptrapthresholdOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snmptrapthreshold) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SnmptrapthresholdArrayOutput struct{ *pulumi.OutputState }

func (SnmptrapthresholdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snmptrapthreshold)(nil)).Elem()
}

func (o SnmptrapthresholdArrayOutput) ToSnmptrapthresholdArrayOutput() SnmptrapthresholdArrayOutput {
	return o
}

func (o SnmptrapthresholdArrayOutput) ToSnmptrapthresholdArrayOutputWithContext(ctx context.Context) SnmptrapthresholdArrayOutput {
	return o
}

func (o SnmptrapthresholdArrayOutput) Index(i pulumi.IntInput) SnmptrapthresholdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snmptrapthreshold {
		return vs[0].([]*Snmptrapthreshold)[vs[1].(int)]
	}).(SnmptrapthresholdOutput)
}

type SnmptrapthresholdMapOutput struct{ *pulumi.OutputState }

func (SnmptrapthresholdMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snmptrapthreshold)(nil)).Elem()
}

func (o SnmptrapthresholdMapOutput) ToSnmptrapthresholdMapOutput() SnmptrapthresholdMapOutput {
	return o
}

func (o SnmptrapthresholdMapOutput) ToSnmptrapthresholdMapOutputWithContext(ctx context.Context) SnmptrapthresholdMapOutput {
	return o
}

func (o SnmptrapthresholdMapOutput) MapIndex(k pulumi.StringInput) SnmptrapthresholdOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snmptrapthreshold {
		return vs[0].(map[string]*Snmptrapthreshold)[vs[1].(string)]
	}).(SnmptrapthresholdOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnmptrapthresholdInput)(nil)).Elem(), &Snmptrapthreshold{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnmptrapthresholdArrayInput)(nil)).Elem(), SnmptrapthresholdArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnmptrapthresholdMapInput)(nil)).Elem(), SnmptrapthresholdMap{})
	pulumi.RegisterOutputType(SnmptrapthresholdOutput{})
	pulumi.RegisterOutputType(SnmptrapthresholdArrayOutput{})
	pulumi.RegisterOutputType(SnmptrapthresholdMapOutput{})
}
