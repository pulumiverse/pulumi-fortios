// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure AntiSpam DNSBL/ORBL. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewSpamfilterDnsbl(ctx, "trname", &fortios.SpamfilterDnsblArgs{
//				Comment: pulumi.String("test"),
//				Entries: fortios.SpamfilterDnsblEntryArray{
//					&fortios.SpamfilterDnsblEntryArgs{
//						Action: pulumi.String("reject"),
//						Server: pulumi.String("1.1.1.1"),
//						Status: pulumi.String("enable"),
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
//
// ## Import
//
// # Spamfilter Dnsbl can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/spamfilterDnsbl:SpamfilterDnsbl labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/spamfilterDnsbl:SpamfilterDnsbl labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SpamfilterDnsbl struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries SpamfilterDnsblEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSpamfilterDnsbl registers a new resource with the given unique name, arguments, and options.
func NewSpamfilterDnsbl(ctx *pulumi.Context,
	name string, args *SpamfilterDnsblArgs, opts ...pulumi.ResourceOption) (*SpamfilterDnsbl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SpamfilterDnsbl
	err := ctx.RegisterResource("fortios:index/spamfilterDnsbl:SpamfilterDnsbl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpamfilterDnsbl gets an existing SpamfilterDnsbl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpamfilterDnsbl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpamfilterDnsblState, opts ...pulumi.ResourceOption) (*SpamfilterDnsbl, error) {
	var resource SpamfilterDnsbl
	err := ctx.ReadResource("fortios:index/spamfilterDnsbl:SpamfilterDnsbl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpamfilterDnsbl resources.
type spamfilterDnsblState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries []SpamfilterDnsblEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SpamfilterDnsblState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries SpamfilterDnsblEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpamfilterDnsblState) ElementType() reflect.Type {
	return reflect.TypeOf((*spamfilterDnsblState)(nil)).Elem()
}

type spamfilterDnsblArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries []SpamfilterDnsblEntry `pulumi:"entries"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SpamfilterDnsbl resource.
type SpamfilterDnsblArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries SpamfilterDnsblEntryArrayInput
	// ID.
	Fosid pulumi.IntInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpamfilterDnsblArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spamfilterDnsblArgs)(nil)).Elem()
}

type SpamfilterDnsblInput interface {
	pulumi.Input

	ToSpamfilterDnsblOutput() SpamfilterDnsblOutput
	ToSpamfilterDnsblOutputWithContext(ctx context.Context) SpamfilterDnsblOutput
}

func (*SpamfilterDnsbl) ElementType() reflect.Type {
	return reflect.TypeOf((**SpamfilterDnsbl)(nil)).Elem()
}

func (i *SpamfilterDnsbl) ToSpamfilterDnsblOutput() SpamfilterDnsblOutput {
	return i.ToSpamfilterDnsblOutputWithContext(context.Background())
}

func (i *SpamfilterDnsbl) ToSpamfilterDnsblOutputWithContext(ctx context.Context) SpamfilterDnsblOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterDnsblOutput)
}

// SpamfilterDnsblArrayInput is an input type that accepts SpamfilterDnsblArray and SpamfilterDnsblArrayOutput values.
// You can construct a concrete instance of `SpamfilterDnsblArrayInput` via:
//
//	SpamfilterDnsblArray{ SpamfilterDnsblArgs{...} }
type SpamfilterDnsblArrayInput interface {
	pulumi.Input

	ToSpamfilterDnsblArrayOutput() SpamfilterDnsblArrayOutput
	ToSpamfilterDnsblArrayOutputWithContext(context.Context) SpamfilterDnsblArrayOutput
}

type SpamfilterDnsblArray []SpamfilterDnsblInput

func (SpamfilterDnsblArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpamfilterDnsbl)(nil)).Elem()
}

func (i SpamfilterDnsblArray) ToSpamfilterDnsblArrayOutput() SpamfilterDnsblArrayOutput {
	return i.ToSpamfilterDnsblArrayOutputWithContext(context.Background())
}

func (i SpamfilterDnsblArray) ToSpamfilterDnsblArrayOutputWithContext(ctx context.Context) SpamfilterDnsblArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterDnsblArrayOutput)
}

// SpamfilterDnsblMapInput is an input type that accepts SpamfilterDnsblMap and SpamfilterDnsblMapOutput values.
// You can construct a concrete instance of `SpamfilterDnsblMapInput` via:
//
//	SpamfilterDnsblMap{ "key": SpamfilterDnsblArgs{...} }
type SpamfilterDnsblMapInput interface {
	pulumi.Input

	ToSpamfilterDnsblMapOutput() SpamfilterDnsblMapOutput
	ToSpamfilterDnsblMapOutputWithContext(context.Context) SpamfilterDnsblMapOutput
}

type SpamfilterDnsblMap map[string]SpamfilterDnsblInput

func (SpamfilterDnsblMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpamfilterDnsbl)(nil)).Elem()
}

func (i SpamfilterDnsblMap) ToSpamfilterDnsblMapOutput() SpamfilterDnsblMapOutput {
	return i.ToSpamfilterDnsblMapOutputWithContext(context.Background())
}

func (i SpamfilterDnsblMap) ToSpamfilterDnsblMapOutputWithContext(ctx context.Context) SpamfilterDnsblMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterDnsblMapOutput)
}

type SpamfilterDnsblOutput struct{ *pulumi.OutputState }

func (SpamfilterDnsblOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpamfilterDnsbl)(nil)).Elem()
}

func (o SpamfilterDnsblOutput) ToSpamfilterDnsblOutput() SpamfilterDnsblOutput {
	return o
}

func (o SpamfilterDnsblOutput) ToSpamfilterDnsblOutputWithContext(ctx context.Context) SpamfilterDnsblOutput {
	return o
}

// Optional comments.
func (o SpamfilterDnsblOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SpamfilterDnsbl) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SpamfilterDnsblOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SpamfilterDnsbl) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
func (o SpamfilterDnsblOutput) Entries() SpamfilterDnsblEntryArrayOutput {
	return o.ApplyT(func(v *SpamfilterDnsbl) SpamfilterDnsblEntryArrayOutput { return v.Entries }).(SpamfilterDnsblEntryArrayOutput)
}

// ID.
func (o SpamfilterDnsblOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *SpamfilterDnsbl) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Name of table.
func (o SpamfilterDnsblOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SpamfilterDnsbl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SpamfilterDnsblOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SpamfilterDnsbl) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SpamfilterDnsblArrayOutput struct{ *pulumi.OutputState }

func (SpamfilterDnsblArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpamfilterDnsbl)(nil)).Elem()
}

func (o SpamfilterDnsblArrayOutput) ToSpamfilterDnsblArrayOutput() SpamfilterDnsblArrayOutput {
	return o
}

func (o SpamfilterDnsblArrayOutput) ToSpamfilterDnsblArrayOutputWithContext(ctx context.Context) SpamfilterDnsblArrayOutput {
	return o
}

func (o SpamfilterDnsblArrayOutput) Index(i pulumi.IntInput) SpamfilterDnsblOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SpamfilterDnsbl {
		return vs[0].([]*SpamfilterDnsbl)[vs[1].(int)]
	}).(SpamfilterDnsblOutput)
}

type SpamfilterDnsblMapOutput struct{ *pulumi.OutputState }

func (SpamfilterDnsblMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpamfilterDnsbl)(nil)).Elem()
}

func (o SpamfilterDnsblMapOutput) ToSpamfilterDnsblMapOutput() SpamfilterDnsblMapOutput {
	return o
}

func (o SpamfilterDnsblMapOutput) ToSpamfilterDnsblMapOutputWithContext(ctx context.Context) SpamfilterDnsblMapOutput {
	return o
}

func (o SpamfilterDnsblMapOutput) MapIndex(k pulumi.StringInput) SpamfilterDnsblOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SpamfilterDnsbl {
		return vs[0].(map[string]*SpamfilterDnsbl)[vs[1].(string)]
	}).(SpamfilterDnsblOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterDnsblInput)(nil)).Elem(), &SpamfilterDnsbl{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterDnsblArrayInput)(nil)).Elem(), SpamfilterDnsblArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterDnsblMapInput)(nil)).Elem(), SpamfilterDnsblMap{})
	pulumi.RegisterOutputType(SpamfilterDnsblOutput{})
	pulumi.RegisterOutputType(SpamfilterDnsblArrayOutput{})
	pulumi.RegisterOutputType(SpamfilterDnsblMapOutput{})
}
