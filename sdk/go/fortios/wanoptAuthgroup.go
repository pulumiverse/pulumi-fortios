// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure WAN optimization authentication groups.
//
// ## Example Usage
//
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
//			_, err := fortios.NewWanoptAuthgroup(ctx, "trname", &fortios.WanoptAuthgroupArgs{
//				AuthMethod: pulumi.String("cert"),
//				Cert:       pulumi.String("Fortinet_CA_SSL"),
//				PeerAccept: pulumi.String("any"),
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
// # Wanopt AuthGroup can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wanoptAuthgroup:WanoptAuthgroup labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wanoptAuthgroup:WanoptAuthgroup labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WanoptAuthgroup struct {
	pulumi.CustomResourceState

	// Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
	AuthMethod pulumi.StringOutput `pulumi:"authMethod"`
	// Name of certificate to identify this peer.
	Cert pulumi.StringOutput `pulumi:"cert"`
	// Auth-group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
	Peer pulumi.StringOutput `pulumi:"peer"`
	// Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
	PeerAccept pulumi.StringOutput `pulumi:"peerAccept"`
	// Pre-shared key used by the peers in this authentication group.
	Psk pulumi.StringPtrOutput `pulumi:"psk"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWanoptAuthgroup registers a new resource with the given unique name, arguments, and options.
func NewWanoptAuthgroup(ctx *pulumi.Context,
	name string, args *WanoptAuthgroupArgs, opts ...pulumi.ResourceOption) (*WanoptAuthgroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cert == nil {
		return nil, errors.New("invalid value for required argument 'Cert'")
	}
	if args.Psk != nil {
		args.Psk = pulumi.ToSecret(args.Psk).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"psk",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource WanoptAuthgroup
	err := ctx.RegisterResource("fortios:index/wanoptAuthgroup:WanoptAuthgroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWanoptAuthgroup gets an existing WanoptAuthgroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWanoptAuthgroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WanoptAuthgroupState, opts ...pulumi.ResourceOption) (*WanoptAuthgroup, error) {
	var resource WanoptAuthgroup
	err := ctx.ReadResource("fortios:index/wanoptAuthgroup:WanoptAuthgroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WanoptAuthgroup resources.
type wanoptAuthgroupState struct {
	// Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
	AuthMethod *string `pulumi:"authMethod"`
	// Name of certificate to identify this peer.
	Cert *string `pulumi:"cert"`
	// Auth-group name.
	Name *string `pulumi:"name"`
	// If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
	Peer *string `pulumi:"peer"`
	// Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
	PeerAccept *string `pulumi:"peerAccept"`
	// Pre-shared key used by the peers in this authentication group.
	Psk *string `pulumi:"psk"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WanoptAuthgroupState struct {
	// Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
	AuthMethod pulumi.StringPtrInput
	// Name of certificate to identify this peer.
	Cert pulumi.StringPtrInput
	// Auth-group name.
	Name pulumi.StringPtrInput
	// If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
	Peer pulumi.StringPtrInput
	// Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
	PeerAccept pulumi.StringPtrInput
	// Pre-shared key used by the peers in this authentication group.
	Psk pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WanoptAuthgroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptAuthgroupState)(nil)).Elem()
}

