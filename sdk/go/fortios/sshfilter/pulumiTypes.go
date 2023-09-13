// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sshfilter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProfileFileFilter struct {
	// File filter entries. The structure of `entries` block is documented below.
	Entries []ProfileFileFilterEntry `pulumi:"entries"`
	// Enable/disable file filter logging. Valid values: `enable`, `disable`.
	Log *string `pulumi:"log"`
	// Enable/disable file filter archive contents scan. Valid values: `enable`, `disable`.
	ScanArchiveContents *string `pulumi:"scanArchiveContents"`
	// Enable/disable file filter. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
}

// ProfileFileFilterInput is an input type that accepts ProfileFileFilterArgs and ProfileFileFilterOutput values.
// You can construct a concrete instance of `ProfileFileFilterInput` via:
//
//	ProfileFileFilterArgs{...}
type ProfileFileFilterInput interface {
	pulumi.Input

	ToProfileFileFilterOutput() ProfileFileFilterOutput
	ToProfileFileFilterOutputWithContext(context.Context) ProfileFileFilterOutput
}

type ProfileFileFilterArgs struct {
	// File filter entries. The structure of `entries` block is documented below.
	Entries ProfileFileFilterEntryArrayInput `pulumi:"entries"`
	// Enable/disable file filter logging. Valid values: `enable`, `disable`.
	Log pulumi.StringPtrInput `pulumi:"log"`
	// Enable/disable file filter archive contents scan. Valid values: `enable`, `disable`.
	ScanArchiveContents pulumi.StringPtrInput `pulumi:"scanArchiveContents"`
	// Enable/disable file filter. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (ProfileFileFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFileFilter)(nil)).Elem()
}

func (i ProfileFileFilterArgs) ToProfileFileFilterOutput() ProfileFileFilterOutput {
	return i.ToProfileFileFilterOutputWithContext(context.Background())
}

func (i ProfileFileFilterArgs) ToProfileFileFilterOutputWithContext(ctx context.Context) ProfileFileFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFileFilterOutput)
}

func (i ProfileFileFilterArgs) ToProfileFileFilterPtrOutput() ProfileFileFilterPtrOutput {
	return i.ToProfileFileFilterPtrOutputWithContext(context.Background())
}

func (i ProfileFileFilterArgs) ToProfileFileFilterPtrOutputWithContext(ctx context.Context) ProfileFileFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFileFilterOutput).ToProfileFileFilterPtrOutputWithContext(ctx)
}

// ProfileFileFilterPtrInput is an input type that accepts ProfileFileFilterArgs, ProfileFileFilterPtr and ProfileFileFilterPtrOutput values.
// You can construct a concrete instance of `ProfileFileFilterPtrInput` via:
//
//	        ProfileFileFilterArgs{...}
//
//	or:
//
//	        nil
type ProfileFileFilterPtrInput interface {
	pulumi.Input

	ToProfileFileFilterPtrOutput() ProfileFileFilterPtrOutput
	ToProfileFileFilterPtrOutputWithContext(context.Context) ProfileFileFilterPtrOutput
}

type profileFileFilterPtrType ProfileFileFilterArgs

func ProfileFileFilterPtr(v *ProfileFileFilterArgs) ProfileFileFilterPtrInput {
	return (*profileFileFilterPtrType)(v)
}

func (*profileFileFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileFileFilter)(nil)).Elem()
}

func (i *profileFileFilterPtrType) ToProfileFileFilterPtrOutput() ProfileFileFilterPtrOutput {
	return i.ToProfileFileFilterPtrOutputWithContext(context.Background())
}

func (i *profileFileFilterPtrType) ToProfileFileFilterPtrOutputWithContext(ctx context.Context) ProfileFileFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFileFilterPtrOutput)
}

type ProfileFileFilterOutput struct{ *pulumi.OutputState }

func (ProfileFileFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFileFilter)(nil)).Elem()
}

func (o ProfileFileFilterOutput) ToProfileFileFilterOutput() ProfileFileFilterOutput {
	return o
}

func (o ProfileFileFilterOutput) ToProfileFileFilterOutputWithContext(ctx context.Context) ProfileFileFilterOutput {
	return o
}

func (o ProfileFileFilterOutput) ToProfileFileFilterPtrOutput() ProfileFileFilterPtrOutput {
	return o.ToProfileFileFilterPtrOutputWithContext(context.Background())
}

