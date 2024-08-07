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

// Configure IPv4 address groups.
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
//			trname1, err := firewall.NewAddress(ctx, "trname1", &firewall.AddressArgs{
//				AllowRouting: pulumi.String("disable"),
//				CacheTtl:     pulumi.Int(0),
//				Color:        pulumi.Int(0),
//				EndIp:        pulumi.String("255.0.0.0"),
//				StartIp:      pulumi.String("12.0.0.0"),
//				Subnet:       pulumi.String("12.0.0.0 255.0.0.0"),
//				Type:         pulumi.String("ipmask"),
//				Visibility:   pulumi.String("enable"),
//				Wildcard:     pulumi.String("12.0.0.0 255.0.0.0"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = firewall.NewAddrgrp(ctx, "trname", &firewall.AddrgrpArgs{
//				AllowRouting: pulumi.String("disable"),
//				Color:        pulumi.Int(0),
//				Exclude:      pulumi.String("disable"),
//				Visibility:   pulumi.String("enable"),
//				Members: firewall.AddrgrpMemberArray{
//					&firewall.AddrgrpMemberArgs{
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
// Firewall Addrgrp can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/addrgrp:Addrgrp labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/addrgrp:Addrgrp labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Addrgrp struct {
	pulumi.CustomResourceState

	// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting pulumi.StringOutput `pulumi:"allowRouting"`
	// Address group category. Valid values: `default`, `ztna-ems-tag`, `ztna-geo-tag`.
	Category pulumi.StringOutput `pulumi:"category"`
	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable address exclusion. Valid values: `enable`, `disable`.
	Exclude pulumi.StringOutput `pulumi:"exclude"`
	// Address exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers AddrgrpExcludeMemberArrayOutput `pulumi:"excludeMembers"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringOutput `pulumi:"fabricObject"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members AddrgrpMemberArrayOutput `pulumi:"members"`
	// Address group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings AddrgrpTaggingArrayOutput `pulumi:"taggings"`
	// Address group type. Valid values: `default`, `folder`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewAddrgrp registers a new resource with the given unique name, arguments, and options.
