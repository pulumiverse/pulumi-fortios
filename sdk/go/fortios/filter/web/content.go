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

// Configure Web filter banned word table.
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
//			_, err := filter.NewContent(ctx, "trname", &filter.ContentArgs{
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
// Webfilter Content can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:filter/web/content:Content labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:filter/web/content:Content labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Content struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Configure banned word entries. The structure of `entries` block is documented below.
	Entries ContentEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewContent registers a new resource with the given unique name, arguments, and options.
func NewContent(ctx *pulumi.Context,
	name string, args *ContentArgs, opts ...pulumi.ResourceOption) (*Content, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Content
	err := ctx.RegisterResource("fortios:filter/web/content:Content", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContent gets an existing Content resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContentState, opts ...pulumi.ResourceOption) (*Content, error) {
	var resource Content
	err := ctx.ReadResource("fortios:filter/web/content:Content", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Content resources.
type contentState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure banned word entries. The structure of `entries` block is documented below.
	Entries []ContentEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ContentState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure banned word entries. The structure of `entries` block is documented below.
	Entries ContentEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ContentState) ElementType() reflect.Type {
	return reflect.TypeOf((*contentState)(nil)).Elem()
}

type contentArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure banned word entries. The structure of `entries` block is documented below.
	Entries []ContentEntry `pulumi:"entries"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Content resource.
type ContentArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure banned word entries. The structure of `entries` block is documented below.
	Entries ContentEntryArrayInput
	// ID.
	Fosid pulumi.IntInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contentArgs)(nil)).Elem()
}

type ContentInput interface {
	pulumi.Input

	ToContentOutput() ContentOutput
	ToContentOutputWithContext(ctx context.Context) ContentOutput
}

func (*Content) ElementType() reflect.Type {
	return reflect.TypeOf((**Content)(nil)).Elem()
}

func (i *Content) ToContentOutput() ContentOutput {
	return i.ToContentOutputWithContext(context.Background())
}

func (i *Content) ToContentOutputWithContext(ctx context.Context) ContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentOutput)
}

// ContentArrayInput is an input type that accepts ContentArray and ContentArrayOutput values.
// You can construct a concrete instance of `ContentArrayInput` via:
//
//	ContentArray{ ContentArgs{...} }
type ContentArrayInput interface {
	pulumi.Input

	ToContentArrayOutput() ContentArrayOutput
	ToContentArrayOutputWithContext(context.Context) ContentArrayOutput
}

type ContentArray []ContentInput

func (ContentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Content)(nil)).Elem()
}

func (i ContentArray) ToContentArrayOutput() ContentArrayOutput {
	return i.ToContentArrayOutputWithContext(context.Background())
}

func (i ContentArray) ToContentArrayOutputWithContext(ctx context.Context) ContentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentArrayOutput)
}

// ContentMapInput is an input type that accepts ContentMap and ContentMapOutput values.
// You can construct a concrete instance of `ContentMapInput` via:
//
//	ContentMap{ "key": ContentArgs{...} }
type ContentMapInput interface {
	pulumi.Input

	ToContentMapOutput() ContentMapOutput
	ToContentMapOutputWithContext(context.Context) ContentMapOutput
}

type ContentMap map[string]ContentInput

func (ContentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Content)(nil)).Elem()
}

func (i ContentMap) ToContentMapOutput() ContentMapOutput {
	return i.ToContentMapOutputWithContext(context.Background())
}

func (i ContentMap) ToContentMapOutputWithContext(ctx context.Context) ContentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentMapOutput)
}

type ContentOutput struct{ *pulumi.OutputState }

func (ContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Content)(nil)).Elem()
}

func (o ContentOutput) ToContentOutput() ContentOutput {
	return o
}

func (o ContentOutput) ToContentOutputWithContext(ctx context.Context) ContentOutput {
	return o
}

// Optional comments.
func (o ContentOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Content) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ContentOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Content) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Configure banned word entries. The structure of `entries` block is documented below.
func (o ContentOutput) Entries() ContentEntryArrayOutput {
	return o.ApplyT(func(v *Content) ContentEntryArrayOutput { return v.Entries }).(ContentEntryArrayOutput)
}

// ID.
func (o ContentOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Content) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Name of table.
func (o ContentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Content) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ContentOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Content) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ContentArrayOutput struct{ *pulumi.OutputState }

func (ContentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Content)(nil)).Elem()
}

func (o ContentArrayOutput) ToContentArrayOutput() ContentArrayOutput {
	return o
}

func (o ContentArrayOutput) ToContentArrayOutputWithContext(ctx context.Context) ContentArrayOutput {
	return o
}

func (o ContentArrayOutput) Index(i pulumi.IntInput) ContentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Content {
		return vs[0].([]*Content)[vs[1].(int)]
	}).(ContentOutput)
}

type ContentMapOutput struct{ *pulumi.OutputState }

func (ContentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Content)(nil)).Elem()
}

func (o ContentMapOutput) ToContentMapOutput() ContentMapOutput {
	return o
}

func (o ContentMapOutput) ToContentMapOutputWithContext(ctx context.Context) ContentMapOutput {
	return o
}

func (o ContentMapOutput) MapIndex(k pulumi.StringInput) ContentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Content {
		return vs[0].(map[string]*Content)[vs[1].(string)]
	}).(ContentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContentInput)(nil)).Elem(), &Content{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentArrayInput)(nil)).Elem(), ContentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentMapInput)(nil)).Elem(), ContentMap{})
	pulumi.RegisterOutputType(ContentOutput{})
	pulumi.RegisterOutputType(ContentArrayOutput{})
	pulumi.RegisterOutputType(ContentMapOutput{})
}
