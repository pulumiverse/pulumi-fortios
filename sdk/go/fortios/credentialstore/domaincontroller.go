// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package credentialstore

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Define known domain controller servers. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,6.4.15,7.0.0`.
//
// ## Import
//
// CredentialStore DomainController can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:credentialstore/domaincontroller:Domaincontroller labelname {{server_name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:credentialstore/domaincontroller:Domaincontroller labelname {{server_name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Domaincontroller struct {
	pulumi.CustomResourceState

	// Fully qualified domain name (FQDN).
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Hostname of the server to connect to.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// IPv4 server address.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// IPv6 server address.
	Ip6 pulumi.StringOutput `pulumi:"ip6"`
	// Password for specified username.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Port number of service. Port number 0 indicates automatic discovery.
	Port pulumi.IntOutput `pulumi:"port"`
	// Name of the server to connect to.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// User name to sign in with. Must have proper permissions for service.
	Username pulumi.StringOutput `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewDomaincontroller registers a new resource with the given unique name, arguments, and options.
func NewDomaincontroller(ctx *pulumi.Context,
	name string, args *DomaincontrollerArgs, opts ...pulumi.ResourceOption) (*Domaincontroller, error) {
	if args == nil {
		args = &DomaincontrollerArgs{}
	}

	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Domaincontroller
	err := ctx.RegisterResource("fortios:credentialstore/domaincontroller:Domaincontroller", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomaincontroller gets an existing Domaincontroller resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomaincontroller(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomaincontrollerState, opts ...pulumi.ResourceOption) (*Domaincontroller, error) {
	var resource Domaincontroller
	err := ctx.ReadResource("fortios:credentialstore/domaincontroller:Domaincontroller", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domaincontroller resources.
type domaincontrollerState struct {
	// Fully qualified domain name (FQDN).
	DomainName *string `pulumi:"domainName"`
	// Hostname of the server to connect to.
	Hostname *string `pulumi:"hostname"`
	// IPv4 server address.
	Ip *string `pulumi:"ip"`
	// IPv6 server address.
	Ip6 *string `pulumi:"ip6"`
	// Password for specified username.
	Password *string `pulumi:"password"`
	// Port number of service. Port number 0 indicates automatic discovery.
	Port *int `pulumi:"port"`
	// Name of the server to connect to.
	ServerName *string `pulumi:"serverName"`
	// User name to sign in with. Must have proper permissions for service.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DomaincontrollerState struct {
	// Fully qualified domain name (FQDN).
	DomainName pulumi.StringPtrInput
	// Hostname of the server to connect to.
	Hostname pulumi.StringPtrInput
	// IPv4 server address.
	Ip pulumi.StringPtrInput
	// IPv6 server address.
	Ip6 pulumi.StringPtrInput
	// Password for specified username.
	Password pulumi.StringPtrInput
	// Port number of service. Port number 0 indicates automatic discovery.
	Port pulumi.IntPtrInput
	// Name of the server to connect to.
	ServerName pulumi.StringPtrInput
	// User name to sign in with. Must have proper permissions for service.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DomaincontrollerState) ElementType() reflect.Type {
	return reflect.TypeOf((*domaincontrollerState)(nil)).Elem()
}

type domaincontrollerArgs struct {
	// Fully qualified domain name (FQDN).
	DomainName *string `pulumi:"domainName"`
	// Hostname of the server to connect to.
	Hostname *string `pulumi:"hostname"`
	// IPv4 server address.
	Ip *string `pulumi:"ip"`
	// IPv6 server address.
	Ip6 *string `pulumi:"ip6"`
	// Password for specified username.
	Password *string `pulumi:"password"`
	// Port number of service. Port number 0 indicates automatic discovery.
	Port *int `pulumi:"port"`
	// Name of the server to connect to.
	ServerName *string `pulumi:"serverName"`
	// User name to sign in with. Must have proper permissions for service.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Domaincontroller resource.
type DomaincontrollerArgs struct {
	// Fully qualified domain name (FQDN).
	DomainName pulumi.StringPtrInput
	// Hostname of the server to connect to.
	Hostname pulumi.StringPtrInput
	// IPv4 server address.
	Ip pulumi.StringPtrInput
	// IPv6 server address.
	Ip6 pulumi.StringPtrInput
	// Password for specified username.
	Password pulumi.StringPtrInput
	// Port number of service. Port number 0 indicates automatic discovery.
	Port pulumi.IntPtrInput
	// Name of the server to connect to.
	ServerName pulumi.StringPtrInput
	// User name to sign in with. Must have proper permissions for service.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DomaincontrollerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domaincontrollerArgs)(nil)).Elem()
}

type DomaincontrollerInput interface {
	pulumi.Input

	ToDomaincontrollerOutput() DomaincontrollerOutput
	ToDomaincontrollerOutputWithContext(ctx context.Context) DomaincontrollerOutput
}

func (*Domaincontroller) ElementType() reflect.Type {
	return reflect.TypeOf((**Domaincontroller)(nil)).Elem()
}

func (i *Domaincontroller) ToDomaincontrollerOutput() DomaincontrollerOutput {
	return i.ToDomaincontrollerOutputWithContext(context.Background())
}

func (i *Domaincontroller) ToDomaincontrollerOutputWithContext(ctx context.Context) DomaincontrollerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomaincontrollerOutput)
}

// DomaincontrollerArrayInput is an input type that accepts DomaincontrollerArray and DomaincontrollerArrayOutput values.
// You can construct a concrete instance of `DomaincontrollerArrayInput` via:
//
//	DomaincontrollerArray{ DomaincontrollerArgs{...} }
type DomaincontrollerArrayInput interface {
	pulumi.Input

	ToDomaincontrollerArrayOutput() DomaincontrollerArrayOutput
	ToDomaincontrollerArrayOutputWithContext(context.Context) DomaincontrollerArrayOutput
}

type DomaincontrollerArray []DomaincontrollerInput

func (DomaincontrollerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domaincontroller)(nil)).Elem()
}

func (i DomaincontrollerArray) ToDomaincontrollerArrayOutput() DomaincontrollerArrayOutput {
	return i.ToDomaincontrollerArrayOutputWithContext(context.Background())
}

func (i DomaincontrollerArray) ToDomaincontrollerArrayOutputWithContext(ctx context.Context) DomaincontrollerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomaincontrollerArrayOutput)
}

