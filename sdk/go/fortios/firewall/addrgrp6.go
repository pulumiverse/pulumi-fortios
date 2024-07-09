// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPv6 address groups.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/firewall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			trname1, err := firewall.NewAddress6(ctx, "trname1", &firewall.Address6Args{
//				CacheTtl:   pulumi.Int(0),
//				Color:      pulumi.Int(0),
//				EndIp:      pulumi.String("::"),
//				Host:       pulumi.String(""),
//				HostType:   pulumi.String("any"),
//				Ip6:        pulumi.String("fdff:ffff::/120"),
//				StartIp:    pulumi.String(""),
//				Type:       pulumi.String("ipprefix"),
//				Visibility: pulumi.String("enable"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = firewall.NewAddrgrp6(ctx, "trname", &firewall.Addrgrp6Args{
//				Color:      pulumi.Int(0),
//				Visibility: pulumi.String("enable"),
//				Members: firewall.Addrgrp6MemberArray{
//					&firewall.Addrgrp6MemberArgs{
//						Name: trname1.Name,
//					},
//				},
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
// Firewall Addrgrp6 can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/addrgrp6:Addrgrp6 labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/addrgrp6:Addrgrp6 labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Addrgrp6 struct {
	pulumi.CustomResourceState

	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable address6 exclusion. Valid values: `enable`, `disable`.
	Exclude pulumi.StringOutput `pulumi:"exclude"`
	// Address6 exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers Addrgrp6ExcludeMemberArrayOutput `pulumi:"excludeMembers"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringOutput `pulumi:"fabricObject"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members Addrgrp6MemberArrayOutput `pulumi:"members"`
	// IPv6 address group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings Addrgrp6TaggingArrayOutput `pulumi:"taggings"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
	// Enable/disable address group6 visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewAddrgrp6 registers a new resource with the given unique name, arguments, and options.
func NewAddrgrp6(ctx *pulumi.Context,
	name string, args *Addrgrp6Args, opts ...pulumi.ResourceOption) (*Addrgrp6, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Addrgrp6
	err := ctx.RegisterResource("fortios:firewall/addrgrp6:Addrgrp6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddrgrp6 gets an existing Addrgrp6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddrgrp6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Addrgrp6State, opts ...pulumi.ResourceOption) (*Addrgrp6, error) {
	var resource Addrgrp6
	err := ctx.ReadResource("fortios:firewall/addrgrp6:Addrgrp6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Addrgrp6 resources.
type addrgrp6State struct {
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable address6 exclusion. Valid values: `enable`, `disable`.
	Exclude *string `pulumi:"exclude"`
	// Address6 exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers []Addrgrp6ExcludeMember `pulumi:"excludeMembers"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members []Addrgrp6Member `pulumi:"members"`
	// IPv6 address group name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []Addrgrp6Tagging `pulumi:"taggings"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address group6 visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

type Addrgrp6State struct {
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable address6 exclusion. Valid values: `enable`, `disable`.
	Exclude pulumi.StringPtrInput
	// Address6 exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers Addrgrp6ExcludeMemberArrayInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members Addrgrp6MemberArrayInput
	// IPv6 address group name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings Addrgrp6TaggingArrayInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable address group6 visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (Addrgrp6State) ElementType() reflect.Type {
	return reflect.TypeOf((*addrgrp6State)(nil)).Elem()
}

type addrgrp6Args struct {
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable address6 exclusion. Valid values: `enable`, `disable`.
	Exclude *string `pulumi:"exclude"`
	// Address6 exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers []Addrgrp6ExcludeMember `pulumi:"excludeMembers"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members []Addrgrp6Member `pulumi:"members"`
	// IPv6 address group name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []Addrgrp6Tagging `pulumi:"taggings"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address group6 visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a Addrgrp6 resource.
type Addrgrp6Args struct {
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable address6 exclusion. Valid values: `enable`, `disable`.
	Exclude pulumi.StringPtrInput
	// Address6 exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers Addrgrp6ExcludeMemberArrayInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members Addrgrp6MemberArrayInput
	// IPv6 address group name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings Addrgrp6TaggingArrayInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable address group6 visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (Addrgrp6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*addrgrp6Args)(nil)).Elem()
}

type Addrgrp6Input interface {
	pulumi.Input

	ToAddrgrp6Output() Addrgrp6Output
	ToAddrgrp6OutputWithContext(ctx context.Context) Addrgrp6Output
}

func (*Addrgrp6) ElementType() reflect.Type {
	return reflect.TypeOf((**Addrgrp6)(nil)).Elem()
}

func (i *Addrgrp6) ToAddrgrp6Output() Addrgrp6Output {
	return i.ToAddrgrp6OutputWithContext(context.Background())
}

func (i *Addrgrp6) ToAddrgrp6OutputWithContext(ctx context.Context) Addrgrp6Output {
	return pulumi.ToOutputWithContext(ctx, i).(Addrgrp6Output)
}

// Addrgrp6ArrayInput is an input type that accepts Addrgrp6Array and Addrgrp6ArrayOutput values.
// You can construct a concrete instance of `Addrgrp6ArrayInput` via:
//
//	Addrgrp6Array{ Addrgrp6Args{...} }
type Addrgrp6ArrayInput interface {
	pulumi.Input

	ToAddrgrp6ArrayOutput() Addrgrp6ArrayOutput
	ToAddrgrp6ArrayOutputWithContext(context.Context) Addrgrp6ArrayOutput
}

type Addrgrp6Array []Addrgrp6Input

func (Addrgrp6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Addrgrp6)(nil)).Elem()
}

func (i Addrgrp6Array) ToAddrgrp6ArrayOutput() Addrgrp6ArrayOutput {
	return i.ToAddrgrp6ArrayOutputWithContext(context.Background())
}

func (i Addrgrp6Array) ToAddrgrp6ArrayOutputWithContext(ctx context.Context) Addrgrp6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Addrgrp6ArrayOutput)
}

// Addrgrp6MapInput is an input type that accepts Addrgrp6Map and Addrgrp6MapOutput values.
// You can construct a concrete instance of `Addrgrp6MapInput` via:
//
//	Addrgrp6Map{ "key": Addrgrp6Args{...} }
type Addrgrp6MapInput interface {
	pulumi.Input

	ToAddrgrp6MapOutput() Addrgrp6MapOutput
	ToAddrgrp6MapOutputWithContext(context.Context) Addrgrp6MapOutput
}

type Addrgrp6Map map[string]Addrgrp6Input

func (Addrgrp6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Addrgrp6)(nil)).Elem()
}

