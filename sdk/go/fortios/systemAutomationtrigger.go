// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Trigger for automation stitches.
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
//			_, err := fortios.NewSystemAutomationtrigger(ctx, "trname", &fortios.SystemAutomationtriggerArgs{
//				EventType:        pulumi.String("event-log"),
//				IocLevel:         pulumi.String("high"),
//				LicenseType:      pulumi.String("forticare-support"),
//				Logid:            pulumi.Int(32002),
//				TriggerFrequency: pulumi.String("daily"),
//				TriggerMinute:    pulumi.Int(60),
//				TriggerType:      pulumi.String("event-based"),
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
// # System AutomationTrigger can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemAutomationtrigger:SystemAutomationtrigger labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemAutomationtrigger:SystemAutomationtrigger labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemAutomationtrigger struct {
	pulumi.CustomResourceState

	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Event type.
	EventType pulumi.StringOutput `pulumi:"eventType"`
	// Fabric connector event handler name.
	FabricEventName pulumi.StringPtrOutput `pulumi:"fabricEventName"`
	// Fabric connector event severity.
	FabricEventSeverity pulumi.StringPtrOutput `pulumi:"fabricEventSeverity"`
	// FortiAnalyzer event handler name.
	FazEventName pulumi.StringPtrOutput `pulumi:"fazEventName"`
	// FortiAnalyzer event severity.
	FazEventSeverity pulumi.StringPtrOutput `pulumi:"fazEventSeverity"`
	// FortiAnalyzer event tags.
	FazEventTags pulumi.StringPtrOutput `pulumi:"fazEventTags"`
	// Customized trigger field settings. The structure of `fields` block is documented below.
	Fields SystemAutomationtriggerFieldArrayOutput `pulumi:"fields"`
	// IOC threat level. Valid values: `medium`, `high`.
	IocLevel pulumi.StringOutput `pulumi:"iocLevel"`
	// License type.
	LicenseType pulumi.StringOutput `pulumi:"licenseType"`
	// Log ID to trigger event.
	Logid pulumi.IntOutput `pulumi:"logid"`
	// Log ID to trigger event. Only applies on FortiOS v7.0.0+. The structure of `logidBlock` block is documented below.
	LogidBlocks SystemAutomationtriggerLogidBlockArrayOutput `pulumi:"logidBlocks"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Security Rating report.
	ReportType pulumi.StringOutput `pulumi:"reportType"`
	// Fabric connector serial number.
	Serial pulumi.StringPtrOutput `pulumi:"serial"`
	// Trigger date and time (YYYY-MM-DD HH:MM:SS).
	TriggerDatetime pulumi.StringOutput `pulumi:"triggerDatetime"`
	// Day within a month to trigger.
	TriggerDay pulumi.IntOutput `pulumi:"triggerDay"`
	// Scheduled trigger frequency (default = daily).
	TriggerFrequency pulumi.StringOutput `pulumi:"triggerFrequency"`
	// Hour of the day on which to trigger (0 - 23, default = 1).
	TriggerHour pulumi.IntOutput `pulumi:"triggerHour"`
	// Minute of the hour on which to trigger (0 - 59, 60 to randomize).
	TriggerMinute pulumi.IntOutput `pulumi:"triggerMinute"`
	// Trigger type. Valid values: `event-based`, `scheduled`.
	TriggerType pulumi.StringOutput `pulumi:"triggerType"`
	// Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	TriggerWeekday pulumi.StringOutput `pulumi:"triggerWeekday"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
	Vdoms SystemAutomationtriggerVdomArrayOutput `pulumi:"vdoms"`
}

