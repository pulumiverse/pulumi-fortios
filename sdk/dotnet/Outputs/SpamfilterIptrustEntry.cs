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
    public sealed class SpamfilterIptrustEntry
    {
        /// <summary>
        /// Type of address. Valid values: `ipv4`, `ipv6`.
        /// </summary>
        public readonly string? AddrType;
        /// <summary>
        /// Trusted IP entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// IPv4 network address or network address/subnet mask bits.
        /// </summary>
        public readonly string? Ip4Subnet;
        /// <summary>
        /// IPv6 network address/subnet mask bits.
        /// </summary>
        public readonly string? Ip6Subnet;
        /// <summary>
        /// Enable/disable status. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private SpamfilterIptrustEntry(
            string? addrType,

            int? id,

            string? ip4Subnet,

            string? ip6Subnet,

            string? status)
        {
            AddrType = addrType;
            Id = id;
            Ip4Subnet = ip4Subnet;
            Ip6Subnet = ip6Subnet;
            Status = status;
        }
    }
}
