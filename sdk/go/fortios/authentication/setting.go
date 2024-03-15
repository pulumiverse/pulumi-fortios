// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentication

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure authentication setting.
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
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := authentication.NewSetting(ctx, "trname", &authentication.SettingArgs{
//				AuthHttps:            pulumi.String("enable"),
//				CaptivePortalIp:      pulumi.String("0.0.0.0"),
//				CaptivePortalIp6:     pulumi.String("::"),
//				CaptivePortalPort:    pulumi.Int(7830),
//				CaptivePortalSslPort: pulumi.Int(7831),
//				CaptivePortalType:    pulumi.String("fqdn"),
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
// Authentication Setting can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:authentication/setting:Setting labelname AuthenticationSetting
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:authentication/setting:Setting labelname AuthenticationSetting
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Setting struct {
	pulumi.CustomResourceState

	// Active authentication method (scheme name).
	ActiveAuthScheme pulumi.StringOutput `pulumi:"activeAuthScheme"`
	// Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
	AuthHttps pulumi.StringOutput `pulumi:"authHttps"`
	// Captive portal host name.
	CaptivePortal pulumi.StringOutput `pulumi:"captivePortal"`
	// IPv6 captive portal host name.
	CaptivePortal6 pulumi.StringOutput `pulumi:"captivePortal6"`
	// Captive portal IP address.
	CaptivePortalIp pulumi.StringOutput `pulumi:"captivePortalIp"`
	// Captive portal IPv6 address.
	CaptivePortalIp6 pulumi.StringOutput `pulumi:"captivePortalIp6"`
	// Captive portal port number (1 - 65535, default = 7830).
	CaptivePortalPort pulumi.IntOutput `pulumi:"captivePortalPort"`
	// Captive portal SSL port number (1 - 65535, default = 7831).
	CaptivePortalSslPort pulumi.IntOutput `pulumi:"captivePortalSslPort"`
	// Captive portal type. Valid values: `fqdn`, `ip`.
	CaptivePortalType pulumi.StringOutput `pulumi:"captivePortalType"`
	// Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable`, `disable`.
	CertAuth pulumi.StringOutput `pulumi:"certAuth"`
	// Certificate captive portal host name.
	CertCaptivePortal pulumi.StringOutput `pulumi:"certCaptivePortal"`
	// Certificate captive portal IP address.
	CertCaptivePortalIp pulumi.StringOutput `pulumi:"certCaptivePortalIp"`
	// Certificate captive portal port number (1 - 65535, default = 7832).
	CertCaptivePortalPort pulumi.IntOutput `pulumi:"certCaptivePortalPort"`
	// Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
	CookieMaxAge pulumi.IntOutput `pulumi:"cookieMaxAge"`
	// Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
	CookieRefreshDiv pulumi.IntOutput `pulumi:"cookieRefreshDiv"`
	// Address range for the IP based device query. The structure of `devRange` block is documented below.
	DevRanges SettingDevRangeArrayOutput `pulumi:"devRanges"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable persistent cookie on IP based web portal authentication (default = disable). Valid values: `enable`, `disable`.
	IpAuthCookie pulumi.StringOutput `pulumi:"ipAuthCookie"`
	// Enable/disable persistent cookie on web portal authentication (default = enable). Valid values: `enable`, `disable`.
	PersistentCookie pulumi.StringOutput `pulumi:"persistentCookie"`
	// Single-Sign-On authentication method (scheme name).
	SsoAuthScheme pulumi.StringOutput `pulumi:"ssoAuthScheme"`
	// Time of the last update.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// CA certificate used for client certificate verification. The structure of `userCertCa` block is documented below.
	UserCertCas SettingUserCertCaArrayOutput `pulumi:"userCertCas"`
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
	err := ctx.RegisterResource("fortios:authentication/setting:Setting", name, args, &resource, opts...)
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
	err := ctx.ReadResource("fortios:authentication/setting:Setting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Setting resources.
type settingState struct {
	// Active authentication method (scheme name).
	ActiveAuthScheme *string `pulumi:"activeAuthScheme"`
	// Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
	AuthHttps *string `pulumi:"authHttps"`
	// Captive portal host name.
	CaptivePortal *string `pulumi:"captivePortal"`
	// IPv6 captive portal host name.
	CaptivePortal6 *string `pulumi:"captivePortal6"`
	// Captive portal IP address.
	CaptivePortalIp *string `pulumi:"captivePortalIp"`
	// Captive portal IPv6 address.
	CaptivePortalIp6 *string `pulumi:"captivePortalIp6"`
	// Captive portal port number (1 - 65535, default = 7830).
	CaptivePortalPort *int `pulumi:"captivePortalPort"`
	// Captive portal SSL port number (1 - 65535, default = 7831).
	CaptivePortalSslPort *int `pulumi:"captivePortalSslPort"`
	// Captive portal type. Valid values: `fqdn`, `ip`.
	CaptivePortalType *string `pulumi:"captivePortalType"`
	// Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable`, `disable`.
	CertAuth *string `pulumi:"certAuth"`
	// Certificate captive portal host name.
	CertCaptivePortal *string `pulumi:"certCaptivePortal"`
	// Certificate captive portal IP address.
	CertCaptivePortalIp *string `pulumi:"certCaptivePortalIp"`
	// Certificate captive portal port number (1 - 65535, default = 7832).
	CertCaptivePortalPort *int `pulumi:"certCaptivePortalPort"`
	// Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
	CookieMaxAge *int `pulumi:"cookieMaxAge"`
	// Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
	CookieRefreshDiv *int `pulumi:"cookieRefreshDiv"`
	// Address range for the IP based device query. The structure of `devRange` block is documented below.
	DevRanges []SettingDevRange `pulumi:"devRanges"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable persistent cookie on IP based web portal authentication (default = disable). Valid values: `enable`, `disable`.
	IpAuthCookie *string `pulumi:"ipAuthCookie"`
	// Enable/disable persistent cookie on web portal authentication (default = enable). Valid values: `enable`, `disable`.
	PersistentCookie *string `pulumi:"persistentCookie"`
	// Single-Sign-On authentication method (scheme name).
	SsoAuthScheme *string `pulumi:"ssoAuthScheme"`
	// Time of the last update.
	UpdateTime *string `pulumi:"updateTime"`
	// CA certificate used for client certificate verification. The structure of `userCertCa` block is documented below.
	UserCertCas []SettingUserCertCa `pulumi:"userCertCas"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SettingState struct {
	// Active authentication method (scheme name).
	ActiveAuthScheme pulumi.StringPtrInput
	// Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
	AuthHttps pulumi.StringPtrInput
	// Captive portal host name.
	CaptivePortal pulumi.StringPtrInput
	// IPv6 captive portal host name.
	CaptivePortal6 pulumi.StringPtrInput
	// Captive portal IP address.
	CaptivePortalIp pulumi.StringPtrInput
	// Captive portal IPv6 address.
	CaptivePortalIp6 pulumi.StringPtrInput
	// Captive portal port number (1 - 65535, default = 7830).
	CaptivePortalPort pulumi.IntPtrInput
	// Captive portal SSL port number (1 - 65535, default = 7831).
	CaptivePortalSslPort pulumi.IntPtrInput
	// Captive portal type. Valid values: `fqdn`, `ip`.
	CaptivePortalType pulumi.StringPtrInput
	// Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable`, `disable`.
	CertAuth pulumi.StringPtrInput
	// Certificate captive portal host name.
	CertCaptivePortal pulumi.StringPtrInput
	// Certificate captive portal IP address.
	CertCaptivePortalIp pulumi.StringPtrInput
	// Certificate captive portal port number (1 - 65535, default = 7832).
	CertCaptivePortalPort pulumi.IntPtrInput
	// Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
	CookieMaxAge pulumi.IntPtrInput
	// Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
	CookieRefreshDiv pulumi.IntPtrInput
	// Address range for the IP based device query. The structure of `devRange` block is documented below.
	DevRanges SettingDevRangeArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable persistent cookie on IP based web portal authentication (default = disable). Valid values: `enable`, `disable`.
	IpAuthCookie pulumi.StringPtrInput
	// Enable/disable persistent cookie on web portal authentication (default = enable). Valid values: `enable`, `disable`.
	PersistentCookie pulumi.StringPtrInput
	// Single-Sign-On authentication method (scheme name).
	SsoAuthScheme pulumi.StringPtrInput
	// Time of the last update.
	UpdateTime pulumi.StringPtrInput
	// CA certificate used for client certificate verification. The structure of `userCertCa` block is documented below.
	UserCertCas SettingUserCertCaArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*settingState)(nil)).Elem()
}

type settingArgs struct {
	// Active authentication method (scheme name).
	ActiveAuthScheme *string `pulumi:"activeAuthScheme"`
	// Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
	AuthHttps *string `pulumi:"authHttps"`
	// Captive portal host name.
	CaptivePortal *string `pulumi:"captivePortal"`
	// IPv6 captive portal host name.
	CaptivePortal6 *string `pulumi:"captivePortal6"`
	// Captive portal IP address.
	CaptivePortalIp *string `pulumi:"captivePortalIp"`
	// Captive portal IPv6 address.
	CaptivePortalIp6 *string `pulumi:"captivePortalIp6"`
	// Captive portal port number (1 - 65535, default = 7830).
	CaptivePortalPort *int `pulumi:"captivePortalPort"`
	// Captive portal SSL port number (1 - 65535, default = 7831).
	CaptivePortalSslPort *int `pulumi:"captivePortalSslPort"`
	// Captive portal type. Valid values: `fqdn`, `ip`.
	CaptivePortalType *string `pulumi:"captivePortalType"`
	// Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable`, `disable`.
	CertAuth *string `pulumi:"certAuth"`
	// Certificate captive portal host name.
	CertCaptivePortal *string `pulumi:"certCaptivePortal"`
	// Certificate captive portal IP address.
	CertCaptivePortalIp *string `pulumi:"certCaptivePortalIp"`
	// Certificate captive portal port number (1 - 65535, default = 7832).
	CertCaptivePortalPort *int `pulumi:"certCaptivePortalPort"`
	// Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
	CookieMaxAge *int `pulumi:"cookieMaxAge"`
	// Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
	CookieRefreshDiv *int `pulumi:"cookieRefreshDiv"`
	// Address range for the IP based device query. The structure of `devRange` block is documented below.
	DevRanges []SettingDevRange `pulumi:"devRanges"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable persistent cookie on IP based web portal authentication (default = disable). Valid values: `enable`, `disable`.
	IpAuthCookie *string `pulumi:"ipAuthCookie"`
	// Enable/disable persistent cookie on web portal authentication (default = enable). Valid values: `enable`, `disable`.
	PersistentCookie *string `pulumi:"persistentCookie"`
	// Single-Sign-On authentication method (scheme name).
	SsoAuthScheme *string `pulumi:"ssoAuthScheme"`
	// Time of the last update.
	UpdateTime *string `pulumi:"updateTime"`
	// CA certificate used for client certificate verification. The structure of `userCertCa` block is documented below.
	UserCertCas []SettingUserCertCa `pulumi:"userCertCas"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Setting resource.
type SettingArgs struct {
	// Active authentication method (scheme name).
	ActiveAuthScheme pulumi.StringPtrInput
	// Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
	AuthHttps pulumi.StringPtrInput
	// Captive portal host name.
	CaptivePortal pulumi.StringPtrInput
	// IPv6 captive portal host name.
	CaptivePortal6 pulumi.StringPtrInput
	// Captive portal IP address.
	CaptivePortalIp pulumi.StringPtrInput
	// Captive portal IPv6 address.
	CaptivePortalIp6 pulumi.StringPtrInput
	// Captive portal port number (1 - 65535, default = 7830).
	CaptivePortalPort pulumi.IntPtrInput
	// Captive portal SSL port number (1 - 65535, default = 7831).
	CaptivePortalSslPort pulumi.IntPtrInput
	// Captive portal type. Valid values: `fqdn`, `ip`.
	CaptivePortalType pulumi.StringPtrInput
	// Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable`, `disable`.
	CertAuth pulumi.StringPtrInput
	// Certificate captive portal host name.
	CertCaptivePortal pulumi.StringPtrInput
	// Certificate captive portal IP address.
	CertCaptivePortalIp pulumi.StringPtrInput
	// Certificate captive portal port number (1 - 65535, default = 7832).
	CertCaptivePortalPort pulumi.IntPtrInput
	// Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
	CookieMaxAge pulumi.IntPtrInput
	// Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
	CookieRefreshDiv pulumi.IntPtrInput
	// Address range for the IP based device query. The structure of `devRange` block is documented below.
	DevRanges SettingDevRangeArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable persistent cookie on IP based web portal authentication (default = disable). Valid values: `enable`, `disable`.
	IpAuthCookie pulumi.StringPtrInput
	// Enable/disable persistent cookie on web portal authentication (default = enable). Valid values: `enable`, `disable`.
	PersistentCookie pulumi.StringPtrInput
	// Single-Sign-On authentication method (scheme name).
	SsoAuthScheme pulumi.StringPtrInput
	// Time of the last update.
	UpdateTime pulumi.StringPtrInput
	// CA certificate used for client certificate verification. The structure of `userCertCa` block is documented below.
	UserCertCas SettingUserCertCaArrayInput
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

// Active authentication method (scheme name).
func (o SettingOutput) ActiveAuthScheme() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.ActiveAuthScheme }).(pulumi.StringOutput)
}

// Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
func (o SettingOutput) AuthHttps() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.AuthHttps }).(pulumi.StringOutput)
}

// Captive portal host name.
func (o SettingOutput) CaptivePortal() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.CaptivePortal }).(pulumi.StringOutput)
}

// IPv6 captive portal host name.
func (o SettingOutput) CaptivePortal6() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.CaptivePortal6 }).(pulumi.StringOutput)
}

// Captive portal IP address.
func (o SettingOutput) CaptivePortalIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.CaptivePortalIp }).(pulumi.StringOutput)
}

// Captive portal IPv6 address.
func (o SettingOutput) CaptivePortalIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.CaptivePortalIp6 }).(pulumi.StringOutput)
}

// Captive portal port number (1 - 65535, default = 7830).
func (o SettingOutput) CaptivePortalPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.CaptivePortalPort }).(pulumi.IntOutput)
}

// Captive portal SSL port number (1 - 65535, default = 7831).
func (o SettingOutput) CaptivePortalSslPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.CaptivePortalSslPort }).(pulumi.IntOutput)
}

// Captive portal type. Valid values: `fqdn`, `ip`.
func (o SettingOutput) CaptivePortalType() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.CaptivePortalType }).(pulumi.StringOutput)
}

// Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable`, `disable`.
func (o SettingOutput) CertAuth() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.CertAuth }).(pulumi.StringOutput)
}

