// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Outputs
{

    [OutputType]
    public sealed class AccessproxyApiGateway6Quic
    {
        /// <summary>
        /// ACK delay exponent (1 - 20, default = 3).
        /// </summary>
        public readonly int? AckDelayExponent;
        /// <summary>
        /// Active connection ID limit (1 - 8, default = 2).
        /// </summary>
        public readonly int? ActiveConnectionIdLimit;
        /// <summary>
        /// Enable/disable active migration (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ActiveMigration;
        /// <summary>
        /// Enable/disable grease QUIC bit (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? GreaseQuicBit;
        /// <summary>
        /// Maximum ACK delay in milliseconds (1 - 16383, default = 25).
        /// </summary>
        public readonly int? MaxAckDelay;
        /// <summary>
        /// Maximum datagram frame size in bytes (1 - 1500, default = 1500).
        /// </summary>
        public readonly int? MaxDatagramFrameSize;
        /// <summary>
        /// Maximum idle timeout milliseconds (1 - 60000, default = 30000).
        /// </summary>
        public readonly int? MaxIdleTimeout;
        /// <summary>
        /// Maximum UDP payload size in bytes (1200 - 1500, default = 1500).
        /// </summary>
        public readonly int? MaxUdpPayloadSize;

        [OutputConstructor]
        private AccessproxyApiGateway6Quic(
            int? ackDelayExponent,

            int? activeConnectionIdLimit,

            string? activeMigration,

            string? greaseQuicBit,

            int? maxAckDelay,

            int? maxDatagramFrameSize,

            int? maxIdleTimeout,

            int? maxUdpPayloadSize)
        {
            AckDelayExponent = ackDelayExponent;
            ActiveConnectionIdLimit = activeConnectionIdLimit;
            ActiveMigration = activeMigration;
            GreaseQuicBit = greaseQuicBit;
            MaxAckDelay = maxAckDelay;
            MaxDatagramFrameSize = maxDatagramFrameSize;
            MaxIdleTimeout = maxIdleTimeout;
            MaxUdpPayloadSize = maxUdpPayloadSize;
        }
    }
}
