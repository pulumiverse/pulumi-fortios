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
    public sealed class GetProfileprotocoloptionsFtpResult
    {
        /// <summary>
        /// Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).
        /// </summary>
        public readonly int ComfortAmount;
        /// <summary>
        /// Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
        /// </summary>
        public readonly int ComfortInterval;
        /// <summary>
        /// Enable/disable FTP redirection for explicit FTPS.
        /// </summary>
        public readonly string ExplicitFtpTls;
        /// <summary>
        /// Enable/disable the inspection of all ports for the protocol.
        /// </summary>
        public readonly string InspectAll;
        /// <summary>
        /// One or more options that can be applied to the session.
        /// </summary>
        public readonly string Options;
        /// <summary>
        /// Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
        /// </summary>
        public readonly int OversizeLimit;
        /// <summary>
        /// Ports to scan for content (1 - 65535, default = 445).
        /// </summary>
        public readonly int Ports;
        /// <summary>
        /// Enable/disable scanning of BZip2 compressed files.
        /// </summary>
        public readonly string ScanBzip2;
        /// <summary>
        /// SSL decryption and encryption performed by an external device.
        /// </summary>
        public readonly string SslOffloaded;
        /// <summary>
        /// Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).
        /// </summary>
        public readonly int StreamBasedUncompressedLimit;
        /// <summary>
        /// Maximum dynamic TCP window size (default = 8MB).
        /// </summary>
        public readonly int TcpWindowMaximum;
        /// <summary>
        /// Minimum dynamic TCP window size (default = 128KB).
        /// </summary>
        public readonly int TcpWindowMinimum;
        /// <summary>
        /// Set TCP static window size (default = 256KB).
        /// </summary>
        public readonly int TcpWindowSize;
        /// <summary>
        /// Specify type of TCP window to use for this protocol.
        /// </summary>
        public readonly string TcpWindowType;
        /// <summary>
        /// Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
        /// </summary>
        public readonly int UncompressedNestLimit;
        /// <summary>
        /// Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
        /// </summary>
        public readonly int UncompressedOversizeLimit;

        [OutputConstructor]
        private GetProfileprotocoloptionsFtpResult(
            int comfortAmount,

            int comfortInterval,

            string explicitFtpTls,

            string inspectAll,

            string options,

            int oversizeLimit,

            int ports,

            string scanBzip2,

            string sslOffloaded,

            string status,

            int streamBasedUncompressedLimit,

            int tcpWindowMaximum,

            int tcpWindowMinimum,

            int tcpWindowSize,

            string tcpWindowType,

            int uncompressedNestLimit,

            int uncompressedOversizeLimit)
        {
            ComfortAmount = comfortAmount;
            ComfortInterval = comfortInterval;
            ExplicitFtpTls = explicitFtpTls;
            InspectAll = inspectAll;
            Options = options;
            OversizeLimit = oversizeLimit;
            Ports = ports;
            ScanBzip2 = scanBzip2;
            SslOffloaded = sslOffloaded;
            Status = status;
            StreamBasedUncompressedLimit = streamBasedUncompressedLimit;
            TcpWindowMaximum = tcpWindowMaximum;
            TcpWindowMinimum = tcpWindowMinimum;
            TcpWindowSize = tcpWindowSize;
            TcpWindowType = tcpWindowType;
            UncompressedNestLimit = uncompressedNestLimit;
            UncompressedOversizeLimit = uncompressedOversizeLimit;
        }
    }
}