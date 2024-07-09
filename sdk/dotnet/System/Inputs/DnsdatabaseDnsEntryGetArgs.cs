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

    public sealed class DnsdatabaseDnsEntryGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Canonical name of the host.
        /// </summary>
        [Input("canonicalName")]
        public Input<string>? CanonicalName { get; set; }

        /// <summary>
        /// Name of the host.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// DNS entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// IPv4 address of the host.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// IPv6 address of the host.
        /// </summary>
        [Input("ipv6")]
        public Input<string>? Ipv6 { get; set; }

        /// <summary>
        /// DNS entry preference (0 - 65535, highest preference = 0, default = 10).
        /// </summary>
        [Input("preference")]
        public Input<int>? Preference { get; set; }

        /// <summary>
        /// Enable/disable resource record status. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Time-to-live for this entry (0 to 2147483647 sec, default = 0).
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// Resource record type. Valid values: `A`, `NS`, `CNAME`, `MX`, `AAAA`, `PTR`, `PTR_V6`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DnsdatabaseDnsEntryGetArgs()
        {
        }
        public static new DnsdatabaseDnsEntryGetArgs Empty => new DnsdatabaseDnsEntryGetArgs();
    }
}
