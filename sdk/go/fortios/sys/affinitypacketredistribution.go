// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure packet redistribution.
//
// ## Import
//
// # System AffinityPacketRedistribution can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:sys/affinitypacketredistribution:Affinitypacketredistribution labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:sys/affinitypacketredistribution:Affinitypacketredistribution labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Affinitypacketredistribution struct {
	pulumi.CustomResourceState

	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask pulumi.StringOutput `pulumi:"affinityCpumask"`
	// ID of the packet redistribution setting.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Physical interface name on which to perform packet redistribution.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
	Rxqid pulumi.IntOutput `pulumi:"rxqid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewAffinitypacketredistribution registers a new resource with the given unique name, arguments, and options.
func NewAffinitypacketredistribution(ctx *pulumi.Context,
	name string, args *AffinitypacketredistributionArgs, opts ...pulumi.ResourceOption) (*Affinitypacketredistribution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AffinityCpumask == nil {
		return nil, errors.New("invalid value for required argument 'AffinityCpumask'")
	}
	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.Rxqid == nil {
		return nil, errors.New("invalid value for required argument 'Rxqid'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Affinitypacketredistribution
	err := ctx.RegisterResource("fortios:sys/affinitypacketredistribution:Affinitypacketredistribution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAffinitypacketredistribution gets an existing Affinitypacketredistribution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAffinitypacketredistribution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AffinitypacketredistributionState, opts ...pulumi.ResourceOption) (*Affinitypacketredistribution, error) {
	var resource Affinitypacketredistribution
	err := ctx.ReadResource("fortios:sys/affinitypacketredistribution:Affinitypacketredistribution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Affinitypacketredistribution resources.
type affinitypacketredistributionState struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask *string `pulumi:"affinityCpumask"`
	// ID of the packet redistribution setting.
	Fosid *int `pulumi:"fosid"`
	// Physical interface name on which to perform packet redistribution.
	Interface *string `pulumi:"interface"`
	// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
	Rxqid *int `pulumi:"rxqid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type AffinitypacketredistributionState struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask pulumi.StringPtrInput
	// ID of the packet redistribution setting.
	Fosid pulumi.IntPtrInput
	// Physical interface name on which to perform packet redistribution.
	Interface pulumi.StringPtrInput
	// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
	Rxqid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AffinitypacketredistributionState) ElementType() reflect.Type {
	return reflect.TypeOf((*affinitypacketredistributionState)(nil)).Elem()
}

