// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure Bluetooth Low Energy profile.
//
// ## Import
//
// WirelessController BleProfile can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/bleprofile:Bleprofile labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/bleprofile:Bleprofile labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Bleprofile struct {
	pulumi.CustomResourceState

	// Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
	Advertising pulumi.StringOutput `pulumi:"advertising"`
	// Beacon interval (default = 100 msec).
	BeaconInterval pulumi.IntOutput `pulumi:"beaconInterval"`
	// Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
	BleScanning pulumi.StringOutput `pulumi:"bleScanning"`
	// Comment.
	Comment pulumi.StringOutput `pulumi:"comment"`
	// Eddystone instance ID.
	EddystoneInstance pulumi.StringOutput `pulumi:"eddystoneInstance"`
	// Eddystone namespace ID.
	EddystoneNamespace pulumi.StringOutput `pulumi:"eddystoneNamespace"`
	// Eddystone URL.
	EddystoneUrl pulumi.StringOutput `pulumi:"eddystoneUrl"`
	// Eddystone encoded URL hexadecimal string
	EddystoneUrlEncodeHex pulumi.StringOutput `pulumi:"eddystoneUrlEncodeHex"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	IbeaconUuid pulumi.StringOutput `pulumi:"ibeaconUuid"`
	// Major ID.
	MajorId pulumi.IntOutput `pulumi:"majorId"`
	// Minor ID.
	MinorId pulumi.IntOutput `pulumi:"minorId"`
	// Bluetooth Low Energy profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Scan Interval (default = 50 msec).
	ScanInterval pulumi.IntOutput `pulumi:"scanInterval"`
	// Scan Period (default = 4000 msec).
	ScanPeriod pulumi.IntOutput `pulumi:"scanPeriod"`
	// Minimum signal level/threshold in dBm required for the AP to report detected BLE device (-95 to -20, default = -90).
	ScanThreshold pulumi.StringOutput `pulumi:"scanThreshold"`
	// Scan Time (default = 1000 msec).
	ScanTime pulumi.IntOutput `pulumi:"scanTime"`
	// Scan Type (default = active). Valid values: `active`, `passive`.
	ScanType pulumi.StringOutput `pulumi:"scanType"`
	// Scan Windows (default = 50 msec).
	ScanWindow pulumi.IntOutput `pulumi:"scanWindow"`
	// Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
	Txpower pulumi.StringOutput `pulumi:"txpower"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewBleprofile registers a new resource with the given unique name, arguments, and options.
