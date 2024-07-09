// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPS URL filter DNS servers.
//
// ## Import
//
// System IpsUrlfilterDns can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/ipsurlfilterdns:Ipsurlfilterdns labelname {{address}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/ipsurlfilterdns:Ipsurlfilterdns labelname {{address}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Ipsurlfilterdns struct {
	pulumi.CustomResourceState

	// DNS server IP address.
	Address pulumi.StringOutput `pulumi:"address"`
	// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
	Ipv6Capability pulumi.StringOutput `pulumi:"ipv6Capability"`
	// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewIpsurlfilterdns registers a new resource with the given unique name, arguments, and options.
func NewIpsurlfilterdns(ctx *pulumi.Context,
	name string, args *IpsurlfilterdnsArgs, opts ...pulumi.ResourceOption) (*Ipsurlfilterdns, error) {
	if args == nil {
		args = &IpsurlfilterdnsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ipsurlfilterdns
	err := ctx.RegisterResource("fortios:system/ipsurlfilterdns:Ipsurlfilterdns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsurlfilterdns gets an existing Ipsurlfilterdns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsurlfilterdns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsurlfilterdnsState, opts ...pulumi.ResourceOption) (*Ipsurlfilterdns, error) {
	var resource Ipsurlfilterdns
	err := ctx.ReadResource("fortios:system/ipsurlfilterdns:Ipsurlfilterdns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipsurlfilterdns resources.
type ipsurlfilterdnsState struct {
	// DNS server IP address.
	Address *string `pulumi:"address"`
	// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
	Ipv6Capability *string `pulumi:"ipv6Capability"`
	// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IpsurlfilterdnsState struct {
	// DNS server IP address.
	Address pulumi.StringPtrInput
	// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
	Ipv6Capability pulumi.StringPtrInput
	// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsurlfilterdnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsurlfilterdnsState)(nil)).Elem()
}

type ipsurlfilterdnsArgs struct {
	// DNS server IP address.
	Address *string `pulumi:"address"`
	// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
	Ipv6Capability *string `pulumi:"ipv6Capability"`
	// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Ipsurlfilterdns resource.
type IpsurlfilterdnsArgs struct {
	// DNS server IP address.
	Address pulumi.StringPtrInput
	// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
	Ipv6Capability pulumi.StringPtrInput
	// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsurlfilterdnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsurlfilterdnsArgs)(nil)).Elem()
}

type IpsurlfilterdnsInput interface {
	pulumi.Input

	ToIpsurlfilterdnsOutput() IpsurlfilterdnsOutput
	ToIpsurlfilterdnsOutputWithContext(ctx context.Context) IpsurlfilterdnsOutput
}

func (*Ipsurlfilterdns) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipsurlfilterdns)(nil)).Elem()
}

func (i *Ipsurlfilterdns) ToIpsurlfilterdnsOutput() IpsurlfilterdnsOutput {
	return i.ToIpsurlfilterdnsOutputWithContext(context.Background())
}

func (i *Ipsurlfilterdns) ToIpsurlfilterdnsOutputWithContext(ctx context.Context) IpsurlfilterdnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsurlfilterdnsOutput)
}

// IpsurlfilterdnsArrayInput is an input type that accepts IpsurlfilterdnsArray and IpsurlfilterdnsArrayOutput values.
// You can construct a concrete instance of `IpsurlfilterdnsArrayInput` via:
//
//	IpsurlfilterdnsArray{ IpsurlfilterdnsArgs{...} }
type IpsurlfilterdnsArrayInput interface {
	pulumi.Input

	ToIpsurlfilterdnsArrayOutput() IpsurlfilterdnsArrayOutput
	ToIpsurlfilterdnsArrayOutputWithContext(context.Context) IpsurlfilterdnsArrayOutput
}

type IpsurlfilterdnsArray []IpsurlfilterdnsInput

func (IpsurlfilterdnsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipsurlfilterdns)(nil)).Elem()
}

func (i IpsurlfilterdnsArray) ToIpsurlfilterdnsArrayOutput() IpsurlfilterdnsArrayOutput {
	return i.ToIpsurlfilterdnsArrayOutputWithContext(context.Background())
}

func (i IpsurlfilterdnsArray) ToIpsurlfilterdnsArrayOutputWithContext(ctx context.Context) IpsurlfilterdnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsurlfilterdnsArrayOutput)
}

