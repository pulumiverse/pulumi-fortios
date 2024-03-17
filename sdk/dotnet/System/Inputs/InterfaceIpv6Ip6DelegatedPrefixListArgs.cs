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

    public sealed class InterfaceIpv6Ip6DelegatedPrefixListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable the autonomous flag. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autonomousFlag")]
        public Input<string>? AutonomousFlag { get; set; }

        /// <summary>
        /// IAID of obtained delegated-prefix from the upstream interface.
        /// </summary>
        [Input("delegatedPrefixIaid")]
        public Input<int>? DelegatedPrefixIaid { get; set; }

        /// <summary>
        /// Enable/disable the onlink flag. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("onlinkFlag")]
        public Input<string>? OnlinkFlag { get; set; }

        /// <summary>
        /// Prefix ID.
        /// </summary>
        [Input("prefixId")]
        public Input<int>? PrefixId { get; set; }

        /// <summary>
        /// Recursive DNS server option.
        /// 
        /// The `dhcp6_iapd_list` block supports:
        /// </summary>
        [Input("rdnss")]
        public Input<string>? Rdnss { get; set; }

        /// <summary>
        /// Recursive DNS service option. Valid values: `delegated`, `default`, `specify`.
        /// </summary>
        [Input("rdnssService")]
        public Input<string>? RdnssService { get; set; }

        /// <summary>
        /// Add subnet ID to routing prefix.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// Name of the interface that provides delegated information.
        /// </summary>
        [Input("upstreamInterface")]
        public Input<string>? UpstreamInterface { get; set; }

        public InterfaceIpv6Ip6DelegatedPrefixListArgs()
        {
        }
        public static new InterfaceIpv6Ip6DelegatedPrefixListArgs Empty => new InterfaceIpv6Ip6DelegatedPrefixListArgs();
    }
}
