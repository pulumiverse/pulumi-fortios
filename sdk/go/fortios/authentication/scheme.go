// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentication

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure Authentication Schemes.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/authentication"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/user"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			trname3, err := user.NewFsso(ctx, "trname3", &user.FssoArgs{
//				Port:      pulumi.Int(8000),
//				Port2:     pulumi.Int(8000),
//				Port3:     pulumi.Int(8000),
//				Port4:     pulumi.Int(8000),
//				Port5:     pulumi.Int(8000),
//				Server:    pulumi.String("1.1.1.1"),
//				SourceIp:  pulumi.String("0.0.0.0"),
//				SourceIp6: pulumi.String("::"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = authentication.NewScheme(ctx, "trname", &authentication.SchemeArgs{
//				FssoAgentForNtlm: trname3.Name,
//				FssoGuest:        pulumi.String("disable"),
//				Method:           pulumi.String("ntlm"),
//				NegotiateNtlm:    pulumi.String("enable"),
//				RequireTfa:       pulumi.String("disable"),
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
// Authentication Scheme can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:authentication/scheme:Scheme labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:authentication/scheme:Scheme labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Scheme struct {
	pulumi.CustomResourceState

	// Domain controller setting.
	DomainController pulumi.StringOutput `pulumi:"domainController"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// FSSO agent to use for NTLM authentication.
	FssoAgentForNtlm pulumi.StringOutput `pulumi:"fssoAgentForNtlm"`
	// Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
	FssoGuest pulumi.StringOutput `pulumi:"fssoGuest"`
	// Kerberos keytab setting.
	KerberosKeytab pulumi.StringOutput `pulumi:"kerberosKeytab"`
	// Authentication methods (default = basic).
	Method pulumi.StringOutput `pulumi:"method"`
	// Authentication scheme name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
	NegotiateNtlm pulumi.StringOutput `pulumi:"negotiateNtlm"`
	// Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
	RequireTfa pulumi.StringOutput `pulumi:"requireTfa"`
	// SAML configuration.
	SamlServer pulumi.StringOutput `pulumi:"samlServer"`
	// SAML authentication timeout in seconds.
	SamlTimeout pulumi.IntOutput `pulumi:"samlTimeout"`
	// SSH CA name.
	SshCa pulumi.StringOutput `pulumi:"sshCa"`
	// Enable/disable authentication with user certificate (default = disable). Valid values: `enable`, `disable`.
	UserCert pulumi.StringOutput `pulumi:"userCert"`
	// Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `userDatabase` block is documented below.
	UserDatabases SchemeUserDatabaseArrayOutput `pulumi:"userDatabases"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewScheme registers a new resource with the given unique name, arguments, and options.
func NewScheme(ctx *pulumi.Context,
	name string, args *SchemeArgs, opts ...pulumi.ResourceOption) (*Scheme, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Method == nil {
		return nil, errors.New("invalid value for required argument 'Method'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Scheme
	err := ctx.RegisterResource("fortios:authentication/scheme:Scheme", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheme gets an existing Scheme resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheme(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemeState, opts ...pulumi.ResourceOption) (*Scheme, error) {
	var resource Scheme
	err := ctx.ReadResource("fortios:authentication/scheme:Scheme", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Scheme resources.
type schemeState struct {
	// Domain controller setting.
	DomainController *string `pulumi:"domainController"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FSSO agent to use for NTLM authentication.
	FssoAgentForNtlm *string `pulumi:"fssoAgentForNtlm"`
	// Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
	FssoGuest *string `pulumi:"fssoGuest"`
	// Kerberos keytab setting.
	KerberosKeytab *string `pulumi:"kerberosKeytab"`
	// Authentication methods (default = basic).
	Method *string `pulumi:"method"`
	// Authentication scheme name.
	Name *string `pulumi:"name"`
	// Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
	NegotiateNtlm *string `pulumi:"negotiateNtlm"`
	// Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
	RequireTfa *string `pulumi:"requireTfa"`
	// SAML configuration.
	SamlServer *string `pulumi:"samlServer"`
	// SAML authentication timeout in seconds.
	SamlTimeout *int `pulumi:"samlTimeout"`
	// SSH CA name.
	SshCa *string `pulumi:"sshCa"`
	// Enable/disable authentication with user certificate (default = disable). Valid values: `enable`, `disable`.
	UserCert *string `pulumi:"userCert"`
	// Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `userDatabase` block is documented below.
	UserDatabases []SchemeUserDatabase `pulumi:"userDatabases"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SchemeState struct {
	// Domain controller setting.
	DomainController pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FSSO agent to use for NTLM authentication.
	FssoAgentForNtlm pulumi.StringPtrInput
	// Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
	FssoGuest pulumi.StringPtrInput
	// Kerberos keytab setting.
	KerberosKeytab pulumi.StringPtrInput
	// Authentication methods (default = basic).
	Method pulumi.StringPtrInput
	// Authentication scheme name.
	Name pulumi.StringPtrInput
	// Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
	NegotiateNtlm pulumi.StringPtrInput
	// Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
	RequireTfa pulumi.StringPtrInput
	// SAML configuration.
	SamlServer pulumi.StringPtrInput
	// SAML authentication timeout in seconds.
	SamlTimeout pulumi.IntPtrInput
	// SSH CA name.
	SshCa pulumi.StringPtrInput
	// Enable/disable authentication with user certificate (default = disable). Valid values: `enable`, `disable`.
	UserCert pulumi.StringPtrInput
	// Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `userDatabase` block is documented below.
	UserDatabases SchemeUserDatabaseArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SchemeState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemeState)(nil)).Elem()
}

type schemeArgs struct {
	// Domain controller setting.
	DomainController *string `pulumi:"domainController"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FSSO agent to use for NTLM authentication.
	FssoAgentForNtlm *string `pulumi:"fssoAgentForNtlm"`
	// Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
	FssoGuest *string `pulumi:"fssoGuest"`
	// Kerberos keytab setting.
	KerberosKeytab *string `pulumi:"kerberosKeytab"`
	// Authentication methods (default = basic).
	Method string `pulumi:"method"`
	// Authentication scheme name.
	Name *string `pulumi:"name"`
	// Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
	NegotiateNtlm *string `pulumi:"negotiateNtlm"`
	// Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
	RequireTfa *string `pulumi:"requireTfa"`
	// SAML configuration.
	SamlServer *string `pulumi:"samlServer"`
	// SAML authentication timeout in seconds.
	SamlTimeout *int `pulumi:"samlTimeout"`
	// SSH CA name.
	SshCa *string `pulumi:"sshCa"`
	// Enable/disable authentication with user certificate (default = disable). Valid values: `enable`, `disable`.
	UserCert *string `pulumi:"userCert"`
	// Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `userDatabase` block is documented below.
	UserDatabases []SchemeUserDatabase `pulumi:"userDatabases"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Scheme resource.
type SchemeArgs struct {
	// Domain controller setting.
	DomainController pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FSSO agent to use for NTLM authentication.
	FssoAgentForNtlm pulumi.StringPtrInput
	// Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
	FssoGuest pulumi.StringPtrInput
	// Kerberos keytab setting.
	KerberosKeytab pulumi.StringPtrInput
	// Authentication methods (default = basic).
	Method pulumi.StringInput
	// Authentication scheme name.
	Name pulumi.StringPtrInput
	// Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
	NegotiateNtlm pulumi.StringPtrInput
	// Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
	RequireTfa pulumi.StringPtrInput
	// SAML configuration.
	SamlServer pulumi.StringPtrInput
	// SAML authentication timeout in seconds.
	SamlTimeout pulumi.IntPtrInput
	// SSH CA name.
	SshCa pulumi.StringPtrInput
	// Enable/disable authentication with user certificate (default = disable). Valid values: `enable`, `disable`.
	UserCert pulumi.StringPtrInput
	// Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `userDatabase` block is documented below.
	UserDatabases SchemeUserDatabaseArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SchemeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemeArgs)(nil)).Elem()
}

type SchemeInput interface {
	pulumi.Input

	ToSchemeOutput() SchemeOutput
	ToSchemeOutputWithContext(ctx context.Context) SchemeOutput
}

func (*Scheme) ElementType() reflect.Type {
	return reflect.TypeOf((**Scheme)(nil)).Elem()
}

func (i *Scheme) ToSchemeOutput() SchemeOutput {
	return i.ToSchemeOutputWithContext(context.Background())
}

func (i *Scheme) ToSchemeOutputWithContext(ctx context.Context) SchemeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemeOutput)
}

// SchemeArrayInput is an input type that accepts SchemeArray and SchemeArrayOutput values.
// You can construct a concrete instance of `SchemeArrayInput` via:
//
//	SchemeArray{ SchemeArgs{...} }
type SchemeArrayInput interface {
	pulumi.Input

	ToSchemeArrayOutput() SchemeArrayOutput
	ToSchemeArrayOutputWithContext(context.Context) SchemeArrayOutput
}

type SchemeArray []SchemeInput

func (SchemeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Scheme)(nil)).Elem()
}

func (i SchemeArray) ToSchemeArrayOutput() SchemeArrayOutput {
	return i.ToSchemeArrayOutputWithContext(context.Background())
}

func (i SchemeArray) ToSchemeArrayOutputWithContext(ctx context.Context) SchemeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemeArrayOutput)
}

