// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Web.Outputs
{

    [OutputType]
    public sealed class ProfileAntiphishCustomPattern
    {
        /// <summary>
        /// Category that the pattern matches. Valid values: `username`, `password`.
        /// </summary>
        public readonly string? Category;
        /// <summary>
        /// Target pattern.
        /// </summary>
        public readonly string? Pattern;
        /// <summary>
        /// Pattern will be treated either as a regex pattern or literal string. Valid values: `regex`, `literal`.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ProfileAntiphishCustomPattern(
            string? category,

            string? pattern,

            string? type)
        {
            Category = category;
            Pattern = pattern;
            Type = type;
        }
    }
}
