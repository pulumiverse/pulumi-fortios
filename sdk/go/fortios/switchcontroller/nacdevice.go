// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,7.0.0`.
//
// ## Import
//
// SwitchController NacDevice can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/nacdevice:Nacdevice labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/nacdevice:Nacdevice labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Nacdevice struct {
	pulumi.CustomResourceState

	// Description for the learned NAC device.
	Description pulumi.StringOutput `pulumi:"description"`
	// Device ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Managed FortiSwitch port where NAC device is last learned.
	LastKnownPort pulumi.StringOutput `pulumi:"lastKnownPort"`
	// Managed FortiSwitch where NAC device is last learned.
	LastKnownSwitch pulumi.StringOutput `pulumi:"lastKnownSwitch"`
	// Device last seen.
	LastSeen pulumi.IntOutput `pulumi:"lastSeen"`
	// MAC address of the learned NAC device.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// MAC policy to be applied on this learned NAC device.
	MacPolicy pulumi.StringOutput `pulumi:"macPolicy"`
	// Matched NAC policy for the learned NAC device.
	MatchedNacPolicy pulumi.StringOutput `pulumi:"matchedNacPolicy"`
	// Port policy to be applied on this learned NAC device.
	PortPolicy pulumi.StringOutput `pulumi:"portPolicy"`
	// Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewNacdevice registers a new resource with the given unique name, arguments, and options.
