// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to configure firewall addresses used in firewall policies of FortiOS.
//
// !> **Warning:** The resource will be deprecated and replaced by new resource `firewall.Address`, we recommend that you use the new resource.
//
// ## Example Usage
// ### Iprange Address
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
//			_, err := firewall.NewObjectAddress(ctx, "s1", &firewall.ObjectAddressArgs{
//				Comment: pulumi.String("dd"),
//				EndIp:   pulumi.String("2.0.0.0"),
//				StartIp: pulumi.String("1.0.0.0"),
//				Type:    pulumi.String("iprange"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Geography Address
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
//			_, err := firewall.NewObjectAddress(ctx, "s2", &firewall.ObjectAddressArgs{
//				Comment: pulumi.String("dd"),
//				Country: pulumi.String("AO"),
//				Type:    pulumi.String("geography"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Fqdn Address
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
//			_, err := firewall.NewObjectAddress(ctx, "s3", &firewall.ObjectAddressArgs{
//				AssociatedInterface:  pulumi.String("port4"),
//				Comment:              pulumi.String("dd"),
//				Fqdn:                 pulumi.String("baid.com"),
//				ShowInAddressList:    pulumi.String("disable"),
//				StaticRouteConfigure: pulumi.String("enable"),
//				Type:                 pulumi.String("fqdn"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Ipmask Address
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
//			_, err := firewall.NewObjectAddress(ctx, "s4", &firewall.ObjectAddressArgs{
//				Comment: pulumi.String("dd"),
//				Subnet:  pulumi.String("0.0.0.0 0.0.0.0"),
//				Type:    pulumi.String("ipmask"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ObjectAddress struct {
	pulumi.CustomResourceState

	// Network interface associated with address.
	AssociatedInterface pulumi.StringOutput `pulumi:"associatedInterface"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// IP addresses associated to a specific country.
	Country pulumi.StringOutput `pulumi:"country"`
	// Final IP address (inclusive) in the range for the address.
	EndIp pulumi.StringOutput `pulumi:"endIp"`
	// Fully Qualified Domain Name address.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Address name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable address visibility in the GUI. default is "enable".
	ShowInAddressList pulumi.StringOutput `pulumi:"showInAddressList"`
	// First IP address (inclusive) in the range for the address.
	StartIp pulumi.StringOutput `pulumi:"startIp"`
	// Enable/disable use of this address in the static route configuration. default is "disable".
	StaticRouteConfigure pulumi.StringOutput `pulumi:"staticRouteConfigure"`
	// IP address and subnet mask of address.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// Type of address(Support ipmask, iprange, fqdn and geography).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewObjectAddress registers a new resource with the given unique name, arguments, and options.
func NewObjectAddress(ctx *pulumi.Context,
	name string, args *ObjectAddressArgs, opts ...pulumi.ResourceOption) (*ObjectAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ObjectAddress
	err := ctx.RegisterResource("fortios:firewall/objectAddress:ObjectAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectAddress gets an existing ObjectAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectAddressState, opts ...pulumi.ResourceOption) (*ObjectAddress, error) {
	var resource ObjectAddress
	err := ctx.ReadResource("fortios:firewall/objectAddress:ObjectAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectAddress resources.
type objectAddressState struct {
	// Network interface associated with address.
	AssociatedInterface *string `pulumi:"associatedInterface"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// IP addresses associated to a specific country.
	Country *string `pulumi:"country"`
	// Final IP address (inclusive) in the range for the address.
	EndIp *string `pulumi:"endIp"`
	// Fully Qualified Domain Name address.
	Fqdn *string `pulumi:"fqdn"`
	// Address name.
	Name *string `pulumi:"name"`
	// Enable/disable address visibility in the GUI. default is "enable".
	ShowInAddressList *string `pulumi:"showInAddressList"`
	// First IP address (inclusive) in the range for the address.
	StartIp *string `pulumi:"startIp"`
	// Enable/disable use of this address in the static route configuration. default is "disable".
	StaticRouteConfigure *string `pulumi:"staticRouteConfigure"`
	// IP address and subnet mask of address.
	Subnet *string `pulumi:"subnet"`
	// Type of address(Support ipmask, iprange, fqdn and geography).
	Type *string `pulumi:"type"`
}

