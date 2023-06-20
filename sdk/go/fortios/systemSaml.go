// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Global settings for SAML authentication. Applies to FortiOS Version `>= 6.2.4`.
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
//			_, err := fortios.NewSystemSaml(ctx, "trname", &fortios.SystemSamlArgs{
//				DefaultLoginPage: pulumi.String("normal"),
//				DefaultProfile:   pulumi.String("admin_no_access"),
//				Life:             pulumi.Int(30),
//				Role:             pulumi.String("service-provider"),
//				Status:           pulumi.String("disable"),
//				Tolerance:        pulumi.Int(5),
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
// # System Saml can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemSaml:SystemSaml labelname SystemSaml
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemSaml:SystemSaml labelname SystemSaml
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemSaml struct {
	pulumi.CustomResourceState

	// IdP Binding protocol. Valid values: `post`, `redirect`.
	BindingProtocol pulumi.StringOutput `pulumi:"bindingProtocol"`
	// Certificate to sign SAML messages.
	Cert pulumi.StringOutput `pulumi:"cert"`
	// Choose default login page. Valid values: `normal`, `sso`.
	DefaultLoginPage pulumi.StringOutput `pulumi:"defaultLoginPage"`
	// Default profile for new SSO admin.
	DefaultProfile pulumi.StringOutput `pulumi:"defaultProfile"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// SP entity ID.
	EntityId pulumi.StringOutput `pulumi:"entityId"`
	// IDP certificate name.
	IdpCert pulumi.StringOutput `pulumi:"idpCert"`
	// IDP entity ID.
	IdpEntityId pulumi.StringOutput `pulumi:"idpEntityId"`
	// IDP single logout URL.
	IdpSingleLogoutUrl pulumi.StringOutput `pulumi:"idpSingleLogoutUrl"`
	// IDP single sign-on URL.
	IdpSingleSignOnUrl pulumi.StringOutput `pulumi:"idpSingleSignOnUrl"`
	// Length of the range of time when the assertion is valid (in minutes).
	Life pulumi.IntOutput `pulumi:"life"`
	// SP portal URL.
	PortalUrl pulumi.StringOutput `pulumi:"portalUrl"`
	// SAML role. Valid values: `identity-provider`, `service-provider`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Server address.
	ServerAddress pulumi.StringOutput `pulumi:"serverAddress"`
	// Authorized service providers. The structure of `serviceProviders` block is documented below.
	ServiceProviders SystemSamlServiceProviderArrayOutput `pulumi:"serviceProviders"`
	// SP single logout URL.
	SingleLogoutUrl pulumi.StringOutput `pulumi:"singleLogoutUrl"`
	// SP single sign-on URL.
	SingleSignOnUrl pulumi.StringOutput `pulumi:"singleSignOnUrl"`
	// Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Tolerance to the range of time when the assertion is valid (in minutes).
	Tolerance pulumi.IntOutput `pulumi:"tolerance"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemSaml registers a new resource with the given unique name, arguments, and options.
func NewSystemSaml(ctx *pulumi.Context,
	name string, args *SystemSamlArgs, opts ...pulumi.ResourceOption) (*SystemSaml, error) {
	if args == nil {
		args = &SystemSamlArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSaml
	err := ctx.RegisterResource("fortios:index/systemSaml:SystemSaml", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSaml gets an existing SystemSaml resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSaml(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSamlState, opts ...pulumi.ResourceOption) (*SystemSaml, error) {
	var resource SystemSaml
	err := ctx.ReadResource("fortios:index/systemSaml:SystemSaml", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSaml resources.
type systemSamlState struct {
	// IdP Binding protocol. Valid values: `post`, `redirect`.
	BindingProtocol *string `pulumi:"bindingProtocol"`
	// Certificate to sign SAML messages.
	Cert *string `pulumi:"cert"`
	// Choose default login page. Valid values: `normal`, `sso`.
	DefaultLoginPage *string `pulumi:"defaultLoginPage"`
	// Default profile for new SSO admin.
	DefaultProfile *string `pulumi:"defaultProfile"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// SP entity ID.
	EntityId *string `pulumi:"entityId"`
	// IDP certificate name.
	IdpCert *string `pulumi:"idpCert"`
	// IDP entity ID.
	IdpEntityId *string `pulumi:"idpEntityId"`
	// IDP single logout URL.
	IdpSingleLogoutUrl *string `pulumi:"idpSingleLogoutUrl"`
	// IDP single sign-on URL.
	IdpSingleSignOnUrl *string `pulumi:"idpSingleSignOnUrl"`
	// Length of the range of time when the assertion is valid (in minutes).
	Life *int `pulumi:"life"`
	// SP portal URL.
	PortalUrl *string `pulumi:"portalUrl"`
	// SAML role. Valid values: `identity-provider`, `service-provider`.
	Role *string `pulumi:"role"`
	// Server address.
	ServerAddress *string `pulumi:"serverAddress"`
	// Authorized service providers. The structure of `serviceProviders` block is documented below.
	ServiceProviders []SystemSamlServiceProvider `pulumi:"serviceProviders"`
	// SP single logout URL.
	SingleLogoutUrl *string `pulumi:"singleLogoutUrl"`
	// SP single sign-on URL.
	SingleSignOnUrl *string `pulumi:"singleSignOnUrl"`
	// Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Tolerance to the range of time when the assertion is valid (in minutes).
	Tolerance *int `pulumi:"tolerance"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemSamlState struct {
	// IdP Binding protocol. Valid values: `post`, `redirect`.
	BindingProtocol pulumi.StringPtrInput
	// Certificate to sign SAML messages.
	Cert pulumi.StringPtrInput
	// Choose default login page. Valid values: `normal`, `sso`.
	DefaultLoginPage pulumi.StringPtrInput
	// Default profile for new SSO admin.
	DefaultProfile pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// SP entity ID.
	EntityId pulumi.StringPtrInput
	// IDP certificate name.
	IdpCert pulumi.StringPtrInput
	// IDP entity ID.
	IdpEntityId pulumi.StringPtrInput
	// IDP single logout URL.
	IdpSingleLogoutUrl pulumi.StringPtrInput
	// IDP single sign-on URL.
	IdpSingleSignOnUrl pulumi.StringPtrInput
	// Length of the range of time when the assertion is valid (in minutes).
	Life pulumi.IntPtrInput
	// SP portal URL.
	PortalUrl pulumi.StringPtrInput
	// SAML role. Valid values: `identity-provider`, `service-provider`.
	Role pulumi.StringPtrInput
	// Server address.
	ServerAddress pulumi.StringPtrInput
	// Authorized service providers. The structure of `serviceProviders` block is documented below.
	ServiceProviders SystemSamlServiceProviderArrayInput
	// SP single logout URL.
	SingleLogoutUrl pulumi.StringPtrInput
	// SP single sign-on URL.
	SingleSignOnUrl pulumi.StringPtrInput
	// Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Tolerance to the range of time when the assertion is valid (in minutes).
	Tolerance pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSamlState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSamlState)(nil)).Elem()
}

type systemSamlArgs struct {
	// IdP Binding protocol. Valid values: `post`, `redirect`.
	BindingProtocol *string `pulumi:"bindingProtocol"`
	// Certificate to sign SAML messages.
	Cert *string `pulumi:"cert"`
	// Choose default login page. Valid values: `normal`, `sso`.
	DefaultLoginPage *string `pulumi:"defaultLoginPage"`
	// Default profile for new SSO admin.
	DefaultProfile *string `pulumi:"defaultProfile"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// SP entity ID.
	EntityId *string `pulumi:"entityId"`
	// IDP certificate name.
	IdpCert *string `pulumi:"idpCert"`
	// IDP entity ID.
	IdpEntityId *string `pulumi:"idpEntityId"`
	// IDP single logout URL.
	IdpSingleLogoutUrl *string `pulumi:"idpSingleLogoutUrl"`
	// IDP single sign-on URL.
	IdpSingleSignOnUrl *string `pulumi:"idpSingleSignOnUrl"`
	// Length of the range of time when the assertion is valid (in minutes).
	Life *int `pulumi:"life"`
	// SP portal URL.
	PortalUrl *string `pulumi:"portalUrl"`
	// SAML role. Valid values: `identity-provider`, `service-provider`.
	Role *string `pulumi:"role"`
	// Server address.
	ServerAddress *string `pulumi:"serverAddress"`
	// Authorized service providers. The structure of `serviceProviders` block is documented below.
	ServiceProviders []SystemSamlServiceProvider `pulumi:"serviceProviders"`
	// SP single logout URL.
	SingleLogoutUrl *string `pulumi:"singleLogoutUrl"`
	// SP single sign-on URL.
	SingleSignOnUrl *string `pulumi:"singleSignOnUrl"`
	// Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Tolerance to the range of time when the assertion is valid (in minutes).
	Tolerance *int `pulumi:"tolerance"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemSaml resource.
type SystemSamlArgs struct {
	// IdP Binding protocol. Valid values: `post`, `redirect`.
	BindingProtocol pulumi.StringPtrInput
	// Certificate to sign SAML messages.
	Cert pulumi.StringPtrInput
	// Choose default login page. Valid values: `normal`, `sso`.
	DefaultLoginPage pulumi.StringPtrInput
	// Default profile for new SSO admin.
	DefaultProfile pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// SP entity ID.
	EntityId pulumi.StringPtrInput
	// IDP certificate name.
	IdpCert pulumi.StringPtrInput
	// IDP entity ID.
	IdpEntityId pulumi.StringPtrInput
	// IDP single logout URL.
	IdpSingleLogoutUrl pulumi.StringPtrInput
	// IDP single sign-on URL.
	IdpSingleSignOnUrl pulumi.StringPtrInput
	// Length of the range of time when the assertion is valid (in minutes).
	Life pulumi.IntPtrInput
	// SP portal URL.
	PortalUrl pulumi.StringPtrInput
	// SAML role. Valid values: `identity-provider`, `service-provider`.
	Role pulumi.StringPtrInput
	// Server address.
	ServerAddress pulumi.StringPtrInput
	// Authorized service providers. The structure of `serviceProviders` block is documented below.
	ServiceProviders SystemSamlServiceProviderArrayInput
	// SP single logout URL.
	SingleLogoutUrl pulumi.StringPtrInput
	// SP single sign-on URL.
	SingleSignOnUrl pulumi.StringPtrInput
	// Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Tolerance to the range of time when the assertion is valid (in minutes).
	Tolerance pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSamlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSamlArgs)(nil)).Elem()
}

type SystemSamlInput interface {
	pulumi.Input

	ToSystemSamlOutput() SystemSamlOutput
	ToSystemSamlOutputWithContext(ctx context.Context) SystemSamlOutput
}

func (*SystemSaml) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSaml)(nil)).Elem()
}

func (i *SystemSaml) ToSystemSamlOutput() SystemSamlOutput {
	return i.ToSystemSamlOutputWithContext(context.Background())
}

func (i *SystemSaml) ToSystemSamlOutputWithContext(ctx context.Context) SystemSamlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSamlOutput)
}