func (i Addrgrp6Map) ToAddrgrp6MapOutput() Addrgrp6MapOutput {
	return i.ToAddrgrp6MapOutputWithContext(context.Background())
}

func (i Addrgrp6Map) ToAddrgrp6MapOutputWithContext(ctx context.Context) Addrgrp6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Addrgrp6MapOutput)
}

type Addrgrp6Output struct{ *pulumi.OutputState }

func (Addrgrp6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Addrgrp6)(nil)).Elem()
}

func (o Addrgrp6Output) ToAddrgrp6Output() Addrgrp6Output {
	return o
}

func (o Addrgrp6Output) ToAddrgrp6OutputWithContext(ctx context.Context) Addrgrp6Output {
	return o
}

// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
func (o Addrgrp6Output) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *Addrgrp6) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o Addrgrp6Output) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Addrgrp6) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Addrgrp6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Addrgrp6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable address6 exclusion. Valid values: `enable`, `disable`.
func (o Addrgrp6Output) Exclude() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp6) pulumi.StringOutput { return v.Exclude }).(pulumi.StringOutput)
}

// Address6 exclusion member. The structure of `excludeMember` block is documented below.
func (o Addrgrp6Output) ExcludeMembers() Addrgrp6ExcludeMemberArrayOutput {
	return o.ApplyT(func(v *Addrgrp6) Addrgrp6ExcludeMemberArrayOutput { return v.ExcludeMembers }).(Addrgrp6ExcludeMemberArrayOutput)
}

// Security Fabric global object setting. Valid values: `enable`, `disable`.
func (o Addrgrp6Output) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp6) pulumi.StringOutput { return v.FabricObject }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o Addrgrp6Output) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Addrgrp6) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Address objects contained within the group. The structure of `member` block is documented below.
func (o Addrgrp6Output) Members() Addrgrp6MemberArrayOutput {
	return o.ApplyT(func(v *Addrgrp6) Addrgrp6MemberArrayOutput { return v.Members }).(Addrgrp6MemberArrayOutput)
}

// IPv6 address group name.
func (o Addrgrp6Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp6) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o Addrgrp6Output) Taggings() Addrgrp6TaggingArrayOutput {
	return o.ApplyT(func(v *Addrgrp6) Addrgrp6TaggingArrayOutput { return v.Taggings }).(Addrgrp6TaggingArrayOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o Addrgrp6Output) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp6) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Addrgrp6Output) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp6) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

// Enable/disable address group6 visibility in the GUI. Valid values: `enable`, `disable`.
func (o Addrgrp6Output) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp6) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type Addrgrp6ArrayOutput struct{ *pulumi.OutputState }

func (Addrgrp6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Addrgrp6)(nil)).Elem()
}

func (o Addrgrp6ArrayOutput) ToAddrgrp6ArrayOutput() Addrgrp6ArrayOutput {
	return o
}

func (o Addrgrp6ArrayOutput) ToAddrgrp6ArrayOutputWithContext(ctx context.Context) Addrgrp6ArrayOutput {
	return o
}

func (o Addrgrp6ArrayOutput) Index(i pulumi.IntInput) Addrgrp6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Addrgrp6 {
		return vs[0].([]*Addrgrp6)[vs[1].(int)]
	}).(Addrgrp6Output)
}

type Addrgrp6MapOutput struct{ *pulumi.OutputState }

func (Addrgrp6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Addrgrp6)(nil)).Elem()
}

func (o Addrgrp6MapOutput) ToAddrgrp6MapOutput() Addrgrp6MapOutput {
	return o
}

func (o Addrgrp6MapOutput) ToAddrgrp6MapOutputWithContext(ctx context.Context) Addrgrp6MapOutput {
	return o
}

func (o Addrgrp6MapOutput) MapIndex(k pulumi.StringInput) Addrgrp6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Addrgrp6 {
		return vs[0].(map[string]*Addrgrp6)[vs[1].(string)]
	}).(Addrgrp6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Addrgrp6Input)(nil)).Elem(), &Addrgrp6{})
	pulumi.RegisterInputType(reflect.TypeOf((*Addrgrp6ArrayInput)(nil)).Elem(), Addrgrp6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Addrgrp6MapInput)(nil)).Elem(), Addrgrp6Map{})
	pulumi.RegisterOutputType(Addrgrp6Output{})
	pulumi.RegisterOutputType(Addrgrp6ArrayOutput{})
	pulumi.RegisterOutputType(Addrgrp6MapOutput{})
}
