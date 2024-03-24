// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ptp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// PTP interface-policy configuration. Applies to FortiOS Version `>= 7.4.1`.
//
// ## Import
//
// SwitchControllerPtp InterfacePolicy can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/ptp/interfacepolicy:Interfacepolicy labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/ptp/interfacepolicy:Interfacepolicy labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Interfacepolicy struct {
	pulumi.CustomResourceState

	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// PTP VLAN.
	Vlan pulumi.StringOutput `pulumi:"vlan"`
	// Configure PTP VLAN priority (0 - 7, default = 4).
	VlanPri pulumi.IntOutput `pulumi:"vlanPri"`
}

// NewInterfacepolicy registers a new resource with the given unique name, arguments, and options.
func NewInterfacepolicy(ctx *pulumi.Context,
	name string, args *InterfacepolicyArgs, opts ...pulumi.ResourceOption) (*Interfacepolicy, error) {
	if args == nil {
		args = &InterfacepolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Interfacepolicy
	err := ctx.RegisterResource("fortios:switchcontroller/ptp/interfacepolicy:Interfacepolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInterfacepolicy gets an existing Interfacepolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInterfacepolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InterfacepolicyState, opts ...pulumi.ResourceOption) (*Interfacepolicy, error) {
	var resource Interfacepolicy
	err := ctx.ReadResource("fortios:switchcontroller/ptp/interfacepolicy:Interfacepolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Interfacepolicy resources.
type interfacepolicyState struct {
	// Description.
	Description *string `pulumi:"description"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// PTP VLAN.
	Vlan *string `pulumi:"vlan"`
	// Configure PTP VLAN priority (0 - 7, default = 4).
	VlanPri *int `pulumi:"vlanPri"`
}

type InterfacepolicyState struct {
	// Description.
	Description pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// PTP VLAN.
	Vlan pulumi.StringPtrInput
	// Configure PTP VLAN priority (0 - 7, default = 4).
	VlanPri pulumi.IntPtrInput
}

func (InterfacepolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*interfacepolicyState)(nil)).Elem()
}

type interfacepolicyArgs struct {
	// Description.
	Description *string `pulumi:"description"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// PTP VLAN.
	Vlan *string `pulumi:"vlan"`
	// Configure PTP VLAN priority (0 - 7, default = 4).
	VlanPri *int `pulumi:"vlanPri"`
}

// The set of arguments for constructing a Interfacepolicy resource.
type InterfacepolicyArgs struct {
	// Description.
	Description pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// PTP VLAN.
	Vlan pulumi.StringPtrInput
	// Configure PTP VLAN priority (0 - 7, default = 4).
	VlanPri pulumi.IntPtrInput
}

func (InterfacepolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*interfacepolicyArgs)(nil)).Elem()
}

type InterfacepolicyInput interface {
	pulumi.Input

	ToInterfacepolicyOutput() InterfacepolicyOutput
	ToInterfacepolicyOutputWithContext(ctx context.Context) InterfacepolicyOutput
}

func (*Interfacepolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**Interfacepolicy)(nil)).Elem()
}

func (i *Interfacepolicy) ToInterfacepolicyOutput() InterfacepolicyOutput {
	return i.ToInterfacepolicyOutputWithContext(context.Background())
}

func (i *Interfacepolicy) ToInterfacepolicyOutputWithContext(ctx context.Context) InterfacepolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfacepolicyOutput)
}

// InterfacepolicyArrayInput is an input type that accepts InterfacepolicyArray and InterfacepolicyArrayOutput values.
// You can construct a concrete instance of `InterfacepolicyArrayInput` via:
//
//	InterfacepolicyArray{ InterfacepolicyArgs{...} }
type InterfacepolicyArrayInput interface {
	pulumi.Input

	ToInterfacepolicyArrayOutput() InterfacepolicyArrayOutput
	ToInterfacepolicyArrayOutputWithContext(context.Context) InterfacepolicyArrayOutput
}

type InterfacepolicyArray []InterfacepolicyInput

func (InterfacepolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Interfacepolicy)(nil)).Elem()
}

func (i InterfacepolicyArray) ToInterfacepolicyArrayOutput() InterfacepolicyArrayOutput {
	return i.ToInterfacepolicyArrayOutputWithContext(context.Background())
}

