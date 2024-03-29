// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fmg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// This resource supports Create/Read/Update/Delete admin user for FortiManager
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/fmg"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fmg.NewSystemAdminUser(ctx, "test1", &fmg.SystemAdminUserArgs{
//				Description: pulumi.String("local user"),
//				Password:    pulumi.String("123"),
//				Profileid:   pulumi.String("Standard_User"),
//				RpcPermit:   pulumi.String("read"),
//				Trusthost1:  pulumi.String("1.1.1.1 255.255.255.255"),
//				UserType:    pulumi.String("local"),
//				Userid:      pulumi.String("user1"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = fmg.NewSystemAdminUser(ctx, "test2", &fmg.SystemAdminUserArgs{
//				Description: pulumi.String("api user"),
//				Password:    pulumi.String("098"),
//				Profileid:   pulumi.String("Standard_User"),
//				RpcPermit:   pulumi.String("read-write"),
//				Trusthost1:  pulumi.String("2.2.2.2 255.255.255.255"),
//				Userid:      pulumi.String("user2"),
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
type SystemAdminUser struct {
	pulumi.CustomResourceState

	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Password.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Profile id.
	Profileid pulumi.StringPtrOutput `pulumi:"profileid"`
	// RADIUS server name.
	RadiusServer pulumi.StringPtrOutput `pulumi:"radiusServer"`
	// Rpc permit, Enum: ["read-write", "none", "read"]
	RpcPermit pulumi.StringPtrOutput `pulumi:"rpcPermit"`
	// Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.
	Trusthost1 pulumi.StringPtrOutput `pulumi:"trusthost1"`
	// Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
	Trusthost2 pulumi.StringPtrOutput `pulumi:"trusthost2"`
	// Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
	Trusthost3 pulumi.StringPtrOutput `pulumi:"trusthost3"`
	// User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]
	UserType pulumi.StringPtrOutput `pulumi:"userType"`
	// User name.
	Userid pulumi.StringOutput `pulumi:"userid"`
}

// NewSystemAdminUser registers a new resource with the given unique name, arguments, and options.
func NewSystemAdminUser(ctx *pulumi.Context,
	name string, args *SystemAdminUserArgs, opts ...pulumi.ResourceOption) (*SystemAdminUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Userid == nil {
		return nil, errors.New("invalid value for required argument 'Userid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemAdminUser
	err := ctx.RegisterResource("fortios:fmg/systemAdminUser:SystemAdminUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAdminUser gets an existing SystemAdminUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAdminUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAdminUserState, opts ...pulumi.ResourceOption) (*SystemAdminUser, error) {
	var resource SystemAdminUser
	err := ctx.ReadResource("fortios:fmg/systemAdminUser:SystemAdminUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAdminUser resources.
type systemAdminUserState struct {
	// Description.
	Description *string `pulumi:"description"`
	// Password.
	Password *string `pulumi:"password"`
	// Profile id.
	Profileid *string `pulumi:"profileid"`
	// RADIUS server name.
	RadiusServer *string `pulumi:"radiusServer"`
	// Rpc permit, Enum: ["read-write", "none", "read"]
	RpcPermit *string `pulumi:"rpcPermit"`
	// Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.
	Trusthost1 *string `pulumi:"trusthost1"`
	// Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
	Trusthost2 *string `pulumi:"trusthost2"`
	// Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
	Trusthost3 *string `pulumi:"trusthost3"`
	// User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]
	UserType *string `pulumi:"userType"`
	// User name.
	Userid *string `pulumi:"userid"`
}

type SystemAdminUserState struct {
	// Description.
	Description pulumi.StringPtrInput
	// Password.
	Password pulumi.StringPtrInput
	// Profile id.
	Profileid pulumi.StringPtrInput
	// RADIUS server name.
	RadiusServer pulumi.StringPtrInput
	// Rpc permit, Enum: ["read-write", "none", "read"]
	RpcPermit pulumi.StringPtrInput
	// Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.
	Trusthost1 pulumi.StringPtrInput
	// Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
	Trusthost2 pulumi.StringPtrInput
	// Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
	Trusthost3 pulumi.StringPtrInput
	// User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]
	UserType pulumi.StringPtrInput
	// User name.
	Userid pulumi.StringPtrInput
}

func (SystemAdminUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAdminUserState)(nil)).Elem()
}

type systemAdminUserArgs struct {
	// Description.
	Description *string `pulumi:"description"`
	// Password.
	Password *string `pulumi:"password"`
	// Profile id.
	Profileid *string `pulumi:"profileid"`
	// RADIUS server name.
	RadiusServer *string `pulumi:"radiusServer"`
	// Rpc permit, Enum: ["read-write", "none", "read"]
	RpcPermit *string `pulumi:"rpcPermit"`
	// Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.
	Trusthost1 *string `pulumi:"trusthost1"`
	// Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
	Trusthost2 *string `pulumi:"trusthost2"`
	// Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
	Trusthost3 *string `pulumi:"trusthost3"`
	// User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]
	UserType *string `pulumi:"userType"`
	// User name.
	Userid string `pulumi:"userid"`
}

// The set of arguments for constructing a SystemAdminUser resource.
type SystemAdminUserArgs struct {
	// Description.
	Description pulumi.StringPtrInput
	// Password.
	Password pulumi.StringPtrInput
	// Profile id.
	Profileid pulumi.StringPtrInput
	// RADIUS server name.
	RadiusServer pulumi.StringPtrInput
	// Rpc permit, Enum: ["read-write", "none", "read"]
	RpcPermit pulumi.StringPtrInput
	// Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.
	Trusthost1 pulumi.StringPtrInput
	// Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
	Trusthost2 pulumi.StringPtrInput
	// Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
	Trusthost3 pulumi.StringPtrInput
	// User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]
	UserType pulumi.StringPtrInput
	// User name.
	Userid pulumi.StringInput
}

func (SystemAdminUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAdminUserArgs)(nil)).Elem()
}

type SystemAdminUserInput interface {
	pulumi.Input

	ToSystemAdminUserOutput() SystemAdminUserOutput
	ToSystemAdminUserOutputWithContext(ctx context.Context) SystemAdminUserOutput
}

func (*SystemAdminUser) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAdminUser)(nil)).Elem()
}

