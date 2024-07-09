// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure global MAC synchronization settings. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Import
//
// SwitchController MacSyncSettings can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/macsyncsettings:Macsyncsettings labelname SwitchControllerMacSyncSettings
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/macsyncsettings:Macsyncsettings labelname SwitchControllerMacSyncSettings
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Macsyncsettings struct {
	pulumi.CustomResourceState

	// Time interval between MAC synchronizations (30 - 1800 sec, default = 60, 0 = disable MAC synchronization).
	MacSyncInterval pulumi.IntOutput `pulumi:"macSyncInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewMacsyncsettings registers a new resource with the given unique name, arguments, and options.
func NewMacsyncsettings(ctx *pulumi.Context,
	name string, args *MacsyncsettingsArgs, opts ...pulumi.ResourceOption) (*Macsyncsettings, error) {
	if args == nil {
		args = &MacsyncsettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Macsyncsettings
	err := ctx.RegisterResource("fortios:switchcontroller/macsyncsettings:Macsyncsettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMacsyncsettings gets an existing Macsyncsettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMacsyncsettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MacsyncsettingsState, opts ...pulumi.ResourceOption) (*Macsyncsettings, error) {
	var resource Macsyncsettings
	err := ctx.ReadResource("fortios:switchcontroller/macsyncsettings:Macsyncsettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Macsyncsettings resources.
type macsyncsettingsState struct {
	// Time interval between MAC synchronizations (30 - 1800 sec, default = 60, 0 = disable MAC synchronization).
	MacSyncInterval *int `pulumi:"macSyncInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type MacsyncsettingsState struct {
	// Time interval between MAC synchronizations (30 - 1800 sec, default = 60, 0 = disable MAC synchronization).
	MacSyncInterval pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (MacsyncsettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*macsyncsettingsState)(nil)).Elem()
}

