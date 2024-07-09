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
    public sealed class InterfaceIpv6
    {
        public readonly string? Autoconf;
        public readonly int? CliConn6Status;
        public readonly string? Dhcp6ClientOptions;
        public readonly ImmutableArray<Outputs.InterfaceIpv6Dhcp6IapdList> Dhcp6IapdLists;
        public readonly string? Dhcp6InformationRequest;
        public readonly string? Dhcp6PrefixDelegation;
        public readonly string? Dhcp6PrefixHint;
        public readonly int? Dhcp6PrefixHintPlt;
        public readonly int? Dhcp6PrefixHintVlt;
        public readonly string? Dhcp6RelayInterfaceId;
        public readonly string? Dhcp6RelayIp;
        public readonly string? Dhcp6RelayService;
        public readonly string? Dhcp6RelaySourceInterface;
        public readonly string? Dhcp6RelaySourceIp;
        public readonly string? Dhcp6RelayType;
        public readonly string? Icmp6SendRedirect;
        public readonly string? InterfaceIdentifier;
        public readonly string? Ip6Address;
        public readonly string? Ip6Allowaccess;
        public readonly int? Ip6DefaultLife;
        public readonly int? Ip6DelegatedPrefixIaid;
        public readonly ImmutableArray<Outputs.InterfaceIpv6Ip6DelegatedPrefixList> Ip6DelegatedPrefixLists;
        public readonly string? Ip6DnsServerOverride;
        public readonly ImmutableArray<Outputs.InterfaceIpv6Ip6ExtraAddr> Ip6ExtraAddrs;
        public readonly int? Ip6HopLimit;
        public readonly int? Ip6LinkMtu;
        public readonly string? Ip6ManageFlag;
        public readonly int? Ip6MaxInterval;
        public readonly int? Ip6MinInterval;
        public readonly string? Ip6Mode;
        public readonly string? Ip6OtherFlag;
        public readonly ImmutableArray<Outputs.InterfaceIpv6Ip6PrefixList> Ip6PrefixLists;
        public readonly string? Ip6PrefixMode;
        public readonly int? Ip6ReachableTime;
        public readonly int? Ip6RetransTime;
        public readonly string? Ip6SendAdv;
        public readonly string? Ip6Subnet;
        public readonly string? Ip6UpstreamInterface;
        public readonly string? NdCert;
        public readonly string? NdCgaModifier;
        public readonly string? NdMode;
        public readonly int? NdSecurityLevel;
        public readonly int? NdTimestampDelta;
        public readonly int? NdTimestampFuzz;
        public readonly string? RaSendMtu;
        public readonly string? UniqueAutoconfAddr;
        public readonly string? Vrip6LinkLocal;
        public readonly ImmutableArray<Outputs.InterfaceIpv6Vrrp6> Vrrp6s;
        public readonly string? VrrpVirtualMac6;

        [OutputConstructor]
        private InterfaceIpv6(
            string? autoconf,

            int? cliConn6Status,

            string? dhcp6ClientOptions,

            ImmutableArray<Outputs.InterfaceIpv6Dhcp6IapdList> dhcp6IapdLists,

            string? dhcp6InformationRequest,

            string? dhcp6PrefixDelegation,

            string? dhcp6PrefixHint,

            int? dhcp6PrefixHintPlt,

            int? dhcp6PrefixHintVlt,

            string? dhcp6RelayInterfaceId,

            string? dhcp6RelayIp,

            string? dhcp6RelayService,

            string? dhcp6RelaySourceInterface,

            string? dhcp6RelaySourceIp,

            string? dhcp6RelayType,

            string? icmp6SendRedirect,

            string? interfaceIdentifier,

            string? ip6Address,

            string? ip6Allowaccess,

            int? ip6DefaultLife,

            int? ip6DelegatedPrefixIaid,

            ImmutableArray<Outputs.InterfaceIpv6Ip6DelegatedPrefixList> ip6DelegatedPrefixLists,

            string? ip6DnsServerOverride,

            ImmutableArray<Outputs.InterfaceIpv6Ip6ExtraAddr> ip6ExtraAddrs,

            int? ip6HopLimit,

            int? ip6LinkMtu,

            string? ip6ManageFlag,

            int? ip6MaxInterval,

            int? ip6MinInterval,

            string? ip6Mode,

            string? ip6OtherFlag,

            ImmutableArray<Outputs.InterfaceIpv6Ip6PrefixList> ip6PrefixLists,

            string? ip6PrefixMode,

            int? ip6ReachableTime,

            int? ip6RetransTime,

            string? ip6SendAdv,

            string? ip6Subnet,

            string? ip6UpstreamInterface,

            string? ndCert,

            string? ndCgaModifier,

            string? ndMode,

            int? ndSecurityLevel,

            int? ndTimestampDelta,

            int? ndTimestampFuzz,

            string? raSendMtu,

            string? uniqueAutoconfAddr,

            string? vrip6LinkLocal,

            ImmutableArray<Outputs.InterfaceIpv6Vrrp6> vrrp6s,

            string? vrrpVirtualMac6)
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