func NewNacdevice(ctx *pulumi.Context,
	name string, args *NacdeviceArgs, opts ...pulumi.ResourceOption) (*Nacdevice, error) {
	if args == nil {
		args = &NacdeviceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Nacdevice
	err := ctx.RegisterResource("fortios:switchcontroller/nacdevice:Nacdevice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNacdevice gets an existing Nacdevice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNacdevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NacdeviceState, opts ...pulumi.ResourceOption) (*Nacdevice, error) {
	var resource Nacdevice
	err := ctx.ReadResource("fortios:switchcontroller/nacdevice:Nacdevice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Nacdevice resources.
type nacdeviceState struct {
	// Description for the learned NAC device.
	Description *string `pulumi:"description"`
	// Device ID.
	Fosid *int `pulumi:"fosid"`
	// Managed FortiSwitch port where NAC device is last learned.
	LastKnownPort *string `pulumi:"lastKnownPort"`
	// Managed FortiSwitch where NAC device is last learned.
	LastKnownSwitch *string `pulumi:"lastKnownSwitch"`
	// Device last seen.
	LastSeen *int `pulumi:"lastSeen"`
	// MAC address of the learned NAC device.
	Mac *string `pulumi:"mac"`
	// MAC policy to be applied on this learned NAC device.
	MacPolicy *string `pulumi:"macPolicy"`
	// Matched NAC policy for the learned NAC device.
	MatchedNacPolicy *string `pulumi:"matchedNacPolicy"`
	// Port policy to be applied on this learned NAC device.
	PortPolicy *string `pulumi:"portPolicy"`
	// Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type NacdeviceState struct {
	// Description for the learned NAC device.
	Description pulumi.StringPtrInput
	// Device ID.
	Fosid pulumi.IntPtrInput
	// Managed FortiSwitch port where NAC device is last learned.
	LastKnownPort pulumi.StringPtrInput
	// Managed FortiSwitch where NAC device is last learned.
	LastKnownSwitch pulumi.StringPtrInput
	// Device last seen.
	LastSeen pulumi.IntPtrInput
	// MAC address of the learned NAC device.
	Mac pulumi.StringPtrInput
	// MAC policy to be applied on this learned NAC device.
	MacPolicy pulumi.StringPtrInput
	// Matched NAC policy for the learned NAC device.
	MatchedNacPolicy pulumi.StringPtrInput
	// Port policy to be applied on this learned NAC device.
	PortPolicy pulumi.StringPtrInput
	// Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NacdeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*nacdeviceState)(nil)).Elem()
}

type nacdeviceArgs struct {
	// Description for the learned NAC device.
	Description *string `pulumi:"description"`
	// Device ID.
	Fosid *int `pulumi:"fosid"`
	// Managed FortiSwitch port where NAC device is last learned.
	LastKnownPort *string `pulumi:"lastKnownPort"`
	// Managed FortiSwitch where NAC device is last learned.
	LastKnownSwitch *string `pulumi:"lastKnownSwitch"`
	// Device last seen.
	LastSeen *int `pulumi:"lastSeen"`
	// MAC address of the learned NAC device.
	Mac *string `pulumi:"mac"`
	// MAC policy to be applied on this learned NAC device.
	MacPolicy *string `pulumi:"macPolicy"`
	// Matched NAC policy for the learned NAC device.
	MatchedNacPolicy *string `pulumi:"matchedNacPolicy"`
	// Port policy to be applied on this learned NAC device.
	PortPolicy *string `pulumi:"portPolicy"`
	// Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Nacdevice resource.
type NacdeviceArgs struct {
	// Description for the learned NAC device.
	Description pulumi.StringPtrInput
	// Device ID.
	Fosid pulumi.IntPtrInput
	// Managed FortiSwitch port where NAC device is last learned.
	LastKnownPort pulumi.StringPtrInput
	// Managed FortiSwitch where NAC device is last learned.
	LastKnownSwitch pulumi.StringPtrInput
	// Device last seen.
	LastSeen pulumi.IntPtrInput
	// MAC address of the learned NAC device.
	Mac pulumi.StringPtrInput
	// MAC policy to be applied on this learned NAC device.
	MacPolicy pulumi.StringPtrInput
	// Matched NAC policy for the learned NAC device.
	MatchedNacPolicy pulumi.StringPtrInput
	// Port policy to be applied on this learned NAC device.
	PortPolicy pulumi.StringPtrInput
	// Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NacdeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nacdeviceArgs)(nil)).Elem()
}

type NacdeviceInput interface {
	pulumi.Input

	ToNacdeviceOutput() NacdeviceOutput
	ToNacdeviceOutputWithContext(ctx context.Context) NacdeviceOutput
}

func (*Nacdevice) ElementType() reflect.Type {
	return reflect.TypeOf((**Nacdevice)(nil)).Elem()
}

func (i *Nacdevice) ToNacdeviceOutput() NacdeviceOutput {
	return i.ToNacdeviceOutputWithContext(context.Background())
}

func (i *Nacdevice) ToNacdeviceOutputWithContext(ctx context.Context) NacdeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NacdeviceOutput)
}

// NacdeviceArrayInput is an input type that accepts NacdeviceArray and NacdeviceArrayOutput values.
// You can construct a concrete instance of `NacdeviceArrayInput` via:
//
//	NacdeviceArray{ NacdeviceArgs{...} }
type NacdeviceArrayInput interface {
	pulumi.Input

	ToNacdeviceArrayOutput() NacdeviceArrayOutput
	ToNacdeviceArrayOutputWithContext(context.Context) NacdeviceArrayOutput
}

type NacdeviceArray []NacdeviceInput

func (NacdeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nacdevice)(nil)).Elem()
}

func (i NacdeviceArray) ToNacdeviceArrayOutput() NacdeviceArrayOutput {
	return i.ToNacdeviceArrayOutputWithContext(context.Background())
}

func (i NacdeviceArray) ToNacdeviceArrayOutputWithContext(ctx context.Context) NacdeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NacdeviceArrayOutput)
}

// NacdeviceMapInput is an input type that accepts NacdeviceMap and NacdeviceMapOutput values.
// You can construct a concrete instance of `NacdeviceMapInput` via:
//
//	NacdeviceMap{ "key": NacdeviceArgs{...} }
type NacdeviceMapInput interface {
	pulumi.Input

	ToNacdeviceMapOutput() NacdeviceMapOutput
	ToNacdeviceMapOutputWithContext(context.Context) NacdeviceMapOutput
}

type NacdeviceMap map[string]NacdeviceInput

func (NacdeviceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nacdevice)(nil)).Elem()
}

func (i NacdeviceMap) ToNacdeviceMapOutput() NacdeviceMapOutput {
	return i.ToNacdeviceMapOutputWithContext(context.Background())
}

