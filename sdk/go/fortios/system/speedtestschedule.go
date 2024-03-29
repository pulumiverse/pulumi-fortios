// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Speed test schedule for each interface. Applies to FortiOS Version `>= 7.0.0`.
//
// ## Import
//
// System SpeedTestSchedule can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/speedtestschedule:Speedtestschedule labelname {{interface}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/speedtestschedule:Speedtestschedule labelname {{interface}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Speedtestschedule struct {
	pulumi.CustomResourceState

	// Port of the controller to get access token.
	CtrlPort pulumi.IntOutput `pulumi:"ctrlPort"`
	// DSCP used for speed test.
	Diffserv pulumi.StringOutput `pulumi:"diffserv"`
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer pulumi.StringOutput `pulumi:"dynamicServer"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Interface name.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Protocol Auto(default), TCP or UDP used for speed test. Valid values: `UDP`, `TCP`, `Auto`.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules SpeedtestscheduleScheduleArrayOutput `pulumi:"schedules"`
	// Speed test server name.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Port of the server to run speed test.
	ServerPort pulumi.IntOutput `pulumi:"serverPort"`
	// Enable/disable scheduled speed test. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateInbandwidth pulumi.StringOutput `pulumi:"updateInbandwidth"`
	// Maximum downloading bandwidth (kbps) to be used in a speed test.
	UpdateInbandwidthMaximum pulumi.IntOutput `pulumi:"updateInbandwidthMaximum"`
	// Minimum downloading bandwidth (kbps) to be considered effective.
	UpdateInbandwidthMinimum pulumi.IntOutput `pulumi:"updateInbandwidthMinimum"`
	// Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateOutbandwidth pulumi.StringOutput `pulumi:"updateOutbandwidth"`
	// Maximum uploading bandwidth (kbps) to be used in a speed test.
	UpdateOutbandwidthMaximum pulumi.IntOutput `pulumi:"updateOutbandwidthMaximum"`
	// Minimum uploading bandwidth (kbps) to be considered effective.
	UpdateOutbandwidthMinimum pulumi.IntOutput `pulumi:"updateOutbandwidthMinimum"`
	// Set egress shaper based on the test result. Valid values: `disable`, `local`, `remote`, `both`.
	UpdateShaper pulumi.StringOutput `pulumi:"updateShaper"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSpeedtestschedule registers a new resource with the given unique name, arguments, and options.
