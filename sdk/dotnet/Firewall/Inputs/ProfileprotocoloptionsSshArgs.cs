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

    public sealed class ProfileprotocoloptionsSshArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of bytes to send in each transmission for client comforting (bytes).
        /// </summary>
        [Input("comfortAmount")]
        public Input<int>? ComfortAmount { get; set; }

        /// <summary>
        /// Interval between successive transmissions of data for client comforting (seconds).
        /// </summary>
        [Input("comfortInterval")]
        public Input<int>? ComfortInterval { get; set; }

        /// <summary>
        /// One or more options that can be applied to the session. Valid values: `oversize`, `clientcomfort`, `servercomfort`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Maximum in-memory file size that can be scanned (MB).
        /// </summary>
        [Input("oversizeLimit")]
        public Input<int>? OversizeLimit { get; set; }

        /// <summary>
        /// Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("scanBzip2")]
        public Input<string>? ScanBzip2 { get; set; }

        /// <summary>
        /// SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.
        /// </summary>
        [Input("sslOffloaded")]
        public Input<string>? SslOffloaded { get; set; }

        /// <summary>
        /// Maximum stream-based uncompressed data size that will be scanned in megabytes. Stream-based uncompression used only under certain conditions (unlimited = 0, default = 0).
        /// </summary>
        [Input("streamBasedUncompressedLimit")]
        public Input<int>? StreamBasedUncompressedLimit { get; set; }

        /// <summary>
        /// Maximum dynamic TCP window size.
        /// </summary>
        [Input("tcpWindowMaximum")]
        public Input<int>? TcpWindowMaximum { get; set; }

        /// <summary>
        /// Minimum dynamic TCP window size.
        /// </summary>
        [Input("tcpWindowMinimum")]
        public Input<int>? TcpWindowMinimum { get; set; }

        /// <summary>
        /// Set TCP static window size.
        /// </summary>
        [Input("tcpWindowSize")]
        public Input<int>? TcpWindowSize { get; set; }

        /// <summary>
        /// TCP window type to use for this protocol.
        /// </summary>
        [Input("tcpWindowType")]
        public Input<string>? TcpWindowType { get; set; }

        /// <summary>
        /// Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
        /// </summary>
        [Input("uncompressedNestLimit")]
        public Input<int>? UncompressedNestLimit { get; set; }

        /// <summary>
        /// Maximum in-memory uncompressed file size that can be scanned (MB).
        /// </summary>
        [Input("uncompressedOversizeLimit")]
        public Input<int>? UncompressedOversizeLimit { get; set; }

        public ProfileprotocoloptionsSshArgs()
        {
        }
        public static new ProfileprotocoloptionsSshArgs Empty => new ProfileprotocoloptionsSshArgs();
    }
}
