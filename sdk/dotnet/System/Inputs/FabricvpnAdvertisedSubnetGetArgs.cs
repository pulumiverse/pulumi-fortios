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

    public sealed class FabricvpnAdvertisedSubnetGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access policy direction. Valid values: `inbound`, `bidirectional`.
        /// </summary>
        [Input("access")]
        public Input<string>? Access { get; set; }

        /// <summary>
        /// Underlying BGP network.
        /// </summary>
        [Input("bgpNetwork")]
        public Input<int>? BgpNetwork { get; set; }

        /// <summary>
        /// Underlying firewall address.
        /// </summary>
        [Input("firewallAddress")]
        public Input<string>? FirewallAddress { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Underlying policies.
        /// </summary>
        [Input("policies")]
        public Input<int>? Policies { get; set; }

        /// <summary>
        /// Network prefix.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public FabricvpnAdvertisedSubnetGetArgs()
        {
        }
        public static new FabricvpnAdvertisedSubnetGetArgs Empty => new FabricvpnAdvertisedSubnetGetArgs();
    }
}
