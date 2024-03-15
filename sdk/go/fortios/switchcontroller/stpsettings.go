// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiSwitch spanning tree protocol (STP).
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/switchcontroller"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := switchcontroller.NewStpsettings(ctx, "trname", &switchcontroller.StpsettingsArgs{
//				ForwardTime:  pulumi.Int(15),
//				HelloTime:    pulumi.Int(2),
//				MaxAge:       pulumi.Int(20),
//				MaxHops:      pulumi.Int(20),
//				PendingTimer: pulumi.Int(4),
//				Revision:     pulumi.Int(0),
//				Status:       pulumi.String("enable"),
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
// SwitchController StpSettings can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/stpsettings:Stpsettings labelname SwitchControllerStpSettings
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/stpsettings:Stpsettings labelname SwitchControllerStpSettings
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Stpsettings struct {
	pulumi.CustomResourceState

	// Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
	ForwardTime pulumi.IntOutput `pulumi:"forwardTime"`
	// Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
	HelloTime pulumi.IntOutput `pulumi:"helloTime"`
	// Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
	MaxAge pulumi.IntOutput `pulumi:"maxAge"`
	// Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
	MaxHops pulumi.IntOutput `pulumi:"maxHops"`
	// Name of global STP settings configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Pending time (1 - 15 sec, default = 4).
	PendingTimer pulumi.IntOutput `pulumi:"pendingTimer"`
	// STP revision number (0 - 65535).
	Revision pulumi.IntOutput `pulumi:"revision"`
	// Enable/disable STP. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewStpsettings registers a new resource with the given unique name, arguments, and options.
func NewStpsettings(ctx *pulumi.Context,
	name string, args *StpsettingsArgs, opts ...pulumi.ResourceOption) (*Stpsettings, error) {
	if args == nil {
		args = &StpsettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Stpsettings
	err := ctx.RegisterResource("fortios:switchcontroller/stpsettings:Stpsettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStpsettings gets an existing Stpsettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStpsettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StpsettingsState, opts ...pulumi.ResourceOption) (*Stpsettings, error) {
	var resource Stpsettings
	err := ctx.ReadResource("fortios:switchcontroller/stpsettings:Stpsettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stpsettings resources.
type stpsettingsState struct {
	// Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
	ForwardTime *int `pulumi:"forwardTime"`
	// Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
	HelloTime *int `pulumi:"helloTime"`
	// Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
	MaxAge *int `pulumi:"maxAge"`
	// Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
	MaxHops *int `pulumi:"maxHops"`
	// Name of global STP settings configuration.
	Name *string `pulumi:"name"`
	// Pending time (1 - 15 sec, default = 4).
	PendingTimer *int `pulumi:"pendingTimer"`
	// STP revision number (0 - 65535).
	Revision *int `pulumi:"revision"`
	// Enable/disable STP. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type StpsettingsState struct {
	// Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
	ForwardTime pulumi.IntPtrInput
	// Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
	HelloTime pulumi.IntPtrInput
	// Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
	MaxAge pulumi.IntPtrInput
	// Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
	MaxHops pulumi.IntPtrInput
	// Name of global STP settings configuration.
	Name pulumi.StringPtrInput
	// Pending time (1 - 15 sec, default = 4).
	PendingTimer pulumi.IntPtrInput
	// STP revision number (0 - 65535).
	Revision pulumi.IntPtrInput
	// Enable/disable STP. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (StpsettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*stpsettingsState)(nil)).Elem()
}

type stpsettingsArgs struct {
	// Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
	ForwardTime *int `pulumi:"forwardTime"`
	// Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
	HelloTime *int `pulumi:"helloTime"`
	// Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
	MaxAge *int `pulumi:"maxAge"`
	// Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
	MaxHops *int `pulumi:"maxHops"`
	// Name of global STP settings configuration.
	Name *string `pulumi:"name"`
	// Pending time (1 - 15 sec, default = 4).
	PendingTimer *int `pulumi:"pendingTimer"`
	// STP revision number (0 - 65535).
	Revision *int `pulumi:"revision"`
	// Enable/disable STP. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Stpsettings resource.
type StpsettingsArgs struct {
	// Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
	ForwardTime pulumi.IntPtrInput
	// Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
	HelloTime pulumi.IntPtrInput
	// Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
	MaxAge pulumi.IntPtrInput
	// Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
	MaxHops pulumi.IntPtrInput
	// Name of global STP settings configuration.
	Name pulumi.StringPtrInput
	// Pending time (1 - 15 sec, default = 4).
	PendingTimer pulumi.IntPtrInput
	// STP revision number (0 - 65535).
	Revision pulumi.IntPtrInput
	// Enable/disable STP. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (StpsettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stpsettingsArgs)(nil)).Elem()
}

type StpsettingsInput interface {
	pulumi.Input

	ToStpsettingsOutput() StpsettingsOutput
	ToStpsettingsOutputWithContext(ctx context.Context) StpsettingsOutput
}

func (*Stpsettings) ElementType() reflect.Type {
	return reflect.TypeOf((**Stpsettings)(nil)).Elem()
}

func (i *Stpsettings) ToStpsettingsOutput() StpsettingsOutput {
	return i.ToStpsettingsOutputWithContext(context.Background())
}

func (i *Stpsettings) ToStpsettingsOutputWithContext(ctx context.Context) StpsettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StpsettingsOutput)
}

// StpsettingsArrayInput is an input type that accepts StpsettingsArray and StpsettingsArrayOutput values.
// You can construct a concrete instance of `StpsettingsArrayInput` via:
//
//	StpsettingsArray{ StpsettingsArgs{...} }
type StpsettingsArrayInput interface {
	pulumi.Input

	ToStpsettingsArrayOutput() StpsettingsArrayOutput
	ToStpsettingsArrayOutputWithContext(context.Context) StpsettingsArrayOutput
}

type StpsettingsArray []StpsettingsInput

func (StpsettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Stpsettings)(nil)).Elem()
}

func (i StpsettingsArray) ToStpsettingsArrayOutput() StpsettingsArrayOutput {
	return i.ToStpsettingsArrayOutputWithContext(context.Background())
}

func (i StpsettingsArray) ToStpsettingsArrayOutputWithContext(ctx context.Context) StpsettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StpsettingsArrayOutput)
}

