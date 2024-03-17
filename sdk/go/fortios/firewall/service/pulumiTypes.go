// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package service

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

var _ = internal.GetEnvOrDefault

type CustomAppCategory struct {
	// Application category id.
	Id *int `pulumi:"id"`
}

// CustomAppCategoryInput is an input type that accepts CustomAppCategoryArgs and CustomAppCategoryOutput values.
// You can construct a concrete instance of `CustomAppCategoryInput` via:
//
//	CustomAppCategoryArgs{...}
type CustomAppCategoryInput interface {
	pulumi.Input

	ToCustomAppCategoryOutput() CustomAppCategoryOutput
	ToCustomAppCategoryOutputWithContext(context.Context) CustomAppCategoryOutput
}

type CustomAppCategoryArgs struct {
	// Application category id.
	Id pulumi.IntPtrInput `pulumi:"id"`
}

func (CustomAppCategoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomAppCategory)(nil)).Elem()
}

func (i CustomAppCategoryArgs) ToCustomAppCategoryOutput() CustomAppCategoryOutput {
	return i.ToCustomAppCategoryOutputWithContext(context.Background())
}

func (i CustomAppCategoryArgs) ToCustomAppCategoryOutputWithContext(ctx context.Context) CustomAppCategoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomAppCategoryOutput)
}

// CustomAppCategoryArrayInput is an input type that accepts CustomAppCategoryArray and CustomAppCategoryArrayOutput values.
// You can construct a concrete instance of `CustomAppCategoryArrayInput` via:
//
//	CustomAppCategoryArray{ CustomAppCategoryArgs{...} }
type CustomAppCategoryArrayInput interface {
	pulumi.Input

	ToCustomAppCategoryArrayOutput() CustomAppCategoryArrayOutput
	ToCustomAppCategoryArrayOutputWithContext(context.Context) CustomAppCategoryArrayOutput
}

type CustomAppCategoryArray []CustomAppCategoryInput

func (CustomAppCategoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomAppCategory)(nil)).Elem()
}

func (i CustomAppCategoryArray) ToCustomAppCategoryArrayOutput() CustomAppCategoryArrayOutput {
	return i.ToCustomAppCategoryArrayOutputWithContext(context.Background())
}

func (i CustomAppCategoryArray) ToCustomAppCategoryArrayOutputWithContext(ctx context.Context) CustomAppCategoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomAppCategoryArrayOutput)
}

type CustomAppCategoryOutput struct{ *pulumi.OutputState }

func (CustomAppCategoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomAppCategory)(nil)).Elem()
}

func (o CustomAppCategoryOutput) ToCustomAppCategoryOutput() CustomAppCategoryOutput {
	return o
}

func (o CustomAppCategoryOutput) ToCustomAppCategoryOutputWithContext(ctx context.Context) CustomAppCategoryOutput {
	return o
}

// Application category id.
func (o CustomAppCategoryOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CustomAppCategory) *int { return v.Id }).(pulumi.IntPtrOutput)
}

type CustomAppCategoryArrayOutput struct{ *pulumi.OutputState }

func (CustomAppCategoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomAppCategory)(nil)).Elem()
}

func (o CustomAppCategoryArrayOutput) ToCustomAppCategoryArrayOutput() CustomAppCategoryArrayOutput {
	return o
}

func (o CustomAppCategoryArrayOutput) ToCustomAppCategoryArrayOutputWithContext(ctx context.Context) CustomAppCategoryArrayOutput {
	return o
}

func (o CustomAppCategoryArrayOutput) Index(i pulumi.IntInput) CustomAppCategoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomAppCategory {
		return vs[0].([]CustomAppCategory)[vs[1].(int)]
	}).(CustomAppCategoryOutput)
}

type CustomApplication struct {
	// Application id.
	Id *int `pulumi:"id"`
}

// CustomApplicationInput is an input type that accepts CustomApplicationArgs and CustomApplicationOutput values.
// You can construct a concrete instance of `CustomApplicationInput` via:
//
//	CustomApplicationArgs{...}
type CustomApplicationInput interface {
	pulumi.Input

	ToCustomApplicationOutput() CustomApplicationOutput
	ToCustomApplicationOutputWithContext(context.Context) CustomApplicationOutput
}

type CustomApplicationArgs struct {
	// Application id.
	Id pulumi.IntPtrInput `pulumi:"id"`
}

func (CustomApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomApplication)(nil)).Elem()
}

