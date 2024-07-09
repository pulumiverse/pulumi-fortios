// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewAutomationtrigger(ctx, "trname", &system.AutomationtriggerArgs{
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
// System AutomationTrigger can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/automationtrigger:Automationtrigger labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/automationtrigger:Automationtrigger labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Automationtrigger struct {
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
	Fields AutomationtriggerFieldArrayOutput `pulumi:"fields"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// IOC threat level. Valid values: `medium`, `high`.
	IocLevel pulumi.StringOutput `pulumi:"iocLevel"`
	// License type.
	LicenseType pulumi.StringOutput `pulumi:"licenseType"`
	// Log ID to trigger event. *Due to the data type change of API, for other versions of FortiOS, please check variable `logidBlock`.*
	Logid pulumi.IntOutput `pulumi:"logid"`
	// Log IDs to trigger event. *Due to the data type change of API, for other versions of FortiOS, please check variable `logid`.* The structure of `logidBlock` block is documented below.
	LogidBlocks AutomationtriggerLogidBlockArrayOutput `pulumi:"logidBlocks"`
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
	// Minute of the hour on which to trigger (0 - 59, default = 0).
	TriggerMinute pulumi.IntOutput `pulumi:"triggerMinute"`
	// Trigger type. Valid values: `event-based`, `scheduled`.
	TriggerType pulumi.StringOutput `pulumi:"triggerType"`
	// Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	TriggerWeekday pulumi.StringOutput `pulumi:"triggerWeekday"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
	// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
	Vdoms AutomationtriggerVdomArrayOutput `pulumi:"vdoms"`
}

