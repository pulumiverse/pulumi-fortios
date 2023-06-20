// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VDOM wireless controller configuration.
//
// ## Import
//
// # WirelessController Setting can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerSetting:WirelesscontrollerSetting labelname WirelessControllerSetting
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerSetting:WirelesscontrollerSetting labelname WirelessControllerSetting
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WirelesscontrollerSetting struct {
	pulumi.CustomResourceState

	// FortiCloud customer account ID.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// Country or region in which the FortiGate is located. The country determines the 802.11 bands and channels that are available.
	Country pulumi.StringOutput `pulumi:"country"`
	// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
	DarrpOptimize pulumi.IntOutput `pulumi:"darrpOptimize"`
	// Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
	DarrpOptimizeSchedules WirelesscontrollerSettingDarrpOptimizeScheduleArrayOutput `pulumi:"darrpOptimizeSchedules"`
	// Lower limit of creation time of device for identification in minutes (0 - 60, default = 5).
	DeviceHoldoff pulumi.IntOutput `pulumi:"deviceHoldoff"`
	// Upper limit of idle time of device for identification in minutes (0 - 14400, default = 1440).
	DeviceIdle pulumi.IntOutput `pulumi:"deviceIdle"`
	// Upper limit of confidence of device for identification (0 - 255, default = 1, 0 = disable).
	DeviceWeight pulumi.IntOutput `pulumi:"deviceWeight"`
	// Enable/disable allowing Virtual Access Points (VAPs) to use the same SSID name in the same VDOM. Valid values: `enable`, `disable`.
	DuplicateSsid pulumi.StringOutput `pulumi:"duplicateSsid"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Actions taken for detected fake SSID. Valid values: `log`, `suppress`.
	FakeSsidAction pulumi.StringOutput `pulumi:"fakeSsidAction"`
	// Enable/disable FAP-C series compatibility. Valid values: `enable`, `disable`.
	FapcCompatibility pulumi.StringOutput `pulumi:"fapcCompatibility"`
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization pulumi.StringOutput `pulumi:"firmwareProvisionOnAuthorization"`
	// Configure offending SSID. The structure of `offendingSsid` block is documented below.
	OffendingSsids WirelesscontrollerSettingOffendingSsidArrayOutput `pulumi:"offendingSsids"`
	// Enable/disable phishing SSID detection. Valid values: `enable`, `disable`.
	PhishingSsidDetect pulumi.StringOutput `pulumi:"phishingSsidDetect"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable WFA compatibility. Valid values: `enable`, `disable`.
	WfaCompatibility pulumi.StringOutput `pulumi:"wfaCompatibility"`
}

