// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure custom Internet Service group.
//
// ## Import
//
// Firewall InternetServiceCustomGroup can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/internetservicecustomgroup:Internetservicecustomgroup labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/internetservicecustomgroup:Internetservicecustomgroup labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Internetservicecustomgroup struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Custom Internet Service group members. The structure of `member` block is documented below.
	Members InternetservicecustomgroupMemberArrayOutput `pulumi:"members"`
	// Custom Internet Service group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewInternetservicecustomgroup registers a new resource with the given unique name, arguments, and options.
func NewInternetservicecustomgroup(ctx *pulumi.Context,
	name string, args *InternetservicecustomgroupArgs, opts ...pulumi.ResourceOption) (*Internetservicecustomgroup, error) {
	if args == nil {
		args = &InternetservicecustomgroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Internetservicecustomgroup
	err := ctx.RegisterResource("fortios:firewall/internetservicecustomgroup:Internetservicecustomgroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternetservicecustomgroup gets an existing Internetservicecustomgroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternetservicecustomgroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternetservicecustomgroupState, opts ...pulumi.ResourceOption) (*Internetservicecustomgroup, error) {
	var resource Internetservicecustomgroup
	err := ctx.ReadResource("fortios:firewall/internetservicecustomgroup:Internetservicecustomgroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Internetservicecustomgroup resources.
type internetservicecustomgroupState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Custom Internet Service group members. The structure of `member` block is documented below.
	Members []InternetservicecustomgroupMember `pulumi:"members"`
	// Custom Internet Service group name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type InternetservicecustomgroupState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Custom Internet Service group members. The structure of `member` block is documented below.
	Members InternetservicecustomgroupMemberArrayInput
	// Custom Internet Service group name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetservicecustomgroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*internetservicecustomgroupState)(nil)).Elem()
}

type internetservicecustomgroupArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Custom Internet Service group members. The structure of `member` block is documented below.
	Members []InternetservicecustomgroupMember `pulumi:"members"`
	// Custom Internet Service group name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Internetservicecustomgroup resource.
type InternetservicecustomgroupArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Custom Internet Service group members. The structure of `member` block is documented below.
	Members InternetservicecustomgroupMemberArrayInput
	// Custom Internet Service group name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetservicecustomgroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internetservicecustomgroupArgs)(nil)).Elem()
}

type InternetservicecustomgroupInput interface {
	pulumi.Input

	ToInternetservicecustomgroupOutput() InternetservicecustomgroupOutput
	ToInternetservicecustomgroupOutputWithContext(ctx context.Context) InternetservicecustomgroupOutput
}

func (*Internetservicecustomgroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetservicecustomgroup)(nil)).Elem()
}

func (i *Internetservicecustomgroup) ToInternetservicecustomgroupOutput() InternetservicecustomgroupOutput {
	return i.ToInternetservicecustomgroupOutputWithContext(context.Background())
}

func (i *Internetservicecustomgroup) ToInternetservicecustomgroupOutputWithContext(ctx context.Context) InternetservicecustomgroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetservicecustomgroupOutput)
}

// InternetservicecustomgroupArrayInput is an input type that accepts InternetservicecustomgroupArray and InternetservicecustomgroupArrayOutput values.
// You can construct a concrete instance of `InternetservicecustomgroupArrayInput` via:
//
//	InternetservicecustomgroupArray{ InternetservicecustomgroupArgs{...} }
type InternetservicecustomgroupArrayInput interface {
	pulumi.Input

	ToInternetservicecustomgroupArrayOutput() InternetservicecustomgroupArrayOutput
	ToInternetservicecustomgroupArrayOutputWithContext(context.Context) InternetservicecustomgroupArrayOutput
}

type InternetservicecustomgroupArray []InternetservicecustomgroupInput

func (InternetservicecustomgroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetservicecustomgroup)(nil)).Elem()
}

func (i InternetservicecustomgroupArray) ToInternetservicecustomgroupArrayOutput() InternetservicecustomgroupArrayOutput {
	return i.ToInternetservicecustomgroupArrayOutputWithContext(context.Background())
}

func (i InternetservicecustomgroupArray) ToInternetservicecustomgroupArrayOutputWithContext(ctx context.Context) InternetservicecustomgroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetservicecustomgroupArrayOutput)
}