// Certificate captive portal host name.
func (o SettingOutput) CertCaptivePortal() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.CertCaptivePortal }).(pulumi.StringOutput)
}

// Certificate captive portal IP address.
func (o SettingOutput) CertCaptivePortalIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.CertCaptivePortalIp }).(pulumi.StringOutput)
}

// Certificate captive portal port number (1 - 65535, default = 7832).
func (o SettingOutput) CertCaptivePortalPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.CertCaptivePortalPort }).(pulumi.IntOutput)
}

// Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
func (o SettingOutput) CookieMaxAge() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.CookieMaxAge }).(pulumi.IntOutput)
}

// Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
func (o SettingOutput) CookieRefreshDiv() pulumi.IntOutput {
	return o.ApplyT(func(v *Setting) pulumi.IntOutput { return v.CookieRefreshDiv }).(pulumi.IntOutput)
}

// Address range for the IP based device query. The structure of `devRange` block is documented below.
func (o SettingOutput) DevRanges() SettingDevRangeArrayOutput {
	return o.ApplyT(func(v *Setting) SettingDevRangeArrayOutput { return v.DevRanges }).(SettingDevRangeArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SettingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable persistent cookie on IP based web portal authentication (default = disable). Valid values: `enable`, `disable`.
func (o SettingOutput) IpAuthCookie() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.IpAuthCookie }).(pulumi.StringOutput)
}

// Enable/disable persistent cookie on web portal authentication (default = enable). Valid values: `enable`, `disable`.
func (o SettingOutput) PersistentCookie() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.PersistentCookie }).(pulumi.StringOutput)
}

// Single-Sign-On authentication method (scheme name).
func (o SettingOutput) SsoAuthScheme() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.SsoAuthScheme }).(pulumi.StringOutput)
}

// Time of the last update.
func (o SettingOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// CA certificate used for client certificate verification. The structure of `userCertCa` block is documented below.
func (o SettingOutput) UserCertCas() SettingUserCertCaArrayOutput {
	return o.ApplyT(func(v *Setting) SettingUserCertCaArrayOutput { return v.UserCertCas }).(SettingUserCertCaArrayOutput)
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