// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

var _ = internal.GetEnvOrDefault

type GetProxypolicyPoolname struct {
	// Group name.
	Name string `pulumi:"name"`
}

// GetProxypolicyPoolnameInput is an input type that accepts GetProxypolicyPoolnameArgs and GetProxypolicyPoolnameOutput values.
// You can construct a concrete instance of `GetProxypolicyPoolnameInput` via:
//
//	GetProxypolicyPoolnameArgs{...}
type GetProxypolicyPoolnameInput interface {
	pulumi.Input

	ToGetProxypolicyPoolnameOutput() GetProxypolicyPoolnameOutput
	ToGetProxypolicyPoolnameOutputWithContext(context.Context) GetProxypolicyPoolnameOutput
}

type GetProxypolicyPoolnameArgs struct {
	// Group name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetProxypolicyPoolnameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicyPoolname)(nil)).Elem()
}

func (i GetProxypolicyPoolnameArgs) ToGetProxypolicyPoolnameOutput() GetProxypolicyPoolnameOutput {
	return i.ToGetProxypolicyPoolnameOutputWithContext(context.Background())
}

func (i GetProxypolicyPoolnameArgs) ToGetProxypolicyPoolnameOutputWithContext(ctx context.Context) GetProxypolicyPoolnameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicyPoolnameOutput)
}

// GetProxypolicyPoolnameArrayInput is an input type that accepts GetProxypolicyPoolnameArray and GetProxypolicyPoolnameArrayOutput values.
// You can construct a concrete instance of `GetProxypolicyPoolnameArrayInput` via:
//
//	GetProxypolicyPoolnameArray{ GetProxypolicyPoolnameArgs{...} }
type GetProxypolicyPoolnameArrayInput interface {
	pulumi.Input

	ToGetProxypolicyPoolnameArrayOutput() GetProxypolicyPoolnameArrayOutput
	ToGetProxypolicyPoolnameArrayOutputWithContext(context.Context) GetProxypolicyPoolnameArrayOutput
}

type GetProxypolicyPoolnameArray []GetProxypolicyPoolnameInput

func (GetProxypolicyPoolnameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicyPoolname)(nil)).Elem()
}

func (i GetProxypolicyPoolnameArray) ToGetProxypolicyPoolnameArrayOutput() GetProxypolicyPoolnameArrayOutput {
	return i.ToGetProxypolicyPoolnameArrayOutputWithContext(context.Background())
}

func (i GetProxypolicyPoolnameArray) ToGetProxypolicyPoolnameArrayOutputWithContext(ctx context.Context) GetProxypolicyPoolnameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicyPoolnameArrayOutput)
}

type GetProxypolicyPoolnameOutput struct{ *pulumi.OutputState }

func (GetProxypolicyPoolnameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicyPoolname)(nil)).Elem()
}

func (o GetProxypolicyPoolnameOutput) ToGetProxypolicyPoolnameOutput() GetProxypolicyPoolnameOutput {
	return o
}

func (o GetProxypolicyPoolnameOutput) ToGetProxypolicyPoolnameOutputWithContext(ctx context.Context) GetProxypolicyPoolnameOutput {
	return o
}

// Group name.
func (o GetProxypolicyPoolnameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxypolicyPoolname) string { return v.Name }).(pulumi.StringOutput)
}

type GetProxypolicyPoolnameArrayOutput struct{ *pulumi.OutputState }

func (GetProxypolicyPoolnameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicyPoolname)(nil)).Elem()
}

func (o GetProxypolicyPoolnameArrayOutput) ToGetProxypolicyPoolnameArrayOutput() GetProxypolicyPoolnameArrayOutput {
	return o
}

func (o GetProxypolicyPoolnameArrayOutput) ToGetProxypolicyPoolnameArrayOutputWithContext(ctx context.Context) GetProxypolicyPoolnameArrayOutput {
	return o
}

