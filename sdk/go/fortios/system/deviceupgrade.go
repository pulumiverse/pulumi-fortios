// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Independent upgrades for managed devices. Applies to FortiOS Version `>= 7.2.4`.
//
// ## Import
//
// System DeviceUpgrade can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/deviceupgrade:Deviceupgrade labelname {{serial}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/deviceupgrade:Deviceupgrade labelname {{serial}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Deviceupgrade struct {
	pulumi.CustomResourceState

	// Fortinet device type.
	DeviceType pulumi.StringOutput `pulumi:"deviceType"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Upgrade failure reason.
	FailureReason pulumi.StringOutput `pulumi:"failureReason"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Serial number of the FortiGate unit that will control the reboot process for the federated upgrade of the HA cluster.
	HaRebootController pulumi.StringOutput `pulumi:"haRebootController"`
	// Known members of the HA cluster. If a member is missing at upgrade time, the upgrade will be cancelled. The structure of `knownHaMembers` block is documented below.
	KnownHaMembers DeviceupgradeKnownHaMemberArrayOutput `pulumi:"knownHaMembers"`
	// Maximum number of minutes to allow for immediate upgrade preparation.
	MaximumMinutes pulumi.IntOutput `pulumi:"maximumMinutes"`
	// Serial number of the node to include.
	Serial pulumi.StringOutput `pulumi:"serial"`
	// Upgrade configuration time in UTC (hh:mm yyyy/mm/dd UTC).
	SetupTime pulumi.StringOutput `pulumi:"setupTime"`
	// Current status of the upgrade. Valid values: `disabled`, `initialized`, `downloading`, `device-disconnected`, `ready`, `coordinating`, `staging`, `final-check`, `upgrade-devices`, `cancelled`, `confirmed`, `done`, `failed`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Scheduled upgrade execution time in UTC (hh:mm yyyy/mm/dd UTC).
	Time pulumi.StringOutput `pulumi:"time"`
	// Run immediately or at a scheduled time. Valid values: `immediate`, `scheduled`.
	Timing pulumi.StringOutput `pulumi:"timing"`
	// Fortinet OS image versions to upgrade through in major-minor-patch format, such as 7-0-4.
	UpgradePath pulumi.StringOutput `pulumi:"upgradePath"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDeviceupgrade registers a new resource with the given unique name, arguments, and options.
func NewDeviceupgrade(ctx *pulumi.Context,
	name string, args *DeviceupgradeArgs, opts ...pulumi.ResourceOption) (*Deviceupgrade, error) {
	if args == nil {
		args = &DeviceupgradeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Deviceupgrade
	err := ctx.RegisterResource("fortios:system/deviceupgrade:Deviceupgrade", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceupgrade gets an existing Deviceupgrade resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceupgrade(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceupgradeState, opts ...pulumi.ResourceOption) (*Deviceupgrade, error) {
	var resource Deviceupgrade
	err := ctx.ReadResource("fortios:system/deviceupgrade:Deviceupgrade", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deviceupgrade resources.
type deviceupgradeState struct {
	// Fortinet device type.
	DeviceType *string `pulumi:"deviceType"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Upgrade failure reason.
	FailureReason *string `pulumi:"failureReason"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Serial number of the FortiGate unit that will control the reboot process for the federated upgrade of the HA cluster.
	HaRebootController *string `pulumi:"haRebootController"`
	// Known members of the HA cluster. If a member is missing at upgrade time, the upgrade will be cancelled. The structure of `knownHaMembers` block is documented below.
	KnownHaMembers []DeviceupgradeKnownHaMember `pulumi:"knownHaMembers"`
	// Maximum number of minutes to allow for immediate upgrade preparation.
	MaximumMinutes *int `pulumi:"maximumMinutes"`
	// Serial number of the node to include.
	Serial *string `pulumi:"serial"`
	// Upgrade configuration time in UTC (hh:mm yyyy/mm/dd UTC).
	SetupTime *string `pulumi:"setupTime"`
	// Current status of the upgrade. Valid values: `disabled`, `initialized`, `downloading`, `device-disconnected`, `ready`, `coordinating`, `staging`, `final-check`, `upgrade-devices`, `cancelled`, `confirmed`, `done`, `failed`.
	Status *string `pulumi:"status"`
	// Scheduled upgrade execution time in UTC (hh:mm yyyy/mm/dd UTC).
	Time *string `pulumi:"time"`
	// Run immediately or at a scheduled time. Valid values: `immediate`, `scheduled`.
	Timing *string `pulumi:"timing"`
	// Fortinet OS image versions to upgrade through in major-minor-patch format, such as 7-0-4.
	UpgradePath *string `pulumi:"upgradePath"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DeviceupgradeState struct {
	// Fortinet device type.
	DeviceType pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Upgrade failure reason.
	FailureReason pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Serial number of the FortiGate unit that will control the reboot process for the federated upgrade of the HA cluster.
	HaRebootController pulumi.StringPtrInput
	// Known members of the HA cluster. If a member is missing at upgrade time, the upgrade will be cancelled. The structure of `knownHaMembers` block is documented below.
	KnownHaMembers DeviceupgradeKnownHaMemberArrayInput
	// Maximum number of minutes to allow for immediate upgrade preparation.
	MaximumMinutes pulumi.IntPtrInput
	// Serial number of the node to include.
	Serial pulumi.StringPtrInput
	// Upgrade configuration time in UTC (hh:mm yyyy/mm/dd UTC).
	SetupTime pulumi.StringPtrInput
	// Current status of the upgrade. Valid values: `disabled`, `initialized`, `downloading`, `device-disconnected`, `ready`, `coordinating`, `staging`, `final-check`, `upgrade-devices`, `cancelled`, `confirmed`, `done`, `failed`.
	Status pulumi.StringPtrInput
	// Scheduled upgrade execution time in UTC (hh:mm yyyy/mm/dd UTC).
	Time pulumi.StringPtrInput
	// Run immediately or at a scheduled time. Valid values: `immediate`, `scheduled`.
	Timing pulumi.StringPtrInput
	// Fortinet OS image versions to upgrade through in major-minor-patch format, such as 7-0-4.
	UpgradePath pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DeviceupgradeState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceupgradeState)(nil)).Elem()
}

