// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replacement messages.
//
// ## Import
//
// # SystemReplacemsg FortiguardWf can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgFortiguardwf:SystemreplacemsgFortiguardwf labelname {{msg_type}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgFortiguardwf:SystemreplacemsgFortiguardwf labelname {{msg_type}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemreplacemsgFortiguardwf struct {
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

// NewSystemreplacemsgFortiguardwf registers a new resource with the given unique name, arguments, and options.
func NewSystemreplacemsgFortiguardwf(ctx *pulumi.Context,
	name string, args *SystemreplacemsgFortiguardwfArgs, opts ...pulumi.ResourceOption) (*SystemreplacemsgFortiguardwf, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemreplacemsgFortiguardwf
	err := ctx.RegisterResource("fortios:index/systemreplacemsgFortiguardwf:SystemreplacemsgFortiguardwf", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemreplacemsgFortiguardwf gets an existing SystemreplacemsgFortiguardwf resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemreplacemsgFortiguardwf(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemreplacemsgFortiguardwfState, opts ...pulumi.ResourceOption) (*SystemreplacemsgFortiguardwf, error) {
	var resource SystemreplacemsgFortiguardwf
	err := ctx.ReadResource("fortios:index/systemreplacemsgFortiguardwf:SystemreplacemsgFortiguardwf", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemreplacemsgFortiguardwf resources.
type systemreplacemsgFortiguardwfState struct {
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

type SystemreplacemsgFortiguardwfState struct {
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

func (SystemreplacemsgFortiguardwfState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgFortiguardwfState)(nil)).Elem()
}

type systemreplacemsgFortiguardwfArgs struct {
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

// The set of arguments for constructing a SystemreplacemsgFortiguardwf resource.
type SystemreplacemsgFortiguardwfArgs struct {
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

func (SystemreplacemsgFortiguardwfArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgFortiguardwfArgs)(nil)).Elem()
}

type SystemreplacemsgFortiguardwfInput interface {
	pulumi.Input

	ToSystemreplacemsgFortiguardwfOutput() SystemreplacemsgFortiguardwfOutput
	ToSystemreplacemsgFortiguardwfOutputWithContext(ctx context.Context) SystemreplacemsgFortiguardwfOutput
}

func (*SystemreplacemsgFortiguardwf) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgFortiguardwf)(nil)).Elem()
}

func (i *SystemreplacemsgFortiguardwf) ToSystemreplacemsgFortiguardwfOutput() SystemreplacemsgFortiguardwfOutput {
	return i.ToSystemreplacemsgFortiguardwfOutputWithContext(context.Background())
}

func (i *SystemreplacemsgFortiguardwf) ToSystemreplacemsgFortiguardwfOutputWithContext(ctx context.Context) SystemreplacemsgFortiguardwfOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgFortiguardwfOutput)
}

// SystemreplacemsgFortiguardwfArrayInput is an input type that accepts SystemreplacemsgFortiguardwfArray and SystemreplacemsgFortiguardwfArrayOutput values.
// You can construct a concrete instance of `SystemreplacemsgFortiguardwfArrayInput` via:
//
//	SystemreplacemsgFortiguardwfArray{ SystemreplacemsgFortiguardwfArgs{...} }
type SystemreplacemsgFortiguardwfArrayInput interface {
	pulumi.Input

	ToSystemreplacemsgFortiguardwfArrayOutput() SystemreplacemsgFortiguardwfArrayOutput
	ToSystemreplacemsgFortiguardwfArrayOutputWithContext(context.Context) SystemreplacemsgFortiguardwfArrayOutput
}

type SystemreplacemsgFortiguardwfArray []SystemreplacemsgFortiguardwfInput

func (SystemreplacemsgFortiguardwfArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgFortiguardwf)(nil)).Elem()
}

func (i SystemreplacemsgFortiguardwfArray) ToSystemreplacemsgFortiguardwfArrayOutput() SystemreplacemsgFortiguardwfArrayOutput {
	return i.ToSystemreplacemsgFortiguardwfArrayOutputWithContext(context.Background())
}

func (i SystemreplacemsgFortiguardwfArray) ToSystemreplacemsgFortiguardwfArrayOutputWithContext(ctx context.Context) SystemreplacemsgFortiguardwfArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgFortiguardwfArrayOutput)
}

