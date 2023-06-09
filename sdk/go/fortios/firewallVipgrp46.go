// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv4 to IPv6 virtual IP groups. Applies to FortiOS Version `<= 7.0.0`.
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
//			trname1, err := fortios.NewFirewallVip46(ctx, "trname1", &fortios.FirewallVip46Args{
//				ArpReply:    pulumi.String("enable"),
//				Color:       pulumi.Int(0),
//				Extip:       pulumi.String("10.202.1.100"),
//				Extport:     pulumi.String("0-65535"),
//				Fosid:       pulumi.Int(0),
//				LdbMethod:   pulumi.String("static"),
//				Mappedip:    pulumi.String("2001:1:1:2::100"),
//				Mappedport:  pulumi.String("0-65535"),
//				Portforward: pulumi.String("disable"),
//				Protocol:    pulumi.String("tcp"),
//				Type:        pulumi.String("static-nat"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = fortios.NewFirewallVipgrp46(ctx, "trname", &fortios.FirewallVipgrp46Args{
//				Color: pulumi.Int(0),
//				Members: fortios.FirewallVipgrp46MemberArray{
//					&fortios.FirewallVipgrp46MemberArgs{
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
// # Firewall Vipgrp46 can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallVipgrp46:FirewallVipgrp46 labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallVipgrp46:FirewallVipgrp46 labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallVipgrp46 struct {
	pulumi.CustomResourceState

	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members FirewallVipgrp46MemberArrayOutput `pulumi:"members"`
	// VIP46 group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallVipgrp46 registers a new resource with the given unique name, arguments, and options.
func NewFirewallVipgrp46(ctx *pulumi.Context,
	name string, args *FirewallVipgrp46Args, opts ...pulumi.ResourceOption) (*FirewallVipgrp46, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallVipgrp46
	err := ctx.RegisterResource("fortios:index/firewallVipgrp46:FirewallVipgrp46", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallVipgrp46 gets an existing FirewallVipgrp46 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallVipgrp46(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallVipgrp46State, opts ...pulumi.ResourceOption) (*FirewallVipgrp46, error) {
	var resource FirewallVipgrp46
	err := ctx.ReadResource("fortios:index/firewallVipgrp46:FirewallVipgrp46", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallVipgrp46 resources.
type firewallVipgrp46State struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color *int `pulumi:"color"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members []FirewallVipgrp46Member `pulumi:"members"`
	// VIP46 group name.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallVipgrp46State struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members FirewallVipgrp46MemberArrayInput
	// VIP46 group name.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallVipgrp46State) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVipgrp46State)(nil)).Elem()
}

type firewallVipgrp46Args struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color *int `pulumi:"color"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members []FirewallVipgrp46Member `pulumi:"members"`
	// VIP46 group name.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallVipgrp46 resource.
type FirewallVipgrp46Args struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members FirewallVipgrp46MemberArrayInput
	// VIP46 group name.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallVipgrp46Args) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVipgrp46Args)(nil)).Elem()
}

type FirewallVipgrp46Input interface {
	pulumi.Input

	ToFirewallVipgrp46Output() FirewallVipgrp46Output
	ToFirewallVipgrp46OutputWithContext(ctx context.Context) FirewallVipgrp46Output
}

func (*FirewallVipgrp46) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVipgrp46)(nil)).Elem()
}

func (i *FirewallVipgrp46) ToFirewallVipgrp46Output() FirewallVipgrp46Output {
	return i.ToFirewallVipgrp46OutputWithContext(context.Background())
}

func (i *FirewallVipgrp46) ToFirewallVipgrp46OutputWithContext(ctx context.Context) FirewallVipgrp46Output {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVipgrp46Output)
}

// FirewallVipgrp46ArrayInput is an input type that accepts FirewallVipgrp46Array and FirewallVipgrp46ArrayOutput values.
// You can construct a concrete instance of `FirewallVipgrp46ArrayInput` via:
//
//	FirewallVipgrp46Array{ FirewallVipgrp46Args{...} }
type FirewallVipgrp46ArrayInput interface {
	pulumi.Input

	ToFirewallVipgrp46ArrayOutput() FirewallVipgrp46ArrayOutput
	ToFirewallVipgrp46ArrayOutputWithContext(context.Context) FirewallVipgrp46ArrayOutput
}

type FirewallVipgrp46Array []FirewallVipgrp46Input

func (FirewallVipgrp46Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallVipgrp46)(nil)).Elem()
}

func (i FirewallVipgrp46Array) ToFirewallVipgrp46ArrayOutput() FirewallVipgrp46ArrayOutput {
	return i.ToFirewallVipgrp46ArrayOutputWithContext(context.Background())
}

