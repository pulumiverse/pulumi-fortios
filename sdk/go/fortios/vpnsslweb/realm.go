// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpnsslweb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Realm.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/vpnsslweb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpnsslweb.NewRealm(ctx, "trname", &vpnsslweb.RealmArgs{
//				LoginPage:         pulumi.String("1.htm"),
//				MaxConcurrentUser: pulumi.Int(33),
//				UrlPath:           pulumi.String("1"),
//				VirtualHost:       pulumi.String("2.2.2.2"),
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
// # VpnSslWeb Realm can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:vpnsslweb/realm:Realm labelname {{url_path}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:vpnsslweb/realm:Realm labelname {{url_path}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Realm struct {
	pulumi.CustomResourceState

	// Replacement HTML for SSL-VPN login page.
	LoginPage pulumi.StringPtrOutput `pulumi:"loginPage"`
	// Maximum concurrent users (0 - 65535, 0 means unlimited).
	MaxConcurrentUser pulumi.IntOutput `pulumi:"maxConcurrentUser"`
	// IP address used as a NAS-IP to communicate with the RADIUS server.
	NasIp pulumi.StringOutput `pulumi:"nasIp"`
	// RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
	RadiusPort pulumi.IntOutput `pulumi:"radiusPort"`
	// RADIUS server associated with realm.
	RadiusServer pulumi.StringOutput `pulumi:"radiusServer"`
	// URL path to access SSL-VPN login page.
	UrlPath pulumi.StringOutput `pulumi:"urlPath"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Virtual host name for realm.
	VirtualHost pulumi.StringPtrOutput `pulumi:"virtualHost"`
	// Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable`, `disable`.
	VirtualHostOnly pulumi.StringOutput `pulumi:"virtualHostOnly"`
	// Name of the server certificate to used for this realm.
	VirtualHostServerCert pulumi.StringOutput `pulumi:"virtualHostServerCert"`
}

