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

// Configure IPv4 IP pools.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := firewall.NewIppool(ctx, "trname", &firewall.IppoolArgs{
//				ArpReply:         pulumi.String("enable"),
//				BlockSize:        pulumi.Int(128),
//				Endip:            pulumi.String("1.0.0.20"),
//				NumBlocksPerUser: pulumi.Int(8),
//				PbaTimeout:       pulumi.Int(30),
//				PermitAnyHost:    pulumi.String("disable"),
//				SourceEndip:      pulumi.String("0.0.0.0"),
//				SourceStartip:    pulumi.String("0.0.0.0"),
//				Startip:          pulumi.String("1.0.0.0"),
//				Type:             pulumi.String("overload"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Firewall Ippool can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/ippool:Ippool labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/ippool:Ippool labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Ippool struct {
	pulumi.CustomResourceState

	// Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.
	AddNat64Route pulumi.StringOutput `pulumi:"addNat64Route"`
	// Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
	ArpIntf pulumi.StringOutput `pulumi:"arpIntf"`
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
	ArpReply pulumi.StringOutput `pulumi:"arpReply"`
	// Associated interface name.
	AssociatedInterface pulumi.StringOutput `pulumi:"associatedInterface"`
	// Number of addresses in a block (64 - 4096, default = 128).
	BlockSize pulumi.IntOutput `pulumi:"blockSize"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Endip pulumi.StringOutput `pulumi:"endip"`
	// Final port number (inclusive) in the range for the address pool (Default: 65533).
	Endport pulumi.IntOutput `pulumi:"endport"`
	// IP pool name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable NAT64. Valid values: `disable`, `enable`.
	Nat64 pulumi.StringOutput `pulumi:"nat64"`
	// Number of addresses blocks that can be used by a user (1 to 128, default = 8).
	NumBlocksPerUser pulumi.IntOutput `pulumi:"numBlocksPerUser"`
	// Port block allocation timeout (seconds).
	PbaTimeout pulumi.IntOutput `pulumi:"pbaTimeout"`
	// Enable/disable full cone NAT. Valid values: `disable`, `enable`.
	PermitAnyHost pulumi.StringOutput `pulumi:"permitAnyHost"`
	// Number of port for each user (32 - 60416, default = 0, which is auto).
	PortPerUser pulumi.IntOutput `pulumi:"portPerUser"`
	// Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceEndip pulumi.StringOutput `pulumi:"sourceEndip"`
	// First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceStartip pulumi.StringOutput `pulumi:"sourceStartip"`
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Startip pulumi.StringOutput `pulumi:"startip"`
	// First port number (inclusive) in the range for the address pool (Default: 5117).
	Startport pulumi.IntOutput `pulumi:"startport"`
	// Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.
	SubnetBroadcastInIppool pulumi.StringOutput `pulumi:"subnetBroadcastInIppool"`
	// IP pool type. On FortiOS versions 6.2.0-7.4.1: overload, one-to-one, fixed port range, or port block allocation. On FortiOS versions >= 7.4.2: overload, one-to-one, fixed-port-range, port-block-allocation, cgn-resource-allocation (hyperscale vdom only). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIppool registers a new resource with the given unique name, arguments, and options.
func NewIppool(ctx *pulumi.Context,
	name string, args *IppoolArgs, opts ...pulumi.ResourceOption) (*Ippool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endip == nil {
		return nil, errors.New("invalid value for required argument 'Endip'")
	}
	if args.Startip == nil {
		return nil, errors.New("invalid value for required argument 'Startip'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ippool
	err := ctx.RegisterResource("fortios:firewall/ippool:Ippool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIppool gets an existing Ippool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIppool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IppoolState, opts ...pulumi.ResourceOption) (*Ippool, error) {
	var resource Ippool
	err := ctx.ReadResource("fortios:firewall/ippool:Ippool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ippool resources.
type ippoolState struct {
	// Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.
	AddNat64Route *string `pulumi:"addNat64Route"`
	// Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
	ArpIntf *string `pulumi:"arpIntf"`
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
	ArpReply *string `pulumi:"arpReply"`
	// Associated interface name.
	AssociatedInterface *string `pulumi:"associatedInterface"`
	// Number of addresses in a block (64 - 4096, default = 128).
	BlockSize *int `pulumi:"blockSize"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Endip *string `pulumi:"endip"`
	// Final port number (inclusive) in the range for the address pool (Default: 65533).
	Endport *int `pulumi:"endport"`
	// IP pool name.
	Name *string `pulumi:"name"`
	// Enable/disable NAT64. Valid values: `disable`, `enable`.
	Nat64 *string `pulumi:"nat64"`
	// Number of addresses blocks that can be used by a user (1 to 128, default = 8).
	NumBlocksPerUser *int `pulumi:"numBlocksPerUser"`
	// Port block allocation timeout (seconds).
	PbaTimeout *int `pulumi:"pbaTimeout"`
	// Enable/disable full cone NAT. Valid values: `disable`, `enable`.
	PermitAnyHost *string `pulumi:"permitAnyHost"`
	// Number of port for each user (32 - 60416, default = 0, which is auto).
	PortPerUser *int `pulumi:"portPerUser"`
	// Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceEndip *string `pulumi:"sourceEndip"`
	// First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceStartip *string `pulumi:"sourceStartip"`
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Startip *string `pulumi:"startip"`
	// First port number (inclusive) in the range for the address pool (Default: 5117).
	Startport *int `pulumi:"startport"`
	// Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.
	SubnetBroadcastInIppool *string `pulumi:"subnetBroadcastInIppool"`
	// IP pool type. On FortiOS versions 6.2.0-7.4.1: overload, one-to-one, fixed port range, or port block allocation. On FortiOS versions >= 7.4.2: overload, one-to-one, fixed-port-range, port-block-allocation, cgn-resource-allocation (hyperscale vdom only). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IppoolState struct {
	// Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.
	AddNat64Route pulumi.StringPtrInput
	// Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
	ArpIntf pulumi.StringPtrInput
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
	ArpReply pulumi.StringPtrInput
	// Associated interface name.
	AssociatedInterface pulumi.StringPtrInput
	// Number of addresses in a block (64 - 4096, default = 128).
	BlockSize pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Endip pulumi.StringPtrInput
	// Final port number (inclusive) in the range for the address pool (Default: 65533).
	Endport pulumi.IntPtrInput
	// IP pool name.
	Name pulumi.StringPtrInput
	// Enable/disable NAT64. Valid values: `disable`, `enable`.
	Nat64 pulumi.StringPtrInput
	// Number of addresses blocks that can be used by a user (1 to 128, default = 8).
	NumBlocksPerUser pulumi.IntPtrInput
	// Port block allocation timeout (seconds).
	PbaTimeout pulumi.IntPtrInput
	// Enable/disable full cone NAT. Valid values: `disable`, `enable`.
	PermitAnyHost pulumi.StringPtrInput
	// Number of port for each user (32 - 60416, default = 0, which is auto).
	PortPerUser pulumi.IntPtrInput
	// Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceEndip pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceStartip pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Startip pulumi.StringPtrInput
	// First port number (inclusive) in the range for the address pool (Default: 5117).
	Startport pulumi.IntPtrInput
	// Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.
	SubnetBroadcastInIppool pulumi.StringPtrInput
	// IP pool type. On FortiOS versions 6.2.0-7.4.1: overload, one-to-one, fixed port range, or port block allocation. On FortiOS versions >= 7.4.2: overload, one-to-one, fixed-port-range, port-block-allocation, cgn-resource-allocation (hyperscale vdom only). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IppoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*ippoolState)(nil)).Elem()
}

