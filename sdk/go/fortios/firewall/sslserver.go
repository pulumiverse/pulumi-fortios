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

// Configure SSL servers.
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
//			_, err := firewall.NewSslserver(ctx, "trname", &firewall.SslserverArgs{
//				AddHeaderXForwardedProto: pulumi.String("enable"),
//				Ip:                       pulumi.String("1.1.1.1"),
//				MappedPort:               pulumi.Int(2234),
//				Port:                     pulumi.Int(32321),
//				SslAlgorithm:             pulumi.String("high"),
//				SslCert:                  pulumi.String("Fortinet_CA_SSL"),
//				SslClientRenegotiation:   pulumi.String("allow"),
//				SslDhBits:                pulumi.String("2048"),
//				SslMaxVersion:            pulumi.String("tls-1.2"),
//				SslMinVersion:            pulumi.String("tls-1.1"),
//				SslMode:                  pulumi.String("half"),
//				SslSendEmptyFrags:        pulumi.String("enable"),
//				UrlRewrite:               pulumi.String("disable"),
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
// Firewall SslServer can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/sslserver:Sslserver labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/sslserver:Sslserver labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Sslserver struct {
	pulumi.CustomResourceState

	// Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
	AddHeaderXForwardedProto pulumi.StringOutput `pulumi:"addHeaderXForwardedProto"`
	// IPv4 address of the SSL server.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Mapped server service port (1 - 65535, default = 80).
	MappedPort pulumi.IntOutput `pulumi:"mappedPort"`
	// Server name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Server service port (1 - 65535, default = 443).
	Port pulumi.IntOutput `pulumi:"port"`
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringOutput `pulumi:"sslAlgorithm"`
	// Name of certificate for SSL connections to this server. On FortiOS versions 6.2.0-7.2.6: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.0-7.4.1: default = "Fortinet_SSL".
	SslCert pulumi.StringOutput `pulumi:"sslCert"`
	// Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
	SslClientRenegotiation pulumi.StringOutput `pulumi:"sslClientRenegotiation"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringOutput `pulumi:"sslDhBits"`
	// Highest SSL/TLS version to negotiate.
	SslMaxVersion pulumi.StringOutput `pulumi:"sslMaxVersion"`
	// Lowest SSL/TLS version to negotiate.
	SslMinVersion pulumi.StringOutput `pulumi:"sslMinVersion"`
	// SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
	SslMode pulumi.StringOutput `pulumi:"sslMode"`
	// Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
	SslSendEmptyFrags pulumi.StringOutput `pulumi:"sslSendEmptyFrags"`
	// Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
	UrlRewrite pulumi.StringOutput `pulumi:"urlRewrite"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSslserver registers a new resource with the given unique name, arguments, and options.
