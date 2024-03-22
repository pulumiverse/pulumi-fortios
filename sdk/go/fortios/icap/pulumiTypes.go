// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package icap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

var _ = internal.GetEnvOrDefault

type ProfileIcapHeader struct {
	// Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.
	Base64Encoding *string `pulumi:"base64Encoding"`
	// HTTP header content.
	Content *string `pulumi:"content"`
	// HTTP forwarded header ID.
	Id *int `pulumi:"id"`
	// HTTP forwarded header name.
	Name *string `pulumi:"name"`
}

// ProfileIcapHeaderInput is an input type that accepts ProfileIcapHeaderArgs and ProfileIcapHeaderOutput values.
// You can construct a concrete instance of `ProfileIcapHeaderInput` via:
//
//	ProfileIcapHeaderArgs{...}
type ProfileIcapHeaderInput interface {
	pulumi.Input

	ToProfileIcapHeaderOutput() ProfileIcapHeaderOutput
	ToProfileIcapHeaderOutputWithContext(context.Context) ProfileIcapHeaderOutput
}

type ProfileIcapHeaderArgs struct {
	// Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.
	Base64Encoding pulumi.StringPtrInput `pulumi:"base64Encoding"`
	// HTTP header content.
	Content pulumi.StringPtrInput `pulumi:"content"`
	// HTTP forwarded header ID.
	Id pulumi.IntPtrInput `pulumi:"id"`
	// HTTP forwarded header name.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ProfileIcapHeaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileIcapHeader)(nil)).Elem()
}

func (i ProfileIcapHeaderArgs) ToProfileIcapHeaderOutput() ProfileIcapHeaderOutput {
	return i.ToProfileIcapHeaderOutputWithContext(context.Background())
}

func (i ProfileIcapHeaderArgs) ToProfileIcapHeaderOutputWithContext(ctx context.Context) ProfileIcapHeaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileIcapHeaderOutput)
}

// ProfileIcapHeaderArrayInput is an input type that accepts ProfileIcapHeaderArray and ProfileIcapHeaderArrayOutput values.
// You can construct a concrete instance of `ProfileIcapHeaderArrayInput` via:
//
//	ProfileIcapHeaderArray{ ProfileIcapHeaderArgs{...} }
type ProfileIcapHeaderArrayInput interface {
	pulumi.Input

	ToProfileIcapHeaderArrayOutput() ProfileIcapHeaderArrayOutput
	ToProfileIcapHeaderArrayOutputWithContext(context.Context) ProfileIcapHeaderArrayOutput
}

type ProfileIcapHeaderArray []ProfileIcapHeaderInput

func (ProfileIcapHeaderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileIcapHeader)(nil)).Elem()
}

func (i ProfileIcapHeaderArray) ToProfileIcapHeaderArrayOutput() ProfileIcapHeaderArrayOutput {
	return i.ToProfileIcapHeaderArrayOutputWithContext(context.Background())
}

func (i ProfileIcapHeaderArray) ToProfileIcapHeaderArrayOutputWithContext(ctx context.Context) ProfileIcapHeaderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileIcapHeaderArrayOutput)
}

type ProfileIcapHeaderOutput struct{ *pulumi.OutputState }

func (ProfileIcapHeaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileIcapHeader)(nil)).Elem()
}

func (o ProfileIcapHeaderOutput) ToProfileIcapHeaderOutput() ProfileIcapHeaderOutput {
	return o
}

func (o ProfileIcapHeaderOutput) ToProfileIcapHeaderOutputWithContext(ctx context.Context) ProfileIcapHeaderOutput {
	return o
}

// Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.
func (o ProfileIcapHeaderOutput) Base64Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileIcapHeader) *string { return v.Base64Encoding }).(pulumi.StringPtrOutput)
}

// HTTP header content.
func (o ProfileIcapHeaderOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileIcapHeader) *string { return v.Content }).(pulumi.StringPtrOutput)
}

// HTTP forwarded header ID.
func (o ProfileIcapHeaderOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProfileIcapHeader) *int { return v.Id }).(pulumi.IntPtrOutput)
}

// HTTP forwarded header name.
func (o ProfileIcapHeaderOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileIcapHeader) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ProfileIcapHeaderArrayOutput struct{ *pulumi.OutputState }

