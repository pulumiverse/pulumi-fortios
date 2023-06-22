// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure initial template for auto-generated VLAN interfaces. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// # SwitchControllerInitialConfig Vlans can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerinitialconfigVlans:SwitchcontrollerinitialconfigVlans labelname SwitchControllerInitialConfigVlans
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerinitialconfigVlans:SwitchcontrollerinitialconfigVlans labelname SwitchControllerInitialConfigVlans
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SwitchcontrollerinitialconfigVlans struct {
	pulumi.CustomResourceState

	// Default VLAN (native) assigned to all switch ports upon discovery.
	DefaultVlan pulumi.StringOutput `pulumi:"defaultVlan"`
	// VLAN for NAC onboarding devices.
	Nac pulumi.StringOutput `pulumi:"nac"`
	// VLAN for NAC segemnt primary interface.
	NacSegment pulumi.StringOutput `pulumi:"nacSegment"`
	// VLAN for quarantined traffic.
	Quarantine pulumi.StringOutput `pulumi:"quarantine"`
	// VLAN for RSPAN/ERSPAN mirrored traffic.
	Rspan pulumi.StringOutput `pulumi:"rspan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// VLAN dedicated for video devices.
	Video pulumi.StringOutput `pulumi:"video"`
	// VLAN dedicated for voice devices.
	Voice pulumi.StringOutput `pulumi:"voice"`
}

// NewSwitchcontrollerinitialconfigVlans registers a new resource with the given unique name, arguments, and options.
func NewSwitchcontrollerinitialconfigVlans(ctx *pulumi.Context,
	name string, args *SwitchcontrollerinitialconfigVlansArgs, opts ...pulumi.ResourceOption) (*SwitchcontrollerinitialconfigVlans, error) {
	if args == nil {
		args = &SwitchcontrollerinitialconfigVlansArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchcontrollerinitialconfigVlans
	err := ctx.RegisterResource("fortios:index/switchcontrollerinitialconfigVlans:SwitchcontrollerinitialconfigVlans", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchcontrollerinitialconfigVlans gets an existing SwitchcontrollerinitialconfigVlans resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchcontrollerinitialconfigVlans(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchcontrollerinitialconfigVlansState, opts ...pulumi.ResourceOption) (*SwitchcontrollerinitialconfigVlans, error) {
	var resource SwitchcontrollerinitialconfigVlans
	err := ctx.ReadResource("fortios:index/switchcontrollerinitialconfigVlans:SwitchcontrollerinitialconfigVlans", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchcontrollerinitialconfigVlans resources.
type switchcontrollerinitialconfigVlansState struct {
	// Default VLAN (native) assigned to all switch ports upon discovery.
	DefaultVlan *string `pulumi:"defaultVlan"`
	// VLAN for NAC onboarding devices.
	Nac *string `pulumi:"nac"`
	// VLAN for NAC segemnt primary interface.
	NacSegment *string `pulumi:"nacSegment"`
	// VLAN for quarantined traffic.
	Quarantine *string `pulumi:"quarantine"`
	// VLAN for RSPAN/ERSPAN mirrored traffic.
	Rspan *string `pulumi:"rspan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN dedicated for video devices.
	Video *string `pulumi:"video"`
	// VLAN dedicated for voice devices.
	Voice *string `pulumi:"voice"`
}

type SwitchcontrollerinitialconfigVlansState struct {
	// Default VLAN (native) assigned to all switch ports upon discovery.
	DefaultVlan pulumi.StringPtrInput
	// VLAN for NAC onboarding devices.
	Nac pulumi.StringPtrInput
	// VLAN for NAC segemnt primary interface.
	NacSegment pulumi.StringPtrInput
	// VLAN for quarantined traffic.
	Quarantine pulumi.StringPtrInput
	// VLAN for RSPAN/ERSPAN mirrored traffic.
	Rspan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN dedicated for video devices.
	Video pulumi.StringPtrInput
	// VLAN dedicated for voice devices.
	Voice pulumi.StringPtrInput
}

func (SwitchcontrollerinitialconfigVlansState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerinitialconfigVlansState)(nil)).Elem()
}