func NewSslserver(ctx *pulumi.Context,
	name string, args *SslserverArgs, opts ...pulumi.ResourceOption) (*Sslserver, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.SslCert == nil {
		return nil, errors.New("invalid value for required argument 'SslCert'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Sslserver
	err := ctx.RegisterResource("fortios:firewall/sslserver:Sslserver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSslserver gets an existing Sslserver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslserver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslserverState, opts ...pulumi.ResourceOption) (*Sslserver, error) {
	var resource Sslserver
	err := ctx.ReadResource("fortios:firewall/sslserver:Sslserver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Sslserver resources.
type sslserverState struct {
	// Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
	AddHeaderXForwardedProto *string `pulumi:"addHeaderXForwardedProto"`
	// IPv4 address of the SSL server.
	Ip *string `pulumi:"ip"`
	// Mapped server service port (1 - 65535, default = 80).
	MappedPort *int `pulumi:"mappedPort"`
	// Server name.
	Name *string `pulumi:"name"`
	// Server service port (1 - 65535, default = 443).
	Port *int `pulumi:"port"`
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm *string `pulumi:"sslAlgorithm"`
	// Name of certificate for SSL connections to this server. On FortiOS versions 6.2.0-7.2.6: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.0-7.4.1: default = "Fortinet_SSL".
	SslCert *string `pulumi:"sslCert"`
	// Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
	SslClientRenegotiation *string `pulumi:"sslClientRenegotiation"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits *string `pulumi:"sslDhBits"`
	// Highest SSL/TLS version to negotiate.
	SslMaxVersion *string `pulumi:"sslMaxVersion"`
	// Lowest SSL/TLS version to negotiate.
	SslMinVersion *string `pulumi:"sslMinVersion"`
	// SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
	SslMode *string `pulumi:"sslMode"`
	// Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
	SslSendEmptyFrags *string `pulumi:"sslSendEmptyFrags"`
	// Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
	UrlRewrite *string `pulumi:"urlRewrite"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SslserverState struct {
	// Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
	AddHeaderXForwardedProto pulumi.StringPtrInput
	// IPv4 address of the SSL server.
	Ip pulumi.StringPtrInput
	// Mapped server service port (1 - 65535, default = 80).
	MappedPort pulumi.IntPtrInput
	// Server name.
	Name pulumi.StringPtrInput
	// Server service port (1 - 65535, default = 443).
	Port pulumi.IntPtrInput
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringPtrInput
	// Name of certificate for SSL connections to this server. On FortiOS versions 6.2.0-7.2.6: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.0-7.4.1: default = "Fortinet_SSL".
	SslCert pulumi.StringPtrInput
	// Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
	SslClientRenegotiation pulumi.StringPtrInput
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringPtrInput
	// Highest SSL/TLS version to negotiate.
	SslMaxVersion pulumi.StringPtrInput
	// Lowest SSL/TLS version to negotiate.
	SslMinVersion pulumi.StringPtrInput
	// SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
	SslMode pulumi.StringPtrInput
	// Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
	SslSendEmptyFrags pulumi.StringPtrInput
	// Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
	UrlRewrite pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SslserverState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslserverState)(nil)).Elem()
}

type sslserverArgs struct {
	// Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
	AddHeaderXForwardedProto *string `pulumi:"addHeaderXForwardedProto"`
	// IPv4 address of the SSL server.
	Ip string `pulumi:"ip"`
	// Mapped server service port (1 - 65535, default = 80).
	MappedPort *int `pulumi:"mappedPort"`
	// Server name.
	Name *string `pulumi:"name"`
	// Server service port (1 - 65535, default = 443).
	Port int `pulumi:"port"`
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm *string `pulumi:"sslAlgorithm"`
	// Name of certificate for SSL connections to this server. On FortiOS versions 6.2.0-7.2.6: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.0-7.4.1: default = "Fortinet_SSL".
	SslCert string `pulumi:"sslCert"`
	// Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
	SslClientRenegotiation *string `pulumi:"sslClientRenegotiation"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits *string `pulumi:"sslDhBits"`
	// Highest SSL/TLS version to negotiate.
	SslMaxVersion *string `pulumi:"sslMaxVersion"`
	// Lowest SSL/TLS version to negotiate.
	SslMinVersion *string `pulumi:"sslMinVersion"`
	// SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
	SslMode *string `pulumi:"sslMode"`
	// Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
	SslSendEmptyFrags *string `pulumi:"sslSendEmptyFrags"`
	// Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
	UrlRewrite *string `pulumi:"urlRewrite"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Sslserver resource.
type SslserverArgs struct {
	// Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
	AddHeaderXForwardedProto pulumi.StringPtrInput
	// IPv4 address of the SSL server.
	Ip pulumi.StringInput
	// Mapped server service port (1 - 65535, default = 80).
	MappedPort pulumi.IntPtrInput
	// Server name.
	Name pulumi.StringPtrInput
	// Server service port (1 - 65535, default = 443).
	Port pulumi.IntInput
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringPtrInput
	// Name of certificate for SSL connections to this server. On FortiOS versions 6.2.0-7.2.6: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.0-7.4.1: default = "Fortinet_SSL".
	SslCert pulumi.StringInput
	// Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
	SslClientRenegotiation pulumi.StringPtrInput
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringPtrInput
	// Highest SSL/TLS version to negotiate.
	SslMaxVersion pulumi.StringPtrInput
	// Lowest SSL/TLS version to negotiate.
	SslMinVersion pulumi.StringPtrInput
	// SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
	SslMode pulumi.StringPtrInput
	// Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
	SslSendEmptyFrags pulumi.StringPtrInput
	// Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
	UrlRewrite pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SslserverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslserverArgs)(nil)).Elem()
}

type SslserverInput interface {
	pulumi.Input

	ToSslserverOutput() SslserverOutput
	ToSslserverOutputWithContext(ctx context.Context) SslserverOutput
}

func (*Sslserver) ElementType() reflect.Type {
	return reflect.TypeOf((**Sslserver)(nil)).Elem()
}

func (i *Sslserver) ToSslserverOutput() SslserverOutput {
	return i.ToSslserverOutputWithContext(context.Background())
}

func (i *Sslserver) ToSslserverOutputWithContext(ctx context.Context) SslserverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslserverOutput)
}

// SslserverArrayInput is an input type that accepts SslserverArray and SslserverArrayOutput values.
// You can construct a concrete instance of `SslserverArrayInput` via:
//
//	SslserverArray{ SslserverArgs{...} }
type SslserverArrayInput interface {
	pulumi.Input

	ToSslserverArrayOutput() SslserverArrayOutput
	ToSslserverArrayOutputWithContext(context.Context) SslserverArrayOutput
}

type SslserverArray []SslserverInput

func (SslserverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sslserver)(nil)).Elem()
}

func (i SslserverArray) ToSslserverArrayOutput() SslserverArrayOutput {
	return i.ToSslserverArrayOutputWithContext(context.Background())
}

func (i SslserverArray) ToSslserverArrayOutputWithContext(ctx context.Context) SslserverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslserverArrayOutput)
}

