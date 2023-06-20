// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallProxypolicyMove struct {
	pulumi.CustomResourceState

	Comment              pulumi.StringPtrOutput `pulumi:"comment"`
	Move                 pulumi.StringOutput    `pulumi:"move"`
	PolicyidDst          pulumi.IntOutput       `pulumi:"policyidDst"`
	PolicyidSrc          pulumi.IntOutput       `pulumi:"policyidSrc"`
	StatePolicySrcdstPos pulumi.StringPtrOutput `pulumi:"statePolicySrcdstPos"`
	Vdomparam            pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallProxypolicyMove registers a new resource with the given unique name, arguments, and options.
func NewFirewallProxypolicyMove(ctx *pulumi.Context,
	name string, args *FirewallProxypolicyMoveArgs, opts ...pulumi.ResourceOption) (*FirewallProxypolicyMove, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Move == nil {
		return nil, errors.New("invalid value for required argument 'Move'")
	}
	if args.PolicyidDst == nil {
		return nil, errors.New("invalid value for required argument 'PolicyidDst'")
	}
	if args.PolicyidSrc == nil {
		return nil, errors.New("invalid value for required argument 'PolicyidSrc'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallProxypolicyMove
	err := ctx.RegisterResource("fortios:index/firewallProxypolicyMove:FirewallProxypolicyMove", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallProxypolicyMove gets an existing FirewallProxypolicyMove resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallProxypolicyMove(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallProxypolicyMoveState, opts ...pulumi.ResourceOption) (*FirewallProxypolicyMove, error) {
	var resource FirewallProxypolicyMove
	err := ctx.ReadResource("fortios:index/firewallProxypolicyMove:FirewallProxypolicyMove", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallProxypolicyMove resources.
type firewallProxypolicyMoveState struct {
	Comment              *string `pulumi:"comment"`
	Move                 *string `pulumi:"move"`
	PolicyidDst          *int    `pulumi:"policyidDst"`
	PolicyidSrc          *int    `pulumi:"policyidSrc"`
	StatePolicySrcdstPos *string `pulumi:"statePolicySrcdstPos"`
	Vdomparam            *string `pulumi:"vdomparam"`
}

type FirewallProxypolicyMoveState struct {
	Comment              pulumi.StringPtrInput
	Move                 pulumi.StringPtrInput
	PolicyidDst          pulumi.IntPtrInput
	PolicyidSrc          pulumi.IntPtrInput
	StatePolicySrcdstPos pulumi.StringPtrInput
	Vdomparam            pulumi.StringPtrInput
}

func (FirewallProxypolicyMoveState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProxypolicyMoveState)(nil)).Elem()
}

type firewallProxypolicyMoveArgs struct {
	Comment              *string `pulumi:"comment"`
	Move                 string  `pulumi:"move"`
	PolicyidDst          int     `pulumi:"policyidDst"`
	PolicyidSrc          int     `pulumi:"policyidSrc"`
	StatePolicySrcdstPos *string `pulumi:"statePolicySrcdstPos"`
	Vdomparam            *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallProxypolicyMove resource.
type FirewallProxypolicyMoveArgs struct {
	Comment              pulumi.StringPtrInput
	Move                 pulumi.StringInput
	PolicyidDst          pulumi.IntInput
	PolicyidSrc          pulumi.IntInput
	StatePolicySrcdstPos pulumi.StringPtrInput
	Vdomparam            pulumi.StringPtrInput
}

func (FirewallProxypolicyMoveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProxypolicyMoveArgs)(nil)).Elem()
}

type FirewallProxypolicyMoveInput interface {
	pulumi.Input

	ToFirewallProxypolicyMoveOutput() FirewallProxypolicyMoveOutput
	ToFirewallProxypolicyMoveOutputWithContext(ctx context.Context) FirewallProxypolicyMoveOutput
}

func (*FirewallProxypolicyMove) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProxypolicyMove)(nil)).Elem()
}

func (i *FirewallProxypolicyMove) ToFirewallProxypolicyMoveOutput() FirewallProxypolicyMoveOutput {
	return i.ToFirewallProxypolicyMoveOutputWithContext(context.Background())
}

func (i *FirewallProxypolicyMove) ToFirewallProxypolicyMoveOutputWithContext(ctx context.Context) FirewallProxypolicyMoveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxypolicyMoveOutput)
}

// FirewallProxypolicyMoveArrayInput is an input type that accepts FirewallProxypolicyMoveArray and FirewallProxypolicyMoveArrayOutput values.
// You can construct a concrete instance of `FirewallProxypolicyMoveArrayInput` via:
//
//	FirewallProxypolicyMoveArray{ FirewallProxypolicyMoveArgs{...} }
type FirewallProxypolicyMoveArrayInput interface {
	pulumi.Input

	ToFirewallProxypolicyMoveArrayOutput() FirewallProxypolicyMoveArrayOutput
	ToFirewallProxypolicyMoveArrayOutputWithContext(context.Context) FirewallProxypolicyMoveArrayOutput
}

type FirewallProxypolicyMoveArray []FirewallProxypolicyMoveInput

func (FirewallProxypolicyMoveArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallProxypolicyMove)(nil)).Elem()
}

func (i FirewallProxypolicyMoveArray) ToFirewallProxypolicyMoveArrayOutput() FirewallProxypolicyMoveArrayOutput {
	return i.ToFirewallProxypolicyMoveArrayOutputWithContext(context.Background())
}

