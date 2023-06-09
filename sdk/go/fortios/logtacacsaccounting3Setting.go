// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings for TACACS+ accounting. Applies to FortiOS Version `>= 7.0.2`.
//
// ## Import
//
// # LogTacacsAccounting3 Setting can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/logtacacsaccounting3Setting:Logtacacsaccounting3Setting labelname LogTacacsAccounting3Setting
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/logtacacsaccounting3Setting:Logtacacsaccounting3Setting labelname LogTacacsAccounting3Setting
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Logtacacsaccounting3Setting struct {
	pulumi.CustomResourceState

	// Address of TACACS+ server.
	Server pulumi.StringOutput `pulumi:"server"`
	// Key to access the TACACS+ server.
	ServerKey pulumi.StringPtrOutput `pulumi:"serverKey"`
	// Enable/disable TACACS+ accounting. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogtacacsaccounting3Setting registers a new resource with the given unique name, arguments, and options.
func NewLogtacacsaccounting3Setting(ctx *pulumi.Context,
	name string, args *Logtacacsaccounting3SettingArgs, opts ...pulumi.ResourceOption) (*Logtacacsaccounting3Setting, error) {
	if args == nil {
		args = &Logtacacsaccounting3SettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Logtacacsaccounting3Setting
	err := ctx.RegisterResource("fortios:index/logtacacsaccounting3Setting:Logtacacsaccounting3Setting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogtacacsaccounting3Setting gets an existing Logtacacsaccounting3Setting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogtacacsaccounting3Setting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Logtacacsaccounting3SettingState, opts ...pulumi.ResourceOption) (*Logtacacsaccounting3Setting, error) {
	var resource Logtacacsaccounting3Setting
	err := ctx.ReadResource("fortios:index/logtacacsaccounting3Setting:Logtacacsaccounting3Setting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Logtacacsaccounting3Setting resources.
type logtacacsaccounting3SettingState struct {
	// Address of TACACS+ server.
	Server *string `pulumi:"server"`
	// Key to access the TACACS+ server.
	ServerKey *string `pulumi:"serverKey"`
	// Enable/disable TACACS+ accounting. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Logtacacsaccounting3SettingState struct {
	// Address of TACACS+ server.
	Server pulumi.StringPtrInput
	// Key to access the TACACS+ server.
	ServerKey pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Logtacacsaccounting3SettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logtacacsaccounting3SettingState)(nil)).Elem()
}

type logtacacsaccounting3SettingArgs struct {
	// Address of TACACS+ server.
	Server *string `pulumi:"server"`
	// Key to access the TACACS+ server.
	ServerKey *string `pulumi:"serverKey"`
	// Enable/disable TACACS+ accounting. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Logtacacsaccounting3Setting resource.
type Logtacacsaccounting3SettingArgs struct {
	// Address of TACACS+ server.
	Server pulumi.StringPtrInput
	// Key to access the TACACS+ server.
	ServerKey pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Logtacacsaccounting3SettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logtacacsaccounting3SettingArgs)(nil)).Elem()
}

type Logtacacsaccounting3SettingInput interface {
	pulumi.Input

	ToLogtacacsaccounting3SettingOutput() Logtacacsaccounting3SettingOutput
	ToLogtacacsaccounting3SettingOutputWithContext(ctx context.Context) Logtacacsaccounting3SettingOutput
}

func (*Logtacacsaccounting3Setting) ElementType() reflect.Type {
	return reflect.TypeOf((**Logtacacsaccounting3Setting)(nil)).Elem()
}

func (i *Logtacacsaccounting3Setting) ToLogtacacsaccounting3SettingOutput() Logtacacsaccounting3SettingOutput {
	return i.ToLogtacacsaccounting3SettingOutputWithContext(context.Background())
}

func (i *Logtacacsaccounting3Setting) ToLogtacacsaccounting3SettingOutputWithContext(ctx context.Context) Logtacacsaccounting3SettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Logtacacsaccounting3SettingOutput)
}

// Logtacacsaccounting3SettingArrayInput is an input type that accepts Logtacacsaccounting3SettingArray and Logtacacsaccounting3SettingArrayOutput values.
// You can construct a concrete instance of `Logtacacsaccounting3SettingArrayInput` via:
//
//	Logtacacsaccounting3SettingArray{ Logtacacsaccounting3SettingArgs{...} }
type Logtacacsaccounting3SettingArrayInput interface {
	pulumi.Input

	ToLogtacacsaccounting3SettingArrayOutput() Logtacacsaccounting3SettingArrayOutput
	ToLogtacacsaccounting3SettingArrayOutputWithContext(context.Context) Logtacacsaccounting3SettingArrayOutput
}

type Logtacacsaccounting3SettingArray []Logtacacsaccounting3SettingInput

func (Logtacacsaccounting3SettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Logtacacsaccounting3Setting)(nil)).Elem()
}

func (i Logtacacsaccounting3SettingArray) ToLogtacacsaccounting3SettingArrayOutput() Logtacacsaccounting3SettingArrayOutput {
	return i.ToLogtacacsaccounting3SettingArrayOutputWithContext(context.Background())
}

