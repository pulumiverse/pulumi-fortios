// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class FirewallInternetserviceadditionEntryGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("addrMode")]
        public Input<string>? AddrMode { get; set; }

        /// <summary>
        /// Entry ID(1-255).
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("portRanges")]
        private InputList<Inputs.FirewallInternetserviceadditionEntryPortRangeGetArgs>? _portRanges;

        /// <summary>
        /// Port ranges in the custom entry. The structure of `port_range` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallInternetserviceadditionEntryPortRangeGetArgs> PortRanges
        {
            get => _portRanges ?? (_portRanges = new InputList<Inputs.FirewallInternetserviceadditionEntryPortRangeGetArgs>());
            set => _portRanges = value;
        }

        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255).
        /// </summary>
        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        public FirewallInternetserviceadditionEntryGetArgs()
        {
        }
        public static new FirewallInternetserviceadditionEntryGetArgs Empty => new FirewallInternetserviceadditionEntryGetArgs();
    }
}