func (ProfileIcapHeaderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileIcapHeader)(nil)).Elem()
}

func (o ProfileIcapHeaderArrayOutput) ToProfileIcapHeaderArrayOutput() ProfileIcapHeaderArrayOutput {
	return o
}

func (o ProfileIcapHeaderArrayOutput) ToProfileIcapHeaderArrayOutputWithContext(ctx context.Context) ProfileIcapHeaderArrayOutput {
	return o
}

func (o ProfileIcapHeaderArrayOutput) Index(i pulumi.IntInput) ProfileIcapHeaderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileIcapHeader {
		return vs[0].([]ProfileIcapHeader)[vs[1].(int)]
	}).(ProfileIcapHeaderOutput)
}

type ProfileRespmodForwardRule struct {
	// Action to be taken for ICAP server. Valid values: `forward`, `bypass`.
	Action *string `pulumi:"action"`
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups []ProfileRespmodForwardRuleHeaderGroup `pulumi:"headerGroups"`
	// Address object for the host.
	Host *string `pulumi:"host"`
	// HTTP response status code. The structure of `httpRespStatusCode` block is documented below.
	HttpRespStatusCodes []ProfileRespmodForwardRuleHttpRespStatusCode `pulumi:"httpRespStatusCodes"`
	// Address name.
	Name *string `pulumi:"name"`
}

// ProfileRespmodForwardRuleInput is an input type that accepts ProfileRespmodForwardRuleArgs and ProfileRespmodForwardRuleOutput values.
// You can construct a concrete instance of `ProfileRespmodForwardRuleInput` via:
//
//	ProfileRespmodForwardRuleArgs{...}
type ProfileRespmodForwardRuleInput interface {
	pulumi.Input

	ToProfileRespmodForwardRuleOutput() ProfileRespmodForwardRuleOutput
	ToProfileRespmodForwardRuleOutputWithContext(context.Context) ProfileRespmodForwardRuleOutput
}

type ProfileRespmodForwardRuleArgs struct {
	// Action to be taken for ICAP server. Valid values: `forward`, `bypass`.
	Action pulumi.StringPtrInput `pulumi:"action"`
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups ProfileRespmodForwardRuleHeaderGroupArrayInput `pulumi:"headerGroups"`
	// Address object for the host.
	Host pulumi.StringPtrInput `pulumi:"host"`
	// HTTP response status code. The structure of `httpRespStatusCode` block is documented below.
	HttpRespStatusCodes ProfileRespmodForwardRuleHttpRespStatusCodeArrayInput `pulumi:"httpRespStatusCodes"`
	// Address name.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ProfileRespmodForwardRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileRespmodForwardRule)(nil)).Elem()
}

func (i ProfileRespmodForwardRuleArgs) ToProfileRespmodForwardRuleOutput() ProfileRespmodForwardRuleOutput {
	return i.ToProfileRespmodForwardRuleOutputWithContext(context.Background())
}

func (i ProfileRespmodForwardRuleArgs) ToProfileRespmodForwardRuleOutputWithContext(ctx context.Context) ProfileRespmodForwardRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileRespmodForwardRuleOutput)
}

// ProfileRespmodForwardRuleArrayInput is an input type that accepts ProfileRespmodForwardRuleArray and ProfileRespmodForwardRuleArrayOutput values.
// You can construct a concrete instance of `ProfileRespmodForwardRuleArrayInput` via:
//
//	ProfileRespmodForwardRuleArray{ ProfileRespmodForwardRuleArgs{...} }
type ProfileRespmodForwardRuleArrayInput interface {
	pulumi.Input

	ToProfileRespmodForwardRuleArrayOutput() ProfileRespmodForwardRuleArrayOutput
	ToProfileRespmodForwardRuleArrayOutputWithContext(context.Context) ProfileRespmodForwardRuleArrayOutput
}

type ProfileRespmodForwardRuleArray []ProfileRespmodForwardRuleInput

func (ProfileRespmodForwardRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileRespmodForwardRule)(nil)).Elem()
}

func (i ProfileRespmodForwardRuleArray) ToProfileRespmodForwardRuleArrayOutput() ProfileRespmodForwardRuleArrayOutput {
	return i.ToProfileRespmodForwardRuleArrayOutputWithContext(context.Background())
}