func (i *SystemAdminUser) ToSystemAdminUserOutput() SystemAdminUserOutput {
	return i.ToSystemAdminUserOutputWithContext(context.Background())
}

func (i *SystemAdminUser) ToSystemAdminUserOutputWithContext(ctx context.Context) SystemAdminUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAdminUserOutput)
}

// SystemAdminUserArrayInput is an input type that accepts SystemAdminUserArray and SystemAdminUserArrayOutput values.
// You can construct a concrete instance of `SystemAdminUserArrayInput` via:
//
//	SystemAdminUserArray{ SystemAdminUserArgs{...} }
type SystemAdminUserArrayInput interface {
	pulumi.Input

	ToSystemAdminUserArrayOutput() SystemAdminUserArrayOutput
	ToSystemAdminUserArrayOutputWithContext(context.Context) SystemAdminUserArrayOutput
}

type SystemAdminUserArray []SystemAdminUserInput

func (SystemAdminUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAdminUser)(nil)).Elem()
}

func (i SystemAdminUserArray) ToSystemAdminUserArrayOutput() SystemAdminUserArrayOutput {
	return i.ToSystemAdminUserArrayOutputWithContext(context.Background())
}

func (i SystemAdminUserArray) ToSystemAdminUserArrayOutputWithContext(ctx context.Context) SystemAdminUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAdminUserArrayOutput)
}

// SystemAdminUserMapInput is an input type that accepts SystemAdminUserMap and SystemAdminUserMapOutput values.
// You can construct a concrete instance of `SystemAdminUserMapInput` via:
//
//	SystemAdminUserMap{ "key": SystemAdminUserArgs{...} }
type SystemAdminUserMapInput interface {
	pulumi.Input

	ToSystemAdminUserMapOutput() SystemAdminUserMapOutput
	ToSystemAdminUserMapOutputWithContext(context.Context) SystemAdminUserMapOutput
}

type SystemAdminUserMap map[string]SystemAdminUserInput

func (SystemAdminUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAdminUser)(nil)).Elem()
}