// NewAutomationtrigger registers a new resource with the given unique name, arguments, and options.
func NewAutomationtrigger(ctx *pulumi.Context,
	name string, args *AutomationtriggerArgs, opts ...pulumi.ResourceOption) (*Automationtrigger, error) {
	if args == nil {
		args = &AutomationtriggerArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Automationtrigger
	err := ctx.RegisterResource("fortios:system/automationtrigger:Automationtrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutomationtrigger gets an existing Automationtrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutomationtrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationtriggerState, opts ...pulumi.ResourceOption) (*Automationtrigger, error) {
	var resource Automationtrigger
	err := ctx.ReadResource("fortios:system/automationtrigger:Automationtrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Automationtrigger resources.
type automationtriggerState struct {
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
	Fields []AutomationtriggerField `pulumi:"fields"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// IOC threat level. Valid values: `medium`, `high`.
	IocLevel *string `pulumi:"iocLevel"`
	// License type.
	LicenseType *string `pulumi:"licenseType"`
	// Log ID to trigger event. *Due to the data type change of API, for other versions of FortiOS, please check variable `logidBlock`.*
	Logid *int `pulumi:"logid"`
	// Log IDs to trigger event. *Due to the data type change of API, for other versions of FortiOS, please check variable `logid`.* The structure of `logidBlock` block is documented below.
	LogidBlocks []AutomationtriggerLogidBlock `pulumi:"logidBlocks"`
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
	// Minute of the hour on which to trigger (0 - 59, default = 0).
	TriggerMinute *int `pulumi:"triggerMinute"`
	// Trigger type. Valid values: `event-based`, `scheduled`.
	TriggerType *string `pulumi:"triggerType"`
	// Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	TriggerWeekday *string `pulumi:"triggerWeekday"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
	Vdoms []AutomationtriggerVdom `pulumi:"vdoms"`
}

type AutomationtriggerState struct {
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
	Fields AutomationtriggerFieldArrayInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// IOC threat level. Valid values: `medium`, `high`.
	IocLevel pulumi.StringPtrInput
	// License type.
	LicenseType pulumi.StringPtrInput
	// Log ID to trigger event. *Due to the data type change of API, for other versions of FortiOS, please check variable `logidBlock`.*
	Logid pulumi.IntPtrInput
	// Log IDs to trigger event. *Due to the data type change of API, for other versions of FortiOS, please check variable `logid`.* The structure of `logidBlock` block is documented below.
	LogidBlocks AutomationtriggerLogidBlockArrayInput
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
	// Minute of the hour on which to trigger (0 - 59, default = 0).
	TriggerMinute pulumi.IntPtrInput
	// Trigger type. Valid values: `event-based`, `scheduled`.
	TriggerType pulumi.StringPtrInput
	// Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	TriggerWeekday pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
	Vdoms AutomationtriggerVdomArrayInput
}

func (AutomationtriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationtriggerState)(nil)).Elem()
}

type automationtriggerArgs struct {
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
	Fields []AutomationtriggerField `pulumi:"fields"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// IOC threat level. Valid values: `medium`, `high`.
	IocLevel *string `pulumi:"iocLevel"`
	// License type.
	LicenseType *string `pulumi:"licenseType"`
	// Log ID to trigger event. *Due to the data type change of API, for other versions of FortiOS, please check variable `logidBlock`.*
	Logid *int `pulumi:"logid"`
	// Log IDs to trigger event. *Due to the data type change of API, for other versions of FortiOS, please check variable `logid`.* The structure of `logidBlock` block is documented below.
	LogidBlocks []AutomationtriggerLogidBlock `pulumi:"logidBlocks"`
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
	// Minute of the hour on which to trigger (0 - 59, default = 0).
	TriggerMinute *int `pulumi:"triggerMinute"`
	// Trigger type. Valid values: `event-based`, `scheduled`.
	TriggerType *string `pulumi:"triggerType"`
	// Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	TriggerWeekday *string `pulumi:"triggerWeekday"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
	Vdoms []AutomationtriggerVdom `pulumi:"vdoms"`
}

// The set of arguments for constructing a Automationtrigger resource.
type AutomationtriggerArgs struct {
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
	Fields AutomationtriggerFieldArrayInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// IOC threat level. Valid values: `medium`, `high`.
	IocLevel pulumi.StringPtrInput
	// License type.
	LicenseType pulumi.StringPtrInput
	// Log ID to trigger event. *Due to the data type change of API, for other versions of FortiOS, please check variable `logidBlock`.*
	Logid pulumi.IntPtrInput
	// Log IDs to trigger event. *Due to the data type change of API, for other versions of FortiOS, please check variable `logid`.* The structure of `logidBlock` block is documented below.
	LogidBlocks AutomationtriggerLogidBlockArrayInput
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
	// Minute of the hour on which to trigger (0 - 59, default = 0).
	TriggerMinute pulumi.IntPtrInput
	// Trigger type. Valid values: `event-based`, `scheduled`.
	TriggerType pulumi.StringPtrInput
	// Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	TriggerWeekday pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
	Vdoms AutomationtriggerVdomArrayInput
}

func (AutomationtriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationtriggerArgs)(nil)).Elem()
}

type AutomationtriggerInput interface {
	pulumi.Input

	ToAutomationtriggerOutput() AutomationtriggerOutput
	ToAutomationtriggerOutputWithContext(ctx context.Context) AutomationtriggerOutput
}

func (*Automationtrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**Automationtrigger)(nil)).Elem()
}

func (i *Automationtrigger) ToAutomationtriggerOutput() AutomationtriggerOutput {
	return i.ToAutomationtriggerOutputWithContext(context.Background())
}

func (i *Automationtrigger) ToAutomationtriggerOutputWithContext(ctx context.Context) AutomationtriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationtriggerOutput)
}

// AutomationtriggerArrayInput is an input type that accepts AutomationtriggerArray and AutomationtriggerArrayOutput values.
// You can construct a concrete instance of `AutomationtriggerArrayInput` via:
//
//	AutomationtriggerArray{ AutomationtriggerArgs{...} }
type AutomationtriggerArrayInput interface {
	pulumi.Input

	ToAutomationtriggerArrayOutput() AutomationtriggerArrayOutput
	ToAutomationtriggerArrayOutputWithContext(context.Context) AutomationtriggerArrayOutput
}

type AutomationtriggerArray []AutomationtriggerInput

func (AutomationtriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Automationtrigger)(nil)).Elem()
}

func (i AutomationtriggerArray) ToAutomationtriggerArrayOutput() AutomationtriggerArrayOutput {
	return i.ToAutomationtriggerArrayOutputWithContext(context.Background())
}

