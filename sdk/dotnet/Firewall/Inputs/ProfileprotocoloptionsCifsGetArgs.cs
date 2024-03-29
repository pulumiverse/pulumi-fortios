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

    public sealed class ProfileprotocoloptionsCifsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain for which to decrypt CIFS traffic.
        /// </summary>
        [Input("domainController")]
        public Input<string>? DomainController { get; set; }

        /// <summary>
        /// One or more options that can be applied to the session. Valid values: `oversize`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Maximum in-memory file size that can be scanned (MB).
        /// </summary>
        [Input("oversizeLimit")]
        public Input<int>? OversizeLimit { get; set; }

        /// <summary>
        /// Ports to scan for content (1 - 65535, default = 445).
        /// </summary>
        [Input("ports")]
        public Input<int>? Ports { get; set; }

        /// <summary>
        /// Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("scanBzip2")]
        public Input<string>? ScanBzip2 { get; set; }

        /// <summary>
        /// CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
        /// </summary>
        [Input("serverCredentialType")]
        public Input<string>? ServerCredentialType { get; set; }

        [Input("serverKeytabs")]
        private InputList<Inputs.ProfileprotocoloptionsCifsServerKeytabGetArgs>? _serverKeytabs;

        /// <summary>
        /// Server keytab. The structure of `server_keytab` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileprotocoloptionsCifsServerKeytabGetArgs> ServerKeytabs
        {
            get => _serverKeytabs ?? (_serverKeytabs = new InputList<Inputs.ProfileprotocoloptionsCifsServerKeytabGetArgs>());
            set => _serverKeytabs = value;
        }

        /// <summary>
        /// Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

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
        /// Specify type of TCP window to use for this protocol.
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

        public ProfileprotocoloptionsCifsGetArgs()
        {
        }
        public static new ProfileprotocoloptionsCifsGetArgs Empty => new ProfileprotocoloptionsCifsGetArgs();
    }
}