func (i NacdeviceMap) ToNacdeviceMapOutputWithContext(ctx context.Context) NacdeviceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NacdeviceMapOutput)
}

type NacdeviceOutput struct{ *pulumi.OutputState }

func (NacdeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nacdevice)(nil)).Elem()
}

func (o NacdeviceOutput) ToNacdeviceOutput() NacdeviceOutput {
	return o
}

func (o NacdeviceOutput) ToNacdeviceOutputWithContext(ctx context.Context) NacdeviceOutput {
	return o
}

// Description for the learned NAC device.
func (o NacdeviceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacdevice) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Device ID.
func (o NacdeviceOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Nacdevice) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Managed FortiSwitch port where NAC device is last learned.
func (o NacdeviceOutput) LastKnownPort() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacdevice) pulumi.StringOutput { return v.LastKnownPort }).(pulumi.StringOutput)
}

// Managed FortiSwitch where NAC device is last learned.
func (o NacdeviceOutput) LastKnownSwitch() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacdevice) pulumi.StringOutput { return v.LastKnownSwitch }).(pulumi.StringOutput)
}

// Device last seen.
func (o NacdeviceOutput) LastSeen() pulumi.IntOutput {
	return o.ApplyT(func(v *Nacdevice) pulumi.IntOutput { return v.LastSeen }).(pulumi.IntOutput)
}

// MAC address of the learned NAC device.
func (o NacdeviceOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacdevice) pulumi.StringOutput { return v.Mac }).(pulumi.StringOutput)
}

// MAC policy to be applied on this learned NAC device.
func (o NacdeviceOutput) MacPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacdevice) pulumi.StringOutput { return v.MacPolicy }).(pulumi.StringOutput)
}

// Matched NAC policy for the learned NAC device.
func (o NacdeviceOutput) MatchedNacPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacdevice) pulumi.StringOutput { return v.MatchedNacPolicy }).(pulumi.StringOutput)
}

// Port policy to be applied on this learned NAC device.
func (o NacdeviceOutput) PortPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacdevice) pulumi.StringOutput { return v.PortPolicy }).(pulumi.StringOutput)
}

// Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
func (o NacdeviceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacdevice) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o NacdeviceOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nacdevice) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type NacdeviceArrayOutput struct{ *pulumi.OutputState }

func (NacdeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nacdevice)(nil)).Elem()
}

func (o NacdeviceArrayOutput) ToNacdeviceArrayOutput() NacdeviceArrayOutput {
	return o
}

func (o NacdeviceArrayOutput) ToNacdeviceArrayOutputWithContext(ctx context.Context) NacdeviceArrayOutput {
	return o
}

func (o NacdeviceArrayOutput) Index(i pulumi.IntInput) NacdeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Nacdevice {
		return vs[0].([]*Nacdevice)[vs[1].(int)]
	}).(NacdeviceOutput)
}

type NacdeviceMapOutput struct{ *pulumi.OutputState }

func (NacdeviceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nacdevice)(nil)).Elem()
}

func (o NacdeviceMapOutput) ToNacdeviceMapOutput() NacdeviceMapOutput {
	return o
}

func (o NacdeviceMapOutput) ToNacdeviceMapOutputWithContext(ctx context.Context) NacdeviceMapOutput {
	return o
}

func (o NacdeviceMapOutput) MapIndex(k pulumi.StringInput) NacdeviceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Nacdevice {
		return vs[0].(map[string]*Nacdevice)[vs[1].(string)]
	}).(NacdeviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NacdeviceInput)(nil)).Elem(), &Nacdevice{})
	pulumi.RegisterInputType(reflect.TypeOf((*NacdeviceArrayInput)(nil)).Elem(), NacdeviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NacdeviceMapInput)(nil)).Elem(), NacdeviceMap{})
	pulumi.RegisterOutputType(NacdeviceOutput{})
	pulumi.RegisterOutputType(NacdeviceArrayOutput{})
	pulumi.RegisterOutputType(NacdeviceMapOutput{})
}
