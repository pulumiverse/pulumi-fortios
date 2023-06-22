// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure WAN metrics.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewWirelesscontrollerhotspot20H2qpwanmetric(ctx, "trname", &fortios.Wirelesscontrollerhotspot20H2qpwanmetricArgs{
//				DownlinkLoad:            pulumi.Int(0),
//				DownlinkSpeed:           pulumi.Int(2400),
//				LinkAtCapacity:          pulumi.String("disable"),
//				LinkStatus:              pulumi.String("up"),
//				LoadMeasurementDuration: pulumi.Int(0),
//				SymmetricWanLink:        pulumi.String("symmetric"),
//				UplinkLoad:              pulumi.Int(0),
//				UplinkSpeed:             pulumi.Int(2400),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # WirelessControllerHotspot20 H2QpWanMetric can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerhotspot20H2qpwanmetric:Wirelesscontrollerhotspot20H2qpwanmetric labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerhotspot20H2qpwanmetric:Wirelesscontrollerhotspot20H2qpwanmetric labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Wirelesscontrollerhotspot20H2qpwanmetric struct {
	pulumi.CustomResourceState

	// Downlink load.
	DownlinkLoad pulumi.IntOutput `pulumi:"downlinkLoad"`
	// Downlink speed (in kilobits/s).
	DownlinkSpeed pulumi.IntOutput `pulumi:"downlinkSpeed"`
	// Link at capacity. Valid values: `enable`, `disable`.
	LinkAtCapacity pulumi.StringOutput `pulumi:"linkAtCapacity"`
	// Link status. Valid values: `up`, `down`, `in-test`.
	LinkStatus pulumi.StringOutput `pulumi:"linkStatus"`
	// Load measurement duration (in tenths of a second).
	LoadMeasurementDuration pulumi.IntOutput `pulumi:"loadMeasurementDuration"`
	// WAN metric name.
	Name pulumi.StringOutput `pulumi:"name"`
	// WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
	SymmetricWanLink pulumi.StringOutput `pulumi:"symmetricWanLink"`
	// Uplink load.
	UplinkLoad pulumi.IntOutput `pulumi:"uplinkLoad"`
	// Uplink speed (in kilobits/s).
	UplinkSpeed pulumi.IntOutput `pulumi:"uplinkSpeed"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelesscontrollerhotspot20H2qpwanmetric registers a new resource with the given unique name, arguments, and options.
func NewWirelesscontrollerhotspot20H2qpwanmetric(ctx *pulumi.Context,
	name string, args *Wirelesscontrollerhotspot20H2qpwanmetricArgs, opts ...pulumi.ResourceOption) (*Wirelesscontrollerhotspot20H2qpwanmetric, error) {
	if args == nil {
		args = &Wirelesscontrollerhotspot20H2qpwanmetricArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Wirelesscontrollerhotspot20H2qpwanmetric
	err := ctx.RegisterResource("fortios:index/wirelesscontrollerhotspot20H2qpwanmetric:Wirelesscontrollerhotspot20H2qpwanmetric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelesscontrollerhotspot20H2qpwanmetric gets an existing Wirelesscontrollerhotspot20H2qpwanmetric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelesscontrollerhotspot20H2qpwanmetric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Wirelesscontrollerhotspot20H2qpwanmetricState, opts ...pulumi.ResourceOption) (*Wirelesscontrollerhotspot20H2qpwanmetric, error) {
	var resource Wirelesscontrollerhotspot20H2qpwanmetric
	err := ctx.ReadResource("fortios:index/wirelesscontrollerhotspot20H2qpwanmetric:Wirelesscontrollerhotspot20H2qpwanmetric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Wirelesscontrollerhotspot20H2qpwanmetric resources.
type wirelesscontrollerhotspot20H2qpwanmetricState struct {
	// Downlink load.
	DownlinkLoad *int `pulumi:"downlinkLoad"`
	// Downlink speed (in kilobits/s).
	DownlinkSpeed *int `pulumi:"downlinkSpeed"`
	// Link at capacity. Valid values: `enable`, `disable`.
	LinkAtCapacity *string `pulumi:"linkAtCapacity"`
	// Link status. Valid values: `up`, `down`, `in-test`.
	LinkStatus *string `pulumi:"linkStatus"`
	// Load measurement duration (in tenths of a second).
	LoadMeasurementDuration *int `pulumi:"loadMeasurementDuration"`
	// WAN metric name.
	Name *string `pulumi:"name"`
	// WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
	SymmetricWanLink *string `pulumi:"symmetricWanLink"`
	// Uplink load.
	UplinkLoad *int `pulumi:"uplinkLoad"`
	// Uplink speed (in kilobits/s).
	UplinkSpeed *int `pulumi:"uplinkSpeed"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Wirelesscontrollerhotspot20H2qpwanmetricState struct {
	// Downlink load.
	DownlinkLoad pulumi.IntPtrInput
	// Downlink speed (in kilobits/s).
	DownlinkSpeed pulumi.IntPtrInput
	// Link at capacity. Valid values: `enable`, `disable`.
	LinkAtCapacity pulumi.StringPtrInput
	// Link status. Valid values: `up`, `down`, `in-test`.
	LinkStatus pulumi.StringPtrInput
	// Load measurement duration (in tenths of a second).
	LoadMeasurementDuration pulumi.IntPtrInput
	// WAN metric name.
	Name pulumi.StringPtrInput
	// WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
	SymmetricWanLink pulumi.StringPtrInput
	// Uplink load.
	UplinkLoad pulumi.IntPtrInput
	// Uplink speed (in kilobits/s).
	UplinkSpeed pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Wirelesscontrollerhotspot20H2qpwanmetricState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerhotspot20H2qpwanmetricState)(nil)).Elem()
}

