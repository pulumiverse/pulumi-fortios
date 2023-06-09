// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Kerberos keytab entries.
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
//			trname2, err := fortios.NewUserLdap(ctx, "trname2", &fortios.UserLdapArgs{
//				AccountKeyFilter:      pulumi.String("(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))"),
//				AccountKeyProcessing:  pulumi.String("same"),
//				Cnid:                  pulumi.String("cn"),
//				Dn:                    pulumi.String("EIWNCIEW"),
//				GroupMemberCheck:      pulumi.String("user-attr"),
//				GroupObjectFilter:     pulumi.String("(&(objectcategory=group)(member=*))"),
//				MemberAttr:            pulumi.String("memberOf"),
//				PasswordExpiryWarning: pulumi.String("disable"),
//				PasswordRenewal:       pulumi.String("disable"),
//				Port:                  pulumi.Int(389),
//				Secure:                pulumi.String("disable"),
//				Server:                pulumi.String("1.1.1.1"),
//				ServerIdentityCheck:   pulumi.String("disable"),
//				SourceIp:              pulumi.String("0.0.0.0"),
//				SslMinProtoVersion:    pulumi.String("default"),
//				Type:                  pulumi.String("simple"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = fortios.NewUserKrbkeytab(ctx, "trname", &fortios.UserKrbkeytabArgs{
//				Keytab:     pulumi.String("ZXdlY2VxcmVxd3Jld3E="),
//				LdapServer: trname2.Name,
//				Principal:  pulumi.String("testprin"),
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
// # User KrbKeytab can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/userKrbkeytab:UserKrbkeytab labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/userKrbkeytab:UserKrbkeytab labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type UserKrbkeytab struct {
	pulumi.CustomResourceState

	// base64 coded keytab file containing a pre-shared key.
	Keytab pulumi.StringOutput `pulumi:"keytab"`
	// LDAP server name.
	LdapServer pulumi.StringOutput `pulumi:"ldapServer"`
	// Kerberos keytab entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
	PacData pulumi.StringOutput `pulumi:"pacData"`
	// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUserKrbkeytab registers a new resource with the given unique name, arguments, and options.
func NewUserKrbkeytab(ctx *pulumi.Context,
	name string, args *UserKrbkeytabArgs, opts ...pulumi.ResourceOption) (*UserKrbkeytab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Keytab == nil {
		return nil, errors.New("invalid value for required argument 'Keytab'")
	}
	if args.LdapServer == nil {
		return nil, errors.New("invalid value for required argument 'LdapServer'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	if args.Keytab != nil {
		args.Keytab = pulumi.ToSecret(args.Keytab).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"keytab",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource UserKrbkeytab
	err := ctx.RegisterResource("fortios:index/userKrbkeytab:UserKrbkeytab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserKrbkeytab gets an existing UserKrbkeytab resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserKrbkeytab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserKrbkeytabState, opts ...pulumi.ResourceOption) (*UserKrbkeytab, error) {
	var resource UserKrbkeytab
	err := ctx.ReadResource("fortios:index/userKrbkeytab:UserKrbkeytab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserKrbkeytab resources.
type userKrbkeytabState struct {
	// base64 coded keytab file containing a pre-shared key.
	Keytab *string `pulumi:"keytab"`
	// LDAP server name.
	LdapServer *string `pulumi:"ldapServer"`
	// Kerberos keytab entry name.
	Name *string `pulumi:"name"`
	// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
	PacData *string `pulumi:"pacData"`
	// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
	Principal *string `pulumi:"principal"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UserKrbkeytabState struct {
	// base64 coded keytab file containing a pre-shared key.
	Keytab pulumi.StringPtrInput
	// LDAP server name.
	LdapServer pulumi.StringPtrInput
	// Kerberos keytab entry name.
	Name pulumi.StringPtrInput
	// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
	PacData pulumi.StringPtrInput
	// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
	Principal pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserKrbkeytabState) ElementType() reflect.Type {
	return reflect.TypeOf((*userKrbkeytabState)(nil)).Elem()
}

type userKrbkeytabArgs struct {
	// base64 coded keytab file containing a pre-shared key.
	Keytab string `pulumi:"keytab"`
	// LDAP server name.
	LdapServer string `pulumi:"ldapServer"`
	// Kerberos keytab entry name.
	Name *string `pulumi:"name"`
	// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
	PacData *string `pulumi:"pacData"`
	// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
	Principal string `pulumi:"principal"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserKrbkeytab resource.
type UserKrbkeytabArgs struct {
	// base64 coded keytab file containing a pre-shared key.
	Keytab pulumi.StringInput
	// LDAP server name.
	LdapServer pulumi.StringInput
	// Kerberos keytab entry name.
	Name pulumi.StringPtrInput
	// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
	PacData pulumi.StringPtrInput
	// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
	Principal pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserKrbkeytabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userKrbkeytabArgs)(nil)).Elem()
}

type UserKrbkeytabInput interface {
	pulumi.Input

	ToUserKrbkeytabOutput() UserKrbkeytabOutput
	ToUserKrbkeytabOutputWithContext(ctx context.Context) UserKrbkeytabOutput
}

func (*UserKrbkeytab) ElementType() reflect.Type {
	return reflect.TypeOf((**UserKrbkeytab)(nil)).Elem()
}

func (i *UserKrbkeytab) ToUserKrbkeytabOutput() UserKrbkeytabOutput {
	return i.ToUserKrbkeytabOutputWithContext(context.Background())
}

func (i *UserKrbkeytab) ToUserKrbkeytabOutputWithContext(ctx context.Context) UserKrbkeytabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserKrbkeytabOutput)
}

// UserKrbkeytabArrayInput is an input type that accepts UserKrbkeytabArray and UserKrbkeytabArrayOutput values.
// You can construct a concrete instance of `UserKrbkeytabArrayInput` via:
//
//	UserKrbkeytabArray{ UserKrbkeytabArgs{...} }
type UserKrbkeytabArrayInput interface {
	pulumi.Input

	ToUserKrbkeytabArrayOutput() UserKrbkeytabArrayOutput
	ToUserKrbkeytabArrayOutputWithContext(context.Context) UserKrbkeytabArrayOutput
}

type UserKrbkeytabArray []UserKrbkeytabInput

func (UserKrbkeytabArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserKrbkeytab)(nil)).Elem()
}

