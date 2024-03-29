// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Video.Outputs
{

    [OutputType]
    public sealed class KeywordWord
    {
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string? Comment;
        /// <summary>
        /// Name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Pattern type. Valid values: `wildcard`, `regex`.
        /// </summary>
        public readonly string? PatternType;
        /// <summary>
        /// Enable(consider)/disable(ignore) this keyword. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private KeywordWord(
            string? comment,

            string? name,

            string? patternType,

            string? status)
        {
            Comment = comment;
            Name = name;
            PatternType = patternType;
            Status = status;
        }
    }
}
