// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logsyslogd3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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

type OverridesettingCustomFieldName struct {
	// Field custom name.
	Custom *string `pulumi:"custom"`
	// Entry ID.
	Id *int `pulumi:"id"`
	// Field name.
	Name *string `pulumi:"name"`
}

// OverridesettingCustomFieldNameInput is an input type that accepts OverridesettingCustomFieldNameArgs and OverridesettingCustomFieldNameOutput values.
// You can construct a concrete instance of `OverridesettingCustomFieldNameInput` via:
//
//	OverridesettingCustomFieldNameArgs{...}
type OverridesettingCustomFieldNameInput interface {
	pulumi.Input

	ToOverridesettingCustomFieldNameOutput() OverridesettingCustomFieldNameOutput
	ToOverridesettingCustomFieldNameOutputWithContext(context.Context) OverridesettingCustomFieldNameOutput
}

type OverridesettingCustomFieldNameArgs struct {
	// Field custom name.
	Custom pulumi.StringPtrInput `pulumi:"custom"`
	// Entry ID.
	Id pulumi.IntPtrInput `pulumi:"id"`
	// Field name.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (OverridesettingCustomFieldNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OverridesettingCustomFieldName)(nil)).Elem()
}

func (i OverridesettingCustomFieldNameArgs) ToOverridesettingCustomFieldNameOutput() OverridesettingCustomFieldNameOutput {
	return i.ToOverridesettingCustomFieldNameOutputWithContext(context.Background())
}

func (i OverridesettingCustomFieldNameArgs) ToOverridesettingCustomFieldNameOutputWithContext(ctx context.Context) OverridesettingCustomFieldNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverridesettingCustomFieldNameOutput)
}

// OverridesettingCustomFieldNameArrayInput is an input type that accepts OverridesettingCustomFieldNameArray and OverridesettingCustomFieldNameArrayOutput values.
// You can construct a concrete instance of `OverridesettingCustomFieldNameArrayInput` via:
//
//	OverridesettingCustomFieldNameArray{ OverridesettingCustomFieldNameArgs{...} }
type OverridesettingCustomFieldNameArrayInput interface {
	pulumi.Input

	ToOverridesettingCustomFieldNameArrayOutput() OverridesettingCustomFieldNameArrayOutput
	ToOverridesettingCustomFieldNameArrayOutputWithContext(context.Context) OverridesettingCustomFieldNameArrayOutput
}

type OverridesettingCustomFieldNameArray []OverridesettingCustomFieldNameInput

func (OverridesettingCustomFieldNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OverridesettingCustomFieldName)(nil)).Elem()
}

func (i OverridesettingCustomFieldNameArray) ToOverridesettingCustomFieldNameArrayOutput() OverridesettingCustomFieldNameArrayOutput {
	return i.ToOverridesettingCustomFieldNameArrayOutputWithContext(context.Background())
}

func (i OverridesettingCustomFieldNameArray) ToOverridesettingCustomFieldNameArrayOutputWithContext(ctx context.Context) OverridesettingCustomFieldNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverridesettingCustomFieldNameArrayOutput)
}

type OverridesettingCustomFieldNameOutput struct{ *pulumi.OutputState }

func (OverridesettingCustomFieldNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OverridesettingCustomFieldName)(nil)).Elem()
}

func (o OverridesettingCustomFieldNameOutput) ToOverridesettingCustomFieldNameOutput() OverridesettingCustomFieldNameOutput {
	return o
}

func (o OverridesettingCustomFieldNameOutput) ToOverridesettingCustomFieldNameOutputWithContext(ctx context.Context) OverridesettingCustomFieldNameOutput {
	return o
}

// Field custom name.
func (o OverridesettingCustomFieldNameOutput) Custom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverridesettingCustomFieldName) *string { return v.Custom }).(pulumi.StringPtrOutput)
}

// Entry ID.
func (o OverridesettingCustomFieldNameOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OverridesettingCustomFieldName) *int { return v.Id }).(pulumi.IntPtrOutput)
}

// Field name.
func (o OverridesettingCustomFieldNameOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverridesettingCustomFieldName) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type OverridesettingCustomFieldNameArrayOutput struct{ *pulumi.OutputState }

func (OverridesettingCustomFieldNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OverridesettingCustomFieldName)(nil)).Elem()
}

func (o OverridesettingCustomFieldNameArrayOutput) ToOverridesettingCustomFieldNameArrayOutput() OverridesettingCustomFieldNameArrayOutput {
	return o
}

func (o OverridesettingCustomFieldNameArrayOutput) ToOverridesettingCustomFieldNameArrayOutputWithContext(ctx context.Context) OverridesettingCustomFieldNameArrayOutput {
	return o
}

func (o OverridesettingCustomFieldNameArrayOutput) Index(i pulumi.IntInput) OverridesettingCustomFieldNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OverridesettingCustomFieldName {
		return vs[0].([]OverridesettingCustomFieldName)[vs[1].(int)]
	}).(OverridesettingCustomFieldNameOutput)
}

type SettingCustomFieldName struct {
	// Field custom name.
	Custom *string `pulumi:"custom"`
	// Entry ID.
	Id *int `pulumi:"id"`
	// Field name.
	Name *string `pulumi:"name"`
}