type wirelesscontrollerhotspot20H2qpwanmetricArgs struct {
	// Downlink load.
	DownlinkLoad *int `pulumi:"downlinkLoad"`
	// Downlink speed (in kilobits/s).
	DownlinkSpeed *int `pulumi:"downlinkSpeed"`
	// Link at capacity. Valid values: `enable`, `disable`.
	LinkAtCapacity *string `pulumi:"linkAtCapacity"`
	// Link status. Valid values: `up`, `down`, `in-test`.
	LinkStatus *string `pulumi:"linkStatus"`
	// Load measurement duration (in tenths of a second).
	LoadMeasurementDuration *int `pulumi:"loadMeasurementDuration"`
	// WAN metric name.
	Name *string `pulumi:"name"`
	// WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
	SymmetricWanLink *string `pulumi:"symmetricWanLink"`
	// Uplink load.
	UplinkLoad *int `pulumi:"uplinkLoad"`
	// Uplink speed (in kilobits/s).
	UplinkSpeed *int `pulumi:"uplinkSpeed"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Wirelesscontrollerhotspot20H2qpwanmetric resource.
type Wirelesscontrollerhotspot20H2qpwanmetricArgs struct {
	// Downlink load.
	DownlinkLoad pulumi.IntPtrInput
	// Downlink speed (in kilobits/s).
	DownlinkSpeed pulumi.IntPtrInput
	// Link at capacity. Valid values: `enable`, `disable`.
	LinkAtCapacity pulumi.StringPtrInput
	// Link status. Valid values: `up`, `down`, `in-test`.
	LinkStatus pulumi.StringPtrInput
	// Load measurement duration (in tenths of a second).
	LoadMeasurementDuration pulumi.IntPtrInput
	// WAN metric name.
	Name pulumi.StringPtrInput
	// WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
	SymmetricWanLink pulumi.StringPtrInput
	// Uplink load.
	UplinkLoad pulumi.IntPtrInput
	// Uplink speed (in kilobits/s).
	UplinkSpeed pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Wirelesscontrollerhotspot20H2qpwanmetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerhotspot20H2qpwanmetricArgs)(nil)).Elem()
}

type Wirelesscontrollerhotspot20H2qpwanmetricInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20H2qpwanmetricOutput() Wirelesscontrollerhotspot20H2qpwanmetricOutput
	ToWirelesscontrollerhotspot20H2qpwanmetricOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qpwanmetricOutput
}

func (*Wirelesscontrollerhotspot20H2qpwanmetric) ElementType() reflect.Type {
	return reflect.TypeOf((**Wirelesscontrollerhotspot20H2qpwanmetric)(nil)).Elem()
}

func (i *Wirelesscontrollerhotspot20H2qpwanmetric) ToWirelesscontrollerhotspot20H2qpwanmetricOutput() Wirelesscontrollerhotspot20H2qpwanmetricOutput {
	return i.ToWirelesscontrollerhotspot20H2qpwanmetricOutputWithContext(context.Background())
}

func (i *Wirelesscontrollerhotspot20H2qpwanmetric) ToWirelesscontrollerhotspot20H2qpwanmetricOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qpwanmetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20H2qpwanmetricOutput)
}

// Wirelesscontrollerhotspot20H2qpwanmetricArrayInput is an input type that accepts Wirelesscontrollerhotspot20H2qpwanmetricArray and Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput values.
// You can construct a concrete instance of `Wirelesscontrollerhotspot20H2qpwanmetricArrayInput` via:
//
//	Wirelesscontrollerhotspot20H2qpwanmetricArray{ Wirelesscontrollerhotspot20H2qpwanmetricArgs{...} }
type Wirelesscontrollerhotspot20H2qpwanmetricArrayInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20H2qpwanmetricArrayOutput() Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput
	ToWirelesscontrollerhotspot20H2qpwanmetricArrayOutputWithContext(context.Context) Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput
}

type Wirelesscontrollerhotspot20H2qpwanmetricArray []Wirelesscontrollerhotspot20H2qpwanmetricInput

func (Wirelesscontrollerhotspot20H2qpwanmetricArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wirelesscontrollerhotspot20H2qpwanmetric)(nil)).Elem()
}

func (i Wirelesscontrollerhotspot20H2qpwanmetricArray) ToWirelesscontrollerhotspot20H2qpwanmetricArrayOutput() Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput {
	return i.ToWirelesscontrollerhotspot20H2qpwanmetricArrayOutputWithContext(context.Background())
}

func (i Wirelesscontrollerhotspot20H2qpwanmetricArray) ToWirelesscontrollerhotspot20H2qpwanmetricArrayOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput)
}

// Wirelesscontrollerhotspot20H2qpwanmetricMapInput is an input type that accepts Wirelesscontrollerhotspot20H2qpwanmetricMap and Wirelesscontrollerhotspot20H2qpwanmetricMapOutput values.
// You can construct a concrete instance of `Wirelesscontrollerhotspot20H2qpwanmetricMapInput` via:
//
//	Wirelesscontrollerhotspot20H2qpwanmetricMap{ "key": Wirelesscontrollerhotspot20H2qpwanmetricArgs{...} }
type Wirelesscontrollerhotspot20H2qpwanmetricMapInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20H2qpwanmetricMapOutput() Wirelesscontrollerhotspot20H2qpwanmetricMapOutput
	ToWirelesscontrollerhotspot20H2qpwanmetricMapOutputWithContext(context.Context) Wirelesscontrollerhotspot20H2qpwanmetricMapOutput
}

type Wirelesscontrollerhotspot20H2qpwanmetricMap map[string]Wirelesscontrollerhotspot20H2qpwanmetricInput

func (Wirelesscontrollerhotspot20H2qpwanmetricMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wirelesscontrollerhotspot20H2qpwanmetric)(nil)).Elem()
}

func (i Wirelesscontrollerhotspot20H2qpwanmetricMap) ToWirelesscontrollerhotspot20H2qpwanmetricMapOutput() Wirelesscontrollerhotspot20H2qpwanmetricMapOutput {
	return i.ToWirelesscontrollerhotspot20H2qpwanmetricMapOutputWithContext(context.Background())
}

func (i Wirelesscontrollerhotspot20H2qpwanmetricMap) ToWirelesscontrollerhotspot20H2qpwanmetricMapOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qpwanmetricMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20H2qpwanmetricMapOutput)
}

type Wirelesscontrollerhotspot20H2qpwanmetricOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20H2qpwanmetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Wirelesscontrollerhotspot20H2qpwanmetric)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20H2qpwanmetricOutput) ToWirelesscontrollerhotspot20H2qpwanmetricOutput() Wirelesscontrollerhotspot20H2qpwanmetricOutput {
	return o
}

func (o Wirelesscontrollerhotspot20H2qpwanmetricOutput) ToWirelesscontrollerhotspot20H2qpwanmetricOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qpwanmetricOutput {
	return o
}

// Downlink load.
func (o Wirelesscontrollerhotspot20H2qpwanmetricOutput) DownlinkLoad() pulumi.IntOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qpwanmetric) pulumi.IntOutput { return v.DownlinkLoad }).(pulumi.IntOutput)
}

// Downlink speed (in kilobits/s).
func (o Wirelesscontrollerhotspot20H2qpwanmetricOutput) DownlinkSpeed() pulumi.IntOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qpwanmetric) pulumi.IntOutput { return v.DownlinkSpeed }).(pulumi.IntOutput)
}

// Link at capacity. Valid values: `enable`, `disable`.
func (o Wirelesscontrollerhotspot20H2qpwanmetricOutput) LinkAtCapacity() pulumi.StringOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qpwanmetric) pulumi.StringOutput { return v.LinkAtCapacity }).(pulumi.StringOutput)
}

// Link status. Valid values: `up`, `down`, `in-test`.
func (o Wirelesscontrollerhotspot20H2qpwanmetricOutput) LinkStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qpwanmetric) pulumi.StringOutput { return v.LinkStatus }).(pulumi.StringOutput)
}

// Load measurement duration (in tenths of a second).
func (o Wirelesscontrollerhotspot20H2qpwanmetricOutput) LoadMeasurementDuration() pulumi.IntOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qpwanmetric) pulumi.IntOutput { return v.LoadMeasurementDuration }).(pulumi.IntOutput)
}

// WAN metric name.
func (o Wirelesscontrollerhotspot20H2qpwanmetricOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qpwanmetric) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
func (o Wirelesscontrollerhotspot20H2qpwanmetricOutput) SymmetricWanLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qpwanmetric) pulumi.StringOutput { return v.SymmetricWanLink }).(pulumi.StringOutput)
}

// Uplink load.
func (o Wirelesscontrollerhotspot20H2qpwanmetricOutput) UplinkLoad() pulumi.IntOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qpwanmetric) pulumi.IntOutput { return v.UplinkLoad }).(pulumi.IntOutput)
}

// Uplink speed (in kilobits/s).
func (o Wirelesscontrollerhotspot20H2qpwanmetricOutput) UplinkSpeed() pulumi.IntOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qpwanmetric) pulumi.IntOutput { return v.UplinkSpeed }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Wirelesscontrollerhotspot20H2qpwanmetricOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20H2qpwanmetric) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wirelesscontrollerhotspot20H2qpwanmetric)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput) ToWirelesscontrollerhotspot20H2qpwanmetricArrayOutput() Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput {
	return o
}

func (o Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput) ToWirelesscontrollerhotspot20H2qpwanmetricArrayOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput {
	return o
}

func (o Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput) Index(i pulumi.IntInput) Wirelesscontrollerhotspot20H2qpwanmetricOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Wirelesscontrollerhotspot20H2qpwanmetric {
		return vs[0].([]*Wirelesscontrollerhotspot20H2qpwanmetric)[vs[1].(int)]
	}).(Wirelesscontrollerhotspot20H2qpwanmetricOutput)
}

type Wirelesscontrollerhotspot20H2qpwanmetricMapOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20H2qpwanmetricMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wirelesscontrollerhotspot20H2qpwanmetric)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20H2qpwanmetricMapOutput) ToWirelesscontrollerhotspot20H2qpwanmetricMapOutput() Wirelesscontrollerhotspot20H2qpwanmetricMapOutput {
	return o
}

func (o Wirelesscontrollerhotspot20H2qpwanmetricMapOutput) ToWirelesscontrollerhotspot20H2qpwanmetricMapOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20H2qpwanmetricMapOutput {
	return o
}

func (o Wirelesscontrollerhotspot20H2qpwanmetricMapOutput) MapIndex(k pulumi.StringInput) Wirelesscontrollerhotspot20H2qpwanmetricOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Wirelesscontrollerhotspot20H2qpwanmetric {
		return vs[0].(map[string]*Wirelesscontrollerhotspot20H2qpwanmetric)[vs[1].(string)]
	}).(Wirelesscontrollerhotspot20H2qpwanmetricOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20H2qpwanmetricInput)(nil)).Elem(), &Wirelesscontrollerhotspot20H2qpwanmetric{})
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20H2qpwanmetricArrayInput)(nil)).Elem(), Wirelesscontrollerhotspot20H2qpwanmetricArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20H2qpwanmetricMapInput)(nil)).Elem(), Wirelesscontrollerhotspot20H2qpwanmetricMap{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20H2qpwanmetricOutput{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20H2qpwanmetricArrayOutput{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20H2qpwanmetricMapOutput{})
}
