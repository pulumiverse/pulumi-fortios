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

// Configure URL filter lists.
//
// ## Example Usage
//
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
//			_, err := filter.NewUrlfilter(ctx, "trname", &filter.UrlfilterArgs{
//				Fosid:              pulumi.Int(1),
//				IpAddrBlock:        pulumi.String("enable"),
//				OneArmIpsUrlfilter: pulumi.String("enable"),
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
// Webfilter Urlfilter can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:filter/web/urlfilter:Urlfilter labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:filter/web/urlfilter:Urlfilter labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Urlfilter struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// URL filter entries. The structure of `entries` block is documented below.
	Entries UrlfilterEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Enable/disable matching of IPv4 mapped IPv6 URLs. Valid values: `enable`, `disable`.
	Ip4MappedIp6 pulumi.StringOutput `pulumi:"ip4MappedIp6"`
	// Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable`, `disable`.
	IpAddrBlock pulumi.StringOutput `pulumi:"ipAddrBlock"`
	// Name of URL filter list.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable`, `disable`.
	OneArmIpsUrlfilter pulumi.StringOutput `pulumi:"oneArmIpsUrlfilter"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUrlfilter registers a new resource with the given unique name, arguments, and options.
func NewUrlfilter(ctx *pulumi.Context,
	name string, args *UrlfilterArgs, opts ...pulumi.ResourceOption) (*Urlfilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Urlfilter
	err := ctx.RegisterResource("fortios:filter/web/urlfilter:Urlfilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUrlfilter gets an existing Urlfilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUrlfilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UrlfilterState, opts ...pulumi.ResourceOption) (*Urlfilter, error) {
	var resource Urlfilter
	err := ctx.ReadResource("fortios:filter/web/urlfilter:Urlfilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Urlfilter resources.
type urlfilterState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// URL filter entries. The structure of `entries` block is documented below.
	Entries []UrlfilterEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable matching of IPv4 mapped IPv6 URLs. Valid values: `enable`, `disable`.
	Ip4MappedIp6 *string `pulumi:"ip4MappedIp6"`
	// Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable`, `disable`.
	IpAddrBlock *string `pulumi:"ipAddrBlock"`
	// Name of URL filter list.
	Name *string `pulumi:"name"`
	// Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable`, `disable`.
	OneArmIpsUrlfilter *string `pulumi:"oneArmIpsUrlfilter"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UrlfilterState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// URL filter entries. The structure of `entries` block is documented below.
	Entries UrlfilterEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable matching of IPv4 mapped IPv6 URLs. Valid values: `enable`, `disable`.
	Ip4MappedIp6 pulumi.StringPtrInput
	// Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable`, `disable`.
	IpAddrBlock pulumi.StringPtrInput
	// Name of URL filter list.
	Name pulumi.StringPtrInput
	// Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable`, `disable`.
	OneArmIpsUrlfilter pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UrlfilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*urlfilterState)(nil)).Elem()
}

type urlfilterArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// URL filter entries. The structure of `entries` block is documented below.
	Entries []UrlfilterEntry `pulumi:"entries"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Enable/disable matching of IPv4 mapped IPv6 URLs. Valid values: `enable`, `disable`.
	Ip4MappedIp6 *string `pulumi:"ip4MappedIp6"`
	// Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable`, `disable`.
	IpAddrBlock *string `pulumi:"ipAddrBlock"`
	// Name of URL filter list.
	Name *string `pulumi:"name"`
	// Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable`, `disable`.
	OneArmIpsUrlfilter *string `pulumi:"oneArmIpsUrlfilter"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Urlfilter resource.
type UrlfilterArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// URL filter entries. The structure of `entries` block is documented below.
	Entries UrlfilterEntryArrayInput
	// ID.
	Fosid pulumi.IntInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Enable/disable matching of IPv4 mapped IPv6 URLs. Valid values: `enable`, `disable`.
	Ip4MappedIp6 pulumi.StringPtrInput
	// Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable`, `disable`.
	IpAddrBlock pulumi.StringPtrInput
	// Name of URL filter list.
	Name pulumi.StringPtrInput
	// Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable`, `disable`.
	OneArmIpsUrlfilter pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UrlfilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*urlfilterArgs)(nil)).Elem()
}

type UrlfilterInput interface {
	pulumi.Input

	ToUrlfilterOutput() UrlfilterOutput
	ToUrlfilterOutputWithContext(ctx context.Context) UrlfilterOutput
}

func (*Urlfilter) ElementType() reflect.Type {
	return reflect.TypeOf((**Urlfilter)(nil)).Elem()
}

func (i *Urlfilter) ToUrlfilterOutput() UrlfilterOutput {
	return i.ToUrlfilterOutputWithContext(context.Background())
}

func (i *Urlfilter) ToUrlfilterOutputWithContext(ctx context.Context) UrlfilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlfilterOutput)
}