func (i Logtacacsaccounting3SettingArray) ToLogtacacsaccounting3SettingArrayOutputWithContext(ctx context.Context) Logtacacsaccounting3SettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Logtacacsaccounting3SettingArrayOutput)
}

// Logtacacsaccounting3SettingMapInput is an input type that accepts Logtacacsaccounting3SettingMap and Logtacacsaccounting3SettingMapOutput values.
// You can construct a concrete instance of `Logtacacsaccounting3SettingMapInput` via:
//
//	Logtacacsaccounting3SettingMap{ "key": Logtacacsaccounting3SettingArgs{...} }
type Logtacacsaccounting3SettingMapInput interface {
	pulumi.Input

	ToLogtacacsaccounting3SettingMapOutput() Logtacacsaccounting3SettingMapOutput
	ToLogtacacsaccounting3SettingMapOutputWithContext(context.Context) Logtacacsaccounting3SettingMapOutput
}

type Logtacacsaccounting3SettingMap map[string]Logtacacsaccounting3SettingInput

func (Logtacacsaccounting3SettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Logtacacsaccounting3Setting)(nil)).Elem()
}

func (i Logtacacsaccounting3SettingMap) ToLogtacacsaccounting3SettingMapOutput() Logtacacsaccounting3SettingMapOutput {
	return i.ToLogtacacsaccounting3SettingMapOutputWithContext(context.Background())
}

func (i Logtacacsaccounting3SettingMap) ToLogtacacsaccounting3SettingMapOutputWithContext(ctx context.Context) Logtacacsaccounting3SettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Logtacacsaccounting3SettingMapOutput)
}

type Logtacacsaccounting3SettingOutput struct{ *pulumi.OutputState }

func (Logtacacsaccounting3SettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Logtacacsaccounting3Setting)(nil)).Elem()
}

func (o Logtacacsaccounting3SettingOutput) ToLogtacacsaccounting3SettingOutput() Logtacacsaccounting3SettingOutput {
	return o
}

func (o Logtacacsaccounting3SettingOutput) ToLogtacacsaccounting3SettingOutputWithContext(ctx context.Context) Logtacacsaccounting3SettingOutput {
	return o
}

// Address of TACACS+ server.
func (o Logtacacsaccounting3SettingOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *Logtacacsaccounting3Setting) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Key to access the TACACS+ server.
func (o Logtacacsaccounting3SettingOutput) ServerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Logtacacsaccounting3Setting) pulumi.StringPtrOutput { return v.ServerKey }).(pulumi.StringPtrOutput)
}

// Enable/disable TACACS+ accounting. Valid values: `enable`, `disable`.
func (o Logtacacsaccounting3SettingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Logtacacsaccounting3Setting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Logtacacsaccounting3SettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Logtacacsaccounting3Setting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Logtacacsaccounting3SettingArrayOutput struct{ *pulumi.OutputState }

func (Logtacacsaccounting3SettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Logtacacsaccounting3Setting)(nil)).Elem()
}

func (o Logtacacsaccounting3SettingArrayOutput) ToLogtacacsaccounting3SettingArrayOutput() Logtacacsaccounting3SettingArrayOutput {
	return o
}

func (o Logtacacsaccounting3SettingArrayOutput) ToLogtacacsaccounting3SettingArrayOutputWithContext(ctx context.Context) Logtacacsaccounting3SettingArrayOutput {
	return o
}

func (o Logtacacsaccounting3SettingArrayOutput) Index(i pulumi.IntInput) Logtacacsaccounting3SettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Logtacacsaccounting3Setting {
		return vs[0].([]*Logtacacsaccounting3Setting)[vs[1].(int)]
	}).(Logtacacsaccounting3SettingOutput)
}

type Logtacacsaccounting3SettingMapOutput struct{ *pulumi.OutputState }

func (Logtacacsaccounting3SettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Logtacacsaccounting3Setting)(nil)).Elem()
}

func (o Logtacacsaccounting3SettingMapOutput) ToLogtacacsaccounting3SettingMapOutput() Logtacacsaccounting3SettingMapOutput {
	return o
}

func (o Logtacacsaccounting3SettingMapOutput) ToLogtacacsaccounting3SettingMapOutputWithContext(ctx context.Context) Logtacacsaccounting3SettingMapOutput {
	return o
}

func (o Logtacacsaccounting3SettingMapOutput) MapIndex(k pulumi.StringInput) Logtacacsaccounting3SettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Logtacacsaccounting3Setting {
		return vs[0].(map[string]*Logtacacsaccounting3Setting)[vs[1].(string)]
	}).(Logtacacsaccounting3SettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Logtacacsaccounting3SettingInput)(nil)).Elem(), &Logtacacsaccounting3Setting{})
	pulumi.RegisterInputType(reflect.TypeOf((*Logtacacsaccounting3SettingArrayInput)(nil)).Elem(), Logtacacsaccounting3SettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Logtacacsaccounting3SettingMapInput)(nil)).Elem(), Logtacacsaccounting3SettingMap{})
	pulumi.RegisterOutputType(Logtacacsaccounting3SettingOutput{})
	pulumi.RegisterOutputType(Logtacacsaccounting3SettingArrayOutput{})
	pulumi.RegisterOutputType(Logtacacsaccounting3SettingMapOutput{})
}
