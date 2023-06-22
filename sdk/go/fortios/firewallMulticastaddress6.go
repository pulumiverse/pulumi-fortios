// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv6 multicast address.
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
//			_, err := fortios.NewFirewallMulticastaddress6(ctx, "trname", &fortios.FirewallMulticastaddress6Args{
//				Color:      pulumi.Int(0),
//				Ip6:        pulumi.String("ff02::1:ff0e:8c6c/128"),
//				Visibility: pulumi.String("enable"),
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
// # Firewall MulticastAddress6 can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallMulticastaddress6:FirewallMulticastaddress6 labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallMulticastaddress6:FirewallMulticastaddress6 labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallMulticastaddress6 struct {
	pulumi.CustomResourceState

	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 pulumi.StringOutput `pulumi:"ip6"`
	// IPv6 multicast address name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallMulticastaddress6TaggingArrayOutput `pulumi:"taggings"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewFirewallMulticastaddress6 registers a new resource with the given unique name, arguments, and options.
func NewFirewallMulticastaddress6(ctx *pulumi.Context,
	name string, args *FirewallMulticastaddress6Args, opts ...pulumi.ResourceOption) (*FirewallMulticastaddress6, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip6 == nil {
		return nil, errors.New("invalid value for required argument 'Ip6'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallMulticastaddress6
	err := ctx.RegisterResource("fortios:index/firewallMulticastaddress6:FirewallMulticastaddress6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallMulticastaddress6 gets an existing FirewallMulticastaddress6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallMulticastaddress6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallMulticastaddress6State, opts ...pulumi.ResourceOption) (*FirewallMulticastaddress6, error) {
	var resource FirewallMulticastaddress6
	err := ctx.ReadResource("fortios:index/firewallMulticastaddress6:FirewallMulticastaddress6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallMulticastaddress6 resources.
type firewallMulticastaddress6State struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 *string `pulumi:"ip6"`
	// IPv6 multicast address name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []FirewallMulticastaddress6Tagging `pulumi:"taggings"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

type FirewallMulticastaddress6State struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 pulumi.StringPtrInput
	// IPv6 multicast address name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallMulticastaddress6TaggingArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallMulticastaddress6State) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallMulticastaddress6State)(nil)).Elem()
}

type firewallMulticastaddress6Args struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 string `pulumi:"ip6"`
	// IPv6 multicast address name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []FirewallMulticastaddress6Tagging `pulumi:"taggings"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a FirewallMulticastaddress6 resource.
type FirewallMulticastaddress6Args struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 pulumi.StringInput
	// IPv6 multicast address name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallMulticastaddress6TaggingArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallMulticastaddress6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallMulticastaddress6Args)(nil)).Elem()
}

type FirewallMulticastaddress6Input interface {
	pulumi.Input

	ToFirewallMulticastaddress6Output() FirewallMulticastaddress6Output
	ToFirewallMulticastaddress6OutputWithContext(ctx context.Context) FirewallMulticastaddress6Output
}

func (*FirewallMulticastaddress6) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallMulticastaddress6)(nil)).Elem()
}

func (i *FirewallMulticastaddress6) ToFirewallMulticastaddress6Output() FirewallMulticastaddress6Output {
	return i.ToFirewallMulticastaddress6OutputWithContext(context.Background())
}

func (i *FirewallMulticastaddress6) ToFirewallMulticastaddress6OutputWithContext(ctx context.Context) FirewallMulticastaddress6Output {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallMulticastaddress6Output)
}

// FirewallMulticastaddress6ArrayInput is an input type that accepts FirewallMulticastaddress6Array and FirewallMulticastaddress6ArrayOutput values.
// You can construct a concrete instance of `FirewallMulticastaddress6ArrayInput` via:
//
//	FirewallMulticastaddress6Array{ FirewallMulticastaddress6Args{...} }
type FirewallMulticastaddress6ArrayInput interface {
	pulumi.Input

	ToFirewallMulticastaddress6ArrayOutput() FirewallMulticastaddress6ArrayOutput
	ToFirewallMulticastaddress6ArrayOutputWithContext(context.Context) FirewallMulticastaddress6ArrayOutput
}

type FirewallMulticastaddress6Array []FirewallMulticastaddress6Input

func (FirewallMulticastaddress6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallMulticastaddress6)(nil)).Elem()
}

func (i FirewallMulticastaddress6Array) ToFirewallMulticastaddress6ArrayOutput() FirewallMulticastaddress6ArrayOutput {
	return i.ToFirewallMulticastaddress6ArrayOutputWithContext(context.Background())
}

func (i FirewallMulticastaddress6Array) ToFirewallMulticastaddress6ArrayOutputWithContext(ctx context.Context) FirewallMulticastaddress6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallMulticastaddress6ArrayOutput)
}

