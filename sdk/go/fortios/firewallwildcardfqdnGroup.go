// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Config global Wildcard FQDN address groups.
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
//			trname1, err := fortios.NewFirewallwildcardfqdnCustom(ctx, "trname1", &fortios.FirewallwildcardfqdnCustomArgs{
//				Color:        pulumi.Int(0),
//				Visibility:   pulumi.String("enable"),
//				WildcardFqdn: pulumi.String("*.ms.com"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = fortios.NewFirewallwildcardfqdnGroup(ctx, "trname", &fortios.FirewallwildcardfqdnGroupArgs{
//				Color:      pulumi.Int(0),
//				Visibility: pulumi.String("enable"),
//				Members: fortios.FirewallwildcardfqdnGroupMemberArray{
//					&fortios.FirewallwildcardfqdnGroupMemberArgs{
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
// # FirewallWildcardFqdn Group can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallwildcardfqdnGroup:FirewallwildcardfqdnGroup labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallwildcardfqdnGroup:FirewallwildcardfqdnGroup labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallwildcardfqdnGroup struct {
	pulumi.CustomResourceState

	// GUI icon color.
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Address group members. The structure of `member` block is documented below.
	Members FirewallwildcardfqdnGroupMemberArrayOutput `pulumi:"members"`
	// Address group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable address visibility. Valid values: `enable`, `disable`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewFirewallwildcardfqdnGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallwildcardfqdnGroup(ctx *pulumi.Context,
	name string, args *FirewallwildcardfqdnGroupArgs, opts ...pulumi.ResourceOption) (*FirewallwildcardfqdnGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallwildcardfqdnGroup
	err := ctx.RegisterResource("fortios:index/firewallwildcardfqdnGroup:FirewallwildcardfqdnGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallwildcardfqdnGroup gets an existing FirewallwildcardfqdnGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallwildcardfqdnGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallwildcardfqdnGroupState, opts ...pulumi.ResourceOption) (*FirewallwildcardfqdnGroup, error) {
	var resource FirewallwildcardfqdnGroup
	err := ctx.ReadResource("fortios:index/firewallwildcardfqdnGroup:FirewallwildcardfqdnGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallwildcardfqdnGroup resources.
type firewallwildcardfqdnGroupState struct {
	// GUI icon color.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Address group members. The structure of `member` block is documented below.
	Members []FirewallwildcardfqdnGroupMember `pulumi:"members"`
	// Address group name.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address visibility. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

type FirewallwildcardfqdnGroupState struct {
	// GUI icon color.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Address group members. The structure of `member` block is documented below.
	Members FirewallwildcardfqdnGroupMemberArrayInput
	// Address group name.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable address visibility. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallwildcardfqdnGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallwildcardfqdnGroupState)(nil)).Elem()
}

type firewallwildcardfqdnGroupArgs struct {
	// GUI icon color.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Address group members. The structure of `member` block is documented below.
	Members []FirewallwildcardfqdnGroupMember `pulumi:"members"`
	// Address group name.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address visibility. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a FirewallwildcardfqdnGroup resource.
type FirewallwildcardfqdnGroupArgs struct {
	// GUI icon color.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Address group members. The structure of `member` block is documented below.
	Members FirewallwildcardfqdnGroupMemberArrayInput
	// Address group name.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable address visibility. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallwildcardfqdnGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallwildcardfqdnGroupArgs)(nil)).Elem()
}

type FirewallwildcardfqdnGroupInput interface {
	pulumi.Input

	ToFirewallwildcardfqdnGroupOutput() FirewallwildcardfqdnGroupOutput
	ToFirewallwildcardfqdnGroupOutputWithContext(ctx context.Context) FirewallwildcardfqdnGroupOutput
}

func (*FirewallwildcardfqdnGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallwildcardfqdnGroup)(nil)).Elem()
}

func (i *FirewallwildcardfqdnGroup) ToFirewallwildcardfqdnGroupOutput() FirewallwildcardfqdnGroupOutput {
	return i.ToFirewallwildcardfqdnGroupOutputWithContext(context.Background())
}

func (i *FirewallwildcardfqdnGroup) ToFirewallwildcardfqdnGroupOutputWithContext(ctx context.Context) FirewallwildcardfqdnGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallwildcardfqdnGroupOutput)
}

// FirewallwildcardfqdnGroupArrayInput is an input type that accepts FirewallwildcardfqdnGroupArray and FirewallwildcardfqdnGroupArrayOutput values.
// You can construct a concrete instance of `FirewallwildcardfqdnGroupArrayInput` via:
//
//	FirewallwildcardfqdnGroupArray{ FirewallwildcardfqdnGroupArgs{...} }
type FirewallwildcardfqdnGroupArrayInput interface {
	pulumi.Input

	ToFirewallwildcardfqdnGroupArrayOutput() FirewallwildcardfqdnGroupArrayOutput
	ToFirewallwildcardfqdnGroupArrayOutputWithContext(context.Context) FirewallwildcardfqdnGroupArrayOutput
}

type FirewallwildcardfqdnGroupArray []FirewallwildcardfqdnGroupInput

func (FirewallwildcardfqdnGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallwildcardfqdnGroup)(nil)).Elem()
}

func (i FirewallwildcardfqdnGroupArray) ToFirewallwildcardfqdnGroupArrayOutput() FirewallwildcardfqdnGroupArrayOutput {
	return i.ToFirewallwildcardfqdnGroupArrayOutputWithContext(context.Background())
}

