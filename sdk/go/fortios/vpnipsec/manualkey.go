// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpnipsec

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPsec manual keys.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/vpnipsec"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpnipsec.NewManualkey(ctx, "trname", &vpnipsec.ManualkeyArgs{
//				Authentication: pulumi.String("md5"),
//				Authkey:        pulumi.String("EE32CA121ECD772A-ECACAABA212345EC"),
//				Enckey:         pulumi.String("-"),
//				Encryption:     pulumi.String("null"),
//				Interface:      pulumi.String("port4"),
//				LocalGw:        pulumi.String("0.0.0.0"),
//				Localspi:       pulumi.String("0x100"),
//				RemoteGw:       pulumi.String("1.1.1.1"),
//				Remotespi:      pulumi.String("0x100"),
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
// # VpnIpsec Manualkey can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:vpnipsec/manualkey:Manualkey labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:vpnipsec/manualkey:Manualkey labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Manualkey struct {
	pulumi.CustomResourceState

	// Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
	Authentication pulumi.StringOutput `pulumi:"authentication"`
	// Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
	Authkey pulumi.StringOutput `pulumi:"authkey"`
	// Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
	Enckey pulumi.StringOutput `pulumi:"enckey"`
	// Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
	Encryption pulumi.StringOutput `pulumi:"encryption"`
	// Name of the physical, aggregate, or VLAN interface.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Local gateway.
	LocalGw pulumi.StringOutput `pulumi:"localGw"`
	// Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
	Localspi pulumi.StringOutput `pulumi:"localspi"`
	// IPsec tunnel name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable NPU offloading. Valid values: `enable`, `disable`.
	NpuOffload pulumi.StringOutput `pulumi:"npuOffload"`
	// Peer gateway.
	RemoteGw pulumi.StringOutput `pulumi:"remoteGw"`
	// Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
	Remotespi pulumi.StringOutput `pulumi:"remotespi"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewManualkey registers a new resource with the given unique name, arguments, and options.
func NewManualkey(ctx *pulumi.Context,
	name string, args *ManualkeyArgs, opts ...pulumi.ResourceOption) (*Manualkey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authentication == nil {
		return nil, errors.New("invalid value for required argument 'Authentication'")
	}
	if args.Encryption == nil {
		return nil, errors.New("invalid value for required argument 'Encryption'")
	}
	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.RemoteGw == nil {
		return nil, errors.New("invalid value for required argument 'RemoteGw'")
	}
	if args.Authkey != nil {
		args.Authkey = pulumi.ToSecret(args.Authkey).(pulumi.StringPtrInput)
	}
	if args.Enckey != nil {
		args.Enckey = pulumi.ToSecret(args.Enckey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"authkey",
		"enckey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Manualkey
	err := ctx.RegisterResource("fortios:vpnipsec/manualkey:Manualkey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManualkey gets an existing Manualkey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManualkey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManualkeyState, opts ...pulumi.ResourceOption) (*Manualkey, error) {
	var resource Manualkey
	err := ctx.ReadResource("fortios:vpnipsec/manualkey:Manualkey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Manualkey resources.
type manualkeyState struct {
	// Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
	Authentication *string `pulumi:"authentication"`
	// Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
	Authkey *string `pulumi:"authkey"`
	// Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
	Enckey *string `pulumi:"enckey"`
	// Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
	Encryption *string `pulumi:"encryption"`
	// Name of the physical, aggregate, or VLAN interface.
	Interface *string `pulumi:"interface"`
	// Local gateway.
	LocalGw *string `pulumi:"localGw"`
	// Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
	Localspi *string `pulumi:"localspi"`
	// IPsec tunnel name.
	Name *string `pulumi:"name"`
	// Enable/disable NPU offloading. Valid values: `enable`, `disable`.
	NpuOffload *string `pulumi:"npuOffload"`
	// Peer gateway.
	RemoteGw *string `pulumi:"remoteGw"`
	// Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
	Remotespi *string `pulumi:"remotespi"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ManualkeyState struct {
	// Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
	Authentication pulumi.StringPtrInput
	// Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
	Authkey pulumi.StringPtrInput
	// Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
	Enckey pulumi.StringPtrInput
	// Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
	Encryption pulumi.StringPtrInput
	// Name of the physical, aggregate, or VLAN interface.
	Interface pulumi.StringPtrInput
	// Local gateway.
	LocalGw pulumi.StringPtrInput
	// Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
	Localspi pulumi.StringPtrInput
	// IPsec tunnel name.
	Name pulumi.StringPtrInput
	// Enable/disable NPU offloading. Valid values: `enable`, `disable`.
	NpuOffload pulumi.StringPtrInput
	// Peer gateway.
	RemoteGw pulumi.StringPtrInput
	// Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
	Remotespi pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ManualkeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*manualkeyState)(nil)).Elem()
}

