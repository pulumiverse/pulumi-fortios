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
    public sealed class Prefixlist6Rule
    {
        /// <summary>
        /// Permit or deny packets that match this rule. Valid values: `permit`, `deny`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Flags.
        /// </summary>
        public readonly int? Flags;
        /// <summary>
        /// Minimum prefix length to be matched (0 - 128).
        /// </summary>
        public readonly int? Ge;
        /// <summary>
        /// Rule ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Maximum prefix length to be matched (0 - 128).
        /// </summary>
        public readonly int? Le;
        /// <summary>
        /// IPv6 prefix to define regular filter criteria, such as "any" or subnets.
        /// </summary>
        public readonly string? Prefix6;

        [OutputConstructor]
        private Prefixlist6Rule(
            string? action,

            int? flags,

            int? ge,

            int? id,

            int? le,

            string? prefix6)
        {
            Action = action;
            Flags = flags;
            Ge = ge;
            Id = id;
            Le = le;
            Prefix6 = prefix6;
        }
    }
}