func (o GetProxypolicyPoolnameArrayOutput) Index(i pulumi.IntInput) GetProxypolicyPoolnameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProxypolicyPoolname {
		return vs[0].([]GetProxypolicyPoolname)[vs[1].(int)]
	}).(GetProxypolicyPoolnameOutput)
}

type GetProxypolicyService struct {
	// Group name.
	Name string `pulumi:"name"`
}

// GetProxypolicyServiceInput is an input type that accepts GetProxypolicyServiceArgs and GetProxypolicyServiceOutput values.
// You can construct a concrete instance of `GetProxypolicyServiceInput` via:
//
//	GetProxypolicyServiceArgs{...}
type GetProxypolicyServiceInput interface {
	pulumi.Input

	ToGetProxypolicyServiceOutput() GetProxypolicyServiceOutput
	ToGetProxypolicyServiceOutputWithContext(context.Context) GetProxypolicyServiceOutput
}

type GetProxypolicyServiceArgs struct {
	// Group name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetProxypolicyServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicyService)(nil)).Elem()
}

func (i GetProxypolicyServiceArgs) ToGetProxypolicyServiceOutput() GetProxypolicyServiceOutput {
	return i.ToGetProxypolicyServiceOutputWithContext(context.Background())
}

func (i GetProxypolicyServiceArgs) ToGetProxypolicyServiceOutputWithContext(ctx context.Context) GetProxypolicyServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicyServiceOutput)
}

// GetProxypolicyServiceArrayInput is an input type that accepts GetProxypolicyServiceArray and GetProxypolicyServiceArrayOutput values.
// You can construct a concrete instance of `GetProxypolicyServiceArrayInput` via:
//
//	GetProxypolicyServiceArray{ GetProxypolicyServiceArgs{...} }
type GetProxypolicyServiceArrayInput interface {
	pulumi.Input

	ToGetProxypolicyServiceArrayOutput() GetProxypolicyServiceArrayOutput
	ToGetProxypolicyServiceArrayOutputWithContext(context.Context) GetProxypolicyServiceArrayOutput
}

type GetProxypolicyServiceArray []GetProxypolicyServiceInput

func (GetProxypolicyServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicyService)(nil)).Elem()
}

func (i GetProxypolicyServiceArray) ToGetProxypolicyServiceArrayOutput() GetProxypolicyServiceArrayOutput {
	return i.ToGetProxypolicyServiceArrayOutputWithContext(context.Background())
}

func (i GetProxypolicyServiceArray) ToGetProxypolicyServiceArrayOutputWithContext(ctx context.Context) GetProxypolicyServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicyServiceArrayOutput)
}

type GetProxypolicyServiceOutput struct{ *pulumi.OutputState }

func (GetProxypolicyServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicyService)(nil)).Elem()
}

func (o GetProxypolicyServiceOutput) ToGetProxypolicyServiceOutput() GetProxypolicyServiceOutput {
	return o
}

func (o GetProxypolicyServiceOutput) ToGetProxypolicyServiceOutputWithContext(ctx context.Context) GetProxypolicyServiceOutput {
	return o
}

// Group name.
func (o GetProxypolicyServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxypolicyService) string { return v.Name }).(pulumi.StringOutput)
}

type GetProxypolicyServiceArrayOutput struct{ *pulumi.OutputState }

func (GetProxypolicyServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicyService)(nil)).Elem()
}

func (o GetProxypolicyServiceArrayOutput) ToGetProxypolicyServiceArrayOutput() GetProxypolicyServiceArrayOutput {
	return o
}

func (o GetProxypolicyServiceArrayOutput) ToGetProxypolicyServiceArrayOutputWithContext(ctx context.Context) GetProxypolicyServiceArrayOutput {
	return o
}

func (o GetProxypolicyServiceArrayOutput) Index(i pulumi.IntInput) GetProxypolicyServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProxypolicyService {
		return vs[0].([]GetProxypolicyService)[vs[1].(int)]
	}).(GetProxypolicyServiceOutput)
}

