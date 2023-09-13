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
    public sealed class CommunityHost
    {
        /// <summary>
        /// Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? HaDirect;
        /// <summary>
        /// Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. Valid values: `any`, `query`, `trap`.
        /// </summary>
        public readonly string? HostType;
        /// <summary>
        /// Host6 entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// IPv4 address of the SNMP manager (host).
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Source IPv4 address for SNMP traps.
        /// </summary>
        public readonly string? SourceIp;

        [OutputConstructor]
        private CommunityHost(
            string? haDirect,

            string? hostType,

            int? id,

            string? ip,

            string? sourceIp)
        {
            HaDirect = haDirect;
            HostType = hostType;
            Id = id;
            Ip = ip;
            SourceIp = sourceIp;
        }
    }
}