func NewAddrgrp(ctx *pulumi.Context,
	name string, args *AddrgrpArgs, opts ...pulumi.ResourceOption) (*Addrgrp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Addrgrp
	err := ctx.RegisterResource("fortios:firewall/addrgrp:Addrgrp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddrgrp gets an existing Addrgrp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddrgrp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddrgrpState, opts ...pulumi.ResourceOption) (*Addrgrp, error) {
	var resource Addrgrp
	err := ctx.ReadResource("fortios:firewall/addrgrp:Addrgrp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Addrgrp resources.
type addrgrpState struct {
	// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting *string `pulumi:"allowRouting"`
	// Address group category. Valid values: `default`, `ztna-ems-tag`, `ztna-geo-tag`.
	Category *string `pulumi:"category"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable address exclusion. Valid values: `enable`, `disable`.
	Exclude *string `pulumi:"exclude"`
	// Address exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers []AddrgrpExcludeMember `pulumi:"excludeMembers"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members []AddrgrpMember `pulumi:"members"`
	// Address group name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []AddrgrpTagging `pulumi:"taggings"`
	// Address group type. Valid values: `default`, `folder`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

type AddrgrpState struct {
	// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting pulumi.StringPtrInput
	// Address group category. Valid values: `default`, `ztna-ems-tag`, `ztna-geo-tag`.
	Category pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable address exclusion. Valid values: `enable`, `disable`.
	Exclude pulumi.StringPtrInput
	// Address exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers AddrgrpExcludeMemberArrayInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members AddrgrpMemberArrayInput
	// Address group name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings AddrgrpTaggingArrayInput
	// Address group type. Valid values: `default`, `folder`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (AddrgrpState) ElementType() reflect.Type {
	return reflect.TypeOf((*addrgrpState)(nil)).Elem()
}

type addrgrpArgs struct {
	// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting *string `pulumi:"allowRouting"`
	// Address group category. Valid values: `default`, `ztna-ems-tag`, `ztna-geo-tag`.
	Category *string `pulumi:"category"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable address exclusion. Valid values: `enable`, `disable`.
	Exclude *string `pulumi:"exclude"`
	// Address exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers []AddrgrpExcludeMember `pulumi:"excludeMembers"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members []AddrgrpMember `pulumi:"members"`
	// Address group name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []AddrgrpTagging `pulumi:"taggings"`
	// Address group type. Valid values: `default`, `folder`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a Addrgrp resource.
type AddrgrpArgs struct {
	// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting pulumi.StringPtrInput
	// Address group category. Valid values: `default`, `ztna-ems-tag`, `ztna-geo-tag`.
	Category pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable address exclusion. Valid values: `enable`, `disable`.
	Exclude pulumi.StringPtrInput
	// Address exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers AddrgrpExcludeMemberArrayInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members AddrgrpMemberArrayInput
	// Address group name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings AddrgrpTaggingArrayInput
	// Address group type. Valid values: `default`, `folder`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (AddrgrpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addrgrpArgs)(nil)).Elem()
}

type AddrgrpInput interface {
	pulumi.Input

	ToAddrgrpOutput() AddrgrpOutput
	ToAddrgrpOutputWithContext(ctx context.Context) AddrgrpOutput
}

func (*Addrgrp) ElementType() reflect.Type {
	return reflect.TypeOf((**Addrgrp)(nil)).Elem()
}

func (i *Addrgrp) ToAddrgrpOutput() AddrgrpOutput {
	return i.ToAddrgrpOutputWithContext(context.Background())
}

func (i *Addrgrp) ToAddrgrpOutputWithContext(ctx context.Context) AddrgrpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddrgrpOutput)
}

// AddrgrpArrayInput is an input type that accepts AddrgrpArray and AddrgrpArrayOutput values.
// You can construct a concrete instance of `AddrgrpArrayInput` via:
//
//	AddrgrpArray{ AddrgrpArgs{...} }
type AddrgrpArrayInput interface {
	pulumi.Input

	ToAddrgrpArrayOutput() AddrgrpArrayOutput
	ToAddrgrpArrayOutputWithContext(context.Context) AddrgrpArrayOutput
}

type AddrgrpArray []AddrgrpInput

func (AddrgrpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Addrgrp)(nil)).Elem()
}

func (i AddrgrpArray) ToAddrgrpArrayOutput() AddrgrpArrayOutput {
	return i.ToAddrgrpArrayOutputWithContext(context.Background())
}

func (i AddrgrpArray) ToAddrgrpArrayOutputWithContext(ctx context.Context) AddrgrpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddrgrpArrayOutput)
}

// AddrgrpMapInput is an input type that accepts AddrgrpMap and AddrgrpMapOutput values.
// You can construct a concrete instance of `AddrgrpMapInput` via:
//
//	AddrgrpMap{ "key": AddrgrpArgs{...} }
type AddrgrpMapInput interface {
	pulumi.Input

	ToAddrgrpMapOutput() AddrgrpMapOutput
	ToAddrgrpMapOutputWithContext(context.Context) AddrgrpMapOutput
}

type AddrgrpMap map[string]AddrgrpInput

func (AddrgrpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Addrgrp)(nil)).Elem()
}

func (i AddrgrpMap) ToAddrgrpMapOutput() AddrgrpMapOutput {
	return i.ToAddrgrpMapOutputWithContext(context.Background())
}

func (i AddrgrpMap) ToAddrgrpMapOutputWithContext(ctx context.Context) AddrgrpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddrgrpMapOutput)
}

type AddrgrpOutput struct{ *pulumi.OutputState }

func (AddrgrpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Addrgrp)(nil)).Elem()
}

func (o AddrgrpOutput) ToAddrgrpOutput() AddrgrpOutput {
	return o
}