type GetProxypolicySrcaddr6 struct {
	// Group name.
	Name string `pulumi:"name"`
}

// GetProxypolicySrcaddr6Input is an input type that accepts GetProxypolicySrcaddr6Args and GetProxypolicySrcaddr6Output values.
// You can construct a concrete instance of `GetProxypolicySrcaddr6Input` via:
//
//	GetProxypolicySrcaddr6Args{...}
type GetProxypolicySrcaddr6Input interface {
	pulumi.Input

	ToGetProxypolicySrcaddr6Output() GetProxypolicySrcaddr6Output
	ToGetProxypolicySrcaddr6OutputWithContext(context.Context) GetProxypolicySrcaddr6Output
}

type GetProxypolicySrcaddr6Args struct {
	// Group name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetProxypolicySrcaddr6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicySrcaddr6)(nil)).Elem()
}

func (i GetProxypolicySrcaddr6Args) ToGetProxypolicySrcaddr6Output() GetProxypolicySrcaddr6Output {
	return i.ToGetProxypolicySrcaddr6OutputWithContext(context.Background())
}

func (i GetProxypolicySrcaddr6Args) ToGetProxypolicySrcaddr6OutputWithContext(ctx context.Context) GetProxypolicySrcaddr6Output {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicySrcaddr6Output)
}

// GetProxypolicySrcaddr6ArrayInput is an input type that accepts GetProxypolicySrcaddr6Array and GetProxypolicySrcaddr6ArrayOutput values.
// You can construct a concrete instance of `GetProxypolicySrcaddr6ArrayInput` via:
//
//	GetProxypolicySrcaddr6Array{ GetProxypolicySrcaddr6Args{...} }
type GetProxypolicySrcaddr6ArrayInput interface {
	pulumi.Input

	ToGetProxypolicySrcaddr6ArrayOutput() GetProxypolicySrcaddr6ArrayOutput
	ToGetProxypolicySrcaddr6ArrayOutputWithContext(context.Context) GetProxypolicySrcaddr6ArrayOutput
}

type GetProxypolicySrcaddr6Array []GetProxypolicySrcaddr6Input

func (GetProxypolicySrcaddr6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicySrcaddr6)(nil)).Elem()
}

func (i GetProxypolicySrcaddr6Array) ToGetProxypolicySrcaddr6ArrayOutput() GetProxypolicySrcaddr6ArrayOutput {
	return i.ToGetProxypolicySrcaddr6ArrayOutputWithContext(context.Background())
}

func (i GetProxypolicySrcaddr6Array) ToGetProxypolicySrcaddr6ArrayOutputWithContext(ctx context.Context) GetProxypolicySrcaddr6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicySrcaddr6ArrayOutput)
}

type GetProxypolicySrcaddr6Output struct{ *pulumi.OutputState }

func (GetProxypolicySrcaddr6Output) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicySrcaddr6)(nil)).Elem()
}

func (o GetProxypolicySrcaddr6Output) ToGetProxypolicySrcaddr6Output() GetProxypolicySrcaddr6Output {
	return o
}

func (o GetProxypolicySrcaddr6Output) ToGetProxypolicySrcaddr6OutputWithContext(ctx context.Context) GetProxypolicySrcaddr6Output {
	return o
}

// Group name.
func (o GetProxypolicySrcaddr6Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxypolicySrcaddr6) string { return v.Name }).(pulumi.StringOutput)
}

type GetProxypolicySrcaddr6ArrayOutput struct{ *pulumi.OutputState }

func (GetProxypolicySrcaddr6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicySrcaddr6)(nil)).Elem()
}

func (o GetProxypolicySrcaddr6ArrayOutput) ToGetProxypolicySrcaddr6ArrayOutput() GetProxypolicySrcaddr6ArrayOutput {
	return o
}

