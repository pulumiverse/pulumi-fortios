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
    public sealed class ProfileprotocoloptionsHttp
    {
        /// <summary>
        /// Enable/disable IP based URL rating. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? AddressIpRating;
        /// <summary>
        /// Code number returned for blocked HTTP pages (non-FortiGuard only) (100 - 599, default = 403).
        /// </summary>
        public readonly int? BlockPageStatusCode;
        /// <summary>
        /// Amount of data to send in a transmission for client comforting (1 - 10240 bytes, default = 1).
        /// </summary>
        public readonly int? ComfortAmount;
        /// <summary>
        /// Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
        /// </summary>
        public readonly int? ComfortInterval;
        /// <summary>
        /// Enable/disable Fortinet bar on HTML content. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? FortinetBar;
        /// <summary>
        /// Port for use by Fortinet Bar (1 - 65535, default = 8011).
        /// </summary>
        public readonly int? FortinetBarPort;
        /// <summary>
        /// Enable/disable h2c HTTP connection upgrade. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? H2c;
        /// <summary>
        /// Enable/disable HTTP policy check. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? HttpPolicy;
        /// <summary>
        /// Enable/disable the inspection of all ports for the protocol. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? InspectAll;
        /// <summary>
        /// One or more options that can be applied to the session. Valid values: `clientcomfort`, `servercomfort`, `oversize`, `chunkedbypass`.
        /// </summary>
        public readonly string? Options;
        /// <summary>
        /// Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
        /// </summary>
        public readonly int? OversizeLimit;
        /// <summary>
        /// Ports to scan for content (1 - 65535, default = 80).
        /// </summary>
        public readonly int? Ports;
        /// <summary>
        /// ID codes for character sets to be used to convert to UTF-8 for banned words and DLP on HTTP posts (maximum of 5 character sets). Valid values: `jisx0201`, `jisx0208`, `jisx0212`, `gb2312`, `ksc5601-ex`, `euc-jp`, `sjis`, `iso2022-jp`, `iso2022-jp-1`, `iso2022-jp-2`, `euc-cn`, `ces-gbk`, `hz`, `ces-big5`, `euc-kr`, `iso2022-jp-3`, `iso8859-1`, `tis620`, `cp874`, `cp1252`, `cp1251`.
        /// </summary>
        public readonly string? PostLang;
        /// <summary>
        /// Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ProxyAfterTcpHandshake;
        /// <summary>
        /// Enable/disable blocking of partial downloads. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? RangeBlock;
        /// <summary>
        /// Number of attempts to retry HTTP connection (0 - 100, default = 0).
        /// </summary>
        public readonly int? RetryCount;
        /// <summary>
        /// Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ScanBzip2;
        /// <summary>
        /// SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.
        /// </summary>
        public readonly string? SslOffloaded;
        /// <summary>
        /// Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).
        /// </summary>
        public readonly int? StreamBasedUncompressedLimit;
        /// <summary>
        /// Enable/disable bypassing of streaming content from buffering. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? StreamingContentBypass;
        /// <summary>
        /// Enable/disable stripping of HTTP X-Forwarded-For header. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? StripXForwardedFor;
        /// <summary>
        /// Bypass from scanning, or block a connection that attempts to switch protocol. Valid values: `bypass`, `block`.
        /// </summary>
        public readonly string? SwitchingProtocols;
        /// <summary>
        /// Maximum dynamic TCP window size (default = 8MB).
        /// </summary>
        public readonly int? TcpWindowMaximum;
        /// <summary>
        /// Minimum dynamic TCP window size (default = 128KB).
        /// </summary>
        public readonly int? TcpWindowMinimum;
        /// <summary>
        /// Set TCP static window size (default = 256KB).
        /// </summary>
        public readonly int? TcpWindowSize;
        /// <summary>
        /// Specify type of TCP window to use for this protocol.
        /// </summary>
        public readonly string? TcpWindowType;
        /// <summary>
        /// Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? TunnelNonHttp;
        /// <summary>
        /// Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
        /// </summary>
        public readonly int? UncompressedNestLimit;
        /// <summary>
        /// Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
        /// </summary>
        public readonly int? UncompressedOversizeLimit;
        /// <summary>
        /// How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1. Valid values: `reject`, `tunnel`, `best-effort`.
        /// </summary>
        public readonly string? UnknownHttpVersion;
        /// <summary>
        /// Enable/disable verification of DNS for policy matching. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? VerifyDnsForPolicyMatching;

        [OutputConstructor]
        private ProfileprotocoloptionsHttp(
            string? addressIpRating,

            int? blockPageStatusCode,

            int? comfortAmount,

            int? comfortInterval,

            string? fortinetBar,

            int? fortinetBarPort,

            string? h2c,

            string? httpPolicy,

            string? inspectAll,

            string? options,

            int? oversizeLimit,

            int? ports,

            string? postLang,

            string? proxyAfterTcpHandshake,

            string? rangeBlock,

            int? retryCount,

            string? scanBzip2,

            string? sslOffloaded,

            string? status,

            int? streamBasedUncompressedLimit,

            string? streamingContentBypass,

            string? stripXForwardedFor,

            string? switchingProtocols,

            int? tcpWindowMaximum,

            int? tcpWindowMinimum,

            int? tcpWindowSize,

            string? tcpWindowType,

            string? tunnelNonHttp,

            int? uncompressedNestLimit,

            int? uncompressedOversizeLimit,

            string? unknownHttpVersion,

            string? verifyDnsForPolicyMatching)
        {
            AddressIpRating = addressIpRating;
            BlockPageStatusCode = blockPageStatusCode;
            ComfortAmount = comfortAmount;
            ComfortInterval = comfortInterval;
            FortinetBar = fortinetBar;
            FortinetBarPort = fortinetBarPort;
            H2c = h2c;
            HttpPolicy = httpPolicy;
            InspectAll = inspectAll;
            Options = options;
            OversizeLimit = oversizeLimit;
            Ports = ports;
            PostLang = postLang;
            ProxyAfterTcpHandshake = proxyAfterTcpHandshake;
            RangeBlock = rangeBlock;
            RetryCount = retryCount;
            ScanBzip2 = scanBzip2;
            SslOffloaded = sslOffloaded;
            Status = status;
            StreamBasedUncompressedLimit = streamBasedUncompressedLimit;
            StreamingContentBypass = streamingContentBypass;
            StripXForwardedFor = stripXForwardedFor;
            SwitchingProtocols = switchingProtocols;
            TcpWindowMaximum = tcpWindowMaximum;
            TcpWindowMinimum = tcpWindowMinimum;
            TcpWindowSize = tcpWindowSize;
            TcpWindowType = tcpWindowType;
            TunnelNonHttp = tunnelNonHttp;
            UncompressedNestLimit = uncompressedNestLimit;
            UncompressedOversizeLimit = uncompressedOversizeLimit;
            UnknownHttpVersion = unknownHttpVersion;
            VerifyDnsForPolicyMatching = verifyDnsForPolicyMatching;
        }
    }
}