func (i ProfileRespmodForwardRuleArray) ToProfileRespmodForwardRuleArrayOutputWithContext(ctx context.Context) ProfileRespmodForwardRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileRespmodForwardRuleArrayOutput)
}

type ProfileRespmodForwardRuleOutput struct{ *pulumi.OutputState }

func (ProfileRespmodForwardRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileRespmodForwardRule)(nil)).Elem()
}

func (o ProfileRespmodForwardRuleOutput) ToProfileRespmodForwardRuleOutput() ProfileRespmodForwardRuleOutput {
	return o
}

func (o ProfileRespmodForwardRuleOutput) ToProfileRespmodForwardRuleOutputWithContext(ctx context.Context) ProfileRespmodForwardRuleOutput {
	return o
}

// Action to be taken for ICAP server. Valid values: `forward`, `bypass`.
func (o ProfileRespmodForwardRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileRespmodForwardRule) *string { return v.Action }).(pulumi.StringPtrOutput)
}

// HTTP header group. The structure of `headerGroup` block is documented below.
func (o ProfileRespmodForwardRuleOutput) HeaderGroups() ProfileRespmodForwardRuleHeaderGroupArrayOutput {
	return o.ApplyT(func(v ProfileRespmodForwardRule) []ProfileRespmodForwardRuleHeaderGroup { return v.HeaderGroups }).(ProfileRespmodForwardRuleHeaderGroupArrayOutput)
}

// Address object for the host.
func (o ProfileRespmodForwardRuleOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileRespmodForwardRule) *string { return v.Host }).(pulumi.StringPtrOutput)
}

// HTTP response status code. The structure of `httpRespStatusCode` block is documented below.
func (o ProfileRespmodForwardRuleOutput) HttpRespStatusCodes() ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput {
	return o.ApplyT(func(v ProfileRespmodForwardRule) []ProfileRespmodForwardRuleHttpRespStatusCode {
		return v.HttpRespStatusCodes
	}).(ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput)
}

// Address name.
func (o ProfileRespmodForwardRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileRespmodForwardRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ProfileRespmodForwardRuleArrayOutput struct{ *pulumi.OutputState }

func (ProfileRespmodForwardRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileRespmodForwardRule)(nil)).Elem()
}

func (o ProfileRespmodForwardRuleArrayOutput) ToProfileRespmodForwardRuleArrayOutput() ProfileRespmodForwardRuleArrayOutput {
	return o
}

func (o ProfileRespmodForwardRuleArrayOutput) ToProfileRespmodForwardRuleArrayOutputWithContext(ctx context.Context) ProfileRespmodForwardRuleArrayOutput {
	return o
}

func (o ProfileRespmodForwardRuleArrayOutput) Index(i pulumi.IntInput) ProfileRespmodForwardRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileRespmodForwardRule {
		return vs[0].([]ProfileRespmodForwardRule)[vs[1].(int)]
	}).(ProfileRespmodForwardRuleOutput)
}

type ProfileRespmodForwardRuleHeaderGroup struct {
	// Enable/disable case sensitivity when matching header. Valid values: `disable`, `enable`.
	CaseSensitivity *string `pulumi:"caseSensitivity"`
	// HTTP header regular expression.
	Header *string `pulumi:"header"`
	// HTTP header.
	HeaderName *string `pulumi:"headerName"`
	// ID.
	Id *int `pulumi:"id"`
}

// ProfileRespmodForwardRuleHeaderGroupInput is an input type that accepts ProfileRespmodForwardRuleHeaderGroupArgs and ProfileRespmodForwardRuleHeaderGroupOutput values.
// You can construct a concrete instance of `ProfileRespmodForwardRuleHeaderGroupInput` via:
//
//	ProfileRespmodForwardRuleHeaderGroupArgs{...}
type ProfileRespmodForwardRuleHeaderGroupInput interface {
	pulumi.Input

	ToProfileRespmodForwardRuleHeaderGroupOutput() ProfileRespmodForwardRuleHeaderGroupOutput
	ToProfileRespmodForwardRuleHeaderGroupOutputWithContext(context.Context) ProfileRespmodForwardRuleHeaderGroupOutput
}

