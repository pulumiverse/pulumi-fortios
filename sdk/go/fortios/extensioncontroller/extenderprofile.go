// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package extensioncontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// FortiExtender extender profile configuration. Applies to FortiOS Version `>= 7.2.1`.
//
// ## Import
//
// # ExtensionController ExtenderProfile can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:extensioncontroller/extenderprofile:Extenderprofile labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:extensioncontroller/extenderprofile:Extenderprofile labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Extenderprofile struct {
	pulumi.CustomResourceState

	// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
	Allowaccess pulumi.StringOutput `pulumi:"allowaccess"`
	// FortiExtender LAN extension bandwidth limit (Mbps).
	BandwidthLimit pulumi.IntOutput `pulumi:"bandwidthLimit"`
	// FortiExtender cellular configuration. The structure of `cellular` block is documented below.
	Cellular ExtenderprofileCellularOutput `pulumi:"cellular"`
	// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
	EnforceBandwidth pulumi.StringOutput `pulumi:"enforceBandwidth"`
	// Extension option. Valid values: `wan-extension`, `lan-extension`.
	Extension pulumi.StringOutput `pulumi:"extension"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// FortiExtender lan extension configuration. The structure of `lanExtension` block is documented below.
	LanExtension ExtenderprofileLanExtensionOutput `pulumi:"lanExtension"`
	// Set the managed extender's administrator password.
	LoginPassword pulumi.StringPtrOutput `pulumi:"loginPassword"`
	// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
	LoginPasswordChange pulumi.StringOutput `pulumi:"loginPasswordChange"`
	// Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.
	Model pulumi.StringOutput `pulumi:"model"`
	// FortiExtender profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewExtenderprofile registers a new resource with the given unique name, arguments, and options.
func NewExtenderprofile(ctx *pulumi.Context,
	name string, args *ExtenderprofileArgs, opts ...pulumi.ResourceOption) (*Extenderprofile, error) {
	if args == nil {
		args = &ExtenderprofileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Extenderprofile
	err := ctx.RegisterResource("fortios:extensioncontroller/extenderprofile:Extenderprofile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtenderprofile gets an existing Extenderprofile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtenderprofile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtenderprofileState, opts ...pulumi.ResourceOption) (*Extenderprofile, error) {
	var resource Extenderprofile
	err := ctx.ReadResource("fortios:extensioncontroller/extenderprofile:Extenderprofile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Extenderprofile resources.
type extenderprofileState struct {
	// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
	Allowaccess *string `pulumi:"allowaccess"`
	// FortiExtender LAN extension bandwidth limit (Mbps).
	BandwidthLimit *int `pulumi:"bandwidthLimit"`
	// FortiExtender cellular configuration. The structure of `cellular` block is documented below.
	Cellular *ExtenderprofileCellular `pulumi:"cellular"`
	// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
	EnforceBandwidth *string `pulumi:"enforceBandwidth"`
	// Extension option. Valid values: `wan-extension`, `lan-extension`.
	Extension *string `pulumi:"extension"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// FortiExtender lan extension configuration. The structure of `lanExtension` block is documented below.
	LanExtension *ExtenderprofileLanExtension `pulumi:"lanExtension"`
	// Set the managed extender's administrator password.
	LoginPassword *string `pulumi:"loginPassword"`
	// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
	LoginPasswordChange *string `pulumi:"loginPasswordChange"`
	// Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.
	Model *string `pulumi:"model"`
	// FortiExtender profile name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ExtenderprofileState struct {
	// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
	Allowaccess pulumi.StringPtrInput
	// FortiExtender LAN extension bandwidth limit (Mbps).
	BandwidthLimit pulumi.IntPtrInput
	// FortiExtender cellular configuration. The structure of `cellular` block is documented below.
	Cellular ExtenderprofileCellularPtrInput
	// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
	EnforceBandwidth pulumi.StringPtrInput
	// Extension option. Valid values: `wan-extension`, `lan-extension`.
	Extension pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntPtrInput
	// FortiExtender lan extension configuration. The structure of `lanExtension` block is documented below.
	LanExtension ExtenderprofileLanExtensionPtrInput
	// Set the managed extender's administrator password.
	LoginPassword pulumi.StringPtrInput
	// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
	LoginPasswordChange pulumi.StringPtrInput
	// Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.
	Model pulumi.StringPtrInput
	// FortiExtender profile name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExtenderprofileState) ElementType() reflect.Type {
	return reflect.TypeOf((*extenderprofileState)(nil)).Elem()
}

