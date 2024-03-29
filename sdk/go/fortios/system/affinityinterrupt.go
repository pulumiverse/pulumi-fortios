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

// Configure interrupt affinity.
//
// ## Import
//
// System AffinityInterrupt can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/affinityinterrupt:Affinityinterrupt labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/affinityinterrupt:Affinityinterrupt labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Affinityinterrupt struct {
	pulumi.CustomResourceState

	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask pulumi.StringOutput `pulumi:"affinityCpumask"`
	// Default affinity setting (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	DefaultAffinityCpumask pulumi.StringOutput `pulumi:"defaultAffinityCpumask"`
	// ID of the interrupt affinity setting.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Interrupt name.
	Interrupt pulumi.StringOutput `pulumi:"interrupt"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewAffinityinterrupt registers a new resource with the given unique name, arguments, and options.
func NewAffinityinterrupt(ctx *pulumi.Context,
	name string, args *AffinityinterruptArgs, opts ...pulumi.ResourceOption) (*Affinityinterrupt, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AffinityCpumask == nil {
		return nil, errors.New("invalid value for required argument 'AffinityCpumask'")
	}
	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	if args.Interrupt == nil {
		return nil, errors.New("invalid value for required argument 'Interrupt'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Affinityinterrupt
	err := ctx.RegisterResource("fortios:system/affinityinterrupt:Affinityinterrupt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAffinityinterrupt gets an existing Affinityinterrupt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAffinityinterrupt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AffinityinterruptState, opts ...pulumi.ResourceOption) (*Affinityinterrupt, error) {
	var resource Affinityinterrupt
	err := ctx.ReadResource("fortios:system/affinityinterrupt:Affinityinterrupt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Affinityinterrupt resources.
type affinityinterruptState struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask *string `pulumi:"affinityCpumask"`
	// Default affinity setting (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	DefaultAffinityCpumask *string `pulumi:"defaultAffinityCpumask"`
	// ID of the interrupt affinity setting.
	Fosid *int `pulumi:"fosid"`
	// Interrupt name.
	Interrupt *string `pulumi:"interrupt"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type AffinityinterruptState struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask pulumi.StringPtrInput
	// Default affinity setting (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	DefaultAffinityCpumask pulumi.StringPtrInput
	// ID of the interrupt affinity setting.
	Fosid pulumi.IntPtrInput
	// Interrupt name.
	Interrupt pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AffinityinterruptState) ElementType() reflect.Type {
	return reflect.TypeOf((*affinityinterruptState)(nil)).Elem()
}

