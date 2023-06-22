// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class DnsfilterProfileDnsTranslation
    {
        /// <summary>
        /// DNS translation type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        /// </summary>
        public readonly string? AddrType;
        /// <summary>
        /// IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
        /// </summary>
        public readonly string? Dst;
        /// <summary>
        /// IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
        /// </summary>
        public readonly string? Dst6;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
        /// </summary>
        public readonly string? Netmask;
        /// <summary>
        /// If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).
        /// </summary>
        public readonly int? Prefix;
        /// <summary>
        /// IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
        /// </summary>
        public readonly string? Src;
        /// <summary>
        /// IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
        /// </summary>
        public readonly string? Src6;
        /// <summary>
        /// Enable/disable this DNS translation entry. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private DnsfilterProfileDnsTranslation(
            string? addrType,

            string? dst,

            string? dst6,

            int? id,

            string? netmask,

            int? prefix,

            string? src,

            string? src6,

            string? status)
        {
            AddrType = addrType;
            Dst = dst;
            Dst6 = dst6;
            Id = id;
            Netmask = netmask;
            Prefix = prefix;
            Src = src;
            Src6 = src6;
            Status = status;
        }
    }
}