// NewSystemAutomationtrigger registers a new resource with the given unique name, arguments, and options.
func NewSystemAutomationtrigger(ctx *pulumi.Context,
	name string, args *SystemAutomationtriggerArgs, opts ...pulumi.ResourceOption) (*SystemAutomationtrigger, error) {
	if args == nil {
		args = &SystemAutomationtriggerArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemAutomationtrigger
	err := ctx.RegisterResource("fortios:index/systemAutomationtrigger:SystemAutomationtrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAutomationtrigger gets an existing SystemAutomationtrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAutomationtrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAutomationtriggerState, opts ...pulumi.ResourceOption) (*SystemAutomationtrigger, error) {
	var resource SystemAutomationtrigger
	err := ctx.ReadResource("fortios:index/systemAutomationtrigger:SystemAutomationtrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAutomationtrigger resources.
type systemAutomationtriggerState struct {
	// Description.
	Description *string `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Event type.
	EventType *string `pulumi:"eventType"`
	// Fabric connector event handler name.
	FabricEventName *string `pulumi:"fabricEventName"`
	// Fabric connector event severity.
	FabricEventSeverity *string `pulumi:"fabricEventSeverity"`
	// FortiAnalyzer event handler name.
	FazEventName *string `pulumi:"fazEventName"`
	// FortiAnalyzer event severity.
	FazEventSeverity *string `pulumi:"fazEventSeverity"`
	// FortiAnalyzer event tags.
	FazEventTags *string `pulumi:"fazEventTags"`
	// Customized trigger field settings. The structure of `fields` block is documented below.
	Fields []SystemAutomationtriggerField `pulumi:"fields"`
	// IOC threat level. Valid values: `medium`, `high`.
	IocLevel *string `pulumi:"iocLevel"`
	// License type.
	LicenseType *string `pulumi:"licenseType"`
	// Log ID to trigger event.
	Logid *int `pulumi:"logid"`
	// Log ID to trigger event. Only applies on FortiOS v7.0.0+. The structure of `logidBlock` block is documented below.
	LogidBlocks []SystemAutomationtriggerLogidBlock `pulumi:"logidBlocks"`
	// Name.
	Name *string `pulumi:"name"`
	// Security Rating report.
	ReportType *string `pulumi:"reportType"`
	// Fabric connector serial number.
	Serial *string `pulumi:"serial"`
	// Trigger date and time (YYYY-MM-DD HH:MM:SS).
	TriggerDatetime *string `pulumi:"triggerDatetime"`
	// Day within a month to trigger.
	TriggerDay *int `pulumi:"triggerDay"`
	// Scheduled trigger frequency (default = daily).
	TriggerFrequency *string `pulumi:"triggerFrequency"`
	// Hour of the day on which to trigger (0 - 23, default = 1).
	TriggerHour *int `pulumi:"triggerHour"`
	// Minute of the hour on which to trigger (0 - 59, 60 to randomize).
	TriggerMinute *int `pulumi:"triggerMinute"`
	// Trigger type. Valid values: `event-based`, `scheduled`.
	TriggerType *string `pulumi:"triggerType"`
	// Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	TriggerWeekday *string `pulumi:"triggerWeekday"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
	Vdoms []SystemAutomationtriggerVdom `pulumi:"vdoms"`
}

type SystemAutomationtriggerState struct {
	// Description.
	Description pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Event type.
	EventType pulumi.StringPtrInput
	// Fabric connector event handler name.
	FabricEventName pulumi.StringPtrInput
	// Fabric connector event severity.
	FabricEventSeverity pulumi.StringPtrInput
	// FortiAnalyzer event handler name.
	FazEventName pulumi.StringPtrInput
	// FortiAnalyzer event severity.
	FazEventSeverity pulumi.StringPtrInput
	// FortiAnalyzer event tags.
	FazEventTags pulumi.StringPtrInput
	// Customized trigger field settings. The structure of `fields` block is documented below.
	Fields SystemAutomationtriggerFieldArrayInput
	// IOC threat level. Valid values: `medium`, `high`.
	IocLevel pulumi.StringPtrInput
	// License type.
	LicenseType pulumi.StringPtrInput
	// Log ID to trigger event.
	Logid pulumi.IntPtrInput
	// Log ID to trigger event. Only applies on FortiOS v7.0.0+. The structure of `logidBlock` block is documented below.
	LogidBlocks SystemAutomationtriggerLogidBlockArrayInput
	// Name.
	Name pulumi.StringPtrInput
	// Security Rating report.
	ReportType pulumi.StringPtrInput
	// Fabric connector serial number.
	Serial pulumi.StringPtrInput
	// Trigger date and time (YYYY-MM-DD HH:MM:SS).
	TriggerDatetime pulumi.StringPtrInput
	// Day within a month to trigger.
	TriggerDay pulumi.IntPtrInput
	// Scheduled trigger frequency (default = daily).
	TriggerFrequency pulumi.StringPtrInput
	// Hour of the day on which to trigger (0 - 23, default = 1).
	TriggerHour pulumi.IntPtrInput
	// Minute of the hour on which to trigger (0 - 59, 60 to randomize).
	TriggerMinute pulumi.IntPtrInput
	// Trigger type. Valid values: `event-based`, `scheduled`.
	TriggerType pulumi.StringPtrInput
	// Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	TriggerWeekday pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
	Vdoms SystemAutomationtriggerVdomArrayInput
}

func (SystemAutomationtriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutomationtriggerState)(nil)).Elem()
}

type systemAutomationtriggerArgs struct {
	// Description.
	Description *string `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Event type.
	EventType *string `pulumi:"eventType"`
	// Fabric connector event handler name.
	FabricEventName *string `pulumi:"fabricEventName"`
	// Fabric connector event severity.
	FabricEventSeverity *string `pulumi:"fabricEventSeverity"`
	// FortiAnalyzer event handler name.
	FazEventName *string `pulumi:"fazEventName"`
	// FortiAnalyzer event severity.
	FazEventSeverity *string `pulumi:"fazEventSeverity"`
	// FortiAnalyzer event tags.
	FazEventTags *string `pulumi:"fazEventTags"`
	// Customized trigger field settings. The structure of `fields` block is documented below.
	Fields []SystemAutomationtriggerField `pulumi:"fields"`
	// IOC threat level. Valid values: `medium`, `high`.
	IocLevel *string `pulumi:"iocLevel"`
	// License type.
	LicenseType *string `pulumi:"licenseType"`
	// Log ID to trigger event.
	Logid *int `pulumi:"logid"`
	// Log ID to trigger event. Only applies on FortiOS v7.0.0+. The structure of `logidBlock` block is documented below.
	LogidBlocks []SystemAutomationtriggerLogidBlock `pulumi:"logidBlocks"`
	// Name.
	Name *string `pulumi:"name"`
	// Security Rating report.
	ReportType *string `pulumi:"reportType"`
	// Fabric connector serial number.
	Serial *string `pulumi:"serial"`
	// Trigger date and time (YYYY-MM-DD HH:MM:SS).
	TriggerDatetime *string `pulumi:"triggerDatetime"`
	// Day within a month to trigger.
	TriggerDay *int `pulumi:"triggerDay"`
	// Scheduled trigger frequency (default = daily).
	TriggerFrequency *string `pulumi:"triggerFrequency"`
	// Hour of the day on which to trigger (0 - 23, default = 1).
	TriggerHour *int `pulumi:"triggerHour"`
	// Minute of the hour on which to trigger (0 - 59, 60 to randomize).
	TriggerMinute *int `pulumi:"triggerMinute"`
	// Trigger type. Valid values: `event-based`, `scheduled`.
	TriggerType *string `pulumi:"triggerType"`
	// Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	TriggerWeekday *string `pulumi:"triggerWeekday"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
	Vdoms []SystemAutomationtriggerVdom `pulumi:"vdoms"`
}

// The set of arguments for constructing a SystemAutomationtrigger resource.
type SystemAutomationtriggerArgs struct {
	// Description.
	Description pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Event type.
	EventType pulumi.StringPtrInput
	// Fabric connector event handler name.
	FabricEventName pulumi.StringPtrInput
	// Fabric connector event severity.
	FabricEventSeverity pulumi.StringPtrInput
	// FortiAnalyzer event handler name.
	FazEventName pulumi.StringPtrInput
	// FortiAnalyzer event severity.
	FazEventSeverity pulumi.StringPtrInput
	// FortiAnalyzer event tags.
	FazEventTags pulumi.StringPtrInput
	// Customized trigger field settings. The structure of `fields` block is documented below.
	Fields SystemAutomationtriggerFieldArrayInput
	// IOC threat level. Valid values: `medium`, `high`.
	IocLevel pulumi.StringPtrInput
	// License type.
	LicenseType pulumi.StringPtrInput
	// Log ID to trigger event.
	Logid pulumi.IntPtrInput
	// Log ID to trigger event. Only applies on FortiOS v7.0.0+. The structure of `logidBlock` block is documented below.
	LogidBlocks SystemAutomationtriggerLogidBlockArrayInput
	// Name.
	Name pulumi.StringPtrInput
	// Security Rating report.
	ReportType pulumi.StringPtrInput
	// Fabric connector serial number.
	Serial pulumi.StringPtrInput
	// Trigger date and time (YYYY-MM-DD HH:MM:SS).
	TriggerDatetime pulumi.StringPtrInput
	// Day within a month to trigger.
	TriggerDay pulumi.IntPtrInput
	// Scheduled trigger frequency (default = daily).
	TriggerFrequency pulumi.StringPtrInput
	// Hour of the day on which to trigger (0 - 23, default = 1).
	TriggerHour pulumi.IntPtrInput
	// Minute of the hour on which to trigger (0 - 59, 60 to randomize).
	TriggerMinute pulumi.IntPtrInput
	// Trigger type. Valid values: `event-based`, `scheduled`.
	TriggerType pulumi.StringPtrInput
	// Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	TriggerWeekday pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
	Vdoms SystemAutomationtriggerVdomArrayInput
}

func (SystemAutomationtriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutomationtriggerArgs)(nil)).Elem()
}

