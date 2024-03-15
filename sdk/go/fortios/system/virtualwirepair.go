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

// Configure virtual wire pairs.
//
// ## Import
//
// System VirtualWirePair can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/virtualwirepair:Virtualwirepair labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/virtualwirepair:Virtualwirepair labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Virtualwirepair struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
	Members VirtualwirepairMemberArrayOutput `pulumi:"members"`
	// Virtual-wire-pair name. Must be a unique interface name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Set VLAN filters.
	VlanFilter pulumi.StringOutput `pulumi:"vlanFilter"`
	// Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
	WildcardVlan pulumi.StringOutput `pulumi:"wildcardVlan"`
}

// NewVirtualwirepair registers a new resource with the given unique name, arguments, and options.
func NewVirtualwirepair(ctx *pulumi.Context,
	name string, args *VirtualwirepairArgs, opts ...pulumi.ResourceOption) (*Virtualwirepair, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Virtualwirepair
	err := ctx.RegisterResource("fortios:system/virtualwirepair:Virtualwirepair", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualwirepair gets an existing Virtualwirepair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualwirepair(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualwirepairState, opts ...pulumi.ResourceOption) (*Virtualwirepair, error) {
	var resource Virtualwirepair
	err := ctx.ReadResource("fortios:system/virtualwirepair:Virtualwirepair", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Virtualwirepair resources.
type virtualwirepairState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
	Members []VirtualwirepairMember `pulumi:"members"`
	// Virtual-wire-pair name. Must be a unique interface name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Set VLAN filters.
	VlanFilter *string `pulumi:"vlanFilter"`
	// Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
	WildcardVlan *string `pulumi:"wildcardVlan"`
}

type VirtualwirepairState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
	Members VirtualwirepairMemberArrayInput
	// Virtual-wire-pair name. Must be a unique interface name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Set VLAN filters.
	VlanFilter pulumi.StringPtrInput
	// Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
	WildcardVlan pulumi.StringPtrInput
}

func (VirtualwirepairState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualwirepairState)(nil)).Elem()
}

type virtualwirepairArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
	Members []VirtualwirepairMember `pulumi:"members"`
	// Virtual-wire-pair name. Must be a unique interface name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Set VLAN filters.
	VlanFilter *string `pulumi:"vlanFilter"`
	// Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
	WildcardVlan *string `pulumi:"wildcardVlan"`
}

// The set of arguments for constructing a Virtualwirepair resource.
type VirtualwirepairArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
	Members VirtualwirepairMemberArrayInput
	// Virtual-wire-pair name. Must be a unique interface name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Set VLAN filters.
	VlanFilter pulumi.StringPtrInput
	// Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
	WildcardVlan pulumi.StringPtrInput
}

func (VirtualwirepairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualwirepairArgs)(nil)).Elem()
}

type VirtualwirepairInput interface {
	pulumi.Input

	ToVirtualwirepairOutput() VirtualwirepairOutput
	ToVirtualwirepairOutputWithContext(ctx context.Context) VirtualwirepairOutput
}

func (*Virtualwirepair) ElementType() reflect.Type {
	return reflect.TypeOf((**Virtualwirepair)(nil)).Elem()
}

func (i *Virtualwirepair) ToVirtualwirepairOutput() VirtualwirepairOutput {
	return i.ToVirtualwirepairOutputWithContext(context.Background())
}

func (i *Virtualwirepair) ToVirtualwirepairOutputWithContext(ctx context.Context) VirtualwirepairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualwirepairOutput)
}

// VirtualwirepairArrayInput is an input type that accepts VirtualwirepairArray and VirtualwirepairArrayOutput values.
// You can construct a concrete instance of `VirtualwirepairArrayInput` via:
//
//	VirtualwirepairArray{ VirtualwirepairArgs{...} }
type VirtualwirepairArrayInput interface {
	pulumi.Input

	ToVirtualwirepairArrayOutput() VirtualwirepairArrayOutput
	ToVirtualwirepairArrayOutputWithContext(context.Context) VirtualwirepairArrayOutput
}

type VirtualwirepairArray []VirtualwirepairInput

func (VirtualwirepairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Virtualwirepair)(nil)).Elem()
}

func (i VirtualwirepairArray) ToVirtualwirepairArrayOutput() VirtualwirepairArrayOutput {
	return i.ToVirtualwirepairArrayOutputWithContext(context.Background())
}