func (i FirewallwildcardfqdnGroupArray) ToFirewallwildcardfqdnGroupArrayOutputWithContext(ctx context.Context) FirewallwildcardfqdnGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallwildcardfqdnGroupArrayOutput)
}

// FirewallwildcardfqdnGroupMapInput is an input type that accepts FirewallwildcardfqdnGroupMap and FirewallwildcardfqdnGroupMapOutput values.
// You can construct a concrete instance of `FirewallwildcardfqdnGroupMapInput` via:
//
//	FirewallwildcardfqdnGroupMap{ "key": FirewallwildcardfqdnGroupArgs{...} }
type FirewallwildcardfqdnGroupMapInput interface {
	pulumi.Input

	ToFirewallwildcardfqdnGroupMapOutput() FirewallwildcardfqdnGroupMapOutput
	ToFirewallwildcardfqdnGroupMapOutputWithContext(context.Context) FirewallwildcardfqdnGroupMapOutput
}

type FirewallwildcardfqdnGroupMap map[string]FirewallwildcardfqdnGroupInput

func (FirewallwildcardfqdnGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallwildcardfqdnGroup)(nil)).Elem()
}

func (i FirewallwildcardfqdnGroupMap) ToFirewallwildcardfqdnGroupMapOutput() FirewallwildcardfqdnGroupMapOutput {
	return i.ToFirewallwildcardfqdnGroupMapOutputWithContext(context.Background())
}

func (i FirewallwildcardfqdnGroupMap) ToFirewallwildcardfqdnGroupMapOutputWithContext(ctx context.Context) FirewallwildcardfqdnGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallwildcardfqdnGroupMapOutput)
}

type FirewallwildcardfqdnGroupOutput struct{ *pulumi.OutputState }

func (FirewallwildcardfqdnGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallwildcardfqdnGroup)(nil)).Elem()
}

func (o FirewallwildcardfqdnGroupOutput) ToFirewallwildcardfqdnGroupOutput() FirewallwildcardfqdnGroupOutput {
	return o
}

func (o FirewallwildcardfqdnGroupOutput) ToFirewallwildcardfqdnGroupOutputWithContext(ctx context.Context) FirewallwildcardfqdnGroupOutput {
	return o
}

// GUI icon color.
func (o FirewallwildcardfqdnGroupOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnGroup) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o FirewallwildcardfqdnGroupOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnGroup) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o FirewallwildcardfqdnGroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnGroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Address group members. The structure of `member` block is documented below.
func (o FirewallwildcardfqdnGroupOutput) Members() FirewallwildcardfqdnGroupMemberArrayOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnGroup) FirewallwildcardfqdnGroupMemberArrayOutput { return v.Members }).(FirewallwildcardfqdnGroupMemberArrayOutput)
}

// Address group name.
func (o FirewallwildcardfqdnGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o FirewallwildcardfqdnGroupOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnGroup) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallwildcardfqdnGroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnGroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable address visibility. Valid values: `enable`, `disable`.
func (o FirewallwildcardfqdnGroupOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallwildcardfqdnGroup) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type FirewallwildcardfqdnGroupArrayOutput struct{ *pulumi.OutputState }

func (FirewallwildcardfqdnGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallwildcardfqdnGroup)(nil)).Elem()
}

func (o FirewallwildcardfqdnGroupArrayOutput) ToFirewallwildcardfqdnGroupArrayOutput() FirewallwildcardfqdnGroupArrayOutput {
	return o
}

func (o FirewallwildcardfqdnGroupArrayOutput) ToFirewallwildcardfqdnGroupArrayOutputWithContext(ctx context.Context) FirewallwildcardfqdnGroupArrayOutput {
	return o
}

func (o FirewallwildcardfqdnGroupArrayOutput) Index(i pulumi.IntInput) FirewallwildcardfqdnGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallwildcardfqdnGroup {
		return vs[0].([]*FirewallwildcardfqdnGroup)[vs[1].(int)]
	}).(FirewallwildcardfqdnGroupOutput)
}

type FirewallwildcardfqdnGroupMapOutput struct{ *pulumi.OutputState }

func (FirewallwildcardfqdnGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallwildcardfqdnGroup)(nil)).Elem()
}

func (o FirewallwildcardfqdnGroupMapOutput) ToFirewallwildcardfqdnGroupMapOutput() FirewallwildcardfqdnGroupMapOutput {
	return o
}

func (o FirewallwildcardfqdnGroupMapOutput) ToFirewallwildcardfqdnGroupMapOutputWithContext(ctx context.Context) FirewallwildcardfqdnGroupMapOutput {
	return o
}

func (o FirewallwildcardfqdnGroupMapOutput) MapIndex(k pulumi.StringInput) FirewallwildcardfqdnGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallwildcardfqdnGroup {
		return vs[0].(map[string]*FirewallwildcardfqdnGroup)[vs[1].(string)]
	}).(FirewallwildcardfqdnGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallwildcardfqdnGroupInput)(nil)).Elem(), &FirewallwildcardfqdnGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallwildcardfqdnGroupArrayInput)(nil)).Elem(), FirewallwildcardfqdnGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallwildcardfqdnGroupMapInput)(nil)).Elem(), FirewallwildcardfqdnGroupMap{})
	pulumi.RegisterOutputType(FirewallwildcardfqdnGroupOutput{})
	pulumi.RegisterOutputType(FirewallwildcardfqdnGroupArrayOutput{})
	pulumi.RegisterOutputType(FirewallwildcardfqdnGroupMapOutput{})
}
