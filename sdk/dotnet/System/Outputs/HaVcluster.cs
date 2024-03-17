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
    public sealed class HaVcluster
    {
        public readonly string? Monitor;
        public readonly string? Override;
        public readonly int? OverrideWaitTime;
        public readonly int? PingserverFailoverThreshold;
        public readonly string? PingserverMonitorInterface;
        public readonly string? PingserverSecondaryForceReset;
        public readonly string? PingserverSlaveForceReset;
        public readonly int? Priority;
        public readonly int? VclusterId;
        public readonly ImmutableArray<Outputs.HaVclusterVdom> Vdoms;

        [OutputConstructor]
        private HaVcluster(
            string? monitor,

            string? @override,

            int? overrideWaitTime,

            int? pingserverFailoverThreshold,

            string? pingserverMonitorInterface,

            string? pingserverSecondaryForceReset,

            string? pingserverSlaveForceReset,

            int? priority,

            int? vclusterId,

            ImmutableArray<Outputs.HaVclusterVdom> vdoms)
        {
            Monitor = monitor;
            Override = @override;
            OverrideWaitTime = overrideWaitTime;
            PingserverFailoverThreshold = pingserverFailoverThreshold;
            PingserverMonitorInterface = pingserverMonitorInterface;
            PingserverSecondaryForceReset = pingserverSecondaryForceReset;
            PingserverSlaveForceReset = pingserverSlaveForceReset;
            Priority = priority;
            VclusterId = vclusterId;
            Vdoms = vdoms;
        }
    }
}