type ProfileRespmodForwardRuleHeaderGroupArgs struct {
	// Enable/disable case sensitivity when matching header. Valid values: `disable`, `enable`.
	CaseSensitivity pulumi.StringPtrInput `pulumi:"caseSensitivity"`
	// HTTP header regular expression.
	Header pulumi.StringPtrInput `pulumi:"header"`
	// HTTP header.
	HeaderName pulumi.StringPtrInput `pulumi:"headerName"`
	// ID.
	Id pulumi.IntPtrInput `pulumi:"id"`
}

func (ProfileRespmodForwardRuleHeaderGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileRespmodForwardRuleHeaderGroup)(nil)).Elem()
}

func (i ProfileRespmodForwardRuleHeaderGroupArgs) ToProfileRespmodForwardRuleHeaderGroupOutput() ProfileRespmodForwardRuleHeaderGroupOutput {
	return i.ToProfileRespmodForwardRuleHeaderGroupOutputWithContext(context.Background())
}

func (i ProfileRespmodForwardRuleHeaderGroupArgs) ToProfileRespmodForwardRuleHeaderGroupOutputWithContext(ctx context.Context) ProfileRespmodForwardRuleHeaderGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileRespmodForwardRuleHeaderGroupOutput)
}

// ProfileRespmodForwardRuleHeaderGroupArrayInput is an input type that accepts ProfileRespmodForwardRuleHeaderGroupArray and ProfileRespmodForwardRuleHeaderGroupArrayOutput values.
// You can construct a concrete instance of `ProfileRespmodForwardRuleHeaderGroupArrayInput` via:
//
//	ProfileRespmodForwardRuleHeaderGroupArray{ ProfileRespmodForwardRuleHeaderGroupArgs{...} }
type ProfileRespmodForwardRuleHeaderGroupArrayInput interface {
	pulumi.Input

	ToProfileRespmodForwardRuleHeaderGroupArrayOutput() ProfileRespmodForwardRuleHeaderGroupArrayOutput
	ToProfileRespmodForwardRuleHeaderGroupArrayOutputWithContext(context.Context) ProfileRespmodForwardRuleHeaderGroupArrayOutput
}

type ProfileRespmodForwardRuleHeaderGroupArray []ProfileRespmodForwardRuleHeaderGroupInput

func (ProfileRespmodForwardRuleHeaderGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileRespmodForwardRuleHeaderGroup)(nil)).Elem()
}

func (i ProfileRespmodForwardRuleHeaderGroupArray) ToProfileRespmodForwardRuleHeaderGroupArrayOutput() ProfileRespmodForwardRuleHeaderGroupArrayOutput {
	return i.ToProfileRespmodForwardRuleHeaderGroupArrayOutputWithContext(context.Background())
}

func (i ProfileRespmodForwardRuleHeaderGroupArray) ToProfileRespmodForwardRuleHeaderGroupArrayOutputWithContext(ctx context.Context) ProfileRespmodForwardRuleHeaderGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileRespmodForwardRuleHeaderGroupArrayOutput)
}

type ProfileRespmodForwardRuleHeaderGroupOutput struct{ *pulumi.OutputState }

func (ProfileRespmodForwardRuleHeaderGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileRespmodForwardRuleHeaderGroup)(nil)).Elem()
}

func (o ProfileRespmodForwardRuleHeaderGroupOutput) ToProfileRespmodForwardRuleHeaderGroupOutput() ProfileRespmodForwardRuleHeaderGroupOutput {
	return o
}

func (o ProfileRespmodForwardRuleHeaderGroupOutput) ToProfileRespmodForwardRuleHeaderGroupOutputWithContext(ctx context.Context) ProfileRespmodForwardRuleHeaderGroupOutput {
	return o
}

// Enable/disable case sensitivity when matching header. Valid values: `disable`, `enable`.
func (o ProfileRespmodForwardRuleHeaderGroupOutput) CaseSensitivity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileRespmodForwardRuleHeaderGroup) *string { return v.CaseSensitivity }).(pulumi.StringPtrOutput)
}

