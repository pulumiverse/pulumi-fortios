// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios system accprofile
func LookupAccprofile(ctx *pulumi.Context, args *LookupAccprofileArgs, opts ...pulumi.InvokeOption) (*LookupAccprofileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccprofileResult
	err := ctx.Invoke("fortios:system/getAccprofile:getAccprofile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccprofile.
type LookupAccprofileArgs struct {
	// Specify the name of the desired system accprofile.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getAccprofile.
type LookupAccprofileResult struct {
	// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
	Admintimeout int `pulumi:"admintimeout"`
	// Enable/disable overriding the global administrator idle timeout.
	AdmintimeoutOverride string `pulumi:"admintimeoutOverride"`
	// Administrator access to Users and Devices.
	Authgrp string `pulumi:"authgrp"`
	// Enable/disable permission to run config commands.
	CliConfig string `pulumi:"cliConfig"`
	// Enable/disable permission to run diagnostic commands.
	CliDiagnose string `pulumi:"cliDiagnose"`
	// Enable/disable permission to run execute commands.
	CliExec string `pulumi:"cliExec"`
	// Enable/disable permission to run get commands.
	CliGet string `pulumi:"cliGet"`
	// Enable/disable permission to run show commands.
	CliShow string `pulumi:"cliShow"`
	// Comment.
	Comments string `pulumi:"comments"`
	// FortiView.
	Ftviewgrp string `pulumi:"ftviewgrp"`
	// Administrator access to the Firewall configuration.
	Fwgrp string `pulumi:"fwgrp"`
	// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
	FwgrpPermissions []GetAccprofileFwgrpPermission `pulumi:"fwgrpPermissions"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Administrator access to Logging and Reporting including viewing log messages.
	Loggrp string `pulumi:"loggrp"`
	// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
	LoggrpPermissions []GetAccprofileLoggrpPermission `pulumi:"loggrpPermissions"`
	// Profile name.
	Name string `pulumi:"name"`
	// Network Configuration.
	Netgrp string `pulumi:"netgrp"`
	// Custom network permission. The structure of `netgrpPermission` block is documented below.
	NetgrpPermissions []GetAccprofileNetgrpPermission `pulumi:"netgrpPermissions"`
	// Scope of admin access: global or specific VDOM(s).
	Scope string `pulumi:"scope"`
	// Security Fabric.
	Secfabgrp string `pulumi:"secfabgrp"`
	// System Configuration.
	Sysgrp string `pulumi:"sysgrp"`
	// Custom system permission. The structure of `sysgrpPermission` block is documented below.
	SysgrpPermissions []GetAccprofileSysgrpPermission `pulumi:"sysgrpPermissions"`
	// Enable/disable permission to run system diagnostic commands.
	SystemDiagnostics string `pulumi:"systemDiagnostics"`
	// Enable/disable permission to execute SSH commands.
	SystemExecuteSsh string `pulumi:"systemExecuteSsh"`
	// Enable/disable permission to execute TELNET commands.
	SystemExecuteTelnet string `pulumi:"systemExecuteTelnet"`
	// Administrator access to Security Profiles.
	Utmgrp string `pulumi:"utmgrp"`
	// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
	UtmgrpPermissions []GetAccprofileUtmgrpPermission `pulumi:"utmgrpPermissions"`
	Vdomparam         *string                         `pulumi:"vdomparam"`
	// Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
	Vpngrp string `pulumi:"vpngrp"`
	// Administrator access to WAN Opt & Cache.
	Wanoptgrp string `pulumi:"wanoptgrp"`
	// Administrator access to the WiFi controller and Switch controller.
	Wifi string `pulumi:"wifi"`
}

func LookupAccprofileOutput(ctx *pulumi.Context, args LookupAccprofileOutputArgs, opts ...pulumi.InvokeOption) LookupAccprofileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccprofileResult, error) {
			args := v.(LookupAccprofileArgs)
			r, err := LookupAccprofile(ctx, &args, opts...)
			var s LookupAccprofileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccprofileResultOutput)
}

// A collection of arguments for invoking getAccprofile.
type LookupAccprofileOutputArgs struct {
	// Specify the name of the desired system accprofile.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupAccprofileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccprofileArgs)(nil)).Elem()
}

// A collection of values returned by getAccprofile.
type LookupAccprofileResultOutput struct{ *pulumi.OutputState }

func (LookupAccprofileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccprofileResult)(nil)).Elem()
}

func (o LookupAccprofileResultOutput) ToLookupAccprofileResultOutput() LookupAccprofileResultOutput {
	return o
}

func (o LookupAccprofileResultOutput) ToLookupAccprofileResultOutputWithContext(ctx context.Context) LookupAccprofileResultOutput {
	return o
}

// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
func (o LookupAccprofileResultOutput) Admintimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAccprofileResult) int { return v.Admintimeout }).(pulumi.IntOutput)
}

// Enable/disable overriding the global administrator idle timeout.
func (o LookupAccprofileResultOutput) AdmintimeoutOverride() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.AdmintimeoutOverride }).(pulumi.StringOutput)
}

// Administrator access to Users and Devices.
func (o LookupAccprofileResultOutput) Authgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Authgrp }).(pulumi.StringOutput)
}

// Enable/disable permission to run config commands.
func (o LookupAccprofileResultOutput) CliConfig() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.CliConfig }).(pulumi.StringOutput)
}

// Enable/disable permission to run diagnostic commands.
func (o LookupAccprofileResultOutput) CliDiagnose() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.CliDiagnose }).(pulumi.StringOutput)
}

// Enable/disable permission to run execute commands.
func (o LookupAccprofileResultOutput) CliExec() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.CliExec }).(pulumi.StringOutput)
}

// Enable/disable permission to run get commands.
func (o LookupAccprofileResultOutput) CliGet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.CliGet }).(pulumi.StringOutput)
}

// Enable/disable permission to run show commands.
func (o LookupAccprofileResultOutput) CliShow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.CliShow }).(pulumi.StringOutput)
}

// Comment.
func (o LookupAccprofileResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Comments }).(pulumi.StringOutput)
}

// FortiView.
func (o LookupAccprofileResultOutput) Ftviewgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Ftviewgrp }).(pulumi.StringOutput)
}

// Administrator access to the Firewall configuration.
func (o LookupAccprofileResultOutput) Fwgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Fwgrp }).(pulumi.StringOutput)
}

// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
func (o LookupAccprofileResultOutput) FwgrpPermissions() GetAccprofileFwgrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupAccprofileResult) []GetAccprofileFwgrpPermission { return v.FwgrpPermissions }).(GetAccprofileFwgrpPermissionArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAccprofileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Id }).(pulumi.StringOutput)
}

// Administrator access to Logging and Reporting including viewing log messages.
func (o LookupAccprofileResultOutput) Loggrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Loggrp }).(pulumi.StringOutput)
}

// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
func (o LookupAccprofileResultOutput) LoggrpPermissions() GetAccprofileLoggrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupAccprofileResult) []GetAccprofileLoggrpPermission { return v.LoggrpPermissions }).(GetAccprofileLoggrpPermissionArrayOutput)
}

// Profile name.
func (o LookupAccprofileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network Configuration.
func (o LookupAccprofileResultOutput) Netgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Netgrp }).(pulumi.StringOutput)
}

// Custom network permission. The structure of `netgrpPermission` block is documented below.
func (o LookupAccprofileResultOutput) NetgrpPermissions() GetAccprofileNetgrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupAccprofileResult) []GetAccprofileNetgrpPermission { return v.NetgrpPermissions }).(GetAccprofileNetgrpPermissionArrayOutput)
}

// Scope of admin access: global or specific VDOM(s).
func (o LookupAccprofileResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Scope }).(pulumi.StringOutput)
}

// Security Fabric.
func (o LookupAccprofileResultOutput) Secfabgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Secfabgrp }).(pulumi.StringOutput)
}

// System Configuration.
func (o LookupAccprofileResultOutput) Sysgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Sysgrp }).(pulumi.StringOutput)
}

// Custom system permission. The structure of `sysgrpPermission` block is documented below.
func (o LookupAccprofileResultOutput) SysgrpPermissions() GetAccprofileSysgrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupAccprofileResult) []GetAccprofileSysgrpPermission { return v.SysgrpPermissions }).(GetAccprofileSysgrpPermissionArrayOutput)
}

// Enable/disable permission to run system diagnostic commands.
func (o LookupAccprofileResultOutput) SystemDiagnostics() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.SystemDiagnostics }).(pulumi.StringOutput)
}

// Enable/disable permission to execute SSH commands.
func (o LookupAccprofileResultOutput) SystemExecuteSsh() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.SystemExecuteSsh }).(pulumi.StringOutput)
}

// Enable/disable permission to execute TELNET commands.
func (o LookupAccprofileResultOutput) SystemExecuteTelnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.SystemExecuteTelnet }).(pulumi.StringOutput)
}

// Administrator access to Security Profiles.
func (o LookupAccprofileResultOutput) Utmgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Utmgrp }).(pulumi.StringOutput)
}

// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
func (o LookupAccprofileResultOutput) UtmgrpPermissions() GetAccprofileUtmgrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupAccprofileResult) []GetAccprofileUtmgrpPermission { return v.UtmgrpPermissions }).(GetAccprofileUtmgrpPermissionArrayOutput)
}

func (o LookupAccprofileResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccprofileResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
func (o LookupAccprofileResultOutput) Vpngrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Vpngrp }).(pulumi.StringOutput)
}

// Administrator access to WAN Opt & Cache.
func (o LookupAccprofileResultOutput) Wanoptgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Wanoptgrp }).(pulumi.StringOutput)
}

// Administrator access to the WiFi controller and Switch controller.
func (o LookupAccprofileResultOutput) Wifi() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccprofileResult) string { return v.Wifi }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccprofileResultOutput{})
}