type SystemAutomationtriggerInput interface {
	pulumi.Input

	ToSystemAutomationtriggerOutput() SystemAutomationtriggerOutput
	ToSystemAutomationtriggerOutputWithContext(ctx context.Context) SystemAutomationtriggerOutput
}

func (*SystemAutomationtrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutomationtrigger)(nil)).Elem()
}

func (i *SystemAutomationtrigger) ToSystemAutomationtriggerOutput() SystemAutomationtriggerOutput {
	return i.ToSystemAutomationtriggerOutputWithContext(context.Background())
}

func (i *SystemAutomationtrigger) ToSystemAutomationtriggerOutputWithContext(ctx context.Context) SystemAutomationtriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutomationtriggerOutput)
}

// SystemAutomationtriggerArrayInput is an input type that accepts SystemAutomationtriggerArray and SystemAutomationtriggerArrayOutput values.
// You can construct a concrete instance of `SystemAutomationtriggerArrayInput` via:
//
//	SystemAutomationtriggerArray{ SystemAutomationtriggerArgs{...} }
type SystemAutomationtriggerArrayInput interface {
	pulumi.Input

	ToSystemAutomationtriggerArrayOutput() SystemAutomationtriggerArrayOutput
	ToSystemAutomationtriggerArrayOutputWithContext(context.Context) SystemAutomationtriggerArrayOutput
}