func (i AutomationtriggerArray) ToAutomationtriggerArrayOutputWithContext(ctx context.Context) AutomationtriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationtriggerArrayOutput)
}

// AutomationtriggerMapInput is an input type that accepts AutomationtriggerMap and AutomationtriggerMapOutput values.
// You can construct a concrete instance of `AutomationtriggerMapInput` via:
//
//	AutomationtriggerMap{ "key": AutomationtriggerArgs{...} }
type AutomationtriggerMapInput interface {
	pulumi.Input

	ToAutomationtriggerMapOutput() AutomationtriggerMapOutput
	ToAutomationtriggerMapOutputWithContext(context.Context) AutomationtriggerMapOutput
}

type AutomationtriggerMap map[string]AutomationtriggerInput

func (AutomationtriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Automationtrigger)(nil)).Elem()
}

func (i AutomationtriggerMap) ToAutomationtriggerMapOutput() AutomationtriggerMapOutput {
	return i.ToAutomationtriggerMapOutputWithContext(context.Background())
}

func (i AutomationtriggerMap) ToAutomationtriggerMapOutputWithContext(ctx context.Context) AutomationtriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationtriggerMapOutput)
}

type AutomationtriggerOutput struct{ *pulumi.OutputState }

func (AutomationtriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Automationtrigger)(nil)).Elem()
}

func (o AutomationtriggerOutput) ToAutomationtriggerOutput() AutomationtriggerOutput {
	return o
}

func (o AutomationtriggerOutput) ToAutomationtriggerOutputWithContext(ctx context.Context) AutomationtriggerOutput {
	return o
}

// Description.
func (o AutomationtriggerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o AutomationtriggerOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Event type.
func (o AutomationtriggerOutput) EventType() pulumi.StringOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringOutput { return v.EventType }).(pulumi.StringOutput)
}

// Fabric connector event handler name.
func (o AutomationtriggerOutput) FabricEventName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringPtrOutput { return v.FabricEventName }).(pulumi.StringPtrOutput)
}

// Fabric connector event severity.
func (o AutomationtriggerOutput) FabricEventSeverity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringPtrOutput { return v.FabricEventSeverity }).(pulumi.StringPtrOutput)
}

// FortiAnalyzer event handler name.
func (o AutomationtriggerOutput) FazEventName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringPtrOutput { return v.FazEventName }).(pulumi.StringPtrOutput)
}

// FortiAnalyzer event severity.
func (o AutomationtriggerOutput) FazEventSeverity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringPtrOutput { return v.FazEventSeverity }).(pulumi.StringPtrOutput)
}

// FortiAnalyzer event tags.
func (o AutomationtriggerOutput) FazEventTags() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringPtrOutput { return v.FazEventTags }).(pulumi.StringPtrOutput)
}

// Customized trigger field settings. The structure of `fields` block is documented below.
func (o AutomationtriggerOutput) Fields() AutomationtriggerFieldArrayOutput {
	return o.ApplyT(func(v *Automationtrigger) AutomationtriggerFieldArrayOutput { return v.Fields }).(AutomationtriggerFieldArrayOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o AutomationtriggerOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// IOC threat level. Valid values: `medium`, `high`.
func (o AutomationtriggerOutput) IocLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringOutput { return v.IocLevel }).(pulumi.StringOutput)
}

// License type.
func (o AutomationtriggerOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringOutput { return v.LicenseType }).(pulumi.StringOutput)
}

// Log ID to trigger event. *Due to the data type change of API, for other versions of FortiOS, please check variable `logidBlock`.*
func (o AutomationtriggerOutput) Logid() pulumi.IntOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.IntOutput { return v.Logid }).(pulumi.IntOutput)
}

// Log IDs to trigger event. *Due to the data type change of API, for other versions of FortiOS, please check variable `logid`.* The structure of `logidBlock` block is documented below.
func (o AutomationtriggerOutput) LogidBlocks() AutomationtriggerLogidBlockArrayOutput {
	return o.ApplyT(func(v *Automationtrigger) AutomationtriggerLogidBlockArrayOutput { return v.LogidBlocks }).(AutomationtriggerLogidBlockArrayOutput)
}