func (o GetProxypolicySrcaddr6ArrayOutput) ToGetProxypolicySrcaddr6ArrayOutputWithContext(ctx context.Context) GetProxypolicySrcaddr6ArrayOutput {
	return o
}

func (o GetProxypolicySrcaddr6ArrayOutput) Index(i pulumi.IntInput) GetProxypolicySrcaddr6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProxypolicySrcaddr6 {
		return vs[0].([]GetProxypolicySrcaddr6)[vs[1].(int)]
	}).(GetProxypolicySrcaddr6Output)
}

type GetProxypolicySrcaddr struct {
	// Group name.
	Name string `pulumi:"name"`
}

// GetProxypolicySrcaddrInput is an input type that accepts GetProxypolicySrcaddrArgs and GetProxypolicySrcaddrOutput values.
// You can construct a concrete instance of `GetProxypolicySrcaddrInput` via:
//
//	GetProxypolicySrcaddrArgs{...}
type GetProxypolicySrcaddrInput interface {
	pulumi.Input

	ToGetProxypolicySrcaddrOutput() GetProxypolicySrcaddrOutput
	ToGetProxypolicySrcaddrOutputWithContext(context.Context) GetProxypolicySrcaddrOutput
}

type GetProxypolicySrcaddrArgs struct {
	// Group name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetProxypolicySrcaddrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicySrcaddr)(nil)).Elem()
}

func (i GetProxypolicySrcaddrArgs) ToGetProxypolicySrcaddrOutput() GetProxypolicySrcaddrOutput {
	return i.ToGetProxypolicySrcaddrOutputWithContext(context.Background())
}

func (i GetProxypolicySrcaddrArgs) ToGetProxypolicySrcaddrOutputWithContext(ctx context.Context) GetProxypolicySrcaddrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicySrcaddrOutput)
}

// GetProxypolicySrcaddrArrayInput is an input type that accepts GetProxypolicySrcaddrArray and GetProxypolicySrcaddrArrayOutput values.
// You can construct a concrete instance of `GetProxypolicySrcaddrArrayInput` via:
//
//	GetProxypolicySrcaddrArray{ GetProxypolicySrcaddrArgs{...} }
type GetProxypolicySrcaddrArrayInput interface {
	pulumi.Input

	ToGetProxypolicySrcaddrArrayOutput() GetProxypolicySrcaddrArrayOutput
	ToGetProxypolicySrcaddrArrayOutputWithContext(context.Context) GetProxypolicySrcaddrArrayOutput
}

type GetProxypolicySrcaddrArray []GetProxypolicySrcaddrInput

func (GetProxypolicySrcaddrArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicySrcaddr)(nil)).Elem()
}

func (i GetProxypolicySrcaddrArray) ToGetProxypolicySrcaddrArrayOutput() GetProxypolicySrcaddrArrayOutput {
	return i.ToGetProxypolicySrcaddrArrayOutputWithContext(context.Background())
}

func (i GetProxypolicySrcaddrArray) ToGetProxypolicySrcaddrArrayOutputWithContext(ctx context.Context) GetProxypolicySrcaddrArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicySrcaddrArrayOutput)
}

type GetProxypolicySrcaddrOutput struct{ *pulumi.OutputState }

func (GetProxypolicySrcaddrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicySrcaddr)(nil)).Elem()
}

func (o GetProxypolicySrcaddrOutput) ToGetProxypolicySrcaddrOutput() GetProxypolicySrcaddrOutput {
	return o
}

func (o GetProxypolicySrcaddrOutput) ToGetProxypolicySrcaddrOutputWithContext(ctx context.Context) GetProxypolicySrcaddrOutput {
	return o
}

// Group name.
func (o GetProxypolicySrcaddrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxypolicySrcaddr) string { return v.Name }).(pulumi.StringOutput)
}

type GetProxypolicySrcaddrArrayOutput struct{ *pulumi.OutputState }

func (GetProxypolicySrcaddrArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicySrcaddr)(nil)).Elem()
}

