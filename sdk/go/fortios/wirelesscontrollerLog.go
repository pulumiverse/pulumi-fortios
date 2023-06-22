// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure wireless controller event log filters. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # WirelessController Log can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerLog:WirelesscontrollerLog labelname WirelessControllerLog
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerLog:WirelesscontrollerLog labelname WirelessControllerLog
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WirelesscontrollerLog struct {
	pulumi.CustomResourceState

	// Lowest severity level to log address group message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	AddrgrpLog pulumi.StringOutput `pulumi:"addrgrpLog"`
	// Lowest severity level to log BLE detection message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	BleLog pulumi.StringOutput `pulumi:"bleLog"`
	// Lowest severity level to log client load balancing message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	ClbLog pulumi.StringOutput `pulumi:"clbLog"`
	// Lowest severity level to log DHCP starvation event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	DhcpStarvLog pulumi.StringOutput `pulumi:"dhcpStarvLog"`
	// Lowest severity level to log LED schedule event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	LedSchedLog pulumi.StringOutput `pulumi:"ledSchedLog"`
	// Lowest severity level to log radio event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	RadioEventLog pulumi.StringOutput `pulumi:"radioEventLog"`
	// Lowest severity level to log rogue AP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	RogueEventLog pulumi.StringOutput `pulumi:"rogueEventLog"`
	// Lowest severity level to log station event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	StaEventLog pulumi.StringOutput `pulumi:"staEventLog"`
	// Lowest severity level to log station locate message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	StaLocateLog pulumi.StringOutput `pulumi:"staLocateLog"`
	// Enable/disable wireless event logging. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Lowest severity level to log WIDS message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	WidsLog pulumi.StringOutput `pulumi:"widsLog"`
	// Lowest severity level to log WTP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	WtpEventLog pulumi.StringOutput `pulumi:"wtpEventLog"`
}

// NewWirelesscontrollerLog registers a new resource with the given unique name, arguments, and options.
func NewWirelesscontrollerLog(ctx *pulumi.Context,
	name string, args *WirelesscontrollerLogArgs, opts ...pulumi.ResourceOption) (*WirelesscontrollerLog, error) {
	if args == nil {
		args = &WirelesscontrollerLogArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelesscontrollerLog
	err := ctx.RegisterResource("fortios:index/wirelesscontrollerLog:WirelesscontrollerLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelesscontrollerLog gets an existing WirelesscontrollerLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelesscontrollerLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelesscontrollerLogState, opts ...pulumi.ResourceOption) (*WirelesscontrollerLog, error) {
	var resource WirelesscontrollerLog
	err := ctx.ReadResource("fortios:index/wirelesscontrollerLog:WirelesscontrollerLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelesscontrollerLog resources.
type wirelesscontrollerLogState struct {
	// Lowest severity level to log address group message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	AddrgrpLog *string `pulumi:"addrgrpLog"`
	// Lowest severity level to log BLE detection message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	BleLog *string `pulumi:"bleLog"`
	// Lowest severity level to log client load balancing message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	ClbLog *string `pulumi:"clbLog"`
	// Lowest severity level to log DHCP starvation event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	DhcpStarvLog *string `pulumi:"dhcpStarvLog"`
	// Lowest severity level to log LED schedule event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	LedSchedLog *string `pulumi:"ledSchedLog"`
	// Lowest severity level to log radio event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	RadioEventLog *string `pulumi:"radioEventLog"`
	// Lowest severity level to log rogue AP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	RogueEventLog *string `pulumi:"rogueEventLog"`
	// Lowest severity level to log station event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	StaEventLog *string `pulumi:"staEventLog"`
	// Lowest severity level to log station locate message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	StaLocateLog *string `pulumi:"staLocateLog"`
	// Enable/disable wireless event logging. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Lowest severity level to log WIDS message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	WidsLog *string `pulumi:"widsLog"`
	// Lowest severity level to log WTP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	WtpEventLog *string `pulumi:"wtpEventLog"`
}

type WirelesscontrollerLogState struct {
	// Lowest severity level to log address group message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	AddrgrpLog pulumi.StringPtrInput
	// Lowest severity level to log BLE detection message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	BleLog pulumi.StringPtrInput
	// Lowest severity level to log client load balancing message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	ClbLog pulumi.StringPtrInput
	// Lowest severity level to log DHCP starvation event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	DhcpStarvLog pulumi.StringPtrInput
	// Lowest severity level to log LED schedule event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	LedSchedLog pulumi.StringPtrInput
	// Lowest severity level to log radio event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	RadioEventLog pulumi.StringPtrInput
	// Lowest severity level to log rogue AP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	RogueEventLog pulumi.StringPtrInput
	// Lowest severity level to log station event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	StaEventLog pulumi.StringPtrInput
	// Lowest severity level to log station locate message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	StaLocateLog pulumi.StringPtrInput
	// Enable/disable wireless event logging. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Lowest severity level to log WIDS message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	WidsLog pulumi.StringPtrInput
	// Lowest severity level to log WTP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	WtpEventLog pulumi.StringPtrInput
}

func (WirelesscontrollerLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerLogState)(nil)).Elem()
}

type wirelesscontrollerLogArgs struct {
	// Lowest severity level to log address group message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	AddrgrpLog *string `pulumi:"addrgrpLog"`
	// Lowest severity level to log BLE detection message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	BleLog *string `pulumi:"bleLog"`
	// Lowest severity level to log client load balancing message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	ClbLog *string `pulumi:"clbLog"`
	// Lowest severity level to log DHCP starvation event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	DhcpStarvLog *string `pulumi:"dhcpStarvLog"`
	// Lowest severity level to log LED schedule event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	LedSchedLog *string `pulumi:"ledSchedLog"`
	// Lowest severity level to log radio event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	RadioEventLog *string `pulumi:"radioEventLog"`
	// Lowest severity level to log rogue AP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	RogueEventLog *string `pulumi:"rogueEventLog"`
	// Lowest severity level to log station event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	StaEventLog *string `pulumi:"staEventLog"`
	// Lowest severity level to log station locate message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	StaLocateLog *string `pulumi:"staLocateLog"`
	// Enable/disable wireless event logging. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Lowest severity level to log WIDS message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	WidsLog *string `pulumi:"widsLog"`
	// Lowest severity level to log WTP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	WtpEventLog *string `pulumi:"wtpEventLog"`
}

// The set of arguments for constructing a WirelesscontrollerLog resource.
type WirelesscontrollerLogArgs struct {
	// Lowest severity level to log address group message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	AddrgrpLog pulumi.StringPtrInput
	// Lowest severity level to log BLE detection message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	BleLog pulumi.StringPtrInput
	// Lowest severity level to log client load balancing message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	ClbLog pulumi.StringPtrInput
	// Lowest severity level to log DHCP starvation event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	DhcpStarvLog pulumi.StringPtrInput
	// Lowest severity level to log LED schedule event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	LedSchedLog pulumi.StringPtrInput
	// Lowest severity level to log radio event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	RadioEventLog pulumi.StringPtrInput
	// Lowest severity level to log rogue AP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	RogueEventLog pulumi.StringPtrInput
	// Lowest severity level to log station event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	StaEventLog pulumi.StringPtrInput
	// Lowest severity level to log station locate message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	StaLocateLog pulumi.StringPtrInput
	// Enable/disable wireless event logging. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Lowest severity level to log WIDS message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	WidsLog pulumi.StringPtrInput
	// Lowest severity level to log WTP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	WtpEventLog pulumi.StringPtrInput
}

func (WirelesscontrollerLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerLogArgs)(nil)).Elem()
}

