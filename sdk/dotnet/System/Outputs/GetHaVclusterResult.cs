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
    public sealed class GetHaVclusterResult
    {
        /// <summary>
        /// Interfaces to check for port monitoring (or link failure).
        /// </summary>
        public readonly string Monitor;
        /// <summary>
        /// Enable and increase the priority of the unit that should always be primary (master).
        /// </summary>
        public readonly string Override;
        /// <summary>
        /// Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.
        /// </summary>
        public readonly int OverrideWaitTime;
        /// <summary>
        /// Remote IP monitoring failover threshold (0 - 50).
        /// </summary>
        public readonly int PingserverFailoverThreshold;
        /// <summary>
        /// Time to wait in minutes before renegotiating after a remote IP monitoring failover.
        /// </summary>
        public readonly int PingserverFlipTimeout;
        /// <summary>
        /// Interfaces to check for remote IP monitoring.
        /// </summary>
        public readonly string PingserverMonitorInterface;
        /// <summary>
        /// Enable to force the cluster to negotiate after a remote IP monitoring failover.
        /// </summary>
        public readonly string PingserverSecondaryForceReset;
        /// <summary>
        /// Enable to force the cluster to negotiate after a remote IP monitoring failover.
        /// </summary>
        public readonly string PingserverSlaveForceReset;
        /// <summary>
        /// Increase the priority to select the primary unit (0 - 255).
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// Cluster ID.
        /// </summary>
        public readonly int VclusterId;
        /// <summary>
        /// VDOMs in virtual cluster 2.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetHaVclusterVdomResult> Vdoms;

        [OutputConstructor]
        private GetHaVclusterResult(
            string monitor,

            string @override,

            int overrideWaitTime,

            int pingserverFailoverThreshold,

            int pingserverFlipTimeout,

            string pingserverMonitorInterface,

            string pingserverSecondaryForceReset,

            string pingserverSlaveForceReset,

            int priority,

            int vclusterId,

            ImmutableArray<Outputs.GetHaVclusterVdomResult> vdoms)
        {
            Monitor = monitor;
            Override = @override;
            OverrideWaitTime = overrideWaitTime;
            PingserverFailoverThreshold = pingserverFailoverThreshold;
            PingserverFlipTimeout = pingserverFlipTimeout;
            PingserverMonitorInterface = pingserverMonitorInterface;
            PingserverSecondaryForceReset = pingserverSecondaryForceReset;
            PingserverSlaveForceReset = pingserverSlaveForceReset;
            Priority = priority;
            VclusterId = vclusterId;
            Vdoms = vdoms;
        }
    }
}
