// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ftpproxy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure explicit FTP proxy settings.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/ftpproxy"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ftpproxy.NewExplicit(ctx, "trname", &ftpproxy.ExplicitArgs{
//				IncomingIp:       pulumi.String("0.0.0.0"),
//				SecDefaultAction: pulumi.String("deny"),
//				Status:           pulumi.String("disable"),
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
// FtpProxy Explicit can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:ftpproxy/explicit:Explicit labelname FtpProxyExplicit
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:ftpproxy/explicit:Explicit labelname FtpProxyExplicit
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Explicit struct {
	pulumi.CustomResourceState

	// Accept incoming FTP requests from this IP address. An interface must have this IP address.
	IncomingIp pulumi.StringOutput `pulumi:"incomingIp"`
	// Accept incoming FTP requests on one or more ports.
	IncomingPort pulumi.StringOutput `pulumi:"incomingPort"`
	// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
	OutgoingIp pulumi.StringOutput `pulumi:"outgoingIp"`
	// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction pulumi.StringOutput `pulumi:"secDefaultAction"`
	// Determine mode of data session on FTP server side. Valid values: `client`, `passive`.
	ServerDataMode pulumi.StringOutput `pulumi:"serverDataMode"`
	// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
	Ssl pulumi.StringOutput `pulumi:"ssl"`
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringOutput `pulumi:"sslAlgorithm"`
	// Name of certificate for SSL connections to this server. On FortiOS versions 6.2.4-7.4.0: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.1: default = "Fortinet_SSL".
	SslCert pulumi.StringOutput `pulumi:"sslCert"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringOutput `pulumi:"sslDhBits"`
	// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewExplicit registers a new resource with the given unique name, arguments, and options.
func NewExplicit(ctx *pulumi.Context,
	name string, args *ExplicitArgs, opts ...pulumi.ResourceOption) (*Explicit, error) {
	if args == nil {
		args = &ExplicitArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Explicit
	err := ctx.RegisterResource("fortios:ftpproxy/explicit:Explicit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExplicit gets an existing Explicit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExplicit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExplicitState, opts ...pulumi.ResourceOption) (*Explicit, error) {
	var resource Explicit
	err := ctx.ReadResource("fortios:ftpproxy/explicit:Explicit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Explicit resources.
type explicitState struct {
	// Accept incoming FTP requests from this IP address. An interface must have this IP address.
	IncomingIp *string `pulumi:"incomingIp"`
	// Accept incoming FTP requests on one or more ports.
	IncomingPort *string `pulumi:"incomingPort"`
	// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
	OutgoingIp *string `pulumi:"outgoingIp"`
	// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction *string `pulumi:"secDefaultAction"`
	// Determine mode of data session on FTP server side. Valid values: `client`, `passive`.
	ServerDataMode *string `pulumi:"serverDataMode"`
	// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
	Ssl *string `pulumi:"ssl"`
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm *string `pulumi:"sslAlgorithm"`
	// Name of certificate for SSL connections to this server. On FortiOS versions 6.2.4-7.4.0: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.1: default = "Fortinet_SSL".
	SslCert *string `pulumi:"sslCert"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits *string `pulumi:"sslDhBits"`
	// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ExplicitState struct {
	// Accept incoming FTP requests from this IP address. An interface must have this IP address.
	IncomingIp pulumi.StringPtrInput
	// Accept incoming FTP requests on one or more ports.
	IncomingPort pulumi.StringPtrInput
	// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
	OutgoingIp pulumi.StringPtrInput
	// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction pulumi.StringPtrInput
	// Determine mode of data session on FTP server side. Valid values: `client`, `passive`.
	ServerDataMode pulumi.StringPtrInput
	// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
	Ssl pulumi.StringPtrInput
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringPtrInput
	// Name of certificate for SSL connections to this server. On FortiOS versions 6.2.4-7.4.0: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.1: default = "Fortinet_SSL".
	SslCert pulumi.StringPtrInput
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringPtrInput
	// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExplicitState) ElementType() reflect.Type {
	return reflect.TypeOf((*explicitState)(nil)).Elem()
}