// NewWirelesscontrollerSetting registers a new resource with the given unique name, arguments, and options.
func NewWirelesscontrollerSetting(ctx *pulumi.Context,
	name string, args *WirelesscontrollerSettingArgs, opts ...pulumi.ResourceOption) (*WirelesscontrollerSetting, error) {
	if args == nil {
		args = &WirelesscontrollerSettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelesscontrollerSetting
	err := ctx.RegisterResource("fortios:index/wirelesscontrollerSetting:WirelesscontrollerSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelesscontrollerSetting gets an existing WirelesscontrollerSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelesscontrollerSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelesscontrollerSettingState, opts ...pulumi.ResourceOption) (*WirelesscontrollerSetting, error) {
	var resource WirelesscontrollerSetting
	err := ctx.ReadResource("fortios:index/wirelesscontrollerSetting:WirelesscontrollerSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelesscontrollerSetting resources.
type wirelesscontrollerSettingState struct {
	// FortiCloud customer account ID.
	AccountId *string `pulumi:"accountId"`
	// Country or region in which the FortiGate is located. The country determines the 802.11 bands and channels that are available.
	Country *string `pulumi:"country"`
	// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
	DarrpOptimize *int `pulumi:"darrpOptimize"`
	// Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
	DarrpOptimizeSchedules []WirelesscontrollerSettingDarrpOptimizeSchedule `pulumi:"darrpOptimizeSchedules"`
	// Lower limit of creation time of device for identification in minutes (0 - 60, default = 5).
	DeviceHoldoff *int `pulumi:"deviceHoldoff"`
	// Upper limit of idle time of device for identification in minutes (0 - 14400, default = 1440).
	DeviceIdle *int `pulumi:"deviceIdle"`
	// Upper limit of confidence of device for identification (0 - 255, default = 1, 0 = disable).
	DeviceWeight *int `pulumi:"deviceWeight"`
	// Enable/disable allowing Virtual Access Points (VAPs) to use the same SSID name in the same VDOM. Valid values: `enable`, `disable`.
	DuplicateSsid *string `pulumi:"duplicateSsid"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Actions taken for detected fake SSID. Valid values: `log`, `suppress`.
	FakeSsidAction *string `pulumi:"fakeSsidAction"`
	// Enable/disable FAP-C series compatibility. Valid values: `enable`, `disable`.
	FapcCompatibility *string `pulumi:"fapcCompatibility"`
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization *string `pulumi:"firmwareProvisionOnAuthorization"`
	// Configure offending SSID. The structure of `offendingSsid` block is documented below.
	OffendingSsids []WirelesscontrollerSettingOffendingSsid `pulumi:"offendingSsids"`
	// Enable/disable phishing SSID detection. Valid values: `enable`, `disable`.
	PhishingSsidDetect *string `pulumi:"phishingSsidDetect"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable WFA compatibility. Valid values: `enable`, `disable`.
	WfaCompatibility *string `pulumi:"wfaCompatibility"`
}

type WirelesscontrollerSettingState struct {
	// FortiCloud customer account ID.
	AccountId pulumi.StringPtrInput
	// Country or region in which the FortiGate is located. The country determines the 802.11 bands and channels that are available.
	Country pulumi.StringPtrInput
	// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
	DarrpOptimize pulumi.IntPtrInput
	// Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
	DarrpOptimizeSchedules WirelesscontrollerSettingDarrpOptimizeScheduleArrayInput
	// Lower limit of creation time of device for identification in minutes (0 - 60, default = 5).
	DeviceHoldoff pulumi.IntPtrInput
	// Upper limit of idle time of device for identification in minutes (0 - 14400, default = 1440).
	DeviceIdle pulumi.IntPtrInput
	// Upper limit of confidence of device for identification (0 - 255, default = 1, 0 = disable).
	DeviceWeight pulumi.IntPtrInput
	// Enable/disable allowing Virtual Access Points (VAPs) to use the same SSID name in the same VDOM. Valid values: `enable`, `disable`.
	DuplicateSsid pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Actions taken for detected fake SSID. Valid values: `log`, `suppress`.
	FakeSsidAction pulumi.StringPtrInput
	// Enable/disable FAP-C series compatibility. Valid values: `enable`, `disable`.
	FapcCompatibility pulumi.StringPtrInput
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization pulumi.StringPtrInput
	// Configure offending SSID. The structure of `offendingSsid` block is documented below.
	OffendingSsids WirelesscontrollerSettingOffendingSsidArrayInput
	// Enable/disable phishing SSID detection. Valid values: `enable`, `disable`.
	PhishingSsidDetect pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable WFA compatibility. Valid values: `enable`, `disable`.
	WfaCompatibility pulumi.StringPtrInput
}

func (WirelesscontrollerSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerSettingState)(nil)).Elem()
}

type wirelesscontrollerSettingArgs struct {
	// FortiCloud customer account ID.
	AccountId *string `pulumi:"accountId"`
	// Country or region in which the FortiGate is located. The country determines the 802.11 bands and channels that are available.
	Country *string `pulumi:"country"`
	// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
	DarrpOptimize *int `pulumi:"darrpOptimize"`
	// Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
	DarrpOptimizeSchedules []WirelesscontrollerSettingDarrpOptimizeSchedule `pulumi:"darrpOptimizeSchedules"`
	// Lower limit of creation time of device for identification in minutes (0 - 60, default = 5).
	DeviceHoldoff *int `pulumi:"deviceHoldoff"`
	// Upper limit of idle time of device for identification in minutes (0 - 14400, default = 1440).
	DeviceIdle *int `pulumi:"deviceIdle"`
	// Upper limit of confidence of device for identification (0 - 255, default = 1, 0 = disable).
	DeviceWeight *int `pulumi:"deviceWeight"`
	// Enable/disable allowing Virtual Access Points (VAPs) to use the same SSID name in the same VDOM. Valid values: `enable`, `disable`.
	DuplicateSsid *string `pulumi:"duplicateSsid"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Actions taken for detected fake SSID. Valid values: `log`, `suppress`.
	FakeSsidAction *string `pulumi:"fakeSsidAction"`
	// Enable/disable FAP-C series compatibility. Valid values: `enable`, `disable`.
	FapcCompatibility *string `pulumi:"fapcCompatibility"`
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization *string `pulumi:"firmwareProvisionOnAuthorization"`
	// Configure offending SSID. The structure of `offendingSsid` block is documented below.
	OffendingSsids []WirelesscontrollerSettingOffendingSsid `pulumi:"offendingSsids"`
	// Enable/disable phishing SSID detection. Valid values: `enable`, `disable`.
	PhishingSsidDetect *string `pulumi:"phishingSsidDetect"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable WFA compatibility. Valid values: `enable`, `disable`.
	WfaCompatibility *string `pulumi:"wfaCompatibility"`
}

// The set of arguments for constructing a WirelesscontrollerSetting resource.
type WirelesscontrollerSettingArgs struct {
	// FortiCloud customer account ID.
	AccountId pulumi.StringPtrInput
	// Country or region in which the FortiGate is located. The country determines the 802.11 bands and channels that are available.
	Country pulumi.StringPtrInput
	// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
	DarrpOptimize pulumi.IntPtrInput
	// Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
	DarrpOptimizeSchedules WirelesscontrollerSettingDarrpOptimizeScheduleArrayInput
	// Lower limit of creation time of device for identification in minutes (0 - 60, default = 5).
	DeviceHoldoff pulumi.IntPtrInput
	// Upper limit of idle time of device for identification in minutes (0 - 14400, default = 1440).
	DeviceIdle pulumi.IntPtrInput
	// Upper limit of confidence of device for identification (0 - 255, default = 1, 0 = disable).
	DeviceWeight pulumi.IntPtrInput
	// Enable/disable allowing Virtual Access Points (VAPs) to use the same SSID name in the same VDOM. Valid values: `enable`, `disable`.
	DuplicateSsid pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Actions taken for detected fake SSID. Valid values: `log`, `suppress`.
	FakeSsidAction pulumi.StringPtrInput
	// Enable/disable FAP-C series compatibility. Valid values: `enable`, `disable`.
	FapcCompatibility pulumi.StringPtrInput
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization pulumi.StringPtrInput
	// Configure offending SSID. The structure of `offendingSsid` block is documented below.
	OffendingSsids WirelesscontrollerSettingOffendingSsidArrayInput
	// Enable/disable phishing SSID detection. Valid values: `enable`, `disable`.
	PhishingSsidDetect pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable WFA compatibility. Valid values: `enable`, `disable`.
	WfaCompatibility pulumi.StringPtrInput
}

func (WirelesscontrollerSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerSettingArgs)(nil)).Elem()
}