type deviceupgradeArgs struct {
	// Fortinet device type.
	DeviceType *string `pulumi:"deviceType"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Upgrade failure reason.
	FailureReason *string `pulumi:"failureReason"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Serial number of the FortiGate unit that will control the reboot process for the federated upgrade of the HA cluster.
	HaRebootController *string `pulumi:"haRebootController"`
	// Known members of the HA cluster. If a member is missing at upgrade time, the upgrade will be cancelled. The structure of `knownHaMembers` block is documented below.
	KnownHaMembers []DeviceupgradeKnownHaMember `pulumi:"knownHaMembers"`
	// Maximum number of minutes to allow for immediate upgrade preparation.
	MaximumMinutes *int `pulumi:"maximumMinutes"`
	// Serial number of the node to include.
	Serial *string `pulumi:"serial"`
	// Upgrade configuration time in UTC (hh:mm yyyy/mm/dd UTC).
	SetupTime *string `pulumi:"setupTime"`
	// Current status of the upgrade. Valid values: `disabled`, `initialized`, `downloading`, `device-disconnected`, `ready`, `coordinating`, `staging`, `final-check`, `upgrade-devices`, `cancelled`, `confirmed`, `done`, `failed`.
	Status *string `pulumi:"status"`
	// Scheduled upgrade execution time in UTC (hh:mm yyyy/mm/dd UTC).
	Time *string `pulumi:"time"`
	// Run immediately or at a scheduled time. Valid values: `immediate`, `scheduled`.
	Timing *string `pulumi:"timing"`
	// Fortinet OS image versions to upgrade through in major-minor-patch format, such as 7-0-4.
	UpgradePath *string `pulumi:"upgradePath"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Deviceupgrade resource.
type DeviceupgradeArgs struct {
	// Fortinet device type.
	DeviceType pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Upgrade failure reason.
	FailureReason pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Serial number of the FortiGate unit that will control the reboot process for the federated upgrade of the HA cluster.
	HaRebootController pulumi.StringPtrInput
	// Known members of the HA cluster. If a member is missing at upgrade time, the upgrade will be cancelled. The structure of `knownHaMembers` block is documented below.
	KnownHaMembers DeviceupgradeKnownHaMemberArrayInput
	// Maximum number of minutes to allow for immediate upgrade preparation.
	MaximumMinutes pulumi.IntPtrInput
	// Serial number of the node to include.
	Serial pulumi.StringPtrInput
	// Upgrade configuration time in UTC (hh:mm yyyy/mm/dd UTC).
	SetupTime pulumi.StringPtrInput
	// Current status of the upgrade. Valid values: `disabled`, `initialized`, `downloading`, `device-disconnected`, `ready`, `coordinating`, `staging`, `final-check`, `upgrade-devices`, `cancelled`, `confirmed`, `done`, `failed`.
	Status pulumi.StringPtrInput
	// Scheduled upgrade execution time in UTC (hh:mm yyyy/mm/dd UTC).
	Time pulumi.StringPtrInput
	// Run immediately or at a scheduled time. Valid values: `immediate`, `scheduled`.
	Timing pulumi.StringPtrInput
	// Fortinet OS image versions to upgrade through in major-minor-patch format, such as 7-0-4.
	UpgradePath pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DeviceupgradeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceupgradeArgs)(nil)).Elem()
}

type DeviceupgradeInput interface {
	pulumi.Input

	ToDeviceupgradeOutput() DeviceupgradeOutput
	ToDeviceupgradeOutputWithContext(ctx context.Context) DeviceupgradeOutput
}

func (*Deviceupgrade) ElementType() reflect.Type {
	return reflect.TypeOf((**Deviceupgrade)(nil)).Elem()
}

func (i *Deviceupgrade) ToDeviceupgradeOutput() DeviceupgradeOutput {
	return i.ToDeviceupgradeOutputWithContext(context.Background())
}

func (i *Deviceupgrade) ToDeviceupgradeOutputWithContext(ctx context.Context) DeviceupgradeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceupgradeOutput)
}

// DeviceupgradeArrayInput is an input type that accepts DeviceupgradeArray and DeviceupgradeArrayOutput values.
// You can construct a concrete instance of `DeviceupgradeArrayInput` via:
//
//	DeviceupgradeArray{ DeviceupgradeArgs{...} }
type DeviceupgradeArrayInput interface {
	pulumi.Input

	ToDeviceupgradeArrayOutput() DeviceupgradeArrayOutput
	ToDeviceupgradeArrayOutputWithContext(context.Context) DeviceupgradeArrayOutput
}

type DeviceupgradeArray []DeviceupgradeInput

func (DeviceupgradeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Deviceupgrade)(nil)).Elem()
}

func (i DeviceupgradeArray) ToDeviceupgradeArrayOutput() DeviceupgradeArrayOutput {
	return i.ToDeviceupgradeArrayOutputWithContext(context.Background())
}

func (i DeviceupgradeArray) ToDeviceupgradeArrayOutputWithContext(ctx context.Context) DeviceupgradeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceupgradeArrayOutput)
}

// DeviceupgradeMapInput is an input type that accepts DeviceupgradeMap and DeviceupgradeMapOutput values.
// You can construct a concrete instance of `DeviceupgradeMapInput` via:
//
//	DeviceupgradeMap{ "key": DeviceupgradeArgs{...} }
type DeviceupgradeMapInput interface {
	pulumi.Input

	ToDeviceupgradeMapOutput() DeviceupgradeMapOutput
	ToDeviceupgradeMapOutputWithContext(context.Context) DeviceupgradeMapOutput
}

type DeviceupgradeMap map[string]DeviceupgradeInput

func (DeviceupgradeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Deviceupgrade)(nil)).Elem()
}

func (i DeviceupgradeMap) ToDeviceupgradeMapOutput() DeviceupgradeMapOutput {
	return i.ToDeviceupgradeMapOutputWithContext(context.Background())
}

func (i DeviceupgradeMap) ToDeviceupgradeMapOutputWithContext(ctx context.Context) DeviceupgradeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceupgradeMapOutput)
}

type DeviceupgradeOutput struct{ *pulumi.OutputState }

func (DeviceupgradeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deviceupgrade)(nil)).Elem()
}

func (o DeviceupgradeOutput) ToDeviceupgradeOutput() DeviceupgradeOutput {
	return o
}

func (o DeviceupgradeOutput) ToDeviceupgradeOutputWithContext(ctx context.Context) DeviceupgradeOutput {
	return o
}

// Fortinet device type.
func (o DeviceupgradeOutput) DeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Deviceupgrade) pulumi.StringOutput { return v.DeviceType }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o DeviceupgradeOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deviceupgrade) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Upgrade failure reason.
func (o DeviceupgradeOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Deviceupgrade) pulumi.StringOutput { return v.FailureReason }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o DeviceupgradeOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deviceupgrade) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Serial number of the FortiGate unit that will control the reboot process for the federated upgrade of the HA cluster.
func (o DeviceupgradeOutput) HaRebootController() pulumi.StringOutput {
	return o.ApplyT(func(v *Deviceupgrade) pulumi.StringOutput { return v.HaRebootController }).(pulumi.StringOutput)
}

// Known members of the HA cluster. If a member is missing at upgrade time, the upgrade will be cancelled. The structure of `knownHaMembers` block is documented below.
func (o DeviceupgradeOutput) KnownHaMembers() DeviceupgradeKnownHaMemberArrayOutput {
	return o.ApplyT(func(v *Deviceupgrade) DeviceupgradeKnownHaMemberArrayOutput { return v.KnownHaMembers }).(DeviceupgradeKnownHaMemberArrayOutput)
}

// Maximum number of minutes to allow for immediate upgrade preparation.
func (o DeviceupgradeOutput) MaximumMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v *Deviceupgrade) pulumi.IntOutput { return v.MaximumMinutes }).(pulumi.IntOutput)
}

// Serial number of the node to include.
func (o DeviceupgradeOutput) Serial() pulumi.StringOutput {
	return o.ApplyT(func(v *Deviceupgrade) pulumi.StringOutput { return v.Serial }).(pulumi.StringOutput)
}

// Upgrade configuration time in UTC (hh:mm yyyy/mm/dd UTC).
func (o DeviceupgradeOutput) SetupTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Deviceupgrade) pulumi.StringOutput { return v.SetupTime }).(pulumi.StringOutput)
}

// Current status of the upgrade. Valid values: `disabled`, `initialized`, `downloading`, `device-disconnected`, `ready`, `coordinating`, `staging`, `final-check`, `upgrade-devices`, `cancelled`, `confirmed`, `done`, `failed`.
func (o DeviceupgradeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Deviceupgrade) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Scheduled upgrade execution time in UTC (hh:mm yyyy/mm/dd UTC).
func (o DeviceupgradeOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v *Deviceupgrade) pulumi.StringOutput { return v.Time }).(pulumi.StringOutput)
}

// Run immediately or at a scheduled time. Valid values: `immediate`, `scheduled`.
func (o DeviceupgradeOutput) Timing() pulumi.StringOutput {
	return o.ApplyT(func(v *Deviceupgrade) pulumi.StringOutput { return v.Timing }).(pulumi.StringOutput)
}

// Fortinet OS image versions to upgrade through in major-minor-patch format, such as 7-0-4.
func (o DeviceupgradeOutput) UpgradePath() pulumi.StringOutput {
	return o.ApplyT(func(v *Deviceupgrade) pulumi.StringOutput { return v.UpgradePath }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DeviceupgradeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deviceupgrade) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DeviceupgradeArrayOutput struct{ *pulumi.OutputState }

func (DeviceupgradeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Deviceupgrade)(nil)).Elem()
}

func (o DeviceupgradeArrayOutput) ToDeviceupgradeArrayOutput() DeviceupgradeArrayOutput {
	return o
}

func (o DeviceupgradeArrayOutput) ToDeviceupgradeArrayOutputWithContext(ctx context.Context) DeviceupgradeArrayOutput {
	return o
}

func (o DeviceupgradeArrayOutput) Index(i pulumi.IntInput) DeviceupgradeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Deviceupgrade {
		return vs[0].([]*Deviceupgrade)[vs[1].(int)]
	}).(DeviceupgradeOutput)
}

type DeviceupgradeMapOutput struct{ *pulumi.OutputState }

func (DeviceupgradeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Deviceupgrade)(nil)).Elem()
}

func (o DeviceupgradeMapOutput) ToDeviceupgradeMapOutput() DeviceupgradeMapOutput {
	return o
}

func (o DeviceupgradeMapOutput) ToDeviceupgradeMapOutputWithContext(ctx context.Context) DeviceupgradeMapOutput {
	return o
}

func (o DeviceupgradeMapOutput) MapIndex(k pulumi.StringInput) DeviceupgradeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Deviceupgrade {
		return vs[0].(map[string]*Deviceupgrade)[vs[1].(string)]
	}).(DeviceupgradeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceupgradeInput)(nil)).Elem(), &Deviceupgrade{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceupgradeArrayInput)(nil)).Elem(), DeviceupgradeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceupgradeMapInput)(nil)).Elem(), DeviceupgradeMap{})
	pulumi.RegisterOutputType(DeviceupgradeOutput{})
	pulumi.RegisterOutputType(DeviceupgradeArrayOutput{})
	pulumi.RegisterOutputType(DeviceupgradeMapOutput{})
}