// IpsurlfilterdnsMapInput is an input type that accepts IpsurlfilterdnsMap and IpsurlfilterdnsMapOutput values.
// You can construct a concrete instance of `IpsurlfilterdnsMapInput` via:
//
//	IpsurlfilterdnsMap{ "key": IpsurlfilterdnsArgs{...} }
type IpsurlfilterdnsMapInput interface {
	pulumi.Input

	ToIpsurlfilterdnsMapOutput() IpsurlfilterdnsMapOutput
	ToIpsurlfilterdnsMapOutputWithContext(context.Context) IpsurlfilterdnsMapOutput
}

type IpsurlfilterdnsMap map[string]IpsurlfilterdnsInput

func (IpsurlfilterdnsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipsurlfilterdns)(nil)).Elem()
}

func (i IpsurlfilterdnsMap) ToIpsurlfilterdnsMapOutput() IpsurlfilterdnsMapOutput {
	return i.ToIpsurlfilterdnsMapOutputWithContext(context.Background())
}

func (i IpsurlfilterdnsMap) ToIpsurlfilterdnsMapOutputWithContext(ctx context.Context) IpsurlfilterdnsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsurlfilterdnsMapOutput)
}

type IpsurlfilterdnsOutput struct{ *pulumi.OutputState }

func (IpsurlfilterdnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipsurlfilterdns)(nil)).Elem()
}

func (o IpsurlfilterdnsOutput) ToIpsurlfilterdnsOutput() IpsurlfilterdnsOutput {
	return o
}

func (o IpsurlfilterdnsOutput) ToIpsurlfilterdnsOutputWithContext(ctx context.Context) IpsurlfilterdnsOutput {
	return o
}

// DNS server IP address.
func (o IpsurlfilterdnsOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipsurlfilterdns) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
func (o IpsurlfilterdnsOutput) Ipv6Capability() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipsurlfilterdns) pulumi.StringOutput { return v.Ipv6Capability }).(pulumi.StringOutput)
}

// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
func (o IpsurlfilterdnsOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipsurlfilterdns) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o IpsurlfilterdnsOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipsurlfilterdns) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type IpsurlfilterdnsArrayOutput struct{ *pulumi.OutputState }

func (IpsurlfilterdnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipsurlfilterdns)(nil)).Elem()
}

func (o IpsurlfilterdnsArrayOutput) ToIpsurlfilterdnsArrayOutput() IpsurlfilterdnsArrayOutput {
	return o
}

func (o IpsurlfilterdnsArrayOutput) ToIpsurlfilterdnsArrayOutputWithContext(ctx context.Context) IpsurlfilterdnsArrayOutput {
	return o
}

func (o IpsurlfilterdnsArrayOutput) Index(i pulumi.IntInput) IpsurlfilterdnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ipsurlfilterdns {
		return vs[0].([]*Ipsurlfilterdns)[vs[1].(int)]
	}).(IpsurlfilterdnsOutput)
}

type IpsurlfilterdnsMapOutput struct{ *pulumi.OutputState }

func (IpsurlfilterdnsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipsurlfilterdns)(nil)).Elem()
}

func (o IpsurlfilterdnsMapOutput) ToIpsurlfilterdnsMapOutput() IpsurlfilterdnsMapOutput {
	return o
}

func (o IpsurlfilterdnsMapOutput) ToIpsurlfilterdnsMapOutputWithContext(ctx context.Context) IpsurlfilterdnsMapOutput {
	return o
}

func (o IpsurlfilterdnsMapOutput) MapIndex(k pulumi.StringInput) IpsurlfilterdnsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ipsurlfilterdns {
		return vs[0].(map[string]*Ipsurlfilterdns)[vs[1].(string)]
	}).(IpsurlfilterdnsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpsurlfilterdnsInput)(nil)).Elem(), &Ipsurlfilterdns{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsurlfilterdnsArrayInput)(nil)).Elem(), IpsurlfilterdnsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsurlfilterdnsMapInput)(nil)).Elem(), IpsurlfilterdnsMap{})
	pulumi.RegisterOutputType(IpsurlfilterdnsOutput{})
	pulumi.RegisterOutputType(IpsurlfilterdnsArrayOutput{})
	pulumi.RegisterOutputType(IpsurlfilterdnsMapOutput{})
}
