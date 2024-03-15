// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure VLAN policy to be applied on the managed FortiSwitch ports through port-policy. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// SwitchController VlanPolicy can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/vlanpolicy:Vlanpolicy labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/vlanpolicy:Vlanpolicy labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Vlanpolicy struct {
	pulumi.CustomResourceState

	// Allowed VLANs to be applied when using this VLAN policy. The structure of `allowedVlans` block is documented below.
	AllowedVlans VlanpolicyAllowedVlanArrayOutput `pulumi:"allowedVlans"`
	// Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
	AllowedVlansAll pulumi.StringOutput `pulumi:"allowedVlansAll"`
	// Description for the VLAN policy.
	Description pulumi.StringOutput `pulumi:"description"`
	// Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
	DiscardMode pulumi.StringOutput `pulumi:"discardMode"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// FortiLink interface for which this VLAN policy belongs to.
	Fortilink pulumi.StringOutput `pulumi:"fortilink"`
	// VLAN policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Untagged VLANs to be applied when using this VLAN policy. The structure of `untaggedVlans` block is documented below.
	UntaggedVlans VlanpolicyUntaggedVlanArrayOutput `pulumi:"untaggedVlans"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Native VLAN to be applied when using this VLAN policy.
	Vlan pulumi.StringOutput `pulumi:"vlan"`
}

// NewVlanpolicy registers a new resource with the given unique name, arguments, and options.
func NewVlanpolicy(ctx *pulumi.Context,
	name string, args *VlanpolicyArgs, opts ...pulumi.ResourceOption) (*Vlanpolicy, error) {
	if args == nil {
		args = &VlanpolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vlanpolicy
	err := ctx.RegisterResource("fortios:switchcontroller/vlanpolicy:Vlanpolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVlanpolicy gets an existing Vlanpolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVlanpolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VlanpolicyState, opts ...pulumi.ResourceOption) (*Vlanpolicy, error) {
	var resource Vlanpolicy
	err := ctx.ReadResource("fortios:switchcontroller/vlanpolicy:Vlanpolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vlanpolicy resources.
type vlanpolicyState struct {
	// Allowed VLANs to be applied when using this VLAN policy. The structure of `allowedVlans` block is documented below.
	AllowedVlans []VlanpolicyAllowedVlan `pulumi:"allowedVlans"`
	// Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
	AllowedVlansAll *string `pulumi:"allowedVlansAll"`
	// Description for the VLAN policy.
	Description *string `pulumi:"description"`
	// Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
	DiscardMode *string `pulumi:"discardMode"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiLink interface for which this VLAN policy belongs to.
	Fortilink *string `pulumi:"fortilink"`
	// VLAN policy name.
	Name *string `pulumi:"name"`
	// Untagged VLANs to be applied when using this VLAN policy. The structure of `untaggedVlans` block is documented below.
	UntaggedVlans []VlanpolicyUntaggedVlan `pulumi:"untaggedVlans"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Native VLAN to be applied when using this VLAN policy.
	Vlan *string `pulumi:"vlan"`
}

type VlanpolicyState struct {
	// Allowed VLANs to be applied when using this VLAN policy. The structure of `allowedVlans` block is documented below.
	AllowedVlans VlanpolicyAllowedVlanArrayInput
	// Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
	AllowedVlansAll pulumi.StringPtrInput
	// Description for the VLAN policy.
	Description pulumi.StringPtrInput
	// Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
	DiscardMode pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiLink interface for which this VLAN policy belongs to.
	Fortilink pulumi.StringPtrInput
	// VLAN policy name.
	Name pulumi.StringPtrInput
	// Untagged VLANs to be applied when using this VLAN policy. The structure of `untaggedVlans` block is documented below.
	UntaggedVlans VlanpolicyUntaggedVlanArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Native VLAN to be applied when using this VLAN policy.
	Vlan pulumi.StringPtrInput
}

func (VlanpolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanpolicyState)(nil)).Elem()
}

type vlanpolicyArgs struct {
	// Allowed VLANs to be applied when using this VLAN policy. The structure of `allowedVlans` block is documented below.
	AllowedVlans []VlanpolicyAllowedVlan `pulumi:"allowedVlans"`
	// Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
	AllowedVlansAll *string `pulumi:"allowedVlansAll"`
	// Description for the VLAN policy.
	Description *string `pulumi:"description"`
	// Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
	DiscardMode *string `pulumi:"discardMode"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiLink interface for which this VLAN policy belongs to.
	Fortilink *string `pulumi:"fortilink"`
	// VLAN policy name.
	Name *string `pulumi:"name"`
	// Untagged VLANs to be applied when using this VLAN policy. The structure of `untaggedVlans` block is documented below.
	UntaggedVlans []VlanpolicyUntaggedVlan `pulumi:"untaggedVlans"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Native VLAN to be applied when using this VLAN policy.
	Vlan *string `pulumi:"vlan"`
}

// The set of arguments for constructing a Vlanpolicy resource.
type VlanpolicyArgs struct {
	// Allowed VLANs to be applied when using this VLAN policy. The structure of `allowedVlans` block is documented below.
	AllowedVlans VlanpolicyAllowedVlanArrayInput
	// Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
	AllowedVlansAll pulumi.StringPtrInput
	// Description for the VLAN policy.
	Description pulumi.StringPtrInput
	// Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
	DiscardMode pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiLink interface for which this VLAN policy belongs to.
	Fortilink pulumi.StringPtrInput
	// VLAN policy name.
	Name pulumi.StringPtrInput
	// Untagged VLANs to be applied when using this VLAN policy. The structure of `untaggedVlans` block is documented below.
	UntaggedVlans VlanpolicyUntaggedVlanArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Native VLAN to be applied when using this VLAN policy.
	Vlan pulumi.StringPtrInput
}

func (VlanpolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanpolicyArgs)(nil)).Elem()
}

