// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Speed test schedule for each interface. Applies to FortiOS Version `>= 7.0.0`.
//
// ## Import
//
// # System SpeedTestSchedule can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemSpeedtestschedule:SystemSpeedtestschedule labelname {{interface}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemSpeedtestschedule:SystemSpeedtestschedule labelname {{interface}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemSpeedtestschedule struct {
	pulumi.CustomResourceState

	// DSCP used for speed test.
	Diffserv pulumi.StringOutput `pulumi:"diffserv"`
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer pulumi.StringOutput `pulumi:"dynamicServer"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Interface name.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules SystemSpeedtestscheduleScheduleArrayOutput `pulumi:"schedules"`
	// Speed test server name.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
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
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemSpeedtestschedule registers a new resource with the given unique name, arguments, and options.
func NewSystemSpeedtestschedule(ctx *pulumi.Context,
	name string, args *SystemSpeedtestscheduleArgs, opts ...pulumi.ResourceOption) (*SystemSpeedtestschedule, error) {
	if args == nil {
		args = &SystemSpeedtestscheduleArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSpeedtestschedule
	err := ctx.RegisterResource("fortios:index/systemSpeedtestschedule:SystemSpeedtestschedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSpeedtestschedule gets an existing SystemSpeedtestschedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSpeedtestschedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSpeedtestscheduleState, opts ...pulumi.ResourceOption) (*SystemSpeedtestschedule, error) {
	var resource SystemSpeedtestschedule
	err := ctx.ReadResource("fortios:index/systemSpeedtestschedule:SystemSpeedtestschedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSpeedtestschedule resources.
type systemSpeedtestscheduleState struct {
	// DSCP used for speed test.
	Diffserv *string `pulumi:"diffserv"`
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer *string `pulumi:"dynamicServer"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules []SystemSpeedtestscheduleSchedule `pulumi:"schedules"`
	// Speed test server name.
	ServerName *string `pulumi:"serverName"`
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
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemSpeedtestscheduleState struct {
	// DSCP used for speed test.
	Diffserv pulumi.StringPtrInput
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules SystemSpeedtestscheduleScheduleArrayInput
	// Speed test server name.
	ServerName pulumi.StringPtrInput
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
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSpeedtestscheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSpeedtestscheduleState)(nil)).Elem()
}

