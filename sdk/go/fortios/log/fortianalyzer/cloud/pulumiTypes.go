// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

var _ = internal.GetEnvOrDefault

type FilterFreeStyle struct {
	// Log category.
	Category *string `pulumi:"category"`
	// Free style filter string.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Entry ID.
	Id *int `pulumi:"id"`
}

// FilterFreeStyleInput is an input type that accepts FilterFreeStyleArgs and FilterFreeStyleOutput values.
// You can construct a concrete instance of `FilterFreeStyleInput` via:
//
//	FilterFreeStyleArgs{...}
type FilterFreeStyleInput interface {
	pulumi.Input

	ToFilterFreeStyleOutput() FilterFreeStyleOutput
	ToFilterFreeStyleOutputWithContext(context.Context) FilterFreeStyleOutput
}

type FilterFreeStyleArgs struct {
	// Log category.
	Category pulumi.StringPtrInput `pulumi:"category"`
	// Free style filter string.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput `pulumi:"filterType"`
	// Entry ID.
	Id pulumi.IntPtrInput `pulumi:"id"`
}

func (FilterFreeStyleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFreeStyle)(nil)).Elem()
}

func (i FilterFreeStyleArgs) ToFilterFreeStyleOutput() FilterFreeStyleOutput {
	return i.ToFilterFreeStyleOutputWithContext(context.Background())
}

func (i FilterFreeStyleArgs) ToFilterFreeStyleOutputWithContext(ctx context.Context) FilterFreeStyleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFreeStyleOutput)
}

// FilterFreeStyleArrayInput is an input type that accepts FilterFreeStyleArray and FilterFreeStyleArrayOutput values.
// You can construct a concrete instance of `FilterFreeStyleArrayInput` via:
//
//	FilterFreeStyleArray{ FilterFreeStyleArgs{...} }
type FilterFreeStyleArrayInput interface {
	pulumi.Input

	ToFilterFreeStyleArrayOutput() FilterFreeStyleArrayOutput
	ToFilterFreeStyleArrayOutputWithContext(context.Context) FilterFreeStyleArrayOutput
}

type FilterFreeStyleArray []FilterFreeStyleInput

func (FilterFreeStyleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterFreeStyle)(nil)).Elem()
}

func (i FilterFreeStyleArray) ToFilterFreeStyleArrayOutput() FilterFreeStyleArrayOutput {
	return i.ToFilterFreeStyleArrayOutputWithContext(context.Background())
}

func (i FilterFreeStyleArray) ToFilterFreeStyleArrayOutputWithContext(ctx context.Context) FilterFreeStyleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFreeStyleArrayOutput)
}

type FilterFreeStyleOutput struct{ *pulumi.OutputState }

func (FilterFreeStyleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFreeStyle)(nil)).Elem()
}

func (o FilterFreeStyleOutput) ToFilterFreeStyleOutput() FilterFreeStyleOutput {
	return o
}

func (o FilterFreeStyleOutput) ToFilterFreeStyleOutputWithContext(ctx context.Context) FilterFreeStyleOutput {
	return o
}

// Log category.
func (o FilterFreeStyleOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilterFreeStyle) *string { return v.Category }).(pulumi.StringPtrOutput)
}

// Free style filter string.
func (o FilterFreeStyleOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilterFreeStyle) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
func (o FilterFreeStyleOutput) FilterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilterFreeStyle) *string { return v.FilterType }).(pulumi.StringPtrOutput)
}

// Entry ID.
func (o FilterFreeStyleOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FilterFreeStyle) *int { return v.Id }).(pulumi.IntPtrOutput)
}

type FilterFreeStyleArrayOutput struct{ *pulumi.OutputState }

func (FilterFreeStyleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterFreeStyle)(nil)).Elem()
}

func (o FilterFreeStyleArrayOutput) ToFilterFreeStyleArrayOutput() FilterFreeStyleArrayOutput {
	return o
}

func (o FilterFreeStyleArrayOutput) ToFilterFreeStyleArrayOutputWithContext(ctx context.Context) FilterFreeStyleArrayOutput {
	return o
}

func (o FilterFreeStyleArrayOutput) Index(i pulumi.IntInput) FilterFreeStyleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilterFreeStyle {
		return vs[0].([]FilterFreeStyle)[vs[1].(int)]
	}).(FilterFreeStyleOutput)
}

