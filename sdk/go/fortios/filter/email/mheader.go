// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package email

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure AntiSpam MIME header. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// Emailfilter Mheader can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:filter/email/mheader:Mheader labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:filter/email/mheader:Mheader labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Mheader struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries MheaderEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewMheader registers a new resource with the given unique name, arguments, and options.
func NewMheader(ctx *pulumi.Context,
	name string, args *MheaderArgs, opts ...pulumi.ResourceOption) (*Mheader, error) {
	if args == nil {
		args = &MheaderArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Mheader
	err := ctx.RegisterResource("fortios:filter/email/mheader:Mheader", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMheader gets an existing Mheader resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMheader(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MheaderState, opts ...pulumi.ResourceOption) (*Mheader, error) {
	var resource Mheader
	err := ctx.ReadResource("fortios:filter/email/mheader:Mheader", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Mheader resources.
type mheaderState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries []MheaderEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type MheaderState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries MheaderEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (MheaderState) ElementType() reflect.Type {
	return reflect.TypeOf((*mheaderState)(nil)).Elem()
}

type mheaderArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries []MheaderEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Mheader resource.
type MheaderArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries MheaderEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (MheaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mheaderArgs)(nil)).Elem()
}

type MheaderInput interface {
	pulumi.Input

	ToMheaderOutput() MheaderOutput
	ToMheaderOutputWithContext(ctx context.Context) MheaderOutput
}

func (*Mheader) ElementType() reflect.Type {
	return reflect.TypeOf((**Mheader)(nil)).Elem()
}

func (i *Mheader) ToMheaderOutput() MheaderOutput {
	return i.ToMheaderOutputWithContext(context.Background())
}

func (i *Mheader) ToMheaderOutputWithContext(ctx context.Context) MheaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MheaderOutput)
}

// MheaderArrayInput is an input type that accepts MheaderArray and MheaderArrayOutput values.
// You can construct a concrete instance of `MheaderArrayInput` via:
//
//	MheaderArray{ MheaderArgs{...} }
type MheaderArrayInput interface {
	pulumi.Input

	ToMheaderArrayOutput() MheaderArrayOutput
	ToMheaderArrayOutputWithContext(context.Context) MheaderArrayOutput
}

type MheaderArray []MheaderInput

func (MheaderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mheader)(nil)).Elem()
}

func (i MheaderArray) ToMheaderArrayOutput() MheaderArrayOutput {
	return i.ToMheaderArrayOutputWithContext(context.Background())
}

func (i MheaderArray) ToMheaderArrayOutputWithContext(ctx context.Context) MheaderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MheaderArrayOutput)
}

// MheaderMapInput is an input type that accepts MheaderMap and MheaderMapOutput values.
// You can construct a concrete instance of `MheaderMapInput` via:
//
//	MheaderMap{ "key": MheaderArgs{...} }
type MheaderMapInput interface {
	pulumi.Input

	ToMheaderMapOutput() MheaderMapOutput
	ToMheaderMapOutputWithContext(context.Context) MheaderMapOutput
}

type MheaderMap map[string]MheaderInput

func (MheaderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mheader)(nil)).Elem()
}

func (i MheaderMap) ToMheaderMapOutput() MheaderMapOutput {
	return i.ToMheaderMapOutputWithContext(context.Background())
}

func (i MheaderMap) ToMheaderMapOutputWithContext(ctx context.Context) MheaderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MheaderMapOutput)
}

type MheaderOutput struct{ *pulumi.OutputState }

func (MheaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Mheader)(nil)).Elem()
}

func (o MheaderOutput) ToMheaderOutput() MheaderOutput {
	return o
}

func (o MheaderOutput) ToMheaderOutputWithContext(ctx context.Context) MheaderOutput {
	return o
}

// Optional comments.
func (o MheaderOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mheader) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o MheaderOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mheader) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Spam filter mime header content. The structure of `entries` block is documented below.
func (o MheaderOutput) Entries() MheaderEntryArrayOutput {
	return o.ApplyT(func(v *Mheader) MheaderEntryArrayOutput { return v.Entries }).(MheaderEntryArrayOutput)
}

// ID.
func (o MheaderOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Mheader) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Name of table.
func (o MheaderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Mheader) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o MheaderOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mheader) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type MheaderArrayOutput struct{ *pulumi.OutputState }

func (MheaderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mheader)(nil)).Elem()
}

func (o MheaderArrayOutput) ToMheaderArrayOutput() MheaderArrayOutput {
	return o
}

func (o MheaderArrayOutput) ToMheaderArrayOutputWithContext(ctx context.Context) MheaderArrayOutput {
	return o
}

func (o MheaderArrayOutput) Index(i pulumi.IntInput) MheaderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Mheader {
		return vs[0].([]*Mheader)[vs[1].(int)]
	}).(MheaderOutput)
}

type MheaderMapOutput struct{ *pulumi.OutputState }

func (MheaderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mheader)(nil)).Elem()
}

func (o MheaderMapOutput) ToMheaderMapOutput() MheaderMapOutput {
	return o
}

func (o MheaderMapOutput) ToMheaderMapOutputWithContext(ctx context.Context) MheaderMapOutput {
	return o
}

func (o MheaderMapOutput) MapIndex(k pulumi.StringInput) MheaderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Mheader {
		return vs[0].(map[string]*Mheader)[vs[1].(string)]
	}).(MheaderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MheaderInput)(nil)).Elem(), &Mheader{})
	pulumi.RegisterInputType(reflect.TypeOf((*MheaderArrayInput)(nil)).Elem(), MheaderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MheaderMapInput)(nil)).Elem(), MheaderMap{})
	pulumi.RegisterOutputType(MheaderOutput{})
	pulumi.RegisterOutputType(MheaderArrayOutput{})
	pulumi.RegisterOutputType(MheaderMapOutput{})
}
