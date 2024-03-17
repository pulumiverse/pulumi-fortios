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
    public sealed class SdwanNeighbor
    {
        /// <summary>
        /// SD-WAN health-check name.
        /// </summary>
        public readonly string? HealthCheck;
        /// <summary>
        /// IP/IPv6 address of neighbor.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Member sequence number.
        /// </summary>
        public readonly int? Member;
        /// <summary>
        /// Minimum number of members which meet SLA when the neighbor is preferred.
        /// </summary>
        public readonly int? MinimumSlaMeetMembers;
        /// <summary>
        /// What metric to select the neighbor. Valid values: `sla`, `speedtest`.
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// Role of neighbor. Valid values: `standalone`, `primary`, `secondary`.
        /// </summary>
        public readonly string? Role;
        /// <summary>
        /// SLA ID.
        /// </summary>
        public readonly int? SlaId;

        [OutputConstructor]
        private SdwanNeighbor(
            string? healthCheck,

            string? ip,

            int? member,

            int? minimumSlaMeetMembers,

            string? mode,

            string? role,

            int? slaId)
        {
            HealthCheck = healthCheck;
            Ip = ip;
            Member = member;
            MinimumSlaMeetMembers = minimumSlaMeetMembers;
            Mode = mode;
            Role = role;
            SlaId = slaId;
        }
    }
}
