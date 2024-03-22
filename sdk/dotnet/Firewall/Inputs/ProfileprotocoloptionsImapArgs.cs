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

    public sealed class ProfileprotocoloptionsImapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable the inspection of all ports for the protocol. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("inspectAll")]
        public Input<string>? InspectAll { get; set; }

        /// <summary>
        /// One or more options that can be applied to the session. Valid values: `fragmail`, `oversize`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Maximum in-memory file size that can be scanned (MB).
        /// </summary>
        [Input("oversizeLimit")]
        public Input<int>? OversizeLimit { get; set; }

        /// <summary>
        /// Ports to scan for content (1 - 65535, default = 143).
        /// </summary>
        [Input("ports")]
        public Input<int>? Ports { get; set; }

        /// <summary>
        /// Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("proxyAfterTcpHandshake")]
        public Input<string>? ProxyAfterTcpHandshake { get; set; }

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
        /// Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

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

        public ProfileprotocoloptionsImapArgs()
        {
        }
        public static new ProfileprotocoloptionsImapArgs Empty => new ProfileprotocoloptionsImapArgs();
    }
}
