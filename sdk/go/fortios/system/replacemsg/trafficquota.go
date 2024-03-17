// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package replacemsg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Replacement messages.
//
// ## Import
//
// SystemReplacemsg TrafficQuota can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/replacemsg/trafficquota:Trafficquota labelname {{msg_type}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/replacemsg/trafficquota:Trafficquota labelname {{msg_type}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Trafficquota struct {
	pulumi.CustomResourceState

	// Message string.
	Buffer pulumi.StringPtrOutput `pulumi:"buffer"`
	// Format flag.
	Format pulumi.StringOutput `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringOutput `pulumi:"header"`
	// Message type.
	MsgType pulumi.StringOutput `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewTrafficquota registers a new resource with the given unique name, arguments, and options.
func NewTrafficquota(ctx *pulumi.Context,
	name string, args *TrafficquotaArgs, opts ...pulumi.ResourceOption) (*Trafficquota, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Trafficquota
	err := ctx.RegisterResource("fortios:system/replacemsg/trafficquota:Trafficquota", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficquota gets an existing Trafficquota resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficquota(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficquotaState, opts ...pulumi.ResourceOption) (*Trafficquota, error) {
	var resource Trafficquota
	err := ctx.ReadResource("fortios:system/replacemsg/trafficquota:Trafficquota", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trafficquota resources.
type trafficquotaState struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType *string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type TrafficquotaState struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (TrafficquotaState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficquotaState)(nil)).Elem()
}

type trafficquotaArgs struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Trafficquota resource.
type TrafficquotaArgs struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (TrafficquotaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficquotaArgs)(nil)).Elem()
}

type TrafficquotaInput interface {
	pulumi.Input

	ToTrafficquotaOutput() TrafficquotaOutput
	ToTrafficquotaOutputWithContext(ctx context.Context) TrafficquotaOutput
}

func (*Trafficquota) ElementType() reflect.Type {
	return reflect.TypeOf((**Trafficquota)(nil)).Elem()
}

func (i *Trafficquota) ToTrafficquotaOutput() TrafficquotaOutput {
	return i.ToTrafficquotaOutputWithContext(context.Background())
}

func (i *Trafficquota) ToTrafficquotaOutputWithContext(ctx context.Context) TrafficquotaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficquotaOutput)
}

// TrafficquotaArrayInput is an input type that accepts TrafficquotaArray and TrafficquotaArrayOutput values.
// You can construct a concrete instance of `TrafficquotaArrayInput` via:
//
//	TrafficquotaArray{ TrafficquotaArgs{...} }
type TrafficquotaArrayInput interface {
	pulumi.Input

	ToTrafficquotaArrayOutput() TrafficquotaArrayOutput
	ToTrafficquotaArrayOutputWithContext(context.Context) TrafficquotaArrayOutput
}

type TrafficquotaArray []TrafficquotaInput

func (TrafficquotaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trafficquota)(nil)).Elem()
}

func (i TrafficquotaArray) ToTrafficquotaArrayOutput() TrafficquotaArrayOutput {
	return i.ToTrafficquotaArrayOutputWithContext(context.Background())
}

func (i TrafficquotaArray) ToTrafficquotaArrayOutputWithContext(ctx context.Context) TrafficquotaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficquotaArrayOutput)
}

// TrafficquotaMapInput is an input type that accepts TrafficquotaMap and TrafficquotaMapOutput values.
// You can construct a concrete instance of `TrafficquotaMapInput` via:
//
//	TrafficquotaMap{ "key": TrafficquotaArgs{...} }
type TrafficquotaMapInput interface {
	pulumi.Input

	ToTrafficquotaMapOutput() TrafficquotaMapOutput
	ToTrafficquotaMapOutputWithContext(context.Context) TrafficquotaMapOutput
}

type TrafficquotaMap map[string]TrafficquotaInput

func (TrafficquotaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trafficquota)(nil)).Elem()
}

func (i TrafficquotaMap) ToTrafficquotaMapOutput() TrafficquotaMapOutput {
	return i.ToTrafficquotaMapOutputWithContext(context.Background())
}

func (i TrafficquotaMap) ToTrafficquotaMapOutputWithContext(ctx context.Context) TrafficquotaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficquotaMapOutput)
}

type TrafficquotaOutput struct{ *pulumi.OutputState }

func (TrafficquotaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trafficquota)(nil)).Elem()
}

func (o TrafficquotaOutput) ToTrafficquotaOutput() TrafficquotaOutput {
	return o
}

func (o TrafficquotaOutput) ToTrafficquotaOutputWithContext(ctx context.Context) TrafficquotaOutput {
	return o
}

// Message string.
func (o TrafficquotaOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trafficquota) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag.
func (o TrafficquotaOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *Trafficquota) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o TrafficquotaOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *Trafficquota) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o TrafficquotaOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *Trafficquota) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o TrafficquotaOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trafficquota) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type TrafficquotaArrayOutput struct{ *pulumi.OutputState }

func (TrafficquotaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trafficquota)(nil)).Elem()
}

func (o TrafficquotaArrayOutput) ToTrafficquotaArrayOutput() TrafficquotaArrayOutput {
	return o
}

func (o TrafficquotaArrayOutput) ToTrafficquotaArrayOutputWithContext(ctx context.Context) TrafficquotaArrayOutput {
	return o
}

func (o TrafficquotaArrayOutput) Index(i pulumi.IntInput) TrafficquotaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Trafficquota {
		return vs[0].([]*Trafficquota)[vs[1].(int)]
	}).(TrafficquotaOutput)
}

type TrafficquotaMapOutput struct{ *pulumi.OutputState }

func (TrafficquotaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trafficquota)(nil)).Elem()
}

func (o TrafficquotaMapOutput) ToTrafficquotaMapOutput() TrafficquotaMapOutput {
	return o
}

func (o TrafficquotaMapOutput) ToTrafficquotaMapOutputWithContext(ctx context.Context) TrafficquotaMapOutput {
	return o
}

func (o TrafficquotaMapOutput) MapIndex(k pulumi.StringInput) TrafficquotaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Trafficquota {
		return vs[0].(map[string]*Trafficquota)[vs[1].(string)]
	}).(TrafficquotaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficquotaInput)(nil)).Elem(), &Trafficquota{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficquotaArrayInput)(nil)).Elem(), TrafficquotaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficquotaMapInput)(nil)).Elem(), TrafficquotaMap{})
	pulumi.RegisterOutputType(TrafficquotaOutput{})
	pulumi.RegisterOutputType(TrafficquotaArrayOutput{})
	pulumi.RegisterOutputType(TrafficquotaMapOutput{})
}
