// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ssl.Web.Outputs
{

    [OutputType]
    public sealed class PortalSplitDn
    {
        /// <summary>
        /// DNS server 1.
        /// </summary>
        public readonly string? DnsServer1;
        /// <summary>
        /// DNS server 2.
        /// </summary>
        public readonly string? DnsServer2;
        /// <summary>
        /// Split DNS domains used for SSL-VPN clients separated by comma(,).
        /// </summary>
        public readonly string? Domains;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// IPv6 DNS server 1.
        /// </summary>
        public readonly string? Ipv6DnsServer1;
        /// <summary>
        /// IPv6 DNS server 2.
        /// </summary>
        public readonly string? Ipv6DnsServer2;

        [OutputConstructor]
        private PortalSplitDn(
            string? dnsServer1,

            string? dnsServer2,

            string? domains,

            int? id,

            string? ipv6DnsServer1,

            string? ipv6DnsServer2)
        {
            DnsServer1 = dnsServer1;
            DnsServer2 = dnsServer2;
            Domains = domains;
            Id = id;
            Ipv6DnsServer1 = ipv6DnsServer1;
            Ipv6DnsServer2 = ipv6DnsServer2;
        }
    }
}