type WirelesscontrollerSettingInput interface {
	pulumi.Input

	ToWirelesscontrollerSettingOutput() WirelesscontrollerSettingOutput
	ToWirelesscontrollerSettingOutputWithContext(ctx context.Context) WirelesscontrollerSettingOutput
}

func (*WirelesscontrollerSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerSetting)(nil)).Elem()
}

func (i *WirelesscontrollerSetting) ToWirelesscontrollerSettingOutput() WirelesscontrollerSettingOutput {
	return i.ToWirelesscontrollerSettingOutputWithContext(context.Background())
}

func (i *WirelesscontrollerSetting) ToWirelesscontrollerSettingOutputWithContext(ctx context.Context) WirelesscontrollerSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerSettingOutput)
}

// WirelesscontrollerSettingArrayInput is an input type that accepts WirelesscontrollerSettingArray and WirelesscontrollerSettingArrayOutput values.
// You can construct a concrete instance of `WirelesscontrollerSettingArrayInput` via:
//
//	WirelesscontrollerSettingArray{ WirelesscontrollerSettingArgs{...} }
type WirelesscontrollerSettingArrayInput interface {
	pulumi.Input

	ToWirelesscontrollerSettingArrayOutput() WirelesscontrollerSettingArrayOutput
	ToWirelesscontrollerSettingArrayOutputWithContext(context.Context) WirelesscontrollerSettingArrayOutput
}