func (i UserKrbkeytabArray) ToUserKrbkeytabArrayOutput() UserKrbkeytabArrayOutput {
	return i.ToUserKrbkeytabArrayOutputWithContext(context.Background())
}

func (i UserKrbkeytabArray) ToUserKrbkeytabArrayOutputWithContext(ctx context.Context) UserKrbkeytabArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserKrbkeytabArrayOutput)
}

// UserKrbkeytabMapInput is an input type that accepts UserKrbkeytabMap and UserKrbkeytabMapOutput values.
// You can construct a concrete instance of `UserKrbkeytabMapInput` via:
//
//	UserKrbkeytabMap{ "key": UserKrbkeytabArgs{...} }
type UserKrbkeytabMapInput interface {
	pulumi.Input

	ToUserKrbkeytabMapOutput() UserKrbkeytabMapOutput
	ToUserKrbkeytabMapOutputWithContext(context.Context) UserKrbkeytabMapOutput
}

type UserKrbkeytabMap map[string]UserKrbkeytabInput

func (UserKrbkeytabMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserKrbkeytab)(nil)).Elem()
}

func (i UserKrbkeytabMap) ToUserKrbkeytabMapOutput() UserKrbkeytabMapOutput {
	return i.ToUserKrbkeytabMapOutputWithContext(context.Background())
}

func (i UserKrbkeytabMap) ToUserKrbkeytabMapOutputWithContext(ctx context.Context) UserKrbkeytabMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserKrbkeytabMapOutput)
}

type UserKrbkeytabOutput struct{ *pulumi.OutputState }

func (UserKrbkeytabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserKrbkeytab)(nil)).Elem()
}

func (o UserKrbkeytabOutput) ToUserKrbkeytabOutput() UserKrbkeytabOutput {
	return o
}

func (o UserKrbkeytabOutput) ToUserKrbkeytabOutputWithContext(ctx context.Context) UserKrbkeytabOutput {
	return o
}

// base64 coded keytab file containing a pre-shared key.
func (o UserKrbkeytabOutput) Keytab() pulumi.StringOutput {
	return o.ApplyT(func(v *UserKrbkeytab) pulumi.StringOutput { return v.Keytab }).(pulumi.StringOutput)
}

// LDAP server name.
func (o UserKrbkeytabOutput) LdapServer() pulumi.StringOutput {
	return o.ApplyT(func(v *UserKrbkeytab) pulumi.StringOutput { return v.LdapServer }).(pulumi.StringOutput)
}

// Kerberos keytab entry name.
func (o UserKrbkeytabOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserKrbkeytab) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
func (o UserKrbkeytabOutput) PacData() pulumi.StringOutput {
	return o.ApplyT(func(v *UserKrbkeytab) pulumi.StringOutput { return v.PacData }).(pulumi.StringOutput)
}

// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
func (o UserKrbkeytabOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *UserKrbkeytab) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o UserKrbkeytabOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserKrbkeytab) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type UserKrbkeytabArrayOutput struct{ *pulumi.OutputState }

func (UserKrbkeytabArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserKrbkeytab)(nil)).Elem()
}

func (o UserKrbkeytabArrayOutput) ToUserKrbkeytabArrayOutput() UserKrbkeytabArrayOutput {
	return o
}

func (o UserKrbkeytabArrayOutput) ToUserKrbkeytabArrayOutputWithContext(ctx context.Context) UserKrbkeytabArrayOutput {
	return o
}

func (o UserKrbkeytabArrayOutput) Index(i pulumi.IntInput) UserKrbkeytabOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserKrbkeytab {
		return vs[0].([]*UserKrbkeytab)[vs[1].(int)]
	}).(UserKrbkeytabOutput)
}

type UserKrbkeytabMapOutput struct{ *pulumi.OutputState }

func (UserKrbkeytabMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserKrbkeytab)(nil)).Elem()
}

func (o UserKrbkeytabMapOutput) ToUserKrbkeytabMapOutput() UserKrbkeytabMapOutput {
	return o
}

func (o UserKrbkeytabMapOutput) ToUserKrbkeytabMapOutputWithContext(ctx context.Context) UserKrbkeytabMapOutput {
	return o
}

func (o UserKrbkeytabMapOutput) MapIndex(k pulumi.StringInput) UserKrbkeytabOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserKrbkeytab {
		return vs[0].(map[string]*UserKrbkeytab)[vs[1].(string)]
	}).(UserKrbkeytabOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserKrbkeytabInput)(nil)).Elem(), &UserKrbkeytab{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserKrbkeytabArrayInput)(nil)).Elem(), UserKrbkeytabArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserKrbkeytabMapInput)(nil)).Elem(), UserKrbkeytabMap{})
	pulumi.RegisterOutputType(UserKrbkeytabOutput{})
	pulumi.RegisterOutputType(UserKrbkeytabArrayOutput{})
	pulumi.RegisterOutputType(UserKrbkeytabMapOutput{})
}
