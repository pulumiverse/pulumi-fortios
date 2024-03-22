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

// Configure anti-spam black/white list. Applies to FortiOS Version `<= 6.2.0`.
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
//			_, err := filter.NewBwl(ctx, "trname", &filter.BwlArgs{
//				Comment: pulumi.String("test"),
//				Entries: spam.BwlEntryArray{
//					&spam.BwlEntryArgs{
//						Action:      pulumi.String("reject"),
//						AddrType:    pulumi.String("ipv4"),
//						Ip4Subnet:   pulumi.String("1.1.1.0 255.255.255.0"),
//						Ip6Subnet:   pulumi.String("::/128"),
//						PatternType: pulumi.String("wildcard"),
//						Status:      pulumi.String("enable"),
//						Type:        pulumi.String("ip"),
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
// Spamfilter Bwl can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:filter/spam/bwl:Bwl labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:filter/spam/bwl:Bwl labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Bwl struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries BwlEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewBwl registers a new resource with the given unique name, arguments, and options.
func NewBwl(ctx *pulumi.Context,
	name string, args *BwlArgs, opts ...pulumi.ResourceOption) (*Bwl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bwl
	err := ctx.RegisterResource("fortios:filter/spam/bwl:Bwl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBwl gets an existing Bwl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBwl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BwlState, opts ...pulumi.ResourceOption) (*Bwl, error) {
	var resource Bwl
	err := ctx.ReadResource("fortios:filter/spam/bwl:Bwl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bwl resources.
type bwlState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries []BwlEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type BwlState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries BwlEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (BwlState) ElementType() reflect.Type {
	return reflect.TypeOf((*bwlState)(nil)).Elem()
}

type bwlArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries []BwlEntry `pulumi:"entries"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Bwl resource.
type BwlArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries BwlEntryArrayInput
	// ID.
	Fosid pulumi.IntInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (BwlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bwlArgs)(nil)).Elem()
}

type BwlInput interface {
	pulumi.Input

	ToBwlOutput() BwlOutput
	ToBwlOutputWithContext(ctx context.Context) BwlOutput
}

func (*Bwl) ElementType() reflect.Type {
	return reflect.TypeOf((**Bwl)(nil)).Elem()
}

func (i *Bwl) ToBwlOutput() BwlOutput {
	return i.ToBwlOutputWithContext(context.Background())
}

func (i *Bwl) ToBwlOutputWithContext(ctx context.Context) BwlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BwlOutput)
}

// BwlArrayInput is an input type that accepts BwlArray and BwlArrayOutput values.
// You can construct a concrete instance of `BwlArrayInput` via:
//
//	BwlArray{ BwlArgs{...} }
type BwlArrayInput interface {
	pulumi.Input

	ToBwlArrayOutput() BwlArrayOutput
	ToBwlArrayOutputWithContext(context.Context) BwlArrayOutput
}

type BwlArray []BwlInput

func (BwlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bwl)(nil)).Elem()
}

func (i BwlArray) ToBwlArrayOutput() BwlArrayOutput {
	return i.ToBwlArrayOutputWithContext(context.Background())
}

func (i BwlArray) ToBwlArrayOutputWithContext(ctx context.Context) BwlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BwlArrayOutput)
}

// BwlMapInput is an input type that accepts BwlMap and BwlMapOutput values.
// You can construct a concrete instance of `BwlMapInput` via:
//
//	BwlMap{ "key": BwlArgs{...} }
type BwlMapInput interface {
	pulumi.Input

	ToBwlMapOutput() BwlMapOutput
	ToBwlMapOutputWithContext(context.Context) BwlMapOutput
}

type BwlMap map[string]BwlInput

func (BwlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bwl)(nil)).Elem()
}

func (i BwlMap) ToBwlMapOutput() BwlMapOutput {
	return i.ToBwlMapOutputWithContext(context.Background())
}

func (i BwlMap) ToBwlMapOutputWithContext(ctx context.Context) BwlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BwlMapOutput)
}

type BwlOutput struct{ *pulumi.OutputState }

func (BwlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bwl)(nil)).Elem()
}

func (o BwlOutput) ToBwlOutput() BwlOutput {
	return o
}

func (o BwlOutput) ToBwlOutputWithContext(ctx context.Context) BwlOutput {
	return o
}

// Optional comments.
func (o BwlOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bwl) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o BwlOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bwl) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Anti-spam black/white list entries. The structure of `entries` block is documented below.
func (o BwlOutput) Entries() BwlEntryArrayOutput {
	return o.ApplyT(func(v *Bwl) BwlEntryArrayOutput { return v.Entries }).(BwlEntryArrayOutput)
}

// ID.
func (o BwlOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Bwl) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o BwlOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bwl) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Name of table.
func (o BwlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bwl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o BwlOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bwl) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type BwlArrayOutput struct{ *pulumi.OutputState }

func (BwlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bwl)(nil)).Elem()
}

func (o BwlArrayOutput) ToBwlArrayOutput() BwlArrayOutput {
	return o
}

func (o BwlArrayOutput) ToBwlArrayOutputWithContext(ctx context.Context) BwlArrayOutput {
	return o
}

func (o BwlArrayOutput) Index(i pulumi.IntInput) BwlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bwl {
		return vs[0].([]*Bwl)[vs[1].(int)]
	}).(BwlOutput)
}

type BwlMapOutput struct{ *pulumi.OutputState }

func (BwlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bwl)(nil)).Elem()
}

func (o BwlMapOutput) ToBwlMapOutput() BwlMapOutput {
	return o
}

func (o BwlMapOutput) ToBwlMapOutputWithContext(ctx context.Context) BwlMapOutput {
	return o
}

func (o BwlMapOutput) MapIndex(k pulumi.StringInput) BwlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bwl {
		return vs[0].(map[string]*Bwl)[vs[1].(string)]
	}).(BwlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BwlInput)(nil)).Elem(), &Bwl{})
	pulumi.RegisterInputType(reflect.TypeOf((*BwlArrayInput)(nil)).Elem(), BwlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BwlMapInput)(nil)).Elem(), BwlMap{})
	pulumi.RegisterOutputType(BwlOutput{})
	pulumi.RegisterOutputType(BwlArrayOutput{})
	pulumi.RegisterOutputType(BwlMapOutput{})
}