// SystemreplacemsgFortiguardwfMapInput is an input type that accepts SystemreplacemsgFortiguardwfMap and SystemreplacemsgFortiguardwfMapOutput values.
// You can construct a concrete instance of `SystemreplacemsgFortiguardwfMapInput` via:
//
//	SystemreplacemsgFortiguardwfMap{ "key": SystemreplacemsgFortiguardwfArgs{...} }
type SystemreplacemsgFortiguardwfMapInput interface {
	pulumi.Input

	ToSystemreplacemsgFortiguardwfMapOutput() SystemreplacemsgFortiguardwfMapOutput
	ToSystemreplacemsgFortiguardwfMapOutputWithContext(context.Context) SystemreplacemsgFortiguardwfMapOutput
}

type SystemreplacemsgFortiguardwfMap map[string]SystemreplacemsgFortiguardwfInput

func (SystemreplacemsgFortiguardwfMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgFortiguardwf)(nil)).Elem()
}

func (i SystemreplacemsgFortiguardwfMap) ToSystemreplacemsgFortiguardwfMapOutput() SystemreplacemsgFortiguardwfMapOutput {
	return i.ToSystemreplacemsgFortiguardwfMapOutputWithContext(context.Background())
}

func (i SystemreplacemsgFortiguardwfMap) ToSystemreplacemsgFortiguardwfMapOutputWithContext(ctx context.Context) SystemreplacemsgFortiguardwfMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgFortiguardwfMapOutput)
}

type SystemreplacemsgFortiguardwfOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgFortiguardwfOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgFortiguardwf)(nil)).Elem()
}

func (o SystemreplacemsgFortiguardwfOutput) ToSystemreplacemsgFortiguardwfOutput() SystemreplacemsgFortiguardwfOutput {
	return o
}

func (o SystemreplacemsgFortiguardwfOutput) ToSystemreplacemsgFortiguardwfOutputWithContext(ctx context.Context) SystemreplacemsgFortiguardwfOutput {
	return o
}

// Message string.
func (o SystemreplacemsgFortiguardwfOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgFortiguardwf) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag.
func (o SystemreplacemsgFortiguardwfOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgFortiguardwf) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o SystemreplacemsgFortiguardwfOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgFortiguardwf) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o SystemreplacemsgFortiguardwfOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgFortiguardwf) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemreplacemsgFortiguardwfOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgFortiguardwf) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemreplacemsgFortiguardwfArrayOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgFortiguardwfArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgFortiguardwf)(nil)).Elem()
}

func (o SystemreplacemsgFortiguardwfArrayOutput) ToSystemreplacemsgFortiguardwfArrayOutput() SystemreplacemsgFortiguardwfArrayOutput {
	return o
}

func (o SystemreplacemsgFortiguardwfArrayOutput) ToSystemreplacemsgFortiguardwfArrayOutputWithContext(ctx context.Context) SystemreplacemsgFortiguardwfArrayOutput {
	return o
}

func (o SystemreplacemsgFortiguardwfArrayOutput) Index(i pulumi.IntInput) SystemreplacemsgFortiguardwfOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemreplacemsgFortiguardwf {
		return vs[0].([]*SystemreplacemsgFortiguardwf)[vs[1].(int)]
	}).(SystemreplacemsgFortiguardwfOutput)
}

type SystemreplacemsgFortiguardwfMapOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgFortiguardwfMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgFortiguardwf)(nil)).Elem()
}

func (o SystemreplacemsgFortiguardwfMapOutput) ToSystemreplacemsgFortiguardwfMapOutput() SystemreplacemsgFortiguardwfMapOutput {
	return o
}

func (o SystemreplacemsgFortiguardwfMapOutput) ToSystemreplacemsgFortiguardwfMapOutputWithContext(ctx context.Context) SystemreplacemsgFortiguardwfMapOutput {
	return o
}

func (o SystemreplacemsgFortiguardwfMapOutput) MapIndex(k pulumi.StringInput) SystemreplacemsgFortiguardwfOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemreplacemsgFortiguardwf {
		return vs[0].(map[string]*SystemreplacemsgFortiguardwf)[vs[1].(string)]
	}).(SystemreplacemsgFortiguardwfOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgFortiguardwfInput)(nil)).Elem(), &SystemreplacemsgFortiguardwf{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgFortiguardwfArrayInput)(nil)).Elem(), SystemreplacemsgFortiguardwfArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgFortiguardwfMapInput)(nil)).Elem(), SystemreplacemsgFortiguardwfMap{})
	pulumi.RegisterOutputType(SystemreplacemsgFortiguardwfOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgFortiguardwfArrayOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgFortiguardwfMapOutput{})
}
