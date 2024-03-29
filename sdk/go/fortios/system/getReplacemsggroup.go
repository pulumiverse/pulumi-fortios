// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios system replacemsggroup
func LookupReplacemsggroup(ctx *pulumi.Context, args *LookupReplacemsggroupArgs, opts ...pulumi.InvokeOption) (*LookupReplacemsggroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReplacemsggroupResult
	err := ctx.Invoke("fortios:system/getReplacemsggroup:getReplacemsggroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReplacemsggroup.
type LookupReplacemsggroupArgs struct {
	// Specify the name of the desired system replacemsggroup.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getReplacemsggroup.
type LookupReplacemsggroupResult struct {
	// Replacement message table entries. The structure of `admin` block is documented below.
	Admins []GetReplacemsggroupAdmin `pulumi:"admins"`
	// Replacement message table entries. The structure of `alertmail` block is documented below.
	Alertmails []GetReplacemsggroupAlertmail `pulumi:"alertmails"`
	// Replacement message table entries. The structure of `auth` block is documented below.
	Auths []GetReplacemsggroupAuth `pulumi:"auths"`
	// Replacement message table entries. The structure of `automation` block is documented below.
	Automations []GetReplacemsggroupAutomation `pulumi:"automations"`
	// Comment.
	Comment string `pulumi:"comment"`
	// Replacement message table entries. The structure of `customMessage` block is documented below.
	CustomMessages []GetReplacemsggroupCustomMessage `pulumi:"customMessages"`
	// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
	DeviceDetectionPortals []GetReplacemsggroupDeviceDetectionPortal `pulumi:"deviceDetectionPortals"`
	// Replacement message table entries. The structure of `ec` block is documented below.
	Ecs []GetReplacemsggroupEc `pulumi:"ecs"`
	// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
	FortiguardWfs []GetReplacemsggroupFortiguardWf `pulumi:"fortiguardWfs"`
	// Replacement message table entries. The structure of `ftp` block is documented below.
	Ftps []GetReplacemsggroupFtp `pulumi:"ftps"`
	// Group type.
	GroupType string `pulumi:"groupType"`
	// Replacement message table entries. The structure of `http` block is documented below.
	Https []GetReplacemsggroupHttp `pulumi:"https"`
	// Replacement message table entries. The structure of `icap` block is documented below.
	Icaps []GetReplacemsggroupIcap `pulumi:"icaps"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Replacement message table entries. The structure of `mail` block is documented below.
	Mails []GetReplacemsggroupMail `pulumi:"mails"`
	// Replacement message table entries. The structure of `nacQuar` block is documented below.
	NacQuars []GetReplacemsggroupNacQuar `pulumi:"nacQuars"`
	// Group name.
	Name string `pulumi:"name"`
	// Replacement message table entries. The structure of `nntp` block is documented below.
	Nntps []GetReplacemsggroupNntp `pulumi:"nntps"`
	// Replacement message table entries. The structure of `spam` block is documented below.
	Spams []GetReplacemsggroupSpam `pulumi:"spams"`
	// Replacement message table entries. The structure of `sslvpn` block is documented below.
	Sslvpns []GetReplacemsggroupSslvpn `pulumi:"sslvpns"`
	// Replacement message table entries. The structure of `trafficQuota` block is documented below.
	TrafficQuotas []GetReplacemsggroupTrafficQuota `pulumi:"trafficQuotas"`
	// Replacement message table entries. The structure of `utm` block is documented below.
	Utms      []GetReplacemsggroupUtm `pulumi:"utms"`
	Vdomparam *string                 `pulumi:"vdomparam"`
	// Replacement message table entries. The structure of `webproxy` block is documented below.
	Webproxies []GetReplacemsggroupWebproxy `pulumi:"webproxies"`
}

func LookupReplacemsggroupOutput(ctx *pulumi.Context, args LookupReplacemsggroupOutputArgs, opts ...pulumi.InvokeOption) LookupReplacemsggroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplacemsggroupResult, error) {
			args := v.(LookupReplacemsggroupArgs)
			r, err := LookupReplacemsggroup(ctx, &args, opts...)
			var s LookupReplacemsggroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplacemsggroupResultOutput)
}

// A collection of arguments for invoking getReplacemsggroup.
type LookupReplacemsggroupOutputArgs struct {
	// Specify the name of the desired system replacemsggroup.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupReplacemsggroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplacemsggroupArgs)(nil)).Elem()
}

// A collection of values returned by getReplacemsggroup.
type LookupReplacemsggroupResultOutput struct{ *pulumi.OutputState }

func (LookupReplacemsggroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplacemsggroupResult)(nil)).Elem()
}

func (o LookupReplacemsggroupResultOutput) ToLookupReplacemsggroupResultOutput() LookupReplacemsggroupResultOutput {
	return o
}

func (o LookupReplacemsggroupResultOutput) ToLookupReplacemsggroupResultOutputWithContext(ctx context.Context) LookupReplacemsggroupResultOutput {
	return o
}

// Replacement message table entries. The structure of `admin` block is documented below.
func (o LookupReplacemsggroupResultOutput) Admins() GetReplacemsggroupAdminArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupAdmin { return v.Admins }).(GetReplacemsggroupAdminArrayOutput)
}

// Replacement message table entries. The structure of `alertmail` block is documented below.
func (o LookupReplacemsggroupResultOutput) Alertmails() GetReplacemsggroupAlertmailArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupAlertmail { return v.Alertmails }).(GetReplacemsggroupAlertmailArrayOutput)
}

// Replacement message table entries. The structure of `auth` block is documented below.
func (o LookupReplacemsggroupResultOutput) Auths() GetReplacemsggroupAuthArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupAuth { return v.Auths }).(GetReplacemsggroupAuthArrayOutput)
}

// Replacement message table entries. The structure of `automation` block is documented below.
func (o LookupReplacemsggroupResultOutput) Automations() GetReplacemsggroupAutomationArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupAutomation { return v.Automations }).(GetReplacemsggroupAutomationArrayOutput)
}

// Comment.
func (o LookupReplacemsggroupResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) string { return v.Comment }).(pulumi.StringOutput)
}

// Replacement message table entries. The structure of `customMessage` block is documented below.
func (o LookupReplacemsggroupResultOutput) CustomMessages() GetReplacemsggroupCustomMessageArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupCustomMessage { return v.CustomMessages }).(GetReplacemsggroupCustomMessageArrayOutput)
}

// Replacement message table entries. The structure of `deviceDetectionPortal` block is documented below.
func (o LookupReplacemsggroupResultOutput) DeviceDetectionPortals() GetReplacemsggroupDeviceDetectionPortalArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupDeviceDetectionPortal {
		return v.DeviceDetectionPortals
	}).(GetReplacemsggroupDeviceDetectionPortalArrayOutput)
}

// Replacement message table entries. The structure of `ec` block is documented below.
func (o LookupReplacemsggroupResultOutput) Ecs() GetReplacemsggroupEcArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupEc { return v.Ecs }).(GetReplacemsggroupEcArrayOutput)
}

// Replacement message table entries. The structure of `fortiguardWf` block is documented below.
func (o LookupReplacemsggroupResultOutput) FortiguardWfs() GetReplacemsggroupFortiguardWfArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupFortiguardWf { return v.FortiguardWfs }).(GetReplacemsggroupFortiguardWfArrayOutput)
}

// Replacement message table entries. The structure of `ftp` block is documented below.
func (o LookupReplacemsggroupResultOutput) Ftps() GetReplacemsggroupFtpArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupFtp { return v.Ftps }).(GetReplacemsggroupFtpArrayOutput)
}

// Group type.
func (o LookupReplacemsggroupResultOutput) GroupType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) string { return v.GroupType }).(pulumi.StringOutput)
}

// Replacement message table entries. The structure of `http` block is documented below.
func (o LookupReplacemsggroupResultOutput) Https() GetReplacemsggroupHttpArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupHttp { return v.Https }).(GetReplacemsggroupHttpArrayOutput)
}

// Replacement message table entries. The structure of `icap` block is documented below.
func (o LookupReplacemsggroupResultOutput) Icaps() GetReplacemsggroupIcapArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupIcap { return v.Icaps }).(GetReplacemsggroupIcapArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupReplacemsggroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Replacement message table entries. The structure of `mail` block is documented below.
func (o LookupReplacemsggroupResultOutput) Mails() GetReplacemsggroupMailArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupMail { return v.Mails }).(GetReplacemsggroupMailArrayOutput)
}

// Replacement message table entries. The structure of `nacQuar` block is documented below.
func (o LookupReplacemsggroupResultOutput) NacQuars() GetReplacemsggroupNacQuarArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupNacQuar { return v.NacQuars }).(GetReplacemsggroupNacQuarArrayOutput)
}

// Group name.
func (o LookupReplacemsggroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Replacement message table entries. The structure of `nntp` block is documented below.
func (o LookupReplacemsggroupResultOutput) Nntps() GetReplacemsggroupNntpArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupNntp { return v.Nntps }).(GetReplacemsggroupNntpArrayOutput)
}

// Replacement message table entries. The structure of `spam` block is documented below.
func (o LookupReplacemsggroupResultOutput) Spams() GetReplacemsggroupSpamArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupSpam { return v.Spams }).(GetReplacemsggroupSpamArrayOutput)
}

// Replacement message table entries. The structure of `sslvpn` block is documented below.
func (o LookupReplacemsggroupResultOutput) Sslvpns() GetReplacemsggroupSslvpnArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupSslvpn { return v.Sslvpns }).(GetReplacemsggroupSslvpnArrayOutput)
}

// Replacement message table entries. The structure of `trafficQuota` block is documented below.
func (o LookupReplacemsggroupResultOutput) TrafficQuotas() GetReplacemsggroupTrafficQuotaArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupTrafficQuota { return v.TrafficQuotas }).(GetReplacemsggroupTrafficQuotaArrayOutput)
}

// Replacement message table entries. The structure of `utm` block is documented below.
func (o LookupReplacemsggroupResultOutput) Utms() GetReplacemsggroupUtmArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupUtm { return v.Utms }).(GetReplacemsggroupUtmArrayOutput)
}

func (o LookupReplacemsggroupResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Replacement message table entries. The structure of `webproxy` block is documented below.
func (o LookupReplacemsggroupResultOutput) Webproxies() GetReplacemsggroupWebproxyArrayOutput {
	return o.ApplyT(func(v LookupReplacemsggroupResult) []GetReplacemsggroupWebproxy { return v.Webproxies }).(GetReplacemsggroupWebproxyArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplacemsggroupResultOutput{})
}