func (i FirewallProxypolicyMoveArray) ToFirewallProxypolicyMoveArrayOutputWithContext(ctx context.Context) FirewallProxypolicyMoveArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxypolicyMoveArrayOutput)
}

// FirewallProxypolicyMoveMapInput is an input type that accepts FirewallProxypolicyMoveMap and FirewallProxypolicyMoveMapOutput values.
// You can construct a concrete instance of `FirewallProxypolicyMoveMapInput` via:
//
//	FirewallProxypolicyMoveMap{ "key": FirewallProxypolicyMoveArgs{...} }
type FirewallProxypolicyMoveMapInput interface {
	pulumi.Input

	ToFirewallProxypolicyMoveMapOutput() FirewallProxypolicyMoveMapOutput
	ToFirewallProxypolicyMoveMapOutputWithContext(context.Context) FirewallProxypolicyMoveMapOutput
}

type FirewallProxypolicyMoveMap map[string]FirewallProxypolicyMoveInput

func (FirewallProxypolicyMoveMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallProxypolicyMove)(nil)).Elem()
}

func (i FirewallProxypolicyMoveMap) ToFirewallProxypolicyMoveMapOutput() FirewallProxypolicyMoveMapOutput {
	return i.ToFirewallProxypolicyMoveMapOutputWithContext(context.Background())
}

func (i FirewallProxypolicyMoveMap) ToFirewallProxypolicyMoveMapOutputWithContext(ctx context.Context) FirewallProxypolicyMoveMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxypolicyMoveMapOutput)
}

type FirewallProxypolicyMoveOutput struct{ *pulumi.OutputState }

func (FirewallProxypolicyMoveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProxypolicyMove)(nil)).Elem()
}

func (o FirewallProxypolicyMoveOutput) ToFirewallProxypolicyMoveOutput() FirewallProxypolicyMoveOutput {
	return o
}

func (o FirewallProxypolicyMoveOutput) ToFirewallProxypolicyMoveOutputWithContext(ctx context.Context) FirewallProxypolicyMoveOutput {
	return o
}

func (o FirewallProxypolicyMoveOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallProxypolicyMove) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o FirewallProxypolicyMoveOutput) Move() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxypolicyMove) pulumi.StringOutput { return v.Move }).(pulumi.StringOutput)
}

func (o FirewallProxypolicyMoveOutput) PolicyidDst() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallProxypolicyMove) pulumi.IntOutput { return v.PolicyidDst }).(pulumi.IntOutput)
}

func (o FirewallProxypolicyMoveOutput) PolicyidSrc() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallProxypolicyMove) pulumi.IntOutput { return v.PolicyidSrc }).(pulumi.IntOutput)
}

func (o FirewallProxypolicyMoveOutput) StatePolicySrcdstPos() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallProxypolicyMove) pulumi.StringPtrOutput { return v.StatePolicySrcdstPos }).(pulumi.StringPtrOutput)
}

func (o FirewallProxypolicyMoveOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallProxypolicyMove) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallProxypolicyMoveArrayOutput struct{ *pulumi.OutputState }

func (FirewallProxypolicyMoveArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallProxypolicyMove)(nil)).Elem()
}

func (o FirewallProxypolicyMoveArrayOutput) ToFirewallProxypolicyMoveArrayOutput() FirewallProxypolicyMoveArrayOutput {
	return o
}

func (o FirewallProxypolicyMoveArrayOutput) ToFirewallProxypolicyMoveArrayOutputWithContext(ctx context.Context) FirewallProxypolicyMoveArrayOutput {
	return o
}

func (o FirewallProxypolicyMoveArrayOutput) Index(i pulumi.IntInput) FirewallProxypolicyMoveOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallProxypolicyMove {
		return vs[0].([]*FirewallProxypolicyMove)[vs[1].(int)]
	}).(FirewallProxypolicyMoveOutput)
}

type FirewallProxypolicyMoveMapOutput struct{ *pulumi.OutputState }

func (FirewallProxypolicyMoveMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallProxypolicyMove)(nil)).Elem()
}

func (o FirewallProxypolicyMoveMapOutput) ToFirewallProxypolicyMoveMapOutput() FirewallProxypolicyMoveMapOutput {
	return o
}

func (o FirewallProxypolicyMoveMapOutput) ToFirewallProxypolicyMoveMapOutputWithContext(ctx context.Context) FirewallProxypolicyMoveMapOutput {
	return o
}

func (o FirewallProxypolicyMoveMapOutput) MapIndex(k pulumi.StringInput) FirewallProxypolicyMoveOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallProxypolicyMove {
		return vs[0].(map[string]*FirewallProxypolicyMove)[vs[1].(string)]
	}).(FirewallProxypolicyMoveOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxypolicyMoveInput)(nil)).Elem(), &FirewallProxypolicyMove{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxypolicyMoveArrayInput)(nil)).Elem(), FirewallProxypolicyMoveArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxypolicyMoveMapInput)(nil)).Elem(), FirewallProxypolicyMoveMap{})
	pulumi.RegisterOutputType(FirewallProxypolicyMoveOutput{})
	pulumi.RegisterOutputType(FirewallProxypolicyMoveArrayOutput{})
	pulumi.RegisterOutputType(FirewallProxypolicyMoveMapOutput{})
}