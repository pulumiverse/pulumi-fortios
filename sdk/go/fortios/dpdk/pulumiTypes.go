// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dpdk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

var _ = internal.GetEnvOrDefault

type GlobalInterface struct {
	// Physical interface name.
	InterfaceName *string `pulumi:"interfaceName"`
}

// GlobalInterfaceInput is an input type that accepts GlobalInterfaceArgs and GlobalInterfaceOutput values.
// You can construct a concrete instance of `GlobalInterfaceInput` via:
//
//	GlobalInterfaceArgs{...}
type GlobalInterfaceInput interface {
	pulumi.Input

	ToGlobalInterfaceOutput() GlobalInterfaceOutput
	ToGlobalInterfaceOutputWithContext(context.Context) GlobalInterfaceOutput
}

type GlobalInterfaceArgs struct {
	// Physical interface name.
	InterfaceName pulumi.StringPtrInput `pulumi:"interfaceName"`
}

func (GlobalInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalInterface)(nil)).Elem()
}

func (i GlobalInterfaceArgs) ToGlobalInterfaceOutput() GlobalInterfaceOutput {
	return i.ToGlobalInterfaceOutputWithContext(context.Background())
}

func (i GlobalInterfaceArgs) ToGlobalInterfaceOutputWithContext(ctx context.Context) GlobalInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalInterfaceOutput)
}

// GlobalInterfaceArrayInput is an input type that accepts GlobalInterfaceArray and GlobalInterfaceArrayOutput values.
// You can construct a concrete instance of `GlobalInterfaceArrayInput` via:
//
//	GlobalInterfaceArray{ GlobalInterfaceArgs{...} }
type GlobalInterfaceArrayInput interface {
	pulumi.Input

	ToGlobalInterfaceArrayOutput() GlobalInterfaceArrayOutput
	ToGlobalInterfaceArrayOutputWithContext(context.Context) GlobalInterfaceArrayOutput
}

type GlobalInterfaceArray []GlobalInterfaceInput

func (GlobalInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GlobalInterface)(nil)).Elem()
}

func (i GlobalInterfaceArray) ToGlobalInterfaceArrayOutput() GlobalInterfaceArrayOutput {
	return i.ToGlobalInterfaceArrayOutputWithContext(context.Background())
}

func (i GlobalInterfaceArray) ToGlobalInterfaceArrayOutputWithContext(ctx context.Context) GlobalInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalInterfaceArrayOutput)
}

type GlobalInterfaceOutput struct{ *pulumi.OutputState }

func (GlobalInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalInterface)(nil)).Elem()
}

func (o GlobalInterfaceOutput) ToGlobalInterfaceOutput() GlobalInterfaceOutput {
	return o
}

func (o GlobalInterfaceOutput) ToGlobalInterfaceOutputWithContext(ctx context.Context) GlobalInterfaceOutput {
	return o
}

// Physical interface name.
func (o GlobalInterfaceOutput) InterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GlobalInterface) *string { return v.InterfaceName }).(pulumi.StringPtrOutput)
}

type GlobalInterfaceArrayOutput struct{ *pulumi.OutputState }

func (GlobalInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GlobalInterface)(nil)).Elem()
}

func (o GlobalInterfaceArrayOutput) ToGlobalInterfaceArrayOutput() GlobalInterfaceArrayOutput {
	return o
}

func (o GlobalInterfaceArrayOutput) ToGlobalInterfaceArrayOutputWithContext(ctx context.Context) GlobalInterfaceArrayOutput {
	return o
}

func (o GlobalInterfaceArrayOutput) Index(i pulumi.IntInput) GlobalInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GlobalInterface {
		return vs[0].([]GlobalInterface)[vs[1].(int)]
	}).(GlobalInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalInterfaceInput)(nil)).Elem(), GlobalInterfaceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalInterfaceArrayInput)(nil)).Elem(), GlobalInterfaceArray{})
	pulumi.RegisterOutputType(GlobalInterfaceOutput{})
	pulumi.RegisterOutputType(GlobalInterfaceArrayOutput{})
}
