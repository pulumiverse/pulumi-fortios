// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles. Applies to FortiOS Version `>= 6.4.2`.
//
// ## Import
//
// WirelessController ArrpProfile can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/arrpprofile:Arrpprofile labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/arrpprofile:Arrpprofile labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Arrpprofile struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Time for running Distributed Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
	DarrpOptimize pulumi.IntOutput `pulumi:"darrpOptimize"`
	// Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
	DarrpOptimizeSchedules ArrpprofileDarrpOptimizeScheduleArrayOutput `pulumi:"darrpOptimizeSchedules"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
	IncludeDfsChannel pulumi.StringOutput `pulumi:"includeDfsChannel"`
	// Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
	IncludeWeatherChannel pulumi.StringOutput `pulumi:"includeWeatherChannel"`
	// Period in seconds to measure average transmit retries and receive errors (default = 300).
	MonitorPeriod pulumi.IntOutput `pulumi:"monitorPeriod"`
	// WiFi ARRP profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable to override setting darrp-optimize and darrp-optimize-schedules (default = disable). Valid values: `enable`, `disable`.
	OverrideDarrpOptimize pulumi.StringOutput `pulumi:"overrideDarrpOptimize"`
	// Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
	SelectionPeriod pulumi.IntOutput `pulumi:"selectionPeriod"`
	// Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
	ThresholdAp pulumi.IntOutput `pulumi:"thresholdAp"`
	// Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
	ThresholdChannelLoad pulumi.IntOutput `pulumi:"thresholdChannelLoad"`
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
	ThresholdNoiseFloor pulumi.StringOutput `pulumi:"thresholdNoiseFloor"`
	// Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
	ThresholdRxErrors pulumi.IntOutput `pulumi:"thresholdRxErrors"`
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
	ThresholdSpectralRssi pulumi.StringOutput `pulumi:"thresholdSpectralRssi"`
	// Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
	ThresholdTxRetries pulumi.IntOutput `pulumi:"thresholdTxRetries"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
	// Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
	WeightChannelLoad pulumi.IntOutput `pulumi:"weightChannelLoad"`
	// Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
	WeightDfsChannel pulumi.IntOutput `pulumi:"weightDfsChannel"`
	// Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
	WeightManagedAp pulumi.IntOutput `pulumi:"weightManagedAp"`
	// Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
	WeightNoiseFloor pulumi.IntOutput `pulumi:"weightNoiseFloor"`
	// Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
	WeightRogueAp pulumi.IntOutput `pulumi:"weightRogueAp"`
	// Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
	WeightSpectralRssi pulumi.IntOutput `pulumi:"weightSpectralRssi"`
	// Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
	WeightWeatherChannel pulumi.IntOutput `pulumi:"weightWeatherChannel"`
}

