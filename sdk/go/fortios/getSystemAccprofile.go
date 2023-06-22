// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system accprofile
func LookupSystemAccprofile(ctx *pulumi.Context, args *LookupSystemAccprofileArgs, opts ...pulumi.InvokeOption) (*LookupSystemAccprofileResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemAccprofileResult
	err := ctx.Invoke("fortios:index/getSystemAccprofile:getSystemAccprofile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemAccprofile.
type LookupSystemAccprofileArgs struct {
	// Specify the name of the desired system accprofile.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemAccprofile.
type LookupSystemAccprofileResult struct {
	// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
	Admintimeout int `pulumi:"admintimeout"`
	// Enable/disable overriding the global administrator idle timeout.
	AdmintimeoutOverride string `pulumi:"admintimeoutOverride"`
	// Administrator access to Users and Devices.
	Authgrp string `pulumi:"authgrp"`
	// Comment.
	Comments string `pulumi:"comments"`
	// FortiView.
	Ftviewgrp string `pulumi:"ftviewgrp"`
	// Administrator access to the Firewall configuration.
	Fwgrp string `pulumi:"fwgrp"`
	// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
	FwgrpPermissions []GetSystemAccprofileFwgrpPermission `pulumi:"fwgrpPermissions"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Administrator access to Logging and Reporting including viewing log messages.
	Loggrp string `pulumi:"loggrp"`
	// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
	LoggrpPermissions []GetSystemAccprofileLoggrpPermission `pulumi:"loggrpPermissions"`
	// Profile name.
	Name string `pulumi:"name"`
	// Network Configuration.
	Netgrp string `pulumi:"netgrp"`
	// Custom network permission. The structure of `netgrpPermission` block is documented below.
	NetgrpPermissions []GetSystemAccprofileNetgrpPermission `pulumi:"netgrpPermissions"`
	// Scope of admin access: global or specific VDOM(s).
	Scope string `pulumi:"scope"`
	// Security Fabric.
	Secfabgrp string `pulumi:"secfabgrp"`
	// System Configuration.
	Sysgrp string `pulumi:"sysgrp"`
	// Custom system permission. The structure of `sysgrpPermission` block is documented below.
	SysgrpPermissions []GetSystemAccprofileSysgrpPermission `pulumi:"sysgrpPermissions"`
	// Enable/disable permission to run system diagnostic commands.
	SystemDiagnostics string `pulumi:"systemDiagnostics"`
	// Enable/disable permission to execute SSH commands.
	SystemExecuteSsh string `pulumi:"systemExecuteSsh"`
	// Enable/disable permission to execute TELNET commands.
	SystemExecuteTelnet string `pulumi:"systemExecuteTelnet"`
	// Administrator access to Security Profiles.
	Utmgrp string `pulumi:"utmgrp"`
	// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
	UtmgrpPermissions []GetSystemAccprofileUtmgrpPermission `pulumi:"utmgrpPermissions"`
	Vdomparam         *string                               `pulumi:"vdomparam"`
	// Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
	Vpngrp string `pulumi:"vpngrp"`
	// Administrator access to WAN Opt & Cache.
	Wanoptgrp string `pulumi:"wanoptgrp"`
	// Administrator access to the WiFi controller and Switch controller.
	Wifi string `pulumi:"wifi"`
}

func LookupSystemAccprofileOutput(ctx *pulumi.Context, args LookupSystemAccprofileOutputArgs, opts ...pulumi.InvokeOption) LookupSystemAccprofileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemAccprofileResult, error) {
			args := v.(LookupSystemAccprofileArgs)
			r, err := LookupSystemAccprofile(ctx, &args, opts...)
			var s LookupSystemAccprofileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemAccprofileResultOutput)
}

// A collection of arguments for invoking getSystemAccprofile.
type LookupSystemAccprofileOutputArgs struct {
	// Specify the name of the desired system accprofile.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemAccprofileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAccprofileArgs)(nil)).Elem()
}

// A collection of values returned by getSystemAccprofile.
type LookupSystemAccprofileResultOutput struct{ *pulumi.OutputState }

func (LookupSystemAccprofileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAccprofileResult)(nil)).Elem()
}

func (o LookupSystemAccprofileResultOutput) ToLookupSystemAccprofileResultOutput() LookupSystemAccprofileResultOutput {
	return o
}

func (o LookupSystemAccprofileResultOutput) ToLookupSystemAccprofileResultOutputWithContext(ctx context.Context) LookupSystemAccprofileResultOutput {
	return o
}

// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
func (o LookupSystemAccprofileResultOutput) Admintimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) int { return v.Admintimeout }).(pulumi.IntOutput)
}

// Enable/disable overriding the global administrator idle timeout.
func (o LookupSystemAccprofileResultOutput) AdmintimeoutOverride() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.AdmintimeoutOverride }).(pulumi.StringOutput)
}

// Administrator access to Users and Devices.
func (o LookupSystemAccprofileResultOutput) Authgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Authgrp }).(pulumi.StringOutput)
}

// Comment.
func (o LookupSystemAccprofileResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Comments }).(pulumi.StringOutput)
}

// FortiView.
func (o LookupSystemAccprofileResultOutput) Ftviewgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Ftviewgrp }).(pulumi.StringOutput)
}

// Administrator access to the Firewall configuration.
func (o LookupSystemAccprofileResultOutput) Fwgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Fwgrp }).(pulumi.StringOutput)
}

// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
func (o LookupSystemAccprofileResultOutput) FwgrpPermissions() GetSystemAccprofileFwgrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) []GetSystemAccprofileFwgrpPermission { return v.FwgrpPermissions }).(GetSystemAccprofileFwgrpPermissionArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemAccprofileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Id }).(pulumi.StringOutput)
}

// Administrator access to Logging and Reporting including viewing log messages.
func (o LookupSystemAccprofileResultOutput) Loggrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Loggrp }).(pulumi.StringOutput)
}

// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
func (o LookupSystemAccprofileResultOutput) LoggrpPermissions() GetSystemAccprofileLoggrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) []GetSystemAccprofileLoggrpPermission { return v.LoggrpPermissions }).(GetSystemAccprofileLoggrpPermissionArrayOutput)
}

// Profile name.
func (o LookupSystemAccprofileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network Configuration.
func (o LookupSystemAccprofileResultOutput) Netgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Netgrp }).(pulumi.StringOutput)
}

// Custom network permission. The structure of `netgrpPermission` block is documented below.
func (o LookupSystemAccprofileResultOutput) NetgrpPermissions() GetSystemAccprofileNetgrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) []GetSystemAccprofileNetgrpPermission { return v.NetgrpPermissions }).(GetSystemAccprofileNetgrpPermissionArrayOutput)
}

// Scope of admin access: global or specific VDOM(s).
func (o LookupSystemAccprofileResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Scope }).(pulumi.StringOutput)
}

// Security Fabric.
func (o LookupSystemAccprofileResultOutput) Secfabgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Secfabgrp }).(pulumi.StringOutput)
}

// System Configuration.
func (o LookupSystemAccprofileResultOutput) Sysgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Sysgrp }).(pulumi.StringOutput)
}

// Custom system permission. The structure of `sysgrpPermission` block is documented below.
func (o LookupSystemAccprofileResultOutput) SysgrpPermissions() GetSystemAccprofileSysgrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) []GetSystemAccprofileSysgrpPermission { return v.SysgrpPermissions }).(GetSystemAccprofileSysgrpPermissionArrayOutput)
}

// Enable/disable permission to run system diagnostic commands.
func (o LookupSystemAccprofileResultOutput) SystemDiagnostics() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.SystemDiagnostics }).(pulumi.StringOutput)
}

// Enable/disable permission to execute SSH commands.
func (o LookupSystemAccprofileResultOutput) SystemExecuteSsh() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.SystemExecuteSsh }).(pulumi.StringOutput)
}

// Enable/disable permission to execute TELNET commands.
func (o LookupSystemAccprofileResultOutput) SystemExecuteTelnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.SystemExecuteTelnet }).(pulumi.StringOutput)
}

// Administrator access to Security Profiles.
func (o LookupSystemAccprofileResultOutput) Utmgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Utmgrp }).(pulumi.StringOutput)
}

// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
func (o LookupSystemAccprofileResultOutput) UtmgrpPermissions() GetSystemAccprofileUtmgrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) []GetSystemAccprofileUtmgrpPermission { return v.UtmgrpPermissions }).(GetSystemAccprofileUtmgrpPermissionArrayOutput)
}

func (o LookupSystemAccprofileResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
func (o LookupSystemAccprofileResultOutput) Vpngrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Vpngrp }).(pulumi.StringOutput)
}

// Administrator access to WAN Opt & Cache.
func (o LookupSystemAccprofileResultOutput) Wanoptgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Wanoptgrp }).(pulumi.StringOutput)
}

// Administrator access to the WiFi controller and Switch controller.
func (o LookupSystemAccprofileResultOutput) Wifi() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Wifi }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemAccprofileResultOutput{})
}
