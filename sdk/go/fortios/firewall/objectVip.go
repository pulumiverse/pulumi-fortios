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

// Provides a resource to configure firewall virtual IPs (VIPs) of FortiOS.
//
// !> **Warning:** The resource will be deprecated and replaced by new resource `firewall.Vip`, we recommend that you use the new resource.
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
//			_, err := firewall.NewObjectVip(ctx, "v11", &firewall.ObjectVipArgs{
//				Comment: pulumi.String("fdsafdsafds"),
//				Extintf: pulumi.String("port3"),
//				Extip:   pulumi.String("11.1.1.1-21.1.1.1"),
//				Extport: pulumi.String("2-3"),
//				Mappedips: pulumi.StringArray{
//					pulumi.String("22.2.2.2-32.2.2.2"),
//				},
//				Mappedport:  pulumi.String("4-5"),
//				Portforward: pulumi.String("enable"),
//				Protocol:    pulumi.String("tcp"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ObjectVip struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	Extintf pulumi.StringOutput `pulumi:"extintf"`
	// IP address or address range on the external interface that you want to map to an address or address range on the
	// destination network.
	Extip pulumi.StringOutput `pulumi:"extip"`
	// Incoming port number range that you want to map to a port number range on the destination network.
	Extport pulumi.StringOutput `pulumi:"extport"`
	// IP address or address range on the destination network to which the external IP address is mapped.
	Mappedips pulumi.StringArrayOutput `pulumi:"mappedips"`
	// Port number range on the destination network to which the external port number range is mapped.
	Mappedport pulumi.StringOutput `pulumi:"mappedport"`
	// Virtual IP name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable port forwarding.
	Portforward pulumi.StringOutput `pulumi:"portforward"`
	// Protocol to use when forwarding packets.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
}

// NewObjectVip registers a new resource with the given unique name, arguments, and options.
func NewObjectVip(ctx *pulumi.Context,
	name string, args *ObjectVipArgs, opts ...pulumi.ResourceOption) (*ObjectVip, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Extip == nil {
		return nil, errors.New("invalid value for required argument 'Extip'")
	}
	if args.Mappedips == nil {
		return nil, errors.New("invalid value for required argument 'Mappedips'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ObjectVip
	err := ctx.RegisterResource("fortios:firewall/objectVip:ObjectVip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectVip gets an existing ObjectVip resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectVip(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectVipState, opts ...pulumi.ResourceOption) (*ObjectVip, error) {
	var resource ObjectVip
	err := ctx.ReadResource("fortios:firewall/objectVip:ObjectVip", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectVip resources.
type objectVipState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	Extintf *string `pulumi:"extintf"`
	// IP address or address range on the external interface that you want to map to an address or address range on the
	// destination network.
	Extip *string `pulumi:"extip"`
	// Incoming port number range that you want to map to a port number range on the destination network.
	Extport *string `pulumi:"extport"`
	// IP address or address range on the destination network to which the external IP address is mapped.
	Mappedips []string `pulumi:"mappedips"`
	// Port number range on the destination network to which the external port number range is mapped.
	Mappedport *string `pulumi:"mappedport"`
	// Virtual IP name.
	Name *string `pulumi:"name"`
	// Enable/disable port forwarding.
	Portforward *string `pulumi:"portforward"`
	// Protocol to use when forwarding packets.
	Protocol *string `pulumi:"protocol"`
}

type ObjectVipState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	Extintf pulumi.StringPtrInput
	// IP address or address range on the external interface that you want to map to an address or address range on the
	// destination network.
	Extip pulumi.StringPtrInput
	// Incoming port number range that you want to map to a port number range on the destination network.
	Extport pulumi.StringPtrInput
	// IP address or address range on the destination network to which the external IP address is mapped.
	Mappedips pulumi.StringArrayInput
	// Port number range on the destination network to which the external port number range is mapped.
	Mappedport pulumi.StringPtrInput
	// Virtual IP name.
	Name pulumi.StringPtrInput
	// Enable/disable port forwarding.
	Portforward pulumi.StringPtrInput
	// Protocol to use when forwarding packets.
	Protocol pulumi.StringPtrInput
}

func (ObjectVipState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectVipState)(nil)).Elem()
}

type objectVipArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	Extintf *string `pulumi:"extintf"`
	// IP address or address range on the external interface that you want to map to an address or address range on the
	// destination network.
	Extip string `pulumi:"extip"`
	// Incoming port number range that you want to map to a port number range on the destination network.
	Extport *string `pulumi:"extport"`
	// IP address or address range on the destination network to which the external IP address is mapped.
	Mappedips []string `pulumi:"mappedips"`
	// Port number range on the destination network to which the external port number range is mapped.
	Mappedport *string `pulumi:"mappedport"`
	// Virtual IP name.
	Name *string `pulumi:"name"`
	// Enable/disable port forwarding.
	Portforward *string `pulumi:"portforward"`
	// Protocol to use when forwarding packets.
	Protocol *string `pulumi:"protocol"`
}

