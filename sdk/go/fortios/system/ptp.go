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

// Configure system PTP information.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewPtp(ctx, "trname", &system.PtpArgs{
//				DelayMechanism:  pulumi.String("E2E"),
//				Interface:       pulumi.String("port3"),
//				Mode:            pulumi.String("multicast"),
//				RequestInterval: pulumi.Int(1),
//				Status:          pulumi.String("enable"),
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
// System Ptp can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/ptp:Ptp labelname SystemPtp
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/ptp:Ptp labelname SystemPtp
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Ptp struct {
	pulumi.CustomResourceState

	// End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
	DelayMechanism pulumi.StringOutput `pulumi:"delayMechanism"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// PTP slave will reply through this interface.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
	RequestInterval pulumi.IntOutput `pulumi:"requestInterval"`
	// FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `serverInterface` block is documented below.
	ServerInterfaces PtpServerInterfaceArrayOutput `pulumi:"serverInterfaces"`
	// Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
	ServerMode pulumi.StringOutput `pulumi:"serverMode"`
	// Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewPtp registers a new resource with the given unique name, arguments, and options.
func NewPtp(ctx *pulumi.Context,
	name string, args *PtpArgs, opts ...pulumi.ResourceOption) (*Ptp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ptp
	err := ctx.RegisterResource("fortios:system/ptp:Ptp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPtp gets an existing Ptp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPtp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PtpState, opts ...pulumi.ResourceOption) (*Ptp, error) {
	var resource Ptp
	err := ctx.ReadResource("fortios:system/ptp:Ptp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ptp resources.
type ptpState struct {
	// End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
	DelayMechanism *string `pulumi:"delayMechanism"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// PTP slave will reply through this interface.
	Interface *string `pulumi:"interface"`
	// Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
	Mode *string `pulumi:"mode"`
	// The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
	RequestInterval *int `pulumi:"requestInterval"`
	// FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `serverInterface` block is documented below.
	ServerInterfaces []PtpServerInterface `pulumi:"serverInterfaces"`
	// Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
	ServerMode *string `pulumi:"serverMode"`
	// Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type PtpState struct {
	// End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
	DelayMechanism pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// PTP slave will reply through this interface.
	Interface pulumi.StringPtrInput
	// Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
	Mode pulumi.StringPtrInput
	// The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
	RequestInterval pulumi.IntPtrInput
	// FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `serverInterface` block is documented below.
	ServerInterfaces PtpServerInterfaceArrayInput
	// Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
	ServerMode pulumi.StringPtrInput
	// Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (PtpState) ElementType() reflect.Type {
	return reflect.TypeOf((*ptpState)(nil)).Elem()
}