// NewArrpprofile registers a new resource with the given unique name, arguments, and options.
func NewArrpprofile(ctx *pulumi.Context,
	name string, args *ArrpprofileArgs, opts ...pulumi.ResourceOption) (*Arrpprofile, error) {
	if args == nil {
		args = &ArrpprofileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Arrpprofile
	err := ctx.RegisterResource("fortios:wirelesscontroller/arrpprofile:Arrpprofile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArrpprofile gets an existing Arrpprofile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArrpprofile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArrpprofileState, opts ...pulumi.ResourceOption) (*Arrpprofile, error) {
	var resource Arrpprofile
	err := ctx.ReadResource("fortios:wirelesscontroller/arrpprofile:Arrpprofile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Arrpprofile resources.
type arrpprofileState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Time for running Distributed Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
	DarrpOptimize *int `pulumi:"darrpOptimize"`
	// Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
	DarrpOptimizeSchedules []ArrpprofileDarrpOptimizeSchedule `pulumi:"darrpOptimizeSchedules"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
	IncludeDfsChannel *string `pulumi:"includeDfsChannel"`
	// Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
	IncludeWeatherChannel *string `pulumi:"includeWeatherChannel"`
	// Period in seconds to measure average transmit retries and receive errors (default = 300).
	MonitorPeriod *int `pulumi:"monitorPeriod"`
	// WiFi ARRP profile name.
	Name *string `pulumi:"name"`
	// Enable to override setting darrp-optimize and darrp-optimize-schedules (default = disable). Valid values: `enable`, `disable`.
	OverrideDarrpOptimize *string `pulumi:"overrideDarrpOptimize"`
	// Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
	SelectionPeriod *int `pulumi:"selectionPeriod"`
	// Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
	ThresholdAp *int `pulumi:"thresholdAp"`
	// Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
	ThresholdChannelLoad *int `pulumi:"thresholdChannelLoad"`
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
	ThresholdNoiseFloor *string `pulumi:"thresholdNoiseFloor"`
	// Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
	ThresholdRxErrors *int `pulumi:"thresholdRxErrors"`
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
	ThresholdSpectralRssi *string `pulumi:"thresholdSpectralRssi"`
	// Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
	ThresholdTxRetries *int `pulumi:"thresholdTxRetries"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
	WeightChannelLoad *int `pulumi:"weightChannelLoad"`
	// Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
	WeightDfsChannel *int `pulumi:"weightDfsChannel"`
	// Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
	WeightManagedAp *int `pulumi:"weightManagedAp"`
	// Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
	WeightNoiseFloor *int `pulumi:"weightNoiseFloor"`
	// Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
	WeightRogueAp *int `pulumi:"weightRogueAp"`
	// Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
	WeightSpectralRssi *int `pulumi:"weightSpectralRssi"`
	// Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
	WeightWeatherChannel *int `pulumi:"weightWeatherChannel"`
}

type ArrpprofileState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Time for running Distributed Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
	DarrpOptimize pulumi.IntPtrInput
	// Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
	DarrpOptimizeSchedules ArrpprofileDarrpOptimizeScheduleArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
	IncludeDfsChannel pulumi.StringPtrInput
	// Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
	IncludeWeatherChannel pulumi.StringPtrInput
	// Period in seconds to measure average transmit retries and receive errors (default = 300).
	MonitorPeriod pulumi.IntPtrInput
	// WiFi ARRP profile name.
	Name pulumi.StringPtrInput
	// Enable to override setting darrp-optimize and darrp-optimize-schedules (default = disable). Valid values: `enable`, `disable`.
	OverrideDarrpOptimize pulumi.StringPtrInput
	// Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
	SelectionPeriod pulumi.IntPtrInput
	// Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
	ThresholdAp pulumi.IntPtrInput
	// Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
	ThresholdChannelLoad pulumi.IntPtrInput
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
	ThresholdNoiseFloor pulumi.StringPtrInput
	// Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
	ThresholdRxErrors pulumi.IntPtrInput
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
	ThresholdSpectralRssi pulumi.StringPtrInput
	// Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
	ThresholdTxRetries pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
	WeightChannelLoad pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
	WeightDfsChannel pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
	WeightManagedAp pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
	WeightNoiseFloor pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
	WeightRogueAp pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
	WeightSpectralRssi pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
	WeightWeatherChannel pulumi.IntPtrInput
}

func (ArrpprofileState) ElementType() reflect.Type {
	return reflect.TypeOf((*arrpprofileState)(nil)).Elem()
}

type arrpprofileArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Time for running Distributed Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
	DarrpOptimize *int `pulumi:"darrpOptimize"`
	// Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
	DarrpOptimizeSchedules []ArrpprofileDarrpOptimizeSchedule `pulumi:"darrpOptimizeSchedules"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
	IncludeDfsChannel *string `pulumi:"includeDfsChannel"`
	// Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
	IncludeWeatherChannel *string `pulumi:"includeWeatherChannel"`
	// Period in seconds to measure average transmit retries and receive errors (default = 300).
	MonitorPeriod *int `pulumi:"monitorPeriod"`
	// WiFi ARRP profile name.
	Name *string `pulumi:"name"`
	// Enable to override setting darrp-optimize and darrp-optimize-schedules (default = disable). Valid values: `enable`, `disable`.
	OverrideDarrpOptimize *string `pulumi:"overrideDarrpOptimize"`
	// Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
	SelectionPeriod *int `pulumi:"selectionPeriod"`
	// Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
	ThresholdAp *int `pulumi:"thresholdAp"`
	// Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
	ThresholdChannelLoad *int `pulumi:"thresholdChannelLoad"`
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
	ThresholdNoiseFloor *string `pulumi:"thresholdNoiseFloor"`
	// Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
	ThresholdRxErrors *int `pulumi:"thresholdRxErrors"`
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
	ThresholdSpectralRssi *string `pulumi:"thresholdSpectralRssi"`
	// Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
	ThresholdTxRetries *int `pulumi:"thresholdTxRetries"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
	WeightChannelLoad *int `pulumi:"weightChannelLoad"`
	// Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
	WeightDfsChannel *int `pulumi:"weightDfsChannel"`
	// Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
	WeightManagedAp *int `pulumi:"weightManagedAp"`
	// Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
	WeightNoiseFloor *int `pulumi:"weightNoiseFloor"`
	// Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
	WeightRogueAp *int `pulumi:"weightRogueAp"`
	// Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
	WeightSpectralRssi *int `pulumi:"weightSpectralRssi"`
	// Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
	WeightWeatherChannel *int `pulumi:"weightWeatherChannel"`
}

// The set of arguments for constructing a Arrpprofile resource.
type ArrpprofileArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Time for running Distributed Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
	DarrpOptimize pulumi.IntPtrInput
	// Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
	DarrpOptimizeSchedules ArrpprofileDarrpOptimizeScheduleArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
	IncludeDfsChannel pulumi.StringPtrInput
	// Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
	IncludeWeatherChannel pulumi.StringPtrInput
	// Period in seconds to measure average transmit retries and receive errors (default = 300).
	MonitorPeriod pulumi.IntPtrInput
	// WiFi ARRP profile name.
	Name pulumi.StringPtrInput
	// Enable to override setting darrp-optimize and darrp-optimize-schedules (default = disable). Valid values: `enable`, `disable`.
	OverrideDarrpOptimize pulumi.StringPtrInput
	// Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
	SelectionPeriod pulumi.IntPtrInput
	// Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
	ThresholdAp pulumi.IntPtrInput
	// Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
	ThresholdChannelLoad pulumi.IntPtrInput
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
	ThresholdNoiseFloor pulumi.StringPtrInput
	// Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
	ThresholdRxErrors pulumi.IntPtrInput
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
	ThresholdSpectralRssi pulumi.StringPtrInput
	// Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
	ThresholdTxRetries pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
	WeightChannelLoad pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
	WeightDfsChannel pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
	WeightManagedAp pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
	WeightNoiseFloor pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
	WeightRogueAp pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
	WeightSpectralRssi pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
	WeightWeatherChannel pulumi.IntPtrInput
}