// HTTP header regular expression.
func (o ProfileRespmodForwardRuleHeaderGroupOutput) Header() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileRespmodForwardRuleHeaderGroup) *string { return v.Header }).(pulumi.StringPtrOutput)
}

// HTTP header.
func (o ProfileRespmodForwardRuleHeaderGroupOutput) HeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileRespmodForwardRuleHeaderGroup) *string { return v.HeaderName }).(pulumi.StringPtrOutput)
}

// ID.
func (o ProfileRespmodForwardRuleHeaderGroupOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProfileRespmodForwardRuleHeaderGroup) *int { return v.Id }).(pulumi.IntPtrOutput)
}

type ProfileRespmodForwardRuleHeaderGroupArrayOutput struct{ *pulumi.OutputState }

func (ProfileRespmodForwardRuleHeaderGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileRespmodForwardRuleHeaderGroup)(nil)).Elem()
}

func (o ProfileRespmodForwardRuleHeaderGroupArrayOutput) ToProfileRespmodForwardRuleHeaderGroupArrayOutput() ProfileRespmodForwardRuleHeaderGroupArrayOutput {
	return o
}

func (o ProfileRespmodForwardRuleHeaderGroupArrayOutput) ToProfileRespmodForwardRuleHeaderGroupArrayOutputWithContext(ctx context.Context) ProfileRespmodForwardRuleHeaderGroupArrayOutput {
	return o
}

func (o ProfileRespmodForwardRuleHeaderGroupArrayOutput) Index(i pulumi.IntInput) ProfileRespmodForwardRuleHeaderGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileRespmodForwardRuleHeaderGroup {
		return vs[0].([]ProfileRespmodForwardRuleHeaderGroup)[vs[1].(int)]
	}).(ProfileRespmodForwardRuleHeaderGroupOutput)
}

type ProfileRespmodForwardRuleHttpRespStatusCode struct {
	// HTTP response status code.
	Code *int `pulumi:"code"`
}

// ProfileRespmodForwardRuleHttpRespStatusCodeInput is an input type that accepts ProfileRespmodForwardRuleHttpRespStatusCodeArgs and ProfileRespmodForwardRuleHttpRespStatusCodeOutput values.
// You can construct a concrete instance of `ProfileRespmodForwardRuleHttpRespStatusCodeInput` via:
//
//	ProfileRespmodForwardRuleHttpRespStatusCodeArgs{...}
type ProfileRespmodForwardRuleHttpRespStatusCodeInput interface {
	pulumi.Input

	ToProfileRespmodForwardRuleHttpRespStatusCodeOutput() ProfileRespmodForwardRuleHttpRespStatusCodeOutput
	ToProfileRespmodForwardRuleHttpRespStatusCodeOutputWithContext(context.Context) ProfileRespmodForwardRuleHttpRespStatusCodeOutput
}

type ProfileRespmodForwardRuleHttpRespStatusCodeArgs struct {
	// HTTP response status code.
	Code pulumi.IntPtrInput `pulumi:"code"`
}

func (ProfileRespmodForwardRuleHttpRespStatusCodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileRespmodForwardRuleHttpRespStatusCode)(nil)).Elem()
}

func (i ProfileRespmodForwardRuleHttpRespStatusCodeArgs) ToProfileRespmodForwardRuleHttpRespStatusCodeOutput() ProfileRespmodForwardRuleHttpRespStatusCodeOutput {
	return i.ToProfileRespmodForwardRuleHttpRespStatusCodeOutputWithContext(context.Background())
}

func (i ProfileRespmodForwardRuleHttpRespStatusCodeArgs) ToProfileRespmodForwardRuleHttpRespStatusCodeOutputWithContext(ctx context.Context) ProfileRespmodForwardRuleHttpRespStatusCodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileRespmodForwardRuleHttpRespStatusCodeOutput)
}

// ProfileRespmodForwardRuleHttpRespStatusCodeArrayInput is an input type that accepts ProfileRespmodForwardRuleHttpRespStatusCodeArray and ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput values.
// You can construct a concrete instance of `ProfileRespmodForwardRuleHttpRespStatusCodeArrayInput` via:
//
//	ProfileRespmodForwardRuleHttpRespStatusCodeArray{ ProfileRespmodForwardRuleHttpRespStatusCodeArgs{...} }
type ProfileRespmodForwardRuleHttpRespStatusCodeArrayInput interface {
	pulumi.Input

	ToProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput() ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput
	ToProfileRespmodForwardRuleHttpRespStatusCodeArrayOutputWithContext(context.Context) ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput
}

