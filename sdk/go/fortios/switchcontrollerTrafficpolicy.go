// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch traffic policy.
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
//			_, err := fortios.NewSwitchcontrollerTrafficpolicy(ctx, "trname", &fortios.SwitchcontrollerTrafficpolicyArgs{
//				GuaranteedBandwidth: pulumi.Int(0),
//				GuaranteedBurst:     pulumi.Int(0),
//				MaximumBurst:        pulumi.Int(0),
//				PolicerStatus:       pulumi.String("enable"),
//				Type:                pulumi.String("ingress"),
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
// # SwitchController TrafficPolicy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerTrafficpolicy:SwitchcontrollerTrafficpolicy labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerTrafficpolicy:SwitchcontrollerTrafficpolicy labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SwitchcontrollerTrafficpolicy struct {
	pulumi.CustomResourceState

	// COS queue(0 - 7), or unset to disable.
	Cos pulumi.IntOutput `pulumi:"cos"`
	// COS queue(0 - 7), or unset to disable.
	CosQueue pulumi.IntOutput `pulumi:"cosQueue"`
	// Description of the traffic policy.
	Description pulumi.StringOutput `pulumi:"description"`
	// FSW Policer id
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Guaranteed bandwidth in kbps (max value = 524287000).
	GuaranteedBandwidth pulumi.IntOutput `pulumi:"guaranteedBandwidth"`
	// Guaranteed burst size in bytes (max value = 4294967295).
	GuaranteedBurst pulumi.IntOutput `pulumi:"guaranteedBurst"`
	// Maximum burst size in bytes (max value = 4294967295).
	MaximumBurst pulumi.IntOutput `pulumi:"maximumBurst"`
	// Traffic policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable policer config on the traffic policy. Valid values: `enable`, `disable`.
	PolicerStatus pulumi.StringOutput `pulumi:"policerStatus"`
	// Configure type of policy(ingress/egress). Valid values: `ingress`, `egress`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchcontrollerTrafficpolicy registers a new resource with the given unique name, arguments, and options.