type SystemAutomationtriggerArray []SystemAutomationtriggerInput

func (SystemAutomationtriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAutomationtrigger)(nil)).Elem()
}

func (i SystemAutomationtriggerArray) ToSystemAutomationtriggerArrayOutput() SystemAutomationtriggerArrayOutput {
	return i.ToSystemAutomationtriggerArrayOutputWithContext(context.Background())
}

func (i SystemAutomationtriggerArray) ToSystemAutomationtriggerArrayOutputWithContext(ctx context.Context) SystemAutomationtriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutomationtriggerArrayOutput)
}

// SystemAutomationtriggerMapInput is an input type that accepts SystemAutomationtriggerMap and SystemAutomationtriggerMapOutput values.
// You can construct a concrete instance of `SystemAutomationtriggerMapInput` via:
//
//	SystemAutomationtriggerMap{ "key": SystemAutomationtriggerArgs{...} }
type SystemAutomationtriggerMapInput interface {
	pulumi.Input

	ToSystemAutomationtriggerMapOutput() SystemAutomationtriggerMapOutput
	ToSystemAutomationtriggerMapOutputWithContext(context.Context) SystemAutomationtriggerMapOutput
}

type SystemAutomationtriggerMap map[string]SystemAutomationtriggerInput

func (SystemAutomationtriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAutomationtrigger)(nil)).Elem()
}

