// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports modifying system admin setting for FortiManager.
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
//			_, err := fortios.NewFmgSystemAdmin(ctx, "test1", &fortios.FmgSystemAdminArgs{
//				HttpPort:    pulumi.Int(80),
//				HttpsPort:   pulumi.Int(443),
//				IdleTimeout: pulumi.Int(20),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FmgSystemAdmin struct {
	pulumi.CustomResourceState

	// Http port.
	HttpPort pulumi.IntPtrOutput `pulumi:"httpPort"`
	// Https port.
	HttpsPort pulumi.IntPtrOutput `pulumi:"httpsPort"`
	// Idle Timeout(1-480 minute).
	IdleTimeout pulumi.IntPtrOutput `pulumi:"idleTimeout"`
}

// NewFmgSystemAdmin registers a new resource with the given unique name, arguments, and options.
func NewFmgSystemAdmin(ctx *pulumi.Context,
	name string, args *FmgSystemAdminArgs, opts ...pulumi.ResourceOption) (*FmgSystemAdmin, error) {
	if args == nil {
		args = &FmgSystemAdminArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FmgSystemAdmin
	err := ctx.RegisterResource("fortios:index/fmgSystemAdmin:FmgSystemAdmin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFmgSystemAdmin gets an existing FmgSystemAdmin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFmgSystemAdmin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FmgSystemAdminState, opts ...pulumi.ResourceOption) (*FmgSystemAdmin, error) {
	var resource FmgSystemAdmin
	err := ctx.ReadResource("fortios:index/fmgSystemAdmin:FmgSystemAdmin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FmgSystemAdmin resources.
type fmgSystemAdminState struct {
	// Http port.
	HttpPort *int `pulumi:"httpPort"`
	// Https port.
	HttpsPort *int `pulumi:"httpsPort"`
	// Idle Timeout(1-480 minute).
	IdleTimeout *int `pulumi:"idleTimeout"`
}

type FmgSystemAdminState struct {
	// Http port.
	HttpPort pulumi.IntPtrInput
	// Https port.
	HttpsPort pulumi.IntPtrInput
	// Idle Timeout(1-480 minute).
	IdleTimeout pulumi.IntPtrInput
}

func (FmgSystemAdminState) ElementType() reflect.Type {
	return reflect.TypeOf((*fmgSystemAdminState)(nil)).Elem()
}

type fmgSystemAdminArgs struct {
	// Http port.
	HttpPort *int `pulumi:"httpPort"`
	// Https port.
	HttpsPort *int `pulumi:"httpsPort"`
	// Idle Timeout(1-480 minute).
	IdleTimeout *int `pulumi:"idleTimeout"`
}

// The set of arguments for constructing a FmgSystemAdmin resource.
type FmgSystemAdminArgs struct {
	// Http port.
	HttpPort pulumi.IntPtrInput
	// Https port.
	HttpsPort pulumi.IntPtrInput
	// Idle Timeout(1-480 minute).
	IdleTimeout pulumi.IntPtrInput
}

func (FmgSystemAdminArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fmgSystemAdminArgs)(nil)).Elem()
}

type FmgSystemAdminInput interface {
	pulumi.Input

	ToFmgSystemAdminOutput() FmgSystemAdminOutput
	ToFmgSystemAdminOutputWithContext(ctx context.Context) FmgSystemAdminOutput
}

func (*FmgSystemAdmin) ElementType() reflect.Type {
	return reflect.TypeOf((**FmgSystemAdmin)(nil)).Elem()
}

func (i *FmgSystemAdmin) ToFmgSystemAdminOutput() FmgSystemAdminOutput {
	return i.ToFmgSystemAdminOutputWithContext(context.Background())
}

func (i *FmgSystemAdmin) ToFmgSystemAdminOutputWithContext(ctx context.Context) FmgSystemAdminOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgSystemAdminOutput)
}

// FmgSystemAdminArrayInput is an input type that accepts FmgSystemAdminArray and FmgSystemAdminArrayOutput values.
// You can construct a concrete instance of `FmgSystemAdminArrayInput` via:
//
//	FmgSystemAdminArray{ FmgSystemAdminArgs{...} }
type FmgSystemAdminArrayInput interface {
	pulumi.Input

	ToFmgSystemAdminArrayOutput() FmgSystemAdminArrayOutput
	ToFmgSystemAdminArrayOutputWithContext(context.Context) FmgSystemAdminArrayOutput
}

type FmgSystemAdminArray []FmgSystemAdminInput

func (FmgSystemAdminArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FmgSystemAdmin)(nil)).Elem()
}

