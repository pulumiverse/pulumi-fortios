// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure log event filters.
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
//			_, err := fortios.NewLogEventfilter(ctx, "trname", &fortios.LogEventfilterArgs{
//				ComplianceCheck:  pulumi.String("enable"),
//				Endpoint:         pulumi.String("enable"),
//				Event:            pulumi.String("enable"),
//				Ha:               pulumi.String("enable"),
//				Router:           pulumi.String("enable"),
//				SecurityRating:   pulumi.String("enable"),
//				System:           pulumi.String("enable"),
//				User:             pulumi.String("enable"),
//				Vpn:              pulumi.String("enable"),
//				WanOpt:           pulumi.String("enable"),
//				WirelessActivity: pulumi.String("enable"),
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
// # Log Eventfilter can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/logEventfilter:LogEventfilter labelname LogEventfilter
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/logEventfilter:LogEventfilter labelname LogEventfilter
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type LogEventfilter struct {
	pulumi.CustomResourceState

	// Enable/disable CIFS logging. Valid values: `enable`, `disable`.
	Cifs pulumi.StringOutput `pulumi:"cifs"`
	// Enable/disable PCI DSS compliance check logging. Valid values: `enable`, `disable`.
	ComplianceCheck pulumi.StringOutput `pulumi:"complianceCheck"`
	// Enable/disable SDN connector logging. Valid values: `enable`, `disable`.
	Connector pulumi.StringOutput `pulumi:"connector"`
	// Enable/disable endpoint event logging. Valid values: `enable`, `disable`.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Enable/disable event logging. Valid values: `enable`, `disable`.
	Event pulumi.StringOutput `pulumi:"event"`
	// Enable/disable FortiExtender logging. Valid values: `enable`, `disable`.
	Fortiextender pulumi.StringOutput `pulumi:"fortiextender"`
	// Enable/disable ha event logging. Valid values: `enable`, `disable`.
	Ha pulumi.StringOutput `pulumi:"ha"`
	// Enable/disable REST API logging. Valid values: `enable`, `disable`.
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// Enable/disable router event logging. Valid values: `enable`, `disable`.
	Router pulumi.StringOutput `pulumi:"router"`
	// Enable/disable SD-WAN logging. Valid values: `enable`, `disable`.
	Sdwan pulumi.StringOutput `pulumi:"sdwan"`
	// Enable/disable Security Rating result logging. Valid values: `enable`, `disable`.
	SecurityRating pulumi.StringOutput `pulumi:"securityRating"`
	// Enable/disable Switch-Controller logging. Valid values: `enable`, `disable`.
	SwitchController pulumi.StringOutput `pulumi:"switchController"`
	// Enable/disable system event logging. Valid values: `enable`, `disable`.
	System pulumi.StringOutput `pulumi:"system"`
	// Enable/disable user authentication event logging. Valid values: `enable`, `disable`.
	User pulumi.StringOutput `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable VPN event logging. Valid values: `enable`, `disable`.
	Vpn pulumi.StringOutput `pulumi:"vpn"`
	// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
	WanOpt pulumi.StringOutput `pulumi:"wanOpt"`
	// Enable/disable wireless event logging. Valid values: `enable`, `disable`.
	WirelessActivity pulumi.StringOutput `pulumi:"wirelessActivity"`
}

// NewLogEventfilter registers a new resource with the given unique name, arguments, and options.
func NewLogEventfilter(ctx *pulumi.Context,
	name string, args *LogEventfilterArgs, opts ...pulumi.ResourceOption) (*LogEventfilter, error) {
	if args == nil {
		args = &LogEventfilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogEventfilter
	err := ctx.RegisterResource("fortios:index/logEventfilter:LogEventfilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogEventfilter gets an existing LogEventfilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogEventfilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogEventfilterState, opts ...pulumi.ResourceOption) (*LogEventfilter, error) {
	var resource LogEventfilter
	err := ctx.ReadResource("fortios:index/logEventfilter:LogEventfilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogEventfilter resources.
type logEventfilterState struct {
	// Enable/disable CIFS logging. Valid values: `enable`, `disable`.
	Cifs *string `pulumi:"cifs"`
	// Enable/disable PCI DSS compliance check logging. Valid values: `enable`, `disable`.
	ComplianceCheck *string `pulumi:"complianceCheck"`
	// Enable/disable SDN connector logging. Valid values: `enable`, `disable`.
	Connector *string `pulumi:"connector"`
	// Enable/disable endpoint event logging. Valid values: `enable`, `disable`.
	Endpoint *string `pulumi:"endpoint"`
	// Enable/disable event logging. Valid values: `enable`, `disable`.
	Event *string `pulumi:"event"`
	// Enable/disable FortiExtender logging. Valid values: `enable`, `disable`.
	Fortiextender *string `pulumi:"fortiextender"`
	// Enable/disable ha event logging. Valid values: `enable`, `disable`.
	Ha *string `pulumi:"ha"`
	// Enable/disable REST API logging. Valid values: `enable`, `disable`.
	RestApi *string `pulumi:"restApi"`
	// Enable/disable router event logging. Valid values: `enable`, `disable`.
	Router *string `pulumi:"router"`
	// Enable/disable SD-WAN logging. Valid values: `enable`, `disable`.
	Sdwan *string `pulumi:"sdwan"`
	// Enable/disable Security Rating result logging. Valid values: `enable`, `disable`.
	SecurityRating *string `pulumi:"securityRating"`
	// Enable/disable Switch-Controller logging. Valid values: `enable`, `disable`.
	SwitchController *string `pulumi:"switchController"`
	// Enable/disable system event logging. Valid values: `enable`, `disable`.
	System *string `pulumi:"system"`
	// Enable/disable user authentication event logging. Valid values: `enable`, `disable`.
	User *string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable VPN event logging. Valid values: `enable`, `disable`.
	Vpn *string `pulumi:"vpn"`
	// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
	WanOpt *string `pulumi:"wanOpt"`
	// Enable/disable wireless event logging. Valid values: `enable`, `disable`.
	WirelessActivity *string `pulumi:"wirelessActivity"`
}

type LogEventfilterState struct {
	// Enable/disable CIFS logging. Valid values: `enable`, `disable`.
	Cifs pulumi.StringPtrInput
	// Enable/disable PCI DSS compliance check logging. Valid values: `enable`, `disable`.
	ComplianceCheck pulumi.StringPtrInput
	// Enable/disable SDN connector logging. Valid values: `enable`, `disable`.
	Connector pulumi.StringPtrInput
	// Enable/disable endpoint event logging. Valid values: `enable`, `disable`.
	Endpoint pulumi.StringPtrInput
	// Enable/disable event logging. Valid values: `enable`, `disable`.
	Event pulumi.StringPtrInput
	// Enable/disable FortiExtender logging. Valid values: `enable`, `disable`.
	Fortiextender pulumi.StringPtrInput
	// Enable/disable ha event logging. Valid values: `enable`, `disable`.
	Ha pulumi.StringPtrInput
	// Enable/disable REST API logging. Valid values: `enable`, `disable`.
	RestApi pulumi.StringPtrInput
	// Enable/disable router event logging. Valid values: `enable`, `disable`.
	Router pulumi.StringPtrInput
	// Enable/disable SD-WAN logging. Valid values: `enable`, `disable`.
	Sdwan pulumi.StringPtrInput
	// Enable/disable Security Rating result logging. Valid values: `enable`, `disable`.
	SecurityRating pulumi.StringPtrInput
	// Enable/disable Switch-Controller logging. Valid values: `enable`, `disable`.
	SwitchController pulumi.StringPtrInput
	// Enable/disable system event logging. Valid values: `enable`, `disable`.
	System pulumi.StringPtrInput
	// Enable/disable user authentication event logging. Valid values: `enable`, `disable`.
	User pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable VPN event logging. Valid values: `enable`, `disable`.
	Vpn pulumi.StringPtrInput
	// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
	WanOpt pulumi.StringPtrInput
	// Enable/disable wireless event logging. Valid values: `enable`, `disable`.
	WirelessActivity pulumi.StringPtrInput
}

func (LogEventfilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logEventfilterState)(nil)).Elem()
}

type logEventfilterArgs struct {
	// Enable/disable CIFS logging. Valid values: `enable`, `disable`.
	Cifs *string `pulumi:"cifs"`
	// Enable/disable PCI DSS compliance check logging. Valid values: `enable`, `disable`.
	ComplianceCheck *string `pulumi:"complianceCheck"`
	// Enable/disable SDN connector logging. Valid values: `enable`, `disable`.
	Connector *string `pulumi:"connector"`
	// Enable/disable endpoint event logging. Valid values: `enable`, `disable`.
	Endpoint *string `pulumi:"endpoint"`
	// Enable/disable event logging. Valid values: `enable`, `disable`.
	Event *string `pulumi:"event"`
	// Enable/disable FortiExtender logging. Valid values: `enable`, `disable`.
	Fortiextender *string `pulumi:"fortiextender"`
	// Enable/disable ha event logging. Valid values: `enable`, `disable`.
	Ha *string `pulumi:"ha"`
	// Enable/disable REST API logging. Valid values: `enable`, `disable`.
	RestApi *string `pulumi:"restApi"`
	// Enable/disable router event logging. Valid values: `enable`, `disable`.
	Router *string `pulumi:"router"`
	// Enable/disable SD-WAN logging. Valid values: `enable`, `disable`.
	Sdwan *string `pulumi:"sdwan"`
	// Enable/disable Security Rating result logging. Valid values: `enable`, `disable`.
	SecurityRating *string `pulumi:"securityRating"`
	// Enable/disable Switch-Controller logging. Valid values: `enable`, `disable`.
	SwitchController *string `pulumi:"switchController"`
	// Enable/disable system event logging. Valid values: `enable`, `disable`.
	System *string `pulumi:"system"`
	// Enable/disable user authentication event logging. Valid values: `enable`, `disable`.
	User *string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable VPN event logging. Valid values: `enable`, `disable`.
	Vpn *string `pulumi:"vpn"`
	// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
	WanOpt *string `pulumi:"wanOpt"`
	// Enable/disable wireless event logging. Valid values: `enable`, `disable`.
	WirelessActivity *string `pulumi:"wirelessActivity"`
}

// The set of arguments for constructing a LogEventfilter resource.
type LogEventfilterArgs struct {
	// Enable/disable CIFS logging. Valid values: `enable`, `disable`.
	Cifs pulumi.StringPtrInput
	// Enable/disable PCI DSS compliance check logging. Valid values: `enable`, `disable`.
	ComplianceCheck pulumi.StringPtrInput
	// Enable/disable SDN connector logging. Valid values: `enable`, `disable`.
	Connector pulumi.StringPtrInput
	// Enable/disable endpoint event logging. Valid values: `enable`, `disable`.
	Endpoint pulumi.StringPtrInput
	// Enable/disable event logging. Valid values: `enable`, `disable`.
	Event pulumi.StringPtrInput
	// Enable/disable FortiExtender logging. Valid values: `enable`, `disable`.
	Fortiextender pulumi.StringPtrInput
	// Enable/disable ha event logging. Valid values: `enable`, `disable`.
	Ha pulumi.StringPtrInput
	// Enable/disable REST API logging. Valid values: `enable`, `disable`.
	RestApi pulumi.StringPtrInput
	// Enable/disable router event logging. Valid values: `enable`, `disable`.
	Router pulumi.StringPtrInput
	// Enable/disable SD-WAN logging. Valid values: `enable`, `disable`.
	Sdwan pulumi.StringPtrInput
	// Enable/disable Security Rating result logging. Valid values: `enable`, `disable`.
	SecurityRating pulumi.StringPtrInput
	// Enable/disable Switch-Controller logging. Valid values: `enable`, `disable`.
	SwitchController pulumi.StringPtrInput
	// Enable/disable system event logging. Valid values: `enable`, `disable`.
	System pulumi.StringPtrInput
	// Enable/disable user authentication event logging. Valid values: `enable`, `disable`.
	User pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable VPN event logging. Valid values: `enable`, `disable`.
	Vpn pulumi.StringPtrInput
	// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
	WanOpt pulumi.StringPtrInput
	// Enable/disable wireless event logging. Valid values: `enable`, `disable`.
	WirelessActivity pulumi.StringPtrInput
}

func (LogEventfilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logEventfilterArgs)(nil)).Elem()
}

type LogEventfilterInput interface {
	pulumi.Input

	ToLogEventfilterOutput() LogEventfilterOutput
	ToLogEventfilterOutputWithContext(ctx context.Context) LogEventfilterOutput
}

func (*LogEventfilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogEventfilter)(nil)).Elem()
}

func (i *LogEventfilter) ToLogEventfilterOutput() LogEventfilterOutput {
	return i.ToLogEventfilterOutputWithContext(context.Background())
}

func (i *LogEventfilter) ToLogEventfilterOutputWithContext(ctx context.Context) LogEventfilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogEventfilterOutput)
}

// LogEventfilterArrayInput is an input type that accepts LogEventfilterArray and LogEventfilterArrayOutput values.
// You can construct a concrete instance of `LogEventfilterArrayInput` via:
//
//	LogEventfilterArray{ LogEventfilterArgs{...} }
type LogEventfilterArrayInput interface {
	pulumi.Input

	ToLogEventfilterArrayOutput() LogEventfilterArrayOutput
	ToLogEventfilterArrayOutputWithContext(context.Context) LogEventfilterArrayOutput
}

type LogEventfilterArray []LogEventfilterInput

func (LogEventfilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogEventfilter)(nil)).Elem()
}

func (i LogEventfilterArray) ToLogEventfilterArrayOutput() LogEventfilterArrayOutput {
	return i.ToLogEventfilterArrayOutputWithContext(context.Background())
}

func (i LogEventfilterArray) ToLogEventfilterArrayOutputWithContext(ctx context.Context) LogEventfilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogEventfilterArrayOutput)
}

// LogEventfilterMapInput is an input type that accepts LogEventfilterMap and LogEventfilterMapOutput values.
// You can construct a concrete instance of `LogEventfilterMapInput` via:
//
//	LogEventfilterMap{ "key": LogEventfilterArgs{...} }
type LogEventfilterMapInput interface {
	pulumi.Input

	ToLogEventfilterMapOutput() LogEventfilterMapOutput
	ToLogEventfilterMapOutputWithContext(context.Context) LogEventfilterMapOutput
}

type LogEventfilterMap map[string]LogEventfilterInput

func (LogEventfilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogEventfilter)(nil)).Elem()
}

func (i LogEventfilterMap) ToLogEventfilterMapOutput() LogEventfilterMapOutput {
	return i.ToLogEventfilterMapOutputWithContext(context.Background())
}

func (i LogEventfilterMap) ToLogEventfilterMapOutputWithContext(ctx context.Context) LogEventfilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogEventfilterMapOutput)
}

type LogEventfilterOutput struct{ *pulumi.OutputState }

func (LogEventfilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogEventfilter)(nil)).Elem()
}

func (o LogEventfilterOutput) ToLogEventfilterOutput() LogEventfilterOutput {
	return o
}

func (o LogEventfilterOutput) ToLogEventfilterOutputWithContext(ctx context.Context) LogEventfilterOutput {
	return o
}

// Enable/disable CIFS logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) Cifs() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.Cifs }).(pulumi.StringOutput)
}

// Enable/disable PCI DSS compliance check logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) ComplianceCheck() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.ComplianceCheck }).(pulumi.StringOutput)
}

// Enable/disable SDN connector logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) Connector() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.Connector }).(pulumi.StringOutput)
}

// Enable/disable endpoint event logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Enable/disable event logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) Event() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.Event }).(pulumi.StringOutput)
}

// Enable/disable FortiExtender logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) Fortiextender() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.Fortiextender }).(pulumi.StringOutput)
}

// Enable/disable ha event logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) Ha() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.Ha }).(pulumi.StringOutput)
}

// Enable/disable REST API logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) RestApi() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.RestApi }).(pulumi.StringOutput)
}

// Enable/disable router event logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) Router() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.Router }).(pulumi.StringOutput)
}

// Enable/disable SD-WAN logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) Sdwan() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.Sdwan }).(pulumi.StringOutput)
}

// Enable/disable Security Rating result logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) SecurityRating() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.SecurityRating }).(pulumi.StringOutput)
}

// Enable/disable Switch-Controller logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) SwitchController() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.SwitchController }).(pulumi.StringOutput)
}

// Enable/disable system event logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) System() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.System }).(pulumi.StringOutput)
}

// Enable/disable user authentication event logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o LogEventfilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable VPN event logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) Vpn() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.Vpn }).(pulumi.StringOutput)
}

// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) WanOpt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.WanOpt }).(pulumi.StringOutput)
}

// Enable/disable wireless event logging. Valid values: `enable`, `disable`.
func (o LogEventfilterOutput) WirelessActivity() pulumi.StringOutput {
	return o.ApplyT(func(v *LogEventfilter) pulumi.StringOutput { return v.WirelessActivity }).(pulumi.StringOutput)
}

type LogEventfilterArrayOutput struct{ *pulumi.OutputState }

func (LogEventfilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogEventfilter)(nil)).Elem()
}

func (o LogEventfilterArrayOutput) ToLogEventfilterArrayOutput() LogEventfilterArrayOutput {
	return o
}

func (o LogEventfilterArrayOutput) ToLogEventfilterArrayOutputWithContext(ctx context.Context) LogEventfilterArrayOutput {
	return o
}

func (o LogEventfilterArrayOutput) Index(i pulumi.IntInput) LogEventfilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogEventfilter {
		return vs[0].([]*LogEventfilter)[vs[1].(int)]
	}).(LogEventfilterOutput)
}

type LogEventfilterMapOutput struct{ *pulumi.OutputState }

func (LogEventfilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogEventfilter)(nil)).Elem()
}

func (o LogEventfilterMapOutput) ToLogEventfilterMapOutput() LogEventfilterMapOutput {
	return o
}

func (o LogEventfilterMapOutput) ToLogEventfilterMapOutputWithContext(ctx context.Context) LogEventfilterMapOutput {
	return o
}

func (o LogEventfilterMapOutput) MapIndex(k pulumi.StringInput) LogEventfilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogEventfilter {
		return vs[0].(map[string]*LogEventfilter)[vs[1].(string)]
	}).(LogEventfilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogEventfilterInput)(nil)).Elem(), &LogEventfilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogEventfilterArrayInput)(nil)).Elem(), LogEventfilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogEventfilterMapInput)(nil)).Elem(), LogEventfilterMap{})
	pulumi.RegisterOutputType(LogEventfilterOutput{})
	pulumi.RegisterOutputType(LogEventfilterArrayOutput{})
	pulumi.RegisterOutputType(LogEventfilterMapOutput{})
}