func (i CustomApplicationArgs) ToCustomApplicationOutput() CustomApplicationOutput {
	return i.ToCustomApplicationOutputWithContext(context.Background())
}

func (i CustomApplicationArgs) ToCustomApplicationOutputWithContext(ctx context.Context) CustomApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomApplicationOutput)
}

// CustomApplicationArrayInput is an input type that accepts CustomApplicationArray and CustomApplicationArrayOutput values.
// You can construct a concrete instance of `CustomApplicationArrayInput` via:
//
//	CustomApplicationArray{ CustomApplicationArgs{...} }
type CustomApplicationArrayInput interface {
	pulumi.Input

	ToCustomApplicationArrayOutput() CustomApplicationArrayOutput
	ToCustomApplicationArrayOutputWithContext(context.Context) CustomApplicationArrayOutput
}

type CustomApplicationArray []CustomApplicationInput

func (CustomApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomApplication)(nil)).Elem()
}

func (i CustomApplicationArray) ToCustomApplicationArrayOutput() CustomApplicationArrayOutput {
	return i.ToCustomApplicationArrayOutputWithContext(context.Background())
}

func (i CustomApplicationArray) ToCustomApplicationArrayOutputWithContext(ctx context.Context) CustomApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomApplicationArrayOutput)
}

type CustomApplicationOutput struct{ *pulumi.OutputState }

func (CustomApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomApplication)(nil)).Elem()
}

func (o CustomApplicationOutput) ToCustomApplicationOutput() CustomApplicationOutput {
	return o
}

func (o CustomApplicationOutput) ToCustomApplicationOutputWithContext(ctx context.Context) CustomApplicationOutput {
	return o
}

// Application id.
func (o CustomApplicationOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CustomApplication) *int { return v.Id }).(pulumi.IntPtrOutput)
}

type CustomApplicationArrayOutput struct{ *pulumi.OutputState }

func (CustomApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomApplication)(nil)).Elem()
}

func (o CustomApplicationArrayOutput) ToCustomApplicationArrayOutput() CustomApplicationArrayOutput {
	return o
}

func (o CustomApplicationArrayOutput) ToCustomApplicationArrayOutputWithContext(ctx context.Context) CustomApplicationArrayOutput {
	return o
}

func (o CustomApplicationArrayOutput) Index(i pulumi.IntInput) CustomApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomApplication {
		return vs[0].([]CustomApplication)[vs[1].(int)]
	}).(CustomApplicationOutput)
}

type GroupMember struct {
	// Address name.
	Name *string `pulumi:"name"`
}

// GroupMemberInput is an input type that accepts GroupMemberArgs and GroupMemberOutput values.
// You can construct a concrete instance of `GroupMemberInput` via:
//
//	GroupMemberArgs{...}
type GroupMemberInput interface {
	pulumi.Input

	ToGroupMemberOutput() GroupMemberOutput
	ToGroupMemberOutputWithContext(context.Context) GroupMemberOutput
}

type GroupMemberArgs struct {
	// Address name.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GroupMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMember)(nil)).Elem()
}

func (i GroupMemberArgs) ToGroupMemberOutput() GroupMemberOutput {
	return i.ToGroupMemberOutputWithContext(context.Background())
}

func (i GroupMemberArgs) ToGroupMemberOutputWithContext(ctx context.Context) GroupMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberOutput)
}

// GroupMemberArrayInput is an input type that accepts GroupMemberArray and GroupMemberArrayOutput values.
// You can construct a concrete instance of `GroupMemberArrayInput` via:
//
//	GroupMemberArray{ GroupMemberArgs{...} }
type GroupMemberArrayInput interface {
	pulumi.Input

	ToGroupMemberArrayOutput() GroupMemberArrayOutput
	ToGroupMemberArrayOutputWithContext(context.Context) GroupMemberArrayOutput
}

type GroupMemberArray []GroupMemberInput

func (GroupMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupMember)(nil)).Elem()
}

func (i GroupMemberArray) ToGroupMemberArrayOutput() GroupMemberArrayOutput {
	return i.ToGroupMemberArrayOutputWithContext(context.Background())
}

func (i GroupMemberArray) ToGroupMemberArrayOutputWithContext(ctx context.Context) GroupMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberArrayOutput)
}

type GroupMemberOutput struct{ *pulumi.OutputState }

func (GroupMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMember)(nil)).Elem()
}

func (o GroupMemberOutput) ToGroupMemberOutput() GroupMemberOutput {
	return o
}