// The set of arguments for constructing a ObjectVip resource.
type ObjectVipArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	Extintf pulumi.StringPtrInput
	// IP address or address range on the external interface that you want to map to an address or address range on the
	// destination network.
	Extip pulumi.StringInput
	// Incoming port number range that you want to map to a port number range on the destination network.
	Extport pulumi.StringPtrInput
	// IP address or address range on the destination network to which the external IP address is mapped.
	Mappedips pulumi.StringArrayInput
	// Port number range on the destination network to which the external port number range is mapped.
	Mappedport pulumi.StringPtrInput
	// Virtual IP name.
	Name pulumi.StringPtrInput
	// Enable/disable port forwarding.
	Portforward pulumi.StringPtrInput
	// Protocol to use when forwarding packets.
	Protocol pulumi.StringPtrInput
}

func (ObjectVipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectVipArgs)(nil)).Elem()
}

type ObjectVipInput interface {
	pulumi.Input

	ToObjectVipOutput() ObjectVipOutput
	ToObjectVipOutputWithContext(ctx context.Context) ObjectVipOutput
}

func (*ObjectVip) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectVip)(nil)).Elem()
}

func (i *ObjectVip) ToObjectVipOutput() ObjectVipOutput {
	return i.ToObjectVipOutputWithContext(context.Background())
}

func (i *ObjectVip) ToObjectVipOutputWithContext(ctx context.Context) ObjectVipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectVipOutput)
}

// ObjectVipArrayInput is an input type that accepts ObjectVipArray and ObjectVipArrayOutput values.
// You can construct a concrete instance of `ObjectVipArrayInput` via:
//
//	ObjectVipArray{ ObjectVipArgs{...} }
type ObjectVipArrayInput interface {
	pulumi.Input

	ToObjectVipArrayOutput() ObjectVipArrayOutput
	ToObjectVipArrayOutputWithContext(context.Context) ObjectVipArrayOutput
}

type ObjectVipArray []ObjectVipInput

func (ObjectVipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectVip)(nil)).Elem()
}

func (i ObjectVipArray) ToObjectVipArrayOutput() ObjectVipArrayOutput {
	return i.ToObjectVipArrayOutputWithContext(context.Background())
}

func (i ObjectVipArray) ToObjectVipArrayOutputWithContext(ctx context.Context) ObjectVipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectVipArrayOutput)
}

// ObjectVipMapInput is an input type that accepts ObjectVipMap and ObjectVipMapOutput values.
// You can construct a concrete instance of `ObjectVipMapInput` via:
//
//	ObjectVipMap{ "key": ObjectVipArgs{...} }
type ObjectVipMapInput interface {
	pulumi.Input

	ToObjectVipMapOutput() ObjectVipMapOutput
	ToObjectVipMapOutputWithContext(context.Context) ObjectVipMapOutput
}