type VlanpolicyInput interface {
	pulumi.Input

	ToVlanpolicyOutput() VlanpolicyOutput
	ToVlanpolicyOutputWithContext(ctx context.Context) VlanpolicyOutput
}

func (*Vlanpolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**Vlanpolicy)(nil)).Elem()
}

func (i *Vlanpolicy) ToVlanpolicyOutput() VlanpolicyOutput {
	return i.ToVlanpolicyOutputWithContext(context.Background())
}

func (i *Vlanpolicy) ToVlanpolicyOutputWithContext(ctx context.Context) VlanpolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanpolicyOutput)
}

// VlanpolicyArrayInput is an input type that accepts VlanpolicyArray and VlanpolicyArrayOutput values.
// You can construct a concrete instance of `VlanpolicyArrayInput` via:
//
//	VlanpolicyArray{ VlanpolicyArgs{...} }
type VlanpolicyArrayInput interface {
	pulumi.Input

	ToVlanpolicyArrayOutput() VlanpolicyArrayOutput
	ToVlanpolicyArrayOutputWithContext(context.Context) VlanpolicyArrayOutput
}

type VlanpolicyArray []VlanpolicyInput

func (VlanpolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vlanpolicy)(nil)).Elem()
}

func (i VlanpolicyArray) ToVlanpolicyArrayOutput() VlanpolicyArrayOutput {
	return i.ToVlanpolicyArrayOutputWithContext(context.Background())
}

func (i VlanpolicyArray) ToVlanpolicyArrayOutputWithContext(ctx context.Context) VlanpolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanpolicyArrayOutput)
}

// VlanpolicyMapInput is an input type that accepts VlanpolicyMap and VlanpolicyMapOutput values.
// You can construct a concrete instance of `VlanpolicyMapInput` via:
//
//	VlanpolicyMap{ "key": VlanpolicyArgs{...} }
type VlanpolicyMapInput interface {
	pulumi.Input

	ToVlanpolicyMapOutput() VlanpolicyMapOutput
	ToVlanpolicyMapOutputWithContext(context.Context) VlanpolicyMapOutput
}

type VlanpolicyMap map[string]VlanpolicyInput

func (VlanpolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vlanpolicy)(nil)).Elem()
}

func (i VlanpolicyMap) ToVlanpolicyMapOutput() VlanpolicyMapOutput {
	return i.ToVlanpolicyMapOutputWithContext(context.Background())
}

