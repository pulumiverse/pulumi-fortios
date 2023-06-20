// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure access profiles for system administrators.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewSystemAccprofile(ctx, "test12", &fortios.SystemAccprofileArgs{
//				Admintimeout:         pulumi.Int(10),
//				AdmintimeoutOverride: pulumi.String("disable"),
//				Authgrp:              pulumi.String("read-write"),
//				Ftviewgrp:            pulumi.String("read-write"),
//				Fwgrp:                pulumi.String("custom"),
//				FwgrpPermission: &fortios.SystemAccprofileFwgrpPermissionArgs{
//					Address:  pulumi.String("read-write"),
//					Policy:   pulumi.String("read-write"),
//					Schedule: pulumi.String("none"),
//					Service:  pulumi.String("none"),
//				},
//				Loggrp: pulumi.String("read-write"),
//				LoggrpPermission: &fortios.SystemAccprofileLoggrpPermissionArgs{
//					Config:       pulumi.String("none"),
//					DataAccess:   pulumi.String("none"),
//					ReportAccess: pulumi.String("none"),
//					ThreatWeight: pulumi.String("none"),
//				},
//				Netgrp: pulumi.String("read-write"),
//				NetgrpPermission: &fortios.SystemAccprofileNetgrpPermissionArgs{
//					Cfg:           pulumi.String("none"),
//					PacketCapture: pulumi.String("none"),
//					RouteCfg:      pulumi.String("none"),
//				},
//				Scope:     pulumi.String("vdom"),
//				Secfabgrp: pulumi.String("read-write"),
//				Sysgrp:    pulumi.String("read-write"),
//				SysgrpPermission: &fortios.SystemAccprofileSysgrpPermissionArgs{
//					Admin: pulumi.String("none"),
//					Cfg:   pulumi.String("none"),
//					Mnt:   pulumi.String("none"),
//					Upd:   pulumi.String("none"),
//				},
//				Utmgrp: pulumi.String("custom"),
//				UtmgrpPermission: &fortios.SystemAccprofileUtmgrpPermissionArgs{
//					Antivirus:          pulumi.String("read-write"),
//					ApplicationControl: pulumi.String("none"),
//					DataLossPrevention: pulumi.String("none"),
//					Dnsfilter:          pulumi.String("none"),
//					EndpointControl:    pulumi.String("none"),
//					Icap:               pulumi.String("none"),
//					Ips:                pulumi.String("read-write"),
//					Voip:               pulumi.String("none"),
//					Waf:                pulumi.String("none"),
//					Webfilter:          pulumi.String("none"),
//				},
//				Vpngrp:    pulumi.String("read-write"),
//				Wanoptgrp: pulumi.String("read-write"),
//				Wifi:      pulumi.String("read-write"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # System Accprofile can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemAccprofile:SystemAccprofile labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemAccprofile:SystemAccprofile labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemAccprofile struct {
	pulumi.CustomResourceState

	// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
	Admintimeout pulumi.IntOutput `pulumi:"admintimeout"`
	// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
	AdmintimeoutOverride pulumi.StringOutput `pulumi:"admintimeoutOverride"`
	// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
	Authgrp pulumi.StringOutput `pulumi:"authgrp"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// FortiView. Valid values: `none`, `read`, `read-write`.
	Ftviewgrp pulumi.StringOutput `pulumi:"ftviewgrp"`
	// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Fwgrp pulumi.StringOutput `pulumi:"fwgrp"`
	// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
	FwgrpPermission SystemAccprofileFwgrpPermissionOutput `pulumi:"fwgrpPermission"`
	// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
	Loggrp pulumi.StringOutput `pulumi:"loggrp"`
	// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
	LoggrpPermission SystemAccprofileLoggrpPermissionOutput `pulumi:"loggrpPermission"`
	// Profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Netgrp pulumi.StringOutput `pulumi:"netgrp"`
	// Custom network permission. The structure of `netgrpPermission` block is documented below.
	NetgrpPermission SystemAccprofileNetgrpPermissionOutput `pulumi:"netgrpPermission"`
	// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// Security Fabric. Valid values: `none`, `read`, `read-write`.
	Secfabgrp pulumi.StringOutput `pulumi:"secfabgrp"`
	// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Sysgrp pulumi.StringOutput `pulumi:"sysgrp"`
	// Custom system permission. The structure of `sysgrpPermission` block is documented below.
	SysgrpPermission SystemAccprofileSysgrpPermissionOutput `pulumi:"sysgrpPermission"`
	// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
	SystemDiagnostics pulumi.StringOutput `pulumi:"systemDiagnostics"`
	// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
	SystemExecuteSsh pulumi.StringOutput `pulumi:"systemExecuteSsh"`
	// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
	SystemExecuteTelnet pulumi.StringOutput `pulumi:"systemExecuteTelnet"`
	// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
	Utmgrp pulumi.StringOutput `pulumi:"utmgrp"`
	// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
	UtmgrpPermission SystemAccprofileUtmgrpPermissionOutput `pulumi:"utmgrpPermission"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
	Vpngrp pulumi.StringOutput `pulumi:"vpngrp"`
	// Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.
	Wanoptgrp pulumi.StringOutput `pulumi:"wanoptgrp"`
	// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
	Wifi pulumi.StringOutput `pulumi:"wifi"`
}