func (o GroupMemberOutput) ToGroupMemberOutputWithContext(ctx context.Context) GroupMemberOutput {
	return o
}

// Address name.
func (o GroupMemberOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupMember) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type GroupMemberArrayOutput struct{ *pulumi.OutputState }

func (GroupMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupMember)(nil)).Elem()
}

func (o GroupMemberArrayOutput) ToGroupMemberArrayOutput() GroupMemberArrayOutput {
	return o
}

func (o GroupMemberArrayOutput) ToGroupMemberArrayOutputWithContext(ctx context.Context) GroupMemberArrayOutput {
	return o
}

func (o GroupMemberArrayOutput) Index(i pulumi.IntInput) GroupMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupMember {
		return vs[0].([]GroupMember)[vs[1].(int)]
	}).(GroupMemberOutput)
}

type GetCustomAppCategory struct {
	// Application id.
	Id int `pulumi:"id"`
}

// GetCustomAppCategoryInput is an input type that accepts GetCustomAppCategoryArgs and GetCustomAppCategoryOutput values.
// You can construct a concrete instance of `GetCustomAppCategoryInput` via:
//
//	GetCustomAppCategoryArgs{...}
type GetCustomAppCategoryInput interface {
	pulumi.Input

	ToGetCustomAppCategoryOutput() GetCustomAppCategoryOutput
	ToGetCustomAppCategoryOutputWithContext(context.Context) GetCustomAppCategoryOutput
}

type GetCustomAppCategoryArgs struct {
	// Application id.
	Id pulumi.IntInput `pulumi:"id"`
}

func (GetCustomAppCategoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomAppCategory)(nil)).Elem()
}

func (i GetCustomAppCategoryArgs) ToGetCustomAppCategoryOutput() GetCustomAppCategoryOutput {
	return i.ToGetCustomAppCategoryOutputWithContext(context.Background())
}

func (i GetCustomAppCategoryArgs) ToGetCustomAppCategoryOutputWithContext(ctx context.Context) GetCustomAppCategoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCustomAppCategoryOutput)
}

// GetCustomAppCategoryArrayInput is an input type that accepts GetCustomAppCategoryArray and GetCustomAppCategoryArrayOutput values.
// You can construct a concrete instance of `GetCustomAppCategoryArrayInput` via:
//
//	GetCustomAppCategoryArray{ GetCustomAppCategoryArgs{...} }
type GetCustomAppCategoryArrayInput interface {
	pulumi.Input

	ToGetCustomAppCategoryArrayOutput() GetCustomAppCategoryArrayOutput
	ToGetCustomAppCategoryArrayOutputWithContext(context.Context) GetCustomAppCategoryArrayOutput
}

type GetCustomAppCategoryArray []GetCustomAppCategoryInput

func (GetCustomAppCategoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCustomAppCategory)(nil)).Elem()
}

func (i GetCustomAppCategoryArray) ToGetCustomAppCategoryArrayOutput() GetCustomAppCategoryArrayOutput {
	return i.ToGetCustomAppCategoryArrayOutputWithContext(context.Background())
}

func (i GetCustomAppCategoryArray) ToGetCustomAppCategoryArrayOutputWithContext(ctx context.Context) GetCustomAppCategoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCustomAppCategoryArrayOutput)
}

type GetCustomAppCategoryOutput struct{ *pulumi.OutputState }

func (GetCustomAppCategoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomAppCategory)(nil)).Elem()
}

func (o GetCustomAppCategoryOutput) ToGetCustomAppCategoryOutput() GetCustomAppCategoryOutput {
	return o
}

func (o GetCustomAppCategoryOutput) ToGetCustomAppCategoryOutputWithContext(ctx context.Context) GetCustomAppCategoryOutput {
	return o
}

// Application id.
func (o GetCustomAppCategoryOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetCustomAppCategory) int { return v.Id }).(pulumi.IntOutput)
}

type GetCustomAppCategoryArrayOutput struct{ *pulumi.OutputState }

func (GetCustomAppCategoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCustomAppCategory)(nil)).Elem()
}

func (o GetCustomAppCategoryArrayOutput) ToGetCustomAppCategoryArrayOutput() GetCustomAppCategoryArrayOutput {
	return o
}

func (o GetCustomAppCategoryArrayOutput) ToGetCustomAppCategoryArrayOutputWithContext(ctx context.Context) GetCustomAppCategoryArrayOutput {
	return o
}

