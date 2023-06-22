// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class RouterMulticastInterfaceIgmpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Groups IGMP hosts are allowed to join.
        /// </summary>
        [Input("accessGroup")]
        public Input<string>? AccessGroup { get; set; }

        /// <summary>
        /// Groups to drop membership for immediately after receiving IGMPv2 leave.
        /// </summary>
        [Input("immediateLeaveGroup")]
        public Input<string>? ImmediateLeaveGroup { get; set; }

        /// <summary>
        /// Number of group specific queries before removing group (2 - 7, default = 2).
        /// </summary>
        [Input("lastMemberQueryCount")]
        public Input<int>? LastMemberQueryCount { get; set; }

        /// <summary>
        /// Timeout between IGMPv2 leave and removing group (1 - 65535 msec, default = 1000).
        /// </summary>
        [Input("lastMemberQueryInterval")]
        public Input<int>? LastMemberQueryInterval { get; set; }

        /// <summary>
        /// Interval between queries to IGMP hosts (1 - 65535 sec, default = 125).
        /// </summary>
        [Input("queryInterval")]
        public Input<int>? QueryInterval { get; set; }

        /// <summary>
        /// Maximum time to wait for a IGMP query response (1 - 25 sec, default = 10).
        /// </summary>
        [Input("queryMaxResponseTime")]
        public Input<int>? QueryMaxResponseTime { get; set; }

        /// <summary>
        /// Timeout between queries before becoming querier for network (60 - 900, default = 255).
        /// </summary>
        [Input("queryTimeout")]
        public Input<int>? QueryTimeout { get; set; }

        /// <summary>
        /// Enable/disable require IGMP packets contain router alert option. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("routerAlertCheck")]
        public Input<string>? RouterAlertCheck { get; set; }

        /// <summary>
        /// Maximum version of IGMP to support. Valid values: `3`, `2`, `1`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public RouterMulticastInterfaceIgmpArgs()
        {
        }
        public static new RouterMulticastInterfaceIgmpArgs Empty => new RouterMulticastInterfaceIgmpArgs();
    }
}