type macsyncsettingsArgs struct {
	// Time interval between MAC synchronizations (30 - 1800 sec, default = 60, 0 = disable MAC synchronization).
	MacSyncInterval *int `pulumi:"macSyncInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Macsyncsettings resource.
type MacsyncsettingsArgs struct {
	// Time interval between MAC synchronizations (30 - 1800 sec, default = 60, 0 = disable MAC synchronization).
	MacSyncInterval pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (MacsyncsettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*macsyncsettingsArgs)(nil)).Elem()
}

type MacsyncsettingsInput interface {
	pulumi.Input

	ToMacsyncsettingsOutput() MacsyncsettingsOutput
	ToMacsyncsettingsOutputWithContext(ctx context.Context) MacsyncsettingsOutput
}

func (*Macsyncsettings) ElementType() reflect.Type {
	return reflect.TypeOf((**Macsyncsettings)(nil)).Elem()
}

func (i *Macsyncsettings) ToMacsyncsettingsOutput() MacsyncsettingsOutput {
	return i.ToMacsyncsettingsOutputWithContext(context.Background())
}

func (i *Macsyncsettings) ToMacsyncsettingsOutputWithContext(ctx context.Context) MacsyncsettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MacsyncsettingsOutput)
}

// MacsyncsettingsArrayInput is an input type that accepts MacsyncsettingsArray and MacsyncsettingsArrayOutput values.
// You can construct a concrete instance of `MacsyncsettingsArrayInput` via:
//
//	MacsyncsettingsArray{ MacsyncsettingsArgs{...} }
type MacsyncsettingsArrayInput interface {
	pulumi.Input

	ToMacsyncsettingsArrayOutput() MacsyncsettingsArrayOutput
	ToMacsyncsettingsArrayOutputWithContext(context.Context) MacsyncsettingsArrayOutput
}

type MacsyncsettingsArray []MacsyncsettingsInput

func (MacsyncsettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Macsyncsettings)(nil)).Elem()
}

func (i MacsyncsettingsArray) ToMacsyncsettingsArrayOutput() MacsyncsettingsArrayOutput {
	return i.ToMacsyncsettingsArrayOutputWithContext(context.Background())
}

func (i MacsyncsettingsArray) ToMacsyncsettingsArrayOutputWithContext(ctx context.Context) MacsyncsettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MacsyncsettingsArrayOutput)
}

// MacsyncsettingsMapInput is an input type that accepts MacsyncsettingsMap and MacsyncsettingsMapOutput values.
// You can construct a concrete instance of `MacsyncsettingsMapInput` via:
//
//	MacsyncsettingsMap{ "key": MacsyncsettingsArgs{...} }
type MacsyncsettingsMapInput interface {
	pulumi.Input

	ToMacsyncsettingsMapOutput() MacsyncsettingsMapOutput
	ToMacsyncsettingsMapOutputWithContext(context.Context) MacsyncsettingsMapOutput
}

type MacsyncsettingsMap map[string]MacsyncsettingsInput

func (MacsyncsettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Macsyncsettings)(nil)).Elem()
}

func (i MacsyncsettingsMap) ToMacsyncsettingsMapOutput() MacsyncsettingsMapOutput {
	return i.ToMacsyncsettingsMapOutputWithContext(context.Background())
}

func (i MacsyncsettingsMap) ToMacsyncsettingsMapOutputWithContext(ctx context.Context) MacsyncsettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MacsyncsettingsMapOutput)
}

type MacsyncsettingsOutput struct{ *pulumi.OutputState }

func (MacsyncsettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Macsyncsettings)(nil)).Elem()
}

func (o MacsyncsettingsOutput) ToMacsyncsettingsOutput() MacsyncsettingsOutput {
	return o
}

func (o MacsyncsettingsOutput) ToMacsyncsettingsOutputWithContext(ctx context.Context) MacsyncsettingsOutput {
	return o
}

// Time interval between MAC synchronizations (30 - 1800 sec, default = 60, 0 = disable MAC synchronization).
func (o MacsyncsettingsOutput) MacSyncInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Macsyncsettings) pulumi.IntOutput { return v.MacSyncInterval }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o MacsyncsettingsOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Macsyncsettings) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type MacsyncsettingsArrayOutput struct{ *pulumi.OutputState }

func (MacsyncsettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Macsyncsettings)(nil)).Elem()
}

func (o MacsyncsettingsArrayOutput) ToMacsyncsettingsArrayOutput() MacsyncsettingsArrayOutput {
	return o
}

func (o MacsyncsettingsArrayOutput) ToMacsyncsettingsArrayOutputWithContext(ctx context.Context) MacsyncsettingsArrayOutput {
	return o
}

func (o MacsyncsettingsArrayOutput) Index(i pulumi.IntInput) MacsyncsettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Macsyncsettings {
		return vs[0].([]*Macsyncsettings)[vs[1].(int)]
	}).(MacsyncsettingsOutput)
}

type MacsyncsettingsMapOutput struct{ *pulumi.OutputState }

func (MacsyncsettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Macsyncsettings)(nil)).Elem()
}

func (o MacsyncsettingsMapOutput) ToMacsyncsettingsMapOutput() MacsyncsettingsMapOutput {
	return o
}

func (o MacsyncsettingsMapOutput) ToMacsyncsettingsMapOutputWithContext(ctx context.Context) MacsyncsettingsMapOutput {
	return o
}

func (o MacsyncsettingsMapOutput) MapIndex(k pulumi.StringInput) MacsyncsettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Macsyncsettings {
		return vs[0].(map[string]*Macsyncsettings)[vs[1].(string)]
	}).(MacsyncsettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MacsyncsettingsInput)(nil)).Elem(), &Macsyncsettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*MacsyncsettingsArrayInput)(nil)).Elem(), MacsyncsettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MacsyncsettingsMapInput)(nil)).Elem(), MacsyncsettingsMap{})
	pulumi.RegisterOutputType(MacsyncsettingsOutput{})
	pulumi.RegisterOutputType(MacsyncsettingsArrayOutput{})
	pulumi.RegisterOutputType(MacsyncsettingsMapOutput{})
}