// StpsettingsMapInput is an input type that accepts StpsettingsMap and StpsettingsMapOutput values.
// You can construct a concrete instance of `StpsettingsMapInput` via:
//
//	StpsettingsMap{ "key": StpsettingsArgs{...} }
type StpsettingsMapInput interface {
	pulumi.Input

	ToStpsettingsMapOutput() StpsettingsMapOutput
	ToStpsettingsMapOutputWithContext(context.Context) StpsettingsMapOutput
}

type StpsettingsMap map[string]StpsettingsInput

func (StpsettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Stpsettings)(nil)).Elem()
}

func (i StpsettingsMap) ToStpsettingsMapOutput() StpsettingsMapOutput {
	return i.ToStpsettingsMapOutputWithContext(context.Background())
}

func (i StpsettingsMap) ToStpsettingsMapOutputWithContext(ctx context.Context) StpsettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StpsettingsMapOutput)
}

type StpsettingsOutput struct{ *pulumi.OutputState }

func (StpsettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stpsettings)(nil)).Elem()
}

func (o StpsettingsOutput) ToStpsettingsOutput() StpsettingsOutput {
	return o
}

func (o StpsettingsOutput) ToStpsettingsOutputWithContext(ctx context.Context) StpsettingsOutput {
	return o
}

// Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
func (o StpsettingsOutput) ForwardTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Stpsettings) pulumi.IntOutput { return v.ForwardTime }).(pulumi.IntOutput)
}

// Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
func (o StpsettingsOutput) HelloTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Stpsettings) pulumi.IntOutput { return v.HelloTime }).(pulumi.IntOutput)
}

// Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
func (o StpsettingsOutput) MaxAge() pulumi.IntOutput {
	return o.ApplyT(func(v *Stpsettings) pulumi.IntOutput { return v.MaxAge }).(pulumi.IntOutput)
}

// Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
func (o StpsettingsOutput) MaxHops() pulumi.IntOutput {
	return o.ApplyT(func(v *Stpsettings) pulumi.IntOutput { return v.MaxHops }).(pulumi.IntOutput)
}

// Name of global STP settings configuration.
func (o StpsettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Stpsettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Pending time (1 - 15 sec, default = 4).
func (o StpsettingsOutput) PendingTimer() pulumi.IntOutput {
	return o.ApplyT(func(v *Stpsettings) pulumi.IntOutput { return v.PendingTimer }).(pulumi.IntOutput)
}

// STP revision number (0 - 65535).
func (o StpsettingsOutput) Revision() pulumi.IntOutput {
	return o.ApplyT(func(v *Stpsettings) pulumi.IntOutput { return v.Revision }).(pulumi.IntOutput)
}

// Enable/disable STP. Valid values: `enable`, `disable`.
func (o StpsettingsOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Stpsettings) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o StpsettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stpsettings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type StpsettingsArrayOutput struct{ *pulumi.OutputState }

func (StpsettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Stpsettings)(nil)).Elem()
}

func (o StpsettingsArrayOutput) ToStpsettingsArrayOutput() StpsettingsArrayOutput {
	return o
}

func (o StpsettingsArrayOutput) ToStpsettingsArrayOutputWithContext(ctx context.Context) StpsettingsArrayOutput {
	return o
}

func (o StpsettingsArrayOutput) Index(i pulumi.IntInput) StpsettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Stpsettings {
		return vs[0].([]*Stpsettings)[vs[1].(int)]
	}).(StpsettingsOutput)
}

type StpsettingsMapOutput struct{ *pulumi.OutputState }

func (StpsettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Stpsettings)(nil)).Elem()
}

func (o StpsettingsMapOutput) ToStpsettingsMapOutput() StpsettingsMapOutput {
	return o
}

func (o StpsettingsMapOutput) ToStpsettingsMapOutputWithContext(ctx context.Context) StpsettingsMapOutput {
	return o
}

func (o StpsettingsMapOutput) MapIndex(k pulumi.StringInput) StpsettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Stpsettings {
		return vs[0].(map[string]*Stpsettings)[vs[1].(string)]
	}).(StpsettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StpsettingsInput)(nil)).Elem(), &Stpsettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*StpsettingsArrayInput)(nil)).Elem(), StpsettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StpsettingsMapInput)(nil)).Elem(), StpsettingsMap{})
	pulumi.RegisterOutputType(StpsettingsOutput{})
	pulumi.RegisterOutputType(StpsettingsArrayOutput{})
	pulumi.RegisterOutputType(StpsettingsMapOutput{})
}