// NewSystemAccprofile registers a new resource with the given unique name, arguments, and options.
func NewSystemAccprofile(ctx *pulumi.Context,
	name string, args *SystemAccprofileArgs, opts ...pulumi.ResourceOption) (*SystemAccprofile, error) {
	if args == nil {
		args = &SystemAccprofileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemAccprofile
	err := ctx.RegisterResource("fortios:index/systemAccprofile:SystemAccprofile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAccprofile gets an existing SystemAccprofile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAccprofile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAccprofileState, opts ...pulumi.ResourceOption) (*SystemAccprofile, error) {
	var resource SystemAccprofile
	err := ctx.ReadResource("fortios:index/systemAccprofile:SystemAccprofile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAccprofile resources.
type systemAccprofileState struct {
	// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
	Admintimeout *int `pulumi:"admintimeout"`
	// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
	AdmintimeoutOverride *string `pulumi:"admintimeoutOverride"`
	// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
	Authgrp *string `pulumi:"authgrp"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// FortiView. Valid values: `none`, `read`, `read-write`.
	Ftviewgrp *string `pulumi:"ftviewgrp"`
	// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Fwgrp *string `pulumi:"fwgrp"`
	// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
	FwgrpPermission *SystemAccprofileFwgrpPermission `pulumi:"fwgrpPermission"`
	// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
	Loggrp *string `pulumi:"loggrp"`
	// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
	LoggrpPermission *SystemAccprofileLoggrpPermission `pulumi:"loggrpPermission"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Netgrp *string `pulumi:"netgrp"`
	// Custom network permission. The structure of `netgrpPermission` block is documented below.
	NetgrpPermission *SystemAccprofileNetgrpPermission `pulumi:"netgrpPermission"`
	// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
	Scope *string `pulumi:"scope"`
	// Security Fabric. Valid values: `none`, `read`, `read-write`.
	Secfabgrp *string `pulumi:"secfabgrp"`
	// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Sysgrp *string `pulumi:"sysgrp"`
	// Custom system permission. The structure of `sysgrpPermission` block is documented below.
	SysgrpPermission *SystemAccprofileSysgrpPermission `pulumi:"sysgrpPermission"`
	// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
	SystemDiagnostics *string `pulumi:"systemDiagnostics"`
	// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
	SystemExecuteSsh *string `pulumi:"systemExecuteSsh"`
	// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
	SystemExecuteTelnet *string `pulumi:"systemExecuteTelnet"`
	// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
	Utmgrp *string `pulumi:"utmgrp"`
	// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
	UtmgrpPermission *SystemAccprofileUtmgrpPermission `pulumi:"utmgrpPermission"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
	Vpngrp *string `pulumi:"vpngrp"`
	// Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.
	Wanoptgrp *string `pulumi:"wanoptgrp"`
	// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
	Wifi *string `pulumi:"wifi"`
}

type SystemAccprofileState struct {
	// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
	Admintimeout pulumi.IntPtrInput
	// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
	AdmintimeoutOverride pulumi.StringPtrInput
	// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
	Authgrp pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// FortiView. Valid values: `none`, `read`, `read-write`.
	Ftviewgrp pulumi.StringPtrInput
	// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Fwgrp pulumi.StringPtrInput
	// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
	FwgrpPermission SystemAccprofileFwgrpPermissionPtrInput
	// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
	Loggrp pulumi.StringPtrInput
	// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
	LoggrpPermission SystemAccprofileLoggrpPermissionPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Netgrp pulumi.StringPtrInput
	// Custom network permission. The structure of `netgrpPermission` block is documented below.
	NetgrpPermission SystemAccprofileNetgrpPermissionPtrInput
	// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
	Scope pulumi.StringPtrInput
	// Security Fabric. Valid values: `none`, `read`, `read-write`.
	Secfabgrp pulumi.StringPtrInput
	// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Sysgrp pulumi.StringPtrInput
	// Custom system permission. The structure of `sysgrpPermission` block is documented below.
	SysgrpPermission SystemAccprofileSysgrpPermissionPtrInput
	// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
	SystemDiagnostics pulumi.StringPtrInput
	// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
	SystemExecuteSsh pulumi.StringPtrInput
	// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
	SystemExecuteTelnet pulumi.StringPtrInput
	// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
	Utmgrp pulumi.StringPtrInput
	// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
	UtmgrpPermission SystemAccprofileUtmgrpPermissionPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
	Vpngrp pulumi.StringPtrInput
	// Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.
	Wanoptgrp pulumi.StringPtrInput
	// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
	Wifi pulumi.StringPtrInput
}

func (SystemAccprofileState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAccprofileState)(nil)).Elem()
}

type systemAccprofileArgs struct {
	// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
	Admintimeout *int `pulumi:"admintimeout"`
	// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
	AdmintimeoutOverride *string `pulumi:"admintimeoutOverride"`
	// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
	Authgrp *string `pulumi:"authgrp"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// FortiView. Valid values: `none`, `read`, `read-write`.
	Ftviewgrp *string `pulumi:"ftviewgrp"`
	// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Fwgrp *string `pulumi:"fwgrp"`
	// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
	FwgrpPermission *SystemAccprofileFwgrpPermission `pulumi:"fwgrpPermission"`
	// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
	Loggrp *string `pulumi:"loggrp"`
	// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
	LoggrpPermission *SystemAccprofileLoggrpPermission `pulumi:"loggrpPermission"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Netgrp *string `pulumi:"netgrp"`
	// Custom network permission. The structure of `netgrpPermission` block is documented below.
	NetgrpPermission *SystemAccprofileNetgrpPermission `pulumi:"netgrpPermission"`
	// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
	Scope *string `pulumi:"scope"`
	// Security Fabric. Valid values: `none`, `read`, `read-write`.
	Secfabgrp *string `pulumi:"secfabgrp"`
	// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Sysgrp *string `pulumi:"sysgrp"`
	// Custom system permission. The structure of `sysgrpPermission` block is documented below.
	SysgrpPermission *SystemAccprofileSysgrpPermission `pulumi:"sysgrpPermission"`
	// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
	SystemDiagnostics *string `pulumi:"systemDiagnostics"`
	// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
	SystemExecuteSsh *string `pulumi:"systemExecuteSsh"`
	// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
	SystemExecuteTelnet *string `pulumi:"systemExecuteTelnet"`
	// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
	Utmgrp *string `pulumi:"utmgrp"`
	// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
	UtmgrpPermission *SystemAccprofileUtmgrpPermission `pulumi:"utmgrpPermission"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
	Vpngrp *string `pulumi:"vpngrp"`
	// Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.
	Wanoptgrp *string `pulumi:"wanoptgrp"`
	// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
	Wifi *string `pulumi:"wifi"`
}

// The set of arguments for constructing a SystemAccprofile resource.
type SystemAccprofileArgs struct {
	// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
	Admintimeout pulumi.IntPtrInput
	// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
	AdmintimeoutOverride pulumi.StringPtrInput
	// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
	Authgrp pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// FortiView. Valid values: `none`, `read`, `read-write`.
	Ftviewgrp pulumi.StringPtrInput
	// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Fwgrp pulumi.StringPtrInput
	// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
	FwgrpPermission SystemAccprofileFwgrpPermissionPtrInput
	// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
	Loggrp pulumi.StringPtrInput
	// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
	LoggrpPermission SystemAccprofileLoggrpPermissionPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Netgrp pulumi.StringPtrInput
	// Custom network permission. The structure of `netgrpPermission` block is documented below.
	NetgrpPermission SystemAccprofileNetgrpPermissionPtrInput
	// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
	Scope pulumi.StringPtrInput
	// Security Fabric. Valid values: `none`, `read`, `read-write`.
	Secfabgrp pulumi.StringPtrInput
	// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Sysgrp pulumi.StringPtrInput
	// Custom system permission. The structure of `sysgrpPermission` block is documented below.
	SysgrpPermission SystemAccprofileSysgrpPermissionPtrInput
	// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
	SystemDiagnostics pulumi.StringPtrInput
	// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
	SystemExecuteSsh pulumi.StringPtrInput
	// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
	SystemExecuteTelnet pulumi.StringPtrInput
	// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
	Utmgrp pulumi.StringPtrInput
	// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
	UtmgrpPermission SystemAccprofileUtmgrpPermissionPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
	Vpngrp pulumi.StringPtrInput
	// Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.
	Wanoptgrp pulumi.StringPtrInput
	// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
	Wifi pulumi.StringPtrInput
}

func (SystemAccprofileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAccprofileArgs)(nil)).Elem()
}

type SystemAccprofileInput interface {
	pulumi.Input

	ToSystemAccprofileOutput() SystemAccprofileOutput
	ToSystemAccprofileOutputWithContext(ctx context.Context) SystemAccprofileOutput
}

func (*SystemAccprofile) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAccprofile)(nil)).Elem()
}