type manualkeyArgs struct {
	// Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
	Authentication string `pulumi:"authentication"`
	// Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
	Authkey *string `pulumi:"authkey"`
	// Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
	Enckey *string `pulumi:"enckey"`
	// Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
	Encryption string `pulumi:"encryption"`
	// Name of the physical, aggregate, or VLAN interface.
	Interface string `pulumi:"interface"`
	// Local gateway.
	LocalGw *string `pulumi:"localGw"`
	// Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
	Localspi *string `pulumi:"localspi"`
	// IPsec tunnel name.
	Name *string `pulumi:"name"`
	// Enable/disable NPU offloading. Valid values: `enable`, `disable`.
	NpuOffload *string `pulumi:"npuOffload"`
	// Peer gateway.
	RemoteGw string `pulumi:"remoteGw"`
	// Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
	Remotespi *string `pulumi:"remotespi"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Manualkey resource.
type ManualkeyArgs struct {
	// Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
	Authentication pulumi.StringInput
	// Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
	Authkey pulumi.StringPtrInput
	// Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
	Enckey pulumi.StringPtrInput
	// Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
	Encryption pulumi.StringInput
	// Name of the physical, aggregate, or VLAN interface.
	Interface pulumi.StringInput
	// Local gateway.
	LocalGw pulumi.StringPtrInput
	// Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
	Localspi pulumi.StringPtrInput
	// IPsec tunnel name.
	Name pulumi.StringPtrInput
	// Enable/disable NPU offloading. Valid values: `enable`, `disable`.
	NpuOffload pulumi.StringPtrInput
	// Peer gateway.
	RemoteGw pulumi.StringInput
	// Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
	Remotespi pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ManualkeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*manualkeyArgs)(nil)).Elem()
}

type ManualkeyInput interface {
	pulumi.Input

	ToManualkeyOutput() ManualkeyOutput
	ToManualkeyOutputWithContext(ctx context.Context) ManualkeyOutput
}

func (*Manualkey) ElementType() reflect.Type {
	return reflect.TypeOf((**Manualkey)(nil)).Elem()
}

func (i *Manualkey) ToManualkeyOutput() ManualkeyOutput {
	return i.ToManualkeyOutputWithContext(context.Background())
}

func (i *Manualkey) ToManualkeyOutputWithContext(ctx context.Context) ManualkeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManualkeyOutput)
}

// ManualkeyArrayInput is an input type that accepts ManualkeyArray and ManualkeyArrayOutput values.
// You can construct a concrete instance of `ManualkeyArrayInput` via:
//
//	ManualkeyArray{ ManualkeyArgs{...} }
type ManualkeyArrayInput interface {
	pulumi.Input

	ToManualkeyArrayOutput() ManualkeyArrayOutput
	ToManualkeyArrayOutputWithContext(context.Context) ManualkeyArrayOutput
}

type ManualkeyArray []ManualkeyInput

func (ManualkeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Manualkey)(nil)).Elem()
}

func (i ManualkeyArray) ToManualkeyArrayOutput() ManualkeyArrayOutput {
	return i.ToManualkeyArrayOutputWithContext(context.Background())
}

func (i ManualkeyArray) ToManualkeyArrayOutputWithContext(ctx context.Context) ManualkeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManualkeyArrayOutput)
}

// ManualkeyMapInput is an input type that accepts ManualkeyMap and ManualkeyMapOutput values.
// You can construct a concrete instance of `ManualkeyMapInput` via:
//
//	ManualkeyMap{ "key": ManualkeyArgs{...} }
type ManualkeyMapInput interface {
	pulumi.Input

	ToManualkeyMapOutput() ManualkeyMapOutput
	ToManualkeyMapOutputWithContext(context.Context) ManualkeyMapOutput
}

type ManualkeyMap map[string]ManualkeyInput

func (ManualkeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Manualkey)(nil)).Elem()
}

func (i ManualkeyMap) ToManualkeyMapOutput() ManualkeyMapOutput {
	return i.ToManualkeyMapOutputWithContext(context.Background())
}

func (i ManualkeyMap) ToManualkeyMapOutputWithContext(ctx context.Context) ManualkeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManualkeyMapOutput)
}

type ManualkeyOutput struct{ *pulumi.OutputState }

func (ManualkeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Manualkey)(nil)).Elem()
}

func (o ManualkeyOutput) ToManualkeyOutput() ManualkeyOutput {
	return o
}

func (o ManualkeyOutput) ToManualkeyOutputWithContext(ctx context.Context) ManualkeyOutput {
	return o
}

// Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
func (o ManualkeyOutput) Authentication() pulumi.StringOutput {
	return o.ApplyT(func(v *Manualkey) pulumi.StringOutput { return v.Authentication }).(pulumi.StringOutput)
}

// Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
func (o ManualkeyOutput) Authkey() pulumi.StringOutput {
	return o.ApplyT(func(v *Manualkey) pulumi.StringOutput { return v.Authkey }).(pulumi.StringOutput)
}

// Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
func (o ManualkeyOutput) Enckey() pulumi.StringOutput {
	return o.ApplyT(func(v *Manualkey) pulumi.StringOutput { return v.Enckey }).(pulumi.StringOutput)
}

// Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
func (o ManualkeyOutput) Encryption() pulumi.StringOutput {
	return o.ApplyT(func(v *Manualkey) pulumi.StringOutput { return v.Encryption }).(pulumi.StringOutput)
}

// Name of the physical, aggregate, or VLAN interface.
func (o ManualkeyOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Manualkey) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Local gateway.
func (o ManualkeyOutput) LocalGw() pulumi.StringOutput {
	return o.ApplyT(func(v *Manualkey) pulumi.StringOutput { return v.LocalGw }).(pulumi.StringOutput)
}

// Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
func (o ManualkeyOutput) Localspi() pulumi.StringOutput {
	return o.ApplyT(func(v *Manualkey) pulumi.StringOutput { return v.Localspi }).(pulumi.StringOutput)
}

// IPsec tunnel name.
func (o ManualkeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Manualkey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable NPU offloading. Valid values: `enable`, `disable`.
func (o ManualkeyOutput) NpuOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *Manualkey) pulumi.StringOutput { return v.NpuOffload }).(pulumi.StringOutput)
}

// Peer gateway.
func (o ManualkeyOutput) RemoteGw() pulumi.StringOutput {
	return o.ApplyT(func(v *Manualkey) pulumi.StringOutput { return v.RemoteGw }).(pulumi.StringOutput)
}

// Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
func (o ManualkeyOutput) Remotespi() pulumi.StringOutput {
	return o.ApplyT(func(v *Manualkey) pulumi.StringOutput { return v.Remotespi }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ManualkeyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Manualkey) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ManualkeyArrayOutput struct{ *pulumi.OutputState }

func (ManualkeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Manualkey)(nil)).Elem()
}

func (o ManualkeyArrayOutput) ToManualkeyArrayOutput() ManualkeyArrayOutput {
	return o
}

func (o ManualkeyArrayOutput) ToManualkeyArrayOutputWithContext(ctx context.Context) ManualkeyArrayOutput {
	return o
}

func (o ManualkeyArrayOutput) Index(i pulumi.IntInput) ManualkeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Manualkey {
		return vs[0].([]*Manualkey)[vs[1].(int)]
	}).(ManualkeyOutput)
}

type ManualkeyMapOutput struct{ *pulumi.OutputState }

func (ManualkeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Manualkey)(nil)).Elem()
}

func (o ManualkeyMapOutput) ToManualkeyMapOutput() ManualkeyMapOutput {
	return o
}

func (o ManualkeyMapOutput) ToManualkeyMapOutputWithContext(ctx context.Context) ManualkeyMapOutput {
	return o
}

func (o ManualkeyMapOutput) MapIndex(k pulumi.StringInput) ManualkeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Manualkey {
		return vs[0].(map[string]*Manualkey)[vs[1].(string)]
	}).(ManualkeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManualkeyInput)(nil)).Elem(), &Manualkey{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManualkeyArrayInput)(nil)).Elem(), ManualkeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManualkeyMapInput)(nil)).Elem(), ManualkeyMap{})
	pulumi.RegisterOutputType(ManualkeyOutput{})
	pulumi.RegisterOutputType(ManualkeyArrayOutput{})
	pulumi.RegisterOutputType(ManualkeyMapOutput{})
}