// SchemeMapInput is an input type that accepts SchemeMap and SchemeMapOutput values.
// You can construct a concrete instance of `SchemeMapInput` via:
//
//	SchemeMap{ "key": SchemeArgs{...} }
type SchemeMapInput interface {
	pulumi.Input

	ToSchemeMapOutput() SchemeMapOutput
	ToSchemeMapOutputWithContext(context.Context) SchemeMapOutput
}

type SchemeMap map[string]SchemeInput

func (SchemeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Scheme)(nil)).Elem()
}

func (i SchemeMap) ToSchemeMapOutput() SchemeMapOutput {
	return i.ToSchemeMapOutputWithContext(context.Background())
}

func (i SchemeMap) ToSchemeMapOutputWithContext(ctx context.Context) SchemeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemeMapOutput)
}

type SchemeOutput struct{ *pulumi.OutputState }

func (SchemeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Scheme)(nil)).Elem()
}

func (o SchemeOutput) ToSchemeOutput() SchemeOutput {
	return o
}

func (o SchemeOutput) ToSchemeOutputWithContext(ctx context.Context) SchemeOutput {
	return o
}

// Domain controller setting.
func (o SchemeOutput) DomainController() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheme) pulumi.StringOutput { return v.DomainController }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SchemeOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Scheme) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// FSSO agent to use for NTLM authentication.
func (o SchemeOutput) FssoAgentForNtlm() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheme) pulumi.StringOutput { return v.FssoAgentForNtlm }).(pulumi.StringOutput)
}

// Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
func (o SchemeOutput) FssoGuest() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheme) pulumi.StringOutput { return v.FssoGuest }).(pulumi.StringOutput)
}

// Kerberos keytab setting.
func (o SchemeOutput) KerberosKeytab() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheme) pulumi.StringOutput { return v.KerberosKeytab }).(pulumi.StringOutput)
}

// Authentication methods (default = basic).
func (o SchemeOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheme) pulumi.StringOutput { return v.Method }).(pulumi.StringOutput)
}

// Authentication scheme name.
func (o SchemeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheme) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
func (o SchemeOutput) NegotiateNtlm() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheme) pulumi.StringOutput { return v.NegotiateNtlm }).(pulumi.StringOutput)
}

// Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
func (o SchemeOutput) RequireTfa() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheme) pulumi.StringOutput { return v.RequireTfa }).(pulumi.StringOutput)
}

// SAML configuration.
func (o SchemeOutput) SamlServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheme) pulumi.StringOutput { return v.SamlServer }).(pulumi.StringOutput)
}

// SAML authentication timeout in seconds.
func (o SchemeOutput) SamlTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Scheme) pulumi.IntOutput { return v.SamlTimeout }).(pulumi.IntOutput)
}

// SSH CA name.
func (o SchemeOutput) SshCa() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheme) pulumi.StringOutput { return v.SshCa }).(pulumi.StringOutput)
}

// Enable/disable authentication with user certificate (default = disable). Valid values: `enable`, `disable`.
func (o SchemeOutput) UserCert() pulumi.StringOutput {
	return o.ApplyT(func(v *Scheme) pulumi.StringOutput { return v.UserCert }).(pulumi.StringOutput)
}

// Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `userDatabase` block is documented below.
func (o SchemeOutput) UserDatabases() SchemeUserDatabaseArrayOutput {
	return o.ApplyT(func(v *Scheme) SchemeUserDatabaseArrayOutput { return v.UserDatabases }).(SchemeUserDatabaseArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SchemeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Scheme) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SchemeArrayOutput struct{ *pulumi.OutputState }

func (SchemeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Scheme)(nil)).Elem()
}

func (o SchemeArrayOutput) ToSchemeArrayOutput() SchemeArrayOutput {
	return o
}

func (o SchemeArrayOutput) ToSchemeArrayOutputWithContext(ctx context.Context) SchemeArrayOutput {
	return o
}

func (o SchemeArrayOutput) Index(i pulumi.IntInput) SchemeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Scheme {
		return vs[0].([]*Scheme)[vs[1].(int)]
	}).(SchemeOutput)
}

type SchemeMapOutput struct{ *pulumi.OutputState }

func (SchemeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Scheme)(nil)).Elem()
}

func (o SchemeMapOutput) ToSchemeMapOutput() SchemeMapOutput {
	return o
}

func (o SchemeMapOutput) ToSchemeMapOutputWithContext(ctx context.Context) SchemeMapOutput {
	return o
}

func (o SchemeMapOutput) MapIndex(k pulumi.StringInput) SchemeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Scheme {
		return vs[0].(map[string]*Scheme)[vs[1].(string)]
	}).(SchemeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchemeInput)(nil)).Elem(), &Scheme{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchemeArrayInput)(nil)).Elem(), SchemeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchemeMapInput)(nil)).Elem(), SchemeMap{})
	pulumi.RegisterOutputType(SchemeOutput{})
	pulumi.RegisterOutputType(SchemeArrayOutput{})
	pulumi.RegisterOutputType(SchemeMapOutput{})
}