type explicitArgs struct {
	// Accept incoming FTP requests from this IP address. An interface must have this IP address.
	IncomingIp *string `pulumi:"incomingIp"`
	// Accept incoming FTP requests on one or more ports.
	IncomingPort *string `pulumi:"incomingPort"`
	// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
	OutgoingIp *string `pulumi:"outgoingIp"`
	// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction *string `pulumi:"secDefaultAction"`
	// Determine mode of data session on FTP server side. Valid values: `client`, `passive`.
	ServerDataMode *string `pulumi:"serverDataMode"`
	// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
	Ssl *string `pulumi:"ssl"`
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm *string `pulumi:"sslAlgorithm"`
	// Name of certificate for SSL connections to this server. On FortiOS versions 6.2.4-7.4.0: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.1: default = "Fortinet_SSL".
	SslCert *string `pulumi:"sslCert"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits *string `pulumi:"sslDhBits"`
	// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Explicit resource.
type ExplicitArgs struct {
	// Accept incoming FTP requests from this IP address. An interface must have this IP address.
	IncomingIp pulumi.StringPtrInput
	// Accept incoming FTP requests on one or more ports.
	IncomingPort pulumi.StringPtrInput
	// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
	OutgoingIp pulumi.StringPtrInput
	// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction pulumi.StringPtrInput
	// Determine mode of data session on FTP server side. Valid values: `client`, `passive`.
	ServerDataMode pulumi.StringPtrInput
	// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
	Ssl pulumi.StringPtrInput
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringPtrInput
	// Name of certificate for SSL connections to this server. On FortiOS versions 6.2.4-7.4.0: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.1: default = "Fortinet_SSL".
	SslCert pulumi.StringPtrInput
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringPtrInput
	// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExplicitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*explicitArgs)(nil)).Elem()
}

type ExplicitInput interface {
	pulumi.Input

	ToExplicitOutput() ExplicitOutput
	ToExplicitOutputWithContext(ctx context.Context) ExplicitOutput
}

func (*Explicit) ElementType() reflect.Type {
	return reflect.TypeOf((**Explicit)(nil)).Elem()
}

func (i *Explicit) ToExplicitOutput() ExplicitOutput {
	return i.ToExplicitOutputWithContext(context.Background())
}

func (i *Explicit) ToExplicitOutputWithContext(ctx context.Context) ExplicitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExplicitOutput)
}

// ExplicitArrayInput is an input type that accepts ExplicitArray and ExplicitArrayOutput values.
// You can construct a concrete instance of `ExplicitArrayInput` via:
//
//	ExplicitArray{ ExplicitArgs{...} }
type ExplicitArrayInput interface {
	pulumi.Input

	ToExplicitArrayOutput() ExplicitArrayOutput
	ToExplicitArrayOutputWithContext(context.Context) ExplicitArrayOutput
}

type ExplicitArray []ExplicitInput

func (ExplicitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Explicit)(nil)).Elem()
}

func (i ExplicitArray) ToExplicitArrayOutput() ExplicitArrayOutput {
	return i.ToExplicitArrayOutputWithContext(context.Background())
}

func (i ExplicitArray) ToExplicitArrayOutputWithContext(ctx context.Context) ExplicitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExplicitArrayOutput)
}

// ExplicitMapInput is an input type that accepts ExplicitMap and ExplicitMapOutput values.
// You can construct a concrete instance of `ExplicitMapInput` via:
//
//	ExplicitMap{ "key": ExplicitArgs{...} }
type ExplicitMapInput interface {
	pulumi.Input

	ToExplicitMapOutput() ExplicitMapOutput
	ToExplicitMapOutputWithContext(context.Context) ExplicitMapOutput
}

type ExplicitMap map[string]ExplicitInput

func (ExplicitMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Explicit)(nil)).Elem()
}

func (i ExplicitMap) ToExplicitMapOutput() ExplicitMapOutput {
	return i.ToExplicitMapOutputWithContext(context.Background())
}

func (i ExplicitMap) ToExplicitMapOutputWithContext(ctx context.Context) ExplicitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExplicitMapOutput)
}

type ExplicitOutput struct{ *pulumi.OutputState }

func (ExplicitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Explicit)(nil)).Elem()
}

func (o ExplicitOutput) ToExplicitOutput() ExplicitOutput {
	return o
}

func (o ExplicitOutput) ToExplicitOutputWithContext(ctx context.Context) ExplicitOutput {
	return o
}

// Accept incoming FTP requests from this IP address. An interface must have this IP address.
func (o ExplicitOutput) IncomingIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Explicit) pulumi.StringOutput { return v.IncomingIp }).(pulumi.StringOutput)
}

// Accept incoming FTP requests on one or more ports.
func (o ExplicitOutput) IncomingPort() pulumi.StringOutput {
	return o.ApplyT(func(v *Explicit) pulumi.StringOutput { return v.IncomingPort }).(pulumi.StringOutput)
}

// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
func (o ExplicitOutput) OutgoingIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Explicit) pulumi.StringOutput { return v.OutgoingIp }).(pulumi.StringOutput)
}

// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
func (o ExplicitOutput) SecDefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v *Explicit) pulumi.StringOutput { return v.SecDefaultAction }).(pulumi.StringOutput)
}

// Determine mode of data session on FTP server side. Valid values: `client`, `passive`.
func (o ExplicitOutput) ServerDataMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Explicit) pulumi.StringOutput { return v.ServerDataMode }).(pulumi.StringOutput)
}

// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
func (o ExplicitOutput) Ssl() pulumi.StringOutput {
	return o.ApplyT(func(v *Explicit) pulumi.StringOutput { return v.Ssl }).(pulumi.StringOutput)
}

// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
func (o ExplicitOutput) SslAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *Explicit) pulumi.StringOutput { return v.SslAlgorithm }).(pulumi.StringOutput)
}

// Name of certificate for SSL connections to this server. On FortiOS versions 6.2.4-7.4.0: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.1: default = "Fortinet_SSL".
func (o ExplicitOutput) SslCert() pulumi.StringOutput {
	return o.ApplyT(func(v *Explicit) pulumi.StringOutput { return v.SslCert }).(pulumi.StringOutput)
}

// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
func (o ExplicitOutput) SslDhBits() pulumi.StringOutput {
	return o.ApplyT(func(v *Explicit) pulumi.StringOutput { return v.SslDhBits }).(pulumi.StringOutput)
}

// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
func (o ExplicitOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Explicit) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ExplicitOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Explicit) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ExplicitArrayOutput struct{ *pulumi.OutputState }

func (ExplicitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Explicit)(nil)).Elem()
}

func (o ExplicitArrayOutput) ToExplicitArrayOutput() ExplicitArrayOutput {
	return o
}

func (o ExplicitArrayOutput) ToExplicitArrayOutputWithContext(ctx context.Context) ExplicitArrayOutput {
	return o
}

func (o ExplicitArrayOutput) Index(i pulumi.IntInput) ExplicitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Explicit {
		return vs[0].([]*Explicit)[vs[1].(int)]
	}).(ExplicitOutput)
}

type ExplicitMapOutput struct{ *pulumi.OutputState }

func (ExplicitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Explicit)(nil)).Elem()
}

func (o ExplicitMapOutput) ToExplicitMapOutput() ExplicitMapOutput {
	return o
}

func (o ExplicitMapOutput) ToExplicitMapOutputWithContext(ctx context.Context) ExplicitMapOutput {
	return o
}

func (o ExplicitMapOutput) MapIndex(k pulumi.StringInput) ExplicitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Explicit {
		return vs[0].(map[string]*Explicit)[vs[1].(string)]
	}).(ExplicitOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExplicitInput)(nil)).Elem(), &Explicit{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExplicitArrayInput)(nil)).Elem(), ExplicitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExplicitMapInput)(nil)).Elem(), ExplicitMap{})
	pulumi.RegisterOutputType(ExplicitOutput{})
	pulumi.RegisterOutputType(ExplicitArrayOutput{})
	pulumi.RegisterOutputType(ExplicitMapOutput{})
}
