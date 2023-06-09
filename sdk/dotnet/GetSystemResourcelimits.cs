// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    public static class GetSystemResourcelimits
    {
        /// <summary>
        /// Use this data source to get information on fortios system resourcelimits
        /// </summary>
        public static Task<GetSystemResourcelimitsResult> InvokeAsync(GetSystemResourcelimitsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemResourcelimitsResult>("fortios:index/getSystemResourcelimits:getSystemResourcelimits", args ?? new GetSystemResourcelimitsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system resourcelimits
        /// </summary>
        public static Output<GetSystemResourcelimitsResult> Invoke(GetSystemResourcelimitsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemResourcelimitsResult>("fortios:index/getSystemResourcelimits:getSystemResourcelimits", args ?? new GetSystemResourcelimitsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemResourcelimitsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemResourcelimitsArgs()
        {
        }
        public static new GetSystemResourcelimitsArgs Empty => new GetSystemResourcelimitsArgs();
    }

    public sealed class GetSystemResourcelimitsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemResourcelimitsInvokeArgs()
        {
        }
        public static new GetSystemResourcelimitsInvokeArgs Empty => new GetSystemResourcelimitsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemResourcelimitsResult
    {
        /// <summary>
        /// Maximum number of firewall custom services.
        /// </summary>
        public readonly int CustomService;
        /// <summary>
        /// Maximum number of dial-up tunnels.
        /// </summary>
        public readonly int DialupTunnel;
        /// <summary>
        /// Maximum number of firewall addresses (IPv4, IPv6, multicast).
        /// </summary>
        public readonly int FirewallAddress;
        /// <summary>
        /// Maximum number of firewall address groups (IPv4, IPv6).
        /// </summary>
        public readonly int FirewallAddrgrp;
        /// <summary>
        /// Maximum number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
        /// </summary>
        public readonly int FirewallPolicy;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Maximum number of VPN IPsec phase1 tunnels.
        /// </summary>
        public readonly int IpsecPhase1;
        /// <summary>
        /// Maximum number of VPN IPsec phase1 interface tunnels.
        /// </summary>
        public readonly int IpsecPhase1Interface;
        /// <summary>
        /// Maximum number of VPN IPsec phase2 tunnels.
        /// </summary>
        public readonly int IpsecPhase2;
        /// <summary>
        /// Maximum number of VPN IPsec phase2 interface tunnels.
        /// </summary>
        public readonly int IpsecPhase2Interface;
        /// <summary>
        /// Log disk quota in MB.
        /// </summary>
        public readonly int LogDiskQuota;
        /// <summary>
        /// Maximum number of firewall one-time schedules.
        /// </summary>
        public readonly int OnetimeSchedule;
        /// <summary>
        /// Maximum number of concurrent proxy users.
        /// </summary>
        public readonly int Proxy;
        /// <summary>
        /// Maximum number of firewall recurring schedules.
        /// </summary>
        public readonly int RecurringSchedule;
        /// <summary>
        /// Maximum number of firewall service groups.
        /// </summary>
        public readonly int ServiceGroup;
        /// <summary>
        /// Maximum number of sessions.
        /// </summary>
        public readonly int Session;
        /// <summary>
        /// Maximum number of SSL-VPN.
        /// </summary>
        public readonly int Sslvpn;
        /// <summary>
        /// Maximum number of local users.
        /// </summary>
        public readonly int User;
        /// <summary>
        /// Maximum number of user groups.
        /// </summary>
        public readonly int UserGroup;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemResourcelimitsResult(
            int customService,

            int dialupTunnel,

            int firewallAddress,

            int firewallAddrgrp,

            int firewallPolicy,

            string id,

            int ipsecPhase1,

            int ipsecPhase1Interface,

            int ipsecPhase2,

            int ipsecPhase2Interface,

            int logDiskQuota,

            int onetimeSchedule,

            int proxy,

            int recurringSchedule,

            int serviceGroup,

            int session,

            int sslvpn,

            int user,

            int userGroup,

            string? vdomparam)
        {
            CustomService = customService;
            DialupTunnel = dialupTunnel;
            FirewallAddress = firewallAddress;
            FirewallAddrgrp = firewallAddrgrp;
            FirewallPolicy = firewallPolicy;
            Id = id;
            IpsecPhase1 = ipsecPhase1;
            IpsecPhase1Interface = ipsecPhase1Interface;
            IpsecPhase2 = ipsecPhase2;
            IpsecPhase2Interface = ipsecPhase2Interface;
            LogDiskQuota = logDiskQuota;
            OnetimeSchedule = onetimeSchedule;
            Proxy = proxy;
            RecurringSchedule = recurringSchedule;
            ServiceGroup = serviceGroup;
            Session = session;
            Sslvpn = sslvpn;
            User = user;
            UserGroup = userGroup;
            Vdomparam = vdomparam;
        }
    }
}