type WirelesscontrollerLogInput interface {
	pulumi.Input

	ToWirelesscontrollerLogOutput() WirelesscontrollerLogOutput
	ToWirelesscontrollerLogOutputWithContext(ctx context.Context) WirelesscontrollerLogOutput
}

func (*WirelesscontrollerLog) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerLog)(nil)).Elem()
}

func (i *WirelesscontrollerLog) ToWirelesscontrollerLogOutput() WirelesscontrollerLogOutput {
	return i.ToWirelesscontrollerLogOutputWithContext(context.Background())
}

func (i *WirelesscontrollerLog) ToWirelesscontrollerLogOutputWithContext(ctx context.Context) WirelesscontrollerLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerLogOutput)
}

// WirelesscontrollerLogArrayInput is an input type that accepts WirelesscontrollerLogArray and WirelesscontrollerLogArrayOutput values.
// You can construct a concrete instance of `WirelesscontrollerLogArrayInput` via:
//
//	WirelesscontrollerLogArray{ WirelesscontrollerLogArgs{...} }
type WirelesscontrollerLogArrayInput interface {
	pulumi.Input

	ToWirelesscontrollerLogArrayOutput() WirelesscontrollerLogArrayOutput
	ToWirelesscontrollerLogArrayOutputWithContext(context.Context) WirelesscontrollerLogArrayOutput
}

type WirelesscontrollerLogArray []WirelesscontrollerLogInput

func (WirelesscontrollerLogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerLog)(nil)).Elem()
}

func (i WirelesscontrollerLogArray) ToWirelesscontrollerLogArrayOutput() WirelesscontrollerLogArrayOutput {
	return i.ToWirelesscontrollerLogArrayOutputWithContext(context.Background())
}

func (i WirelesscontrollerLogArray) ToWirelesscontrollerLogArrayOutputWithContext(ctx context.Context) WirelesscontrollerLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerLogArrayOutput)
}

// WirelesscontrollerLogMapInput is an input type that accepts WirelesscontrollerLogMap and WirelesscontrollerLogMapOutput values.
// You can construct a concrete instance of `WirelesscontrollerLogMapInput` via:
//
//	WirelesscontrollerLogMap{ "key": WirelesscontrollerLogArgs{...} }
type WirelesscontrollerLogMapInput interface {
	pulumi.Input

	ToWirelesscontrollerLogMapOutput() WirelesscontrollerLogMapOutput
	ToWirelesscontrollerLogMapOutputWithContext(context.Context) WirelesscontrollerLogMapOutput
}

