// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure log event filters.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/log"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := log.NewEventfilter(ctx, "trname", &log.EventfilterArgs{
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Log Eventfilter can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:log/eventfilter:Eventfilter labelname LogEventfilter
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:log/eventfilter:Eventfilter labelname LogEventfilter
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Eventfilter struct {
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

// NewEventfilter registers a new resource with the given unique name, arguments, and options.
func NewEventfilter(ctx *pulumi.Context,
	name string, args *EventfilterArgs, opts ...pulumi.ResourceOption) (*Eventfilter, error) {
	if args == nil {
		args = &EventfilterArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Eventfilter
	err := ctx.RegisterResource("fortios:log/eventfilter:Eventfilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventfilter gets an existing Eventfilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventfilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventfilterState, opts ...pulumi.ResourceOption) (*Eventfilter, error) {
	var resource Eventfilter
	err := ctx.ReadResource("fortios:log/eventfilter:Eventfilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Eventfilter resources.
type eventfilterState struct {
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

type EventfilterState struct {
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

func (EventfilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventfilterState)(nil)).Elem()
}

type eventfilterArgs struct {
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

// The set of arguments for constructing a Eventfilter resource.
type EventfilterArgs struct {
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

func (EventfilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventfilterArgs)(nil)).Elem()
}

type EventfilterInput interface {
	pulumi.Input

	ToEventfilterOutput() EventfilterOutput
	ToEventfilterOutputWithContext(ctx context.Context) EventfilterOutput
}

func (*Eventfilter) ElementType() reflect.Type {
	return reflect.TypeOf((**Eventfilter)(nil)).Elem()
}

func (i *Eventfilter) ToEventfilterOutput() EventfilterOutput {
	return i.ToEventfilterOutputWithContext(context.Background())
}

func (i *Eventfilter) ToEventfilterOutputWithContext(ctx context.Context) EventfilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventfilterOutput)
}

// EventfilterArrayInput is an input type that accepts EventfilterArray and EventfilterArrayOutput values.
// You can construct a concrete instance of `EventfilterArrayInput` via:
//
//	EventfilterArray{ EventfilterArgs{...} }
type EventfilterArrayInput interface {
	pulumi.Input

	ToEventfilterArrayOutput() EventfilterArrayOutput
	ToEventfilterArrayOutputWithContext(context.Context) EventfilterArrayOutput
}

type EventfilterArray []EventfilterInput

func (EventfilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eventfilter)(nil)).Elem()
}

func (i EventfilterArray) ToEventfilterArrayOutput() EventfilterArrayOutput {
	return i.ToEventfilterArrayOutputWithContext(context.Background())
}

func (i EventfilterArray) ToEventfilterArrayOutputWithContext(ctx context.Context) EventfilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventfilterArrayOutput)
}

// EventfilterMapInput is an input type that accepts EventfilterMap and EventfilterMapOutput values.
// You can construct a concrete instance of `EventfilterMapInput` via:
//
//	EventfilterMap{ "key": EventfilterArgs{...} }
type EventfilterMapInput interface {
	pulumi.Input

	ToEventfilterMapOutput() EventfilterMapOutput
	ToEventfilterMapOutputWithContext(context.Context) EventfilterMapOutput
}

type EventfilterMap map[string]EventfilterInput

func (EventfilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eventfilter)(nil)).Elem()
}

func (i EventfilterMap) ToEventfilterMapOutput() EventfilterMapOutput {
	return i.ToEventfilterMapOutputWithContext(context.Background())
}

func (i EventfilterMap) ToEventfilterMapOutputWithContext(ctx context.Context) EventfilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventfilterMapOutput)
}

type EventfilterOutput struct{ *pulumi.OutputState }

func (EventfilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Eventfilter)(nil)).Elem()
}

func (o EventfilterOutput) ToEventfilterOutput() EventfilterOutput {
	return o
}

func (o EventfilterOutput) ToEventfilterOutputWithContext(ctx context.Context) EventfilterOutput {
	return o
}

// Enable/disable CIFS logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) Cifs() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.Cifs }).(pulumi.StringOutput)
}

// Enable/disable PCI DSS compliance check logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) ComplianceCheck() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.ComplianceCheck }).(pulumi.StringOutput)
}

// Enable/disable SDN connector logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) Connector() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.Connector }).(pulumi.StringOutput)
}

// Enable/disable endpoint event logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Enable/disable event logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) Event() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.Event }).(pulumi.StringOutput)
}

// Enable/disable FortiExtender logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) Fortiextender() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.Fortiextender }).(pulumi.StringOutput)
}

// Enable/disable ha event logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) Ha() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.Ha }).(pulumi.StringOutput)
}

// Enable/disable REST API logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) RestApi() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.RestApi }).(pulumi.StringOutput)
}

// Enable/disable router event logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) Router() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.Router }).(pulumi.StringOutput)
}

// Enable/disable SD-WAN logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) Sdwan() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.Sdwan }).(pulumi.StringOutput)
}

// Enable/disable Security Rating result logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) SecurityRating() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.SecurityRating }).(pulumi.StringOutput)
}

// Enable/disable Switch-Controller logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) SwitchController() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.SwitchController }).(pulumi.StringOutput)
}

// Enable/disable system event logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) System() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.System }).(pulumi.StringOutput)
}

// Enable/disable user authentication event logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o EventfilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable VPN event logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) Vpn() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.Vpn }).(pulumi.StringOutput)
}

// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) WanOpt() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.WanOpt }).(pulumi.StringOutput)
}

// Enable/disable wireless event logging. Valid values: `enable`, `disable`.
func (o EventfilterOutput) WirelessActivity() pulumi.StringOutput {
	return o.ApplyT(func(v *Eventfilter) pulumi.StringOutput { return v.WirelessActivity }).(pulumi.StringOutput)
}

type EventfilterArrayOutput struct{ *pulumi.OutputState }

func (EventfilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eventfilter)(nil)).Elem()
}

func (o EventfilterArrayOutput) ToEventfilterArrayOutput() EventfilterArrayOutput {
	return o
}

func (o EventfilterArrayOutput) ToEventfilterArrayOutputWithContext(ctx context.Context) EventfilterArrayOutput {
	return o
}

func (o EventfilterArrayOutput) Index(i pulumi.IntInput) EventfilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Eventfilter {
		return vs[0].([]*Eventfilter)[vs[1].(int)]
	}).(EventfilterOutput)
}

type EventfilterMapOutput struct{ *pulumi.OutputState }

func (EventfilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eventfilter)(nil)).Elem()
}

func (o EventfilterMapOutput) ToEventfilterMapOutput() EventfilterMapOutput {
	return o
}

func (o EventfilterMapOutput) ToEventfilterMapOutputWithContext(ctx context.Context) EventfilterMapOutput {
	return o
}

func (o EventfilterMapOutput) MapIndex(k pulumi.StringInput) EventfilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Eventfilter {
		return vs[0].(map[string]*Eventfilter)[vs[1].(string)]
	}).(EventfilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventfilterInput)(nil)).Elem(), &Eventfilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventfilterArrayInput)(nil)).Elem(), EventfilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventfilterMapInput)(nil)).Elem(), EventfilterMap{})
	pulumi.RegisterOutputType(EventfilterOutput{})
	pulumi.RegisterOutputType(EventfilterArrayOutput{})
	pulumi.RegisterOutputType(EventfilterMapOutput{})
}