type systemSpeedtestscheduleArgs struct {
	// DSCP used for speed test.
	Diffserv *string `pulumi:"diffserv"`
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer *string `pulumi:"dynamicServer"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules []SystemSpeedtestscheduleSchedule `pulumi:"schedules"`
	// Speed test server name.
	ServerName *string `pulumi:"serverName"`
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
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemSpeedtestschedule resource.
type SystemSpeedtestscheduleArgs struct {
	// DSCP used for speed test.
	Diffserv pulumi.StringPtrInput
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules SystemSpeedtestscheduleScheduleArrayInput
	// Speed test server name.
	ServerName pulumi.StringPtrInput
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
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSpeedtestscheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSpeedtestscheduleArgs)(nil)).Elem()
}

type SystemSpeedtestscheduleInput interface {
	pulumi.Input

	ToSystemSpeedtestscheduleOutput() SystemSpeedtestscheduleOutput
	ToSystemSpeedtestscheduleOutputWithContext(ctx context.Context) SystemSpeedtestscheduleOutput
}

func (*SystemSpeedtestschedule) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSpeedtestschedule)(nil)).Elem()
}

func (i *SystemSpeedtestschedule) ToSystemSpeedtestscheduleOutput() SystemSpeedtestscheduleOutput {
	return i.ToSystemSpeedtestscheduleOutputWithContext(context.Background())
}

func (i *SystemSpeedtestschedule) ToSystemSpeedtestscheduleOutputWithContext(ctx context.Context) SystemSpeedtestscheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSpeedtestscheduleOutput)
}

// SystemSpeedtestscheduleArrayInput is an input type that accepts SystemSpeedtestscheduleArray and SystemSpeedtestscheduleArrayOutput values.
// You can construct a concrete instance of `SystemSpeedtestscheduleArrayInput` via:
//
//	SystemSpeedtestscheduleArray{ SystemSpeedtestscheduleArgs{...} }
type SystemSpeedtestscheduleArrayInput interface {
	pulumi.Input

	ToSystemSpeedtestscheduleArrayOutput() SystemSpeedtestscheduleArrayOutput
	ToSystemSpeedtestscheduleArrayOutputWithContext(context.Context) SystemSpeedtestscheduleArrayOutput
}

type SystemSpeedtestscheduleArray []SystemSpeedtestscheduleInput

func (SystemSpeedtestscheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSpeedtestschedule)(nil)).Elem()
}

func (i SystemSpeedtestscheduleArray) ToSystemSpeedtestscheduleArrayOutput() SystemSpeedtestscheduleArrayOutput {
	return i.ToSystemSpeedtestscheduleArrayOutputWithContext(context.Background())
}

func (i SystemSpeedtestscheduleArray) ToSystemSpeedtestscheduleArrayOutputWithContext(ctx context.Context) SystemSpeedtestscheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSpeedtestscheduleArrayOutput)
}

// SystemSpeedtestscheduleMapInput is an input type that accepts SystemSpeedtestscheduleMap and SystemSpeedtestscheduleMapOutput values.
// You can construct a concrete instance of `SystemSpeedtestscheduleMapInput` via:
//
//	SystemSpeedtestscheduleMap{ "key": SystemSpeedtestscheduleArgs{...} }
type SystemSpeedtestscheduleMapInput interface {
	pulumi.Input

	ToSystemSpeedtestscheduleMapOutput() SystemSpeedtestscheduleMapOutput
	ToSystemSpeedtestscheduleMapOutputWithContext(context.Context) SystemSpeedtestscheduleMapOutput
}

type SystemSpeedtestscheduleMap map[string]SystemSpeedtestscheduleInput

func (SystemSpeedtestscheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSpeedtestschedule)(nil)).Elem()
}

func (i SystemSpeedtestscheduleMap) ToSystemSpeedtestscheduleMapOutput() SystemSpeedtestscheduleMapOutput {
	return i.ToSystemSpeedtestscheduleMapOutputWithContext(context.Background())
}

func (i SystemSpeedtestscheduleMap) ToSystemSpeedtestscheduleMapOutputWithContext(ctx context.Context) SystemSpeedtestscheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSpeedtestscheduleMapOutput)
}

type SystemSpeedtestscheduleOutput struct{ *pulumi.OutputState }

func (SystemSpeedtestscheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSpeedtestschedule)(nil)).Elem()
}

func (o SystemSpeedtestscheduleOutput) ToSystemSpeedtestscheduleOutput() SystemSpeedtestscheduleOutput {
	return o
}

func (o SystemSpeedtestscheduleOutput) ToSystemSpeedtestscheduleOutputWithContext(ctx context.Context) SystemSpeedtestscheduleOutput {
	return o
}

// DSCP used for speed test.
func (o SystemSpeedtestscheduleOutput) Diffserv() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) pulumi.StringOutput { return v.Diffserv }).(pulumi.StringOutput)
}

// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
func (o SystemSpeedtestscheduleOutput) DynamicServer() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) pulumi.StringOutput { return v.DynamicServer }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SystemSpeedtestscheduleOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Interface name.
func (o SystemSpeedtestscheduleOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Schedules for the interface. The structure of `schedules` block is documented below.
func (o SystemSpeedtestscheduleOutput) Schedules() SystemSpeedtestscheduleScheduleArrayOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) SystemSpeedtestscheduleScheduleArrayOutput { return v.Schedules }).(SystemSpeedtestscheduleScheduleArrayOutput)
}

// Speed test server name.
func (o SystemSpeedtestscheduleOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

// Enable/disable scheduled speed test. Valid values: `disable`, `enable`.
func (o SystemSpeedtestscheduleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
func (o SystemSpeedtestscheduleOutput) UpdateInbandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) pulumi.StringOutput { return v.UpdateInbandwidth }).(pulumi.StringOutput)
}

// Maximum downloading bandwidth (kbps) to be used in a speed test.
func (o SystemSpeedtestscheduleOutput) UpdateInbandwidthMaximum() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) pulumi.IntOutput { return v.UpdateInbandwidthMaximum }).(pulumi.IntOutput)
}

// Minimum downloading bandwidth (kbps) to be considered effective.
func (o SystemSpeedtestscheduleOutput) UpdateInbandwidthMinimum() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) pulumi.IntOutput { return v.UpdateInbandwidthMinimum }).(pulumi.IntOutput)
}

// Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
func (o SystemSpeedtestscheduleOutput) UpdateOutbandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) pulumi.StringOutput { return v.UpdateOutbandwidth }).(pulumi.StringOutput)
}

// Maximum uploading bandwidth (kbps) to be used in a speed test.
func (o SystemSpeedtestscheduleOutput) UpdateOutbandwidthMaximum() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) pulumi.IntOutput { return v.UpdateOutbandwidthMaximum }).(pulumi.IntOutput)
}

// Minimum uploading bandwidth (kbps) to be considered effective.
func (o SystemSpeedtestscheduleOutput) UpdateOutbandwidthMinimum() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) pulumi.IntOutput { return v.UpdateOutbandwidthMinimum }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemSpeedtestscheduleOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemSpeedtestschedule) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemSpeedtestscheduleArrayOutput struct{ *pulumi.OutputState }

func (SystemSpeedtestscheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSpeedtestschedule)(nil)).Elem()
}

func (o SystemSpeedtestscheduleArrayOutput) ToSystemSpeedtestscheduleArrayOutput() SystemSpeedtestscheduleArrayOutput {
	return o
}

func (o SystemSpeedtestscheduleArrayOutput) ToSystemSpeedtestscheduleArrayOutputWithContext(ctx context.Context) SystemSpeedtestscheduleArrayOutput {
	return o
}

func (o SystemSpeedtestscheduleArrayOutput) Index(i pulumi.IntInput) SystemSpeedtestscheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSpeedtestschedule {
		return vs[0].([]*SystemSpeedtestschedule)[vs[1].(int)]
	}).(SystemSpeedtestscheduleOutput)
}

type SystemSpeedtestscheduleMapOutput struct{ *pulumi.OutputState }

func (SystemSpeedtestscheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSpeedtestschedule)(nil)).Elem()
}

func (o SystemSpeedtestscheduleMapOutput) ToSystemSpeedtestscheduleMapOutput() SystemSpeedtestscheduleMapOutput {
	return o
}

func (o SystemSpeedtestscheduleMapOutput) ToSystemSpeedtestscheduleMapOutputWithContext(ctx context.Context) SystemSpeedtestscheduleMapOutput {
	return o
}

func (o SystemSpeedtestscheduleMapOutput) MapIndex(k pulumi.StringInput) SystemSpeedtestscheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSpeedtestschedule {
		return vs[0].(map[string]*SystemSpeedtestschedule)[vs[1].(string)]
	}).(SystemSpeedtestscheduleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSpeedtestscheduleInput)(nil)).Elem(), &SystemSpeedtestschedule{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSpeedtestscheduleArrayInput)(nil)).Elem(), SystemSpeedtestscheduleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSpeedtestscheduleMapInput)(nil)).Elem(), SystemSpeedtestscheduleMap{})
	pulumi.RegisterOutputType(SystemSpeedtestscheduleOutput{})
	pulumi.RegisterOutputType(SystemSpeedtestscheduleArrayOutput{})
	pulumi.RegisterOutputType(SystemSpeedtestscheduleMapOutput{})
}