func (i FirewallVipgrp46Array) ToFirewallVipgrp46ArrayOutputWithContext(ctx context.Context) FirewallVipgrp46ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVipgrp46ArrayOutput)
}

// FirewallVipgrp46MapInput is an input type that accepts FirewallVipgrp46Map and FirewallVipgrp46MapOutput values.
// You can construct a concrete instance of `FirewallVipgrp46MapInput` via:
//
//	FirewallVipgrp46Map{ "key": FirewallVipgrp46Args{...} }
type FirewallVipgrp46MapInput interface {
	pulumi.Input

	ToFirewallVipgrp46MapOutput() FirewallVipgrp46MapOutput
	ToFirewallVipgrp46MapOutputWithContext(context.Context) FirewallVipgrp46MapOutput
}

type FirewallVipgrp46Map map[string]FirewallVipgrp46Input

func (FirewallVipgrp46Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallVipgrp46)(nil)).Elem()
}

func (i FirewallVipgrp46Map) ToFirewallVipgrp46MapOutput() FirewallVipgrp46MapOutput {
	return i.ToFirewallVipgrp46MapOutputWithContext(context.Background())
}

func (i FirewallVipgrp46Map) ToFirewallVipgrp46MapOutputWithContext(ctx context.Context) FirewallVipgrp46MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVipgrp46MapOutput)
}

type FirewallVipgrp46Output struct{ *pulumi.OutputState }

func (FirewallVipgrp46Output) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVipgrp46)(nil)).Elem()
}

func (o FirewallVipgrp46Output) ToFirewallVipgrp46Output() FirewallVipgrp46Output {
	return o
}

func (o FirewallVipgrp46Output) ToFirewallVipgrp46OutputWithContext(ctx context.Context) FirewallVipgrp46Output {
	return o
}

// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
func (o FirewallVipgrp46Output) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallVipgrp46) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o FirewallVipgrp46Output) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallVipgrp46) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o FirewallVipgrp46Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallVipgrp46) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
func (o FirewallVipgrp46Output) Members() FirewallVipgrp46MemberArrayOutput {
	return o.ApplyT(func(v *FirewallVipgrp46) FirewallVipgrp46MemberArrayOutput { return v.Members }).(FirewallVipgrp46MemberArrayOutput)
}

// VIP46 group name.
func (o FirewallVipgrp46Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVipgrp46) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o FirewallVipgrp46Output) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVipgrp46) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallVipgrp46Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallVipgrp46) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallVipgrp46ArrayOutput struct{ *pulumi.OutputState }

func (FirewallVipgrp46ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallVipgrp46)(nil)).Elem()
}

func (o FirewallVipgrp46ArrayOutput) ToFirewallVipgrp46ArrayOutput() FirewallVipgrp46ArrayOutput {
	return o
}

func (o FirewallVipgrp46ArrayOutput) ToFirewallVipgrp46ArrayOutputWithContext(ctx context.Context) FirewallVipgrp46ArrayOutput {
	return o
}

func (o FirewallVipgrp46ArrayOutput) Index(i pulumi.IntInput) FirewallVipgrp46Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallVipgrp46 {
		return vs[0].([]*FirewallVipgrp46)[vs[1].(int)]
	}).(FirewallVipgrp46Output)
}

type FirewallVipgrp46MapOutput struct{ *pulumi.OutputState }

func (FirewallVipgrp46MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallVipgrp46)(nil)).Elem()
}

func (o FirewallVipgrp46MapOutput) ToFirewallVipgrp46MapOutput() FirewallVipgrp46MapOutput {
	return o
}

func (o FirewallVipgrp46MapOutput) ToFirewallVipgrp46MapOutputWithContext(ctx context.Context) FirewallVipgrp46MapOutput {
	return o
}

func (o FirewallVipgrp46MapOutput) MapIndex(k pulumi.StringInput) FirewallVipgrp46Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallVipgrp46 {
		return vs[0].(map[string]*FirewallVipgrp46)[vs[1].(string)]
	}).(FirewallVipgrp46Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVipgrp46Input)(nil)).Elem(), &FirewallVipgrp46{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVipgrp46ArrayInput)(nil)).Elem(), FirewallVipgrp46Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVipgrp46MapInput)(nil)).Elem(), FirewallVipgrp46Map{})
	pulumi.RegisterOutputType(FirewallVipgrp46Output{})
	pulumi.RegisterOutputType(FirewallVipgrp46ArrayOutput{})
	pulumi.RegisterOutputType(FirewallVipgrp46MapOutput{})
}
