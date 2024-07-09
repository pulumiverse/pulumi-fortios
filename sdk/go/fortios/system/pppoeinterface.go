// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure the PPPoE interfaces.
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
//			_, err := system.NewPppoeinterface(ctx, "trname", &system.PppoeinterfaceArgs{
//				AuthType:                 pulumi.String("auto"),
//				Device:                   pulumi.String("port3"),
//				DialOnDemand:             pulumi.String("disable"),
//				DiscRetryTimeout:         pulumi.Int(1),
//				IdleTimeout:              pulumi.Int(0),
//				Ipunnumbered:             pulumi.String("0.0.0.0"),
//				Ipv6:                     pulumi.String("disable"),
//				LcpEchoInterval:          pulumi.Int(5),
//				LcpMaxEchoFails:          pulumi.Int(3),
//				PadtRetryTimeout:         pulumi.Int(1),
//				PppoeUnnumberedNegotiate: pulumi.String("enable"),
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
// System PppoeInterface can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/pppoeinterface:Pppoeinterface labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/pppoeinterface:Pppoeinterface labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Pppoeinterface struct {
	pulumi.CustomResourceState

	// PPPoE AC name.
	AcName pulumi.StringOutput `pulumi:"acName"`
	// PPP authentication type to use. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.
	AuthType pulumi.StringOutput `pulumi:"authType"`
	// Name for the physical interface.
	Device pulumi.StringOutput `pulumi:"device"`
	// Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface. Valid values: `enable`, `disable`.
	DialOnDemand pulumi.StringOutput `pulumi:"dialOnDemand"`
	// PPPoE discovery init timeout value in (0-4294967295 sec).
	DiscRetryTimeout pulumi.IntOutput `pulumi:"discRetryTimeout"`
	// PPPoE auto disconnect after idle timeout (0-4294967295 sec).
	IdleTimeout pulumi.IntOutput `pulumi:"idleTimeout"`
	// PPPoE unnumbered IP.
	Ipunnumbered pulumi.StringOutput `pulumi:"ipunnumbered"`
	// Enable/disable IPv6 Control Protocol (IPv6CP). Valid values: `enable`, `disable`.
	Ipv6 pulumi.StringOutput `pulumi:"ipv6"`
	// Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
	LcpEchoInterval pulumi.IntOutput `pulumi:"lcpEchoInterval"`
	// Maximum missed LCP echo messages before disconnect.
	LcpMaxEchoFails pulumi.IntOutput `pulumi:"lcpMaxEchoFails"`
	// Name of the PPPoE interface.
	Name pulumi.StringOutput `pulumi:"name"`
	// PPPoE terminate timeout value in (0-4294967295 sec).
	PadtRetryTimeout pulumi.IntOutput `pulumi:"padtRetryTimeout"`
	// Enter the password.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Enable/disable PPPoE unnumbered negotiation. Valid values: `enable`, `disable`.
	PppoeUnnumberedNegotiate pulumi.StringOutput `pulumi:"pppoeUnnumberedNegotiate"`
	// PPPoE service name.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// User name.
	Username pulumi.StringOutput `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewPppoeinterface registers a new resource with the given unique name, arguments, and options.
func NewPppoeinterface(ctx *pulumi.Context,
	name string, args *PppoeinterfaceArgs, opts ...pulumi.ResourceOption) (*Pppoeinterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Device == nil {
		return nil, errors.New("invalid value for required argument 'Device'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Pppoeinterface
	err := ctx.RegisterResource("fortios:system/pppoeinterface:Pppoeinterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPppoeinterface gets an existing Pppoeinterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPppoeinterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PppoeinterfaceState, opts ...pulumi.ResourceOption) (*Pppoeinterface, error) {
	var resource Pppoeinterface
	err := ctx.ReadResource("fortios:system/pppoeinterface:Pppoeinterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pppoeinterface resources.
type pppoeinterfaceState struct {
	// PPPoE AC name.
	AcName *string `pulumi:"acName"`
	// PPP authentication type to use. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.
	AuthType *string `pulumi:"authType"`
	// Name for the physical interface.
	Device *string `pulumi:"device"`
	// Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface. Valid values: `enable`, `disable`.
	DialOnDemand *string `pulumi:"dialOnDemand"`
	// PPPoE discovery init timeout value in (0-4294967295 sec).
	DiscRetryTimeout *int `pulumi:"discRetryTimeout"`
	// PPPoE auto disconnect after idle timeout (0-4294967295 sec).
	IdleTimeout *int `pulumi:"idleTimeout"`
	// PPPoE unnumbered IP.
	Ipunnumbered *string `pulumi:"ipunnumbered"`
	// Enable/disable IPv6 Control Protocol (IPv6CP). Valid values: `enable`, `disable`.
	Ipv6 *string `pulumi:"ipv6"`
	// Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
	LcpEchoInterval *int `pulumi:"lcpEchoInterval"`
	// Maximum missed LCP echo messages before disconnect.
	LcpMaxEchoFails *int `pulumi:"lcpMaxEchoFails"`
	// Name of the PPPoE interface.
	Name *string `pulumi:"name"`
	// PPPoE terminate timeout value in (0-4294967295 sec).
	PadtRetryTimeout *int `pulumi:"padtRetryTimeout"`
	// Enter the password.
	Password *string `pulumi:"password"`
	// Enable/disable PPPoE unnumbered negotiation. Valid values: `enable`, `disable`.
	PppoeUnnumberedNegotiate *string `pulumi:"pppoeUnnumberedNegotiate"`
	// PPPoE service name.
	ServiceName *string `pulumi:"serviceName"`
	// User name.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type PppoeinterfaceState struct {
	// PPPoE AC name.
	AcName pulumi.StringPtrInput
	// PPP authentication type to use. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.
	AuthType pulumi.StringPtrInput
	// Name for the physical interface.
	Device pulumi.StringPtrInput
	// Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface. Valid values: `enable`, `disable`.
	DialOnDemand pulumi.StringPtrInput
	// PPPoE discovery init timeout value in (0-4294967295 sec).
	DiscRetryTimeout pulumi.IntPtrInput
	// PPPoE auto disconnect after idle timeout (0-4294967295 sec).
	IdleTimeout pulumi.IntPtrInput
	// PPPoE unnumbered IP.
	Ipunnumbered pulumi.StringPtrInput
	// Enable/disable IPv6 Control Protocol (IPv6CP). Valid values: `enable`, `disable`.
	Ipv6 pulumi.StringPtrInput
	// Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
	LcpEchoInterval pulumi.IntPtrInput
	// Maximum missed LCP echo messages before disconnect.
	LcpMaxEchoFails pulumi.IntPtrInput
	// Name of the PPPoE interface.
	Name pulumi.StringPtrInput
	// PPPoE terminate timeout value in (0-4294967295 sec).
	PadtRetryTimeout pulumi.IntPtrInput
	// Enter the password.
	Password pulumi.StringPtrInput
	// Enable/disable PPPoE unnumbered negotiation. Valid values: `enable`, `disable`.
	PppoeUnnumberedNegotiate pulumi.StringPtrInput
	// PPPoE service name.
	ServiceName pulumi.StringPtrInput
	// User name.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (PppoeinterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*pppoeinterfaceState)(nil)).Elem()
}

type pppoeinterfaceArgs struct {
	// PPPoE AC name.
	AcName *string `pulumi:"acName"`
	// PPP authentication type to use. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.
	AuthType *string `pulumi:"authType"`
	// Name for the physical interface.
	Device string `pulumi:"device"`
	// Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface. Valid values: `enable`, `disable`.
	DialOnDemand *string `pulumi:"dialOnDemand"`
	// PPPoE discovery init timeout value in (0-4294967295 sec).
	DiscRetryTimeout *int `pulumi:"discRetryTimeout"`
	// PPPoE auto disconnect after idle timeout (0-4294967295 sec).
	IdleTimeout *int `pulumi:"idleTimeout"`
	// PPPoE unnumbered IP.
	Ipunnumbered *string `pulumi:"ipunnumbered"`
	// Enable/disable IPv6 Control Protocol (IPv6CP). Valid values: `enable`, `disable`.
	Ipv6 *string `pulumi:"ipv6"`
	// Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
	LcpEchoInterval *int `pulumi:"lcpEchoInterval"`
	// Maximum missed LCP echo messages before disconnect.
	LcpMaxEchoFails *int `pulumi:"lcpMaxEchoFails"`
	// Name of the PPPoE interface.
	Name *string `pulumi:"name"`
	// PPPoE terminate timeout value in (0-4294967295 sec).
	PadtRetryTimeout *int `pulumi:"padtRetryTimeout"`
	// Enter the password.
	Password *string `pulumi:"password"`
	// Enable/disable PPPoE unnumbered negotiation. Valid values: `enable`, `disable`.
	PppoeUnnumberedNegotiate *string `pulumi:"pppoeUnnumberedNegotiate"`
	// PPPoE service name.
	ServiceName *string `pulumi:"serviceName"`
	// User name.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Pppoeinterface resource.
type PppoeinterfaceArgs struct {
	// PPPoE AC name.
	AcName pulumi.StringPtrInput
	// PPP authentication type to use. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.
	AuthType pulumi.StringPtrInput
	// Name for the physical interface.
	Device pulumi.StringInput
	// Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface. Valid values: `enable`, `disable`.
	DialOnDemand pulumi.StringPtrInput
	// PPPoE discovery init timeout value in (0-4294967295 sec).
	DiscRetryTimeout pulumi.IntPtrInput
	// PPPoE auto disconnect after idle timeout (0-4294967295 sec).
	IdleTimeout pulumi.IntPtrInput
	// PPPoE unnumbered IP.
	Ipunnumbered pulumi.StringPtrInput
	// Enable/disable IPv6 Control Protocol (IPv6CP). Valid values: `enable`, `disable`.
	Ipv6 pulumi.StringPtrInput
	// Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
	LcpEchoInterval pulumi.IntPtrInput
	// Maximum missed LCP echo messages before disconnect.
	LcpMaxEchoFails pulumi.IntPtrInput
	// Name of the PPPoE interface.
	Name pulumi.StringPtrInput
	// PPPoE terminate timeout value in (0-4294967295 sec).
	PadtRetryTimeout pulumi.IntPtrInput
	// Enter the password.
	Password pulumi.StringPtrInput
	// Enable/disable PPPoE unnumbered negotiation. Valid values: `enable`, `disable`.
	PppoeUnnumberedNegotiate pulumi.StringPtrInput
	// PPPoE service name.
	ServiceName pulumi.StringPtrInput
	// User name.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (PppoeinterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pppoeinterfaceArgs)(nil)).Elem()
}

type PppoeinterfaceInput interface {
	pulumi.Input

	ToPppoeinterfaceOutput() PppoeinterfaceOutput
	ToPppoeinterfaceOutputWithContext(ctx context.Context) PppoeinterfaceOutput
}

func (*Pppoeinterface) ElementType() reflect.Type {
	return reflect.TypeOf((**Pppoeinterface)(nil)).Elem()
}

func (i *Pppoeinterface) ToPppoeinterfaceOutput() PppoeinterfaceOutput {
	return i.ToPppoeinterfaceOutputWithContext(context.Background())
}

func (i *Pppoeinterface) ToPppoeinterfaceOutputWithContext(ctx context.Context) PppoeinterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PppoeinterfaceOutput)
}

// PppoeinterfaceArrayInput is an input type that accepts PppoeinterfaceArray and PppoeinterfaceArrayOutput values.
// You can construct a concrete instance of `PppoeinterfaceArrayInput` via:
//
//	PppoeinterfaceArray{ PppoeinterfaceArgs{...} }
type PppoeinterfaceArrayInput interface {
	pulumi.Input

	ToPppoeinterfaceArrayOutput() PppoeinterfaceArrayOutput
	ToPppoeinterfaceArrayOutputWithContext(context.Context) PppoeinterfaceArrayOutput
}

type PppoeinterfaceArray []PppoeinterfaceInput

func (PppoeinterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pppoeinterface)(nil)).Elem()
}

func (i PppoeinterfaceArray) ToPppoeinterfaceArrayOutput() PppoeinterfaceArrayOutput {
	return i.ToPppoeinterfaceArrayOutputWithContext(context.Background())
}

func (i PppoeinterfaceArray) ToPppoeinterfaceArrayOutputWithContext(ctx context.Context) PppoeinterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PppoeinterfaceArrayOutput)
}

// PppoeinterfaceMapInput is an input type that accepts PppoeinterfaceMap and PppoeinterfaceMapOutput values.
// You can construct a concrete instance of `PppoeinterfaceMapInput` via:
//
//	PppoeinterfaceMap{ "key": PppoeinterfaceArgs{...} }
type PppoeinterfaceMapInput interface {
	pulumi.Input

	ToPppoeinterfaceMapOutput() PppoeinterfaceMapOutput
	ToPppoeinterfaceMapOutputWithContext(context.Context) PppoeinterfaceMapOutput
}

type PppoeinterfaceMap map[string]PppoeinterfaceInput

func (PppoeinterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pppoeinterface)(nil)).Elem()
}

func (i PppoeinterfaceMap) ToPppoeinterfaceMapOutput() PppoeinterfaceMapOutput {
	return i.ToPppoeinterfaceMapOutputWithContext(context.Background())
}

func (i PppoeinterfaceMap) ToPppoeinterfaceMapOutputWithContext(ctx context.Context) PppoeinterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PppoeinterfaceMapOutput)
}

type PppoeinterfaceOutput struct{ *pulumi.OutputState }

func (PppoeinterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pppoeinterface)(nil)).Elem()
}

func (o PppoeinterfaceOutput) ToPppoeinterfaceOutput() PppoeinterfaceOutput {
	return o
}

func (o PppoeinterfaceOutput) ToPppoeinterfaceOutputWithContext(ctx context.Context) PppoeinterfaceOutput {
	return o
}

// PPPoE AC name.
func (o PppoeinterfaceOutput) AcName() pulumi.StringOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.StringOutput { return v.AcName }).(pulumi.StringOutput)
}

// PPP authentication type to use. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.
func (o PppoeinterfaceOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.StringOutput { return v.AuthType }).(pulumi.StringOutput)
}

// Name for the physical interface.
func (o PppoeinterfaceOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.StringOutput { return v.Device }).(pulumi.StringOutput)
}

// Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface. Valid values: `enable`, `disable`.
func (o PppoeinterfaceOutput) DialOnDemand() pulumi.StringOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.StringOutput { return v.DialOnDemand }).(pulumi.StringOutput)
}

// PPPoE discovery init timeout value in (0-4294967295 sec).
func (o PppoeinterfaceOutput) DiscRetryTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.IntOutput { return v.DiscRetryTimeout }).(pulumi.IntOutput)
}

// PPPoE auto disconnect after idle timeout (0-4294967295 sec).
func (o PppoeinterfaceOutput) IdleTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.IntOutput { return v.IdleTimeout }).(pulumi.IntOutput)
}

// PPPoE unnumbered IP.
func (o PppoeinterfaceOutput) Ipunnumbered() pulumi.StringOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.StringOutput { return v.Ipunnumbered }).(pulumi.StringOutput)
}

// Enable/disable IPv6 Control Protocol (IPv6CP). Valid values: `enable`, `disable`.
func (o PppoeinterfaceOutput) Ipv6() pulumi.StringOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.StringOutput { return v.Ipv6 }).(pulumi.StringOutput)
}

// Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
func (o PppoeinterfaceOutput) LcpEchoInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.IntOutput { return v.LcpEchoInterval }).(pulumi.IntOutput)
}

// Maximum missed LCP echo messages before disconnect.
func (o PppoeinterfaceOutput) LcpMaxEchoFails() pulumi.IntOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.IntOutput { return v.LcpMaxEchoFails }).(pulumi.IntOutput)
}

// Name of the PPPoE interface.
func (o PppoeinterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// PPPoE terminate timeout value in (0-4294967295 sec).
func (o PppoeinterfaceOutput) PadtRetryTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.IntOutput { return v.PadtRetryTimeout }).(pulumi.IntOutput)
}

// Enter the password.
func (o PppoeinterfaceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Enable/disable PPPoE unnumbered negotiation. Valid values: `enable`, `disable`.
func (o PppoeinterfaceOutput) PppoeUnnumberedNegotiate() pulumi.StringOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.StringOutput { return v.PppoeUnnumberedNegotiate }).(pulumi.StringOutput)
}

// PPPoE service name.
func (o PppoeinterfaceOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// User name.
func (o PppoeinterfaceOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o PppoeinterfaceOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Pppoeinterface) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type PppoeinterfaceArrayOutput struct{ *pulumi.OutputState }

func (PppoeinterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pppoeinterface)(nil)).Elem()
}

func (o PppoeinterfaceArrayOutput) ToPppoeinterfaceArrayOutput() PppoeinterfaceArrayOutput {
	return o
}

func (o PppoeinterfaceArrayOutput) ToPppoeinterfaceArrayOutputWithContext(ctx context.Context) PppoeinterfaceArrayOutput {
	return o
}

func (o PppoeinterfaceArrayOutput) Index(i pulumi.IntInput) PppoeinterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Pppoeinterface {
		return vs[0].([]*Pppoeinterface)[vs[1].(int)]
	}).(PppoeinterfaceOutput)
}

type PppoeinterfaceMapOutput struct{ *pulumi.OutputState }

func (PppoeinterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pppoeinterface)(nil)).Elem()
}

func (o PppoeinterfaceMapOutput) ToPppoeinterfaceMapOutput() PppoeinterfaceMapOutput {
	return o
}

func (o PppoeinterfaceMapOutput) ToPppoeinterfaceMapOutputWithContext(ctx context.Context) PppoeinterfaceMapOutput {
	return o
}

func (o PppoeinterfaceMapOutput) MapIndex(k pulumi.StringInput) PppoeinterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Pppoeinterface {
		return vs[0].(map[string]*Pppoeinterface)[vs[1].(string)]
	}).(PppoeinterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PppoeinterfaceInput)(nil)).Elem(), &Pppoeinterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*PppoeinterfaceArrayInput)(nil)).Elem(), PppoeinterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PppoeinterfaceMapInput)(nil)).Elem(), PppoeinterfaceMap{})
	pulumi.RegisterOutputType(PppoeinterfaceOutput{})
	pulumi.RegisterOutputType(PppoeinterfaceArrayOutput{})
	pulumi.RegisterOutputType(PppoeinterfaceMapOutput{})
}