type ProfileRespmodForwardRuleHttpRespStatusCodeArray []ProfileRespmodForwardRuleHttpRespStatusCodeInput

func (ProfileRespmodForwardRuleHttpRespStatusCodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileRespmodForwardRuleHttpRespStatusCode)(nil)).Elem()
}

func (i ProfileRespmodForwardRuleHttpRespStatusCodeArray) ToProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput() ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput {
	return i.ToProfileRespmodForwardRuleHttpRespStatusCodeArrayOutputWithContext(context.Background())
}

func (i ProfileRespmodForwardRuleHttpRespStatusCodeArray) ToProfileRespmodForwardRuleHttpRespStatusCodeArrayOutputWithContext(ctx context.Context) ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput)
}

type ProfileRespmodForwardRuleHttpRespStatusCodeOutput struct{ *pulumi.OutputState }

func (ProfileRespmodForwardRuleHttpRespStatusCodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileRespmodForwardRuleHttpRespStatusCode)(nil)).Elem()
}

func (o ProfileRespmodForwardRuleHttpRespStatusCodeOutput) ToProfileRespmodForwardRuleHttpRespStatusCodeOutput() ProfileRespmodForwardRuleHttpRespStatusCodeOutput {
	return o
}

func (o ProfileRespmodForwardRuleHttpRespStatusCodeOutput) ToProfileRespmodForwardRuleHttpRespStatusCodeOutputWithContext(ctx context.Context) ProfileRespmodForwardRuleHttpRespStatusCodeOutput {
	return o
}

// HTTP response status code.
func (o ProfileRespmodForwardRuleHttpRespStatusCodeOutput) Code() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProfileRespmodForwardRuleHttpRespStatusCode) *int { return v.Code }).(pulumi.IntPtrOutput)
}

type ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput struct{ *pulumi.OutputState }

func (ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileRespmodForwardRuleHttpRespStatusCode)(nil)).Elem()
}

func (o ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput) ToProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput() ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput {
	return o
}

func (o ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput) ToProfileRespmodForwardRuleHttpRespStatusCodeArrayOutputWithContext(ctx context.Context) ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput {
	return o
}

func (o ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput) Index(i pulumi.IntInput) ProfileRespmodForwardRuleHttpRespStatusCodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileRespmodForwardRuleHttpRespStatusCode {
		return vs[0].([]ProfileRespmodForwardRuleHttpRespStatusCode)[vs[1].(int)]
	}).(ProfileRespmodForwardRuleHttpRespStatusCodeOutput)
}

type ServergroupServerList struct {
	// ICAP server name.
	Name *string `pulumi:"name"`
	// Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10).
	Weight *int `pulumi:"weight"`
}

// ServergroupServerListInput is an input type that accepts ServergroupServerListArgs and ServergroupServerListOutput values.
// You can construct a concrete instance of `ServergroupServerListInput` via:
//
//	ServergroupServerListArgs{...}
type ServergroupServerListInput interface {
	pulumi.Input

	ToServergroupServerListOutput() ServergroupServerListOutput
	ToServergroupServerListOutputWithContext(context.Context) ServergroupServerListOutput
}

type ServergroupServerListArgs struct {
	// ICAP server name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10).
	Weight pulumi.IntPtrInput `pulumi:"weight"`
}

func (ServergroupServerListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServergroupServerList)(nil)).Elem()
}

func (i ServergroupServerListArgs) ToServergroupServerListOutput() ServergroupServerListOutput {
	return i.ToServergroupServerListOutputWithContext(context.Background())
}

func (i ServergroupServerListArgs) ToServergroupServerListOutputWithContext(ctx context.Context) ServergroupServerListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServergroupServerListOutput)
}