func NewSwitchcontrollerTrafficpolicy(ctx *pulumi.Context,
	name string, args *SwitchcontrollerTrafficpolicyArgs, opts ...pulumi.ResourceOption) (*SwitchcontrollerTrafficpolicy, error) {
	if args == nil {
		args = &SwitchcontrollerTrafficpolicyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchcontrollerTrafficpolicy
	err := ctx.RegisterResource("fortios:index/switchcontrollerTrafficpolicy:SwitchcontrollerTrafficpolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchcontrollerTrafficpolicy gets an existing SwitchcontrollerTrafficpolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchcontrollerTrafficpolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchcontrollerTrafficpolicyState, opts ...pulumi.ResourceOption) (*SwitchcontrollerTrafficpolicy, error) {
	var resource SwitchcontrollerTrafficpolicy
	err := ctx.ReadResource("fortios:index/switchcontrollerTrafficpolicy:SwitchcontrollerTrafficpolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchcontrollerTrafficpolicy resources.
type switchcontrollerTrafficpolicyState struct {
	// COS queue(0 - 7), or unset to disable.
	Cos *int `pulumi:"cos"`
	// COS queue(0 - 7), or unset to disable.
	CosQueue *int `pulumi:"cosQueue"`
	// Description of the traffic policy.
	Description *string `pulumi:"description"`
	// FSW Policer id
	Fosid *int `pulumi:"fosid"`
	// Guaranteed bandwidth in kbps (max value = 524287000).
	GuaranteedBandwidth *int `pulumi:"guaranteedBandwidth"`
	// Guaranteed burst size in bytes (max value = 4294967295).
	GuaranteedBurst *int `pulumi:"guaranteedBurst"`
	// Maximum burst size in bytes (max value = 4294967295).
	MaximumBurst *int `pulumi:"maximumBurst"`
	// Traffic policy name.
	Name *string `pulumi:"name"`
	// Enable/disable policer config on the traffic policy. Valid values: `enable`, `disable`.
	PolicerStatus *string `pulumi:"policerStatus"`
	// Configure type of policy(ingress/egress). Valid values: `ingress`, `egress`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchcontrollerTrafficpolicyState struct {
	// COS queue(0 - 7), or unset to disable.
	Cos pulumi.IntPtrInput
	// COS queue(0 - 7), or unset to disable.
	CosQueue pulumi.IntPtrInput
	// Description of the traffic policy.
	Description pulumi.StringPtrInput
	// FSW Policer id
	Fosid pulumi.IntPtrInput
	// Guaranteed bandwidth in kbps (max value = 524287000).
	GuaranteedBandwidth pulumi.IntPtrInput
	// Guaranteed burst size in bytes (max value = 4294967295).
	GuaranteedBurst pulumi.IntPtrInput
	// Maximum burst size in bytes (max value = 4294967295).
	MaximumBurst pulumi.IntPtrInput
	// Traffic policy name.
	Name pulumi.StringPtrInput
	// Enable/disable policer config on the traffic policy. Valid values: `enable`, `disable`.
	PolicerStatus pulumi.StringPtrInput
	// Configure type of policy(ingress/egress). Valid values: `ingress`, `egress`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerTrafficpolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerTrafficpolicyState)(nil)).Elem()
}

type switchcontrollerTrafficpolicyArgs struct {
	// COS queue(0 - 7), or unset to disable.
	Cos *int `pulumi:"cos"`
	// COS queue(0 - 7), or unset to disable.
	CosQueue *int `pulumi:"cosQueue"`
	// Description of the traffic policy.
	Description *string `pulumi:"description"`
	// FSW Policer id
	Fosid *int `pulumi:"fosid"`
	// Guaranteed bandwidth in kbps (max value = 524287000).
	GuaranteedBandwidth *int `pulumi:"guaranteedBandwidth"`
	// Guaranteed burst size in bytes (max value = 4294967295).
	GuaranteedBurst *int `pulumi:"guaranteedBurst"`
	// Maximum burst size in bytes (max value = 4294967295).
	MaximumBurst *int `pulumi:"maximumBurst"`
	// Traffic policy name.
	Name *string `pulumi:"name"`
	// Enable/disable policer config on the traffic policy. Valid values: `enable`, `disable`.
	PolicerStatus *string `pulumi:"policerStatus"`
	// Configure type of policy(ingress/egress). Valid values: `ingress`, `egress`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchcontrollerTrafficpolicy resource.
type SwitchcontrollerTrafficpolicyArgs struct {
	// COS queue(0 - 7), or unset to disable.
	Cos pulumi.IntPtrInput
	// COS queue(0 - 7), or unset to disable.
	CosQueue pulumi.IntPtrInput
	// Description of the traffic policy.
	Description pulumi.StringPtrInput
	// FSW Policer id
	Fosid pulumi.IntPtrInput
	// Guaranteed bandwidth in kbps (max value = 524287000).
	GuaranteedBandwidth pulumi.IntPtrInput
	// Guaranteed burst size in bytes (max value = 4294967295).
	GuaranteedBurst pulumi.IntPtrInput
	// Maximum burst size in bytes (max value = 4294967295).
	MaximumBurst pulumi.IntPtrInput
	// Traffic policy name.
	Name pulumi.StringPtrInput
	// Enable/disable policer config on the traffic policy. Valid values: `enable`, `disable`.
	PolicerStatus pulumi.StringPtrInput
	// Configure type of policy(ingress/egress). Valid values: `ingress`, `egress`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerTrafficpolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerTrafficpolicyArgs)(nil)).Elem()
}

type SwitchcontrollerTrafficpolicyInput interface {
	pulumi.Input

	ToSwitchcontrollerTrafficpolicyOutput() SwitchcontrollerTrafficpolicyOutput
	ToSwitchcontrollerTrafficpolicyOutputWithContext(ctx context.Context) SwitchcontrollerTrafficpolicyOutput
}

func (*SwitchcontrollerTrafficpolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerTrafficpolicy)(nil)).Elem()
}

func (i *SwitchcontrollerTrafficpolicy) ToSwitchcontrollerTrafficpolicyOutput() SwitchcontrollerTrafficpolicyOutput {
	return i.ToSwitchcontrollerTrafficpolicyOutputWithContext(context.Background())
}

func (i *SwitchcontrollerTrafficpolicy) ToSwitchcontrollerTrafficpolicyOutputWithContext(ctx context.Context) SwitchcontrollerTrafficpolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerTrafficpolicyOutput)
}

// SwitchcontrollerTrafficpolicyArrayInput is an input type that accepts SwitchcontrollerTrafficpolicyArray and SwitchcontrollerTrafficpolicyArrayOutput values.
// You can construct a concrete instance of `SwitchcontrollerTrafficpolicyArrayInput` via:
//
//	SwitchcontrollerTrafficpolicyArray{ SwitchcontrollerTrafficpolicyArgs{...} }
type SwitchcontrollerTrafficpolicyArrayInput interface {
	pulumi.Input

	ToSwitchcontrollerTrafficpolicyArrayOutput() SwitchcontrollerTrafficpolicyArrayOutput
	ToSwitchcontrollerTrafficpolicyArrayOutputWithContext(context.Context) SwitchcontrollerTrafficpolicyArrayOutput
}

type SwitchcontrollerTrafficpolicyArray []SwitchcontrollerTrafficpolicyInput

func (SwitchcontrollerTrafficpolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerTrafficpolicy)(nil)).Elem()
}