func (i FmgSystemAdminArray) ToFmgSystemAdminArrayOutput() FmgSystemAdminArrayOutput {
	return i.ToFmgSystemAdminArrayOutputWithContext(context.Background())
}

func (i FmgSystemAdminArray) ToFmgSystemAdminArrayOutputWithContext(ctx context.Context) FmgSystemAdminArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgSystemAdminArrayOutput)
}

// FmgSystemAdminMapInput is an input type that accepts FmgSystemAdminMap and FmgSystemAdminMapOutput values.
// You can construct a concrete instance of `FmgSystemAdminMapInput` via:
//
//	FmgSystemAdminMap{ "key": FmgSystemAdminArgs{...} }
type FmgSystemAdminMapInput interface {
	pulumi.Input

	ToFmgSystemAdminMapOutput() FmgSystemAdminMapOutput
	ToFmgSystemAdminMapOutputWithContext(context.Context) FmgSystemAdminMapOutput
}

type FmgSystemAdminMap map[string]FmgSystemAdminInput

func (FmgSystemAdminMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FmgSystemAdmin)(nil)).Elem()
}

func (i FmgSystemAdminMap) ToFmgSystemAdminMapOutput() FmgSystemAdminMapOutput {
	return i.ToFmgSystemAdminMapOutputWithContext(context.Background())
}

func (i FmgSystemAdminMap) ToFmgSystemAdminMapOutputWithContext(ctx context.Context) FmgSystemAdminMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgSystemAdminMapOutput)
}

type FmgSystemAdminOutput struct{ *pulumi.OutputState }

func (FmgSystemAdminOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FmgSystemAdmin)(nil)).Elem()
}

func (o FmgSystemAdminOutput) ToFmgSystemAdminOutput() FmgSystemAdminOutput {
	return o
}

func (o FmgSystemAdminOutput) ToFmgSystemAdminOutputWithContext(ctx context.Context) FmgSystemAdminOutput {
	return o
}

// Http port.
func (o FmgSystemAdminOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FmgSystemAdmin) pulumi.IntPtrOutput { return v.HttpPort }).(pulumi.IntPtrOutput)
}

// Https port.
func (o FmgSystemAdminOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FmgSystemAdmin) pulumi.IntPtrOutput { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

// Idle Timeout(1-480 minute).
func (o FmgSystemAdminOutput) IdleTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FmgSystemAdmin) pulumi.IntPtrOutput { return v.IdleTimeout }).(pulumi.IntPtrOutput)
}

type FmgSystemAdminArrayOutput struct{ *pulumi.OutputState }

func (FmgSystemAdminArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FmgSystemAdmin)(nil)).Elem()
}

func (o FmgSystemAdminArrayOutput) ToFmgSystemAdminArrayOutput() FmgSystemAdminArrayOutput {
	return o
}

func (o FmgSystemAdminArrayOutput) ToFmgSystemAdminArrayOutputWithContext(ctx context.Context) FmgSystemAdminArrayOutput {
	return o
}

func (o FmgSystemAdminArrayOutput) Index(i pulumi.IntInput) FmgSystemAdminOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FmgSystemAdmin {
		return vs[0].([]*FmgSystemAdmin)[vs[1].(int)]
	}).(FmgSystemAdminOutput)
}

type FmgSystemAdminMapOutput struct{ *pulumi.OutputState }

func (FmgSystemAdminMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FmgSystemAdmin)(nil)).Elem()
}

func (o FmgSystemAdminMapOutput) ToFmgSystemAdminMapOutput() FmgSystemAdminMapOutput {
	return o
}

func (o FmgSystemAdminMapOutput) ToFmgSystemAdminMapOutputWithContext(ctx context.Context) FmgSystemAdminMapOutput {
	return o
}

func (o FmgSystemAdminMapOutput) MapIndex(k pulumi.StringInput) FmgSystemAdminOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FmgSystemAdmin {
		return vs[0].(map[string]*FmgSystemAdmin)[vs[1].(string)]
	}).(FmgSystemAdminOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FmgSystemAdminInput)(nil)).Elem(), &FmgSystemAdmin{})
	pulumi.RegisterInputType(reflect.TypeOf((*FmgSystemAdminArrayInput)(nil)).Elem(), FmgSystemAdminArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FmgSystemAdminMapInput)(nil)).Elem(), FmgSystemAdminMap{})
	pulumi.RegisterOutputType(FmgSystemAdminOutput{})
	pulumi.RegisterOutputType(FmgSystemAdminArrayOutput{})
	pulumi.RegisterOutputType(FmgSystemAdminMapOutput{})
}
