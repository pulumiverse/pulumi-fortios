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

    public sealed class InterfaceIpv6Dhcp6IapdListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identity association identifier.
        /// </summary>
        [Input("iaid")]
        public Input<int>? Iaid { get; set; }

        /// <summary>
        /// DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
        /// </summary>
        [Input("prefixHint")]
        public Input<string>? PrefixHint { get; set; }

        /// <summary>
        /// DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
        /// </summary>
        [Input("prefixHintPlt")]
        public Input<int>? PrefixHintPlt { get; set; }

        /// <summary>
        /// DHCPv6 prefix hint valid life time (sec).
        /// 
        /// The `vrrp6` block supports:
        /// </summary>
        [Input("prefixHintVlt")]
        public Input<int>? PrefixHintVlt { get; set; }

        public InterfaceIpv6Dhcp6IapdListGetArgs()
        {
        }
        public static new InterfaceIpv6Dhcp6IapdListGetArgs Empty => new InterfaceIpv6Dhcp6IapdListGetArgs();
    }
}