func (ArrpprofileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*arrpprofileArgs)(nil)).Elem()
}

type ArrpprofileInput interface {
	pulumi.Input

	ToArrpprofileOutput() ArrpprofileOutput
	ToArrpprofileOutputWithContext(ctx context.Context) ArrpprofileOutput
}

func (*Arrpprofile) ElementType() reflect.Type {
	return reflect.TypeOf((**Arrpprofile)(nil)).Elem()
}

func (i *Arrpprofile) ToArrpprofileOutput() ArrpprofileOutput {
	return i.ToArrpprofileOutputWithContext(context.Background())
}

func (i *Arrpprofile) ToArrpprofileOutputWithContext(ctx context.Context) ArrpprofileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArrpprofileOutput)
}

// ArrpprofileArrayInput is an input type that accepts ArrpprofileArray and ArrpprofileArrayOutput values.
// You can construct a concrete instance of `ArrpprofileArrayInput` via:
//
//	ArrpprofileArray{ ArrpprofileArgs{...} }
type ArrpprofileArrayInput interface {
	pulumi.Input

	ToArrpprofileArrayOutput() ArrpprofileArrayOutput
	ToArrpprofileArrayOutputWithContext(context.Context) ArrpprofileArrayOutput
}

type ArrpprofileArray []ArrpprofileInput

