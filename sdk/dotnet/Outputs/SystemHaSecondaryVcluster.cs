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
    public sealed class SystemHaSecondaryVcluster
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
        public readonly string? Vdom;

        [OutputConstructor]
        private SystemHaSecondaryVcluster(
            string? monitor,

            string? @override,

            int? overrideWaitTime,

            int? pingserverFailoverThreshold,

            string? pingserverMonitorInterface,

            string? pingserverSecondaryForceReset,

            string? pingserverSlaveForceReset,

            int? priority,

            int? vclusterId,

            string? vdom)
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
            Vdom = vdom;
        }
    }
}
