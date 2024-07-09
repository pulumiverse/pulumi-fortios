// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package qos

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiSwitch QoS policy.
//
// ## Import
//
// SwitchControllerQos QosPolicy can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/qos/qospolicy:Qospolicy labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/qos/qospolicy:Qospolicy labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Qospolicy struct {
	pulumi.CustomResourceState

	// Default cos queue for untagged packets.
	DefaultCos pulumi.IntOutput `pulumi:"defaultCos"`
	// QoS policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// QoS egress queue policy.
	QueuePolicy pulumi.StringOutput `pulumi:"queuePolicy"`
	// QoS trust 802.1p map.
	TrustDot1pMap pulumi.StringOutput `pulumi:"trustDot1pMap"`
	// QoS trust ip dscp map.
	TrustIpDscpMap pulumi.StringOutput `pulumi:"trustIpDscpMap"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewQospolicy registers a new resource with the given unique name, arguments, and options.
func NewQospolicy(ctx *pulumi.Context,
	name string, args *QospolicyArgs, opts ...pulumi.ResourceOption) (*Qospolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultCos == nil {
		return nil, errors.New("invalid value for required argument 'DefaultCos'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Qospolicy
	err := ctx.RegisterResource("fortios:switchcontroller/qos/qospolicy:Qospolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQospolicy gets an existing Qospolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQospolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QospolicyState, opts ...pulumi.ResourceOption) (*Qospolicy, error) {
	var resource Qospolicy
	err := ctx.ReadResource("fortios:switchcontroller/qos/qospolicy:Qospolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Qospolicy resources.
type qospolicyState struct {
	// Default cos queue for untagged packets.
	DefaultCos *int `pulumi:"defaultCos"`
	// QoS policy name.
	Name *string `pulumi:"name"`
	// QoS egress queue policy.
	QueuePolicy *string `pulumi:"queuePolicy"`
	// QoS trust 802.1p map.
	TrustDot1pMap *string `pulumi:"trustDot1pMap"`
	// QoS trust ip dscp map.
	TrustIpDscpMap *string `pulumi:"trustIpDscpMap"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type QospolicyState struct {
	// Default cos queue for untagged packets.
	DefaultCos pulumi.IntPtrInput
	// QoS policy name.
	Name pulumi.StringPtrInput
	// QoS egress queue policy.
	QueuePolicy pulumi.StringPtrInput
	// QoS trust 802.1p map.
	TrustDot1pMap pulumi.StringPtrInput
	// QoS trust ip dscp map.
	TrustIpDscpMap pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (QospolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*qospolicyState)(nil)).Elem()
}

