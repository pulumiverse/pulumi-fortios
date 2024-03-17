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
    public sealed class StandaloneclusterClusterPeerSessionSyncFilter
    {
        /// <summary>
        /// Only sessions using these custom services are synchronized. Use source and destination port ranges to define these custom services. The structure of `custom_service` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.StandaloneclusterClusterPeerSessionSyncFilterCustomService> CustomServices;
        /// <summary>
        /// Only sessions to this IPv4 address are synchronized.
        /// </summary>
        public readonly string? Dstaddr;
        /// <summary>
        /// Only sessions to this IPv6 address are synchronized.
        /// </summary>
        public readonly string? Dstaddr6;
        /// <summary>
        /// Only sessions to this interface are synchronized.
        /// </summary>
        public readonly string? Dstintf;
        /// <summary>
        /// Only sessions from this IPv4 address are synchronized.
        /// </summary>
        public readonly string? Srcaddr;
        /// <summary>
        /// Only sessions from this IPv6 address are synchronized.
        /// </summary>
        public readonly string? Srcaddr6;
        /// <summary>
        /// Only sessions from this interface are synchronized.
        /// </summary>
        public readonly string? Srcintf;

        [OutputConstructor]
        private StandaloneclusterClusterPeerSessionSyncFilter(
            ImmutableArray<Outputs.StandaloneclusterClusterPeerSessionSyncFilterCustomService> customServices,

            string? dstaddr,

            string? dstaddr6,

            string? dstintf,

            string? srcaddr,

            string? srcaddr6,

            string? srcintf)
        {
            CustomServices = customServices;
            Dstaddr = dstaddr;
            Dstaddr6 = dstaddr6;
            Dstintf = dstintf;
            Srcaddr = srcaddr;
            Srcaddr6 = srcaddr6;
            Srcintf = srcintf;
        }
    }
}
