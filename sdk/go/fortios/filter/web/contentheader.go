// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure content types used by Web filter.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/filter"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := filter.NewContentheader(ctx, "trname", &filter.ContentheaderArgs{
//				Fosid: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Webfilter ContentHeader can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:filter/web/contentheader:Contentheader labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:filter/web/contentheader:Contentheader labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Contentheader struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Configure content types used by web filter. The structure of `entries` block is documented below.
	Entries ContentheaderEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewContentheader registers a new resource with the given unique name, arguments, and options.
func NewContentheader(ctx *pulumi.Context,
	name string, args *ContentheaderArgs, opts ...pulumi.ResourceOption) (*Contentheader, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Contentheader
	err := ctx.RegisterResource("fortios:filter/web/contentheader:Contentheader", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContentheader gets an existing Contentheader resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContentheader(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContentheaderState, opts ...pulumi.ResourceOption) (*Contentheader, error) {
	var resource Contentheader
	err := ctx.ReadResource("fortios:filter/web/contentheader:Contentheader", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Contentheader resources.
type contentheaderState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure content types used by web filter. The structure of `entries` block is documented below.
	Entries []ContentheaderEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ContentheaderState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure content types used by web filter. The structure of `entries` block is documented below.
	Entries ContentheaderEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ContentheaderState) ElementType() reflect.Type {
	return reflect.TypeOf((*contentheaderState)(nil)).Elem()
}

type contentheaderArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure content types used by web filter. The structure of `entries` block is documented below.
	Entries []ContentheaderEntry `pulumi:"entries"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Contentheader resource.
type ContentheaderArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure content types used by web filter. The structure of `entries` block is documented below.
	Entries ContentheaderEntryArrayInput
	// ID.
	Fosid pulumi.IntInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ContentheaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contentheaderArgs)(nil)).Elem()
}

type ContentheaderInput interface {
	pulumi.Input

	ToContentheaderOutput() ContentheaderOutput
	ToContentheaderOutputWithContext(ctx context.Context) ContentheaderOutput
}

func (*Contentheader) ElementType() reflect.Type {
	return reflect.TypeOf((**Contentheader)(nil)).Elem()
}

func (i *Contentheader) ToContentheaderOutput() ContentheaderOutput {
	return i.ToContentheaderOutputWithContext(context.Background())
}

func (i *Contentheader) ToContentheaderOutputWithContext(ctx context.Context) ContentheaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentheaderOutput)
}

// ContentheaderArrayInput is an input type that accepts ContentheaderArray and ContentheaderArrayOutput values.
// You can construct a concrete instance of `ContentheaderArrayInput` via:
//
//	ContentheaderArray{ ContentheaderArgs{...} }
type ContentheaderArrayInput interface {
	pulumi.Input

	ToContentheaderArrayOutput() ContentheaderArrayOutput
	ToContentheaderArrayOutputWithContext(context.Context) ContentheaderArrayOutput
}

type ContentheaderArray []ContentheaderInput

func (ContentheaderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Contentheader)(nil)).Elem()
}

func (i ContentheaderArray) ToContentheaderArrayOutput() ContentheaderArrayOutput {
	return i.ToContentheaderArrayOutputWithContext(context.Background())
}

func (i ContentheaderArray) ToContentheaderArrayOutputWithContext(ctx context.Context) ContentheaderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentheaderArrayOutput)
}

// ContentheaderMapInput is an input type that accepts ContentheaderMap and ContentheaderMapOutput values.
// You can construct a concrete instance of `ContentheaderMapInput` via:
//
//	ContentheaderMap{ "key": ContentheaderArgs{...} }
type ContentheaderMapInput interface {
	pulumi.Input

	ToContentheaderMapOutput() ContentheaderMapOutput
	ToContentheaderMapOutputWithContext(context.Context) ContentheaderMapOutput
}

type ContentheaderMap map[string]ContentheaderInput

func (ContentheaderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Contentheader)(nil)).Elem()
}

func (i ContentheaderMap) ToContentheaderMapOutput() ContentheaderMapOutput {
	return i.ToContentheaderMapOutputWithContext(context.Background())
}

func (i ContentheaderMap) ToContentheaderMapOutputWithContext(ctx context.Context) ContentheaderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentheaderMapOutput)
}

type ContentheaderOutput struct{ *pulumi.OutputState }

func (ContentheaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Contentheader)(nil)).Elem()
}

func (o ContentheaderOutput) ToContentheaderOutput() ContentheaderOutput {
	return o
}

func (o ContentheaderOutput) ToContentheaderOutputWithContext(ctx context.Context) ContentheaderOutput {
	return o
}

// Optional comments.
func (o ContentheaderOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contentheader) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ContentheaderOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contentheader) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Configure content types used by web filter. The structure of `entries` block is documented below.
func (o ContentheaderOutput) Entries() ContentheaderEntryArrayOutput {
	return o.ApplyT(func(v *Contentheader) ContentheaderEntryArrayOutput { return v.Entries }).(ContentheaderEntryArrayOutput)
}

// ID.
func (o ContentheaderOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Contentheader) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Name of table.
func (o ContentheaderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Contentheader) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ContentheaderOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Contentheader) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ContentheaderArrayOutput struct{ *pulumi.OutputState }

func (ContentheaderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Contentheader)(nil)).Elem()
}

func (o ContentheaderArrayOutput) ToContentheaderArrayOutput() ContentheaderArrayOutput {
	return o
}

func (o ContentheaderArrayOutput) ToContentheaderArrayOutputWithContext(ctx context.Context) ContentheaderArrayOutput {
	return o
}

func (o ContentheaderArrayOutput) Index(i pulumi.IntInput) ContentheaderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Contentheader {
		return vs[0].([]*Contentheader)[vs[1].(int)]
	}).(ContentheaderOutput)
}

type ContentheaderMapOutput struct{ *pulumi.OutputState }

func (ContentheaderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Contentheader)(nil)).Elem()
}

func (o ContentheaderMapOutput) ToContentheaderMapOutput() ContentheaderMapOutput {
	return o
}

func (o ContentheaderMapOutput) ToContentheaderMapOutputWithContext(ctx context.Context) ContentheaderMapOutput {
	return o
}

func (o ContentheaderMapOutput) MapIndex(k pulumi.StringInput) ContentheaderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Contentheader {
		return vs[0].(map[string]*Contentheader)[vs[1].(string)]
	}).(ContentheaderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContentheaderInput)(nil)).Elem(), &Contentheader{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentheaderArrayInput)(nil)).Elem(), ContentheaderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentheaderMapInput)(nil)).Elem(), ContentheaderMap{})
	pulumi.RegisterOutputType(ContentheaderOutput{})
	pulumi.RegisterOutputType(ContentheaderArrayOutput{})
	pulumi.RegisterOutputType(ContentheaderMapOutput{})
}