// SettingCustomFieldNameInput is an input type that accepts SettingCustomFieldNameArgs and SettingCustomFieldNameOutput values.
// You can construct a concrete instance of `SettingCustomFieldNameInput` via:
//
//	SettingCustomFieldNameArgs{...}
type SettingCustomFieldNameInput interface {
	pulumi.Input

	ToSettingCustomFieldNameOutput() SettingCustomFieldNameOutput
	ToSettingCustomFieldNameOutputWithContext(context.Context) SettingCustomFieldNameOutput
}

type SettingCustomFieldNameArgs struct {
	// Field custom name.
	Custom pulumi.StringPtrInput `pulumi:"custom"`
	// Entry ID.
	Id pulumi.IntPtrInput `pulumi:"id"`
	// Field name.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SettingCustomFieldNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingCustomFieldName)(nil)).Elem()
}

func (i SettingCustomFieldNameArgs) ToSettingCustomFieldNameOutput() SettingCustomFieldNameOutput {
	return i.ToSettingCustomFieldNameOutputWithContext(context.Background())
}

func (i SettingCustomFieldNameArgs) ToSettingCustomFieldNameOutputWithContext(ctx context.Context) SettingCustomFieldNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingCustomFieldNameOutput)
}

// SettingCustomFieldNameArrayInput is an input type that accepts SettingCustomFieldNameArray and SettingCustomFieldNameArrayOutput values.
// You can construct a concrete instance of `SettingCustomFieldNameArrayInput` via:
//
//	SettingCustomFieldNameArray{ SettingCustomFieldNameArgs{...} }
type SettingCustomFieldNameArrayInput interface {
	pulumi.Input

	ToSettingCustomFieldNameArrayOutput() SettingCustomFieldNameArrayOutput
	ToSettingCustomFieldNameArrayOutputWithContext(context.Context) SettingCustomFieldNameArrayOutput
}

type SettingCustomFieldNameArray []SettingCustomFieldNameInput

func (SettingCustomFieldNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingCustomFieldName)(nil)).Elem()
}

func (i SettingCustomFieldNameArray) ToSettingCustomFieldNameArrayOutput() SettingCustomFieldNameArrayOutput {
	return i.ToSettingCustomFieldNameArrayOutputWithContext(context.Background())
}

func (i SettingCustomFieldNameArray) ToSettingCustomFieldNameArrayOutputWithContext(ctx context.Context) SettingCustomFieldNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingCustomFieldNameArrayOutput)
}

type SettingCustomFieldNameOutput struct{ *pulumi.OutputState }

func (SettingCustomFieldNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingCustomFieldName)(nil)).Elem()
}

func (o SettingCustomFieldNameOutput) ToSettingCustomFieldNameOutput() SettingCustomFieldNameOutput {
	return o
}

func (o SettingCustomFieldNameOutput) ToSettingCustomFieldNameOutputWithContext(ctx context.Context) SettingCustomFieldNameOutput {
	return o
}

// Field custom name.
func (o SettingCustomFieldNameOutput) Custom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SettingCustomFieldName) *string { return v.Custom }).(pulumi.StringPtrOutput)
}

// Entry ID.
func (o SettingCustomFieldNameOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SettingCustomFieldName) *int { return v.Id }).(pulumi.IntPtrOutput)
}

// Field name.
func (o SettingCustomFieldNameOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SettingCustomFieldName) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SettingCustomFieldNameArrayOutput struct{ *pulumi.OutputState }

func (SettingCustomFieldNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingCustomFieldName)(nil)).Elem()
}

func (o SettingCustomFieldNameArrayOutput) ToSettingCustomFieldNameArrayOutput() SettingCustomFieldNameArrayOutput {
	return o
}

func (o SettingCustomFieldNameArrayOutput) ToSettingCustomFieldNameArrayOutputWithContext(ctx context.Context) SettingCustomFieldNameArrayOutput {
	return o
}

func (o SettingCustomFieldNameArrayOutput) Index(i pulumi.IntInput) SettingCustomFieldNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingCustomFieldName {
		return vs[0].([]SettingCustomFieldName)[vs[1].(int)]
	}).(SettingCustomFieldNameOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FilterFreeStyleInput)(nil)).Elem(), FilterFreeStyleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterFreeStyleArrayInput)(nil)).Elem(), FilterFreeStyleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OverridefilterFreeStyleInput)(nil)).Elem(), OverridefilterFreeStyleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OverridefilterFreeStyleArrayInput)(nil)).Elem(), OverridefilterFreeStyleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OverridesettingCustomFieldNameInput)(nil)).Elem(), OverridesettingCustomFieldNameArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OverridesettingCustomFieldNameArrayInput)(nil)).Elem(), OverridesettingCustomFieldNameArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingCustomFieldNameInput)(nil)).Elem(), SettingCustomFieldNameArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingCustomFieldNameArrayInput)(nil)).Elem(), SettingCustomFieldNameArray{})
	pulumi.RegisterOutputType(FilterFreeStyleOutput{})
	pulumi.RegisterOutputType(FilterFreeStyleArrayOutput{})
	pulumi.RegisterOutputType(OverridefilterFreeStyleOutput{})
	pulumi.RegisterOutputType(OverridefilterFreeStyleArrayOutput{})
	pulumi.RegisterOutputType(OverridesettingCustomFieldNameOutput{})
	pulumi.RegisterOutputType(OverridesettingCustomFieldNameArrayOutput{})
	pulumi.RegisterOutputType(SettingCustomFieldNameOutput{})
	pulumi.RegisterOutputType(SettingCustomFieldNameArrayOutput{})
}