func (i SwitchcontrollerTrafficpolicyArray) ToSwitchcontrollerTrafficpolicyArrayOutput() SwitchcontrollerTrafficpolicyArrayOutput {
	return i.ToSwitchcontrollerTrafficpolicyArrayOutputWithContext(context.Background())
}

func (i SwitchcontrollerTrafficpolicyArray) ToSwitchcontrollerTrafficpolicyArrayOutputWithContext(ctx context.Context) SwitchcontrollerTrafficpolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerTrafficpolicyArrayOutput)
}

// SwitchcontrollerTrafficpolicyMapInput is an input type that accepts SwitchcontrollerTrafficpolicyMap and SwitchcontrollerTrafficpolicyMapOutput values.
// You can construct a concrete instance of `SwitchcontrollerTrafficpolicyMapInput` via:
//
//	SwitchcontrollerTrafficpolicyMap{ "key": SwitchcontrollerTrafficpolicyArgs{...} }
type SwitchcontrollerTrafficpolicyMapInput interface {
	pulumi.Input

	ToSwitchcontrollerTrafficpolicyMapOutput() SwitchcontrollerTrafficpolicyMapOutput
	ToSwitchcontrollerTrafficpolicyMapOutputWithContext(context.Context) SwitchcontrollerTrafficpolicyMapOutput
}

type SwitchcontrollerTrafficpolicyMap map[string]SwitchcontrollerTrafficpolicyInput

func (SwitchcontrollerTrafficpolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerTrafficpolicy)(nil)).Elem()
}

func (i SwitchcontrollerTrafficpolicyMap) ToSwitchcontrollerTrafficpolicyMapOutput() SwitchcontrollerTrafficpolicyMapOutput {
	return i.ToSwitchcontrollerTrafficpolicyMapOutputWithContext(context.Background())
}

func (i SwitchcontrollerTrafficpolicyMap) ToSwitchcontrollerTrafficpolicyMapOutputWithContext(ctx context.Context) SwitchcontrollerTrafficpolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerTrafficpolicyMapOutput)
}

type SwitchcontrollerTrafficpolicyOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerTrafficpolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerTrafficpolicy)(nil)).Elem()
}

func (o SwitchcontrollerTrafficpolicyOutput) ToSwitchcontrollerTrafficpolicyOutput() SwitchcontrollerTrafficpolicyOutput {
	return o
}

func (o SwitchcontrollerTrafficpolicyOutput) ToSwitchcontrollerTrafficpolicyOutputWithContext(ctx context.Context) SwitchcontrollerTrafficpolicyOutput {
	return o
}

// COS queue(0 - 7), or unset to disable.
func (o SwitchcontrollerTrafficpolicyOutput) Cos() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerTrafficpolicy) pulumi.IntOutput { return v.Cos }).(pulumi.IntOutput)
}

