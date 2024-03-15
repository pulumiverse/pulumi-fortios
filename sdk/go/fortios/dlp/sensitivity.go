// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dlp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// Dlp Sensitivity can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:dlp/sensitivity:Sensitivity labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:dlp/sensitivity:Sensitivity labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Sensitivity struct {
	pulumi.CustomResourceState

	// DLP Sensitivity Levels.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSensitivity registers a new resource with the given unique name, arguments, and options.
func NewSensitivity(ctx *pulumi.Context,
	name string, args *SensitivityArgs, opts ...pulumi.ResourceOption) (*Sensitivity, error) {
	if args == nil {
		args = &SensitivityArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Sensitivity
	err := ctx.RegisterResource("fortios:dlp/sensitivity:Sensitivity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSensitivity gets an existing Sensitivity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSensitivity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SensitivityState, opts ...pulumi.ResourceOption) (*Sensitivity, error) {
	var resource Sensitivity
	err := ctx.ReadResource("fortios:dlp/sensitivity:Sensitivity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Sensitivity resources.
type sensitivityState struct {
	// DLP Sensitivity Levels.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SensitivityState struct {
	// DLP Sensitivity Levels.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SensitivityState) ElementType() reflect.Type {
	return reflect.TypeOf((*sensitivityState)(nil)).Elem()
}

type sensitivityArgs struct {
	// DLP Sensitivity Levels.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Sensitivity resource.
type SensitivityArgs struct {
	// DLP Sensitivity Levels.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SensitivityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sensitivityArgs)(nil)).Elem()
}

type SensitivityInput interface {
	pulumi.Input

	ToSensitivityOutput() SensitivityOutput
	ToSensitivityOutputWithContext(ctx context.Context) SensitivityOutput
}

func (*Sensitivity) ElementType() reflect.Type {
	return reflect.TypeOf((**Sensitivity)(nil)).Elem()
}

func (i *Sensitivity) ToSensitivityOutput() SensitivityOutput {
	return i.ToSensitivityOutputWithContext(context.Background())
}

func (i *Sensitivity) ToSensitivityOutputWithContext(ctx context.Context) SensitivityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SensitivityOutput)
}

// SensitivityArrayInput is an input type that accepts SensitivityArray and SensitivityArrayOutput values.
// You can construct a concrete instance of `SensitivityArrayInput` via:
//
//	SensitivityArray{ SensitivityArgs{...} }
type SensitivityArrayInput interface {
	pulumi.Input

	ToSensitivityArrayOutput() SensitivityArrayOutput
	ToSensitivityArrayOutputWithContext(context.Context) SensitivityArrayOutput
}

type SensitivityArray []SensitivityInput

func (SensitivityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sensitivity)(nil)).Elem()
}

func (i SensitivityArray) ToSensitivityArrayOutput() SensitivityArrayOutput {
	return i.ToSensitivityArrayOutputWithContext(context.Background())
}

func (i SensitivityArray) ToSensitivityArrayOutputWithContext(ctx context.Context) SensitivityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SensitivityArrayOutput)
}

// SensitivityMapInput is an input type that accepts SensitivityMap and SensitivityMapOutput values.
// You can construct a concrete instance of `SensitivityMapInput` via:
//
//	SensitivityMap{ "key": SensitivityArgs{...} }
type SensitivityMapInput interface {
	pulumi.Input

	ToSensitivityMapOutput() SensitivityMapOutput
	ToSensitivityMapOutputWithContext(context.Context) SensitivityMapOutput
}

type SensitivityMap map[string]SensitivityInput

func (SensitivityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sensitivity)(nil)).Elem()
}

func (i SensitivityMap) ToSensitivityMapOutput() SensitivityMapOutput {
	return i.ToSensitivityMapOutputWithContext(context.Background())
}

func (i SensitivityMap) ToSensitivityMapOutputWithContext(ctx context.Context) SensitivityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SensitivityMapOutput)
}

type SensitivityOutput struct{ *pulumi.OutputState }

func (SensitivityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sensitivity)(nil)).Elem()
}

func (o SensitivityOutput) ToSensitivityOutput() SensitivityOutput {
	return o
}

func (o SensitivityOutput) ToSensitivityOutputWithContext(ctx context.Context) SensitivityOutput {
	return o
}

// DLP Sensitivity Levels.
func (o SensitivityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Sensitivity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SensitivityOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sensitivity) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SensitivityArrayOutput struct{ *pulumi.OutputState }

func (SensitivityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sensitivity)(nil)).Elem()
}

func (o SensitivityArrayOutput) ToSensitivityArrayOutput() SensitivityArrayOutput {
	return o
}

func (o SensitivityArrayOutput) ToSensitivityArrayOutputWithContext(ctx context.Context) SensitivityArrayOutput {
	return o
}

func (o SensitivityArrayOutput) Index(i pulumi.IntInput) SensitivityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Sensitivity {
		return vs[0].([]*Sensitivity)[vs[1].(int)]
	}).(SensitivityOutput)
}

type SensitivityMapOutput struct{ *pulumi.OutputState }

func (SensitivityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sensitivity)(nil)).Elem()
}

func (o SensitivityMapOutput) ToSensitivityMapOutput() SensitivityMapOutput {
	return o
}

func (o SensitivityMapOutput) ToSensitivityMapOutputWithContext(ctx context.Context) SensitivityMapOutput {
	return o
}

func (o SensitivityMapOutput) MapIndex(k pulumi.StringInput) SensitivityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Sensitivity {
		return vs[0].(map[string]*Sensitivity)[vs[1].(string)]
	}).(SensitivityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SensitivityInput)(nil)).Elem(), &Sensitivity{})
	pulumi.RegisterInputType(reflect.TypeOf((*SensitivityArrayInput)(nil)).Elem(), SensitivityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SensitivityMapInput)(nil)).Elem(), SensitivityMap{})
	pulumi.RegisterOutputType(SensitivityOutput{})
	pulumi.RegisterOutputType(SensitivityArrayOutput{})
	pulumi.RegisterOutputType(SensitivityMapOutput{})
}