// UrlfilterArrayInput is an input type that accepts UrlfilterArray and UrlfilterArrayOutput values.
// You can construct a concrete instance of `UrlfilterArrayInput` via:
//
//	UrlfilterArray{ UrlfilterArgs{...} }
type UrlfilterArrayInput interface {
	pulumi.Input

	ToUrlfilterArrayOutput() UrlfilterArrayOutput
	ToUrlfilterArrayOutputWithContext(context.Context) UrlfilterArrayOutput
}

type UrlfilterArray []UrlfilterInput

func (UrlfilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Urlfilter)(nil)).Elem()
}

func (i UrlfilterArray) ToUrlfilterArrayOutput() UrlfilterArrayOutput {
	return i.ToUrlfilterArrayOutputWithContext(context.Background())
}

func (i UrlfilterArray) ToUrlfilterArrayOutputWithContext(ctx context.Context) UrlfilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlfilterArrayOutput)
}

// UrlfilterMapInput is an input type that accepts UrlfilterMap and UrlfilterMapOutput values.
// You can construct a concrete instance of `UrlfilterMapInput` via:
//
//	UrlfilterMap{ "key": UrlfilterArgs{...} }
type UrlfilterMapInput interface {
	pulumi.Input

	ToUrlfilterMapOutput() UrlfilterMapOutput
	ToUrlfilterMapOutputWithContext(context.Context) UrlfilterMapOutput
}

type UrlfilterMap map[string]UrlfilterInput

func (UrlfilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Urlfilter)(nil)).Elem()
}

func (i UrlfilterMap) ToUrlfilterMapOutput() UrlfilterMapOutput {
	return i.ToUrlfilterMapOutputWithContext(context.Background())
}

func (i UrlfilterMap) ToUrlfilterMapOutputWithContext(ctx context.Context) UrlfilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlfilterMapOutput)
}

type UrlfilterOutput struct{ *pulumi.OutputState }

func (UrlfilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Urlfilter)(nil)).Elem()
}

func (o UrlfilterOutput) ToUrlfilterOutput() UrlfilterOutput {
	return o
}

func (o UrlfilterOutput) ToUrlfilterOutputWithContext(ctx context.Context) UrlfilterOutput {
	return o
}

// Optional comments.
func (o UrlfilterOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Urlfilter) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o UrlfilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Urlfilter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// URL filter entries. The structure of `entries` block is documented below.
func (o UrlfilterOutput) Entries() UrlfilterEntryArrayOutput {
	return o.ApplyT(func(v *Urlfilter) UrlfilterEntryArrayOutput { return v.Entries }).(UrlfilterEntryArrayOutput)
}

// ID.
func (o UrlfilterOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Urlfilter) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o UrlfilterOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Urlfilter) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Enable/disable matching of IPv4 mapped IPv6 URLs. Valid values: `enable`, `disable`.
func (o UrlfilterOutput) Ip4MappedIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *Urlfilter) pulumi.StringOutput { return v.Ip4MappedIp6 }).(pulumi.StringOutput)
}

// Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable`, `disable`.
func (o UrlfilterOutput) IpAddrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *Urlfilter) pulumi.StringOutput { return v.IpAddrBlock }).(pulumi.StringOutput)
}

// Name of URL filter list.
func (o UrlfilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Urlfilter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable`, `disable`.
func (o UrlfilterOutput) OneArmIpsUrlfilter() pulumi.StringOutput {
	return o.ApplyT(func(v *Urlfilter) pulumi.StringOutput { return v.OneArmIpsUrlfilter }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o UrlfilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Urlfilter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type UrlfilterArrayOutput struct{ *pulumi.OutputState }

func (UrlfilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Urlfilter)(nil)).Elem()
}

func (o UrlfilterArrayOutput) ToUrlfilterArrayOutput() UrlfilterArrayOutput {
	return o
}

func (o UrlfilterArrayOutput) ToUrlfilterArrayOutputWithContext(ctx context.Context) UrlfilterArrayOutput {
	return o
}

func (o UrlfilterArrayOutput) Index(i pulumi.IntInput) UrlfilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Urlfilter {
		return vs[0].([]*Urlfilter)[vs[1].(int)]
	}).(UrlfilterOutput)
}

type UrlfilterMapOutput struct{ *pulumi.OutputState }

func (UrlfilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Urlfilter)(nil)).Elem()
}

func (o UrlfilterMapOutput) ToUrlfilterMapOutput() UrlfilterMapOutput {
	return o
}

func (o UrlfilterMapOutput) ToUrlfilterMapOutputWithContext(ctx context.Context) UrlfilterMapOutput {
	return o
}

func (o UrlfilterMapOutput) MapIndex(k pulumi.StringInput) UrlfilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Urlfilter {
		return vs[0].(map[string]*Urlfilter)[vs[1].(string)]
	}).(UrlfilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UrlfilterInput)(nil)).Elem(), &Urlfilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*UrlfilterArrayInput)(nil)).Elem(), UrlfilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UrlfilterMapInput)(nil)).Elem(), UrlfilterMap{})
	pulumi.RegisterOutputType(UrlfilterOutput{})
	pulumi.RegisterOutputType(UrlfilterArrayOutput{})
	pulumi.RegisterOutputType(UrlfilterMapOutput{})
}
