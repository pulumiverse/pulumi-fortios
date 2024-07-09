// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure VDOM links.
//
// ## Import
//
// System VdomLink can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/vdomlink:Vdomlink labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/vdomlink:Vdomlink labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Vdomlink struct {
	pulumi.CustomResourceState

	// VDOM link name. On FortiOS versions 6.2.0-6.4.0: maximum = 8 characters. On FortiOS versions >= 6.4.1: maximum = 11 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// VDOM link type: PPP or Ethernet.
	Type pulumi.StringOutput `pulumi:"type"`
	// Virtual cluster. Valid values: `vcluster1`, `vcluster2`.
	Vcluster pulumi.StringOutput `pulumi:"vcluster"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewVdomlink registers a new resource with the given unique name, arguments, and options.
func NewVdomlink(ctx *pulumi.Context,
	name string, args *VdomlinkArgs, opts ...pulumi.ResourceOption) (*Vdomlink, error) {
	if args == nil {
		args = &VdomlinkArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vdomlink
	err := ctx.RegisterResource("fortios:system/vdomlink:Vdomlink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVdomlink gets an existing Vdomlink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVdomlink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VdomlinkState, opts ...pulumi.ResourceOption) (*Vdomlink, error) {
	var resource Vdomlink
	err := ctx.ReadResource("fortios:system/vdomlink:Vdomlink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vdomlink resources.
type vdomlinkState struct {
	// VDOM link name. On FortiOS versions 6.2.0-6.4.0: maximum = 8 characters. On FortiOS versions >= 6.4.1: maximum = 11 characters.
	Name *string `pulumi:"name"`
	// VDOM link type: PPP or Ethernet.
	Type *string `pulumi:"type"`
	// Virtual cluster. Valid values: `vcluster1`, `vcluster2`.
	Vcluster *string `pulumi:"vcluster"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type VdomlinkState struct {
	// VDOM link name. On FortiOS versions 6.2.0-6.4.0: maximum = 8 characters. On FortiOS versions >= 6.4.1: maximum = 11 characters.
	Name pulumi.StringPtrInput
	// VDOM link type: PPP or Ethernet.
	Type pulumi.StringPtrInput
	// Virtual cluster. Valid values: `vcluster1`, `vcluster2`.
	Vcluster pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VdomlinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vdomlinkState)(nil)).Elem()
}

