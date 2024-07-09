// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure integrated NAC settings for FortiSwitch. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,6.4.15,7.0.0`.
//
// ## Import
//
// SwitchController NacSettings can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/nacsettings:Nacsettings labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/nacsettings:Nacsettings labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Nacsettings struct {
	pulumi.CustomResourceState

	// Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
	AutoAuth pulumi.StringOutput `pulumi:"autoAuth"`
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceNacPort pulumi.StringOutput `pulumi:"bounceNacPort"`
	// Time interval(minutes, 0 = no expiry) to be included in the inactive NAC devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
	InactiveTimer pulumi.IntOutput `pulumi:"inactiveTimer"`
	// Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush pulumi.StringOutput `pulumi:"linkDownFlush"`
	// Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// NAC settings name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Default NAC Onboarding VLAN when NAC devices are discovered.
	OnboardingVlan pulumi.StringOutput `pulumi:"onboardingVlan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewNacsettings registers a new resource with the given unique name, arguments, and options.
func NewNacsettings(ctx *pulumi.Context,
	name string, args *NacsettingsArgs, opts ...pulumi.ResourceOption) (*Nacsettings, error) {
	if args == nil {
		args = &NacsettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Nacsettings
	err := ctx.RegisterResource("fortios:switchcontroller/nacsettings:Nacsettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNacsettings gets an existing Nacsettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNacsettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NacsettingsState, opts ...pulumi.ResourceOption) (*Nacsettings, error) {
	var resource Nacsettings
	err := ctx.ReadResource("fortios:switchcontroller/nacsettings:Nacsettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Nacsettings resources.
type nacsettingsState struct {
	// Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
	AutoAuth *string `pulumi:"autoAuth"`
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceNacPort *string `pulumi:"bounceNacPort"`
	// Time interval(minutes, 0 = no expiry) to be included in the inactive NAC devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
	InactiveTimer *int `pulumi:"inactiveTimer"`
	// Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush *string `pulumi:"linkDownFlush"`
	// Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
	Mode *string `pulumi:"mode"`
	// NAC settings name.
	Name *string `pulumi:"name"`
	// Default NAC Onboarding VLAN when NAC devices are discovered.
	OnboardingVlan *string `pulumi:"onboardingVlan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type NacsettingsState struct {
	// Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
	AutoAuth pulumi.StringPtrInput
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceNacPort pulumi.StringPtrInput
	// Time interval(minutes, 0 = no expiry) to be included in the inactive NAC devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
	InactiveTimer pulumi.IntPtrInput
	// Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush pulumi.StringPtrInput
	// Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
	Mode pulumi.StringPtrInput
	// NAC settings name.
	Name pulumi.StringPtrInput
	// Default NAC Onboarding VLAN when NAC devices are discovered.
	OnboardingVlan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NacsettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*nacsettingsState)(nil)).Elem()
}

