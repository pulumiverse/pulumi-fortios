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
    public sealed class ProfileFilter
    {
        /// <summary>
        /// VideoFilter action. Valid values: `allow`, `monitor`, `block`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// FortiGuard category ID.
        /// </summary>
        public readonly string? Category;
        /// <summary>
        /// Channel ID.
        /// </summary>
        public readonly string? Channel;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string? Comment;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Video filter keyword ID.
        /// </summary>
        public readonly int? Keyword;
        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Log;
        /// <summary>
        /// Filter type. Valid values: `category`, `channel`, `title`, `description`.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ProfileFilter(
            string? action,

            string? category,

            string? channel,

            string? comment,

            int? id,

            int? keyword,

            string? log,

            string? type)
        {
            Action = action;
            Category = category;
            Channel = channel;
            Comment = comment;
            Id = id;
            Keyword = keyword;
            Log = log;
            Type = type;
        }
    }
}
