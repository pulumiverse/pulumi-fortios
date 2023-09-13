// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Systemsnmp.Outputs
{

    [OutputType]
    public sealed class CommunityHosts6
    {
        /// <summary>
        /// Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? HaDirect;
        /// <summary>
        /// Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. Valid values: `any`, `query`, `trap`.
        /// 
        /// The `hosts6` block supports:
        /// </summary>
        public readonly string? HostType;
        /// <summary>
        /// Host entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// SNMP manager IPv6 address prefix.
        /// </summary>
        public readonly string? Ipv6;
        /// <summary>
        /// Source IPv6 address for SNMP traps.
        /// </summary>
        public readonly string? SourceIpv6;

        [OutputConstructor]
        private CommunityHosts6(
            string? haDirect,

            string? hostType,

            int? id,

            string? ipv6,

            string? sourceIpv6)
        {
            HaDirect = haDirect;
            HostType = hostType;
            Id = id;
            Ipv6 = ipv6;
            SourceIpv6 = sourceIpv6;
        }
    }
}
