// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure identity based routing.
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
//			_, err := fortios.NewFirewallIdentitybasedroute(ctx, "trname", &fortios.FirewallIdentitybasedrouteArgs{
//				Comments: pulumi.String("test"),
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
// # Firewall IdentityBasedRoute can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallIdentitybasedroute:FirewallIdentitybasedroute labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallIdentitybasedroute:FirewallIdentitybasedroute labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallIdentitybasedroute struct {
	pulumi.CustomResourceState

	// Comments.
	Comments pulumi.StringOutput `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules FirewallIdentitybasedrouteRuleArrayOutput `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallIdentitybasedroute registers a new resource with the given unique name, arguments, and options.
func NewFirewallIdentitybasedroute(ctx *pulumi.Context,
	name string, args *FirewallIdentitybasedrouteArgs, opts ...pulumi.ResourceOption) (*FirewallIdentitybasedroute, error) {
	if args == nil {
		args = &FirewallIdentitybasedrouteArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallIdentitybasedroute
	err := ctx.RegisterResource("fortios:index/firewallIdentitybasedroute:FirewallIdentitybasedroute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallIdentitybasedroute gets an existing FirewallIdentitybasedroute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallIdentitybasedroute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallIdentitybasedrouteState, opts ...pulumi.ResourceOption) (*FirewallIdentitybasedroute, error) {
	var resource FirewallIdentitybasedroute
	err := ctx.ReadResource("fortios:index/firewallIdentitybasedroute:FirewallIdentitybasedroute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallIdentitybasedroute resources.
type firewallIdentitybasedrouteState struct {
	// Comments.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules []FirewallIdentitybasedrouteRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallIdentitybasedrouteState struct {
	// Comments.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Rule. The structure of `rule` block is documented below.
	Rules FirewallIdentitybasedrouteRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallIdentitybasedrouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallIdentitybasedrouteState)(nil)).Elem()
}

type firewallIdentitybasedrouteArgs struct {
	// Comments.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules []FirewallIdentitybasedrouteRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallIdentitybasedroute resource.
type FirewallIdentitybasedrouteArgs struct {
	// Comments.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Rule. The structure of `rule` block is documented below.
	Rules FirewallIdentitybasedrouteRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallIdentitybasedrouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallIdentitybasedrouteArgs)(nil)).Elem()
}

type FirewallIdentitybasedrouteInput interface {
	pulumi.Input

	ToFirewallIdentitybasedrouteOutput() FirewallIdentitybasedrouteOutput
	ToFirewallIdentitybasedrouteOutputWithContext(ctx context.Context) FirewallIdentitybasedrouteOutput
}

func (*FirewallIdentitybasedroute) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallIdentitybasedroute)(nil)).Elem()
}

func (i *FirewallIdentitybasedroute) ToFirewallIdentitybasedrouteOutput() FirewallIdentitybasedrouteOutput {
	return i.ToFirewallIdentitybasedrouteOutputWithContext(context.Background())
}

func (i *FirewallIdentitybasedroute) ToFirewallIdentitybasedrouteOutputWithContext(ctx context.Context) FirewallIdentitybasedrouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIdentitybasedrouteOutput)
}

// FirewallIdentitybasedrouteArrayInput is an input type that accepts FirewallIdentitybasedrouteArray and FirewallIdentitybasedrouteArrayOutput values.
// You can construct a concrete instance of `FirewallIdentitybasedrouteArrayInput` via:
//
//	FirewallIdentitybasedrouteArray{ FirewallIdentitybasedrouteArgs{...} }
type FirewallIdentitybasedrouteArrayInput interface {
	pulumi.Input

	ToFirewallIdentitybasedrouteArrayOutput() FirewallIdentitybasedrouteArrayOutput
	ToFirewallIdentitybasedrouteArrayOutputWithContext(context.Context) FirewallIdentitybasedrouteArrayOutput
}

type FirewallIdentitybasedrouteArray []FirewallIdentitybasedrouteInput

func (FirewallIdentitybasedrouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallIdentitybasedroute)(nil)).Elem()
}

func (i FirewallIdentitybasedrouteArray) ToFirewallIdentitybasedrouteArrayOutput() FirewallIdentitybasedrouteArrayOutput {
	return i.ToFirewallIdentitybasedrouteArrayOutputWithContext(context.Background())
}

func (i FirewallIdentitybasedrouteArray) ToFirewallIdentitybasedrouteArrayOutputWithContext(ctx context.Context) FirewallIdentitybasedrouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIdentitybasedrouteArrayOutput)
}

