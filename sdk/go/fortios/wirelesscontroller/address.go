// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure the client with its MAC address. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.0.7,7.0.8,7.0.9,7.0.10,7.0.11,7.0.12,7.0.13`.
//
// ## Import
//
// WirelessController Address can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/address:Address labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:wirelesscontroller/address:Address labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Address struct {
	pulumi.CustomResourceState

	// ID.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// MAC address.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewAddress registers a new resource with the given unique name, arguments, and options.
func NewAddress(ctx *pulumi.Context,
	name string, args *AddressArgs, opts ...pulumi.ResourceOption) (*Address, error) {
	if args == nil {
		args = &AddressArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Address
	err := ctx.RegisterResource("fortios:wirelesscontroller/address:Address", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddress gets an existing Address resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddressState, opts ...pulumi.ResourceOption) (*Address, error) {
	var resource Address
	err := ctx.ReadResource("fortios:wirelesscontroller/address:Address", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Address resources.
type addressState struct {
	// ID.
	Fosid *string `pulumi:"fosid"`
	// MAC address.
	Mac *string `pulumi:"mac"`
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy *string `pulumi:"policy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type AddressState struct {
	// ID.
	Fosid pulumi.StringPtrInput
	// MAC address.
	Mac pulumi.StringPtrInput
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*addressState)(nil)).Elem()
}

type addressArgs struct {
	// ID.
	Fosid *string `pulumi:"fosid"`
	// MAC address.
	Mac *string `pulumi:"mac"`
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy *string `pulumi:"policy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Address resource.
type AddressArgs struct {
	// ID.
	Fosid pulumi.StringPtrInput
	// MAC address.
	Mac pulumi.StringPtrInput
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addressArgs)(nil)).Elem()
}

type AddressInput interface {
	pulumi.Input

	ToAddressOutput() AddressOutput
	ToAddressOutputWithContext(ctx context.Context) AddressOutput
}

func (*Address) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (i *Address) ToAddressOutput() AddressOutput {
	return i.ToAddressOutputWithContext(context.Background())
}

func (i *Address) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput)
}

// AddressArrayInput is an input type that accepts AddressArray and AddressArrayOutput values.
// You can construct a concrete instance of `AddressArrayInput` via:
//
//	AddressArray{ AddressArgs{...} }
type AddressArrayInput interface {
	pulumi.Input

	ToAddressArrayOutput() AddressArrayOutput
	ToAddressArrayOutputWithContext(context.Context) AddressArrayOutput
}

type AddressArray []AddressInput

func (AddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Address)(nil)).Elem()
}

func (i AddressArray) ToAddressArrayOutput() AddressArrayOutput {
	return i.ToAddressArrayOutputWithContext(context.Background())
}

func (i AddressArray) ToAddressArrayOutputWithContext(ctx context.Context) AddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressArrayOutput)
}

// AddressMapInput is an input type that accepts AddressMap and AddressMapOutput values.
// You can construct a concrete instance of `AddressMapInput` via:
//
//	AddressMap{ "key": AddressArgs{...} }
type AddressMapInput interface {
	pulumi.Input

	ToAddressMapOutput() AddressMapOutput
	ToAddressMapOutputWithContext(context.Context) AddressMapOutput
}

type AddressMap map[string]AddressInput

func (AddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Address)(nil)).Elem()
}

func (i AddressMap) ToAddressMapOutput() AddressMapOutput {
	return i.ToAddressMapOutputWithContext(context.Background())
}

func (i AddressMap) ToAddressMapOutputWithContext(ctx context.Context) AddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressMapOutput)
}

type AddressOutput struct{ *pulumi.OutputState }

func (AddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (o AddressOutput) ToAddressOutput() AddressOutput {
	return o
}

func (o AddressOutput) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return o
}

// ID.
func (o AddressOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

// MAC address.
func (o AddressOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Mac }).(pulumi.StringOutput)
}

// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
func (o AddressOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AddressOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AddressArrayOutput struct{ *pulumi.OutputState }

func (AddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Address)(nil)).Elem()
}

func (o AddressArrayOutput) ToAddressArrayOutput() AddressArrayOutput {
	return o
}

func (o AddressArrayOutput) ToAddressArrayOutputWithContext(ctx context.Context) AddressArrayOutput {
	return o
}

func (o AddressArrayOutput) Index(i pulumi.IntInput) AddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Address {
		return vs[0].([]*Address)[vs[1].(int)]
	}).(AddressOutput)
}

type AddressMapOutput struct{ *pulumi.OutputState }

func (AddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Address)(nil)).Elem()
}

func (o AddressMapOutput) ToAddressMapOutput() AddressMapOutput {
	return o
}

func (o AddressMapOutput) ToAddressMapOutputWithContext(ctx context.Context) AddressMapOutput {
	return o
}

func (o AddressMapOutput) MapIndex(k pulumi.StringInput) AddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Address {
		return vs[0].(map[string]*Address)[vs[1].(string)]
	}).(AddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddressInput)(nil)).Elem(), &Address{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressArrayInput)(nil)).Elem(), AddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressMapInput)(nil)).Elem(), AddressMap{})
	pulumi.RegisterOutputType(AddressOutput{})
	pulumi.RegisterOutputType(AddressArrayOutput{})
	pulumi.RegisterOutputType(AddressMapOutput{})
}