// NewRealm registers a new resource with the given unique name, arguments, and options.
func NewRealm(ctx *pulumi.Context,
	name string, args *RealmArgs, opts ...pulumi.ResourceOption) (*Realm, error) {
	if args == nil {
		args = &RealmArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Realm
	err := ctx.RegisterResource("fortios:vpnsslweb/realm:Realm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRealm gets an existing Realm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RealmState, opts ...pulumi.ResourceOption) (*Realm, error) {
	var resource Realm
	err := ctx.ReadResource("fortios:vpnsslweb/realm:Realm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Realm resources.
type realmState struct {
	// Replacement HTML for SSL-VPN login page.
	LoginPage *string `pulumi:"loginPage"`
	// Maximum concurrent users (0 - 65535, 0 means unlimited).
	MaxConcurrentUser *int `pulumi:"maxConcurrentUser"`
	// IP address used as a NAS-IP to communicate with the RADIUS server.
	NasIp *string `pulumi:"nasIp"`
	// RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
	RadiusPort *int `pulumi:"radiusPort"`
	// RADIUS server associated with realm.
	RadiusServer *string `pulumi:"radiusServer"`
	// URL path to access SSL-VPN login page.
	UrlPath *string `pulumi:"urlPath"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual host name for realm.
	VirtualHost *string `pulumi:"virtualHost"`
	// Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable`, `disable`.
	VirtualHostOnly *string `pulumi:"virtualHostOnly"`
	// Name of the server certificate to used for this realm.
	VirtualHostServerCert *string `pulumi:"virtualHostServerCert"`
}

type RealmState struct {
	// Replacement HTML for SSL-VPN login page.
	LoginPage pulumi.StringPtrInput
	// Maximum concurrent users (0 - 65535, 0 means unlimited).
	MaxConcurrentUser pulumi.IntPtrInput
	// IP address used as a NAS-IP to communicate with the RADIUS server.
	NasIp pulumi.StringPtrInput
	// RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
	RadiusPort pulumi.IntPtrInput
	// RADIUS server associated with realm.
	RadiusServer pulumi.StringPtrInput
	// URL path to access SSL-VPN login page.
	UrlPath pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual host name for realm.
	VirtualHost pulumi.StringPtrInput
	// Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable`, `disable`.
	VirtualHostOnly pulumi.StringPtrInput
	// Name of the server certificate to used for this realm.
	VirtualHostServerCert pulumi.StringPtrInput
}

func (RealmState) ElementType() reflect.Type {
	return reflect.TypeOf((*realmState)(nil)).Elem()
}

type realmArgs struct {
	// Replacement HTML for SSL-VPN login page.
	LoginPage *string `pulumi:"loginPage"`
	// Maximum concurrent users (0 - 65535, 0 means unlimited).
	MaxConcurrentUser *int `pulumi:"maxConcurrentUser"`
	// IP address used as a NAS-IP to communicate with the RADIUS server.
	NasIp *string `pulumi:"nasIp"`
	// RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
	RadiusPort *int `pulumi:"radiusPort"`
	// RADIUS server associated with realm.
	RadiusServer *string `pulumi:"radiusServer"`
	// URL path to access SSL-VPN login page.
	UrlPath *string `pulumi:"urlPath"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual host name for realm.
	VirtualHost *string `pulumi:"virtualHost"`
	// Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable`, `disable`.
	VirtualHostOnly *string `pulumi:"virtualHostOnly"`
	// Name of the server certificate to used for this realm.
	VirtualHostServerCert *string `pulumi:"virtualHostServerCert"`
}

// The set of arguments for constructing a Realm resource.
type RealmArgs struct {
	// Replacement HTML for SSL-VPN login page.
	LoginPage pulumi.StringPtrInput
	// Maximum concurrent users (0 - 65535, 0 means unlimited).
	MaxConcurrentUser pulumi.IntPtrInput
	// IP address used as a NAS-IP to communicate with the RADIUS server.
	NasIp pulumi.StringPtrInput
	// RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
	RadiusPort pulumi.IntPtrInput
	// RADIUS server associated with realm.
	RadiusServer pulumi.StringPtrInput
	// URL path to access SSL-VPN login page.
	UrlPath pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual host name for realm.
	VirtualHost pulumi.StringPtrInput
	// Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable`, `disable`.
	VirtualHostOnly pulumi.StringPtrInput
	// Name of the server certificate to used for this realm.
	VirtualHostServerCert pulumi.StringPtrInput
}

func (RealmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*realmArgs)(nil)).Elem()
}

type RealmInput interface {
	pulumi.Input

	ToRealmOutput() RealmOutput
	ToRealmOutputWithContext(ctx context.Context) RealmOutput
}

func (*Realm) ElementType() reflect.Type {
	return reflect.TypeOf((**Realm)(nil)).Elem()
}

func (i *Realm) ToRealmOutput() RealmOutput {
	return i.ToRealmOutputWithContext(context.Background())
}

func (i *Realm) ToRealmOutputWithContext(ctx context.Context) RealmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmOutput)
}

// RealmArrayInput is an input type that accepts RealmArray and RealmArrayOutput values.
// You can construct a concrete instance of `RealmArrayInput` via:
//
//	RealmArray{ RealmArgs{...} }
type RealmArrayInput interface {
	pulumi.Input

	ToRealmArrayOutput() RealmArrayOutput
	ToRealmArrayOutputWithContext(context.Context) RealmArrayOutput
}

type RealmArray []RealmInput

func (RealmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Realm)(nil)).Elem()
}

func (i RealmArray) ToRealmArrayOutput() RealmArrayOutput {
	return i.ToRealmArrayOutputWithContext(context.Background())
}

func (i RealmArray) ToRealmArrayOutputWithContext(ctx context.Context) RealmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmArrayOutput)
}

// RealmMapInput is an input type that accepts RealmMap and RealmMapOutput values.
// You can construct a concrete instance of `RealmMapInput` via:
//
//	RealmMap{ "key": RealmArgs{...} }
type RealmMapInput interface {
	pulumi.Input

	ToRealmMapOutput() RealmMapOutput
	ToRealmMapOutputWithContext(context.Context) RealmMapOutput
}

