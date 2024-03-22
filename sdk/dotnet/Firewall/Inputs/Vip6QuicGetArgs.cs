// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class Vip6QuicGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ACK delay exponent (1 - 20, default = 3).
        /// </summary>
        [Input("ackDelayExponent")]
        public Input<int>? AckDelayExponent { get; set; }

        /// <summary>
        /// Active connection ID limit (1 - 8, default = 2).
        /// </summary>
        [Input("activeConnectionIdLimit")]
        public Input<int>? ActiveConnectionIdLimit { get; set; }

        /// <summary>
        /// Enable/disable active migration (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("activeMigration")]
        public Input<string>? ActiveMigration { get; set; }

        /// <summary>
        /// Enable/disable grease QUIC bit (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("greaseQuicBit")]
        public Input<string>? GreaseQuicBit { get; set; }

        /// <summary>
        /// Maximum ACK delay in milliseconds (1 - 16383, default = 25).
        /// </summary>
        [Input("maxAckDelay")]
        public Input<int>? MaxAckDelay { get; set; }

        /// <summary>
        /// Maximum datagram frame size in bytes (1 - 1500, default = 1500).
        /// </summary>
        [Input("maxDatagramFrameSize")]
        public Input<int>? MaxDatagramFrameSize { get; set; }

        /// <summary>
        /// Maximum idle timeout milliseconds (1 - 60000, default = 30000).
        /// </summary>
        [Input("maxIdleTimeout")]
        public Input<int>? MaxIdleTimeout { get; set; }

        /// <summary>
        /// Maximum UDP payload size in bytes (1200 - 1500, default = 1500).
        /// </summary>
        [Input("maxUdpPayloadSize")]
        public Input<int>? MaxUdpPayloadSize { get; set; }

        public Vip6QuicGetArgs()
        {
        }
        public static new Vip6QuicGetArgs Empty => new Vip6QuicGetArgs();
    }
}