type WirelesscontrollerLogMap map[string]WirelesscontrollerLogInput

func (WirelesscontrollerLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerLog)(nil)).Elem()
}

func (i WirelesscontrollerLogMap) ToWirelesscontrollerLogMapOutput() WirelesscontrollerLogMapOutput {
	return i.ToWirelesscontrollerLogMapOutputWithContext(context.Background())
}

func (i WirelesscontrollerLogMap) ToWirelesscontrollerLogMapOutputWithContext(ctx context.Context) WirelesscontrollerLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerLogMapOutput)
}

type WirelesscontrollerLogOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerLog)(nil)).Elem()
}

func (o WirelesscontrollerLogOutput) ToWirelesscontrollerLogOutput() WirelesscontrollerLogOutput {
	return o
}

func (o WirelesscontrollerLogOutput) ToWirelesscontrollerLogOutputWithContext(ctx context.Context) WirelesscontrollerLogOutput {
	return o
}

// Lowest severity level to log address group message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o WirelesscontrollerLogOutput) AddrgrpLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerLog) pulumi.StringOutput { return v.AddrgrpLog }).(pulumi.StringOutput)
}

// Lowest severity level to log BLE detection message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o WirelesscontrollerLogOutput) BleLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerLog) pulumi.StringOutput { return v.BleLog }).(pulumi.StringOutput)
}

// Lowest severity level to log client load balancing message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o WirelesscontrollerLogOutput) ClbLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerLog) pulumi.StringOutput { return v.ClbLog }).(pulumi.StringOutput)
}

// Lowest severity level to log DHCP starvation event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o WirelesscontrollerLogOutput) DhcpStarvLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerLog) pulumi.StringOutput { return v.DhcpStarvLog }).(pulumi.StringOutput)
}

// Lowest severity level to log LED schedule event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o WirelesscontrollerLogOutput) LedSchedLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerLog) pulumi.StringOutput { return v.LedSchedLog }).(pulumi.StringOutput)
}

// Lowest severity level to log radio event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o WirelesscontrollerLogOutput) RadioEventLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerLog) pulumi.StringOutput { return v.RadioEventLog }).(pulumi.StringOutput)
}

// Lowest severity level to log rogue AP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o WirelesscontrollerLogOutput) RogueEventLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerLog) pulumi.StringOutput { return v.RogueEventLog }).(pulumi.StringOutput)
}

// Lowest severity level to log station event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o WirelesscontrollerLogOutput) StaEventLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerLog) pulumi.StringOutput { return v.StaEventLog }).(pulumi.StringOutput)
}

// Lowest severity level to log station locate message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o WirelesscontrollerLogOutput) StaLocateLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerLog) pulumi.StringOutput { return v.StaLocateLog }).(pulumi.StringOutput)
}

// Enable/disable wireless event logging. Valid values: `enable`, `disable`.
func (o WirelesscontrollerLogOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerLog) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WirelesscontrollerLogOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelesscontrollerLog) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Lowest severity level to log WIDS message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o WirelesscontrollerLogOutput) WidsLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerLog) pulumi.StringOutput { return v.WidsLog }).(pulumi.StringOutput)
}

// Lowest severity level to log WTP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o WirelesscontrollerLogOutput) WtpEventLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerLog) pulumi.StringOutput { return v.WtpEventLog }).(pulumi.StringOutput)
}

type WirelesscontrollerLogArrayOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerLog)(nil)).Elem()
}

func (o WirelesscontrollerLogArrayOutput) ToWirelesscontrollerLogArrayOutput() WirelesscontrollerLogArrayOutput {
	return o
}

func (o WirelesscontrollerLogArrayOutput) ToWirelesscontrollerLogArrayOutputWithContext(ctx context.Context) WirelesscontrollerLogArrayOutput {
	return o
}

func (o WirelesscontrollerLogArrayOutput) Index(i pulumi.IntInput) WirelesscontrollerLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelesscontrollerLog {
		return vs[0].([]*WirelesscontrollerLog)[vs[1].(int)]
	}).(WirelesscontrollerLogOutput)
}

type WirelesscontrollerLogMapOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerLog)(nil)).Elem()
}

func (o WirelesscontrollerLogMapOutput) ToWirelesscontrollerLogMapOutput() WirelesscontrollerLogMapOutput {
	return o
}

func (o WirelesscontrollerLogMapOutput) ToWirelesscontrollerLogMapOutputWithContext(ctx context.Context) WirelesscontrollerLogMapOutput {
	return o
}

func (o WirelesscontrollerLogMapOutput) MapIndex(k pulumi.StringInput) WirelesscontrollerLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelesscontrollerLog {
		return vs[0].(map[string]*WirelesscontrollerLog)[vs[1].(string)]
	}).(WirelesscontrollerLogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerLogInput)(nil)).Elem(), &WirelesscontrollerLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerLogArrayInput)(nil)).Elem(), WirelesscontrollerLogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerLogMapInput)(nil)).Elem(), WirelesscontrollerLogMap{})
	pulumi.RegisterOutputType(WirelesscontrollerLogOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerLogArrayOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerLogMapOutput{})
}
