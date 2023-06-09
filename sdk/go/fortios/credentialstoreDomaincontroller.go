// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define known domain controller servers. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,7.0.0`.
//
// ## Import
//
// # CredentialStore DomainController can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/credentialstoreDomaincontroller:CredentialstoreDomaincontroller labelname {{server_name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/credentialstoreDomaincontroller:CredentialstoreDomaincontroller labelname {{server_name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type CredentialstoreDomaincontroller struct {
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
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewCredentialstoreDomaincontroller registers a new resource with the given unique name, arguments, and options.
func NewCredentialstoreDomaincontroller(ctx *pulumi.Context,
	name string, args *CredentialstoreDomaincontrollerArgs, opts ...pulumi.ResourceOption) (*CredentialstoreDomaincontroller, error) {
	if args == nil {
		args = &CredentialstoreDomaincontrollerArgs{}
	}

	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource CredentialstoreDomaincontroller
	err := ctx.RegisterResource("fortios:index/credentialstoreDomaincontroller:CredentialstoreDomaincontroller", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCredentialstoreDomaincontroller gets an existing CredentialstoreDomaincontroller resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCredentialstoreDomaincontroller(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CredentialstoreDomaincontrollerState, opts ...pulumi.ResourceOption) (*CredentialstoreDomaincontroller, error) {
	var resource CredentialstoreDomaincontroller
	err := ctx.ReadResource("fortios:index/credentialstoreDomaincontroller:CredentialstoreDomaincontroller", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CredentialstoreDomaincontroller resources.
type credentialstoreDomaincontrollerState struct {
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

type CredentialstoreDomaincontrollerState struct {
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

func (CredentialstoreDomaincontrollerState) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialstoreDomaincontrollerState)(nil)).Elem()
}

type credentialstoreDomaincontrollerArgs struct {
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

// The set of arguments for constructing a CredentialstoreDomaincontroller resource.
type CredentialstoreDomaincontrollerArgs struct {
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

func (CredentialstoreDomaincontrollerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialstoreDomaincontrollerArgs)(nil)).Elem()
}

type CredentialstoreDomaincontrollerInput interface {
	pulumi.Input

	ToCredentialstoreDomaincontrollerOutput() CredentialstoreDomaincontrollerOutput
	ToCredentialstoreDomaincontrollerOutputWithContext(ctx context.Context) CredentialstoreDomaincontrollerOutput
}

func (*CredentialstoreDomaincontroller) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialstoreDomaincontroller)(nil)).Elem()
}

func (i *CredentialstoreDomaincontroller) ToCredentialstoreDomaincontrollerOutput() CredentialstoreDomaincontrollerOutput {
	return i.ToCredentialstoreDomaincontrollerOutputWithContext(context.Background())
}

func (i *CredentialstoreDomaincontroller) ToCredentialstoreDomaincontrollerOutputWithContext(ctx context.Context) CredentialstoreDomaincontrollerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialstoreDomaincontrollerOutput)
}

// CredentialstoreDomaincontrollerArrayInput is an input type that accepts CredentialstoreDomaincontrollerArray and CredentialstoreDomaincontrollerArrayOutput values.
// You can construct a concrete instance of `CredentialstoreDomaincontrollerArrayInput` via:
//
//	CredentialstoreDomaincontrollerArray{ CredentialstoreDomaincontrollerArgs{...} }
type CredentialstoreDomaincontrollerArrayInput interface {
	pulumi.Input

	ToCredentialstoreDomaincontrollerArrayOutput() CredentialstoreDomaincontrollerArrayOutput
	ToCredentialstoreDomaincontrollerArrayOutputWithContext(context.Context) CredentialstoreDomaincontrollerArrayOutput
}

type CredentialstoreDomaincontrollerArray []CredentialstoreDomaincontrollerInput

func (CredentialstoreDomaincontrollerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CredentialstoreDomaincontroller)(nil)).Elem()
}

func (i CredentialstoreDomaincontrollerArray) ToCredentialstoreDomaincontrollerArrayOutput() CredentialstoreDomaincontrollerArrayOutput {
	return i.ToCredentialstoreDomaincontrollerArrayOutputWithContext(context.Background())
}

func (i CredentialstoreDomaincontrollerArray) ToCredentialstoreDomaincontrollerArrayOutputWithContext(ctx context.Context) CredentialstoreDomaincontrollerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialstoreDomaincontrollerArrayOutput)
}

// CredentialstoreDomaincontrollerMapInput is an input type that accepts CredentialstoreDomaincontrollerMap and CredentialstoreDomaincontrollerMapOutput values.
// You can construct a concrete instance of `CredentialstoreDomaincontrollerMapInput` via:
//
//	CredentialstoreDomaincontrollerMap{ "key": CredentialstoreDomaincontrollerArgs{...} }
type CredentialstoreDomaincontrollerMapInput interface {
	pulumi.Input

	ToCredentialstoreDomaincontrollerMapOutput() CredentialstoreDomaincontrollerMapOutput
	ToCredentialstoreDomaincontrollerMapOutputWithContext(context.Context) CredentialstoreDomaincontrollerMapOutput
}

