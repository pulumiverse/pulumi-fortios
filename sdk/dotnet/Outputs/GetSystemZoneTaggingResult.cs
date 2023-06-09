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
    public sealed class GetSystemZoneTaggingResult
    {
        /// <summary>
        /// Tag category.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Specify the name of the desired system zone.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Tags. The structure of `tags` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemZoneTaggingTagResult> Tags;

        [OutputConstructor]
        private GetSystemZoneTaggingResult(
            string category,

            string name,

            ImmutableArray<Outputs.GetSystemZoneTaggingTagResult> tags)
        {
            Category = category;
            Name = name;
            Tags = tags;
        }
    }
}
