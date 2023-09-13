// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure DNS translation.
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
//			_, err := firewall.NewDnstranslation(ctx, "trname", &firewall.DnstranslationArgs{
//				Dst:     pulumi.String("2.2.2.2"),
//				Fosid:   pulumi.Int(1),
//				Netmask: pulumi.String("255.0.0.0"),
//				Src:     pulumi.String("1.1.1.1"),
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
// # Firewall Dnstranslation can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:firewall/dnstranslation:Dnstranslation labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:firewall/dnstranslation:Dnstranslation labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Dnstranslation struct {
	pulumi.CustomResourceState

	// IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
	Dst pulumi.StringOutput `pulumi:"dst"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
	Netmask pulumi.StringOutput `pulumi:"netmask"`
	// IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
	Src pulumi.StringOutput `pulumi:"src"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDnstranslation registers a new resource with the given unique name, arguments, and options.
func NewDnstranslation(ctx *pulumi.Context,
	name string, args *DnstranslationArgs, opts ...pulumi.ResourceOption) (*Dnstranslation, error) {
	if args == nil {
		args = &DnstranslationArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Dnstranslation
	err := ctx.RegisterResource("fortios:firewall/dnstranslation:Dnstranslation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnstranslation gets an existing Dnstranslation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnstranslation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnstranslationState, opts ...pulumi.ResourceOption) (*Dnstranslation, error) {
	var resource Dnstranslation
	err := ctx.ReadResource("fortios:firewall/dnstranslation:Dnstranslation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dnstranslation resources.
type dnstranslationState struct {
	// IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
	Dst *string `pulumi:"dst"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
	Netmask *string `pulumi:"netmask"`
	// IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
	Src *string `pulumi:"src"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DnstranslationState struct {
	// IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
	Dst pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntPtrInput
	// If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
	Netmask pulumi.StringPtrInput
	// IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
	Src pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DnstranslationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnstranslationState)(nil)).Elem()
}

type dnstranslationArgs struct {
	// IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
	Dst *string `pulumi:"dst"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
	Netmask *string `pulumi:"netmask"`
	// IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
	Src *string `pulumi:"src"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Dnstranslation resource.
type DnstranslationArgs struct {
	// IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
	Dst pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntPtrInput
	// If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
	Netmask pulumi.StringPtrInput
	// IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
	Src pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DnstranslationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnstranslationArgs)(nil)).Elem()
}

type DnstranslationInput interface {
	pulumi.Input

	ToDnstranslationOutput() DnstranslationOutput
	ToDnstranslationOutputWithContext(ctx context.Context) DnstranslationOutput
}

func (*Dnstranslation) ElementType() reflect.Type {
	return reflect.TypeOf((**Dnstranslation)(nil)).Elem()
}

func (i *Dnstranslation) ToDnstranslationOutput() DnstranslationOutput {
	return i.ToDnstranslationOutputWithContext(context.Background())
}

func (i *Dnstranslation) ToDnstranslationOutputWithContext(ctx context.Context) DnstranslationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnstranslationOutput)
}

// DnstranslationArrayInput is an input type that accepts DnstranslationArray and DnstranslationArrayOutput values.
// You can construct a concrete instance of `DnstranslationArrayInput` via:
//
//	DnstranslationArray{ DnstranslationArgs{...} }
type DnstranslationArrayInput interface {
	pulumi.Input

	ToDnstranslationArrayOutput() DnstranslationArrayOutput
	ToDnstranslationArrayOutputWithContext(context.Context) DnstranslationArrayOutput
}

type DnstranslationArray []DnstranslationInput

func (DnstranslationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dnstranslation)(nil)).Elem()
}

func (i DnstranslationArray) ToDnstranslationArrayOutput() DnstranslationArrayOutput {
	return i.ToDnstranslationArrayOutputWithContext(context.Background())
}

func (i DnstranslationArray) ToDnstranslationArrayOutputWithContext(ctx context.Context) DnstranslationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnstranslationArrayOutput)
}

// DnstranslationMapInput is an input type that accepts DnstranslationMap and DnstranslationMapOutput values.
// You can construct a concrete instance of `DnstranslationMapInput` via:
//
//	DnstranslationMap{ "key": DnstranslationArgs{...} }
type DnstranslationMapInput interface {
	pulumi.Input

	ToDnstranslationMapOutput() DnstranslationMapOutput
	ToDnstranslationMapOutputWithContext(context.Context) DnstranslationMapOutput
}

type DnstranslationMap map[string]DnstranslationInput

func (DnstranslationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dnstranslation)(nil)).Elem()
}

func (i DnstranslationMap) ToDnstranslationMapOutput() DnstranslationMapOutput {
	return i.ToDnstranslationMapOutputWithContext(context.Background())
}

func (i DnstranslationMap) ToDnstranslationMapOutputWithContext(ctx context.Context) DnstranslationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnstranslationMapOutput)
}

type DnstranslationOutput struct{ *pulumi.OutputState }

func (DnstranslationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dnstranslation)(nil)).Elem()
}

func (o DnstranslationOutput) ToDnstranslationOutput() DnstranslationOutput {
	return o
}

func (o DnstranslationOutput) ToDnstranslationOutputWithContext(ctx context.Context) DnstranslationOutput {
	return o
}

// IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
func (o DnstranslationOutput) Dst() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnstranslation) pulumi.StringOutput { return v.Dst }).(pulumi.StringOutput)
}

// ID.
func (o DnstranslationOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Dnstranslation) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
func (o DnstranslationOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnstranslation) pulumi.StringOutput { return v.Netmask }).(pulumi.StringOutput)
}

// IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
func (o DnstranslationOutput) Src() pulumi.StringOutput {
	return o.ApplyT(func(v *Dnstranslation) pulumi.StringOutput { return v.Src }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DnstranslationOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dnstranslation) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DnstranslationArrayOutput struct{ *pulumi.OutputState }

func (DnstranslationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dnstranslation)(nil)).Elem()
}

func (o DnstranslationArrayOutput) ToDnstranslationArrayOutput() DnstranslationArrayOutput {
	return o
}

func (o DnstranslationArrayOutput) ToDnstranslationArrayOutputWithContext(ctx context.Context) DnstranslationArrayOutput {
	return o
}

func (o DnstranslationArrayOutput) Index(i pulumi.IntInput) DnstranslationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dnstranslation {
		return vs[0].([]*Dnstranslation)[vs[1].(int)]
	}).(DnstranslationOutput)
}

type DnstranslationMapOutput struct{ *pulumi.OutputState }

func (DnstranslationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dnstranslation)(nil)).Elem()
}

func (o DnstranslationMapOutput) ToDnstranslationMapOutput() DnstranslationMapOutput {
	return o
}

func (o DnstranslationMapOutput) ToDnstranslationMapOutputWithContext(ctx context.Context) DnstranslationMapOutput {
	return o
}

func (o DnstranslationMapOutput) MapIndex(k pulumi.StringInput) DnstranslationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dnstranslation {
		return vs[0].(map[string]*Dnstranslation)[vs[1].(string)]
	}).(DnstranslationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnstranslationInput)(nil)).Elem(), &Dnstranslation{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnstranslationArrayInput)(nil)).Elem(), DnstranslationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnstranslationMapInput)(nil)).Elem(), DnstranslationMap{})
	pulumi.RegisterOutputType(DnstranslationOutput{})
	pulumi.RegisterOutputType(DnstranslationArrayOutput{})
	pulumi.RegisterOutputType(DnstranslationMapOutput{})
}