func NewSpeedtestschedule(ctx *pulumi.Context,
	name string, args *SpeedtestscheduleArgs, opts ...pulumi.ResourceOption) (*Speedtestschedule, error) {
	if args == nil {
		args = &SpeedtestscheduleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Speedtestschedule
	err := ctx.RegisterResource("fortios:system/speedtestschedule:Speedtestschedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpeedtestschedule gets an existing Speedtestschedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpeedtestschedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpeedtestscheduleState, opts ...pulumi.ResourceOption) (*Speedtestschedule, error) {
	var resource Speedtestschedule
	err := ctx.ReadResource("fortios:system/speedtestschedule:Speedtestschedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Speedtestschedule resources.
type speedtestscheduleState struct {
	// Port of the controller to get access token.
	CtrlPort *int `pulumi:"ctrlPort"`
	// DSCP used for speed test.
	Diffserv *string `pulumi:"diffserv"`
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer *string `pulumi:"dynamicServer"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// Protocol Auto(default), TCP or UDP used for speed test. Valid values: `UDP`, `TCP`, `Auto`.
	Mode *string `pulumi:"mode"`
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules []SpeedtestscheduleSchedule `pulumi:"schedules"`
	// Speed test server name.
	ServerName *string `pulumi:"serverName"`
	// Port of the server to run speed test.
	ServerPort *int `pulumi:"serverPort"`
	// Enable/disable scheduled speed test. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateInbandwidth *string `pulumi:"updateInbandwidth"`
	// Maximum downloading bandwidth (kbps) to be used in a speed test.
	UpdateInbandwidthMaximum *int `pulumi:"updateInbandwidthMaximum"`
	// Minimum downloading bandwidth (kbps) to be considered effective.
	UpdateInbandwidthMinimum *int `pulumi:"updateInbandwidthMinimum"`
	// Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateOutbandwidth *string `pulumi:"updateOutbandwidth"`
	// Maximum uploading bandwidth (kbps) to be used in a speed test.
	UpdateOutbandwidthMaximum *int `pulumi:"updateOutbandwidthMaximum"`
	// Minimum uploading bandwidth (kbps) to be considered effective.
	UpdateOutbandwidthMinimum *int `pulumi:"updateOutbandwidthMinimum"`
	// Set egress shaper based on the test result. Valid values: `disable`, `local`, `remote`, `both`.
	UpdateShaper *string `pulumi:"updateShaper"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SpeedtestscheduleState struct {
	// Port of the controller to get access token.
	CtrlPort pulumi.IntPtrInput
	// DSCP used for speed test.
	Diffserv pulumi.StringPtrInput
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// Protocol Auto(default), TCP or UDP used for speed test. Valid values: `UDP`, `TCP`, `Auto`.
	Mode pulumi.StringPtrInput
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules SpeedtestscheduleScheduleArrayInput
	// Speed test server name.
	ServerName pulumi.StringPtrInput
	// Port of the server to run speed test.
	ServerPort pulumi.IntPtrInput
	// Enable/disable scheduled speed test. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateInbandwidth pulumi.StringPtrInput
	// Maximum downloading bandwidth (kbps) to be used in a speed test.
	UpdateInbandwidthMaximum pulumi.IntPtrInput
	// Minimum downloading bandwidth (kbps) to be considered effective.
	UpdateInbandwidthMinimum pulumi.IntPtrInput
	// Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateOutbandwidth pulumi.StringPtrInput
	// Maximum uploading bandwidth (kbps) to be used in a speed test.
	UpdateOutbandwidthMaximum pulumi.IntPtrInput
	// Minimum uploading bandwidth (kbps) to be considered effective.
	UpdateOutbandwidthMinimum pulumi.IntPtrInput
	// Set egress shaper based on the test result. Valid values: `disable`, `local`, `remote`, `both`.
	UpdateShaper pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpeedtestscheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*speedtestscheduleState)(nil)).Elem()
}

type speedtestscheduleArgs struct {
	// Port of the controller to get access token.
	CtrlPort *int `pulumi:"ctrlPort"`
	// DSCP used for speed test.
	Diffserv *string `pulumi:"diffserv"`
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer *string `pulumi:"dynamicServer"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// Protocol Auto(default), TCP or UDP used for speed test. Valid values: `UDP`, `TCP`, `Auto`.
	Mode *string `pulumi:"mode"`
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules []SpeedtestscheduleSchedule `pulumi:"schedules"`
	// Speed test server name.
	ServerName *string `pulumi:"serverName"`
	// Port of the server to run speed test.
	ServerPort *int `pulumi:"serverPort"`
	// Enable/disable scheduled speed test. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateInbandwidth *string `pulumi:"updateInbandwidth"`
	// Maximum downloading bandwidth (kbps) to be used in a speed test.
	UpdateInbandwidthMaximum *int `pulumi:"updateInbandwidthMaximum"`
	// Minimum downloading bandwidth (kbps) to be considered effective.
	UpdateInbandwidthMinimum *int `pulumi:"updateInbandwidthMinimum"`
	// Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateOutbandwidth *string `pulumi:"updateOutbandwidth"`
	// Maximum uploading bandwidth (kbps) to be used in a speed test.
	UpdateOutbandwidthMaximum *int `pulumi:"updateOutbandwidthMaximum"`
	// Minimum uploading bandwidth (kbps) to be considered effective.
	UpdateOutbandwidthMinimum *int `pulumi:"updateOutbandwidthMinimum"`
	// Set egress shaper based on the test result. Valid values: `disable`, `local`, `remote`, `both`.
	UpdateShaper *string `pulumi:"updateShaper"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Speedtestschedule resource.
type SpeedtestscheduleArgs struct {
	// Port of the controller to get access token.
	CtrlPort pulumi.IntPtrInput
	// DSCP used for speed test.
	Diffserv pulumi.StringPtrInput
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// Protocol Auto(default), TCP or UDP used for speed test. Valid values: `UDP`, `TCP`, `Auto`.
	Mode pulumi.StringPtrInput
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules SpeedtestscheduleScheduleArrayInput
	// Speed test server name.
	ServerName pulumi.StringPtrInput
	// Port of the server to run speed test.
	ServerPort pulumi.IntPtrInput
	// Enable/disable scheduled speed test. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateInbandwidth pulumi.StringPtrInput
	// Maximum downloading bandwidth (kbps) to be used in a speed test.
	UpdateInbandwidthMaximum pulumi.IntPtrInput
	// Minimum downloading bandwidth (kbps) to be considered effective.
	UpdateInbandwidthMinimum pulumi.IntPtrInput
	// Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateOutbandwidth pulumi.StringPtrInput
	// Maximum uploading bandwidth (kbps) to be used in a speed test.
	UpdateOutbandwidthMaximum pulumi.IntPtrInput
	// Minimum uploading bandwidth (kbps) to be considered effective.
	UpdateOutbandwidthMinimum pulumi.IntPtrInput
	// Set egress shaper based on the test result. Valid values: `disable`, `local`, `remote`, `both`.
	UpdateShaper pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpeedtestscheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*speedtestscheduleArgs)(nil)).Elem()
}

type SpeedtestscheduleInput interface {
	pulumi.Input

	ToSpeedtestscheduleOutput() SpeedtestscheduleOutput
	ToSpeedtestscheduleOutputWithContext(ctx context.Context) SpeedtestscheduleOutput
}

func (*Speedtestschedule) ElementType() reflect.Type {
	return reflect.TypeOf((**Speedtestschedule)(nil)).Elem()
}

func (i *Speedtestschedule) ToSpeedtestscheduleOutput() SpeedtestscheduleOutput {
	return i.ToSpeedtestscheduleOutputWithContext(context.Background())
}

func (i *Speedtestschedule) ToSpeedtestscheduleOutputWithContext(ctx context.Context) SpeedtestscheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpeedtestscheduleOutput)
}

