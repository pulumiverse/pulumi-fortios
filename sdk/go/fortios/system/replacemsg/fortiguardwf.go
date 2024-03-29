// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package replacemsg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Replacement messages.
//
// ## Import
//
// SystemReplacemsg FortiguardWf can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/replacemsg/fortiguardwf:Fortiguardwf labelname {{msg_type}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/replacemsg/fortiguardwf:Fortiguardwf labelname {{msg_type}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Fortiguardwf struct {
	pulumi.CustomResourceState

	// Message string.
	Buffer pulumi.StringPtrOutput `pulumi:"buffer"`
	// Format flag.
	Format pulumi.StringOutput `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringOutput `pulumi:"header"`
	// Message type.
	MsgType pulumi.StringOutput `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFortiguardwf registers a new resource with the given unique name, arguments, and options.
func NewFortiguardwf(ctx *pulumi.Context,
	name string, args *FortiguardwfArgs, opts ...pulumi.ResourceOption) (*Fortiguardwf, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Fortiguardwf
	err := ctx.RegisterResource("fortios:system/replacemsg/fortiguardwf:Fortiguardwf", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortiguardwf gets an existing Fortiguardwf resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortiguardwf(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortiguardwfState, opts ...pulumi.ResourceOption) (*Fortiguardwf, error) {
	var resource Fortiguardwf
	err := ctx.ReadResource("fortios:system/replacemsg/fortiguardwf:Fortiguardwf", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fortiguardwf resources.
type fortiguardwfState struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType *string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FortiguardwfState struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FortiguardwfState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortiguardwfState)(nil)).Elem()
}

type fortiguardwfArgs struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Fortiguardwf resource.
type FortiguardwfArgs struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FortiguardwfArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortiguardwfArgs)(nil)).Elem()
}

type FortiguardwfInput interface {
	pulumi.Input

	ToFortiguardwfOutput() FortiguardwfOutput
	ToFortiguardwfOutputWithContext(ctx context.Context) FortiguardwfOutput
}

func (*Fortiguardwf) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortiguardwf)(nil)).Elem()
}

func (i *Fortiguardwf) ToFortiguardwfOutput() FortiguardwfOutput {
	return i.ToFortiguardwfOutputWithContext(context.Background())
}

func (i *Fortiguardwf) ToFortiguardwfOutputWithContext(ctx context.Context) FortiguardwfOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortiguardwfOutput)
}

// FortiguardwfArrayInput is an input type that accepts FortiguardwfArray and FortiguardwfArrayOutput values.
// You can construct a concrete instance of `FortiguardwfArrayInput` via:
//
//	FortiguardwfArray{ FortiguardwfArgs{...} }
type FortiguardwfArrayInput interface {
	pulumi.Input

	ToFortiguardwfArrayOutput() FortiguardwfArrayOutput
	ToFortiguardwfArrayOutputWithContext(context.Context) FortiguardwfArrayOutput
}

type FortiguardwfArray []FortiguardwfInput

func (FortiguardwfArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortiguardwf)(nil)).Elem()
}

func (i FortiguardwfArray) ToFortiguardwfArrayOutput() FortiguardwfArrayOutput {
	return i.ToFortiguardwfArrayOutputWithContext(context.Background())
}

func (i FortiguardwfArray) ToFortiguardwfArrayOutputWithContext(ctx context.Context) FortiguardwfArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortiguardwfArrayOutput)
}

// FortiguardwfMapInput is an input type that accepts FortiguardwfMap and FortiguardwfMapOutput values.
// You can construct a concrete instance of `FortiguardwfMapInput` via:
//
//	FortiguardwfMap{ "key": FortiguardwfArgs{...} }
type FortiguardwfMapInput interface {
	pulumi.Input

	ToFortiguardwfMapOutput() FortiguardwfMapOutput
	ToFortiguardwfMapOutputWithContext(context.Context) FortiguardwfMapOutput
}

type FortiguardwfMap map[string]FortiguardwfInput

func (FortiguardwfMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortiguardwf)(nil)).Elem()
}

func (i FortiguardwfMap) ToFortiguardwfMapOutput() FortiguardwfMapOutput {
	return i.ToFortiguardwfMapOutputWithContext(context.Background())
}

func (i FortiguardwfMap) ToFortiguardwfMapOutputWithContext(ctx context.Context) FortiguardwfMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortiguardwfMapOutput)
}

type FortiguardwfOutput struct{ *pulumi.OutputState }

func (FortiguardwfOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fortiguardwf)(nil)).Elem()
}

func (o FortiguardwfOutput) ToFortiguardwfOutput() FortiguardwfOutput {
	return o
}

func (o FortiguardwfOutput) ToFortiguardwfOutputWithContext(ctx context.Context) FortiguardwfOutput {
	return o
}

// Message string.
func (o FortiguardwfOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fortiguardwf) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag.
func (o FortiguardwfOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortiguardwf) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o FortiguardwfOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortiguardwf) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o FortiguardwfOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *Fortiguardwf) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FortiguardwfOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fortiguardwf) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FortiguardwfArrayOutput struct{ *pulumi.OutputState }

func (FortiguardwfArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fortiguardwf)(nil)).Elem()
}

func (o FortiguardwfArrayOutput) ToFortiguardwfArrayOutput() FortiguardwfArrayOutput {
	return o
}

func (o FortiguardwfArrayOutput) ToFortiguardwfArrayOutputWithContext(ctx context.Context) FortiguardwfArrayOutput {
	return o
}

func (o FortiguardwfArrayOutput) Index(i pulumi.IntInput) FortiguardwfOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fortiguardwf {
		return vs[0].([]*Fortiguardwf)[vs[1].(int)]
	}).(FortiguardwfOutput)
}

type FortiguardwfMapOutput struct{ *pulumi.OutputState }

func (FortiguardwfMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fortiguardwf)(nil)).Elem()
}

func (o FortiguardwfMapOutput) ToFortiguardwfMapOutput() FortiguardwfMapOutput {
	return o
}

func (o FortiguardwfMapOutput) ToFortiguardwfMapOutputWithContext(ctx context.Context) FortiguardwfMapOutput {
	return o
}

func (o FortiguardwfMapOutput) MapIndex(k pulumi.StringInput) FortiguardwfOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fortiguardwf {
		return vs[0].(map[string]*Fortiguardwf)[vs[1].(string)]
	}).(FortiguardwfOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortiguardwfInput)(nil)).Elem(), &Fortiguardwf{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortiguardwfArrayInput)(nil)).Elem(), FortiguardwfArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortiguardwfMapInput)(nil)).Elem(), FortiguardwfMap{})
	pulumi.RegisterOutputType(FortiguardwfOutput{})
	pulumi.RegisterOutputType(FortiguardwfArrayOutput{})
	pulumi.RegisterOutputType(FortiguardwfMapOutput{})
}
