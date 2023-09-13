// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Systemdhcp.Outputs
{

    [OutputType]
    public sealed class ServerExcludeRange
    {
        /// <summary>
        /// End of IP range.
        /// </summary>
        public readonly string? EndIp;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Start of IP range.
        /// </summary>
        public readonly string? StartIp;
        /// <summary>
        /// Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? VciMatch;
        /// <summary>
        /// One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServerExcludeRangeVciString> VciStrings;

        [OutputConstructor]
        private ServerExcludeRange(
            string? endIp,

            int? id,

            string? startIp,

            string? vciMatch,

            ImmutableArray<Outputs.ServerExcludeRangeVciString> vciStrings)
        {
            EndIp = endIp;
            Id = id;
            StartIp = startIp;
            VciMatch = vciMatch;
            VciStrings = vciStrings;
        }
    }
}