// COS queue(0 - 7), or unset to disable.
func (o SwitchcontrollerTrafficpolicyOutput) CosQueue() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerTrafficpolicy) pulumi.IntOutput { return v.CosQueue }).(pulumi.IntOutput)
}

// Description of the traffic policy.
func (o SwitchcontrollerTrafficpolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerTrafficpolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// FSW Policer id
func (o SwitchcontrollerTrafficpolicyOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerTrafficpolicy) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Guaranteed bandwidth in kbps (max value = 524287000).
func (o SwitchcontrollerTrafficpolicyOutput) GuaranteedBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerTrafficpolicy) pulumi.IntOutput { return v.GuaranteedBandwidth }).(pulumi.IntOutput)
}

// Guaranteed burst size in bytes (max value = 4294967295).
func (o SwitchcontrollerTrafficpolicyOutput) GuaranteedBurst() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerTrafficpolicy) pulumi.IntOutput { return v.GuaranteedBurst }).(pulumi.IntOutput)
}

// Maximum burst size in bytes (max value = 4294967295).
func (o SwitchcontrollerTrafficpolicyOutput) MaximumBurst() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerTrafficpolicy) pulumi.IntOutput { return v.MaximumBurst }).(pulumi.IntOutput)
}

// Traffic policy name.
func (o SwitchcontrollerTrafficpolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerTrafficpolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable policer config on the traffic policy. Valid values: `enable`, `disable`.
func (o SwitchcontrollerTrafficpolicyOutput) PolicerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerTrafficpolicy) pulumi.StringOutput { return v.PolicerStatus }).(pulumi.StringOutput)
}

// Configure type of policy(ingress/egress). Valid values: `ingress`, `egress`.
func (o SwitchcontrollerTrafficpolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerTrafficpolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchcontrollerTrafficpolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerTrafficpolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchcontrollerTrafficpolicyArrayOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerTrafficpolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerTrafficpolicy)(nil)).Elem()
}

func (o SwitchcontrollerTrafficpolicyArrayOutput) ToSwitchcontrollerTrafficpolicyArrayOutput() SwitchcontrollerTrafficpolicyArrayOutput {
	return o
}

func (o SwitchcontrollerTrafficpolicyArrayOutput) ToSwitchcontrollerTrafficpolicyArrayOutputWithContext(ctx context.Context) SwitchcontrollerTrafficpolicyArrayOutput {
	return o
}

func (o SwitchcontrollerTrafficpolicyArrayOutput) Index(i pulumi.IntInput) SwitchcontrollerTrafficpolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchcontrollerTrafficpolicy {
		return vs[0].([]*SwitchcontrollerTrafficpolicy)[vs[1].(int)]
	}).(SwitchcontrollerTrafficpolicyOutput)
}

type SwitchcontrollerTrafficpolicyMapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerTrafficpolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerTrafficpolicy)(nil)).Elem()
}

func (o SwitchcontrollerTrafficpolicyMapOutput) ToSwitchcontrollerTrafficpolicyMapOutput() SwitchcontrollerTrafficpolicyMapOutput {
	return o
}

func (o SwitchcontrollerTrafficpolicyMapOutput) ToSwitchcontrollerTrafficpolicyMapOutputWithContext(ctx context.Context) SwitchcontrollerTrafficpolicyMapOutput {
	return o
}

func (o SwitchcontrollerTrafficpolicyMapOutput) MapIndex(k pulumi.StringInput) SwitchcontrollerTrafficpolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchcontrollerTrafficpolicy {
		return vs[0].(map[string]*SwitchcontrollerTrafficpolicy)[vs[1].(string)]
	}).(SwitchcontrollerTrafficpolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerTrafficpolicyInput)(nil)).Elem(), &SwitchcontrollerTrafficpolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerTrafficpolicyArrayInput)(nil)).Elem(), SwitchcontrollerTrafficpolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerTrafficpolicyMapInput)(nil)).Elem(), SwitchcontrollerTrafficpolicyMap{})
	pulumi.RegisterOutputType(SwitchcontrollerTrafficpolicyOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerTrafficpolicyArrayOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerTrafficpolicyMapOutput{})
}