// SystemSamlArrayInput is an input type that accepts SystemSamlArray and SystemSamlArrayOutput values.
// You can construct a concrete instance of `SystemSamlArrayInput` via:
//
//	SystemSamlArray{ SystemSamlArgs{...} }
type SystemSamlArrayInput interface {
	pulumi.Input

	ToSystemSamlArrayOutput() SystemSamlArrayOutput
	ToSystemSamlArrayOutputWithContext(context.Context) SystemSamlArrayOutput
}

type SystemSamlArray []SystemSamlInput

func (SystemSamlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSaml)(nil)).Elem()
}

func (i SystemSamlArray) ToSystemSamlArrayOutput() SystemSamlArrayOutput {
	return i.ToSystemSamlArrayOutputWithContext(context.Background())
}

func (i SystemSamlArray) ToSystemSamlArrayOutputWithContext(ctx context.Context) SystemSamlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSamlArrayOutput)
}

// SystemSamlMapInput is an input type that accepts SystemSamlMap and SystemSamlMapOutput values.
// You can construct a concrete instance of `SystemSamlMapInput` via:
//
//	SystemSamlMap{ "key": SystemSamlArgs{...} }
type SystemSamlMapInput interface {
	pulumi.Input

	ToSystemSamlMapOutput() SystemSamlMapOutput
	ToSystemSamlMapOutputWithContext(context.Context) SystemSamlMapOutput
}

