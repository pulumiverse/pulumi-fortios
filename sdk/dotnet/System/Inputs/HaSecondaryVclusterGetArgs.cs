// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Inputs
{

    public sealed class HaSecondaryVclusterGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interfaces to check for port monitoring (or link failure).
        /// </summary>
        [Input("monitor")]
        public Input<string>? Monitor { get; set; }

        /// <summary>
        /// Enable and increase the priority of the unit that should always be primary (master). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("override")]
        public Input<string>? Override { get; set; }

        /// <summary>
        /// Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.
        /// </summary>
        [Input("overrideWaitTime")]
        public Input<int>? OverrideWaitTime { get; set; }

        /// <summary>
        /// Remote IP monitoring failover threshold (0 - 50).
        /// </summary>
        [Input("pingserverFailoverThreshold")]
        public Input<int>? PingserverFailoverThreshold { get; set; }

        /// <summary>
        /// Interfaces to check for remote IP monitoring.
        /// </summary>
        [Input("pingserverMonitorInterface")]
        public Input<string>? PingserverMonitorInterface { get; set; }

        /// <summary>
        /// Enable to force the cluster to negotiate after a remote IP monitoring failover. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pingserverSecondaryForceReset")]
        public Input<string>? PingserverSecondaryForceReset { get; set; }

        /// <summary>
        /// Enable to force the cluster to negotiate after a remote IP monitoring failover. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pingserverSlaveForceReset")]
        public Input<string>? PingserverSlaveForceReset { get; set; }

        /// <summary>
        /// Increase the priority to select the primary unit (0 - 255).
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("vclusterId")]
        public Input<int>? VclusterId { get; set; }

        /// <summary>
        /// VDOMs in virtual cluster 2.
        /// </summary>
        [Input("vdom")]
        public Input<string>? Vdom { get; set; }

        public HaSecondaryVclusterGetArgs()
        {
        }
        public static new HaSecondaryVclusterGetArgs Empty => new HaSecondaryVclusterGetArgs();
    }
}