type OverridefilterFreeStyle struct {
	// Log category.
	Category *string `pulumi:"category"`
	// Free style filter string.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Entry ID.
	Id *int `pulumi:"id"`
}

// OverridefilterFreeStyleInput is an input type that accepts OverridefilterFreeStyleArgs and OverridefilterFreeStyleOutput values.
// You can construct a concrete instance of `OverridefilterFreeStyleInput` via:
//
//	OverridefilterFreeStyleArgs{...}
type OverridefilterFreeStyleInput interface {
	pulumi.Input

	ToOverridefilterFreeStyleOutput() OverridefilterFreeStyleOutput
	ToOverridefilterFreeStyleOutputWithContext(context.Context) OverridefilterFreeStyleOutput
}

type OverridefilterFreeStyleArgs struct {
	// Log category.
	Category pulumi.StringPtrInput `pulumi:"category"`
	// Free style filter string.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput `pulumi:"filterType"`
	// Entry ID.
	Id pulumi.IntPtrInput `pulumi:"id"`
}

func (OverridefilterFreeStyleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OverridefilterFreeStyle)(nil)).Elem()
}

func (i OverridefilterFreeStyleArgs) ToOverridefilterFreeStyleOutput() OverridefilterFreeStyleOutput {
	return i.ToOverridefilterFreeStyleOutputWithContext(context.Background())
}

func (i OverridefilterFreeStyleArgs) ToOverridefilterFreeStyleOutputWithContext(ctx context.Context) OverridefilterFreeStyleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverridefilterFreeStyleOutput)
}

// OverridefilterFreeStyleArrayInput is an input type that accepts OverridefilterFreeStyleArray and OverridefilterFreeStyleArrayOutput values.
// You can construct a concrete instance of `OverridefilterFreeStyleArrayInput` via:
//
//	OverridefilterFreeStyleArray{ OverridefilterFreeStyleArgs{...} }
type OverridefilterFreeStyleArrayInput interface {
	pulumi.Input

	ToOverridefilterFreeStyleArrayOutput() OverridefilterFreeStyleArrayOutput
	ToOverridefilterFreeStyleArrayOutputWithContext(context.Context) OverridefilterFreeStyleArrayOutput
}

type OverridefilterFreeStyleArray []OverridefilterFreeStyleInput

func (OverridefilterFreeStyleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OverridefilterFreeStyle)(nil)).Elem()
}

func (i OverridefilterFreeStyleArray) ToOverridefilterFreeStyleArrayOutput() OverridefilterFreeStyleArrayOutput {
	return i.ToOverridefilterFreeStyleArrayOutputWithContext(context.Background())
}

func (i OverridefilterFreeStyleArray) ToOverridefilterFreeStyleArrayOutputWithContext(ctx context.Context) OverridefilterFreeStyleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverridefilterFreeStyleArrayOutput)
}

type OverridefilterFreeStyleOutput struct{ *pulumi.OutputState }

func (OverridefilterFreeStyleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OverridefilterFreeStyle)(nil)).Elem()
}

func (o OverridefilterFreeStyleOutput) ToOverridefilterFreeStyleOutput() OverridefilterFreeStyleOutput {
	return o
}

func (o OverridefilterFreeStyleOutput) ToOverridefilterFreeStyleOutputWithContext(ctx context.Context) OverridefilterFreeStyleOutput {
	return o
}

// Log category.
func (o OverridefilterFreeStyleOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverridefilterFreeStyle) *string { return v.Category }).(pulumi.StringPtrOutput)
}

// Free style filter string.
func (o OverridefilterFreeStyleOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverridefilterFreeStyle) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
func (o OverridefilterFreeStyleOutput) FilterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverridefilterFreeStyle) *string { return v.FilterType }).(pulumi.StringPtrOutput)
}

// Entry ID.
func (o OverridefilterFreeStyleOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OverridefilterFreeStyle) *int { return v.Id }).(pulumi.IntPtrOutput)
}

type OverridefilterFreeStyleArrayOutput struct{ *pulumi.OutputState }

func (OverridefilterFreeStyleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OverridefilterFreeStyle)(nil)).Elem()
}

func (o OverridefilterFreeStyleArrayOutput) ToOverridefilterFreeStyleArrayOutput() OverridefilterFreeStyleArrayOutput {
	return o
}