type CredentialstoreDomaincontrollerMap map[string]CredentialstoreDomaincontrollerInput

func (CredentialstoreDomaincontrollerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CredentialstoreDomaincontroller)(nil)).Elem()
}

func (i CredentialstoreDomaincontrollerMap) ToCredentialstoreDomaincontrollerMapOutput() CredentialstoreDomaincontrollerMapOutput {
	return i.ToCredentialstoreDomaincontrollerMapOutputWithContext(context.Background())
}

func (i CredentialstoreDomaincontrollerMap) ToCredentialstoreDomaincontrollerMapOutputWithContext(ctx context.Context) CredentialstoreDomaincontrollerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialstoreDomaincontrollerMapOutput)
}

type CredentialstoreDomaincontrollerOutput struct{ *pulumi.OutputState }

func (CredentialstoreDomaincontrollerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialstoreDomaincontroller)(nil)).Elem()
}

func (o CredentialstoreDomaincontrollerOutput) ToCredentialstoreDomaincontrollerOutput() CredentialstoreDomaincontrollerOutput {
	return o
}

func (o CredentialstoreDomaincontrollerOutput) ToCredentialstoreDomaincontrollerOutputWithContext(ctx context.Context) CredentialstoreDomaincontrollerOutput {
	return o
}

// Fully qualified domain name (FQDN).
func (o CredentialstoreDomaincontrollerOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialstoreDomaincontroller) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// Hostname of the server to connect to.
func (o CredentialstoreDomaincontrollerOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialstoreDomaincontroller) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// IPv4 server address.
func (o CredentialstoreDomaincontrollerOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialstoreDomaincontroller) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// IPv6 server address.
func (o CredentialstoreDomaincontrollerOutput) Ip6() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialstoreDomaincontroller) pulumi.StringOutput { return v.Ip6 }).(pulumi.StringOutput)
}

// Password for specified username.
func (o CredentialstoreDomaincontrollerOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CredentialstoreDomaincontroller) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Port number of service. Port number 0 indicates automatic discovery.
func (o CredentialstoreDomaincontrollerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *CredentialstoreDomaincontroller) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Name of the server to connect to.
func (o CredentialstoreDomaincontrollerOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialstoreDomaincontroller) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

// User name to sign in with. Must have proper permissions for service.
func (o CredentialstoreDomaincontrollerOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialstoreDomaincontroller) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o CredentialstoreDomaincontrollerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CredentialstoreDomaincontroller) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type CredentialstoreDomaincontrollerArrayOutput struct{ *pulumi.OutputState }

func (CredentialstoreDomaincontrollerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CredentialstoreDomaincontroller)(nil)).Elem()
}

func (o CredentialstoreDomaincontrollerArrayOutput) ToCredentialstoreDomaincontrollerArrayOutput() CredentialstoreDomaincontrollerArrayOutput {
	return o
}

func (o CredentialstoreDomaincontrollerArrayOutput) ToCredentialstoreDomaincontrollerArrayOutputWithContext(ctx context.Context) CredentialstoreDomaincontrollerArrayOutput {
	return o
}

func (o CredentialstoreDomaincontrollerArrayOutput) Index(i pulumi.IntInput) CredentialstoreDomaincontrollerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CredentialstoreDomaincontroller {
		return vs[0].([]*CredentialstoreDomaincontroller)[vs[1].(int)]
	}).(CredentialstoreDomaincontrollerOutput)
}

type CredentialstoreDomaincontrollerMapOutput struct{ *pulumi.OutputState }

func (CredentialstoreDomaincontrollerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CredentialstoreDomaincontroller)(nil)).Elem()
}

func (o CredentialstoreDomaincontrollerMapOutput) ToCredentialstoreDomaincontrollerMapOutput() CredentialstoreDomaincontrollerMapOutput {
	return o
}

func (o CredentialstoreDomaincontrollerMapOutput) ToCredentialstoreDomaincontrollerMapOutputWithContext(ctx context.Context) CredentialstoreDomaincontrollerMapOutput {
	return o
}

func (o CredentialstoreDomaincontrollerMapOutput) MapIndex(k pulumi.StringInput) CredentialstoreDomaincontrollerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CredentialstoreDomaincontroller {
		return vs[0].(map[string]*CredentialstoreDomaincontroller)[vs[1].(string)]
	}).(CredentialstoreDomaincontrollerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialstoreDomaincontrollerInput)(nil)).Elem(), &CredentialstoreDomaincontroller{})
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialstoreDomaincontrollerArrayInput)(nil)).Elem(), CredentialstoreDomaincontrollerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialstoreDomaincontrollerMapInput)(nil)).Elem(), CredentialstoreDomaincontrollerMap{})
	pulumi.RegisterOutputType(CredentialstoreDomaincontrollerOutput{})
	pulumi.RegisterOutputType(CredentialstoreDomaincontrollerArrayOutput{})
	pulumi.RegisterOutputType(CredentialstoreDomaincontrollerMapOutput{})
}
