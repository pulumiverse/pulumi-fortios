// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system automationtrigger
func LookupSystemAutomationtrigger(ctx *pulumi.Context, args *LookupSystemAutomationtriggerArgs, opts ...pulumi.InvokeOption) (*LookupSystemAutomationtriggerResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemAutomationtriggerResult
	err := ctx.Invoke("fortios:index/getSystemAutomationtrigger:getSystemAutomationtrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemAutomationtrigger.
type LookupSystemAutomationtriggerArgs struct {
	// Specify the name of the desired system automationtrigger.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemAutomationtrigger.
type LookupSystemAutomationtriggerResult struct {
	// Description.
	Description string `pulumi:"description"`
	// Event type.
	EventType string `pulumi:"eventType"`
	// Fabric connector event handler name.
	FabricEventName string `pulumi:"fabricEventName"`
	// Fabric connector event severity.
	FabricEventSeverity string `pulumi:"fabricEventSeverity"`
	// FortiAnalyzer event handler name.
	FazEventName string `pulumi:"fazEventName"`
	// FortiAnalyzer event severity.
	FazEventSeverity string `pulumi:"fazEventSeverity"`
	// FortiAnalyzer event tags.
	FazEventTags string `pulumi:"fazEventTags"`
	// Customized trigger field settings. The structure of `fields` block is documented below.
	Fields []GetSystemAutomationtriggerField `pulumi:"fields"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IOC threat level.
	IocLevel string `pulumi:"iocLevel"`
	// License type.
	LicenseType string `pulumi:"licenseType"`
	// Log ID to trigger event.
	Logid int `pulumi:"logid"`
	// Log ID to trigger event. Only applies on FortiOS v7.0.0+. The structure of `logidBlock` block is documented below.
	LogidBlocks []GetSystemAutomationtriggerLogidBlock `pulumi:"logidBlocks"`
	// Name.
	Name string `pulumi:"name"`
	// Security Rating report.
	ReportType string `pulumi:"reportType"`
	// Fabric connector serial number.
	Serial string `pulumi:"serial"`
	// Trigger date and time (YYYY-MM-DD HH:MM:SS).
	TriggerDatetime string `pulumi:"triggerDatetime"`
	// Day within a month to trigger.
	TriggerDay int `pulumi:"triggerDay"`
	// Scheduled trigger frequency (default = daily).
	TriggerFrequency string `pulumi:"triggerFrequency"`
	// Hour of the day on which to trigger (0 - 23, default = 1).
	TriggerHour int `pulumi:"triggerHour"`
	// Minute of the hour on which to trigger (0 - 59, 60 to randomize).
	TriggerMinute int `pulumi:"triggerMinute"`
	// Trigger type.
	TriggerType string `pulumi:"triggerType"`
	// Day of week for trigger.
	TriggerWeekday string  `pulumi:"triggerWeekday"`
	Vdomparam      *string `pulumi:"vdomparam"`
	// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
	Vdoms []GetSystemAutomationtriggerVdom `pulumi:"vdoms"`
}

func LookupSystemAutomationtriggerOutput(ctx *pulumi.Context, args LookupSystemAutomationtriggerOutputArgs, opts ...pulumi.InvokeOption) LookupSystemAutomationtriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemAutomationtriggerResult, error) {
			args := v.(LookupSystemAutomationtriggerArgs)
			r, err := LookupSystemAutomationtrigger(ctx, &args, opts...)
			var s LookupSystemAutomationtriggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemAutomationtriggerResultOutput)
}

// A collection of arguments for invoking getSystemAutomationtrigger.
type LookupSystemAutomationtriggerOutputArgs struct {
	// Specify the name of the desired system automationtrigger.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemAutomationtriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAutomationtriggerArgs)(nil)).Elem()
}

// A collection of values returned by getSystemAutomationtrigger.
type LookupSystemAutomationtriggerResultOutput struct{ *pulumi.OutputState }

func (LookupSystemAutomationtriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAutomationtriggerResult)(nil)).Elem()
}

func (o LookupSystemAutomationtriggerResultOutput) ToLookupSystemAutomationtriggerResultOutput() LookupSystemAutomationtriggerResultOutput {
	return o
}

func (o LookupSystemAutomationtriggerResultOutput) ToLookupSystemAutomationtriggerResultOutputWithContext(ctx context.Context) LookupSystemAutomationtriggerResultOutput {
	return o
}

// Description.
func (o LookupSystemAutomationtriggerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.Description }).(pulumi.StringOutput)
}

// Event type.
func (o LookupSystemAutomationtriggerResultOutput) EventType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.EventType }).(pulumi.StringOutput)
}

// Fabric connector event handler name.
func (o LookupSystemAutomationtriggerResultOutput) FabricEventName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.FabricEventName }).(pulumi.StringOutput)
}

// Fabric connector event severity.
func (o LookupSystemAutomationtriggerResultOutput) FabricEventSeverity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.FabricEventSeverity }).(pulumi.StringOutput)
}

// FortiAnalyzer event handler name.
func (o LookupSystemAutomationtriggerResultOutput) FazEventName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.FazEventName }).(pulumi.StringOutput)
}

// FortiAnalyzer event severity.
func (o LookupSystemAutomationtriggerResultOutput) FazEventSeverity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.FazEventSeverity }).(pulumi.StringOutput)
}

// FortiAnalyzer event tags.
func (o LookupSystemAutomationtriggerResultOutput) FazEventTags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.FazEventTags }).(pulumi.StringOutput)
}

// Customized trigger field settings. The structure of `fields` block is documented below.
func (o LookupSystemAutomationtriggerResultOutput) Fields() GetSystemAutomationtriggerFieldArrayOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) []GetSystemAutomationtriggerField { return v.Fields }).(GetSystemAutomationtriggerFieldArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemAutomationtriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

// IOC threat level.
func (o LookupSystemAutomationtriggerResultOutput) IocLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.IocLevel }).(pulumi.StringOutput)
}

// License type.
func (o LookupSystemAutomationtriggerResultOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.LicenseType }).(pulumi.StringOutput)
}

// Log ID to trigger event.
func (o LookupSystemAutomationtriggerResultOutput) Logid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) int { return v.Logid }).(pulumi.IntOutput)
}

// Log ID to trigger event. Only applies on FortiOS v7.0.0+. The structure of `logidBlock` block is documented below.
func (o LookupSystemAutomationtriggerResultOutput) LogidBlocks() GetSystemAutomationtriggerLogidBlockArrayOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) []GetSystemAutomationtriggerLogidBlock {
		return v.LogidBlocks
	}).(GetSystemAutomationtriggerLogidBlockArrayOutput)
}

// Name.
func (o LookupSystemAutomationtriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Security Rating report.
func (o LookupSystemAutomationtriggerResultOutput) ReportType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.ReportType }).(pulumi.StringOutput)
}

// Fabric connector serial number.
func (o LookupSystemAutomationtriggerResultOutput) Serial() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.Serial }).(pulumi.StringOutput)
}

// Trigger date and time (YYYY-MM-DD HH:MM:SS).
func (o LookupSystemAutomationtriggerResultOutput) TriggerDatetime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.TriggerDatetime }).(pulumi.StringOutput)
}

// Day within a month to trigger.
func (o LookupSystemAutomationtriggerResultOutput) TriggerDay() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) int { return v.TriggerDay }).(pulumi.IntOutput)
}

// Scheduled trigger frequency (default = daily).
func (o LookupSystemAutomationtriggerResultOutput) TriggerFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.TriggerFrequency }).(pulumi.StringOutput)
}

// Hour of the day on which to trigger (0 - 23, default = 1).
func (o LookupSystemAutomationtriggerResultOutput) TriggerHour() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) int { return v.TriggerHour }).(pulumi.IntOutput)
}

// Minute of the hour on which to trigger (0 - 59, 60 to randomize).
func (o LookupSystemAutomationtriggerResultOutput) TriggerMinute() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) int { return v.TriggerMinute }).(pulumi.IntOutput)
}

// Trigger type.
func (o LookupSystemAutomationtriggerResultOutput) TriggerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.TriggerType }).(pulumi.StringOutput)
}

// Day of week for trigger.
func (o LookupSystemAutomationtriggerResultOutput) TriggerWeekday() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) string { return v.TriggerWeekday }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationtriggerResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
func (o LookupSystemAutomationtriggerResultOutput) Vdoms() GetSystemAutomationtriggerVdomArrayOutput {
	return o.ApplyT(func(v LookupSystemAutomationtriggerResult) []GetSystemAutomationtriggerVdom { return v.Vdoms }).(GetSystemAutomationtriggerVdomArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemAutomationtriggerResultOutput{})
}
