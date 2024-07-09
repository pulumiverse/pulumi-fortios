// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// SSL proxy settings.
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
//			_, err := firewall.NewSetting(ctx, "trname", &firewall.SettingArgs{
//				AbbreviateHandshake:    pulumi.String("enable"),
//				CertCacheCapacity:      pulumi.Int(200),
//				CertCacheTimeout:       pulumi.Int(10),
//				KxpQueueThreshold:      pulumi.Int(16),
//				NoMatchingCipherAction: pulumi.String("bypass"),
//				ProxyConnectTimeout:    pulumi.Int(30),
//				SessionCacheCapacity:   pulumi.Int(500),
//				SessionCacheTimeout:    pulumi.Int(20),
//				SslDhBits:              pulumi.String("2048"),
//				SslQueueThreshold:      pulumi.Int(32),
//				SslSendEmptyFrags:      pulumi.String("enable"),
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
// FirewallSsl Setting can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/ssl/setting:Setting labelname FirewallSslSetting
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/ssl/setting:Setting labelname FirewallSslSetting
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Setting struct {
	pulumi.CustomResourceState

	// Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
	AbbreviateHandshake pulumi.StringOutput `pulumi:"abbreviateHandshake"`
	// Maximum capacity of the host certificate cache (0 - 500, default = 200).
	CertCacheCapacity pulumi.IntOutput `pulumi:"certCacheCapacity"`
	// Time limit to keep certificate cache (1 - 120 min, default = 10).
	CertCacheTimeout pulumi.IntOutput `pulumi:"certCacheTimeout"`
	// Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
	KxpQueueThreshold pulumi.IntOutput `pulumi:"kxpQueueThreshold"`
	// Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
	NoMatchingCipherAction pulumi.StringOutput `pulumi:"noMatchingCipherAction"`
	// Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
	ProxyConnectTimeout pulumi.IntOutput `pulumi:"proxyConnectTimeout"`
	// Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
	SessionCacheCapacity pulumi.IntOutput `pulumi:"sessionCacheCapacity"`
	// Time limit to keep SSL session state (1 - 60 min, default = 20).
	SessionCacheTimeout pulumi.IntOutput `pulumi:"sessionCacheTimeout"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringOutput `pulumi:"sslDhBits"`
	// Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
	SslQueueThreshold pulumi.IntOutput `pulumi:"sslQueueThreshold"`
	// Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
	SslSendEmptyFrags pulumi.StringOutput `pulumi:"sslSendEmptyFrags"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewSetting registers a new resource with the given unique name, arguments, and options.
func NewSetting(ctx *pulumi.Context,
	name string, args *SettingArgs, opts ...pulumi.ResourceOption) (*Setting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertCacheCapacity == nil {
		return nil, errors.New("invalid value for required argument 'CertCacheCapacity'")
	}
	if args.CertCacheTimeout == nil {
		return nil, errors.New("invalid value for required argument 'CertCacheTimeout'")
	}
	if args.NoMatchingCipherAction == nil {
		return nil, errors.New("invalid value for required argument 'NoMatchingCipherAction'")
	}
	if args.ProxyConnectTimeout == nil {
		return nil, errors.New("invalid value for required argument 'ProxyConnectTimeout'")
	}
	if args.SessionCacheCapacity == nil {
		return nil, errors.New("invalid value for required argument 'SessionCacheCapacity'")
	}
	if args.SessionCacheTimeout == nil {
		return nil, errors.New("invalid value for required argument 'SessionCacheTimeout'")
	}
	if args.SslDhBits == nil {
		return nil, errors.New("invalid value for required argument 'SslDhBits'")
	}
	if args.SslSendEmptyFrags == nil {
		return nil, errors.New("invalid value for required argument 'SslSendEmptyFrags'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Setting
	err := ctx.RegisterResource("fortios:firewall/ssl/setting:Setting", name, args, &resource, opts...)
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
	err := ctx.ReadResource("fortios:firewall/ssl/setting:Setting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Setting resources.
type settingState struct {
	// Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
	AbbreviateHandshake *string `pulumi:"abbreviateHandshake"`
	// Maximum capacity of the host certificate cache (0 - 500, default = 200).
	CertCacheCapacity *int `pulumi:"certCacheCapacity"`
	// Time limit to keep certificate cache (1 - 120 min, default = 10).
	CertCacheTimeout *int `pulumi:"certCacheTimeout"`
	// Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
	KxpQueueThreshold *int `pulumi:"kxpQueueThreshold"`
	// Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
	NoMatchingCipherAction *string `pulumi:"noMatchingCipherAction"`
	// Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
	ProxyConnectTimeout *int `pulumi:"proxyConnectTimeout"`
	// Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
	SessionCacheCapacity *int `pulumi:"sessionCacheCapacity"`
	// Time limit to keep SSL session state (1 - 60 min, default = 20).
	SessionCacheTimeout *int `pulumi:"sessionCacheTimeout"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits *string `pulumi:"sslDhBits"`
	// Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
	SslQueueThreshold *int `pulumi:"sslQueueThreshold"`
	// Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
	SslSendEmptyFrags *string `pulumi:"sslSendEmptyFrags"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SettingState struct {
	// Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
	AbbreviateHandshake pulumi.StringPtrInput
	// Maximum capacity of the host certificate cache (0 - 500, default = 200).
	CertCacheCapacity pulumi.IntPtrInput
	// Time limit to keep certificate cache (1 - 120 min, default = 10).
	CertCacheTimeout pulumi.IntPtrInput
	// Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
	KxpQueueThreshold pulumi.IntPtrInput
	// Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
	NoMatchingCipherAction pulumi.StringPtrInput
	// Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
	ProxyConnectTimeout pulumi.IntPtrInput
	// Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
	SessionCacheCapacity pulumi.IntPtrInput
	// Time limit to keep SSL session state (1 - 60 min, default = 20).
	SessionCacheTimeout pulumi.IntPtrInput
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringPtrInput
	// Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
	SslQueueThreshold pulumi.IntPtrInput
	// Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
	SslSendEmptyFrags pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*settingState)(nil)).Elem()
}

type settingArgs struct {
	// Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
	AbbreviateHandshake *string `pulumi:"abbreviateHandshake"`
	// Maximum capacity of the host certificate cache (0 - 500, default = 200).
	CertCacheCapacity int `pulumi:"certCacheCapacity"`
	// Time limit to keep certificate cache (1 - 120 min, default = 10).
	CertCacheTimeout int `pulumi:"certCacheTimeout"`
	// Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
	KxpQueueThreshold *int `pulumi:"kxpQueueThreshold"`
	// Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
	NoMatchingCipherAction string `pulumi:"noMatchingCipherAction"`
	// Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
	ProxyConnectTimeout int `pulumi:"proxyConnectTimeout"`
	// Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
	SessionCacheCapacity int `pulumi:"sessionCacheCapacity"`
	// Time limit to keep SSL session state (1 - 60 min, default = 20).
	SessionCacheTimeout int `pulumi:"sessionCacheTimeout"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits string `pulumi:"sslDhBits"`
	// Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
	SslQueueThreshold *int `pulumi:"sslQueueThreshold"`
	// Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
	SslSendEmptyFrags string `pulumi:"sslSendEmptyFrags"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Setting resource.
type SettingArgs struct {
	// Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
	AbbreviateHandshake pulumi.StringPtrInput
	// Maximum capacity of the host certificate cache (0 - 500, default = 200).
	CertCacheCapacity pulumi.IntInput
	// Time limit to keep certificate cache (1 - 120 min, default = 10).
	CertCacheTimeout pulumi.IntInput
	// Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
	KxpQueueThreshold pulumi.IntPtrInput
	// Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
	NoMatchingCipherAction pulumi.StringInput
	// Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
	ProxyConnectTimeout pulumi.IntInput
	// Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
	SessionCacheCapacity pulumi.IntInput
	// Time limit to keep SSL session state (1 - 60 min, default = 20).
	SessionCacheTimeout pulumi.IntInput
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringInput
	// Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
	SslQueueThreshold pulumi.IntPtrInput
	// Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
	SslSendEmptyFrags pulumi.StringInput
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

// Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
func (o SettingOutput) AbbreviateHandshake() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.AbbreviateHandshake }).(pulumi.StringOutput)
}

// Maximum capacity of the host certificate cache (0 - 500, default = 200).
func (o SettingOutput) CertCacheCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.CertCacheCapacity }).(pulumi.IntOutput)
}

// Time limit to keep certificate cache (1 - 120 min, default = 10).
func (o SettingOutput) CertCacheTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.CertCacheTimeout }).(pulumi.IntOutput)
}

// Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
func (o SettingOutput) KxpQueueThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.KxpQueueThreshold }).(pulumi.IntOutput)
}

// Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
func (o SettingOutput) NoMatchingCipherAction() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.NoMatchingCipherAction }).(pulumi.StringOutput)
}

// Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
func (o SettingOutput) ProxyConnectTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.ProxyConnectTimeout }).(pulumi.IntOutput)
}

// Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
func (o SettingOutput) SessionCacheCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.SessionCacheCapacity }).(pulumi.IntOutput)
}

// Time limit to keep SSL session state (1 - 60 min, default = 20).
func (o SettingOutput) SessionCacheTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.SessionCacheTimeout }).(pulumi.IntOutput)
}

// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
func (o SettingOutput) SslDhBits() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.SslDhBits }).(pulumi.StringOutput)
}

// Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
func (o SettingOutput) SslQueueThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.SslQueueThreshold }).(pulumi.IntOutput)
}

// Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
func (o SettingOutput) SslSendEmptyFrags() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.SslSendEmptyFrags }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SettingOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
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