func NewBleprofile(ctx *pulumi.Context,
	name string, args *BleprofileArgs, opts ...pulumi.ResourceOption) (*Bleprofile, error) {
	if args == nil {
		args = &BleprofileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bleprofile
	err := ctx.RegisterResource("fortios:wirelesscontroller/bleprofile:Bleprofile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBleprofile gets an existing Bleprofile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBleprofile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BleprofileState, opts ...pulumi.ResourceOption) (*Bleprofile, error) {
	var resource Bleprofile
	err := ctx.ReadResource("fortios:wirelesscontroller/bleprofile:Bleprofile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bleprofile resources.
type bleprofileState struct {
	// Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
	Advertising *string `pulumi:"advertising"`
	// Beacon interval (default = 100 msec).
	BeaconInterval *int `pulumi:"beaconInterval"`
	// Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
	BleScanning *string `pulumi:"bleScanning"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Eddystone instance ID.
	EddystoneInstance *string `pulumi:"eddystoneInstance"`
	// Eddystone namespace ID.
	EddystoneNamespace *string `pulumi:"eddystoneNamespace"`
	// Eddystone URL.
	EddystoneUrl *string `pulumi:"eddystoneUrl"`
	// Eddystone encoded URL hexadecimal string
	EddystoneUrlEncodeHex *string `pulumi:"eddystoneUrlEncodeHex"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	IbeaconUuid *string `pulumi:"ibeaconUuid"`
	// Major ID.
	MajorId *int `pulumi:"majorId"`
	// Minor ID.
	MinorId *int `pulumi:"minorId"`
	// Bluetooth Low Energy profile name.
	Name *string `pulumi:"name"`
	// Scan Interval (default = 50 msec).
	ScanInterval *int `pulumi:"scanInterval"`
	// Scan Period (default = 4000 msec).
	ScanPeriod *int `pulumi:"scanPeriod"`
	// Minimum signal level/threshold in dBm required for the AP to report detected BLE device (-95 to -20, default = -90).
	ScanThreshold *string `pulumi:"scanThreshold"`
	// Scan Time (default = 1000 msec).
	ScanTime *int `pulumi:"scanTime"`
	// Scan Type (default = active). Valid values: `active`, `passive`.
	ScanType *string `pulumi:"scanType"`
	// Scan Windows (default = 50 msec).
	ScanWindow *int `pulumi:"scanWindow"`
	// Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
	Txpower *string `pulumi:"txpower"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type BleprofileState struct {
	// Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
	Advertising pulumi.StringPtrInput
	// Beacon interval (default = 100 msec).
	BeaconInterval pulumi.IntPtrInput
	// Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
	BleScanning pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Eddystone instance ID.
	EddystoneInstance pulumi.StringPtrInput
	// Eddystone namespace ID.
	EddystoneNamespace pulumi.StringPtrInput
	// Eddystone URL.
	EddystoneUrl pulumi.StringPtrInput
	// Eddystone encoded URL hexadecimal string
	EddystoneUrlEncodeHex pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	IbeaconUuid pulumi.StringPtrInput
	// Major ID.
	MajorId pulumi.IntPtrInput
	// Minor ID.
	MinorId pulumi.IntPtrInput
	// Bluetooth Low Energy profile name.
	Name pulumi.StringPtrInput
	// Scan Interval (default = 50 msec).
	ScanInterval pulumi.IntPtrInput
	// Scan Period (default = 4000 msec).
	ScanPeriod pulumi.IntPtrInput
	// Minimum signal level/threshold in dBm required for the AP to report detected BLE device (-95 to -20, default = -90).
	ScanThreshold pulumi.StringPtrInput
	// Scan Time (default = 1000 msec).
	ScanTime pulumi.IntPtrInput
	// Scan Type (default = active). Valid values: `active`, `passive`.
	ScanType pulumi.StringPtrInput
	// Scan Windows (default = 50 msec).
	ScanWindow pulumi.IntPtrInput
	// Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
	Txpower pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (BleprofileState) ElementType() reflect.Type {
	return reflect.TypeOf((*bleprofileState)(nil)).Elem()
}

type bleprofileArgs struct {
	// Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
	Advertising *string `pulumi:"advertising"`
	// Beacon interval (default = 100 msec).
	BeaconInterval *int `pulumi:"beaconInterval"`
	// Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
	BleScanning *string `pulumi:"bleScanning"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Eddystone instance ID.
	EddystoneInstance *string `pulumi:"eddystoneInstance"`
	// Eddystone namespace ID.
	EddystoneNamespace *string `pulumi:"eddystoneNamespace"`
	// Eddystone URL.
	EddystoneUrl *string `pulumi:"eddystoneUrl"`
	// Eddystone encoded URL hexadecimal string
	EddystoneUrlEncodeHex *string `pulumi:"eddystoneUrlEncodeHex"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	IbeaconUuid *string `pulumi:"ibeaconUuid"`
	// Major ID.
	MajorId *int `pulumi:"majorId"`
	// Minor ID.
	MinorId *int `pulumi:"minorId"`
	// Bluetooth Low Energy profile name.
	Name *string `pulumi:"name"`
	// Scan Interval (default = 50 msec).
	ScanInterval *int `pulumi:"scanInterval"`
	// Scan Period (default = 4000 msec).
	ScanPeriod *int `pulumi:"scanPeriod"`
	// Minimum signal level/threshold in dBm required for the AP to report detected BLE device (-95 to -20, default = -90).
	ScanThreshold *string `pulumi:"scanThreshold"`
	// Scan Time (default = 1000 msec).
	ScanTime *int `pulumi:"scanTime"`
	// Scan Type (default = active). Valid values: `active`, `passive`.
	ScanType *string `pulumi:"scanType"`
	// Scan Windows (default = 50 msec).
	ScanWindow *int `pulumi:"scanWindow"`
	// Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
	Txpower *string `pulumi:"txpower"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Bleprofile resource.
type BleprofileArgs struct {
	// Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
	Advertising pulumi.StringPtrInput
	// Beacon interval (default = 100 msec).
	BeaconInterval pulumi.IntPtrInput
	// Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
	BleScanning pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Eddystone instance ID.
	EddystoneInstance pulumi.StringPtrInput
	// Eddystone namespace ID.
	EddystoneNamespace pulumi.StringPtrInput
	// Eddystone URL.
	EddystoneUrl pulumi.StringPtrInput
	// Eddystone encoded URL hexadecimal string
	EddystoneUrlEncodeHex pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	IbeaconUuid pulumi.StringPtrInput
	// Major ID.
	MajorId pulumi.IntPtrInput
	// Minor ID.
	MinorId pulumi.IntPtrInput
	// Bluetooth Low Energy profile name.
	Name pulumi.StringPtrInput
	// Scan Interval (default = 50 msec).
	ScanInterval pulumi.IntPtrInput
	// Scan Period (default = 4000 msec).
	ScanPeriod pulumi.IntPtrInput
	// Minimum signal level/threshold in dBm required for the AP to report detected BLE device (-95 to -20, default = -90).
	ScanThreshold pulumi.StringPtrInput
	// Scan Time (default = 1000 msec).
	ScanTime pulumi.IntPtrInput
	// Scan Type (default = active). Valid values: `active`, `passive`.
	ScanType pulumi.StringPtrInput
	// Scan Windows (default = 50 msec).
	ScanWindow pulumi.IntPtrInput
	// Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
	Txpower pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (BleprofileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bleprofileArgs)(nil)).Elem()
}

type BleprofileInput interface {
	pulumi.Input

	ToBleprofileOutput() BleprofileOutput
	ToBleprofileOutputWithContext(ctx context.Context) BleprofileOutput
}

func (*Bleprofile) ElementType() reflect.Type {
	return reflect.TypeOf((**Bleprofile)(nil)).Elem()
}

func (i *Bleprofile) ToBleprofileOutput() BleprofileOutput {
	return i.ToBleprofileOutputWithContext(context.Background())
}

func (i *Bleprofile) ToBleprofileOutputWithContext(ctx context.Context) BleprofileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BleprofileOutput)
}

// BleprofileArrayInput is an input type that accepts BleprofileArray and BleprofileArrayOutput values.
// You can construct a concrete instance of `BleprofileArrayInput` via:
//
//	BleprofileArray{ BleprofileArgs{...} }
type BleprofileArrayInput interface {
	pulumi.Input

	ToBleprofileArrayOutput() BleprofileArrayOutput
	ToBleprofileArrayOutputWithContext(context.Context) BleprofileArrayOutput
}

type BleprofileArray []BleprofileInput

func (BleprofileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bleprofile)(nil)).Elem()
}