func (o OverridefilterFreeStyleArrayOutput) ToOverridefilterFreeStyleArrayOutputWithContext(ctx context.Context) OverridefilterFreeStyleArrayOutput {
	return o
}

func (o OverridefilterFreeStyleArrayOutput) Index(i pulumi.IntInput) OverridefilterFreeStyleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OverridefilterFreeStyle {
		return vs[0].([]OverridefilterFreeStyle)[vs[1].(int)]
	}).(OverridefilterFreeStyleOutput)
}

type SettingSerial struct {
	// Serial Number.
	Name *string `pulumi:"name"`
}

// SettingSerialInput is an input type that accepts SettingSerialArgs and SettingSerialOutput values.
// You can construct a concrete instance of `SettingSerialInput` via:
//
//	SettingSerialArgs{...}
type SettingSerialInput interface {
	pulumi.Input

	ToSettingSerialOutput() SettingSerialOutput
	ToSettingSerialOutputWithContext(context.Context) SettingSerialOutput
}

type SettingSerialArgs struct {
	// Serial Number.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SettingSerialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingSerial)(nil)).Elem()
}

func (i SettingSerialArgs) ToSettingSerialOutput() SettingSerialOutput {
	return i.ToSettingSerialOutputWithContext(context.Background())
}

func (i SettingSerialArgs) ToSettingSerialOutputWithContext(ctx context.Context) SettingSerialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingSerialOutput)
}

// SettingSerialArrayInput is an input type that accepts SettingSerialArray and SettingSerialArrayOutput values.
// You can construct a concrete instance of `SettingSerialArrayInput` via:
//
//	SettingSerialArray{ SettingSerialArgs{...} }
type SettingSerialArrayInput interface {
	pulumi.Input

	ToSettingSerialArrayOutput() SettingSerialArrayOutput
	ToSettingSerialArrayOutputWithContext(context.Context) SettingSerialArrayOutput
}

type SettingSerialArray []SettingSerialInput

func (SettingSerialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingSerial)(nil)).Elem()
}

func (i SettingSerialArray) ToSettingSerialArrayOutput() SettingSerialArrayOutput {
	return i.ToSettingSerialArrayOutputWithContext(context.Background())
}

func (i SettingSerialArray) ToSettingSerialArrayOutputWithContext(ctx context.Context) SettingSerialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingSerialArrayOutput)
}

type SettingSerialOutput struct{ *pulumi.OutputState }

func (SettingSerialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingSerial)(nil)).Elem()
}

func (o SettingSerialOutput) ToSettingSerialOutput() SettingSerialOutput {
	return o
}

func (o SettingSerialOutput) ToSettingSerialOutputWithContext(ctx context.Context) SettingSerialOutput {
	return o
}

// Serial Number.
func (o SettingSerialOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SettingSerial) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SettingSerialArrayOutput struct{ *pulumi.OutputState }

func (SettingSerialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingSerial)(nil)).Elem()
}

func (o SettingSerialArrayOutput) ToSettingSerialArrayOutput() SettingSerialArrayOutput {
	return o
}

func (o SettingSerialArrayOutput) ToSettingSerialArrayOutputWithContext(ctx context.Context) SettingSerialArrayOutput {
	return o
}

func (o SettingSerialArrayOutput) Index(i pulumi.IntInput) SettingSerialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingSerial {
		return vs[0].([]SettingSerial)[vs[1].(int)]
	}).(SettingSerialOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FilterFreeStyleInput)(nil)).Elem(), FilterFreeStyleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterFreeStyleArrayInput)(nil)).Elem(), FilterFreeStyleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OverridefilterFreeStyleInput)(nil)).Elem(), OverridefilterFreeStyleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OverridefilterFreeStyleArrayInput)(nil)).Elem(), OverridefilterFreeStyleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingSerialInput)(nil)).Elem(), SettingSerialArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingSerialArrayInput)(nil)).Elem(), SettingSerialArray{})
	pulumi.RegisterOutputType(FilterFreeStyleOutput{})
	pulumi.RegisterOutputType(FilterFreeStyleArrayOutput{})
	pulumi.RegisterOutputType(OverridefilterFreeStyleOutput{})
	pulumi.RegisterOutputType(OverridefilterFreeStyleArrayOutput{})
	pulumi.RegisterOutputType(SettingSerialOutput{})
	pulumi.RegisterOutputType(SettingSerialArrayOutput{})
}