// SpeedtestscheduleArrayInput is an input type that accepts SpeedtestscheduleArray and SpeedtestscheduleArrayOutput values.
// You can construct a concrete instance of `SpeedtestscheduleArrayInput` via:
//
//	SpeedtestscheduleArray{ SpeedtestscheduleArgs{...} }
type SpeedtestscheduleArrayInput interface {
	pulumi.Input

	ToSpeedtestscheduleArrayOutput() SpeedtestscheduleArrayOutput
	ToSpeedtestscheduleArrayOutputWithContext(context.Context) SpeedtestscheduleArrayOutput
}

type SpeedtestscheduleArray []SpeedtestscheduleInput

func (SpeedtestscheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Speedtestschedule)(nil)).Elem()
}

func (i SpeedtestscheduleArray) ToSpeedtestscheduleArrayOutput() SpeedtestscheduleArrayOutput {
	return i.ToSpeedtestscheduleArrayOutputWithContext(context.Background())
}

func (i SpeedtestscheduleArray) ToSpeedtestscheduleArrayOutputWithContext(ctx context.Context) SpeedtestscheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpeedtestscheduleArrayOutput)
}

// SpeedtestscheduleMapInput is an input type that accepts SpeedtestscheduleMap and SpeedtestscheduleMapOutput values.
// You can construct a concrete instance of `SpeedtestscheduleMapInput` via:
//
//	SpeedtestscheduleMap{ "key": SpeedtestscheduleArgs{...} }
type SpeedtestscheduleMapInput interface {
	pulumi.Input

	ToSpeedtestscheduleMapOutput() SpeedtestscheduleMapOutput
	ToSpeedtestscheduleMapOutputWithContext(context.Context) SpeedtestscheduleMapOutput
}

type SpeedtestscheduleMap map[string]SpeedtestscheduleInput

func (SpeedtestscheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Speedtestschedule)(nil)).Elem()
}

func (i SpeedtestscheduleMap) ToSpeedtestscheduleMapOutput() SpeedtestscheduleMapOutput {
	return i.ToSpeedtestscheduleMapOutputWithContext(context.Background())
}

func (i SpeedtestscheduleMap) ToSpeedtestscheduleMapOutputWithContext(ctx context.Context) SpeedtestscheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpeedtestscheduleMapOutput)
}

type SpeedtestscheduleOutput struct{ *pulumi.OutputState }

func (SpeedtestscheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Speedtestschedule)(nil)).Elem()
}

func (o SpeedtestscheduleOutput) ToSpeedtestscheduleOutput() SpeedtestscheduleOutput {
	return o
}

func (o SpeedtestscheduleOutput) ToSpeedtestscheduleOutputWithContext(ctx context.Context) SpeedtestscheduleOutput {
	return o
}

// Port of the controller to get access token.
func (o SpeedtestscheduleOutput) CtrlPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.IntOutput { return v.CtrlPort }).(pulumi.IntOutput)
}

// DSCP used for speed test.
func (o SpeedtestscheduleOutput) Diffserv() pulumi.StringOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.StringOutput { return v.Diffserv }).(pulumi.StringOutput)
}

// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
func (o SpeedtestscheduleOutput) DynamicServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.StringOutput { return v.DynamicServer }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SpeedtestscheduleOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o SpeedtestscheduleOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Interface name.
func (o SpeedtestscheduleOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Protocol Auto(default), TCP or UDP used for speed test. Valid values: `UDP`, `TCP`, `Auto`.
func (o SpeedtestscheduleOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Schedules for the interface. The structure of `schedules` block is documented below.
func (o SpeedtestscheduleOutput) Schedules() SpeedtestscheduleScheduleArrayOutput {
	return o.ApplyT(func(v *Speedtestschedule) SpeedtestscheduleScheduleArrayOutput { return v.Schedules }).(SpeedtestscheduleScheduleArrayOutput)
}

// Speed test server name.
func (o SpeedtestscheduleOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

// Port of the server to run speed test.
func (o SpeedtestscheduleOutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.IntOutput { return v.ServerPort }).(pulumi.IntOutput)
}

// Enable/disable scheduled speed test. Valid values: `disable`, `enable`.
func (o SpeedtestscheduleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
func (o SpeedtestscheduleOutput) UpdateInbandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.StringOutput { return v.UpdateInbandwidth }).(pulumi.StringOutput)
}

// Maximum downloading bandwidth (kbps) to be used in a speed test.
func (o SpeedtestscheduleOutput) UpdateInbandwidthMaximum() pulumi.IntOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.IntOutput { return v.UpdateInbandwidthMaximum }).(pulumi.IntOutput)
}

// Minimum downloading bandwidth (kbps) to be considered effective.
func (o SpeedtestscheduleOutput) UpdateInbandwidthMinimum() pulumi.IntOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.IntOutput { return v.UpdateInbandwidthMinimum }).(pulumi.IntOutput)
}

// Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
func (o SpeedtestscheduleOutput) UpdateOutbandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.StringOutput { return v.UpdateOutbandwidth }).(pulumi.StringOutput)
}

// Maximum uploading bandwidth (kbps) to be used in a speed test.
func (o SpeedtestscheduleOutput) UpdateOutbandwidthMaximum() pulumi.IntOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.IntOutput { return v.UpdateOutbandwidthMaximum }).(pulumi.IntOutput)
}

// Minimum uploading bandwidth (kbps) to be considered effective.
func (o SpeedtestscheduleOutput) UpdateOutbandwidthMinimum() pulumi.IntOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.IntOutput { return v.UpdateOutbandwidthMinimum }).(pulumi.IntOutput)
}

// Set egress shaper based on the test result. Valid values: `disable`, `local`, `remote`, `both`.
func (o SpeedtestscheduleOutput) UpdateShaper() pulumi.StringOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.StringOutput { return v.UpdateShaper }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SpeedtestscheduleOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Speedtestschedule) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SpeedtestscheduleArrayOutput struct{ *pulumi.OutputState }

func (SpeedtestscheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Speedtestschedule)(nil)).Elem()
}

func (o SpeedtestscheduleArrayOutput) ToSpeedtestscheduleArrayOutput() SpeedtestscheduleArrayOutput {
	return o
}

func (o SpeedtestscheduleArrayOutput) ToSpeedtestscheduleArrayOutputWithContext(ctx context.Context) SpeedtestscheduleArrayOutput {
	return o
}

func (o SpeedtestscheduleArrayOutput) Index(i pulumi.IntInput) SpeedtestscheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Speedtestschedule {
		return vs[0].([]*Speedtestschedule)[vs[1].(int)]
	}).(SpeedtestscheduleOutput)
}

type SpeedtestscheduleMapOutput struct{ *pulumi.OutputState }

func (SpeedtestscheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Speedtestschedule)(nil)).Elem()
}

func (o SpeedtestscheduleMapOutput) ToSpeedtestscheduleMapOutput() SpeedtestscheduleMapOutput {
	return o
}

func (o SpeedtestscheduleMapOutput) ToSpeedtestscheduleMapOutputWithContext(ctx context.Context) SpeedtestscheduleMapOutput {
	return o
}

func (o SpeedtestscheduleMapOutput) MapIndex(k pulumi.StringInput) SpeedtestscheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Speedtestschedule {
		return vs[0].(map[string]*Speedtestschedule)[vs[1].(string)]
	}).(SpeedtestscheduleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpeedtestscheduleInput)(nil)).Elem(), &Speedtestschedule{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpeedtestscheduleArrayInput)(nil)).Elem(), SpeedtestscheduleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpeedtestscheduleMapInput)(nil)).Elem(), SpeedtestscheduleMap{})
	pulumi.RegisterOutputType(SpeedtestscheduleOutput{})
	pulumi.RegisterOutputType(SpeedtestscheduleArrayOutput{})
	pulumi.RegisterOutputType(SpeedtestscheduleMapOutput{})
}
