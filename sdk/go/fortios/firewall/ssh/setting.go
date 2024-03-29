// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// SSH proxy settings.
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
//			_, err := firewall.NewSetting(ctx, "trname", &firewall.SettingArgs{
//				Caname:              pulumi.String("Fortinet_SSH_CA"),
//				HostTrustedChecking: pulumi.String("enable"),
//				HostkeyDsa1024:      pulumi.String("Fortinet_SSH_DSA1024"),
//				HostkeyEcdsa256:     pulumi.String("Fortinet_SSH_ECDSA256"),
//				HostkeyEcdsa384:     pulumi.String("Fortinet_SSH_ECDSA384"),
//				HostkeyEcdsa521:     pulumi.String("Fortinet_SSH_ECDSA521"),
//				HostkeyEd25519:      pulumi.String("Fortinet_SSH_ED25519"),
//				HostkeyRsa2048:      pulumi.String("Fortinet_SSH_RSA2048"),
//				UntrustedCaname:     pulumi.String("Fortinet_SSH_CA_Untrusted"),
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
// FirewallSsh Setting can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/ssh/setting:Setting labelname FirewallSshSetting
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/ssh/setting:Setting labelname FirewallSshSetting
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Setting struct {
	pulumi.CustomResourceState

	// CA certificate used by SSH Inspection.
	Caname pulumi.StringOutput `pulumi:"caname"`
	// Enable/disable host trusted checking. Valid values: `enable`, `disable`.
	HostTrustedChecking pulumi.StringOutput `pulumi:"hostTrustedChecking"`
	// DSA certificate used by SSH proxy.
	HostkeyDsa1024 pulumi.StringOutput `pulumi:"hostkeyDsa1024"`
	// ECDSA nid256 certificate used by SSH proxy.
	HostkeyEcdsa256 pulumi.StringOutput `pulumi:"hostkeyEcdsa256"`
	// ECDSA nid384 certificate used by SSH proxy.
	HostkeyEcdsa384 pulumi.StringOutput `pulumi:"hostkeyEcdsa384"`
	// ECDSA nid384 certificate used by SSH proxy.
	HostkeyEcdsa521 pulumi.StringOutput `pulumi:"hostkeyEcdsa521"`
	// ED25519 hostkey used by SSH proxy.
	HostkeyEd25519 pulumi.StringOutput `pulumi:"hostkeyEd25519"`
	// RSA certificate used by SSH proxy.
	HostkeyRsa2048 pulumi.StringOutput `pulumi:"hostkeyRsa2048"`
	// Untrusted CA certificate used by SSH Inspection.
	UntrustedCaname pulumi.StringOutput `pulumi:"untrustedCaname"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSetting registers a new resource with the given unique name, arguments, and options.
func NewSetting(ctx *pulumi.Context,
	name string, args *SettingArgs, opts ...pulumi.ResourceOption) (*Setting, error) {
	if args == nil {
		args = &SettingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Setting
	err := ctx.RegisterResource("fortios:firewall/ssh/setting:Setting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSetting gets an existing Setting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SettingState, opts ...pulumi.ResourceOption) (*Setting, error) {
	var resource Setting
	err := ctx.ReadResource("fortios:firewall/ssh/setting:Setting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Setting resources.
type settingState struct {
	// CA certificate used by SSH Inspection.
	Caname *string `pulumi:"caname"`
	// Enable/disable host trusted checking. Valid values: `enable`, `disable`.
	HostTrustedChecking *string `pulumi:"hostTrustedChecking"`
	// DSA certificate used by SSH proxy.
	HostkeyDsa1024 *string `pulumi:"hostkeyDsa1024"`
	// ECDSA nid256 certificate used by SSH proxy.
	HostkeyEcdsa256 *string `pulumi:"hostkeyEcdsa256"`
	// ECDSA nid384 certificate used by SSH proxy.
	HostkeyEcdsa384 *string `pulumi:"hostkeyEcdsa384"`
	// ECDSA nid384 certificate used by SSH proxy.
	HostkeyEcdsa521 *string `pulumi:"hostkeyEcdsa521"`
	// ED25519 hostkey used by SSH proxy.
	HostkeyEd25519 *string `pulumi:"hostkeyEd25519"`
	// RSA certificate used by SSH proxy.
	HostkeyRsa2048 *string `pulumi:"hostkeyRsa2048"`
	// Untrusted CA certificate used by SSH Inspection.
	UntrustedCaname *string `pulumi:"untrustedCaname"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SettingState struct {
	// CA certificate used by SSH Inspection.
	Caname pulumi.StringPtrInput
	// Enable/disable host trusted checking. Valid values: `enable`, `disable`.
	HostTrustedChecking pulumi.StringPtrInput
	// DSA certificate used by SSH proxy.
	HostkeyDsa1024 pulumi.StringPtrInput
	// ECDSA nid256 certificate used by SSH proxy.
	HostkeyEcdsa256 pulumi.StringPtrInput
	// ECDSA nid384 certificate used by SSH proxy.
	HostkeyEcdsa384 pulumi.StringPtrInput
	// ECDSA nid384 certificate used by SSH proxy.
	HostkeyEcdsa521 pulumi.StringPtrInput
	// ED25519 hostkey used by SSH proxy.
	HostkeyEd25519 pulumi.StringPtrInput
	// RSA certificate used by SSH proxy.
	HostkeyRsa2048 pulumi.StringPtrInput
	// Untrusted CA certificate used by SSH Inspection.
	UntrustedCaname pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*settingState)(nil)).Elem()
}

