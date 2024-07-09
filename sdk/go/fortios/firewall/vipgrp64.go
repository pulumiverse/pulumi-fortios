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

// Configure IPv6 to IPv4 virtual IP groups. Applies to FortiOS Version `<= 7.0.0`.
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
//			trname1, err := firewall.NewVip64(ctx, "trname1", &firewall.Vip64Args{
//				ArpReply:    pulumi.String("enable"),
//				Color:       pulumi.Int(0),
//				Extip:       pulumi.String("2001:db8:99:503::22"),
//				Extport:     pulumi.String("0-65535"),
//				Fosid:       pulumi.Int(0),
//				LdbMethod:   pulumi.String("static"),
//				Mappedip:    pulumi.String("1.1.3.1"),
//				Mappedport:  pulumi.String("0-65535"),
//				Portforward: pulumi.String("disable"),
//				Protocol:    pulumi.String("tcp"),
//				Type:        pulumi.String("static-nat"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = firewall.NewVipgrp64(ctx, "trname", &firewall.Vipgrp64Args{
//				Color: pulumi.Int(0),
//				Members: firewall.Vipgrp64MemberArray{
//					&firewall.Vipgrp64MemberArgs{
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
// Firewall Vipgrp64 can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/vipgrp64:Vipgrp64 labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/vipgrp64:Vipgrp64 labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Vipgrp64 struct {
	pulumi.CustomResourceState

	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members Vipgrp64MemberArrayOutput `pulumi:"members"`
	// VIP64 group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewVipgrp64 registers a new resource with the given unique name, arguments, and options.
func NewVipgrp64(ctx *pulumi.Context,
	name string, args *Vipgrp64Args, opts ...pulumi.ResourceOption) (*Vipgrp64, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vipgrp64
	err := ctx.RegisterResource("fortios:firewall/vipgrp64:Vipgrp64", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVipgrp64 gets an existing Vipgrp64 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVipgrp64(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Vipgrp64State, opts ...pulumi.ResourceOption) (*Vipgrp64, error) {
	var resource Vipgrp64
	err := ctx.ReadResource("fortios:firewall/vipgrp64:Vipgrp64", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vipgrp64 resources.
type vipgrp64State struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color *int `pulumi:"color"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members []Vipgrp64Member `pulumi:"members"`
	// VIP64 group name.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Vipgrp64State struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members Vipgrp64MemberArrayInput
	// VIP64 group name.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Vipgrp64State) ElementType() reflect.Type {
	return reflect.TypeOf((*vipgrp64State)(nil)).Elem()
}

type vipgrp64Args struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color *int `pulumi:"color"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members []Vipgrp64Member `pulumi:"members"`
	// VIP64 group name.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Vipgrp64 resource.
type Vipgrp64Args struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members Vipgrp64MemberArrayInput
	// VIP64 group name.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Vipgrp64Args) ElementType() reflect.Type {
	return reflect.TypeOf((*vipgrp64Args)(nil)).Elem()
}

type Vipgrp64Input interface {
	pulumi.Input

	ToVipgrp64Output() Vipgrp64Output
	ToVipgrp64OutputWithContext(ctx context.Context) Vipgrp64Output
}

func (*Vipgrp64) ElementType() reflect.Type {
	return reflect.TypeOf((**Vipgrp64)(nil)).Elem()
}

func (i *Vipgrp64) ToVipgrp64Output() Vipgrp64Output {
	return i.ToVipgrp64OutputWithContext(context.Background())
}

func (i *Vipgrp64) ToVipgrp64OutputWithContext(ctx context.Context) Vipgrp64Output {
	return pulumi.ToOutputWithContext(ctx, i).(Vipgrp64Output)
}

// Vipgrp64ArrayInput is an input type that accepts Vipgrp64Array and Vipgrp64ArrayOutput values.
// You can construct a concrete instance of `Vipgrp64ArrayInput` via:
//
//	Vipgrp64Array{ Vipgrp64Args{...} }
type Vipgrp64ArrayInput interface {
	pulumi.Input

	ToVipgrp64ArrayOutput() Vipgrp64ArrayOutput
	ToVipgrp64ArrayOutputWithContext(context.Context) Vipgrp64ArrayOutput
}

type Vipgrp64Array []Vipgrp64Input

func (Vipgrp64Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vipgrp64)(nil)).Elem()
}

func (i Vipgrp64Array) ToVipgrp64ArrayOutput() Vipgrp64ArrayOutput {
	return i.ToVipgrp64ArrayOutputWithContext(context.Background())
}

func (i Vipgrp64Array) ToVipgrp64ArrayOutputWithContext(ctx context.Context) Vipgrp64ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Vipgrp64ArrayOutput)
}

