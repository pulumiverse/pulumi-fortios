// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class FirewallProxyaddressTagging
    {
        /// <summary>
        /// Tag category.
        /// </summary>
        public readonly string? Category;
        /// <summary>
        /// Tagging entry name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Tags. The structure of `tags` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallProxyaddressTaggingTag> Tags;

        [OutputConstructor]
        private FirewallProxyaddressTagging(
            string? category,

            string? name,

            ImmutableArray<Outputs.FirewallProxyaddressTaggingTag> tags)
        {
            Category = category;
            Name = name;
            Tags = tags;
        }
    }
}