func (i *SystemAccprofile) ToSystemAccprofileOutput() SystemAccprofileOutput {
	return i.ToSystemAccprofileOutputWithContext(context.Background())
}

func (i *SystemAccprofile) ToSystemAccprofileOutputWithContext(ctx context.Context) SystemAccprofileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAccprofileOutput)
}

// SystemAccprofileArrayInput is an input type that accepts SystemAccprofileArray and SystemAccprofileArrayOutput values.
// You can construct a concrete instance of `SystemAccprofileArrayInput` via:
//
//	SystemAccprofileArray{ SystemAccprofileArgs{...} }
type SystemAccprofileArrayInput interface {
	pulumi.Input

	ToSystemAccprofileArrayOutput() SystemAccprofileArrayOutput
	ToSystemAccprofileArrayOutputWithContext(context.Context) SystemAccprofileArrayOutput
}

type SystemAccprofileArray []SystemAccprofileInput

func (SystemAccprofileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAccprofile)(nil)).Elem()
}

func (i SystemAccprofileArray) ToSystemAccprofileArrayOutput() SystemAccprofileArrayOutput {
	return i.ToSystemAccprofileArrayOutputWithContext(context.Background())
}

func (i SystemAccprofileArray) ToSystemAccprofileArrayOutputWithContext(ctx context.Context) SystemAccprofileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAccprofileArrayOutput)
}