type vdomlinkArgs struct {
	// VDOM link name. On FortiOS versions 6.2.0-6.4.0: maximum = 8 characters. On FortiOS versions >= 6.4.1: maximum = 11 characters.
	Name *string `pulumi:"name"`
	// VDOM link type: PPP or Ethernet.
	Type *string `pulumi:"type"`
	// Virtual cluster. Valid values: `vcluster1`, `vcluster2`.
	Vcluster *string `pulumi:"vcluster"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Vdomlink resource.
type VdomlinkArgs struct {
	// VDOM link name. On FortiOS versions 6.2.0-6.4.0: maximum = 8 characters. On FortiOS versions >= 6.4.1: maximum = 11 characters.
	Name pulumi.StringPtrInput
	// VDOM link type: PPP or Ethernet.
	Type pulumi.StringPtrInput
	// Virtual cluster. Valid values: `vcluster1`, `vcluster2`.
	Vcluster pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VdomlinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vdomlinkArgs)(nil)).Elem()
}

type VdomlinkInput interface {
	pulumi.Input

	ToVdomlinkOutput() VdomlinkOutput
	ToVdomlinkOutputWithContext(ctx context.Context) VdomlinkOutput
}

func (*Vdomlink) ElementType() reflect.Type {
	return reflect.TypeOf((**Vdomlink)(nil)).Elem()
}

func (i *Vdomlink) ToVdomlinkOutput() VdomlinkOutput {
	return i.ToVdomlinkOutputWithContext(context.Background())
}

func (i *Vdomlink) ToVdomlinkOutputWithContext(ctx context.Context) VdomlinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VdomlinkOutput)
}

// VdomlinkArrayInput is an input type that accepts VdomlinkArray and VdomlinkArrayOutput values.
// You can construct a concrete instance of `VdomlinkArrayInput` via:
//
//	VdomlinkArray{ VdomlinkArgs{...} }
type VdomlinkArrayInput interface {
	pulumi.Input

	ToVdomlinkArrayOutput() VdomlinkArrayOutput
	ToVdomlinkArrayOutputWithContext(context.Context) VdomlinkArrayOutput
}

type VdomlinkArray []VdomlinkInput

func (VdomlinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vdomlink)(nil)).Elem()
}

func (i VdomlinkArray) ToVdomlinkArrayOutput() VdomlinkArrayOutput {
	return i.ToVdomlinkArrayOutputWithContext(context.Background())
}

func (i VdomlinkArray) ToVdomlinkArrayOutputWithContext(ctx context.Context) VdomlinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VdomlinkArrayOutput)
}

// VdomlinkMapInput is an input type that accepts VdomlinkMap and VdomlinkMapOutput values.
// You can construct a concrete instance of `VdomlinkMapInput` via:
//
//	VdomlinkMap{ "key": VdomlinkArgs{...} }
type VdomlinkMapInput interface {
	pulumi.Input

	ToVdomlinkMapOutput() VdomlinkMapOutput
	ToVdomlinkMapOutputWithContext(context.Context) VdomlinkMapOutput
}

type VdomlinkMap map[string]VdomlinkInput

func (VdomlinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vdomlink)(nil)).Elem()
}

func (i VdomlinkMap) ToVdomlinkMapOutput() VdomlinkMapOutput {
	return i.ToVdomlinkMapOutputWithContext(context.Background())
}

func (i VdomlinkMap) ToVdomlinkMapOutputWithContext(ctx context.Context) VdomlinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VdomlinkMapOutput)
}

type VdomlinkOutput struct{ *pulumi.OutputState }

func (VdomlinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vdomlink)(nil)).Elem()
}

func (o VdomlinkOutput) ToVdomlinkOutput() VdomlinkOutput {
	return o
}

func (o VdomlinkOutput) ToVdomlinkOutputWithContext(ctx context.Context) VdomlinkOutput {
	return o
}

// VDOM link name. On FortiOS versions 6.2.0-6.4.0: maximum = 8 characters. On FortiOS versions >= 6.4.1: maximum = 11 characters.
func (o VdomlinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomlink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// VDOM link type: PPP or Ethernet.
func (o VdomlinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomlink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Virtual cluster. Valid values: `vcluster1`, `vcluster2`.
func (o VdomlinkOutput) Vcluster() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomlink) pulumi.StringOutput { return v.Vcluster }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o VdomlinkOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomlink) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type VdomlinkArrayOutput struct{ *pulumi.OutputState }

func (VdomlinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vdomlink)(nil)).Elem()
}

func (o VdomlinkArrayOutput) ToVdomlinkArrayOutput() VdomlinkArrayOutput {
	return o
}

func (o VdomlinkArrayOutput) ToVdomlinkArrayOutputWithContext(ctx context.Context) VdomlinkArrayOutput {
	return o
}

func (o VdomlinkArrayOutput) Index(i pulumi.IntInput) VdomlinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vdomlink {
		return vs[0].([]*Vdomlink)[vs[1].(int)]
	}).(VdomlinkOutput)
}

type VdomlinkMapOutput struct{ *pulumi.OutputState }

func (VdomlinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vdomlink)(nil)).Elem()
}

func (o VdomlinkMapOutput) ToVdomlinkMapOutput() VdomlinkMapOutput {
	return o
}

func (o VdomlinkMapOutput) ToVdomlinkMapOutputWithContext(ctx context.Context) VdomlinkMapOutput {
	return o
}

func (o VdomlinkMapOutput) MapIndex(k pulumi.StringInput) VdomlinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vdomlink {
		return vs[0].(map[string]*Vdomlink)[vs[1].(string)]
	}).(VdomlinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VdomlinkInput)(nil)).Elem(), &Vdomlink{})
	pulumi.RegisterInputType(reflect.TypeOf((*VdomlinkArrayInput)(nil)).Elem(), VdomlinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VdomlinkMapInput)(nil)).Elem(), VdomlinkMap{})
	pulumi.RegisterOutputType(VdomlinkOutput{})
	pulumi.RegisterOutputType(VdomlinkArrayOutput{})
	pulumi.RegisterOutputType(VdomlinkMapOutput{})
}
