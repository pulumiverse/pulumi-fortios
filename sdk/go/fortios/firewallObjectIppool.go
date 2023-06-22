// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to configure IPv4 IP address pools of FortiOS.
//
// !> **Warning:** The resource will be deprecated and replaced by new resource `FirewallIppool`, we recommend that you use the new resource.
//
// ## Example Usage
// ### Overload Ippool
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
//			_, err := fortios.NewFirewallObjectIppool(ctx, "s1", &fortios.FirewallObjectIppoolArgs{
//				ArpReply: pulumi.String("enable"),
//				Comments: pulumi.String("fdsaf"),
//				Endip:    pulumi.String("22.0.0.0"),
//				Startip:  pulumi.String("11.0.0.0"),
//				Type:     pulumi.String("overload"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### One-To-One Ippool
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
//			_, err := fortios.NewFirewallObjectIppool(ctx, "s2", &fortios.FirewallObjectIppoolArgs{
//				ArpReply: pulumi.String("enable"),
//				Comments: pulumi.String("fdsaf"),
//				Endip:    pulumi.String("222.0.0.0"),
//				Startip:  pulumi.String("121.0.0.0"),
//				Type:     pulumi.String("one-to-one"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FirewallObjectIppool struct {
	pulumi.CustomResourceState

	// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
	ArpReply pulumi.StringOutput `pulumi:"arpReply"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Endip pulumi.StringOutput `pulumi:"endip"`
	// IP pool name.
	Name pulumi.StringOutput `pulumi:"name"`
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Startip pulumi.StringOutput `pulumi:"startip"`
	// IP pool type(Support overload and one-to-one).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallObjectIppool registers a new resource with the given unique name, arguments, and options.
func NewFirewallObjectIppool(ctx *pulumi.Context,
	name string, args *FirewallObjectIppoolArgs, opts ...pulumi.ResourceOption) (*FirewallObjectIppool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endip == nil {
		return nil, errors.New("invalid value for required argument 'Endip'")
	}
	if args.Startip == nil {
		return nil, errors.New("invalid value for required argument 'Startip'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallObjectIppool
	err := ctx.RegisterResource("fortios:index/firewallObjectIppool:FirewallObjectIppool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallObjectIppool gets an existing FirewallObjectIppool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallObjectIppool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallObjectIppoolState, opts ...pulumi.ResourceOption) (*FirewallObjectIppool, error) {
	var resource FirewallObjectIppool
	err := ctx.ReadResource("fortios:index/firewallObjectIppool:FirewallObjectIppool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallObjectIppool resources.
type firewallObjectIppoolState struct {
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
	ArpReply *string `pulumi:"arpReply"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Endip *string `pulumi:"endip"`
	// IP pool name.
	Name *string `pulumi:"name"`
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Startip *string `pulumi:"startip"`
	// IP pool type(Support overload and one-to-one).
	Type *string `pulumi:"type"`
}

type FirewallObjectIppoolState struct {
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
	ArpReply pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Endip pulumi.StringPtrInput
	// IP pool name.
	Name pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Startip pulumi.StringPtrInput
	// IP pool type(Support overload and one-to-one).
	Type pulumi.StringPtrInput
}

func (FirewallObjectIppoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectIppoolState)(nil)).Elem()
}