func (i VirtualwirepairArray) ToVirtualwirepairArrayOutputWithContext(ctx context.Context) VirtualwirepairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualwirepairArrayOutput)
}

// VirtualwirepairMapInput is an input type that accepts VirtualwirepairMap and VirtualwirepairMapOutput values.
// You can construct a concrete instance of `VirtualwirepairMapInput` via:
//
//	VirtualwirepairMap{ "key": VirtualwirepairArgs{...} }
type VirtualwirepairMapInput interface {
	pulumi.Input

	ToVirtualwirepairMapOutput() VirtualwirepairMapOutput
	ToVirtualwirepairMapOutputWithContext(context.Context) VirtualwirepairMapOutput
}

type VirtualwirepairMap map[string]VirtualwirepairInput

func (VirtualwirepairMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Virtualwirepair)(nil)).Elem()
}

func (i VirtualwirepairMap) ToVirtualwirepairMapOutput() VirtualwirepairMapOutput {
	return i.ToVirtualwirepairMapOutputWithContext(context.Background())
}

func (i VirtualwirepairMap) ToVirtualwirepairMapOutputWithContext(ctx context.Context) VirtualwirepairMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualwirepairMapOutput)
}

type VirtualwirepairOutput struct{ *pulumi.OutputState }

func (VirtualwirepairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Virtualwirepair)(nil)).Elem()
}

func (o VirtualwirepairOutput) ToVirtualwirepairOutput() VirtualwirepairOutput {
	return o
}

func (o VirtualwirepairOutput) ToVirtualwirepairOutputWithContext(ctx context.Context) VirtualwirepairOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o VirtualwirepairOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Virtualwirepair) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
func (o VirtualwirepairOutput) Members() VirtualwirepairMemberArrayOutput {
	return o.ApplyT(func(v *Virtualwirepair) VirtualwirepairMemberArrayOutput { return v.Members }).(VirtualwirepairMemberArrayOutput)
}

// Virtual-wire-pair name. Must be a unique interface name.
func (o VirtualwirepairOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Virtualwirepair) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o VirtualwirepairOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Virtualwirepair) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Set VLAN filters.
func (o VirtualwirepairOutput) VlanFilter() pulumi.StringOutput {
	return o.ApplyT(func(v *Virtualwirepair) pulumi.StringOutput { return v.VlanFilter }).(pulumi.StringOutput)
}

// Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
func (o VirtualwirepairOutput) WildcardVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *Virtualwirepair) pulumi.StringOutput { return v.WildcardVlan }).(pulumi.StringOutput)
}

type VirtualwirepairArrayOutput struct{ *pulumi.OutputState }

func (VirtualwirepairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Virtualwirepair)(nil)).Elem()
}

func (o VirtualwirepairArrayOutput) ToVirtualwirepairArrayOutput() VirtualwirepairArrayOutput {
	return o
}

func (o VirtualwirepairArrayOutput) ToVirtualwirepairArrayOutputWithContext(ctx context.Context) VirtualwirepairArrayOutput {
	return o
}

func (o VirtualwirepairArrayOutput) Index(i pulumi.IntInput) VirtualwirepairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Virtualwirepair {
		return vs[0].([]*Virtualwirepair)[vs[1].(int)]
	}).(VirtualwirepairOutput)
}

type VirtualwirepairMapOutput struct{ *pulumi.OutputState }

func (VirtualwirepairMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Virtualwirepair)(nil)).Elem()
}

func (o VirtualwirepairMapOutput) ToVirtualwirepairMapOutput() VirtualwirepairMapOutput {
	return o
}

func (o VirtualwirepairMapOutput) ToVirtualwirepairMapOutputWithContext(ctx context.Context) VirtualwirepairMapOutput {
	return o
}

func (o VirtualwirepairMapOutput) MapIndex(k pulumi.StringInput) VirtualwirepairOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Virtualwirepair {
		return vs[0].(map[string]*Virtualwirepair)[vs[1].(string)]
	}).(VirtualwirepairOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualwirepairInput)(nil)).Elem(), &Virtualwirepair{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualwirepairArrayInput)(nil)).Elem(), VirtualwirepairArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualwirepairMapInput)(nil)).Elem(), VirtualwirepairMap{})
	pulumi.RegisterOutputType(VirtualwirepairOutput{})
	pulumi.RegisterOutputType(VirtualwirepairArrayOutput{})
	pulumi.RegisterOutputType(VirtualwirepairMapOutput{})
}