func (o GetProxypolicySrcaddrArrayOutput) ToGetProxypolicySrcaddrArrayOutput() GetProxypolicySrcaddrArrayOutput {
	return o
}

func (o GetProxypolicySrcaddrArrayOutput) ToGetProxypolicySrcaddrArrayOutputWithContext(ctx context.Context) GetProxypolicySrcaddrArrayOutput {
	return o
}

func (o GetProxypolicySrcaddrArrayOutput) Index(i pulumi.IntInput) GetProxypolicySrcaddrOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProxypolicySrcaddr {
		return vs[0].([]GetProxypolicySrcaddr)[vs[1].(int)]
	}).(GetProxypolicySrcaddrOutput)
}

type GetProxypolicySrcintf struct {
	// Group name.
	Name string `pulumi:"name"`
}

// GetProxypolicySrcintfInput is an input type that accepts GetProxypolicySrcintfArgs and GetProxypolicySrcintfOutput values.
// You can construct a concrete instance of `GetProxypolicySrcintfInput` via:
//
//	GetProxypolicySrcintfArgs{...}
type GetProxypolicySrcintfInput interface {
	pulumi.Input

	ToGetProxypolicySrcintfOutput() GetProxypolicySrcintfOutput
	ToGetProxypolicySrcintfOutputWithContext(context.Context) GetProxypolicySrcintfOutput
}

type GetProxypolicySrcintfArgs struct {
	// Group name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetProxypolicySrcintfArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicySrcintf)(nil)).Elem()
}

func (i GetProxypolicySrcintfArgs) ToGetProxypolicySrcintfOutput() GetProxypolicySrcintfOutput {
	return i.ToGetProxypolicySrcintfOutputWithContext(context.Background())
}

func (i GetProxypolicySrcintfArgs) ToGetProxypolicySrcintfOutputWithContext(ctx context.Context) GetProxypolicySrcintfOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicySrcintfOutput)
}

// GetProxypolicySrcintfArrayInput is an input type that accepts GetProxypolicySrcintfArray and GetProxypolicySrcintfArrayOutput values.
// You can construct a concrete instance of `GetProxypolicySrcintfArrayInput` via:
//
//	GetProxypolicySrcintfArray{ GetProxypolicySrcintfArgs{...} }
type GetProxypolicySrcintfArrayInput interface {
	pulumi.Input

	ToGetProxypolicySrcintfArrayOutput() GetProxypolicySrcintfArrayOutput
	ToGetProxypolicySrcintfArrayOutputWithContext(context.Context) GetProxypolicySrcintfArrayOutput
}

type GetProxypolicySrcintfArray []GetProxypolicySrcintfInput

func (GetProxypolicySrcintfArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicySrcintf)(nil)).Elem()
}

func (i GetProxypolicySrcintfArray) ToGetProxypolicySrcintfArrayOutput() GetProxypolicySrcintfArrayOutput {
	return i.ToGetProxypolicySrcintfArrayOutputWithContext(context.Background())
}

func (i GetProxypolicySrcintfArray) ToGetProxypolicySrcintfArrayOutputWithContext(ctx context.Context) GetProxypolicySrcintfArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicySrcintfArrayOutput)
}

type GetProxypolicySrcintfOutput struct{ *pulumi.OutputState }

func (GetProxypolicySrcintfOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicySrcintf)(nil)).Elem()
}

func (o GetProxypolicySrcintfOutput) ToGetProxypolicySrcintfOutput() GetProxypolicySrcintfOutput {
	return o
}

func (o GetProxypolicySrcintfOutput) ToGetProxypolicySrcintfOutputWithContext(ctx context.Context) GetProxypolicySrcintfOutput {
	return o
}

// Group name.
func (o GetProxypolicySrcintfOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxypolicySrcintf) string { return v.Name }).(pulumi.StringOutput)
}

type GetProxypolicySrcintfArrayOutput struct{ *pulumi.OutputState }

func (GetProxypolicySrcintfArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicySrcintf)(nil)).Elem()
}

func (o GetProxypolicySrcintfArrayOutput) ToGetProxypolicySrcintfArrayOutput() GetProxypolicySrcintfArrayOutput {
	return o
}

func (o GetProxypolicySrcintfArrayOutput) ToGetProxypolicySrcintfArrayOutputWithContext(ctx context.Context) GetProxypolicySrcintfArrayOutput {
	return o
}

func (o GetProxypolicySrcintfArrayOutput) Index(i pulumi.IntInput) GetProxypolicySrcintfOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProxypolicySrcintf {
		return vs[0].([]GetProxypolicySrcintf)[vs[1].(int)]
	}).(GetProxypolicySrcintfOutput)
}

type GetProxypolicyUser struct {
	// Group name.
	Name string `pulumi:"name"`
}

// GetProxypolicyUserInput is an input type that accepts GetProxypolicyUserArgs and GetProxypolicyUserOutput values.
// You can construct a concrete instance of `GetProxypolicyUserInput` via:
//
//	GetProxypolicyUserArgs{...}
type GetProxypolicyUserInput interface {
	pulumi.Input

	ToGetProxypolicyUserOutput() GetProxypolicyUserOutput
	ToGetProxypolicyUserOutputWithContext(context.Context) GetProxypolicyUserOutput
}

type GetProxypolicyUserArgs struct {
	// Group name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetProxypolicyUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicyUser)(nil)).Elem()
}

func (i GetProxypolicyUserArgs) ToGetProxypolicyUserOutput() GetProxypolicyUserOutput {
	return i.ToGetProxypolicyUserOutputWithContext(context.Background())
}

func (i GetProxypolicyUserArgs) ToGetProxypolicyUserOutputWithContext(ctx context.Context) GetProxypolicyUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicyUserOutput)
}

// GetProxypolicyUserArrayInput is an input type that accepts GetProxypolicyUserArray and GetProxypolicyUserArrayOutput values.
// You can construct a concrete instance of `GetProxypolicyUserArrayInput` via:
//
//	GetProxypolicyUserArray{ GetProxypolicyUserArgs{...} }
type GetProxypolicyUserArrayInput interface {
	pulumi.Input

	ToGetProxypolicyUserArrayOutput() GetProxypolicyUserArrayOutput
	ToGetProxypolicyUserArrayOutputWithContext(context.Context) GetProxypolicyUserArrayOutput
}

type GetProxypolicyUserArray []GetProxypolicyUserInput

func (GetProxypolicyUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicyUser)(nil)).Elem()
}

func (i GetProxypolicyUserArray) ToGetProxypolicyUserArrayOutput() GetProxypolicyUserArrayOutput {
	return i.ToGetProxypolicyUserArrayOutputWithContext(context.Background())
}

func (i GetProxypolicyUserArray) ToGetProxypolicyUserArrayOutputWithContext(ctx context.Context) GetProxypolicyUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicyUserArrayOutput)
}

type GetProxypolicyUserOutput struct{ *pulumi.OutputState }

func (GetProxypolicyUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicyUser)(nil)).Elem()
}

func (o GetProxypolicyUserOutput) ToGetProxypolicyUserOutput() GetProxypolicyUserOutput {
	return o
}

func (o GetProxypolicyUserOutput) ToGetProxypolicyUserOutputWithContext(ctx context.Context) GetProxypolicyUserOutput {
	return o
}

// Group name.
func (o GetProxypolicyUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxypolicyUser) string { return v.Name }).(pulumi.StringOutput)
}

type GetProxypolicyUserArrayOutput struct{ *pulumi.OutputState }

func (GetProxypolicyUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicyUser)(nil)).Elem()
}

func (o GetProxypolicyUserArrayOutput) ToGetProxypolicyUserArrayOutput() GetProxypolicyUserArrayOutput {
	return o
}