type nacsettingsArgs struct {
	// Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
	AutoAuth *string `pulumi:"autoAuth"`
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceNacPort *string `pulumi:"bounceNacPort"`
	// Time interval(minutes, 0 = no expiry) to be included in the inactive NAC devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
	InactiveTimer *int `pulumi:"inactiveTimer"`
	// Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush *string `pulumi:"linkDownFlush"`
	// Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
	Mode *string `pulumi:"mode"`
	// NAC settings name.
	Name *string `pulumi:"name"`
	// Default NAC Onboarding VLAN when NAC devices are discovered.
	OnboardingVlan *string `pulumi:"onboardingVlan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Nacsettings resource.
type NacsettingsArgs struct {
	// Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
	AutoAuth pulumi.StringPtrInput
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceNacPort pulumi.StringPtrInput
	// Time interval(minutes, 0 = no expiry) to be included in the inactive NAC devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
	InactiveTimer pulumi.IntPtrInput
	// Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush pulumi.StringPtrInput
	// Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
	Mode pulumi.StringPtrInput
	// NAC settings name.
	Name pulumi.StringPtrInput
	// Default NAC Onboarding VLAN when NAC devices are discovered.
	OnboardingVlan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NacsettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nacsettingsArgs)(nil)).Elem()
}

type NacsettingsInput interface {
	pulumi.Input

	ToNacsettingsOutput() NacsettingsOutput
	ToNacsettingsOutputWithContext(ctx context.Context) NacsettingsOutput
}

func (*Nacsettings) ElementType() reflect.Type {
	return reflect.TypeOf((**Nacsettings)(nil)).Elem()
}

func (i *Nacsettings) ToNacsettingsOutput() NacsettingsOutput {
	return i.ToNacsettingsOutputWithContext(context.Background())
}

func (i *Nacsettings) ToNacsettingsOutputWithContext(ctx context.Context) NacsettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NacsettingsOutput)
}

// NacsettingsArrayInput is an input type that accepts NacsettingsArray and NacsettingsArrayOutput values.
// You can construct a concrete instance of `NacsettingsArrayInput` via:
//
//	NacsettingsArray{ NacsettingsArgs{...} }
type NacsettingsArrayInput interface {
	pulumi.Input

	ToNacsettingsArrayOutput() NacsettingsArrayOutput
	ToNacsettingsArrayOutputWithContext(context.Context) NacsettingsArrayOutput
}

type NacsettingsArray []NacsettingsInput

func (NacsettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nacsettings)(nil)).Elem()
}

func (i NacsettingsArray) ToNacsettingsArrayOutput() NacsettingsArrayOutput {
	return i.ToNacsettingsArrayOutputWithContext(context.Background())
}

func (i NacsettingsArray) ToNacsettingsArrayOutputWithContext(ctx context.Context) NacsettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NacsettingsArrayOutput)
}

// NacsettingsMapInput is an input type that accepts NacsettingsMap and NacsettingsMapOutput values.
// You can construct a concrete instance of `NacsettingsMapInput` via:
//
//	NacsettingsMap{ "key": NacsettingsArgs{...} }
type NacsettingsMapInput interface {
	pulumi.Input

	ToNacsettingsMapOutput() NacsettingsMapOutput
	ToNacsettingsMapOutputWithContext(context.Context) NacsettingsMapOutput
}

type NacsettingsMap map[string]NacsettingsInput

func (NacsettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nacsettings)(nil)).Elem()
}

func (i NacsettingsMap) ToNacsettingsMapOutput() NacsettingsMapOutput {
	return i.ToNacsettingsMapOutputWithContext(context.Background())
}

func (i NacsettingsMap) ToNacsettingsMapOutputWithContext(ctx context.Context) NacsettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NacsettingsMapOutput)
}

type NacsettingsOutput struct{ *pulumi.OutputState }

func (NacsettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nacsettings)(nil)).Elem()
}

func (o NacsettingsOutput) ToNacsettingsOutput() NacsettingsOutput {
	return o
}

func (o NacsettingsOutput) ToNacsettingsOutputWithContext(ctx context.Context) NacsettingsOutput {
	return o
}

// Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
func (o NacsettingsOutput) AutoAuth() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacsettings) pulumi.StringOutput { return v.AutoAuth }).(pulumi.StringOutput)
}

// Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
func (o NacsettingsOutput) BounceNacPort() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacsettings) pulumi.StringOutput { return v.BounceNacPort }).(pulumi.StringOutput)
}

// Time interval(minutes, 0 = no expiry) to be included in the inactive NAC devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
func (o NacsettingsOutput) InactiveTimer() pulumi.IntOutput {
	return o.ApplyT(func(v *Nacsettings) pulumi.IntOutput { return v.InactiveTimer }).(pulumi.IntOutput)
}

// Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
func (o NacsettingsOutput) LinkDownFlush() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacsettings) pulumi.StringOutput { return v.LinkDownFlush }).(pulumi.StringOutput)
}

// Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
func (o NacsettingsOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacsettings) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// NAC settings name.
func (o NacsettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacsettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Default NAC Onboarding VLAN when NAC devices are discovered.
func (o NacsettingsOutput) OnboardingVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacsettings) pulumi.StringOutput { return v.OnboardingVlan }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o NacsettingsOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacsettings) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type NacsettingsArrayOutput struct{ *pulumi.OutputState }

func (NacsettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nacsettings)(nil)).Elem()
}

func (o NacsettingsArrayOutput) ToNacsettingsArrayOutput() NacsettingsArrayOutput {
	return o
}

func (o NacsettingsArrayOutput) ToNacsettingsArrayOutputWithContext(ctx context.Context) NacsettingsArrayOutput {
	return o
}

func (o NacsettingsArrayOutput) Index(i pulumi.IntInput) NacsettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Nacsettings {
		return vs[0].([]*Nacsettings)[vs[1].(int)]
	}).(NacsettingsOutput)
}

type NacsettingsMapOutput struct{ *pulumi.OutputState }

func (NacsettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nacsettings)(nil)).Elem()
}

func (o NacsettingsMapOutput) ToNacsettingsMapOutput() NacsettingsMapOutput {
	return o
}

func (o NacsettingsMapOutput) ToNacsettingsMapOutputWithContext(ctx context.Context) NacsettingsMapOutput {
	return o
}

func (o NacsettingsMapOutput) MapIndex(k pulumi.StringInput) NacsettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Nacsettings {
		return vs[0].(map[string]*Nacsettings)[vs[1].(string)]
	}).(NacsettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NacsettingsInput)(nil)).Elem(), &Nacsettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*NacsettingsArrayInput)(nil)).Elem(), NacsettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NacsettingsMapInput)(nil)).Elem(), NacsettingsMap{})
	pulumi.RegisterOutputType(NacsettingsOutput{})
	pulumi.RegisterOutputType(NacsettingsArrayOutput{})
	pulumi.RegisterOutputType(NacsettingsMapOutput{})
}
