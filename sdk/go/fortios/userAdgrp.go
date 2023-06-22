// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FSSO groups.
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
//			trname1, err := fortios.NewUserFsso(ctx, "trname1", &fortios.UserFssoArgs{
//				Port:      pulumi.Int(32381),
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
//			_, err = fortios.NewUserAdgrp(ctx, "trname", &fortios.UserAdgrpArgs{
//				ServerName: trname1.Name,
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
// # User Adgrp can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/userAdgrp:UserAdgrp labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/userAdgrp:UserAdgrp labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type UserAdgrp struct {
	pulumi.CustomResourceState

	// FSSO connector source.
	ConnectorSource pulumi.StringOutput `pulumi:"connectorSource"`
	// Group ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// FSSO agent name.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUserAdgrp registers a new resource with the given unique name, arguments, and options.
func NewUserAdgrp(ctx *pulumi.Context,
	name string, args *UserAdgrpArgs, opts ...pulumi.ResourceOption) (*UserAdgrp, error) {
	if args == nil {
		args = &UserAdgrpArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource UserAdgrp
	err := ctx.RegisterResource("fortios:index/userAdgrp:UserAdgrp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserAdgrp gets an existing UserAdgrp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserAdgrp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserAdgrpState, opts ...pulumi.ResourceOption) (*UserAdgrp, error) {
	var resource UserAdgrp
	err := ctx.ReadResource("fortios:index/userAdgrp:UserAdgrp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserAdgrp resources.
type userAdgrpState struct {
	// FSSO connector source.
	ConnectorSource *string `pulumi:"connectorSource"`
	// Group ID.
	Fosid *int `pulumi:"fosid"`
	// Name.
	Name *string `pulumi:"name"`
	// FSSO agent name.
	ServerName *string `pulumi:"serverName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UserAdgrpState struct {
	// FSSO connector source.
	ConnectorSource pulumi.StringPtrInput
	// Group ID.
	Fosid pulumi.IntPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// FSSO agent name.
	ServerName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserAdgrpState) ElementType() reflect.Type {
	return reflect.TypeOf((*userAdgrpState)(nil)).Elem()
}

type userAdgrpArgs struct {
	// FSSO connector source.
	ConnectorSource *string `pulumi:"connectorSource"`
	// Group ID.
	Fosid *int `pulumi:"fosid"`
	// Name.
	Name *string `pulumi:"name"`
	// FSSO agent name.
	ServerName *string `pulumi:"serverName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserAdgrp resource.
type UserAdgrpArgs struct {
	// FSSO connector source.
	ConnectorSource pulumi.StringPtrInput
	// Group ID.
	Fosid pulumi.IntPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// FSSO agent name.
	ServerName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserAdgrpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userAdgrpArgs)(nil)).Elem()
}

type UserAdgrpInput interface {
	pulumi.Input

	ToUserAdgrpOutput() UserAdgrpOutput
	ToUserAdgrpOutputWithContext(ctx context.Context) UserAdgrpOutput
}

func (*UserAdgrp) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAdgrp)(nil)).Elem()
}

func (i *UserAdgrp) ToUserAdgrpOutput() UserAdgrpOutput {
	return i.ToUserAdgrpOutputWithContext(context.Background())
}

func (i *UserAdgrp) ToUserAdgrpOutputWithContext(ctx context.Context) UserAdgrpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAdgrpOutput)
}

// UserAdgrpArrayInput is an input type that accepts UserAdgrpArray and UserAdgrpArrayOutput values.
// You can construct a concrete instance of `UserAdgrpArrayInput` via:
//
//	UserAdgrpArray{ UserAdgrpArgs{...} }
type UserAdgrpArrayInput interface {
	pulumi.Input

	ToUserAdgrpArrayOutput() UserAdgrpArrayOutput
	ToUserAdgrpArrayOutputWithContext(context.Context) UserAdgrpArrayOutput
}

type UserAdgrpArray []UserAdgrpInput

func (UserAdgrpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserAdgrp)(nil)).Elem()
}