type affinitypacketredistributionArgs struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask string `pulumi:"affinityCpumask"`
	// ID of the packet redistribution setting.
	Fosid int `pulumi:"fosid"`
	// Physical interface name on which to perform packet redistribution.
	Interface string `pulumi:"interface"`
	// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
	Rxqid int `pulumi:"rxqid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Affinitypacketredistribution resource.
type AffinitypacketredistributionArgs struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask pulumi.StringInput
	// ID of the packet redistribution setting.
	Fosid pulumi.IntInput
	// Physical interface name on which to perform packet redistribution.
	Interface pulumi.StringInput
	// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
	Rxqid pulumi.IntInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AffinitypacketredistributionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*affinitypacketredistributionArgs)(nil)).Elem()
}

type AffinitypacketredistributionInput interface {
	pulumi.Input

	ToAffinitypacketredistributionOutput() AffinitypacketredistributionOutput
	ToAffinitypacketredistributionOutputWithContext(ctx context.Context) AffinitypacketredistributionOutput
}

func (*Affinitypacketredistribution) ElementType() reflect.Type {
	return reflect.TypeOf((**Affinitypacketredistribution)(nil)).Elem()
}

func (i *Affinitypacketredistribution) ToAffinitypacketredistributionOutput() AffinitypacketredistributionOutput {
	return i.ToAffinitypacketredistributionOutputWithContext(context.Background())
}

func (i *Affinitypacketredistribution) ToAffinitypacketredistributionOutputWithContext(ctx context.Context) AffinitypacketredistributionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AffinitypacketredistributionOutput)
}

// AffinitypacketredistributionArrayInput is an input type that accepts AffinitypacketredistributionArray and AffinitypacketredistributionArrayOutput values.
// You can construct a concrete instance of `AffinitypacketredistributionArrayInput` via:
//
//	AffinitypacketredistributionArray{ AffinitypacketredistributionArgs{...} }
type AffinitypacketredistributionArrayInput interface {
	pulumi.Input

	ToAffinitypacketredistributionArrayOutput() AffinitypacketredistributionArrayOutput
	ToAffinitypacketredistributionArrayOutputWithContext(context.Context) AffinitypacketredistributionArrayOutput
}

type AffinitypacketredistributionArray []AffinitypacketredistributionInput

func (AffinitypacketredistributionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Affinitypacketredistribution)(nil)).Elem()
}

func (i AffinitypacketredistributionArray) ToAffinitypacketredistributionArrayOutput() AffinitypacketredistributionArrayOutput {
	return i.ToAffinitypacketredistributionArrayOutputWithContext(context.Background())
}

func (i AffinitypacketredistributionArray) ToAffinitypacketredistributionArrayOutputWithContext(ctx context.Context) AffinitypacketredistributionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AffinitypacketredistributionArrayOutput)
}

// AffinitypacketredistributionMapInput is an input type that accepts AffinitypacketredistributionMap and AffinitypacketredistributionMapOutput values.
// You can construct a concrete instance of `AffinitypacketredistributionMapInput` via:
//
//	AffinitypacketredistributionMap{ "key": AffinitypacketredistributionArgs{...} }
type AffinitypacketredistributionMapInput interface {
	pulumi.Input

	ToAffinitypacketredistributionMapOutput() AffinitypacketredistributionMapOutput
	ToAffinitypacketredistributionMapOutputWithContext(context.Context) AffinitypacketredistributionMapOutput
}

type AffinitypacketredistributionMap map[string]AffinitypacketredistributionInput

func (AffinitypacketredistributionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Affinitypacketredistribution)(nil)).Elem()
}

func (i AffinitypacketredistributionMap) ToAffinitypacketredistributionMapOutput() AffinitypacketredistributionMapOutput {
	return i.ToAffinitypacketredistributionMapOutputWithContext(context.Background())
}

func (i AffinitypacketredistributionMap) ToAffinitypacketredistributionMapOutputWithContext(ctx context.Context) AffinitypacketredistributionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AffinitypacketredistributionMapOutput)
}

type AffinitypacketredistributionOutput struct{ *pulumi.OutputState }

func (AffinitypacketredistributionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Affinitypacketredistribution)(nil)).Elem()
}

func (o AffinitypacketredistributionOutput) ToAffinitypacketredistributionOutput() AffinitypacketredistributionOutput {
	return o
}

func (o AffinitypacketredistributionOutput) ToAffinitypacketredistributionOutputWithContext(ctx context.Context) AffinitypacketredistributionOutput {
	return o
}

// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
func (o AffinitypacketredistributionOutput) AffinityCpumask() pulumi.StringOutput {
	return o.ApplyT(func(v *Affinitypacketredistribution) pulumi.StringOutput { return v.AffinityCpumask }).(pulumi.StringOutput)
}

// ID of the packet redistribution setting.
func (o AffinitypacketredistributionOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Affinitypacketredistribution) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Physical interface name on which to perform packet redistribution.
func (o AffinitypacketredistributionOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Affinitypacketredistribution) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
func (o AffinitypacketredistributionOutput) Rxqid() pulumi.IntOutput {
	return o.ApplyT(func(v *Affinitypacketredistribution) pulumi.IntOutput { return v.Rxqid }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AffinitypacketredistributionOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Affinitypacketredistribution) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AffinitypacketredistributionArrayOutput struct{ *pulumi.OutputState }

func (AffinitypacketredistributionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Affinitypacketredistribution)(nil)).Elem()
}

func (o AffinitypacketredistributionArrayOutput) ToAffinitypacketredistributionArrayOutput() AffinitypacketredistributionArrayOutput {
	return o
}

func (o AffinitypacketredistributionArrayOutput) ToAffinitypacketredistributionArrayOutputWithContext(ctx context.Context) AffinitypacketredistributionArrayOutput {
	return o
}

func (o AffinitypacketredistributionArrayOutput) Index(i pulumi.IntInput) AffinitypacketredistributionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Affinitypacketredistribution {
		return vs[0].([]*Affinitypacketredistribution)[vs[1].(int)]
	}).(AffinitypacketredistributionOutput)
}

type AffinitypacketredistributionMapOutput struct{ *pulumi.OutputState }

func (AffinitypacketredistributionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Affinitypacketredistribution)(nil)).Elem()
}

func (o AffinitypacketredistributionMapOutput) ToAffinitypacketredistributionMapOutput() AffinitypacketredistributionMapOutput {
	return o
}

func (o AffinitypacketredistributionMapOutput) ToAffinitypacketredistributionMapOutputWithContext(ctx context.Context) AffinitypacketredistributionMapOutput {
	return o
}

func (o AffinitypacketredistributionMapOutput) MapIndex(k pulumi.StringInput) AffinitypacketredistributionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Affinitypacketredistribution {
		return vs[0].(map[string]*Affinitypacketredistribution)[vs[1].(string)]
	}).(AffinitypacketredistributionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AffinitypacketredistributionInput)(nil)).Elem(), &Affinitypacketredistribution{})
	pulumi.RegisterInputType(reflect.TypeOf((*AffinitypacketredistributionArrayInput)(nil)).Elem(), AffinitypacketredistributionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AffinitypacketredistributionMapInput)(nil)).Elem(), AffinitypacketredistributionMap{})
	pulumi.RegisterOutputType(AffinitypacketredistributionOutput{})
	pulumi.RegisterOutputType(AffinitypacketredistributionArrayOutput{})
	pulumi.RegisterOutputType(AffinitypacketredistributionMapOutput{})
}
