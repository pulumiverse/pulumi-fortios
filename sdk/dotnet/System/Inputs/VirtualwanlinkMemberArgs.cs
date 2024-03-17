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

    public sealed class VirtualwanlinkMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).
        /// </summary>
        [Input("cost")]
        public Input<int>? Cost { get; set; }

        /// <summary>
        /// The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// IPv6 gateway.
        /// </summary>
        [Input("gateway6")]
        public Input<string>? Gateway6 { get; set; }

        /// <summary>
        /// Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
        /// </summary>
        [Input("ingressSpilloverThreshold")]
        public Input<int>? IngressSpilloverThreshold { get; set; }

        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Priority of the interface (0 - 4294967295). Used for SD-WAN rules or priority rules.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Member sequence number.
        /// </summary>
        [Input("seqNum")]
        public Input<int>? SeqNum { get; set; }

        /// <summary>
        /// Source IP address used in the health-check packet to the server.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Source IPv6 address used in the health-check packet to the server.
        /// </summary>
        [Input("source6")]
        public Input<string>? Source6 { get; set; }

        /// <summary>
        /// Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
        /// </summary>
        [Input("spilloverThreshold")]
        public Input<int>? SpilloverThreshold { get; set; }

        /// <summary>
        /// Enable/disable this interface in the SD-WAN. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Measured volume ratio (this value / sum of all values = percentage of link volume, 0 - 255).
        /// </summary>
        [Input("volumeRatio")]
        public Input<int>? VolumeRatio { get; set; }

        /// <summary>
        /// Weight of this interface for weighted load balancing. (0 - 255) More traffic is directed to interfaces with higher weights.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public VirtualwanlinkMemberArgs()
        {
        }
        public static new VirtualwanlinkMemberArgs Empty => new VirtualwanlinkMemberArgs();
    }
}