// Name.
func (o AutomationtriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Security Rating report.
func (o AutomationtriggerOutput) ReportType() pulumi.StringOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringOutput { return v.ReportType }).(pulumi.StringOutput)
}

// Fabric connector serial number.
func (o AutomationtriggerOutput) Serial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringPtrOutput { return v.Serial }).(pulumi.StringPtrOutput)
}

// Trigger date and time (YYYY-MM-DD HH:MM:SS).
func (o AutomationtriggerOutput) TriggerDatetime() pulumi.StringOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringOutput { return v.TriggerDatetime }).(pulumi.StringOutput)
}

// Day within a month to trigger.
func (o AutomationtriggerOutput) TriggerDay() pulumi.IntOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.IntOutput { return v.TriggerDay }).(pulumi.IntOutput)
}

// Scheduled trigger frequency (default = daily).
func (o AutomationtriggerOutput) TriggerFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringOutput { return v.TriggerFrequency }).(pulumi.StringOutput)
}

// Hour of the day on which to trigger (0 - 23, default = 1).
func (o AutomationtriggerOutput) TriggerHour() pulumi.IntOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.IntOutput { return v.TriggerHour }).(pulumi.IntOutput)
}

// Minute of the hour on which to trigger (0 - 59, default = 0).
func (o AutomationtriggerOutput) TriggerMinute() pulumi.IntOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.IntOutput { return v.TriggerMinute }).(pulumi.IntOutput)
}

// Trigger type. Valid values: `event-based`, `scheduled`.
func (o AutomationtriggerOutput) TriggerType() pulumi.StringOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringOutput { return v.TriggerType }).(pulumi.StringOutput)
}

// Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
func (o AutomationtriggerOutput) TriggerWeekday() pulumi.StringOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringOutput { return v.TriggerWeekday }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AutomationtriggerOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Automationtrigger) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
func (o AutomationtriggerOutput) Vdoms() AutomationtriggerVdomArrayOutput {
	return o.ApplyT(func(v *Automationtrigger) AutomationtriggerVdomArrayOutput { return v.Vdoms }).(AutomationtriggerVdomArrayOutput)
}

type AutomationtriggerArrayOutput struct{ *pulumi.OutputState }

func (AutomationtriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Automationtrigger)(nil)).Elem()
}

func (o AutomationtriggerArrayOutput) ToAutomationtriggerArrayOutput() AutomationtriggerArrayOutput {
	return o
}

func (o AutomationtriggerArrayOutput) ToAutomationtriggerArrayOutputWithContext(ctx context.Context) AutomationtriggerArrayOutput {
	return o
}

func (o AutomationtriggerArrayOutput) Index(i pulumi.IntInput) AutomationtriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Automationtrigger {
		return vs[0].([]*Automationtrigger)[vs[1].(int)]
	}).(AutomationtriggerOutput)
}

type AutomationtriggerMapOutput struct{ *pulumi.OutputState }

func (AutomationtriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Automationtrigger)(nil)).Elem()
}

func (o AutomationtriggerMapOutput) ToAutomationtriggerMapOutput() AutomationtriggerMapOutput {
	return o
}

func (o AutomationtriggerMapOutput) ToAutomationtriggerMapOutputWithContext(ctx context.Context) AutomationtriggerMapOutput {
	return o
}

func (o AutomationtriggerMapOutput) MapIndex(k pulumi.StringInput) AutomationtriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Automationtrigger {
		return vs[0].(map[string]*Automationtrigger)[vs[1].(string)]
	}).(AutomationtriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationtriggerInput)(nil)).Elem(), &Automationtrigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationtriggerArrayInput)(nil)).Elem(), AutomationtriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationtriggerMapInput)(nil)).Elem(), AutomationtriggerMap{})
	pulumi.RegisterOutputType(AutomationtriggerOutput{})
	pulumi.RegisterOutputType(AutomationtriggerArrayOutput{})
	pulumi.RegisterOutputType(AutomationtriggerMapOutput{})
}