func (i InterfacepolicyArray) ToInterfacepolicyArrayOutputWithContext(ctx context.Context) InterfacepolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfacepolicyArrayOutput)
}

// InterfacepolicyMapInput is an input type that accepts InterfacepolicyMap and InterfacepolicyMapOutput values.
// You can construct a concrete instance of `InterfacepolicyMapInput` via:
//
//	InterfacepolicyMap{ "key": InterfacepolicyArgs{...} }
type InterfacepolicyMapInput interface {
	pulumi.Input

	ToInterfacepolicyMapOutput() InterfacepolicyMapOutput
	ToInterfacepolicyMapOutputWithContext(context.Context) InterfacepolicyMapOutput
}

type InterfacepolicyMap map[string]InterfacepolicyInput

func (InterfacepolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Interfacepolicy)(nil)).Elem()
}

func (i InterfacepolicyMap) ToInterfacepolicyMapOutput() InterfacepolicyMapOutput {
	return i.ToInterfacepolicyMapOutputWithContext(context.Background())
}

func (i InterfacepolicyMap) ToInterfacepolicyMapOutputWithContext(ctx context.Context) InterfacepolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfacepolicyMapOutput)
}

type InterfacepolicyOutput struct{ *pulumi.OutputState }

func (InterfacepolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Interfacepolicy)(nil)).Elem()
}

func (o InterfacepolicyOutput) ToInterfacepolicyOutput() InterfacepolicyOutput {
	return o
}

func (o InterfacepolicyOutput) ToInterfacepolicyOutputWithContext(ctx context.Context) InterfacepolicyOutput {
	return o
}

// Description.
func (o InterfacepolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Interfacepolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Policy name.
func (o InterfacepolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Interfacepolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o InterfacepolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Interfacepolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// PTP VLAN.
func (o InterfacepolicyOutput) Vlan() pulumi.StringOutput {
	return o.ApplyT(func(v *Interfacepolicy) pulumi.StringOutput { return v.Vlan }).(pulumi.StringOutput)
}

// Configure PTP VLAN priority (0 - 7, default = 4).
func (o InterfacepolicyOutput) VlanPri() pulumi.IntOutput {
	return o.ApplyT(func(v *Interfacepolicy) pulumi.IntOutput { return v.VlanPri }).(pulumi.IntOutput)
}

type InterfacepolicyArrayOutput struct{ *pulumi.OutputState }

func (InterfacepolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Interfacepolicy)(nil)).Elem()
}

func (o InterfacepolicyArrayOutput) ToInterfacepolicyArrayOutput() InterfacepolicyArrayOutput {
	return o
}

func (o InterfacepolicyArrayOutput) ToInterfacepolicyArrayOutputWithContext(ctx context.Context) InterfacepolicyArrayOutput {
	return o
}

func (o InterfacepolicyArrayOutput) Index(i pulumi.IntInput) InterfacepolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Interfacepolicy {
		return vs[0].([]*Interfacepolicy)[vs[1].(int)]
	}).(InterfacepolicyOutput)
}

type InterfacepolicyMapOutput struct{ *pulumi.OutputState }

func (InterfacepolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Interfacepolicy)(nil)).Elem()
}

func (o InterfacepolicyMapOutput) ToInterfacepolicyMapOutput() InterfacepolicyMapOutput {
	return o
}

func (o InterfacepolicyMapOutput) ToInterfacepolicyMapOutputWithContext(ctx context.Context) InterfacepolicyMapOutput {
	return o
}

func (o InterfacepolicyMapOutput) MapIndex(k pulumi.StringInput) InterfacepolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Interfacepolicy {
		return vs[0].(map[string]*Interfacepolicy)[vs[1].(string)]
	}).(InterfacepolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InterfacepolicyInput)(nil)).Elem(), &Interfacepolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*InterfacepolicyArrayInput)(nil)).Elem(), InterfacepolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InterfacepolicyMapInput)(nil)).Elem(), InterfacepolicyMap{})
	pulumi.RegisterOutputType(InterfacepolicyOutput{})
	pulumi.RegisterOutputType(InterfacepolicyArrayOutput{})
	pulumi.RegisterOutputType(InterfacepolicyMapOutput{})
}