func (i VlanpolicyMap) ToVlanpolicyMapOutputWithContext(ctx context.Context) VlanpolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanpolicyMapOutput)
}

type VlanpolicyOutput struct{ *pulumi.OutputState }

func (VlanpolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vlanpolicy)(nil)).Elem()
}

func (o VlanpolicyOutput) ToVlanpolicyOutput() VlanpolicyOutput {
	return o
}

func (o VlanpolicyOutput) ToVlanpolicyOutputWithContext(ctx context.Context) VlanpolicyOutput {
	return o
}

// Allowed VLANs to be applied when using this VLAN policy. The structure of `allowedVlans` block is documented below.
func (o VlanpolicyOutput) AllowedVlans() VlanpolicyAllowedVlanArrayOutput {
	return o.ApplyT(func(v *Vlanpolicy) VlanpolicyAllowedVlanArrayOutput { return v.AllowedVlans }).(VlanpolicyAllowedVlanArrayOutput)
}

// Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
func (o VlanpolicyOutput) AllowedVlansAll() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlanpolicy) pulumi.StringOutput { return v.AllowedVlansAll }).(pulumi.StringOutput)
}

// Description for the VLAN policy.
func (o VlanpolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlanpolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
func (o VlanpolicyOutput) DiscardMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlanpolicy) pulumi.StringOutput { return v.DiscardMode }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o VlanpolicyOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vlanpolicy) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// FortiLink interface for which this VLAN policy belongs to.
func (o VlanpolicyOutput) Fortilink() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlanpolicy) pulumi.StringOutput { return v.Fortilink }).(pulumi.StringOutput)
}

// VLAN policy name.
func (o VlanpolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlanpolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Untagged VLANs to be applied when using this VLAN policy. The structure of `untaggedVlans` block is documented below.
func (o VlanpolicyOutput) UntaggedVlans() VlanpolicyUntaggedVlanArrayOutput {
	return o.ApplyT(func(v *Vlanpolicy) VlanpolicyUntaggedVlanArrayOutput { return v.UntaggedVlans }).(VlanpolicyUntaggedVlanArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o VlanpolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vlanpolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Native VLAN to be applied when using this VLAN policy.
func (o VlanpolicyOutput) Vlan() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlanpolicy) pulumi.StringOutput { return v.Vlan }).(pulumi.StringOutput)
}

type VlanpolicyArrayOutput struct{ *pulumi.OutputState }

func (VlanpolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vlanpolicy)(nil)).Elem()
}

func (o VlanpolicyArrayOutput) ToVlanpolicyArrayOutput() VlanpolicyArrayOutput {
	return o
}

func (o VlanpolicyArrayOutput) ToVlanpolicyArrayOutputWithContext(ctx context.Context) VlanpolicyArrayOutput {
	return o
}

func (o VlanpolicyArrayOutput) Index(i pulumi.IntInput) VlanpolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vlanpolicy {
		return vs[0].([]*Vlanpolicy)[vs[1].(int)]
	}).(VlanpolicyOutput)
}

type VlanpolicyMapOutput struct{ *pulumi.OutputState }

func (VlanpolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vlanpolicy)(nil)).Elem()
}

func (o VlanpolicyMapOutput) ToVlanpolicyMapOutput() VlanpolicyMapOutput {
	return o
}

func (o VlanpolicyMapOutput) ToVlanpolicyMapOutputWithContext(ctx context.Context) VlanpolicyMapOutput {
	return o
}

func (o VlanpolicyMapOutput) MapIndex(k pulumi.StringInput) VlanpolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vlanpolicy {
		return vs[0].(map[string]*Vlanpolicy)[vs[1].(string)]
	}).(VlanpolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VlanpolicyInput)(nil)).Elem(), &Vlanpolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*VlanpolicyArrayInput)(nil)).Elem(), VlanpolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VlanpolicyMapInput)(nil)).Elem(), VlanpolicyMap{})
	pulumi.RegisterOutputType(VlanpolicyOutput{})
	pulumi.RegisterOutputType(VlanpolicyArrayOutput{})
	pulumi.RegisterOutputType(VlanpolicyMapOutput{})
}
