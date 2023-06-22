// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system resourcelimits
func LookupSystemResourcelimits(ctx *pulumi.Context, args *LookupSystemResourcelimitsArgs, opts ...pulumi.InvokeOption) (*LookupSystemResourcelimitsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemResourcelimitsResult
	err := ctx.Invoke("fortios:index/getSystemResourcelimits:getSystemResourcelimits", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemResourcelimits.
type LookupSystemResourcelimitsArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemResourcelimits.
type LookupSystemResourcelimitsResult struct {
	// Maximum number of firewall custom services.
	CustomService int `pulumi:"customService"`
	// Maximum number of dial-up tunnels.
	DialupTunnel int `pulumi:"dialupTunnel"`
	// Maximum number of firewall addresses (IPv4, IPv6, multicast).
	FirewallAddress int `pulumi:"firewallAddress"`
	// Maximum number of firewall address groups (IPv4, IPv6).
	FirewallAddrgrp int `pulumi:"firewallAddrgrp"`
	// Maximum number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
	FirewallPolicy int `pulumi:"firewallPolicy"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Maximum number of VPN IPsec phase1 tunnels.
	IpsecPhase1 int `pulumi:"ipsecPhase1"`
	// Maximum number of VPN IPsec phase1 interface tunnels.
	IpsecPhase1Interface int `pulumi:"ipsecPhase1Interface"`
	// Maximum number of VPN IPsec phase2 tunnels.
	IpsecPhase2 int `pulumi:"ipsecPhase2"`
	// Maximum number of VPN IPsec phase2 interface tunnels.
	IpsecPhase2Interface int `pulumi:"ipsecPhase2Interface"`
	// Log disk quota in MB.
	LogDiskQuota int `pulumi:"logDiskQuota"`
	// Maximum number of firewall one-time schedules.
	OnetimeSchedule int `pulumi:"onetimeSchedule"`
	// Maximum number of concurrent proxy users.
	Proxy int `pulumi:"proxy"`
	// Maximum number of firewall recurring schedules.
	RecurringSchedule int `pulumi:"recurringSchedule"`
	// Maximum number of firewall service groups.
	ServiceGroup int `pulumi:"serviceGroup"`
	// Maximum number of sessions.
	Session int `pulumi:"session"`
	// Maximum number of SSL-VPN.
	Sslvpn int `pulumi:"sslvpn"`
	// Maximum number of local users.
	User int `pulumi:"user"`
	// Maximum number of user groups.
	UserGroup int     `pulumi:"userGroup"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemResourcelimitsOutput(ctx *pulumi.Context, args LookupSystemResourcelimitsOutputArgs, opts ...pulumi.InvokeOption) LookupSystemResourcelimitsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemResourcelimitsResult, error) {
			args := v.(LookupSystemResourcelimitsArgs)
			r, err := LookupSystemResourcelimits(ctx, &args, opts...)
			var s LookupSystemResourcelimitsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemResourcelimitsResultOutput)
}

// A collection of arguments for invoking getSystemResourcelimits.
type LookupSystemResourcelimitsOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemResourcelimitsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemResourcelimitsArgs)(nil)).Elem()
}

// A collection of values returned by getSystemResourcelimits.
type LookupSystemResourcelimitsResultOutput struct{ *pulumi.OutputState }

func (LookupSystemResourcelimitsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemResourcelimitsResult)(nil)).Elem()
}

func (o LookupSystemResourcelimitsResultOutput) ToLookupSystemResourcelimitsResultOutput() LookupSystemResourcelimitsResultOutput {
	return o
}

func (o LookupSystemResourcelimitsResultOutput) ToLookupSystemResourcelimitsResultOutputWithContext(ctx context.Context) LookupSystemResourcelimitsResultOutput {
	return o
}

// Maximum number of firewall custom services.
func (o LookupSystemResourcelimitsResultOutput) CustomService() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.CustomService }).(pulumi.IntOutput)
}

// Maximum number of dial-up tunnels.
func (o LookupSystemResourcelimitsResultOutput) DialupTunnel() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.DialupTunnel }).(pulumi.IntOutput)
}

// Maximum number of firewall addresses (IPv4, IPv6, multicast).
func (o LookupSystemResourcelimitsResultOutput) FirewallAddress() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.FirewallAddress }).(pulumi.IntOutput)
}

// Maximum number of firewall address groups (IPv4, IPv6).
func (o LookupSystemResourcelimitsResultOutput) FirewallAddrgrp() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.FirewallAddrgrp }).(pulumi.IntOutput)
}

// Maximum number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
func (o LookupSystemResourcelimitsResultOutput) FirewallPolicy() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.FirewallPolicy }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemResourcelimitsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Maximum number of VPN IPsec phase1 tunnels.
func (o LookupSystemResourcelimitsResultOutput) IpsecPhase1() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.IpsecPhase1 }).(pulumi.IntOutput)
}

// Maximum number of VPN IPsec phase1 interface tunnels.
func (o LookupSystemResourcelimitsResultOutput) IpsecPhase1Interface() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.IpsecPhase1Interface }).(pulumi.IntOutput)
}

// Maximum number of VPN IPsec phase2 tunnels.
func (o LookupSystemResourcelimitsResultOutput) IpsecPhase2() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.IpsecPhase2 }).(pulumi.IntOutput)
}

// Maximum number of VPN IPsec phase2 interface tunnels.
func (o LookupSystemResourcelimitsResultOutput) IpsecPhase2Interface() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.IpsecPhase2Interface }).(pulumi.IntOutput)
}

// Log disk quota in MB.
func (o LookupSystemResourcelimitsResultOutput) LogDiskQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.LogDiskQuota }).(pulumi.IntOutput)
}

// Maximum number of firewall one-time schedules.
func (o LookupSystemResourcelimitsResultOutput) OnetimeSchedule() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.OnetimeSchedule }).(pulumi.IntOutput)
}

// Maximum number of concurrent proxy users.
func (o LookupSystemResourcelimitsResultOutput) Proxy() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.Proxy }).(pulumi.IntOutput)
}

// Maximum number of firewall recurring schedules.
func (o LookupSystemResourcelimitsResultOutput) RecurringSchedule() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.RecurringSchedule }).(pulumi.IntOutput)
}

// Maximum number of firewall service groups.
func (o LookupSystemResourcelimitsResultOutput) ServiceGroup() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.ServiceGroup }).(pulumi.IntOutput)
}

// Maximum number of sessions.
func (o LookupSystemResourcelimitsResultOutput) Session() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.Session }).(pulumi.IntOutput)
}

// Maximum number of SSL-VPN.
func (o LookupSystemResourcelimitsResultOutput) Sslvpn() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.Sslvpn }).(pulumi.IntOutput)
}

// Maximum number of local users.
func (o LookupSystemResourcelimitsResultOutput) User() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.User }).(pulumi.IntOutput)
}

// Maximum number of user groups.
func (o LookupSystemResourcelimitsResultOutput) UserGroup() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) int { return v.UserGroup }).(pulumi.IntOutput)
}

func (o LookupSystemResourcelimitsResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemResourcelimitsResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemResourcelimitsResultOutput{})
}
