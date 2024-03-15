// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Hidden table for datasource.
//
// ## Import
//
// Waf SubClass can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:waf/subclass:Subclass labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:waf/subclass:Subclass labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Subclass struct {
	pulumi.CustomResourceState

	// Signature subclass ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Signature subclass name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSubclass registers a new resource with the given unique name, arguments, and options.
func NewSubclass(ctx *pulumi.Context,
	name string, args *SubclassArgs, opts ...pulumi.ResourceOption) (*Subclass, error) {
	if args == nil {
		args = &SubclassArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Subclass
	err := ctx.RegisterResource("fortios:waf/subclass:Subclass", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubclass gets an existing Subclass resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubclass(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubclassState, opts ...pulumi.ResourceOption) (*Subclass, error) {
	var resource Subclass
	err := ctx.ReadResource("fortios:waf/subclass:Subclass", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subclass resources.
type subclassState struct {
	// Signature subclass ID.
	Fosid *int `pulumi:"fosid"`
	// Signature subclass name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SubclassState struct {
	// Signature subclass ID.
	Fosid pulumi.IntPtrInput
	// Signature subclass name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SubclassState) ElementType() reflect.Type {
	return reflect.TypeOf((*subclassState)(nil)).Elem()
}

type subclassArgs struct {
	// Signature subclass ID.
	Fosid *int `pulumi:"fosid"`
	// Signature subclass name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Subclass resource.
type SubclassArgs struct {
	// Signature subclass ID.
	Fosid pulumi.IntPtrInput
	// Signature subclass name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SubclassArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subclassArgs)(nil)).Elem()
}

type SubclassInput interface {
	pulumi.Input

	ToSubclassOutput() SubclassOutput
	ToSubclassOutputWithContext(ctx context.Context) SubclassOutput
}

func (*Subclass) ElementType() reflect.Type {
	return reflect.TypeOf((**Subclass)(nil)).Elem()
}

func (i *Subclass) ToSubclassOutput() SubclassOutput {
	return i.ToSubclassOutputWithContext(context.Background())
}

func (i *Subclass) ToSubclassOutputWithContext(ctx context.Context) SubclassOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubclassOutput)
}

// SubclassArrayInput is an input type that accepts SubclassArray and SubclassArrayOutput values.
// You can construct a concrete instance of `SubclassArrayInput` via:
//
//	SubclassArray{ SubclassArgs{...} }
type SubclassArrayInput interface {
	pulumi.Input

	ToSubclassArrayOutput() SubclassArrayOutput
	ToSubclassArrayOutputWithContext(context.Context) SubclassArrayOutput
}

type SubclassArray []SubclassInput

func (SubclassArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Subclass)(nil)).Elem()
}

func (i SubclassArray) ToSubclassArrayOutput() SubclassArrayOutput {
	return i.ToSubclassArrayOutputWithContext(context.Background())
}

func (i SubclassArray) ToSubclassArrayOutputWithContext(ctx context.Context) SubclassArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubclassArrayOutput)
}

// SubclassMapInput is an input type that accepts SubclassMap and SubclassMapOutput values.
// You can construct a concrete instance of `SubclassMapInput` via:
//
//	SubclassMap{ "key": SubclassArgs{...} }
type SubclassMapInput interface {
	pulumi.Input

	ToSubclassMapOutput() SubclassMapOutput
	ToSubclassMapOutputWithContext(context.Context) SubclassMapOutput
}

type SubclassMap map[string]SubclassInput

func (SubclassMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Subclass)(nil)).Elem()
}

func (i SubclassMap) ToSubclassMapOutput() SubclassMapOutput {
	return i.ToSubclassMapOutputWithContext(context.Background())
}

func (i SubclassMap) ToSubclassMapOutputWithContext(ctx context.Context) SubclassMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubclassMapOutput)
}

type SubclassOutput struct{ *pulumi.OutputState }

func (SubclassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subclass)(nil)).Elem()
}

func (o SubclassOutput) ToSubclassOutput() SubclassOutput {
	return o
}

func (o SubclassOutput) ToSubclassOutputWithContext(ctx context.Context) SubclassOutput {
	return o
}

// Signature subclass ID.
func (o SubclassOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Subclass) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Signature subclass name.
func (o SubclassOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Subclass) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SubclassOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subclass) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SubclassArrayOutput struct{ *pulumi.OutputState }

func (SubclassArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Subclass)(nil)).Elem()
}

func (o SubclassArrayOutput) ToSubclassArrayOutput() SubclassArrayOutput {
	return o
}

func (o SubclassArrayOutput) ToSubclassArrayOutputWithContext(ctx context.Context) SubclassArrayOutput {
	return o
}

func (o SubclassArrayOutput) Index(i pulumi.IntInput) SubclassOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Subclass {
		return vs[0].([]*Subclass)[vs[1].(int)]
	}).(SubclassOutput)
}

type SubclassMapOutput struct{ *pulumi.OutputState }

func (SubclassMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Subclass)(nil)).Elem()
}

func (o SubclassMapOutput) ToSubclassMapOutput() SubclassMapOutput {
	return o
}

func (o SubclassMapOutput) ToSubclassMapOutputWithContext(ctx context.Context) SubclassMapOutput {
	return o
}

func (o SubclassMapOutput) MapIndex(k pulumi.StringInput) SubclassOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Subclass {
		return vs[0].(map[string]*Subclass)[vs[1].(string)]
	}).(SubclassOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubclassInput)(nil)).Elem(), &Subclass{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubclassArrayInput)(nil)).Elem(), SubclassArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubclassMapInput)(nil)).Elem(), SubclassMap{})
	pulumi.RegisterOutputType(SubclassOutput{})
	pulumi.RegisterOutputType(SubclassArrayOutput{})
	pulumi.RegisterOutputType(SubclassMapOutput{})
}