type SystemSamlMap map[string]SystemSamlInput

func (SystemSamlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSaml)(nil)).Elem()
}

func (i SystemSamlMap) ToSystemSamlMapOutput() SystemSamlMapOutput {
	return i.ToSystemSamlMapOutputWithContext(context.Background())
}

func (i SystemSamlMap) ToSystemSamlMapOutputWithContext(ctx context.Context) SystemSamlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSamlMapOutput)
}

type SystemSamlOutput struct{ *pulumi.OutputState }

func (SystemSamlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSaml)(nil)).Elem()
}

func (o SystemSamlOutput) ToSystemSamlOutput() SystemSamlOutput {
	return o
}

func (o SystemSamlOutput) ToSystemSamlOutputWithContext(ctx context.Context) SystemSamlOutput {
	return o
}

// IdP Binding protocol. Valid values: `post`, `redirect`.
func (o SystemSamlOutput) BindingProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.BindingProtocol }).(pulumi.StringOutput)
}

// Certificate to sign SAML messages.
func (o SystemSamlOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.Cert }).(pulumi.StringOutput)
}

// Choose default login page. Valid values: `normal`, `sso`.
func (o SystemSamlOutput) DefaultLoginPage() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.DefaultLoginPage }).(pulumi.StringOutput)
}