type RealmMap map[string]RealmInput

func (RealmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Realm)(nil)).Elem()
}

func (i RealmMap) ToRealmMapOutput() RealmMapOutput {
	return i.ToRealmMapOutputWithContext(context.Background())
}

func (i RealmMap) ToRealmMapOutputWithContext(ctx context.Context) RealmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmMapOutput)
}

type RealmOutput struct{ *pulumi.OutputState }

func (RealmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Realm)(nil)).Elem()
}

func (o RealmOutput) ToRealmOutput() RealmOutput {
	return o
}

func (o RealmOutput) ToRealmOutputWithContext(ctx context.Context) RealmOutput {
	return o
}

// Replacement HTML for SSL-VPN login page.
func (o RealmOutput) LoginPage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Realm) pulumi.StringPtrOutput { return v.LoginPage }).(pulumi.StringPtrOutput)
}

// Maximum concurrent users (0 - 65535, 0 means unlimited).
func (o RealmOutput) MaxConcurrentUser() pulumi.IntOutput {
	return o.ApplyT(func(v *Realm) pulumi.IntOutput { return v.MaxConcurrentUser }).(pulumi.IntOutput)
}

// IP address used as a NAS-IP to communicate with the RADIUS server.
func (o RealmOutput) NasIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Realm) pulumi.StringOutput { return v.NasIp }).(pulumi.StringOutput)
}

// RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
func (o RealmOutput) RadiusPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Realm) pulumi.IntOutput { return v.RadiusPort }).(pulumi.IntOutput)
}

// RADIUS server associated with realm.
func (o RealmOutput) RadiusServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Realm) pulumi.StringOutput { return v.RadiusServer }).(pulumi.StringOutput)
}

// URL path to access SSL-VPN login page.
func (o RealmOutput) UrlPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Realm) pulumi.StringOutput { return v.UrlPath }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RealmOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Realm) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Virtual host name for realm.
func (o RealmOutput) VirtualHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Realm) pulumi.StringPtrOutput { return v.VirtualHost }).(pulumi.StringPtrOutput)
}

// Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable`, `disable`.
func (o RealmOutput) VirtualHostOnly() pulumi.StringOutput {
	return o.ApplyT(func(v *Realm) pulumi.StringOutput { return v.VirtualHostOnly }).(pulumi.StringOutput)
}

// Name of the server certificate to used for this realm.
func (o RealmOutput) VirtualHostServerCert() pulumi.StringOutput {
	return o.ApplyT(func(v *Realm) pulumi.StringOutput { return v.VirtualHostServerCert }).(pulumi.StringOutput)
}

type RealmArrayOutput struct{ *pulumi.OutputState }

func (RealmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Realm)(nil)).Elem()
}

func (o RealmArrayOutput) ToRealmArrayOutput() RealmArrayOutput {
	return o
}

func (o RealmArrayOutput) ToRealmArrayOutputWithContext(ctx context.Context) RealmArrayOutput {
	return o
}

func (o RealmArrayOutput) Index(i pulumi.IntInput) RealmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Realm {
		return vs[0].([]*Realm)[vs[1].(int)]
	}).(RealmOutput)
}

type RealmMapOutput struct{ *pulumi.OutputState }

func (RealmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Realm)(nil)).Elem()
}

func (o RealmMapOutput) ToRealmMapOutput() RealmMapOutput {
	return o
}

func (o RealmMapOutput) ToRealmMapOutputWithContext(ctx context.Context) RealmMapOutput {
	return o
}

func (o RealmMapOutput) MapIndex(k pulumi.StringInput) RealmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Realm {
		return vs[0].(map[string]*Realm)[vs[1].(string)]
	}).(RealmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RealmInput)(nil)).Elem(), &Realm{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmArrayInput)(nil)).Elem(), RealmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmMapInput)(nil)).Elem(), RealmMap{})
	pulumi.RegisterOutputType(RealmOutput{})
	pulumi.RegisterOutputType(RealmArrayOutput{})
	pulumi.RegisterOutputType(RealmMapOutput{})
}