type ObjectVipMap map[string]ObjectVipInput

func (ObjectVipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectVip)(nil)).Elem()
}

func (i ObjectVipMap) ToObjectVipMapOutput() ObjectVipMapOutput {
	return i.ToObjectVipMapOutputWithContext(context.Background())
}

func (i ObjectVipMap) ToObjectVipMapOutputWithContext(ctx context.Context) ObjectVipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectVipMapOutput)
}

type ObjectVipOutput struct{ *pulumi.OutputState }

func (ObjectVipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectVip)(nil)).Elem()
}

func (o ObjectVipOutput) ToObjectVipOutput() ObjectVipOutput {
	return o
}

func (o ObjectVipOutput) ToObjectVipOutputWithContext(ctx context.Context) ObjectVipOutput {
	return o
}

// Comment.
func (o ObjectVipOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectVip) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
func (o ObjectVipOutput) Extintf() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectVip) pulumi.StringOutput { return v.Extintf }).(pulumi.StringOutput)
}

// IP address or address range on the external interface that you want to map to an address or address range on the
// destination network.
func (o ObjectVipOutput) Extip() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectVip) pulumi.StringOutput { return v.Extip }).(pulumi.StringOutput)
}

// Incoming port number range that you want to map to a port number range on the destination network.
func (o ObjectVipOutput) Extport() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectVip) pulumi.StringOutput { return v.Extport }).(pulumi.StringOutput)
}

// IP address or address range on the destination network to which the external IP address is mapped.
func (o ObjectVipOutput) Mappedips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ObjectVip) pulumi.StringArrayOutput { return v.Mappedips }).(pulumi.StringArrayOutput)
}

// Port number range on the destination network to which the external port number range is mapped.
func (o ObjectVipOutput) Mappedport() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectVip) pulumi.StringOutput { return v.Mappedport }).(pulumi.StringOutput)
}

// Virtual IP name.
func (o ObjectVipOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectVip) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable port forwarding.
func (o ObjectVipOutput) Portforward() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectVip) pulumi.StringOutput { return v.Portforward }).(pulumi.StringOutput)
}

// Protocol to use when forwarding packets.
func (o ObjectVipOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectVip) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

type ObjectVipArrayOutput struct{ *pulumi.OutputState }

func (ObjectVipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectVip)(nil)).Elem()
}

func (o ObjectVipArrayOutput) ToObjectVipArrayOutput() ObjectVipArrayOutput {
	return o
}

func (o ObjectVipArrayOutput) ToObjectVipArrayOutputWithContext(ctx context.Context) ObjectVipArrayOutput {
	return o
}

func (o ObjectVipArrayOutput) Index(i pulumi.IntInput) ObjectVipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectVip {
		return vs[0].([]*ObjectVip)[vs[1].(int)]
	}).(ObjectVipOutput)
}

type ObjectVipMapOutput struct{ *pulumi.OutputState }

func (ObjectVipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectVip)(nil)).Elem()
}

func (o ObjectVipMapOutput) ToObjectVipMapOutput() ObjectVipMapOutput {
	return o
}

func (o ObjectVipMapOutput) ToObjectVipMapOutputWithContext(ctx context.Context) ObjectVipMapOutput {
	return o
}

func (o ObjectVipMapOutput) MapIndex(k pulumi.StringInput) ObjectVipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectVip {
		return vs[0].(map[string]*ObjectVip)[vs[1].(string)]
	}).(ObjectVipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectVipInput)(nil)).Elem(), &ObjectVip{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectVipArrayInput)(nil)).Elem(), ObjectVipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectVipMapInput)(nil)).Elem(), ObjectVipMap{})
	pulumi.RegisterOutputType(ObjectVipOutput{})
	pulumi.RegisterOutputType(ObjectVipArrayOutput{})
	pulumi.RegisterOutputType(ObjectVipMapOutput{})
}