func (i BleprofileArray) ToBleprofileArrayOutput() BleprofileArrayOutput {
	return i.ToBleprofileArrayOutputWithContext(context.Background())
}

func (i BleprofileArray) ToBleprofileArrayOutputWithContext(ctx context.Context) BleprofileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BleprofileArrayOutput)
}

// BleprofileMapInput is an input type that accepts BleprofileMap and BleprofileMapOutput values.
// You can construct a concrete instance of `BleprofileMapInput` via:
//
//	BleprofileMap{ "key": BleprofileArgs{...} }
type BleprofileMapInput interface {
	pulumi.Input

	ToBleprofileMapOutput() BleprofileMapOutput
	ToBleprofileMapOutputWithContext(context.Context) BleprofileMapOutput
}

type BleprofileMap map[string]BleprofileInput

func (BleprofileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bleprofile)(nil)).Elem()
}

func (i BleprofileMap) ToBleprofileMapOutput() BleprofileMapOutput {
	return i.ToBleprofileMapOutputWithContext(context.Background())
}

func (i BleprofileMap) ToBleprofileMapOutputWithContext(ctx context.Context) BleprofileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BleprofileMapOutput)
}

type BleprofileOutput struct{ *pulumi.OutputState }

func (BleprofileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bleprofile)(nil)).Elem()
}

func (o BleprofileOutput) ToBleprofileOutput() BleprofileOutput {
	return o
}

func (o BleprofileOutput) ToBleprofileOutputWithContext(ctx context.Context) BleprofileOutput {
	return o
}

// Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
func (o BleprofileOutput) Advertising() pulumi.StringOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.StringOutput { return v.Advertising }).(pulumi.StringOutput)
}

// Beacon interval (default = 100 msec).
func (o BleprofileOutput) BeaconInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.IntOutput { return v.BeaconInterval }).(pulumi.IntOutput)
}

// Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
func (o BleprofileOutput) BleScanning() pulumi.StringOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.StringOutput { return v.BleScanning }).(pulumi.StringOutput)
}