type ippoolArgs struct {
	// Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.
	AddNat64Route *string `pulumi:"addNat64Route"`
	// Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
	ArpIntf *string `pulumi:"arpIntf"`
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
	ArpReply *string `pulumi:"arpReply"`
	// Associated interface name.
	AssociatedInterface *string `pulumi:"associatedInterface"`
	// Number of addresses in a block (64 - 4096, default = 128).
	BlockSize *int `pulumi:"blockSize"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Endip string `pulumi:"endip"`
	// Final port number (inclusive) in the range for the address pool (Default: 65533).
	Endport *int `pulumi:"endport"`
	// IP pool name.
	Name *string `pulumi:"name"`
	// Enable/disable NAT64. Valid values: `disable`, `enable`.
	Nat64 *string `pulumi:"nat64"`
	// Number of addresses blocks that can be used by a user (1 to 128, default = 8).
	NumBlocksPerUser *int `pulumi:"numBlocksPerUser"`
	// Port block allocation timeout (seconds).
	PbaTimeout *int `pulumi:"pbaTimeout"`
	// Enable/disable full cone NAT. Valid values: `disable`, `enable`.
	PermitAnyHost *string `pulumi:"permitAnyHost"`
	// Number of port for each user (32 - 60416, default = 0, which is auto).
	PortPerUser *int `pulumi:"portPerUser"`
	// Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceEndip *string `pulumi:"sourceEndip"`
	// First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceStartip *string `pulumi:"sourceStartip"`
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Startip string `pulumi:"startip"`
	// First port number (inclusive) in the range for the address pool (Default: 5117).
	Startport *int `pulumi:"startport"`
	// Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.
	SubnetBroadcastInIppool *string `pulumi:"subnetBroadcastInIppool"`
	// IP pool type. On FortiOS versions 6.2.0-7.4.1: overload, one-to-one, fixed port range, or port block allocation. On FortiOS versions >= 7.4.2: overload, one-to-one, fixed-port-range, port-block-allocation, cgn-resource-allocation (hyperscale vdom only). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Ippool resource.
type IppoolArgs struct {
	// Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.
	AddNat64Route pulumi.StringPtrInput
	// Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
	ArpIntf pulumi.StringPtrInput
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
	ArpReply pulumi.StringPtrInput
	// Associated interface name.
	AssociatedInterface pulumi.StringPtrInput
	// Number of addresses in a block (64 - 4096, default = 128).
	BlockSize pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Endip pulumi.StringInput
	// Final port number (inclusive) in the range for the address pool (Default: 65533).
	Endport pulumi.IntPtrInput
	// IP pool name.
	Name pulumi.StringPtrInput
	// Enable/disable NAT64. Valid values: `disable`, `enable`.
	Nat64 pulumi.StringPtrInput
	// Number of addresses blocks that can be used by a user (1 to 128, default = 8).
	NumBlocksPerUser pulumi.IntPtrInput
	// Port block allocation timeout (seconds).
	PbaTimeout pulumi.IntPtrInput
	// Enable/disable full cone NAT. Valid values: `disable`, `enable`.
	PermitAnyHost pulumi.StringPtrInput
	// Number of port for each user (32 - 60416, default = 0, which is auto).
	PortPerUser pulumi.IntPtrInput
	// Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceEndip pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceStartip pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Startip pulumi.StringInput
	// First port number (inclusive) in the range for the address pool (Default: 5117).
	Startport pulumi.IntPtrInput
	// Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.
	SubnetBroadcastInIppool pulumi.StringPtrInput
	// IP pool type. On FortiOS versions 6.2.0-7.4.1: overload, one-to-one, fixed port range, or port block allocation. On FortiOS versions >= 7.4.2: overload, one-to-one, fixed-port-range, port-block-allocation, cgn-resource-allocation (hyperscale vdom only). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IppoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ippoolArgs)(nil)).Elem()
}

type IppoolInput interface {
	pulumi.Input

	ToIppoolOutput() IppoolOutput
	ToIppoolOutputWithContext(ctx context.Context) IppoolOutput
}

func (*Ippool) ElementType() reflect.Type {
	return reflect.TypeOf((**Ippool)(nil)).Elem()
}

func (i *Ippool) ToIppoolOutput() IppoolOutput {
	return i.ToIppoolOutputWithContext(context.Background())
}

func (i *Ippool) ToIppoolOutputWithContext(ctx context.Context) IppoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IppoolOutput)
}

// IppoolArrayInput is an input type that accepts IppoolArray and IppoolArrayOutput values.
// You can construct a concrete instance of `IppoolArrayInput` via:
//
//	IppoolArray{ IppoolArgs{...} }
type IppoolArrayInput interface {
	pulumi.Input

	ToIppoolArrayOutput() IppoolArrayOutput
	ToIppoolArrayOutputWithContext(context.Context) IppoolArrayOutput
}

type IppoolArray []IppoolInput

func (IppoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ippool)(nil)).Elem()
}

func (i IppoolArray) ToIppoolArrayOutput() IppoolArrayOutput {
	return i.ToIppoolArrayOutputWithContext(context.Background())
}

func (i IppoolArray) ToIppoolArrayOutputWithContext(ctx context.Context) IppoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IppoolArrayOutput)
}

// IppoolMapInput is an input type that accepts IppoolMap and IppoolMapOutput values.
// You can construct a concrete instance of `IppoolMapInput` via:
//
//	IppoolMap{ "key": IppoolArgs{...} }
type IppoolMapInput interface {
	pulumi.Input

	ToIppoolMapOutput() IppoolMapOutput
	ToIppoolMapOutputWithContext(context.Context) IppoolMapOutput
}

type IppoolMap map[string]IppoolInput

func (IppoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ippool)(nil)).Elem()
}

func (i IppoolMap) ToIppoolMapOutput() IppoolMapOutput {
	return i.ToIppoolMapOutputWithContext(context.Background())
}

func (i IppoolMap) ToIppoolMapOutputWithContext(ctx context.Context) IppoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IppoolMapOutput)
}

type IppoolOutput struct{ *pulumi.OutputState }

func (IppoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ippool)(nil)).Elem()
}

func (o IppoolOutput) ToIppoolOutput() IppoolOutput {
	return o
}

func (o IppoolOutput) ToIppoolOutputWithContext(ctx context.Context) IppoolOutput {
	return o
}

// Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.
func (o IppoolOutput) AddNat64Route() pulumi.StringOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringOutput { return v.AddNat64Route }).(pulumi.StringOutput)
}

// Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
func (o IppoolOutput) ArpIntf() pulumi.StringOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringOutput { return v.ArpIntf }).(pulumi.StringOutput)
}

// Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
func (o IppoolOutput) ArpReply() pulumi.StringOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringOutput { return v.ArpReply }).(pulumi.StringOutput)
}

// Associated interface name.
func (o IppoolOutput) AssociatedInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringOutput { return v.AssociatedInterface }).(pulumi.StringOutput)
}

// Number of addresses in a block (64 - 4096, default = 128).
func (o IppoolOutput) BlockSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Ippool) pulumi.IntOutput { return v.BlockSize }).(pulumi.IntOutput)
}

// Comment.
func (o IppoolOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
func (o IppoolOutput) Endip() pulumi.StringOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringOutput { return v.Endip }).(pulumi.StringOutput)
}

// Final port number (inclusive) in the range for the address pool (Default: 65533).
func (o IppoolOutput) Endport() pulumi.IntOutput {
	return o.ApplyT(func(v *Ippool) pulumi.IntOutput { return v.Endport }).(pulumi.IntOutput)
}

// IP pool name.
func (o IppoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable NAT64. Valid values: `disable`, `enable`.
func (o IppoolOutput) Nat64() pulumi.StringOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringOutput { return v.Nat64 }).(pulumi.StringOutput)
}

// Number of addresses blocks that can be used by a user (1 to 128, default = 8).
func (o IppoolOutput) NumBlocksPerUser() pulumi.IntOutput {
	return o.ApplyT(func(v *Ippool) pulumi.IntOutput { return v.NumBlocksPerUser }).(pulumi.IntOutput)
}

// Port block allocation timeout (seconds).
func (o IppoolOutput) PbaTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Ippool) pulumi.IntOutput { return v.PbaTimeout }).(pulumi.IntOutput)
}

// Enable/disable full cone NAT. Valid values: `disable`, `enable`.
func (o IppoolOutput) PermitAnyHost() pulumi.StringOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringOutput { return v.PermitAnyHost }).(pulumi.StringOutput)
}

// Number of port for each user (32 - 60416, default = 0, which is auto).
func (o IppoolOutput) PortPerUser() pulumi.IntOutput {
	return o.ApplyT(func(v *Ippool) pulumi.IntOutput { return v.PortPerUser }).(pulumi.IntOutput)
}

// Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
func (o IppoolOutput) SourceEndip() pulumi.StringOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringOutput { return v.SourceEndip }).(pulumi.StringOutput)
}

// First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
func (o IppoolOutput) SourceStartip() pulumi.StringOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringOutput { return v.SourceStartip }).(pulumi.StringOutput)
}

// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
func (o IppoolOutput) Startip() pulumi.StringOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringOutput { return v.Startip }).(pulumi.StringOutput)
}

// First port number (inclusive) in the range for the address pool (Default: 5117).
func (o IppoolOutput) Startport() pulumi.IntOutput {
	return o.ApplyT(func(v *Ippool) pulumi.IntOutput { return v.Startport }).(pulumi.IntOutput)
}

// Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.
func (o IppoolOutput) SubnetBroadcastInIppool() pulumi.StringOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringOutput { return v.SubnetBroadcastInIppool }).(pulumi.StringOutput)
}

// IP pool type. On FortiOS versions 6.2.0-7.4.1: overload, one-to-one, fixed port range, or port block allocation. On FortiOS versions >= 7.4.2: overload, one-to-one, fixed-port-range, port-block-allocation, cgn-resource-allocation (hyperscale vdom only). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
func (o IppoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o IppoolOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ippool) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type IppoolArrayOutput struct{ *pulumi.OutputState }

func (IppoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ippool)(nil)).Elem()
}

func (o IppoolArrayOutput) ToIppoolArrayOutput() IppoolArrayOutput {
	return o
}

func (o IppoolArrayOutput) ToIppoolArrayOutputWithContext(ctx context.Context) IppoolArrayOutput {
	return o
}

func (o IppoolArrayOutput) Index(i pulumi.IntInput) IppoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ippool {
		return vs[0].([]*Ippool)[vs[1].(int)]
	}).(IppoolOutput)
}

type IppoolMapOutput struct{ *pulumi.OutputState }

func (IppoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ippool)(nil)).Elem()
}

func (o IppoolMapOutput) ToIppoolMapOutput() IppoolMapOutput {
	return o
}

func (o IppoolMapOutput) ToIppoolMapOutputWithContext(ctx context.Context) IppoolMapOutput {
	return o
}

func (o IppoolMapOutput) MapIndex(k pulumi.StringInput) IppoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ippool {
		return vs[0].(map[string]*Ippool)[vs[1].(string)]
	}).(IppoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IppoolInput)(nil)).Elem(), &Ippool{})
	pulumi.RegisterInputType(reflect.TypeOf((*IppoolArrayInput)(nil)).Elem(), IppoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IppoolMapInput)(nil)).Elem(), IppoolMap{})
	pulumi.RegisterOutputType(IppoolOutput{})
	pulumi.RegisterOutputType(IppoolArrayOutput{})
	pulumi.RegisterOutputType(IppoolMapOutput{})
}