type ptpArgs struct {
	// End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
	DelayMechanism *string `pulumi:"delayMechanism"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// PTP slave will reply through this interface.
	Interface string `pulumi:"interface"`
	// Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
	Mode *string `pulumi:"mode"`
	// The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
	RequestInterval *int `pulumi:"requestInterval"`
	// FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `serverInterface` block is documented below.
	ServerInterfaces []PtpServerInterface `pulumi:"serverInterfaces"`
	// Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
	ServerMode *string `pulumi:"serverMode"`
	// Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Ptp resource.
type PtpArgs struct {
	// End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
	DelayMechanism pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// PTP slave will reply through this interface.
	Interface pulumi.StringInput
	// Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
	Mode pulumi.StringPtrInput
	// The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
	RequestInterval pulumi.IntPtrInput
	// FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `serverInterface` block is documented below.
	ServerInterfaces PtpServerInterfaceArrayInput
	// Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
	ServerMode pulumi.StringPtrInput
	// Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (PtpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ptpArgs)(nil)).Elem()
}

type PtpInput interface {
	pulumi.Input

	ToPtpOutput() PtpOutput
	ToPtpOutputWithContext(ctx context.Context) PtpOutput
}

func (*Ptp) ElementType() reflect.Type {
	return reflect.TypeOf((**Ptp)(nil)).Elem()
}

func (i *Ptp) ToPtpOutput() PtpOutput {
	return i.ToPtpOutputWithContext(context.Background())
}

func (i *Ptp) ToPtpOutputWithContext(ctx context.Context) PtpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PtpOutput)
}

// PtpArrayInput is an input type that accepts PtpArray and PtpArrayOutput values.
// You can construct a concrete instance of `PtpArrayInput` via:
//
//	PtpArray{ PtpArgs{...} }
type PtpArrayInput interface {
	pulumi.Input

	ToPtpArrayOutput() PtpArrayOutput
	ToPtpArrayOutputWithContext(context.Context) PtpArrayOutput
}

type PtpArray []PtpInput

func (PtpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ptp)(nil)).Elem()
}

func (i PtpArray) ToPtpArrayOutput() PtpArrayOutput {
	return i.ToPtpArrayOutputWithContext(context.Background())
}

func (i PtpArray) ToPtpArrayOutputWithContext(ctx context.Context) PtpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PtpArrayOutput)
}

// PtpMapInput is an input type that accepts PtpMap and PtpMapOutput values.
// You can construct a concrete instance of `PtpMapInput` via:
//
//	PtpMap{ "key": PtpArgs{...} }
type PtpMapInput interface {
	pulumi.Input

	ToPtpMapOutput() PtpMapOutput
	ToPtpMapOutputWithContext(context.Context) PtpMapOutput
}

type PtpMap map[string]PtpInput

func (PtpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ptp)(nil)).Elem()
}

func (i PtpMap) ToPtpMapOutput() PtpMapOutput {
	return i.ToPtpMapOutputWithContext(context.Background())
}

func (i PtpMap) ToPtpMapOutputWithContext(ctx context.Context) PtpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PtpMapOutput)
}

type PtpOutput struct{ *pulumi.OutputState }

func (PtpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ptp)(nil)).Elem()
}

func (o PtpOutput) ToPtpOutput() PtpOutput {
	return o
}

func (o PtpOutput) ToPtpOutputWithContext(ctx context.Context) PtpOutput {
	return o
}

// End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
func (o PtpOutput) DelayMechanism() pulumi.StringOutput {
	return o.ApplyT(func(v *Ptp) pulumi.StringOutput { return v.DelayMechanism }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o PtpOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ptp) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// PTP slave will reply through this interface.
func (o PtpOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Ptp) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
func (o PtpOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *Ptp) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
func (o PtpOutput) RequestInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Ptp) pulumi.IntOutput { return v.RequestInterval }).(pulumi.IntOutput)
}

// FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `serverInterface` block is documented below.
func (o PtpOutput) ServerInterfaces() PtpServerInterfaceArrayOutput {
	return o.ApplyT(func(v *Ptp) PtpServerInterfaceArrayOutput { return v.ServerInterfaces }).(PtpServerInterfaceArrayOutput)
}

// Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
func (o PtpOutput) ServerMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Ptp) pulumi.StringOutput { return v.ServerMode }).(pulumi.StringOutput)
}

// Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
func (o PtpOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Ptp) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o PtpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ptp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type PtpArrayOutput struct{ *pulumi.OutputState }

func (PtpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ptp)(nil)).Elem()
}

func (o PtpArrayOutput) ToPtpArrayOutput() PtpArrayOutput {
	return o
}

func (o PtpArrayOutput) ToPtpArrayOutputWithContext(ctx context.Context) PtpArrayOutput {
	return o
}

func (o PtpArrayOutput) Index(i pulumi.IntInput) PtpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ptp {
		return vs[0].([]*Ptp)[vs[1].(int)]
	}).(PtpOutput)
}

type PtpMapOutput struct{ *pulumi.OutputState }

func (PtpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ptp)(nil)).Elem()
}

func (o PtpMapOutput) ToPtpMapOutput() PtpMapOutput {
	return o
}

func (o PtpMapOutput) ToPtpMapOutputWithContext(ctx context.Context) PtpMapOutput {
	return o
}

func (o PtpMapOutput) MapIndex(k pulumi.StringInput) PtpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ptp {
		return vs[0].(map[string]*Ptp)[vs[1].(string)]
	}).(PtpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PtpInput)(nil)).Elem(), &Ptp{})
	pulumi.RegisterInputType(reflect.TypeOf((*PtpArrayInput)(nil)).Elem(), PtpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PtpMapInput)(nil)).Elem(), PtpMap{})
	pulumi.RegisterOutputType(PtpOutput{})
	pulumi.RegisterOutputType(PtpArrayOutput{})
	pulumi.RegisterOutputType(PtpMapOutput{})
}
