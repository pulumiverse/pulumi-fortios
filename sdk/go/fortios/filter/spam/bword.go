// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure AntiSpam banned word list. Applies to FortiOS Version `<= 6.2.0`.
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
//			_, err := filter.NewBword(ctx, "trname", &filter.BwordArgs{
//				Comment: pulumi.String("test"),
//				Entries: spam.BwordEntryArray{
//					&spam.BwordEntryArgs{
//						Action:      pulumi.String("clear"),
//						Language:    pulumi.String("western"),
//						Pattern:     pulumi.String("test*patten"),
//						PatternType: pulumi.String("wildcard"),
//						Score:       pulumi.Int(10),
//						Status:      pulumi.String("enable"),
//						Where:       pulumi.String("subject"),
//					},
//				},
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
// Spamfilter Bword can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:filter/spam/bword:Bword labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:filter/spam/bword:Bword labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Bword struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Spam filter banned word. The structure of `entries` block is documented below.
	Entries BwordEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewBword registers a new resource with the given unique name, arguments, and options.
func NewBword(ctx *pulumi.Context,
	name string, args *BwordArgs, opts ...pulumi.ResourceOption) (*Bword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bword
	err := ctx.RegisterResource("fortios:filter/spam/bword:Bword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBword gets an existing Bword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BwordState, opts ...pulumi.ResourceOption) (*Bword, error) {
	var resource Bword
	err := ctx.ReadResource("fortios:filter/spam/bword:Bword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bword resources.
type bwordState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter banned word. The structure of `entries` block is documented below.
	Entries []BwordEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type BwordState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter banned word. The structure of `entries` block is documented below.
	Entries BwordEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (BwordState) ElementType() reflect.Type {
	return reflect.TypeOf((*bwordState)(nil)).Elem()
}

type bwordArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter banned word. The structure of `entries` block is documented below.
	Entries []BwordEntry `pulumi:"entries"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Bword resource.
type BwordArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter banned word. The structure of `entries` block is documented below.
	Entries BwordEntryArrayInput
	// ID.
	Fosid pulumi.IntInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (BwordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bwordArgs)(nil)).Elem()
}

type BwordInput interface {
	pulumi.Input

	ToBwordOutput() BwordOutput
	ToBwordOutputWithContext(ctx context.Context) BwordOutput
}

func (*Bword) ElementType() reflect.Type {
	return reflect.TypeOf((**Bword)(nil)).Elem()
}

func (i *Bword) ToBwordOutput() BwordOutput {
	return i.ToBwordOutputWithContext(context.Background())
}

func (i *Bword) ToBwordOutputWithContext(ctx context.Context) BwordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BwordOutput)
}

// BwordArrayInput is an input type that accepts BwordArray and BwordArrayOutput values.
// You can construct a concrete instance of `BwordArrayInput` via:
//
//	BwordArray{ BwordArgs{...} }
type BwordArrayInput interface {
	pulumi.Input

	ToBwordArrayOutput() BwordArrayOutput
	ToBwordArrayOutputWithContext(context.Context) BwordArrayOutput
}

type BwordArray []BwordInput

func (BwordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bword)(nil)).Elem()
}

func (i BwordArray) ToBwordArrayOutput() BwordArrayOutput {
	return i.ToBwordArrayOutputWithContext(context.Background())
}

func (i BwordArray) ToBwordArrayOutputWithContext(ctx context.Context) BwordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BwordArrayOutput)
}

// BwordMapInput is an input type that accepts BwordMap and BwordMapOutput values.
// You can construct a concrete instance of `BwordMapInput` via:
//
//	BwordMap{ "key": BwordArgs{...} }
type BwordMapInput interface {
	pulumi.Input

	ToBwordMapOutput() BwordMapOutput
	ToBwordMapOutputWithContext(context.Context) BwordMapOutput
}

type BwordMap map[string]BwordInput

func (BwordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bword)(nil)).Elem()
}

func (i BwordMap) ToBwordMapOutput() BwordMapOutput {
	return i.ToBwordMapOutputWithContext(context.Background())
}

func (i BwordMap) ToBwordMapOutputWithContext(ctx context.Context) BwordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BwordMapOutput)
}

type BwordOutput struct{ *pulumi.OutputState }

func (BwordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bword)(nil)).Elem()
}

func (o BwordOutput) ToBwordOutput() BwordOutput {
	return o
}

func (o BwordOutput) ToBwordOutputWithContext(ctx context.Context) BwordOutput {
	return o
}

// Optional comments.
func (o BwordOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bword) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o BwordOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bword) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Spam filter banned word. The structure of `entries` block is documented below.
func (o BwordOutput) Entries() BwordEntryArrayOutput {
	return o.ApplyT(func(v *Bword) BwordEntryArrayOutput { return v.Entries }).(BwordEntryArrayOutput)
}

// ID.
func (o BwordOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Bword) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o BwordOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bword) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Name of table.
func (o BwordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bword) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o BwordOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bword) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type BwordArrayOutput struct{ *pulumi.OutputState }

func (BwordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bword)(nil)).Elem()
}

func (o BwordArrayOutput) ToBwordArrayOutput() BwordArrayOutput {
	return o
}

func (o BwordArrayOutput) ToBwordArrayOutputWithContext(ctx context.Context) BwordArrayOutput {
	return o
}

func (o BwordArrayOutput) Index(i pulumi.IntInput) BwordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bword {
		return vs[0].([]*Bword)[vs[1].(int)]
	}).(BwordOutput)
}

type BwordMapOutput struct{ *pulumi.OutputState }

func (BwordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bword)(nil)).Elem()
}

func (o BwordMapOutput) ToBwordMapOutput() BwordMapOutput {
	return o
}

func (o BwordMapOutput) ToBwordMapOutputWithContext(ctx context.Context) BwordMapOutput {
	return o
}

func (o BwordMapOutput) MapIndex(k pulumi.StringInput) BwordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bword {
		return vs[0].(map[string]*Bword)[vs[1].(string)]
	}).(BwordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BwordInput)(nil)).Elem(), &Bword{})
	pulumi.RegisterInputType(reflect.TypeOf((*BwordArrayInput)(nil)).Elem(), BwordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BwordMapInput)(nil)).Elem(), BwordMap{})
	pulumi.RegisterOutputType(BwordOutput{})
	pulumi.RegisterOutputType(BwordArrayOutput{})
	pulumi.RegisterOutputType(BwordMapOutput{})
}