func (o GetCustomAppCategoryArrayOutput) Index(i pulumi.IntInput) GetCustomAppCategoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCustomAppCategory {
		return vs[0].([]GetCustomAppCategory)[vs[1].(int)]
	}).(GetCustomAppCategoryOutput)
}

type GetCustomApplication struct {
	// Application id.
	Id int `pulumi:"id"`
}

// GetCustomApplicationInput is an input type that accepts GetCustomApplicationArgs and GetCustomApplicationOutput values.
// You can construct a concrete instance of `GetCustomApplicationInput` via:
//
//	GetCustomApplicationArgs{...}
type GetCustomApplicationInput interface {
	pulumi.Input

	ToGetCustomApplicationOutput() GetCustomApplicationOutput
	ToGetCustomApplicationOutputWithContext(context.Context) GetCustomApplicationOutput
}

type GetCustomApplicationArgs struct {
	// Application id.
	Id pulumi.IntInput `pulumi:"id"`
}

func (GetCustomApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomApplication)(nil)).Elem()
}

func (i GetCustomApplicationArgs) ToGetCustomApplicationOutput() GetCustomApplicationOutput {
	return i.ToGetCustomApplicationOutputWithContext(context.Background())
}

func (i GetCustomApplicationArgs) ToGetCustomApplicationOutputWithContext(ctx context.Context) GetCustomApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCustomApplicationOutput)
}

// GetCustomApplicationArrayInput is an input type that accepts GetCustomApplicationArray and GetCustomApplicationArrayOutput values.
// You can construct a concrete instance of `GetCustomApplicationArrayInput` via:
//
//	GetCustomApplicationArray{ GetCustomApplicationArgs{...} }
type GetCustomApplicationArrayInput interface {
	pulumi.Input

	ToGetCustomApplicationArrayOutput() GetCustomApplicationArrayOutput
	ToGetCustomApplicationArrayOutputWithContext(context.Context) GetCustomApplicationArrayOutput
}

type GetCustomApplicationArray []GetCustomApplicationInput

func (GetCustomApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCustomApplication)(nil)).Elem()
}

func (i GetCustomApplicationArray) ToGetCustomApplicationArrayOutput() GetCustomApplicationArrayOutput {
	return i.ToGetCustomApplicationArrayOutputWithContext(context.Background())
}

func (i GetCustomApplicationArray) ToGetCustomApplicationArrayOutputWithContext(ctx context.Context) GetCustomApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCustomApplicationArrayOutput)
}

type GetCustomApplicationOutput struct{ *pulumi.OutputState }

func (GetCustomApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomApplication)(nil)).Elem()
}

func (o GetCustomApplicationOutput) ToGetCustomApplicationOutput() GetCustomApplicationOutput {
	return o
}

func (o GetCustomApplicationOutput) ToGetCustomApplicationOutputWithContext(ctx context.Context) GetCustomApplicationOutput {
	return o
}

// Application id.
func (o GetCustomApplicationOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetCustomApplication) int { return v.Id }).(pulumi.IntOutput)
}

type GetCustomApplicationArrayOutput struct{ *pulumi.OutputState }

func (GetCustomApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCustomApplication)(nil)).Elem()
}

func (o GetCustomApplicationArrayOutput) ToGetCustomApplicationArrayOutput() GetCustomApplicationArrayOutput {
	return o
}

func (o GetCustomApplicationArrayOutput) ToGetCustomApplicationArrayOutputWithContext(ctx context.Context) GetCustomApplicationArrayOutput {
	return o
}

func (o GetCustomApplicationArrayOutput) Index(i pulumi.IntInput) GetCustomApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCustomApplication {
		return vs[0].([]GetCustomApplication)[vs[1].(int)]
	}).(GetCustomApplicationOutput)
}

type GetGroupMember struct {
	// Specify the name of the desired firewallservice group.
	Name string `pulumi:"name"`
}

// GetGroupMemberInput is an input type that accepts GetGroupMemberArgs and GetGroupMemberOutput values.
// You can construct a concrete instance of `GetGroupMemberInput` via:
//
//	GetGroupMemberArgs{...}
type GetGroupMemberInput interface {
	pulumi.Input

	ToGetGroupMemberOutput() GetGroupMemberOutput
	ToGetGroupMemberOutputWithContext(context.Context) GetGroupMemberOutput
}

type GetGroupMemberArgs struct {
	// Specify the name of the desired firewallservice group.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetGroupMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupMember)(nil)).Elem()
}