// Vipgrp64MapInput is an input type that accepts Vipgrp64Map and Vipgrp64MapOutput values.
// You can construct a concrete instance of `Vipgrp64MapInput` via:
//
//	Vipgrp64Map{ "key": Vipgrp64Args{...} }
type Vipgrp64MapInput interface {
	pulumi.Input

	ToVipgrp64MapOutput() Vipgrp64MapOutput
	ToVipgrp64MapOutputWithContext(context.Context) Vipgrp64MapOutput
}

type Vipgrp64Map map[string]Vipgrp64Input

func (Vipgrp64Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vipgrp64)(nil)).Elem()
}

func (i Vipgrp64Map) ToVipgrp64MapOutput() Vipgrp64MapOutput {
	return i.ToVipgrp64MapOutputWithContext(context.Background())
}

func (i Vipgrp64Map) ToVipgrp64MapOutputWithContext(ctx context.Context) Vipgrp64MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Vipgrp64MapOutput)
}

type Vipgrp64Output struct{ *pulumi.OutputState }

func (Vipgrp64Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Vipgrp64)(nil)).Elem()
}

func (o Vipgrp64Output) ToVipgrp64Output() Vipgrp64Output {
	return o
}

func (o Vipgrp64Output) ToVipgrp64OutputWithContext(ctx context.Context) Vipgrp64Output {
	return o
}

// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
func (o Vipgrp64Output) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *Vipgrp64) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o Vipgrp64Output) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vipgrp64) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Vipgrp64Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vipgrp64) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o Vipgrp64Output) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vipgrp64) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
func (o Vipgrp64Output) Members() Vipgrp64MemberArrayOutput {
	return o.ApplyT(func(v *Vipgrp64) Vipgrp64MemberArrayOutput { return v.Members }).(Vipgrp64MemberArrayOutput)
}

// VIP64 group name.
func (o Vipgrp64Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vipgrp64) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o Vipgrp64Output) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Vipgrp64) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Vipgrp64Output) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Vipgrp64) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type Vipgrp64ArrayOutput struct{ *pulumi.OutputState }

func (Vipgrp64ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vipgrp64)(nil)).Elem()
}

func (o Vipgrp64ArrayOutput) ToVipgrp64ArrayOutput() Vipgrp64ArrayOutput {
	return o
}

func (o Vipgrp64ArrayOutput) ToVipgrp64ArrayOutputWithContext(ctx context.Context) Vipgrp64ArrayOutput {
	return o
}

func (o Vipgrp64ArrayOutput) Index(i pulumi.IntInput) Vipgrp64Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vipgrp64 {
		return vs[0].([]*Vipgrp64)[vs[1].(int)]
	}).(Vipgrp64Output)
}

type Vipgrp64MapOutput struct{ *pulumi.OutputState }

func (Vipgrp64MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vipgrp64)(nil)).Elem()
}

func (o Vipgrp64MapOutput) ToVipgrp64MapOutput() Vipgrp64MapOutput {
	return o
}

func (o Vipgrp64MapOutput) ToVipgrp64MapOutputWithContext(ctx context.Context) Vipgrp64MapOutput {
	return o
}

func (o Vipgrp64MapOutput) MapIndex(k pulumi.StringInput) Vipgrp64Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vipgrp64 {
		return vs[0].(map[string]*Vipgrp64)[vs[1].(string)]
	}).(Vipgrp64Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Vipgrp64Input)(nil)).Elem(), &Vipgrp64{})
	pulumi.RegisterInputType(reflect.TypeOf((*Vipgrp64ArrayInput)(nil)).Elem(), Vipgrp64Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Vipgrp64MapInput)(nil)).Elem(), Vipgrp64Map{})
	pulumi.RegisterOutputType(Vipgrp64Output{})
	pulumi.RegisterOutputType(Vipgrp64ArrayOutput{})
	pulumi.RegisterOutputType(Vipgrp64MapOutput{})
}