// FirewallIdentitybasedrouteMapInput is an input type that accepts FirewallIdentitybasedrouteMap and FirewallIdentitybasedrouteMapOutput values.
// You can construct a concrete instance of `FirewallIdentitybasedrouteMapInput` via:
//
//	FirewallIdentitybasedrouteMap{ "key": FirewallIdentitybasedrouteArgs{...} }
type FirewallIdentitybasedrouteMapInput interface {
	pulumi.Input

	ToFirewallIdentitybasedrouteMapOutput() FirewallIdentitybasedrouteMapOutput
	ToFirewallIdentitybasedrouteMapOutputWithContext(context.Context) FirewallIdentitybasedrouteMapOutput
}

type FirewallIdentitybasedrouteMap map[string]FirewallIdentitybasedrouteInput

func (FirewallIdentitybasedrouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallIdentitybasedroute)(nil)).Elem()
}

func (i FirewallIdentitybasedrouteMap) ToFirewallIdentitybasedrouteMapOutput() FirewallIdentitybasedrouteMapOutput {
	return i.ToFirewallIdentitybasedrouteMapOutputWithContext(context.Background())
}

func (i FirewallIdentitybasedrouteMap) ToFirewallIdentitybasedrouteMapOutputWithContext(ctx context.Context) FirewallIdentitybasedrouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIdentitybasedrouteMapOutput)
}

type FirewallIdentitybasedrouteOutput struct{ *pulumi.OutputState }

func (FirewallIdentitybasedrouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallIdentitybasedroute)(nil)).Elem()
}

func (o FirewallIdentitybasedrouteOutput) ToFirewallIdentitybasedrouteOutput() FirewallIdentitybasedrouteOutput {
	return o
}

func (o FirewallIdentitybasedrouteOutput) ToFirewallIdentitybasedrouteOutputWithContext(ctx context.Context) FirewallIdentitybasedrouteOutput {
	return o
}

// Comments.
func (o FirewallIdentitybasedrouteOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallIdentitybasedroute) pulumi.StringOutput { return v.Comments }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o FirewallIdentitybasedrouteOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallIdentitybasedroute) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Name.
func (o FirewallIdentitybasedrouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallIdentitybasedroute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Rule. The structure of `rule` block is documented below.
func (o FirewallIdentitybasedrouteOutput) Rules() FirewallIdentitybasedrouteRuleArrayOutput {
	return o.ApplyT(func(v *FirewallIdentitybasedroute) FirewallIdentitybasedrouteRuleArrayOutput { return v.Rules }).(FirewallIdentitybasedrouteRuleArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallIdentitybasedrouteOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallIdentitybasedroute) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallIdentitybasedrouteArrayOutput struct{ *pulumi.OutputState }

func (FirewallIdentitybasedrouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallIdentitybasedroute)(nil)).Elem()
}

func (o FirewallIdentitybasedrouteArrayOutput) ToFirewallIdentitybasedrouteArrayOutput() FirewallIdentitybasedrouteArrayOutput {
	return o
}

func (o FirewallIdentitybasedrouteArrayOutput) ToFirewallIdentitybasedrouteArrayOutputWithContext(ctx context.Context) FirewallIdentitybasedrouteArrayOutput {
	return o
}

func (o FirewallIdentitybasedrouteArrayOutput) Index(i pulumi.IntInput) FirewallIdentitybasedrouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallIdentitybasedroute {
		return vs[0].([]*FirewallIdentitybasedroute)[vs[1].(int)]
	}).(FirewallIdentitybasedrouteOutput)
}

type FirewallIdentitybasedrouteMapOutput struct{ *pulumi.OutputState }

func (FirewallIdentitybasedrouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallIdentitybasedroute)(nil)).Elem()
}

func (o FirewallIdentitybasedrouteMapOutput) ToFirewallIdentitybasedrouteMapOutput() FirewallIdentitybasedrouteMapOutput {
	return o
}

func (o FirewallIdentitybasedrouteMapOutput) ToFirewallIdentitybasedrouteMapOutputWithContext(ctx context.Context) FirewallIdentitybasedrouteMapOutput {
	return o
}

func (o FirewallIdentitybasedrouteMapOutput) MapIndex(k pulumi.StringInput) FirewallIdentitybasedrouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallIdentitybasedroute {
		return vs[0].(map[string]*FirewallIdentitybasedroute)[vs[1].(string)]
	}).(FirewallIdentitybasedrouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIdentitybasedrouteInput)(nil)).Elem(), &FirewallIdentitybasedroute{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIdentitybasedrouteArrayInput)(nil)).Elem(), FirewallIdentitybasedrouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIdentitybasedrouteMapInput)(nil)).Elem(), FirewallIdentitybasedrouteMap{})
	pulumi.RegisterOutputType(FirewallIdentitybasedrouteOutput{})
	pulumi.RegisterOutputType(FirewallIdentitybasedrouteArrayOutput{})
	pulumi.RegisterOutputType(FirewallIdentitybasedrouteMapOutput{})
}