type settingArgs struct {
	// CA certificate used by SSH Inspection.
	Caname *string `pulumi:"caname"`
	// Enable/disable host trusted checking. Valid values: `enable`, `disable`.
	HostTrustedChecking *string `pulumi:"hostTrustedChecking"`
	// DSA certificate used by SSH proxy.
	HostkeyDsa1024 *string `pulumi:"hostkeyDsa1024"`
	// ECDSA nid256 certificate used by SSH proxy.
	HostkeyEcdsa256 *string `pulumi:"hostkeyEcdsa256"`
	// ECDSA nid384 certificate used by SSH proxy.
	HostkeyEcdsa384 *string `pulumi:"hostkeyEcdsa384"`
	// ECDSA nid384 certificate used by SSH proxy.
	HostkeyEcdsa521 *string `pulumi:"hostkeyEcdsa521"`
	// ED25519 hostkey used by SSH proxy.
	HostkeyEd25519 *string `pulumi:"hostkeyEd25519"`
	// RSA certificate used by SSH proxy.
	HostkeyRsa2048 *string `pulumi:"hostkeyRsa2048"`
	// Untrusted CA certificate used by SSH Inspection.
	UntrustedCaname *string `pulumi:"untrustedCaname"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Setting resource.
type SettingArgs struct {
	// CA certificate used by SSH Inspection.
	Caname pulumi.StringPtrInput
	// Enable/disable host trusted checking. Valid values: `enable`, `disable`.
	HostTrustedChecking pulumi.StringPtrInput
	// DSA certificate used by SSH proxy.
	HostkeyDsa1024 pulumi.StringPtrInput
	// ECDSA nid256 certificate used by SSH proxy.
	HostkeyEcdsa256 pulumi.StringPtrInput
	// ECDSA nid384 certificate used by SSH proxy.
	HostkeyEcdsa384 pulumi.StringPtrInput
	// ECDSA nid384 certificate used by SSH proxy.
	HostkeyEcdsa521 pulumi.StringPtrInput
	// ED25519 hostkey used by SSH proxy.
	HostkeyEd25519 pulumi.StringPtrInput
	// RSA certificate used by SSH proxy.
	HostkeyRsa2048 pulumi.StringPtrInput
	// Untrusted CA certificate used by SSH Inspection.
	UntrustedCaname pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*settingArgs)(nil)).Elem()
}

type SettingInput interface {
	pulumi.Input

	ToSettingOutput() SettingOutput
	ToSettingOutputWithContext(ctx context.Context) SettingOutput
}

func (*Setting) ElementType() reflect.Type {
	return reflect.TypeOf((**Setting)(nil)).Elem()
}

func (i *Setting) ToSettingOutput() SettingOutput {
	return i.ToSettingOutputWithContext(context.Background())
}

func (i *Setting) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingOutput)
}

// SettingArrayInput is an input type that accepts SettingArray and SettingArrayOutput values.
// You can construct a concrete instance of `SettingArrayInput` via:
//
//	SettingArray{ SettingArgs{...} }
type SettingArrayInput interface {
	pulumi.Input

	ToSettingArrayOutput() SettingArrayOutput
	ToSettingArrayOutputWithContext(context.Context) SettingArrayOutput
}

type SettingArray []SettingInput

func (SettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Setting)(nil)).Elem()
}

func (i SettingArray) ToSettingArrayOutput() SettingArrayOutput {
	return i.ToSettingArrayOutputWithContext(context.Background())
}

func (i SettingArray) ToSettingArrayOutputWithContext(ctx context.Context) SettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingArrayOutput)
}

// SettingMapInput is an input type that accepts SettingMap and SettingMapOutput values.
// You can construct a concrete instance of `SettingMapInput` via:
//
//	SettingMap{ "key": SettingArgs{...} }
type SettingMapInput interface {
	pulumi.Input

	ToSettingMapOutput() SettingMapOutput
	ToSettingMapOutputWithContext(context.Context) SettingMapOutput
}

type SettingMap map[string]SettingInput

func (SettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Setting)(nil)).Elem()
}

func (i SettingMap) ToSettingMapOutput() SettingMapOutput {
	return i.ToSettingMapOutputWithContext(context.Background())
}

func (i SettingMap) ToSettingMapOutputWithContext(ctx context.Context) SettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingMapOutput)
}

type SettingOutput struct{ *pulumi.OutputState }

func (SettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Setting)(nil)).Elem()
}

func (o SettingOutput) ToSettingOutput() SettingOutput {
	return o
}

func (o SettingOutput) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return o
}

// CA certificate used by SSH Inspection.
func (o SettingOutput) Caname() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Caname }).(pulumi.StringOutput)
}

// Enable/disable host trusted checking. Valid values: `enable`, `disable`.
func (o SettingOutput) HostTrustedChecking() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.HostTrustedChecking }).(pulumi.StringOutput)
}

// DSA certificate used by SSH proxy.
func (o SettingOutput) HostkeyDsa1024() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.HostkeyDsa1024 }).(pulumi.StringOutput)
}

// ECDSA nid256 certificate used by SSH proxy.
func (o SettingOutput) HostkeyEcdsa256() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.HostkeyEcdsa256 }).(pulumi.StringOutput)
}

// ECDSA nid384 certificate used by SSH proxy.
func (o SettingOutput) HostkeyEcdsa384() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.HostkeyEcdsa384 }).(pulumi.StringOutput)
}

// ECDSA nid384 certificate used by SSH proxy.
func (o SettingOutput) HostkeyEcdsa521() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.HostkeyEcdsa521 }).(pulumi.StringOutput)
}

// ED25519 hostkey used by SSH proxy.
func (o SettingOutput) HostkeyEd25519() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.HostkeyEd25519 }).(pulumi.StringOutput)
}

// RSA certificate used by SSH proxy.
func (o SettingOutput) HostkeyRsa2048() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.HostkeyRsa2048 }).(pulumi.StringOutput)
}

// Untrusted CA certificate used by SSH Inspection.
func (o SettingOutput) UntrustedCaname() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.UntrustedCaname }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SettingArrayOutput struct{ *pulumi.OutputState }

func (SettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Setting)(nil)).Elem()
}

func (o SettingArrayOutput) ToSettingArrayOutput() SettingArrayOutput {
	return o
}

func (o SettingArrayOutput) ToSettingArrayOutputWithContext(ctx context.Context) SettingArrayOutput {
	return o
}

func (o SettingArrayOutput) Index(i pulumi.IntInput) SettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Setting {
		return vs[0].([]*Setting)[vs[1].(int)]
	}).(SettingOutput)
}

type SettingMapOutput struct{ *pulumi.OutputState }

func (SettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Setting)(nil)).Elem()
}

func (o SettingMapOutput) ToSettingMapOutput() SettingMapOutput {
	return o
}

func (o SettingMapOutput) ToSettingMapOutputWithContext(ctx context.Context) SettingMapOutput {
	return o
}

func (o SettingMapOutput) MapIndex(k pulumi.StringInput) SettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Setting {
		return vs[0].(map[string]*Setting)[vs[1].(string)]
	}).(SettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SettingInput)(nil)).Elem(), &Setting{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingArrayInput)(nil)).Elem(), SettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingMapInput)(nil)).Elem(), SettingMap{})
	pulumi.RegisterOutputType(SettingOutput{})
	pulumi.RegisterOutputType(SettingArrayOutput{})
	pulumi.RegisterOutputType(SettingMapOutput{})
}