// SystemAccprofileMapInput is an input type that accepts SystemAccprofileMap and SystemAccprofileMapOutput values.
// You can construct a concrete instance of `SystemAccprofileMapInput` via:
//
//	SystemAccprofileMap{ "key": SystemAccprofileArgs{...} }
type SystemAccprofileMapInput interface {
	pulumi.Input

	ToSystemAccprofileMapOutput() SystemAccprofileMapOutput
	ToSystemAccprofileMapOutputWithContext(context.Context) SystemAccprofileMapOutput
}

type SystemAccprofileMap map[string]SystemAccprofileInput

func (SystemAccprofileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAccprofile)(nil)).Elem()
}

func (i SystemAccprofileMap) ToSystemAccprofileMapOutput() SystemAccprofileMapOutput {
	return i.ToSystemAccprofileMapOutputWithContext(context.Background())
}

func (i SystemAccprofileMap) ToSystemAccprofileMapOutputWithContext(ctx context.Context) SystemAccprofileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAccprofileMapOutput)
}

type SystemAccprofileOutput struct{ *pulumi.OutputState }

func (SystemAccprofileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAccprofile)(nil)).Elem()
}

func (o SystemAccprofileOutput) ToSystemAccprofileOutput() SystemAccprofileOutput {
	return o
}

func (o SystemAccprofileOutput) ToSystemAccprofileOutputWithContext(ctx context.Context) SystemAccprofileOutput {
	return o
}

// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
func (o SystemAccprofileOutput) Admintimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.IntOutput { return v.Admintimeout }).(pulumi.IntOutput)
}

// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
func (o SystemAccprofileOutput) AdmintimeoutOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.AdmintimeoutOverride }).(pulumi.StringOutput)
}

// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
func (o SystemAccprofileOutput) Authgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.Authgrp }).(pulumi.StringOutput)
}

// Comment.
func (o SystemAccprofileOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// FortiView. Valid values: `none`, `read`, `read-write`.
func (o SystemAccprofileOutput) Ftviewgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.Ftviewgrp }).(pulumi.StringOutput)
}

// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
func (o SystemAccprofileOutput) Fwgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.Fwgrp }).(pulumi.StringOutput)
}

// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
func (o SystemAccprofileOutput) FwgrpPermission() SystemAccprofileFwgrpPermissionOutput {
	return o.ApplyT(func(v *SystemAccprofile) SystemAccprofileFwgrpPermissionOutput { return v.FwgrpPermission }).(SystemAccprofileFwgrpPermissionOutput)
}

// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
func (o SystemAccprofileOutput) Loggrp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.Loggrp }).(pulumi.StringOutput)
}

// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
func (o SystemAccprofileOutput) LoggrpPermission() SystemAccprofileLoggrpPermissionOutput {
	return o.ApplyT(func(v *SystemAccprofile) SystemAccprofileLoggrpPermissionOutput { return v.LoggrpPermission }).(SystemAccprofileLoggrpPermissionOutput)
}