type qospolicyArgs struct {
	// Default cos queue for untagged packets.
	DefaultCos int `pulumi:"defaultCos"`
	// QoS policy name.
	Name *string `pulumi:"name"`
	// QoS egress queue policy.
	QueuePolicy *string `pulumi:"queuePolicy"`
	// QoS trust 802.1p map.
	TrustDot1pMap *string `pulumi:"trustDot1pMap"`
	// QoS trust ip dscp map.
	TrustIpDscpMap *string `pulumi:"trustIpDscpMap"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Qospolicy resource.
type QospolicyArgs struct {
	// Default cos queue for untagged packets.
	DefaultCos pulumi.IntInput
	// QoS policy name.
	Name pulumi.StringPtrInput
	// QoS egress queue policy.
	QueuePolicy pulumi.StringPtrInput
	// QoS trust 802.1p map.
	TrustDot1pMap pulumi.StringPtrInput
	// QoS trust ip dscp map.
	TrustIpDscpMap pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (QospolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*qospolicyArgs)(nil)).Elem()
}

type QospolicyInput interface {
	pulumi.Input

	ToQospolicyOutput() QospolicyOutput
	ToQospolicyOutputWithContext(ctx context.Context) QospolicyOutput
}

func (*Qospolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**Qospolicy)(nil)).Elem()
}

func (i *Qospolicy) ToQospolicyOutput() QospolicyOutput {
	return i.ToQospolicyOutputWithContext(context.Background())
}

func (i *Qospolicy) ToQospolicyOutputWithContext(ctx context.Context) QospolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QospolicyOutput)
}

// QospolicyArrayInput is an input type that accepts QospolicyArray and QospolicyArrayOutput values.
// You can construct a concrete instance of `QospolicyArrayInput` via:
//
//	QospolicyArray{ QospolicyArgs{...} }
type QospolicyArrayInput interface {
	pulumi.Input

	ToQospolicyArrayOutput() QospolicyArrayOutput
	ToQospolicyArrayOutputWithContext(context.Context) QospolicyArrayOutput
}

type QospolicyArray []QospolicyInput

func (QospolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Qospolicy)(nil)).Elem()
}

func (i QospolicyArray) ToQospolicyArrayOutput() QospolicyArrayOutput {
	return i.ToQospolicyArrayOutputWithContext(context.Background())
}

func (i QospolicyArray) ToQospolicyArrayOutputWithContext(ctx context.Context) QospolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QospolicyArrayOutput)
}

// QospolicyMapInput is an input type that accepts QospolicyMap and QospolicyMapOutput values.
// You can construct a concrete instance of `QospolicyMapInput` via:
//
//	QospolicyMap{ "key": QospolicyArgs{...} }
type QospolicyMapInput interface {
	pulumi.Input

	ToQospolicyMapOutput() QospolicyMapOutput
	ToQospolicyMapOutputWithContext(context.Context) QospolicyMapOutput
}

type QospolicyMap map[string]QospolicyInput

func (QospolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Qospolicy)(nil)).Elem()
}

func (i QospolicyMap) ToQospolicyMapOutput() QospolicyMapOutput {
	return i.ToQospolicyMapOutputWithContext(context.Background())
}

func (i QospolicyMap) ToQospolicyMapOutputWithContext(ctx context.Context) QospolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QospolicyMapOutput)
}

type QospolicyOutput struct{ *pulumi.OutputState }

func (QospolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Qospolicy)(nil)).Elem()
}

func (o QospolicyOutput) ToQospolicyOutput() QospolicyOutput {
	return o
}

func (o QospolicyOutput) ToQospolicyOutputWithContext(ctx context.Context) QospolicyOutput {
	return o
}

// Default cos queue for untagged packets.
func (o QospolicyOutput) DefaultCos() pulumi.IntOutput {
	return o.ApplyT(func(v *Qospolicy) pulumi.IntOutput { return v.DefaultCos }).(pulumi.IntOutput)
}

// QoS policy name.
func (o QospolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Qospolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// QoS egress queue policy.
func (o QospolicyOutput) QueuePolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Qospolicy) pulumi.StringOutput { return v.QueuePolicy }).(pulumi.StringOutput)
}

// QoS trust 802.1p map.
func (o QospolicyOutput) TrustDot1pMap() pulumi.StringOutput {
	return o.ApplyT(func(v *Qospolicy) pulumi.StringOutput { return v.TrustDot1pMap }).(pulumi.StringOutput)
}

// QoS trust ip dscp map.
func (o QospolicyOutput) TrustIpDscpMap() pulumi.StringOutput {
	return o.ApplyT(func(v *Qospolicy) pulumi.StringOutput { return v.TrustIpDscpMap }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o QospolicyOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Qospolicy) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type QospolicyArrayOutput struct{ *pulumi.OutputState }

func (QospolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Qospolicy)(nil)).Elem()
}

func (o QospolicyArrayOutput) ToQospolicyArrayOutput() QospolicyArrayOutput {
	return o
}

func (o QospolicyArrayOutput) ToQospolicyArrayOutputWithContext(ctx context.Context) QospolicyArrayOutput {
	return o
}

func (o QospolicyArrayOutput) Index(i pulumi.IntInput) QospolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Qospolicy {
		return vs[0].([]*Qospolicy)[vs[1].(int)]
	}).(QospolicyOutput)
}

type QospolicyMapOutput struct{ *pulumi.OutputState }

func (QospolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Qospolicy)(nil)).Elem()
}

func (o QospolicyMapOutput) ToQospolicyMapOutput() QospolicyMapOutput {
	return o
}

func (o QospolicyMapOutput) ToQospolicyMapOutputWithContext(ctx context.Context) QospolicyMapOutput {
	return o
}

func (o QospolicyMapOutput) MapIndex(k pulumi.StringInput) QospolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Qospolicy {
		return vs[0].(map[string]*Qospolicy)[vs[1].(string)]
	}).(QospolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QospolicyInput)(nil)).Elem(), &Qospolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*QospolicyArrayInput)(nil)).Elem(), QospolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QospolicyMapInput)(nil)).Elem(), QospolicyMap{})
	pulumi.RegisterOutputType(QospolicyOutput{})
	pulumi.RegisterOutputType(QospolicyArrayOutput{})
	pulumi.RegisterOutputType(QospolicyMapOutput{})
}
