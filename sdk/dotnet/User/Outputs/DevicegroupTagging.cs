// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User.Outputs
{

    [OutputType]
    public sealed class DevicegroupTagging
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
        public readonly ImmutableArray<Outputs.DevicegroupTaggingTag> Tags;

        [OutputConstructor]
        private DevicegroupTagging(
            string? category,

            string? name,

            ImmutableArray<Outputs.DevicegroupTaggingTag> tags)
        {
            Category = category;
            Name = name;
            Tags = tags;
        }
    }
}