// Profile name.
func (o SystemAccprofileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
func (o SystemAccprofileOutput) Netgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.Netgrp }).(pulumi.StringOutput)
}

// Custom network permission. The structure of `netgrpPermission` block is documented below.
func (o SystemAccprofileOutput) NetgrpPermission() SystemAccprofileNetgrpPermissionOutput {
	return o.ApplyT(func(v *SystemAccprofile) SystemAccprofileNetgrpPermissionOutput { return v.NetgrpPermission }).(SystemAccprofileNetgrpPermissionOutput)
}

// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
func (o SystemAccprofileOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

// Security Fabric. Valid values: `none`, `read`, `read-write`.
func (o SystemAccprofileOutput) Secfabgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.Secfabgrp }).(pulumi.StringOutput)
}

// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
func (o SystemAccprofileOutput) Sysgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.Sysgrp }).(pulumi.StringOutput)
}

// Custom system permission. The structure of `sysgrpPermission` block is documented below.
func (o SystemAccprofileOutput) SysgrpPermission() SystemAccprofileSysgrpPermissionOutput {
	return o.ApplyT(func(v *SystemAccprofile) SystemAccprofileSysgrpPermissionOutput { return v.SysgrpPermission }).(SystemAccprofileSysgrpPermissionOutput)
}

// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
func (o SystemAccprofileOutput) SystemDiagnostics() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.SystemDiagnostics }).(pulumi.StringOutput)
}

// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
func (o SystemAccprofileOutput) SystemExecuteSsh() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.SystemExecuteSsh }).(pulumi.StringOutput)
}

// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
func (o SystemAccprofileOutput) SystemExecuteTelnet() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.SystemExecuteTelnet }).(pulumi.StringOutput)
}

// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
func (o SystemAccprofileOutput) Utmgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.Utmgrp }).(pulumi.StringOutput)
}

// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
func (o SystemAccprofileOutput) UtmgrpPermission() SystemAccprofileUtmgrpPermissionOutput {
	return o.ApplyT(func(v *SystemAccprofile) SystemAccprofileUtmgrpPermissionOutput { return v.UtmgrpPermission }).(SystemAccprofileUtmgrpPermissionOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemAccprofileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
func (o SystemAccprofileOutput) Vpngrp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.Vpngrp }).(pulumi.StringOutput)
}

// Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.
func (o SystemAccprofileOutput) Wanoptgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.Wanoptgrp }).(pulumi.StringOutput)
}

// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
func (o SystemAccprofileOutput) Wifi() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccprofile) pulumi.StringOutput { return v.Wifi }).(pulumi.StringOutput)
}

type SystemAccprofileArrayOutput struct{ *pulumi.OutputState }

func (SystemAccprofileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAccprofile)(nil)).Elem()
}

func (o SystemAccprofileArrayOutput) ToSystemAccprofileArrayOutput() SystemAccprofileArrayOutput {
	return o
}

func (o SystemAccprofileArrayOutput) ToSystemAccprofileArrayOutputWithContext(ctx context.Context) SystemAccprofileArrayOutput {
	return o
}

func (o SystemAccprofileArrayOutput) Index(i pulumi.IntInput) SystemAccprofileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAccprofile {
		return vs[0].([]*SystemAccprofile)[vs[1].(int)]
	}).(SystemAccprofileOutput)
}

type SystemAccprofileMapOutput struct{ *pulumi.OutputState }

func (SystemAccprofileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAccprofile)(nil)).Elem()
}

func (o SystemAccprofileMapOutput) ToSystemAccprofileMapOutput() SystemAccprofileMapOutput {
	return o
}

func (o SystemAccprofileMapOutput) ToSystemAccprofileMapOutputWithContext(ctx context.Context) SystemAccprofileMapOutput {
	return o
}

func (o SystemAccprofileMapOutput) MapIndex(k pulumi.StringInput) SystemAccprofileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAccprofile {
		return vs[0].(map[string]*SystemAccprofile)[vs[1].(string)]
	}).(SystemAccprofileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAccprofileInput)(nil)).Elem(), &SystemAccprofile{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAccprofileArrayInput)(nil)).Elem(), SystemAccprofileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAccprofileMapInput)(nil)).Elem(), SystemAccprofileMap{})
	pulumi.RegisterOutputType(SystemAccprofileOutput{})
	pulumi.RegisterOutputType(SystemAccprofileArrayOutput{})
	pulumi.RegisterOutputType(SystemAccprofileMapOutput{})
}