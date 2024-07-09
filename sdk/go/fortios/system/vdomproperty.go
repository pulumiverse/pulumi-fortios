// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure VDOM property.
//
// ## Import
//
// System VdomProperty can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/vdomproperty:Vdomproperty labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/vdomproperty:Vdomproperty labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Vdomproperty struct {
	pulumi.CustomResourceState

	// Maximum guaranteed number of firewall custom services.
	CustomService pulumi.StringOutput `pulumi:"customService"`
	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Maximum guaranteed number of dial-up tunnels.
	DialupTunnel pulumi.StringOutput `pulumi:"dialupTunnel"`
	// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
	FirewallAddress pulumi.StringOutput `pulumi:"firewallAddress"`
	// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
	FirewallAddrgrp pulumi.StringOutput `pulumi:"firewallAddrgrp"`
	// Maximum guaranteed number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).
	FirewallPolicy pulumi.StringOutput `pulumi:"firewallPolicy"`
	// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
	IpsecPhase1 pulumi.StringOutput `pulumi:"ipsecPhase1"`
	// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
	IpsecPhase1Interface pulumi.StringOutput `pulumi:"ipsecPhase1Interface"`
	// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
	IpsecPhase2 pulumi.StringOutput `pulumi:"ipsecPhase2"`
	// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
	IpsecPhase2Interface pulumi.StringOutput `pulumi:"ipsecPhase2Interface"`
	// Log disk quota in megabytes (MB). Range depends on how much disk space is available.
	LogDiskQuota pulumi.StringOutput `pulumi:"logDiskQuota"`
	// VDOM name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Maximum guaranteed number of firewall one-time schedules.
	OnetimeSchedule pulumi.StringOutput `pulumi:"onetimeSchedule"`
	// Maximum guaranteed number of concurrent proxy users.
	Proxy pulumi.StringOutput `pulumi:"proxy"`
	// Maximum guaranteed number of firewall recurring schedules.
	RecurringSchedule pulumi.StringOutput `pulumi:"recurringSchedule"`
	// Maximum guaranteed number of firewall service groups.
	ServiceGroup pulumi.StringOutput `pulumi:"serviceGroup"`
	// Maximum guaranteed number of sessions.
	Session pulumi.StringOutput `pulumi:"session"`
	// Permanent SNMP Index of the virtual domain. On FortiOS versions 6.2.0-6.2.6: 0 - 4294967295. On FortiOS versions >= 6.4.0: 1 - 2147483647.
	SnmpIndex pulumi.IntOutput `pulumi:"snmpIndex"`
	// Maximum guaranteed number of SSL-VPNs.
	Sslvpn pulumi.StringOutput `pulumi:"sslvpn"`
	// Maximum guaranteed number of local users.
	User pulumi.StringOutput `pulumi:"user"`
	// Maximum guaranteed number of user groups.
	UserGroup pulumi.StringOutput `pulumi:"userGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringOutput `pulumi:"vdomparam"`
}

// NewVdomproperty registers a new resource with the given unique name, arguments, and options.
func NewVdomproperty(ctx *pulumi.Context,
	name string, args *VdompropertyArgs, opts ...pulumi.ResourceOption) (*Vdomproperty, error) {
	if args == nil {
		args = &VdompropertyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vdomproperty
	err := ctx.RegisterResource("fortios:system/vdomproperty:Vdomproperty", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVdomproperty gets an existing Vdomproperty resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVdomproperty(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VdompropertyState, opts ...pulumi.ResourceOption) (*Vdomproperty, error) {
	var resource Vdomproperty
	err := ctx.ReadResource("fortios:system/vdomproperty:Vdomproperty", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vdomproperty resources.
type vdompropertyState struct {
	// Maximum guaranteed number of firewall custom services.
	CustomService *string `pulumi:"customService"`
	// Description.
	Description *string `pulumi:"description"`
	// Maximum guaranteed number of dial-up tunnels.
	DialupTunnel *string `pulumi:"dialupTunnel"`
	// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
	FirewallAddress *string `pulumi:"firewallAddress"`
	// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
	FirewallAddrgrp *string `pulumi:"firewallAddrgrp"`
	// Maximum guaranteed number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).
	FirewallPolicy *string `pulumi:"firewallPolicy"`
	// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
	IpsecPhase1 *string `pulumi:"ipsecPhase1"`
	// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
	IpsecPhase1Interface *string `pulumi:"ipsecPhase1Interface"`
	// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
	IpsecPhase2 *string `pulumi:"ipsecPhase2"`
	// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
	IpsecPhase2Interface *string `pulumi:"ipsecPhase2Interface"`
	// Log disk quota in megabytes (MB). Range depends on how much disk space is available.
	LogDiskQuota *string `pulumi:"logDiskQuota"`
	// VDOM name.
	Name *string `pulumi:"name"`
	// Maximum guaranteed number of firewall one-time schedules.
	OnetimeSchedule *string `pulumi:"onetimeSchedule"`
	// Maximum guaranteed number of concurrent proxy users.
	Proxy *string `pulumi:"proxy"`
	// Maximum guaranteed number of firewall recurring schedules.
	RecurringSchedule *string `pulumi:"recurringSchedule"`
	// Maximum guaranteed number of firewall service groups.
	ServiceGroup *string `pulumi:"serviceGroup"`
	// Maximum guaranteed number of sessions.
	Session *string `pulumi:"session"`
	// Permanent SNMP Index of the virtual domain. On FortiOS versions 6.2.0-6.2.6: 0 - 4294967295. On FortiOS versions >= 6.4.0: 1 - 2147483647.
	SnmpIndex *int `pulumi:"snmpIndex"`
	// Maximum guaranteed number of SSL-VPNs.
	Sslvpn *string `pulumi:"sslvpn"`
	// Maximum guaranteed number of local users.
	User *string `pulumi:"user"`
	// Maximum guaranteed number of user groups.
	UserGroup *string `pulumi:"userGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type VdompropertyState struct {
	// Maximum guaranteed number of firewall custom services.
	CustomService pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Maximum guaranteed number of dial-up tunnels.
	DialupTunnel pulumi.StringPtrInput
	// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
	FirewallAddress pulumi.StringPtrInput
	// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
	FirewallAddrgrp pulumi.StringPtrInput
	// Maximum guaranteed number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).
	FirewallPolicy pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
	IpsecPhase1 pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
	IpsecPhase1Interface pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
	IpsecPhase2 pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
	IpsecPhase2Interface pulumi.StringPtrInput
	// Log disk quota in megabytes (MB). Range depends on how much disk space is available.
	LogDiskQuota pulumi.StringPtrInput
	// VDOM name.
	Name pulumi.StringPtrInput
	// Maximum guaranteed number of firewall one-time schedules.
	OnetimeSchedule pulumi.StringPtrInput
	// Maximum guaranteed number of concurrent proxy users.
	Proxy pulumi.StringPtrInput
	// Maximum guaranteed number of firewall recurring schedules.
	RecurringSchedule pulumi.StringPtrInput
	// Maximum guaranteed number of firewall service groups.
	ServiceGroup pulumi.StringPtrInput
	// Maximum guaranteed number of sessions.
	Session pulumi.StringPtrInput
	// Permanent SNMP Index of the virtual domain. On FortiOS versions 6.2.0-6.2.6: 0 - 4294967295. On FortiOS versions >= 6.4.0: 1 - 2147483647.
	SnmpIndex pulumi.IntPtrInput
	// Maximum guaranteed number of SSL-VPNs.
	Sslvpn pulumi.StringPtrInput
	// Maximum guaranteed number of local users.
	User pulumi.StringPtrInput
	// Maximum guaranteed number of user groups.
	UserGroup pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VdompropertyState) ElementType() reflect.Type {
	return reflect.TypeOf((*vdompropertyState)(nil)).Elem()
}