// Comment.
func (o BleprofileOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.StringOutput { return v.Comment }).(pulumi.StringOutput)
}

// Eddystone instance ID.
func (o BleprofileOutput) EddystoneInstance() pulumi.StringOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.StringOutput { return v.EddystoneInstance }).(pulumi.StringOutput)
}

// Eddystone namespace ID.
func (o BleprofileOutput) EddystoneNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.StringOutput { return v.EddystoneNamespace }).(pulumi.StringOutput)
}

// Eddystone URL.
func (o BleprofileOutput) EddystoneUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.StringOutput { return v.EddystoneUrl }).(pulumi.StringOutput)
}

// Eddystone encoded URL hexadecimal string
func (o BleprofileOutput) EddystoneUrlEncodeHex() pulumi.StringOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.StringOutput { return v.EddystoneUrlEncodeHex }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o BleprofileOutput) IbeaconUuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.StringOutput { return v.IbeaconUuid }).(pulumi.StringOutput)
}

// Major ID.
func (o BleprofileOutput) MajorId() pulumi.IntOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.IntOutput { return v.MajorId }).(pulumi.IntOutput)
}

// Minor ID.
func (o BleprofileOutput) MinorId() pulumi.IntOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.IntOutput { return v.MinorId }).(pulumi.IntOutput)
}

// Bluetooth Low Energy profile name.
func (o BleprofileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Scan Interval (default = 50 msec).
func (o BleprofileOutput) ScanInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.IntOutput { return v.ScanInterval }).(pulumi.IntOutput)
}

// Scan Period (default = 4000 msec).
func (o BleprofileOutput) ScanPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.IntOutput { return v.ScanPeriod }).(pulumi.IntOutput)
}

// Minimum signal level/threshold in dBm required for the AP to report detected BLE device (-95 to -20, default = -90).
func (o BleprofileOutput) ScanThreshold() pulumi.StringOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.StringOutput { return v.ScanThreshold }).(pulumi.StringOutput)
}

// Scan Time (default = 1000 msec).
func (o BleprofileOutput) ScanTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.IntOutput { return v.ScanTime }).(pulumi.IntOutput)
}

// Scan Type (default = active). Valid values: `active`, `passive`.
func (o BleprofileOutput) ScanType() pulumi.StringOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.StringOutput { return v.ScanType }).(pulumi.StringOutput)
}

// Scan Windows (default = 50 msec).
func (o BleprofileOutput) ScanWindow() pulumi.IntOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.IntOutput { return v.ScanWindow }).(pulumi.IntOutput)
}

// Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
func (o BleprofileOutput) Txpower() pulumi.StringOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.StringOutput { return v.Txpower }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o BleprofileOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Bleprofile) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type BleprofileArrayOutput struct{ *pulumi.OutputState }

func (BleprofileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bleprofile)(nil)).Elem()
}

func (o BleprofileArrayOutput) ToBleprofileArrayOutput() BleprofileArrayOutput {
	return o
}

func (o BleprofileArrayOutput) ToBleprofileArrayOutputWithContext(ctx context.Context) BleprofileArrayOutput {
	return o
}

func (o BleprofileArrayOutput) Index(i pulumi.IntInput) BleprofileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bleprofile {
		return vs[0].([]*Bleprofile)[vs[1].(int)]
	}).(BleprofileOutput)
}

type BleprofileMapOutput struct{ *pulumi.OutputState }

func (BleprofileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bleprofile)(nil)).Elem()
}

func (o BleprofileMapOutput) ToBleprofileMapOutput() BleprofileMapOutput {
	return o
}

func (o BleprofileMapOutput) ToBleprofileMapOutputWithContext(ctx context.Context) BleprofileMapOutput {
	return o
}

func (o BleprofileMapOutput) MapIndex(k pulumi.StringInput) BleprofileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bleprofile {
		return vs[0].(map[string]*Bleprofile)[vs[1].(string)]
	}).(BleprofileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BleprofileInput)(nil)).Elem(), &Bleprofile{})
	pulumi.RegisterInputType(reflect.TypeOf((*BleprofileArrayInput)(nil)).Elem(), BleprofileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BleprofileMapInput)(nil)).Elem(), BleprofileMap{})
	pulumi.RegisterOutputType(BleprofileOutput{})
	pulumi.RegisterOutputType(BleprofileArrayOutput{})
	pulumi.RegisterOutputType(BleprofileMapOutput{})
}