// FirewallMulticastaddress6MapInput is an input type that accepts FirewallMulticastaddress6Map and FirewallMulticastaddress6MapOutput values.
// You can construct a concrete instance of `FirewallMulticastaddress6MapInput` via:
//
//	FirewallMulticastaddress6Map{ "key": FirewallMulticastaddress6Args{...} }
type FirewallMulticastaddress6MapInput interface {
	pulumi.Input

	ToFirewallMulticastaddress6MapOutput() FirewallMulticastaddress6MapOutput
	ToFirewallMulticastaddress6MapOutputWithContext(context.Context) FirewallMulticastaddress6MapOutput
}

type FirewallMulticastaddress6Map map[string]FirewallMulticastaddress6Input

func (FirewallMulticastaddress6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallMulticastaddress6)(nil)).Elem()
}

func (i FirewallMulticastaddress6Map) ToFirewallMulticastaddress6MapOutput() FirewallMulticastaddress6MapOutput {
	return i.ToFirewallMulticastaddress6MapOutputWithContext(context.Background())
}

func (i FirewallMulticastaddress6Map) ToFirewallMulticastaddress6MapOutputWithContext(ctx context.Context) FirewallMulticastaddress6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallMulticastaddress6MapOutput)
}

type FirewallMulticastaddress6Output struct{ *pulumi.OutputState }

func (FirewallMulticastaddress6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallMulticastaddress6)(nil)).Elem()
}

func (o FirewallMulticastaddress6Output) ToFirewallMulticastaddress6Output() FirewallMulticastaddress6Output {
	return o
}

func (o FirewallMulticastaddress6Output) ToFirewallMulticastaddress6OutputWithContext(ctx context.Context) FirewallMulticastaddress6Output {
	return o
}

// Color of icon on the GUI.
func (o FirewallMulticastaddress6Output) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallMulticastaddress6) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o FirewallMulticastaddress6Output) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallMulticastaddress6) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o FirewallMulticastaddress6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallMulticastaddress6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
func (o FirewallMulticastaddress6Output) Ip6() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallMulticastaddress6) pulumi.StringOutput { return v.Ip6 }).(pulumi.StringOutput)
}

// IPv6 multicast address name.
func (o FirewallMulticastaddress6Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallMulticastaddress6) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o FirewallMulticastaddress6Output) Taggings() FirewallMulticastaddress6TaggingArrayOutput {
	return o.ApplyT(func(v *FirewallMulticastaddress6) FirewallMulticastaddress6TaggingArrayOutput { return v.Taggings }).(FirewallMulticastaddress6TaggingArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallMulticastaddress6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallMulticastaddress6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
func (o FirewallMulticastaddress6Output) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallMulticastaddress6) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type FirewallMulticastaddress6ArrayOutput struct{ *pulumi.OutputState }

func (FirewallMulticastaddress6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallMulticastaddress6)(nil)).Elem()
}

func (o FirewallMulticastaddress6ArrayOutput) ToFirewallMulticastaddress6ArrayOutput() FirewallMulticastaddress6ArrayOutput {
	return o
}

func (o FirewallMulticastaddress6ArrayOutput) ToFirewallMulticastaddress6ArrayOutputWithContext(ctx context.Context) FirewallMulticastaddress6ArrayOutput {
	return o
}

func (o FirewallMulticastaddress6ArrayOutput) Index(i pulumi.IntInput) FirewallMulticastaddress6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallMulticastaddress6 {
		return vs[0].([]*FirewallMulticastaddress6)[vs[1].(int)]
	}).(FirewallMulticastaddress6Output)
}

type FirewallMulticastaddress6MapOutput struct{ *pulumi.OutputState }

func (FirewallMulticastaddress6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallMulticastaddress6)(nil)).Elem()
}

func (o FirewallMulticastaddress6MapOutput) ToFirewallMulticastaddress6MapOutput() FirewallMulticastaddress6MapOutput {
	return o
}

func (o FirewallMulticastaddress6MapOutput) ToFirewallMulticastaddress6MapOutputWithContext(ctx context.Context) FirewallMulticastaddress6MapOutput {
	return o
}

func (o FirewallMulticastaddress6MapOutput) MapIndex(k pulumi.StringInput) FirewallMulticastaddress6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallMulticastaddress6 {
		return vs[0].(map[string]*FirewallMulticastaddress6)[vs[1].(string)]
	}).(FirewallMulticastaddress6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallMulticastaddress6Input)(nil)).Elem(), &FirewallMulticastaddress6{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallMulticastaddress6ArrayInput)(nil)).Elem(), FirewallMulticastaddress6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallMulticastaddress6MapInput)(nil)).Elem(), FirewallMulticastaddress6Map{})
	pulumi.RegisterOutputType(FirewallMulticastaddress6Output{})
	pulumi.RegisterOutputType(FirewallMulticastaddress6ArrayOutput{})
	pulumi.RegisterOutputType(FirewallMulticastaddress6MapOutput{})
}