type switchcontrollerinitialconfigVlansArgs struct {
	// Default VLAN (native) assigned to all switch ports upon discovery.
	DefaultVlan *string `pulumi:"defaultVlan"`
	// VLAN for NAC onboarding devices.
	Nac *string `pulumi:"nac"`
	// VLAN for NAC segemnt primary interface.
	NacSegment *string `pulumi:"nacSegment"`
	// VLAN for quarantined traffic.
	Quarantine *string `pulumi:"quarantine"`
	// VLAN for RSPAN/ERSPAN mirrored traffic.
	Rspan *string `pulumi:"rspan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN dedicated for video devices.
	Video *string `pulumi:"video"`
	// VLAN dedicated for voice devices.
	Voice *string `pulumi:"voice"`
}

// The set of arguments for constructing a SwitchcontrollerinitialconfigVlans resource.
type SwitchcontrollerinitialconfigVlansArgs struct {
	// Default VLAN (native) assigned to all switch ports upon discovery.
	DefaultVlan pulumi.StringPtrInput
	// VLAN for NAC onboarding devices.
	Nac pulumi.StringPtrInput
	// VLAN for NAC segemnt primary interface.
	NacSegment pulumi.StringPtrInput
	// VLAN for quarantined traffic.
	Quarantine pulumi.StringPtrInput
	// VLAN for RSPAN/ERSPAN mirrored traffic.
	Rspan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN dedicated for video devices.
	Video pulumi.StringPtrInput
	// VLAN dedicated for voice devices.
	Voice pulumi.StringPtrInput
}

func (SwitchcontrollerinitialconfigVlansArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerinitialconfigVlansArgs)(nil)).Elem()
}

type SwitchcontrollerinitialconfigVlansInput interface {
	pulumi.Input

	ToSwitchcontrollerinitialconfigVlansOutput() SwitchcontrollerinitialconfigVlansOutput
	ToSwitchcontrollerinitialconfigVlansOutputWithContext(ctx context.Context) SwitchcontrollerinitialconfigVlansOutput
}

func (*SwitchcontrollerinitialconfigVlans) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerinitialconfigVlans)(nil)).Elem()
}

func (i *SwitchcontrollerinitialconfigVlans) ToSwitchcontrollerinitialconfigVlansOutput() SwitchcontrollerinitialconfigVlansOutput {
	return i.ToSwitchcontrollerinitialconfigVlansOutputWithContext(context.Background())
}

func (i *SwitchcontrollerinitialconfigVlans) ToSwitchcontrollerinitialconfigVlansOutputWithContext(ctx context.Context) SwitchcontrollerinitialconfigVlansOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerinitialconfigVlansOutput)
}

// SwitchcontrollerinitialconfigVlansArrayInput is an input type that accepts SwitchcontrollerinitialconfigVlansArray and SwitchcontrollerinitialconfigVlansArrayOutput values.
// You can construct a concrete instance of `SwitchcontrollerinitialconfigVlansArrayInput` via:
//
//	SwitchcontrollerinitialconfigVlansArray{ SwitchcontrollerinitialconfigVlansArgs{...} }
type SwitchcontrollerinitialconfigVlansArrayInput interface {
	pulumi.Input

	ToSwitchcontrollerinitialconfigVlansArrayOutput() SwitchcontrollerinitialconfigVlansArrayOutput
	ToSwitchcontrollerinitialconfigVlansArrayOutputWithContext(context.Context) SwitchcontrollerinitialconfigVlansArrayOutput
}

type SwitchcontrollerinitialconfigVlansArray []SwitchcontrollerinitialconfigVlansInput

func (SwitchcontrollerinitialconfigVlansArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerinitialconfigVlans)(nil)).Elem()
}

func (i SwitchcontrollerinitialconfigVlansArray) ToSwitchcontrollerinitialconfigVlansArrayOutput() SwitchcontrollerinitialconfigVlansArrayOutput {
	return i.ToSwitchcontrollerinitialconfigVlansArrayOutputWithContext(context.Background())
}

func (i SwitchcontrollerinitialconfigVlansArray) ToSwitchcontrollerinitialconfigVlansArrayOutputWithContext(ctx context.Context) SwitchcontrollerinitialconfigVlansArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerinitialconfigVlansArrayOutput)
}

// SwitchcontrollerinitialconfigVlansMapInput is an input type that accepts SwitchcontrollerinitialconfigVlansMap and SwitchcontrollerinitialconfigVlansMapOutput values.
// You can construct a concrete instance of `SwitchcontrollerinitialconfigVlansMapInput` via:
//
//	SwitchcontrollerinitialconfigVlansMap{ "key": SwitchcontrollerinitialconfigVlansArgs{...} }
type SwitchcontrollerinitialconfigVlansMapInput interface {
	pulumi.Input

	ToSwitchcontrollerinitialconfigVlansMapOutput() SwitchcontrollerinitialconfigVlansMapOutput
	ToSwitchcontrollerinitialconfigVlansMapOutputWithContext(context.Context) SwitchcontrollerinitialconfigVlansMapOutput
}

type SwitchcontrollerinitialconfigVlansMap map[string]SwitchcontrollerinitialconfigVlansInput

func (SwitchcontrollerinitialconfigVlansMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerinitialconfigVlans)(nil)).Elem()
}

func (i SwitchcontrollerinitialconfigVlansMap) ToSwitchcontrollerinitialconfigVlansMapOutput() SwitchcontrollerinitialconfigVlansMapOutput {
	return i.ToSwitchcontrollerinitialconfigVlansMapOutputWithContext(context.Background())
}

func (i SwitchcontrollerinitialconfigVlansMap) ToSwitchcontrollerinitialconfigVlansMapOutputWithContext(ctx context.Context) SwitchcontrollerinitialconfigVlansMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerinitialconfigVlansMapOutput)
}

type SwitchcontrollerinitialconfigVlansOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerinitialconfigVlansOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerinitialconfigVlans)(nil)).Elem()
}

func (o SwitchcontrollerinitialconfigVlansOutput) ToSwitchcontrollerinitialconfigVlansOutput() SwitchcontrollerinitialconfigVlansOutput {
	return o
}

func (o SwitchcontrollerinitialconfigVlansOutput) ToSwitchcontrollerinitialconfigVlansOutputWithContext(ctx context.Context) SwitchcontrollerinitialconfigVlansOutput {
	return o
}

// Default VLAN (native) assigned to all switch ports upon discovery.
func (o SwitchcontrollerinitialconfigVlansOutput) DefaultVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerinitialconfigVlans) pulumi.StringOutput { return v.DefaultVlan }).(pulumi.StringOutput)
}

// VLAN for NAC onboarding devices.
func (o SwitchcontrollerinitialconfigVlansOutput) Nac() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerinitialconfigVlans) pulumi.StringOutput { return v.Nac }).(pulumi.StringOutput)
}

// VLAN for NAC segemnt primary interface.
func (o SwitchcontrollerinitialconfigVlansOutput) NacSegment() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerinitialconfigVlans) pulumi.StringOutput { return v.NacSegment }).(pulumi.StringOutput)
}

// VLAN for quarantined traffic.
func (o SwitchcontrollerinitialconfigVlansOutput) Quarantine() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerinitialconfigVlans) pulumi.StringOutput { return v.Quarantine }).(pulumi.StringOutput)
}

// VLAN for RSPAN/ERSPAN mirrored traffic.
func (o SwitchcontrollerinitialconfigVlansOutput) Rspan() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerinitialconfigVlans) pulumi.StringOutput { return v.Rspan }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchcontrollerinitialconfigVlansOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerinitialconfigVlans) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// VLAN dedicated for video devices.
func (o SwitchcontrollerinitialconfigVlansOutput) Video() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerinitialconfigVlans) pulumi.StringOutput { return v.Video }).(pulumi.StringOutput)
}

// VLAN dedicated for voice devices.
func (o SwitchcontrollerinitialconfigVlansOutput) Voice() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerinitialconfigVlans) pulumi.StringOutput { return v.Voice }).(pulumi.StringOutput)
}

type SwitchcontrollerinitialconfigVlansArrayOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerinitialconfigVlansArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerinitialconfigVlans)(nil)).Elem()
}

func (o SwitchcontrollerinitialconfigVlansArrayOutput) ToSwitchcontrollerinitialconfigVlansArrayOutput() SwitchcontrollerinitialconfigVlansArrayOutput {
	return o
}

func (o SwitchcontrollerinitialconfigVlansArrayOutput) ToSwitchcontrollerinitialconfigVlansArrayOutputWithContext(ctx context.Context) SwitchcontrollerinitialconfigVlansArrayOutput {
	return o
}

func (o SwitchcontrollerinitialconfigVlansArrayOutput) Index(i pulumi.IntInput) SwitchcontrollerinitialconfigVlansOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchcontrollerinitialconfigVlans {
		return vs[0].([]*SwitchcontrollerinitialconfigVlans)[vs[1].(int)]
	}).(SwitchcontrollerinitialconfigVlansOutput)
}

type SwitchcontrollerinitialconfigVlansMapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerinitialconfigVlansMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerinitialconfigVlans)(nil)).Elem()
}

func (o SwitchcontrollerinitialconfigVlansMapOutput) ToSwitchcontrollerinitialconfigVlansMapOutput() SwitchcontrollerinitialconfigVlansMapOutput {
	return o
}

func (o SwitchcontrollerinitialconfigVlansMapOutput) ToSwitchcontrollerinitialconfigVlansMapOutputWithContext(ctx context.Context) SwitchcontrollerinitialconfigVlansMapOutput {
	return o
}

func (o SwitchcontrollerinitialconfigVlansMapOutput) MapIndex(k pulumi.StringInput) SwitchcontrollerinitialconfigVlansOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchcontrollerinitialconfigVlans {
		return vs[0].(map[string]*SwitchcontrollerinitialconfigVlans)[vs[1].(string)]
	}).(SwitchcontrollerinitialconfigVlansOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerinitialconfigVlansInput)(nil)).Elem(), &SwitchcontrollerinitialconfigVlans{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerinitialconfigVlansArrayInput)(nil)).Elem(), SwitchcontrollerinitialconfigVlansArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerinitialconfigVlansMapInput)(nil)).Elem(), SwitchcontrollerinitialconfigVlansMap{})
	pulumi.RegisterOutputType(SwitchcontrollerinitialconfigVlansOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerinitialconfigVlansArrayOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerinitialconfigVlansMapOutput{})
}