func (o ProfileFileFilterOutput) ToProfileFileFilterPtrOutputWithContext(ctx context.Context) ProfileFileFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProfileFileFilter) *ProfileFileFilter {
		return &v
	}).(ProfileFileFilterPtrOutput)
}

// File filter entries. The structure of `entries` block is documented below.
func (o ProfileFileFilterOutput) Entries() ProfileFileFilterEntryArrayOutput {
	return o.ApplyT(func(v ProfileFileFilter) []ProfileFileFilterEntry { return v.Entries }).(ProfileFileFilterEntryArrayOutput)
}

// Enable/disable file filter logging. Valid values: `enable`, `disable`.
func (o ProfileFileFilterOutput) Log() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileFileFilter) *string { return v.Log }).(pulumi.StringPtrOutput)
}

// Enable/disable file filter archive contents scan. Valid values: `enable`, `disable`.
func (o ProfileFileFilterOutput) ScanArchiveContents() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileFileFilter) *string { return v.ScanArchiveContents }).(pulumi.StringPtrOutput)
}

// Enable/disable file filter. Valid values: `enable`, `disable`.
func (o ProfileFileFilterOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileFileFilter) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ProfileFileFilterPtrOutput struct{ *pulumi.OutputState }

func (ProfileFileFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileFileFilter)(nil)).Elem()
}

func (o ProfileFileFilterPtrOutput) ToProfileFileFilterPtrOutput() ProfileFileFilterPtrOutput {
	return o
}

func (o ProfileFileFilterPtrOutput) ToProfileFileFilterPtrOutputWithContext(ctx context.Context) ProfileFileFilterPtrOutput {
	return o
}

func (o ProfileFileFilterPtrOutput) Elem() ProfileFileFilterOutput {
	return o.ApplyT(func(v *ProfileFileFilter) ProfileFileFilter {
		if v != nil {
			return *v
		}
		var ret ProfileFileFilter
		return ret
	}).(ProfileFileFilterOutput)
}

// File filter entries. The structure of `entries` block is documented below.
func (o ProfileFileFilterPtrOutput) Entries() ProfileFileFilterEntryArrayOutput {
	return o.ApplyT(func(v *ProfileFileFilter) []ProfileFileFilterEntry {
		if v == nil {
			return nil
		}
		return v.Entries
	}).(ProfileFileFilterEntryArrayOutput)
}

// Enable/disable file filter logging. Valid values: `enable`, `disable`.
func (o ProfileFileFilterPtrOutput) Log() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfileFileFilter) *string {
		if v == nil {
			return nil
		}
		return v.Log
	}).(pulumi.StringPtrOutput)
}

// Enable/disable file filter archive contents scan. Valid values: `enable`, `disable`.
func (o ProfileFileFilterPtrOutput) ScanArchiveContents() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfileFileFilter) *string {
		if v == nil {
			return nil
		}
		return v.ScanArchiveContents
	}).(pulumi.StringPtrOutput)
}