type affinityinterruptArgs struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask string `pulumi:"affinityCpumask"`
	// Default affinity setting (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	DefaultAffinityCpumask *string `pulumi:"defaultAffinityCpumask"`
	// ID of the interrupt affinity setting.
	Fosid int `pulumi:"fosid"`
	// Interrupt name.
	Interrupt string `pulumi:"interrupt"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Affinityinterrupt resource.
type AffinityinterruptArgs struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask pulumi.StringInput
	// Default affinity setting (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	DefaultAffinityCpumask pulumi.StringPtrInput
	// ID of the interrupt affinity setting.
	Fosid pulumi.IntInput
	// Interrupt name.
	Interrupt pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AffinityinterruptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*affinityinterruptArgs)(nil)).Elem()
}

type AffinityinterruptInput interface {
	pulumi.Input

	ToAffinityinterruptOutput() AffinityinterruptOutput
	ToAffinityinterruptOutputWithContext(ctx context.Context) AffinityinterruptOutput
}

func (*Affinityinterrupt) ElementType() reflect.Type {
	return reflect.TypeOf((**Affinityinterrupt)(nil)).Elem()
}

func (i *Affinityinterrupt) ToAffinityinterruptOutput() AffinityinterruptOutput {
	return i.ToAffinityinterruptOutputWithContext(context.Background())
}

func (i *Affinityinterrupt) ToAffinityinterruptOutputWithContext(ctx context.Context) AffinityinterruptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AffinityinterruptOutput)
}

// AffinityinterruptArrayInput is an input type that accepts AffinityinterruptArray and AffinityinterruptArrayOutput values.
// You can construct a concrete instance of `AffinityinterruptArrayInput` via:
//
//	AffinityinterruptArray{ AffinityinterruptArgs{...} }
type AffinityinterruptArrayInput interface {
	pulumi.Input

	ToAffinityinterruptArrayOutput() AffinityinterruptArrayOutput
	ToAffinityinterruptArrayOutputWithContext(context.Context) AffinityinterruptArrayOutput
}

type AffinityinterruptArray []AffinityinterruptInput

func (AffinityinterruptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Affinityinterrupt)(nil)).Elem()
}

func (i AffinityinterruptArray) ToAffinityinterruptArrayOutput() AffinityinterruptArrayOutput {
	return i.ToAffinityinterruptArrayOutputWithContext(context.Background())
}

func (i AffinityinterruptArray) ToAffinityinterruptArrayOutputWithContext(ctx context.Context) AffinityinterruptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AffinityinterruptArrayOutput)
}

// AffinityinterruptMapInput is an input type that accepts AffinityinterruptMap and AffinityinterruptMapOutput values.
// You can construct a concrete instance of `AffinityinterruptMapInput` via:
//
//	AffinityinterruptMap{ "key": AffinityinterruptArgs{...} }
type AffinityinterruptMapInput interface {
	pulumi.Input

	ToAffinityinterruptMapOutput() AffinityinterruptMapOutput
	ToAffinityinterruptMapOutputWithContext(context.Context) AffinityinterruptMapOutput
}

type AffinityinterruptMap map[string]AffinityinterruptInput

func (AffinityinterruptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Affinityinterrupt)(nil)).Elem()
}

func (i AffinityinterruptMap) ToAffinityinterruptMapOutput() AffinityinterruptMapOutput {
	return i.ToAffinityinterruptMapOutputWithContext(context.Background())
}

func (i AffinityinterruptMap) ToAffinityinterruptMapOutputWithContext(ctx context.Context) AffinityinterruptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AffinityinterruptMapOutput)
}

type AffinityinterruptOutput struct{ *pulumi.OutputState }

func (AffinityinterruptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Affinityinterrupt)(nil)).Elem()
}

func (o AffinityinterruptOutput) ToAffinityinterruptOutput() AffinityinterruptOutput {
	return o
}

func (o AffinityinterruptOutput) ToAffinityinterruptOutputWithContext(ctx context.Context) AffinityinterruptOutput {
	return o
}

// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
func (o AffinityinterruptOutput) AffinityCpumask() pulumi.StringOutput {
	return o.ApplyT(func(v *Affinityinterrupt) pulumi.StringOutput { return v.AffinityCpumask }).(pulumi.StringOutput)
}

// Default affinity setting (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
func (o AffinityinterruptOutput) DefaultAffinityCpumask() pulumi.StringOutput {
	return o.ApplyT(func(v *Affinityinterrupt) pulumi.StringOutput { return v.DefaultAffinityCpumask }).(pulumi.StringOutput)
}

// ID of the interrupt affinity setting.
func (o AffinityinterruptOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Affinityinterrupt) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Interrupt name.
func (o AffinityinterruptOutput) Interrupt() pulumi.StringOutput {
	return o.ApplyT(func(v *Affinityinterrupt) pulumi.StringOutput { return v.Interrupt }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AffinityinterruptOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Affinityinterrupt) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AffinityinterruptArrayOutput struct{ *pulumi.OutputState }

func (AffinityinterruptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Affinityinterrupt)(nil)).Elem()
}

func (o AffinityinterruptArrayOutput) ToAffinityinterruptArrayOutput() AffinityinterruptArrayOutput {
	return o
}

func (o AffinityinterruptArrayOutput) ToAffinityinterruptArrayOutputWithContext(ctx context.Context) AffinityinterruptArrayOutput {
	return o
}

func (o AffinityinterruptArrayOutput) Index(i pulumi.IntInput) AffinityinterruptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Affinityinterrupt {
		return vs[0].([]*Affinityinterrupt)[vs[1].(int)]
	}).(AffinityinterruptOutput)
}

type AffinityinterruptMapOutput struct{ *pulumi.OutputState }

func (AffinityinterruptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Affinityinterrupt)(nil)).Elem()
}

func (o AffinityinterruptMapOutput) ToAffinityinterruptMapOutput() AffinityinterruptMapOutput {
	return o
}

func (o AffinityinterruptMapOutput) ToAffinityinterruptMapOutputWithContext(ctx context.Context) AffinityinterruptMapOutput {
	return o
}

func (o AffinityinterruptMapOutput) MapIndex(k pulumi.StringInput) AffinityinterruptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Affinityinterrupt {
		return vs[0].(map[string]*Affinityinterrupt)[vs[1].(string)]
	}).(AffinityinterruptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AffinityinterruptInput)(nil)).Elem(), &Affinityinterrupt{})
	pulumi.RegisterInputType(reflect.TypeOf((*AffinityinterruptArrayInput)(nil)).Elem(), AffinityinterruptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AffinityinterruptMapInput)(nil)).Elem(), AffinityinterruptMap{})
	pulumi.RegisterOutputType(AffinityinterruptOutput{})
	pulumi.RegisterOutputType(AffinityinterruptArrayOutput{})
	pulumi.RegisterOutputType(AffinityinterruptMapOutput{})
}
