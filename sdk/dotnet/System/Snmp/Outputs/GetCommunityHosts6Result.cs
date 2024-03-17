// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Snmp.Outputs
{

    [OutputType]
    public sealed class GetCommunityHosts6Result
    {
        /// <summary>
        /// Enable/disable direct management of HA cluster members.
        /// </summary>
        public readonly string HaDirect;
        /// <summary>
        /// Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both.
        /// </summary>
        public readonly string HostType;
        /// <summary>
        /// Host6 entry ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// SNMP manager IPv6 address prefix.
        /// </summary>
        public readonly string Ipv6;
        /// <summary>
        /// Source IPv6 address for SNMP traps.
        /// </summary>
        public readonly string SourceIpv6;

        [OutputConstructor]
        private GetCommunityHosts6Result(
            string haDirect,

            string hostType,

            int id,

            string ipv6,

            string sourceIpv6)
        {
            HaDirect = haDirect;
            HostType = hostType;
            Id = id;
            Ipv6 = ipv6;
            SourceIpv6 = sourceIpv6;
        }
    }
}