type WirelesscontrollerSettingArray []WirelesscontrollerSettingInput

func (WirelesscontrollerSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerSetting)(nil)).Elem()
}

func (i WirelesscontrollerSettingArray) ToWirelesscontrollerSettingArrayOutput() WirelesscontrollerSettingArrayOutput {
	return i.ToWirelesscontrollerSettingArrayOutputWithContext(context.Background())
}

func (i WirelesscontrollerSettingArray) ToWirelesscontrollerSettingArrayOutputWithContext(ctx context.Context) WirelesscontrollerSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerSettingArrayOutput)
}

// WirelesscontrollerSettingMapInput is an input type that accepts WirelesscontrollerSettingMap and WirelesscontrollerSettingMapOutput values.
// You can construct a concrete instance of `WirelesscontrollerSettingMapInput` via:
//
//	WirelesscontrollerSettingMap{ "key": WirelesscontrollerSettingArgs{...} }
type WirelesscontrollerSettingMapInput interface {
	pulumi.Input

	ToWirelesscontrollerSettingMapOutput() WirelesscontrollerSettingMapOutput
	ToWirelesscontrollerSettingMapOutputWithContext(context.Context) WirelesscontrollerSettingMapOutput
}

type WirelesscontrollerSettingMap map[string]WirelesscontrollerSettingInput

func (WirelesscontrollerSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerSetting)(nil)).Elem()
}

func (i WirelesscontrollerSettingMap) ToWirelesscontrollerSettingMapOutput() WirelesscontrollerSettingMapOutput {
	return i.ToWirelesscontrollerSettingMapOutputWithContext(context.Background())
}

func (i WirelesscontrollerSettingMap) ToWirelesscontrollerSettingMapOutputWithContext(ctx context.Context) WirelesscontrollerSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerSettingMapOutput)
}

type WirelesscontrollerSettingOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerSetting)(nil)).Elem()
}

func (o WirelesscontrollerSettingOutput) ToWirelesscontrollerSettingOutput() WirelesscontrollerSettingOutput {
	return o
}

func (o WirelesscontrollerSettingOutput) ToWirelesscontrollerSettingOutputWithContext(ctx context.Context) WirelesscontrollerSettingOutput {
	return o
}

// FortiCloud customer account ID.
func (o WirelesscontrollerSettingOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// Country or region in which the FortiGate is located. The country determines the 802.11 bands and channels that are available.
func (o WirelesscontrollerSettingOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.StringOutput { return v.Country }).(pulumi.StringOutput)
}

// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
func (o WirelesscontrollerSettingOutput) DarrpOptimize() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.IntOutput { return v.DarrpOptimize }).(pulumi.IntOutput)
}

// Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
func (o WirelesscontrollerSettingOutput) DarrpOptimizeSchedules() WirelesscontrollerSettingDarrpOptimizeScheduleArrayOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) WirelesscontrollerSettingDarrpOptimizeScheduleArrayOutput {
		return v.DarrpOptimizeSchedules
	}).(WirelesscontrollerSettingDarrpOptimizeScheduleArrayOutput)
}

// Lower limit of creation time of device for identification in minutes (0 - 60, default = 5).
func (o WirelesscontrollerSettingOutput) DeviceHoldoff() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.IntOutput { return v.DeviceHoldoff }).(pulumi.IntOutput)
}

// Upper limit of idle time of device for identification in minutes (0 - 14400, default = 1440).
func (o WirelesscontrollerSettingOutput) DeviceIdle() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.IntOutput { return v.DeviceIdle }).(pulumi.IntOutput)
}