type wanoptAuthgroupArgs struct {
	// Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
	AuthMethod *string `pulumi:"authMethod"`
	// Name of certificate to identify this peer.
	Cert string `pulumi:"cert"`
	// Auth-group name.
	Name *string `pulumi:"name"`
	// If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
	Peer *string `pulumi:"peer"`
	// Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
	PeerAccept *string `pulumi:"peerAccept"`
	// Pre-shared key used by the peers in this authentication group.
	Psk *string `pulumi:"psk"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WanoptAuthgroup resource.
type WanoptAuthgroupArgs struct {
	// Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
	AuthMethod pulumi.StringPtrInput
	// Name of certificate to identify this peer.
	Cert pulumi.StringInput
	// Auth-group name.
	Name pulumi.StringPtrInput
	// If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
	Peer pulumi.StringPtrInput
	// Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
	PeerAccept pulumi.StringPtrInput
	// Pre-shared key used by the peers in this authentication group.
	Psk pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WanoptAuthgroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptAuthgroupArgs)(nil)).Elem()
}

type WanoptAuthgroupInput interface {
	pulumi.Input

	ToWanoptAuthgroupOutput() WanoptAuthgroupOutput
	ToWanoptAuthgroupOutputWithContext(ctx context.Context) WanoptAuthgroupOutput
}

func (*WanoptAuthgroup) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptAuthgroup)(nil)).Elem()
}

func (i *WanoptAuthgroup) ToWanoptAuthgroupOutput() WanoptAuthgroupOutput {
	return i.ToWanoptAuthgroupOutputWithContext(context.Background())
}

func (i *WanoptAuthgroup) ToWanoptAuthgroupOutputWithContext(ctx context.Context) WanoptAuthgroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptAuthgroupOutput)
}

// WanoptAuthgroupArrayInput is an input type that accepts WanoptAuthgroupArray and WanoptAuthgroupArrayOutput values.
// You can construct a concrete instance of `WanoptAuthgroupArrayInput` via:
//
//	WanoptAuthgroupArray{ WanoptAuthgroupArgs{...} }
type WanoptAuthgroupArrayInput interface {
	pulumi.Input

	ToWanoptAuthgroupArrayOutput() WanoptAuthgroupArrayOutput
	ToWanoptAuthgroupArrayOutputWithContext(context.Context) WanoptAuthgroupArrayOutput
}

type WanoptAuthgroupArray []WanoptAuthgroupInput

func (WanoptAuthgroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WanoptAuthgroup)(nil)).Elem()
}

func (i WanoptAuthgroupArray) ToWanoptAuthgroupArrayOutput() WanoptAuthgroupArrayOutput {
	return i.ToWanoptAuthgroupArrayOutputWithContext(context.Background())
}

func (i WanoptAuthgroupArray) ToWanoptAuthgroupArrayOutputWithContext(ctx context.Context) WanoptAuthgroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptAuthgroupArrayOutput)
}

// WanoptAuthgroupMapInput is an input type that accepts WanoptAuthgroupMap and WanoptAuthgroupMapOutput values.
// You can construct a concrete instance of `WanoptAuthgroupMapInput` via:
//
//	WanoptAuthgroupMap{ "key": WanoptAuthgroupArgs{...} }
type WanoptAuthgroupMapInput interface {
	pulumi.Input

	ToWanoptAuthgroupMapOutput() WanoptAuthgroupMapOutput
	ToWanoptAuthgroupMapOutputWithContext(context.Context) WanoptAuthgroupMapOutput
}

type WanoptAuthgroupMap map[string]WanoptAuthgroupInput

func (WanoptAuthgroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WanoptAuthgroup)(nil)).Elem()
}

func (i WanoptAuthgroupMap) ToWanoptAuthgroupMapOutput() WanoptAuthgroupMapOutput {
	return i.ToWanoptAuthgroupMapOutputWithContext(context.Background())
}

func (i WanoptAuthgroupMap) ToWanoptAuthgroupMapOutputWithContext(ctx context.Context) WanoptAuthgroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptAuthgroupMapOutput)
}

type WanoptAuthgroupOutput struct{ *pulumi.OutputState }

func (WanoptAuthgroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptAuthgroup)(nil)).Elem()
}

func (o WanoptAuthgroupOutput) ToWanoptAuthgroupOutput() WanoptAuthgroupOutput {
	return o
}

func (o WanoptAuthgroupOutput) ToWanoptAuthgroupOutputWithContext(ctx context.Context) WanoptAuthgroupOutput {
	return o
}

// Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
func (o WanoptAuthgroupOutput) AuthMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *WanoptAuthgroup) pulumi.StringOutput { return v.AuthMethod }).(pulumi.StringOutput)
}

// Name of certificate to identify this peer.
func (o WanoptAuthgroupOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v *WanoptAuthgroup) pulumi.StringOutput { return v.Cert }).(pulumi.StringOutput)
}

// Auth-group name.
func (o WanoptAuthgroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WanoptAuthgroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
func (o WanoptAuthgroupOutput) Peer() pulumi.StringOutput {
	return o.ApplyT(func(v *WanoptAuthgroup) pulumi.StringOutput { return v.Peer }).(pulumi.StringOutput)
}

// Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
func (o WanoptAuthgroupOutput) PeerAccept() pulumi.StringOutput {
	return o.ApplyT(func(v *WanoptAuthgroup) pulumi.StringOutput { return v.PeerAccept }).(pulumi.StringOutput)
}

// Pre-shared key used by the peers in this authentication group.
func (o WanoptAuthgroupOutput) Psk() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WanoptAuthgroup) pulumi.StringPtrOutput { return v.Psk }).(pulumi.StringPtrOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WanoptAuthgroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WanoptAuthgroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WanoptAuthgroupArrayOutput struct{ *pulumi.OutputState }

func (WanoptAuthgroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WanoptAuthgroup)(nil)).Elem()
}

func (o WanoptAuthgroupArrayOutput) ToWanoptAuthgroupArrayOutput() WanoptAuthgroupArrayOutput {
	return o
}

func (o WanoptAuthgroupArrayOutput) ToWanoptAuthgroupArrayOutputWithContext(ctx context.Context) WanoptAuthgroupArrayOutput {
	return o
}

func (o WanoptAuthgroupArrayOutput) Index(i pulumi.IntInput) WanoptAuthgroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WanoptAuthgroup {
		return vs[0].([]*WanoptAuthgroup)[vs[1].(int)]
	}).(WanoptAuthgroupOutput)
}

type WanoptAuthgroupMapOutput struct{ *pulumi.OutputState }

func (WanoptAuthgroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WanoptAuthgroup)(nil)).Elem()
}

func (o WanoptAuthgroupMapOutput) ToWanoptAuthgroupMapOutput() WanoptAuthgroupMapOutput {
	return o
}

func (o WanoptAuthgroupMapOutput) ToWanoptAuthgroupMapOutputWithContext(ctx context.Context) WanoptAuthgroupMapOutput {
	return o
}

func (o WanoptAuthgroupMapOutput) MapIndex(k pulumi.StringInput) WanoptAuthgroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WanoptAuthgroup {
		return vs[0].(map[string]*WanoptAuthgroup)[vs[1].(string)]
	}).(WanoptAuthgroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptAuthgroupInput)(nil)).Elem(), &WanoptAuthgroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptAuthgroupArrayInput)(nil)).Elem(), WanoptAuthgroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptAuthgroupMapInput)(nil)).Elem(), WanoptAuthgroupMap{})
	pulumi.RegisterOutputType(WanoptAuthgroupOutput{})
	pulumi.RegisterOutputType(WanoptAuthgroupArrayOutput{})
	pulumi.RegisterOutputType(WanoptAuthgroupMapOutput{})
}