type vdompropertyArgs struct {
	// Maximum guaranteed number of firewall custom services.
	CustomService *string `pulumi:"customService"`
	// Description.
	Description *string `pulumi:"description"`
	// Maximum guaranteed number of dial-up tunnels.
	DialupTunnel *string `pulumi:"dialupTunnel"`
	// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
	FirewallAddress *string `pulumi:"firewallAddress"`
	// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
	FirewallAddrgrp *string `pulumi:"firewallAddrgrp"`
	// Maximum guaranteed number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).
	FirewallPolicy *string `pulumi:"firewallPolicy"`
	// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
	IpsecPhase1 *string `pulumi:"ipsecPhase1"`
	// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
	IpsecPhase1Interface *string `pulumi:"ipsecPhase1Interface"`
	// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
	IpsecPhase2 *string `pulumi:"ipsecPhase2"`
	// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
	IpsecPhase2Interface *string `pulumi:"ipsecPhase2Interface"`
	// Log disk quota in megabytes (MB). Range depends on how much disk space is available.
	LogDiskQuota *string `pulumi:"logDiskQuota"`
	// VDOM name.
	Name *string `pulumi:"name"`
	// Maximum guaranteed number of firewall one-time schedules.
	OnetimeSchedule *string `pulumi:"onetimeSchedule"`
	// Maximum guaranteed number of concurrent proxy users.
	Proxy *string `pulumi:"proxy"`
	// Maximum guaranteed number of firewall recurring schedules.
	RecurringSchedule *string `pulumi:"recurringSchedule"`
	// Maximum guaranteed number of firewall service groups.
	ServiceGroup *string `pulumi:"serviceGroup"`
	// Maximum guaranteed number of sessions.
	Session *string `pulumi:"session"`
	// Permanent SNMP Index of the virtual domain. On FortiOS versions 6.2.0-6.2.6: 0 - 4294967295. On FortiOS versions >= 6.4.0: 1 - 2147483647.
	SnmpIndex *int `pulumi:"snmpIndex"`
	// Maximum guaranteed number of SSL-VPNs.
	Sslvpn *string `pulumi:"sslvpn"`
	// Maximum guaranteed number of local users.
	User *string `pulumi:"user"`
	// Maximum guaranteed number of user groups.
	UserGroup *string `pulumi:"userGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Vdomproperty resource.
type VdompropertyArgs struct {
	// Maximum guaranteed number of firewall custom services.
	CustomService pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Maximum guaranteed number of dial-up tunnels.
	DialupTunnel pulumi.StringPtrInput
	// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
	FirewallAddress pulumi.StringPtrInput
	// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
	FirewallAddrgrp pulumi.StringPtrInput
	// Maximum guaranteed number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).
	FirewallPolicy pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
	IpsecPhase1 pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
	IpsecPhase1Interface pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
	IpsecPhase2 pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
	IpsecPhase2Interface pulumi.StringPtrInput
	// Log disk quota in megabytes (MB). Range depends on how much disk space is available.
	LogDiskQuota pulumi.StringPtrInput
	// VDOM name.
	Name pulumi.StringPtrInput
	// Maximum guaranteed number of firewall one-time schedules.
	OnetimeSchedule pulumi.StringPtrInput
	// Maximum guaranteed number of concurrent proxy users.
	Proxy pulumi.StringPtrInput
	// Maximum guaranteed number of firewall recurring schedules.
	RecurringSchedule pulumi.StringPtrInput
	// Maximum guaranteed number of firewall service groups.
	ServiceGroup pulumi.StringPtrInput
	// Maximum guaranteed number of sessions.
	Session pulumi.StringPtrInput
	// Permanent SNMP Index of the virtual domain. On FortiOS versions 6.2.0-6.2.6: 0 - 4294967295. On FortiOS versions >= 6.4.0: 1 - 2147483647.
	SnmpIndex pulumi.IntPtrInput
	// Maximum guaranteed number of SSL-VPNs.
	Sslvpn pulumi.StringPtrInput
	// Maximum guaranteed number of local users.
	User pulumi.StringPtrInput
	// Maximum guaranteed number of user groups.
	UserGroup pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VdompropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vdompropertyArgs)(nil)).Elem()
}

type VdompropertyInput interface {
	pulumi.Input

	ToVdompropertyOutput() VdompropertyOutput
	ToVdompropertyOutputWithContext(ctx context.Context) VdompropertyOutput
}

func (*Vdomproperty) ElementType() reflect.Type {
	return reflect.TypeOf((**Vdomproperty)(nil)).Elem()
}

func (i *Vdomproperty) ToVdompropertyOutput() VdompropertyOutput {
	return i.ToVdompropertyOutputWithContext(context.Background())
}

func (i *Vdomproperty) ToVdompropertyOutputWithContext(ctx context.Context) VdompropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VdompropertyOutput)
}

// VdompropertyArrayInput is an input type that accepts VdompropertyArray and VdompropertyArrayOutput values.
// You can construct a concrete instance of `VdompropertyArrayInput` via:
//
//	VdompropertyArray{ VdompropertyArgs{...} }
type VdompropertyArrayInput interface {
	pulumi.Input

	ToVdompropertyArrayOutput() VdompropertyArrayOutput
	ToVdompropertyArrayOutputWithContext(context.Context) VdompropertyArrayOutput
}

type VdompropertyArray []VdompropertyInput

func (VdompropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vdomproperty)(nil)).Elem()
}

func (i VdompropertyArray) ToVdompropertyArrayOutput() VdompropertyArrayOutput {
	return i.ToVdompropertyArrayOutputWithContext(context.Background())
}

func (i VdompropertyArray) ToVdompropertyArrayOutputWithContext(ctx context.Context) VdompropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VdompropertyArrayOutput)
}

// VdompropertyMapInput is an input type that accepts VdompropertyMap and VdompropertyMapOutput values.
// You can construct a concrete instance of `VdompropertyMapInput` via:
//
//	VdompropertyMap{ "key": VdompropertyArgs{...} }
type VdompropertyMapInput interface {
	pulumi.Input

	ToVdompropertyMapOutput() VdompropertyMapOutput
	ToVdompropertyMapOutputWithContext(context.Context) VdompropertyMapOutput
}

type VdompropertyMap map[string]VdompropertyInput

func (VdompropertyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vdomproperty)(nil)).Elem()
}

func (i VdompropertyMap) ToVdompropertyMapOutput() VdompropertyMapOutput {
	return i.ToVdompropertyMapOutputWithContext(context.Background())
}

func (i VdompropertyMap) ToVdompropertyMapOutputWithContext(ctx context.Context) VdompropertyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VdompropertyMapOutput)
}

type VdompropertyOutput struct{ *pulumi.OutputState }

func (VdompropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vdomproperty)(nil)).Elem()
}

func (o VdompropertyOutput) ToVdompropertyOutput() VdompropertyOutput {
	return o
}

func (o VdompropertyOutput) ToVdompropertyOutputWithContext(ctx context.Context) VdompropertyOutput {
	return o
}

// Maximum guaranteed number of firewall custom services.
func (o VdompropertyOutput) CustomService() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.CustomService }).(pulumi.StringOutput)
}

// Description.
func (o VdompropertyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Maximum guaranteed number of dial-up tunnels.
func (o VdompropertyOutput) DialupTunnel() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.DialupTunnel }).(pulumi.StringOutput)
}

// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
func (o VdompropertyOutput) FirewallAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.FirewallAddress }).(pulumi.StringOutput)
}

// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
func (o VdompropertyOutput) FirewallAddrgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.FirewallAddrgrp }).(pulumi.StringOutput)
}

// Maximum guaranteed number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).
func (o VdompropertyOutput) FirewallPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.FirewallPolicy }).(pulumi.StringOutput)
}

// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
func (o VdompropertyOutput) IpsecPhase1() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.IpsecPhase1 }).(pulumi.StringOutput)
}

// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
func (o VdompropertyOutput) IpsecPhase1Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.IpsecPhase1Interface }).(pulumi.StringOutput)
}

// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
func (o VdompropertyOutput) IpsecPhase2() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.IpsecPhase2 }).(pulumi.StringOutput)
}

// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
func (o VdompropertyOutput) IpsecPhase2Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.IpsecPhase2Interface }).(pulumi.StringOutput)
}

// Log disk quota in megabytes (MB). Range depends on how much disk space is available.
func (o VdompropertyOutput) LogDiskQuota() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.LogDiskQuota }).(pulumi.StringOutput)
}

// VDOM name.
func (o VdompropertyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Maximum guaranteed number of firewall one-time schedules.
func (o VdompropertyOutput) OnetimeSchedule() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.OnetimeSchedule }).(pulumi.StringOutput)
}

// Maximum guaranteed number of concurrent proxy users.
func (o VdompropertyOutput) Proxy() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.Proxy }).(pulumi.StringOutput)
}

// Maximum guaranteed number of firewall recurring schedules.
func (o VdompropertyOutput) RecurringSchedule() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.RecurringSchedule }).(pulumi.StringOutput)
}

// Maximum guaranteed number of firewall service groups.
func (o VdompropertyOutput) ServiceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.ServiceGroup }).(pulumi.StringOutput)
}

// Maximum guaranteed number of sessions.
func (o VdompropertyOutput) Session() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.Session }).(pulumi.StringOutput)
}

// Permanent SNMP Index of the virtual domain. On FortiOS versions 6.2.0-6.2.6: 0 - 4294967295. On FortiOS versions >= 6.4.0: 1 - 2147483647.
func (o VdompropertyOutput) SnmpIndex() pulumi.IntOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.IntOutput { return v.SnmpIndex }).(pulumi.IntOutput)
}

// Maximum guaranteed number of SSL-VPNs.
func (o VdompropertyOutput) Sslvpn() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.Sslvpn }).(pulumi.StringOutput)
}

// Maximum guaranteed number of local users.
func (o VdompropertyOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

// Maximum guaranteed number of user groups.
func (o VdompropertyOutput) UserGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.UserGroup }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o VdompropertyOutput) Vdomparam() pulumi.StringOutput {
	return o.ApplyT(func(v *Vdomproperty) pulumi.StringOutput { return v.Vdomparam }).(pulumi.StringOutput)
}

type VdompropertyArrayOutput struct{ *pulumi.OutputState }

func (VdompropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vdomproperty)(nil)).Elem()
}

func (o VdompropertyArrayOutput) ToVdompropertyArrayOutput() VdompropertyArrayOutput {
	return o
}

func (o VdompropertyArrayOutput) ToVdompropertyArrayOutputWithContext(ctx context.Context) VdompropertyArrayOutput {
	return o
}

func (o VdompropertyArrayOutput) Index(i pulumi.IntInput) VdompropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vdomproperty {
		return vs[0].([]*Vdomproperty)[vs[1].(int)]
	}).(VdompropertyOutput)
}

type VdompropertyMapOutput struct{ *pulumi.OutputState }

func (VdompropertyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vdomproperty)(nil)).Elem()
}

func (o VdompropertyMapOutput) ToVdompropertyMapOutput() VdompropertyMapOutput {
	return o
}

func (o VdompropertyMapOutput) ToVdompropertyMapOutputWithContext(ctx context.Context) VdompropertyMapOutput {
	return o
}

func (o VdompropertyMapOutput) MapIndex(k pulumi.StringInput) VdompropertyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vdomproperty {
		return vs[0].(map[string]*Vdomproperty)[vs[1].(string)]
	}).(VdompropertyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VdompropertyInput)(nil)).Elem(), &Vdomproperty{})
	pulumi.RegisterInputType(reflect.TypeOf((*VdompropertyArrayInput)(nil)).Elem(), VdompropertyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VdompropertyMapInput)(nil)).Elem(), VdompropertyMap{})
	pulumi.RegisterOutputType(VdompropertyOutput{})
	pulumi.RegisterOutputType(VdompropertyArrayOutput{})
	pulumi.RegisterOutputType(VdompropertyMapOutput{})
}