type firewallObjectIppoolArgs struct {
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
	ArpReply *string `pulumi:"arpReply"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Endip string `pulumi:"endip"`
	// IP pool name.
	Name *string `pulumi:"name"`
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Startip string `pulumi:"startip"`
	// IP pool type(Support overload and one-to-one).
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a FirewallObjectIppool resource.
type FirewallObjectIppoolArgs struct {
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
	ArpReply pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Endip pulumi.StringInput
	// IP pool name.
	Name pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Startip pulumi.StringInput
	// IP pool type(Support overload and one-to-one).
	Type pulumi.StringInput
}

func (FirewallObjectIppoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectIppoolArgs)(nil)).Elem()
}

type FirewallObjectIppoolInput interface {
	pulumi.Input

	ToFirewallObjectIppoolOutput() FirewallObjectIppoolOutput
	ToFirewallObjectIppoolOutputWithContext(ctx context.Context) FirewallObjectIppoolOutput
}

func (*FirewallObjectIppool) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectIppool)(nil)).Elem()
}

func (i *FirewallObjectIppool) ToFirewallObjectIppoolOutput() FirewallObjectIppoolOutput {
	return i.ToFirewallObjectIppoolOutputWithContext(context.Background())
}

func (i *FirewallObjectIppool) ToFirewallObjectIppoolOutputWithContext(ctx context.Context) FirewallObjectIppoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectIppoolOutput)
}

// FirewallObjectIppoolArrayInput is an input type that accepts FirewallObjectIppoolArray and FirewallObjectIppoolArrayOutput values.
// You can construct a concrete instance of `FirewallObjectIppoolArrayInput` via:
//
//	FirewallObjectIppoolArray{ FirewallObjectIppoolArgs{...} }
type FirewallObjectIppoolArrayInput interface {
	pulumi.Input

	ToFirewallObjectIppoolArrayOutput() FirewallObjectIppoolArrayOutput
	ToFirewallObjectIppoolArrayOutputWithContext(context.Context) FirewallObjectIppoolArrayOutput
}

type FirewallObjectIppoolArray []FirewallObjectIppoolInput

func (FirewallObjectIppoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallObjectIppool)(nil)).Elem()
}

func (i FirewallObjectIppoolArray) ToFirewallObjectIppoolArrayOutput() FirewallObjectIppoolArrayOutput {
	return i.ToFirewallObjectIppoolArrayOutputWithContext(context.Background())
}

func (i FirewallObjectIppoolArray) ToFirewallObjectIppoolArrayOutputWithContext(ctx context.Context) FirewallObjectIppoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectIppoolArrayOutput)
}

// FirewallObjectIppoolMapInput is an input type that accepts FirewallObjectIppoolMap and FirewallObjectIppoolMapOutput values.
// You can construct a concrete instance of `FirewallObjectIppoolMapInput` via:
//
//	FirewallObjectIppoolMap{ "key": FirewallObjectIppoolArgs{...} }
type FirewallObjectIppoolMapInput interface {
	pulumi.Input

	ToFirewallObjectIppoolMapOutput() FirewallObjectIppoolMapOutput
	ToFirewallObjectIppoolMapOutputWithContext(context.Context) FirewallObjectIppoolMapOutput
}

type FirewallObjectIppoolMap map[string]FirewallObjectIppoolInput

func (FirewallObjectIppoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallObjectIppool)(nil)).Elem()
}

func (i FirewallObjectIppoolMap) ToFirewallObjectIppoolMapOutput() FirewallObjectIppoolMapOutput {
	return i.ToFirewallObjectIppoolMapOutputWithContext(context.Background())
}

func (i FirewallObjectIppoolMap) ToFirewallObjectIppoolMapOutputWithContext(ctx context.Context) FirewallObjectIppoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectIppoolMapOutput)
}

type FirewallObjectIppoolOutput struct{ *pulumi.OutputState }

func (FirewallObjectIppoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectIppool)(nil)).Elem()
}

func (o FirewallObjectIppoolOutput) ToFirewallObjectIppoolOutput() FirewallObjectIppoolOutput {
	return o
}

func (o FirewallObjectIppoolOutput) ToFirewallObjectIppoolOutputWithContext(ctx context.Context) FirewallObjectIppoolOutput {
	return o
}

// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
func (o FirewallObjectIppoolOutput) ArpReply() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectIppool) pulumi.StringOutput { return v.ArpReply }).(pulumi.StringOutput)
}

// Comment.
func (o FirewallObjectIppoolOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallObjectIppool) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
func (o FirewallObjectIppoolOutput) Endip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectIppool) pulumi.StringOutput { return v.Endip }).(pulumi.StringOutput)
}

// IP pool name.
func (o FirewallObjectIppoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectIppool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
func (o FirewallObjectIppoolOutput) Startip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectIppool) pulumi.StringOutput { return v.Startip }).(pulumi.StringOutput)
}

// IP pool type(Support overload and one-to-one).
func (o FirewallObjectIppoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectIppool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type FirewallObjectIppoolArrayOutput struct{ *pulumi.OutputState }

func (FirewallObjectIppoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallObjectIppool)(nil)).Elem()
}

func (o FirewallObjectIppoolArrayOutput) ToFirewallObjectIppoolArrayOutput() FirewallObjectIppoolArrayOutput {
	return o
}

func (o FirewallObjectIppoolArrayOutput) ToFirewallObjectIppoolArrayOutputWithContext(ctx context.Context) FirewallObjectIppoolArrayOutput {
	return o
}

func (o FirewallObjectIppoolArrayOutput) Index(i pulumi.IntInput) FirewallObjectIppoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallObjectIppool {
		return vs[0].([]*FirewallObjectIppool)[vs[1].(int)]
	}).(FirewallObjectIppoolOutput)
}

type FirewallObjectIppoolMapOutput struct{ *pulumi.OutputState }

func (FirewallObjectIppoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallObjectIppool)(nil)).Elem()
}

func (o FirewallObjectIppoolMapOutput) ToFirewallObjectIppoolMapOutput() FirewallObjectIppoolMapOutput {
	return o
}

func (o FirewallObjectIppoolMapOutput) ToFirewallObjectIppoolMapOutputWithContext(ctx context.Context) FirewallObjectIppoolMapOutput {
	return o
}

func (o FirewallObjectIppoolMapOutput) MapIndex(k pulumi.StringInput) FirewallObjectIppoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallObjectIppool {
		return vs[0].(map[string]*FirewallObjectIppool)[vs[1].(string)]
	}).(FirewallObjectIppoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectIppoolInput)(nil)).Elem(), &FirewallObjectIppool{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectIppoolArrayInput)(nil)).Elem(), FirewallObjectIppoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectIppoolMapInput)(nil)).Elem(), FirewallObjectIppoolMap{})
	pulumi.RegisterOutputType(FirewallObjectIppoolOutput{})
	pulumi.RegisterOutputType(FirewallObjectIppoolArrayOutput{})
	pulumi.RegisterOutputType(FirewallObjectIppoolMapOutput{})
}
