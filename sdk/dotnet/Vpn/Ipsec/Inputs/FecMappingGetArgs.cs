// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ipsec.Inputs
{

    public sealed class FecMappingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Apply FEC parameters when available bi-bandwidth is &gt;= threshold (kbps, 0 means no threshold).
        /// </summary>
        [Input("bandwidthBiThreshold")]
        public Input<int>? BandwidthBiThreshold { get; set; }

        /// <summary>
        /// Apply FEC parameters when available down bandwidth is &gt;= threshold (kbps, 0 means no threshold).
        /// </summary>
        [Input("bandwidthDownThreshold")]
        public Input<int>? BandwidthDownThreshold { get; set; }

        /// <summary>
        /// Apply FEC parameters when available up bandwidth is &gt;= threshold (kbps, 0 means no threshold).
        /// </summary>
        [Input("bandwidthUpThreshold")]
        public Input<int>? BandwidthUpThreshold { get; set; }

        /// <summary>
        /// Number of base FEC packets (1 - 20).
        /// </summary>
        [Input("base")]
        public Input<int>? Base { get; set; }

        /// <summary>
        /// Apply FEC parameters when latency is &lt;= threshold (0 means no threshold).
        /// </summary>
        [Input("latencyThreshold")]
        public Input<int>? LatencyThreshold { get; set; }

        /// <summary>
        /// Apply FEC parameters when packet loss is &gt;= threshold (0 - 100, 0 means no threshold).
        /// </summary>
        [Input("packetLossThreshold")]
        public Input<int>? PacketLossThreshold { get; set; }

        /// <summary>
        /// Number of redundant FEC packets (1 - 5).
        /// </summary>
        [Input("redundant")]
        public Input<int>? Redundant { get; set; }

        /// <summary>
        /// Sequence number (1 - 64).
        /// </summary>
        [Input("seqno")]
        public Input<int>? Seqno { get; set; }

        public FecMappingGetArgs()
        {
        }
        public static new FecMappingGetArgs Empty => new FecMappingGetArgs();
    }
}
