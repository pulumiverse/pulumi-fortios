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
    public sealed class ProfileprotocoloptionsSmtp
    {
        /// <summary>
        /// Enable/disable the inspection of all ports for the protocol. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? InspectAll;
        /// <summary>
        /// One or more options that can be applied to the session. Valid values: `fragmail`, `oversize`, `splice`.
        /// </summary>
        public readonly string? Options;
        /// <summary>
        /// Maximum in-memory file size that can be scanned (MB).
        /// </summary>
        public readonly int? OversizeLimit;
        /// <summary>
        /// Ports to scan for content (1 - 65535, default = 25).
        /// </summary>
        public readonly int? Ports;
        /// <summary>
        /// Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ProxyAfterTcpHandshake;
        /// <summary>
        /// Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ScanBzip2;
        /// <summary>
        /// Enable/disable SMTP server busy when server not available. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ServerBusy;
        /// <summary>
        /// SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.
        /// </summary>
        public readonly string? SslOffloaded;
        /// <summary>
        /// Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
        /// </summary>
        public readonly int? UncompressedNestLimit;
        /// <summary>
        /// Maximum in-memory uncompressed file size that can be scanned (MB).
        /// </summary>
        public readonly int? UncompressedOversizeLimit;

        [OutputConstructor]
        private ProfileprotocoloptionsSmtp(
            string? inspectAll,

            string? options,

            int? oversizeLimit,

            int? ports,

            string? proxyAfterTcpHandshake,

            string? scanBzip2,

            string? serverBusy,

            string? sslOffloaded,

            string? status,

            int? uncompressedNestLimit,

            int? uncompressedOversizeLimit)
        {
            InspectAll = inspectAll;
            Options = options;
            OversizeLimit = oversizeLimit;
            Ports = ports;
            ProxyAfterTcpHandshake = proxyAfterTcpHandshake;
            ScanBzip2 = scanBzip2;
            ServerBusy = serverBusy;
            SslOffloaded = sslOffloaded;
            Status = status;
            UncompressedNestLimit = uncompressedNestLimit;
            UncompressedOversizeLimit = uncompressedOversizeLimit;
        }
    }
}