// InternetservicecustomgroupMapInput is an input type that accepts InternetservicecustomgroupMap and InternetservicecustomgroupMapOutput values.
// You can construct a concrete instance of `InternetservicecustomgroupMapInput` via:
//
//	InternetservicecustomgroupMap{ "key": InternetservicecustomgroupArgs{...} }
type InternetservicecustomgroupMapInput interface {
	pulumi.Input

	ToInternetservicecustomgroupMapOutput() InternetservicecustomgroupMapOutput
	ToInternetservicecustomgroupMapOutputWithContext(context.Context) InternetservicecustomgroupMapOutput
}

type InternetservicecustomgroupMap map[string]InternetservicecustomgroupInput

func (InternetservicecustomgroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetservicecustomgroup)(nil)).Elem()
}

func (i InternetservicecustomgroupMap) ToInternetservicecustomgroupMapOutput() InternetservicecustomgroupMapOutput {
	return i.ToInternetservicecustomgroupMapOutputWithContext(context.Background())
}

func (i InternetservicecustomgroupMap) ToInternetservicecustomgroupMapOutputWithContext(ctx context.Context) InternetservicecustomgroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetservicecustomgroupMapOutput)
}

type InternetservicecustomgroupOutput struct{ *pulumi.OutputState }

func (InternetservicecustomgroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetservicecustomgroup)(nil)).Elem()
}

func (o InternetservicecustomgroupOutput) ToInternetservicecustomgroupOutput() InternetservicecustomgroupOutput {
	return o
}

func (o InternetservicecustomgroupOutput) ToInternetservicecustomgroupOutputWithContext(ctx context.Context) InternetservicecustomgroupOutput {
	return o
}

// Comment.
func (o InternetservicecustomgroupOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Internetservicecustomgroup) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o InternetservicecustomgroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Internetservicecustomgroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Custom Internet Service group members. The structure of `member` block is documented below.
func (o InternetservicecustomgroupOutput) Members() InternetservicecustomgroupMemberArrayOutput {
	return o.ApplyT(func(v *Internetservicecustomgroup) InternetservicecustomgroupMemberArrayOutput { return v.Members }).(InternetservicecustomgroupMemberArrayOutput)
}

// Custom Internet Service group name.
func (o InternetservicecustomgroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Internetservicecustomgroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o InternetservicecustomgroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Internetservicecustomgroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type InternetservicecustomgroupArrayOutput struct{ *pulumi.OutputState }

func (InternetservicecustomgroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetservicecustomgroup)(nil)).Elem()
}

func (o InternetservicecustomgroupArrayOutput) ToInternetservicecustomgroupArrayOutput() InternetservicecustomgroupArrayOutput {
	return o
}

func (o InternetservicecustomgroupArrayOutput) ToInternetservicecustomgroupArrayOutputWithContext(ctx context.Context) InternetservicecustomgroupArrayOutput {
	return o
}

func (o InternetservicecustomgroupArrayOutput) Index(i pulumi.IntInput) InternetservicecustomgroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Internetservicecustomgroup {
		return vs[0].([]*Internetservicecustomgroup)[vs[1].(int)]
	}).(InternetservicecustomgroupOutput)
}

type InternetservicecustomgroupMapOutput struct{ *pulumi.OutputState }

func (InternetservicecustomgroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetservicecustomgroup)(nil)).Elem()
}

func (o InternetservicecustomgroupMapOutput) ToInternetservicecustomgroupMapOutput() InternetservicecustomgroupMapOutput {
	return o
}

func (o InternetservicecustomgroupMapOutput) ToInternetservicecustomgroupMapOutputWithContext(ctx context.Context) InternetservicecustomgroupMapOutput {
	return o
}

func (o InternetservicecustomgroupMapOutput) MapIndex(k pulumi.StringInput) InternetservicecustomgroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Internetservicecustomgroup {
		return vs[0].(map[string]*Internetservicecustomgroup)[vs[1].(string)]
	}).(InternetservicecustomgroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InternetservicecustomgroupInput)(nil)).Elem(), &Internetservicecustomgroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetservicecustomgroupArrayInput)(nil)).Elem(), InternetservicecustomgroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetservicecustomgroupMapInput)(nil)).Elem(), InternetservicecustomgroupMap{})
	pulumi.RegisterOutputType(InternetservicecustomgroupOutput{})
	pulumi.RegisterOutputType(InternetservicecustomgroupArrayOutput{})
	pulumi.RegisterOutputType(InternetservicecustomgroupMapOutput{})
}