// Enable/disable file filter. Valid values: `enable`, `disable`.
func (o ProfileFileFilterPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfileFileFilter) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ProfileFileFilterEntry struct {
	// Action taken for matched file. Valid values: `log`, `block`.
	Action *string `pulumi:"action"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Match files transmitted in the session's originating or reply direction. Valid values: `incoming`, `outgoing`, `any`.
	Direction *string `pulumi:"direction"`
	// Select file type. The structure of `fileType` block is documented below.
	FileTypes []ProfileFileFilterEntryFileType `pulumi:"fileTypes"`
	// Add a file filter.
	Filter *string `pulumi:"filter"`
	// Match password-protected files. Valid values: `yes`, `any`.
	PasswordProtected *string `pulumi:"passwordProtected"`
}

// ProfileFileFilterEntryInput is an input type that accepts ProfileFileFilterEntryArgs and ProfileFileFilterEntryOutput values.
// You can construct a concrete instance of `ProfileFileFilterEntryInput` via:
//
//	ProfileFileFilterEntryArgs{...}
type ProfileFileFilterEntryInput interface {
	pulumi.Input

	ToProfileFileFilterEntryOutput() ProfileFileFilterEntryOutput
	ToProfileFileFilterEntryOutputWithContext(context.Context) ProfileFileFilterEntryOutput
}

type ProfileFileFilterEntryArgs struct {
	// Action taken for matched file. Valid values: `log`, `block`.
	Action pulumi.StringPtrInput `pulumi:"action"`
	// Comment.
	Comment pulumi.StringPtrInput `pulumi:"comment"`
	// Match files transmitted in the session's originating or reply direction. Valid values: `incoming`, `outgoing`, `any`.
	Direction pulumi.StringPtrInput `pulumi:"direction"`
	// Select file type. The structure of `fileType` block is documented below.
	FileTypes ProfileFileFilterEntryFileTypeArrayInput `pulumi:"fileTypes"`
	// Add a file filter.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Match password-protected files. Valid values: `yes`, `any`.
	PasswordProtected pulumi.StringPtrInput `pulumi:"passwordProtected"`
}

func (ProfileFileFilterEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFileFilterEntry)(nil)).Elem()
}

func (i ProfileFileFilterEntryArgs) ToProfileFileFilterEntryOutput() ProfileFileFilterEntryOutput {
	return i.ToProfileFileFilterEntryOutputWithContext(context.Background())
}

func (i ProfileFileFilterEntryArgs) ToProfileFileFilterEntryOutputWithContext(ctx context.Context) ProfileFileFilterEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFileFilterEntryOutput)
}

// ProfileFileFilterEntryArrayInput is an input type that accepts ProfileFileFilterEntryArray and ProfileFileFilterEntryArrayOutput values.
// You can construct a concrete instance of `ProfileFileFilterEntryArrayInput` via:
//
//	ProfileFileFilterEntryArray{ ProfileFileFilterEntryArgs{...} }
type ProfileFileFilterEntryArrayInput interface {
	pulumi.Input

	ToProfileFileFilterEntryArrayOutput() ProfileFileFilterEntryArrayOutput
	ToProfileFileFilterEntryArrayOutputWithContext(context.Context) ProfileFileFilterEntryArrayOutput
}

type ProfileFileFilterEntryArray []ProfileFileFilterEntryInput

func (ProfileFileFilterEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileFileFilterEntry)(nil)).Elem()
}

func (i ProfileFileFilterEntryArray) ToProfileFileFilterEntryArrayOutput() ProfileFileFilterEntryArrayOutput {
	return i.ToProfileFileFilterEntryArrayOutputWithContext(context.Background())
}

func (i ProfileFileFilterEntryArray) ToProfileFileFilterEntryArrayOutputWithContext(ctx context.Context) ProfileFileFilterEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFileFilterEntryArrayOutput)
}

type ProfileFileFilterEntryOutput struct{ *pulumi.OutputState }

func (ProfileFileFilterEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFileFilterEntry)(nil)).Elem()
}

func (o ProfileFileFilterEntryOutput) ToProfileFileFilterEntryOutput() ProfileFileFilterEntryOutput {
	return o
}

func (o ProfileFileFilterEntryOutput) ToProfileFileFilterEntryOutputWithContext(ctx context.Context) ProfileFileFilterEntryOutput {
	return o
}

// Action taken for matched file. Valid values: `log`, `block`.
func (o ProfileFileFilterEntryOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileFileFilterEntry) *string { return v.Action }).(pulumi.StringPtrOutput)
}

// Comment.
func (o ProfileFileFilterEntryOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileFileFilterEntry) *string { return v.Comment }).(pulumi.StringPtrOutput)
}

// Match files transmitted in the session's originating or reply direction. Valid values: `incoming`, `outgoing`, `any`.
func (o ProfileFileFilterEntryOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileFileFilterEntry) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

// Select file type. The structure of `fileType` block is documented below.
func (o ProfileFileFilterEntryOutput) FileTypes() ProfileFileFilterEntryFileTypeArrayOutput {
	return o.ApplyT(func(v ProfileFileFilterEntry) []ProfileFileFilterEntryFileType { return v.FileTypes }).(ProfileFileFilterEntryFileTypeArrayOutput)
}

// Add a file filter.
func (o ProfileFileFilterEntryOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileFileFilterEntry) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// Match password-protected files. Valid values: `yes`, `any`.
func (o ProfileFileFilterEntryOutput) PasswordProtected() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileFileFilterEntry) *string { return v.PasswordProtected }).(pulumi.StringPtrOutput)
}

type ProfileFileFilterEntryArrayOutput struct{ *pulumi.OutputState }

func (ProfileFileFilterEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileFileFilterEntry)(nil)).Elem()
}

func (o ProfileFileFilterEntryArrayOutput) ToProfileFileFilterEntryArrayOutput() ProfileFileFilterEntryArrayOutput {
	return o
}

func (o ProfileFileFilterEntryArrayOutput) ToProfileFileFilterEntryArrayOutputWithContext(ctx context.Context) ProfileFileFilterEntryArrayOutput {
	return o
}

func (o ProfileFileFilterEntryArrayOutput) Index(i pulumi.IntInput) ProfileFileFilterEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileFileFilterEntry {
		return vs[0].([]ProfileFileFilterEntry)[vs[1].(int)]
	}).(ProfileFileFilterEntryOutput)
}

type ProfileFileFilterEntryFileType struct {
	// File type name.
	Name *string `pulumi:"name"`
}

// ProfileFileFilterEntryFileTypeInput is an input type that accepts ProfileFileFilterEntryFileTypeArgs and ProfileFileFilterEntryFileTypeOutput values.
// You can construct a concrete instance of `ProfileFileFilterEntryFileTypeInput` via:
//
//	ProfileFileFilterEntryFileTypeArgs{...}
type ProfileFileFilterEntryFileTypeInput interface {
	pulumi.Input

	ToProfileFileFilterEntryFileTypeOutput() ProfileFileFilterEntryFileTypeOutput
	ToProfileFileFilterEntryFileTypeOutputWithContext(context.Context) ProfileFileFilterEntryFileTypeOutput
}

type ProfileFileFilterEntryFileTypeArgs struct {
	// File type name.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ProfileFileFilterEntryFileTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFileFilterEntryFileType)(nil)).Elem()
}

func (i ProfileFileFilterEntryFileTypeArgs) ToProfileFileFilterEntryFileTypeOutput() ProfileFileFilterEntryFileTypeOutput {
	return i.ToProfileFileFilterEntryFileTypeOutputWithContext(context.Background())
}

func (i ProfileFileFilterEntryFileTypeArgs) ToProfileFileFilterEntryFileTypeOutputWithContext(ctx context.Context) ProfileFileFilterEntryFileTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFileFilterEntryFileTypeOutput)
}

// ProfileFileFilterEntryFileTypeArrayInput is an input type that accepts ProfileFileFilterEntryFileTypeArray and ProfileFileFilterEntryFileTypeArrayOutput values.
// You can construct a concrete instance of `ProfileFileFilterEntryFileTypeArrayInput` via:
//
//	ProfileFileFilterEntryFileTypeArray{ ProfileFileFilterEntryFileTypeArgs{...} }
type ProfileFileFilterEntryFileTypeArrayInput interface {
	pulumi.Input

	ToProfileFileFilterEntryFileTypeArrayOutput() ProfileFileFilterEntryFileTypeArrayOutput
	ToProfileFileFilterEntryFileTypeArrayOutputWithContext(context.Context) ProfileFileFilterEntryFileTypeArrayOutput
}

type ProfileFileFilterEntryFileTypeArray []ProfileFileFilterEntryFileTypeInput

func (ProfileFileFilterEntryFileTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileFileFilterEntryFileType)(nil)).Elem()
}

func (i ProfileFileFilterEntryFileTypeArray) ToProfileFileFilterEntryFileTypeArrayOutput() ProfileFileFilterEntryFileTypeArrayOutput {
	return i.ToProfileFileFilterEntryFileTypeArrayOutputWithContext(context.Background())
}

func (i ProfileFileFilterEntryFileTypeArray) ToProfileFileFilterEntryFileTypeArrayOutputWithContext(ctx context.Context) ProfileFileFilterEntryFileTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFileFilterEntryFileTypeArrayOutput)
}

type ProfileFileFilterEntryFileTypeOutput struct{ *pulumi.OutputState }

func (ProfileFileFilterEntryFileTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFileFilterEntryFileType)(nil)).Elem()
}

func (o ProfileFileFilterEntryFileTypeOutput) ToProfileFileFilterEntryFileTypeOutput() ProfileFileFilterEntryFileTypeOutput {
	return o
}

func (o ProfileFileFilterEntryFileTypeOutput) ToProfileFileFilterEntryFileTypeOutputWithContext(ctx context.Context) ProfileFileFilterEntryFileTypeOutput {
	return o
}

// File type name.
func (o ProfileFileFilterEntryFileTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileFileFilterEntryFileType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ProfileFileFilterEntryFileTypeArrayOutput struct{ *pulumi.OutputState }

func (ProfileFileFilterEntryFileTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileFileFilterEntryFileType)(nil)).Elem()
}

func (o ProfileFileFilterEntryFileTypeArrayOutput) ToProfileFileFilterEntryFileTypeArrayOutput() ProfileFileFilterEntryFileTypeArrayOutput {
	return o
}

func (o ProfileFileFilterEntryFileTypeArrayOutput) ToProfileFileFilterEntryFileTypeArrayOutputWithContext(ctx context.Context) ProfileFileFilterEntryFileTypeArrayOutput {
	return o
}

func (o ProfileFileFilterEntryFileTypeArrayOutput) Index(i pulumi.IntInput) ProfileFileFilterEntryFileTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileFileFilterEntryFileType {
		return vs[0].([]ProfileFileFilterEntryFileType)[vs[1].(int)]
	}).(ProfileFileFilterEntryFileTypeOutput)
}

type ProfileShellCommand struct {
	// Action to take for URL filter matches. Valid values: `block`, `allow`.
	Action *string `pulumi:"action"`
	// Enable/disable alert. Valid values: `enable`, `disable`.
	Alert *string `pulumi:"alert"`
	// Id.
	Id *int `pulumi:"id"`
	// Enable/disable logging. Valid values: `enable`, `disable`.
	Log *string `pulumi:"log"`
	// SSH shell command pattern.
	Pattern *string `pulumi:"pattern"`
	// Log severity. Valid values: `low`, `medium`, `high`, `critical`.
	Severity *string `pulumi:"severity"`
	// Matching type. Valid values: `simple`, `regex`.
	Type *string `pulumi:"type"`
}

// ProfileShellCommandInput is an input type that accepts ProfileShellCommandArgs and ProfileShellCommandOutput values.
// You can construct a concrete instance of `ProfileShellCommandInput` via:
//
//	ProfileShellCommandArgs{...}
type ProfileShellCommandInput interface {
	pulumi.Input

	ToProfileShellCommandOutput() ProfileShellCommandOutput
	ToProfileShellCommandOutputWithContext(context.Context) ProfileShellCommandOutput
}

type ProfileShellCommandArgs struct {
	// Action to take for URL filter matches. Valid values: `block`, `allow`.
	Action pulumi.StringPtrInput `pulumi:"action"`
	// Enable/disable alert. Valid values: `enable`, `disable`.
	Alert pulumi.StringPtrInput `pulumi:"alert"`
	// Id.
	Id pulumi.IntPtrInput `pulumi:"id"`
	// Enable/disable logging. Valid values: `enable`, `disable`.
	Log pulumi.StringPtrInput `pulumi:"log"`
	// SSH shell command pattern.
	Pattern pulumi.StringPtrInput `pulumi:"pattern"`
	// Log severity. Valid values: `low`, `medium`, `high`, `critical`.
	Severity pulumi.StringPtrInput `pulumi:"severity"`
	// Matching type. Valid values: `simple`, `regex`.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ProfileShellCommandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileShellCommand)(nil)).Elem()
}

func (i ProfileShellCommandArgs) ToProfileShellCommandOutput() ProfileShellCommandOutput {
	return i.ToProfileShellCommandOutputWithContext(context.Background())
}

func (i ProfileShellCommandArgs) ToProfileShellCommandOutputWithContext(ctx context.Context) ProfileShellCommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileShellCommandOutput)
}

// ProfileShellCommandArrayInput is an input type that accepts ProfileShellCommandArray and ProfileShellCommandArrayOutput values.
// You can construct a concrete instance of `ProfileShellCommandArrayInput` via:
//
//	ProfileShellCommandArray{ ProfileShellCommandArgs{...} }
type ProfileShellCommandArrayInput interface {
	pulumi.Input

	ToProfileShellCommandArrayOutput() ProfileShellCommandArrayOutput
	ToProfileShellCommandArrayOutputWithContext(context.Context) ProfileShellCommandArrayOutput
}

type ProfileShellCommandArray []ProfileShellCommandInput

func (ProfileShellCommandArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileShellCommand)(nil)).Elem()
}

func (i ProfileShellCommandArray) ToProfileShellCommandArrayOutput() ProfileShellCommandArrayOutput {
	return i.ToProfileShellCommandArrayOutputWithContext(context.Background())
}

func (i ProfileShellCommandArray) ToProfileShellCommandArrayOutputWithContext(ctx context.Context) ProfileShellCommandArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileShellCommandArrayOutput)
}

type ProfileShellCommandOutput struct{ *pulumi.OutputState }

func (ProfileShellCommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileShellCommand)(nil)).Elem()
}

func (o ProfileShellCommandOutput) ToProfileShellCommandOutput() ProfileShellCommandOutput {
	return o
}

func (o ProfileShellCommandOutput) ToProfileShellCommandOutputWithContext(ctx context.Context) ProfileShellCommandOutput {
	return o
}

// Action to take for URL filter matches. Valid values: `block`, `allow`.
func (o ProfileShellCommandOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileShellCommand) *string { return v.Action }).(pulumi.StringPtrOutput)
}

// Enable/disable alert. Valid values: `enable`, `disable`.
func (o ProfileShellCommandOutput) Alert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileShellCommand) *string { return v.Alert }).(pulumi.StringPtrOutput)
}

// Id.
func (o ProfileShellCommandOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProfileShellCommand) *int { return v.Id }).(pulumi.IntPtrOutput)
}

// Enable/disable logging. Valid values: `enable`, `disable`.
func (o ProfileShellCommandOutput) Log() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileShellCommand) *string { return v.Log }).(pulumi.StringPtrOutput)
}

// SSH shell command pattern.
func (o ProfileShellCommandOutput) Pattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileShellCommand) *string { return v.Pattern }).(pulumi.StringPtrOutput)
}

// Log severity. Valid values: `low`, `medium`, `high`, `critical`.
func (o ProfileShellCommandOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileShellCommand) *string { return v.Severity }).(pulumi.StringPtrOutput)
}

// Matching type. Valid values: `simple`, `regex`.
func (o ProfileShellCommandOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileShellCommand) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ProfileShellCommandArrayOutput struct{ *pulumi.OutputState }

func (ProfileShellCommandArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileShellCommand)(nil)).Elem()
}

func (o ProfileShellCommandArrayOutput) ToProfileShellCommandArrayOutput() ProfileShellCommandArrayOutput {
	return o
}

func (o ProfileShellCommandArrayOutput) ToProfileShellCommandArrayOutputWithContext(ctx context.Context) ProfileShellCommandArrayOutput {
	return o
}

func (o ProfileShellCommandArrayOutput) Index(i pulumi.IntInput) ProfileShellCommandOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileShellCommand {
		return vs[0].([]ProfileShellCommand)[vs[1].(int)]
	}).(ProfileShellCommandOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileFileFilterInput)(nil)).Elem(), ProfileFileFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileFileFilterPtrInput)(nil)).Elem(), ProfileFileFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileFileFilterEntryInput)(nil)).Elem(), ProfileFileFilterEntryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileFileFilterEntryArrayInput)(nil)).Elem(), ProfileFileFilterEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileFileFilterEntryFileTypeInput)(nil)).Elem(), ProfileFileFilterEntryFileTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileFileFilterEntryFileTypeArrayInput)(nil)).Elem(), ProfileFileFilterEntryFileTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileShellCommandInput)(nil)).Elem(), ProfileShellCommandArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileShellCommandArrayInput)(nil)).Elem(), ProfileShellCommandArray{})
	pulumi.RegisterOutputType(ProfileFileFilterOutput{})
	pulumi.RegisterOutputType(ProfileFileFilterPtrOutput{})
	pulumi.RegisterOutputType(ProfileFileFilterEntryOutput{})
	pulumi.RegisterOutputType(ProfileFileFilterEntryArrayOutput{})
	pulumi.RegisterOutputType(ProfileFileFilterEntryFileTypeOutput{})
	pulumi.RegisterOutputType(ProfileFileFilterEntryFileTypeArrayOutput{})
	pulumi.RegisterOutputType(ProfileShellCommandOutput{})
	pulumi.RegisterOutputType(ProfileShellCommandArrayOutput{})
}