// ServergroupServerListArrayInput is an input type that accepts ServergroupServerListArray and ServergroupServerListArrayOutput values.
// You can construct a concrete instance of `ServergroupServerListArrayInput` via:
//
//	ServergroupServerListArray{ ServergroupServerListArgs{...} }
type ServergroupServerListArrayInput interface {
	pulumi.Input

	ToServergroupServerListArrayOutput() ServergroupServerListArrayOutput
	ToServergroupServerListArrayOutputWithContext(context.Context) ServergroupServerListArrayOutput
}

type ServergroupServerListArray []ServergroupServerListInput

func (ServergroupServerListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServergroupServerList)(nil)).Elem()
}

func (i ServergroupServerListArray) ToServergroupServerListArrayOutput() ServergroupServerListArrayOutput {
	return i.ToServergroupServerListArrayOutputWithContext(context.Background())
}

func (i ServergroupServerListArray) ToServergroupServerListArrayOutputWithContext(ctx context.Context) ServergroupServerListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServergroupServerListArrayOutput)
}

type ServergroupServerListOutput struct{ *pulumi.OutputState }

func (ServergroupServerListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServergroupServerList)(nil)).Elem()
}

func (o ServergroupServerListOutput) ToServergroupServerListOutput() ServergroupServerListOutput {
	return o
}

func (o ServergroupServerListOutput) ToServergroupServerListOutputWithContext(ctx context.Context) ServergroupServerListOutput {
	return o
}

// ICAP server name.
func (o ServergroupServerListOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServergroupServerList) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10).
func (o ServergroupServerListOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServergroupServerList) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type ServergroupServerListArrayOutput struct{ *pulumi.OutputState }

func (ServergroupServerListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServergroupServerList)(nil)).Elem()
}

func (o ServergroupServerListArrayOutput) ToServergroupServerListArrayOutput() ServergroupServerListArrayOutput {
	return o
}

func (o ServergroupServerListArrayOutput) ToServergroupServerListArrayOutputWithContext(ctx context.Context) ServergroupServerListArrayOutput {
	return o
}

func (o ServergroupServerListArrayOutput) Index(i pulumi.IntInput) ServergroupServerListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServergroupServerList {
		return vs[0].([]ServergroupServerList)[vs[1].(int)]
	}).(ServergroupServerListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileIcapHeaderInput)(nil)).Elem(), ProfileIcapHeaderArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileIcapHeaderArrayInput)(nil)).Elem(), ProfileIcapHeaderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileRespmodForwardRuleInput)(nil)).Elem(), ProfileRespmodForwardRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileRespmodForwardRuleArrayInput)(nil)).Elem(), ProfileRespmodForwardRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileRespmodForwardRuleHeaderGroupInput)(nil)).Elem(), ProfileRespmodForwardRuleHeaderGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileRespmodForwardRuleHeaderGroupArrayInput)(nil)).Elem(), ProfileRespmodForwardRuleHeaderGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileRespmodForwardRuleHttpRespStatusCodeInput)(nil)).Elem(), ProfileRespmodForwardRuleHttpRespStatusCodeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileRespmodForwardRuleHttpRespStatusCodeArrayInput)(nil)).Elem(), ProfileRespmodForwardRuleHttpRespStatusCodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServergroupServerListInput)(nil)).Elem(), ServergroupServerListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServergroupServerListArrayInput)(nil)).Elem(), ServergroupServerListArray{})
	pulumi.RegisterOutputType(ProfileIcapHeaderOutput{})
	pulumi.RegisterOutputType(ProfileIcapHeaderArrayOutput{})
	pulumi.RegisterOutputType(ProfileRespmodForwardRuleOutput{})
	pulumi.RegisterOutputType(ProfileRespmodForwardRuleArrayOutput{})
	pulumi.RegisterOutputType(ProfileRespmodForwardRuleHeaderGroupOutput{})
	pulumi.RegisterOutputType(ProfileRespmodForwardRuleHeaderGroupArrayOutput{})
	pulumi.RegisterOutputType(ProfileRespmodForwardRuleHttpRespStatusCodeOutput{})
	pulumi.RegisterOutputType(ProfileRespmodForwardRuleHttpRespStatusCodeArrayOutput{})
	pulumi.RegisterOutputType(ServergroupServerListOutput{})
	pulumi.RegisterOutputType(ServergroupServerListArrayOutput{})
}
