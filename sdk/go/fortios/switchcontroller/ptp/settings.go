// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ptp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Global PTP settings. Applies to FortiOS Version `>= 6.4.2`.
//
// ## Import
//
// SwitchControllerPtp Settings can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/ptp/settings:Settings labelname SwitchControllerPtpSettings
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/ptp/settings:Settings labelname SwitchControllerPtpSettings
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Settings struct {
	pulumi.CustomResourceState

	// Enable/disable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSettings registers a new resource with the given unique name, arguments, and options.
func NewSettings(ctx *pulumi.Context,
	name string, args *SettingsArgs, opts ...pulumi.ResourceOption) (*Settings, error) {
	if args == nil {
		args = &SettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Settings
	err := ctx.RegisterResource("fortios:switchcontroller/ptp/settings:Settings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSettings gets an existing Settings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SettingsState, opts ...pulumi.ResourceOption) (*Settings, error) {
	var resource Settings
	err := ctx.ReadResource("fortios:switchcontroller/ptp/settings:Settings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Settings resources.
type settingsState struct {
	// Enable/disable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.
	Mode *string `pulumi:"mode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SettingsState struct {
	// Enable/disable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.
	Mode pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*settingsState)(nil)).Elem()
}

type settingsArgs struct {
	// Enable/disable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.
	Mode *string `pulumi:"mode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Settings resource.
type SettingsArgs struct {
	// Enable/disable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.
	Mode pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*settingsArgs)(nil)).Elem()
}

type SettingsInput interface {
	pulumi.Input

	ToSettingsOutput() SettingsOutput
	ToSettingsOutputWithContext(ctx context.Context) SettingsOutput
}

func (*Settings) ElementType() reflect.Type {
	return reflect.TypeOf((**Settings)(nil)).Elem()
}

func (i *Settings) ToSettingsOutput() SettingsOutput {
	return i.ToSettingsOutputWithContext(context.Background())
}

func (i *Settings) ToSettingsOutputWithContext(ctx context.Context) SettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsOutput)
}

// SettingsArrayInput is an input type that accepts SettingsArray and SettingsArrayOutput values.
// You can construct a concrete instance of `SettingsArrayInput` via:
//
//	SettingsArray{ SettingsArgs{...} }
type SettingsArrayInput interface {
	pulumi.Input

	ToSettingsArrayOutput() SettingsArrayOutput
	ToSettingsArrayOutputWithContext(context.Context) SettingsArrayOutput
}

type SettingsArray []SettingsInput

func (SettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Settings)(nil)).Elem()
}

func (i SettingsArray) ToSettingsArrayOutput() SettingsArrayOutput {
	return i.ToSettingsArrayOutputWithContext(context.Background())
}

func (i SettingsArray) ToSettingsArrayOutputWithContext(ctx context.Context) SettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsArrayOutput)
}

// SettingsMapInput is an input type that accepts SettingsMap and SettingsMapOutput values.
// You can construct a concrete instance of `SettingsMapInput` via:
//
//	SettingsMap{ "key": SettingsArgs{...} }
type SettingsMapInput interface {
	pulumi.Input

	ToSettingsMapOutput() SettingsMapOutput
	ToSettingsMapOutputWithContext(context.Context) SettingsMapOutput
}

type SettingsMap map[string]SettingsInput

func (SettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Settings)(nil)).Elem()
}

func (i SettingsMap) ToSettingsMapOutput() SettingsMapOutput {
	return i.ToSettingsMapOutputWithContext(context.Background())
}

func (i SettingsMap) ToSettingsMapOutputWithContext(ctx context.Context) SettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsMapOutput)
}

type SettingsOutput struct{ *pulumi.OutputState }

func (SettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Settings)(nil)).Elem()
}

func (o SettingsOutput) ToSettingsOutput() SettingsOutput {
	return o
}

func (o SettingsOutput) ToSettingsOutputWithContext(ctx context.Context) SettingsOutput {
	return o
}

// Enable/disable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.
func (o SettingsOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *Settings) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Settings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SettingsArrayOutput struct{ *pulumi.OutputState }

func (SettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Settings)(nil)).Elem()
}

func (o SettingsArrayOutput) ToSettingsArrayOutput() SettingsArrayOutput {
	return o
}

func (o SettingsArrayOutput) ToSettingsArrayOutputWithContext(ctx context.Context) SettingsArrayOutput {
	return o
}

func (o SettingsArrayOutput) Index(i pulumi.IntInput) SettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Settings {
		return vs[0].([]*Settings)[vs[1].(int)]
	}).(SettingsOutput)
}

type SettingsMapOutput struct{ *pulumi.OutputState }

func (SettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Settings)(nil)).Elem()
}

func (o SettingsMapOutput) ToSettingsMapOutput() SettingsMapOutput {
	return o
}

func (o SettingsMapOutput) ToSettingsMapOutputWithContext(ctx context.Context) SettingsMapOutput {
	return o
}

func (o SettingsMapOutput) MapIndex(k pulumi.StringInput) SettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Settings {
		return vs[0].(map[string]*Settings)[vs[1].(string)]
	}).(SettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsInput)(nil)).Elem(), &Settings{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsArrayInput)(nil)).Elem(), SettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsMapInput)(nil)).Elem(), SettingsMap{})
	pulumi.RegisterOutputType(SettingsOutput{})
	pulumi.RegisterOutputType(SettingsArrayOutput{})
	pulumi.RegisterOutputType(SettingsMapOutput{})
}