func (i UserAdgrpArray) ToUserAdgrpArrayOutput() UserAdgrpArrayOutput {
	return i.ToUserAdgrpArrayOutputWithContext(context.Background())
}

func (i UserAdgrpArray) ToUserAdgrpArrayOutputWithContext(ctx context.Context) UserAdgrpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAdgrpArrayOutput)
}

// UserAdgrpMapInput is an input type that accepts UserAdgrpMap and UserAdgrpMapOutput values.
// You can construct a concrete instance of `UserAdgrpMapInput` via:
//
//	UserAdgrpMap{ "key": UserAdgrpArgs{...} }
type UserAdgrpMapInput interface {
	pulumi.Input

	ToUserAdgrpMapOutput() UserAdgrpMapOutput
	ToUserAdgrpMapOutputWithContext(context.Context) UserAdgrpMapOutput
}

type UserAdgrpMap map[string]UserAdgrpInput

func (UserAdgrpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserAdgrp)(nil)).Elem()
}

func (i UserAdgrpMap) ToUserAdgrpMapOutput() UserAdgrpMapOutput {
	return i.ToUserAdgrpMapOutputWithContext(context.Background())
}

func (i UserAdgrpMap) ToUserAdgrpMapOutputWithContext(ctx context.Context) UserAdgrpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAdgrpMapOutput)
}

type UserAdgrpOutput struct{ *pulumi.OutputState }

func (UserAdgrpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAdgrp)(nil)).Elem()
}

func (o UserAdgrpOutput) ToUserAdgrpOutput() UserAdgrpOutput {
	return o
}

func (o UserAdgrpOutput) ToUserAdgrpOutputWithContext(ctx context.Context) UserAdgrpOutput {
	return o
}

// FSSO connector source.
func (o UserAdgrpOutput) ConnectorSource() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAdgrp) pulumi.StringOutput { return v.ConnectorSource }).(pulumi.StringOutput)
}

// Group ID.
func (o UserAdgrpOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *UserAdgrp) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Name.
func (o UserAdgrpOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAdgrp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// FSSO agent name.
func (o UserAdgrpOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAdgrp) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o UserAdgrpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAdgrp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type UserAdgrpArrayOutput struct{ *pulumi.OutputState }

func (UserAdgrpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserAdgrp)(nil)).Elem()
}

func (o UserAdgrpArrayOutput) ToUserAdgrpArrayOutput() UserAdgrpArrayOutput {
	return o
}

func (o UserAdgrpArrayOutput) ToUserAdgrpArrayOutputWithContext(ctx context.Context) UserAdgrpArrayOutput {
	return o
}

func (o UserAdgrpArrayOutput) Index(i pulumi.IntInput) UserAdgrpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserAdgrp {
		return vs[0].([]*UserAdgrp)[vs[1].(int)]
	}).(UserAdgrpOutput)
}

type UserAdgrpMapOutput struct{ *pulumi.OutputState }

func (UserAdgrpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserAdgrp)(nil)).Elem()
}

func (o UserAdgrpMapOutput) ToUserAdgrpMapOutput() UserAdgrpMapOutput {
	return o
}

func (o UserAdgrpMapOutput) ToUserAdgrpMapOutputWithContext(ctx context.Context) UserAdgrpMapOutput {
	return o
}

func (o UserAdgrpMapOutput) MapIndex(k pulumi.StringInput) UserAdgrpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserAdgrp {
		return vs[0].(map[string]*UserAdgrp)[vs[1].(string)]
	}).(UserAdgrpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserAdgrpInput)(nil)).Elem(), &UserAdgrp{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAdgrpArrayInput)(nil)).Elem(), UserAdgrpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAdgrpMapInput)(nil)).Elem(), UserAdgrpMap{})
	pulumi.RegisterOutputType(UserAdgrpOutput{})
	pulumi.RegisterOutputType(UserAdgrpArrayOutput{})
	pulumi.RegisterOutputType(UserAdgrpMapOutput{})
}