type extenderprofileArgs struct {
	// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
	Allowaccess *string `pulumi:"allowaccess"`
	// FortiExtender LAN extension bandwidth limit (Mbps).
	BandwidthLimit *int `pulumi:"bandwidthLimit"`
	// FortiExtender cellular configuration. The structure of `cellular` block is documented below.
	Cellular *ExtenderprofileCellular `pulumi:"cellular"`
	// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
	EnforceBandwidth *string `pulumi:"enforceBandwidth"`
	// Extension option. Valid values: `wan-extension`, `lan-extension`.
	Extension *string `pulumi:"extension"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// FortiExtender lan extension configuration. The structure of `lanExtension` block is documented below.
	LanExtension *ExtenderprofileLanExtension `pulumi:"lanExtension"`
	// Set the managed extender's administrator password.
	LoginPassword *string `pulumi:"loginPassword"`
	// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
	LoginPasswordChange *string `pulumi:"loginPasswordChange"`
	// Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.
	Model *string `pulumi:"model"`
	// FortiExtender profile name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Extenderprofile resource.
type ExtenderprofileArgs struct {
	// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
	Allowaccess pulumi.StringPtrInput
	// FortiExtender LAN extension bandwidth limit (Mbps).
	BandwidthLimit pulumi.IntPtrInput
	// FortiExtender cellular configuration. The structure of `cellular` block is documented below.
	Cellular ExtenderprofileCellularPtrInput
	// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
	EnforceBandwidth pulumi.StringPtrInput
	// Extension option. Valid values: `wan-extension`, `lan-extension`.
	Extension pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntPtrInput
	// FortiExtender lan extension configuration. The structure of `lanExtension` block is documented below.
	LanExtension ExtenderprofileLanExtensionPtrInput
	// Set the managed extender's administrator password.
	LoginPassword pulumi.StringPtrInput
	// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
	LoginPasswordChange pulumi.StringPtrInput
	// Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.
	Model pulumi.StringPtrInput
	// FortiExtender profile name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExtenderprofileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extenderprofileArgs)(nil)).Elem()
}

type ExtenderprofileInput interface {
	pulumi.Input

	ToExtenderprofileOutput() ExtenderprofileOutput
	ToExtenderprofileOutputWithContext(ctx context.Context) ExtenderprofileOutput
}

func (*Extenderprofile) ElementType() reflect.Type {
	return reflect.TypeOf((**Extenderprofile)(nil)).Elem()
}

func (i *Extenderprofile) ToExtenderprofileOutput() ExtenderprofileOutput {
	return i.ToExtenderprofileOutputWithContext(context.Background())
}

func (i *Extenderprofile) ToExtenderprofileOutputWithContext(ctx context.Context) ExtenderprofileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtenderprofileOutput)
}

// ExtenderprofileArrayInput is an input type that accepts ExtenderprofileArray and ExtenderprofileArrayOutput values.
// You can construct a concrete instance of `ExtenderprofileArrayInput` via:
//
//	ExtenderprofileArray{ ExtenderprofileArgs{...} }
type ExtenderprofileArrayInput interface {
	pulumi.Input

	ToExtenderprofileArrayOutput() ExtenderprofileArrayOutput
	ToExtenderprofileArrayOutputWithContext(context.Context) ExtenderprofileArrayOutput
}

type ExtenderprofileArray []ExtenderprofileInput

func (ExtenderprofileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Extenderprofile)(nil)).Elem()
}

func (i ExtenderprofileArray) ToExtenderprofileArrayOutput() ExtenderprofileArrayOutput {
	return i.ToExtenderprofileArrayOutputWithContext(context.Background())
}

func (i ExtenderprofileArray) ToExtenderprofileArrayOutputWithContext(ctx context.Context) ExtenderprofileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtenderprofileArrayOutput)
}

// ExtenderprofileMapInput is an input type that accepts ExtenderprofileMap and ExtenderprofileMapOutput values.
// You can construct a concrete instance of `ExtenderprofileMapInput` via:
//
//	ExtenderprofileMap{ "key": ExtenderprofileArgs{...} }
type ExtenderprofileMapInput interface {
	pulumi.Input

	ToExtenderprofileMapOutput() ExtenderprofileMapOutput
	ToExtenderprofileMapOutputWithContext(context.Context) ExtenderprofileMapOutput
}

type ExtenderprofileMap map[string]ExtenderprofileInput

func (ExtenderprofileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Extenderprofile)(nil)).Elem()
}

func (i ExtenderprofileMap) ToExtenderprofileMapOutput() ExtenderprofileMapOutput {
	return i.ToExtenderprofileMapOutputWithContext(context.Background())
}

func (i ExtenderprofileMap) ToExtenderprofileMapOutputWithContext(ctx context.Context) ExtenderprofileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtenderprofileMapOutput)
}

type ExtenderprofileOutput struct{ *pulumi.OutputState }

func (ExtenderprofileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Extenderprofile)(nil)).Elem()
}

func (o ExtenderprofileOutput) ToExtenderprofileOutput() ExtenderprofileOutput {
	return o
}

func (o ExtenderprofileOutput) ToExtenderprofileOutputWithContext(ctx context.Context) ExtenderprofileOutput {
	return o
}

// Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
func (o ExtenderprofileOutput) Allowaccess() pulumi.StringOutput {
	return o.ApplyT(func(v *Extenderprofile) pulumi.StringOutput { return v.Allowaccess }).(pulumi.StringOutput)
}

// FortiExtender LAN extension bandwidth limit (Mbps).
func (o ExtenderprofileOutput) BandwidthLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Extenderprofile) pulumi.IntOutput { return v.BandwidthLimit }).(pulumi.IntOutput)
}

// FortiExtender cellular configuration. The structure of `cellular` block is documented below.
func (o ExtenderprofileOutput) Cellular() ExtenderprofileCellularOutput {
	return o.ApplyT(func(v *Extenderprofile) ExtenderprofileCellularOutput { return v.Cellular }).(ExtenderprofileCellularOutput)
}

// Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
func (o ExtenderprofileOutput) EnforceBandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *Extenderprofile) pulumi.StringOutput { return v.EnforceBandwidth }).(pulumi.StringOutput)
}

// Extension option. Valid values: `wan-extension`, `lan-extension`.
func (o ExtenderprofileOutput) Extension() pulumi.StringOutput {
	return o.ApplyT(func(v *Extenderprofile) pulumi.StringOutput { return v.Extension }).(pulumi.StringOutput)
}

// ID.
func (o ExtenderprofileOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Extenderprofile) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// FortiExtender lan extension configuration. The structure of `lanExtension` block is documented below.
func (o ExtenderprofileOutput) LanExtension() ExtenderprofileLanExtensionOutput {
	return o.ApplyT(func(v *Extenderprofile) ExtenderprofileLanExtensionOutput { return v.LanExtension }).(ExtenderprofileLanExtensionOutput)
}

// Set the managed extender's administrator password.
func (o ExtenderprofileOutput) LoginPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extenderprofile) pulumi.StringPtrOutput { return v.LoginPassword }).(pulumi.StringPtrOutput)
}

// Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
func (o ExtenderprofileOutput) LoginPasswordChange() pulumi.StringOutput {
	return o.ApplyT(func(v *Extenderprofile) pulumi.StringOutput { return v.LoginPasswordChange }).(pulumi.StringOutput)
}

// Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.
func (o ExtenderprofileOutput) Model() pulumi.StringOutput {
	return o.ApplyT(func(v *Extenderprofile) pulumi.StringOutput { return v.Model }).(pulumi.StringOutput)
}

// FortiExtender profile name.
func (o ExtenderprofileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Extenderprofile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ExtenderprofileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extenderprofile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ExtenderprofileArrayOutput struct{ *pulumi.OutputState }

func (ExtenderprofileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Extenderprofile)(nil)).Elem()
}

func (o ExtenderprofileArrayOutput) ToExtenderprofileArrayOutput() ExtenderprofileArrayOutput {
	return o
}

func (o ExtenderprofileArrayOutput) ToExtenderprofileArrayOutputWithContext(ctx context.Context) ExtenderprofileArrayOutput {
	return o
}

func (o ExtenderprofileArrayOutput) Index(i pulumi.IntInput) ExtenderprofileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Extenderprofile {
		return vs[0].([]*Extenderprofile)[vs[1].(int)]
	}).(ExtenderprofileOutput)
}

type ExtenderprofileMapOutput struct{ *pulumi.OutputState }

func (ExtenderprofileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Extenderprofile)(nil)).Elem()
}

func (o ExtenderprofileMapOutput) ToExtenderprofileMapOutput() ExtenderprofileMapOutput {
	return o
}

func (o ExtenderprofileMapOutput) ToExtenderprofileMapOutputWithContext(ctx context.Context) ExtenderprofileMapOutput {
	return o
}

func (o ExtenderprofileMapOutput) MapIndex(k pulumi.StringInput) ExtenderprofileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Extenderprofile {
		return vs[0].(map[string]*Extenderprofile)[vs[1].(string)]
	}).(ExtenderprofileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExtenderprofileInput)(nil)).Elem(), &Extenderprofile{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtenderprofileArrayInput)(nil)).Elem(), ExtenderprofileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtenderprofileMapInput)(nil)).Elem(), ExtenderprofileMap{})
	pulumi.RegisterOutputType(ExtenderprofileOutput{})
	pulumi.RegisterOutputType(ExtenderprofileArrayOutput{})
	pulumi.RegisterOutputType(ExtenderprofileMapOutput{})
}