func (o AddrgrpOutput) ToAddrgrpOutputWithContext(ctx context.Context) AddrgrpOutput {
	return o
}

// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
func (o AddrgrpOutput) AllowRouting() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp) pulumi.StringOutput { return v.AllowRouting }).(pulumi.StringOutput)
}

// Address group category. Valid values: `default`, `ztna-ems-tag`, `ztna-geo-tag`.
func (o AddrgrpOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

// Color of icon on the GUI.
func (o AddrgrpOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *Addrgrp) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o AddrgrpOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Addrgrp) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o AddrgrpOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Addrgrp) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable address exclusion. Valid values: `enable`, `disable`.
func (o AddrgrpOutput) Exclude() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp) pulumi.StringOutput { return v.Exclude }).(pulumi.StringOutput)
}

// Address exclusion member. The structure of `excludeMember` block is documented below.
func (o AddrgrpOutput) ExcludeMembers() AddrgrpExcludeMemberArrayOutput {
	return o.ApplyT(func(v *Addrgrp) AddrgrpExcludeMemberArrayOutput { return v.ExcludeMembers }).(AddrgrpExcludeMemberArrayOutput)
}

// Security Fabric global object setting. Valid values: `enable`, `disable`.
func (o AddrgrpOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp) pulumi.StringOutput { return v.FabricObject }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o AddrgrpOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Addrgrp) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Address objects contained within the group. The structure of `member` block is documented below.
func (o AddrgrpOutput) Members() AddrgrpMemberArrayOutput {
	return o.ApplyT(func(v *Addrgrp) AddrgrpMemberArrayOutput { return v.Members }).(AddrgrpMemberArrayOutput)
}

// Address group name.
func (o AddrgrpOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o AddrgrpOutput) Taggings() AddrgrpTaggingArrayOutput {
	return o.ApplyT(func(v *Addrgrp) AddrgrpTaggingArrayOutput { return v.Taggings }).(AddrgrpTaggingArrayOutput)
}

// Address group type. Valid values: `default`, `folder`.
func (o AddrgrpOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o AddrgrpOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AddrgrpOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
func (o AddrgrpOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *Addrgrp) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type AddrgrpArrayOutput struct{ *pulumi.OutputState }

func (AddrgrpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Addrgrp)(nil)).Elem()
}

func (o AddrgrpArrayOutput) ToAddrgrpArrayOutput() AddrgrpArrayOutput {
	return o
}

func (o AddrgrpArrayOutput) ToAddrgrpArrayOutputWithContext(ctx context.Context) AddrgrpArrayOutput {
	return o
}

func (o AddrgrpArrayOutput) Index(i pulumi.IntInput) AddrgrpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Addrgrp {
		return vs[0].([]*Addrgrp)[vs[1].(int)]
	}).(AddrgrpOutput)
}

type AddrgrpMapOutput struct{ *pulumi.OutputState }

func (AddrgrpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Addrgrp)(nil)).Elem()
}

func (o AddrgrpMapOutput) ToAddrgrpMapOutput() AddrgrpMapOutput {
	return o
}

func (o AddrgrpMapOutput) ToAddrgrpMapOutputWithContext(ctx context.Context) AddrgrpMapOutput {
	return o
}

func (o AddrgrpMapOutput) MapIndex(k pulumi.StringInput) AddrgrpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Addrgrp {
		return vs[0].(map[string]*Addrgrp)[vs[1].(string)]
	}).(AddrgrpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddrgrpInput)(nil)).Elem(), &Addrgrp{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddrgrpArrayInput)(nil)).Elem(), AddrgrpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddrgrpMapInput)(nil)).Elem(), AddrgrpMap{})
	pulumi.RegisterOutputType(AddrgrpOutput{})
	pulumi.RegisterOutputType(AddrgrpArrayOutput{})
	pulumi.RegisterOutputType(AddrgrpMapOutput{})
}