func (ArrpprofileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Arrpprofile)(nil)).Elem()
}

func (i ArrpprofileArray) ToArrpprofileArrayOutput() ArrpprofileArrayOutput {
	return i.ToArrpprofileArrayOutputWithContext(context.Background())
}

func (i ArrpprofileArray) ToArrpprofileArrayOutputWithContext(ctx context.Context) ArrpprofileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArrpprofileArrayOutput)
}

// ArrpprofileMapInput is an input type that accepts ArrpprofileMap and ArrpprofileMapOutput values.
// You can construct a concrete instance of `ArrpprofileMapInput` via:
//
//	ArrpprofileMap{ "key": ArrpprofileArgs{...} }
type ArrpprofileMapInput interface {
	pulumi.Input

	ToArrpprofileMapOutput() ArrpprofileMapOutput
	ToArrpprofileMapOutputWithContext(context.Context) ArrpprofileMapOutput
}

type ArrpprofileMap map[string]ArrpprofileInput

func (ArrpprofileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Arrpprofile)(nil)).Elem()
}

func (i ArrpprofileMap) ToArrpprofileMapOutput() ArrpprofileMapOutput {
	return i.ToArrpprofileMapOutputWithContext(context.Background())
}

func (i ArrpprofileMap) ToArrpprofileMapOutputWithContext(ctx context.Context) ArrpprofileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArrpprofileMapOutput)
}

type ArrpprofileOutput struct{ *pulumi.OutputState }

func (ArrpprofileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Arrpprofile)(nil)).Elem()
}

func (o ArrpprofileOutput) ToArrpprofileOutput() ArrpprofileOutput {
	return o
}

func (o ArrpprofileOutput) ToArrpprofileOutputWithContext(ctx context.Context) ArrpprofileOutput {
	return o
}

// Comment.
func (o ArrpprofileOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Time for running Distributed Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
func (o ArrpprofileOutput) DarrpOptimize() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.DarrpOptimize }).(pulumi.IntOutput)
}

// Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
func (o ArrpprofileOutput) DarrpOptimizeSchedules() ArrpprofileDarrpOptimizeScheduleArrayOutput {
	return o.ApplyT(func(v *Arrpprofile) ArrpprofileDarrpOptimizeScheduleArrayOutput { return v.DarrpOptimizeSchedules }).(ArrpprofileDarrpOptimizeScheduleArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ArrpprofileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o ArrpprofileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
func (o ArrpprofileOutput) IncludeDfsChannel() pulumi.StringOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.StringOutput { return v.IncludeDfsChannel }).(pulumi.StringOutput)
}

// Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
func (o ArrpprofileOutput) IncludeWeatherChannel() pulumi.StringOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.StringOutput { return v.IncludeWeatherChannel }).(pulumi.StringOutput)
}

// Period in seconds to measure average transmit retries and receive errors (default = 300).
func (o ArrpprofileOutput) MonitorPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.MonitorPeriod }).(pulumi.IntOutput)
}

// WiFi ARRP profile name.
func (o ArrpprofileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable to override setting darrp-optimize and darrp-optimize-schedules (default = disable). Valid values: `enable`, `disable`.
func (o ArrpprofileOutput) OverrideDarrpOptimize() pulumi.StringOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.StringOutput { return v.OverrideDarrpOptimize }).(pulumi.StringOutput)
}

// Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
func (o ArrpprofileOutput) SelectionPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.SelectionPeriod }).(pulumi.IntOutput)
}

// Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
func (o ArrpprofileOutput) ThresholdAp() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.ThresholdAp }).(pulumi.IntOutput)
}

// Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
func (o ArrpprofileOutput) ThresholdChannelLoad() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.ThresholdChannelLoad }).(pulumi.IntOutput)
}

// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
func (o ArrpprofileOutput) ThresholdNoiseFloor() pulumi.StringOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.StringOutput { return v.ThresholdNoiseFloor }).(pulumi.StringOutput)
}

// Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
func (o ArrpprofileOutput) ThresholdRxErrors() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.ThresholdRxErrors }).(pulumi.IntOutput)
}

// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
func (o ArrpprofileOutput) ThresholdSpectralRssi() pulumi.StringOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.StringOutput { return v.ThresholdSpectralRssi }).(pulumi.StringOutput)
}

// Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
func (o ArrpprofileOutput) ThresholdTxRetries() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.ThresholdTxRetries }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ArrpprofileOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

// Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
func (o ArrpprofileOutput) WeightChannelLoad() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.WeightChannelLoad }).(pulumi.IntOutput)
}

// Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
func (o ArrpprofileOutput) WeightDfsChannel() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.WeightDfsChannel }).(pulumi.IntOutput)
}

// Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
func (o ArrpprofileOutput) WeightManagedAp() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.WeightManagedAp }).(pulumi.IntOutput)
}

// Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
func (o ArrpprofileOutput) WeightNoiseFloor() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.WeightNoiseFloor }).(pulumi.IntOutput)
}

// Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
func (o ArrpprofileOutput) WeightRogueAp() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.WeightRogueAp }).(pulumi.IntOutput)
}

// Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
func (o ArrpprofileOutput) WeightSpectralRssi() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.WeightSpectralRssi }).(pulumi.IntOutput)
}

// Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
func (o ArrpprofileOutput) WeightWeatherChannel() pulumi.IntOutput {
	return o.ApplyT(func(v *Arrpprofile) pulumi.IntOutput { return v.WeightWeatherChannel }).(pulumi.IntOutput)
}

type ArrpprofileArrayOutput struct{ *pulumi.OutputState }

func (ArrpprofileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Arrpprofile)(nil)).Elem()
}

func (o ArrpprofileArrayOutput) ToArrpprofileArrayOutput() ArrpprofileArrayOutput {
	return o
}

func (o ArrpprofileArrayOutput) ToArrpprofileArrayOutputWithContext(ctx context.Context) ArrpprofileArrayOutput {
	return o
}

func (o ArrpprofileArrayOutput) Index(i pulumi.IntInput) ArrpprofileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Arrpprofile {
		return vs[0].([]*Arrpprofile)[vs[1].(int)]
	}).(ArrpprofileOutput)
}

type ArrpprofileMapOutput struct{ *pulumi.OutputState }

func (ArrpprofileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Arrpprofile)(nil)).Elem()
}

func (o ArrpprofileMapOutput) ToArrpprofileMapOutput() ArrpprofileMapOutput {
	return o
}

func (o ArrpprofileMapOutput) ToArrpprofileMapOutputWithContext(ctx context.Context) ArrpprofileMapOutput {
	return o
}

func (o ArrpprofileMapOutput) MapIndex(k pulumi.StringInput) ArrpprofileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Arrpprofile {
		return vs[0].(map[string]*Arrpprofile)[vs[1].(string)]
	}).(ArrpprofileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ArrpprofileInput)(nil)).Elem(), &Arrpprofile{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArrpprofileArrayInput)(nil)).Elem(), ArrpprofileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArrpprofileMapInput)(nil)).Elem(), ArrpprofileMap{})
	pulumi.RegisterOutputType(ArrpprofileOutput{})
	pulumi.RegisterOutputType(ArrpprofileArrayOutput{})
	pulumi.RegisterOutputType(ArrpprofileMapOutput{})
}
