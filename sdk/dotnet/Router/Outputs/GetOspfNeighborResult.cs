// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class GetOspfNeighborResult
    {
        /// <summary>
        /// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
        /// </summary>
        public readonly int Cost;
        /// <summary>
        /// Distribute list entry ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Interface IP address of the neighbor.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// Poll interval time in seconds.
        /// </summary>
        public readonly int PollInterval;
        /// <summary>
        /// Priority.
        /// </summary>
        public readonly int Priority;

        [OutputConstructor]
        private GetOspfNeighborResult(
            int cost,

            int id,

            string ip,

            int pollInterval,

            int priority)
        {
            Cost = cost;
            Id = id;
            Ip = ip;
            PollInterval = pollInterval;
            Priority = priority;
        }
    }
}