func (o GetProxypolicyUserArrayOutput) ToGetProxypolicyUserArrayOutputWithContext(ctx context.Context) GetProxypolicyUserArrayOutput {
	return o
}

func (o GetProxypolicyUserArrayOutput) Index(i pulumi.IntInput) GetProxypolicyUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProxypolicyUser {
		return vs[0].([]GetProxypolicyUser)[vs[1].(int)]
	}).(GetProxypolicyUserOutput)
}

type GetProxypolicyZtnaEmsTag struct {
	// Group name.
	Name string `pulumi:"name"`
}

// GetProxypolicyZtnaEmsTagInput is an input type that accepts GetProxypolicyZtnaEmsTagArgs and GetProxypolicyZtnaEmsTagOutput values.
// You can construct a concrete instance of `GetProxypolicyZtnaEmsTagInput` via:
//
//	GetProxypolicyZtnaEmsTagArgs{...}
type GetProxypolicyZtnaEmsTagInput interface {
	pulumi.Input

	ToGetProxypolicyZtnaEmsTagOutput() GetProxypolicyZtnaEmsTagOutput
	ToGetProxypolicyZtnaEmsTagOutputWithContext(context.Context) GetProxypolicyZtnaEmsTagOutput
}

type GetProxypolicyZtnaEmsTagArgs struct {
	// Group name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetProxypolicyZtnaEmsTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicyZtnaEmsTag)(nil)).Elem()
}

func (i GetProxypolicyZtnaEmsTagArgs) ToGetProxypolicyZtnaEmsTagOutput() GetProxypolicyZtnaEmsTagOutput {
	return i.ToGetProxypolicyZtnaEmsTagOutputWithContext(context.Background())
}

func (i GetProxypolicyZtnaEmsTagArgs) ToGetProxypolicyZtnaEmsTagOutputWithContext(ctx context.Context) GetProxypolicyZtnaEmsTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicyZtnaEmsTagOutput)
}

// GetProxypolicyZtnaEmsTagArrayInput is an input type that accepts GetProxypolicyZtnaEmsTagArray and GetProxypolicyZtnaEmsTagArrayOutput values.
// You can construct a concrete instance of `GetProxypolicyZtnaEmsTagArrayInput` via:
//
//	GetProxypolicyZtnaEmsTagArray{ GetProxypolicyZtnaEmsTagArgs{...} }
type GetProxypolicyZtnaEmsTagArrayInput interface {
	pulumi.Input

	ToGetProxypolicyZtnaEmsTagArrayOutput() GetProxypolicyZtnaEmsTagArrayOutput
	ToGetProxypolicyZtnaEmsTagArrayOutputWithContext(context.Context) GetProxypolicyZtnaEmsTagArrayOutput
}

type GetProxypolicyZtnaEmsTagArray []GetProxypolicyZtnaEmsTagInput

func (GetProxypolicyZtnaEmsTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicyZtnaEmsTag)(nil)).Elem()
}

func (i GetProxypolicyZtnaEmsTagArray) ToGetProxypolicyZtnaEmsTagArrayOutput() GetProxypolicyZtnaEmsTagArrayOutput {
	return i.ToGetProxypolicyZtnaEmsTagArrayOutputWithContext(context.Background())
}

func (i GetProxypolicyZtnaEmsTagArray) ToGetProxypolicyZtnaEmsTagArrayOutputWithContext(ctx context.Context) GetProxypolicyZtnaEmsTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProxypolicyZtnaEmsTagArrayOutput)
}

type GetProxypolicyZtnaEmsTagOutput struct{ *pulumi.OutputState }

func (GetProxypolicyZtnaEmsTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxypolicyZtnaEmsTag)(nil)).Elem()
}

func (o GetProxypolicyZtnaEmsTagOutput) ToGetProxypolicyZtnaEmsTagOutput() GetProxypolicyZtnaEmsTagOutput {
	return o
}

