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
    public sealed class InternetservicedefinitionEntry
    {
        /// <summary>
        /// Internet Service category ID.
        /// </summary>
        public readonly int? CategoryId;
        /// <summary>
        /// Internet Service name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Integer value for ending TCP/UDP/SCTP destination port in range (0 to 65535). 0 means undefined.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Port ranges in the definition entry. The structure of `port_range` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.InternetservicedefinitionEntryPortRange> PortRanges;
        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255).
        /// </summary>
        public readonly int? Protocol;
        /// <summary>
        /// Entry sequence number.
        /// </summary>
        public readonly int? SeqNum;

        [OutputConstructor]
        private InternetservicedefinitionEntry(
            int? categoryId,

            string? name,

            int? port,

            ImmutableArray<Outputs.InternetservicedefinitionEntryPortRange> portRanges,

            int? protocol,

            int? seqNum)
        {
            CategoryId = categoryId;
            Name = name;
            Port = port;
            PortRanges = portRanges;
            Protocol = protocol;
            SeqNum = seqNum;
        }
    }
}