// Default profile for new SSO admin.
func (o SystemSamlOutput) DefaultProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.DefaultProfile }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SystemSamlOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// SP entity ID.
func (o SystemSamlOutput) EntityId() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.EntityId }).(pulumi.StringOutput)
}

// IDP certificate name.
func (o SystemSamlOutput) IdpCert() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.IdpCert }).(pulumi.StringOutput)
}

// IDP entity ID.
func (o SystemSamlOutput) IdpEntityId() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.IdpEntityId }).(pulumi.StringOutput)
}

// IDP single logout URL.
func (o SystemSamlOutput) IdpSingleLogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.IdpSingleLogoutUrl }).(pulumi.StringOutput)
}

// IDP single sign-on URL.
func (o SystemSamlOutput) IdpSingleSignOnUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.IdpSingleSignOnUrl }).(pulumi.StringOutput)
}

// Length of the range of time when the assertion is valid (in minutes).
func (o SystemSamlOutput) Life() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.IntOutput { return v.Life }).(pulumi.IntOutput)
}

// SP portal URL.
func (o SystemSamlOutput) PortalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.PortalUrl }).(pulumi.StringOutput)
}

// SAML role. Valid values: `identity-provider`, `service-provider`.
func (o SystemSamlOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Server address.
func (o SystemSamlOutput) ServerAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.ServerAddress }).(pulumi.StringOutput)
}

// Authorized service providers. The structure of `serviceProviders` block is documented below.
func (o SystemSamlOutput) ServiceProviders() SystemSamlServiceProviderArrayOutput {
	return o.ApplyT(func(v *SystemSaml) SystemSamlServiceProviderArrayOutput { return v.ServiceProviders }).(SystemSamlServiceProviderArrayOutput)
}

// SP single logout URL.
func (o SystemSamlOutput) SingleLogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.SingleLogoutUrl }).(pulumi.StringOutput)
}

// SP single sign-on URL.
func (o SystemSamlOutput) SingleSignOnUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.SingleSignOnUrl }).(pulumi.StringOutput)
}

// Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
func (o SystemSamlOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tolerance to the range of time when the assertion is valid (in minutes).
func (o SystemSamlOutput) Tolerance() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.IntOutput { return v.Tolerance }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemSamlOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemSaml) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemSamlArrayOutput struct{ *pulumi.OutputState }

func (SystemSamlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSaml)(nil)).Elem()
}

func (o SystemSamlArrayOutput) ToSystemSamlArrayOutput() SystemSamlArrayOutput {
	return o
}

func (o SystemSamlArrayOutput) ToSystemSamlArrayOutputWithContext(ctx context.Context) SystemSamlArrayOutput {
	return o
}

func (o SystemSamlArrayOutput) Index(i pulumi.IntInput) SystemSamlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSaml {
		return vs[0].([]*SystemSaml)[vs[1].(int)]
	}).(SystemSamlOutput)
}

type SystemSamlMapOutput struct{ *pulumi.OutputState }

func (SystemSamlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSaml)(nil)).Elem()
}

func (o SystemSamlMapOutput) ToSystemSamlMapOutput() SystemSamlMapOutput {
	return o
}

func (o SystemSamlMapOutput) ToSystemSamlMapOutputWithContext(ctx context.Context) SystemSamlMapOutput {
	return o
}

func (o SystemSamlMapOutput) MapIndex(k pulumi.StringInput) SystemSamlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSaml {
		return vs[0].(map[string]*SystemSaml)[vs[1].(string)]
	}).(SystemSamlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSamlInput)(nil)).Elem(), &SystemSaml{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSamlArrayInput)(nil)).Elem(), SystemSamlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSamlMapInput)(nil)).Elem(), SystemSamlMap{})
	pulumi.RegisterOutputType(SystemSamlOutput{})
	pulumi.RegisterOutputType(SystemSamlArrayOutput{})
	pulumi.RegisterOutputType(SystemSamlMapOutput{})
}