// SslserverMapInput is an input type that accepts SslserverMap and SslserverMapOutput values.
// You can construct a concrete instance of `SslserverMapInput` via:
//
//	SslserverMap{ "key": SslserverArgs{...} }
type SslserverMapInput interface {
	pulumi.Input

	ToSslserverMapOutput() SslserverMapOutput
	ToSslserverMapOutputWithContext(context.Context) SslserverMapOutput
}

type SslserverMap map[string]SslserverInput

func (SslserverMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sslserver)(nil)).Elem()
}

func (i SslserverMap) ToSslserverMapOutput() SslserverMapOutput {
	return i.ToSslserverMapOutputWithContext(context.Background())
}

func (i SslserverMap) ToSslserverMapOutputWithContext(ctx context.Context) SslserverMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslserverMapOutput)
}

type SslserverOutput struct{ *pulumi.OutputState }

func (SslserverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sslserver)(nil)).Elem()
}

func (o SslserverOutput) ToSslserverOutput() SslserverOutput {
	return o
}

func (o SslserverOutput) ToSslserverOutputWithContext(ctx context.Context) SslserverOutput {
	return o
}

// Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
func (o SslserverOutput) AddHeaderXForwardedProto() pulumi.StringOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.StringOutput { return v.AddHeaderXForwardedProto }).(pulumi.StringOutput)
}

// IPv4 address of the SSL server.
func (o SslserverOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Mapped server service port (1 - 65535, default = 80).
func (o SslserverOutput) MappedPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.IntOutput { return v.MappedPort }).(pulumi.IntOutput)
}

// Server name.
func (o SslserverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Server service port (1 - 65535, default = 443).
func (o SslserverOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
func (o SslserverOutput) SslAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.StringOutput { return v.SslAlgorithm }).(pulumi.StringOutput)
}

// Name of certificate for SSL connections to this server. On FortiOS versions 6.2.0-7.2.6: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.0-7.4.1: default = "Fortinet_SSL".
func (o SslserverOutput) SslCert() pulumi.StringOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.StringOutput { return v.SslCert }).(pulumi.StringOutput)
}

// Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
func (o SslserverOutput) SslClientRenegotiation() pulumi.StringOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.StringOutput { return v.SslClientRenegotiation }).(pulumi.StringOutput)
}

// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
func (o SslserverOutput) SslDhBits() pulumi.StringOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.StringOutput { return v.SslDhBits }).(pulumi.StringOutput)
}

// Highest SSL/TLS version to negotiate.
func (o SslserverOutput) SslMaxVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.StringOutput { return v.SslMaxVersion }).(pulumi.StringOutput)
}

// Lowest SSL/TLS version to negotiate.
func (o SslserverOutput) SslMinVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.StringOutput { return v.SslMinVersion }).(pulumi.StringOutput)
}

// SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
func (o SslserverOutput) SslMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.StringOutput { return v.SslMode }).(pulumi.StringOutput)
}

// Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
func (o SslserverOutput) SslSendEmptyFrags() pulumi.StringOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.StringOutput { return v.SslSendEmptyFrags }).(pulumi.StringOutput)
}

// Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
func (o SslserverOutput) UrlRewrite() pulumi.StringOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.StringOutput { return v.UrlRewrite }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SslserverOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sslserver) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SslserverArrayOutput struct{ *pulumi.OutputState }

func (SslserverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sslserver)(nil)).Elem()
}

func (o SslserverArrayOutput) ToSslserverArrayOutput() SslserverArrayOutput {
	return o
}

func (o SslserverArrayOutput) ToSslserverArrayOutputWithContext(ctx context.Context) SslserverArrayOutput {
	return o
}

func (o SslserverArrayOutput) Index(i pulumi.IntInput) SslserverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Sslserver {
		return vs[0].([]*Sslserver)[vs[1].(int)]
	}).(SslserverOutput)
}

type SslserverMapOutput struct{ *pulumi.OutputState }

func (SslserverMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sslserver)(nil)).Elem()
}

func (o SslserverMapOutput) ToSslserverMapOutput() SslserverMapOutput {
	return o
}

func (o SslserverMapOutput) ToSslserverMapOutputWithContext(ctx context.Context) SslserverMapOutput {
	return o
}

func (o SslserverMapOutput) MapIndex(k pulumi.StringInput) SslserverOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Sslserver {
		return vs[0].(map[string]*Sslserver)[vs[1].(string)]
	}).(SslserverOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SslserverInput)(nil)).Elem(), &Sslserver{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslserverArrayInput)(nil)).Elem(), SslserverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslserverMapInput)(nil)).Elem(), SslserverMap{})
	pulumi.RegisterOutputType(SslserverOutput{})
	pulumi.RegisterOutputType(SslserverArrayOutput{})
	pulumi.RegisterOutputType(SslserverMapOutput{})
}
