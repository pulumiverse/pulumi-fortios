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
    public sealed class VirtualwanlinkNeighbor
    {
        /// <summary>
        /// SD-WAN health-check name.
        /// </summary>
        public readonly string? HealthCheck;
        /// <summary>
        /// IP address of neighbor.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Member sequence number.
        /// </summary>
        public readonly int? Member;
        /// <summary>
        /// Role of neighbor. Valid values: `standalone`, `primary`, `secondary`.
        /// </summary>
        public readonly string? Role;
        /// <summary>
        /// SLA ID.
        /// </summary>
        public readonly int? SlaId;

        [OutputConstructor]
        private VirtualwanlinkNeighbor(
            string? healthCheck,

            string? ip,

            int? member,

            string? role,

            int? slaId)
        {
            HealthCheck = healthCheck;
            Ip = ip;
            Member = member;
            Role = role;
            SlaId = slaId;
        }
    }
}