type ObjectAddressState struct {
	// Network interface associated with address.
	AssociatedInterface pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// IP addresses associated to a specific country.
	Country pulumi.StringPtrInput
	// Final IP address (inclusive) in the range for the address.
	EndIp pulumi.StringPtrInput
	// Fully Qualified Domain Name address.
	Fqdn pulumi.StringPtrInput
	// Address name.
	Name pulumi.StringPtrInput
	// Enable/disable address visibility in the GUI. default is "enable".
	ShowInAddressList pulumi.StringPtrInput
	// First IP address (inclusive) in the range for the address.
	StartIp pulumi.StringPtrInput
	// Enable/disable use of this address in the static route configuration. default is "disable".
	StaticRouteConfigure pulumi.StringPtrInput
	// IP address and subnet mask of address.
	Subnet pulumi.StringPtrInput
	// Type of address(Support ipmask, iprange, fqdn and geography).
	Type pulumi.StringPtrInput
}

func (ObjectAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAddressState)(nil)).Elem()
}

type objectAddressArgs struct {
	// Network interface associated with address.
	AssociatedInterface *string `pulumi:"associatedInterface"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// IP addresses associated to a specific country.
	Country *string `pulumi:"country"`
	// Final IP address (inclusive) in the range for the address.
	EndIp *string `pulumi:"endIp"`
	// Fully Qualified Domain Name address.
	Fqdn *string `pulumi:"fqdn"`
	// Address name.
	Name *string `pulumi:"name"`
	// Enable/disable address visibility in the GUI. default is "enable".
	ShowInAddressList *string `pulumi:"showInAddressList"`
	// First IP address (inclusive) in the range for the address.
	StartIp *string `pulumi:"startIp"`
	// Enable/disable use of this address in the static route configuration. default is "disable".
	StaticRouteConfigure *string `pulumi:"staticRouteConfigure"`
	// IP address and subnet mask of address.
	Subnet *string `pulumi:"subnet"`
	// Type of address(Support ipmask, iprange, fqdn and geography).
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a ObjectAddress resource.
type ObjectAddressArgs struct {
	// Network interface associated with address.
	AssociatedInterface pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// IP addresses associated to a specific country.
	Country pulumi.StringPtrInput
	// Final IP address (inclusive) in the range for the address.
	EndIp pulumi.StringPtrInput
	// Fully Qualified Domain Name address.
	Fqdn pulumi.StringPtrInput
	// Address name.
	Name pulumi.StringPtrInput
	// Enable/disable address visibility in the GUI. default is "enable".
	ShowInAddressList pulumi.StringPtrInput
	// First IP address (inclusive) in the range for the address.
	StartIp pulumi.StringPtrInput
	// Enable/disable use of this address in the static route configuration. default is "disable".
	StaticRouteConfigure pulumi.StringPtrInput
	// IP address and subnet mask of address.
	Subnet pulumi.StringPtrInput
	// Type of address(Support ipmask, iprange, fqdn and geography).
	Type pulumi.StringInput
}

func (ObjectAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAddressArgs)(nil)).Elem()
}

type ObjectAddressInput interface {
	pulumi.Input

	ToObjectAddressOutput() ObjectAddressOutput
	ToObjectAddressOutputWithContext(ctx context.Context) ObjectAddressOutput
}

func (*ObjectAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectAddress)(nil)).Elem()
}

func (i *ObjectAddress) ToObjectAddressOutput() ObjectAddressOutput {
	return i.ToObjectAddressOutputWithContext(context.Background())
}

func (i *ObjectAddress) ToObjectAddressOutputWithContext(ctx context.Context) ObjectAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectAddressOutput)
}

// ObjectAddressArrayInput is an input type that accepts ObjectAddressArray and ObjectAddressArrayOutput values.
// You can construct a concrete instance of `ObjectAddressArrayInput` via:
//
//	ObjectAddressArray{ ObjectAddressArgs{...} }
type ObjectAddressArrayInput interface {
	pulumi.Input

	ToObjectAddressArrayOutput() ObjectAddressArrayOutput
	ToObjectAddressArrayOutputWithContext(context.Context) ObjectAddressArrayOutput
}

type ObjectAddressArray []ObjectAddressInput

func (ObjectAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectAddress)(nil)).Elem()
}

func (i ObjectAddressArray) ToObjectAddressArrayOutput() ObjectAddressArrayOutput {
	return i.ToObjectAddressArrayOutputWithContext(context.Background())
}

func (i ObjectAddressArray) ToObjectAddressArrayOutputWithContext(ctx context.Context) ObjectAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectAddressArrayOutput)
}

// ObjectAddressMapInput is an input type that accepts ObjectAddressMap and ObjectAddressMapOutput values.
// You can construct a concrete instance of `ObjectAddressMapInput` via:
//
//	ObjectAddressMap{ "key": ObjectAddressArgs{...} }
type ObjectAddressMapInput interface {
	pulumi.Input

	ToObjectAddressMapOutput() ObjectAddressMapOutput
	ToObjectAddressMapOutputWithContext(context.Context) ObjectAddressMapOutput
}

type ObjectAddressMap map[string]ObjectAddressInput

func (ObjectAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectAddress)(nil)).Elem()
}

func (i ObjectAddressMap) ToObjectAddressMapOutput() ObjectAddressMapOutput {
	return i.ToObjectAddressMapOutputWithContext(context.Background())
}

func (i ObjectAddressMap) ToObjectAddressMapOutputWithContext(ctx context.Context) ObjectAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectAddressMapOutput)
}

type ObjectAddressOutput struct{ *pulumi.OutputState }

func (ObjectAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectAddress)(nil)).Elem()
}

func (o ObjectAddressOutput) ToObjectAddressOutput() ObjectAddressOutput {
	return o
}

func (o ObjectAddressOutput) ToObjectAddressOutputWithContext(ctx context.Context) ObjectAddressOutput {
	return o
}

// Network interface associated with address.
func (o ObjectAddressOutput) AssociatedInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAddress) pulumi.StringOutput { return v.AssociatedInterface }).(pulumi.StringOutput)
}

// Comment.
func (o ObjectAddressOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectAddress) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// IP addresses associated to a specific country.
func (o ObjectAddressOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAddress) pulumi.StringOutput { return v.Country }).(pulumi.StringOutput)
}

// Final IP address (inclusive) in the range for the address.
func (o ObjectAddressOutput) EndIp() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAddress) pulumi.StringOutput { return v.EndIp }).(pulumi.StringOutput)
}

// Fully Qualified Domain Name address.
func (o ObjectAddressOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAddress) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// Address name.
func (o ObjectAddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAddress) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable address visibility in the GUI. default is "enable".
func (o ObjectAddressOutput) ShowInAddressList() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAddress) pulumi.StringOutput { return v.ShowInAddressList }).(pulumi.StringOutput)
}

// First IP address (inclusive) in the range for the address.
func (o ObjectAddressOutput) StartIp() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAddress) pulumi.StringOutput { return v.StartIp }).(pulumi.StringOutput)
}

// Enable/disable use of this address in the static route configuration. default is "disable".
func (o ObjectAddressOutput) StaticRouteConfigure() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAddress) pulumi.StringOutput { return v.StaticRouteConfigure }).(pulumi.StringOutput)
}

// IP address and subnet mask of address.
func (o ObjectAddressOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAddress) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// Type of address(Support ipmask, iprange, fqdn and geography).
func (o ObjectAddressOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAddress) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ObjectAddressArrayOutput struct{ *pulumi.OutputState }

func (ObjectAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectAddress)(nil)).Elem()
}

func (o ObjectAddressArrayOutput) ToObjectAddressArrayOutput() ObjectAddressArrayOutput {
	return o
}

func (o ObjectAddressArrayOutput) ToObjectAddressArrayOutputWithContext(ctx context.Context) ObjectAddressArrayOutput {
	return o
}

func (o ObjectAddressArrayOutput) Index(i pulumi.IntInput) ObjectAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectAddress {
		return vs[0].([]*ObjectAddress)[vs[1].(int)]
	}).(ObjectAddressOutput)
}

type ObjectAddressMapOutput struct{ *pulumi.OutputState }

func (ObjectAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectAddress)(nil)).Elem()
}

func (o ObjectAddressMapOutput) ToObjectAddressMapOutput() ObjectAddressMapOutput {
	return o
}

func (o ObjectAddressMapOutput) ToObjectAddressMapOutputWithContext(ctx context.Context) ObjectAddressMapOutput {
	return o
}

func (o ObjectAddressMapOutput) MapIndex(k pulumi.StringInput) ObjectAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectAddress {
		return vs[0].(map[string]*ObjectAddress)[vs[1].(string)]
	}).(ObjectAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectAddressInput)(nil)).Elem(), &ObjectAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectAddressArrayInput)(nil)).Elem(), ObjectAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectAddressMapInput)(nil)).Elem(), ObjectAddressMap{})
	pulumi.RegisterOutputType(ObjectAddressOutput{})
	pulumi.RegisterOutputType(ObjectAddressArrayOutput{})
	pulumi.RegisterOutputType(ObjectAddressMapOutput{})
}