func (i SystemAdminUserMap) ToSystemAdminUserMapOutput() SystemAdminUserMapOutput {
	return i.ToSystemAdminUserMapOutputWithContext(context.Background())
}

func (i SystemAdminUserMap) ToSystemAdminUserMapOutputWithContext(ctx context.Context) SystemAdminUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAdminUserMapOutput)
}

type SystemAdminUserOutput struct{ *pulumi.OutputState }

func (SystemAdminUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAdminUser)(nil)).Elem()
}

func (o SystemAdminUserOutput) ToSystemAdminUserOutput() SystemAdminUserOutput {
	return o
}

func (o SystemAdminUserOutput) ToSystemAdminUserOutputWithContext(ctx context.Context) SystemAdminUserOutput {
	return o
}

// Description.
func (o SystemAdminUserOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdminUser) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Password.
func (o SystemAdminUserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdminUser) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Profile id.
func (o SystemAdminUserOutput) Profileid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdminUser) pulumi.StringPtrOutput { return v.Profileid }).(pulumi.StringPtrOutput)
}

// RADIUS server name.
func (o SystemAdminUserOutput) RadiusServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdminUser) pulumi.StringPtrOutput { return v.RadiusServer }).(pulumi.StringPtrOutput)
}

// Rpc permit, Enum: ["read-write", "none", "read"]
func (o SystemAdminUserOutput) RpcPermit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdminUser) pulumi.StringPtrOutput { return v.RpcPermit }).(pulumi.StringPtrOutput)
}

// Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.
func (o SystemAdminUserOutput) Trusthost1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdminUser) pulumi.StringPtrOutput { return v.Trusthost1 }).(pulumi.StringPtrOutput)
}

// Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
func (o SystemAdminUserOutput) Trusthost2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdminUser) pulumi.StringPtrOutput { return v.Trusthost2 }).(pulumi.StringPtrOutput)
}

// Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
func (o SystemAdminUserOutput) Trusthost3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdminUser) pulumi.StringPtrOutput { return v.Trusthost3 }).(pulumi.StringPtrOutput)
}

// User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]
func (o SystemAdminUserOutput) UserType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAdminUser) pulumi.StringPtrOutput { return v.UserType }).(pulumi.StringPtrOutput)
}

// User name.
func (o SystemAdminUserOutput) Userid() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAdminUser) pulumi.StringOutput { return v.Userid }).(pulumi.StringOutput)
}

type SystemAdminUserArrayOutput struct{ *pulumi.OutputState }

func (SystemAdminUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAdminUser)(nil)).Elem()
}

func (o SystemAdminUserArrayOutput) ToSystemAdminUserArrayOutput() SystemAdminUserArrayOutput {
	return o
}

func (o SystemAdminUserArrayOutput) ToSystemAdminUserArrayOutputWithContext(ctx context.Context) SystemAdminUserArrayOutput {
	return o
}

func (o SystemAdminUserArrayOutput) Index(i pulumi.IntInput) SystemAdminUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAdminUser {
		return vs[0].([]*SystemAdminUser)[vs[1].(int)]
	}).(SystemAdminUserOutput)
}

type SystemAdminUserMapOutput struct{ *pulumi.OutputState }

func (SystemAdminUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAdminUser)(nil)).Elem()
}

func (o SystemAdminUserMapOutput) ToSystemAdminUserMapOutput() SystemAdminUserMapOutput {
	return o
}

func (o SystemAdminUserMapOutput) ToSystemAdminUserMapOutputWithContext(ctx context.Context) SystemAdminUserMapOutput {
	return o
}

func (o SystemAdminUserMapOutput) MapIndex(k pulumi.StringInput) SystemAdminUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAdminUser {
		return vs[0].(map[string]*SystemAdminUser)[vs[1].(string)]
	}).(SystemAdminUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAdminUserInput)(nil)).Elem(), &SystemAdminUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAdminUserArrayInput)(nil)).Elem(), SystemAdminUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAdminUserMapInput)(nil)).Elem(), SystemAdminUserMap{})
	pulumi.RegisterOutputType(SystemAdminUserOutput{})
	pulumi.RegisterOutputType(SystemAdminUserArrayOutput{})
	pulumi.RegisterOutputType(SystemAdminUserMapOutput{})
}