// Upper limit of confidence of device for identification (0 - 255, default = 1, 0 = disable).
func (o WirelesscontrollerSettingOutput) DeviceWeight() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.IntOutput { return v.DeviceWeight }).(pulumi.IntOutput)
}

// Enable/disable allowing Virtual Access Points (VAPs) to use the same SSID name in the same VDOM. Valid values: `enable`, `disable`.
func (o WirelesscontrollerSettingOutput) DuplicateSsid() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.StringOutput { return v.DuplicateSsid }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o WirelesscontrollerSettingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Actions taken for detected fake SSID. Valid values: `log`, `suppress`.
func (o WirelesscontrollerSettingOutput) FakeSsidAction() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.StringOutput { return v.FakeSsidAction }).(pulumi.StringOutput)
}

// Enable/disable FAP-C series compatibility. Valid values: `enable`, `disable`.
func (o WirelesscontrollerSettingOutput) FapcCompatibility() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.StringOutput { return v.FapcCompatibility }).(pulumi.StringOutput)
}

// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
func (o WirelesscontrollerSettingOutput) FirmwareProvisionOnAuthorization() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.StringOutput { return v.FirmwareProvisionOnAuthorization }).(pulumi.StringOutput)
}

// Configure offending SSID. The structure of `offendingSsid` block is documented below.
func (o WirelesscontrollerSettingOutput) OffendingSsids() WirelesscontrollerSettingOffendingSsidArrayOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) WirelesscontrollerSettingOffendingSsidArrayOutput {
		return v.OffendingSsids
	}).(WirelesscontrollerSettingOffendingSsidArrayOutput)
}

// Enable/disable phishing SSID detection. Valid values: `enable`, `disable`.
func (o WirelesscontrollerSettingOutput) PhishingSsidDetect() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.StringOutput { return v.PhishingSsidDetect }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WirelesscontrollerSettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable WFA compatibility. Valid values: `enable`, `disable`.
func (o WirelesscontrollerSettingOutput) WfaCompatibility() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerSetting) pulumi.StringOutput { return v.WfaCompatibility }).(pulumi.StringOutput)
}

type WirelesscontrollerSettingArrayOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerSetting)(nil)).Elem()
}

func (o WirelesscontrollerSettingArrayOutput) ToWirelesscontrollerSettingArrayOutput() WirelesscontrollerSettingArrayOutput {
	return o
}

func (o WirelesscontrollerSettingArrayOutput) ToWirelesscontrollerSettingArrayOutputWithContext(ctx context.Context) WirelesscontrollerSettingArrayOutput {
	return o
}

func (o WirelesscontrollerSettingArrayOutput) Index(i pulumi.IntInput) WirelesscontrollerSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelesscontrollerSetting {
		return vs[0].([]*WirelesscontrollerSetting)[vs[1].(int)]
	}).(WirelesscontrollerSettingOutput)
}

type WirelesscontrollerSettingMapOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerSetting)(nil)).Elem()
}

func (o WirelesscontrollerSettingMapOutput) ToWirelesscontrollerSettingMapOutput() WirelesscontrollerSettingMapOutput {
	return o
}

func (o WirelesscontrollerSettingMapOutput) ToWirelesscontrollerSettingMapOutputWithContext(ctx context.Context) WirelesscontrollerSettingMapOutput {
	return o
}

func (o WirelesscontrollerSettingMapOutput) MapIndex(k pulumi.StringInput) WirelesscontrollerSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelesscontrollerSetting {
		return vs[0].(map[string]*WirelesscontrollerSetting)[vs[1].(string)]
	}).(WirelesscontrollerSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerSettingInput)(nil)).Elem(), &WirelesscontrollerSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerSettingArrayInput)(nil)).Elem(), WirelesscontrollerSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerSettingMapInput)(nil)).Elem(), WirelesscontrollerSettingMap{})
	pulumi.RegisterOutputType(WirelesscontrollerSettingOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerSettingArrayOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerSettingMapOutput{})
}