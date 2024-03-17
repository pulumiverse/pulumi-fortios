// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class OspfSummaryAddress
    {
        /// <summary>
        /// Enable/disable advertise status. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Advertise;
        /// <summary>
        /// Summary address entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Prefix.
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// Tag value.
        /// </summary>
        public readonly int? Tag;

        [OutputConstructor]
        private OspfSummaryAddress(
            string? advertise,

            int? id,

            string? prefix,

            int? tag)
        {
            Advertise = advertise;
            Id = id;
            Prefix = prefix;
            Tag = tag;
        }
    }
}
