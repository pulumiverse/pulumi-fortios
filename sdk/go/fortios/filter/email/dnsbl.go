// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package email

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure AntiSpam DNSBL/ORBL. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// Emailfilter Dnsbl can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:filter/email/dnsbl:Dnsbl labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:filter/email/dnsbl:Dnsbl labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Dnsbl struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries DnsblEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDnsbl registers a new resource with the given unique name, arguments, and options.
func NewDnsbl(ctx *pulumi.Context,
	name string, args *DnsblArgs, opts ...pulumi.ResourceOption) (*Dnsbl, error) {
	if args == nil {
		args = &DnsblArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Dnsbl
	err := ctx.RegisterResource("fortios:filter/email/dnsbl:Dnsbl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsbl gets an existing Dnsbl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsbl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsblState, opts ...pulumi.ResourceOption) (*Dnsbl, error) {
	var resource Dnsbl
	err := ctx.ReadResource("fortios:filter/email/dnsbl:Dnsbl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dnsbl resources.
type dnsblState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries []DnsblEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DnsblState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries DnsblEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DnsblState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsblState)(nil)).Elem()
}

type dnsblArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries []DnsblEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Dnsbl resource.
type DnsblArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries DnsblEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DnsblArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsblArgs)(nil)).Elem()
}

type DnsblInput interface {
	pulumi.Input

	ToDnsblOutput() DnsblOutput
	ToDnsblOutputWithContext(ctx context.Context) DnsblOutput
}

func (*Dnsbl) ElementType() reflect.Type {
	return reflect.TypeOf((**Dnsbl)(nil)).Elem()
}

func (i *Dnsbl) ToDnsblOutput() DnsblOutput {
	return i.ToDnsblOutputWithContext(context.Background())
}

func (i *Dnsbl) ToDnsblOutputWithContext(ctx context.Context) DnsblOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsblOutput)
}

// DnsblArrayInput is an input type that accepts DnsblArray and DnsblArrayOutput values.
// You can construct a concrete instance of `DnsblArrayInput` via:
//
//	DnsblArray{ DnsblArgs{...} }
type DnsblArrayInput interface {
	pulumi.Input

	ToDnsblArrayOutput() DnsblArrayOutput
	ToDnsblArrayOutputWithContext(context.Context) DnsblArrayOutput
}

type DnsblArray []DnsblInput

func (DnsblArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dnsbl)(nil)).Elem()
}

func (i DnsblArray) ToDnsblArrayOutput() DnsblArrayOutput {
	return i.ToDnsblArrayOutputWithContext(context.Background())
}

func (i DnsblArray) ToDnsblArrayOutputWithContext(ctx context.Context) DnsblArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsblArrayOutput)
}

// DnsblMapInput is an input type that accepts DnsblMap and DnsblMapOutput values.
// You can construct a concrete instance of `DnsblMapInput` via:
//
//	DnsblMap{ "key": DnsblArgs{...} }
type DnsblMapInput interface {
	pulumi.Input

	ToDnsblMapOutput() DnsblMapOutput
	ToDnsblMapOutputWithContext(context.Context) DnsblMapOutput
}

type DnsblMap map[string]DnsblInput

func (DnsblMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dnsbl)(nil)).Elem()
}

func (i DnsblMap) ToDnsblMapOutput() DnsblMapOutput {
	return i.ToDnsblMapOutputWithContext(context.Background())
}

func (i DnsblMap) ToDnsblMapOutputWithContext(ctx context.Context) DnsblMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsblMapOutput)
}

type DnsblOutput struct{ *pulumi.OutputState }

func (DnsblOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dnsbl)(nil)).Elem()
}

func (o DnsblOutput) ToDnsblOutput() DnsblOutput {
	return o
}

func (o DnsblOutput) ToDnsblOutputWithContext(ctx context.Context) DnsblOutput {
	return o
}

// Optional comments.
func (o DnsblOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dnsbl) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o DnsblOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dnsbl) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
func (o DnsblOutput) Entries() DnsblEntryArrayOutput {
	return o.ApplyT(func(v *Dnsbl) DnsblEntryArrayOutput { return v.Entries }).(DnsblEntryArrayOutput)
}

// ID.
func (o DnsblOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Dnsbl) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Name of table.
func (o DnsblOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnsbl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DnsblOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dnsbl) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DnsblArrayOutput struct{ *pulumi.OutputState }

func (DnsblArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dnsbl)(nil)).Elem()
}

func (o DnsblArrayOutput) ToDnsblArrayOutput() DnsblArrayOutput {
	return o
}

func (o DnsblArrayOutput) ToDnsblArrayOutputWithContext(ctx context.Context) DnsblArrayOutput {
	return o
}

func (o DnsblArrayOutput) Index(i pulumi.IntInput) DnsblOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dnsbl {
		return vs[0].([]*Dnsbl)[vs[1].(int)]
	}).(DnsblOutput)
}

type DnsblMapOutput struct{ *pulumi.OutputState }

func (DnsblMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dnsbl)(nil)).Elem()
}

func (o DnsblMapOutput) ToDnsblMapOutput() DnsblMapOutput {
	return o
}

func (o DnsblMapOutput) ToDnsblMapOutputWithContext(ctx context.Context) DnsblMapOutput {
	return o
}

func (o DnsblMapOutput) MapIndex(k pulumi.StringInput) DnsblOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dnsbl {
		return vs[0].(map[string]*Dnsbl)[vs[1].(string)]
	}).(DnsblOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsblInput)(nil)).Elem(), &Dnsbl{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsblArrayInput)(nil)).Elem(), DnsblArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsblMapInput)(nil)).Elem(), DnsblMap{})
	pulumi.RegisterOutputType(DnsblOutput{})
	pulumi.RegisterOutputType(DnsblArrayOutput{})
	pulumi.RegisterOutputType(DnsblMapOutput{})
}
