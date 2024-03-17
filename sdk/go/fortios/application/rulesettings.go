// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package application

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure application rule settings.
//
// ## Import
//
// Application RuleSettings can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:application/rulesettings:Rulesettings labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:application/rulesettings:Rulesettings labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Rulesettings struct {
	pulumi.CustomResourceState

	// Rule ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRulesettings registers a new resource with the given unique name, arguments, and options.
func NewRulesettings(ctx *pulumi.Context,
	name string, args *RulesettingsArgs, opts ...pulumi.ResourceOption) (*Rulesettings, error) {
	if args == nil {
		args = &RulesettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rulesettings
	err := ctx.RegisterResource("fortios:application/rulesettings:Rulesettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRulesettings gets an existing Rulesettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRulesettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RulesettingsState, opts ...pulumi.ResourceOption) (*Rulesettings, error) {
	var resource Rulesettings
	err := ctx.ReadResource("fortios:application/rulesettings:Rulesettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rulesettings resources.
type rulesettingsState struct {
	// Rule ID.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RulesettingsState struct {
	// Rule ID.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RulesettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*rulesettingsState)(nil)).Elem()
}

type rulesettingsArgs struct {
	// Rule ID.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Rulesettings resource.
type RulesettingsArgs struct {
	// Rule ID.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RulesettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rulesettingsArgs)(nil)).Elem()
}

type RulesettingsInput interface {
	pulumi.Input

	ToRulesettingsOutput() RulesettingsOutput
	ToRulesettingsOutputWithContext(ctx context.Context) RulesettingsOutput
}

func (*Rulesettings) ElementType() reflect.Type {
	return reflect.TypeOf((**Rulesettings)(nil)).Elem()
}

func (i *Rulesettings) ToRulesettingsOutput() RulesettingsOutput {
	return i.ToRulesettingsOutputWithContext(context.Background())
}

func (i *Rulesettings) ToRulesettingsOutputWithContext(ctx context.Context) RulesettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesettingsOutput)
}

// RulesettingsArrayInput is an input type that accepts RulesettingsArray and RulesettingsArrayOutput values.
// You can construct a concrete instance of `RulesettingsArrayInput` via:
//
//	RulesettingsArray{ RulesettingsArgs{...} }
type RulesettingsArrayInput interface {
	pulumi.Input

	ToRulesettingsArrayOutput() RulesettingsArrayOutput
	ToRulesettingsArrayOutputWithContext(context.Context) RulesettingsArrayOutput
}

type RulesettingsArray []RulesettingsInput

func (RulesettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rulesettings)(nil)).Elem()
}

func (i RulesettingsArray) ToRulesettingsArrayOutput() RulesettingsArrayOutput {
	return i.ToRulesettingsArrayOutputWithContext(context.Background())
}

func (i RulesettingsArray) ToRulesettingsArrayOutputWithContext(ctx context.Context) RulesettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesettingsArrayOutput)
}

// RulesettingsMapInput is an input type that accepts RulesettingsMap and RulesettingsMapOutput values.
// You can construct a concrete instance of `RulesettingsMapInput` via:
//
//	RulesettingsMap{ "key": RulesettingsArgs{...} }
type RulesettingsMapInput interface {
	pulumi.Input

	ToRulesettingsMapOutput() RulesettingsMapOutput
	ToRulesettingsMapOutputWithContext(context.Context) RulesettingsMapOutput
}

type RulesettingsMap map[string]RulesettingsInput

func (RulesettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rulesettings)(nil)).Elem()
}

func (i RulesettingsMap) ToRulesettingsMapOutput() RulesettingsMapOutput {
	return i.ToRulesettingsMapOutputWithContext(context.Background())
}

func (i RulesettingsMap) ToRulesettingsMapOutputWithContext(ctx context.Context) RulesettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesettingsMapOutput)
}

type RulesettingsOutput struct{ *pulumi.OutputState }

func (RulesettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rulesettings)(nil)).Elem()
}

func (o RulesettingsOutput) ToRulesettingsOutput() RulesettingsOutput {
	return o
}

func (o RulesettingsOutput) ToRulesettingsOutputWithContext(ctx context.Context) RulesettingsOutput {
	return o
}

// Rule ID.
func (o RulesettingsOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Rulesettings) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RulesettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rulesettings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RulesettingsArrayOutput struct{ *pulumi.OutputState }

func (RulesettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rulesettings)(nil)).Elem()
}

func (o RulesettingsArrayOutput) ToRulesettingsArrayOutput() RulesettingsArrayOutput {
	return o
}

func (o RulesettingsArrayOutput) ToRulesettingsArrayOutputWithContext(ctx context.Context) RulesettingsArrayOutput {
	return o
}

func (o RulesettingsArrayOutput) Index(i pulumi.IntInput) RulesettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rulesettings {
		return vs[0].([]*Rulesettings)[vs[1].(int)]
	}).(RulesettingsOutput)
}

type RulesettingsMapOutput struct{ *pulumi.OutputState }

func (RulesettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rulesettings)(nil)).Elem()
}

func (o RulesettingsMapOutput) ToRulesettingsMapOutput() RulesettingsMapOutput {
	return o
}

func (o RulesettingsMapOutput) ToRulesettingsMapOutputWithContext(ctx context.Context) RulesettingsMapOutput {
	return o
}

func (o RulesettingsMapOutput) MapIndex(k pulumi.StringInput) RulesettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rulesettings {
		return vs[0].(map[string]*Rulesettings)[vs[1].(string)]
	}).(RulesettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RulesettingsInput)(nil)).Elem(), &Rulesettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*RulesettingsArrayInput)(nil)).Elem(), RulesettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RulesettingsMapInput)(nil)).Elem(), RulesettingsMap{})
	pulumi.RegisterOutputType(RulesettingsOutput{})
	pulumi.RegisterOutputType(RulesettingsArrayOutput{})
	pulumi.RegisterOutputType(RulesettingsMapOutput{})
}