func (i SystemAutomationtriggerMap) ToSystemAutomationtriggerMapOutput() SystemAutomationtriggerMapOutput {
	return i.ToSystemAutomationtriggerMapOutputWithContext(context.Background())
}

func (i SystemAutomationtriggerMap) ToSystemAutomationtriggerMapOutputWithContext(ctx context.Context) SystemAutomationtriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutomationtriggerMapOutput)
}

type SystemAutomationtriggerOutput struct{ *pulumi.OutputState }

func (SystemAutomationtriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutomationtrigger)(nil)).Elem()
}

func (o SystemAutomationtriggerOutput) ToSystemAutomationtriggerOutput() SystemAutomationtriggerOutput {
	return o
}

func (o SystemAutomationtriggerOutput) ToSystemAutomationtriggerOutputWithContext(ctx context.Context) SystemAutomationtriggerOutput {
	return o
}

// Description.
func (o SystemAutomationtriggerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SystemAutomationtriggerOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Event type.
func (o SystemAutomationtriggerOutput) EventType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringOutput { return v.EventType }).(pulumi.StringOutput)
}

// Fabric connector event handler name.
func (o SystemAutomationtriggerOutput) FabricEventName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringPtrOutput { return v.FabricEventName }).(pulumi.StringPtrOutput)
}

// Fabric connector event severity.
func (o SystemAutomationtriggerOutput) FabricEventSeverity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringPtrOutput { return v.FabricEventSeverity }).(pulumi.StringPtrOutput)
}

// FortiAnalyzer event handler name.
func (o SystemAutomationtriggerOutput) FazEventName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringPtrOutput { return v.FazEventName }).(pulumi.StringPtrOutput)
}

// FortiAnalyzer event severity.
func (o SystemAutomationtriggerOutput) FazEventSeverity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringPtrOutput { return v.FazEventSeverity }).(pulumi.StringPtrOutput)
}

// FortiAnalyzer event tags.
func (o SystemAutomationtriggerOutput) FazEventTags() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringPtrOutput { return v.FazEventTags }).(pulumi.StringPtrOutput)
}

// Customized trigger field settings. The structure of `fields` block is documented below.
func (o SystemAutomationtriggerOutput) Fields() SystemAutomationtriggerFieldArrayOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) SystemAutomationtriggerFieldArrayOutput { return v.Fields }).(SystemAutomationtriggerFieldArrayOutput)
}

// IOC threat level. Valid values: `medium`, `high`.
func (o SystemAutomationtriggerOutput) IocLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringOutput { return v.IocLevel }).(pulumi.StringOutput)
}

// License type.
func (o SystemAutomationtriggerOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringOutput { return v.LicenseType }).(pulumi.StringOutput)
}

// Log ID to trigger event.
func (o SystemAutomationtriggerOutput) Logid() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.IntOutput { return v.Logid }).(pulumi.IntOutput)
}

// Log ID to trigger event. Only applies on FortiOS v7.0.0+. The structure of `logidBlock` block is documented below.
func (o SystemAutomationtriggerOutput) LogidBlocks() SystemAutomationtriggerLogidBlockArrayOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) SystemAutomationtriggerLogidBlockArrayOutput { return v.LogidBlocks }).(SystemAutomationtriggerLogidBlockArrayOutput)
}

// Name.
func (o SystemAutomationtriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Security Rating report.
func (o SystemAutomationtriggerOutput) ReportType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringOutput { return v.ReportType }).(pulumi.StringOutput)
}