func (o GetProxypolicyZtnaEmsTagOutput) ToGetProxypolicyZtnaEmsTagOutputWithContext(ctx context.Context) GetProxypolicyZtnaEmsTagOutput {
	return o
}

// Group name.
func (o GetProxypolicyZtnaEmsTagOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxypolicyZtnaEmsTag) string { return v.Name }).(pulumi.StringOutput)
}

type GetProxypolicyZtnaEmsTagArrayOutput struct{ *pulumi.OutputState }

func (GetProxypolicyZtnaEmsTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProxypolicyZtnaEmsTag)(nil)).Elem()
}

func (o GetProxypolicyZtnaEmsTagArrayOutput) ToGetProxypolicyZtnaEmsTagArrayOutput() GetProxypolicyZtnaEmsTagArrayOutput {
	return o
}

func (o GetProxypolicyZtnaEmsTagArrayOutput) ToGetProxypolicyZtnaEmsTagArrayOutputWithContext(ctx context.Context) GetProxypolicyZtnaEmsTagArrayOutput {
	return o
}

func (o GetProxypolicyZtnaEmsTagArrayOutput) Index(i pulumi.IntInput) GetProxypolicyZtnaEmsTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProxypolicyZtnaEmsTag {
		return vs[0].([]GetProxypolicyZtnaEmsTag)[vs[1].(int)]
	}).(GetProxypolicyZtnaEmsTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicyPoolnameInput)(nil)).Elem(), GetProxypolicyPoolnameArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicyPoolnameArrayInput)(nil)).Elem(), GetProxypolicyPoolnameArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicyServiceInput)(nil)).Elem(), GetProxypolicyServiceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicyServiceArrayInput)(nil)).Elem(), GetProxypolicyServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicySrcaddr6Input)(nil)).Elem(), GetProxypolicySrcaddr6Args{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicySrcaddr6ArrayInput)(nil)).Elem(), GetProxypolicySrcaddr6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicySrcaddrInput)(nil)).Elem(), GetProxypolicySrcaddrArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicySrcaddrArrayInput)(nil)).Elem(), GetProxypolicySrcaddrArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicySrcintfInput)(nil)).Elem(), GetProxypolicySrcintfArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicySrcintfArrayInput)(nil)).Elem(), GetProxypolicySrcintfArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicyUserInput)(nil)).Elem(), GetProxypolicyUserArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicyUserArrayInput)(nil)).Elem(), GetProxypolicyUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicyZtnaEmsTagInput)(nil)).Elem(), GetProxypolicyZtnaEmsTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProxypolicyZtnaEmsTagArrayInput)(nil)).Elem(), GetProxypolicyZtnaEmsTagArray{})
	pulumi.RegisterOutputType(GetProxypolicyPoolnameOutput{})
	pulumi.RegisterOutputType(GetProxypolicyPoolnameArrayOutput{})
	pulumi.RegisterOutputType(GetProxypolicyServiceOutput{})
	pulumi.RegisterOutputType(GetProxypolicyServiceArrayOutput{})
	pulumi.RegisterOutputType(GetProxypolicySrcaddr6Output{})
	pulumi.RegisterOutputType(GetProxypolicySrcaddr6ArrayOutput{})
	pulumi.RegisterOutputType(GetProxypolicySrcaddrOutput{})
	pulumi.RegisterOutputType(GetProxypolicySrcaddrArrayOutput{})
	pulumi.RegisterOutputType(GetProxypolicySrcintfOutput{})
	pulumi.RegisterOutputType(GetProxypolicySrcintfArrayOutput{})
	pulumi.RegisterOutputType(GetProxypolicyUserOutput{})
	pulumi.RegisterOutputType(GetProxypolicyUserArrayOutput{})
	pulumi.RegisterOutputType(GetProxypolicyZtnaEmsTagOutput{})
	pulumi.RegisterOutputType(GetProxypolicyZtnaEmsTagArrayOutput{})
}