// DomaincontrollerMapInput is an input type that accepts DomaincontrollerMap and DomaincontrollerMapOutput values.
// You can construct a concrete instance of `DomaincontrollerMapInput` via:
//
//	DomaincontrollerMap{ "key": DomaincontrollerArgs{...} }
type DomaincontrollerMapInput interface {
	pulumi.Input

	ToDomaincontrollerMapOutput() DomaincontrollerMapOutput
	ToDomaincontrollerMapOutputWithContext(context.Context) DomaincontrollerMapOutput
}

type DomaincontrollerMap map[string]DomaincontrollerInput

func (DomaincontrollerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domaincontroller)(nil)).Elem()
}

func (i DomaincontrollerMap) ToDomaincontrollerMapOutput() DomaincontrollerMapOutput {
	return i.ToDomaincontrollerMapOutputWithContext(context.Background())
}

func (i DomaincontrollerMap) ToDomaincontrollerMapOutputWithContext(ctx context.Context) DomaincontrollerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomaincontrollerMapOutput)
}

type DomaincontrollerOutput struct{ *pulumi.OutputState }

func (DomaincontrollerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domaincontroller)(nil)).Elem()
}

func (o DomaincontrollerOutput) ToDomaincontrollerOutput() DomaincontrollerOutput {
	return o
}

func (o DomaincontrollerOutput) ToDomaincontrollerOutputWithContext(ctx context.Context) DomaincontrollerOutput {
	return o
}

// Fully qualified domain name (FQDN).
func (o DomaincontrollerOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Domaincontroller) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// Hostname of the server to connect to.
func (o DomaincontrollerOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Domaincontroller) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// IPv4 server address.
func (o DomaincontrollerOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Domaincontroller) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// IPv6 server address.
func (o DomaincontrollerOutput) Ip6() pulumi.StringOutput {
	return o.ApplyT(func(v *Domaincontroller) pulumi.StringOutput { return v.Ip6 }).(pulumi.StringOutput)
}

// Password for specified username.
func (o DomaincontrollerOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domaincontroller) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Port number of service. Port number 0 indicates automatic discovery.
func (o DomaincontrollerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Domaincontroller) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Name of the server to connect to.
func (o DomaincontrollerOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *Domaincontroller) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

// User name to sign in with. Must have proper permissions for service.
func (o DomaincontrollerOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Domaincontroller) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DomaincontrollerOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Domaincontroller) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type DomaincontrollerArrayOutput struct{ *pulumi.OutputState }

func (DomaincontrollerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domaincontroller)(nil)).Elem()
}

func (o DomaincontrollerArrayOutput) ToDomaincontrollerArrayOutput() DomaincontrollerArrayOutput {
	return o
}

func (o DomaincontrollerArrayOutput) ToDomaincontrollerArrayOutputWithContext(ctx context.Context) DomaincontrollerArrayOutput {
	return o
}

func (o DomaincontrollerArrayOutput) Index(i pulumi.IntInput) DomaincontrollerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Domaincontroller {
		return vs[0].([]*Domaincontroller)[vs[1].(int)]
	}).(DomaincontrollerOutput)
}

type DomaincontrollerMapOutput struct{ *pulumi.OutputState }

func (DomaincontrollerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domaincontroller)(nil)).Elem()
}

func (o DomaincontrollerMapOutput) ToDomaincontrollerMapOutput() DomaincontrollerMapOutput {
	return o
}

func (o DomaincontrollerMapOutput) ToDomaincontrollerMapOutputWithContext(ctx context.Context) DomaincontrollerMapOutput {
	return o
}

func (o DomaincontrollerMapOutput) MapIndex(k pulumi.StringInput) DomaincontrollerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Domaincontroller {
		return vs[0].(map[string]*Domaincontroller)[vs[1].(string)]
	}).(DomaincontrollerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomaincontrollerInput)(nil)).Elem(), &Domaincontroller{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomaincontrollerArrayInput)(nil)).Elem(), DomaincontrollerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomaincontrollerMapInput)(nil)).Elem(), DomaincontrollerMap{})
	pulumi.RegisterOutputType(DomaincontrollerOutput{})
	pulumi.RegisterOutputType(DomaincontrollerArrayOutput{})
	pulumi.RegisterOutputType(DomaincontrollerMapOutput{})
}