// Fabric connector serial number.
func (o SystemAutomationtriggerOutput) Serial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringPtrOutput { return v.Serial }).(pulumi.StringPtrOutput)
}

// Trigger date and time (YYYY-MM-DD HH:MM:SS).
func (o SystemAutomationtriggerOutput) TriggerDatetime() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringOutput { return v.TriggerDatetime }).(pulumi.StringOutput)
}

// Day within a month to trigger.
func (o SystemAutomationtriggerOutput) TriggerDay() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.IntOutput { return v.TriggerDay }).(pulumi.IntOutput)
}

// Scheduled trigger frequency (default = daily).
func (o SystemAutomationtriggerOutput) TriggerFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringOutput { return v.TriggerFrequency }).(pulumi.StringOutput)
}

// Hour of the day on which to trigger (0 - 23, default = 1).
func (o SystemAutomationtriggerOutput) TriggerHour() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.IntOutput { return v.TriggerHour }).(pulumi.IntOutput)
}

// Minute of the hour on which to trigger (0 - 59, 60 to randomize).
func (o SystemAutomationtriggerOutput) TriggerMinute() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.IntOutput { return v.TriggerMinute }).(pulumi.IntOutput)
}

// Trigger type. Valid values: `event-based`, `scheduled`.
func (o SystemAutomationtriggerOutput) TriggerType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringOutput { return v.TriggerType }).(pulumi.StringOutput)
}

// Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
func (o SystemAutomationtriggerOutput) TriggerWeekday() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringOutput { return v.TriggerWeekday }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemAutomationtriggerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
func (o SystemAutomationtriggerOutput) Vdoms() SystemAutomationtriggerVdomArrayOutput {
	return o.ApplyT(func(v *SystemAutomationtrigger) SystemAutomationtriggerVdomArrayOutput { return v.Vdoms }).(SystemAutomationtriggerVdomArrayOutput)
}

type SystemAutomationtriggerArrayOutput struct{ *pulumi.OutputState }

func (SystemAutomationtriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAutomationtrigger)(nil)).Elem()
}

func (o SystemAutomationtriggerArrayOutput) ToSystemAutomationtriggerArrayOutput() SystemAutomationtriggerArrayOutput {
	return o
}

func (o SystemAutomationtriggerArrayOutput) ToSystemAutomationtriggerArrayOutputWithContext(ctx context.Context) SystemAutomationtriggerArrayOutput {
	return o
}

func (o SystemAutomationtriggerArrayOutput) Index(i pulumi.IntInput) SystemAutomationtriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAutomationtrigger {
		return vs[0].([]*SystemAutomationtrigger)[vs[1].(int)]
	}).(SystemAutomationtriggerOutput)
}

type SystemAutomationtriggerMapOutput struct{ *pulumi.OutputState }

func (SystemAutomationtriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAutomationtrigger)(nil)).Elem()
}

func (o SystemAutomationtriggerMapOutput) ToSystemAutomationtriggerMapOutput() SystemAutomationtriggerMapOutput {
	return o
}

func (o SystemAutomationtriggerMapOutput) ToSystemAutomationtriggerMapOutputWithContext(ctx context.Context) SystemAutomationtriggerMapOutput {
	return o
}

func (o SystemAutomationtriggerMapOutput) MapIndex(k pulumi.StringInput) SystemAutomationtriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAutomationtrigger {
		return vs[0].(map[string]*SystemAutomationtrigger)[vs[1].(string)]
	}).(SystemAutomationtriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutomationtriggerInput)(nil)).Elem(), &SystemAutomationtrigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutomationtriggerArrayInput)(nil)).Elem(), SystemAutomationtriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutomationtriggerMapInput)(nil)).Elem(), SystemAutomationtriggerMap{})
	pulumi.RegisterOutputType(SystemAutomationtriggerOutput{})
	pulumi.RegisterOutputType(SystemAutomationtriggerArrayOutput{})
	pulumi.RegisterOutputType(SystemAutomationtriggerMapOutput{})
}