func (i GetGroupMemberArgs) ToGetGroupMemberOutput() GetGroupMemberOutput {
	return i.ToGetGroupMemberOutputWithContext(context.Background())
}

func (i GetGroupMemberArgs) ToGetGroupMemberOutputWithContext(ctx context.Context) GetGroupMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupMemberOutput)
}

// GetGroupMemberArrayInput is an input type that accepts GetGroupMemberArray and GetGroupMemberArrayOutput values.
// You can construct a concrete instance of `GetGroupMemberArrayInput` via:
//
//	GetGroupMemberArray{ GetGroupMemberArgs{...} }
type GetGroupMemberArrayInput interface {
	pulumi.Input

	ToGetGroupMemberArrayOutput() GetGroupMemberArrayOutput
	ToGetGroupMemberArrayOutputWithContext(context.Context) GetGroupMemberArrayOutput
}

type GetGroupMemberArray []GetGroupMemberInput

func (GetGroupMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupMember)(nil)).Elem()
}

func (i GetGroupMemberArray) ToGetGroupMemberArrayOutput() GetGroupMemberArrayOutput {
	return i.ToGetGroupMemberArrayOutputWithContext(context.Background())
}

func (i GetGroupMemberArray) ToGetGroupMemberArrayOutputWithContext(ctx context.Context) GetGroupMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupMemberArrayOutput)
}

type GetGroupMemberOutput struct{ *pulumi.OutputState }

func (GetGroupMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupMember)(nil)).Elem()
}

func (o GetGroupMemberOutput) ToGetGroupMemberOutput() GetGroupMemberOutput {
	return o
}

func (o GetGroupMemberOutput) ToGetGroupMemberOutputWithContext(ctx context.Context) GetGroupMemberOutput {
	return o
}

// Specify the name of the desired firewallservice group.
func (o GetGroupMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupMember) string { return v.Name }).(pulumi.StringOutput)
}

type GetGroupMemberArrayOutput struct{ *pulumi.OutputState }

func (GetGroupMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupMember)(nil)).Elem()
}

func (o GetGroupMemberArrayOutput) ToGetGroupMemberArrayOutput() GetGroupMemberArrayOutput {
	return o
}

func (o GetGroupMemberArrayOutput) ToGetGroupMemberArrayOutputWithContext(ctx context.Context) GetGroupMemberArrayOutput {
	return o
}

func (o GetGroupMemberArrayOutput) Index(i pulumi.IntInput) GetGroupMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetGroupMember {
		return vs[0].([]GetGroupMember)[vs[1].(int)]
	}).(GetGroupMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomAppCategoryInput)(nil)).Elem(), CustomAppCategoryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomAppCategoryArrayInput)(nil)).Elem(), CustomAppCategoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomApplicationInput)(nil)).Elem(), CustomApplicationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomApplicationArrayInput)(nil)).Elem(), CustomApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMemberInput)(nil)).Elem(), GroupMemberArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMemberArrayInput)(nil)).Elem(), GroupMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCustomAppCategoryInput)(nil)).Elem(), GetCustomAppCategoryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCustomAppCategoryArrayInput)(nil)).Elem(), GetCustomAppCategoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCustomApplicationInput)(nil)).Elem(), GetCustomApplicationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCustomApplicationArrayInput)(nil)).Elem(), GetCustomApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGroupMemberInput)(nil)).Elem(), GetGroupMemberArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGroupMemberArrayInput)(nil)).Elem(), GetGroupMemberArray{})
	pulumi.RegisterOutputType(CustomAppCategoryOutput{})
	pulumi.RegisterOutputType(CustomAppCategoryArrayOutput{})
	pulumi.RegisterOutputType(CustomApplicationOutput{})
	pulumi.RegisterOutputType(CustomApplicationArrayOutput{})
	pulumi.RegisterOutputType(GroupMemberOutput{})
	pulumi.RegisterOutputType(GroupMemberArrayOutput{})
	pulumi.RegisterOutputType(GetCustomAppCategoryOutput{})
	pulumi.RegisterOutputType(GetCustomAppCategoryArrayOutput{})
	pulumi.RegisterOutputType(GetCustomApplicationOutput{})
	pulumi.RegisterOutputType(GetCustomApplicationArrayOutput{})
	pulumi.RegisterOutputType(GetGroupMemberOutput{})
	pulumi.RegisterOutputType(GetGroupMemberArrayOutput{})
}
