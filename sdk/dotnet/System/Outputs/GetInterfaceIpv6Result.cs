// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class GetInterfaceIpv6Result
    {
        /// <summary>
        /// Enable/disable address auto config.
        /// </summary>
        public readonly string Autoconf;
        /// <summary>
        /// CLI IPv6 connection status.
        /// </summary>
        public readonly int CliConn6Status;
        /// <summary>
        /// DHCPv6 client options.
        /// </summary>
        public readonly string Dhcp6ClientOptions;
        /// <summary>
        /// DHCPv6 IA-PD list The structure of `dhcp6_iapd_list` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInterfaceIpv6Dhcp6IapdListResult> Dhcp6IapdLists;
        /// <summary>
        /// Enable/disable DHCPv6 information request.
        /// </summary>
        public readonly string Dhcp6InformationRequest;
        /// <summary>
        /// Enable/disable DHCPv6 prefix delegation.
        /// </summary>
        public readonly string Dhcp6PrefixDelegation;
        /// <summary>
        /// DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
        /// </summary>
        public readonly string Dhcp6PrefixHint;
        /// <summary>
        /// DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
        /// </summary>
        public readonly int Dhcp6PrefixHintPlt;
        /// <summary>
        /// DHCPv6 prefix hint valid life time (sec).
        /// </summary>
        public readonly int Dhcp6PrefixHintVlt;
        /// <summary>
        /// DHCP6 relay interface ID.
        /// </summary>
        public readonly string Dhcp6RelayInterfaceId;
        /// <summary>
        /// DHCPv6 relay IP address.
        /// </summary>
        public readonly string Dhcp6RelayIp;
        /// <summary>
        /// Enable/disable DHCPv6 relay.
        /// </summary>
        public readonly string Dhcp6RelayService;
        /// <summary>
        /// Enable/disable use of address on this interface as the source address of the relay message.
        /// </summary>
        public readonly string Dhcp6RelaySourceInterface;
        /// <summary>
        /// IPv6 address used by the DHCP6 relay as its source IP.
        /// </summary>
        public readonly string Dhcp6RelaySourceIp;
        /// <summary>
        /// DHCPv6 relay type.
        /// </summary>
        public readonly string Dhcp6RelayType;
        /// <summary>
        /// Enable/disable sending of ICMPv6 redirects.
        /// </summary>
        public readonly string Icmp6SendRedirect;
        /// <summary>
        /// IPv6 interface identifier.
        /// </summary>
        public readonly string InterfaceIdentifier;
        /// <summary>
        /// Primary IPv6 address prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx
        /// </summary>
        public readonly string Ip6Address;
        /// <summary>
        /// Allow management access to the interface.
        /// </summary>
        public readonly string Ip6Allowaccess;
        /// <summary>
        /// Default life (sec).
        /// </summary>
        public readonly int Ip6DefaultLife;
        /// <summary>
        /// IAID of obtained delegated-prefix from the upstream interface.
        /// </summary>
        public readonly int Ip6DelegatedPrefixIaid;
        /// <summary>
        /// Advertised IPv6 delegated prefix list. The structure of `ip6_delegated_prefix_list` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInterfaceIpv6Ip6DelegatedPrefixListResult> Ip6DelegatedPrefixLists;
        /// <summary>
        /// Enable/disable using the DNS server acquired by DHCP.
        /// </summary>
        public readonly string Ip6DnsServerOverride;
        /// <summary>
        /// Extra IPv6 address prefixes of interface. The structure of `ip6_extra_addr` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInterfaceIpv6Ip6ExtraAddrResult> Ip6ExtraAddrs;
        /// <summary>
        /// Hop limit (0 means unspecified).
        /// </summary>
        public readonly int Ip6HopLimit;
        /// <summary>
        /// IPv6 link MTU.
        /// </summary>
        public readonly int Ip6LinkMtu;
        /// <summary>
        /// Enable/disable the managed flag.
        /// </summary>
        public readonly string Ip6ManageFlag;
        /// <summary>
        /// IPv6 maximum interval (4 to 1800 sec).
        /// </summary>
        public readonly int Ip6MaxInterval;
        /// <summary>
        /// IPv6 minimum interval (3 to 1350 sec).
        /// </summary>
        public readonly int Ip6MinInterval;
        /// <summary>
        /// Addressing mode (static, DHCP, delegated).
        /// </summary>
        public readonly string Ip6Mode;
        /// <summary>
        /// Enable/disable the other IPv6 flag.
        /// </summary>
        public readonly string Ip6OtherFlag;
        /// <summary>
        /// Advertised prefix list. The structure of `ip6_prefix_list` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInterfaceIpv6Ip6PrefixListResult> Ip6PrefixLists;
        /// <summary>
        /// Assigning a prefix from DHCP or RA.
        /// </summary>
        public readonly string Ip6PrefixMode;
        /// <summary>
        /// IPv6 reachable time (milliseconds; 0 means unspecified).
        /// </summary>
        public readonly int Ip6ReachableTime;
        /// <summary>
        /// IPv6 retransmit time (milliseconds; 0 means unspecified).
        /// </summary>
        public readonly int Ip6RetransTime;
        /// <summary>
        /// Enable/disable sending advertisements about the interface.
        /// </summary>
        public readonly string Ip6SendAdv;
        /// <summary>
        /// Subnet to routing prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx
        /// </summary>
        public readonly string Ip6Subnet;
        /// <summary>
        /// Interface name providing delegated information.
        /// </summary>
        public readonly string Ip6UpstreamInterface;
        /// <summary>
        /// Neighbor discovery certificate.
        /// </summary>
        public readonly string NdCert;
        /// <summary>
        /// Neighbor discovery CGA modifier.
        /// </summary>
        public readonly string NdCgaModifier;
        /// <summary>
        /// Neighbor discovery mode.
        /// </summary>
        public readonly string NdMode;
        /// <summary>
        /// Neighbor discovery security level (0 - 7; 0 = least secure, default = 0).
        /// </summary>
        public readonly int NdSecurityLevel;
        /// <summary>
        /// Neighbor discovery timestamp delta value (1 - 3600 sec; default = 300).
        /// </summary>
        public readonly int NdTimestampDelta;
        /// <summary>
        /// Neighbor discovery timestamp fuzz factor (1 - 60 sec; default = 1).
        /// </summary>
        public readonly int NdTimestampFuzz;
        /// <summary>
        /// Enable/disable sending link MTU in RA packet.
        /// </summary>
        public readonly string RaSendMtu;
        /// <summary>
        /// Enable/disable unique auto config address.
        /// </summary>
        public readonly string UniqueAutoconfAddr;
        /// <summary>
        /// Link-local IPv6 address of virtual router.
        /// </summary>
        public readonly string Vrip6LinkLocal;
        /// <summary>
        /// IPv6 VRRP configuration. The structure of `vrrp6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInterfaceIpv6Vrrp6Result> Vrrp6s;
        /// <summary>
        /// Enable/disable virtual MAC for VRRP.
        /// </summary>
        public readonly string VrrpVirtualMac6;

        [OutputConstructor]
        private GetInterfaceIpv6Result(
            string autoconf,

            int cliConn6Status,

            string dhcp6ClientOptions,

            ImmutableArray<Outputs.GetInterfaceIpv6Dhcp6IapdListResult> dhcp6IapdLists,

            string dhcp6InformationRequest,

            string dhcp6PrefixDelegation,

            string dhcp6PrefixHint,

            int dhcp6PrefixHintPlt,

            int dhcp6PrefixHintVlt,

            string dhcp6RelayInterfaceId,

            string dhcp6RelayIp,

            string dhcp6RelayService,

            string dhcp6RelaySourceInterface,

            string dhcp6RelaySourceIp,

            string dhcp6RelayType,

            string icmp6SendRedirect,

            string interfaceIdentifier,

            string ip6Address,

            string ip6Allowaccess,

            int ip6DefaultLife,

            int ip6DelegatedPrefixIaid,

            ImmutableArray<Outputs.GetInterfaceIpv6Ip6DelegatedPrefixListResult> ip6DelegatedPrefixLists,

            string ip6DnsServerOverride,

            ImmutableArray<Outputs.GetInterfaceIpv6Ip6ExtraAddrResult> ip6ExtraAddrs,

            int ip6HopLimit,

            int ip6LinkMtu,

            string ip6ManageFlag,

            int ip6MaxInterval,

            int ip6MinInterval,

            string ip6Mode,

            string ip6OtherFlag,

            ImmutableArray<Outputs.GetInterfaceIpv6Ip6PrefixListResult> ip6PrefixLists,

            string ip6PrefixMode,

            int ip6ReachableTime,

            int ip6RetransTime,

            string ip6SendAdv,

            string ip6Subnet,

            string ip6UpstreamInterface,

            string ndCert,

            string ndCgaModifier,

            string ndMode,

            int ndSecurityLevel,

            int ndTimestampDelta,

            int ndTimestampFuzz,

            string raSendMtu,

            string uniqueAutoconfAddr,

            string vrip6LinkLocal,

            ImmutableArray<Outputs.GetInterfaceIpv6Vrrp6Result> vrrp6s,

            string vrrpVirtualMac6)
        {
            Autoconf = autoconf;
            CliConn6Status = cliConn6Status;
            Dhcp6ClientOptions = dhcp6ClientOptions;
            Dhcp6IapdLists = dhcp6IapdLists;
            Dhcp6InformationRequest = dhcp6InformationRequest;
            Dhcp6PrefixDelegation = dhcp6PrefixDelegation;
            Dhcp6PrefixHint = dhcp6PrefixHint;
            Dhcp6PrefixHintPlt = dhcp6PrefixHintPlt;
            Dhcp6PrefixHintVlt = dhcp6PrefixHintVlt;
            Dhcp6RelayInterfaceId = dhcp6RelayInterfaceId;
            Dhcp6RelayIp = dhcp6RelayIp;
            Dhcp6RelayService = dhcp6RelayService;
            Dhcp6RelaySourceInterface = dhcp6RelaySourceInterface;
            Dhcp6RelaySourceIp = dhcp6RelaySourceIp;
            Dhcp6RelayType = dhcp6RelayType;
            Icmp6SendRedirect = icmp6SendRedirect;
            InterfaceIdentifier = interfaceIdentifier;
            Ip6Address = ip6Address;
            Ip6Allowaccess = ip6Allowaccess;
            Ip6DefaultLife = ip6DefaultLife;
            Ip6DelegatedPrefixIaid = ip6DelegatedPrefixIaid;
            Ip6DelegatedPrefixLists = ip6DelegatedPrefixLists;
            Ip6DnsServerOverride = ip6DnsServerOverride;
            Ip6ExtraAddrs = ip6ExtraAddrs;
            Ip6HopLimit = ip6HopLimit;
            Ip6LinkMtu = ip6LinkMtu;
            Ip6ManageFlag = ip6ManageFlag;
            Ip6MaxInterval = ip6MaxInterval;
            Ip6MinInterval = ip6MinInterval;
            Ip6Mode = ip6Mode;
            Ip6OtherFlag = ip6OtherFlag;
            Ip6PrefixLists = ip6PrefixLists;
            Ip6PrefixMode = ip6PrefixMode;
            Ip6ReachableTime = ip6ReachableTime;
            Ip6RetransTime = ip6RetransTime;
            Ip6SendAdv = ip6SendAdv;
            Ip6Subnet = ip6Subnet;
            Ip6UpstreamInterface = ip6UpstreamInterface;
            NdCert = ndCert;
            NdCgaModifier = ndCgaModifier;
            NdMode = ndMode;
            NdSecurityLevel = ndSecurityLevel;
            NdTimestampDelta = ndTimestampDelta;
            NdTimestampFuzz = ndTimestampFuzz;
            RaSendMtu = raSendMtu;
            UniqueAutoconfAddr = uniqueAutoconfAddr;
            Vrip6LinkLocal = vrip6LinkLocal;
            Vrrp6s = vrrp6s;
            VrrpVirtualMac6 = vrrpVirtualMac6;
        }
    }
}
